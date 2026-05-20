package generation2halftraceboundarycoordinateweightaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HalfTraceBoundaryCoordinateWeightAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 656 — Half-Trace Boundary Coordinate Weight Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate656 half-trace boundary coordinate audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate655 internal-only Fano-Hitchin seal", Passed: a.Inherited.FanoSealDefined && a.Inherited.FanoSealInternalOnly && a.Inherited.FanoStructuresNumerator && a.Inherited.NoBoundaryInterface && a.Inherited.NoSevenOver72Theorem && a.Inherited.NoBoundaryStress && a.Inherited.NoScalarFlavorMap && a.Inherited.NoHistoryLoopSource && !a.Inherited.ClaimsBoundaryStress && !a.Inherited.ClaimsSevenOver72 && !a.Inherited.ClaimsScalarFlavor && !a.Inherited.ClaimsHistoryLoopUnit && a.Inherited.Gate655Firewall, Detail: FormatInherited(a.Inherited)},
			{Name: "audit source type of 7/144", Passed: near(a.SourceType.FullWeight, 7.0/72.0) && near(a.SourceType.HalfWeight, 7.0/144.0) && a.SourceType.SevenTyped && a.SourceType.SeventyTwoTyped && a.SourceType.HalfTyped && !a.SourceType.HalfNative && a.SourceType.AllFactorsTyped && !a.SourceType.CertifiedHalfTraceMap, Detail: FormatSourceType(a.SourceType)},
			{Name: "compare 7/144 to boundary coordinates without certification", Passed: len(a.Boundary.Rows) == 3 && a.Boundary.ClosestTarget == "|lambda(Lambda_12)|" && near(a.Boundary.ClosestResidual, absLambda-wHalf) && !a.Boundary.CertifiedMatch && a.Boundary.NoProximityCertification, Detail: FormatBoundaryComparison(a.Boundary)},
			{Name: "audit mean-stress interpretation", Passed: near(a.MeanStress.HalfWeight, wHalf) && near(a.MeanStress.XiBoundary, xiBoundary) && a.MeanStress.SignedResidual > 0 && a.MeanStress.ExistingMeanStressBetter && a.MeanStress.AntiAlignmentSealStronger, Detail: FormatMeanStress(a.MeanStress)},
			{Name: "audit two-coordinate split routes", Passed: near(a.Split.FullWeight, wFull) && near(a.Split.HalfWeight, wHalf) && near(a.Split.SignedPair[0], wHalf) && near(a.Split.SignedPair[1], -wHalf) && a.Split.FullWeightTyped && a.Split.PerCoordinateTyped && a.Split.SignedPairTyped && a.Split.MeanStressTyped && !a.Split.SuppliesBoundaryMap && !a.Split.SuppliesTraceTheorem, Detail: FormatSplit(a.Split)},
			{Name: "audit relations to prior seals", Passed: len(a.Relations.Rows) == 4 && a.Relations.FanoHitchinSource && !a.Relations.HistoryLoopSource && !a.Relations.BoundaryStressSource && !a.Relations.OrientationBalanceSource, Detail: FormatRelations(a.Relations)},
			{Name: "preserve missing half-trace boundary map", Passed: !a.BoundaryMap.HasHalfTraceMap && !a.BoundaryMap.HasSevenOver72Map && !a.BoundaryMap.HasBoundaryStressMap && !a.BoundaryMap.CanDeriveBoundaryStress && !a.BoundaryMap.CanDeriveLambdaOrR3, Detail: FormatBoundaryMap(a.BoundaryMap)},
			{Name: "preserve boundary/history/physics firewalls", Passed: !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsLambdaR3 && !a.Firewalls.ClaimsSevenOver144 && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsHistoryLoopUnit && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate656Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		if !strings.Contains(a.SourceType.Verdict, StatusNoNativeHalfTraceMap) {
			notes = append(notes, "WARNING_MISSING_HALF_TRACE_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
