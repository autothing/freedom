package project

func init() {
	content["/domain/default.go"] = servicesTemplate()
}

func servicesInterfaceTemplate() string {
	return `package domain`
}

func servicesTemplate() string {
	return `
	// Code generated by 'freedom new-project {{.PackagePath}}'
	package domain

	import (
		"github.com/8treenet/freedom"
		"{{.PackagePath}}/adapter/repository"
	)
	
	func init() {
		freedom.Prepare(func(initiator freedom.Initiator) {
			initiator.BindService(func() *Default {
				return &Default{}
			})
			initiator.InjectController(func(ctx freedom.Context) (service *Default) {
				initiator.GetService(ctx, &service)
				return
			})
		})
	}
	
	// Default .
	type Default struct {
		Worker   freedom.Worker
		DefRepo   *repository.Default
	}
	
	// RemoteInfo .
	func (s *Default) RemoteInfo() (result struct {
		Ip string
		Ua string
	}) {
		s.Worker.Logger().Infof("我是service")
		result.Ip = s.DefRepo.GetIP()
		result.Ua = s.DefRepo.GetUA()
		return
	}

	`
}
