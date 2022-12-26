package envstruct

import (
	"os"
	"strings"
)

func parseEnv() {
	osenv := os.Environ()
	r := map[string]string{}
	for _, v := range osenv {
		p := strings.SplitN(v, "=", 2)
		r[p[0]] = p[1]
	}
}
