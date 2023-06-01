package pojo

// SysBaseMenus 结构体
type SysBaseMenus struct {
	BasePojo
	ActiveName  string `json:"activeName" form:"activeName" gorm:"column:active_name;comment:附加属性;size:191;"`
	CloseTab    *bool  `json:"closeTab" form:"closeTab" gorm:"column:close_tab;comment:附加属性;"`
	Component   string `json:"component" form:"component" gorm:"column:component;comment:对应前端文件路径;size:191;"`
	DefaultMenu *bool  `json:"defaultMenu" form:"defaultMenu" gorm:"column:default_menu;comment:附加属性;"`
	Hidden      *bool  `json:"hidden" form:"hidden" gorm:"column:hidden;comment:是否在列表隐藏;"`
	Icon        string `json:"icon" form:"icon" gorm:"column:icon;comment:附加属性;size:191;"`
	KeepAlive   *bool  `json:"keepAlive" form:"keepAlive" gorm:"column:keep_alive;comment:附加属性;"`
	MenuLevel   *int   `json:"menuLevel" form:"menuLevel" gorm:"column:menu_level;comment:;size:20;"`
	Name        string `json:"name" form:"name" gorm:"column:name;comment:路由name;size:191;"`
	ParentId    string `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父菜单ID;size:191;"`
	Path        string `json:"path" form:"path" gorm:"column:path;comment:路由path;size:191;"`
	Sort        *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记;size:19;"`
	Title       string `json:"title" form:"title" gorm:"column:title;comment:附加属性;size:191;"`
}

// TableName SysBaseMenus 表名
func (SysBaseMenus) TableName() string {
	return "sys_base_menus"
}
