package generation2quarternormalizedphasetransportsourcetypeaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate723QuarterPhaseTransportSourceType(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate722.Inherited || !a.Gate722.SocketInterfacesWithOneFormLane || !a.Gate722.ScalarProxyInterfacesWithHistoryLoop || !a.Gate722.OneOver8PiAfterScalarProxyNotRepresentation || !a.Gate722.NoNativeHistoryLoopUnit || !a.Gate722.NAndQSealedNotDerived {
		t.Fatalf("bad Gate722 inheritance: %+v", a.Gate722)
	}
	if !a.PhaseLoop.Candidate || a.PhaseLoop.NativeHistoryTransportUsesMeasure || math.Abs(a.PhaseLoop.PhaseLoopUnit-1/(2*math.Pi)) > 1e-18 {
		t.Fatalf("bad phase-loop source candidate: %+v", a.PhaseLoop)
	}
	if a.Quarter.RealCarrierDimension != 4 || a.Quarter.ComplexCarrierDimension != 2 || math.Abs(a.Quarter.QuarterFactor-0.25) > 1e-18 || math.Abs(a.Quarter.CandidateValue-1/(8*math.Pi)) > 1e-18 || !a.Quarter.EqualsHistoryLoopUnit || a.Quarter.ScalarTransportAveragesOverFourComponents {
		t.Fatalf("bad quarter normalization candidate: %+v", a.Quarter)
	}
	if !a.Placement.BelongsAfterScalarProxy || a.Placement.DerivedFromRepresentationSocketAlone {
		t.Fatalf("bad scalar transport placement: %+v", a.Placement)
	}
}

func TestGate723FirewallsAndLedger(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.QFirewall.QRescalesPhysicalChargeGenerator || !a.QFirewall.GeometricCircleUnitIndependentOfQ || a.QFirewall.NativeRelationQToL {
		t.Fatalf("bad q firewall: %+v", a.QFirewall)
	}
	if !a.NFirewall.PhaseLineDependsOnN || !a.NFirewall.LoopMeasureUniformOverTwistorLines || a.NFirewall.LSelectsN {
		t.Fatalf("bad n firewall: %+v", a.NFirewall)
	}
	if math.Abs(a.SevenFirewall.EventProbability-float64(k7Dim)/float64(h72Dim)) > 1e-18 || a.SevenFirewall.SameObject || a.SevenFirewall.SevenOver72SourcesOneOver8Pi {
		t.Fatalf("bad 7/72 firewall: %+v", a.SevenFirewall)
	}
	if math.Abs(a.Ledger.LCandidate-1/(8*math.Pi)) > 1e-18 || math.Abs(a.Ledger.TransportResidual) > 1e-15 || a.Ledger.KappaLambda < 0.044 || a.Ledger.KappaLambda > 0.045 || a.Ledger.RhoLambdaMatch < 0.038 || a.Ledger.RhoLambdaMatch > 0.039 {
		t.Fatalf("bad scalar matching ledger: %+v", a.Ledger)
	}
	if a.Firewall.NativeHistoryLoopUnitSourceTheorem || a.Firewall.NativeScalarProxyToRuntimeTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem || a.Firewall.NativeQSource || a.Firewall.NativeNSelector || a.Firewall.Native7Over72ToLTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	res := Generation2QuarterNormalizedPhaseTransportSourceTypeAuditTheorem().Verify()
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
