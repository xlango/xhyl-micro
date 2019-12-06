package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
	"xhyl-micro/service/common/constant/micro_name"
	"xhyl-micro/service/ticket/proto"
)

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.10.33:8500",
		}
	})

	userService := micro.NewService(
		micro.Name(micro_name.MicroNameUser),
		micro.Registry(reg),
		micro.RegisterInterval(15*time.Second),
		micro.RegisterTTL(30*time.Second),
	)
	userService.Init()

	ticketClient := ticket.NewTicketService(micro_name.MicroNameTicket, userService.Client())

	for i := 0; i < 40; i++ {
		go func(j int) {
			res, err := ticketClient.Buy(context.TODO(), &ticket.TicketRequest{Username: fmt.Sprintf("xxx%d", j)})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf(res.Ticket)
			}
		}(i)
	}

	if err := userService.Run(); err != nil {
		log.Fatal(err)
	}
}
