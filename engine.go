package boxes

type Engine struct {
	ResourceManager *ResourceManager
}

func LoadEngine() *Engine {
	e := &Engine{}
	e.initResourceManager()

	return e
}
