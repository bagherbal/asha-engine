package generation2minimalrightmodulefinitediracedgeskeletonaudit

import (
	"strings"
	"testing"
)

func TestGate844DomainAndTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Domain.InheritedFromGate843 || !a.Domain.MinimalAbsenceSeal || a.Domain.ActiveRank != 7 || a.Domain.PunctureRank != 1 || !a.Domain.ActiveIsFullMinusPuncture {
		t.Fatalf("bad domain: %s", FormatDomain(a.Domain))
	}
	if a.Domain.BMinusLActive != 1 || a.Domain.BMinusLPuncture != -1 || a.Domain.BMinusLFull != 0 {
		t.Fatalf("bad B-L compensation: %s", FormatDomain(a.Domain))
	}
	if !a.Target.Complete || !a.Target.LeptoColorPreserved || a.Target.Rank != 8 || a.Target.ColorRank != 6 || a.Target.LeptonRank != 2 {
		t.Fatalf("bad left target: %s", FormatTarget(a.Target))
	}
}

func TestGate844SymbolicEdgeSupport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edges.ConstructedAtSealLevel || !a.Edges.EdgeSupportOnly || a.Edges.DomainRank != 7 || a.Edges.TargetRank != 8 || len(a.Edges.Edges) != 3 {
		t.Fatalf("bad edge skeleton: %s", FormatEdges(a.Edges))
	}
	for _, e := range a.Edges.Edges {
		if !e.ColorLeptonPreserving || !e.SymbolicOnly || e.DomainRank <= 0 || e.TargetRank <= 0 {
			t.Fatalf("bad edge: %+v", e)
		}
	}
	if a.Edges.Edges[0].TargetExpression != "C_L^2 tensor P_3" || a.Edges.Edges[1].TargetExpression != "C_L^2 tensor P_3" || a.Edges.Edges[2].TargetExpression != "C_L^2 tensor P_1" {
		t.Fatalf("edge support failed to preserve lepto-color blocks: %s", FormatEdges(a.Edges))
	}
}

func TestGate844PunctureAndDFNativeFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Edges.PunctureInDomain || !a.Edges.PunctureAbsenceCompatible || a.Edges.PunctureNullEdgeCertified || a.Edges.PunctureAbsenceDerivedFromDFNullEdge {
		t.Fatalf("puncture status over/under-certified: %s", FormatEdges(a.Edges))
	}
	if a.Edges.NativeDFMatrixCertified || a.Edges.ExplicitDFMatrixCertified || a.Edges.FirstOrderConditionCertified || a.Edges.BimoduleCommutantCertified || a.Edges.YukawaMagnitudes || a.Edges.NumericalValues {
		t.Fatalf("D_F edge skeleton over-certified: %s", FormatEdges(a.Edges))
	}
}

func TestGate844ShadowAndLedgerFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Shadow.FiniteBodyLocationAtSealLevel || !a.Shadow.EdgeSupportSealLevel || a.Shadow.NativeCompressionTheorem || a.Shadow.AlphaDerived || a.Shadow.TraceMagnitudeReadout || a.Shadow.R3 || a.Shadow.R4 {
		t.Fatalf("shadow invalid: %s", FormatShadow(a.Shadow))
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		t.Fatalf("ledger invalid: %s", FormatLedger(a.Ledger))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.DFSupportSealNotNative || !a.Firewalls.NoExplicitDFMatrix || !a.Firewalls.NoFirstOrderProof || !a.Firewalls.NoBimoduleCommutantProof || !a.Firewalls.PunctureAbsenceNotFromNullEdge || !a.Firewalls.DFEdgeSupportNotYukawa || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate844 {
		t.Fatalf("firewalls invalid: %+v", a.Firewalls)
	}
}

func TestGate844Theorem(t *testing.T) {
	res := Generation2MinimalRightModuleFiniteDiracEdgeSkeletonAuditTheorem().Verify()
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
