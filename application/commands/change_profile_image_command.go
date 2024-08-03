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

type ChangeProfileImageCommand struct {
	UserId     string    `json:"userId"`
	FileName   string    `json:"fileName"`
	FileReader io.Reader `json:"fileReader"`
}

func NewChangeProfileImageCommand(userId, fileName string, fileReader io.Reader) *ChangeProfileImageCommand {
	return &ChangeProfileImageCommand{
		UserId:     userId,
		FileName:   fileName,
		FileReader: fileReader,
	}
}

func (c ChangeProfileImageCommand) Execute(ctx context.Context) error {
	userId, err := uuid.Parse(c.UserId)
	if err != nil {
		return err
	}
	path := filepath.Join(projectpath.Root, "files", "profile_img", c.FileName)
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
	resource := fmt.Sprintf("%s/%s", "profile_img", c.FileName)
	url := fmt.Sprintf("%s/%s", dns, resource)
	reg := repositories.MapperRegistryInstance()
	r, ok := reg.Get("User").(repositories.UserRepository)
	if !ok {
		return ErrWrongTypeAssertion
	}
	user, err := r.Get(ctx, userId)
	if err != nil {
		return err
	}
	err = user.ChangeProfileIMG(ctx, url)
	if err != nil {
		return err
	}
	err = r.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
