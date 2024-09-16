package main

import (
	"path"

	"google.golang.org/protobuf/compiler/protogen"
)

// Service data regarding the service to be implemented.
type Service struct {
	// FileGoImportPath used for services.
	//
	// Note: you do not need to wrap this with quotations marks as this is handled by
	// default due to the protogen pkg's behaviour.
	FileGoImportPath protogen.GoImportPath
	// ConnectGoServicePath connect go import path.
	ConnectGoServicePath protogen.GoImportPath
	// FileGoPkgName go package for the file.
	FileGoPkgName string
	// ServiceName
	ServiceName string
	// ServiceFullName full service name e.g foo.bar.service.
	ServiceFullName string
	// Methods the methods for the service.
	Methods []*protogen.Method
	// Service the protogen Service.
	Service *protogen.Service
	// File to which this method belongs.
	File *protogen.File
	// Extra additional info passed in to the generation.
	Extra map[string]any
}

// connectPath used to get the connect path for a service file.
func connectPath(file *protogen.File) protogen.GoImportPath {
	connectFileName := file.GoPackageName + "connect"
	importP := protogen.GoImportPath(path.Join(
		string(file.GoImportPath),
		string(connectFileName),
	))
	return importP
}
