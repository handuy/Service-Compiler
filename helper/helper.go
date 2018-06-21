package helper

import (
	"bufio"
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/rs/xid"
	"os"
	"io"
)

//Trả về file name đã được khởi tạo
func CreateFileComplie(req *model.CompileRequest) (string, error) {
	uuid := xid.New().String()
	var folder,filePath string
	switch req.Language {
	case "go","golang":
		filePath = fmt.Sprintf("%s/%s.%s", cons.RootGo, uuid, req.Language)
		folder=cons.RootGo
		break
	case "javascript","nodejs","js":
		filePath = fmt.Sprintf("%s/%s.%s", cons.RootJS, uuid, req.Language)
		folder=cons.RootJS
		break
	}

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, os.FileMode(0777))
	}

	_, err := SaveToFile(filePath, req)
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

func CopyFile(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	err = copyFileContents(src, dst)
	return
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}