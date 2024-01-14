package ui

type Model struct {
	packaging    string
	javaVersion  string
	language     string
	bootVersion  string
	groupId      string
	artifactId   string
	name         string
	description  string
	packageName  string
	version      string
	typE         string
	dependencies []string
}

func NewModel() *Model {
	return &Model{
		packaging:    "jar",
		javaVersion:  "17",
		language:     "java",
		bootVersion:  "3.2.1",
		groupId:      "com.example",
		artifactId:   "demo",
		name:         "demo",
		description:  "Demo project for Spring Boot",
		packageName:  "com.example.demo",
		version:      "0.0.1-SNAPSHOT",
		typE:         "maven-project",
		dependencies: []string{"web", "lombok"},
	}
}
