package generation2minimalnulledgeorientationprincipleaudit

import (
	"strings"
	"testing"
)

func TestGate894NullEdgeMinimization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NullEdge
	if !n.YPlus1Zero || n.AmbientRightRank != 8 || n.ActiveRightRank != 7 || n.RankReduction != 1 || !n.MinimalRankSevenCandidate || n.NativePrinciple {
		t.Fatalf("bad null-edge minimization: %s", FormatNullEdgeMinimization(n))
	}
	if !containsAll(n.Supports, []string{SupportMinimalNullEdgeSelectsHPlusCandidate, SupportMinimalRankSevenEdgeDomain}) || !containsAll(n.Failures, []string{FailureNoNativeMinimalNullEdgeOrientation}) {
		t.Fatalf("missing null-edge statuses: %s", FormatNullEdgeMinimization(n))
	}
}

func TestGate894ImageKernelReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	k := a.ImageKernel
	if k.HLeftRank != 8 || k.ImageRank != 7 || k.Kernel != LeftKernel || k.KernelRank != 1 || !k.QuotientIsHPlusP1 || !k.ReconstructsKernel || k.SelectsFrameNonCircularly {
		t.Fatalf("bad image/kernel audit: %s", FormatImageKernel(k))
	}
	if !containsAll(k.Failures, []string{FailureKernelLineDependsOnEdgeSupportChoice, FailureDFPatternRestatesOrientation}) {
		t.Fatalf("missing image/kernel firewalls: %s", FormatImageKernel(k))
	}
}

func TestGate894NonCircularityStillBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonCircularity
	if n.CanDefineHPlusFromKernelWithoutPriorFrame || !n.EdgeSupportAssumesFrame || !n.RequiresVariationalMinimality || n.NativeSelectorFunctional {
		t.Fatalf("noncircularity leak: %s", FormatNonCircularity(n))
	}
	if !strings.Contains(n.MissingObject, "WeakSocketSelectorFunctional") || !strings.Contains(n.MissingObject, "MinimalNullEdgeOrientationPrinciple") {
		t.Fatalf("wrong missing object: %s", FormatNonCircularity(n))
	}
}

func TestGate894FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate894Theorem(t *testing.T) {
	res := Generation2MinimalNullEdgeOrientationPrincipleAuditTheorem().Verify()
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
