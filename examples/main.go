package main

import (
	"fmt"
	"log"
	"os"

    // Adjust this import path to match your actual module name
	"github.com/alex-hughes-miris/go-tinyusdz" 
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file.usdz>")
		os.Exit(1)
	}

	file := os.Args[1]
	fmt.Printf("Attempting to load: %s\n", file)

	// 1. Load the Stage
	// Notice: No manual memory management for strings or error buffers here!
	stage, err := tinyusdz.LoadUSDZ(file)
	if err != nil {
		log.Fatalf("Error loading USDZ: %v", err)
	}
	
	// Ensure we free the stage when we are done
	defer stage.Free()

	// 2. Use the Stage
    roots, _ := stage.GetRootPrims()

    // 2. Iterate
    for _, root := range roots {
        fmt.Printf("Root Prim: %s (%s)\n", root.ElementName(), root.Type())

        // 3. Go deeper
        children, _ := root.Children()
        for _, child := range children {
             fmt.Printf("  -> Child: %s\n", child.ElementName())
        }
    }
	fmt.Println("--------------------------------------------------")
	fmt.Println("Successfully loaded!")
	fmt.Println("Stage Description:")
	fmt.Println(stage.String())
	fmt.Println("--------------------------------------------------")
}
