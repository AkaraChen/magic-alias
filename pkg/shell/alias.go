package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/hairyhenderson/go-which"
)

var (
	validAliasRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

func AddAlias(alias string, command string) error {
	if alias == "" {
		return fmt.Errorf("alias cannot be empty")
	}
	if command == "" {
		return fmt.Errorf("command cannot be empty")
	}

	// Validate alias name
	if !validAliasRegex.MatchString(alias) {
		return fmt.Errorf("alias must contain only alphanumeric characters, underscores, and hyphens")
	}

	// Validate command exists
	if which.Which(command) == "" {
		return fmt.Errorf("command not found: %s", command)
	}

	aliasPath := filepath.Join(MagicAliasPath, alias)
	content := fmt.Sprintf("#!/usr/bin/env bash\n\nexec %s \"$@\"\n", command)
	if _, err := os.Stat(MagicAliasPath); err != nil {
		if err := os.Mkdir(MagicAliasPath, 0755); err != nil {
			return err
		}
	}
	err := os.WriteFile(aliasPath, []byte(content), 0755)
	if err != nil {
		return fmt.Errorf("failed to write alias file: %w", err)
	}
	return nil
}

func RemoveAlias(alias string) error {
	aliasPath := filepath.Join(MagicAliasPath, alias)
	if _, err := os.Stat(aliasPath); os.IsNotExist(err) {
		return fmt.Errorf("alias '%s' not found - it may not have been added yet", alias)
	}
	err := os.Remove(aliasPath)
	if err != nil {
		return fmt.Errorf("failed to remove alias: %w", err)
	}
	return nil
}

func ListAliases() ([]string, error) {
	aliases, err := os.ReadDir(MagicAliasPath)
	if err != nil {
		return nil, err
	}
	var aliasNames []string
	for _, alias := range aliases {
		if !alias.IsDir() {
			aliasNames = append(aliasNames, alias.Name())
		}
	}
	return aliasNames, nil
}
