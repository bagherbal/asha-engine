package normalizationfactoraudit

import (
	"math"
	"testing"
)

func TestBuildDefaultNormalizationAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	au := a.Audit
	if !au.Executed || len(au.Factors) != 6 || !au.ChannelSeparated {
		t.Fatalf("bad audit: %+v", au)
	}
	if au.ClosesEH || au.FullNumericalClosure {
		t.Fatalf("firewall breach: %+v", au)
	}
	if math.Abs(au.Factors[0].Numeric-1/(16*math.Pi*math.Pi)) > 1e-15 {
		t.Fatalf("bad heat factor")
	}
	if math.Abs(au.Factors[1].Numeric-1.0/12.0) > 1e-15 {
		t.Fatalf("bad Dirac a2 factor %.18g", au.Factors[1].Numeric)
	}
	wantDoubled := 96.0 * (1 / (16 * math.Pi * math.Pi)) * (1.0 / 12.0) * (math.Pi / 64.0)
	if math.Abs(au.EHWithDoubledTrace.CoefficientPerMP2-wantDoubled) > 1e-15 {
		t.Fatalf("bad doubled EH %.18g want %.18g", au.EHWithDoubledTrace.CoefficientPerMP2, wantDoubled)
	}
	if math.Abs(au.EHWithDoubledTrace.RequiredF2LambdaOverMP2-math.Pi*math.Pi) > 1e-12 {
		t.Fatalf("bad required f2 doubled %.18g", au.EHWithDoubledTrace.RequiredF2LambdaOverMP2)
	}
	if math.Abs(au.EHWithDoubledTrace.CurrentF2ShortBy-64*math.Pi) > 1e-12 {
		t.Fatalf("bad short factor doubled %.18g", au.EHWithDoubledTrace.CurrentF2ShortBy)
	}
	if math.Abs(au.EHWithRealityHalfTrace.CurrentF2ShortBy-128*math.Pi) > 1e-12 {
		t.Fatalf("bad short factor half %.18g", au.EHWithRealityHalfTrace.CurrentF2ShortBy)
	}
}

func TestConstants(t *testing.T) {
	c := Constants()
	if math.Abs(c["required_f2_doubled"]-math.Pi*math.Pi) > 1e-12 {
		t.Fatalf("bad required_f2_doubled: %.18g", c["required_f2_doubled"])
	}
	if math.Abs(c["required_f2_half"]-2*math.Pi*math.Pi) > 1e-12 {
		t.Fatalf("bad required_f2_half: %.18g", c["required_f2_half"])
	}
	if !(c["gauge_required_K_for_alpha_8pi_if_f0_7"] > 45 && c["gauge_required_K_for_alpha_8pi_if_f0_7"] < 46) {
		t.Fatalf("bad K: %.18g", c["gauge_required_K_for_alpha_8pi_if_f0_7"])
	}
}

func TestTheoremPasses(t *testing.T) {
	res := CompleteNormalizationFactorAuditProductSpectralActionConventionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
