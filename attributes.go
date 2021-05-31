package tables

import (
	"fmt"
	"strings"
)

type Attributes map[string]string

func (a Attributes) Slice() []string {
	var attrs []string
	for k, v := range a {
		if len(v) > 0 {
			attrs = append(attrs, fmt.Sprintf("%s=%q", k, v))
		} else {
			attrs = append(attrs, k)
		}
	}
	return attrs
}

func (a Attributes) String() string {
	return strings.Join(a.Slice(), ` `)
}
