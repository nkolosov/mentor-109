package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/nkolosov/mentor-109/internal/app"
)

func main() {
	application := app.NewApplication()
	application.Start()
}
