package helper

import (
	"bufio"
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"strings"
)

//Trả về path của folder đã được khởi tạo
func CreateImageGo(req *model.CompileRequest) (string, error) {
	uuid := xid.New().String()
	uFolder := fmt.Sprintf("%s/%s", cons.RootTemp, uuid)
	uDocker := fmt.Sprintf("%s/%s", uFolder, "Dockerfile")
	uGoFile := fmt.Sprintf("%s/%s.go", uFolder, uuid)

	if _, err := os.Stat(uFolder); os.IsNotExist(err) {
		os.MkdirAll(uFolder, os.FileMode(0777))
	}

	dockerFile := strings.Replace(cons.GoDocker, "temp_folder_id", uuid, -1)
	err := ioutil.WriteFile(uDocker, []byte(dockerFile), 07777)
	if err != nil {
		return "", err
	}

	_, err = SaveToFile(uGoFile, req)
	if err != nil {
		return "", err
	}
	shell := strings.Replace(cons.BaseBash, "temp_folder_id", uuid, -1)
	err = ioutil.WriteFile(uFolder+"/run.sh", []byte(shell), 0777)
	if err != nil {
		return "", err
	}

	return uFolder, nil
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
