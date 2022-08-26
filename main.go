package main

import (
	"log"
	"os"

	"github.com/Linrena/go_gen/internal"

	"github.com/Linrena/go_gen/internal/command"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "go_gen",
		Description: "golang code generator",
		Commands: []*cli.Command{
			command.ModelGenerate(),
		},
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Version: internal.Version,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
