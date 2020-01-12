package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func Run() {

	configYamlPath := os.Getenv("configYaml")
	if configYamlPath == "" {
		configYamlPath = "/downloads/config.yaml"
	}

	downloadDirectory := os.Getenv("downloadDir")
	if downloadDirectory == "" {
		downloadDirectory = "/downloads/"
	}
	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	// router.Handle("/public", http.FileServer(http.Dir("frontend/")))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("frontend/"))))
	router.PathPrefix("/downloads/").Handler(http.StripPrefix("/downloads/", http.FileServer(http.Dir(downloadDirectory+"/"))))
	ServeAllPodcasts(router, configYamlPath, downloadDirectory+"/", port)
	ServePodcastInfo(router, configYamlPath)

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
