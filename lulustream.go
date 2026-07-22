// Package lulustream provides a wrapper for lulustream.com
package lulustream

import (
	"fmt"
	"html"
	"os"
	"os/exec"
	"regexp"

	"github.com/dop251/goja"
	"github.com/kissanjamgit/ext"
	"resty.dev/v3"
)

var header = map[string]string{
	`User-Agent`:      `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:152.0) Gecko/20100101 Firefox/152.0`,
	`Accept`:          `*/*`,
	`Accept-Language`: `en-US,en;q=0.9`,
}

type Lulustream struct {
	Source string
}

func New(source string) ext.Site {
	return &Lulustream{Source: source}
}

func (l *Lulustream) Resource(client *resty.Client) (cr ext.ContentResource, err error) {
	res, err := client.R().SetHeaders(header).Get(l.Source)
	if err != nil {
		return
	}

	scrabbledFunSubmatch := regexp.MustCompile(`<script type='text/javascript'>eval\((.*)\)\n`).FindStringSubmatch(res.String())
	if len(scrabbledFunSubmatch) < 2 {
		err = fmt.Errorf(`len(scrabbledFunSubmatch) < 2 `)
		return
	}

	ja := goja.New()
	ja.Set("console", map[string]any{
		"log": func(call goja.FunctionCall) goja.Value {
			return nil
		},
	})
	value, err := ja.RunString(`(` + scrabbledFunSubmatch[1] + `)`)
	if err != nil {
		return
	}
	var URL string
	if urlSubmatch := regexp.MustCompile(`file:"([^"]+\.m3u8[^"]*)"`).FindStringSubmatch(value.String()); len(urlSubmatch) < 2 {
		err = fmt.Errorf("len(urlSubmatch) <=2")
		return
	} else {
		URL = urlSubmatch[1]
	}

	var Name string

	nameRe := regexp.MustCompile(`<h1\s+class="h5">([^<]+)<`)
	if nameSubmatch := nameRe.FindStringSubmatch(res.String()); len(nameSubmatch) < 2 {
		err = fmt.Errorf("len(nameSubmatch) <=2")
		return
	} else {
		Name = nameSubmatch[1]
	}
	Name = html.UnescapeString(Name)

	cr = ext.ContentResource{URL: URL, Name: Name}
	return
}

func (l *Lulustream) Download(cr ext.ContentResource) (err error) {
	cmd := exec.Command("yt-dlp.exe", cr.URL, "-o", cr.Name+`.mp4`)
	var arg []string
	for k, v := range header {
		arg = append(arg, "--add-header", fmt.Sprintf("%s:%s", k, v))
	}

	cmd.Args = append(cmd.Args, arg...)

	fmt.Println(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	return
}
