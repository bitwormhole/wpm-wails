package boot

import (
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm-wails/app"
)

// Loader for boot
type Loader struct {
	markup.Component `class:"server-loader life"`

	Port       int                 `inject:"${server.port}"`
	WithGui    bool                `inject:"${wpm.options.run-with-gui}"`
	WithServer bool                `inject:"${wpm.options.run-with-server}"`
	AC         application.Context `inject:"context"`

	ctx         app.Ctx
	countBgTask int
}

func (inst *Loader) _Impl() application.LifeRegistry {
	return inst
}

func (inst *Loader) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit:  inst.init,
		OnStart: inst.start,
		OnStop:  inst.stop,
		Looper:  inst,
	}
}

func (inst *Loader) init() error {
	atts := inst.AC.GetAttributes()
	ctx, err := app.GetCtx(atts)
	if err != nil {
		return err
	}
	inst.ctx = ctx
	ctx.SetPort(inst.Port)
	return nil
}

func (inst *Loader) start() error {
	if inst.WithGui && inst.WithServer {
		inst.countBgTask++
		go func() {
			inst.runServerBackground()
		}()
	}
	return nil
}

func (inst *Loader) Loop() error {
	if inst.WithGui {
		return inst.ctx.RunGUI()
	} else if inst.WithServer {
		return inst.ctx.RunServer()
	}
	return nil
}

func (inst *Loader) stop() error {
	for inst.countBgTask > 0 {
		time.Sleep(time.Second * 2)
	}
	return nil
}

func (inst *Loader) runServerBackground() {
	defer func() {
		x := recover()
		inst.handleErrorX(x)
		inst.countBgTask--
	}()
	err := inst.ctx.RunServer()
	inst.handleError(err)
}

func (inst *Loader) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Error(err)
}

func (inst *Loader) handleErrorX(e any) {
	if e == nil {
		return
	}
	err, ok := e.(error)
	if ok {
		inst.handleError(err)
	}
}
