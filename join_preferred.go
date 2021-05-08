package vangogh_properties

func joinNotDesirable() []string {
	jnd := ImageId()
	jnd = append(jnd, IsRequiredByGamesProperty, IncludesGamesProperty, RequiresGamesProperty)
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
