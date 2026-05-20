package generation2selfconsistentrestconcentrationlawandboundaryalphamapaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate817BoundaryAlphaAndExactClosure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.DeltaNBFN-6*a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("Delta_BFN != 6 alpha_B: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Closure.BetaByFormula-a.Closure.BetaBySimplifiedIdentity) > 1e-15 {
		t.Fatalf("beta identity failed: %s", FormatClosure(a.Closure))
	}
	if math.Abs(a.Closure.QRest-1/a.Ledger.NEffBFN) > 1e-18 {
		t.Fatalf("q rest not inverse N_eff_BFN: %s", FormatClosure(a.Closure))
	}
	if !containsAll(a.Alpha.Failures, []string{StatusAlphaNotTheorem}) || !containsAll(a.Closure.Failures, []string{StatusQNoIndependentSource}) {
		t.Fatalf("missing theorem firewalls")
	}
}

func TestGate817PositiveSpectrumConstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	three := spectrumByName(a.Spectra, "diffuse three-rest construction")
	if three.Exact || three.Q <= a.Ledger.QRestB {
		t.Fatalf("three equal rest atoms should be close but above target q: %s", FormatSpectra(a.Spectra))
	}
	one := spectrumByName(a.Spectra, "concentrated one-rest construction")
	if one.Exact || math.Abs(one.Q-1) > 1e-15 {
		t.Fatalf("one rest atom should fail with q=1: %s", FormatSpectra(a.Spectra))
	}
	mixed := spectrumByName(a.Spectra, "mixed four-rest construction small-support branch")
	if !mixed.Exact || math.Abs(mixed.Sum-1) > 1e-15 || math.Abs(mixed.Q-a.Ledger.QRestB) > 1e-15 {
		t.Fatalf("mixed four-rest construction should exactly realize target q: %s", FormatSpectra(a.Spectra))
	}
}

func TestGate817StatusImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Map.Level, "partial R2") || a.Map.ConstructsSectorLedger || a.Map.NativeYukawaTheorem {
		t.Fatalf("bad map status: %+v", a.Map)
	}
	if math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) > 1e-15 || math.Abs(a.Impact.CHiggsBFN-1.0372205108665146) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	res := Generation2SelfConsistentRestConcentrationLawAndBoundaryAlphaMapAuditTheorem().Verify()
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
