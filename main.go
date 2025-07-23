package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/amirizzudin4/seek-assessment/utils"
)

func getFileExtension(filename string) string {
	return filepath.Ext(filename)
}

func isYamlFile(filename string) bool {
	ext := getFileExtension(filename)
	return ext == ".yaml" || ext == ".yml"
}

func main() {
	fmt.Println("starting app...")

	// declare input directory where manifest file is located
	inputPath := "./input"

	files, err := os.ReadDir(inputPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v \n", err)
		return
	}

	fmt.Printf("Files in %s \n", inputPath)
	for _, file := range files {
		if !file.IsDir() && isYamlFile(file.Name()) {
			fmt.Print("\n ------------------------------- \n")
			err := utils.ParseKubernetesManifest(fmt.Sprintf("%s/%s", inputPath, file.Name()))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("\n ------------------------------- \n")
		}
	}
}
