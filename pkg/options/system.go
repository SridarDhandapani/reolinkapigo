package options

import "github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"

type DeviceNameOption func(osd *models.DeviceName)

// WithDeviceNameOptionName Set the camera device name
func WithDeviceNameOptionName(name string) DeviceNameOption {
	return func(dn *models.DeviceName) {
		dn.Name = name
	}
}
