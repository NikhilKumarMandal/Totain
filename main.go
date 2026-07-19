package main

import (
  "fmt"
  "os"
  tea "charm.land/bubbletea/v2"
)


type model struct {
  msg string

}


func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m model) View() string {
    return m.msg
}

func initialModel() model {
  return model{msg: "Hello, World!"}
}

func main() {
  p := tea.NewProgram(initialModel())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Alas, there's been an error: %v", err)
    os.Exit(1)
  }
  fmt.Println("Welcome to Totain app.")
}