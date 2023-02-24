// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IGoods interface {
		Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error)
		Delete(ctx context.Context, id uint) (err error)
		Update(ctx context.Context, in model.GoodsUpdateInput) error
		GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error)
		Detail(ctx context.Context, in model.GoodsDetailInput) (out model.GoodsDetailOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
