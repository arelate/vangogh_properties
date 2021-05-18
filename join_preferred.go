package vangogh_properties

func joinNotDesirable() []string {
	jnd := ImageId()
	jnd = append(jnd,
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
	)
	return jnd
}

func JoinPreferred(property string) bool {
	for _, nd := range joinNotDesirable() {
		if property == nd {
			return false
		}
	}
	return true
}
