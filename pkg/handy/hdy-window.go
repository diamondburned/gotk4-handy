// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #cgo pkg-config: libhandy-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <handy.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.hdy_window_get_type()), F: marshalWindower},
	})
}

type Window struct {
	gtk.Window
}

func wrapWindow(obj *externglib.Object) *Window {
	return &Window{
		Window: gtk.Window{
			Bin: gtk.Bin{
				Container: gtk.Container{
					Widget: gtk.Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: gtk.Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalWindower(p uintptr) (interface{}, error) {
	return wrapWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindow creates a new Window.
func NewWindow() *Window {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_window_new()

	var _window *Window // out

	_window = wrapWindow(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _window
}
