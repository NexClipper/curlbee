package util

import "strings"

func SplitParameter(p string) map[string]string {
	parameters := make(map[string]string)
	if p != "" {
		splits := strings.Split(p, ",")
		for _, s := range splits {
			kv := strings.Split(s, "=")
			if len(kv) == 2 {
				parameters[kv[0]] = kv[1]
			}
		}

		return parameters
	}

	return nil
}
