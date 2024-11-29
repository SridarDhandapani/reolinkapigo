package api

import (
	"encoding/json"
	"fmt"
	"github.com/ReolinkCameraAPI/reolinkapigo/internal/pkg/models"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/network/rest"
	"github.com/ReolinkCameraAPI/reolinkapigo/pkg/options"
)

type NetworkMixin struct {
}

// SetNetworkPort Set the camera network ports using the NetworkPortOption<prop> functions
// Defaults are automatically for the excluded networkPortOptions
// http: 80
// https: 443
// media: 9000
// onvif: 8000
// rtmp: 1935
// rtsp: 554
// Only http & rtsp ports are enabled by default
func (nm *NetworkMixin) SetNetworkPort(networkPortOptions ...options.NetworkPortOption) func(handler *rest.RestHandler) (bool,
	error) {

	// Defaults
	networkPorts := &models.NetworkPort{
		HttpEnable:  enum.Enabled,
		HttpPort:    80,
		HttpsEnable: enum.Disabled,
		HttpsPort:   443,
		MediaPort:   9000,
		OnvifEnable: enum.Disabled,
		OnvifPort:   8000,
		RtmpEnable:  enum.Disabled,
		RtmpPort:    1935,
		RtspEnable:  enum.Enabled,
		RtspPort:    554,
	}

	for _, op := range networkPortOptions {
		op(networkPorts)
	}

	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetNetPort",
			"action": 0,
			"param": map[string]interface{}{
				"NetPort": map[string]interface{}{
					"httpEnable":  networkPorts.HttpsEnable,
					"httpPort":    networkPorts.HttpPort,
					"httpsEnable": networkPorts.HttpsEnable,
					"httpsPort":   networkPorts.HttpsPort,
					"mediaPort":   networkPorts.MediaPort,
					"onvifEnable": networkPorts.OnvifEnable,
					"onvifPort":   networkPorts.OnvifPort,
					"rtmpEnable":  networkPorts.RtmpEnable,
					"rtmpPort":    networkPorts.RtmpPort,
					"rtspEnable":  networkPorts.RtspEnable,
					"rtspPort":    networkPorts.RtspPort,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetNetPort")

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

		return false, fmt.Errorf("camera could not set network port(s). camera responded with %v", result.Value)
	}
}

// Set the camera's wifi settings
func (nm *NetworkMixin) SetWifi(ssid string, password string) func(handler *rest.RestHandler) (bool, error) {
	return func(handler *rest.RestHandler) (bool, error) {
		payload := map[string]interface{}{
			"cmd":    "SetWifi",
			"action": 0,
			"param": map[string]interface{}{
				"Wifi": map[string]interface{}{
					"ssid":     ssid,
					"password": password,
				},
			},
		}

		result, err := handler.Request("POST", payload, "SetWifi")

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

		return false, fmt.Errorf("camera could not set wifi. camera responded with %v", result.Value)
	}
}

// Get the current camera's wifi settings
func (nm *NetworkMixin) GetWifi() func(handler *rest.RestHandler) (*models.Wifi, error) {
	return func(handler *rest.RestHandler) (*models.Wifi, error) {
		payload := map[string]interface{}{
			"cmd":    "GetWifi",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "GetWifi")

		if err != nil {
			return nil, err
		}

		var wifi *models.Wifi

		err = json.Unmarshal(result.Value["Wifi"], &wifi)

		if err != nil {
			return nil, err
		}

		return wifi, nil
	}
}

// Scan the current camera's wifi network
func (nm *NetworkMixin) ScanWifi() func(handler *rest.RestHandler) (*models.ScanWifi, error) {
	return func(handler *rest.RestHandler) (*models.ScanWifi, error) {
		payload := map[string]interface{}{
			"cmd":    "ScanWifi",
			"action": 1,
			"param":  map[string]interface{}{},
		}

		result, err := handler.Request("POST", payload, "ScanWifi")

		if err != nil {
			return nil, err
		}

		var scanWifi *models.ScanWifi

		err = json.Unmarshal(result.Value["ScanWifi"], &scanWifi)

		if err != nil {
			return nil, err
		}

		return scanWifi, nil
	}
}

// GetNetworkGeneral Get the camera's general network information
func (nm *NetworkMixin) GetNetworkGeneral() func(handler *rest.RestHandler) (*models.NetworkGeneral, error) {
	return func(handler *rest.RestHandler) (*models.NetworkGeneral, error) {
		payload := map[string]interface{}{
			"cmd":    "GetLocalLink",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetLocalLink")

		if err != nil {
			return nil, err
		}

		var networkGeneral *models.NetworkGeneral

		err = json.Unmarshal(resp.Value["LocalLink"], &networkGeneral)

		if err != nil {
			return nil, err
		}

		return networkGeneral, nil
	}
}

// GetNetworkPort Get the camera's network ports status and value
func (nm *NetworkMixin) GetNetworkPort() func(handler *rest.RestHandler) (*models.NetworkPort, error) {
	return func(handler *rest.RestHandler) (*models.NetworkPort, error) {
		payload := map[string]interface{}{
			"cmd":    "GetNetPort",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetNetPort")

		if err != nil {
			return nil, err
		}

		var networkPort *models.NetworkPort

		err = json.Unmarshal(resp.Value["NetPort"], &networkPort)

		if err != nil {
			return nil, err
		}

		return networkPort, nil
	}
}

// Get the camera's network DDNS information
func (nm *NetworkMixin) GetNetworkDDNS() func(handler *rest.RestHandler) (*models.NetworkDDNS, error) {
	return func(handler *rest.RestHandler) (*models.NetworkDDNS, error) {
		payload := map[string]interface{}{
			"cmd":    "GetDdns",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetDdns")

		if err != nil {
			return nil, err
		}

		var networkDdns *models.NetworkDDNS

		err = json.Unmarshal(resp.Value["Ddns"], &networkDdns)

		if err != nil {
			return nil, err
		}

		return networkDdns, nil
	}
}

// Get the camera's network NTP information
func (nm *NetworkMixin) GetNetworkNTP() func(handler *rest.RestHandler) (*models.NetworkNTP, error) {
	return func(handler *rest.RestHandler) (*models.NetworkNTP, error) {
		payload := map[string]interface{}{
			"cmd":    "GetNtp",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetNtp")

		if err != nil {
			return nil, err
		}

		var networkNtp *models.NetworkNTP

		err = json.Unmarshal(resp.Value["Ntp"], &networkNtp)

		if err != nil {
			return nil, err
		}

		return networkNtp, nil
	}
}

// Get the camera's network Email information
func (nm *NetworkMixin) GetNetworkEmail() func(handler *rest.RestHandler) (*models.NetworkEmail, error) {
	return func(handler *rest.RestHandler) (*models.NetworkEmail, error) {
		payload := map[string]interface{}{
			"cmd":    "GetEmail",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetEmail")

		if err != nil {
			return nil, err
		}

		var networkEmail *models.NetworkEmail

		err = json.Unmarshal(resp.Value["Email"], &networkEmail)

		if err != nil {
			return nil, err
		}

		return networkEmail, nil
	}
}

// Get the camera's network FTP information
func (nm *NetworkMixin) GetNetworkFTP() func(handler *rest.RestHandler) (*models.NetworkFTP, error) {
	return func(handler *rest.RestHandler) (*models.NetworkFTP, error) {
		payload := map[string]interface{}{
			"cmd":    "GetFtp",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetFtp")

		if err != nil {
			return nil, err
		}

		var networkFtp *models.NetworkFTP

		err = json.Unmarshal(resp.Value["Ftp"], &networkFtp)

		if err != nil {
			return nil, err
		}

		return networkFtp, nil
	}
}

// Get the camera's network Push information
func (nm *NetworkMixin) GetNetworkPush() func(handler *rest.RestHandler) (*models.NetworkPush, error) {
	return func(handler *rest.RestHandler) (*models.NetworkPush, error) {
		payload := map[string]interface{}{
			"cmd":    "GetPush",
			"action": 0,
			"param":  map[string]interface{}{},
		}

		resp, err := handler.Request("POST", payload, "GetPush")

		if err != nil {
			return nil, err
		}

		var networkPush *models.NetworkPush

		err = json.Unmarshal(resp.Value["Push"], &networkPush)

		if err != nil {
			return nil, err
		}

		return networkPush, nil
	}
}

// Get the camera's network Status information is just a wrapper for networkGeneral
// TODO: revise this, exactly copied from the reolink-python-api project.
func (nm *NetworkMixin) GetNetworkStatus() func(handler *rest.RestHandler) (*models.NetworkGeneral, error) {
	return func(handler *rest.RestHandler) (*models.NetworkGeneral, error) {
		return nm.GetNetworkGeneral()(handler)
	}
}
