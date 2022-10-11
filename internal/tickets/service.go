package tickets

import (
	"context"
)


type service struct {
	respository Repository
}

func NewService(r Repository) Service {
	return &service {
		respository: r,
	}
}

type Service interface {
	AverageDestination(ctx context.Context, destination string) (float64, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	//El  otro método debe devolver el promedio de personas que viajan a un país determinado en un dia
	ticketsXDest, err := s.respository.GetTicketByDestination(ctx, destination)
	if err!= nil {
		return 0, err
	}
	totalTickets, err := s.respository.GetAll(ctx)
	if err!= nil {
		return 0, err
	}
	return (float64(len(ticketsXDest)) / float64(len(totalTickets))), nil
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	//debe devolver la cantidad de tickets de un destino.
	tickets, err := s.respository.GetTicketByDestination(ctx, destination)
	if err!= nil {
		return 0, err
	}
	
	return len(tickets), nil
}