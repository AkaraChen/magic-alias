package shell

import (
	"testing"
)

func TestRemoveScriptContent(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name: "remove script content",
			content: "Some content before\n" +
				ScriptFirstLine + "\n" +
				"export PATH=\"$PATH:$HOME/.magic-alias\"\n" +
				"eval \"$(ma completion bash)\"\n" +
				ScriptFinalLine + "\n" +
				"Some content after",
			expected: "Some content before\nSome content after",
		},
		{
			name: "no script content",
			content: "Some content\n" +
				"No script markers here\n" +
				"Just regular text",
			expected: "Some content\nNo script markers here\nJust regular text",
		},
		{
			name:     "empty content",
			content:  "",
			expected: "",
		},
		{
			name: "only first line no final line",
			content: "Some content\n" +
				ScriptFirstLine + "\n" +
				"export PATH=\"$PATH:$HOME/.magic-alias\"\n" +
				"Some more content",
			expected: "Some content\nSome more content",
		},
		{
			name: "only final line no first line",
			content: "Some content\n" +
				"export PATH=\"$PATH:$HOME/.magic-alias\"\n" +
				ScriptFinalLine + "\n" +
				"Some more content",
			expected: "Some content\nexport PATH=\"$PATH:$HOME/.magic-alias\"\nSome more content",
		},
		{
			name: "multiple script blocks",
			content: "Before first block\n" +
				ScriptFirstLine + "\n" +
				"First block content\n" +
				ScriptFinalLine + "\n" +
				"Between blocks\n" +
				ScriptFirstLine + "\n" +
				"Second block content\n" +
				ScriptFinalLine + "\n" +
				"After second block",
			expected: "Before first block\nBetween blocks\nAfter second block",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveScriptContent(tt.content)
			if result != tt.expected {
				t.Errorf("expected:\n%q\ngot:\n%q", tt.expected, result)
			}
		})
	}
}
