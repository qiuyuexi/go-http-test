package model

import (
	"sync"
)

var instance *TestModel

var once sync.Once

type TestModel struct {
	SingleDb
}

type TestTemplate struct {
	Id   int    `json:"id"`
	T    int    `json:"t"`
	Test string `json:"test"`
}

func newModel() *TestModel {
	model := new(TestModel)
	model.TableName = "t"
	model.Template = TestTemplate{}
	model.configFileName = "test.json"
	return model
}

/**
单例模式调用
 */
func GetTestModelInstance() *TestModel {
	once.Do(func() {
		instance = newModel()
	})
	return instance
}
