package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Method contains all info for method generation.
type Method struct {
	// MethodName is the name of the RPC being implemented.
	MethodName string
	// MethodFullName full method name.
	//
	// example.service.MethodName
	MethodFullName string
	// MethodType is the type of RPC that for this method.
	MethodType MethodType
	// FileGoPkgName go package for the file.
	FileGoPkgName string
	// ServiceName is the name of the service to which the method belongs.
	ServiceName string
	// Input is the *protogen.Message for the input of the gRPC method.
	Input *protogen.Message
	// Output is the *protogen.Message for the output of the gRPC method.
	Output *protogen.Message
	// Method *protogen.Method.
	Method *protogen.Method
	// Service to which this belongs to.
	Service *protogen.Service
	// File to which this method belongs.
	File *protogen.File
	// Extra additional info passed in to the generation.
	Extra map[string]any
}

// MethodType the type of rpc that the method is.
type MethodType string

const (
	Unary           MethodType = "unary"
	ClientStreaming MethodType = "client_streaming"
	ServerStreaming MethodType = "server_streaming"
	Bidirectional   MethodType = "bidirectional"
)

func methodType(m protoreflect.MethodDescriptor) MethodType {
	switch {
	case m.IsStreamingServer() && m.IsStreamingClient():
		return Bidirectional
	case m.IsStreamingServer():
		return ServerStreaming
	case m.IsStreamingClient():
		return ClientStreaming
	default:
		return Unary
	}
}
