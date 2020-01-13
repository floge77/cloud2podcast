package fileUtils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/floge77/cloud2podcast/model"
)

type FileInfoExtractor struct {
}

func (f *FileInfoExtractor) GetPodcastItemsInformationForDir(dir string) (itemInfos []*model.PodcastItem, err error) {

	fileNames, err := f.readDir(dir)
	if err != nil {
		return nil, err
	}
	for _, name := range fileNames {
		// create Info structs for podcastsItems
		if strings.Contains(name, ".mp3") {
			item := f.getPodcastItemInfosFromFileName(dir, name)
			itemInfos = append(itemInfos, item)
		}
	}
	return
}

func (f *FileInfoExtractor) getPodcastItemInfosFromFileName(dir string, filename string) (item *model.PodcastItem) {

	s := strings.Replace(filename, ".mp3", "", -1)
	fields := strings.Split(s, "__")
	item = &model.PodcastItem{}
	item.Title = fields[0]
	item.Channel = fields[1]
	item.ReleaseDate = f.getReleaseDateFromString(fields[2])
	fileSize, _ := f.extractFileSize(dir, filename)
	item.FileSize = fileSize
	item.FileName = filename
	return
}

func (*FileInfoExtractor) extractFileSize(dir string, filename string) (fileSize int64, err error) {
	file, err := os.Stat(dir + "/" + filename)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
}

func (*FileInfoExtractor) getReleaseDateFromString(date string) *time.Time {
	t, _ := time.Parse("20060102", date)
	return &t
}

func (*FileInfoExtractor) readDir(dirname string) (list []string, err error) {
	file, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	list, err = file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		fmt.Printf("Could not read directory %v", dirname)
	}
	return
}
