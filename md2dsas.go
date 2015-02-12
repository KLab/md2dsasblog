package main

import (
	"bytes"
	"flag"
	"github.com/russross/blackfriday"
	"html"
	"io/ioutil"
	"log"
	"os"
)

type renderer struct {
	*blackfriday.Html
}

func newRenderer() *renderer {
	return &renderer{Html: blackfriday.HtmlRenderer(0, "", "").(*blackfriday.Html)}
}

func (r *renderer) Header(out *bytes.Buffer, text func() bool, level int, x string) {
	if level == 1 {
		text()
		out.WriteString("\n")
		return
	}
	r.Html.Header(out, text, level+1, x)
}

func (r *renderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	class := "prog prettyprint lang-" + lang
	if lang == "console" || lang == "terminal" {
		class = "terminal"
	}

	out.WriteString(`<pre class="` + class + `">` + "\n")
	out.WriteString(html.EscapeString(string(text)))
	out.WriteString("</pre>\n")
}

func main() {
	ext := 0
	ext |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	ext |= blackfriday.EXTENSION_TABLES
	ext |= blackfriday.EXTENSION_FENCED_CODE
	ext |= blackfriday.EXTENSION_AUTOLINK
	ext |= blackfriday.EXTENSION_STRIKETHROUGH
	ext |= blackfriday.EXTENSION_SPACE_HEADERS

	var output string
	flag.StringVar(&output, "out", "-", "output file name")
	flag.Parse()

	inFile := os.Stdin
	if flag.NArg() == 1 {
		var err error
		inFile, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	}

	in, err := ioutil.ReadAll(inFile)
	if err != nil {
		log.Fatal(err)
	}
	out := blackfriday.Markdown(in, newRenderer(), ext)
	if output == "-" {
		os.Stdout.Write(out)
		return
	}
	if f, err := os.Create(output); err != nil {
		f.Write(out)
		f.Close()
		return
	}
	log.Fatal(err)
}
