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
		{T: externglib.Type(C.hdy_preferences_page_get_type()), F: marshalPreferencesPager},
	})
}

type PreferencesPage struct {
	gtk.Bin
}

func wrapPreferencesPage(obj *externglib.Object) *PreferencesPage {
	return &PreferencesPage{
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

func marshalPreferencesPager(p uintptr) (interface{}, error) {
	return wrapPreferencesPage(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesPage creates a new PreferencesPage.
func NewPreferencesPage() *PreferencesPage {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_preferences_page_new()

	var _preferencesPage *PreferencesPage // out

	_preferencesPage = wrapPreferencesPage(externglib.Take(unsafe.Pointer(_cret)))

	return _preferencesPage
}

// IconName gets the icon name for self, or NULL.
func (self *PreferencesPage) IconName() string {
	var _arg0 *C.HdyPreferencesPage // out
	var _cret *C.gchar              // in

	_arg0 = (*C.HdyPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_preferences_page_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title gets the title of self, or NULL.
func (self *PreferencesPage) Title() string {
	var _arg0 *C.HdyPreferencesPage // out
	var _cret *C.gchar              // in

	_arg0 = (*C.HdyPreferencesPage)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_preferences_page_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetIconName sets the icon name for self.
//
// The function takes the following parameters:
//
//    - iconName: icon name, or NULL.
//
func (self *PreferencesPage) SetIconName(iconName string) {
	var _arg0 *C.HdyPreferencesPage // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.HdyPreferencesPage)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_preferences_page_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetTitle sets the title of self.
//
// The function takes the following parameters:
//
//    - title of the page, or NULL.
//
func (self *PreferencesPage) SetTitle(title string) {
	var _arg0 *C.HdyPreferencesPage // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.HdyPreferencesPage)(unsafe.Pointer(self.Native()))
	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_preferences_page_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}