package gobbledeegook

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Generate(language string, in string) string {
	decoded := make(map[string]string)
	json.Unmarshal([]byte(in), &decoded)

	fields := make([]string, 0)
	inits := make([]string, 0)
	for key, value := range decoded {
		fields = append(fields, fmt.Sprintf("\t%s string", key))
		inits = append(inits, "\""+value+"\"")
	}

	return "var " + language + " = struct {\n" + strings.Join(fields, "\n") + "\n}{" + strings.Join(inits, ",") + "}"
}
