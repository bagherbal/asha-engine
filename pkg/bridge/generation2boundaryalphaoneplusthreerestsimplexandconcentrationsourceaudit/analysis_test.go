package generation2boundaryalphaoneplusthreerestsimplexandconcentrationsourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate818PriorAlphaSimplex(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.DeltaNBFN-6*a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("Delta_BFN != 6 alpha_B: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Prior.NEffResidual) > 3e-16 || math.Abs(a.Prior.SymbolicResidual+2.107593378826735e-16) > 1e-27 {
		t.Fatalf("bad fifth-order residual: %s", FormatPrior(a.Prior))
	}
	if !(a.Prior.QSimplex > a.Ledger.QRestB) {
		t.Fatalf("q_simplex(alpha_B) should be slightly above q_rest_B: %s", FormatPrior(a.Prior))
	}
}

func TestGate818TStarControlsAndSpectrum(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.TStar.TStar-a.Ledger.AlphaB-2.336e-10) > 5e-14 || math.Abs(a.TStar.QResidual) > 1e-15 {
		t.Fatalf("bad t_star branch: %s", FormatTStar(a.TStar))
	}
	three := controlByName(a.Controls, "three equal rest atoms")
	one := controlByName(a.Controls, "one concentrated rest atom")
	if three.Exact || three.Q <= a.Ledger.QRestB || one.Exact || math.Abs(one.Q-1) > 1e-15 {
		t.Fatalf("bad controls: %s", FormatControls(a.Controls))
	}
	if !a.Spectrum.Realizable || math.Abs(a.Spectrum.Sum-1) > 1e-15 || a.Spectrum.Q <= 0 || a.Spectrum.Q >= 1 {
		t.Fatalf("bad abstract spectrum: %s", FormatSpectrum(a.Spectrum))
	}
}

func TestGate818StatusImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Status.Level, "strengthened partial R2") || a.Status.ConstructsSectorLedger || a.Status.NativeYukawaTheorem {
		t.Fatalf("bad status: %+v", a.Status)
	}
	if math.Abs(a.Impact.CYukawaCandidate-0.999224809692266) > 1e-15 || math.Abs(a.Impact.CHiggsCandidate-1.0372205108665148) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	res := Generation2BoundaryAlphaOnePlusThreeRestSimplexAndConcentrationSourceAuditTheorem().Verify()
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
