package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yosssi/gohtml"
)

var basic = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>TITLE</title>
</head>
<body>
`

var basicEnd = `
</body>
</html>`

func writeTag(b *strings.Builder, tag string, attrs, children []string) {
	// write tag
	b.WriteRune('<')
	b.WriteString(tag)

	// write attributes
	attrsCount := len(attrs)
	if attrsCount > 0 {
		b.WriteString(" ")
		for i, a := range attrs {
			if strings.Contains(a, "=") {
				attr := strings.Split(a, "=")
				fmt.Fprintf(b, `%s="%s"`, attr[0], attr[1])
			} else {
				b.WriteString(a)
			}
			if i < attrsCount-1 {
				b.WriteRune(' ')
			}
		}
	}

	if len(children) > 0 {
		b.WriteRune('>')
		// write children
		b.WriteString(strings.Join(children, " "))

		fmt.Fprintf(b, "</%s>", tag)
	} else {
		b.WriteString("/>")
	}
}

func copyStdin(b *strings.Builder) {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// stdin is pipe
		io.Copy(b, os.Stdin)
	}
}

func main() {
	attrs := pflag.StringArrayP("attribute", "a", nil, "list of attributes for the tag")
	tpl := pflag.BoolP("basic", "b", false, "basic HTML template")
	title := pflag.StringP("title", "t", "Document", "title for basic template")
	pflag.Parse()

	tag := pflag.Arg(0)
	b := &strings.Builder{}

	if !*tpl {
		copyStdin(b)
		writeTag(b, tag, *attrs, pflag.Args()[1:])
	} else {
		b.WriteString(strings.Replace(basic, "TITLE", *title, 1))
		copyStdin(b)
		b.WriteString(basicEnd)
	}

	fmt.Println(gohtml.Format(b.String()))
}
