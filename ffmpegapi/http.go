package ffmpegapi

import (
	"encoding/json"
	"log"
	"net/http"
)

// ApplyEffectsHandler This is an example of how to convert a JSON request into a system command and return it in a JSON response.
var ApplyEffectsHandler = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST is supported", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var request *ApplyEffectRequest
	response := &ApplyEffectResponse{}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	encoder := json.NewEncoder(w)

	err := decoder.Decode(&request)

	// NOTE could add much more user-friendly logging here
	if err != nil || request.Video == nil || request.TextEffect == nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Default().Print(err.Error())

		response.Error = "invalid JSON request"
		encoder.Encode(response)
		return
	}

	result, err := request.Video.ApplyTextEffect(request.TextEffect)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = err.Error()
	} else {
		response.Cmd = result
	}

	encoder.Encode(response)
}
