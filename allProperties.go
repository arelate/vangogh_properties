package vangogh_properties

func AllTextProperties() []string {
	return []string{
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
	}
}

func AllImageIdProperties() []string {
	return []string{
		ImageProperty,
		BoxArtProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		IconProperty,
		LogoProperty,
		ScreenshotsProperty,
	}
}

func AllExtractedProperties() []string {
	all := AllTextProperties()
	return append(all, AllImageIdProperties()...)
}

func AllProperties() []string {
	all := []string{IdProperty}
	return append(all, AllExtractedProperties()...)
}
