package generation2higgssocketmissingsealindependenceandsourcecandidateaudit

import (
	"strings"
	"testing"
)

func TestGate720MissingSealSourcesAndIndependence(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate719.Inherited || !a.Gate719.SocketAssembled || !a.Gate719.RequiresN || !a.Gate719.RequiresQ || !a.Gate719.RepresentationCompatible || a.Gate719.NativeTwistorSelector || a.Gate719.NativeHyperchargeNorm || a.Gate719.CanonicalThetaH || a.Gate719.PhysicalHiggsDoubletTheorem || a.Gate719.HiggsMassOrRuntime || a.Gate719.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad Gate719 inheritance: %+v", a.Gate719)
	}
	if !a.NSelector.LivesInS2K7Minus || !a.NSelector.SelectsJH || !a.NSelector.SelectsPhaseLine || !a.NSelector.SelectsComplexCarrier || len(a.NSelector.Candidates) != 7 || a.NSelector.NativeSelectorFound || !a.NSelector.RequiresSelectorSeal {
		t.Fatalf("bad n selector source audit: %+v", a.NSelector)
	}
	if !a.QNorm.LivesInRNonzero || !a.QNorm.NormalizesPhaseGenerator || len(a.QNorm.Candidates) != 4 || !a.QNorm.CanMatchTargetConvention || a.QNorm.NativeQDerived || !a.QNorm.RequiresNormalizationSeal {
		t.Fatalf("bad q normalization audit: %+v", a.QNorm)
	}
	if !a.Types.TypeDistinct || !a.Types.ChangingNChangesLine || !a.Types.ChangingQRescalesLine || a.Types.NCanDetermineQ || a.Types.QCanDetermineN || !a.Types.IndependentAtLevel {
		t.Fatalf("bad type distinction audit: %+v", a.Types)
	}
}

func TestGate720ForbiddenShortcutsSealsAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Shortcuts.Shortcuts) != 5 || a.Shortcuts.ScalarQuantitiesSelectN || a.Shortcuts.EventProbabilityFixesQ || !a.Shortcuts.AllShortcutsRejected {
		t.Fatalf("bad forbidden shortcut audit: %+v", a.Shortcuts)
	}
	if len(a.Seals.Seals) != 2 || !strings.Contains(a.Seals.TwistorSelectorSeal, "TwistorSelectorSeal") || !strings.Contains(a.Seals.HyperchargeNormalizationSeal, "HyperchargeNormalizationSeal") || !a.Seals.ConditionalSocketRemains || a.Seals.DerivedNative {
		t.Fatalf("bad seal classification: %+v", a.Seals)
	}
	if a.Physical.ConditionalSocketPhysicalHiggsTheorem || a.Physical.MatchedQDerivedHypercharge || a.Physical.ChosenNDerivedVacuumOrientation || a.Physical.K7MinusSelectorFlavorHierarchy || a.Physical.K7PlusPhysicalHiggsMassTheorem || a.Physical.ScalarPotential || a.Physical.QuarticRuntimeLambda || a.Physical.HiggsPoleMass || a.Physical.YukawaOperators || a.Physical.FlavorHierarchy || a.Physical.CKMPMNS {
		t.Fatalf("physical firewall failed: %+v", a.Physical)
	}
	res := Generation2HiggsSocketMissingSealIndependenceAndSourceCandidateAuditTheorem().Verify()
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
