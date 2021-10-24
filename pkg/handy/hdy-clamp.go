// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"runtime"
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
		{T: externglib.Type(C.hdy_clamp_get_type()), F: marshalClamper},
	})
}

type Clamp struct {
	gtk.Bin

	gtk.Orientable
	*externglib.Object
}

func wrapClamp(obj *externglib.Object) *Clamp {
	return &Clamp{
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
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalClamper(p uintptr) (interface{}, error) {
	return wrapClamp(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewClamp creates a new Clamp.
func NewClamp() *Clamp {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_clamp_new()

	var _clamp *Clamp // out

	_clamp = wrapClamp(externglib.Take(unsafe.Pointer(_cret)))

	return _clamp
}

// MaximumSize gets the maximum size to allocate to the contained child. It is
// the width if self is horizontal, or the height if it is vertical.
func (self *Clamp) MaximumSize() int {
	var _arg0 *C.HdyClamp // out
	var _cret C.gint      // in

	_arg0 = (*C.HdyClamp)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_clamp_get_maximum_size(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TighteningThreshold gets the size starting from which the clamp will tighten
// its grip on the child.
func (self *Clamp) TighteningThreshold() int {
	var _arg0 *C.HdyClamp // out
	var _cret C.gint      // in

	_arg0 = (*C.HdyClamp)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_clamp_get_tightening_threshold(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetMaximumSize sets the maximum size to allocate to the contained child. It
// is the width if self is horizontal, or the height if it is vertical.
//
// The function takes the following parameters:
//
//    - maximumSize: maximum size.
//
func (self *Clamp) SetMaximumSize(maximumSize int) {
	var _arg0 *C.HdyClamp // out
	var _arg1 C.gint      // out

	_arg0 = (*C.HdyClamp)(unsafe.Pointer(self.Native()))
	_arg1 = C.gint(maximumSize)

	C.hdy_clamp_set_maximum_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maximumSize)
}

// SetTighteningThreshold sets the size starting from which the clamp will
// tighten its grip on the child.
//
// The function takes the following parameters:
//
//    - tighteningThreshold: tightening threshold.
//
func (self *Clamp) SetTighteningThreshold(tighteningThreshold int) {
	var _arg0 *C.HdyClamp // out
	var _arg1 C.gint      // out

	_arg0 = (*C.HdyClamp)(unsafe.Pointer(self.Native()))
	_arg1 = C.gint(tighteningThreshold)

	C.hdy_clamp_set_tightening_threshold(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tighteningThreshold)
}
