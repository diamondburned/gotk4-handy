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
// #include <stdlib.h>
// #include <glib-object.h>
// #include <handy.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.hdy_keypad_get_type()), F: marshalKeypadder},
	})
}

type Keypad struct {
	gtk.Bin
}

var (
	_ gtk.Binner = (*Keypad)(nil)
)

func wrapKeypad(obj *externglib.Object) *Keypad {
	return &Keypad{
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

func marshalKeypadder(p uintptr) (interface{}, error) {
	return wrapKeypad(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewKeypad: create a new Keypad widget.
//
// The function takes the following parameters:
//
//    - symbolsVisible: whether the hash, plus, and asterisk symbols should be
//    visible.
//    - lettersVisible: whether the letters below the digits should be visible.
//
func NewKeypad(symbolsVisible, lettersVisible bool) *Keypad {
	var _arg1 C.gboolean   // out
	var _arg2 C.gboolean   // out
	var _cret *C.GtkWidget // in

	if symbolsVisible {
		_arg1 = C.TRUE
	}
	if lettersVisible {
		_arg2 = C.TRUE
	}

	_cret = C.hdy_keypad_new(_arg1, _arg2)
	runtime.KeepAlive(symbolsVisible)
	runtime.KeepAlive(lettersVisible)

	var _keypad *Keypad // out

	_keypad = wrapKeypad(externglib.Take(unsafe.Pointer(_cret)))

	return _keypad
}

// ColumnSpacing returns the amount of space between the columns of self.
func (self *Keypad) ColumnSpacing() uint {
	var _arg0 *C.HdyKeypad // out
	var _cret C.guint      // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_column_spacing(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EndAction returns the widget for the lower right corner (or left, in RTL
// locales) of self.
func (self *Keypad) EndAction() gtk.Widgetter {
	var _arg0 *C.HdyKeypad // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_end_action(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Entry: get the connected entry. See hdy_keypad_set_entry() for details.
func (self *Keypad) Entry() *gtk.Entry {
	var _arg0 *C.HdyKeypad // out
	var _cret *C.GtkEntry  // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_entry(_arg0)
	runtime.KeepAlive(self)

	var _entry *gtk.Entry // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_entry = &gtk.Entry{
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
			Object: obj,
			CellEditable: gtk.CellEditable{
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
			Editable: gtk.Editable{
				Object: obj,
			},
		}
	}

	return _entry
}

// LettersVisible returns whether self should display the standard letters below
// the digits on its buttons.
func (self *Keypad) LettersVisible() bool {
	var _arg0 *C.HdyKeypad // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_letters_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowSpacing returns the amount of space between the rows of self.
func (self *Keypad) RowSpacing() uint {
	var _arg0 *C.HdyKeypad // out
	var _cret C.guint      // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_row_spacing(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// StartAction returns the widget for the lower left corner (or right, in RTL
// locales) of self.
func (self *Keypad) StartAction() gtk.Widgetter {
	var _arg0 *C.HdyKeypad // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_start_action(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gtk.Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SymbolsVisible returns whether self should display the standard letters below
// the digits on its buttons.
//
// Returns Whether self should display the hash and asterisk buttons, and should
// display the plus symbol at the bottom of its 0 button.
func (self *Keypad) SymbolsVisible() bool {
	var _arg0 *C.HdyKeypad // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_keypad_get_symbols_visible(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetColumnSpacing sets the amount of space between columns of self.
//
// The function takes the following parameters:
//
//    - spacing: amount of space to insert between columns.
//
func (self *Keypad) SetColumnSpacing(spacing uint) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 C.guint      // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(spacing)

	C.hdy_keypad_set_column_spacing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(spacing)
}

// SetEndAction sets the widget for the lower right corner (or left, in RTL
// locales) of self.
//
// The function takes the following parameters:
//
//    - endAction: end action widget.
//
func (self *Keypad) SetEndAction(endAction gtk.Widgetter) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	if endAction != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(endAction.Native()))
	}

	C.hdy_keypad_set_end_action(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(endAction)
}

// SetEntry binds entry to self and blocks any input which wouldn't be possible
// to type with with the keypad.
//
// The function takes the following parameters:
//
//    - entry: Entry.
//
func (self *Keypad) SetEntry(entry *gtk.Entry) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 *C.GtkEntry  // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	if entry != nil {
		_arg1 = (*C.GtkEntry)(unsafe.Pointer(entry.Native()))
	}

	C.hdy_keypad_set_entry(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(entry)
}

// SetLettersVisible sets whether self should display the standard letters below
// the digits on its buttons.
//
// The function takes the following parameters:
//
//    - lettersVisible: whether the letters below the digits should be visible.
//
func (self *Keypad) SetLettersVisible(lettersVisible bool) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	if lettersVisible {
		_arg1 = C.TRUE
	}

	C.hdy_keypad_set_letters_visible(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(lettersVisible)
}

// SetRowSpacing sets the amount of space between rows of self.
//
// The function takes the following parameters:
//
//    - spacing: amount of space to insert between rows.
//
func (self *Keypad) SetRowSpacing(spacing uint) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 C.guint      // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(spacing)

	C.hdy_keypad_set_row_spacing(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(spacing)
}

// SetStartAction sets the widget for the lower left corner (or right, in RTL
// locales) of self.
//
// The function takes the following parameters:
//
//    - startAction: start action widget.
//
func (self *Keypad) SetStartAction(startAction gtk.Widgetter) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	if startAction != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(startAction.Native()))
	}

	C.hdy_keypad_set_start_action(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(startAction)
}

// SetSymbolsVisible sets whether self should display the hash and asterisk
// buttons, and should display the plus symbol at the bottom of its 0 button.
//
// The function takes the following parameters:
//
//    - symbolsVisible: whether the hash, plus, and asterisk symbols should be
//    visible.
//
func (self *Keypad) SetSymbolsVisible(symbolsVisible bool) {
	var _arg0 *C.HdyKeypad // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyKeypad)(unsafe.Pointer(self.Native()))
	if symbolsVisible {
		_arg1 = C.TRUE
	}

	C.hdy_keypad_set_symbols_visible(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(symbolsVisible)
}
