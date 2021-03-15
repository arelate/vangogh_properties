package vangogh_properties

import (
	"github.com/arelate/vangogh_types"
)

func SupportsProperty(pt vangogh_types.ProductType, property string) bool {
	it := ToImageType(property)
	if it == vangogh_types.UnknownImageType {
		return false
	}
	return vangogh_types.SupportsImageType(pt, it)
}
