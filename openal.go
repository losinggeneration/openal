// Package openal provides some AL defines that golang.org/x/mobile/exp/audio/al
// doesn't currently provide. It also provides some ALC functionality, namely
// contexts.
package openal

/*
#cgo darwin   CFLAGS:  -DGOOS_darwin
#cgo linux    CFLAGS:  -DGOOS_linux
#cgo darwin   LDFLAGS: -framework OpenAL
#cgo linux    LDFLAGS: -lopenal

#include <stdlib.h>

#ifdef GOOS_darwin
#include <OpenAL/al.h>
#include <OpenAL/alc.h>
#endif

#ifdef GOOS_linux
#include <AL/al.h>
#include <AL/alc.h>
#endif
*/
import "C"
import "unsafe"

const (
	TRUE  = C.AL_TRUE
	FALSE = C.AL_FALSE

	SOURCE_RELATIVE           = C.AL_SOURCE_RELATIVE
	CONE_INNER_ANGLE          = C.AL_CONE_INNER_ANGLE
	CONE_OUTER_ANGLE          = C.AL_CONE_OUTER_ANGLE
	PITCH                     = C.AL_PITCH
	POSITION                  = C.AL_POSITION
	DIRECTION                 = C.AL_DIRECTION
	VELOCITY                  = C.AL_VELOCITY
	LOOPING                   = C.AL_LOOPING
	BUFFER                    = C.AL_BUFFER
	GAIN                      = C.AL_GAIN
	MIN_GAIN                  = C.AL_MIN_GAIN
	MAX_GAIN                  = C.AL_MAX_GAIN
	ORIENTATION               = C.AL_ORIENTATION
	SOURCE_STATE              = C.AL_SOURCE_STATE
	INITIAL                   = C.AL_INITIAL
	PLAYING                   = C.AL_PLAYING
	PAUSED                    = C.AL_PAUSED
	STOPPED                   = C.AL_STOPPED
	BUFFERS_QUEUED            = C.AL_BUFFERS_QUEUED
	BUFFERS_PROCESSED         = C.AL_BUFFERS_PROCESSED
	REFERENCE_DISTANCE        = C.AL_REFERENCE_DISTANCE
	ROLLOFF_FACTOR            = C.AL_ROLLOFF_FACTOR
	CONE_OUTER_GAIN           = C.AL_CONE_OUTER_GAIN
	MAX_DISTANCE              = C.AL_MAX_DISTANCE
	SEC_OFFSET                = C.AL_SEC_OFFSET
	SAMPLE_OFFSET             = C.AL_SAMPLE_OFFSET
	BYTE_OFFSET               = C.AL_BYTE_OFFSET
	SOURCE_TYPE               = C.AL_SOURCE_TYPE
	STATIC                    = C.AL_STATIC
	STREAMING                 = C.AL_STREAMING
	UNDETERMINED              = C.AL_UNDETERMINED
	FORMAT_MONO8              = C.AL_FORMAT_MONO8
	FORMAT_MONO16             = C.AL_FORMAT_MONO16
	FORMAT_STEREO8            = C.AL_FORMAT_STEREO8
	FORMAT_STEREO16           = C.AL_FORMAT_STEREO16
	FREQUENCY                 = C.AL_FREQUENCY
	BITS                      = C.AL_BITS
	CHANNELS                  = C.AL_CHANNELS
	SIZE                      = C.AL_SIZE
	UNUSED                    = C.AL_UNUSED
	PENDING                   = C.AL_PENDING
	PROCESSED                 = C.AL_PROCESSED
	NO_ERROR                  = C.AL_NO_ERROR
	INVALID_NAME              = C.AL_INVALID_NAME
	INVALID_ENUM              = C.AL_INVALID_ENUM
	INVALID_VALUE             = C.AL_INVALID_VALUE
	INVALID_OPERATION         = C.AL_INVALID_OPERATION
	OUT_OF_MEMORY             = C.AL_OUT_OF_MEMORY
	VENDOR                    = C.AL_VENDOR
	VERSION                   = C.AL_VERSION
	RENDERER                  = C.AL_RENDERER
	EXTENSIONS                = C.AL_EXTENSIONS
	DOPPLER_FACTOR            = C.AL_DOPPLER_FACTOR
	DOPPLER_VELOCITY          = C.AL_DOPPLER_VELOCITY
	SPEED_OF_SOUND            = C.AL_SPEED_OF_SOUND
	DISTANCE_MODEL            = C.AL_DISTANCE_MODEL
	INVERSE_DISTANCE          = C.AL_INVERSE_DISTANCE
	INVERSE_DISTANCE_CLAMPED  = C.AL_INVERSE_DISTANCE_CLAMPED
	LINEAR_DISTANCE           = C.AL_LINEAR_DISTANCE
	LINEAR_DISTANCE_CLAMPED   = C.AL_LINEAR_DISTANCE_CLAMPED
	EXPONENT_DISTANCE         = C.AL_EXPONENT_DISTANCE
	EXPONENT_DISTANCE_CLAMPED = C.AL_EXPONENT_DISTANCE_CLAMPED
)

type Device struct {
	device *C.ALCdevice
}

func OpenDevice(devicename *string) *Device {
	var d Device
	if devicename != nil {
		str := C.CString(*devicename)
		defer C.free(unsafe.Pointer(str))
		d.device = C.alcOpenDevice((*C.ALCchar)(unsafe.Pointer(str)))
	} else {
		d.device = C.alcOpenDevice(nil)
	}
	if d.device == nil {
		return nil
	}

	return &d
}

func (d *Device) Close() bool {
	return C.alcCloseDevice(d.device) == 1
}

type Context struct {
	context *C.ALCcontext
}

func NewContext(device *Device, attrs ...int) *Context {
	l := make([]C.ALCint, len(attrs)+1)
	for i, a := range attrs {
		l[i] = C.ALCint(a)
	}
	l[len(attrs)] = 0

	c := C.alcCreateContext(device.device, (*C.ALCint)(unsafe.Pointer(&l[0])))
	if c == nil {
		return nil
	}

	return &Context{context: c}
}

func (c *Context) MakeCurrent() bool {
	return C.alcMakeContextCurrent(c.context) == 1
}

func (c *Context) Process() {
	C.alcProcessContext(c.context)
}

func (c *Context) Suspend() {
	C.alcSuspendContext(c.context)
}

func (c *Context) Destroy() {
	C.alcDestroyContext(c.context)
}

func GetCurrentContext() *Context {
	c := C.alcGetCurrentContext()
	if c == nil {
		return nil
	}
	return &Context{context: c}
}

func (c *Context) GetDevice() *Device {
	d := C.alcGetContextsDevice(c.context)
	if d == nil {
		return nil
	}
	return &Device{device: d}
}
