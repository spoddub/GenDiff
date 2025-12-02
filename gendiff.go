package code

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"code/parsers"
)

const indentSize = 4

type nodeType string

const (
	nodeAdded     nodeType = "added"
	nodeRemoved   nodeType = "removed"
	nodeUnchanged nodeType = "unchanged"
	nodeUpdated   nodeType = "updated"
	nodeNested    nodeType = "nested"
)

type node struct {
	Key      string
	Type     nodeType
	Value    any
	OldValue any
	NewValue any
	Children []node
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

	tree := buildDiff(data1, data2)

	switch format {
	case "", "stylish":
		return formatStylish(tree), nil
	default:
		return "", fmt.Errorf("unsupported format %q", format)
	}
}

func buildDiff(data1, data2 map[string]any) []node {
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

	result := make([]node, 0, len(keys))

	for _, key := range keys {
		v1, ok1 := data1[key]
		v2, ok2 := data2[key]

		switch {
		case ok1 && ok2:
			m1, m1ok := toMap(v1)
			m2, m2ok := toMap(v2)

			if m1ok && m2ok {
				children := buildDiff(m1, m2)
				result = append(result, node{
					Key:      key,
					Type:     nodeNested,
					Children: children,
				})
			} else if reflect.DeepEqual(v1, v2) {
				result = append(result, node{
					Key:   key,
					Type:  nodeUnchanged,
					Value: v1,
				})
			} else {
				result = append(result, node{
					Key:      key,
					Type:     nodeUpdated,
					OldValue: v1,
					NewValue: v2,
				})
			}

		case ok1 && !ok2:
			result = append(result, node{
				Key:   key,
				Type:  nodeRemoved,
				Value: v1,
			})

		case !ok1 && ok2:
			result = append(result, node{
				Key:   key,
				Type:  nodeAdded,
				Value: v2,
			})
		}
	}

	return result
}

func toMap(v any) (map[string]any, bool) {
	if v == nil {
		return nil, false
	}

	if m, ok := v.(map[string]any); ok {
		return m, true
	}

	if m, ok := v.(map[string]interface{}); ok {
		res := make(map[string]any, len(m))
		for k, vv := range m {
			res[k] = vv
		}
		return res, true
	}

	return nil, false
}

func formatStylish(tree []node) string {
	var b strings.Builder
	b.WriteString("{\n")
	formatNodes(&b, tree, 1)
	b.WriteString("}")
	return b.String()
}

func formatNodes(b *strings.Builder, nodes []node, depth int) {
	indent := strings.Repeat(" ", depth*indentSize-2)

	for _, n := range nodes {
		switch n.Type {
		case nodeNested:
			fmt.Fprintf(b, "%s  %s: {\n", indent, n.Key)
			formatNodes(b, n.Children, depth+1)
			closingIndent := strings.Repeat(" ", depth*indentSize)
			fmt.Fprintf(b, "%s}\n", closingIndent)

		case nodeUnchanged:
			valStr := stringify(n.Value, depth)
			fmt.Fprintf(b, "%s  %s: %s\n", indent, n.Key, valStr)

		case nodeAdded:
			valStr := stringify(n.Value, depth)
			fmt.Fprintf(b, "%s+ %s: %s\n", indent, n.Key, valStr)

		case nodeRemoved:
			valStr := stringify(n.Value, depth)
			fmt.Fprintf(b, "%s- %s: %s\n", indent, n.Key, valStr)

		case nodeUpdated:
			oldValStr := stringify(n.OldValue, depth)
			newValStr := stringify(n.NewValue, depth)
			fmt.Fprintf(b, "%s- %s: %s\n", indent, n.Key, oldValStr)
			fmt.Fprintf(b, "%s+ %s: %s\n", indent, n.Key, newValStr)
		}
	}
}

func stringify(value any, depth int) string {
	if value == nil {
		return "null"
	}

	m, ok := toMap(value)
	if !ok {
		return fmt.Sprintf("%v", value)
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("{\n")

	for _, key := range keys {
		val := m[key]
		indent := strings.Repeat(" ", (depth+1)*indentSize)
		b.WriteString(fmt.Sprintf("%s%s: %s\n", indent, key, stringify(val, depth+1)))
	}

	closingIndent := strings.Repeat(" ", depth*indentSize)
	b.WriteString(fmt.Sprintf("%s}", closingIndent))

	return b.String()
}
