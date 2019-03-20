// todo.go init go with vuejs programmation
package  main

import (
	"github.com/labstack/echo"
	"github.com/labstack/engine/standard"
)


// main function
func main() {
	// create new instance of Echo
	e := echo.New()


	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Tasks "+c.Param("id")) })

	// Start as a web server
	Run(standard.New(":8000"))
}
