// Code generated by goctl. DO NOT EDIT!
// Source: admintoken.proto

package admintoken

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Admintoken interface {
		GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
		GetExistToken(ctx context.Context, in *GetExistTokenRequest, opts ...grpc.CallOption) (*GetExistTokenResponse, error)
		DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
	}

	defaultAdmintoken struct {
		cli zrpc.Client
	}
)

func NewAdmintoken(cli zrpc.Client) Admintoken {
	return &defaultAdmintoken{
		cli: cli,
	}
}

func (m *defaultAdmintoken) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	client := NewAdmintokenClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultAdmintoken) GetExistToken(ctx context.Context, in *GetExistTokenRequest, opts ...grpc.CallOption) (*GetExistTokenResponse, error) {
	client := NewAdmintokenClient(m.cli.Conn())
	return client.GetExistToken(ctx, in, opts...)
}

func (m *defaultAdmintoken) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	client := NewAdmintokenClient(m.cli.Conn())
	return client.DeleteToken(ctx, in, opts...)
}
