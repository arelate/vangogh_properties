package vangogh_properties

const (
	IdProperty               = "id"
	TitleProperty            = "title"
	DevelopersProperty       = "developers"
	PublisherProperty        = "publisher"
	ImageProperty            = "image"
	BoxArtProperty           = "box-art"
	BackgroundProperty       = "background"
	GalaxyBackgroundProperty = "galaxy-background"
	IconProperty             = "icon"
	LogoProperty             = "logo"
	ScreenshotsProperty      = "screenshots"
	AllTextProperties        = "text"
	AllImageIdProperties     = "image-id"
	RatingProperty           = "rating"
)

func ValidProperty(property string) bool {
	for _, p := range All() {
		if p == property {
			return true
		}
	}
	return false
}
