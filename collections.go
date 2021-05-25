package vangogh_properties

const (
	IdProperty                = "id"
	TitleProperty             = "title"
	DevelopersProperty        = "developers"
	PublisherProperty         = "publisher"
	ImageProperty             = "image"
	BoxArtProperty            = "box-art"
	BackgroundProperty        = "background"
	GalaxyBackgroundProperty  = "galaxy-background"
	IconProperty              = "icon"
	LogoProperty              = "logo"
	ScreenshotsProperty       = "screenshots"
	RatingProperty            = "rating"
	IncludesGamesProperty     = "includes-games"
	IsIncludedByGamesProperty = "is-included-by-games"
	RequiresGamesProperty     = "requires-games"
	IsRequiredByGamesProperty = "is-required-by-games"
	GenresProperty            = "genres"
	FeaturesProperty          = "features"
	SeriesProperty            = "series"
	TagIdProperty             = "tag"
	TagNameProperty           = "tag-name"
	VideoIdProperty           = "video-id"
	OperatingSystemsProperty  = "os"
	LanguageCodeProperty      = "lang-code"
	LanguageNameProperty      = "lang-name"
	SlugProperty              = "slug"
	TextProperties            = "text"
	AllTextProperties         = "all-text"
	ImageIdProperties         = "image-id"
	TypesProperty             = "types"
)

func Text() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
	}
}

func AllText() []string {
	return append(Text(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		RatingProperty,
		TagIdProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		SlugProperty,
	)
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
	all := AllText()
	all = append(all, VideoId()...)
	all = append(all, Computed()...)
	return append(all, ImageId()...)
}

func Computed() []string {
	return []string{
		TypesProperty,
	}
}

func All() []string {
	all := []string{IdProperty}
	return append(all, Extracted()...)
}

func Searchable() []string {
	searchable := make([]string, 0, len(Extracted())+3)
	searchable = append(searchable, TextProperties, AllTextProperties, ImageIdProperties)
	searchable = append(searchable, Extracted()...)
	return searchable
}

func Digestible() []string {
	return []string{
		DevelopersProperty,
		PublisherProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		TagIdProperty,
		LanguageCodeProperty,
	}
}
