package main

import (
	"context"
	"go-template/cmd"
	"go-template/config"
	"go-template/internal/api/router"
	"log/slog"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func Init() cobra.Command {
	command := cobra.Command{
		Use: "server",
		RunE: func(c *cobra.Command, args []string) error {
			ctx := context.TODO()
			if path, err := c.Flags().GetString("config"); err != nil {
				return err
			} else if err := cmd.Init(path); err != nil {
				return err
			}
			myFigure := figure.NewFigure(config.C.APP, "", true)
			myFigure.Print()
			err := router.Run(ctx)
			return err
		},
	}
	command.PersistentFlags().String("config", "config.yaml", "Specify config file path")
	return command
}

func main() {
	command := Init()
	err := command.Execute()
	if err != nil {
		panic(err)
	}
	slog.Debug("Bye~")
}
