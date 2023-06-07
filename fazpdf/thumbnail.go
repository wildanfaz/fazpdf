package fazpdf

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"os/exec"

	"github.com/nfnt/resize"
)

type ThumbnailResolution struct {
	Height uint
	Width  uint
}

func GenerateThumbnail(fc FileConfig) error {
	cmd := exec.Command("pdftoppm", "-jpeg", "-f", "1", "-l", "1", fmt.Sprintf("./result/%s.pdf", fc.ResultName), fmt.Sprintf("./result/%s", fc.ResultName))

	err := cmd.Run()

	if err != nil {
		return err
	}

	file, err := os.Open(fmt.Sprintf("./result/%s-1.jpg", fc.ResultName))
	defer file.Close()

	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	img = resize.Resize(fc.ThumbnailResolution.Width, fc.ThumbnailResolution.Height, img, resize.Lanczos3)

	for quality := 100; quality >= 1; quality-- {
		w, err := os.Create(fmt.Sprintf("./result/%s-1.jpg", fc.ResultName))
		defer w.Close()

		if err != nil {
			return err
		}

		err = jpeg.Encode(w, img, &jpeg.Options{Quality: quality})

		if err != nil {
			return err
		}

		fs, err := w.Stat()

		if err != nil {
			return err
		}

		if fs.Size() < 100*980 {
			err = os.Rename(fmt.Sprintf("./result/%s-1.jpg", fc.ResultName), fmt.Sprintf("./result/thumbnail_%s-1.jpg", fc.ResultName))

			if err != nil {
				return err
			}

			break
		}
	}

	return nil
}
