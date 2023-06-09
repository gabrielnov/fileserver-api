package main

import (
	"uploads-api/webserver"
)

func main() {
	r := webserver.InitRouter()
	r.Run("localhost:8888")
}
