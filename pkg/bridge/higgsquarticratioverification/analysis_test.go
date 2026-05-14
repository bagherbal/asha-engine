package higgsquarticratioverification

import "testing"

func TestRatioInheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.RatioInherited || a.Input.RatioNumerator != 1197 || a.Input.RatioDenominator != 4624 || a.Input.UsesGStarSquaredOne {
		t.Fatalf("bad ratio inheritance: %s", FormatRatioInheritance(a.Input))
	}
	if a.Input.ExactRatio <= 0.25 || a.Input.ExactRatio >= 0.26 {
		t.Fatalf("unexpected exact ratio: %s", FormatRatioInheritance(a.Input))
	}
}

func TestEmpiricalLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Ledger.QuarantinedInput || a.Ledger.DerivedFromFiniteCore || !a.Ledger.ReplacesDiagnosticSeal {
		t.Fatalf("ledger firewall failure: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.AlphaGUT <= 0.039 || a.Ledger.AlphaGUT >= 0.041 || a.Ledger.GStarSquared <= 0.50 || a.Ledger.GStarSquared >= 0.51 {
		t.Fatalf("bad empirical coupling values: %s", FormatLedger(a.Ledger))
	}
}

func TestQuarticPrediction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Prediction.PredictedLambda <= 0.129 || a.Prediction.PredictedLambda >= 0.132 {
		t.Fatalf("lambda not near Higgs proxy: %s", FormatPrediction(a.Prediction))
	}
	if a.Prediction.PredictedMassGeV <= 125.0 || a.Prediction.PredictedMassGeV >= 126.5 || a.Prediction.OldSealTreeMassGeV <= 175 || !a.Prediction.OldSealRejected {
		t.Fatalf("bad mass proxy / old seal handling: %s", FormatPrediction(a.Prediction))
	}
}

func TestEmpiricalProxyComparison(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Comparison.WithinRatioTolerance || !a.Comparison.WithinTreeMassTolerance || a.Comparison.RatioPercentError >= comparisonTolerancePct || a.Comparison.MassPercentError >= comparisonTolerancePct {
		t.Fatalf("comparison outside tolerance: %s", FormatComparison(a.Comparison))
	}
	if !a.Comparison.ComparisonIsTreeProxyOnly || a.Comparison.FullGUTRGERunExecuted || a.Comparison.PoleMassMatched {
		t.Fatalf("comparison overclaimed precision physics: %s", FormatComparison(a.Comparison))
	}
}

func TestBoundaryRatioCatalog(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Catalog.BothAreRatios || !a.Catalog.NoAbsoluteCouplingClaim || a.Catalog.AlgebraicRatioCount != 2 || !a.Catalog.SecondRatioCataloged {
		t.Fatalf("bad ratio catalog: %s", FormatCatalog(a.Catalog))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoAlphaGUTDerivationClaimed || !a.Firewalls.NoFullRGTransportClaimed || !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoThresholdMatchingClaimed || !a.Firewalls.NoObservedMassUsedAsDerivation || !a.Firewalls.NoGStarOnePhysicalClaim || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := EmpiricalHiggsQuarticRatioVerificationTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
