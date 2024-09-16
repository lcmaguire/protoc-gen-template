package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the config used for this type.
type Config struct {
	Generations []*Generation `yaml:"generations"`
}

// Generation for this plugin.
type Generation struct {
	// OutPath is where the output of this goes.
	OutPath string `yaml:"outPath"`
	// TemplatePath to where the template is.
	TemplatePath string `yaml:"templatePath"`
	// Type the type of generation to perform.
	Type GenerationType `yaml:"type"`
	// Extra additional info you wish to parse in to the generation.
	Extra map[string]any `yaml:"extra"`
}

// GenerationType the type of generation to be performed.
type GenerationType string

const (
	// MethodGenerationType will generate for all methods included in protoc invocation.
	MethodGenerationType GenerationType = "method"
	// ServiceType will generate for all services included in protoc invocation.
	ServiceType GenerationType = "service"
	// FileType will generate for each file included in protoc invocation.
	FileType GenerationType = "file"
	// AllType will generate one file for all files included in protoc invocation.
	AllType GenerationType = "all"
)

func loadConfig(path string) (*Config, error) {
	bites, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	out := &Config{}
	err = yaml.Unmarshal(bites, out)
	return out, err
}
