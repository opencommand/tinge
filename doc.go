// Copyright 2025 The Tinge Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package tinge provides a fluent API for creating styled text output in terminal applications.
// It uses lipgloss for styling and offers a chainable interface for building formatted text.
//
// Example usage:
//
//	package main
//
//	import "github.com/your-username/tinge"
//
//	func main() {
//		// Basic colored text
//		tinge.Styled().
//			Green("Hello, ").
//			Bold("World!").
//			Newline().
//			Write()
//
//		// Complex formatting with indentation
//		tinge.Styled().
//			Bold("Project Status:").
//			Newline().
//			Indent(2).
//			Green("✓ ").
//			Text("Tests passing").
//			Newline().
//			Indent(2).
//			Red("✗ ").
//			Text("Build failed").
//			Newline().
//			Indent(2).
//			Yellow("⚠ ").
//			Text("Warnings found").
//			Write()
//
//		// Using custom styles
//		tinge.Styled().
//			With(tinge.Blue, tinge.Bold).
//			Text("Important notice").
//			Space().
//			With(tinge.Grey).
//			Text("(read carefully)").
//			Write()
//
//		// Building strings for later use
//		message := tinge.Styled().
//			Pink("Welcome to ").
//			BoldItalic("Tinge").
//			Space().
//			BlueDark("v1.0.0").
//			String()
//
//		println(message)
//	}
//
// Output examples:
//
//	Hello, World!
//
//	Project Status:
//	  ✓ Tests passing
//	  ✗ Build failed
//	  ⚠ Warnings found
//
//	Important notice (read carefully)
//
//	Welcome to Tinge v1.0.0
package tinge
