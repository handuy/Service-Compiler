package main

import (
	"context"
	"fmt"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/micro/go-micro"
	"log"
	"os/exec"
	"git.hocngay.com/techmaster/service-complier/helper"
	"time"
)

type Compiler struct{}


func (g *Compiler) Run(ctx context.Context, req *model.CompileRequest, rsp *model.CompileResponse) error {
	path, err := helper.CreateImageGo(req)
	if err != nil {
		return err
	}

	switch req.Language {
	case "go":
		compileGo(path)
	case "js":
		compileNode(path)
	case "py":
	}
	return nil
}

func compileGo(path string) error {
	go func() {
		fmt.Println(time.Now().Second())
		out, err := exec.Command("docker","build","-t","test",path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n",out)
		out, err = exec.Command("docker","run","test" ).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\n",time.Now().Second()," ",out)
	}()
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
	service := micro.NewService(
		micro.Name("compiler"),
	)

	service.Init()
	model.RegisterCompileHandler(service.Server(), new(Compiler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
