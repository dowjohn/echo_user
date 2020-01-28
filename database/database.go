package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"user/model"
)

type Connection interface {
	UserSaver
	UserGetter
	UserUpdater
	UserDeleter
}

type ConnectionImpl struct {
	UserSaverImpl
	UserGetterImpl
	UserUpdaterImpl
	UserDeleterImpl
	*mongo.Client
}

type UserSaver interface{
	Save(model.User) (*model.User, error)
}

type UserSaverImpl struct {
	collection *mongo.Collection
}

func (s *UserSaverImpl) Save(user model.User) (*model.User, error) {
	return nil, nil
}

type UserGetter interface{
	Get(string) (*model.User, error)
}

type UserGetterImpl struct {
	collection *mongo.Collection
}

func (s *UserGetterImpl) Get(userId string) (*model.User, error) {
	u := new(model.User)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": userId}
	err := s.collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		log.Print(err)
	}

	return u, err
}

type UserUpdater interface{
	Update(*model.User) (*model.User, error)
}

type UserUpdaterImpl struct {
	collection *mongo.Collection
}

func (s *UserUpdaterImpl) Update(*model.User) (*model.User, error) {
	return nil, nil
}

type UserDeleter interface{
	Delete(string) (*string, error)
}

type UserDeleterImpl struct {
	collection *mongo.Collection
}

func (s *UserDeleterImpl) Delete(string) (*string, error) {
	return nil, nil
}

// todo the following functions should be the only code in this file
func Init() (Connection, error) {
	println("initializing database")

	return initializeConnection(), nil
}

func initializeConnection() Connection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("failed to secure database connection")
	}

	db := client.Database("AtlasDb")

	return &ConnectionImpl{
		UserSaverImpl:   UserSaverImpl{
			db.Collection("user"),
		},
		UserGetterImpl:  UserGetterImpl{
			db.Collection("user"),
		},
		UserUpdaterImpl: UserUpdaterImpl{
			db.Collection("user"),
		},
		UserDeleterImpl: UserDeleterImpl{
			db.Collection("user"),
		},
		Client:          client,
	}
}
