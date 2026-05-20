package generation2ferminormalizedhiggsratioandscalecancellationaudit

import (
	"strings"
	"testing"
)

func TestGate779InheritanceAndRatio(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate778.Inherited || a.Gate778.TreeTowerFormula != "m_H_tree=(v/2)sqrt(C_Higgs)" || a.Gate778.FermiVEVSeal != "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)" || !closeRel(a.Gate778.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Gate778.VEVGeV, 246.2196508, 1e-15) || !closeRel(a.Gate778.DilationFactor, 1.0184402389953279, 1e-15) || !closeRel(a.Gate778.TreeMassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Gate778.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) || a.Gate778.NativeFermiTheorem || a.Gate778.NativeElectroweakScale || a.Gate778.PoleMassTheorem {
		t.Fatalf("bad Gate778 inheritance: %+v", a.Gate778)
	}
	if !a.Ratio.Defined || a.Ratio.SquaredTreeFormula != "m_H_tree^2=(v^2/4)C_Higgs" || a.Ratio.VEVCancelledFormula != "m_H_tree^2/v^2=C_Higgs/4" || a.Ratio.FermiConvention != "1/v^2=sqrt(2)G_F" || a.Ratio.NormalizedIdentity != "4sqrt(2)G_F m_H_tree^2=C_Higgs" || !a.Ratio.UsesExternalGFSeal || a.Ratio.DerivesGF || a.Ratio.DerivesVEV {
		t.Fatalf("bad ratio definition: %+v", a.Ratio)
	}
}

func TestGate779CancellationAndLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Cancellation.Computed || !closeRel(a.Cancellation.TreeMassSquaredOverVSquared, 0.2593051301012151, 1e-15) || !closeRel(a.Cancellation.CHiggsOverFour, 0.2593051301012151, 1e-15) || !closeRel(a.Cancellation.Sqrt2GFTreeMassSquared, 0.2593051301012151, 1e-15) || !closeRel(a.Cancellation.FourSqrt2GFTreeMassSquared, 1.0372205204048603, 1e-15) || !a.Cancellation.MatchesCHiggs || !a.Cancellation.ScaleCancelledToDimensionless {
		t.Fatalf("bad cancellation: %+v", a.Cancellation)
	}
	if !a.Ledger.Finite || !closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Ledger.VEVGeV, 246.2196508, 1e-15) || !closeRel(a.Ledger.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) || !closeRel(a.Ledger.TreeMassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Ledger.TreeMassSquaredGeV2, 15720.144400764586, 1e-15) || !closeRel(a.Ledger.TreeMassOverVEV, 0.5092201194976639, 1e-15) || !closeRel(a.Ledger.Sqrt2GFTreeMassSquared, 0.2593051301012151, 1e-15) || !closeRel(a.Ledger.FourSqrt2GFTreeMassSquared, 1.0372205204048603, 1e-15) {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
}

func TestGate779TaskSeparationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Tasks.Separated || a.Tasks.DimensionlessTask != "derive or reduce C_Higgs natively" || a.Tasks.ScaleTask != "derive or seal G_F / v" || !a.Tasks.RequiresBothForMass || !a.Tasks.RatioDoesNotDeriveGF || !a.Tasks.RatioDoesNotDeriveVEV {
		t.Fatalf("bad task separation: %+v", a.Tasks)
	}
	if !a.Firewalls.Audited || a.Firewalls.RatioPoleMassTheorem || a.Firewalls.GFAShaNativeInput || a.Firewalls.FermiNormalizedRatioMeasuredPrediction || a.Firewalls.CHiggsNativeHiggsTheorem || a.Firewalls.TreeProxyPoleMass || a.Firewalls.DimensionlessRatioElectroweakScale || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate779TheoremStatuses(t *testing.T) {
	res := Generation2FermiNormalizedHiggsRatioAndScaleCancellationAuditTheorem().Verify()
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
