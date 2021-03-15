module github.com/arelate/vangogh_properties

go 1.16

require (
	github.com/arelate/gog_types v0.1.8-alpha
	github.com/arelate/gog_urls v0.1.8-alpha
	github.com/arelate/vangogh_types v0.1.3-alpha
	github.com/arelate/vangogh_values v0.1.2-alpha
	github.com/boggydigital/froth v0.1.0-alpha
	github.com/arelate/vangogh_urls v0.1.0-alpha
)

replace (
	github.com/arelate/gog_types => ../gog_types
	github.com/arelate/gog_urls => ../gog_urls
	github.com/arelate/vangogh_types => ../vangogh_types
	github.com/arelate/vangogh_values => ../vangogh_values
	github.com/boggydigital/froth => ../froth
	github.com/arelate/vangogh_urls => ../vangogh_urls
)
