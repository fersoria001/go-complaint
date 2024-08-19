package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _      = runtime.Caller(0)
	Root            = filepath.Join(filepath.Dir(b), "..")
	ProfileImgsPath = filepath.Join(Root, "files", "profile_img")
	LogoImgsPath    = filepath.Join(Root, "files", "logo_img")
	BannerImgsPath  = filepath.Join(Root, "files", "banner_img")
	PresenterPath   = filepath.Join(Root, "presenter", "dist")
	IndexPath       = filepath.Join(Root, "presenter", "dist")
	EmailsPath      = filepath.Join(Root, "presentation", "email")
	CertPath = filepath.Join(Root, "server.crt")
	KeyPath = filepath.Join(Root, "server.key")
)
