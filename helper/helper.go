package helper

import (
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"strings"
)


//Trả về path của folder đã được khởi tạo
func CreateImageGo(data string) (string, error) {
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

	err = ioutil.WriteFile(uGoFile, []byte(data), 07777)
	if err != nil {
		return "", err
	}

	return uFolder, nil
}
