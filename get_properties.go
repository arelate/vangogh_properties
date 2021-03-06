package vangogh_properties

import (
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_values"
	"time"
)

func GetProperties(
	id string,
	reader *vangogh_values.ValueReader,
	properties []string) (propValues map[string][]string, err error) {
	supProps := Supported(reader.ProductType(), properties)
	value, err := reader.ReadValue(id)
	return fillProperties(value, supProps), err
}

func fillProperties(value interface{}, properties []string) map[string][]string {
	propValues := make(map[string][]string, 0)
	for _, prop := range properties {
		propValues[prop] = getPropertyValues(value, prop)
	}
	return propValues
}

func getPropertyValues(value interface{}, property string) []string {
	switch property {
	case BackgroundProperty:
		return getImageIdSlice(value.(gog_types.BackgroundGetter).GetBackground)
	case BoxArtProperty:
		return getImageIdSlice(value.(gog_types.BoxArtGetter).GetBoxArt)
	case DevelopersProperty:
		return value.(gog_types.DevelopersGetter).GetDevelopers()
	case FeaturesProperty:
		return value.(gog_types.FeaturesGetter).GetFeatures()
	case IconProperty:
		return getImageIdSlice(value.(gog_types.IconGetter).GetIcon)
	case ImageProperty:
		return getImageIdSlice(value.(gog_types.ImageGetter).GetImage)
	case IncludesGamesProperty:
		return value.(gog_types.IncludesGamesGetter).GetIncludesGames()
	case IsIncludedByGamesProperty:
		return value.(gog_types.IsIncludedInGamesGetter).GetIsIncludedInGames()
	case IsRequiredByGamesProperty:
		return value.(gog_types.IsRequiredByGamesGetter).GetIsRequiredByGames()
	case GalaxyBackgroundProperty:
		return getImageIdSlice(value.(gog_types.GalaxyBackgroundGetter).GetGalaxyBackground)
	case GenresProperty:
		return value.(gog_types.GenresGetter).GetGenres()
	case GlobalReleaseDate:
		return dateSlice(value.(gog_types.GlobalReleaseGetter).GetGlobalRelease)
	case GOGReleaseDate:
		return dateSlice(value.(gog_types.GOGReleaseGetter).GetGOGRelease)
	case LanguageCodeProperty:
		return value.(gog_types.LanguageCodesGetter).GetLanguageCodes()
	case LogoProperty:
		return getImageIdSlice(value.(gog_types.LogoGetter).GetLogo)
	case OperatingSystemsProperty:
		return value.(gog_types.OperatingSystemsGetter).GetOperatingSystems()
	case PublisherProperty:
		return getSlice(value.(gog_types.PublisherGetter).GetPublisher)
	case RatingProperty:
		return getSlice(value.(gog_types.RatingGetter).GetRating)
	case RequiresGamesProperty:
		return value.(gog_types.RequiresGamesGetter).GetRequiresGames()
	case SeriesProperty:
		return getSlice(value.(gog_types.SeriesGetter).GetSeries)
	case ScreenshotsProperty:
		return getScreenshots(value)
	case SlugProperty:
		return getSlice(value.(gog_types.SlugGetter).GetSlug)
	case TagIdProperty:
		return value.(gog_types.TagIdsGetter).GetTagIds()
	case TitleProperty:
		return getSlice(value.(gog_types.TitleGetter).GetTitle)
	case VideoIdProperty:
		return value.(gog_types.VideoIdsGetter).GetVideoIds()
	default:
		return []string{}
	}
}

func dateSlice(timestamper func() int64) []string {
	dates := make([]string, 0)
	if timestamper != nil {
		val := timestamper()
		if val > 0 {
			date := time.Unix(val, 0)
			dates = append(dates, date.Format("2006-01-02"))
		}
	}
	return dates
}

func getSlice(stringer func() string) []string {
	values := make([]string, 0)
	if stringer != nil {
		val := stringer()
		if val != "" {
			values = append(values, val)
		}
	}
	return values
}

func getImageIdSlice(stringer func() string) []string {
	strings := getSlice(stringer)
	imageIds := make([]string, 0, len(strings))
	for _, str := range strings {
		imageIds = append(imageIds, gog_urls.ImageId(str))
	}
	return imageIds
}

func getScreenshots(value interface{}) []string {
	screenshotsGetter := value.(gog_types.ScreenshotsGetter)
	if screenshotsGetter != nil {
		screenshots := screenshotsGetter.GetScreenshots()
		imageIds := make([]string, 0, len(screenshots))
		for _, scr := range screenshots {
			imageIds = append(imageIds, gog_urls.ImageId(scr))
		}
		return imageIds
	}
	return []string{}
}
