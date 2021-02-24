package ffmpegapi

type ApplyEffectRequest struct {
	Video      *Video      `json:"video"`
	TextEffect *TextEffect `json:"text_effect"`
	// NOTE could add more effects, such as a video overlay, here
}

type ApplyEffectResponse struct {
	Cmd   string `json:"cmd"`
	Error string `json:"error"`
}

type Video struct {
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
	Duration   string `json:"duration"`
	Resolution string `json:"resolution"`
}

type TextEffect struct {
	Text      string `json:"text"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	FontSize  int    `json:"font_size"`
	FontColor string `json:"font_color"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
