package main

import "gorm.io/gorm"

type Example struct {
	gorm.Model

	Name   string  `json:"name,omitempty"`
	Scenes []Scene `json:"scenes,omitempty"`
}

func (Example) TableName() string { return "examples" }

type Scene struct {
	gorm.Model

	Name      string `json:"name,omitempty"`
	ExampleID int    `json:"exampleID,omitempty"`

	// ResourceIDs []int `json:"resourceIDs,omitempty"`
}

func (Scene) TableName() string { return "scenes" }

type SceneResourceBinding struct {
	gorm.Model

	SceneID    int `json:"sceneID,omitempty"`
	ResourceID int `json:"resourceID,omitempty"`
}

func (SceneResourceBinding) TableName() string { return "scene_resource_bindings" }
