package model

import "github.com/gogf/gf/v2/os/gtime"

// PermissionCreateUpdateBase 创建/修改内容基类
type PermissionCreateUpdateBase struct {
	Name string
	Path string
}

// PermissionCreateInput 创建内容
type PermissionCreateInput struct {
	PermissionCreateUpdateBase
}

// PermissionCreateOutput 创建内容返回结果
type PermissionCreateOutput struct {
	PermissionId uint `json:"permission_id"`
}

// PermissionUpdateInput 修改内容
type PermissionUpdateInput struct {
	PermissionCreateUpdateBase
	Id uint
}

// PermissionGetListInput 获取内容列表
type PermissionGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PermissionGetListOutput 查询列表结果
type PermissionGetListOutput struct {
	List  []PermissionGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

//type PermissionGetListOutputItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	Name      string      `json:"name"`
//	Path      string      `json:"path"`
//	Component string      `json:"component"`  // 前端文件路径
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}

type PermissionSearchOutputItem struct {
	PermissionGetListOutputItem
}

type PermissionGetListOutputItem struct {
	ID        int64                          `json:"id"` // 自增ID
	ParentID  int64                          `json:"parent_id"`
	Path      string                         `json:"path"`
	Component string                         `json:"component"`
	Redirect  string                         `json:"redirect"`
	Name      string                         `json:"name"`
	Meta      *Meta                          `json:"meta"`
	Children  []*PermissionGetListOutputItem `json:"children"`
	CreatedAt *gtime.Time                    `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time                    `json:"updated_at"` // 修改时间
}

type Meta struct {
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	AlwaysShow bool   `json:"alwaysShow"`
	NoCache    bool   `json:"noCache"`
}

type PermissionGetListOutItem struct {
	ID         int64       `json:"id"` // 自增ID
	ParentID   int64       `json:"parent_id"`
	Path       string      `json:"path"`
	Component  string      `json:"component"`
	Redirect   string      `json:"redirect"`
	Name       string      `json:"name"`
	Title      string      `json:"title"`
	Icon       string      `json:"icon"`
	AlwaysShow uint        `json:"alwaysShow"`
	NoCache    uint        `json:"noCache"`
	CreatedAt  *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"` // 修改时间
}
