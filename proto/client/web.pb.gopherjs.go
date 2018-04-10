// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: proto/web.proto

/*
	Package client is a generated protocol buffer package.

	Web exposes a backend server over gRPC.

	It is generated from these files:
		proto/web.proto

	It has these top-level messages:
*/
package client

import jspb "github.com/johanbrandhorst/protobuf/jspb"

import (
	context "context"

	grpcweb "github.com/johanbrandhorst/protobuf/grpcweb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpcweb.Client

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcweb package it is being compiled against.
const _ = grpcweb.GrpcWebPackageIsVersion3

// Client API for Backend service

// Backend defines the interface exposed by the backend.
// TODO: Define functionality exposed by backend.
type BackendClient interface {
}

type backendClient struct {
	client *grpcweb.Client
}

// NewBackendClient creates a new gRPC-Web client.
func NewBackendClient(hostname string, opts ...grpcweb.DialOption) BackendClient {
	return &backendClient{
		client: grpcweb.NewClient(hostname, "web.Backend", opts...),
	}
}
