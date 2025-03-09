package shell

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddAlias(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	MagicAliasPath = tmpDir

	tests := []struct {
		name    string
		alias   string
		command string
		wantErr bool
	}{
		{"valid alias", "m", "sh", false},
		{"empty alias", "", "git", true},
		{"empty command", "m", "", true},
		{"invalid alias special chars", "m@", "git", true},
		{"invalid alias spaces", "m m", "git", true},
		{"non-existent command", "m", "nonexistentcmd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddAlias(tt.alias, tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddAlias() error = %v, wantErr %v, test case %s", err, tt.wantErr, tt.name)
				return
			}

			if !tt.wantErr {
				aliasPath := filepath.Join(MagicAliasPath, tt.alias)
				if _, err := os.Stat(aliasPath); os.IsNotExist(err) {
					t.Errorf("alias file was not created")
				}
			}
		})
	}
}

func TestRemoveAlias(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	MagicAliasPath = tmpDir

	// Create test alias
	err := AddAlias("m", "git")
	if err != nil {
		t.Fatalf("failed to create test alias: %v", err)
	}

	tests := []struct {
		name    string
		alias   string
		wantErr bool
	}{
		{"existing alias", "m", false},
		{"non-existent alias", "nonexistent", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RemoveAlias(tt.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveAlias() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListAliases(t *testing.T) {
	// Setup
	tmpDir := t.TempDir()
	MagicAliasPath = tmpDir

	// Create test aliases
	testAliases := []string{"m", "g", "s"}
	for _, alias := range testAliases {
		err := AddAlias(alias, "git")
		if err != nil {
			t.Fatalf("failed to create test alias: %v", err)
		}
	}

	t.Run("list aliases", func(t *testing.T) {
		aliases, err := ListAliases()
		if err != nil {
			t.Errorf("ListAliases() error = %v", err)
			return
		}

		if len(aliases) != len(testAliases) {
			t.Errorf("expected %d aliases, got %d", len(testAliases), len(aliases))
		}
	})
}
