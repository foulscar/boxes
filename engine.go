package boxes

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

	e.RuntimeHandler(e)
}
