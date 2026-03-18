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

// 添加角色参数结构体
type AddRoleRequest struct {
	Name    string `json:"name"`
	Sort    int64  `json:"sort"`
	IsAdmin int8   `json:"is_admin"`
	Remarks string `json:"remarks"`
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
