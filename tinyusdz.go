package tinyusdz

// This directive tells Go to run the generator we just made.
// It assumes you have the header in the 'include' folder.

//go:generate go run cmd/gen/main.go

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -Llib -lctinyusd -lstdc++ -lm
#include "include/c-tinyusd.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

// Stage wraps the C tinyusd stage.
type Stage struct {
	// We keep this unexported so users can't mess up the memory
	ptr *C.CTinyUSDStage
}

// NewStage creates a new Stage and sets a finalizer to free it automatically.
func NewStage() *Stage {
	ptr := StageNew() // Call the generated raw function
	s := &Stage{ptr: ptr}

	// Optional: Use runtime.SetFinalizer to handle cleanup automatically.
	// However, explicit Free() is often safer for C libraries.
	runtime.SetFinalizer(s, func(obj *Stage) {
		obj.Free()
	})
	return s
}

// Free manually releases the memory.
func (s *Stage) Free() {
	if s.ptr != nil {
		StageFree(s.ptr)
		s.ptr = nil
	}
}

// String returns the string representation of the stage (for debugging).
func (s *Stage) String() string {
	if s.ptr == nil {
		return "<nil>"
	}
	
	// Create a C string buffer to hold the output
	cStrObj := StringNewEmpty()
	defer StringFree(cStrObj)

	StageToString(s.ptr, cStrObj)

	// Extract the raw char* and convert to Go string
	return cStringToGo(StringStr(cStrObj))
}

// LoadUSDZ is the high-level function you actually want to use.
func LoadUSDZ(filename string) (*Stage, error) {
	stage := NewStage()

	// 1. Prepare C-side string buffers for errors
	warnObj := StringNewEmpty()
	defer StringFree(warnObj)

	errObj := StringNewEmpty()
	defer StringFree(errObj)

	// 2. Call the generated raw function
	// Note: The generator handles the 'filename' string conversion for us!
	ret := LoadUsdzFromFile(filename, stage.ptr, warnObj, errObj)

	// 3. Check result (0 = failure in tinyusdz)
	if ret == 0 {
		// Convert the C error string to a Go error
		errMsg := cStringToGo(StringStr(errObj))
		stage.Free() // Clean up the failed stage
		return nil, errors.New(errMsg)
	}

	// Check for warnings (optional)
	warnMsg := cStringToGo(StringStr(warnObj))
	if warnMsg != "" {
		fmt.Printf("Warning loading USDZ: %s\n", warnMsg)
	}

	return stage, nil
}

// Helper to convert *C.char to Go string
func cStringToGo(cStr *C.char) string {
	if cStr == nil {
		return ""
	}
	return C.GoString(cStr)
}
