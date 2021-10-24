// Code generated by girgen. DO NOT EDIT.

package handy

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v3"
)

// #cgo pkg-config: libhandy-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <handy.h>
// GdkPixbuf* _gotk4_handy1_AvatarImageLoadFunc(gint, gpointer);
// extern void callbackDelete(gpointer);
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.hdy_avatar_get_type()), F: marshalAvatarrer},
	})
}

// AvatarImageLoadFunc: returned Pixbuf is expected to be square with width and
// height set to size. The image is cropped to a circle without any scaling or
// transformation.
//
// Deprecated: use hdy_avatar_set_loadable_icon() instead.
type AvatarImageLoadFunc func(size int) (pixbuf *gdkpixbuf.Pixbuf)

//export _gotk4_handy1_AvatarImageLoadFunc
func _gotk4_handy1_AvatarImageLoadFunc(arg0 C.gint, arg1 C.gpointer) (cret *C.GdkPixbuf) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var size int // out

	size = int(arg0)

	fn := v.(AvatarImageLoadFunc)
	pixbuf := fn(size)

	if pixbuf != nil {
		cret = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
		C.g_object_ref(C.gpointer(pixbuf.Native()))
	}

	return cret
}

type Avatar struct {
	gtk.DrawingArea
}

func wrapAvatar(obj *externglib.Object) *Avatar {
	return &Avatar{
		DrawingArea: gtk.DrawingArea{
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

func marshalAvatarrer(p uintptr) (interface{}, error) {
	return wrapAvatar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAvatar creates a new Avatar.
//
// The function takes the following parameters:
//
//    - size of the avatar.
//    - text used to generate the color and initials if show_initials is TRUE.
//    The color is selected at random if text is empty.
//    - showInitials: whether to show the initials or the fallback icon on top
//    of the color generated based on text.
//
func NewAvatar(size int, text string, showInitials bool) *Avatar {
	var _arg1 C.gint       // out
	var _arg2 *C.gchar     // out
	var _arg3 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gint(size)
	if text != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if showInitials {
		_arg3 = C.TRUE
	}

	_cret = C.hdy_avatar_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(size)
	runtime.KeepAlive(text)
	runtime.KeepAlive(showInitials)

	var _avatar *Avatar // out

	_avatar = wrapAvatar(externglib.Take(unsafe.Pointer(_cret)))

	return _avatar
}

// DrawToPixbuf renders self into a pixbuf at size and scale_factor. This can be
// used to export the fallback avatar.
//
// The function takes the following parameters:
//
//    - size of the pixbuf.
//    - scaleFactor: scale factor.
//
func (self *Avatar) DrawToPixbuf(size, scaleFactor int) *gdkpixbuf.Pixbuf {
	var _arg0 *C.HdyAvatar // out
	var _arg1 C.gint       // out
	var _arg2 C.gint       // out
	var _cret *C.GdkPixbuf // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	_arg1 = C.gint(size)
	_arg2 = C.gint(scaleFactor)

	_cret = C.hdy_avatar_draw_to_pixbuf(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scaleFactor)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// DrawToPixbufAsync renders asynchronously self into a pixbuf at size and
// scale_factor. This can be used to export the fallback avatar.
//
// The function takes the following parameters:
//
//    - ctx: optional #GCancellable object, NULL to ignore.
//    - size of the pixbuf.
//    - scaleFactor: scale factor.
//    - callback to call when the avatar is generated.
//
func (self *Avatar) DrawToPixbufAsync(ctx context.Context, size, scaleFactor int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.HdyAvatar          // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.gint                // out
	var _arg2 C.gint                // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gint(size)
	_arg2 = C.gint(scaleFactor)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.hdy_avatar_draw_to_pixbuf_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(size)
	runtime.KeepAlive(scaleFactor)
	runtime.KeepAlive(callback)
}

// DrawToPixbufFinish finishes an asynchronous draw of an avatar to a pixbuf.
//
// The function takes the following parameters:
//
//    - asyncResult: Result.
//
func (self *Avatar) DrawToPixbufFinish(asyncResult gio.AsyncResulter) *gdkpixbuf.Pixbuf {
	var _arg0 *C.HdyAvatar    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(asyncResult.Native()))

	_cret = C.hdy_avatar_draw_to_pixbuf_finish(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(asyncResult)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// IconName gets the name of the icon in the icon theme to use when the icon
// should be displayed.
func (self *Avatar) IconName() string {
	var _arg0 *C.HdyAvatar // out
	var _cret *C.gchar     // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_avatar_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LoadableIcon gets the Icon set via hdy_avatar_set_loadable_icon().
func (self *Avatar) LoadableIcon() gio.LoadableIconner {
	var _arg0 *C.HdyAvatar     // out
	var _cret *C.GLoadableIcon // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_avatar_get_loadable_icon(_arg0)
	runtime.KeepAlive(self)

	var _loadableIcon gio.LoadableIconner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gio.LoadableIconner)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.LoadableIconner")
			}
			_loadableIcon = rv
		}
	}

	return _loadableIcon
}

// ShowInitials returns whether initials are used for the fallback or the icon.
func (self *Avatar) ShowInitials() bool {
	var _arg0 *C.HdyAvatar // out
	var _cret C.gboolean   // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_avatar_get_show_initials(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Size returns the size of the avatar.
func (self *Avatar) Size() int {
	var _arg0 *C.HdyAvatar // out
	var _cret C.gint       // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_avatar_get_size(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Text: get the text used to generate the fallback initials and color.
func (self *Avatar) Text() string {
	var _arg0 *C.HdyAvatar // out
	var _cret *C.gchar     // in

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))

	_cret = C.hdy_avatar_get_text(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetIconName sets the name of the icon in the icon theme to use when the icon
// should be displayed. If no name is set, the avatar-default-symbolic icon will
// be used. If the name doesn't match a valid icon, it is an error and no icon
// will be displayed. If the icon theme is changed, the image will be updated
// automatically.
//
// The function takes the following parameters:
//
//    - iconName: name of the icon from the icon theme.
//
func (self *Avatar) SetIconName(iconName string) {
	var _arg0 *C.HdyAvatar // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_avatar_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetImageLoadFunc: callback which is called when the custom image need to be
// reloaded for some reason (e.g. scale-factor changes).
//
// Deprecated: use hdy_avatar_set_loadable_icon() instead.
//
// The function takes the following parameters:
//
//    - loadImage: callback to set a custom image.
//
func (self *Avatar) SetImageLoadFunc(loadImage AvatarImageLoadFunc) {
	var _arg0 *C.HdyAvatar             // out
	var _arg1 C.HdyAvatarImageLoadFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	if loadImage != nil {
		_arg1 = (*[0]byte)(C._gotk4_handy1_AvatarImageLoadFunc)
		_arg2 = C.gpointer(gbox.Assign(loadImage))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.hdy_avatar_set_image_load_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(loadImage)
}

// SetLoadableIcon sets the Icon to use as an avatar. The previous avatar is
// displayed till the new avatar is loaded, to immediately remove the custom
// avatar set the loadable-icon to NULL.
//
// The Icon set via this function is prefered over a set AvatarImageLoadFunc.
//
// The function takes the following parameters:
//
//    - icon: Icon.
//
func (self *Avatar) SetLoadableIcon(icon gio.LoadableIconner) {
	var _arg0 *C.HdyAvatar     // out
	var _arg1 *C.GLoadableIcon // out

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	if icon != nil {
		_arg1 = (*C.GLoadableIcon)(unsafe.Pointer(icon.Native()))
	}

	C.hdy_avatar_set_loadable_icon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(icon)
}

// SetShowInitials sets whether the initials should be shown on the fallback
// avatar or the icon.
//
// The function takes the following parameters:
//
//    - showInitials: whether the initials should be shown on the fallback
//    avatar or the icon.
//
func (self *Avatar) SetShowInitials(showInitials bool) {
	var _arg0 *C.HdyAvatar // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	if showInitials {
		_arg1 = C.TRUE
	}

	C.hdy_avatar_set_show_initials(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showInitials)
}

// SetSize sets the size of the avatar.
//
// The function takes the following parameters:
//
//    - size to be used for the avatar.
//
func (self *Avatar) SetSize(size int) {
	var _arg0 *C.HdyAvatar // out
	var _arg1 C.gint       // out

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	_arg1 = C.gint(size)

	C.hdy_avatar_set_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(size)
}

// SetText: set the text used to generate the fallback initials color.
//
// The function takes the following parameters:
//
//    - text used to get the initials and color.
//
func (self *Avatar) SetText(text string) {
	var _arg0 *C.HdyAvatar // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.HdyAvatar)(unsafe.Pointer(self.Native()))
	if text != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.hdy_avatar_set_text(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(text)
}
