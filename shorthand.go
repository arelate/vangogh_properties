package vangogh_properties

func Shorthand(property string) string {
	for _, tp := range Text() {
		if tp == property {
			return TextProperties
		}
	}
	for _, ip := range ImageId() {
		if ip == property {
			return ImageIdProperties
		}
	}
	return ""
}
