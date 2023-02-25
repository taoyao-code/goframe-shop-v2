package role

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop-v2/internal/dao"
	"goframe-shop-v2/internal/model"
	"goframe-shop-v2/internal/model/entity"
	"goframe-shop-v2/internal/service"
	"golang.org/x/net/context"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	// 删除内容
	_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		dao.PermissionInfo.Columns().Id: id,
	}).Unscoped().Delete()
	return err
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err
}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	//listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PermissionInfo
	if err := m.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	var a []model.PermissionGetListOutItem
	//不指定item的键名用：Scan
	if err := m.Scan(&a); err != nil {
		return out, err
	}
	menus, err := GetItemMenus(a)
	//menus, err := GetItemMenus(out.List)
	out.List = menus

	return
}

func GetItemMenus(Data []model.PermissionGetListOutItem) ([]model.PermissionGetListOutputItem, error) {
	// 构建 map 保存每个 Menu 对象
	menuMap := make(map[int64]*model.PermissionGetListOutputItem)
	var tree []*model.PermissionGetListOutputItem

	for _, item := range Data {
		var treeItem model.PermissionGetListOutputItem
		// 主数据
		treeItem.ID = item.ID
		treeItem.ParentID = item.ParentID
		treeItem.Path = item.Path
		treeItem.Component = item.Component
		treeItem.Redirect = item.Redirect
		treeItem.Name = item.Name

		fmt.Println(item.AlwaysShow)
		treeItem.Meta = &model.Meta{
			Title:      item.Title,
			Icon:       item.Icon,
			AlwaysShow: item.AlwaysShow == 1,
			NoCache:    item.NoCache == 1,
		}
		treeItem.Children = []*model.PermissionGetListOutputItem{}

		// 根节点收集
		if item.ParentID == 0 {
			tree = append(tree, &treeItem)
		} else {
			// 子节点收集
			menuMap[item.ParentID].Children = append(menuMap[item.ParentID].Children, &treeItem)
		}
		// 把子节点映射到map表
		menuMap[item.ID] = &treeItem
	}

	jsonRes, _ := json.Marshal(tree)
	var p []model.PermissionGetListOutputItem
	err := json.Unmarshal(jsonRes, &p)
	if err != nil {
		return nil, err
	}
	return p, nil

}
