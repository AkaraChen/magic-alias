package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var magicAliasPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	magicAliasPath = filepath.Join(homeDir, ".magic-alias")
}

const magicAliasLine = `# Magic Alias
export PATH="$PATH:$HOME/.magic-alias"
`

func WriteMagicAliasToRc(rcPath string) error {
	// Check if the line already exists in the rc file
	content, err := os.ReadFile(rcPath)
	if err != nil {
		return fmt.Errorf("failed to read rc file: %w", err)
	}

	if strings.Contains(string(content), magicAliasLine) {
		return nil // Already exists, nothing to do
	}

	// Append the line to the rc file
	f, err := os.OpenFile(rcPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open rc file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString("\n" + magicAliasLine); err != nil {
		return fmt.Errorf("failed to write to rc file: %w", err)
	}

	return nil
}

func pathInPath(targetPath string) bool {
	pathEnv := os.Getenv("PATH")
	pathDirs := strings.Split(pathEnv, ":")

	for _, dir := range pathDirs {
		if dir == targetPath {
			return true
		}
	}
	return false
}

func IsMagicAliasInPath() (bool, error) {
	return pathInPath(magicAliasPath), nil
}
