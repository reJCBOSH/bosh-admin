package dept

import (
	"strconv"

	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/global"
	"bosh-admin/model"
)

type SysDeptSvc struct{}

func NewSysDeptSvc() *SysDeptSvc {
	return &SysDeptSvc{}
}

func (svc *SysDeptSvc) GetDeptTree() ([]model.SysDept, error) {
	treeMap, err := queryDeptTreeMap()
	if err != nil {
		return nil, err
	}
	deptTree := treeMap[0]
	for i := 0; i < len(deptTree); i++ {
		err = getDeptChildrenList(&deptTree[i], treeMap)
	}
	if err != nil {
		return nil, exception.NewException("查询部门树失败", err)
	}
	return deptTree, nil
}

// queryDeptTreeMap 查询部门树map
func queryDeptTreeMap() (map[uint][]model.SysDept, error) {
	var allDept []model.SysDept
	treeMap := make(map[uint][]model.SysDept)
	err := db.GormDB().Order("display_order DESC").Find(&allDept).Error
	if err != nil {
		return nil, exception.NewException("查询部门列表失败", err)
	}
	for _, v := range allDept {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, nil
}

// getDeptChildrenList 获取子部门列表
func getDeptChildrenList(dept *model.SysDept, treeMap map[uint][]model.SysDept) error {
	var err error
	dept.Children = treeMap[dept.Id]
	for i := 0; i < len(dept.Children); i++ {
		err = getDeptChildrenList(&dept.Children[i], treeMap)
	}
	return err
}

func (svc *SysDeptSvc) GetDeptList(deptName, deptCode string, status *int, pageNo, pageSize int) ([]model.SysDept, int64, error) {
	var list []model.SysDept
	var total int64
	var err error
	query := db.GormDB().Model(&model.SysDept{})
	if deptName != "" {
		query = query.Where("dept_name LIKE ?", "%"+deptName+"%")
	}
	if deptCode != "" {
		query = query.Where("dept_code LIKE ?", "%"+deptCode+"%")
	}
	if status != nil {
		query = query.Where("status = ?", status)
	}
	if pageNo > 0 && pageSize > 0 {
		err = query.Count(&total).Error
		if err != nil {
			return nil, 0, exception.NewException("查询部门数量失败", err)
		}
		query = query.Scopes(db.PageScope(pageSize, pageNo))
	}
	err = query.Order("display_order DESC").Find(&list).Error
	if err != nil {
		return nil, 0, exception.NewException("查询部门列表失败", err)
	}
	return list, total, nil
}

func (svc *SysDeptSvc) GetDeptById(id uint) (*model.SysDept, error) {
	data, err := db.QueryById[model.SysDept](id)
	if err != nil {
		return nil, exception.NewException("查询部门失败", err)
	}
	return data, nil
}

func (svc *SysDeptSvc) AddDept(dept AddDeptReq) error {
	var count int64
	err := db.GormDB().Model(&model.SysDept{}).Where("dept_code = ?", dept.DeptCode).Count(&count).Error
	if err != nil {
		return exception.NewException("查询部门数量失败", err)
	}
	if count > 0 {
		return exception.NewException("部门编码已存在")
	}
	if dept.ParentId == 0 {
		dept.DeptPath = "0"
	} else {
		var parentDept *model.SysDept
		parentDept, err = db.QueryById[model.SysDept](dept.ParentId)
		if err != nil {
			return exception.NewException("查询父部门失败", err)
		}
		dept.DeptPath = parentDept.DeptPath + "," + strconv.Itoa(int(dept.ParentId))
	}
	err = db.Create(&dept, new(model.SysDept).TableName())
	if err != nil {
		return exception.NewException("新增部门失败", err)
	}
	return nil
}

func (svc *SysDeptSvc) EditDept(dept EditDeptReq) error {
	originDept, err := db.QueryById[model.SysDept](dept.Id)
	if err != nil {
		return exception.NewException("查询部门失败", err)
	}
	if originDept.DeptCode == global.SystemAdmin {
		return exception.NewException("系统管理员部门不允许修改")
	}
	err = db.Updates(&dept, new(model.SysDept).TableName())
	if err != nil {
		return exception.NewException("修改部门失败", err)
	}
	return nil
}

func (svc *SysDeptSvc) DelDept(id any) error {
	originDept, err := db.QueryById[model.SysDept](id)
	if err != nil {
		return exception.NewException("查询部门失败", err)
	}
	if originDept.DeptCode == global.SystemAdmin {
		return exception.NewException("系统管理员部门不允许删除")
	}
	var count int64
	err = db.GormDB().Model(&model.SysDept{}).Where("parent_id = ?", originDept.Id).Count(&count).Error
	if err != nil {
		return exception.NewException("查询子部门数量失败", err)
	}
	if count > 0 {
		return exception.NewException("存在子部门，不允许删除")
	}
	err = db.DelById[model.SysDept](id)
	if err != nil {
		return exception.NewException("删除部门失败", err)
	}
	return nil
}
