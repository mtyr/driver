package service

import (
	"github.com/hal-ms/driver/building/db"
	"github.com/hal-ms/driver/building/model"
)

var Scene = sceneService{}

type sceneService struct {
}

func (s *sceneService) Write(scene string, data model.Scene) error {
	db := db.DB.Get()
	var d model.Scene
	err := db.Read("scene", scene, &d)
	if err != nil {
		panic(err)
	}

	db.Write("scene", scene, data)

	return nil
}
