package generation2sealedhiggssockettooneformscalarproxyandhistorylooptransportcompatibilityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate722SocketOneFormScalarProxyTransport(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate721.Inherited || !a.Gate721.PackageMinimal || !a.Gate721.SealedInterfaceDefined || !a.Gate721.ReadyOnlyAfterNQSeals || !a.Gate721.NNotDerived || !a.Gate721.QNotDerived || !a.Gate721.NoPhysicalHiggsTheorem || !a.Gate721.NoScalarPotentialOrRuntimeLambda || !a.Gate721.NoHiggsMassTheorem || !a.Gate721.NoYukawa {
		t.Fatalf("bad Gate721 inheritance: %+v", a.Gate721)
	}
	if !a.Socket.HasNSeal || !a.Socket.HasQSeal || a.Socket.ComplexDimension != 2 || !a.Socket.SU2DoubletCompatibility || !a.Socket.U1PhaseCompatibility || !a.Socket.RepresentationInterfaceAvailable {
		t.Fatalf("bad sealed socket: %+v", a.Socket)
	}
	if !a.OneForm.FiniteHiggsOneFormLaneIdentified || !a.OneForm.ComplexDimensionMatch || !a.OneForm.SU2SideCompatible || !a.OneForm.U1SideCompatible || !a.OneForm.Compatible || a.OneForm.DerivesOneForm {
		t.Fatalf("bad one-form lane: %+v", a.OneForm)
	}
	if !a.ScalarProxy.OneFormCanFeedProxyLane || a.ScalarProxy.ProxyDerivedFromSocket || a.ScalarProxy.RuntimeLambdaDerived || !a.ScalarProxy.CompatibilityOnly || math.Abs(a.ScalarProxy.LambdaProxyMZ-lambdaProxyMZ) > 1e-14 {
		t.Fatalf("bad scalar proxy lane: %+v", a.ScalarProxy)
	}
	if math.Abs(a.Transport.LoopUnit-1/(8*math.Pi)) > 1e-18 || !a.Transport.UsesHistoryLoopTransport || a.Transport.NativeHistoryLoopSource || a.Transport.NativeRuntimeTheorem || !strings.Contains(a.Transport.SubstitutedFormula, "W_72") {
		t.Fatalf("bad HistoryLoop transport: %+v", a.Transport)
	}
	if !a.LSource.PhaseUnitCandidate || !a.LSource.FourRealComponentCandidate || a.LSource.NativeFourComponentSourceProof || a.LSource.NativeHistoryLoopUnitTheorem {
		t.Fatalf("bad L source audit: %+v", a.LSource)
	}
}

func TestGate722BoundaryCompatibilityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Boundary.ResponseCoefficient-float64(k7Dim)/float64(h72Dim)) > 1e-18 || !a.Boundary.ScalarLaneConnectsHistoryWall || a.Boundary.NativeScalarFlavorBoundaryMap || math.Abs(a.Boundary.ResidualE1) > 1e-8 {
		t.Fatalf("bad boundary compatibility: %+v", a.Boundary)
	}
	if a.Firewall.SealedSocketScalarPotentialTheorem || a.Firewall.LDerivedFromHiggsRepresentation || a.Firewall.OneOver8PiNativeLoopTheorem || a.Firewall.LambdaProxyHiggsMassTheorem || a.Firewall.RuntimeLambdaPoleMassTheorem || a.Firewall.FanoK7YukawaOperatorFamily || a.Firewall.NAndQDerived {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	res := Generation2SealedHiggsSocketToOneFormScalarProxyAndHistoryLoopTransportCompatibilityAuditTheorem().Verify()
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
