package vangogh_properties

func Extracted() []string {
	all := AllText()
	all = append(all, VideoId()...)
	all = append(all, Computed()...)
	return append(all, ImageId()...)
}
