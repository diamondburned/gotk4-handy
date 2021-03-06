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
// extern void _gotk4_handy1_Carousel_ConnectPageChanged(gpointer, guint, guintptr);
import "C"

// glib.Type values for hdy-carousel.go.
var GTypeCarousel = externglib.Type(C.hdy_carousel_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeCarousel, F: marshalCarousel},
	})
}

// CarouselOverrider contains methods that are overridable.
type CarouselOverrider interface {
}

type Carousel struct {
	_ [0]func() // equal guard
	gtk.EventBox

	*externglib.Object
	gtk.Orientable
	Swipeable
}

var (
	_ externglib.Objector = (*Carousel)(nil)
	_ gtk.Binner          = (*Carousel)(nil)
)

func classInitCarouseller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapCarousel(obj *externglib.Object) *Carousel {
	return &Carousel{
		EventBox: gtk.EventBox{
			Bin: gtk.Bin{
				Container: gtk.Container{
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
			},
		},
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Swipeable: Swipeable{
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
	}
}

func marshalCarousel(p uintptr) (interface{}, error) {
	return wrapCarousel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_handy1_Carousel_ConnectPageChanged
func _gotk4_handy1_Carousel_ConnectPageChanged(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(index uint)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(index uint))
	}

	var _index uint // out

	_index = uint(arg1)

	f(_index)
}

// ConnectPageChanged: this signal is emitted after a page has been changed.
// This can be used to implement "infinite scrolling" by connecting to this
// signal and amending the pages.
func (self *Carousel) ConnectPageChanged(f func(index uint)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "page-changed", false, unsafe.Pointer(C._gotk4_handy1_Carousel_ConnectPageChanged), f)
}

// NewCarousel: create a new Carousel widget.
//
// The function returns the following values:
//
//    - carousel: newly created Carousel widget.
//
func NewCarousel() *Carousel {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_carousel_new()

	var _carousel *Carousel // out

	_carousel = wrapCarousel(externglib.Take(unsafe.Pointer(_cret)))

	return _carousel
}

