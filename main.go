package main

import (
	"context"
	"fmt"
	"git.hocngay.com/techmaster/service-complier/cons"
	"git.hocngay.com/techmaster/service-complier/helper"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
	"log"
	"os/exec"
)

type Compiler struct{}

func (g *Compiler) Run(ctx context.Context, req *model.CompileRequest, rsp *model.CompileResponse) error {
	cErr := make(chan error, 1)
	path, err := helper.CreateFileComplie(req)
	if err != nil {
		return err
	}
	go func() {

		switch req.Language {
		case "go":
			cErr <- compileGo(path, rsp)
		case "js":
			cErr <- compileNode(path)
		case "py":
		}
	}()
	return <-cErr
}

func compileGo(path string, rsp *model.CompileResponse) error {

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

func compileNode(path string) error {
	go func() {
		out, err := exec.Command("node", path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}()
	return nil
}

/*func compilePython(fileID string) error {
}
func compileC(fileID string) error {
}*/

func main() {
	helper.Init()
	service := micro.NewService(
		micro.Name("compiler"),
	)

	service.Init()
	model.RegisterCompileHandler(service.Server(), new(Compiler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
