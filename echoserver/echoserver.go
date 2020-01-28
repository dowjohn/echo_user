package echoserver

import (
	"github.com/labstack/echo"
	"net/http"
	"user/database"
	"user/model"
)

var db database.Connection

func Init(connection database.Connection) {
	println("initializing echoserver")
	db = connection
	e := echo.New()
	initializeEndpoints(*e)

	// start http server
	if err := e.Start(":1323"); err != nil {
		e.Logger.Fatal(err)
	}
}

func initializeEndpoints(e echo.Echo) {
	// User
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	u, err := db.Get(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, u)
}

// e.POST("/save", save)
func saveUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// todo write to db
	user, err := db.Save(*u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func updateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := db.Save(*u)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	_, err := db.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
