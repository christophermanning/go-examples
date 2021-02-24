package ffmpegapi

import (
	"errors"
	"testing"
)

func TestParseResolution(t *testing.T) {
	tests := []struct {
		label      string
		resolution string
		x          int
		y          int
	}{
		{
			"works",
			"1920x1080",
			1920,
			1080,
		},
	}

	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			x, y, err := (&Video{Resolution: tc.resolution}).ParseResolution()
			if err != nil {
				t.Errorf(err.Error())
			}
			if tc.x != x || tc.y != y {
				t.Errorf("incorrect output: %d %d expected: %d %d", x, y, tc.x, tc.y)
			}
		})
	}
}

func TestApplyTextEffect(t *testing.T) {
	tests := []struct {
		label       string
		video       *Video
		textEffect  *TextEffect
		expected    string
		expectedErr error
	}{
		{
			"error: no input video",
			&Video{},
			&TextEffect{},
			"",
			ErrVideo{Err: errors.New("Invalid Input Path")},
		},
		{
			"test1: basic output",
			&Video{
				InputPath:  "test_input1.mp4",
				OutputPath: "test_output1.mp4",
				Duration:   "60.0",
				Resolution: "1920x1080",
			},
			&TextEffect{
				Text:      "I’m sOoOo good at this game! xD",
				X:         200,
				Y:         100,
				FontSize:  64,
				FontColor: "0xFFFFFF",
				StartTime: "23.0",
				EndTime:   "40.0",
			},
			`ffmpeg -i test_input1.mp4 -vf drawtext="enable='between(t,23.0,40.0)':text='I’m sOoOo good at this game! xD':fontcolor=0xFFFFFF:fontsize=64:x=200:y=100" test_output1.mp4`,
			nil,
		},
		{
			"test2: basic output2",
			&Video{
				InputPath:  "test_input2.mp4",
				OutputPath: "test_output2.mp4",
				Duration:   "60.0",
				Resolution: "1920x1080",
			},
			&TextEffect{
				Text:      "Brutal, Savage, Rekt",
				X:         0,
				Y:         0,
				FontSize:  48,
				FontColor: "0x000000",
				StartTime: "0.0",
				EndTime:   "60.0",
			},
			`ffmpeg -i test_input2.mp4 -vf drawtext="enable='between(t,0.0,60.0)':text='Brutal, Savage, Rekt':fontcolor=0x000000:fontsize=48:x=0:y=0" test_output2.mp4`,
			nil,
		},
		{
			"test3: error invalid end time",
			&Video{
				InputPath:  "test_input3.mp4",
				OutputPath: "test_output3.mp4",
				Duration:   "60.0",
				Resolution: "1920x1080",
			},
			&TextEffect{
				Text:      "RIP",
				X:         100,
				Y:         200,
				FontSize:  32,
				FontColor: "0xFFFFFF",
				StartTime: "24.0",
				EndTime:   "75.0",
			},
			``,
			ErrVideo{Err: errors.New("Invalid End Time")},
		},
		{
			"test4: error invalid X,Y",
			&Video{
				InputPath:  "test_input4.mp4",
				OutputPath: "test_output4.mp4",
				Duration:   "60.0",
				Resolution: "1920x1080",
			},
			&TextEffect{
				Text:      "RIP",
				X:         100,
				Y:         9999,
				FontSize:  48,
				FontColor: "0x777777",
				StartTime: "24.0",
				EndTime:   "48.0",
			},
			``,
			ErrTextEffect{Err: errors.New("Invalid X,Y coordinate")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			output, err := tc.video.ApplyTextEffect(tc.textEffect)
			if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("incorrect error message: %s expected: %s", err.Error(), tc.expectedErr.Error())
			}

			// errors.As() overwrites tc.expectedErr, so do this after checking .Error()
			if (err != nil && !errors.As(err, &tc.expectedErr)) || (err == nil && tc.expectedErr != nil) {
				t.Errorf("incorrect error: %s expected: %s", err, tc.expectedErr)
			}

			if tc.expected != output {
				t.Errorf("incorrect output: %s expected: %s", output, tc.expected)
			}
		})
	}
}
