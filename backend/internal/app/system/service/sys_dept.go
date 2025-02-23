// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type ISysDept interface {
	GetList(ctx context.Context, req *system.DeptSearchReq) (list []*entity.SysDept, err error)
	GetFromCache(ctx context.Context) (list []*entity.SysDept, err error)
	Add(ctx context.Context, req *system.DeptAddReq) (err error)
	Edit(ctx context.Context, req *system.DeptEditReq) (err error)
	Delete(ctx context.Context, id uint64) (err error)
	FindSonByParentId(deptList []*entity.SysDept, deptId uint64) []*entity.SysDept
	GetListTree(pid uint64, list []*entity.SysDept) (deptTree []*model.SysDeptTreeRes)
	GetByDeptId(ctx context.Context, deptId uint64) (dept *entity.SysDept, err error)
}

var localSysDept ISysDept

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}
