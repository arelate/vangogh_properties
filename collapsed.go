package vangogh_properties

var collapsedExpanded = map[string][]string{
	TextProperties:    Text(),
	AllTextProperties: AllText(),
	ImageIdProperties: ImageId(),
}

func IsCollapsed(property string) bool {
	for c, _ := range collapsedExpanded {
		if c == property {
			return true
		}
	}
	return false
}

func Collapsed() []string {
	collapsed := make([]string, 0)
	for c, _ := range collapsedExpanded {
		collapsed = append(collapsed, c)
	}
	return collapsed
}

func Expand(property string) []string {
	return collapsedExpanded[property]
}

func ExpandAll(properties []string) []string {
	expanded := make(map[string]bool, 0)

	for _, prop := range properties {
		if IsCollapsed(prop) {
			for _, ep := range Expand(prop) {
				expanded[ep] = true
			}
		} else {
			expanded[prop] = true
		}
	}

	keys := make([]string, 0, len(expanded))
	for ep, _ := range expanded {
		keys = append(keys, ep)
	}

	return keys
}

func Collapse(property string) string {
	for c, expanded := range collapsedExpanded {
		for _, e := range expanded {
			if e == property {
				return c
			}
		}
	}
	return ""
}
