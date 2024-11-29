package models

import "github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"

// TODO: update with its actual data structure
type ScanWifi struct {
}

// TODO: update with its actual data structure
type Wifi struct {
}

type NetworkGeneralDns struct {
	Auto int    `json:"auto"`
	Dns1 string `json:"dns1"`
	Dns2 string `json:"dns2"`
}

type NetworkGeneralStatic struct {
	Gateway string `json:"gateway"`
	Ip      string `json:"ip"`
	Mask    string `json:"mask"`
}

type NetworkGeneral struct {
	ActiveLink string               `json:"activeLink"`
	Dns        NetworkGeneralDns    `json:"dns"`
	Mac        string               `json:"mac"`
	Static     NetworkGeneralStatic `json:"static"`
	Type       string               `json:"type"`
}

type NetworkDDNS struct {
	Domain   string `json:"domain"`
	Enable   bool   `json:"enable"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Username string `json:"userName"`
}

type NetworkNTP struct {
	Enable   bool   `json:"enable"`
	Interval int    `json:"interval"`
	Port     int    `json:"port"`
	Server   string `json:"server"`
}

type NetworkEmail struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Addr1      string   `json:"addr1"`
	Addr2      string   `json:"addr2"`
	Addr3      string   `json:"addr3"`
	Attachment string   `json:"attachment"`
	Interval   string   `json:"interval"`
	Nickname   string   `json:"nickName"`
	Schedule   Schedule `json:"schedule"`
	SmtpPort   int      `json:"smtpPort"`
	SmtpServer string   `json:"smtpServer"`
	SSL        bool     `json:"ssl"`
}

type NetworkFTP struct {
	Username   string   `json:"userName"`
	Password   string   `json:"password"`
	Anonymous  bool     `json:"anonymous"`
	Interval   int      `json:"interval"`
	MaxSize    int      `json:"maxSize"`
	Mode       int      `json:"mode"`
	Port       int      `json:"port"`
	RemoteDir  string   `json:"remoteDir"`
	Schedule   Schedule `json:"schedule"`
	Server     string   `json:"server"`
	StreamType int      `json:"streamType"`
}

type NetworkPush struct {
	Schedule Schedule `json:"schedule"`
}

type NetworkPort struct {
	HttpEnable  enum.Toggle `json:"httpEnable"`
	HttpPort    int         `json:"httpPort"`
	HttpsEnable enum.Toggle `json:"httpsEnable"`
	HttpsPort   int         `json:"httpsPort"`
	MediaPort   int         `json:"mediaPort"`
	OnvifEnable enum.Toggle `json:"onvifEnable"`
	OnvifPort   int         `json:"onvifPort"`
	RtmpEnable  enum.Toggle `json:"rtmpEnable"`
	RtmpPort    int         `json:"rtmpPort"`
	RtspEnable  enum.Toggle `json:"rtspEnable"`
	RtspPort    int         `json:"rtspPort"`
}
