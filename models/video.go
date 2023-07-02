package models

//import (
//	"fmt"
//	"google.golang.org/api/googleapi/transport"
//	"log"
//	"net/http"
//)
//
//func youtube() {
//	client := &http.Client{
//		Transport: &transport.APIKey{Key: developerKey},
//	}
//
//	service, err := youtube.New(client)
//	if err != nil {
//		log.Fatalf("Error creating new YouTube client: %v", err)
//	}
//
//	// Make the API call to YouTube.
//	call := service.Search.List("id,snippet").
//		Q(*query).
//		MaxResults(*maxResults)
//	response, err := call.Do()
//	handleError(err, "")
//
//	// Group video, channel, and playlist results in separate lists.
//	videos := make(map[string]string)
//	channels := make(map[string]string)
//	playlists := make(map[string]string)
//
//	// Iterate through each item and add it to the correct list.
//	for _, item := range response.Items {
//		switch item.Id.Kind {
//		case "youtube#video":
//			videos[item.Id.VideoId] = item.Snippet.Title
//		case "youtube#channel":
//			channels[item.Id.ChannelId] = item.Snippet.Title
//		case "youtube#playlist":
//			playlists[item.Id.PlaylistId] = item.Snippet.Title
//		}
//	}
//
//	printIDs("Videos", videos)
//	printIDs("Channels", channels)
//	printIDs("Playlists", playlists)
//}
//
//func printIDs(sectionName string, matches map[string]string) {
//	fmt.Printf("%v:\n", sectionName)
//	for id, title := range matches {
//		fmt.Printf("[%v] %v\n", id, title)
//	}
//	fmt.Printf("\n\n")
//}
