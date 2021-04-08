package vangogh_properties

var joinNotDesirable = AllImageId()

func JoinPreferred(property string) bool {
	for _, nd := range joinNotDesirable {
		if property == nd {
			return false
		}
	}
	return true
}
