package main

import (
	"app/cmd"
	"app/internal/global"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	global.SetupEnv()
	cmd.RunService()
}
