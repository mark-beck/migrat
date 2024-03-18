package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	// get module location from args
	path := os.Args[1]

	// load module
	module, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading module file")
		panic(err)
	}

	// setup environment
	ctx := context.Background()

	fs := os.DirFS("/")

	buffer := new(strings.Builder)
	errBuffer := new(strings.Builder)
	moduleConfig := wazero.NewModuleConfig().WithFS(fs).WithStdout(buffer).WithStderr(errBuffer)

	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	mod, err := r.InstantiateWithConfig(ctx, module, moduleConfig)
	if err != nil {
		fmt.Println("Error instantiating module")
		panic(err)
	}

	// call the run function
	result, err := mod.ExportedFunction("run").Call(ctx, 1)
	if err != nil {
		fmt.Println("Error calling run function")
		panic(err)
	}

	fmt.Println("Result:", result)
	fmt.Println("Out:\n", buffer.String())
	fmt.Println("Err:\n", errBuffer.String())

}
