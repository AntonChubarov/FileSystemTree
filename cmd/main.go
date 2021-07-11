package main

import (
	"../app"
	"../config"
	"fmt"
	colorText "github.com/daviddengcn/go-colortext"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	appConfig := config.New()

	colorText.Foreground(appConfig.TextColor, false)

	fmt.Println(appConfig.IntitalFolder)

	app.PrintFolder(appConfig.IntitalFolder, true, true, make([]string, 1), appConfig.FileExtensions)

	fmt.Println("Press Enter to exit")

	colorText.ResetColor()

	_, err := fmt.Scanln()
	if err != nil {
		log.Println(err)
	}
}
