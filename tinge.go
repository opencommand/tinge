package tinge

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type TextStyle = lipgloss.Style

type StyledText struct {
	parts           []string
	currLine        string
	indent          int
	pendingNewlines int
}

func Styled() *StyledText {
	return &StyledText{}
}

func (s *StyledText) With(styles ...TextStyle) *StyledTextBuilder {
	combined := lipgloss.NewStyle()
	for _, style := range styles {
		combined = combined.Inherit(style)
	}
	return &StyledTextBuilder{parent: s, style: combined}
}

func (s *StyledText) Space(n ...int) *StyledText {
	s.ensureLineStart()
	if len(n) > 0 {
		s.currLine += strings.Repeat(" ", n[0])
	} else {
		s.currLine += " "
	}
	return s
}

func (s *StyledText) Newline() *StyledText {
	if s.currLine != "" {
		s.parts = append(s.parts, s.currLine)
		s.currLine = ""
	}
	s.pendingNewlines++
	return s
}

func (s *StyledText) Indent(spaces int) *StyledText {
	s.indent = spaces
	// if s.currLine == "" {
	// 	s.currLine = strings.Repeat(" ", spaces)
	// }
	return s
}

func (s *StyledText) ensureLineStart() {
	if s.pendingNewlines > 0 {
		if s.currLine != "" {
			s.parts = append(s.parts, s.currLine)
			s.currLine = ""
		}
		// add newline
		for i := 0; i < s.pendingNewlines; i++ {
			s.parts = append(s.parts, "")
		}
		s.pendingNewlines = 0
		s.currLine = strings.Repeat(" ", s.indent)
	}
}

func (s *StyledText) Grey(text string) *StyledText {
	return s.With(Grey).Text(text)
}

func (s *StyledText) GreyDark(text string) *StyledText {
	return s.With(GreyDark).Text(text)
}

func (s *StyledText) Red(text string) *StyledText {
	return s.With(Red).Text(text)
}

func (s *StyledText) Green(text string) *StyledText {
	return s.With(Green).Text(text)
}

func (s *StyledText) GreenLight(text string) *StyledText {
	return s.With(GreenLight).Text(text)
}

func (s *StyledText) GreenDark(text string) *StyledText {
	return s.With(GreenDark).Text(text)
}

func (s *StyledText) Pink(text string) *StyledText {
	return s.With(Pink).Text(text)
}

func (s *StyledText) Yellow(text string) *StyledText {
	return s.With(Yellow).Text(text)
}

func (s *StyledText) Blue(text string) *StyledText {
	return s.With(Blue).Text(text)
}

func (s *StyledText) BlueDark(text string) *StyledText {
	return s.With(BlueDark).Text(text)
}

func (s *StyledText) Bold(text string) *StyledText {
	return s.With(Bold).Text(text)
}

func (s *StyledText) Italic(text string) *StyledText {
	return s.With(Italic).Text(text)
}

func (s *StyledText) BoldItalic(text string) *StyledText {
	return s.With(Bold, Italic).Text(text)
}

func (s *StyledText) Text(text string) *StyledText {
	s.ensureLineStart()
	return s.With().Text(text)
}

func (s *StyledText) String() string {
	s.ensureLineStart()
	if s.currLine != "" {
		s.parts = append(s.parts, s.currLine)
		s.currLine = ""
	}
	return strings.Join(s.parts, "\n")
}

type StyledTextBuilder struct {
	parent *StyledText
	style  lipgloss.Style
}

func (b *StyledTextBuilder) Text(content string) *StyledText {
	b.parent.ensureLineStart()
	b.parent.currLine += b.style.Render(content)
	return b.parent
}

var (
	Grey       = lipgloss.NewStyle().Foreground(lipgloss.Color("#909194"))
	GreyDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("#454e6d"))
	Green      = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B"))
	GreenLight = lipgloss.NewStyle().Foreground(lipgloss.Color("#3fed7b"))
	GreenDark  = lipgloss.NewStyle().Foreground(lipgloss.Color("#3C9258"))
	Red        = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff5555"))
	Pink       = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff79c6"))
	Yellow     = lipgloss.NewStyle().Foreground(lipgloss.Color("#f1fa8c"))
	Blue       = lipgloss.NewStyle().Foreground(lipgloss.Color("#a4ffff"))
	BlueDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("#8be9fd"))

	Bold   = lipgloss.NewStyle().Bold(true)
	Italic = lipgloss.NewStyle().Italic(true)
)
