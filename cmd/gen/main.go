package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

// Config
const (
	HeaderFile  = "include/c-tinyusd.h"
	PackageName = "tinyusdz"
	OutputFile  = "tinyusdz_wrapper.go"
)

// TemplateData is the struct passed to the template to fix the panic
type TemplateData struct {
	PackageName string
	Funcs       []FuncDef
}

type FuncDef struct {
	Name       string // C Name
	GoName     string // Go Name
	ReturnType string // Go Return Type
	CReturnType string // C Return Type
	Params     []ParamDef
	HasReturn  bool
}

type ParamDef struct {
	Name   string
	GoType string
	CType  string // The type cast needed for C call
	IsStr  bool   // Helper to generate C.CString logic
}

func main() {
	// 1. Read Header
	contentBytes, err := os.ReadFile(HeaderFile)
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		os.Exit(1)
	}
	content := string(contentBytes)

	// 2. Parse Functions
	funcs := parseFunctions(content)

	// 3. Generate Code
	generateGoFile(funcs)
}

// Regex to handle C_TINYUSD_EXPORT, return types, names, and multi-line arguments
// Group 1: Return Type
// Group 2: Function Name
// Group 3: Arguments
var funcRegex = regexp.MustCompile(`(?m)^C_TINYUSD_EXPORT\s+([\w\s\*]+?)\s+(c_tinyusd_\w+)\s*\(([\s\S]*?)\);`)

func parseFunctions(header string) []FuncDef {
	var funcs []FuncDef

	matches := funcRegex.FindAllStringSubmatch(header, -1)
	for _, match := range matches {
		rawRet := strings.TrimSpace(match[1])
		cName := match[2]
		rawArgs := match[3]

		// Skip wide char functions for now (ending in _w) as they require special Go handling
		if strings.HasSuffix(cName, "_w") {
			continue
		}

		// Convert c_tinyusd_load_usdz -> LoadUsdz
		goName := toPascalCase(strings.TrimPrefix(cName, "c_tinyusd_"))

		// Map Return Type
		goRet, cRet := mapReturnType(rawRet)

		f := FuncDef{
			Name:        cName,
			GoName:      goName,
			ReturnType:  goRet,
			CReturnType: cRet,
			HasReturn:   goRet != "",
		}

		// Parse Parameters
		// 1. Remove newlines and extra spaces
		cleanArgs := strings.ReplaceAll(rawArgs, "\n", " ")
		cleanArgs = strings.Join(strings.Fields(cleanArgs), " ")

		if cleanArgs != "" && cleanArgs != "void" {
			args := strings.Split(cleanArgs, ",")
			for _, arg := range args {
				arg = strings.TrimSpace(arg)
				if arg == "" {
					continue
				}

				// Split Type and Name (Last word is name, rest is type)
				// Handle pointer * sticking to name or type
				parts := strings.Fields(arg)
				
				// Handle "const char *filename" vs "int *val"
				var name, rawType string
				
				// Find where the type ends and name starts
				// Simple heuristic: The last token is the name, unless it's a pointer like "type *name"
				last := parts[len(parts)-1]
				if strings.HasPrefix(last, "*") {
					name = strings.TrimPrefix(last, "*")
					rawType = strings.Join(parts[:len(parts)-1], " ") + "*"
				} else {
					name = last
					rawType = strings.Join(parts[:len(parts)-1], " ")
				}
				
				// Clean up pointer spacing "char *" -> "char*"
				rawType = strings.ReplaceAll(rawType, " *", "*")

				goType, cTypeCast, isStr := mapParamType(rawType)

				f.Params = append(f.Params, ParamDef{
					Name:   name,
					GoType: goType,
					CType:  cTypeCast,
					IsStr:  isStr,
				})
			}
		}
		funcs = append(funcs, f)
	}
	return funcs
}

// Maps C return types to Go types
func mapReturnType(cType string) (string, string) {
	cType = strings.TrimSpace(cType)
	if cType == "void" {
		return "", ""
	}
	
	// Check for Pointers
	if strings.HasSuffix(cType, "*") {
		// Remove the space before * if exists
		clean := strings.ReplaceAll(cType, " *", "*")
		// In Cgo, struct pointers usually map to *C.struct_Name
		return "*" + "C." + strings.TrimSuffix(clean, "*"), ""
	}

	// Primitives
	switch cType {
	case "int":
		return "int", "int"
	case "uint32_t":
		return "uint32", "uint32"
	case "uint64_t", "size_t":
		return "uint64", "uint64"
	case "float":
		return "float32", "float32"
	case "double":
		return "float64", "float64"
	case "CTinyUSDFormat", "CTinyUSDPrimType", "CTinyUSDValueType":
		return "int", "int" // Enums are ints in Cgo usually
	default:
		return "C." + cType, "" // Fallback
	}
}

// Maps C param types to Go types
func mapParamType(cType string) (goType string, cTypeCast string, isStr bool) {
	if strings.Contains(cType, "const char*") || strings.Contains(cType, "const char *") {
		return "string", "*C.char", true
	}
	
	// Pointers to specific tinyusd structs
	if strings.HasSuffix(cType, "*") {
		// e.g. CTinyUSDStage* -> *C.CTinyUSDStage
		// We keep it as a pointer to the C type for type safety in Go
		clean := strings.ReplaceAll(cType, " *", "*")
		clean = strings.TrimSuffix(clean, "*") // Remove * for the C prefix
		
		// Handle const prefix
		clean = strings.TrimPrefix(clean, "const ")
		
		return "*C." + clean, "", false
	}

	switch cType {
	case "int":
		return "int", "C.int", false
	case "uint32_t":
		return "uint32", "C.uint32_t", false
	case "uint64_t", "size_t":
		return "uint64", "C.size_t", false
	case "float":
		return "float32", "C.float", false
	case "double":
		return "float64", "C.double", false
	case "CTinyUSDFormat", "CTinyUSDPrimType", "CTinyUSDValueType":
		return "C."+cType, "", false
	default:
		// Fallback for structs passed by value or unknown types
		return "C." + cType, "", false
	}
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func generateGoFile(funcs []FuncDef) {
	t := template.Must(template.New("go").Parse(tmpl))
	
	data := TemplateData{
		PackageName: PackageName,
		Funcs:       funcs,
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}
	os.WriteFile(OutputFile, buf.Bytes(), 0644)
	fmt.Println("Generated", OutputFile)
}

const tmpl = `package {{.PackageName}}

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

{{range .Funcs}}
// {{.GoName}} wraps {{.Name}}
func {{.GoName}}({{range .Params}}{{.Name}} {{.GoType}}, {{end}}) {{.ReturnType}} {
    {{range .Params}}
        {{if .IsStr}}
            c_{{.Name}} := C.CString({{.Name}})
            defer C.free(unsafe.Pointer(c_{{.Name}}))
        {{end}}
    {{end}}

    {{if .HasReturn}}ret := {{end}}C.{{.Name}}(
        {{range .Params}}
            {{if .IsStr}}c_{{.Name}}{{else}}{{if .CType}}{{.CType}}({{.Name}}){{else}}{{.Name}}{{end}}{{end}},
        {{end}}
    )
    {{if .HasReturn}}
        {{if .CReturnType}}return {{.CReturnType}}(ret){{else}}return ret{{end}}
    {{end}}
}
{{end}}
`
