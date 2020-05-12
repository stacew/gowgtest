package g

type Scene interface {
	StartUp()
	Update()
}

type scenemanager struct {
	broadCastScene Scene
}

var manager *scenemanager

func init() {
	manager = &scenemanager{}
}

func Update() error {
	if manager.broadCastScene != nil {
		manager.broadCastScene.Update()
	}
	return nil
}

func ReisterScene(scene Scene) {
	//manager.currentScene = scene
	scene.StartUp()
}

func UnRegisterScene(scene Scene) {

}
