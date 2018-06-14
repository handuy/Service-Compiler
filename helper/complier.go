package helper

import (
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/sirupsen/logrus"
	"os/exec"
)

func CompileGo(path string, rsp *model.CompileResponse) error {
	filePath := fmt.Sprintf("%s:%s/%s", cons.ComplierGO, cons.PathFileGO, path)

	_, err := exec.Command("docker", "cp", cons.RootGo+"/"+path, filePath).Output()
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	logrus.Infof(" docker exec %s %s %s %s", cons.ComplierGO, "go", "run",
		fmt.Sprintf("%s/%s", cons.PathFileGO, path))
	out, err := exec.Command("docker", "exec", cons.ComplierGO, "go", "run",
		fmt.Sprintf("%s/%s", cons.PathFileGO, path)).Output()
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	rsp.Result = string(out)
	return nil
}

func CompileNode(path string) error {
	out, err := exec.Command("node", path).Output()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return nil
}
