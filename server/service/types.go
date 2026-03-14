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
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

// 返回管理员信息结构体
type GetUserListReply struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Avatar string `json:"avatar"`
	Phone string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
