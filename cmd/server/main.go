package main

import (
	"go-template/cmd"
	"go-template/config"
	"log/slog"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func Init() cobra.Command {
	command := cobra.Command{
		Use: "server",
		RunE: func(c *cobra.Command, args []string) error {
			if path, err := c.Flags().GetString("config"); err != nil {
				return err
			} else if err := cmd.Init(path); err != nil {
				return err
			}
			myFigure := figure.NewFigure(config.C.APP, "", true)
			myFigure.Print()
			return nil
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
