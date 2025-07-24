package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadMarkdownFiles(dir string, files []string) {
	fmt.Println("\n--- Markdown File Contents ---")
	for _, path := range files {
		if strings.HasSuffix(path, ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", path, err)
				continue
			}
			fmt.Printf("\n== %s ==\n%s\n", path, string(content))
		}
	}
}
