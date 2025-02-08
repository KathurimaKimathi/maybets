package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/helpers"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/presentation"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx := context.Background()

	usecases, err := presentation.ConfigureStartUpDependencies()
	if err != nil {
		log.Fatalf("Failed to configure start up dependencies: %v", err)
	}

	app := &cli.App{
		Name:  "betting-analytics",
		Usage: "Analyze betting data",
		Commands: []*cli.Command{
			{
				Name:  "process",
				Usage: "Process betting data from a file",
				Action: func(c *cli.Context) error {
					filename := c.Args().First()
					bets, err := helpers.LoadBetsFromFile(filename)
					if err != nil {
						return err
					}

					if err = usecases.ProcessBets(ctx, bets); err != nil {
						return fmt.Errorf("failed to process bets: %w", err)
					}

					fmt.Println("Processing complete!")
					return nil
				},
			},
			{
				Name:  "runserver",
				Usage: "Start the analytics API server",
				Action: func(_ *cli.Context) error {
					port, err := helpers.ConvertPortToInt()
					if err != nil {
						return err
					}

					err = presentation.StartServer(ctx, port)
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "generate",
				Usage: "Generate test bet data",
				Action: func(c *cli.Context) error {
					numRecords := c.Int("betdata")
					filename := c.Args().First()
					err := helpers.GenerateTestData(filename, numRecords)
					if err != nil {
						return err
					}

					fmt.Printf("Generated %d records in %s\n", numRecords, filename)
					return nil
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "betdata",
						Value: 10000,
						Usage: "Total number of bet data records to generate",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
