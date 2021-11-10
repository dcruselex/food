package main

import (
	"net/http"

	"github.com/dcruselex/food/internal/router"
)

func main() {
	http.ListenAndServe(":8888", router.GetRouter())
}
