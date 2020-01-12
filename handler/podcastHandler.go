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
	"github.com/floge77/cloud2podcastnew/fileUtils"
	"github.com/floge77/cloud2podcastnew/model"
	"github.com/gorilla/mux"
)

func ServeAllPodcasts(router *mux.Router, configYamlPath string, downloadDirectory string, port string) {
	yamlReader := fileUtils.YamlReader{}
	config := yamlReader.GetConfig(configYamlPath)

	dirs, err := ioutil.ReadDir(downloadDirectory)
	if err != nil {
		log.Fatalf("Could not read %v. Error: %v", downloadDirectory, err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("Serving Podcast: " + dir.Name())

			podcastInfo := &model.PodcastInfo{
				Channel: dir.Name(),
			}
			addConfigInfo(podcastInfo, config, dir.Name())
			addAllPodcastItemsToPodcastFeed(podcastInfo, config, downloadDirectory, dir.Name())
			podcastFeed := buildPodcastFeed(podcastInfo, port, dir.Name())

			handlerFunc := handleSinglePodcast(podcastFeed)
			router.HandleFunc("/"+dir.Name(), handlerFunc).Methods("GET")
		}

	}

}

func addConfigInfo(podcastInfo *model.PodcastInfo, config *model.PodcastConfigYaml, dir string) {
	for _, podcast := range config.Podcasts {
		if podcast.Channel == dir {
			podcastInfo.ChannelURL = podcast.ChannelURL
			podcastInfo.ChannelImageURL = podcast.ChannelImageURL
			podcastInfo.PlaylistToDownloadURL = podcast.PlaylistToDownloadURL
		}
	}
}

func addAllPodcastItemsToPodcastFeed(podcastInfo *model.PodcastInfo, config *model.PodcastConfigYaml, downloadDirectory string, dir string) {
	var err error
	fileReader := fileUtils.FileInfoExtractor{}
	podcastInfo.Items, err = fileReader.GetPodcastItemsInformationForDir(downloadDirectory + dir)
	if err != nil {
		log.Fatal("Could not read item infos")
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

func buildPodcastFeed(podcastInfo *model.PodcastInfo, port string, dir string) *podcast.Podcast {
	podcastToServe := getInitializedPodcast(podcastInfo)
	// hostName := os.Hostname()
	hostIP := os.Getenv("HOST_IP")
	for _, item := range podcastInfo.Items {
		downloadURL := &url.URL{
			Scheme: "http",
			Host:   hostIP + ":" + port,
			Path:   path.Join("downloads", dir, item.FileName),
		}

		// "http://" + hostIP + ":" + config.Port + dir.Name() + "/"
		appendPodcastItem(podcastToServe, item, downloadURL)
	}
	return podcastToServe
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
