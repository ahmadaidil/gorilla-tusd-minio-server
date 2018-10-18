package view

import "github.com/thedevsaddam/renderer"

var render *renderer.Render

func init() {
	render = renderer.New(renderer.Options{
		ParseGlobPattern: "template/*.html",
	})
}
