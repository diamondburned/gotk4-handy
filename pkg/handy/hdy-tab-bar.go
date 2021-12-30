// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <handy.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.hdy_tab_bar_get_type()), F: marshalTabBarrer},
	})
}

type TabBar struct {
	_ [0]func() // equal guard
	gtk.Bin
}

var (
	_ gtk.Binner = (*TabBar)(nil)
)

func wrapTabBar(obj *externglib.Object) *TabBar {
	return &TabBar{
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
	}
}

func marshalTabBarrer(p uintptr) (interface{}, error) {
	return wrapTabBar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectExtraDragDataReceived: this signal is emitted when content allowed via
// TabBar:extra-drag-dest-targets is dropped onto a tab.
//
// See Widget::drag-data-received.
func (self *TabBar) ConnectExtraDragDataReceived(f func(page TabPage, context gdk.DragContext, data *gtk.SelectionData, info, time uint)) externglib.SignalHandle {
	return self.Connect("extra-drag-data-received", f)
}

// NewTabBar creates a new TabBar widget.
//
// The function returns the following values:
//
//    - tabBar: new TabBar.
//
func NewTabBar() *TabBar {
	var _cret *C.HdyTabBar // in

	_cret = C.hdy_tab_bar_new()

	var _tabBar *TabBar // out

	_tabBar = wrapTabBar(externglib.Take(unsafe.Pointer(_cret)))

	return _tabBar
}

// Autohide gets whether the tabs automatically hide, see
// hdy_tab_bar_set_autohide().
//
// The function returns the following values:
//
//    - ok: whether the tabs automatically hide.
//
func (self *TabBar) Autohide() bool {
	var _arg0 *C.HdyTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_autohide(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EndActionWidget gets the widget shown after the tabs.
//
// The function returns the following values:
//
//    - widget (optional) shown after the tabs, or NULL.
//
func (self *TabBar) EndActionWidget() gtk.Widgetter {
	var _arg0 *C.HdyTabBar // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_end_action_widget(_arg0)
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

// ExpandTabs gets whether tabs should expand, see
// hdy_tab_bar_set_expand_tabs().
//
// The function returns the following values:
//
//    - ok: whether tabs should expand.
//
func (self *TabBar) ExpandTabs() bool {
	var _arg0 *C.HdyTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_expand_tabs(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtraDragDestTargets gets extra drag destination targets, see
// hdy_tab_bar_set_extra_drag_dest_targets().
//
// The function returns the following values:
//
//    - targetList (optional): extra drag targets, or NULL.
//
func (self *TabBar) ExtraDragDestTargets() *gtk.TargetList {
	var _arg0 *C.HdyTabBar     // out
	var _cret *C.GtkTargetList // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_extra_drag_dest_targets(_arg0)
	runtime.KeepAlive(self)

	var _targetList *gtk.TargetList // out

	if _cret != nil {
		_targetList = (*gtk.TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gtk_target_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_targetList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gtk_target_list_unref((*C.GtkTargetList)(intern.C))
			},
		)
	}

	return _targetList
}

// Inverted gets whether tabs use inverted layout, see
// hdy_tab_bar_set_inverted().
//
// The function returns the following values:
//
//    - ok: whether tabs use inverted layout.
//
func (self *TabBar) Inverted() bool {
	var _arg0 *C.HdyTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_inverted(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsOverflowing gets whether self is overflowing.
//
// The function returns the following values:
//
//    - ok: whether self is overflowing.
//
func (self *TabBar) IsOverflowing() bool {
	var _arg0 *C.HdyTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_is_overflowing(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartActionWidget gets the widget shown before the tabs.
//
// The function returns the following values:
//
//    - widget (optional) shown before the tabs, or NULL.
//
func (self *TabBar) StartActionWidget() gtk.Widgetter {
	var _arg0 *C.HdyTabBar // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_start_action_widget(_arg0)
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

// TabsRevealed gets the value of the TabBar:tabs-revealed property.
//
// The function returns the following values:
//
//    - ok: whether the tabs are current revealed.
//
func (self *TabBar) TabsRevealed() bool {
	var _arg0 *C.HdyTabBar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_tabs_revealed(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// View gets the TabView self controls.
//
// The function returns the following values:
//
//    - tabView (optional) self controls.
//
func (self *TabBar) View() *TabView {
	var _arg0 *C.HdyTabBar  // out
	var _cret *C.HdyTabView // in

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_tab_bar_get_view(_arg0)
	runtime.KeepAlive(self)

	var _tabView *TabView // out

	if _cret != nil {
		_tabView = wrapTabView(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _tabView
}

// SetAutohide sets whether the tabs automatically hide.
//
// If autohide is TRUE, the tab bar disappears when the associated TabView has 0
// or 1 tab, no pinned tabs, and no tab is being transferred.
//
// Autohide is enabled by default.
//
// See TabBar:tabs-revealed.
//
// The function takes the following parameters:
//
//    - autohide: whether the tabs automatically hide.
//
func (self *TabBar) SetAutohide(autohide bool) {
	var _arg0 *C.HdyTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if autohide {
		_arg1 = C.TRUE
	}

	C.hdy_tab_bar_set_autohide(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(autohide)
}

// SetEndActionWidget sets the widget to show after the tabs.
//
// The function takes the following parameters:
//
//    - widget (optional) to show after the tabs, or NULL.
//
func (self *TabBar) SetEndActionWidget(widget gtk.Widgetter) {
	var _arg0 *C.HdyTabBar // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.hdy_tab_bar_set_end_action_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetExpandTabs sets whether tabs should expand.
//
// If expand_tabs is TRUE, the tabs will always vary width filling the whole
// width when possible, otherwise tabs will always have the minimum possible
// size.
//
// Expand is enabled by default.
//
// The function takes the following parameters:
//
//    - expandTabs: whether to expand tabs.
//
func (self *TabBar) SetExpandTabs(expandTabs bool) {
	var _arg0 *C.HdyTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if expandTabs {
		_arg1 = C.TRUE
	}

	C.hdy_tab_bar_set_expand_tabs(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expandTabs)
}

// SetExtraDragDestTargets sets extra drag destination targets.
//
// This allows to drag arbitrary content onto tabs, for example URLs in a web
// browser.
//
// If a tab is hovered for a certain period of time while dragging the content,
// it will be automatically selected.
//
// After content is dropped, the TabBar::extra-drag-data-received signal can be
// used to retrieve and process the drag data.
//
// The function takes the following parameters:
//
//    - extraDragDestTargets (optional): extra drag targets, or NULL.
//
func (self *TabBar) SetExtraDragDestTargets(extraDragDestTargets *gtk.TargetList) {
	var _arg0 *C.HdyTabBar     // out
	var _arg1 *C.GtkTargetList // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if extraDragDestTargets != nil {
		_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(extraDragDestTargets)))
	}

	C.hdy_tab_bar_set_extra_drag_dest_targets(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(extraDragDestTargets)
}

// SetInverted sets whether tabs tabs use inverted layout.
//
// If inverted is TRUE, non-pinned tabs will have the close button at the
// beginning and the indicator at the end rather than the opposite.
//
// The function takes the following parameters:
//
//    - inverted: whether tabs use inverted layout.
//
func (self *TabBar) SetInverted(inverted bool) {
	var _arg0 *C.HdyTabBar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.hdy_tab_bar_set_inverted(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(inverted)
}

// SetStartActionWidget sets the widget to show before the tabs.
//
// The function takes the following parameters:
//
//    - widget (optional) to show before the tabs, or NULL.
//
func (self *TabBar) SetStartActionWidget(widget gtk.Widgetter) {
	var _arg0 *C.HdyTabBar // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.hdy_tab_bar_set_start_action_widget(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// SetView sets the TabView self controls.
//
// The function takes the following parameters:
//
//    - view (optional): TabView.
//
func (self *TabBar) SetView(view *TabView) {
	var _arg0 *C.HdyTabBar  // out
	var _arg1 *C.HdyTabView // out

	_arg0 = (*C.HdyTabBar)(unsafe.Pointer(self.Native()))
	if view != nil {
		_arg1 = (*C.HdyTabView)(unsafe.Pointer(view.Native()))
	}

	C.hdy_tab_bar_set_view(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(view)
}
