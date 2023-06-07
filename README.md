# fazpdf

Simple Library for Generate Pdf and Place QR in Pdf / Png / Jpg / Docx

## Table of Contents

- [Installation](#installation)
- [Requirements](#requirements)
- [Usage](#usage)
- [Features](#features)

## Installation
```
git clone https://github.com/wildanfaz/fazpdf.git
```

## Requirements (mac)
1. pdfcpu
manipulate pdf in golang
```
brew install pdfcpu
```
2. exiftool
change metadata
```
brew install exiftool
```
3. pandoc
convert docx to pdf
```
brew install pandoc
```
4. basictex
convert docx to pdf
```
brew install --cask basictex
```
5. pdftoppm
generate thumbnail from pdf
```
brew install poppler
```

## Usage
```
go run main.go
```

## Features

- Generate QR
- Watermarked QR
- Input pdf then place qr on last page
- Custom qr position and offset on last page
- Convert png / jpg / docx to pdf
- Generate thumbnail from first page pdf
- Unlock encrypted pdf
- Custom metadata pdf
