package tinyusdz

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../lib -lctinyusd -lstdc++ -lm
#include "c-tinyusd.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)


// Free wraps c_tinyusd_free
func Free(ptr *C.void, ) int {
    
        
    

    ret := C.c_tinyusd_free(
        
            ptr,
        
    )
    
        return int(ret)
    
}

// TokenSize wraps c_tinyusd_token_size
func TokenSize(tok *C.c_tinyusd_token_t, ) uint64 {
    
        
    

    ret := C.c_tinyusd_token_size(
        
            tok,
        
    )
    
        return uint64(ret)
    
}

// TokenFree wraps c_tinyusd_token_free
func TokenFree(tok *C.c_tinyusd_token_t, ) int {
    
        
    

    ret := C.c_tinyusd_token_free(
        
            tok,
        
    )
    
        return int(ret)
    
}

// TokenVectorFree wraps c_tinyusd_token_vector_free
func TokenVectorFree(sv *C.c_tinyusd_token_vector_t, ) int {
    
        
    

    ret := C.c_tinyusd_token_vector_free(
        
            sv,
        
    )
    
        return int(ret)
    
}

// TokenVectorSize wraps c_tinyusd_token_vector_size
func TokenVectorSize(sv *C.c_tinyusd_token_vector_t, ) uint64 {
    
        
    

    ret := C.c_tinyusd_token_vector_size(
        
            sv,
        
    )
    
        return uint64(ret)
    
}

// TokenVectorClear wraps c_tinyusd_token_vector_clear
func TokenVectorClear(sv *C.c_tinyusd_token_vector_t, ) int {
    
        
    

    ret := C.c_tinyusd_token_vector_clear(
        
            sv,
        
    )
    
        return int(ret)
    
}

// TokenVectorResize wraps c_tinyusd_token_vector_resize
func TokenVectorResize(sv *C.c_tinyusd_token_vector_t, n C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_token_vector_resize(
        
            sv,
        
            n,
        
    )
    
        return int(ret)
    
}

// TokenVectorReplace wraps c_tinyusd_token_vector_replace
func TokenVectorReplace(sv *C.c_tinyusd_token_vector_t, idx C.const size_t, str string, ) int {
    
        
    
        
    
        
            c_str := C.CString(str)
            defer C.free(unsafe.Pointer(c_str))
        
    

    ret := C.c_tinyusd_token_vector_replace(
        
            sv,
        
            idx,
        
            c_str,
        
    )
    
        return int(ret)
    
}

// StringSize wraps c_tinyusd_string_size
func StringSize(s *C.c_tinyusd_string_t, ) uint64 {
    
        
    

    ret := C.c_tinyusd_string_size(
        
            s,
        
    )
    
        return uint64(ret)
    
}

// StringReplace wraps c_tinyusd_string_replace
func StringReplace(s *C.c_tinyusd_string_t, str string, ) int {
    
        
    
        
            c_str := C.CString(str)
            defer C.free(unsafe.Pointer(c_str))
        
    

    ret := C.c_tinyusd_string_replace(
        
            s,
        
            c_str,
        
    )
    
        return int(ret)
    
}

// StringFree wraps c_tinyusd_string_free
func StringFree(s *C.c_tinyusd_string_t, ) int {
    
        
    

    ret := C.c_tinyusd_string_free(
        
            s,
        
    )
    
        return int(ret)
    
}

// StringVectorNewEmpty wraps c_tinyusd_string_vector_new_empty
func StringVectorNewEmpty(sv *C.c_tinyusd_string_vector, n C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_string_vector_new_empty(
        
            sv,
        
            n,
        
    )
    
        return int(ret)
    
}

// StringVectorNew wraps c_tinyusd_string_vector_new
func StringVectorNew(sv *C.c_tinyusd_string_vector, n C.const size_t, *strs string, ) int {
    
        
    
        
    
        
            c_*strs := C.CString(*strs)
            defer C.free(unsafe.Pointer(c_*strs))
        
    

    ret := C.c_tinyusd_string_vector_new(
        
            sv,
        
            n,
        
            c_*strs,
        
    )
    
        return int(ret)
    
}

