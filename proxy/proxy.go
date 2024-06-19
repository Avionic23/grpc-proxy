// Copyright 2021 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package proxy

import (
	"google.golang.org/grpc"
)

// NewProxy sets up a simple proxy that forwards all requests to dst.
func NewProxy(dst *grpc.ClientConn, opts ...grpc.ServerOption) *grpc.Server {
	opts = append(opts, DefaultProxyOpt(dst))
	// Set up the proxy server and then serve from it like in step one.
	return grpc.NewServer(opts...)
}

// DefaultProxyOpt returns an grpc.UnknownServiceHandler with a DefaultDirector.
func DefaultProxyOpt(cc *grpc.ClientConn) grpc.ServerOption {
	return grpc.UnknownServiceHandler(TransparentHandler(DefaultDirector(cc), DefaultErrorHandler()))
}

// DefaultDirector returns a very simple forwarding StreamDirector that forwards all
// calls.
func DefaultDirector(cc *grpc.ClientConn) StreamDirector {
	return func() string {
		return cc.Target()
	}
}

func DefaultErrorHandler() ErrorHandler {
	return func(addr string) {}
}
