package shell

import (
	"os"
	"testing"
)

func TestGetShell(t *testing.T) {
	// Save the original SHELL value to restore it later
	originalShell := os.Getenv("SHELL")
	defer os.Setenv("SHELL", originalShell)

	tests := []struct {
		name        string
		shellValue  string
		expectError bool
	}{
		{
			name:        "valid shell path",
			shellValue:  "/bin/bash",
			expectError: false,
		},
		{
			name:        "empty shell path",
			shellValue:  "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variable to test value
			os.Setenv("SHELL", tc.shellValue)

			// Call the function
			shell, err := GetShell()

			// Check error condition
			if tc.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			// If no error expected, check the returned shell
			if !tc.expectError {
				if shell != tc.shellValue {
					t.Errorf("expected shell to be %q, got %q", tc.shellValue, shell)
				}
			}
		})
	}
}

func TestGetShellName(t *testing.T) {
	// Save the original SHELL value to restore it later
	originalShell := os.Getenv("SHELL")
	defer os.Setenv("SHELL", originalShell)

	tests := []struct {
		name        string
		shellValue  string
		expected    string
		expectError bool
	}{
		{
			name:        "bash shell",
			shellValue:  "/bin/bash",
			expected:    "bash",
			expectError: false,
		},
		{
			name:        "zsh shell",
			shellValue:  "/usr/bin/zsh",
			expected:    "zsh",
			expectError: false,
		},
		{
			name:        "shell with multiple slashes",
			shellValue:  "/usr/local/bin/fish",
			expected:    "fish",
			expectError: false,
		},
		{
			name:        "empty shell path",
			shellValue:  "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variable to test value
			os.Setenv("SHELL", tc.shellValue)

			// Call the function
			shellName, err := GetShellName()

			// Check error condition
			if tc.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			// If no error expected, check the returned shell name
			if !tc.expectError {
				if shellName != tc.expected {
					t.Errorf("expected shell name to be %q, got %q", tc.expected, shellName)
				}
			}
		})
	}
}
