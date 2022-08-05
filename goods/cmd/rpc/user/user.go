// Code generated by goctl. DO NOT EDIT!
// Source: goods.proto

package user

import (
	"context"

	"github.com/zhovdawei/gozero-micro-service/goods/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GoodsAttrsResp          = pb.GoodsAttrsResp
	GoodsAttrsVO            = pb.GoodsAttrsVO
	GoodsDetailResp         = pb.GoodsDetailResp
	GoodsResp               = pb.GoodsResp
	GoodsSpecVO             = pb.GoodsSpecVO
	GoodsSpecsResp          = pb.GoodsSpecsResp
	GoodsVO                 = pb.GoodsVO
	QueryGoodsAttrsReq      = pb.QueryGoodsAttrsReq
	QueryGoodsDetailReq     = pb.QueryGoodsDetailReq
	QueryGoodsListReq       = pb.QueryGoodsListReq
	QueryGoodsSpecByIdReq   = pb.QueryGoodsSpecByIdReq
	QueryGoodsSpecBySpecReq = pb.QueryGoodsSpecBySpecReq
	QueryGoodsSpecsReq      = pb.QueryGoodsSpecsReq

	User interface {
		QueryGoodsList(ctx context.Context, in *QueryGoodsListReq, opts ...grpc.CallOption) (*GoodsResp, error)
		QueryGoodsAttrs(ctx context.Context, in *QueryGoodsAttrsReq, opts ...grpc.CallOption) (*GoodsAttrsResp, error)
		QueryGoodsDetail(ctx context.Context, in *QueryGoodsDetailReq, opts ...grpc.CallOption) (*GoodsDetailResp, error)
		QueryGoodsSpecById(ctx context.Context, in *QueryGoodsSpecByIdReq, opts ...grpc.CallOption) (*GoodsSpecVO, error)
		QueryGoodsSpecs(ctx context.Context, in *QueryGoodsSpecsReq, opts ...grpc.CallOption) (*GoodsSpecsResp, error)
		QueryGoodsSpecBySpec(ctx context.Context, in *QueryGoodsSpecBySpecReq, opts ...grpc.CallOption) (*GoodsSpecVO, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) QueryGoodsList(ctx context.Context, in *QueryGoodsListReq, opts ...grpc.CallOption) (*GoodsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsList(ctx, in, opts...)
}

func (m *defaultUser) QueryGoodsAttrs(ctx context.Context, in *QueryGoodsAttrsReq, opts ...grpc.CallOption) (*GoodsAttrsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsAttrs(ctx, in, opts...)
}

func (m *defaultUser) QueryGoodsDetail(ctx context.Context, in *QueryGoodsDetailReq, opts ...grpc.CallOption) (*GoodsDetailResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsDetail(ctx, in, opts...)
}

func (m *defaultUser) QueryGoodsSpecById(ctx context.Context, in *QueryGoodsSpecByIdReq, opts ...grpc.CallOption) (*GoodsSpecVO, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsSpecById(ctx, in, opts...)
}

func (m *defaultUser) QueryGoodsSpecs(ctx context.Context, in *QueryGoodsSpecsReq, opts ...grpc.CallOption) (*GoodsSpecsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsSpecs(ctx, in, opts...)
}

func (m *defaultUser) QueryGoodsSpecBySpec(ctx context.Context, in *QueryGoodsSpecBySpecReq, opts ...grpc.CallOption) (*GoodsSpecVO, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.QueryGoodsSpecBySpec(ctx, in, opts...)
}
