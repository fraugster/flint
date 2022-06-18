package imports

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	one     = regexp.MustCompile(`^import ([a-z_][a-zA-Z0-9_]* )?"[^"]+"(\s+//.*)?$`)
	more    = regexp.MustCompile(`^\t(\. |[a-z_][a-zA-Z0-9_]* )?"([^"]+)"(\s+//.*)?$`)
	comment = regexp.MustCompile(`^\s+//.*`)
)

func oneLiner(in string) error {
	if !one.MatchString(in) {
		return fmt.Errorf("invalid one liner import : %s", in)
	}
	return nil
}

func moreLiner(in string) (string, error) {
	res := more.FindAllStringSubmatch(in, -1)
	if len(res) != 1 {
		return "", fmt.Errorf("invalid multi line import : %s", in)
	}

	return res[0][2], nil
}

func multiLiner(in string) error {
	all := strings.Split(in, "\n")
	if all[0] != "import (" || all[len(all)-1] != ")" {
		return fmt.Errorf("invalid begin or end for import block was : %s", in)
	}

	var state int
	for _, i := range all[1 : len(all)-1] {
		err := handleExpectation(i, &state)
		if err != nil {
			return err
		}
	}

	return nil
}

func handleExpectation(line string, state *int) error {
	// ignore comment line. they are not empty lines
	if comment.MatchString(line) {
		return nil
	}
	if strings.Trim(line, "\n\t ") == "" {
		switch *state {
		case 0:
			return errors.New("start the import with empty line")
		case 1:
			*state = 2
			return nil
		case 2:
			return errors.New("the 2nd empty line is not allowed inside import")
		}
	}
	imprt, err := moreLiner(line)
	if err != nil {
		return err
	}
	host := strings.SplitN(imprt, "/", 2)
	// simplest way is look for a dot
	external := strings.Contains(host[0], ".")
	if external && *state == 1 {
		return errors.New("no space between external and internal imports")
	}

	if *state == 0 {
		if external {
			*state = 2
		} else {
			*state = 1
		}
	}

	return nil
}
