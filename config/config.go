package config

import (
	"log"
	"os"
	"strings"
)

type FileSystemTreeConfig struct {
	IntitalFolder  string
	TextColor      string
	FileExtensions []string
}

func New() *FileSystemTreeConfig {
	currentFolder, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return &FileSystemTreeConfig{
		IntitalFolder:  getEnvAsString("INITIAL_FOLDER", currentFolder),
		TextColor:      getEnvAsString("CONSOLE_TEXT_COLOR", "None"),
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
