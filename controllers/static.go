package controllers

import (
	"html/template"
	"net/http"

	"gallery/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "How are you doing?",
			Answer:   "I am alright",
		},
		{
			Question: "How are you doing?",
			Answer:   "I am alright",
		},
		{
			Question: "How are you doing?",
			Answer:   "I am alright",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
