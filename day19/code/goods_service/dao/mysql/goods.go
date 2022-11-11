package mysql

import (
	"context"
	"goods_service/errno"
	"goods_service/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// dao 层用来执行数据库相关的操作

// GetGoodsByRoomId 根据roomID查直播间绑定的所有的商品 Id
func GetGoodsByRoomId(ctx context.Context, roomId int64) ([]*model.RoomGoods, error) {
	// 通过gorm去数据库中获取数据
	var data []*model.RoomGoods
	err := db.WithContext(ctx).
		Model(&model.RoomGoods{}).
		Where("room_id = ?", roomId).
		Order("weight").
		Find(&data).Error
	// 如果查询出错且不是空数据的错
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errno.ErrQueryFailed
	}
	return data, nil
}

// GetGoodsById 根据id查询商品信息
func GetGoodsById(ctx context.Context, idList []int64) ([]*model.Goods, error) {
	var data []*model.Goods
	err := db.WithContext(ctx).
		Model(&model.Goods{}).
		Where("goods_id in ?", idList). // 会按照idList顺序返回吗?
		Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: "FIELD(goods_id,?)", Vars: []interface{}{idList}, WithoutParentheses: true},
		}).
		Find(&data).Error
	if err != nil && err != gorm.ErrEmptySlice {
		return nil, errno.ErrQueryFailed
	}
	return data, nil
}

// select * from xx_goods where goods_id in (3,2,1) order by field (goods_id, 3,2,1);
