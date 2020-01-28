package database

import "user/model"

var mdb interface{}

type Connection interface {
	UserSaver
	UserGetter
}

type ConnectionImpl struct {
	UserSaverImpl
	UserGetterImpl
}

type UserSaver interface{
	Save(model.User) (*model.User, error)
}

type UserSaverImpl struct {

}

func (s *UserSaverImpl) Save(user model.User) (*model.User, error) {
	return nil, nil
}

type UserGetter interface{
	Get(string) (*model.User, error)
}

type UserGetterImpl struct {

}

func (s *UserGetterImpl) Get(userId string) (*model.User, error) {
	return nil, nil
}




func Init() (Connection, error) {
	println("initializing database")

	return initializeConnection(), nil
}

func initializeConnection() Connection {
	return &ConnectionImpl{}
}
