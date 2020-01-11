package model

import "time"

type PodcastItem struct {
	Title, Channel, FileName string
	FileSize                 int64
	ReleaseDate              *time.Time
}

type PodcastInfo struct {
	Channel               string `yaml:"channelName"`
	ChannelURL            string `yaml:"channelURL"`
	ChannelImageURL       string `yaml:"channelImageURL"`
	PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
	Provider              string
	DownloadDirectory     string
	Items                 []*PodcastItem
}

type Podcast struct {
	Channel               string `yaml:"channelName"`
	ChannelURL            string `yaml:"channelURL"`
	ChannelImageURL       string `yaml:"channelImageURL"`
	PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
	Items                 []*PodcastItem
}

type PodcastConfigYaml struct {
	MinLength         int `yaml:"minLength"`
	DownloadDirectory string
	Port              string
	Podcasts          []*Podcast `yaml:"podcasts"`
}
