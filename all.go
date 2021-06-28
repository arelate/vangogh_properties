package vangogh_properties

const (
	IdProperty                 = "id"
	TitleProperty              = "title"
	DevelopersProperty         = "developers"
	PublisherProperty          = "publisher"
	ImageProperty              = "image"
	BoxArtProperty             = "box-art"
	BackgroundProperty         = "background"
	GalaxyBackgroundProperty   = "galaxy-background"
	IconProperty               = "icon"
	LogoProperty               = "logo"
	ScreenshotsProperty        = "screenshots"
	RatingProperty             = "rating"
	IncludesGamesProperty      = "includes-games"
	IsIncludedByGamesProperty  = "is-included-by-games"
	RequiresGamesProperty      = "requires-games"
	IsRequiredByGamesProperty  = "is-required-by-games"
	GenresProperty             = "genres"
	FeaturesProperty           = "features"
	SeriesProperty             = "series"
	TagIdProperty              = "tag"
	TagNameProperty            = "tag-name"
	VideoIdProperty            = "video-id"
	MissingVideoUrlProperty    = "missing-video-url"
	OperatingSystemsProperty   = "os"
	LanguageCodeProperty       = "lang-code"
	LanguageNameProperty       = "lang-name"
	NativeLanguageNameProperty = "native-lang-name"
	SlugProperty               = "slug"
	GOGReleaseDate             = "gog-release-date"
	GlobalReleaseDate          = "global-release-date"
	TextProperties             = "text"
	AllTextProperties          = "all-text"
	ImageIdProperties          = "image-id"
	TypesProperty              = "types"
)

func All() []string {
	all := []string{IdProperty}
	return append(all, Extracted()...)
}
