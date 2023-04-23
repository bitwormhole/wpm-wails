package server

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// Loader for server
type Loader struct {
	markup.Component `class:"server-loader life"`
}

func (inst *Loader) _Impl() application.LifeRegistry {
	return inst
}

func (inst *Loader) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{}
}
