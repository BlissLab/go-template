package cmd

import (
	"go-template/config"
	"go-template/utils/log"
)

func Init(filename string) error {
	config.LoadFile(filename)
	log.Init()
	return nil
}
