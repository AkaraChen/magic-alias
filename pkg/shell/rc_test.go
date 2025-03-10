package shell

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	magicAliasLine string
	shell          string
)

func init() {
	var err error
	shell, err = GetShellName()
	if err != nil {
		panic(err)
	}
	magicAliasLine, err = RenderScriptContent(shell)
	if err != nil {
		panic(err)
	}
}

func TestPathInPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		pathEnv  string
		expected bool
	}{
		{"path exists", "/test/path", "/bin:/usr/bin:/test/path", true},
		{"path does not exist", "/test/path", "/bin:/usr/bin", false},
		{"empty path", "", "/bin:/usr/bin", false},
		{"empty PATH", "/test/path", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("PATH", tt.pathEnv)
			result := pathInPath(tt.path)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsMagicAliasInPath(t *testing.T) {
	// Setup
	originalPath := os.Getenv("PATH")
	defer os.Setenv("PATH", originalPath)

	tests := []struct {
		name     string
		pathEnv  string
		expected bool
	}{
		{"in PATH", ":/bin:/usr/bin:" + MagicAliasPath, true},
		{"not in PATH", "/bin:/usr/bin", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("PATH", tt.pathEnv)
			result, err := IsMagicAliasInPath()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestWriteMagicAliasToRc(t *testing.T) {
	// Create a temporary rc file for testing
	tmpDir := t.TempDir()
	tmpRc := filepath.Join(tmpDir, "testrc")

	// Test cases
	tests := []struct {
		name            string
		initialContent  string
		expectedContent string
		expectError     bool
	}{
		{
			name:            "add to empty file",
			initialContent:  "",
			expectedContent: magicAliasLine,
			expectError:     false,
		},
		{
			name:            "add to existing content",
			initialContent:  "existing content\n",
			expectedContent: "existing content\n\n" + magicAliasLine,
			expectError:     false,
		},
		{
			name:            "already exists",
			initialContent:  magicAliasLine,
			expectedContent: magicAliasLine,
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test file
			err := os.WriteFile(tmpRc, []byte(tt.initialContent), 0644)
			if err != nil {
				t.Fatalf("failed to create test rc file: %v", err)
			}

			// Test the function with the temporary rc file path
			err = WriteMagicAliasToRcPath(shell, tmpRc)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			// Verify file contents
			content, err := os.ReadFile(tmpRc)
			if err != nil {
				t.Fatalf("failed to read test rc file: %v", err)
			}
			if string(content) != tt.expectedContent {
				t.Errorf("expected content:\n%q\ngot:\n%q", tt.expectedContent, string(content))
			}
		})
	}
}

func TestRemoveMagicAliasFromRc(t *testing.T) {
	// Create a temporary rc file for testing
	tmpDir := t.TempDir()
	tmpRc := filepath.Join(tmpDir, "testrc")

	// Test cases
	tests := []struct {
		name            string
		initialContent  string
		expectedContent string
		expectError     bool
	}{
		{
			name:            "remove from file with only magic alias",
			initialContent:  magicAliasLine,
			expectedContent: "",
			expectError:     false,
		},
		{
			name:            "remove from file with magic alias and newline",
			initialContent:  "existing content\n" + magicAliasLine,
			expectedContent: "existing content", // No trailing newline due to how RemoveMagicAliasFromRc works
			expectError:     false,
		},
		{
			name:            "remove from file with magic alias at beginning",
			initialContent:  magicAliasLine + "\nexisting content",
			expectedContent: "existing content", // Keeps the newline after removing magicAliasLine
			expectError:     false,
		},
		{
			name:            "file without magic alias",
			initialContent:  "existing content\n",
			expectedContent: "existing content\n",
			expectError:     true,
		},
		{
			name:            "empty file",
			initialContent:  "",
			expectedContent: "",
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup test file
			err := os.WriteFile(tmpRc, []byte(tt.initialContent), 0644)
			if err != nil {
				t.Fatalf("failed to create test rc file: %v", err)
			}

			// Test the function with the temporary rc file path
			err = RemoveMagicAliasFromRc(tmpRc)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			// Verify file contents
			content, err := os.ReadFile(tmpRc)
			if err != nil {
				t.Fatalf("failed to read test rc file: %v", err)
			}
			if string(content) != tt.expectedContent {
				t.Errorf("expected content:\n%q\ngot:\n%q", tt.expectedContent, string(content))
			}
		})
	}
}
