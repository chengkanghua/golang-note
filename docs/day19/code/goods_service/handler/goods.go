package handler

import (
	"context"
	"fmt"
	"goods_service/biz/goods"
	"goods_service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RPC的入口

type GoodsSrv struct {
	proto.UnimplementedGoodsServer
}

// GetGoodsByRoom 根据room_id获取直播间的商品列表
func (s *GoodsSrv) GetGoodsByRoom(ctx context.Context, req *proto.GetGoodsByRoomReq) (*proto.GoodsListResp, error) {
	// 参数处理
	fmt.Println(req.RoomId)
	if req.GetRoomId() <= 0 {
		// 无效的请求
		return nil, status.Error(codes.InvalidArgument, "请求参数有误")
	}
	// 去查询数据并封装返回的响应数据 --> 业务逻辑
	data, err := goods.GetGoodsByRoomId(ctx, req.GetRoomId())
	if err != nil {
		return nil, status.Error(codes.Internal, "内部错误")
	}
	return data, nil
}

// GetGoodsDetail 根据goods_id获取商品详情
func (s *GoodsSrv) GetGoodsDetail(context.Context, *proto.GetGoodsDetailReq) (*proto.GoodsDetail, error) {
	// 参数处理
	return &proto.GoodsDetail{}, nil
}
