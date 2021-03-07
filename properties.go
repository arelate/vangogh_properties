package vangogh_properties

const (
	IdProperty        = "id"
	TitleProperty     = "title"
	DeveloperProperty = "developer"
	PublisherProperty = "publisher"
)

func ValidProperty(property string) bool {
	switch property {
	case IdProperty:
		fallthrough
	case TitleProperty:
		fallthrough
	case DeveloperProperty:
		fallthrough
	case PublisherProperty:
		return true
	default:
		return false
	}
}
