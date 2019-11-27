package main

import (
	"context"
	"github.com/go-acme/lego/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"xhyl-micro/service/common/constant/micro_name"
	"xhyl-micro/service/ticket/proto"
	"xhyl-micro/service/ticket/service"
)

type TicketHandler struct {
}

func (t *TicketHandler) Buy(cxt context.Context, req *ticket.TicketRequest, resp *ticket.TicketResponse) error {
	response, err := service.NewTicketService().Buy(cxt, req)
	resp = response
	return err
}

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.10.33:8500",
		}
	})

	ticketService := micro.NewService(
		micro.Name(micro_name.MicroNameTicket),
		micro.Registry(reg),
		micro.Version("v1.0.0"),
	)

	ticketService.Init()

	ticket.RegisterTicketHandler(ticketService.Server(), &TicketHandler{})

	if err := ticketService.Run(); err != nil {
		log.Fatal(err)
	}
}
