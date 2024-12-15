package boxes

import (
	"log"
)

type Engine struct {
	ResourceManager *ResourceManager
	RuntimeHandler  func(*Engine)
}

func LoadEngine() *Engine {
	e := &Engine{}
	e.InitResourceManager()

	return e
}

func (e *Engine) SetRuntimeHandler(handler func(*Engine)) {
	e.RuntimeHandler = handler
}

func (e *Engine) Run() {
	defer e.ResourceManager.Unload()
	if e.RuntimeHandler == nil {
		log.Fatal("Engine runtime handler has not been set")
	}

	e.RuntimeHandler(e)
}