// StringVectorSize wraps c_tinyusd_string_vector_size
func StringVectorSize(sv *C.c_tinyusd_string_vector, ) uint64 {
    
        
    

    ret := C.c_tinyusd_string_vector_size(
        
            sv,
        
    )
    
        return uint64(ret)
    
}

// StringVectorClear wraps c_tinyusd_string_vector_clear
func StringVectorClear(sv *C.c_tinyusd_string_vector, ) int {
    
        
    

    ret := C.c_tinyusd_string_vector_clear(
        
            sv,
        
    )
    
        return int(ret)
    
}

// StringVectorResize wraps c_tinyusd_string_vector_resize
func StringVectorResize(sv *C.c_tinyusd_string_vector, n C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_string_vector_resize(
        
            sv,
        
            n,
        
    )
    
        return int(ret)
    
}

// StringVectorReplace wraps c_tinyusd_string_vector_replace
func StringVectorReplace(sv *C.c_tinyusd_string_vector, idx C.const size_t, str string, ) int {
    
        
    
        
    
        
            c_str := C.CString(str)
            defer C.free(unsafe.Pointer(c_str))
        
    

    ret := C.c_tinyusd_string_vector_replace(
        
            sv,
        
            idx,
        
            c_str,
        
    )
    
        return int(ret)
    
}

// StringVectorFree wraps c_tinyusd_string_vector_free
func StringVectorFree(sv *C.c_tinyusd_string_vector, ) int {
    
        
    

    ret := C.c_tinyusd_string_vector_free(
        
            sv,
        
    )
    
        return int(ret)
    
}

// ValueTypeIsNumeric wraps c_tinyusd_value_type_is_numeric
func ValueTypeIsNumeric(value_type C.CTinyUSDValueType, ) uint32 {
    
        
    

    ret := C.c_tinyusd_value_type_is_numeric(
        
            value_type,
        
    )
    
        return uint32(ret)
    
}

// ValueTypeSizeof wraps c_tinyusd_value_type_sizeof
func ValueTypeSizeof(value_type C.CTinyUSDValueType, ) uint32 {
    
        
    

    ret := C.c_tinyusd_value_type_sizeof(
        
            value_type,
        
    )
    
        return uint32(ret)
    
}

// ValueTypeComponents wraps c_tinyusd_value_type_components
func ValueTypeComponents(value_type C.CTinyUSDValueType, ) uint32 {
    
        
    

    ret := C.c_tinyusd_value_type_components(
        
            value_type,
        
    )
    
        return uint32(ret)
    
}

// ValueType wraps c_tinyusd_value_type
func ValueType(value *C.CTinyUSDValue, ) int {
    
        
    

    ret := C.c_tinyusd_value_type(
        
            value,
        
    )
    
        return int(ret)
    
}

