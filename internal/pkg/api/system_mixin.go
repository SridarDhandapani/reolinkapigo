package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/options"
)

type SystemMixin struct{}

// Get the general system information
func (sm *SystemMixin) GetGeneralSystem() func(handler *rest.RestHandler) (*models.DeviceGeneralInformation, error) {
	return func(handler *rest.RestHandler) (*models.DeviceGeneralInformation, error) {
		payloadTime := map[string]interface{}{
			"cmd":    "GetTime",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		payloadNorm := map[string]interface{}{
			"cmd":    "GetNorm",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		resultTime, err := handler.Request("POST", payloadTime, "GetTime")

		if err != nil {
			return nil, err
		}

		resultNorm, err := handler.Request("POST", payloadNorm, "GetNorm")

		if err != nil {
			return nil, err
		}

		var timeData *models.TimeInformation
		var dstData *models.DstInformation
		var normData string

		err = json.Unmarshal(resultTime.Value["Time"], &timeData)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resultTime.Value["Dst"], &dstData)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(resultNorm.Value["norm"], &normData)

		if err != nil {
			return nil, err
		}

		return &models.DeviceGeneralInformation{
			Time: timeData,
			Dst:  dstData,
			Norm: &models.DeviceNorm{
				Norm: normData,
			},
		}, nil
	}
}

// Get the camera performance information
// See examples/responses/GetPerformance.json for example response data
func (sm *SystemMixin) GetPerformance() func(handler *rest.RestHandler) (*models.DevicePerformanceInformation, error) {
	return func(handler *rest.RestHandler) (*models.DevicePerformanceInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetPerformance",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetPerformance")

		if err != nil {
			return nil, err
		}

		var devicePerformance *models.DevicePerformanceInformation

		err = json.Unmarshal(result.Value["Performance"], &devicePerformance)

		if err != nil {
			return nil, err
		}

		return devicePerformance, nil
	}
}

// Get the camera device information
// See examples/responses/GetDevInfo.json for example response data
func (sm *SystemMixin) GetDeviceInformation() func(handler *rest.RestHandler) (*models.DeviceInformation, error) {
	return func(handler *rest.RestHandler) (*models.DeviceInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetDevInfo",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetDevInfo")

		if err != nil {
			return nil, err
		}

		var deviceInfo *models.DeviceInformation

		err = json.Unmarshal(result.Value["DevInfo"], &deviceInfo)

		if err != nil {
			return nil, err
		}

		return deviceInfo, nil
	}
}

// Reboot the camera
func (sm *SystemMixin) RebootCamera() func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "Reboot",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "Reboot")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not reboot camera. camera responded with %v", result.Value)
	}
}

// Get the camera DST information
// See examples/response/GetDSTInfo.json for example response data
func (sm *SystemMixin) GetDstInformation() func(handler *rest.RestHandler) (*models.DstInformation,
	*models.TimeInformation, error) {
	return func(handler *rest.RestHandler) (*models.DstInformation, *models.TimeInformation, error) {
		payload := map[string]interface{}{
			"cmd":    "GetTime",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetTime")

		if err != nil {
			return nil, nil, err
		}

		var dstData *models.DstInformation
		var timeData *models.TimeInformation

		err = json.Unmarshal(result.Value["Dst"], &dstData)

		if err != nil {
			return nil, nil, err
		}

		err = json.Unmarshal(result.Value["Time"], &timeData)

		if err != nil {
			return nil, nil, err
		}
		return dstData, timeData, nil
	}
}

// GetDeviceName Get the camera name
func (sm *SystemMixin) GetDeviceName() func(handler *rest.RestHandler) (*models.DeviceName, error) {
	return func(handler *rest.RestHandler) (*models.DeviceName, error) {
		payload := map[string]interface{}{
			"cmd":    "GetDevName",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetDevName")

		if err != nil {
			return nil, err
		}

		var devName *models.DeviceName

		err = json.Unmarshal(result.Value["DevName"], &devName)

		if err != nil {
			return nil, err
		}

		return devName, nil
	}
}

// SetDeviceName updates the device name.
func (nm *NetworkMixin) SetDeviceName(deviceNameOption ...options.DeviceNameOption) func(handler *rest.RestHandler) (bool,
	error) {

	// Defaults
	deviceName := &models.DeviceName{
		Name: "Camera1",
	}
	for _, op := range deviceNameOption {
		op(deviceName)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetDevName",
			"action": 0,
			"param": map[string]interface{}{
				"DevName": map[string]interface{}{
					"name": deviceName.Name,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetDevName")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not set device name. camera responded with %v", result.Value)
	}
}

// SetDeviceTime updates the device name.
func (nm *NetworkMixin) SetDeviceTime(deviceTimeOption ...options.DeviceTimeOption) func(handler *rest.RestHandler) (bool,
	error) {

	// Defaults
	deviceTime := &models.DeviceTime{}
	for _, op := range deviceTimeOption {
		op(deviceTime)
	}

	data, _ := json.Marshal(deviceTime)
	var param map[string]interface{}
	_ = json.Unmarshal(data, &param)
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetTime",
			"action": 0,
			"param":  param,
		}

		result, err := handler.Request("POST", payload, "SetTime")

		if err != nil {
			return false, err
		}

		var respCode int

		err = json.Unmarshal(result.Value["rspCode"], &respCode)

		if err != nil {
			return false, err
		}

		if respCode == 200 {
			return true, nil
		}

		return false, fmt.Errorf("camera could not set device time. camera responded with %v", result.Value)
	}
}
