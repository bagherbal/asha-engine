package generation2boundarystresssplitpullbackcorrectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryStressSplitPullbackCorrectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 672 — BoundaryStressSplit Pullback Correction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate672 stress-split pullback audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate671 normal-vector firewall", Passed: a.Inherited.NormalVectorInherited && a.Inherited.NormalVectorBestTypedExact && a.Inherited.CoordinateSealed && a.Inherited.NoNativeNormalVectorTheorem && a.Inherited.NoNativeSevenOver72Theorem && a.Inherited.NoWallDistanceAirlockTheorem && a.Inherited.NoBoundaryStressDerivation && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "decompose n_72 into base plus stress-split pullback", Passed: a.Decomposition.DecompositionPasses && math.Abs(a.Decomposition.Weight-sevenOver72) < 1e-15 && strings.Contains(a.Decomposition.Verdict, StatusNormalVectorStressSplitDecomposed), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "compute base scalar/flavor closure", Passed: math.Abs(a.BaseClosure.DBase-0.00012565520996836) < 1e-14 && strings.Contains(a.BaseClosure.Verdict, StatusBaseScalarFlavorClosureComputed), Detail: FormatBase(a.BaseClosure)},
			{Name: "compute boundary stress split", Passed: math.Abs(a.StressSplit.SSplit-0.0012924448188163) < 1e-14 && strings.Contains(a.StressSplit.Verdict, StatusBoundaryStressSplitComputed), Detail: FormatStress(a.StressSplit)},
			{Name: "test 7/72 pullback", Passed: a.Pullback.PassesBridgeWindow && math.Abs(a.Pullback.Pullback-0.000125654357384641) < 1e-14 && math.Abs(a.Pullback.Residual-8.52583727234e-10) < 1e-14 && strings.Contains(a.Pullback.Verdict, StatusSevenOver72PullbackTested), Detail: FormatPullback(a.Pullback)},
			{Name: "reconstruct HistoryWallBalance normal", Passed: a.Reconstruction.EquivalentToGate670Normal && math.Abs(a.Reconstruction.DBaseMinusPullback-a.Reconstruction.HistoryWallBalance) < 1e-15 && strings.Contains(a.Reconstruction.Verdict, StatusNormalVectorReconstructionComputed), Detail: FormatReconstruction(a.Reconstruction)},
			{Name: "audit source types and missing maps", Passed: len(a.Source.CandidateSupport) == 3 && len(a.Source.RequiredMissingMaps) == 4 && strings.Contains(a.Source.Verdict, StatusStressSplitCorrectedScalarFlavorClosure) && strings.Contains(a.Source.Verdict, StatusSevenOver72ActsOnBoundaryStressSplit), Detail: FormatSource(a.Source)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeStressSplitPullback && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsWallDistanceAirlock && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate672Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
