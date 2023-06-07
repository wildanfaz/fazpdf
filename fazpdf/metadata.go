package fazpdf

import (
	"os/exec"
)

type Metadata struct {
	Title    string
	Author   string
	Creator  string
	Producer string
}

func ChangeMetadata(src string, metadata Metadata) error {
	cmd := exec.Command("exiftool", "-Title="+metadata.Title, "-Author="+metadata.Author, "-Creator="+metadata.Creator, "-Producer="+metadata.Producer, "-overwrite_original", src)

	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
