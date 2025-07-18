# Tinge

A lightweight Go library for creating beautifully styled terminal text output with an intuitive fluent API.

## Features

- 🎨 **Rich Text Styling**: Support for colors, bold, italic, and combined styles
- 🔤 **Fluent API**: Chain methods for readable and expressive code
- 📝 **Flexible Output**: Write to any `io.Writer` or use default stdout
- 🎯 **Simple Integration**: Easy to integrate into any Go project
- 🌈 **Dracula Theme**: Built-in color palette inspired by Dracula theme

## Installation

```bash
go get github.com/opencommand/tinge
```

## Quick Start

```go
package main

import "github.com/opencommand/tinge"

func main() {
    tinge.Styled().
        Bold("Hello, ").
        Green("World!").
        Newline().
        Grey("This is a ").
        Italic("styled").
        Grey(" message.").
        Write()
}
```

## Usage

### Basic Styling

```go
// Create a new styled text instance
text := tinge.Styled()

// Add styled content
text.Bold("Bold text").
    Space().
    Italic("Italic text").
    Newline().
    Red("Red text").
    Green("Green text")
```

### Available Colors

- `Grey()` - Light grey text
- `GreyDark()` - Dark grey text
- `Red()` - Red text
- `Green()` - Bright green text
- `GreenLight()` - Light green text
- `GreenDark()` - Dark green text
- `Pink()` - Pink text
- `Yellow()` - Yellow text
- `Blue()` - Light blue text
- `BlueDark()` - Dark blue text

### Text Formatting

- `Bold()` - Bold text
- `Italic()` - Italic text
- `BoldItalic()` - Bold and italic text

### Layout Control

```go
tinge.Styled().
    Text("First line").
    Newline().
    Indent(4).
    Text("Indented line").
    Newline().
    Space(8).
    Text("Spaced text")
```

### Custom Styling

Use the `With()` method to apply custom styles:

```go
customStyle := lipgloss.NewStyle().Underline(true)
tinge.Styled().
    With(customStyle).
    Text("Custom styled text")
```

### Output Control

```go
// Write to stdout (default)
text.Write()

// Write to a specific writer
var buf strings.Builder
text.Write(&buf)

// Set default output writer
tinge.SetWriter(os.Stderr)
```

### Complete Example

```go
package main

import (
    "github.com/opencommand/tinge"
)

func main() {
    tinge.Styled().
        Bold("Welcome to ").
        Blue("Tinge").
        Bold("!").
        Newline().
        Newline().
        Grey("This library provides:").
        Newline().
        Indent(2).
        Green("✓ ").
        Text("Beautiful terminal output").
        Newline().
        Indent(2).
        Green("✓ ").
        Text("Fluent API design").
        Newline().
        Indent(2).
        Green("✓ ").
        Text("Easy integration").
        Newline().
        Newline().
        Yellow("Get started today!").
        Write()
}
```

## API Reference

### Core Methods

- `Styled()` - Create a new styled text instance
- `Text(string)` - Add plain text
- `Space(int?)` - Add spaces (default: 1)
- `Newline()` - Add a newline
- `Indent(int)` - Set indentation for subsequent lines
- `Write(...io.Writer)` - Output the styled text

### Color Methods

- `Grey(string)` - Light grey text
- `GreyDark(string)` - Dark grey text
- `Red(string)` - Red text
- `Green(string)` - Bright green text
- `GreenLight(string)` - Light green text
- `GreenDark(string)` - Dark green text
- `Pink(string)` - Pink text
- `Yellow(string)` - Yellow text
- `Blue(string)` - Light blue text
- `BlueDark(string)` - Dark blue text

### Formatting Methods

- `Bold(string)` - Bold text
- `Italic(string)` - Italic text
- `BoldItalic(string)` - Bold and italic text

### Advanced Methods

- `With(...TextStyle)` - Apply custom styles
- `String()` - Get the rendered string
- `SetWriter(io.Writer)` - Set default output writer

## Dependencies

- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling library

## Requirements

- Go 1.23.5 or later

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Examples

Check out the examples in the repository for more usage patterns and advanced features. 