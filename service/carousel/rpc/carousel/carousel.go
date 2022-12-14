// Code generated by goctl. DO NOT EDIT!
// Source: carousel.proto

package carousel

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Carousel interface {
		GetCarouselsForIndex(ctx context.Context, in *GetCarouselsForIndexRequest, opts ...grpc.CallOption) (*GetCarouselsForIndexResponse, error)
		AddCarousel(ctx context.Context, in *AddCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		UpdateCarousel(ctx context.Context, in *UpdateCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		DeleteCarousel(ctx context.Context, in *DeleteCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
		GetCarousel(ctx context.Context, in *GetCarouselRequest, opts ...grpc.CallOption) (*GetCarouselResponse, error)
		GetCarouselList(ctx context.Context, in *GetCarouselListRequest, opts ...grpc.CallOption) (*GetCarouselListResponse, error)
	}

	defaultCarousel struct {
		cli zrpc.Client
	}
)

func NewCarousel(cli zrpc.Client) Carousel {
	return &defaultCarousel{
		cli: cli,
	}
}

func (m *defaultCarousel) GetCarouselsForIndex(ctx context.Context, in *GetCarouselsForIndexRequest, opts ...grpc.CallOption) (*GetCarouselsForIndexResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.GetCarouselsForIndex(ctx, in, opts...)
}

func (m *defaultCarousel) AddCarousel(ctx context.Context, in *AddCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.AddCarousel(ctx, in, opts...)
}

func (m *defaultCarousel) UpdateCarousel(ctx context.Context, in *UpdateCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.UpdateCarousel(ctx, in, opts...)
}

func (m *defaultCarousel) DeleteCarousel(ctx context.Context, in *DeleteCarouselRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.DeleteCarousel(ctx, in, opts...)
}

func (m *defaultCarousel) GetCarousel(ctx context.Context, in *GetCarouselRequest, opts ...grpc.CallOption) (*GetCarouselResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.GetCarousel(ctx, in, opts...)
}

func (m *defaultCarousel) GetCarouselList(ctx context.Context, in *GetCarouselListRequest, opts ...grpc.CallOption) (*GetCarouselListResponse, error) {
	client := NewCarouselClient(m.cli.Conn())
	return client.GetCarouselList(ctx, in, opts...)
}
