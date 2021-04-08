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
	AllTextProperties        = "text"
	AllImageIdProperties     = "image-id"
)

func AllText() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		IncludesGamesProperty,
		GenresProperty,
		FeaturesProperty,
	}
}

func AllNumerical() []string {
	return []string{
		RatingProperty,
	}
}

func AllImageId() []string {
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

func AllExtracted() []string {
	all := AllText()
	all = append(all, AllNumerical()...)
	return append(all, AllImageId()...)
}

func All() []string {
	all := []string{IdProperty}
	return append(all, AllExtracted()...)
}

func AllQuery() map[string][]string {
	return map[string][]string{
		AllTextProperties:    AllText(),
		AllImageIdProperties: AllImageId(),
		TitleProperty:        {TitleProperty},
		DevelopersProperty:   {DevelopersProperty},
		PublisherProperty:    {PublisherProperty},
		GenresProperty:       {GenresProperty},
		FeaturesProperty:     {FeaturesProperty},
		RatingProperty:       {RatingProperty},
	}
}

func AllSearchable() []string {
	allQuery := AllQuery()
	keys := make([]string, 0, len(allQuery))
	for key, _ := range allQuery {
		keys = append(keys, key)
	}
	return keys
}
