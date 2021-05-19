package vangogh_properties

var collapsedExpanded = map[string][]string{
	TextProperties:    Text(),
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

func ExpandAll(properties map[string]bool) map[string]bool {
	//expanded := make([]string, 0, len(properties))
	for prop, ok := range properties {
		if !ok {
			continue
		}
		if IsCollapsed(prop) {
			for _, ep := range Expand(prop) {
				properties[ep] = true
			}
		}
	}
	return properties
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
