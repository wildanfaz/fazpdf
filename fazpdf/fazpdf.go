package fazpdf

import (
	"errors"
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type FileConfig struct {
	Src                 string
	TargetPath          string
	ResultName          string
	WatermarkPath       string
	Password            string
	EncryptedPdf        bool
	UpdateMetadata      bool
	GenerateThumbnail   bool
	WatermarkPosition   int
	Dx                  int
	Dy                  int
	Metadata            Metadata
	ThumbnailResolution ThumbnailResolution
}

/*
Src : source input file,
TargetPath : target path for result,
ResultName : name for result file,
WatermarkPath : path watermark for qr,
Password : if encrypted true password is used for decrypt,
EncryptedPdf : is need to decrypt or not,
UpdateMetadata : is need to update metadata or not,
GenerateThumbnail : is need to generate thumbnail or not,
WatermarkPosition : 1 = top-left, 2 = top-right, 3 = bottom-right, 4 = bottom-left,
Dx : offset horizontal,
Dy : offset vertical,
Metadata : change the metadata,
ThumbnailResolution : set the resolution of the generated thumbnail
*/
func (fc FileConfig) NewPDF() error {
	var (
		conf = model.NewDefaultConfiguration()
		imp  = pdfcpu.DefaultImportConfig()
		desc string
	)

	if fc.WatermarkPosition > 4 {
		return errors.New("maximum watermark position is 4")
	}

	switch fc.WatermarkPosition {
	case 1:
		desc = fmt.Sprintf("scale:0.15, rotation:0, position:tl, offset:%d %d", fc.Dx, fc.Dy)
	case 2:
		desc = fmt.Sprintf("scale:0.15, rotation:0, position:tr, offset:%d %d", fc.Dx, fc.Dy)
	case 3:
		desc = fmt.Sprintf("scale:0.15, rotation:0, position:br, offset:%d %d", fc.Dx, fc.Dy)
	case 4:
		desc = fmt.Sprintf("scale:0.15, rotation:0, position:bl, offset:%d %d", fc.Dx, fc.Dy)
	}

	CheckFolder()

	name, err := CheckFilename(fc.ResultName)

	if err != nil {
		return err
	}

	fc.ResultName = name

	uuid, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	file, err := os.Open(fc.Src)
	defer file.Close()

	if err != nil {
		return err
	}

	stat, err := file.Stat()

	if err != nil {
		return err
	}

	filename := stat.Name()

	err = NewQR(uuid.String(), fc.WatermarkPath)

	if err != nil {
		return err
	}

	checkQr, err := os.Open(fmt.Sprintf("./qr/%s.png", uuid))
	defer checkQr.Close()

	if err != nil {
		return err
	}

	check, _, err := image.DecodeConfig(checkQr)

	if err != nil {
		return err
	}

	fmt.Println("Input File :", filename)
	fmt.Println(fmt.Sprintf("QR Heigh : %d, QR Width : %d, Scale Into PDF 0.15 : 123", check.Height, check.Width))

	// Process PDF
	if strings.Contains(filename[len(filename)-5:], ".pdf") {
		if fc.EncryptedPdf {
			err = ProcessEncryptedPdf(fc.Src, uuid.String(), fc.ResultName, fc.Password, desc, fc, conf)

			if err != nil {
				return err
			}

			if fc.GenerateThumbnail {
				err = GenerateThumbnail(fc)

				if err != nil {
					return err
				}
			}

			return nil
		}

		err = ProcessPdf(uuid.String(), desc, fc, conf)

		if err != nil {
			return err
		}

		if fc.GenerateThumbnail {
			err = GenerateThumbnail(fc)

			if err != nil {
				return err
			}
		}

		return nil
	}

	// Process Image PNG or JPG
	if strings.Contains(filename[len(filename)-5:], ".jpg") || strings.Contains(filename[len(filename)-5:], ".png") {
		err := ProcessImage(uuid.String(), desc, fc, conf, imp)

		if err != nil {
			return err
		}

		if fc.GenerateThumbnail {
			err = GenerateThumbnail(fc)

			if err != nil {
				return err
			}
		}

		return nil
	}

	// Process Docx
	if strings.Contains(filename[len(filename)-6:], ".docx") {
		err := ProcessDocx(uuid.String(), desc, fc, conf)

		if err != nil {
			return err
		}

		if fc.GenerateThumbnail {
			err = GenerateThumbnail(fc)

			if err != nil {
				return err
			}
		}

		return nil
	}

	fmt.Println("Success")
	return nil
}

func NewQR(uuid, watermarkPath string) error {
	url := fmt.Sprintf("https://privy.id/verify/%s", uuid)

	qr, err := qrcode.New(url)

	if err != nil {
		return err
	}

	reader, err := os.Open(watermarkPath)
	defer reader.Close()

	if err != nil {
		return err
	}

	img, _, err := image.Decode(reader)

	if err != nil {
		return err
	}

	img = resize.Resize(125, 125, img, resize.NearestNeighbor)

	w, err := standard.New(fmt.Sprintf("./qr/%s.png", uuid), standard.WithLogoImage(img))

	if err != nil {
		return err
	}

	err = qr.Save(w)

	if err != nil {
		return err
	}

	return nil
}

func CheckFolder() {
	_, err := os.Stat("./qr")

	if os.IsNotExist(err) {
		os.Mkdir("./qr", os.ModePerm)
	}

	_, err = os.Stat("./result")

	if os.IsNotExist(err) {
		os.Mkdir("./result", os.ModePerm)
	}
}

func CheckFilename(resultName string) (string, error) {
	var (
		name string = resultName
		i           = 2
	)

	for {
		_, err := os.Stat(fmt.Sprintf("./result/%s.pdf", name))

		if os.IsNotExist(err) {
			return name, nil
		} else if err != nil {
			fmt.Println(err)
			return "", err
		}

		name = fmt.Sprintf("%s_%d", resultName, i)
		i++
	}
}
