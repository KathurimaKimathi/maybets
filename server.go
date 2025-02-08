package main

import (
	"context"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/helpers"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/presentation"
)

const (
	JaegerCodllectorEndpoint = "JAEGER_URL"
)

// StartApplication is used to start the application server
func StartApplication(ctx context.Context) error {
	port, err := helpers.ConvertPortToInt()
	if err != nil {
		return err
	}

	return presentation.StartServer(ctx, port)
}

func main() {
	ctx := context.Background()

	err := StartApplication(ctx)
	if err != nil {
		panic(err)
	}
}
