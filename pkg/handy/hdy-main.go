// Code generated by girgen. DO NOT EDIT.

package handy

// #cgo pkg-config: libhandy-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <handy.h>
import "C"

// Init: call this function just after initializing GTK, if you are using
// Application it means it must be called when the #GApplication::startup signal
// is emitted. If libhandy has already been initialized, the function will
// simply return.
//
// This makes sure translations, types, themes, and icons for the Handy library
// are set up properly.
func Init() {
	C.hdy_init()
}