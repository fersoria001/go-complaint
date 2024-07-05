package cmd_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestCorrectFilePath(t *testing.T) {
	dns := os.Getenv("DNS")
	log.Println("FILEDNS", dns)
	name := "/home/ec2-user/docker-go-complaint/files/png/profile_img/upload-3648084910.jpg"
	_, after, win := strings.Cut(name, "\\png")
	var s string
	if !win {
		_, after, lin := strings.Cut(name, "/png")
		if !lin {
			t.Log("path not correct in win or lin os ")
		}
		s = strings.Replace(after, "/", "", 1)
	} else {
		s = strings.ReplaceAll(after, "\\", "/")
		s = strings.Replace(s, "/", "", 1)
	}
	url := fmt.Sprintf("%s/%s", dns, s)
	log.Println("url", url)
}
