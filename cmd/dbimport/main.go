package main

import (
	"github.com/rherlt/reval/internal/config"
	"github.com/rherlt/reval/internal/data"
	"github.com/rherlt/reval/internal/persistence"
)

func main() {

	config.Configure()
	persistence.SetupDb()
	data.ImportData()
	persistence.CloseClient()
}
