package config

import "fmt"

type API struct {
	Port    int    `yaml:"port"`
	Address string `yaml:"address"`
}

func (a *API) ListenAddr() string {
	return fmt.Sprintf("%s:%d", a.Address, a.Port)
}
