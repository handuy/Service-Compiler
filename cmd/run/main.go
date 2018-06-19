package main

import (
	"context"
	"errors"
	"fmt"
	"git.hocngay.com/techmaster/service-complier/helper"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"log"
	"time"
)

type Compiler struct{}

func (g *Compiler) Run(ctx context.Context, req *model.CompileRequest, rsp *model.CompileResponse) error {
	cErr := make(chan error, 1)
	path, err := helper.CreateFileComplie(req)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(1 * time.Second)
	//Chạy complier theo từng ngôn ngữ
	go func() {
		fmt.Println(time.Now())
		switch req.Language {
		case "c":
			cErr <- helper.ComplieC(path, rsp)
			return
		//case "c++":
		//	cErr <- helper.ComplieCPlus(path, rsp)
		//	return
		//case "c#":
		//	cErr <- helper.ComplieDotNet(path, rsp)
		//	return
		//case "php":
		//	cErr <- helper.ComplieDotNet(path, rsp)
		//	return
		//case "java":
		//	cErr <- helper.ComplieDotNet(path, rsp)
		//	return
		case "go":
			cErr <- helper.CompileGo(path, rsp)
			return
		case "js":
			cErr <- helper.CompileNode(path, rsp)
			return
		case "py":

		}
	}()

	//Nếu hàm compiler quá 3s thì sẽ báo lỗi
	go func() {
		for c := range ticker.C {
			cErr <- errors.New(fmt.Sprintf("%s Run too long \n", c.String()))
		}
	}()

	fmt.Println(time.Now())
	return <-cErr
}

func main() {
	helper.Init()
	service := micro.NewService(
		micro.Name("compiler"),
		micro.Registry(registry.NewRegistry(registry.Timeout(4*time.Second))),
	)

	service.Init()
	model.RegisterCompileHandler(service.Server(), new(Compiler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
