package generation2ssplittoreducedboundarypairresponsetransportaudit

import (
	"strings"
	"testing"
)

func TestGate916InheritedSubobjectsDoNotPromoteNativeR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ResponseShapeCertified || !a.Inherited.SelectorShapeTyped || !a.Inherited.CrossLanesBlocked || a.Inherited.DerivesAlpha || a.Inherited.UpdatesOfficialLedger || a.Inherited.PromotesNativeR3 {
		t.Fatalf("bad inherited state: %s", FormatInherited(a.Inherited))
	}
}

func TestGate916TransportTargetIsRBParameterOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Target.Target != TransportTarget || a.Target.TargetsAlphaDirectly || a.Target.TargetsSocketMag || !a.Target.UsesRBOnly || a.Target.NativeMap {
		t.Fatalf("bad transport target: %s", FormatTarget(a.Target))
	}
}

func TestGate916SingleInsertionGeneratesQuadraticTerm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Insertion.Insertions != 2 || !a.Insertion.ScalarInsertedPerFactor || a.Insertion.SeparateQuadraticTransport || !a.Insertion.QuadraticFromProduct || a.Insertion.NativeUniformLaw {
		t.Fatalf("bad insertion audit: %s", FormatInsertion(a.Insertion))
	}
}

func TestGate916ScalarCompatibilityAndBasepointReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Scalar.Dimensionless || !a.Scalar.CanMultiplyGenerators || a.Scalar.ScalarTypeNative {
		t.Fatalf("bad scalar compatibility: %s", FormatScalar(a.Scalar))
	}
	if !a.Reduction.TransportAppliesToActive || a.Reduction.IdentityTransported || !a.Reduction.BasepointRemoved || a.Reduction.NativeBasepointTheorem {
		t.Fatalf("bad reduction: %s", FormatReduction(a.Reduction))
	}
}

func TestGate916SelectorCompatibilityDoesNotReopenCrossLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selector.DegreeSelectorCompatible || !a.Selector.FeedsCorrectAlphaLanes || a.Selector.ReopensCrossLanePollution || a.Selector.NativeTransportTheorem {
		t.Fatalf("bad selector compatibility: %s", FormatSelector(a.Selector))
	}
}

func TestGate916AlphaReconstructionUnderTransportSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Alpha.AlphaLinear, AlphaLinear) || !near(a.Alpha.AlphaQuadratic, AlphaQuad) || !near(a.Alpha.AlphaTotal, AlphaB) || !a.Alpha.TransportSealAssumed || !a.Alpha.PriorSubobjectsAssumed || a.Alpha.NativeAlphaTheorem {
		t.Fatalf("bad alpha reconstruction: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate916GlobalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeTsMap, FailureNoNativeTransportToZ2Airlock, FailureNoTypedSSplitExteriorParameterMap, FailureNoNativeUniformInsertionReason, FailureSSplitScalarTypeSealed, FailureNoNativeBasepointReduction, FailureSelectorCompatibilityNotNative, FailureAlphaReconstructionNotNative, FailureDenominatorNormalizationExternal, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s from %v", want, a.FirewallsList())
		}
	}
}

func TestGate916Theorem(t *testing.T) {
	res := Generation2SSplitToReducedBoundaryPairResponseTransportAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s in notes: %s", want, joined)
		}
	}
}
