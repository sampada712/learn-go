package main

import (
	"net/http"

	"github.com/sampada712/learn-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
