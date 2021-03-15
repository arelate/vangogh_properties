package vangogh_properties

import (
	"github.com/arelate/vangogh_urls"
	"github.com/boggydigital/froth"
	"log"
)

func PropStashes(properties []string) (map[string]*froth.Stash, error) {

	propStashes := make(map[string]*froth.Stash, len(properties))
	distStashUrl := vangogh_urls.DistilledStashUrl()

	for _, prop := range properties {
		if prop == IdProperty {
			continue
		}
		if !ValidProperty(prop) {
			log.Printf("vangogh: invalid property %s", prop)
			continue
		}
		stash, err := froth.NewStash(distStashUrl, prop)
		if err != nil {
			return propStashes, err
		}
		propStashes[prop] = stash
	}

	return propStashes, nil
}
