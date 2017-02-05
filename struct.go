package main

type Question struct {
	Title string
	Explanation string
	Type string // checkbox(string), dropdown(string), string
	Examplelist []string
	Answer string
}

type Survey struct {
	Title string
	Questions []Question
}
