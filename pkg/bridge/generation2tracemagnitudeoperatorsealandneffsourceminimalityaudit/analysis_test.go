package generation2tracemagnitudeoperatorsealandneffsourceminimalityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate807SealFormulasAndTopColorLimit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Seal.Defined || a.Seal.Name != "TraceMagnitudeOperatorSeal" {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	for _, want := range []string{"sector Hermitian operators H_u,H_d,H_e,H_nu", "positive spectra Spec(H_f)", "top-dominant block selector", StatusSealNotNative} {
		if !strings.Contains(FormatSeal(a.Seal), want) {
			t.Fatalf("seal missing %s: %s", want, FormatSeal(a.Seal))
		}
	}
	if !strings.Contains(FormatFormulas(a.Formulas), "N_eff=1/sum_i w_i²") {
		t.Fatalf("participation identity missing: %s", FormatFormulas(a.Formulas))
	}
	neffTop, err := TopColorNEff(0.9)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(neffTop-3) > 1e-12 || math.Abs(a.TopColor.NEffTop-3) > 1e-12 {
		t.Fatalf("bad top-color limit: %.17g / %.17g", neffTop, a.TopColor.NEffTop)
	}
}

func TestGate807OrientationInvisibilityAndRestPressure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Orientation.Audited || !containsAll(a.Orientation.Failures, []string{StatusNoPMNSCKMFromNEff, StatusNoKappaOrientFromTrace}) {
		t.Fatalf("bad orientation audit: %s", FormatOrientation(a.Orientation))
	}
	neff, err := RestPressureNEff(0.01, 0.005)
	if err != nil {
		t.Fatal(err)
	}
	want := 3 * (1.01 * 1.01) / 1.005
	if math.Abs(neff-want) > 1e-15 {
		t.Fatalf("bad rest formula: got %.17g want %.17g", neff, want)
	}
	delta, err := RestPressureDelta(0.01, 0.005)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(delta-(neff-3)) > 1e-15 {
		t.Fatalf("bad delta: %.17g %.17g", delta, neff-3)
	}
	if !containsAll(a.Rest.Failures, []string{StatusNoAlphaBetaWithoutT, StatusNoSectorRestAssignment}) {
		t.Fatalf("missing rest failures: %s", FormatRest(a.Rest))
	}
}

func TestGate807SourceAuditsScaleCHiggsAndBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.NonIdentifiability.Failures, []string{StatusABNoOperators, StatusABNoTopChannel, StatusABNoRestSectors}) {
		t.Fatalf("bad non-identifiability: %s", FormatNonID(a.NonIdentifiability))
	}
	if !containsAll(a.FiniteTriple.Supports, []string{StatusFSTSuppliesTraceShape}) || !containsAll(a.FiniteTriple.Failures, []string{StatusFSTNoOperators}) {
		t.Fatalf("bad finite triple: %s", FormatSource(a.FiniteTriple))
	}
	if !containsAll(a.External.Supports, []string{StatusExternalCanPopulate}) || !containsAll(a.External.Failures, []string{StatusExternalNotNative}) {
		t.Fatalf("bad external ledger: %s", FormatSource(a.External))
	}
	if !containsAll(a.TD4.Failures, []string{StatusTD4NoMagnitudes, StatusTD4NoNEff}) {
		t.Fatalf("bad D4 audit: %s", FormatSource(a.TD4))
	}
	if !containsAll(a.K7Projective.Failures, []string{StatusK7NotMagnitudeOperator, StatusProjectiveNotNEff}) {
		t.Fatalf("bad K7/projective audit: %s", FormatSource(a.K7Projective))
	}
	if !strings.Contains(FormatScale(a.Scale), "2 d ln a - d ln b") || !containsAll(a.Scale.Failures, []string{StatusNoScaleStability, StatusMZScaleSealed}) {
		t.Fatalf("bad scale audit: %s", FormatScale(a.Scale))
	}
	if !containsAll(a.CHiggs.Failures, []string{StatusSealAloneNoNativeCHiggs, StatusCHiggsLevelB}) {
		t.Fatalf("bad C_Higgs audit: %s", FormatCHiggs(a.CHiggs))
	}
	if !strings.Contains(a.Branch.Next, "RankThreeTopColorBlock") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
}

func TestGate807Theorem(t *testing.T) {
	res := Generation2TraceMagnitudeOperatorSealAndNEffSourceMinimalityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
