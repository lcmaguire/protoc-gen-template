# protoc-gen-go-boilerplate

plugin to generate through the protoc compiler based upon custom go templates.

can be used for the following 

- Boilerplate go code generation akin to [protoc-gen-go-boilerplate](https://github.com/lcmaguire/protoc-gen-go-boilerplate)
- generate Markdown Documentation
- generate static curl / gRPCurl commands.

## installation

```sh

go install github.com/lcmaguire/protoc-gen-template@latest

```

## usage

Heavily recommended to use buf cli v2


within buf.gen.yaml include it as a plugin like below. 

```yaml
plugins:
  - local: protoc-gen-template # local installation executable
    out: example # where you want it output
    opt:
      - config=config.yaml # path to your config.
```

## Config
contains 

| Type         | Description                                                                                                      | Example |
|--------------|------------------------------------------------------------------------------------------------------------------|---------|
| OutPath      | the output path for the provided file, is a Go template. Will output to the proto out path + OutPath defined in  |         |
| TemplatePath | path to the template to be ran.                                                                                  |         |
| Type         | the type of generation to perform [see type](#generation-types)                                                  |         |
| Extra        | Additional values you want to access within the template                                                         |         |


## generation types

currently supports generating templates for the following scenarios.

| Type    | Description                                                        | Example |
|---------|--------------------------------------------------------------------|---------|
| Method  | Will generate for all methods included in the protoc request       |         |
| Service | Will generate for all services included in the protoc request      |         |
| File    | Will generate for each files included in the protoc request        |         |
| All     | Will generate with the template using [] Files included in request |         |


## Examples

....