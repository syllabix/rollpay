package media

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/disintegration/imaging"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
)

const maxDimen = 150

var ErrImageProcessFailed = errors.New("an error occurred while processing the image")

// Process takes the provided file and smart crops it within a 150x150 pixel square,
// and encodes it to a jpeg
func Process(ctx context.Context, file io.Reader) ([]byte, error) {
	img, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		return failure(err)
	}

	analyser := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	crop, err := analyser.FindBestCrop(img, maxDimen, maxDimen)
	if err != nil {
		return failure(err)
	}

	img = imaging.Crop(img, crop)
	img = imaging.Fill(img, maxDimen, maxDimen, imaging.Center, imaging.Lanczos)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, imaging.JPEG, imaging.JPEGQuality(75))
	if err != nil {
		return failure(err)
	}

	return buf.Bytes(), nil
}

func failure(reason error) ([]byte, error) {
	return nil, fmt.Errorf("%w: %v", ErrImageProcessFailed, reason)
}
