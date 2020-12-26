package main 

import (
	//"io"
	"os"
	"net/http"
	//"html/template"
	//"io/ioutil"
	//"log"
	"time"
	"github.com/julienschmidt/httprouter"
)

//Streaming handler: formatting the video into binary bit stream, and use stream pass it to client side
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//Get video id for video identifier in streaming
	vid := p.ByName("vid-id")
	//Read video path from defs.go
	vl := VIDEO_DIR + vid

	//Open video
	video, err := os.Open(vl)

	//If fail to open video, throw 500 
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	//Convert video type to mp4, so the browser at client side will be able to parse the content
	//After parsing video, browser will be able to combine them to a video
	w.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent would be able to pass the video file as binary stream to client(browser)
	//client side browser will be able to parse the content and play the video
	//(w: response, r: request, name: "", time, videoFile)
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
}