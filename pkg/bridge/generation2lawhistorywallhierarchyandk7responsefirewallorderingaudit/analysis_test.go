package generation2lawhistorywallhierarchyandk7responsefirewallorderingaudit

import (
	"strings"
	"testing"
)

func TestGate749WallHierarchyAndK7Roles(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate748.Inherited || !a.Gate748.ResidualCompression || !a.Gate748.BoundaryStressReappears || !a.Gate748.FlavorFirewallKept {
		t.Fatalf("bad Gate748 inheritance: %+v", a.Gate748)
	}
	if a.Hierarchy.Count < 16 || !containsWall(a.Hierarchy, "Boolean wall") || !containsWall(a.Hierarchy, "K7 second-moment stress wall") || !containsWall(a.Hierarchy, "Tree/pole wall") {
		t.Fatalf("bad wall hierarchy: %+v", a.Hierarchy)
	}
	if !a.K7Roles.BoundaryVectorMapBlocked || !a.K7Roles.FlavorPromotionBlocked || !a.K7Roles.HiggsPromotionBlocked || !strings.Contains(a.K7Roles.Verdict, StatusNoK7ToBoundaryVectorMap) {
		t.Fatalf("bad K7 role separation: %+v", a.K7Roles)
	}
	if a.Firewall.Count != 12 || a.Firewall.Steps[0].Name != "Native law-space firewall" || a.Firewall.Steps[11].Name != "Tree/pole mass firewall" {
		t.Fatalf("bad firewall ordering: %+v", a.Firewall)
	}
}

func TestGate749ResonanceMomentsAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Resonance.UsesOrientation || !a.Resonance.UsesFiveThirds || !a.Resonance.UsesXiBoundary || !a.Resonance.UsesK7Moment || a.Resonance.IsFlavorTheorem {
		t.Fatalf("bad Gate748 resonance: %+v", a.Resonance)
	}
	if a.Moment.K7EventWeight < 0.09 || a.Moment.K7EventWeight > 0.1 || a.Moment.XiBoundary < 0.05 || a.Moment.M2Wall < 1e-7 || !strings.Contains(a.Moment.MomentFormula, "p_K7") {
		t.Fatalf("bad moment coordinate: %+v", a.Moment)
	}
	if !a.Reduction.DoNotChaseResidual || len(a.Reduction.RecommendedTargets) != 4 || !a.Reduction.StabilizedBeforeNext {
		t.Fatalf("bad reduction priority: %+v", a.Reduction)
	}
	if !a.Physical.NoBoundaryVectorMap || !a.Physical.NoFlavorTheorem || !a.Physical.NoHistoryLoopTheorem || !a.Physical.NoScalarRuntimeTheorem || !a.Physical.NoHiggsPoleMassTheorem || !a.Physical.NoYukawaTheorem {
		t.Fatalf("bad physical firewall: %+v", a.Physical)
	}
	res := Generation2LawHistoryWallHierarchyAndK7ResponseFirewallOrderingAuditTheorem().Verify()
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
