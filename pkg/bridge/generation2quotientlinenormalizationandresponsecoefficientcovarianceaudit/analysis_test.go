package generation2quotientlinenormalizationandresponsecoefficientcovarianceaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate700InheritanceAndRescalingCovariance(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	inh := a.Inherited
	if !inh.InheritedConditionalLaw || !nearly(inh.SigmaBoundary, canonicalSBoundary()) || !nearly(inh.SigmaHistory, canonicalSHistory()) || !nearly(inh.Expectation, canonicalExpectation()) || math.Abs(inh.ResidualE1-canonicalResidual()) > strictTol || !inh.PremisesNonredundant || !inh.NoNativePrinciple || !inh.NoNativeSevenOver72 || inh.Verdict != StatusGate700ConditionalHistoryResponseLawInherited {
		t.Fatalf("bad inheritance: %+v", inh)
	}
	r := a.Rescaling
	if r.BoundaryScale != 2 || r.HistoryScale != 3 || !nearly(r.OriginalCoefficient, eventProbK7) || !nearly(r.TransformedCoefficient, (3.0/2.0)*eventProbK7) || !strings.Contains(r.Formula, "beta/alpha") || !r.CoefficientCovariant || r.CoefficientInvariant {
		t.Fatalf("bad rescaling: %+v", r)
	}
	for _, want := range []string{StatusQuotientLineRescalingDefined, StatusResponseCoefficientTransformationComputed, StatusResponseCoeffNotInvariantUnderArbitraryQuotientRescaling} {
		if !strings.Contains(r.Verdict, want) {
			t.Fatalf("missing %s in %q", want, r.Verdict)
		}
	}
	unit := BuildRescaling(5, 5)
	if !unit.CoefficientInvariant || !nearly(unit.TransformedCoefficient, eventProbK7) {
		t.Fatalf("equal quotient rescaling should preserve numeric coefficient: %+v", unit)
	}
}

func TestProbabilityInvariantAndWallNormalization(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Probability
	if p.ProbabilityName != "p_K7" || !nearly(p.Probability, 7.0/72.0) || !strings.Contains(p.TraceFormula, "rho_72") || !strings.Contains(p.TraceFormula, "P_K7") || !p.IndependentOfAlpha || !p.IndependentOfBeta || !p.InvariantObjectSeparated || !p.ResponseCoefficientNeedsCoordinates {
		t.Fatalf("bad probability audit: %+v", p)
	}
	if !strings.Contains(p.Verdict, StatusEventProbabilityInvariantSeparated) || !strings.Contains(p.Verdict, StatusSevenOver72InvariantAsK7EventProbability) {
		t.Fatalf("bad probability verdict: %q", p.Verdict)
	}
	w := a.WallNormalize
	if !strings.Contains(w.BoundaryCoordinate, "lambda+(R_3-1)") || !strings.Contains(w.HistoryCoordinate, "kappa_lambda+kappa_e+lambda") || !w.BoundaryUsesUnitWallCoefficients || !w.HistoryUsesUnitWallCoefficients || !w.SameWallDistanceFamily || w.CanonicalAlpha != 1 || w.CanonicalBeta != 1 || !w.CoefficientEqualsProbability || !w.EqualScaleRequiredForDirectEquality {
		t.Fatalf("bad wall normalization: %+v", w)
	}
	for _, want := range []string{StatusWallCoordinateNormalizationAudited, StatusResponseCoefficientEqualsEventProbabilityOnlyCanonicalWalls, StatusGate700LawCoordinateSealedNotCoordinateFree} {
		if !strings.Contains(w.Verdict, want) {
			t.Fatalf("missing %s in %q", want, w.Verdict)
		}
	}
}

func TestAlternativeNormalizationsAndSourceSeparation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	alts := a.Alternatives
	if len(alts.Examples) != 3 || !alts.AllComputed || !alts.NonInvariantSeen || !alts.CanonicalSeen || alts.Verdict != StatusAlternativeNormalizationExamplesComputed {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	expected := []float64{7.0 / 144.0, 7.0 / 36.0, 7.0 / 72.0}
	for i, ex := range alts.Examples {
		if !nearly(ex.TransformedCoefficient, expected[i]) || !nearly(ex.ExpectedCoefficient, expected[i]) || !ex.MatchesExpected {
			t.Fatalf("bad alternative %d: %+v", i, ex)
		}
	}
	s := a.Source
	if !strings.Contains(s.InvariantObject, "p_K7") || !strings.Contains(s.CoordinateSealedObject, "wall-distance") || !s.DoesNotWeakenGate700 || !s.ClarifiesSourceType {
		t.Fatalf("bad source separation: %+v", s)
	}
	if !strings.Contains(s.Verdict, StatusEventProbabilityInvariantSeparated) || !strings.Contains(s.Verdict, StatusGate700LawCoordinateSealedNotCoordinateFree) {
		t.Fatalf("bad source verdict: %q", s.Verdict)
	}
}

func TestMissingTheoremFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	m := a.Missing
	if len(m.Theorems) != 6 || !strings.Contains(m.PrecisePair, "p_K7=7/72") || !strings.Contains(m.PrecisePair, "aligned unit wall-distance") {
		t.Fatalf("bad missing theorem: %+v", m)
	}
	for _, want := range []string{StatusNoNativeWallCoordinateNormalizationAlignmentTheorem, StatusNoNativeBoundaryHistoryResponsePrinciple, StatusNoNativeSevenOver72Theorem} {
		if !strings.Contains(m.Verdict, want) {
			t.Fatalf("missing %s in %q", want, m.Verdict)
		}
	}
	f := a.Firewall
	if f.ClaimsResponseCoefficientCoordinateFree || f.ClaimsNativeWallNormalization || f.ClaimsNativeBoundaryHistoryPrinciple || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate701QuotientNormalizationBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	res := Generation2QuotientLineNormalizationAndResponseCoefficientCovarianceAuditTheorem().Verify()
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
