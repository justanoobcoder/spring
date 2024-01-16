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
	inputApplicationName
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
	Type         string
	BootVersion  string
	Language     string
	GroupId      string
	ArtifactId   string
	Name         string
	Description  string
	PackageName  string
	Version      string
	Packaging    string
	JavaVersion  string
	Dependencies []string
	springBoot   springboot.SpringBoot
	textInput    textinput.Model
	state        state
	quitting     bool
}

func NewModel() *Model {
	sp, err := springboot.NewSpringBoot()
	if err != nil {
		panic(err)
	}
	l := springlist.NewNormalListModel(
		"Project Type",
		getProjectTypes(sp),
		sp.Type.Default,
		0, 0,
	)
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
		springBoot:   sp,
		state:        chooseProjectType,
		list:         l,
		quitting:     false,
	}
}
