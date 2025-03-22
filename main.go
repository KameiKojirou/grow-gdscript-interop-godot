package main

import (
	"main/example"

	"graphics.gd/classdb"
	"graphics.gd/startup"
)

func main() {
    startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality

    classdb.Register[example.MyGoNode]()
    startup.Scene() // starts up the scene and blocks until the engine shuts down.
}