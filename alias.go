package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// Handle command line flags
	flag.Parse()
	root := flag.Arg(0)
	// Create slice to store entries
	e := []string{}
	// Function to execute when walk encounters a hit
	var visit = func(file_path string, _ os.FileInfo, _ error) error {
		if filepath.Ext(file_path) == ".projlib" || filepath.Ext(file_path) == ".jar" {
			e = append(e, (entry(path.Base(file_path))))
		}
		return nil
	}
	// Walk filepath
	err := filepath.Walk(root, visit)
	if err != nil {
		panic(err)
	}
	// Write entries to output file
	write(e)
}

// Function to build entries
func entry(path string) string {
	// Entry components
	prefix := "tibco.alias."
	// Replacing filepath separators is only nessasay for non forward slashes and this can thus be more efficient
	key := strings.Replace(strings.SplitAfterN(path, "components" + string(filepath.Separator), 2)[1], string(filepath.Separator), "/", -1)
	equals := "="
	value := ""
	// Windows aka "\" only
	if filepath.Separator == 92 {
		value = strings.Replace(path, "\\", "\\\\", -1)
	} else {
		value = path
	}
	// Build entry
	entry := strings.Join([]string{prefix, key, equals, value}, "")
	return entry
}

func write(entries []string) {
	// Create file
	fo, err := os.Create("FileAliases.properties")
	if err != nil {
		panic(err)
	}
	// Write entries
	for _, entry := range entries {
		_, err := fo.WriteString(strings.TrimSpace(entry) + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
