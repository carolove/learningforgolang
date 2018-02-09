// Code generated by protoc-gen-auxo 0.1, DO NOT EDIT.
// source: msg.proto

package proto

import (
	"context"

	"github.com/cuigh/auxo/net/rpc"
)

var (
	helloService = &helloServiceClient{rpc.LazyClient{Name: ""}}
	testService  = &testServiceClient{rpc.LazyClient{Name: ""}}
	demoService  = &demoServiceClient{rpc.LazyClient{Name: ""}}
)

// HelloService comment
type HelloService interface {
	// Hello method comment
	Hello(context.Context, *HelloRequest) (*HelloResponse, error) // abc
}

func GetHelloService() HelloService {
	return helloService
}

type helloServiceClient struct {
	rpc.LazyClient
}

func (s *helloServiceClient) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(HelloResponse)
	err = c.Call(ctx, "HelloService", "Hello", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TestService comment
type TestService interface {
	Hello1(context.Context, *HelloRequest) (*HelloResponse, error) // Hello1 method comment
	Hello2(context.Context, *HelloRequest) (*HelloResponse, error)
}

func GetTestService() TestService {
	return testService
}

type testServiceClient struct {
	rpc.LazyClient
}

func (s *testServiceClient) Hello1(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(HelloResponse)
	err = c.Call(ctx, "TestService", "Hello1", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *testServiceClient) Hello2(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	c, err := s.Try()
	if err != nil {
		return nil, err
	}

	resp := new(HelloResponse)
	err = c.Call(ctx, "TestService", "Hello2", []interface{}{req}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Block comment attached to
// DemoService.
type DemoService interface {
}

func GetDemoService() DemoService {
	return demoService
}

type demoServiceClient struct {
	rpc.LazyClient
}