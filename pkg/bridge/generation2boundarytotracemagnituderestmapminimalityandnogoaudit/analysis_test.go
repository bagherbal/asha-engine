package generation2boundarytotracemagnituderestmapminimalityandnogoaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate814InheritanceAndMapDefinition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate813Inherited || !a.Inheritance.SecondMomentSelected || !a.Inheritance.PositiveCompatibilityInherited {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if math.Abs(a.Inheritance.ResidualImprovement-34.2924) > 1e-3 {
		t.Fatalf("bad residual improvement: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Map.Defined || len(a.Map.Objects) < 12 || !strings.Contains(a.Map.TargetChain, "positive rest atoms") {
		t.Fatalf("bad map definition: %s", FormatMap(a.Map))
	}
	if !containsAll(a.Map.Supports, []string{StatusExactMissingObject}) {
		t.Fatalf("missing map support: %+v", a.Map.Supports)
	}
}

func TestGate814MinimalityAndSourceFailures(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	for key, want := range map[string]string{
		"s":                    StatusNoScaleWithoutS,
		"p":                    StatusNoM2WithoutP,
		"5/3":                  StatusNoNineFiveWithoutHyper,
		"color 3":              StatusNoTopBaselineWithoutColor,
		"boundary pair 2":      StatusNoSixWithoutBoundaryTwo,
		"top selector":         StatusNoAlphaBetaWithoutTop,
		"alpha beta q maps":    StatusDirectClosureNoAtoms,
		"positive spectrum":    StatusNoYukawaWithoutPositive,
		"scale scheme":         StatusNoScaleLocalWithoutScheme,
		"noncircularity proof": StatusNoPredictionNoNonCirc,
	} {
		if got := a.Minimality.RemovalFailures[key]; got != want {
			t.Fatalf("bad minimality failure for %s: got %q want %q", key, got, want)
		}
	}
	if !containsAll(a.Sources.Failures, []string{StatusBoundaryPairNoMap, StatusK7NoYukawaAtoms, StatusFiniteTripleNoMap, StatusExternalNotNative, StatusD4NotMap, StatusChiralityNotSource}) {
		t.Fatalf("missing source failures: %s", FormatSources(a.Sources))
	}
}

func TestGate814LevelsImpactAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Levels.CurrentStatus, "R1") || !strings.Contains(a.Levels.CurrentStatus, "partial R2") {
		t.Fatalf("bad levels: %s", FormatLevels(a.Levels))
	}
	if math.Abs(a.Impact.OfficialCYukawa-CYukawa) > 1e-15 || math.Abs(a.Impact.OfficialCHiggs-CHiggs) > 1e-15 {
		t.Fatalf("official ledger changed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate814 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2BoundaryToTraceMagnitudeRestMapMinimalityAndNoGoAuditTheorem().Verify()
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
			t.Fatalf("missing status %s", want)
		}
	}
}
