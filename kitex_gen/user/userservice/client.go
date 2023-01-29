// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	user "github.com/TremblingV5/DouTok/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error)
	GetUserById(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kUserServiceClient) Login(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kUserServiceClient) GetUserById(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserById(ctx, Req)
}
