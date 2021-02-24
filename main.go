package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/christophermanning/go-examples/ffmpegapi"
)

var portFlag = flag.Int("p", 8000, "port")

func main() {
	flag.Parse()

	http.HandleFunc("/ffmpeg-api/apply-effects", ffmpegapi.ApplyEffectsHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), nil))
}
