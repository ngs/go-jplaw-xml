package jplaw

type LawType string

const (
	LawTypeConstitution         LawType = "Constitution"
	LawTypeAct                  LawType = "Act"
	LawTypeCabinetOrder         LawType = "CabinetOrder"
	LawTypeImperialOrder        LawType = "ImperialOrder"
	LawTypeMinisterialOrdinance LawType = "MinisterialOrdinance"
	LawTypeRule                 LawType = "Rule"
	LawTypeMisc                 LawType = "Misc"
)
