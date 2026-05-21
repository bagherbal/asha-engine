package generation2positivephasegeneratorcharacterweightorientationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_POSITIVE_PHASE_GENERATOR_CHARACTER_WEIGHT_ORIENTATION_AUDIT"
	theoremName = "Gate 906 — PositivePhase Generator and CharacterWeight Orientation Audit"
)

func Generation2PositivePhaseGeneratorCharacterWeightOrientationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 905 inherited pair-not-order wound", Passed: a.Inherited.PhaseModuleInduced && a.Inherited.PairCertified && !a.Inherited.OrderCertified && containsAll(a.Inherited.Failures, []string{FailureTwoCharPairNotOrder}), Detail: FormatInherited(a.Inherited)},
			{Name: "Q_phi exists and lambda/barlambda order is equivalent to its sign orientation", Passed: a.WeightOperator.Exists && a.WeightOperator.OrderEquivalentToSign && !a.WeightOperator.PositiveSignNative && containsAll(a.WeightOperator.Supports, []string{SupportQPhiExists, SupportOrderEquivalentQPhi, SupportWoundToQPhi}) && containsAll(a.WeightOperator.Failures, []string{FailureQPhiSignNotNative, FailureNoLambdaSelection}), Detail: FormatWeightOperator(a.WeightOperator)},
			{Name: "Hopf Reeb is strongest positive phase generator candidate but not typed action on C_R2", Passed: a.HopfReeb.SuppliesPositiveGeneratorCandidate && a.HopfReeb.SelectsLambdaWeightIfSealed && !a.HopfReeb.TypedActionOnCR2 && !a.HopfReeb.NativeSelector && containsAll(a.HopfReeb.Supports, []string{SupportHopfReebStrongest, SupportHopfPositiveWeight}) && containsAll(a.HopfReeb.Failures, []string{FailureNoTypedHopfReebToCR2, FailureNoNativePositiveGenerator}), Detail: FormatHopfReeb(a.HopfReeb)},
			{Name: "Cl17 chirality sign matches phase weight sign but lacks typed map to Q_phi", Passed: a.CL17Sign.SignMatchesPhaseWeight && a.CL17Sign.CanSourceQPhiIfTyped && !a.CL17Sign.TypedMapToQPhi && !a.CL17Sign.SelectsSocketOrder && containsAll(a.CL17Sign.Supports, []string{SupportCL17SignMatches, SupportGammaChiCanSource}) && containsAll(a.CL17Sign.Failures, []string{FailureNoTypedGammaChiToQPhi, FailureCL17DoesNotSelect}), Detail: FormatCL17Sign(a.CL17Sign)},
			{Name: "J confirms conjugate pair but does not orient Q_phi", Passed: a.JConjugation.ExchangesWeights && a.JConjugation.ExplainsPair && !a.JConjugation.OrientsSign && containsAll(a.JConjugation.Failures, []string{FailureJDoesNotOrient}), Detail: FormatJConjugation(a.JConjugation)},
			{Name: "boundary orientation does not select phase weight sign", Passed: a.BoundaryOrientation.OrientsExteriorDegree && !a.BoundaryOrientation.SelectsPhaseWeightSign && containsAll(a.BoundaryOrientation.Failures, []string{FailureBoundaryNoPhaseSign}), Detail: FormatBoundaryOrientation(a.BoundaryOrientation)},
			{Name: "master wound reduced to positive phase generator selection", Passed: !a.Wound.NativeSolved && containsAll(a.Wound.Supports, []string{SupportOrderedAirlockIfSealed, SupportR3WoundToGenerator}) && containsAll(a.Wound.Failures, []string{FailureNoNativePositiveGenerator, FailurePhaseAnchorSealed}), Detail: FormatWound(a.Wound)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, phase generator, alpha, Higgs orientation, physical-sector, generation/flavor, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativePositiveGenerator, FailureNoTypedHopfReebToCR2, FailureNoTypedGammaChiToQPhi, FailureQPhiSignNotNative, FailureNotNativeR3}), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatWeightOperator(a.WeightOperator), FormatHopfReeb(a.HopfReeb), FormatCL17Sign(a.CL17Sign), FormatJConjugation(a.JConjugation), FormatBoundaryOrientation(a.BoundaryOrientation), FormatWound(a.Wound), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
