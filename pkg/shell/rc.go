package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var MagicAliasPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	MagicAliasPath = filepath.Join(homeDir, ".magic-alias")
}

func WriteMagicAliasToRc(shell string) error {
	rcPath, err := GetShellRcPath()
	// Check if the line already exists in the rc file
	content, err := os.ReadFile(rcPath)
	if err != nil {
		return fmt.Errorf("failed to read rc file: %w", err)
	}

	if RemoveScriptContent(string(content)) != string(content) {
		return fmt.Errorf("magic alias line already exists in rc file")
	}

	// Append the line to the rc file
	f, err := os.OpenFile(rcPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open rc file: %w", err)
	}
	defer f.Close()

	scriptContent, err := RenderScriptContent(shell)
	if err != nil {
		return fmt.Errorf("failed to render script content: %w", err)
	}
	if _, err := f.WriteString("\n" + scriptContent); err != nil {
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
	return pathInPath(MagicAliasPath), nil
}

// RemoveMagicAliasFromRc removes the magic alias PATH export from the rc file
func RemoveMagicAliasFromRc(rcPath string) error {
	// Read the rc file content
	content, err := os.ReadFile(rcPath)
	if err != nil {
		return fmt.Errorf("failed to read rc file: %w", err)
	}

	contentStr := string(content)

	// Remove the magic alias line
	newContent := RemoveScriptContent(contentStr)

	if newContent == contentStr {
		return fmt.Errorf("magic alias line not found in rc file")
	}

	// Write the updated content back to the file
	err = os.WriteFile(rcPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated rc file: %w", err)
	}

	return nil
}
