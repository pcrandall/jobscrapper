package main

import "github.com/charmbracelet/bubbles/textinput"

type model struct {
	index         int
	keywordInput  textinput.Model
	locationInput textinput.Model
	submitButton  string
}

type extractedJob struct {
	Id       string
	Title    string
	Location string
	Salary   string
	Summary  string
	Date     string
	FullDesc string
}

type SearchConfig struct {
	Baseurl    string `yaml:"baseurl"`
	Baselimit  string `yaml:"baselimit"`
	Maxresults int    `yaml:"maxresults"`
	Jobs       []struct {
		Job      interface{} `yaml:"job"`
		Keyword  string      `yaml:"keyword"`
		Location []string    `yaml:"location"`
	} `yaml:"jobs"`
}