// ValueFree wraps c_tinyusd_value_free
func ValueFree(val *C.CTinyUSDValue, ) int {
    
        
    

    ret := C.c_tinyusd_value_free(
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueToString wraps c_tinyusd_value_to_string
func ValueToString(val *C.CTinyUSDValue, str *C.c_tinyusd_string_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_to_string(
        
            val,
        
            str,
        
    )
    
        return int(ret)
    
}

// ValueFree wraps c_tinyusd_value_free
func ValueFree(val *C.CTinyUSDValue, ) int {
    
        
    

    ret := C.c_tinyusd_value_free(
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueIsType wraps c_tinyusd_value_is_type
func ValueIsType(value *C.CTinyUSDValue, value_type C.CTinyUSDValueType, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_is_type(
        
            value,
        
            value_type,
        
    )
    
        return int(ret)
    
}

// ValueAsInt wraps c_tinyusd_value_as_int
func ValueAsInt(value *C.CTinyUSDValue, val *C.int, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_int(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsInt2 wraps c_tinyusd_value_as_int2
func ValueAsInt2(value *C.CTinyUSDValue, val *C.c_tinyusd_int2_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_int2(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsInt3 wraps c_tinyusd_value_as_int3
func ValueAsInt3(value *C.CTinyUSDValue, val *C.c_tinyusd_int3_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_int3(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsInt4 wraps c_tinyusd_value_as_int4
func ValueAsInt4(value *C.CTinyUSDValue, val *C.c_tinyusd_int4_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_int4(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsFloat wraps c_tinyusd_value_as_float
func ValueAsFloat(value *C.CTinyUSDValue, val *C.float, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_float(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsFloat2 wraps c_tinyusd_value_as_float2
func ValueAsFloat2(value *C.CTinyUSDValue, val *C.c_tinyusd_float2_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_float2(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsFloat3 wraps c_tinyusd_value_as_float3
func ValueAsFloat3(value *C.CTinyUSDValue, val *C.c_tinyusd_float3_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_float3(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// ValueAsFloat4 wraps c_tinyusd_value_as_float4
func ValueAsFloat4(value *C.CTinyUSDValue, val *C.c_tinyusd_float4_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_value_as_float4(
        
            value,
        
            val,
        
    )
    
        return int(ret)
    
}

// RelationsipFree wraps c_tinyusd_relationsip_free
func RelationsipFree(rel *C.CTinyUSDRelationship, ) int {
    
        
    

    ret := C.c_tinyusd_relationsip_free(
        
            rel,
        
    )
    
        return int(ret)
    
}

// RelationsipIsBlocked wraps c_tinyusd_relationsip_is_blocked
func RelationsipIsBlocked(rel *C.CTinyUSDRelationship, ) int {
    
        
    

    ret := C.c_tinyusd_relationsip_is_blocked(
        
            rel,
        
    )
    
        return int(ret)
    
}

// RelationsipNumTargetPaths wraps c_tinyusd_relationsip_num_targetPaths
func RelationsipNumTargetPaths(rel *C.CTinyUSDRelationship, ) uint32 {
    
        
    

    ret := C.c_tinyusd_relationsip_num_targetPaths(
        
            rel,
        
    )
    
        return uint32(ret)
    
}

// RelationsipGetTargetPath wraps c_tinyusd_relationsip_get_targetPath
func RelationsipGetTargetPath(rel *C.CTinyUSDRelationship, i uint32, targetPath *C.CTinyUSDPath, ) int {
    
        
    
        
    
        
    

    ret := C.c_tinyusd_relationsip_get_targetPath(
        
            rel,
        
            C.uint32_t(i),
        
            targetPath,
        
    )
    
        return int(ret)
    
}

// AttributeConnectionSet wraps c_tinyusd_attribute_connection_set
func AttributeConnectionSet(attr *C.CTinyUSDAttribute, connectionPath *C.CTinyUSDPath, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_attribute_connection_set(
        
            attr,
        
            connectionPath,
        
    )
    
        return int(ret)
    
}

// AttributeConnectionsSet wraps c_tinyusd_attribute_connections_set
func AttributeConnectionsSet(attr *C.CTinyUSDAttribute, n uint32, connectionPaths *C.CTinyUSDPath, ) int {
    
        
    
        
    
        
    

    ret := C.c_tinyusd_attribute_connections_set(
        
            attr,
        
            C.uint32_t(n),
        
            connectionPaths,
        
    )
    
        return int(ret)
    
}

// AttributeMetaSet wraps c_tinyusd_attribute_meta_set
func AttributeMetaSet(attr *C.CTinyUSDAttribute, meta_name string, value *C.CTinyUSDValue, ) int {
    
        
    
        
            c_meta_name := C.CString(meta_name)
            defer C.free(unsafe.Pointer(c_meta_name))
        
    
        
    

    ret := C.c_tinyusd_attribute_meta_set(
        
            attr,
        
            c_meta_name,
        
            value,
        
    )
    
        return int(ret)
    
}

// AttributeMetaGet wraps c_tinyusd_attribute_meta_get
func AttributeMetaGet(attr *C.CTinyUSDAttribute, meta_name string, *value *C.CTinyUSDValue, ) int {
    
        
    
        
            c_meta_name := C.CString(meta_name)
            defer C.free(unsafe.Pointer(c_meta_name))
        
    
        
    

    ret := C.c_tinyusd_attribute_meta_get(
        
            attr,
        
            c_meta_name,
        
            *value,
        
    )
    
        return int(ret)
    
}

// AttributeConnectionGet wraps c_tinyusd_attribute_connection_get
func AttributeConnectionGet(attr *C.CTinyUSDAttribute, n uint32, connectionPaths *C.CTinyUSDPath, ) int {
    
        
    
        
    
        
    

    ret := C.c_tinyusd_attribute_connection_get(
        
            attr,
        
            C.uint32_t(n),
        
            connectionPaths,
        
    )
    
        return int(ret)
    
}

// PropertyNew wraps c_tinyusd_property_new
func PropertyNew(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_new(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertyNewAttribute wraps c_tinyusd_property_new_attribute
func PropertyNewAttribute(prop *C.CTinyUSDProperty, attr *C.CTinyUSDAttribute, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_property_new_attribute(
        
            prop,
        
            attr,
        
    )
    
        return int(ret)
    
}

// PropertyNewRelationship wraps c_tinyusd_property_new_relationship
func PropertyNewRelationship(prop *C.CTinyUSDProperty, rel *C.CTinyUSDRelationship, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_property_new_relationship(
        
            prop,
        
            rel,
        
    )
    
        return int(ret)
    
}

// PropertyFree wraps c_tinyusd_property_free
func PropertyFree(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_free(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertySetAttribute wraps c_tinyusd_property_set_attribute
func PropertySetAttribute(prop *C.CTinyUSDProperty, attr *C.CTinyUSDAttribute, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_property_set_attribute(
        
            prop,
        
            attr,
        
    )
    
        return int(ret)
    
}

// PropertySetRelationship wraps c_tinyusd_property_set_relationship
func PropertySetRelationship(prop *C.CTinyUSDProperty, rel *C.CTinyUSDRelationship, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_property_set_relationship(
        
            prop,
        
            rel,
        
    )
    
        return int(ret)
    
}

// PropertyIsAttribute wraps c_tinyusd_property_is_attribute
func PropertyIsAttribute(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_is_attribute(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertyIsAttributeConnection wraps c_tinyusd_property_is_attribute_connection
func PropertyIsAttributeConnection(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_is_attribute_connection(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertyIsRelationship wraps c_tinyusd_property_is_relationship
func PropertyIsRelationship(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_is_relationship(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertyIsCustom wraps c_tinyusd_property_is_custom
func PropertyIsCustom(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_is_custom(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PropertyIsVarying wraps c_tinyusd_property_is_varying
func PropertyIsVarying(prop *C.CTinyUSDProperty, ) int {
    
        
    

    ret := C.c_tinyusd_property_is_varying(
        
            prop,
        
    )
    
        return int(ret)
    
}

// PrimToString wraps c_tinyusd_prim_to_string
func PrimToString(prim *C.CTinyUSDPrim, str *C.c_tinyusd_string_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_prim_to_string(
        
            prim,
        
            str,
        
    )
    
        return int(ret)
    
}

// PrimFree wraps c_tinyusd_prim_free
func PrimFree(prim *C.CTinyUSDPrim, ) int {
    
        
    

    ret := C.c_tinyusd_prim_free(
        
            prim,
        
    )
    
        return int(ret)
    
}

// PrimGetPropertyNames wraps c_tinyusd_prim_get_property_names
func PrimGetPropertyNames(prim *C.CTinyUSDPrim, prop_names_out *C.c_tinyusd_token_vector_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_prim_get_property_names(
        
            prim,
        
            prop_names_out,
        
    )
    
        return int(ret)
    
}

// PrimPropertyGet wraps c_tinyusd_prim_property_get
func PrimPropertyGet(prim *C.CTinyUSDPrim, prop_name string, prop *C.CTinyUSDProperty, ) int {
    
        
    
        
            c_prop_name := C.CString(prop_name)
            defer C.free(unsafe.Pointer(c_prop_name))
        
    
        
    

    ret := C.c_tinyusd_prim_property_get(
        
            prim,
        
            c_prop_name,
        
            prop,
        
    )
    
        return int(ret)
    
}

// PrimPropertyAdd wraps c_tinyusd_prim_property_add
func PrimPropertyAdd(prim *C.CTinyUSDPrim, prop_name string, prop *C.CTinyUSDProperty, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
            c_prop_name := C.CString(prop_name)
            defer C.free(unsafe.Pointer(c_prop_name))
        
    
        
    
        
    

    ret := C.c_tinyusd_prim_property_add(
        
            prim,
        
            c_prop_name,
        
            prop,
        
            err,
        
    )
    
        return int(ret)
    
}

// PrimPropertyDel wraps c_tinyusd_prim_property_del
func PrimPropertyDel(prim *C.CTinyUSDPrim, prop_name string, ) int {
    
        
    
        
            c_prop_name := C.CString(prop_name)
            defer C.free(unsafe.Pointer(c_prop_name))
        
    

    ret := C.c_tinyusd_prim_property_del(
        
            prim,
        
            c_prop_name,
        
    )
    
        return int(ret)
    
}

// PrimMetaSet wraps c_tinyusd_prim_meta_set
func PrimMetaSet(prim *C.CTinyUSDPrim, meta_name string, value *C.CTinyUSDValue, ) int {
    
        
    
        
            c_meta_name := C.CString(meta_name)
            defer C.free(unsafe.Pointer(c_meta_name))
        
    
        
    

    ret := C.c_tinyusd_prim_meta_set(
        
            prim,
        
            c_meta_name,
        
            value,
        
    )
    
        return int(ret)
    
}

// PrimMetaGet wraps c_tinyusd_prim_meta_get
func PrimMetaGet(prim *C.CTinyUSDPrim, meta_name string, *value *C.CTinyUSDValue, ) int {
    
        
    
        
            c_meta_name := C.CString(meta_name)
            defer C.free(unsafe.Pointer(c_meta_name))
        
    
        
    

    ret := C.c_tinyusd_prim_meta_get(
        
            prim,
        
            c_meta_name,
        
            *value,
        
    )
    
        return int(ret)
    
}

// PrimMetaAuthored wraps c_tinyusd_prim_meta_authored
func PrimMetaAuthored(prim *C.CTinyUSDPrim, meta_name string, ) int {
    
        
    
        
            c_meta_name := C.CString(meta_name)
            defer C.free(unsafe.Pointer(c_meta_name))
        
    

    ret := C.c_tinyusd_prim_meta_authored(
        
            prim,
        
            c_meta_name,
        
    )
    
        return int(ret)
    
}

// PrimAppendChild wraps c_tinyusd_prim_append_child
func PrimAppendChild(prim *C.CTinyUSDPrim, child *C.CTinyUSDPrim, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_prim_append_child(
        
            prim,
        
            child,
        
    )
    
        return int(ret)
    
}

// PrimAppendChildMove wraps c_tinyusd_prim_append_child_move
func PrimAppendChildMove(prim *C.CTinyUSDPrim, child *C.CTinyUSDPrim, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_prim_append_child_move(
        
            prim,
        
            child,
        
    )
    
        return int(ret)
    
}

// PrimDelChild wraps c_tinyusd_prim_del_child
func PrimDelChild(prim *C.CTinyUSDPrim, child_index uint64, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_prim_del_child(
        
            prim,
        
            C.size_t(child_index),
        
    )
    
        return int(ret)
    
}

// PrimNumChildren wraps c_tinyusd_prim_num_children
func PrimNumChildren(prim *C.CTinyUSDPrim, ) uint64 {
    
        
    

    ret := C.c_tinyusd_prim_num_children(
        
            prim,
        
    )
    
        return uint64(ret)
    
}

// PrimGetChild wraps c_tinyusd_prim_get_child
func PrimGetChild(prim *C.CTinyUSDPrim, child_index uint64, *child_prim *C.CTinyUSDPrim, ) int {
    
        
    
        
    
        
    

    ret := C.c_tinyusd_prim_get_child(
        
            prim,
        
            C.size_t(child_index),
        
            *child_prim,
        
    )
    
        return int(ret)
    
}

// StageToString wraps c_tinyusd_stage_to_string
func StageToString(stage *C.CTinyUSDStage, str *C.c_tinyusd_string_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_stage_to_string(
        
            stage,
        
            str,
        
    )
    
        return int(ret)
    
}

// StageFree wraps c_tinyusd_stage_free
func StageFree(stage *C.CTinyUSDStage, ) int {
    
        
    

    ret := C.c_tinyusd_stage_free(
        
            stage,
        
    )
    
        return int(ret)
    
}

// StageTraverse wraps c_tinyusd_stage_traverse
func StageTraverse(stage *C.CTinyUSDStage, callback_fun C.CTinyUSDTraversalFunction, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
    
        
    

    ret := C.c_tinyusd_stage_traverse(
        
            stage,
        
            callback_fun,
        
            err,
        
    )
    
        return int(ret)
    
}

// IsUsdFile wraps c_tinyusd_is_usd_file
func IsUsdFile(filename string, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    

    ret := C.c_tinyusd_is_usd_file(
        
            c_filename,
        
    )
    
        return int(ret)
    
}

// IsUsdaFile wraps c_tinyusd_is_usda_file
func IsUsdaFile(filename string, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    

    ret := C.c_tinyusd_is_usda_file(
        
            c_filename,
        
    )
    
        return int(ret)
    
}

// IsUsdcFile wraps c_tinyusd_is_usdc_file
func IsUsdcFile(filename string, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    

    ret := C.c_tinyusd_is_usdc_file(
        
            c_filename,
        
    )
    
        return int(ret)
    
}

// IsUsdzFile wraps c_tinyusd_is_usdz_file
func IsUsdzFile(filename string, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    

    ret := C.c_tinyusd_is_usdz_file(
        
            c_filename,
        
    )
    
        return int(ret)
    
}

// IsUsdMemory wraps c_tinyusd_is_usd_memory
func IsUsdMemory(addr *C.uint8_t, nbytes C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_is_usd_memory(
        
            addr,
        
            nbytes,
        
    )
    
        return int(ret)
    
}

// IsUsdaMemory wraps c_tinyusd_is_usda_memory
func IsUsdaMemory(addr *C.uint8_t, nbytes C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_is_usda_memory(
        
            addr,
        
            nbytes,
        
    )
    
        return int(ret)
    
}

// IsUsdcMemory wraps c_tinyusd_is_usdc_memory
func IsUsdcMemory(addr *C.uint8_t, nbytes C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_is_usdc_memory(
        
            addr,
        
            nbytes,
        
    )
    
        return int(ret)
    
}

// IsUsdzMemory wraps c_tinyusd_is_usdz_memory
func IsUsdzMemory(addr *C.uint8_t, nbytes C.const size_t, ) int {
    
        
    
        
    

    ret := C.c_tinyusd_is_usdz_memory(
        
            addr,
        
            nbytes,
        
    )
    
        return int(ret)
    
}

// LoadUsdFromFile wraps c_tinyusd_load_usd_from_file
func LoadUsdFromFile(filename string, stage *C.CTinyUSDStage, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usd_from_file(
        
            c_filename,
        
            stage,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdaFromFile wraps c_tinyusd_load_usda_from_file
func LoadUsdaFromFile(filename string, stage *C.CTinyUSDStage, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usda_from_file(
        
            c_filename,
        
            stage,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdcFromFile wraps c_tinyusd_load_usdc_from_file
func LoadUsdcFromFile(filename string, stage *C.CTinyUSDStage, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usdc_from_file(
        
            c_filename,
        
            stage,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdzFromFile wraps c_tinyusd_load_usdz_from_file
func LoadUsdzFromFile(filename string, stage *C.CTinyUSDStage, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
            c_filename := C.CString(filename)
            defer C.free(unsafe.Pointer(c_filename))
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usdz_from_file(
        
            c_filename,
        
            stage,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdFromMemory wraps c_tinyusd_load_usd_from_memory
func LoadUsdFromMemory(addr *C.uint8_t, nbytes C.const size_t, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usd_from_memory(
        
            addr,
        
            nbytes,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdaFromMemory wraps c_tinyusd_load_usda_from_memory
func LoadUsdaFromMemory(addr *C.uint8_t, nbytes C.const size_t, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usda_from_memory(
        
            addr,
        
            nbytes,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdcFromMemory wraps c_tinyusd_load_usdc_from_memory
func LoadUsdcFromMemory(addr *C.uint8_t, nbytes C.const size_t, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usdc_from_memory(
        
            addr,
        
            nbytes,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

// LoadUsdzFromMemory wraps c_tinyusd_load_usdz_from_memory
func LoadUsdzFromMemory(addr *C.uint8_t, nbytes C.const size_t, warn *C.c_tinyusd_string_t, err *C.c_tinyusd_string_t, ) int {
    
        
    
        
    
        
    
        
    

    ret := C.c_tinyusd_load_usdz_from_memory(
        
            addr,
        
            nbytes,
        
            warn,
        
            err,
        
    )
    
        return int(ret)
    
}

