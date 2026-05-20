package generation2oneeighthscalarbaselineandmultiplicativecorrectionfactorizationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate758InheritanceAndFactors(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate757.Inherited || !a.Gate757.EffectiveParticipationAudited || a.Gate757.IndependentScalarRuntimeTheorem {
		t.Fatalf("bad Gate757 inheritance: %+v", a.Gate757)
	}
	if math.Abs(a.Gate757.NEff-nEffMZ) > 1e-15 || math.Abs(a.Gate757.TraceRatio-bOverA2MZ) > 1e-15 || math.Abs(a.Gate757.LambdaProxy-lambdaProxyMZ) > 1e-15 || math.Abs(a.Gate757.RuntimeTransportBracket-cHistoryMZ) > 1e-12 {
		t.Fatalf("bad inherited numerics: %+v", a.Gate757)
	}
	if !a.Factors.FactorsDefined || !a.Factors.CYukawaBelowOne || !a.Factors.CHistoryAboveOne {
		t.Fatalf("bad factor definitions: %+v", a.Factors)
	}
	if math.Abs(a.Factors.CYukawa-0.9992248188812008) > 1e-15 || math.Abs(a.Factors.CYukawaFromTrace-a.Factors.CYukawa) > 1e-15 || math.Abs(a.Factors.CHistory-1.038025177923625) > 1e-12 {
		t.Fatalf("bad factor numerics: %+v", a.Factors)
	}
}

func TestGate758OneEighthFactorizationAndTreeProxy(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Factorization.Computed || a.Factorization.IndependentRuntimeTheorem {
		t.Fatalf("bad factorization typing: %+v", a.Factorization)
	}
	if math.Abs(a.Factorization.Baseline-oneEighth) > 1e-15 || math.Abs(a.Factorization.TotalCorrection-1.0372205204048603) > 1e-15 || math.Abs(a.Factorization.LambdaRuntimeEff-0.12965256505060754) > 1e-15 || math.Abs(a.Factorization.FactorizationResidual) > 1e-18 {
		t.Fatalf("bad factorization numerics: %+v", a.Factorization)
	}
	if !a.TreeProxy.Computed || a.TreeProxy.PoleMassPrediction {
		t.Fatalf("bad tree proxy typing: %+v", a.TreeProxy)
	}
	if math.Abs(a.TreeProxy.BaselineVOverTwoGeV-123.1098254) > 1e-10 || math.Abs(a.TreeProxy.SqrtTotalCorrection-1.0184402389953278) > 1e-15 || math.Abs(a.TreeProxy.TreeProxyGeV-125.38000000304908) > 1e-9 || math.Abs(a.TreeProxy.TreeProxyResidualGeV) > 1e-12 {
		t.Fatalf("bad tree proxy numerics: %+v", a.TreeProxy)
	}
}

func TestGate758SourceRolesAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Interpretation.Recorded || !a.Interpretation.CYukawaLowersProxy || !a.Interpretation.CHistoryLiftsRuntime || a.Interpretation.OneEighthPotentialLaw {
		t.Fatalf("bad source interpretation: %+v", a.Interpretation)
	}
	if !a.Roles.LayerSeparationAudited || !a.Roles.FactorsMultiplyAfterScalarCollapse || a.Roles.OperatorsOnSameNativeSpace || a.Roles.CYukawaNativeYukawaTheorem || a.Roles.CHistoryNativeHistoryLoopTheorem {
		t.Fatalf("bad role audit: %+v", a.Roles)
	}
	if a.Firewalls.CYukawaNativeYukawaTheorem || a.Firewalls.CHistoryNativeHistoryLoopTheorem || a.Firewalls.ProductIndependentScalarRuntimeTheorem || a.Firewalls.TreeProxyPoleMassPrediction || a.Firewalls.OneEighthScalarPotentialTheorem || a.Firewalls.ClaimsYukawaEigenvaluesDerived || a.Firewalls.ClaimsFlavorHierarchyDerived || a.Firewalls.ClaimsCKMPMNSDerived || a.Firewalls.ClaimsHiggsMassTheorem || a.Firewalls.ClaimsPoleMassTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate758TheoremVerdictStatuses(t *testing.T) {
	res := Generation2OneEighthScalarBaselineAndMultiplicativeCorrectionFactorizationAuditTheorem().Verify()
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
			t.Fatalf("missing status note %s", want)
		}
	}
}
