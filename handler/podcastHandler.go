package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/eduncan911/podcast"
	"github.com/floge77/cloud2podcastnew/model"
	"github.com/floge77/cloud2podcastnew/reader"
	"github.com/gorilla/mux"
)

func ServeAllPodcasts(router *mux.Router, config *model.PodcastConfigYaml) {
	dirs, err := ioutil.ReadDir(config.DownloadDirectory)
	if err != nil {
		log.Fatalf("Could not read %v. Error: %v", config.DownloadDirectory, err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("Serving Podcast: " + dir.Name())

			podcastInfo := &model.PodcastInfo{
				Channel: dir.Name(),
			}

			for _, podcast := range config.Podcasts {
				if podcast.Channel == dir.Name() {
					podcastInfo.ChannelURL = podcast.ChannelURL
					podcastInfo.ChannelImageURL = podcast.ChannelImageURL
					podcastInfo.PlaylistToDownloadURL = podcast.PlaylistToDownloadURL
				}
			}
			fileReader := reader.FileInfoExtractor{}
			podcastInfo.Items, err = fileReader.GetPodcastItemsInformationForDir(config.DownloadDirectory + dir.Name())
			if err != nil {
				log.Fatal("Could not read item infos")
			}
			fmt.Println(podcastInfo)

			podcastToServe := getInitializedPodcast(podcastInfo)
			// hostName := os.Hostname()
			hostIP := os.Getenv("HOST_IP")
			for _, item := range podcastInfo.Items {
				downloadURL := &url.URL{
					Scheme: "http",
					Host:   hostIP + ":" + config.Port,
					Path:   path.Join("downloads", dir.Name(), item.FileName),
				}

				// "http://" + hostIP + ":" + config.Port + dir.Name() + "/"
				appendPodcastItem(podcastToServe, item, downloadURL)
			}
			handlerFunc := handleSinglePodcast(podcastToServe)
			router.HandleFunc("/"+dir.Name(), handlerFunc).Methods("GET")
			fmt.Println(podcastToServe)
		}

	}

}

func getInitializedPodcast(podcastInfo *model.PodcastInfo) *podcast.Podcast {
	channel := podcastInfo.Channel
	imageURL := podcastInfo.ChannelImageURL
	title := channel + "-Podcast"

	p := podcast.New(
		title,
		podcastInfo.ChannelURL,
		"",
		nil, nil,
	)

	p.ISubtitle = title
	p.AddSummary("Podcast from " + channel + " channel")
	p.AddImage(imageURL)
	p.AddAuthor(channel, channel+"@email.com")

	return &p
}

func appendPodcastItem(podcastToAppend *podcast.Podcast, itemToAdd *model.PodcastItem, downloadURL *url.URL) {

	title := itemToAdd.Title
	channel := itemToAdd.Channel
	releaseDate := itemToAdd.ReleaseDate
	fileSize := itemToAdd.FileSize

	item := podcast.Item{
		Title:       title,
		Description: channel,
		ISubtitle:   "",
		PubDate:     releaseDate,
	}
	item.AddSummary(title)
	item.AddEnclosure(downloadURL.String(), podcast.MP3, fileSize)

	_, err := podcastToAppend.AddItem(item)
	if err != nil {
		fmt.Println(item.Title, ": error", err.Error())
	}
}

func handleSinglePodcast(podcast *podcast.Podcast) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")

		if err := podcast.Encode(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
