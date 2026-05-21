package generation2boundaryactivationmeasurenativeconstraintsourceaudit

import (
	"strings"
	"testing"
)

func TestGate922ConstraintSourceLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.ReducedResponse.SourceStatus != SourceBridgeStrong || !a.Ledger.ReducedResponse.BridgeLawful || a.Ledger.ReducedResponse.Native {
		t.Fatalf("bad reduced response source: %s", FormatConstraint(a.Ledger.ReducedResponse))
	}
	if a.Ledger.DegreeRespect.SourceStatus != SourceNativeShapeStrong || !a.Ledger.DegreeRespect.Native || !a.Ledger.DegreeRespect.BridgeLawful {
		t.Fatalf("bad degree source: %s", FormatConstraint(a.Ledger.DegreeRespect))
	}
	if a.Ledger.SelectorFunctionhood.SourceStatus != SourceBridgeCandidateNotNative || !a.Ledger.SelectorFunctionhood.Primary || a.Ledger.SelectorFunctionhood.Native {
		t.Fatalf("bad selector source: %s", FormatConstraint(a.Ledger.SelectorFunctionhood))
	}
	if a.Ledger.CrossLaneExclusion.SourceStatus != SourceDependentOnSelector || !a.Ledger.CrossLaneExclusion.Dependent || a.Ledger.CrossLaneExclusion.Native {
		t.Fatalf("bad cross-lane source: %s", FormatConstraint(a.Ledger.CrossLaneExclusion))
	}
}

func TestGate922ChamberZ2PositivityAndNativeStatus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Ledger.ChamberNormalization.SourceStatus != SourceBridgeStrong || !a.Ledger.ChamberNormalization.BridgeLawful || a.Ledger.ChamberNormalization.Native {
		t.Fatalf("bad chamber source: %s", FormatConstraint(a.Ledger.ChamberNormalization))
	}
	if a.Ledger.Z2Independence.SourceStatus != SourceBridgeStrongOrientation || !a.Ledger.Z2Independence.BridgeLawful || a.Ledger.Z2Independence.Native {
		t.Fatalf("bad z2 source: %s", FormatConstraint(a.Ledger.Z2Independence))
	}
	if a.Ledger.Positivity.SourceStatus != SourceCompatibilityOnly || !a.Ledger.Positivity.Compatible || a.Ledger.Positivity.Native {
		t.Fatalf("bad positivity source: %s", FormatConstraint(a.Ledger.Positivity))
	}
	if !a.NativeStatus.ConstraintsPartlySourced || !a.NativeStatus.DegreeRespectStrongestNativeShape || !a.NativeStatus.SelectorFunctionhoodPrimaryGap || a.NativeStatus.NativeBoundaryActivationMeasure || a.NativeStatus.NativeAlpha || a.NativeStatus.NativeR3 {
		t.Fatalf("bad native status: %s", FormatNativeStatus(a.NativeStatus))
	}
}

func TestGate922AlphaAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.Formula != MeasureFormula || a.Alpha.BoundaryFormula != BoundaryAlphaFormula || a.Alpha.NativeAlpha {
		t.Fatalf("bad alpha metadata: %s", FormatAlpha(a.Alpha))
	}
	if !near(a.Alpha.LinearContribution, AlphaLinear) || !near(a.Alpha.QuadraticContribution, AlphaQuad) || !near(a.Alpha.Alpha, AlphaB) {
		t.Fatalf("bad alpha values: %s", FormatAlpha(a.Alpha))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeCrossLaneWithoutSelector, FailureNoNativeLaneLocalityToChamber, FailureNoNativeResponseChamberNormalization, FailureNoNativeGlobalPhaseZ2Equivariance, FailureNoNativeBasepointDeviation, FailurePositivityNotSelectionTheorem, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate922Theorem(t *testing.T) {
	res := Generation2BoundaryActivationMeasureNativeConstraintSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate, BoundaryMeasureObject, MeasureFormula, BranchMeasureFormula, BoundaryAlphaFormula, StrongestNativeShapeSource, StrongestBridgeSources, PrimaryGapSelectorFunctionhood} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
