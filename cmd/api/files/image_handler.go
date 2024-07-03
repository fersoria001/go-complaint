package files

import (
	"context"
	"fmt"
	"go-complaint/application/application_services"
	"go-complaint/application/commands"
	projectpath "go-complaint/project_path"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Folder int

const (
	PROFILE_IMG Folder = iota
	LOGO_IMG
	BANNER_IMG
)

func (f Folder) String() string {
	switch f {
	case PROFILE_IMG:
		return "profile_img"
	case LOGO_IMG:
		return "logo_img"
	case BANNER_IMG:
		return "banner_img"
	default:
		return ""
	}
}
func ParseFolder(folder string) Folder {
	switch folder {
	case "profile_img":
		return PROFILE_IMG
	case "logo_img":
		return LOGO_IMG
	case "banner_img":
		return BANNER_IMG
	default:
		return -1
	}
}
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	f := r.URL.Query().Get("folder")
	id := r.URL.Query().Get("id")
	folder := ParseFolder(f)
	if folder < 0 {
		log.Println("error at folder")
		w.WriteHeader(http.StatusBadRequest)
	}
	file, handler, err := r.FormFile(folder.String())
	if err != nil {
		log.Println("error at formfile", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	defer file.Close()
	fileType := filepath.Ext(handler.Filename)[1:]
	path := filepath.Join(projectpath.Root, "files", "png", folder.String())
	tempFile, err := os.CreateTemp(path, fmt.Sprintf("upload-*.%s", fileType))
	if err != nil {
		log.Println("error at tempfile", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	defer tempFile.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("error at read", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	name := tempFile.Name()
	tempFile.Write(fileBytes)

	//
	err = dispatchToCommand(r.Context(), folder, name, id)
	if err != nil {
		log.Println("error at dspatch", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func dispatchToCommand(ctx context.Context, folder Folder, name string, id string) error {
	_, after, found := strings.Cut(name, "\\png")
	if !found {
		return fmt.Errorf("path %s is incorrect", name)
	}
	s := strings.ReplaceAll(after, "\\", "/")
	s = strings.Replace(s, "/", "", 1)
	url := fmt.Sprintf("%s/%s", "https://docker-go-complaint-server-latest.onrender.com", s)
	switch folder {
	case PROFILE_IMG:
		credentials, err := application_services.AuthorizationApplicationServiceInstance().Credentials(ctx)
		if err != nil {
			return err
		}
		c := commands.UserCommand{
			ProfileIMG: url,
			Email:      credentials.Email,
			UpdateType: "profileIMG",
		}
		return c.UpdatePersonalData(ctx)
	case LOGO_IMG:
		if id == "" {
			return fmt.Errorf("enterprise_id is required")
		}
		_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
			ctx,
			"Enterprise",
			id,
			application_services.WRITE,
			"OWNER",
		)
		if err != nil {
			return err
		}
		c := commands.EnterpriseCommand{
			LogoIMG:    url,
			Name:       id,
			UpdateType: "logoIMG",
		}
		return c.UpdateEnterprise(ctx)
	case BANNER_IMG:
		if id == "" {
			return fmt.Errorf("enterprise_id is required")
		}
		_, err := application_services.AuthorizationApplicationServiceInstance().ResourceAccess(
			ctx,
			"Enterprise",
			id,
			application_services.WRITE,
			"OWNER",
		)
		if err != nil {
			return err
		}
		c := commands.EnterpriseCommand{
			BannerIMG:  url,
			Name:       id,
			UpdateType: "bannerIMG",
		}
		return c.UpdateEnterprise(ctx)
	default:
		return fmt.Errorf("invalid folder")
	}
}
