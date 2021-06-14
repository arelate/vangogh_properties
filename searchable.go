package vangogh_properties

func Searchable() []string {
	searchable := make([]string, 0, len(Extracted())+3)
	searchable = append(searchable, TextProperties, AllTextProperties, ImageIdProperties)
	searchable = append(searchable, Extracted()...)
	return searchable
}
