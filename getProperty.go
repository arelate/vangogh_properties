package vangogh_properties

import (
	"fmt"
	"github.com/arelate/gog_types"
	"github.com/arelate/vangogh_types"
)

func GetStrProperty(id string, pt vangogh_types.ProductType, mt gog_types.Media, property string) (value string, err error) {
	if !SupportsProperty(pt, property) {
		return "", fmt.Errorf("vangogh_properties: %s doesn't support property %s", pt, property)
	}

	var getStrProperty func(string, vangogh_types.ProductType, gog_types.Media) (string, error)

	switch property {
	case TitleProperty:
		getStrProperty = getTitle
	case DeveloperProperty:
		getStrProperty = getDeveloper
	case PublisherProperty:
		getStrProperty = getPublisher
	case ImageProperty:
		getStrProperty = getImage
	case BoxArtProperty:
		getStrProperty = getBoxArt
	case BackgroundImageProperty:
		getStrProperty = getBackgroundImage
	case GalaxyBackgroundImageProperty:
		getStrProperty = getGalaxyBackgroundImage
	case LogoProperty:
		getStrProperty = getLogo
	case IconProperty:
		getStrProperty = getIcon
	default:
		return value, fmt.Errorf("vangogh_properties: unknown property %s", property)
	}

	return getStrProperty(id, pt, mt)
}
