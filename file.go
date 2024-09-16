package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// File message type used for generating for a file.
type File struct {
	File          *protogen.File
	Name          string
	FullName      string
	GoPackageName string
	// Extra additional info passed in to the generation.
	Extra map[string]any
}

// All will perform for all files within the protoc invocation.
type All struct {
	Files []*protogen.File
	// Extra additional info passed in to the generation.
	Extra map[string]any
}
