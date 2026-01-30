/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package color

import (
	"strings"
)

// ANSI color codes
const (
	Reset  = "\x1b[0m"
	Bold   = "\x1b[1m"
	Cyan   = "\x1b[36m"
	Green  = "\x1b[32m"
	Yellow = "\x1b[33m"
)

// ApplyColors adds ANSI color codes to kubectl describe output based on indentation levels.
// This function is called AFTER tabwriter has formatted the output, so alignment is preserved.
func ApplyColors(text string) string {
	var result strings.Builder
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Count leading spaces to determine indentation level
		indent := countLeadingSpaces(line)
		level := indent / 2 // Each level is 2 spaces

		// Apply color based on indentation level
		coloredLine := colorizeLine(line, level)
		result.WriteString(coloredLine)
		result.WriteString("\n")
	}

	return result.String()
}

// countLeadingSpaces counts the number of leading spaces in a line
func countLeadingSpaces(line string) int {
	count := 0
	for _, ch := range line {
		if ch == ' ' {
			count++
		} else {
			break
		}
	}
	return count
}

// colorizeLine applies color to a line based on its indentation level
func colorizeLine(line string, level int) string {
	// Don't color empty lines
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return line
	}

	// Extract the indentation (leading spaces)
	indent := line[:len(line)-len(strings.TrimLeft(line, " "))]

	// The rest of the line after indentation
	content := line[len(indent):]

	// Apply color based on level
	var color string

	// Detect continuation lines (deep indentation for alignment, likely wrapped values)
	// Normal nesting levels are 0, 1, 2, 3 (0, 2, 4, 6 spaces)
	// Anything beyond level 4 (8+ spaces) is likely a continuation line for alignment
	if level >= 4 {
		// This is a continuation line (e.g., wrapped Labels, Tolerations values)
		// Use the same color as level 0 (top-level fields)
		color = Bold + Cyan
	} else {
		switch level {
		case 0:
			// Top level (Name, Namespace, etc.) - Bold Cyan
			color = Bold + Cyan
		case 1:
			// First level sections - Bold
			color = Bold
		case 2:
			// Second level - Green
			color = Green
		case 3:
			// Third level - Yellow
			color = Yellow
		}
	}

	// Apply color to the content, preserve indentation
	return indent + color + content + Reset
}
