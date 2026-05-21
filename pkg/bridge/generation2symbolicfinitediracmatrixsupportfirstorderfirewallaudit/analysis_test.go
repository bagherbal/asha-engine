package generation2symbolicfinitediracmatrixsupportfirstorderfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate848YSupportMatrix(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Y.SupportOnly || a.Y.ActiveFamilies != 3 || len(a.Y.Edges) != 3 || !a.Y.PreservesLeptoColor || a.Y.HasNumericalValues {
		t.Fatalf("bad Y support: %s", FormatY(a.Y))
	}
	if !allActiveEdgesSymbolicNoMagnitude(a.Y.Edges) {
		t.Fatalf("bad active symbolic edges: %s", FormatY(a.Y))
	}
	if !a.Y.PunctureCoefficientZero || a.Y.PunctureCoefficient != "y_+1" || !a.Y.MissingEdge.Puncture || a.Y.MissingEdge.Present || a.Y.MissingEdge.Coefficient != "y_+1=0" {
		t.Fatalf("bad puncture coefficient: %s", FormatEdge(a.Y.MissingEdge))
	}
}

func TestGate848ChiralDiracSupportMatrix(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Dirac.ExplicitSupportMatrix || a.Dirac.NativeDFMatrix || a.Dirac.NumericalDFMatrix {
		t.Fatalf("D_F support matrix promoted incorrectly: %s", FormatDirac(a.Dirac))
	}
	if a.Dirac.LeftRank != HLRank || a.Dirac.RightRank != HRMinRank || a.Dirac.TotalRank != ChiralTotalDim || a.Dirac.BlockRows != ChiralTotalDim || a.Dirac.BlockCols != ChiralTotalDim {
		t.Fatalf("bad chiral matrix dimensions: %s", FormatDirac(a.Dirac))
	}
	if !a.Dirac.UsesAdjointBlock || !a.Dirac.SelfAdjointByConstruction || !a.Dirac.ChiralOddByConstruction {
		t.Fatalf("self-adjoint/chiral support checks failed: %s", FormatDirac(a.Dirac))
	}
}

func TestGate848FirstOrderAndMagnitudeFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Dirac.FirstOrderCertified || a.Dirac.BimoduleCommutantProof || a.Dirac.JOppositeCompatibilityProof {
		t.Fatalf("first-order/J/bimodule overpromoted: %s", FormatDirac(a.Dirac))
	}
	if a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing {
		t.Fatalf("impact overpromoted: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoFirstOrderProof || !a.Firewalls.NoBimoduleCommutantProof || !a.Firewalls.NoJOppositeCompatibilityProof || !a.Firewalls.YSymbolsNotYukawaValues || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate848 {
		t.Fatalf("firewalls invalid: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate848Theorem(t *testing.T) {
	res := Generation2SymbolicFiniteDiracMatrixSupportFirstOrderFirewallAuditTheorem().Verify()
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
