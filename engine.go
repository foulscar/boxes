package boxes

type Engine struct {
	ResourceManager *ResourceManager
	RuntimeHandler func(*Engine)
}

func LoadEngine() *Engine {
	e := &Engine{}
	e.initResourceManager()

	return e
}

func (e *Engine) SetRuntimeHandler(handler func(*Engine)) {
	e.RuntimeHandler = handler
}

func (e *Engine) Run() {
	defer e.ResourceManager.unload()

	e.RuntimeHandler(e)
}
