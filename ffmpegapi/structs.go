package ffmpegapi

type Video struct {
	inputPath  string
	outputPath string
	duration   string
	resolution string
}

type TextEffect struct {
	text      string
	x         int
	y         int
	fontSize  int
	fontColor string
	startTime string
	endTime   string
}
