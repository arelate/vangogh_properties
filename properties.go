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
	ScreenshotsProperty           = "screenshots"
)

func ValidProperty(property string) bool {
	for _, p := range AllProperties() {
		if p == property {
			return true
		}
	}
	return false
}
