package vangogh_properties

func Valid(property string) bool {
	for _, p := range All() {
		if p == property {
			return true
		}
	}
	return false
}
