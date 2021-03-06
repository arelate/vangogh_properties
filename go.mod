module github.com/arelate/vangogh_properties

go 1.16

require (
	github.com/arelate/gog_types v0.1.8-alpha
	github.com/arelate/vangogh_types v0.1.3-alpha
	github.com/arelate/vangogh_values v0.1.2-alpha
)

replace (
	github.com/arelate/gog_types => ../gog_types
	github.com/arelate/vangogh_types => ../vangogh_types
	github.com/arelate/vangogh_values => ../vangogh_values
)
