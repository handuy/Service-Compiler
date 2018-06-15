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

func CompileNode(path string, rsp *model.CompileResponse) error {
	filePath := fmt.Sprintf("%s:%s/%s", cons.ComplierJS, cons.PathFileJS, path)

	_, err := exec.Command("docker", "cp", cons.RootJS+"/"+path, filePath).Output()
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	logrus.Infof(" docker exec %s %s %s %s", cons.ComplierJS, "go", "run",
		fmt.Sprintf("%s/%s", cons.PathFileJS, path))
	out, err := exec.Command("docker", "exec", cons.ComplierJS, "node",
		fmt.Sprintf("%s/%s", cons.PathFileJS, path)).Output()
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}

	rsp.Result = string(out)
	return nil
}
