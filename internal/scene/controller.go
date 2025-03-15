package scene

type Controller struct {
	scenes map[string]Scene
}

func NewController(scenes ...Scene) *Controller {
	sceneMap := make(map[string]Scene)
	for _, scene := range scenes {
		sceneMap[scene.Key()] = scene
	}
	return &Controller{sceneMap}
}

func (c *Controller) Transition(sceneKey string) error {
	if scene, ok := c.scenes[sceneKey]; ok {
		return scene.Run(c)
	}
	return ErrInvalidTransition
}
