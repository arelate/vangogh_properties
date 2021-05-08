package vangogh_properties

import "fmt"

var replacementProperties = map[string]string{
	IncludesGamesProperty:     TitleProperty,
	IsRequiredByGamesProperty: TitleProperty,
	RequiresGamesProperty:     TitleProperty,
}

const replOrigValuesFormat = "%s (%s)"

var replacementFormats = map[string]string{
	IncludesGamesProperty:     replOrigValuesFormat,
	IsRequiredByGamesProperty: replOrigValuesFormat,
	RequiresGamesProperty:     replOrigValuesFormat,
}

func IsReplaced(property string) bool {
	for prop, _ := range replacementProperties {
		if property == prop {
			return true
		}
	}
	return false
}

func Replace(property string) string {
	return replacementProperties[property]
}

func FmtReplacement(property, origVal, replVal string) string {
	format := replacementFormats[property]
	if format != "" {
		return fmt.Sprintf(format, replVal, origVal)
	}
	return replVal
}
