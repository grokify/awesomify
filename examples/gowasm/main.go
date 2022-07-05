package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/grokify/awesomify/schema"
)

func main() {
	lines := []string{}
	entries := EntriesGoRuntimes()
	for _, entry := range entries {
		md, err := entry.Markdown()
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, "- "+md)
	}
	fmt.Println(strings.Join(lines, "\n"))
}

func EntriesGoRuntimes() schema.Entries {
	entries := schema.Entries{
		{
			Name:        "WasmEdge-go",
			URL:         "https://github.com/second-state/WasmEdge-go",
			Description: "Go WebAssembly run time for [WasmEdge](https://github.com/WasmEdge/WasmEdge) (C++)",
		},
		{
			Name:        "wasmer-go",
			URL:         "https://github.com/wasmerio/wasmer-go",
			Description: "Go WebAssembly run time for [Wasmer](https://github.com/wasmerio/wasmer) (Rust)",
			Shields: schema.Shields{
				{
					ImageName: schema.SchemaNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/wasmerio/wasmer-go/wasmer",
					ImageURL:  "https://pkg.go.dev/badge/github.com/wasmerio/wasmer-go/wasmer",
				},
				{
					ImageName: schema.ShieldNameBuildStatus,
					ImageURL:  `https://github.com/wasmerio/wasmer-go/workflows/Build%20and%20Test/badge.svg`,
					LinkURL:   "https://github.com/wasmerio/wasmer-go/actions?query=workflow%3A%22Build+and+Test%22",
				},
			},
		},
		{
			Name:        "wasmtime-go",
			URL:         "https://github.com/bytecodealliance/wasmtime-go",
			Description: "Go WebAssembly run time for [Wastime](https://github.com/bytecodealliance/wasmtime) (Rust)",
			Shields: schema.Shields{
				{
					ImageName: schema.SchemaNameAPIDocs,
					LinkURL:   "https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go",
					ImageURL:  "https://pkg.go.dev/badge/github.com/bytecodealliance/wasmtime-go",
				},
				{
					ImageName: schema.ShieldNameBuildStatus,
					LinkURL:   "https://github.com/bytecodealliance/wasmtime-go/actions?query=workflow%3ACI",
					ImageURL:  "https://github.com/bytecodealliance/wasmtime-go/workflows/CI/badge.svg",
				},
			},
		},
		{
			Name:        "wazero",
			URL:         "https://github.com/tetratelabs/wazero",
			Description: "Zero dependency Go WebAssembly runtime (no CGO)",
			Shields: schema.Shields{
				{
					ImageName: schema.SchemaNameAPIDocs,
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
	for i, entry := range entries {
		entry.CategoryPath = []string{"Go Runtimes"}
		entries[i] = entry
	}
	return entries
}
