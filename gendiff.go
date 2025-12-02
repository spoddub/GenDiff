package code

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"code/parsers"
)

func formatLine(sign rune, key string, value any) string {
	var b strings.Builder

	if sign == ' ' {
		b.WriteString("    ")
	} else {
		b.WriteString("  ")
		b.WriteRune(sign)
		b.WriteRune(' ')
	}

	b.WriteString(fmt.Sprintf("%s: %v\n", key, value))
	return b.String()
}

func GenDiff(path1, path2, format string) (string, error) {
	data1, err := parsers.Parse(path1)
	if err != nil {
		return "", err
	}
	data2, err := parsers.Parse(path2)
	if err != nil {
		return "", err
	}

	keysSet := make(map[string]struct{})
	for k := range data1 {
		keysSet[k] = struct{}{}
	}
	for k := range data2 {
		keysSet[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysSet))
	for k := range keysSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("{\n")

	for _, key := range keys {
		v1, ok1 := data1[key]
		v2, ok2 := data2[key]

		switch {
		case ok1 && ok2:
			if reflect.DeepEqual(v1, v2) {
				b.WriteString(formatLine(' ', key, v1))
			} else {
				b.WriteString(formatLine('-', key, v1))
				b.WriteString(formatLine('+', key, v2))
			}
		case ok1 && !ok2:
			b.WriteString(formatLine('-', key, v1))
		case !ok1 && ok2:
			b.WriteString(formatLine('+', key, v2))
		}
	}

	b.WriteString("}")
	return b.String(), nil
}
