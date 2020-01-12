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
	} else if port == "80" {
		port = ""
	}

	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("frontend/"))))
	router.PathPrefix("/downloads/").Handler(http.StripPrefix("/downloads/", http.FileServer(http.Dir(downloadDirectory+"/"))))
	ServeAllPodcasts(router, configYamlPath, downloadDirectory+"/", port)
	ServePodcastInfo(router, configYamlPath)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	server := &http.Server{
		Handler: router,
		Addr:    ":" + "8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Router running at Port " + port)
	log.Fatal(server.ListenAndServe())
}
