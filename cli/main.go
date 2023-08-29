package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mifer",
		Usage: "generate a random data queries",
		Action: func(cCtx *cli.Context) error {
			fmt.Println(cCtx.Args())
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "input_file",
				Aliases: []string{"in"},
				Usage:   "create table sql query file",
				Action: func(cCtx *cli.Context) error {
					sqlFilePath := cCtx.Args().Get(0)
					f, err := os.Stat(sqlFilePath)
					if err != nil {
						return fmt.Errorf("not found %v sql file", sqlFilePath)
					}
					if f.IsDir() {
						return fmt.Errorf("argument is directory")
					}
					if ext := filepath.Ext(sqlFilePath); ext != ".sql" {
						return fmt.Errorf("argument path must be sql extension")
					}

					return nil
				},
			},
			{
				Name:    "output_file",
				Aliases: []string{"out"},
				Usage:   "output generated insert query file",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("output:", cCtx.Args().Get(0))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
