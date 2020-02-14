package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/eduncan911/podcast"
	"github.com/floge77/cloud2podcast/fileUtils"
	"github.com/floge77/cloud2podcast/model"
	"github.com/gorilla/mux"
)

//ServeAllPodcasts serves all podcasts the app can find in the given downloadDir
func ServeAllPodcasts(router *mux.Router, configYamlPath string, downloadDirectory string, port string) {

	dirs, err := ioutil.ReadDir(downloadDirectory)
	if err != nil {
		log.Printf("Could not read %v. Error: %v", downloadDirectory, err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println("Serving Podcast: " + dir.Name())

			handlerFunc := handleSinglePodcast(configYamlPath, dir.Name(), port, downloadDirectory)
			router.HandleFunc("/podcasts/"+dir.Name(), handlerFunc).Methods("GET")
		}
	}
	router.HandleFunc("/availablePodcasts", handleAvailablePodcasts(dirs)).Methods("GET")
}

func addConfigInfo(podcastInfo *model.Podcast, config *model.ConfigYaml, dir string) {
	for _, podcast := range config.Podcasts {
		if podcast.Channel == dir {
			podcastInfo.ChannelURL = podcast.ChannelURL
			podcastInfo.ChannelImageURL = podcast.ChannelImageURL
			podcastInfo.PlaylistToDownloadURL = podcast.PlaylistToDownloadURL
		}
	}
}

func addAllPodcastItemsToPodcastFeed(podcastInfo *model.Podcast, config *model.ConfigYaml, downloadDirectory string, dir string) {
	var err error
	fileReader := fileUtils.FileInfoExtractor{}
	podcastInfo.Items, err = fileReader.GetPodcastItemsInformationForDir(downloadDirectory + dir)
	if err != nil {
		log.Fatal("Could not read item infos")
	}
}

func getInitializedPodcast(podcastInfo *model.Podcast) *podcast.Podcast {
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

func buildPodcastFeed(podcastInfo *model.Podcast, port string, dir string) *podcast.Podcast {
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

func handleSinglePodcast(configYamlPath, dir, port, downloadDirectory string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		yamlUtil := fileUtils.YamlUtil{}
		config := yamlUtil.GetConfig(configYamlPath)

		podcastInfo := &model.Podcast{
			Channel: dir,
		}
		addConfigInfo(podcastInfo, config, dir)
		addAllPodcastItemsToPodcastFeed(podcastInfo, config, downloadDirectory, dir)
		podcastFeed := buildPodcastFeed(podcastInfo, port, dir)

		w.Header().Set("Content-Type", "application/xml")

		if err := podcastFeed.Encode(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func handleAvailablePodcasts(dirs []os.FileInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var podcastURIs []string
		for _, dir := range dirs {
			if dir.IsDir() {
				podcastURIs = append(podcastURIs, "/"+dir.Name())
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(podcastURIs)
	}
}
