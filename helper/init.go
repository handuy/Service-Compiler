package helper

import (
	"git.hocngay.com/techmaster/service-complier/cons"
	"github.com/sirupsen/logrus"
	"os/exec"
)

func Init() {
	initGo()
	initNodeJS()
}

func initGo() {
	logrus.Infoln("Init compiler GO container !")
	out, err := exec.Command("docker", "build", "-t", cons.ComplierGO, cons.GoDockerfile).Output()
	if err != nil {
		logrus.Errorf("%s", err)
	}
	logrus.Infof("%s", out)

	out, err = exec.Command("docker", "run", "-d", "--rm", "--name", cons.ComplierGO, cons.ComplierGO).Output()
	if err != nil {
		logrus.Errorf("%s\n", err)
	}
	logrus.Infof("%s\n", out)
}

func initNodeJS() {
	logrus.Infoln("Init compiler JS container !")
	out, err := exec.Command("docker", "build", "-t", cons.ComplierJS, cons.JSDockerfile).Output()
	if err != nil {
		logrus.Errorf("%s", err)
	}
	logrus.Infof("%s", out)

	out, err = exec.Command("docker", "run", "-d", "--rm", "--name", cons.ComplierJS, cons.ComplierJS).Output()
	if err != nil {
		logrus.Errorf("%s\n", err)
	}
	logrus.Infof("%s\n", out)
}
