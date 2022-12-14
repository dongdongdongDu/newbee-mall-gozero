// Code generated by goctl. DO NOT EDIT!
// Source: indexconfig.proto

package indexconfig

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Indexconfig interface {
		GetConfigForIndex(ctx context.Context, in *GetConfigForIndexRequest, opts ...grpc.CallOption) (*GetConfigForIndexResponse, error)
		AddConfig(ctx context.Context, in *AddConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
		GetConfigList(ctx context.Context, in *GetConfigListRequest, opts ...grpc.CallOption) (*GetConfigListResponse, error)
	}

	defaultIndexconfig struct {
		cli zrpc.Client
	}
)

func NewIndexconfig(cli zrpc.Client) Indexconfig {
	return &defaultIndexconfig{
		cli: cli,
	}
}

func (m *defaultIndexconfig) GetConfigForIndex(ctx context.Context, in *GetConfigForIndexRequest, opts ...grpc.CallOption) (*GetConfigForIndexResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.GetConfigForIndex(ctx, in, opts...)
}

func (m *defaultIndexconfig) AddConfig(ctx context.Context, in *AddConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.AddConfig(ctx, in, opts...)
}

func (m *defaultIndexconfig) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.UpdateConfig(ctx, in, opts...)
}

func (m *defaultIndexconfig) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.DeleteConfig(ctx, in, opts...)
}

func (m *defaultIndexconfig) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.GetConfig(ctx, in, opts...)
}

func (m *defaultIndexconfig) GetConfigList(ctx context.Context, in *GetConfigListRequest, opts ...grpc.CallOption) (*GetConfigListResponse, error) {
	client := NewIndexconfigClient(m.cli.Conn())
	return client.GetConfigList(ctx, in, opts...)
}
