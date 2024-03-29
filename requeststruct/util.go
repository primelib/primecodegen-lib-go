package requeststruct

import (
	"strings"
)

func parseKVTags(tag string) map[string]string {
	kv := map[string]string{}

	for _, part := range strings.Split(tag, ",") {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			continue
		}

		kv[parts[0]] = parts[1]
	}

	return kv
}
