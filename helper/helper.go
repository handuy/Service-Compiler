package helper

import (
	"bufio"
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/rs/xid"
	"os"
)

//Trả về path của folder đã được khởi tạo
func CreateFileComplie(req *model.CompileRequest) (string, error) {
	uuid := xid.New().String()
	uGoFile := fmt.Sprintf("%s/%s.%s", cons.RootGo, uuid, req.Language)

	if _, err := os.Stat(cons.RootGo); os.IsNotExist(err) {
		os.MkdirAll(cons.RootGo, os.FileMode(0777))
	}

	_, err := SaveToFile(uGoFile, req)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", uuid, req.Language), nil
}

func SaveToFile(path string, req *model.CompileRequest) (string, error) {
	file, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range req.SourceCode {
		fmt.Fprintln(w, line)
	}

	return path, w.Flush()
}
