package model

import "time"

type PodcastItem struct {
	Title, Channel, FileName string
	FileSize                 int64
	ReleaseDate              *time.Time
}

type Podcastinfo struct {
	Channel               string `yaml:"channelName"`
	ChannelURL            string `yaml:"channelURL"`
	ChannelImageURL       string `yaml:"channelImageURL"`
	PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
	Provider              string
	DownloadDirectory     string
	Items                 []*PodcastItem
}

type PodcastConfigYaml struct {
	minLength int `yaml:"minLength"`
	Podcasts  []struct {
		Channel               string `yaml:"channelName"`
		ChannelURL            string `yaml:"channelURL"`
		ChannelImageURL       string `yaml:"channelImageURL"`
		PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
		Items                 []*PodcastItem
	} `yaml:"podcasts"`
}
