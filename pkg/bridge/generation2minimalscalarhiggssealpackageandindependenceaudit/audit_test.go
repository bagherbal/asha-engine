package generation2minimalscalarhiggssealpackageandindependenceaudit

import (
	"strings"
	"testing"
)

func TestGate738RolesIndependenceAndMinimality(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate737.Inherited || !a.Gate737.PRadTypeDistinctSeal || !a.Gate737.NoNativePRadSelector || !a.Gate737.HistoryLoopConditional {
		t.Fatalf("bad Gate737 inheritance: %+v", a.Gate737)
	}
	if len(a.Roles.Roles) != 3 || !a.Roles.NSelectsJH || !a.Roles.NDefinesHopfPhase || !a.Roles.QNormalizesU1 || !a.Roles.PRadSelectsRadial || !a.Roles.PRadEnablesSplits {
		t.Fatalf("bad roles: %+v", a.Roles)
	}
	if len(a.Independence.Substitutions) != 8 || !a.Independence.NQTypeDistinct || !a.Independence.NPRadTypeDistinct || !a.Independence.QPRadTypeDistinct || a.Independence.RhoPlusDeterminesAny || a.Independence.PK7DeterminesPRad {
		t.Fatalf("bad independence: %+v", a.Independence)
	}
	for _, s := range a.Independence.Substitutions {
		if s.Allowed {
			t.Fatalf("forbidden substitution unexpectedly allowed: %+v", s)
		}
	}
	if !a.Minimality.AllThreeRequired || len(a.Minimality.RemoveNConsequences) != 3 || len(a.Minimality.RemoveQConsequences) != 2 || len(a.Minimality.RemovePRadConsequences) != 3 {
		t.Fatalf("bad minimality: %+v", a.Minimality)
	}
}

func TestGate738AvailableRemainingFirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Available.Structures) != 6 || !a.Available.NQPRadSupplied || !a.Available.HistoryLoopAvailable || !a.Available.RuntimeBridgeCompatible {
		t.Fatalf("bad available structures: %+v", a.Available)
	}
	if len(a.Remaining.Dependencies) != 6 || !a.Remaining.AllStillBridgeOrSealed {
		t.Fatalf("bad remaining dependencies: %+v", a.Remaining)
	}
	if a.Firewall.PackageIsPhysicalHiggsTheorem || a.Firewall.PRadIsElectroweakVacuumTheorem || a.Firewall.NIsNativeComplexStructureTheorem || a.Firewall.QIsNativeHyperchargeDerivation || a.Firewall.LIsNativeHistoryLoopTheorem || a.Firewall.RuntimeBridgeIsHiggsMassPrediction || a.Firewall.FWall3IsNativeBoundaryResponseTheorem {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}
	res := Generation2MinimalScalarHiggsSealPackageAndIndependenceAuditTheorem().Verify()
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
