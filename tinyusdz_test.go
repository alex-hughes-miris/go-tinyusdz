package tinyusdz

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// TestStageLifecycle verifies that we can create and free stages without leaks or crashes.
// It specifically tests the "Double Free" safety mechanism in our wrapper.
func TestStageLifecycle(t *testing.T) {
	t.Log("Creating Stage...")
	stage := NewStage()
	if stage == nil {
		t.Fatal("Expected NewStage to return a pointer, got nil")
	}

	t.Log("Freeing Stage...")
	stage.Free()

	t.Log("Attempting Double Free (should be safe)...")
	// This should not panic or segfault because our wrapper checks for nil
	stage.Free()
}

// TestLoadNonExistentFile checks the error boundary.
// It ensures that C++ exceptions or error codes are correctly translated to Go errors
// without crashing the runtime.
func TestLoadNonExistentFile(t *testing.T) {
	stage, err := LoadUSDZ("this_file_definitely_does_not_exist.usdz")
	if err == nil {
		stage.Free()
		t.Fatal("Expected error when loading non-existent file, got nil")
	}

	t.Logf("Got expected error: %v", err)

	// Ensure we can't use a nil stage returned from a failed load
	if stage != nil {
		t.Error("Expected nil stage on failure")
	}
}

// TestConcurrency checks if the library is thread-safe for basic allocation.
// Many C libraries use static globals that race when initialized in parallel.
func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	count := 50 // Run enough iterations to likely trigger a race if one exists

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Stress test allocation and deallocation
			s := NewStage()
			// Simulate some tiny work duration
			time.Sleep(time.Millisecond) 
			s.Free()
		}(i)
	}

	wg.Wait()
	t.Logf("Successfully created and destroyed %d stages concurrently", count)
}

// TestValueMarshaling verifies the low-level C -> Go type conversion.
// We construct C values manually using the generated raw API and check 
// if ToInterface() correctly converts them to Go types.
func TestValueMarshaling(t *testing.T) {
	// 1. Test Int Marshaling
	// We use the raw generated API here (assuming ValueNewInt exists from generator)
	cInt := ValueNewInt(42)
	defer ValueFree(cInt)

	val := &Value{ptr: cInt}
	goInterface := val.ToInterface()

	if goInt, ok := goInterface.(int); !ok {
		t.Errorf("Expected int type, got %T", goInterface)
	} else if goInt != 42 {
		t.Errorf("Expected 42, got %d", goInt)
	}

	// 2. Test Float Marshaling
	cFloat := ValueNewFloat(3.14)
	defer ValueFree(cFloat)

	valFloat := &Value{ptr: cFloat}
	goInterfaceFloat := valFloat.ToInterface()

	// Note: C float maps to Go float32
	if goFloat, ok := goInterfaceFloat.(float32); !ok {
		t.Errorf("Expected float32 type, got %T", goInterfaceFloat)
	} else if fmt.Sprintf("%.2f", goFloat) != "3.14" {
		t.Errorf("Expected 3.14, got %f", goFloat)
	}
}

// TestGCStress ensures that the Go Garbage Collector doesn't collect our 
// wrapper struct while the C pointer is still in use.
func TestGCStress(t *testing.T) {
	for i := 0; i < 1000; i++ {
		func() {
			s := NewStage()
			// We don't defer Free() here to test finalizers if you implemented them,
			// or simply to churn memory.
			// Ideally, we explicitly Free() to be safe.
			defer s.Free()
			
			// Trigger aggressive GC to see if it eats our pointer prematurely
			runtime.GC()
			
			// Accessing the pointer after GC should be safe
			_ = s.String() 
		}()
	}
}
