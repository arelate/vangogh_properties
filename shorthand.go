package vangogh_properties

func Shorthand(property string) string {
	for _, tp := range AllText() {
		if tp == property {
			return AllTextProperties
		}
	}
	for _, ip := range AllImageId() {
		if ip == property {
			return AllImageIdProperties
		}
	}
	return ""
}
