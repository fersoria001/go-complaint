package projectpath

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _      = runtime.Caller(0)
	Root            = filepath.Join(filepath.Dir(b), "..")
	ProfileImgsPath = filepath.Join(Root, "files", "png", "profile_img")
	LogoImgsPath    = filepath.Join(Root, "files", "png", "logo_img")
	BannerImgsPath  = filepath.Join(Root, "files", "png", "banner_img")
	PresenterPath   = filepath.Join(Root, "presenter", "dist")
	IndexPath       = filepath.Join(Root, "presenter", "dist")
)
