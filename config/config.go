package config

import (
	colorText "github.com/daviddengcn/go-colortext"
	"log"
	"os"
	"strings"
)

type FileSystemTreeConfig struct {
	IntitalFolder  string
	TextColor      colorText.Color
	FileExtensions []string
}

func New() *FileSystemTreeConfig {
	currentFolder, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return &FileSystemTreeConfig{
		IntitalFolder:  getEnvAsString("INITIAL_FOLDER", currentFolder),
		TextColor:      getEnvAsColor("CONSOLE_TEXT_COLOR", colorText.None),
		FileExtensions: getEnvAsSlice("FILE_TYPES_TO_DISPLAY", []string{"*.go"}, " "),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsString(name string, defaultVal string) string {
	valueStr := getEnv(name, defaultVal)
	return valueStr
}

//func getEnvAsInt(name string, defaultVal int) int {
//	valueStr := getEnv(name, "")
//	if value, err := strconv.Atoi(valueStr); err == nil {
//		return value
//	}
//
//	return defaultVal
//}
//
//func getEnvAsBool(name string, defaultVal bool) bool {
//	valStr := getEnv(name, "")
//	if val, err := strconv.ParseBool(valStr); err == nil {
//		return val
//	}
//
//	return defaultVal
//}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

func getEnvAsColor(name string, defaultVal colorText.Color) colorText.Color {
	valStr := getEnv(name, "")

	var colorMap = map[string]int{
		"None":    0,
		"Black":   1,
		"Red":     2,
		"Green":   3,
		"Yellow":  4,
		"Blue":    5,
		"Magenta": 6,
		"Cyan":    7,
		"White":   8,
	}

	if valStr == "" {
		return defaultVal
	}

	if val, ok := colorMap[valStr]; ok {
		return colorText.Color(val)
	}

	return defaultVal
}
