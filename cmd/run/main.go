package run

import (
	"context"
	"git.hocngay.com/techmaster/service-complier/helper"
	model "git.hocngay.com/techmaster/service-complier/proto"
	"github.com/micro/go-micro"
	"log"
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
	return <-cErr
}

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
