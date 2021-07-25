package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yosssi/gohtml"
)

func main() {
	attrs := pflag.StringArrayP("attribute", "a", nil, "list of attributes for the tag")
	pflag.Parse()

	tag := pflag.Arg(0)
	b := &strings.Builder{}

	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// stdin is pipe
		io.Copy(b, os.Stdin)
	}

	// write tag
	b.WriteRune('<')
	b.WriteString(tag)

	// write attributes
	attrsCount := len(*attrs)
	if attrsCount > 0 {
		b.WriteString(" ")
		for i, a := range *attrs {
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

	if len(pflag.Args()) > 1 {
		b.WriteRune('>')

		// write children
		b.WriteString(strings.Join(pflag.Args()[1:], " "))

		fmt.Fprintf(b, "<%s/>", tag)
	} else {
		b.WriteString("/>")
	}

	fmt.Println(gohtml.Format(b.String()))
}
