# Generator


## Usage

Put a `gommon.yml` in folder, gommon will traverse folders and generate go file based on instruction

Supported generators

- logger, generate methods to match `log.LoggableStruct` interface
- go template, render template using `text/template`
- shell, call external commands like `protoc`

````yaml
loggers:
  - struct: "*Server"
    receiver: srv
gotmpls:
    - src: abc.go.tmpl
      dst: abc.go
      data:
        foo: bar
````

NOTE

- remember to quote string with `*` in YAML, `*Foo` means reference while `"*Foo"` is a normal string


## TODO

logger

- [ ] use assert in test
- [ ] generate interface check
- [x] return error in Render, better error handling
- [x] write to buffer and then run go format like https://github.com/dyweb/Ayi/blob/master/cmd/gotmpl.go

gotmpl

- [x] replace Ayi's gotmpl

shell

- [ ] support using `sh -c`