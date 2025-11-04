package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ColorTheme struct {
	Primary    lipgloss.AdaptiveColor
	Secondary  lipgloss.AdaptiveColor
	Accent     lipgloss.AdaptiveColor
	Background lipgloss.AdaptiveColor
	Foreground lipgloss.AdaptiveColor
	Success    lipgloss.AdaptiveColor
	Error      lipgloss.AdaptiveColor
}

var (
	DefaultTheme = ColorTheme{
		Primary:   lipgloss.AdaptiveColor{Dark: "#1e1e2e", Light: "#6c7086"},
		Secondary: lipgloss.AdaptiveColor{Dark: "#6c7086", Light: "#acb0be"},
		Accent:    lipgloss.AdaptiveColor{Dark: "#89b4fa", Light: "#1e66f5"},
	}

	CattpuccinMocha = ColorTheme{
		Primary:    lipgloss.AdaptiveColor{Dark: "#1e1e2e", Light: "#eff1f5"},
		Secondary:  lipgloss.AdaptiveColor{Dark: "#6c7086", Light: "#9ca0b0"},
		Accent:     lipgloss.AdaptiveColor{Dark: "#89b4fa", Light: "#1e66f5"},
		Background: lipgloss.AdaptiveColor{Dark: "#181825", Light: "#dce0e8"},
		Foreground: lipgloss.AdaptiveColor{Dark: "#cdd6f4", Light: "#4c4f69"},
		Success:    lipgloss.AdaptiveColor{Dark: "#a6e3a1", Light: "#40a02b"},
		Error:      lipgloss.AdaptiveColor{Dark: "#f38ba8", Light: "#d20f39"},
	}

	Dracula = ColorTheme{
		Primary:    lipgloss.AdaptiveColor{Dark: "#282a36", Light: "#f8f8f2"},
		Secondary:  lipgloss.AdaptiveColor{Dark: "#44475a", Light: "#6272a4"},
		Accent:     lipgloss.AdaptiveColor{Dark: "#bd93f9", Light: "#6272a4"},
		Background: lipgloss.AdaptiveColor{Dark: "#1e1f29", Light: "#f1f2f8"},
		Foreground: lipgloss.AdaptiveColor{Dark: "#f8f8f2", Light: "#282a36"},
		Success:    lipgloss.AdaptiveColor{Dark: "#50fa7b", Light: "#2ecc71"},
		Error:      lipgloss.AdaptiveColor{Dark: "#ff5555", Light: "#ff6b6b"},
	}

	Gruvbox = ColorTheme{
		Primary:    lipgloss.AdaptiveColor{Dark: "#282828", Light: "#fbf1c7"},
		Secondary:  lipgloss.AdaptiveColor{Dark: "#3c3836", Light: "#ebdbb2"},
		Accent:     lipgloss.AdaptiveColor{Dark: "#d79921", Light: "#b57614"},
		Background: lipgloss.AdaptiveColor{Dark: "#1d2021", Light: "#f9f5d7"},
		Foreground: lipgloss.AdaptiveColor{Dark: "#ebdbb2", Light: "#3c3836"},
		Success:    lipgloss.AdaptiveColor{Dark: "#b8bb26", Light: "#79740e"},
		Error:      lipgloss.AdaptiveColor{Dark: "#fb4934", Light: "#9d0006"},
	}

	Nord = ColorTheme{
		Primary:    lipgloss.AdaptiveColor{Dark: "#2e3440", Light: "#e5e9f0"},
		Secondary:  lipgloss.AdaptiveColor{Dark: "#3b4252", Light: "#d8dee9"},
		Accent:     lipgloss.AdaptiveColor{Dark: "#88c0d0", Light: "#5e81ac"},
		Background: lipgloss.AdaptiveColor{Dark: "#2e3440", Light: "#eceff4"},
		Foreground: lipgloss.AdaptiveColor{Dark: "#e5e9f0", Light: "#2e3440"},
		Success:    lipgloss.AdaptiveColor{Dark: "#a3be8c", Light: "#8fbcbb"},
		Error:      lipgloss.AdaptiveColor{Dark: "#bf616a", Light: "#bf616a"},
	}

	Solarized = ColorTheme{
		Primary:    lipgloss.AdaptiveColor{Dark: "#002b36", Light: "#fdf6e3"},
		Secondary:  lipgloss.AdaptiveColor{Dark: "#073642", Light: "#eee8d5"},
		Accent:     lipgloss.AdaptiveColor{Dark: "#268bd2", Light: "#268bd2"},
		Background: lipgloss.AdaptiveColor{Dark: "#002b36", Light: "#fdf6e3"},
		Foreground: lipgloss.AdaptiveColor{Dark: "#93a1a1", Light: "#657b83"},
		Success:    lipgloss.AdaptiveColor{Dark: "#859900", Light: "#859900"},
		Error:      lipgloss.AdaptiveColor{Dark: "#dc322f", Light: "#dc322f"},
	}
)

var (
	inactiveTabBorder = lipgloss.Border{Bottom: "─", BottomLeft: "─", BottomRight: "─"}
	activeTabBorder   = lipgloss.Border{Top: "─", Bottom: " ", Left: "│", Right: "│", TopLeft: "╭", TopRight: "╮", BottomLeft: "┘", BottomRight: "└"}
	tabGapBorderLeft  = lipgloss.Border{Bottom: "─", BottomLeft: "╭", BottomRight: "─"}
	tabGapBorderRight = lipgloss.Border{Bottom: "─", BottomLeft: "─", BottomRight: "╮"}
	docStyle          = lipgloss.NewStyle().Padding(0).Align(lipgloss.Center)
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Border(activeTabBorder, true)
	tabGapLeft        = inactiveTabStyle.Border(tabGapBorderLeft, true)
	tabGapRight       = inactiveTabStyle.Border(tabGapBorderRight, true)
	windowStyle       = lipgloss.NewStyle().Padding(2, 2).Margin(4, 4).Align(lipgloss.Center).Border(lipgloss.NormalBorder())
	wallStyle         = lipgloss.NewStyle().Foreground(DefaultTheme.Secondary)
	// playerStyle       = lipgloss.NewStyle().Foreground(DefaultTheme.Accent)
)

func SetCurrentTheme(t ColorTheme) func() tea.Msg {
	return func() tea.Msg {
		inactiveTabStyle = inactiveTabStyle.BorderForeground(t.Accent)
		activeTabStyle = activeTabStyle.BorderForeground(t.Accent)
		tabGapLeft = tabGapLeft.BorderForeground(t.Accent)
		tabGapRight = tabGapRight.BorderForeground(t.Accent)
		windowStyle = windowStyle.BorderForeground(t.Secondary)
		wallStyle = wallStyle.Foreground(t.Secondary)
		return nil
	}
}
