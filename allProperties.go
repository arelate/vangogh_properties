package vangogh_properties

func AllProperties() []string {
	return []string{
		IdProperty, // must be first to be removed when not needed
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
		ImageProperty,
		BoxArtProperty,
		BackgroundImageProperty,
		GalaxyBackgroundImageProperty,
		IconProperty,
		LogoProperty,
	}
}

func AllStashedProperties() []string {
	// stashed properties
	all := AllProperties()
	// remove first property (IdProperty)
	all[0] = all[len(all)-1]
	return all[:len(all)-1]
}
