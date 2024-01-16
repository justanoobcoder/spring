package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/justanoobcoder/spring/springboot"
	springlist "github.com/justanoobcoder/spring/springlist"
)

type errMsg struct{ error }

func (e errMsg) Error() string { return e.error.Error() }

type state int

const (
	chooseProjectType state = iota
	chooseBootVersion
	chooseLanguage
	inputGroupId
	inputArtifactId
	inputName
	inputDescription
	inputPackageName
	inputVersion
	choosePackaging
	chooseJavaVersion
	chooseDependencies
	downloadFileState
)

type Model struct {
	list         list.Model
	PackageName  string
	Version      string
	BootVersion  string
	GroupId      string
	ArtifactId   string
	Packaging    string
	Description  string
	Language     string
	Name         string
	Type         string
	JavaVersion  string
	springboot   springboot.SpringBoot
	Dependencies []string
	textInput    textinput.Model
	state        state
	quitting     bool
}

func NewModel() *Model {
	sp := springboot.GetSpringBoot()
	l := springlist.NewNormalListModel("Choose Project Type", getProjectTypes(sp), sp.Type.Default, 0, 0)
	ti := textinput.New()
	ti.Placeholder = "Group ID"
	ti.SetValue(sp.GroupID.Default)
	ti.Focus()
	return &Model{
		Packaging:    sp.Packaging.Default,
		JavaVersion:  sp.JavaVersion.Default,
		Language:     sp.Language.Default,
		BootVersion:  sp.BootVersion.Default,
		GroupId:      sp.GroupID.Default,
		ArtifactId:   sp.ArtifactID.Default,
		Name:         sp.Name.Default,
		Description:  sp.Description.Default,
		PackageName:  sp.PackageName.Default,
		Version:      sp.Version.Default,
		Type:         sp.Type.Default,
		Dependencies: []string{},
		list:         l,
		textInput:    ti,
		springboot:   sp,
		state:        chooseProjectType,
		quitting:     false,
	}
}
