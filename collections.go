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
	RatingProperty           = "rating"
	IncludesGamesProperty    = "includes-games"
	GenresProperty           = "genres"
	FeaturesProperty         = "features"
	SeriesProperty           = "series"
	TagIdProperty            = "tag"
	VideoIdProperty          = "video-id"
	OperatingSystemsProperty = "os"
	TextProperties           = "text"
	ImageIdProperties        = "image-id"
	DLCsProperty             = "dlcs"
)

func Text() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		IncludesGamesProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		RatingProperty,
		TagIdProperty,
		OperatingSystemsProperty,
		DLCsProperty,
	}
}

func ImageId() []string {
	return []string{
		ImageProperty,
		BoxArtProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		IconProperty,
		LogoProperty,
		ScreenshotsProperty,
	}
}

func VideoId() []string {
	return []string{
		VideoIdProperty,
	}
}

func Extracted() []string {
	all := Text()
	all = append(all, VideoId()...)
	return append(all, ImageId()...)
}

func All() []string {
	all := []string{IdProperty}
	return append(all, Extracted()...)
}

func Query() map[string][]string {
	query := make(map[string][]string)

	query[TextProperties] = Text()
	query[ImageIdProperties] = ImageId()
	for _, textProp := range Text() {
		query[textProp] = []string{textProp}
	}
	query[VideoIdProperty] = []string{VideoIdProperty}

	return query
}

func Searchable() []string {
	query := Query()
	keys := make([]string, 0, len(query))
	//TextProperties needs to be the first one
	keys = append(keys, TextProperties)
	for key, _ := range query {
		if key == TextProperties {
			continue
		}
		keys = append(keys, key)
	}
	return keys
}

func Digestible() []string {
	return []string{
		DevelopersProperty,
		PublisherProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
	}
}
