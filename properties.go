package vangogh_properties

const (
	IdProperty                    = "id"
	TitleProperty                 = "title"
	DeveloperProperty             = "developer"
	PublisherProperty             = "publisher"
	ImageProperty                 = "image"
	BoxArtProperty                = "box-art"
	BackgroundImageProperty       = "background-image"
	GalaxyBackgroundImageProperty = "galaxy-background-image"
	IconProperty                  = "icon"
	LogoProperty                  = "logo"
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
		fallthrough
	case BoxArtProperty:
		fallthrough
	case BackgroundImageProperty:
		fallthrough
	case GalaxyBackgroundImageProperty:
		fallthrough
	case LogoProperty:
		fallthrough
	case IconProperty:
		return true
	default:
		return false
	}
}
