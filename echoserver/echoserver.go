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
	e.Logger.Fatal(e.Start(":1323"))
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
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.POST("/save", save)
func saveUser(c echo.Context) error {
	u := model.User{
		Name:  c.FormValue("name"),
		Email: c.FormValue("email"),
	}

	// todo write to db
	if err := db.SaveAny(); err != nil {
		return err
	}

	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func updateUser(c echo.Context) error {
	return nil
}

func deleteUser(c echo.Context) error {
	return nil
}
