// Copyright 2025 The Tinge Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package tinge

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"testing"
)

// stripANSI removes ANSI escape sequences from a string
func stripANSI(s string) string {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return ansiRegex.ReplaceAllString(s, "")
}

func TestStyled(t *testing.T) {
	styled := Styled()
	if styled == nil {
		t.Error("Styled() should not return nil")
	}
}

func TestBasicColors(t *testing.T) {
	tests := []struct {
		name     string
		method   func(string) *StyledText
		expected string
	}{
		{"Grey", func(s string) *StyledText { return Styled().Grey(s) }, "grey text"},
		{"GreyDark", func(s string) *StyledText { return Styled().GreyDark(s) }, "dark grey text"},
		{"Red", func(s string) *StyledText { return Styled().Red(s) }, "red text"},
		{"Green", func(s string) *StyledText { return Styled().Green(s) }, "green text"},
		{"GreenLight", func(s string) *StyledText { return Styled().GreenLight(s) }, "light green text"},
		{"GreenDark", func(s string) *StyledText { return Styled().GreenDark(s) }, "dark green text"},
		{"Pink", func(s string) *StyledText { return Styled().Pink(s) }, "pink text"},
		{"Yellow", func(s string) *StyledText { return Styled().Yellow(s) }, "yellow text"},
		{"Blue", func(s string) *StyledText { return Styled().Blue(s) }, "blue text"},
		{"BlueDark", func(s string) *StyledText { return Styled().BlueDark(s) }, "dark blue text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method(tt.expected)
			output := stripANSI(result.String())
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected output to contain '%s', got '%s'", tt.expected, output)
			}
		})
	}
}

func TestTextFormatting(t *testing.T) {
	tests := []struct {
		name     string
		method   func(string) *StyledText
		expected string
	}{
		{"Bold", func(s string) *StyledText { return Styled().Bold(s) }, "bold text"},
		{"Italic", func(s string) *StyledText { return Styled().Italic(s) }, "italic text"},
		{"BoldItalic", func(s string) *StyledText { return Styled().BoldItalic(s) }, "bold italic text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.method(tt.expected)
			output := stripANSI(result.String())
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Expected output to contain '%s', got '%s'", tt.expected, output)
			}
		})
	}
}

func TestSpace(t *testing.T) {
	result := Styled().Text("hello").Space().Text("world")
	output := stripANSI(result.String())
	expected := "hello world"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestSpaceWithCount(t *testing.T) {
	result := Styled().Text("hello").Space(3).Text("world")
	output := stripANSI(result.String())
	expected := "hello   world"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestNewline(t *testing.T) {
	result := Styled().Text("hello").Newline().Text("world")
	output := stripANSI(result.String())
	expected := "hello\n\nworld"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestIndent(t *testing.T) {
	result := Styled().Indent(2).Text("indented")
	output := stripANSI(result.String())
	expected := "indented"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestIndentWithNewline(t *testing.T) {
	result := Styled().Text("header").Newline().Indent(2).Text("indented")
	output := stripANSI(result.String())
	expected := "header\n\n  indented"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestWith(t *testing.T) {
	result := Styled().With(Red, Bold).Text("red bold")
	output := stripANSI(result.String())
	if !strings.Contains(output, "red bold") {
		t.Errorf("Expected output to contain 'red bold', got '%s'", output)
	}
}

func TestChaining(t *testing.T) {
	result := Styled().
		Bold("Hello").
		Space().
		Green("World").
		Newline().
		Indent(2).
		Red("Indented").
		Space().
		Blue("Text")

	output := stripANSI(result.String())
	expected := "Hello World\n\n  Indented Text"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestString(t *testing.T) {
	result := Styled().Text("hello").Space().Text("world")
	output := stripANSI(result.String())
	expected := "hello world"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestWrite(t *testing.T) {
	var buf bytes.Buffer
	result := Styled().Text("hello").Space().Text("world")
	result.Write(&buf)
	output := buf.String()
	expected := "hello world"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestWriteWithCustomWriter(t *testing.T) {
	var buf bytes.Buffer
	result := Styled().Text("test")
	result.Write(&buf)
	output := buf.String()
	expected := "test"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestSetWriter(t *testing.T) {
	var buf bytes.Buffer

	// Test setting custom writer
	SetWriter(&buf)
	Styled().Text("test").Write()

	output := buf.String()
	expected := "test"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}

	// Reset to stdout
	SetWriter(os.Stdout)
}

func TestMultipleNewlines(t *testing.T) {
	result := Styled().Text("line1").Newline().Newline().Text("line3")
	output := stripANSI(result.String())
	expected := "line1\n\n\nline3"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestEmptyString(t *testing.T) {
	result := Styled()
	output := stripANSI(result.String())
	expected := ""
	if output != expected {
		t.Errorf("Expected empty string, got '%s'", output)
	}
}

func TestComplexFormatting(t *testing.T) {
	result := Styled().
		Bold("Status Report").
		Newline().
		Indent(2).
		Green("✓ ").
		Text("Tests passing").
		Newline().
		Indent(2).
		Red("✗ ").
		Text("Build failed").
		Newline().
		Indent(2).
		Yellow("⚠ ").
		Text("Warnings found")

	output := stripANSI(result.String())
	expected := "Status Report\n\n  ✓ Tests passing\n\n  ✗ Build failed\n\n  ⚠ Warnings found"
	if output != expected {
		t.Errorf("Expected '%s', got '%s'", expected, output)
	}
}

func TestStyledTextBuilder(t *testing.T) {
	builder := Styled().With(Red, Bold)
	if builder == nil {
		t.Error("StyledTextBuilder should not be nil")
	}
}

func TestStyledTextBuilderText(t *testing.T) {
	result := Styled().With(Red, Bold).Text("styled text")
	output := stripANSI(result.String())
	if !strings.Contains(output, "styled text") {
		t.Errorf("Expected output to contain 'styled text', got '%s'", output)
	}
}

func TestPredefinedStyles(t *testing.T) {
	styles := []TextStyle{
		Grey, GreyDark, Green, GreenLight, GreenDark,
		Red, Pink, Yellow, Blue, BlueDark, Bold, Italic,
	}

	for i, style := range styles {
		// Check if the style is properly initialized by testing its string representation
		styleStr := style.String()
		if styleStr == "" {
			t.Errorf("Style at index %d appears to be empty", i)
		}
	}
}
