# **go-tinyusdz**

Idiomatic Go bindings for [tinyusdz](https://github.com/lighttransport/tinyusdz), a lightweight and dependency-free Universal Scene Description (USD) library.

## **Overview**

go-tinyusdz allows you to load, inspect, and traverse .usdz, .usdc, and .usda files directly in Go. It bridges the C++ tinyusdz library using Cgo, providing a memory-safe and garbage-collected API while maintaining high performance.

**Features:**

* **Safe API**: Automatic memory management (no manual free() required for Stages).
* **Type Safety**: Converts USD types (tokens, floats, vectors) to native Go interfaces.
* **Concurrency**: Thread-safe access for scene traversal.
* **Auto-Generated**: Low-level bindings are automatically generated from the C header.

## **Prerequisites**

Because this library binds to C++ code, you must have a C++ compiler and CMake installed on your system.

* **Linux**: sudo apt install build-essential cmake
* **macOS**: xcode-select \--install && brew install cmake
* **Windows**: Visual Studio with C++ support or MinGW.

## **Installation**

go get [github.com/alex-hughes-miris/go-tinyusdz](https://github.com/alex-hughes-miris/go-tinyusdz)

*Note: You may need to build the C library first if you are checking this out from source. See the [Building](https://www.google.com/search?q=%23building-from-source) section.*

## **Usage**

### **1. Loading a USDZ File**

The LoadUSDZ function handles memory allocation and error string conversion for you.
```go
package main

import (
	"fmt"
	"log"
	"github.com/alex-hughes-miris/go-tinyusdz"
)

func main() {
	// Load the stage (automatically freed by GC, or use defer .Free() for determinism)
	stage, err := tinyusdz.LoadUSDZ("assets/teapot.usdz")
	if err != nil {
		log.Fatalf("Failed to load: %v", err)
	}
	defer stage.Free()

	fmt.Println("Successfully loaded USDZ file!")
	fmt.Println(stage.String()) // Prints the C++ string representation
}
```
### **2. Traversing the Scene Graph**

Since C++ callbacks cannot directly capture Go closures, this library uses a global mutex-protected gateway to allow safe, idiomatic Go traversal.
```go
// Get all root Prims
roots, err := stage.GetRootPrims()
if err != nil {
    log.Fatal(err)
}

for _, root := range roots {
    fmt.Printf("Root Prim: %s (Type: %s)\n", root.ElementName(), root.Type())

    // Recursive helper to print children
    printChildren(root, 1)
}

func printChildren(p *tinyusdz.Prim, depth int) {
    children, _ := p.Children()
    for _, child := range children {
        indent := "" // build indentation string...
        fmt.Printf("%s-> %s\n", indent, child.ElementName())
        printChildren(child, depth+1)
    }
}
```
### **3. Reading Attributes**

Use the ToInterface() helper to convert raw C values into Go types (int, float32, \[3\]float32, etc.).
```go
// Assuming you have a Prim...
props, _ := prim.GetPropertyNames()
for _, name := range props {
    // Logic to get property/attribute (Not fully implemented in example wrapper yet)
    // val := prim.GetAttribute(name).Value()
    // fmt.Println(name, val.ToInterface())
}
```
## **Building from Source**

If you are developing this library or need to rebuild the C++ core:

1. **Clone the Repo**:
   git clone [https://github.com/alex-hughes-miris/go-tinyusdz.git](https://github.com/alex-hughes-miris/go-tinyusdz.git)
   cd go-tinyusdz

2. Fetch Dependencies:
   This repo expects the tinyusdz source code to be available or the library to be pre-compiled. A helper script or CI job typically handles this.
   *See .github/workflows/build.yml for the exact build steps.*
3. Generate Bindings:
   If you modified c-tinyusd.h, regenerate the Go wrapper:
   `go generate ./...`

4. **Run Tests**:
   ```bash
   # Ensure the compiled library is in your library path
   export LD_LIBRARY_PATH=$(pwd)/lib
   go test -v ./...
   ```

## **Architecture**

This repository is split into three layers:

1. **tinyusdz_wrapper.go (Generated)**:
   * Created by cmd/gen.
   * Contains raw Cgo functions (e.g., C.c_tinyusd_stage_new).
   * **Do not edit manually.**
2. **tinyusdz.go & types.go (Handwritten)**:
   * Wraps unsafe C pointers in Go structs.
   * Handles finalizers, error strings, and type marshaling.
3. **traversal.go (Handwritten)**:
   * Implements the Go-to-C-to-Go callback bridge for traversing the stage.

## **Contributing**

1. Fork the repository.
2. Add your changes (if modifying C code, add a patch to build/patches).
3. Run go generate if you changed the headers.
4. Submit a Pull Request.

## **License**

MIT License. See LICENSE for details.
Underlying tinyusdz library is Apache 2.0 / MIT.
