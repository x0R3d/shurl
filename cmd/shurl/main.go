package main

import (
	"fmt"
	"shurl/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog

	// TODO: init db: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}