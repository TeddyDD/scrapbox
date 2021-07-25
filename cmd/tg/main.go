package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/yosssi/gohtml"
)

func main() {
	attrs := pflag.StringArrayP("attribute", "a", nil, "list of attributes for the tag")
	pflag.Parse()

	tag := pflag.Arg(0)
	b := &strings.Builder{}

	b.WriteRune('<')
	b.WriteString(tag)

	if len(*attrs) > 0 {
		b.WriteString(" ")
		for _, a := range *attrs {
			if strings.Contains(a, "=") {
				attr := strings.Split(a, "=")
				fmt.Fprintf(b, `%s="%s"`, attr[0], attr[1])
			} else {
				b.WriteString(a)
			}
			b.WriteRune(' ')
		}
	}

	if len(pflag.Args()) > 1 {
		b.WriteRune('>')
		b.WriteString(strings.Join(pflag.Args()[1:], " "))
		fmt.Fprintf(b, "<%s/>", tag)
	} else {
		b.WriteString("/>")
	}

	fmt.Println(gohtml.Format(b.String()))
}
