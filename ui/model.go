package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/justanoobcoder/spring/springboot"
)

type state int

const (
	chooseProjectType state = iota
	chooseBootVersion
	chooseLanguage
	chooseJavaVersion
)

type Model struct {
	list         list.Model
	packageName  string
	version      string
	bootVersion  string
	groupId      string
	artifactId   string
	packaging    string
	description  string
	language     string
	name         string
	typE         string
	javaVersion  string
	springboot   springboot.SpringBoot
	dependencies []string
	choice       string
	textInput    textinput.Model
	state        state
	quitting     bool
}

func NewModel() *Model {
	sp := springboot.GetSpringBoot()
	l := list.New(getProjectTypes(sp), itemDelegate{}, 100, listHeight)
	l.Title = "Choose Project Type"
	items := l.Items()
	for i := range items {
		if items[i].(item).id == sp.Type.Default {
			l.Select(i)
			break
		}
	}
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	ti := textinput.New()
	ti.Focus()
	return &Model{
		packaging:    sp.Packaging.Default,
		javaVersion:  sp.JavaVersion.Default,
		language:     sp.Language.Default,
		bootVersion:  sp.BootVersion.Default,
		groupId:      sp.GroupID.Default,
		artifactId:   sp.ArtifactID.Default,
		name:         sp.Name.Default,
		description:  sp.Description.Default,
		packageName:  sp.PackageName.Default,
		version:      sp.Version.Default,
		typE:         sp.Type.Default,
		dependencies: []string{},
		list:         l,
		textInput:    ti,
		springboot:   sp,
		state:        chooseProjectType,
		quitting:     false,
	}
}
