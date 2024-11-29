package options

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"
)

type NetworkPortOption func(ports *models.NetworkPort)

// WithNetworkPortOptionHttpEnable An option for SetNetworkPort to set the httpEnable
func WithNetworkPortOptionHttpEnable(enable enum.Toggle) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.HttpEnable = enable
	}
}

// WithNetworkPortOptionHttpPort An option for SetNetworkPort to set the httpPort
// Default value of httpPort is 80
func WithNetworkPortOptionHttpPort(httpPort int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.HttpPort = httpPort
	}
}

// WithNetworkPortOptionHttpsEnable An option for SetNetworkPort to set the httpsEnable
func WithNetworkPortOptionHttpsEnable(enable enum.Toggle) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.HttpsEnable = enable
	}
}

// WithNetworkPortOptionHttpsPort An option for SetNetworkPort to set the httpsPort
// Default value of httpsPort is 443
func WithNetworkPortOptionHttpsPort(https int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.HttpsPort = https
	}
}

// WithNetworkPortOptionMediaPort An option for SetNetworkPort to set the mediaPort
// Default value of mediaPort is 9000
func WithNetworkPortOptionMediaPort(media int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.MediaPort = media
	}
}

// WithNetworkPortOptionOnvifEnable An option for SetNetworkPort to set the onvifEnable
func WithNetworkPortOptionOnvifEnable(enable enum.Toggle) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.OnvifEnable = enable
	}
}

// WithNetworkPortOptionOnvifPort An option for SetNetworkPort to set the onvifPort
// Default value of onvifPort is 8000
func WithNetworkPortOptionOnvifPort(onvif int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.OnvifPort = onvif
	}
}

// WithNetworkPortOptionRtmpEnable An option for SetNetworkPort to set the rtmpEnable
func WithNetworkPortOptionRtmpEnable(enable enum.Toggle) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.RtmpEnable = enable
	}
}

// WithNetworkPortOptionRtmpPort An option for SetNetworkPort to set the rtmpPort
// Default value of rtmpPort is 1935
func WithNetworkPortOptionRtmpPort(rtmp int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.RtmpPort = rtmp
	}
}

// WithNetworkPortOptionRtspEnable An option for SetNetworkPort to set the rtspEnable
func WithNetworkPortOptionRtspEnable(enable enum.Toggle) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.RtspEnable = enable
	}
}

// WithNetworkPortOptionRtspPort An option for SetNetworkPort to set the rtspPort
// Default value of rtspPort is 554
func WithNetworkPortOptionRtspPort(rtsp int) NetworkPortOption {
	return func(nm *models.NetworkPort) {
		nm.RtspPort = rtsp
	}
}
