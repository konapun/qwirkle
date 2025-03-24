package scene

const (
	SceneActionUnknown = -1
)

type Scene interface {
	Key() string
	Run(*Controller) error
}

type RuntimeScene struct {
	key    string
	runner func(*Controller) error
}

func NewScene(key string, runner func(*Controller) error) *RuntimeScene {
	return &RuntimeScene{key, runner}
}

func (s *RuntimeScene) Key() string {
	return s.key
}

func (s *RuntimeScene) Run(controller *Controller) error {
	return s.runner(controller)
}
