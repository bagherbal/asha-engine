package generation2spectralactioncoefficientgrammaraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SpectralActionCoefficientGrammarAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spectral-action coefficient grammar for gauge-scalar boundary stress seal audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate615 coefficient grammar audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate614 stress source-type audit", Passed: a.Inherited.XiBoundary > 0 && a.Inherited.Verdict == StatusGate614Inherited, Detail: FormatInherited(a.Inherited)},
			{Name: "build coefficient dependency table", Passed: len(a.Dependencies) >= 7 && hasCoefficient(a.Dependencies, "C_i") && hasCoefficient(a.Dependencies, "lambda") && hasCoefficient(a.Dependencies, "f0"), Detail: FormatDependencies(a.Dependencies)},
			{Name: "audit shared coefficient grammar", Passed: len(a.SharedAudits) == 3 && hasSharedVerdict(a.SharedAudits, StatusNoSU3OnlyDeformation), Detail: FormatSharedAudits(a.SharedAudits)},
			{Name: "audit color-specific deformation", Passed: a.ColorDeformation.BridgeExpressible && !a.ColorDeformation.NativeRepresentationTrace && a.ColorDeformation.RequiresSectorSplitF0 && a.ColorDeformation.RequiresThresholdMatching, Detail: FormatColorDeformation(a.ColorDeformation)},
			{Name: "audit scalar quartic correction", Passed: a.ScalarCorrection.BridgeExpressible && a.ScalarCorrection.ViaBA2 && a.ScalarCorrection.ViaF0 && a.ScalarCorrection.ViaMatching && !a.ScalarCorrection.Native, Detail: FormatScalarCorrection(a.ScalarCorrection)},
			{Name: "audit joint deformation", Passed: a.JointDeformation.BridgeExpressible && !a.JointDeformation.KnownNativeRelation && !a.JointDeformation.ForcesStressEquation && a.JointDeformation.ResidualOverXi < 0.03, Detail: FormatJointDeformation(a.JointDeformation)},
			{Name: "check type-safe normalized shadows", Passed: !a.TypeConsistency.RawComparisonSafe && a.TypeConsistency.NormalizedSafe && len(a.TypeConsistency.NormalizedForms) == 3, Detail: FormatTypeConsistency(a.TypeConsistency)},
			{Name: "record native obstruction ledger", Passed: !a.NativeObstructions.NativeXi && !a.NativeObstructions.NativeSU3Only && !a.NativeObstructions.NativeC3LambdaLaw && !a.NativeObstructions.NativeF0Split && !a.NativeObstructions.NativeLambdaBC && !a.NativeObstructions.NativeThresholds, Detail: FormatNativeObstructions(a.NativeObstructions)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsXiNative && !a.Firewalls.ClaimsLambdaZero && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsThresholdExistence && !a.Firewalls.ClaimsNativeCorrection && !a.Firewalls.ClaimsObservedEndpointDerivation, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasCoefficient(rows []CoefficientDependency, name string) bool {
	for _, r := range rows {
		if r.Coefficient == name {
			return true
		}
	}
	return false
}

func hasSharedVerdict(rows []SharedCoefficientAudit, verdict string) bool {
	for _, r := range rows {
		if r.Verdict == verdict {
			return true
		}
	}
	return false
}
