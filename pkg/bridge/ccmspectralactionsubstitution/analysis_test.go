package ccmspectralactionsubstitution

import (
	"math"
	"testing"
)

func TestBuildDefaultCCMSubstitution(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.Executed || !c.StructuralClosure || c.FullNumericalTOEClosure {
		t.Fatalf("bad closure flags: %+v", c)
	}
	if math.Abs(c.Einstein.CoefficientWithPreviousF2NoC-1.0/(16.0*math.Pi)) > 1e-15 {
		t.Fatalf("bad previous-f2 coefficient: %.18g", c.Einstein.CoefficientWithPreviousF2NoC)
	}
	if math.Abs(c.Einstein.GapFactor-8.0*math.Pi) > 1e-12 {
		t.Fatalf("bad gap factor: %.18g", c.Einstein.GapFactor)
	}
	if math.Abs(c.Einstein.RequiredF2LambdaMP2Leading-math.Pi*math.Pi/8.0) > 1e-15 {
		t.Fatalf("bad required F2: %.18g", c.Einstein.RequiredF2LambdaMP2Leading)
	}
}

func TestHiggsReadoffConventions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	h := a.Calculation.Higgs
	wantNoPi := (1197.0 / 4624.0) / (2.0 * 7.0)
	wantPi := math.Pi * math.Pi * wantNoPi
	if math.Abs(h.QuarticNoOuterPiConvention-wantNoPi) > 1e-15 {
		t.Fatalf("bad no-pi quartic: %.18g", h.QuarticNoOuterPiConvention)
	}
	if math.Abs(h.QuarticCanonicalOuterPiConvention-wantPi) > 1e-15 {
		t.Fatalf("bad pi quartic: %.18g", h.QuarticCanonicalOuterPiConvention)
	}
	if h.QuarticNoOuterPiConvention == h.PreviousAshaRatio || h.QuarticCanonicalOuterPiConvention == h.PreviousAshaRatio {
		t.Fatalf("expected CCM quartic to differ from raw finite ratio")
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["ccm_gap_factor"]-8.0*math.Pi) > 1e-12 {
		t.Fatalf("bad constants: %+v", m)
	}
	if math.Abs(m["ccm_required_f2_lambda_over_mp2_leading"]-math.Pi*math.Pi/8.0) > 1e-15 {
		t.Fatalf("bad required f2: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := CCMSpectralActionDirectSubstitutionCompleteCoefficientLedgerTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
