package generation2electroweakmeetingdeficitclosuredualrootaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ElectroweakMeetingDeficitClosureDualRootAlignmentAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 664 — ElectroweakMeeting DeficitClosure Dual-Root Alignment Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate664 dual-root audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate663 zero-crossing result", Passed: a.Inherited.ZeroCrossingInherited && a.Inherited.NoStationaryClaim && a.Inherited.NoNativeScale && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoUncertainty && a.Inherited.NoBoundaryStress && math.Abs(a.Inherited.E72AtLambda12-8.525835107353608e-10) < 5e-16, Detail: FormatInherited(a.Inherited)},
			{Name: "construct v1 transport seed", Passed: a.Seed.Mu0GeV > 0 && a.Seed.Lambda12GeV > 9e13 && len(a.Seed.InitialVector) == 13, Detail: FormatSeed(a.Seed)},
			{Name: "define electroweak meeting functions", Passed: math.Abs(a.Meeting.T12Analytic-a.Seed.T12) < 1e-12 && math.Abs(a.Meeting.F12AtRoot) < 1e-12 && math.Abs(a.Meeting.U12AtRoot) < 1e-10, Detail: FormatMeeting(a.Meeting)},
			{Name: "compute E72 closure root", Passed: a.ClosureRoot.ClosureIsTransverse && math.Abs(a.ClosureRoot.E72AtClosureZero) < 1e-12, Detail: FormatClosureRoot(a.ClosureRoot)},
			{Name: "compute dual-root offset", Passed: a.DualRoot.AlignedInV1 && math.Abs(a.DualRoot.DeltaLogMuEOverMu12) < 1e-5 && math.Abs(a.DualRoot.MuEOverMu12-0.9999991071689) < 1e-9, Detail: FormatDualRoot(a.DualRoot)},
			{Name: "audit transversality", Passed: a.Transversality.F12Transverse && a.Transversality.U12Transverse && a.Transversality.E72Transverse && !a.Transversality.SlopeTied && a.Transversality.DE72DtAtLambda12 > 9e-4, Detail: FormatTransversality(a.Transversality)},
			{Name: "audit local proportionality", Passed: a.Proportionality.Samples == 5 && a.Proportionality.RelativeResidualF12 < 0.2 && a.Proportionality.RelativeResidualU12 < 0.2, Detail: FormatProportionality(a.Proportionality)},
			{Name: "audit gauge residual conventions", Passed: len(a.Conventions.Rows) == 5 && a.Conventions.DirectCouplingConventionsPass >= 4 && a.Conventions.ConventionStable, Detail: FormatConventions(a.Conventions)},
			{Name: "audit best weight root", Passed: len(a.WeightRoot.Rows) == 5 && math.Abs(a.WeightRoot.WBestMinus7Over72AtLambda12) < 1e-6 && a.WeightRoot.CrossesSevenOver72NearLambda && !a.WeightRoot.WeightIndependentlySelected, Detail: FormatWeightRoot(a.WeightRoot)},
			{Name: "classify source type and preserve firewalls", Passed: len(a.Source.Outcomes) == 5 && !a.Discipline.ClaimsNativeDualRootTheorem && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsFullUncertaintyPropagation && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate664Boundary, Detail: FormatSource(a.Source) + " | " + FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func _containsAll(haystack string, needles []string) bool {
	for _, n := range needles {
		if !strings.Contains(haystack, n) {
			return false
		}
	}
	return true
}
