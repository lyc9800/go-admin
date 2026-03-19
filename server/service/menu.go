package service

import (
	"fmt"
	"net/http"
	"server/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetMenuList 获取菜单列表
func GetMenuList(c *gin.Context) {
	Menus(c)
}

// 获取菜单列表数据
func Menus(c *gin.Context) {
	data := make([]*MenuReplay, 0)
	allMenus := make([]*AllMenu, 0)
	tx := models.GetMenuList().Debug()
	err := tx.Find(&allMenus).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 转换前
	fmt.Println("=== 转换前（扁平数据）===")
	for i, menu := range allMenus {
		fmt.Printf("[%d] ID: %d, Name: %s, ParentId: %d, ComponentName: '%s'\n",
			i, menu.ID, menu.Name, menu.ParentId, menu.ComponentName)
	}

	data = allMenuToMenuReplay(allMenus)

	// 转换后
	fmt.Println("\n=== 转换后（树形数据）===")
	printTree(data, 0)

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "菜单列表获取成功",
		"result": data,
	})
}

// 递归打印树形结构
func printTree(menus []*MenuReplay, depth int) {
	prefix := strings.Repeat("  ", depth)
	for _, menu := range menus {
		fmt.Printf("%s├─ [%d] %s (Path: %s, Component: '%s')\n",
			prefix, menu.ID, menu.Name, menu.Path, menu.ComponentName)
		if len(menu.SubMenus) > 0 {
			printTree(menu.SubMenus, depth+1)
		}
	}
}

// 创建转化函数，将数据库平面数据转化为树形结构
func allMenuToMenuReplay(allMenus []*AllMenu) []*MenuReplay {
	replay := make([]*MenuReplay, 0)

	// 1. 找出所有顶级菜单（ParentId == 0）
	for _, v := range allMenus {
		if v.ParentId == 0 {
			menuReplay := &MenuReplay{
				ID:            v.ID,
				Name:          v.Name,
				Path:          v.Path,
				WebIcon:       v.WebIcon,
				Sort:          v.Sort,
				Level:         v.Level,
				ComponentName: v.ComponentName, // 添加这行
				ParentId:      v.ParentId,
				SubMenus:      getChildrenMenu(v.ID, allMenus),
			}
			replay = append(replay, menuReplay)
		}
	}
	return replay
}

func getChildrenMenu(parentId uint, allMenus []*AllMenu) []*MenuReplay {
	children := make([]*MenuReplay, 0)
	for _, v := range allMenus {
		if v.ParentId == parentId {
			childMenu := &MenuReplay{
				ID:            v.ID,
				Name:          v.Name,
				Path:          v.Path,
				WebIcon:       v.WebIcon,
				Sort:          v.Sort,
				Level:         v.Level,
				ComponentName: v.ComponentName, // 添加这行
				ParentId:      v.ParentId,
				SubMenus:      getChildrenMenu(v.ID, allMenus),
			}
			children = append(children, childMenu)
		}
	}
	return children
}

// 增加菜单
func AddMenu(c *gin.Context) {
	in := &AddMenuRequest{}
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	// 验证参数
	err = models.DB.Create(&models.SysMenu{
		Level:         in.Level,
		Name:          in.Name,
		ParentId:      in.ParentId,
		Path:          in.Path,
		Sort:          in.Sort,
		WebIcon:       in.WebIcon,
		ComponentName: in.ComponentName,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "菜单添加成功",
		"result": nil,
	})
}
