package vangogh_properties

func AllText() []string {
	return []string{
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
	}
}

func AllImageId() []string {
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

func AllExtracted() []string {
	all := AllText()
	return append(all, AllImageId()...)
}

func All() []string {
	all := []string{IdProperty}
	return append(all, AllExtracted()...)
}
