package generation2minimalrightmoduleweakdoubletsocketedgeoperatoraudit

import (
	"strings"
	"testing"
)

func TestGate847WeakSocketSplit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Weak.Complete || !a.Weak.Orthogonal || a.Weak.RankHPlus != 1 || a.Weak.RankHMinus != 1 || a.Weak.WeakDim != 2 {
		t.Fatalf("bad weak split: %s", FormatWeak(a.Weak))
	}
	if !a.Weak.OrientationSeal || a.Weak.NativeSplit || a.Weak.HiggsOrientationCertified {
		t.Fatalf("weak split promoted incorrectly: %s", FormatWeak(a.Weak))
	}
}

func TestGate847SymbolicEdges(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Edges.ActiveEdges != 3 || len(a.Edges.Edges) != 3 || !a.Edges.PreservesLeptoColor || !a.Edges.ReconstructsGate846Cells {
		t.Fatalf("bad edge operator: %s", FormatEdges(a.Edges))
	}
	if !allEdgesPresentNoMagnitude(a.Edges.Edges) || !edgesPreserveLeptoColor(a.Edges.Edges) {
		t.Fatalf("bad active edges: %s", FormatEdges(a.Edges))
	}
	if !a.Edges.MissingEdge.Puncture || a.Edges.MissingEdge.Present || a.Edges.MissingEdge.Domain != "e_+ tensor P_1" || a.Edges.MissingEdge.Target != "h_+ tensor P_1" {
		t.Fatalf("bad missing edge: %s", FormatEdge(a.Edges.MissingEdge))
	}
}

func TestGate847Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edges.SupportOnly || a.Edges.ExplicitDFMatrix || a.Edges.NativeDFMatrix || a.Edges.FirstOrderCertified || a.Edges.BimoduleCommutantProof {
		t.Fatalf("edge operator overpromoted: %s", FormatEdges(a.Edges))
	}
	if a.Edges.Magnitudes || a.Edges.AlphaDerived || a.Edges.ParticleAssignment || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing {
		t.Fatalf("magnitude/alpha firewall failed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoExplicitDFMatrix || !a.Firewalls.NoFirstOrderProof || !a.Firewalls.EdgeSupportNotYukawa || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate847 {
		t.Fatalf("firewalls invalid: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate847Theorem(t *testing.T) {
	res := Generation2MinimalRightModuleWeakDoubletSocketEdgeOperatorAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
