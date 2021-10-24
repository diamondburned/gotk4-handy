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
		{T: externglib.Type(C.hdy_view_switcher_title_get_type()), F: marshalViewSwitcherTitler},
	})
}

type ViewSwitcherTitle struct {
	gtk.Bin
}

func wrapViewSwitcherTitle(obj *externglib.Object) *ViewSwitcherTitle {
	return &ViewSwitcherTitle{
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
	}
}

func marshalViewSwitcherTitler(p uintptr) (interface{}, error) {
	return wrapViewSwitcherTitle(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewViewSwitcherTitle creates a new ViewSwitcherTitle widget.
func NewViewSwitcherTitle() *ViewSwitcherTitle {
	var _cret *C.HdyViewSwitcherTitle // in

	_cret = C.hdy_view_switcher_title_new()

	var _viewSwitcherTitle *ViewSwitcherTitle // out

	_viewSwitcherTitle = wrapViewSwitcherTitle(externglib.Take(unsafe.Pointer(_cret)))

	return _viewSwitcherTitle
}

// Policy gets the policy of self.
func (self *ViewSwitcherTitle) Policy() ViewSwitcherPolicy {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret C.HdyViewSwitcherPolicy // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_policy(_arg0)
	runtime.KeepAlive(self)

	var _viewSwitcherPolicy ViewSwitcherPolicy // out

	_viewSwitcherPolicy = ViewSwitcherPolicy(_cret)

	return _viewSwitcherPolicy
}

// Stack: get the Stack being controlled by the ViewSwitcher.
func (self *ViewSwitcherTitle) Stack() *gtk.Stack {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret *C.GtkStack             // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_stack(_arg0)
	runtime.KeepAlive(self)

	var _stack *gtk.Stack // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_stack = &gtk.Stack{
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
			}
		}
	}

	return _stack
}

// Subtitle gets the subtitle of self. See
// hdy_view_switcher_title_set_subtitle().
func (self *ViewSwitcherTitle) Subtitle() string {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret *C.gchar                // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self. See hdy_view_switcher_title_set_title().
func (self *ViewSwitcherTitle) Title() string {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret *C.gchar                // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TitleVisible: get whether the title label of self is visible.
func (self *ViewSwitcherTitle) TitleVisible() bool {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret C.gboolean              // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_title_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ViewSwitcherEnabled gets whether self's view switcher is enabled.
//
// See hdy_view_switcher_title_set_view_switcher_enabled().
func (self *ViewSwitcherTitle) ViewSwitcherEnabled() bool {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _cret C.gboolean              // in

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_view_switcher_title_get_view_switcher_enabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetPolicy sets the policy of self.
//
// The function takes the following parameters:
//
//    - policy: new policy.
//
func (self *ViewSwitcherTitle) SetPolicy(policy ViewSwitcherPolicy) {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _arg1 C.HdyViewSwitcherPolicy // out

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	_arg1 = C.HdyViewSwitcherPolicy(policy)

	C.hdy_view_switcher_title_set_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(policy)
}

// SetStack sets the Stack to control.
//
// The function takes the following parameters:
//
//    - stack: Stack.
//
func (self *ViewSwitcherTitle) SetStack(stack *gtk.Stack) {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _arg1 *C.GtkStack             // out

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if stack != nil {
		_arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	}

	C.hdy_view_switcher_title_set_stack(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(stack)
}

// SetSubtitle sets the subtitle of self. The subtitle should give a user
// additional details.
//
// The function takes the following parameters:
//
//    - subtitle: subtitle, or NULL.
//
func (self *ViewSwitcherTitle) SetSubtitle(subtitle string) {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if subtitle != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(subtitle)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_view_switcher_title_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of self. The title should give a user additional
// details. A good title should not include the application name.
//
// The function takes the following parameters:
//
//    - title: title, or NULL.
//
func (self *ViewSwitcherTitle) SetTitle(title string) {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_view_switcher_title_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetViewSwitcherEnabled: make self enable or disable its view switcher. If it
// is disabled, the title will be displayed instead. This allows to
// programmatically and prematurely hide the view switcher of self even if it
// fits in the available space.
//
// This can be used e.g. to ensure the view switcher is hidden below a certain
// window width, or any other constraint you find suitable.
//
// The function takes the following parameters:
//
//    - enabled: TRUE to enable the view switcher, FALSE to disable it.
//
func (self *ViewSwitcherTitle) SetViewSwitcherEnabled(enabled bool) {
	var _arg0 *C.HdyViewSwitcherTitle // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.HdyViewSwitcherTitle)(unsafe.Pointer(self.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.hdy_view_switcher_title_set_view_switcher_enabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enabled)
}
