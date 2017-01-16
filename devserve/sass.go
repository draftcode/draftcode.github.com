package main

// #cgo LDFLAGS: -lsass
// #include <sass.h>
// #include <stdlib.h>
import "C"
import "fmt"
import "unsafe"

func CompileSass(path string) (string, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	ctx := C.sass_make_file_context(p)
	defer C.sass_delete_file_context(ctx)

	if C.sass_compile_file_context(ctx) != 0 {
		msg := C.sass_context_get_error_message((*C.struct_Sass_Context)(ctx))
		return "", fmt.Errorf(C.GoString(msg))
	}
	return C.GoString(C.sass_context_get_output_string((*C.struct_Sass_Context)(ctx))), nil
}
