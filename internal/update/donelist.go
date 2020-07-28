package update

import (
	"fmt"
	"strings"
)

// A DoneList is a list of completed tasks.
type DoneList []string

func (dl DoneList) String() string {
	return dl.PrefixedString("")
}

// PrefixedString prints the DoneList with the given prefix on each line.
func (dl DoneList) PrefixedString(prefix string) string {
	var out strings.Builder
	for i, did := range dl {
		fmt.Fprintf(&out, "%s%d: %s\n", prefix, i+1, did)
	}
	return strings.TrimSuffix(out.String(), "\n")
}
