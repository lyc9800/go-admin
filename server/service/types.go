package service

// 接受登陆参数的结构体
type LoginPassWordRequest struct {
	UserName string `json:"userName"` // 登陆用户名
	Password string `json:"password"` // 登陆密码
}

// 登陆成功后的token结构体
type LoginPasswordReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// 获取管理员列表参数结构体
type GetUserListRequest struct {
	*QueryRequest
}

// 关键字和分页信息结构体
type QueryRequest struct {
	Page    int    `json:"pageIndex" form:"pageIndex"`
	Size    int    `json:"pageSize" form:"pageSize"`
	Keyword string `json:"keyword" form:"keyword"`
}

// 返回管理员信息结构体
type GetUserListReply struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Remarks   string `json:"remarks"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// 接收添加管理员表单数据的结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
	RoleId   uint   `json:"roleId"`
}

// 返回用户详细信息结构体
type GetUserDetailReply struct {
	ID uint `json:"id"`
	AddUserRequest
}

// 接收更新用户信息结构体
type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}

// 角色列表查询
type GetRoleListRequest struct {
	*QueryRequest
}

// 角色列表返回
type GetRoleListReply struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Sort      int    `json:"sort"`
	IsAdmin   int8   `json:"is_admin"`
	Remarks   string `json:"remarks"`
}

// 返回角色数据的结构体
type AllRoleListReply struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// 添加角色参数结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Sort    int64  `json:"sort"`
	IsAdmin int8   `json:"is_admin"`
	Remarks string `json:"remarks"`
	MenuId  []uint `json:"menuId"`
}

// 返回角色信息结构体
type GetRoleDetailReply struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// 更新角色参数结构体
type UpdateRoleRequest struct {
	ID uint `json:"id"`
	AddRoleRequest
}

// 获取菜单列表返回参数结构体
type MenuReplay struct {
	ID            uint          `json:"id"`
	ParentId      uint          `json:"parent_id"`
	Name          string        `json:"name"`
	WebIcon       string        `json:"web_icon"`
	Sort          int           `json:"sort"`
	Path          string        `json:"path"`
	Level         int8          `json:"level"`
	ComponentName string        `json:"component_name"`
	SubMenus      []*MenuReplay `json:"sub_menus"`
}

// 获取所有菜单结构体
type AllMenu struct {
	ID            uint   `json:"id"`
	ParentId      uint   `json:"parent_id"`
	Name          string `json:"name"`
	WebIcon       string `json:"web_icon"`
	Sort          int    `json:"sort"`
	Path          string `json:"path"`
	Level         int8   `json:"level"`
	ComponentName string `json:"component_name"`
}

// 增加菜单参数结构体
type AddMenuRequest struct {
	ParentId      uint   `json:"parent_id"`      // 父级菜单ID
	Name          string `json:"name"`           // 菜单名称
	WebIcon       string `json:"web_icon"`       // 图标
	Sort          int    `json:"sort"`           // 排序
	Path          string `json:"path"`           // 路径
	Level         uint   `json:"level"`          // 菜单等级 {0: 目录, 1: 菜单, 3: 按钮}
	ComponentName string `json:"component_name"` // 组件路径
}

// 修改菜单参数结构体
type UpdateMenuRequest struct {
	ID uint `json:"id"`
	AddMenuRequest
}
