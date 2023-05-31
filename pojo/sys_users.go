package pojo

// SysUsers 结构体
type SysUsers struct {
	BasePojo
	ActiveColor string `json:"activeColor" form:"activeColor" gorm:"column:active_color;comment:活跃颜色;size:191;"`
	AuthorityId *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;size:20;"`
	BaseColor   string `json:"baseColor" form:"baseColor" gorm:"column:base_color;comment:基础颜色;size:191;"`
	Email       string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`
	Enable      *int   `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否被冻结 1正常 2冻结;size:19;"`
	HeaderImg   string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`
	NickName    string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`
	Password    string `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:191;"`
	Phone       string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`
	SideMode    string `json:"sideMode" form:"sideMode" gorm:"column:side_mode;comment:用户侧边主题;size:191;"`
	Username    string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`
	Uuid        string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:191;"`
}

// TableName SysUsers 表名
func (SysUsers) TableName() string {
	return "sys_users"
}
