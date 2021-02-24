package ffmpegapi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrVideo custom error to distinguish between error types
// alternatively `var ErrVideo = errors.New("top level message")` could be used
// if error wrapping was needed
type ErrVideo struct {
	Err error
}

func (e ErrVideo) Error() string { return e.Err.Error() }

type ErrTextEffect struct {
	Err error
}

func (e ErrTextEffect) Error() string { return e.Err.Error() }

func (v Video) ParseResolution() (int, int, error) {
	// TODO add more robust parsing and error handling
	parts := strings.Split(v.Resolution, "x")
	if len(parts) != 2 {
		return -1, -1, errors.New("invalid format")
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return -1, -1, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return -1, -1, err
	}

	return x, y, nil
}

func (v *Video) ApplyTextEffect(t *TextEffect) (string, error) {
	if v.InputPath == "" {
		// NOTE could wrap this error with `fmt.Errorf("%w: Invalid Input Path")`
		// but the returned error message would include the top level error string
		return "", ErrVideo{Err: errors.New("Invalid Input Path")}
	}

	tEndTime, err := strconv.ParseFloat(t.EndTime, 64)
	if err != nil {
		return "", ErrTextEffect{Err: errors.New("Invalid End Time")}
	}

	vDuration, err := strconv.ParseFloat(v.Duration, 64)
	if err != nil {
		return "", ErrVideo{Err: errors.New("Invalid Input Path")}
	}

	if tEndTime > vDuration {
		return "", ErrTextEffect{Err: errors.New("Invalid End Time")}
	}

	x, y, err := v.ParseResolution()
	if err != nil {
		return "", ErrTextEffect{Err: errors.New("Invalid Resolution")}
	}

	if t.X > x || t.Y > y || t.X < 0 || t.Y < 0 {
		return "", ErrTextEffect{Err: errors.New("Invalid X,Y coordinate")}
	}

	// the drawtext value should have no spaces in it
	return fmt.Sprintf(`ffmpeg -i %s -vf drawtext="enable='between(t,%s,%s)':text='%s':fontcolor=%s:fontsize=%d:x=%d:y=%d" %s`,
		v.InputPath,
		t.StartTime,
		t.EndTime,
		t.Text,
		t.FontColor,
		t.FontSize,
		t.X,
		t.Y,
		v.OutputPath,
	), nil
}
