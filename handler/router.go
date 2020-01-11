package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/floge77/cloud2podcastnew/reader"
	"github.com/gorilla/mux"
)

func Run() {

	yamlReader := reader.YamlReader{}
	var configYamlPath string
	configYamlPath = os.Getenv("configYaml")
	if configYamlPath == "" {
		configYamlPath = "/downloads/config.yaml"
	}

	config := yamlReader.GetConfig(configYamlPath)
	downloadDirectory := os.Getenv("downloadDir")
	if downloadDirectory == "" {
		downloadDirectory = "/downloads/"
	}
	config.DownloadDirectory = downloadDirectory + "/"
	fmt.Println(downloadDirectory)
	fmt.Println(config.Podcasts[0])

	router := mux.NewRouter()
	port := "8080"

	config.Port = port

	router.PathPrefix("/downloads/").Handler(http.StripPrefix("/downloads/", http.FileServer(http.Dir(downloadDirectory+"/"))))
	ServeAllPodcasts(router, config)

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Router running at Port " + port)
	log.Fatal(server.ListenAndServe())

}
