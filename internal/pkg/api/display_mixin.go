package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/options"
)

type DisplayMixin struct {
}

//type osdChannel struct {
//	Enable int
//	Name   string
//	Pos    string
//}
//
//type osdTime struct {
//	Enable int
//	Pos    string
//}
//
//type osd struct {
//	BgColor    int
//	Channel    int
//	OsdChannel osdChannel
//	OsdTime    osdTime
//	Watermark  int
//}

// Get the camera's Osd information
func (dm *DisplayMixin) GetOSD() func(handler *rest.RestHandler) (*models.Osd, error) {
	return func(handler *rest.RestHandler) (*models.Osd, error) {
		payload := map[string]interface{}{
			"cmd":    "GetOsd",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("POST", payload, "GetOsd")

		if err != nil {
			return nil, err
		}

		var osdData *models.Osd

		err = json.Unmarshal(result.Value["Osd"], &osdData)

		if err != nil {
			return nil, err
		}

		return osdData, nil
	}
}

// Get the camera's mask information
func (dm *DisplayMixin) GetMask() func(handler *rest.RestHandler) (*models.MaskData, error) {
	return func(handler *rest.RestHandler) (*models.MaskData, error) {
		payload := map[string]interface{}{
			"cmd":    "GetMask",
			"action": 1,
			"param": map[string]interface{}{
				"channel": 0,
			},
		}

		result, err := handler.Request("POST", payload, "GetMask")

		if err != nil {
			return nil, err
		}

		var maskData *models.MaskData

		err = json.Unmarshal(result.Value["Mask"], &maskData)

		if err != nil {
			return nil, err
		}

		return maskData, nil
	}
}

// SetOSD Set the camera's on-screen display
func (dm *DisplayMixin) SetOSD(osdOption ...options.OsdOption) func(handler *rest.RestHandler) (bool,
	error) {

	osd := &models.Osd{
		BgColor: enum.Disabled,
		Channel: 0,
		OsdChannel: models.OsdChannel{
			Enable: enum.Disabled,
			Name:   "Camera1",
			Pos:    "Lower Right",
		},
		OsdTime: models.OsdTime{
			Enable: enum.Disabled,
			Pos:    "Top Center",
		},
		Watermark: enum.Disabled,
	}

	for _, op := range osdOption {
		op(osd)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetOsd",
			"action": 1,
			"param": map[string]interface{}{
				"Osd": map[string]interface{}{
					"bgcolor": osd.BgColor,
					"channel": osd.Channel,
					"osdChannel": map[string]interface{}{
						"enable": osd.OsdChannel.Enable,
						"name":   osd.OsdChannel.Name,
						"pos":    osd.OsdChannel.Pos,
					},
					"osdTime": map[string]interface{}{
						"enable": osd.OsdTime.Enable,
						"pos":    osd.OsdTime.Pos,
					},
					"watermark": osd.Watermark,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetOsd")

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

		return false, fmt.Errorf("camera could not set osd. camera responded with %v", result.Value)
	}
}
