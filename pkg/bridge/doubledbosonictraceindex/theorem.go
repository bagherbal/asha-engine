package doubledbosonictraceindex

import "github.com/bagherbal/asha-engine/pkg/theorem"

func DoubledBosonicTraceIndexJMirrorGaugeCapacityAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-DOUBLED-BOSONIC-TRACE-INDEX-J-MIRROR-GAUGE-CAPACITY-AUDIT"
	const name = "Doubled Bosonic Trace Index / J-Mirror Gauge Capacity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 329 doubled bosonic trace index audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inputs inherit Gate 328 factor-two obligation without empirical fitting", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && nearlyEqual(a.Inputs.RequiredMultiplier, 2.0, 1e-12) && !a.Inputs.AddsEmpiricalFit, Detail: FormatInputs(a.Inputs)},
			{Name: "J-mirror gauge carrier gives equal positive curvature contribution", Passed: nearlyEqual(a.Mirror.FullDoubledTraceIndex, 2.0, 1e-12) && a.Mirror.CurvaturesHaveSameF2Sign && a.Mirror.ComplexConjugationNeutral, Detail: FormatMirror(a.Mirror)},
			{Name: "full doubled bosonic trace supplies eight pi branch", Passed: a.Doubled.MatchesEightPi && nearlyEqual(a.Doubled.GStarSquared, 0.5, 1e-12) && a.Doubled.MatchesHiggs && a.Doubled.Promoted, Detail: FormatLane(a.Doubled)},
			{Name: "single-carrier and quotient lanes kept visible", Passed: nearlyEqual(a.BaseLane.AlphaInverse, 4.0*3.141592653589793, 1e-12) && nearlyEqual(a.Quotient.TraceMultiplier, 0.5, 1e-12) && a.QuotientA.WrongDirectionForEightPi && !a.QuotientA.CanExplainEightPi, Detail: FormatQuotient(a.QuotientA)},
			{Name: "promotion remains conditional on bosonic trace convention", Passed: a.Promotion.MultiplierMatches && a.Promotion.ConditionalPromotion && !a.Promotion.UnconditionalDerivation && !a.Promotion.FullBosonicTraceNativeHere && a.Promotion.QuotientConventionRejected, Detail: FormatPromotion(a.Promotion)},
			{Name: "firewalls preserved", Passed: a.Audit.NoEmpiricalAlphaInserted && a.Audit.NoObservedHiggsFitInserted && a.Audit.NoPoleMassClaimed && a.Audit.NoFinalColliderMassClaimed && a.Audit.QuotientLaneKeptVisible && a.Audit.TraceConventionStillConditional && !a.Audit.FiniteCorePolluted, Detail: FormatAudit(a.Audit)},
			{Name: "summary distinguishes capacity from unconditional alpha derivation", Passed: a.Summary.FactorTwoSuppliedAsCapacity && a.Summary.EightPiConditionallyWorks && a.Summary.QuotientLaneFails && !a.Summary.NativeAlphaClosed && !a.Summary.FinalMassClaimed, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 329 conditionally identifies the missing factor two with the full doubled bosonic trace over particle and J-mirror antiparticle gauge carriers.", "It does not yet close α_GUT unconditionally: the next theorem must prove the bosonic spectral action uses the full doubled trace while fermionic real-structure quotienting does not halve the gauge kinetic a4 coefficient."}}
	}}
}
