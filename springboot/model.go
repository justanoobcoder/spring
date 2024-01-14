package springboot

import (
	"io"
	"log"
	"net/http"
	"time"
)

type SpringBoot struct {
	GroupID      ArtifactID             `json:"groupId"`
	PackageName  ArtifactID             `json:"packageName"`
	Description  ArtifactID             `json:"description"`
	Name         ArtifactID             `json:"name"`
	Version      ArtifactID             `json:"version"`
	ArtifactID   ArtifactID             `json:"artifactId"`
	Links        SpringBootLinks        `json:"_links"`
	Packaging    BootVersion            `json:"packaging"`
	BootVersion  BootVersion            `json:"bootVersion"`
	Language     BootVersion            `json:"language"`
	JavaVersion  BootVersion            `json:"javaVersion"`
	Type         Type                   `json:"type"`
	Dependencies SpringBootDependencies `json:"dependencies"`
}

type ArtifactID struct {
	Type    string `json:"type"`
	Default string `json:"default"`
}

type BootVersion struct {
	Type    string             `json:"type"`
	Default string             `json:"default"`
	Values  []BootVersionValue `json:"values"`
}

type BootVersionValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SpringBootDependencies struct {
	Type   string              `json:"type"`
	Values []DependenciesValue `json:"values"`
}

type DependenciesValue struct {
	Name   string       `json:"name"`
	Values []ValueValue `json:"values"`
}

type ValueValue struct {
	VersionRange *string     `json:"versionRange,omitempty"`
	Links        *ValueLinks `json:"_links,omitempty"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
}

type ReferenceUnion struct {
	ReferenceClass *ReferenceClass
	HomeArray      []Home
}

type ValueLinks struct {
	Reference *ReferenceUnion `json:"reference"`
	Guide     *Guide          `json:"guide"`
	Home      *Home           `json:"home,omitempty"`
	Sample    *Home           `json:"sample,omitempty"`
}

type Home struct {
	Title *string `json:"title,omitempty"`
	Href  string  `json:"href"`
}

type ReferenceClass struct {
	Templated *bool   `json:"templated,omitempty"`
	Title     *string `json:"title,omitempty"`
	Href      string  `json:"href"`
}

type SpringBootLinks struct {
	GradleProject       GradleBuildClass `json:"gradle-project"`
	GradleProjectKotlin GradleBuildClass `json:"gradle-project-kotlin"`
	GradleBuild         GradleBuildClass `json:"gradle-build"`
	MavenProject        GradleBuildClass `json:"maven-project"`
	MavenBuild          GradleBuildClass `json:"maven-build"`
	Dependencies        GradleBuildClass `json:"dependencies"`
}

type GradleBuildClass struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type Type struct {
	Type    string      `json:"type"`
	Default string      `json:"default"`
	Values  []TypeValue `json:"values"`
}

type TypeValue struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Action      string `json:"action"`
	Tags        Tags   `json:"tags"`
}

type Tags struct {
	Build   string  `json:"build"`
	Dialect *string `json:"dialect,omitempty"`
	Format  string  `json:"format"`
}

type Guide struct {
	Home      *Home
	HomeArray []Home
}

func GetSpringBoot() SpringBoot {
	const url = "https://start.spring.io"
	http.DefaultClient.Timeout = 10 * time.Second
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	springBoot, err := UnmarshalSpringBoot(body)
	if err != nil {
		log.Fatal(err)
	}
	return springBoot
}
