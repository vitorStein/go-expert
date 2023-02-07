package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("Agendamento no hotel cancelado, tempo excedido!")
		return

	case <-time.After(time.Second * 5):
		fmt.Println("Reserva agendada!")
		return
	}

}
