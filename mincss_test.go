package mincss

import "testing"

func TestMinifyCSS(t *testing.T) {
	cssPathes := []string{
		"../pali/tipitaka/app/css/tipitaka-latn.css",
		"../pali/tipitaka/app/css/app.css",
	}
	minifiedCss := MinifyCSS(cssPathes)
	println(minifiedCss)
}
