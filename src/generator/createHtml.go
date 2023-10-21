// metabypass
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/generator/createHtml.go
// Original timestamp: 2023/10/21 09:22

package generator

import (
	"metabypass/templates"
	"os"
	"path/filepath"
	"strings"
)

func Create(htmlfile, targetURL string) error {
	var err error

	if !strings.HasSuffix(".html", htmlfile) && !strings.HasSuffix(".htm", htmlfile) {
		htmlfile += ".html"
	}
	// We need to ensure that the directory exists
	if err := os.MkdirAll(filepath.Dir(htmlfile), os.ModePerm); err != nil {
		return err
	}

	// vars to be substituted (only one, so far)
	placeholders := map[string]string{
		"{{ TARGET URL }}": targetURL,
	}
	// Now we generate a very simple html page and css file
	// only twofiles for now
	if err = templates.ProcessEmbeddedAsset("stylesheet.css", filepath.Join(filepath.Dir(htmlfile), "stylesheet.css"), placeholders); err == nil {
		if err = templates.ProcessEmbeddedAsset("source.html", htmlfile, placeholders); err != nil {
			return err
		}
	}
	return err
}
