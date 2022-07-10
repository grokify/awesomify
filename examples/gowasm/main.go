package main

import (
	"fmt"
	"log"

	"github.com/grokify/awesomely/schema"
)

func main() {
	awe := AwesomeGoWasm()
	md, err := awe.Markdown()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(md)
}

const (
	CategoryGoGuestExamples     = "Go Guest Examples"
	CategoryTinyGoGuestExamples = "TinyGo Guest Examples"
	CategoryLibraries           = "Libraries"
	CategoryRuntimes            = "Runtimes"
)

func AwesomeGoWasm() schema.Awesome {
	return schema.Awesome{
		Categories: schema.Categories{
			{
				Path: []string{CategoryGoGuestExamples},
			},
			{
				Path: []string{CategoryTinyGoGuestExamples},
			},
			{
				Path: []string{CategoryLibraries},
			},
			{
				Path: []string{CategoryRuntimes},
			},
		},
		Entries: EntriesGoRuntimes(),
	}
}

func EntriesGoRuntimes() schema.Entries {
	entries := schema.Entries{
		{
			Category:    schema.Category{Path: []string{CategoryGoGuestExamples}},
			Name:        "go-wasm-examples",
			URL:         "https://github.com/danieljoos/go-wasm-examples",
			Description: "Some small examples of using Go and WebAssembly.",
		},
		{
			Category:    schema.Category{Path: []string{CategoryGoGuestExamples}},
			Name:        "Go WebAssembly Tutorial - Building a Calculator Tutorial",
			URL:         "https://tutorialedge.net/golang/go-webassembly-tutorial",
			Description: `build "building a really simple calculator to give us an idea as to how we can write functions that can be exposed to the frontend, evaluate DOM elements and subsequently update any DOM elements with the results from any functions we call."`,
		},

		{
			Category:    schema.Category{Path: []string{CategoryLibraries}},
			Name:        "Wat2Wasm",
			URL:         "https://github.com/bytecodealliance/wasmtime-go/blob/main/wat2wasm.go",
			Description: "Wat2Wasm converts the text format of WebAssembly to the binary format.",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go#Wat2Wasm",
					ImageURL:  "https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go",
				},
				{
					ImageName: schema.BadgeNameBuildStatus,
					LinkURL:   "https://github.com/bytecodealliance/wasmtime-go/actions?query=workflow%3ACI",
					ImageURL:  "https://github.com/bytecodealliance/wasmtime-go/workflows/CI/badge.svg",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryLibraries}},
			Name:        "vugu",
			URL:         "https://github.com/vugu/vugu",
			Description: "Vugu is an experimental library for web UIs written in Go and targeting webassembly.",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/vugu/vugu",
					ImageURL:  "https://pkg.go.dev/badge/github.com/vugu/vugu",
				},
				{
					ImageName: schema.BadgeNameBuildStatus,
					LinkURL:   "https://travis-ci.org/github/vugu/vugu",
					ImageURL:  "https://api.travis-ci.org/vugu/vugu.svg?branch=master",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryLibraries}},
			Name:        "wasm-fetch",
			URL:         "https://github.com/marwan-at-work/wasm-fetch",
			Description: "A go-wasm library that wraps the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API). This is useful since TinyGo does not support `net/http`.",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/marwan-at-work/wasm-fetch",
					ImageURL:  "https://pkg.go.dev/badge/github.com/marwan-at-work/wasm-fetch",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryRuntimes}},
			Name:        "WasmEdge-go",
			URL:         "https://github.com/second-state/WasmEdge-go",
			Description: "Go WebAssembly run time for [WasmEdge](https://github.com/WasmEdge/WasmEdge) (C++)",

			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/second-state/WasmEdge-go",
					ImageURL:  "https://pkg.go.dev/badge/second-state/WasmEdge-go",
				},
				{
					ImageName: schema.BadgeNameBuildStatus,
					ImageURL:  `https://github.com/wasmerio/wasmer-go/workflows/Build%20and%20Test/badge.svg`,
					LinkURL:   "https://github.com/wasmerio/wasmer-go/actions?query=workflow%3A%22Build+and+Test%22",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryRuntimes}},
			Name:        "wasmer-go",
			URL:         "https://github.com/wasmerio/wasmer-go",
			Description: "Go WebAssembly run time for [Wasmer](https://github.com/wasmerio/wasmer) (Rust)",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/wasmerio/wasmer-go/wasmer",
					ImageURL:  "https://pkg.go.dev/badge/github.com/wasmerio/wasmer-go/wasmer",
				},
				{
					ImageName: schema.BadgeNameBuildStatus,
					ImageURL:  `https://github.com/wasmerio/wasmer-go/workflows/Build%20and%20Test/badge.svg`,
					LinkURL:   "https://github.com/wasmerio/wasmer-go/actions?query=workflow%3A%22Build+and+Test%22",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryRuntimes}},
			Name:        "wasmtime-go",
			URL:         "https://github.com/bytecodealliance/wasmtime-go",
			Description: "Go WebAssembly run time for [Wastime](https://github.com/bytecodealliance/wasmtime) (Rust)",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go",
					ImageURL:  "https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go",
				},
				{
					ImageName: schema.BadgeNameBuildStatus,
					LinkURL:   "https://github.com/bytecodealliance/wasmtime-go/actions?query=workflow%3ACI",
					ImageURL:  "https://github.com/bytecodealliance/wasmtime-go/workflows/CI/badge.svg",
				},
			},
		},
		{
			Category:    schema.Category{Path: []string{CategoryRuntimes}},
			Name:        "wazero",
			URL:         "https://github.com/tetratelabs/wazero",
			Description: "Zero dependency Go WebAssembly runtime (no CGO)",
			Badges: schema.Badges{
				{
					ImageName: schema.BadgeNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/tetratelabs/wazero",
					ImageURL:  "https://pkg.go.dev/badge//github.com/tetratelabs/wazero",
				},
				{
					ImageName: "WebAssembly Core Specification Test",
					ImageURL:  "https://github.com/tetratelabs/wazero/actions/workflows/spectest.yaml/badge.svg",
					LinkURL:   "https://github.com/tetratelabs/wazero/actions/workflows/spectest.yaml",
				},
			},
		},
	}
	return entries
}
