package fazpdf

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func ProcessPdf(uuid, desc string, fc FileConfig, conf *model.Configuration) error {
	wm, err := pdfcpu.ParseImageWatermarkDetails(fmt.Sprintf("./qr/%s.png", uuid), desc, true, types.POINTS)

	if err != nil {
		return err
	}

	ctx, err := api.ReadContextFile(fc.Src)

	if err != nil {
		return err
	}

	err = api.AddWatermarksFile(fc.Src, fmt.Sprintf("./result/%s.pdf", fc.ResultName), []string{fmt.Sprintf("%d", ctx.PageCount)}, wm, conf)

	if err != nil {
		return err
	}

	if fc.UpdateMetadata {
		err = ChangeMetadata(fmt.Sprintf("./result/%s.pdf", fc.ResultName), fc.Metadata)

		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessImage(uuid, desc string, fc FileConfig, conf *model.Configuration, imp *pdfcpu.Import) error {
	err := api.ImportImagesFile([]string{fc.Src}, fmt.Sprintf("./result/%s.pdf", fc.ResultName), imp, conf)

	if err != nil {
		return err
	}

	wm, err := pdfcpu.ParseImageWatermarkDetails(fmt.Sprintf("./qr/%s.png", uuid), desc, true, types.POINTS)

	if err != nil {
		return err
	}

	ctx, err := api.ReadContextFile(fmt.Sprintf("./result/%s.pdf", fc.ResultName))

	if err != nil {
		return err
	}

	err = api.AddWatermarksFile(fmt.Sprintf("./result/%s.pdf", fc.ResultName), fmt.Sprintf("./result/%s.pdf", fc.ResultName), []string{fmt.Sprintf("%d", ctx.PageCount)}, wm, conf)

	if err != nil {
		return err
	}

	if fc.UpdateMetadata {
		err = ChangeMetadata(fmt.Sprintf("./result/%s.pdf", fc.ResultName), fc.Metadata)

		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessDocx(uuid, desc string, fc FileConfig, conf *model.Configuration) error {
	cmd := exec.Command("pandoc", "-s", fc.Src, "-o", fmt.Sprintf("./result/%s.pdf", fc.ResultName))

	err := cmd.Run()

	if err != nil {
		return errors.New("failed to process docx make sure pandoc and mactex are installed")
	}

	wm, err := pdfcpu.ParseImageWatermarkDetails(fmt.Sprintf("./qr/%s.png", uuid), desc, true, types.POINTS)

	if err != nil {
		return err
	}

	ctx, err := api.ReadContextFile(fmt.Sprintf("./result/%s.pdf", fc.ResultName))

	if err != nil {
		return err
	}

	err = api.AddWatermarksFile(fmt.Sprintf("./result/%s.pdf", fc.ResultName), fmt.Sprintf("./result/%s.pdf", fc.ResultName), []string{fmt.Sprintf("%d", ctx.PageCount)}, wm, conf)

	if err != nil {
		return err
	}

	if fc.UpdateMetadata {
		err = ChangeMetadata(fmt.Sprintf("./result/%s.pdf", fc.ResultName), fc.Metadata)

		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessEncryptedPdf(src, uuid, resultName, password, desc string, fc FileConfig, conf *model.Configuration) error {
	conf.Cmd = model.DECRYPT
	conf.OwnerPW = password
	conf.UserPW = password
	err := api.DecryptFile(src, fmt.Sprintf("./result/%s.pdf", resultName), conf)

	if err != nil {
		return err
	}

	wm, err := pdfcpu.ParseImageWatermarkDetails(fmt.Sprintf("./qr/%s.png", uuid), desc, true, types.POINTS)

	if err != nil {
		return err
	}

	ctx, err := api.ReadContextFile(fmt.Sprintf("./result/%s.pdf", resultName))

	if err != nil {
		return err
	}

	err = api.AddWatermarksFile(fmt.Sprintf("./result/%s.pdf", resultName), fmt.Sprintf("./result/%s.pdf", resultName), []string{fmt.Sprintf("%d", ctx.PageCount)}, wm, conf)

	if err != nil {
		return err
	}

	if fc.UpdateMetadata {
		err = ChangeMetadata(fmt.Sprintf("./result/%s.pdf", resultName), fc.Metadata)

		if err != nil {
			return err
		}
	}

	return nil
}
