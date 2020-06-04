package main

import (
	"github.com/mopeneko/novel-gamest/api/infrastructure"
)

func main() {
	infrastructure.InitDB()
	infrastructure.InitRouter()
	infrastructure.Run()
}
