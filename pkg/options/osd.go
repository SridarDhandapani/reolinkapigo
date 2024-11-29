package options

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"
)

type OsdOption func(osd *models.Osd)

// WithOsdOptionBgColor Set the OSD background color on or off
func WithOsdOptionBgColor(bgColor enum.Toggle) OsdOption {
	return func(o *models.Osd) {
		o.BgColor = bgColor
	}
}

// WithOsdOptionChannel Set the OSD channel
func WithOsdOptionChannel(channel int) OsdOption {
	return func(o *models.Osd) {
		o.Channel = channel
	}
}

// WithOsdOptionOsdChannelEnable Set the OSD channel on or off
func WithOsdOptionOsdChannelEnable(enable enum.Toggle) OsdOption {
	return func(o *models.Osd) {
		o.OsdChannel.Enable = enable
	}
}

// WithOsdOptionOsdChannelName Set the OSD channel name
func WithOsdOptionOsdChannelName(name string) OsdOption {
	return func(o *models.Osd) {
		o.OsdChannel.Name = name
	}
}

// WithOsdOptionOsdChannelPos Set the OSD channel position
func WithOsdOptionOsdChannelPos(position enum.OsdPosition) OsdOption {
	return func(o *models.Osd) {
		o.OsdChannel.Pos = position.Value()
	}
}

// WithOsdOptionOsdTimeEnable Set the OSD time as on or off
func WithOsdOptionOsdTimeEnable(enable enum.Toggle) OsdOption {
	return func(o *models.Osd) {
		o.OsdTime.Enable = enable
	}
}

// WithOsdOptionOsdTimePos Set the OSD time position
func WithOsdOptionOsdTimePos(position enum.OsdPosition) OsdOption {
	return func(o *models.Osd) {
		o.OsdTime.Pos = position.Value()
	}
}

// WithOsdOptionWatermark Set the OSD watermark
func WithOsdOptionWatermark(enable enum.Toggle) OsdOption {
	return func(o *models.Osd) {
		o.Watermark = enable
	}
}
