# Golang Summary

# Features

- Updates every 6 months
- Has a garbage collector
- Rich stdlib:
  - Web server lib
  - Testing tool `go test`
- 静的解析 is possible
- Finding the error without compilation is possible; this is unlikely for static-typed lang
- Single-binary
  - To run the compiled Go binary (mostly command line tool), users don't need to install the run time
  - However the binary size is big
- cross-compile
  - e.g. Can be compiled for 32 bit Windows on Mac machine
  - `GOOS` to specify the OS
  - `GOARCH` to specify the architecture (32bit 64bit)
- Simple syntax

## Purpose

- Major use
  - server side
  - network & DB fields
- Minor use
  - Some use for IoT (TinyGo)
  - GopherJS: Alt-JS for frontend. Compile Go into JS
  - WASM: WebAssembly
  - Go Mobile
  - Ebiten: Gaming

## Has a parallelism

- Go Routine
  - lightweight thread
  - `GOMAXPROCS`
- Channel
  - Pipe between go routines

# Go Versions

## Go1

- Compatibility within Go1.0 series is guaranteed

## Go2

- Not released yet
- Will be released when the large change with which 

# CUI

```sh
go build
go test

# Genrate documentation
go doc, godoc

# Code format
gofmt, goimports

# Code checker
go vet

# LSP (used for code autocompletion)
gopls
```

# MISC

## Google App Engine for Go

- Really fast to start the instance on the PaaS

## Famous service written in Golang

- gVisor
- docker
- kubernetes

## Resources

- Tour of Go
- go.dev
- gopher道場
- Go Programming Language (Alan, Kernighan)
- golang.org official document
- Go Code Review Comments (basic)
- Effective Go (detailed)
- Gopher Con

### Community

- Go Beginners
- Go Conference
- Gophers Slack

## Keywords

- gRPC
  - RPC: Remote Procedure Call
- protocol buffer
- gVisor
- WebAssembly

## Code

- `main()` is the entry point for all the Go program
- For web server, `main()` is always running
  - Most meaningful process is done by HTTP handler

## Package

- main package
  - Has `main()`
  - All he executable Go program has this
- standard package
- 3rd party package

## IDE

- GoLand (JetBrains) is the most popular
- VS Code