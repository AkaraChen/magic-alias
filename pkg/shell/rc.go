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

const magicAliasLine = `# Magic Alias (ma)
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
	return pathInPath(MagicAliasPath), nil
}

// RemoveMagicAliasFromRc removes the magic alias PATH export from the rc file
func RemoveMagicAliasFromRc(rcPath string) error {
	// Read the rc file content
	content, err := os.ReadFile(rcPath)
	if err != nil {
		return fmt.Errorf("failed to read rc file: %w", err)
	}

	// Check if the magic alias line exists in the file
	contentStr := string(content)
	if !strings.Contains(contentStr, magicAliasLine) {
		return nil // Nothing to remove
	}

	// Remove the magic alias line
	newContent := strings.Replace(contentStr, "\n"+magicAliasLine, "", 1)
	// If the replacement didn't work (possibly due to no newline), try without newline
	if newContent == contentStr {
		newContent = strings.Replace(contentStr, magicAliasLine, "", 1)
	}

	// Write the updated content back to the file
	err = os.WriteFile(rcPath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated rc file: %w", err)
	}

	return nil
}
