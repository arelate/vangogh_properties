package vangogh_properties

func AllText() []string {
	return []string{
		TitleProperty,
		DeveloperProperty,
		PublisherProperty,
	}
}

func AllNumerical() []string {
	return []string{
		RatingProperty,
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
	all = append(all, AllNumerical()...)
	return append(all, AllImageId()...)
}

func All() []string {
	all := []string{IdProperty}
	return append(all, AllExtracted()...)
}
