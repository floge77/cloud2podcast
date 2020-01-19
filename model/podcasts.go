package model

import "time"

// PodcastItem represents a single Podcastitem
type PodcastItem struct {
	Title, Channel, FileName string
	FileSize                 int64
	ReleaseDate              *time.Time
}

// Podcast represents a complete Podcastfeed
type Podcast struct {
	Channel               string `yaml:"channelName"`
	ChannelURL            string `yaml:"channelURL"`
	ChannelImageURL       string `yaml:"channelImageURL"`
	PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
	Items                 []*PodcastItem
}


type ConfigYaml struct {
	MinLength int        `yaml:"minLength"`
	Podcasts  []*Podcast `yaml:"podcasts"`
}
