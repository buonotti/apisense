package tui

import (
	"fmt"
	"github.com/buonotti/apisense/errors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type editConfigModel struct {
	keymap    keymap
	err       error
	textInput textinput.Model
}

func EditConfigModel() tea.Model {

	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.TextStyle = styleHelp.Copy()

	return editConfigModel{
		keymap:    DefaultKeyMap,
		err:       nil,
		textInput: ti,
	}
}

func (e editConfigModel) Init() tea.Cmd {
	return nil
}

func (e editConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if choiceConfigModel != "selectConfigModel" {
		if updateEditConfigField {
			ti := textinput.New()
			ti.Placeholder = selectedField
			ti.Focus()
			ti.CharLimit = 156
			ti.Width = 20
			e.textInput = ti
			e.textInput.SetValue(viper.GetString(selectedField))
			updateEditConfigField = false
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, e.keymap.back):
				if choiceConfigModel == "editConfigModel" {
					choiceConfigModel = "selectConfigModel"
				}
			case key.Matches(msg, e.keymap.quit):
				return e, tea.Quit
			case key.Matches(msg, e.keymap.choose):
				if choiceConfigModel == "editConfigModel" {
					choiceConfigModel = "selectConfigModel"
				}
				viper.Set(selectedField, e.textInput.Value())
				err := viper.WriteConfig()
				errors.CheckErr(err)
				updateSelectConfigRows = true
				return e, cmd
			}

		case errMsg:
			e.err = msg
			errors.CheckErr(e.err)
		}

		e.textInput, cmd = e.textInput.Update(msg)
		return e, cmd
	}
	return e, nil
}

func (e editConfigModel) View() string {
	return lipgloss.NewStyle().PaddingRight(13).PaddingTop(2).MarginLeft(13).Render(fmt.Sprintf("%v", selectedField) + "\n" +
		e.textInput.View() + "\n")
}
