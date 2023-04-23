package app

import (
	"fmt"

	"github.com/bitwormhole/starter/collection"
)

const theBindAttrName = "github.com/bitwormhole/wpm-wails/app/Ctx#84424d1e1c#binding"

type Ctx interface {
	GetPort() int

	SetPort(port int)

	RunServer() error

	RunGUI() error
}

func BindCtx(c1 Ctx, fn func(name string, c2 Ctx)) {
	fn(theBindAttrName, c1)
}

func GetCtx(atts collection.Attributes) (Ctx, error) {
	o, err := atts.GetAttributeRequired(theBindAttrName)
	if err != nil {
		return nil, err
	}
	ctx, ok := o.(Ctx)
	if !ok {
		return nil, fmt.Errorf("bad cast of app.Ctx")
	}
	return ctx, nil
}
