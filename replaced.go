package vangogh_properties

import "fmt"

var replacementProperties = map[string]string{
	IncludesGamesProperty:     TitleProperty,
	IsIncludedByGamesProperty: TitleProperty,
	RequiresGamesProperty:     TitleProperty,
	IsRequiredByGamesProperty: TitleProperty,
	TagIdProperty:             TagNameProperty,
	LanguageCodeProperty:      NativeLanguageNameProperty,
}

const replOrigValuesFormat = "%s (%s)"

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

func Fmt(origVal, replVal string) string {
	return fmt.Sprintf(replOrigValuesFormat, replVal, origVal)
}
