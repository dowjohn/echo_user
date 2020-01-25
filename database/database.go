package database

import "user/model"

var mdb interface{}

type Connection interface {
	UserSaver
}

type ConnectionImpl struct {
	UserSaverImpl
}

type UserSaver interface{
	Save(model.User) error
}

type UserSaverImpl struct {
}

func (s *UserSaverImpl) Save(user model.User) error {

	return nil
}

func Init() (Connection, error) {
	println("initializing database")

	return initializeConnection(), nil
}

func initializeConnection() Connection {
	return &ConnectionImpl{}
}
