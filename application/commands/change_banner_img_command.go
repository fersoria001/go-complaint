package commands

import (
	"context"
	"fmt"
	"go-complaint/infrastructure/persistence/repositories"
	projectpath "go-complaint/project_path"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type ChangeBannerImageCommand struct {
	EnterpriseId string    `json:"enterpriseId"`
	FileName     string    `json:"fileName"`
	FileReader   io.Reader `json:"fileReader"`
}

func NewChangeBannerImageCommand(enterpriseId, fileName string, fileReader io.Reader) *ChangeBannerImageCommand {
	return &ChangeBannerImageCommand{
		EnterpriseId: enterpriseId,
		FileName:     fileName,
		FileReader:   fileReader,
	}
}

func (c ChangeBannerImageCommand) Execute(ctx context.Context) error {
	enterpriseId, err := uuid.Parse(c.EnterpriseId)
	if err != nil {
		return err
	}
	path := filepath.Join(projectpath.Root, "files", "banner_img", c.FileName)
	permissions := 0644
	b, err := io.ReadAll(c.FileReader)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, b, fs.FileMode(permissions))
	if err != nil {
		return err
	}
	dns := os.Getenv("DNS")
	resource := fmt.Sprintf("%s/%s", "banner_img", c.FileName)
	url := fmt.Sprintf("%s/%s", dns, resource)
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("Enterprise").(repositories.EnterpriseRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	dbEnterprise, err := r.Get(ctx, enterpriseId)
	if err != nil {
		return err
	}
	err = dbEnterprise.ChangeBannerIMG(ctx, url)
	if err != nil {
		return err
	}
	err = r.Update(ctx, dbEnterprise)
	if err != nil {
		return err
	}
	return nil
}
