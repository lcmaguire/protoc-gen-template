package main

import (
	"bytes"
	"errors"
	"flag"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const pluginName = "protoc-gen-boilerplate"

func main() {
	var flags flag.FlagSet
	configPath := flags.String("config", "", "path to config")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		if configPath == nil || *configPath == "" {
			return errors.New("missing config path")
		}

		config, err := loadConfig(*configPath)
		if err != nil {
			return err
		}

		for _, step := range config.Generations {
			var err error
			switch step.Type {
			case MethodGenerationType:
				err = methodGen(gen, step)
			case ServiceType:
				err = serviceGen(gen, step)
			case FileType:
				err = fileGen(gen, step)
			case AllType:
				err = allGen(gen, step)
			default:
				continue
			}

			if err != nil {
				return err
			}
		}

		return nil
	})
}

func methodGen(gen *protogen.Plugin, step *Generation) error {
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}

		for _, service := range file.Services {
			for _, meth := range service.Methods {
				currMethod := Method{
					MethodName:     meth.GoName,
					MethodFullName: string(meth.Desc.FullName()),
					MethodType:     methodType(meth.Desc),
					ServiceName:    service.GoName,
					FileGoPkgName:  string(file.GoPackageName),
					Input:          meth.Input,
					Output:         meth.Output,
					Method:         meth,
					Service:        service,
					File:           file,
					Extra:          step.Extra,
				}

				err := generateFiles(gen, step, currMethod)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func serviceGen(gen *protogen.Plugin, step *Generation) error {
	for _, file := range gen.Files {
		// perhaps include a field within Step to determine filtering.
		if !file.Generate {
			continue
		}

		for _, service := range file.Services {
			currService := Service{
				FileGoImportPath:     file.GoImportPath,
				ConnectGoServicePath: connectPath(file),
				FileGoPkgName:        string(file.GoPackageName),
				ServiceName:          string(service.Desc.Name()),
				ServiceFullName:      string(service.Desc.FullName()),
				Methods:              service.Methods,
				Service:              service,
				File:                 file,
				Extra:                step.Extra,
			}

			err := generateFiles(gen, step, currService)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func allGen(gen *protogen.Plugin, step *Generation) error {
	a := All{
		Files: gen.Files,
		Extra: step.Extra,
	}
	return generateFiles(gen, step, a)
}

func fileGen(gen *protogen.Plugin, step *Generation) error {
	for _, file := range gen.Files {
		// perhaps include a field within Step to determine filtering.
		if !file.Generate {
			continue
		}

		currFile := File{
			File:          file,
			Name:          string(file.Desc.Name()),
			FullName:      string(file.Desc.FullName()),
			GoPackageName: string(file.GoPackageName),
			Extra:         step.Extra,
		}
		err := generateFiles(gen, step, currFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateFiles(gen *protogen.Plugin, step *Generation, data any) error {
	// parse step.OutPath as a template.
	outPathTemplate, err := template.New("").Parse(step.OutPath)
	if err != nil {
		return err
	}

	outPathBuffer := bytes.NewBuffer([]byte{})
	if err := outPathTemplate.Execute(outPathBuffer, data); err != nil {
		return err
	}

	gf := gen.NewGeneratedFile(outPathBuffer.String(), ".")

	currMethodTemplate, err := template.ParseFiles(step.TemplatePath)
	if err != nil {
		return err
	}

	currFileBuffer := bytes.NewBuffer([]byte{})
	if err := currMethodTemplate.Execute(currFileBuffer, data); err != nil {
		return err
	}
	gf.P(currFileBuffer.String())
	return nil
}
