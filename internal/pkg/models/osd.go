package models

import "github.com/ReolinkCameraAPI/reolinkapigo/pkg/enum"

type OsdChannel struct {
	Enable enum.Toggle `json:"enable"`
	Name   string      `json:"name"`
	Pos    string      `json:"pos"`
}

type OsdTime struct {
	Enable enum.Toggle `json:"enable"`
	Pos    string      `json:"pos"`
}

type Osd struct {
	BgColor    enum.Toggle `json:"bgcolor"`
	Channel    int         `json:"channel"`
	OsdChannel OsdChannel  `json:"osdChannel"`
	OsdTime    OsdTime     `json:"osdTime"`
	Watermark  enum.Toggle `json:"watermark"`
}
