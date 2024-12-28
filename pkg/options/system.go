package options

import (
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"time"
)

type DeviceNameOption func(osd *models.DeviceName)

// WithDeviceNameOptionName set the camera device name.
func WithDeviceNameOptionName(name string) DeviceNameOption {
	return func(dn *models.DeviceName) {
		dn.Name = name
	}
}

type DeviceTimeOption func(osd *models.DeviceTime)

// WithDeviceTimeOptionTime sets the camera device time.
func WithDeviceTimeOptionTime(t time.Time) DeviceTimeOption {
	return func(dt *models.DeviceTime) {
		_, offSet := t.Zone()
		t := models.TimeInformation{
			Day:      t.Day(),
			Hour:     t.Hour(),
			HourFmt:  0,
			Min:      t.Minute(),
			Mon:      int(t.Month()),
			Sec:      t.Second(),
			TimeFmt:  "DD/MM/YYYY",
			TimeZone: offSet,
			Year:     t.Year(),
		}
		dt.Time = &t
	}
}
