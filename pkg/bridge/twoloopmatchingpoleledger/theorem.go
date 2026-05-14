package twoloopmatchingpoleledger

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TwoLoopMatchingPoleMassConversionLedgerAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TWO-LOOP-MATCHING-POLE-MASS-CONVERSION-LEDGER-AUDIT"
	const name = "Two-Loop / Matching / Pole-Mass Conversion Ledger Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 310 higher-order transport ledger", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 309 one-loop Higgs-mass diagnostic is inherited as tension, not final prediction", Passed: a.Inheritance.InheritedAsDiagnostic && a.Inheritance.OneLoopOnly && a.Inheritance.PrimaryLambdaAtV > 0.90 && a.Inheritance.PrimaryLambdaAtV < 0.91 && a.Inheritance.PrimaryMassGeV > 330 && a.Inheritance.PrimaryMassGeV < 333 && a.Inheritance.PureSMHighScaleRejected && a.Inheritance.ThresholdMatchingOmitted && a.Inheritance.PoleMassMatchingOmitted && !a.Inheritance.FinalColliderMassClaimed, Detail: FormatInheritance(a.Inheritance)},
			{Name: "two-loop RG ledger is formalized without installing an unverified coefficient table", Passed: a.TwoLoop.Formalized && a.TwoLoop.PositiveBetaSoftensDownward && a.TwoLoop.NegativeBetaAmplifiesDownward && len(a.TwoLoop.RepresentativeTerms) >= 4 && !a.TwoLoop.ExactFullSystemInstalled && !a.TwoLoop.TwoLoopIntegrationExecuted && !a.TwoLoop.CanResolveAlone, Detail: FormatTwoLoop(a.TwoLoop)},
			{Name: "threshold matching ledger formalizes finite jumps and identifies the only correction class with enough capacity", Passed: a.Thresholds.Formalized && a.Thresholds.HasCapacityToResolve && !a.Thresholds.ValuesDerived && !a.Thresholds.ThresholdsExecuted && a.Thresholds.RequiredIRLambdaShift < -0.7 && len(a.Thresholds.Sources) >= 3, Detail: FormatThresholds(a.Thresholds)},
			{Name: "pole-mass conversion ledger is formalized but cannot erase the one-loop tension alone", Passed: a.PoleMass.Formalized && a.PoleMass.RequiredLambdaAtV > 0.12 && a.PoleMass.RequiredLambdaAtV < 0.14 && a.PoleMass.RequiredLambdaShiftAtV < -0.7 && !a.PoleMass.CanResolveAlone && !a.PoleMass.SelfEnergiesComputed && !a.PoleMass.UsesMeasuredMassForFit, Detail: FormatPole(a.PoleMass)},
			{Name: "tension capacity audit separates precision corrections from structural resolution classes", Passed: a.Tension.Formalized && a.Tension.LambdaExcessAtV > 0.7 && a.Tension.MassExcessGeV > 200 && !a.Tension.TwoLoopCanResolveAlone && !a.Tension.PoleMassCanResolveAlone && a.Tension.ThresholdsCanResolveInPrinciple && a.Tension.ModifiedTopSectorMayBeRequired && a.Tension.NeedsFullPrecisionRun && !a.Tension.FinalMassResolved, Detail: FormatTension(a.Tension)},
			{Name: "firewalls prevent inserting threshold, two-loop, pole, Higgs, or top data as a fit", Passed: a.Firewalls.NoTwoLoopNumericalTransportRun && a.Firewalls.NoThresholdJumpInserted && a.Firewalls.NoPoleSelfEnergyInserted && a.Firewalls.NoObservedHiggsUsedAsFit && a.Firewalls.NoObservedTopUsedAsFit && a.Firewalls.NoFinalMassClaimed && a.Firewalls.NoFiniteCorePolluted && len(a.Firewalls.Obligations) >= 4, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary authorizes the higher-order ledger but not a resolved Higgs mass", Passed: a.Summary.Gate309Inherited && a.Summary.TwoLoopLedgerReady && a.Summary.ThresholdLedgerReady && a.Summary.PoleMassLedgerReady && a.Summary.CapacityAssessed && !a.Summary.FinalMassResolved && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 310 is a precision-infrastructure theorem: it does not numerically correct the Gate 309 mass.", "The next legal move is a threshold-sensitivity matrix or an explicitly derived top-sector deformation, not empirical tuning."}}
	}}
}
