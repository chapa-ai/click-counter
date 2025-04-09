package model

import "time"

type Click struct {
	BannerID  int       `json:"bannerID"`
	Timestamp time.Time `json:"timestamp"`
	TsFrom    time.Time `json:"tsFrom"`
	TsTo      time.Time `json:"tsTo"`
}

type Stats struct {
	BannerID  int       `json:"bannerID"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

type StatResponse struct {
	Timestamp string `json:"ts"`
	Value     int    `json:"v"`
}

type StatsResponse struct {
	Stats []StatResponse `json:"stats"`
}
