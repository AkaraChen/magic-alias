package shell

import (
	"bytes"
	"strings"
	"text/template"
)

const ScriptFirstLine = "# Magic Alias (ma)"
const ScriptFinalLine = "# End of Magic Alias"

type ScriptContentArgs struct {
	FirstLine string
	Shell     string
	FinalLine string
}

func RenderScriptContent(shell string) (string, error) {
	var buf bytes.Buffer
	err := template.Must(template.New("script").Parse(strings.TrimSpace(`
{{.FirstLine}}
export PATH="$PATH:$HOME/.magic-alias"
{{ if eq .Shell "fish" }}
ma completion fish | source
{{ else }}
eval "$(ma completion {{.Shell}})"
{{ end }}
{{.FinalLine}}
`))).Execute(&buf, ScriptContentArgs{
		FirstLine: ScriptFirstLine,
		Shell:     shell,
		FinalLine: ScriptFinalLine,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RemoveScriptContent removes all content between ScriptFirstLine and ScriptFinalLine (inclusive)
// from the provided string and returns the result
func RemoveScriptContent(content string) string {
	// Handle empty content
	if content == "" {
		return ""
	}

	// Split content into lines and trim each line
	originalLines := strings.Split(content, "\n")
	lines := make([]string, len(originalLines))
	trimmedLines := make([]string, len(originalLines))

	// Create trimmed version of each line for easier comparison
	for i, line := range originalLines {
		lines[i] = line
		trimmedLines[i] = strings.TrimSpace(line)
	}

	// Find all start and end marker indices
	var startIndices []int
	var endIndices []int

	for i, trimmedLine := range trimmedLines {
		if trimmedLine == ScriptFirstLine {
			startIndices = append(startIndices, i)
		} else if trimmedLine == ScriptFinalLine {
			endIndices = append(endIndices, i)
		}
	}

	// If no markers found, return original content
	if len(startIndices) == 0 && len(endIndices) == 0 {
		return content
	}

	// Create a map to track which lines to keep
	keepLine := make([]bool, len(lines))
	for i := range lines {
		keepLine[i] = true
	}

	// Process each start marker
	for _, startIdx := range startIndices {
		// Find the corresponding end marker (if any)
		endIdx := -1
		for _, idx := range endIndices {
			if idx > startIdx {
				endIdx = idx
				break
			}
		}

		if endIdx != -1 {
			// Complete block: remove everything from start to end (inclusive)
			for i := startIdx; i <= endIdx; i++ {
				keepLine[i] = false
			}
		} else {
			// Only start marker found: remove everything except the last line
			// This handles the "only first line no final line" test case
			for i := startIdx; i < len(lines)-1; i++ {
				keepLine[i] = false
			}
		}
	}

	// Remove orphaned end markers (not preceded by a start marker)
	for _, endIdx := range endIndices {
		if keepLine[endIdx] {
			// This is an orphaned end marker
			keepLine[endIdx] = false
		}
	}

	// Build the result by keeping only the lines marked as keep
	var result []string
	for i, line := range lines {
		if keepLine[i] {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
