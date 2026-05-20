package generation2firsttraceresidualandquadraticspectralcorrectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate690Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.FirstTraceSelectionInherited || !a.Inherited.FirstTraceActive || a.Inherited.QuadraticLeadingActive {
		t.Fatalf("bad Gate689 inheritance: %+v", a.Inherited)
	}
	if a.Discipline.ClaimsNativeSpectralExpansion || a.Discipline.ClaimsKappaECorrectionCertified || a.Discipline.PerformsArbitraryRationalSearch {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestFirstTraceResidualAndQuadraticScale(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Residual.Residual-8.525834398014336e-10) > residualTolerance {
		t.Fatalf("bad E1 residual: %+v", a.Residual)
	}
	if math.Abs(a.Quadratic.F2-1.624013231638281e-7) > 1e-21 {
		t.Fatalf("bad F2: %+v", a.Quadratic)
	}
	if math.Abs(a.Quadratic.ResidualOverF2-0.005249855254820553) > coefficientTolerance {
		t.Fatalf("bad residual/F2 coefficient: %+v", a.Quadratic)
	}
	if !a.Quadratic.ResidualOverF2Small || !a.Quadratic.F2SecondOrder || !a.Quadratic.F2MuchLargerThanResidual || !a.Quadratic.QuadraticStillNotLeading {
		t.Fatalf("bad quadratic scale classification: %+v", a.Quadratic)
	}
}

func TestTypedCoefficientCandidates(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Coefficients.CandidateCount != 7 || len(a.Coefficients.Candidates) != 7 {
		t.Fatalf("expected seven typed coefficient candidates: %+v", a.Coefficients)
	}
	if !a.Coefficients.KappaEClosest || a.Coefficients.BestCandidate != "kappa_e" {
		t.Fatalf("kappa_e should be closest typed candidate: %+v", a.Coefficients)
	}
	if !a.Coefficients.NoArbitraryRationalSearch || !a.Coefficients.AllCandidatesAlreadyTyped {
		t.Fatalf("coefficient audit should avoid arbitrary search and use typed quantities only: %+v", a.Coefficients)
	}
	if !strings.Contains(a.Coefficients.Verdict, StatusKappaEQuadraticNotIndependentlyCertified) {
		t.Fatalf("missing noncertification verdict: %+v", a.Coefficients)
	}
}

func TestFlavorDeficitComparisonIsClueOnly(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FlavorDeficit.KappaECloserThanF2Alone || !a.FlavorDeficit.KappaEOrientCloserThanF2Alone {
		t.Fatalf("flavor coefficients should improve over raw F2 alone: %+v", a.FlavorDeficit)
	}
	if a.FlavorDeficit.KappaEExact || a.FlavorDeficit.KappaEOrientExact || !a.FlavorDeficit.ResidualClueOnly {
		t.Fatalf("flavor comparison must remain clue-only, not exact theorem: %+v", a.FlavorDeficit)
	}
	if math.Abs(a.FlavorDeficit.KappaEAbsResidual-4.1201043014107086e-11) > 1e-20 {
		t.Fatalf("unexpected kappa_e residual: %+v", a.FlavorDeficit)
	}
}

func TestNoncircularityAndExpansionFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Noncircularity.DBaseContainsKappaE || !a.Noncircularity.CorrectionUsesKappaE || !a.Noncircularity.KappaEExplanationPartiallyDependent {
		t.Fatalf("kappa_e dependency was not recorded: %+v", a.Noncircularity)
	}
	if a.Noncircularity.IndependentEvidence || a.Noncircularity.NativeCorrectionCertified || a.Noncircularity.PromoteCorrection {
		t.Fatalf("noncircularity firewall violated: %+v", a.Noncircularity)
	}
	if !a.Expansion.ExactC2IsDefinitionOnly || a.Expansion.KappaEFormulaPromoted || a.Expansion.ExpansionTheoremCertified {
		t.Fatalf("spectral expansion was promoted incorrectly: %+v", a.Expansion)
	}
	if math.Abs(a.Expansion.ResidualWithExactC2) > residualTolerance {
		t.Fatalf("exact c2 should close only definitionally: %+v", a.Expansion)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2FirstTraceResidualAndQuadraticSpectralCorrectionAuditTheorem().Verify()
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
