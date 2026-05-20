package generation2twistorinvariantsu2doubletsocketrepresentationaudit

import (
	"strings"
	"testing"
)

func TestGate715ComplexLinearityTraceZeroAndDoubletShape(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TwistorInvariantSocketInherited || a.Inherited.CommonCommutantDimension != 3 || !a.Inherited.CommonCommutantInAllSockets || !a.Inherited.IntersectionEqualsCommutant || !a.Inherited.PhaseLineSelectorDependent || a.Inherited.SelectorIndependentU1Line {
		t.Fatalf("bad Gate714 inheritance: %+v", a.Inherited)
	}
	if !a.ComplexLinearity.CommutesWithEveryJH || !a.ComplexLinearity.ComplexLinearEveryJH || !a.ComplexLinearity.ActsOnEachC2Carrier || a.ComplexLinearity.PhysicalSU2LCertified || !strings.Contains(a.ComplexLinearity.Verdict, StatusCCommutantComplexLinearForEveryJH) {
		t.Fatalf("bad complex-linearity audit: %+v", a.ComplexLinearity)
	}
	if !a.UnitaryAction.CSubsetSO4 || !a.UnitaryAction.SkewForRealMetric || !a.UnitaryAction.LiesInU2EveryJH || a.UnitaryAction.U2Dimension != 4 {
		t.Fatalf("bad unitary action audit: %+v", a.UnitaryAction)
	}
	if !a.TraceZero.ComplexTraceZero || !a.TraceZero.LiesInSU2EveryJH || a.TraceZero.CommutantDimension != 3 || !a.TraceZero.PhaseLineExcluded || a.TraceZero.HyperchargeCertified {
		t.Fatalf("bad trace-zero audit: %+v", a.TraceZero)
	}
	if a.Doublet.RealDimension != 4 || a.Doublet.ComplexDimension != 2 || !a.Doublet.CClosesAsSU2Like || !a.Doublet.ComplexIrreducible || !a.Doublet.DoubletShapeCertified || a.Doublet.PhysicalDoubletMap {
		t.Fatalf("bad doublet audit: %+v", a.Doublet)
	}
}

func TestGate715TwistorInvariancePhysicalFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Twistor.CommonCommutantIndependentOfN || !a.Twistor.IncludedForEveryJH || !a.Twistor.PhaseLineMovesWithN || !a.Twistor.SU2SocketTwistorInvariant || !a.Twistor.U1PhaseSelectorDependent || !strings.Contains(a.Twistor.Verdict, StatusTwistorInvarianceOfCAudited) {
		t.Fatalf("bad twistor invariance audit: %+v", a.Twistor)
	}
	f := a.PhysicalFirewall
	if f.InternalDoubletSocketPhysicalSU2L || f.TypedThetaSU2Intertwiner || f.U1HyperchargeSelectorIndependent || f.HyperchargeAssignment || f.HyperchargeNormalization || f.TypedHiggsDoubletMap || f.YukawaOperator || f.YukawaEigenvalues || f.HiggsMass || f.ScalarRuntime || len(f.MissingMaps) != 4 || !strings.Contains(f.Verdict, StatusNoTypedThetaSU2Intertwiner) || !strings.Contains(f.Verdict, StatusNoHyperchargeAssignmentOrNormalization) {
		t.Fatalf("physical firewall violated: %+v", f)
	}
	if !a.Strategy.AirlockReady || !strings.Contains(a.Strategy.SU2Side, "complex-linearly") || !strings.Contains(a.Strategy.U1Side, "selector-dependent") || !strings.Contains(a.Strategy.YukawaSide, "Yukawa") {
		t.Fatalf("bad strategy: %+v", a.Strategy)
	}
	res := Generation2TwistorInvariantSU2DoubletSocketRepresentationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
