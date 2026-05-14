package propagatorspectrum

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FinitePropagatorSpectrumSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-PROPAGATOR-SPECTRUM-SEARCH"
	const name = "finite propagator from B-sector/contact spectrum search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite propagator spectrum audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 63 exchange diagnostic input", Passed: a.Exchange.UnitPropagatorBranchAvailable, Detail: fmt.Sprintf("unit attractive diagnostic G_hat=%.10f across %d current sectors", a.Exchange.UnitAttractiveKernel, a.CurrentSectorCount)},
			{Name: "finite spectral anchors available", Passed: a.SpectralDenominatorsAvailable, Detail: fmt.Sprintf("threshold candidates=%d, positive spectral anchors=%d", a.ThresholdCandidateCount, a.PositiveSpectralAnchorCount)},
			{Name: "candidate denominator families", Passed: len(a.Families) > 0, Detail: FormatFamilies(a.Families)},
			{Name: "diagnostic kernel sensitivity", Passed: a.StrongestDiagnosticKernel > 0, Detail: fmt.Sprintf("smallest rho=%.10f, largest rho=%.10f, strongest diagnostic=%s with G=%.10f", a.SmallestDenominator, a.LargestDenominator, a.StrongestDiagnosticFamily, a.StrongestDiagnosticKernel)},
			{Name: "sector spectral assignment", Passed: a.SectorSpectralAssignmentDerived, Detail: "open; no theorem maps finite spectral anchors to central/color/B-L/leptoquark current sectors"},
			{Name: "current-sector representation map", Passed: a.CurrentSectorRepresentationMapDerived, Detail: "open; B-sector/contact/scalar spectra are not yet represented as current-sector propagator operators"},
			{Name: "propagator denominators", Passed: a.PropagatorDenominatorsDerived, Detail: "open; all rho_A choices are diagnostic families, not derived denominators"},
			{Name: "exchange kernel updated", Passed: a.ExchangeKernelUpdated, Detail: "open; G_hat is not updated because no denominator family is selected"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveKernelDerived, Detail: "open; denominator candidates do not derive exchange sign, couplings, or attraction"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no spectral denominator distinguishes top-like up from bottom-like down"},
			{Name: "regulator criticality", Passed: a.RegulatorCriticalityDerived, Detail: "open; C_reg for the finite NJL gap equation is still missing"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no top condensation, Higgs VEV, or mass scale is claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed threshold, mass, coupling, y_t, v, or Higgs mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
