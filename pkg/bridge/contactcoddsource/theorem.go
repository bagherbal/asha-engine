package contactcoddsource

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactCoddSourceFunctionalFiniteSignedCurrentConstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-CODD-SOURCE-FUNCTIONAL-FINITE-SIGNED-CURRENT-CONSTRUCTION"
	const name = "Contact C-odd source functional / finite signed-current construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact C-odd signed-current construction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 143 unbroken C-degeneracy and closed beta firewall are inherited", Passed: a.ContactRows == 7 && a.LargestGapHighRows == 3 && a.LargestGapLowRows == 4 && a.OrientationCandidates == 2 && a.BetaPermissionFirewallClosed, Detail: fmt.Sprintf("split=%s rows=%d", a.SplitPattern, a.ContactRows)},
			{Name: "centered contact spectral functional is canonical, trace-zero, signed, and matches the 3|4 diagnostic split", Passed: a.CenteredFunctional.CanonicalAsDiagnostic && !a.CenteredFunctional.PhysicalCoddSource && absf(a.CenteredFunctional.Trace) < 1e-9 && a.CenteredPositiveRows == 3 && a.CenteredNegativeRows == 4 && a.CenteredZeroRows == 0 && a.CenteredFunctional.MatchesLargestGap, Detail: FormatCenteredFunctional(a.CenteredFunctional)},
			{Name: "binary and balanced gap split functionals do not become T3R or a selected physical branch", Passed: !a.BinarySplit.BinaryT3RAvailable && a.BinarySplit.TracelessSignedAvailable && !a.BinarySplit.T3RSemantic && !a.BinarySplit.SelectedPhysicalBranch, Detail: FormatBinarySplit(a.BinarySplit)},
			{Name: "signed-current audit distinguishes finite diagnostics from missing C-odd source currents", Passed: a.SignedSourcesAudited == 7 && a.AvailableSignedDiagnostics == 3 && a.TraceZeroDiagnostics == 2 && a.CanonicalSignedDiagnostics == 1 && a.CoddContactFunctionals == 0 && a.CBreakingSources == 0 && a.SourcesSelectingPhysicalSign == 0, Detail: FormatSignedSources(a.SourceAudits)},
			{Name: "B-L, chirality, T3R, SU2L, hypercharge, local field, mass activation, and decoupling remain absent", Passed: a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.BinaryT3RRowsDerived == 0, Detail: Join(a.RemainingUnknowns)},
			{Name: "contact beta firewall, S6 ambiguity, nullity, and no-observed-input discipline remain sealed", Passed: a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: a.TruthStatement},
		}}
	}}
}

func absf(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
