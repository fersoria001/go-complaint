package infrastructure_test

import (
	"testing"
)

func TestApi(t *testing.T) {
	// // Open a file
	// file, err := os.Open("book.pdf")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// defer file.Close()
	// buf := bytes.NewBuffer(make([]byte, 0))
	// //s, err := io.ReadAll(buf)
	// written, err := io.Copy(buf, file)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Logf("written %d", written)
	// content, err := io.ReadAll(buf)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// //
	// var msg = server.SubmitFileMessage{
	// 	FileName: "copy_book.pdf",
	// 	FileType: "pdf",
	// 	File:     content,
	// }

	// fileContent, err := infrastructure.NewFileContent(
	// 	msg.FileName,
	// 	msg.FileType,
	// 	msg.File,
	// )
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// err = fileContent.Save()
	// if err != nil {
	// 	t.Fatal(err)
	// }
}

// func TestTXT(t *testing.T) {
// 	var msg = server.SubmitFileMessage{
// 		FileName: "test.txt",
// 		FileType: "txt",
// 		File:     []byte("Hello World"),
// 		Try:      1,
// 	}
// 	sub := server.Subscriber{
// 		User: dto.UserDescriptor{
// 			Email: "bercho001@gmail.com",
// 		},
// 		Enterprise: dto.Enterprise{
// 			Name: "test",
// 		},
// 	}
// 	pub := server.Publisher{}
// 	msgs, err := pub.ProccessFile(&sub, msg)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("msg: %v", string(msgs.Content))
// }
