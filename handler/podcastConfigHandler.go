package handler

import (
	"encoding/json"
	"net/http"

	"github.com/floge77/cloud2podcast/fileUtils"
	"github.com/floge77/cloud2podcast/model"
	"github.com/gorilla/mux"
)

func ServePodcastInfo(router *mux.Router, configYamlPath string) {
	yamlReader := fileUtils.YamlReader{}
	config := yamlReader.GetConfig(configYamlPath)
	handleGetAllPodcastInfo(router, config)
}

func handleGetAllPodcastInfo(router *mux.Router, config *model.PodcastConfigYaml) {
	router.HandleFunc("/podcasts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(config.Podcasts)
	}).Methods("GET")
}

func handleUpdatePodcastInfo() {

}

func handleAddPodcastInfo() {

}
