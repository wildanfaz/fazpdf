package main

import (
	"fazpdf/fazpdf"
	"fmt"
)

func main() {
	// make sure the source file (Src) is in pdf / png / jpg / docx format
	fc := fazpdf.FileConfig{
		Src:               "./file/example.pdf",
		TargetPath:        "./result/",
		ResultName:        "test",
		WatermarkPath:     "./file/privyid-favicon.jpg",
		Password:          "",
		EncryptedPdf:      false,
		UpdateMetadata:    true,
		GenerateThumbnail: true,
		WatermarkPosition: 3,
		Dx:                -10,
		Dy:                10,
		Metadata: fazpdf.Metadata{
			Title:    "Test",
			Author:   "Muhamad Wildan Faz",
			Creator:  "https://github.com/wildanfaz",
			Producer: "fazpdf",
		},
		ThumbnailResolution: fazpdf.ThumbnailResolution{
			Height: 1600,
			Width:  1300,
		},
	}

	err := fc.NewPDF()
	if err != nil {
		fmt.Println(err)
	}
}

/*
Test 1 - pdf
fc := fazpdf.FileConfig{
	Src:               "./file/example.pdf",
	TargetPath:        "./result/",
	ResultName:        "test",
	WatermarkPath:     "./file/privyid-favicon.jpg",
	Password:          "",
	EncryptedPdf:      false,
	UpdateMetadata:    true,
	GenerateThumbnail: true,
	WatermarkPosition: 3,
	Dx:                -10,
	Dy:                10,
	Metadata: fazpdf.Metadata{
		Title:    "Test",
		Author:   "Muhamad Wildan Faz",
		Creator:  "https://github.com/wildanfaz",
		Producer: "fazpdf",
	},
	ThumbnailResolution: fazpdf.ThumbnailResolution{
		Height: 1600,
		Width:  1300,
	},
}

Test 2 - encrypted pdf
fc := fazpdf.FileConfig{
		Src:               "./file/example_protected.pdf",
		TargetPath:        "./result/",
		ResultName:        "test",
		WatermarkPath:     "./file/privyid-favicon.jpg",
		Password:          "helloworld",
		EncryptedPdf:      true,
		UpdateMetadata:    true,
		GenerateThumbnail: true,
		WatermarkPosition: 3,
		Dx:                -10,
		Dy:                10,
		Metadata: fazpdf.Metadata{
			Title:    "Test",
			Author:   "Muhamad Wildan Faz",
			Creator:  "https://github.com/wildanfaz",
			Producer: "fazpdf",
		},
		ThumbnailResolution: fazpdf.ThumbnailResolution{
			Height: 1600,
			Width:  1300,
		},
	}

	Test 3 - png or jpg
	fc := fazpdf.FileConfig{
		Src:               "./file/cedric-letsch-7Ew3m_d75Os-unsplash.jpg",
		TargetPath:        "./result/",
		ResultName:        "test",
		WatermarkPath:     "./file/privyid-favicon.jpg",
		Password:          "",
		EncryptedPdf:      false,
		UpdateMetadata:    true,
		GenerateThumbnail: true,
		WatermarkPosition: 3,
		Dx:                -10,
		Dy:                10,
		Metadata: fazpdf.Metadata{
			Title:    "Test",
			Author:   "Muhamad Wildan Faz",
			Creator:  "https://github.com/wildanfaz",
			Producer: "fazpdf",
		},
		ThumbnailResolution: fazpdf.ThumbnailResolution{
			Height: 1600,
			Width:  1300,
		},
	}

	Test 4 - docx
	fc := fazpdf.FileConfig{
		Src:               "./file/Free_Test_Data_100KB_DOCX.docx",
		TargetPath:        "./result/",
		ResultName:        "test",
		WatermarkPath:     "./file/privyid-favicon.jpg",
		Password:          "",
		EncryptedPdf:      false,
		UpdateMetadata:    true,
		GenerateThumbnail: true,
		WatermarkPosition: 3,
		Dx:                -10,
		Dy:                10,
		Metadata: fazpdf.Metadata{
			Title:    "Test",
			Author:   "Muhamad Wildan Faz",
			Creator:  "https://github.com/wildanfaz",
			Producer: "fazpdf",
		},
		ThumbnailResolution: fazpdf.ThumbnailResolution{
			Height: 1600,
			Width:  1300,
		},
	}
*/
