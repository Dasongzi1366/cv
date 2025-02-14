package cv

func (c VideoCaptureAPI) String() string {
	switch c {
	case VideoCaptureAny:
		return "video-capture-any"
	case VideoCaptureV4L2:
		return "video-capture-v4l2"
	case VideoCaptureFirewire:
		return "video-capture-firewire"
	case VideoCaptureQT:
		return "video-capture-qt"
	case VideoCaptureUnicap:
		return "video-capture-unicap"
	case VideoCaptureDshow:
		return "video-capture-dshow"
	case VideoCapturePvAPI:
		return "video-capture-pvapi"
	case VideoCaptureOpenNI:
		return "video-capture-openni"
	case VideoCaptureOpenNIAsus:
		return "video-capture-openni-asus"
	case VideoCaptureAndroid:
		return "video-capture-android"
	case VideoCaptureXiAPI:
		return "video-capture-xiapi"
	case VideoCaptureAVFoundation:
		return "video-capture-av-foundation"
	case VideoCaptureGiganetix:
		return "video-capture-giganetix"
	case VideoCaptureMSMF:
		return "video-capture-msmf"
	case VideoCaptureWinRT:
		return "video-capture-winrt"
	case VideoCaptureIntelPerc:
		return "video-capture-intel-perc"
	case VideoCaptureOpenNI2:
		return "video-capture-openni2"
	case VideoCaptureOpenNI2Asus:
		return "video-capture-openni2-asus"
	case VideoCaptureGPhoto2:
		return "video-capture-gphoto2"
	case VideoCaptureGstreamer:
		return "video-capture-gstreamer"
	case VideoCaptureFFmpeg:
		return "video-capture-ffmpeg"
	case VideoCaptureImages:
		return "video-capture-images"
	case VideoCaptureAravis:
		return "video-capture-aravis"
	case VideoCaptureOpencvMjpeg:
		return "video-capture-opencv-mjpeg"
	case VideoCaptureIntelMFX:
		return "video-capture-intel-mfx"
	case VideoCaptureXINE:
		return "video-capture-xine"
	}
	return ""
}

func (c VideoCaptureProperties) String() string {
	switch c {
	case VideoCapturePosMsec:
		return "video-capture-pos-msec"
	case VideoCapturePosFrames:
		return "video-capture-pos-frames"
	case VideoCapturePosAVIRatio:
		return "video-capture-pos-avi-ratio"
	case VideoCaptureFrameWidth:
		return "video-capture-frame-width"
	case VideoCaptureFrameHeight:
		return "video-capture-frame-height"
	case VideoCaptureFPS:
		return "video-capture-fps"
	case VideoCaptureFOURCC:
		return "video-capture-fourcc"
	case VideoCaptureFrameCount:
		return "video-capture-frame-count"
	case VideoCaptureFormat:
		return "video-capture-format"
	case VideoCaptureMode:
		return "video-capture-mode"
	case VideoCaptureBrightness:
		return "video-capture-brightness"
	case VideoCaptureContrast:
		return "video-capture-contrast"
	case VideoCaptureSaturation:
		return "video-capture-saturation"
	case VideoCaptureHue:
		return "video-capture-hue"
	case VideoCaptureGain:
		return "video-capture-gain"
	case VideoCaptureExposure:
		return "video-capture-exposure"
	case VideoCaptureConvertRGB:
		return "video-capture-convert-rgb"
	case VideoCaptureWhiteBalanceBlueU:
		return "video-capture-white-balanced-blue-u"
	case VideoCaptureWhiteBalanceRedV:
		return "video-capture-white-balanced-red-v"
	case VideoCaptureRectification:
		return "video-capture-rectification"
	case VideoCaptureMonochrome:
		return "video-capture-monochrome"
	case VideoCaptureSharpness:
		return "video-capture-sharpness"
	case VideoCaptureAutoExposure:
		return "video-capture-auto-exposure"
	case VideoCaptureGamma:
		return "video-capture-gamma"
	case VideoCaptureTemperature:
		return "video-capture-temperature"
	case VideoCaptureTrigger:
		return "video-capture-trigger"
	case VideoCaptureTriggerDelay:
		return "video-capture-trigger-delay"
	case VideoCaptureZoom:
		return "video-capture-zoom"
	case VideoCaptureFocus:
		return "video-capture-focus"
	case VideoCaptureGUID:
		return "video-capture-guid"
	case VideoCaptureISOSpeed:
		return "video-capture-iso-speed"
	case VideoCaptureBacklight:
		return "video-capture-backlight"
	case VideoCapturePan:
		return "video-capture-pan"
	case VideoCaptureTilt:
		return "video-capture-tilt"
	case VideoCaptureRoll:
		return "video-capture-roll"
	case VideoCaptureIris:
		return "video-capture-iris"
	case VideoCaptureSettings:
		return "video-capture-settings"
	case VideoCaptureBufferSize:
		return "video-capture-buffer-size"
	case VideoCaptureAutoFocus:
		return "video-capture-auto-focus"
	case VideoCaptureSarNumerator:
		return "video-capture-sar-numerator"
	case VideoCaptureSarDenominator:
		return "video-capture-sar-denominator"
	case VideoCaptureBackend:
		return "video-capture-backend"
	case VideoCaptureChannel:
		return "video-capture-channel"
	case VideoCaptureAutoWB:
		return "video-capture-auto-wb"
	case VideoCaptureWBTemperature:
		return "video-capture-wb-temperature"
	case VideoCaptureCodecPixelFormat:
		return "video-capture-pixel-format"
	case VideoCaptureBitrate:
		return "video-capture-bitrate"
	}
	return ""
}
