package service

import (
	"context"
	"fmt"
	"xhyl-micro/service/common/lock/cache"
	"xhyl-micro/service/ticket/proto"
)

type TicketService struct {
}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (t *TicketService) Buy(ctx context.Context, in *ticket.TicketRequest) (*ticket.TicketResponse, error) {

	//lock.SetMysqlLock("ticket.buy")
	lock.SetRedisLock("ticket.buy")
	fmt.Printf("%v用户请求获取到锁\n", in.Username)
	//defer lock.ReleaseMysqlLock("ticket.buy")
	defer lock.ReleaseRedisLock("ticket.buy")

	return &ticket.TicketResponse{
		Ticket: "票" + in.Username,
	}, nil
}
