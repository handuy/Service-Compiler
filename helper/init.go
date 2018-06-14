package helper

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func Init() {
	logrus.Infoln("Init compiler container !")
	out, err := exec.Command("docker", "build", "-t", "complier-go", "./build/godoc/.").Output()
	if err != nil {
		logrus.Errorf("%s",err)
	}
	logrus.Infof("%s",out)

	out, err = exec.Command("docker", "run","-d","--rm", "--name","complier-go", "complier-go").Output()
	if err != nil {
		logrus.Errorf("%s\n",err)
	}
	logrus.Infof("%s\n",out)

}
