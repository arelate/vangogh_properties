package vangogh_properties

import (
	"github.com/arelate/vangogh_urls"
	"github.com/boggydigital/froth"
	"log"
)

func PropExtracts(properties []string) (map[string]*froth.Stash, error) {

	propExtracts := make(map[string]*froth.Stash, len(properties))
	extractsDir := vangogh_urls.ExtractsDir()

	for _, prop := range properties {
		if prop == IdProperty {
			continue
		}
		if !Valid(prop) {
			log.Printf("vangogh: invalid property %s", prop)
			continue
		}

		extracts, err := froth.NewStash(extractsDir, prop)
		if err != nil {
			return propExtracts, err
		}

		propExtracts[prop] = extracts
	}

	return propExtracts, nil
}
