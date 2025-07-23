package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// function to read yaml file and parse yaml into go object
func parseYamlFile(filePath string, target interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file at path %s: %v", filePath, err)
	}

	err = yaml.Unmarshal(data, target)

	if err != nil {
		return fmt.Errorf("error parsing YAML file at path %s: %v", filePath, err)
	}

	return nil
}

// function to convert Kubernetes manifest struct to json and write as file
func saveManifestAsJSON(manifest KubernetesManifest, fileNamePath string) error {
	outputDir := "./output"

	fileName := filepath.Base(fileNamePath)
	outputPath := filepath.Join(outputDir, fileName)

	jsonData, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %v", err)
	}

	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON file to %s: %v", outputPath, err)
	}

	fmt.Printf("JSON file saved to: %s\n", outputPath)
	return nil
}

func ParseKubernetesManifest(filePath string) error {
	fmt.Printf("parsing and validating file in %s \n", filePath)

	// declare manifest according KubernetesManifest struct,
	// if there is any wrong pattern detected during converting
	// output to manifest it will return error with fault on the line
	var manifest KubernetesManifest
	err := parseYamlFile(filePath, &manifest)
	if err != nil {
		return err
	}

	err = saveManifestAsJSON(manifest, filePath)
	if err != nil {
		return err
	}

	return nil
}
