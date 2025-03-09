package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "", fmt.Errorf("SHELL environment variable is not set")
	}
	return shell, nil
}

func GetShellName() (string, error) {
	shell, err := GetShell()
	if err != nil {
		return "", err
	}
	name := strings.Split(shell, "/")[len(strings.Split(shell, "/"))-1]
	return name, nil
}

func GetShellRcPath() (string, error) {
	shell, err := GetShellName()
	if err != nil {
		return "", err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	switch shell {
	case "bash":
		return filepath.Join(homeDir, ".bashrc"), nil
	case "zsh":
		return filepath.Join(homeDir, ".zshrc"), nil
	case "fish":
		return filepath.Join(homeDir, ".config", "fish", "config.fish"), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}
