// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <handy.h>
import "C"

// glib.Type values for hdy-swipe-group.go.
var GTypeSwipeGroup = externglib.Type(C.hdy_swipe_group_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSwipeGroup, F: marshalSwipeGroup},
	})
}

// SwipeGroupOverrider contains methods that are overridable.
type SwipeGroupOverrider interface {
}

type SwipeGroup struct {
	_ [0]func() // equal guard
	*externglib.Object

	gtk.Buildable
}

var (
	_ externglib.Objector = (*SwipeGroup)(nil)
)

func classInitSwipeGrouper(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSwipeGroup(obj *externglib.Object) *SwipeGroup {
	return &SwipeGroup{
		Object: obj,
		Buildable: gtk.Buildable{
			Object: obj,
		},
	}
}

func marshalSwipeGroup(p uintptr) (interface{}, error) {
	return wrapSwipeGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSwipeGroup: create a new SwipeGroup object.
//
// The function returns the following values:
//
//    - swipeGroup: newly created SwipeGroup object.
//
func NewSwipeGroup() *SwipeGroup {
	var _cret *C.HdySwipeGroup // in

	_cret = C.hdy_swipe_group_new()

	var _swipeGroup *SwipeGroup // out

	_swipeGroup = wrapSwipeGroup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _swipeGroup
}

// AddSwipeable: when the widget is destroyed or no longer referenced elsewhere,
// it will be removed from the swipe group.
//
// The function takes the following parameters:
//
//    - swipeable to add.
//
func (self *SwipeGroup) AddSwipeable(swipeable Swipeabler) {
	var _arg0 *C.HdySwipeGroup // out
	var _arg1 *C.HdySwipeable  // out

	_arg0 = (*C.HdySwipeGroup)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.HdySwipeable)(unsafe.Pointer(externglib.InternObject(swipeable).Native()))

	C.hdy_swipe_group_add_swipeable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(swipeable)
}

// Swipeables returns the list of swipeables associated with self.
//
// The function returns the following values:
//
//    - sList of swipeables. The list is owned by libhandy and should not be
//      modified.
//
func (self *SwipeGroup) Swipeables() []*Swipeable {
	var _arg0 *C.HdySwipeGroup // out
	var _cret *C.GSList        // in

	_arg0 = (*C.HdySwipeGroup)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.hdy_swipe_group_get_swipeables(_arg0)
	runtime.KeepAlive(self)

	var _sList []*Swipeable // out

	_sList = make([]*Swipeable, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.HdySwipeable)(v)
		var dst *Swipeable // out
		dst = wrapSwipeable(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// RemoveSwipeable removes a widget from a SwipeGroup.
//
// The function takes the following parameters:
//
//    - swipeable to remove.
//
func (self *SwipeGroup) RemoveSwipeable(swipeable Swipeabler) {
	var _arg0 *C.HdySwipeGroup // out
	var _arg1 *C.HdySwipeable  // out

	_arg0 = (*C.HdySwipeGroup)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.HdySwipeable)(unsafe.Pointer(externglib.InternObject(swipeable).Native()))

	C.hdy_swipe_group_remove_swipeable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(swipeable)
}
