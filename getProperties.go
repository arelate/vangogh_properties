package vangogh_properties

import (
	"github.com/arelate/gog_types"
	"github.com/arelate/gog_urls"
	"github.com/arelate/vangogh_values"
	"strconv"
	"strings"
)

func GetProperties(
	id string,
	reader *vangogh_values.ValueReader,
	properties []string) (propValues map[string]string, err error) {
	value, err := reader.ProductType(id)
	return fillProperties(value, properties), err
}

func fillProperties(value interface{}, properties []string) map[string]string {
	propValues := make(map[string]string, 0)
	for _, prop := range properties {
		propValues[prop] = getPropertyValue(value, prop)
	}
	return propValues
}

func getPropertyValue(value interface{}, property string) string {
	switch property {
	case BackgroundProperty:
		return getBackground(value)
	case BoxArtProperty:
		return getBoxArt(value)
	case DeveloperProperty:
		return getDeveloper(value)
	case GalaxyBackgroundProperty:
		return getGalaxyBackground(value)
	case IconProperty:
		return getIcon(value)
	case ImageProperty:
		return getImage(value)
	case LogoProperty:
		return getLogo(value)
	case PublisherProperty:
		return getPublisher(value)
	case TitleProperty:
		return getTitle(value)
	case ScreenshotsProperty:
		return getScreenshots(value)
	case RatingProperty:
		return getRating(value)
	default:
		return ""
	}
}

func getTitle(value interface{}) string {
	titleGetter := value.(gog_types.TitleGetter)
	if titleGetter != nil {
		return titleGetter.GetTitle()
	}
	return ""
}

func getDeveloper(value interface{}) string {
	developerGetter := value.(gog_types.DeveloperGetter)
	if developerGetter != nil {
		return developerGetter.GetDeveloper()
	}
	return ""
}

func getPublisher(value interface{}) string {
	publisherGetter := value.(gog_types.PublisherGetter)
	if publisherGetter != nil {
		return publisherGetter.GetPublisher()
	}
	return ""
}

func getBackground(value interface{}) string {
	backgroundGetter := value.(gog_types.BackgroundGetter)
	if backgroundGetter != nil {
		return gog_urls.ImageId(backgroundGetter.GetBackground())
	}
	return ""
}

func getGalaxyBackground(value interface{}) string {
	galaxyBackgroundGetter := value.(gog_types.GalaxyBackgroundGetter)
	if galaxyBackgroundGetter != nil {
		return gog_urls.ImageId(galaxyBackgroundGetter.GetGalaxyBackground())
	}
	return ""
}

func getBoxArt(value interface{}) string {
	boxArtGetter := value.(gog_types.BoxArtGetter)
	if boxArtGetter != nil {
		return gog_urls.ImageId(boxArtGetter.GetBoxArt())
	}
	return ""
}

func getIcon(value interface{}) string {
	iconGetter := value.(gog_types.IconGetter)
	if iconGetter != nil {
		return gog_urls.ImageId(iconGetter.GetIcon())
	}
	return ""
}

func getImage(value interface{}) string {
	imageGetter := value.(gog_types.ImageGetter)
	if imageGetter != nil {
		return gog_urls.ImageId(imageGetter.GetImage())
	}
	return ""
}

func getLogo(value interface{}) string {
	logoGetter := value.(gog_types.LogoGetter)
	if logoGetter != nil {
		return gog_urls.ImageId(logoGetter.GetLogo())
	}
	return ""
}

func getScreenshots(value interface{}) string {
	screenshotsGetter := value.(gog_types.ScreenshotsGetter)
	if screenshotsGetter != nil {
		screenshots := screenshotsGetter.GetScreenshots()
		imageIds := make([]string, 0, len(screenshots))
		for _, scr := range screenshots {
			imageIds = append(imageIds, gog_urls.ImageId(scr))
		}
		return strings.Join(imageIds, ",")
	}
	return ""
}

func getRating(value interface{}) string {
	ratingGetter := value.(gog_types.RatingGetter)
	if ratingGetter != nil {
		return strconv.Itoa(ratingGetter.GetRating())
	}
	return "0"
}
