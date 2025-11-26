package tinyusdz

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -Llib -lctinyusd -lstdc++ -lm
#include "include/c-tinyusd.h"
#include <stdlib.h>

// Forward declaration of the Go function exported below
extern int goTraversalProxy(const CTinyUSDPrim *prim, const CTinyUSDPath *path);
*/
import "C"
import (
	"fmt"
	"sync"
)

// Global state to bridge C callback to Go closure
var (
	traversalMu   sync.Mutex
	// This function variable holds the callback for the *current* traversal operation
	currentTraversalFunc func(*Prim) bool
)

// export goTraversalProxy
//
//export goTraversalProxy
func goTraversalProxy(prim *C.CTinyUSDPrim, path *C.CTinyUSDPath) C.int {
	if currentTraversalFunc == nil {
		return 0 // Stop traversal if no handler
	}
	
	// Create a temporary Go wrapper for the Prim
	// Note: We do not store this permanently as the C pointer might be transient
	goPrim := &Prim{ptr: prim}
	
	// Call the user-provided Go function
	shouldContinue := currentTraversalFunc(goPrim)
	
	if shouldContinue {
		return 1
	}
	return 0
}

// TraverseRoots visits all root Prims in the stage.
//
// NOTE: This function is thread-safe on the Go side (it uses a Mutex),
// effectively serializing all Traversal operations globally.
func (s *Stage) TraverseRoots(callback func(prim *Prim) bool) error {
	// 1. Lock to ensure no other goroutine overwrites 'currentTraversalFunc'
	traversalMu.Lock()
	defer traversalMu.Unlock()

	// 2. Register the callback
	currentTraversalFunc = callback
	
	// 3. Clean up callback after we are done
	defer func() { currentTraversalFunc = nil }()

	// 4. Prepare error buffer
	errObj := StringNewEmpty()
	defer StringFree(errObj)

	// 5. Call C traversal, passing our C proxy function
	// The C proxy function (goTraversalProxy) will call 'currentTraversalFunc'
	ret := C.c_tinyusd_stage_traverse(s.ptr, (C.CTinyUSDTraversalFunction)(C.goTraversalProxy), errObj)

	if ret == 0 {
		errMsg := cStringToGo(C.c_tinyusd_string_str(errObj))
		return fmt.Errorf("traversal failed: %s", errMsg)
	}

	return nil
}

// GetRootPrims is a helper that returns all root prims as a slice.
func (s *Stage) GetRootPrims() ([]*Prim, error) {
	var roots []*Prim
	
	err := s.TraverseRoots(func(p *Prim) bool {
		// We need to store the prim pointer.
		// WARNING: The validity of 'p.ptr' depends on the stage lifetime.
		// Since 'p' here is created inside the callback, we must create a copy
		// of the struct to return it.
		roots = append(roots, &Prim{ptr: p.ptr})
		return true // Continue traversal
	})
	
	if err != nil {
		return nil, err
	}
	
	return roots, nil
}
