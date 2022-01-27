package constkit

import (
	"github.com/bartmika/constkit/health"
)

func ListAllAlergies() []*health.Allergy {
	return health.Allergies
}

func ListAllDiseases() []*health.Disease {
	return health.Diseases
}
