package backend

import (
	"context"
	"media-down/backend/internal/cmd"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GetBinds 获取所有需要绑定的对象
func (a *App) GetBinds() []interface{} {
	return []interface{}{
		cmd.NewCmd(),
	}
}
