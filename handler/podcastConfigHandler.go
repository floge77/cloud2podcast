package handler

import (
	"encoding/json"
	"net/http"

	"github.com/floge77/cloud2podcast/fileUtils"
	"github.com/floge77/cloud2podcast/model"
	"github.com/gorilla/mux"
)

// ServePodcastInfo serves all configured podcasts
func ServePodcastInfo(router *mux.Router, configYamlPath string) {
	handleGetAllPodcastInfo(router, configYamlPath)
	handleAddPodcastInfo(router, configYamlPath)
}

func handleGetAllPodcastInfo(router *mux.Router, configYamlPath string) {
	router.HandleFunc("/podcasts", func(w http.ResponseWriter, r *http.Request) {
		yamlUtil := fileUtils.YamlUtil{}
		config := yamlUtil.GetConfig(configYamlPath)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config.Podcasts)
	}).Methods("GET")
}

// func handleUpdatePodcastInfo(router *mux.Router, configYamlPath string) {
// 	router.HandleFunc("/podcasts", func(w http.ResponseWriter, r *http.Request) {
// 		YamlUtil := fileUtils.YamlUtil{}
// 		config := YamlUtil.GetConfig(configYamlPath)
// 		w.WriteHeader(http.StatusOK)
// 	}).Methods("PUT")
// }

// curl --header "Content-Type: application/json" --request POST --data '{"Channel": "MyChannel", "ChannelImageURL": "MyChannel.jpq", "ChannelURL": "mixcloud.com/MyChannel", "PlaylistToDownloadURL": "MyChannel.com"}' 127.0.0.1:8000/podcasts
func handleAddPodcastInfo(router *mux.Router, configYamlPath string) {
	router.HandleFunc("/podcasts", func(w http.ResponseWriter, r *http.Request) {
		yamlUtil := fileUtils.YamlUtil{}
		configYaml := yamlUtil.GetConfig(configYamlPath)
		decoder := json.NewDecoder(r.Body)
		var podcast model.Podcast

		if err := decoder.Decode(&podcast); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid ")
			return
		}
		if podcast.Channel == "" || podcast.ChannelURL == "" || podcast.PlaylistToDownloadURL == "" || podcast.ChannelImageURL == "" {
			RespondWithError(w, http.StatusBadRequest, "Invalid Request")
			return
		}

		defer r.Body.Close()
		configYaml.Podcasts = append(configYaml.Podcasts, &podcast)
		yamlUtil.WriteConfig(configYaml, configYamlPath)

		w.WriteHeader(http.StatusOK)
	}).Methods("POST")
}

// func handleDeletePodcastInfo(router *mux.Router, configYamlPath string) {
// 	router.HandleFunc("/podcasts", func(w http.ResponseWriter, r *http.Request) {
// 		YamlUtil := fileUtils.YamlUtil{}
// 		config := YamlUtil.GetConfig(configYamlPath)
// 		w.WriteHeader(http.StatusOK)
// 	}).Methods("DELETE")
// }
