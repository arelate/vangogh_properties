package vangogh_properties

const (
	IdProperty        = "id"
	TitleProperty     = "title"
	DeveloperProperty = "developer"
	PublisherProperty = "publisher"
	ImageProperty     = "image"
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
		fallthrough
	case ImageProperty:
		return true
	default:
		return false
	}
}
