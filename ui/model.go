package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/justanoobcoder/spring/springboot"
	springlist "github.com/justanoobcoder/spring/springlist"
)

type state int

const (
	chooseProjectType state = iota
	chooseBootVersion
	chooseLanguage
	inputMetaData
	choosePackaging
	chooseJavaVersion
	chooseDependencies
	downloadFile
)

const (
	groupId = iota
	artifactId
	applicationName
	description
	packageName
	version
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
	message      string
	Dependencies []string
	inputs       []textinput.Model
	springBoot   springboot.SpringBoot
	focused      int
	state        state
	failed       bool
	quitting     bool
}

func NewModel() *Model {
	sp, err := springboot.New()
	if err != nil {
		panic(err)
	}
	l := springlist.NewNormalListModel(
		"Project Type",
		getProjectTypes(sp),
		sp.Type.Default,
		0, 0,
	)

	inputs := make([]textinput.Model, 6)

	inputs[groupId] = textinput.New()
	inputs[groupId].Placeholder = sp.GroupID.Default
	inputs[groupId].SetValue(sp.GroupID.Default)
	inputs[groupId].Focus()
	inputs[groupId].Width = 60
	inputs[groupId].Prompt = ""

	inputs[artifactId] = textinput.New()
	inputs[artifactId].Placeholder = sp.ArtifactID.Default
	inputs[artifactId].SetValue(sp.ArtifactID.Default)
	inputs[artifactId].Focus()
	inputs[artifactId].Width = 60
	inputs[artifactId].Prompt = ""

	inputs[applicationName] = textinput.New()
	inputs[applicationName].Placeholder = sp.Name.Default
	inputs[applicationName].SetValue(sp.Name.Default)
	inputs[applicationName].Focus()
	inputs[applicationName].Width = 60
	inputs[applicationName].Prompt = ""

	inputs[description] = textinput.New()
	inputs[description].Placeholder = sp.Description.Default
	inputs[description].SetValue(sp.Description.Default)
	inputs[description].Focus()
	inputs[description].Width = 60
	inputs[description].Prompt = ""

	inputs[packageName] = textinput.New()
	inputs[packageName].Placeholder = sp.PackageName.Default
	inputs[packageName].SetValue(sp.PackageName.Default)
	inputs[packageName].Focus()
	inputs[packageName].Width = 60
	inputs[packageName].Prompt = ""

	inputs[version] = textinput.New()
	inputs[version].Placeholder = sp.Version.Default
	inputs[version].SetValue(sp.Version.Default)
	inputs[version].Focus()
	inputs[version].Width = 60
	inputs[version].Prompt = ""

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
		inputs:       inputs,
		focused:      0,
		failed:       false,
		quitting:     false,
	}
}
