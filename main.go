package main

import (
	"main/example"
	"graphics.gd/classdb"
	"graphics.gd/startup"
)

func main() {
    classdb.Register[example.MyGoNode]() // this needs to happen before startup
    startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
    startup.Scene() // starts up the scene and blocks until the engine shuts down.
}