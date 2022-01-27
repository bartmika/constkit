package health

type Disease struct {
	Name string
}

// https://www.nhsinform.scot/illnesses-and-conditions/a-to-z
// https://en.wikipedia.org/wiki/Lists_of_diseases

var Diseases = []*Disease{
	&Disease{Name: "Abdominal aortic aneurysm"},
	&Disease{Name: "Acne"},
	&Disease{Name: "Acute cholecystitis"},
	&Disease{Name: "Acute lymphoblastic leukaemia"},
	&Disease{Name: "Acute myeloid leukaemia"},
	// ... TODO ...
}