// AllowLongSwipes: whether to allow swiping for more than one page at a time.
// If the value is FALSE, each swipe can only move to the adjacent pages.
//
// The function returns the following values:
//
//    - ok: TRUE if long swipes are allowed, FALSE otherwise.
//
func (self *Carousel) AllowLongSwipes() bool {
	var _arg0 *C.HdyCarousel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_allow_long_swipes(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AllowMouseDrag sets whether self can be dragged with mouse pointer.
//
// The function returns the following values:
//
//    - ok: TRUE if self can be dragged with mouse.
//
func (self *Carousel) AllowMouseDrag() bool {
	var _arg0 *C.HdyCarousel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_allow_mouse_drag(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AnimationDuration gets animation duration used by hdy_carousel_scroll_to().
//
// The function returns the following values:
//
//    - guint: animation duration in milliseconds.
//
func (self *Carousel) AnimationDuration() uint {
	var _arg0 *C.HdyCarousel // out
	var _cret C.guint        // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_animation_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Interactive gets whether self can be navigated.
//
// The function returns the following values:
//
//    - ok: TRUE if self can be swiped.
//
func (self *Carousel) Interactive() bool {
	var _arg0 *C.HdyCarousel // out
	var _cret C.gboolean     // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_interactive(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NPages gets the number of pages in self.
//
// The function returns the following values:
//
//    - guint: number of pages in self.
//
func (self *Carousel) NPages() uint {
	var _arg0 *C.HdyCarousel // out
	var _cret C.guint        // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_n_pages(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Position gets current scroll position in self. It's unitless, 1 matches 1
// page.
//
// The function returns the following values:
//
//    - gdouble: scroll position.
//
func (self *Carousel) Position() float64 {
	var _arg0 *C.HdyCarousel // out
	var _cret C.gdouble      // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_position(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RevealDuration gets duration of the animation used when adding or removing
// pages in milliseconds.
//
// The function returns the following values:
//
//    - guint: page reveal duration.
//
func (self *Carousel) RevealDuration() uint {
	var _arg0 *C.HdyCarousel // out
	var _cret C.guint        // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_reveal_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Spacing gets spacing between pages in pixels.
//
// The function returns the following values:
//
//    - guint: spacing between pages.
//
func (self *Carousel) Spacing() uint {
	var _arg0 *C.HdyCarousel // out
	var _cret C.guint        // in

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_carousel_get_spacing(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Insert inserts child into self at position position.
//
// If position is -1, or larger than the number of pages, child will be appended
// to the end.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - position to insert child in.
//
func (self *Carousel) Insert(child gtk.Widgetter, position int) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint         // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	_arg2 = C.gint(position)

	C.hdy_carousel_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// Prepend prepends child to self.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//
func (self *Carousel) Prepend(child gtk.Widgetter) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))

	C.hdy_carousel_prepend(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// Reorder moves child into position position.
//
// If position is -1, or larger than the number of pages, child will be moved to
// the end.
//
// The function takes the following parameters:
//
//    - child: widget to add.
//    - position to move child to.
//
func (self *Carousel) Reorder(child gtk.Widgetter, position int) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint         // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	_arg2 = C.gint(position)

	C.hdy_carousel_reorder(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
	runtime.KeepAlive(position)
}

// ScrollTo scrolls to widget position with an animation.
// Carousel:animation-duration property can be used for controlling the
// duration.
//
// The function takes the following parameters:
//
//    - widget: child of self.
//
func (self *Carousel) ScrollTo(widget gtk.Widgetter) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	C.hdy_carousel_scroll_to(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// ScrollToFull scrolls to widget position with an animation.
//
// The function takes the following parameters:
//
//    - widget: child of self.
//    - duration: animation duration in milliseconds.
//
func (self *Carousel) ScrollToFull(widget gtk.Widgetter, duration int64) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gint64       // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))
	_arg2 = C.gint64(duration)

	C.hdy_carousel_scroll_to_full(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(duration)
}

// SetAllowLongSwipes sets whether to allow swiping for more than one page at a
// time. If the value is FALSE, each swipe can only move to the adjacent pages.
//
// The function takes the following parameters:
//
//    - allowLongSwipes: whether to allow long swipes.
//
func (self *Carousel) SetAllowLongSwipes(allowLongSwipes bool) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if allowLongSwipes {
		_arg1 = C.TRUE
	}

	C.hdy_carousel_set_allow_long_swipes(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowLongSwipes)
}

// SetAllowMouseDrag sets whether self can be dragged with mouse pointer. If
// allow_mouse_drag is FALSE, dragging is only available on touch.
//
// The function takes the following parameters:
//
//    - allowMouseDrag: whether self can be dragged with mouse pointer.
//
func (self *Carousel) SetAllowMouseDrag(allowMouseDrag bool) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if allowMouseDrag {
		_arg1 = C.TRUE
	}

	C.hdy_carousel_set_allow_mouse_drag(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(allowMouseDrag)
}

// SetAnimationDuration sets animation duration used by
// hdy_carousel_scroll_to().
//
// The function takes the following parameters:
//
//    - duration: animation duration in milliseconds.
//
func (self *Carousel) SetAnimationDuration(duration uint) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.guint        // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(duration)

	C.hdy_carousel_set_animation_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}

// SetInteractive sets whether self can be navigated. This can be used to
// temporarily disable a Carousel to only allow swiping in a certain state.
//
// The function takes the following parameters:
//
//    - interactive: whether self can be swiped.
//
func (self *Carousel) SetInteractive(interactive bool) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if interactive {
		_arg1 = C.TRUE
	}

	C.hdy_carousel_set_interactive(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interactive)
}

// SetRevealDuration sets duration of the animation used when adding or removing
// pages in milliseconds.
//
// The function takes the following parameters:
//
//    - revealDuration: new reveal duration value.
//
func (self *Carousel) SetRevealDuration(revealDuration uint) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.guint        // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(revealDuration)

	C.hdy_carousel_set_reveal_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(revealDuration)
}

// SetSpacing sets spacing between pages in pixels.
//
// The function takes the following parameters:
//
//    - spacing: new spacing value.
//
func (self *Carousel) SetSpacing(spacing uint) {
	var _arg0 *C.HdyCarousel // out
	var _arg1 C.guint        // out

	_arg0 = (*C.HdyCarousel)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(spacing)

	C.hdy_carousel_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(spacing)
}
