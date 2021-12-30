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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.hdy_header_bar_get_type()), F: marshalHeaderBarrer},
	})
}

type HeaderBar struct {
	_ [0]func() // equal guard
	gtk.Container
}

var (
	_ gtk.Containerer = (*HeaderBar)(nil)
)

func wrapHeaderBar(obj *externglib.Object) *HeaderBar {
	return &HeaderBar{
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
	}
}

func marshalHeaderBarrer(p uintptr) (interface{}, error) {
	return wrapHeaderBar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHeaderBar creates a new HeaderBar widget.
//
// The function returns the following values:
//
//    - headerBar: new HeaderBar.
//
func NewHeaderBar() *HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.hdy_header_bar_new()

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// CenteringPolicy gets the policy self follows to horizontally align its center
// widget.
//
// The function returns the following values:
//
//    - centeringPolicy: centering policy.
//
func (self *HeaderBar) CenteringPolicy() CenteringPolicy {
	var _arg0 *C.HdyHeaderBar      // out
	var _cret C.HdyCenteringPolicy // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_centering_policy(_arg0)
	runtime.KeepAlive(self)

	var _centeringPolicy CenteringPolicy // out

	_centeringPolicy = CenteringPolicy(_cret)

	return _centeringPolicy
}

// CustomTitle retrieves the custom title widget of the header. See
// hdy_header_bar_set_custom_title().
//
// The function returns the following values:
//
//    - widget (optional): custom title widget of the header, or NULL if none has
//      been set explicitly.
//
func (self *HeaderBar) CustomTitle() gtk.Widgetter {
	var _arg0 *C.HdyHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_custom_title(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// DecorationLayout gets the decoration layout set with
// hdy_header_bar_set_decoration_layout().
//
// The function returns the following values:
//
//    - utf8: decoration layout.
//
func (self *HeaderBar) DecorationLayout() string {
	var _arg0 *C.HdyHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_decoration_layout(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasSubtitle retrieves whether the header bar reserves space for a subtitle,
// regardless if one is currently set or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the header bar reserves space for a subtitle.
//
func (self *HeaderBar) HasSubtitle() bool {
	var _arg0 *C.HdyHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_has_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize gets whether self should interpolate its size on visible
// child change.
//
// See hdy_header_bar_set_interpolate_size().
//
// The function returns the following values:
//
//    - ok: TRUE if self interpolates its size on visible child change, FALSE if
//      not.
//
func (self *HeaderBar) InterpolateSize() bool {
	var _arg0 *C.HdyHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_interpolate_size(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether this header bar shows the standard window
// decorations.
//
// The function returns the following values:
//
//    - ok: TRUE if the decorations are shown.
//
func (self *HeaderBar) ShowCloseButton() bool {
	var _arg0 *C.HdyHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_show_close_button(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Subtitle retrieves the subtitle of the header. See
// hdy_header_bar_set_subtitle().
//
// The function returns the following values:
//
//    - utf8 (optional): subtitle of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (self *HeaderBar) Subtitle() string {
	var _arg0 *C.HdyHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title retrieves the title of the header. See hdy_header_bar_set_title().
//
// The function returns the following values:
//
//    - utf8 (optional): title of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (self *HeaderBar) Title() string {
	var _arg0 *C.HdyHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in self will take.
//
// The function returns the following values:
//
//    - guint: transition duration.
//
func (self *HeaderBar) TransitionDuration() uint {
	var _arg0 *C.HdyHeaderBar // out
	var _cret C.guint         // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_transition_duration(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionRunning returns whether the self is currently in a transition from
// one page to another.
//
// The function returns the following values:
//
//    - ok: TRUE if the transition is currently running, FALSE otherwise.
//
func (self *HeaderBar) TransitionRunning() bool {
	var _arg0 *C.HdyHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_header_bar_get_transition_running(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PackEnd adds child to self:, packed with reference to the end of the self:.
//
// The function takes the following parameters:
//
//    - child to be added to self:.
//
func (self *HeaderBar) PackEnd(child gtk.Widgetter) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.hdy_header_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// PackStart adds child to self:, packed with reference to the start of the
// self:.
//
// The function takes the following parameters:
//
//    - child to be added to self:.
//
func (self *HeaderBar) PackStart(child gtk.Widgetter) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.hdy_header_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetCenteringPolicy sets the policy self must follow to horizontally align its
// center widget.
//
// The function takes the following parameters:
//
//    - centeringPolicy: centering policy.
//
func (self *HeaderBar) SetCenteringPolicy(centeringPolicy CenteringPolicy) {
	var _arg0 *C.HdyHeaderBar      // out
	var _arg1 C.HdyCenteringPolicy // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.HdyCenteringPolicy(centeringPolicy)

	C.hdy_header_bar_set_centering_policy(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(centeringPolicy)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes any
// title set by hdy_header_bar_set_title() or hdy_header_bar_set_subtitle(). To
// achieve the same style as the builtin title and subtitle, use the “title” and
// “subtitle” style classes.
//
// You should set the custom title to NULL, for the header title label to be
// visible again.
//
// The function takes the following parameters:
//
//    - titleWidget (optional): custom widget to use for a title.
//
func (self *HeaderBar) SetCustomTitle(titleWidget gtk.Widgetter) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if titleWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))
	}

	C.hdy_header_bar_set_custom_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(titleWidget)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close, icon (the window icon)
// and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left, and
// minimize, maximize and close buttons on the right.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (self *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if layout != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_header_bar_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(layout)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
//
// The function takes the following parameters:
//
//    - setting: TRUE to reserve space for a subtitle.
//
func (self *HeaderBar) SetHasSubtitle(setting bool) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.hdy_header_bar_set_has_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetInterpolateSize sets whether or not self will interpolate the size of its
// opposing orientation when changing the visible child. If TRUE, self will
// interpolate its size between the one of the previous visible child and the
// one of the new visible child, according to the set transition duration and
// the orientation, e.g. if self is horizontal, it will interpolate the its
// height.
//
// The function takes the following parameters:
//
//    - interpolateSize: TRUE to interpolate the size.
//
func (self *HeaderBar) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.hdy_header_bar_set_interpolate_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(interpolateSize)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
//
// The function takes the following parameters:
//
//    - setting: TRUE to show standard window decorations.
//
func (self *HeaderBar) SetShowCloseButton(setting bool) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.hdy_header_bar_set_show_close_button(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(setting)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a user
// an additional detail to help them identify the current view.
//
// Note that HdyHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the HeaderBar:has-subtitle
// property to FALSE.
//
// The function takes the following parameters:
//
//    - subtitle (optional): subtitle, or NULL.
//
func (self *HeaderBar) SetSubtitle(subtitle string) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if subtitle != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(subtitle)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_header_bar_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the application
// name.
//
// The function takes the following parameters:
//
//    - title (optional): title, or NULL.
//
func (self *HeaderBar) SetTitle(title string) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_header_bar_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetTransitionDuration sets the duration that transitions between pages in
// self will take.
//
// The function takes the following parameters:
//
//    - duration: new duration, in milliseconds.
//
func (self *HeaderBar) SetTransitionDuration(duration uint) {
	var _arg0 *C.HdyHeaderBar // out
	var _arg1 C.guint         // out

	_arg0 = (*C.HdyHeaderBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(duration)

	C.hdy_header_bar_set_transition_duration(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(duration)
}
