// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <handy.h>
import "C"

// glib.Type values for hdy-carousel-indicator-dots.go.
var GTypeCarouselIndicatorDots = externglib.Type(C.hdy_carousel_indicator_dots_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCarouselIndicatorDots, F: marshalCarouselIndicatorDots},
	})
}

// CarouselIndicatorDotsOverrider contains methods that are overridable.
type CarouselIndicatorDotsOverrider interface {
}

type CarouselIndicatorDots struct {
	_ [0]func() // equal guard
	gtk.DrawingArea

	*externglib.Object
	gtk.Orientable
}

var (
	_ externglib.Objector = (*CarouselIndicatorDots)(nil)
	_ gtk.Widgetter       = (*CarouselIndicatorDots)(nil)
)

func classInitCarouselIndicatorDotser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCarouselIndicatorDots(obj *externglib.Object) *CarouselIndicatorDots {
	return &CarouselIndicatorDots{
		DrawingArea: gtk.DrawingArea{
			Widget: gtk.Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: gtk.Buildable{
					Object: obj,
				},
			},
		},
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
	}
}

func marshalCarouselIndicatorDots(p uintptr) (interface{}, error) {
	return wrapCarouselIndicatorDots(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCarouselIndicatorDots: create a new CarouselIndicatorDots widget.
//
// The function returns the following values:
//
//    - carouselIndicatorDots: newly created CarouselIndicatorDots widget.
//
func NewCarouselIndicatorDots() *CarouselIndicatorDots {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_carousel_indicator_dots_new()

	var _carouselIndicatorDots *CarouselIndicatorDots // out

	_carouselIndicatorDots = wrapCarouselIndicatorDots(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _carouselIndicatorDots
}

// Carousel: get the Carousel the indicator uses.
//
// See: hdy_carousel_indicator_dots_set_carousel().
//
// The function returns the following values:
//
//    - carousel (optional) or NULL if none has been set.
//
func (self *CarouselIndicatorDots) Carousel() *Carousel {
	var _arg0 *C.HdyCarouselIndicatorDots // out
	var _cret *C.HdyCarousel              // in

	_arg0 = (*C.HdyCarouselIndicatorDots)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_indicator_dots_get_carousel(_arg0)
	runtime.KeepAlive(self)

	var _carousel *Carousel // out

	if _cret != nil {
		_carousel = wrapCarousel(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _carousel
}

// SetCarousel sets the Carousel to use.
//
// The function takes the following parameters:
//
//    - carousel (optional): Carousel.
//
func (self *CarouselIndicatorDots) SetCarousel(carousel *Carousel) {
	var _arg0 *C.HdyCarouselIndicatorDots // out
	var _arg1 *C.HdyCarousel              // out

	_arg0 = (*C.HdyCarouselIndicatorDots)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if carousel != nil {
		_arg1 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(carousel).Native()))
	}

	C.hdy_carousel_indicator_dots_set_carousel(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(carousel)
}
