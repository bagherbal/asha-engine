package contactincidence

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactIncidenceFiberFunctorSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-INCIDENCE-FIBER-FUNCTOR-SEARCH"
	const name = "contact incidence/fiber functor search from Fano/contact geometry"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact incidence/fiber functor search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 115 local-bundle obstruction inherited", Passed: a.LocalBundleObstructionInherited && a.ContactBundle.LocalBundleObstructionDerived && a.ContactBundle.RepresentationCompleteRows == 0 && a.ContactBundle.ContactBetaRowsAllowed == 0, Detail: "seven contact rows remain bundle-open before the incidence/functor search"},
			{Name: "exact Fano incidence carrier available", Passed: a.FanoIncidenceAvailable && a.FanoPointCount == 7 && a.FanoLineCount == 7 && a.EveryFanoPointDegreeThree && a.EveryFanoLineSizeThree, Detail: FormatFano(a.FanoLines)},
			{Name: "Fano/contact cardinality resonance exposed", Passed: a.FanoContactCardinalityMatch && a.FanoIncidenceResonance && a.ContactRows == 7 && a.PositiveFiniteContactRows == 7, Detail: fmt.Sprintf("contact rows=%d; Fano points=%d; Fano lines=%d", a.ContactRows, a.FanoPointCount, a.FanoLineCount)},
			{Name: "canonical contact-to-Fano map not derived", Passed: a.ContactToFanoMapAttempted && !a.CanonicalContactToFanoMap, Detail: "7! compatible bijections remain unless a naturality/automorphism-breaking rule is derived"},
			{Name: "fiber functor/chart atlas not derived", Passed: a.FiberFunctorConstructionTried && !a.FiberFunctorDerived && !a.ChartAtlasDerived && !a.TransitionCocycleDerived && !a.SectionMapDerived, Detail: FormatCriteria(a.Criteria)},
			{Name: "representation and threshold permission remain closed", Passed: !a.GaugeRepresentationDerived && !a.LorentzKineticDerived && !a.MassActivationDerived && !a.DecouplingRuleDerived && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0, Detail: FormatRows(a.Rows, 10)},
			{Name: "incidence resonance is not a physical bridge leak", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
