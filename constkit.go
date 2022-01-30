package constkit

import (
	"github.com/bartmika/go-constkit/health"
)

func ListAllAlergies() []*health.Allergy {
	return health.Allergies
}

func ListAllDiseases() []*health.Disease {
	return health.Diseases
}

// ListAllNHSIllnesses return a list of common disease, illnesses and conditions according to the Scotland's national health information service (NHS).
func ListAllNHSIllnesses() []*health.NHSIllness {
	return health.NHSIllnesses
}
