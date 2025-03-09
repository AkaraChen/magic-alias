package shell

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddAlias(alias string, command string) error {
	if alias == "" {
		return fmt.Errorf("alias cannot be empty")
	}
	if command == "" {
		return fmt.Errorf("command cannot be empty")
	}

	aliasPath := filepath.Join(MagicAliasPath, alias)
	content := fmt.Sprintf("#!/usr/bin/env bash\n\nexec %s \"$@\"\n", command)
	if _, err := os.Stat(MagicAliasPath); err != nil {
		if err := os.Mkdir(MagicAliasPath, 0755); err != nil {
			return err
		}
	}
	err := os.WriteFile(aliasPath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func RemoveAlias(alias string) error {
	aliasPath := filepath.Join(MagicAliasPath, alias)
	err := os.Remove(aliasPath)
	return err
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
