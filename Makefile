run: test
	go run main.go -p 8000

curl:
	curl http://localhost:8000/ffmpeg-api/apply-effects -X POST -d '{"video": {"input_path": "video.mp4", "output_path": "video_out.mp4", "duration": "10.0", "resolution": "1920x1080"}, "text_effect": {"text":"hello", "x":10, "y":10, "font_size":10, "font_color":"0x000000", "start_time":"0.0", "end_time":"5.0" }}' -k -H "Content-Type: application/json"

test:
	go test ./...
