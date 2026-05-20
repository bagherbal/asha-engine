package generation2electroweakscalesourcecandidatesandfermivevairlockaudit

import (
	"strings"
	"testing"
)

func TestGate778InheritanceAndFermiLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate777.Inherited || a.Gate777.TreeTowerFormula != "m_H_tree=(v/2)sqrt(C_Higgs)" || !closeRel(a.Gate777.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Gate777.DilationFactor, 1.0184402389953279, 1e-15) || !closeRel(a.Gate777.VEVScaleGeV, 246.2196508, 1e-15) || !closeRel(a.Gate777.TreeMassGeV, 125.38000000304908, 1e-15) || !a.Gate777.DimensionlessCorrectionExists || a.Gate777.DimensionfulScaleSeal != "v" || a.Gate777.NativeVEVTheorem || a.Gate777.PoleMassTheorem {
		t.Fatalf("bad Gate777 inheritance: %+v", a.Gate777)
	}
	if !a.FermiLane.Audited || a.FermiLane.SealName != "FermiVEVScaleSeal" || a.FermiLane.Formula != "v=(sqrt(2)G_F)^(-1/2)" || a.FermiLane.Input != "G_F" || a.FermiLane.Output != "v" || !closeRel(a.FermiLane.VEVGeV, 246.2196508, 1e-15) || !closeRel(a.FermiLane.EquivalentGFGeVMinus2, 1.1663786999444556e-05, 1e-15) || a.FermiLane.NativeFermiTheorem || a.FermiLane.NativeVEVTheorem || !a.FermiLane.LawfulExternalAirlock {
		t.Fatalf("bad Fermi lane: %+v", a.FermiLane)
	}
}

func TestGate778CandidateLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.WLane.Audited || a.WLane.Formula != "v=2m_W/g" || !a.WLane.RequiresAbsoluteWeakG || !a.WLane.RequiresWMass || !a.WLane.GaugeRatiosOrganized || a.WLane.AbsoluteWeakScaleNative || a.WLane.WMassNative || !a.WLane.LaneSealed {
		t.Fatalf("bad W lane: %+v", a.WLane)
	}
	if !a.PotentialLane.Audited || a.PotentialLane.Formula != "v^2=-mu^2/lambda_H" || !a.PotentialLane.LambdaAirlockExists || a.PotentialLane.MuSquaredIndependentlySourced || !a.PotentialLane.CircularWithoutMuSource || a.PotentialLane.NativeMuSquaredSource || a.PotentialLane.DeterminesVEV {
		t.Fatalf("bad potential lane: %+v", a.PotentialLane)
	}
	if !a.SpectralLane.Audited || a.SpectralLane.Candidate != "dimensionful spectral-action scale or cutoff" || !a.SpectralLane.DimensionfulScaleCouldSetUnits || !a.SpectralLane.CurrentBridgeDimensionlessOnly || a.SpectralLane.MapsSpectralScaleToVEV || !a.SpectralLane.LaneCandidateOnly {
		t.Fatalf("bad spectral lane: %+v", a.SpectralLane)
	}
	if !a.BoundaryLane.Audited || !a.BoundaryLane.BoundaryScaleSealExists || !a.BoundaryLane.ScalarWallDataExists || a.BoundaryLane.DeterminesElectroweakVEV || a.BoundaryLane.BoundaryScaleEqualsVEVTheorem {
		t.Fatalf("bad boundary lane: %+v", a.BoundaryLane)
	}
}

func TestGate778RankingLedgerAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ranking.Recorded || a.Ranking.BestCurrentLawfulSource != "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)" || !containsAll(a.Ranking.BestFutureNativeTargets, []string{"mu^2 source theorem", "absolute electroweak scale theorem"}) || !containsAll(a.Ranking.BlockedShortcuts, []string{"C_Higgs does not determine v", "lambda_runtime_eff does not determine v without mu^2", "P_rad does not determine v", "HistoryLoopUnit does not determine v", "7/72 does not determine v", "1/(8pi) does not determine v"}) {
		t.Fatalf("bad ranking: %+v", a.Ranking)
	}
	if !a.Ledger.Finite || !closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Ledger.DilationFactor, 1.0184402389953279, 1e-15) || !closeRel(a.Ledger.VEVGeV, 246.2196508, 1e-15) || !closeRel(a.Ledger.EquivalentGF, 1.1663786999444556e-05, 1e-15) || !closeRel(a.Ledger.VHalfGeV, 123.1098254, 1e-15) || !closeRel(a.Ledger.TreeMassGeV, 125.38000000304908, 1e-15) {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
	if !a.Firewalls.Audited || a.Firewalls.FermiScaleNativeTheorem || a.Firewalls.VEVDerivedFromCHiggs || a.Firewalls.VEVDerivedFromLambdaRuntimeOnly || a.Firewalls.MuSquaredBridgeNativeSource || a.Firewalls.WRelationNativeWithoutInputs || a.Firewalls.TreeProxyPoleMass || a.Firewalls.DimensionlessTowerMassScaleTheorem || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate778TheoremStatuses(t *testing.T) {
	res := Generation2ElectroweakScaleSourceCandidatesAndFermiVEVAirlockAuditTheorem().Verify()
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
