package tinyusdz

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -Llib -lctinyusd -lstdc++ -lm
#include "include/c-tinyusd.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Value wraps CTinyUSDValue
type Value struct {
	ptr *C.CTinyUSDValue
}

// Prim wraps CTinyUSDPrim
type Prim struct {
	ptr *C.CTinyUSDPrim
}

// Type returns the Prim type (e.g., "Xform", "Mesh")
func (p *Prim) Type() string {
	cStr := C.c_tinyusd_prim_type(p.ptr)
	return C.GoString(cStr)
}

// ElementName returns the name of the Prim (e.g., "world", "mesh_0")
func (p *Prim) ElementName() string {
	cStr := C.c_tinyusd_prim_element_name(p.ptr)
	return C.GoString(cStr)
}

// GetPropertyNames returns a list of property names for this Prim
func (p *Prim) GetPropertyNames() ([]string, error) {
	cTokVec := C.c_tinyusd_token_vector_new_empty()
	defer C.c_tinyusd_token_vector_free(cTokVec)

	if C.c_tinyusd_prim_get_property_names(p.ptr, cTokVec) == 0 {
		return nil, fmt.Errorf("failed to get property names")
	}

	n := C.c_tinyusd_token_vector_size(cTokVec)
	result := make([]string, n)
	for i := 0; i < int(n); i++ {
		cStr := C.c_tinyusd_token_vector_str(cTokVec, C.size_t(i))
		result[i] = C.GoString(cStr)
	}
	return result, nil
}

// NumChildren returns the number of child prims
func (p *Prim) NumChildren() int {
	return int(C.c_tinyusd_prim_num_children(p.ptr))
}

// Children returns a slice of all child Prims
func (p *Prim) Children() ([]*Prim, error) {
	count := p.NumChildren()
	children := make([]*Prim, 0, count)

	for i := 0; i < count; i++ {
		var childPtr *C.CTinyUSDPrim
		
		// Note: The C API returns a pointer to a pointer in the last argument
		// We pass the address of our 'childPtr' variable
		ret := C.c_tinyusd_prim_get_child(p.ptr, C.uint64_t(i), &childPtr)
		if ret == 0 {
			return nil, fmt.Errorf("failed to get child at index %d", i)
		}
		
		// Important: The C documentation says "child's content is just a pointer...
		// so do not call Prim deleter". We treat this as a borrowed reference.
		children = append(children, &Prim{ptr: childPtr})
	}
	return children, nil
}

// ToInterface converts the underlying C Value to a Go interface{}.
func (v *Value) ToInterface() interface{} {
	if v.ptr == nil {
		return nil
	}

	cType := C.c_tinyusd_value_type(v.ptr)

	switch cType {
	case C.C_TINYUSD_VALUE_BOOL:
		var out C.int
		if C.c_tinyusd_value_as_int(v.ptr, &out) != 0 {
			return out != 0
		}
	case C.C_TINYUSD_VALUE_INT:
		var out C.int
		if C.c_tinyusd_value_as_int(v.ptr, &out) != 0 {
			return int(out)
		}
	case C.C_TINYUSD_VALUE_FLOAT:
		var out C.float
		if C.c_tinyusd_value_as_float(v.ptr, &out) != 0 {
			return float32(out)
		}
	case C.C_TINYUSD_VALUE_DOUBLE:
		var out C.float
		if C.c_tinyusd_value_as_float(v.ptr, &out) != 0 {
			return float64(out)
		}
	case C.C_TINYUSD_VALUE_FLOAT3:
		var out C.c_tinyusd_float3_t
		if C.c_tinyusd_value_as_float3(v.ptr, &out) != 0 {
			return [3]float32{float32(out.x), float32(out.y), float32(out.z)}
		}
	// Add other types as needed...
	}
	
	return nil
}

