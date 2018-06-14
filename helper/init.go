package helper

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func Init() {
	initGo()
	initNodeJS()
}

func initGo() {
	logrus.Infoln("Init compiler GO container !")
	out, err := exec.Command("docker", "build", "-t", "complier-go", "./build/godoc/.").Output()
	if err != nil {
		logrus.Errorf("%s", err)
	}
	logrus.Infof("%s", out)

	out, err = exec.Command("docker", "run", "-d", "--rm", "--name", "complier-go", "complier-go").Output()
	if err != nil {
		logrus.Errorf("%s\n", err)
	}
	logrus.Infof("%s\n", out)
}

func initNodeJS()  {
	logrus.Infoln("Init compiler GO container !")
	out, err := exec.Command("docker", "build", "-t", "complier-nodejs", "./build/nodedoc/.").Output()
	if err != nil {
		logrus.Errorf("%s", err)
	}
	logrus.Infof("%s", out)

	out, err = exec.Command("docker", "run", "-d", "--rm", "--name", "complier-nodejs","complier-nodejs").Output()
	if err != nil {
		logrus.Errorf("%s\n", err)
	}
	logrus.Infof("%s\n", out)
}