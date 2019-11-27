package service

import (
	"context"
	"xhyl-micro/service/common/lock/db"
	"xhyl-micro/service/ticket/proto"
)

type TicketService struct {
}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (t *TicketService) Buy(ctx context.Context, in *ticket.TicketRequest) (*ticket.TicketResponse, error) {

	lock.SetMysqlLock("ticket.buy")
	defer lock.ReleaseMysqlLock("ticket.buy")

	return &ticket.TicketResponse{
		Ticket: "1",
	}, nil
}
