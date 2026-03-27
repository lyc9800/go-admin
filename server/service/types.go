package service

import "server/models"

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
	RoleName  string `json:"role_name"`
}

// 接收添加管理员表单数据的结构体
type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
	RoleId   uint   `json:"roleId"`
	Avatar   string `json:"avatar"`
	Sex      string `json:"sex"`
	RoleName string `json:"role_name"`
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
	RoleName  string `json:"role_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Sort      int    `json:"sort"`
	IsAdmin   int8   `json:"is_admin"`
	Remarks   string `json:"remarks"`
}

// 返回角色数据的结构体
type AllRoleListReply struct {
	ID       uint   `json:"id"`
	RoleName string `json:"role_name"`
}

// 添加角色参数结构体
type AddRoleRequest struct {
	RoleName string `json:"role_name"`
	Sort     int64  `json:"sort"`
	IsAdmin  int8   `json:"is_admin"`
	Remarks  string `json:"remarks"`
	MenuId   []uint `json:"menuId"`
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

// 验证验证码请求参数
type VerifyCodeRequest struct {
	Email string `json:"email" form:"email" binding:"required,email"`     // 邮箱地址
	Code  string `json:"code" form:"code" binding:"required,min=6,max=6"` // 6位验证码
}

// 验证验证码返回结果
type VerifyCodeReply struct {
	Valid   bool   `json:"valid"`   // 验证是否成功
	Email   string `json:"email"`   // 验证的邮箱
	Message string `json:"message"` // 返回消息
}

// 更换绑定邮箱请求参数
type ChangeEmailRequest struct {
	Email string `json:"email" binding:"required,email"`      // 新邮箱地址
	Code  string `json:"code" binding:"required,min=6,max=6"` // 6位验证码
}

// 更换邮箱返回结果
type ChangeEmailReply struct {
	Success  bool   `json:"success"`             // 是否成功
	Message  string `json:"message"`             // 返回消息
	OldEmail string `json:"old_email,omitempty"` // 旧邮箱
	NewEmail string `json:"new_email,omitempty"` // 新邮箱
}

// 发送验证码请求参数
type SendEmailRequest struct {
	Email string `json:"email" form:"email" binding:"required,email"` // 邮箱地址
}

// 发送验证码返回结果
type SendEmailReply struct {
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 返回消息
	Email   string `json:"email"`   // 发送到的邮箱
}

// 更新用户信息请求（包含邮箱）
type UpdateUserInfoRequest struct {
	Username string `json:"username,omitempty"` // 用户名
	Phone    string `json:"phone,omitempty"`    // 手机号
	Email    string `json:"email,omitempty"`    // 邮箱
	Remarks  string `json:"remarks,omitempty"`  // 备注
	Avatar   string `json:"avatar,omitempty"`   // 头像
	Sex      string `json:"sex,omitempty"`      // 性别
	Code     string `json:"code,omitempty"`     // 验证码（更换邮箱时需要）
}

// 更新用户信息返回
type UpdateUserInfoReply struct {
	Success bool             `json:"success"`        // 是否成功
	Message string           `json:"message"`        // 返回消息
	User    GetUserListReply `json:"user,omitempty"` // 更新后的用户信息
}

// 修改密码请求参数
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required,min=6,max=20"` // 当前密码
	NewPassword string `json:"newPassword" binding:"required,min=6,max=20"` // 新密码
}

// 修改密码返回结果
type ChangePasswordReply struct {
	Success     bool   `json:"success"`       // 是否成功
	Message     string `json:"message"`       // 返回消息
	NeedReLogin bool   `json:"need_re_login"` // 是否需要重新登录
}

// 用户信息（包含角色）
type UserWithRole struct {
	models.SysUser
	RoleName string `gorm:"column:role_name" json:"role_name"`
}
