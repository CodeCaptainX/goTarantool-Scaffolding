package internal

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// replacePlaceholders replaces specified placeholders in the file content
func replacePlaceholders(content, projectName string) string {
	// Replace common placeholders with the project name
	replacements := map[string]string{
		"{{PROJECT_NAME}}": projectName,
		"${PROJECT_NAME}":  projectName,
	}

	for placeholder, replacement := range replacements {
		content = strings.ReplaceAll(content, placeholder, replacement)
	}

	return content
}

func copyFile(srcPath, destPath, projectName string) error {
	// Ensure the destination directory exists
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("error creating destination directory %s: %v", destDir, err)
	}

	// Open source file
	src, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("error opening source file %s: %v", srcPath, err)
	}
	defer src.Close()

	// Read file content
	content, err := io.ReadAll(src)
	if err != nil {
		return fmt.Errorf("error reading source file %s: %v", srcPath, err)
	}

	// Replace placeholders
	modifiedContent := replacePlaceholders(string(content), projectName)

	// Create destination file
	dest, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("error creating destination file %s: %v", destPath, err)
	}
	defer dest.Close()

	// Write modified content
	_, err = dest.WriteString(modifiedContent)
	if err != nil {
		return fmt.Errorf("error writing to destination file %s: %v", destPath, err)
	}

	return nil
}

func CreateProjectStructure(basePath, projectName string, structure map[string]interface{}) error {
	for name, content := range structure {

		path := filepath.Join(basePath, name)

		switch v := content.(type) {
		case map[string]interface{}:
			// This is a directory
			err := os.MkdirAll(path, 0755)
			if err != nil {
				return fmt.Errorf("error creating directory %s: %v", path, err)
			}

			// Recursively create subdirectories and files
			if len(v) > 0 {
				if err := CreateProjectStructure(path, projectName, v); err != nil {
					return err
				}
			}

		case string:
			// This is a template file
			if strings.HasPrefix(v, "template/") {
				// Remove the "template/" prefix
				templatePath := filepath.Join("/home/khomchomroeun/backend/go/generate_project/template", strings.TrimPrefix(v, "template/"))
				fmt.Println("🚀 ~ file: createProjectStructure.go ~ line 90 ~ ifstrings.HasPrefix ~ templatePath : ", templatePath)

				// Copy the template file with placeholder replacement
				err := copyFile(templatePath, path, projectName)
				if err != nil {
					return fmt.Errorf("error copying template file %s to %s: %v", templatePath, path, err)
				}
				// fmt.Printf("Created file: %s (from template %s)\n", path, templatePath)
			}

		case nil:
			// This is an empty file
			file, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("error creating empty file %s: %v", path, err)
			}
			file.Close()
			// fmt.Printf("Created empty file: %s\n", path)
		}
	}
	return nil
}
