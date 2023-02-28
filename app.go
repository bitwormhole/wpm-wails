package main

import (
	"context"
	"strconv"
	"time"
)

// App struct
type App struct {
	context *myContext
	ctx     context.Context
}

// NewApp creates a new App application struct
func NewApp(c *myContext) *App {
	return &App{context: c}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// HomePageURL 取主页地址
func (a *App) HomePageURL() string {

	time.Sleep(3 * time.Second)

	port := a.context.port
	return "http://localhost:" + strconv.Itoa(port)
}
