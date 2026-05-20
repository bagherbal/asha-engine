package generation2k7twistorspherehiggssocketbundleandvacuumselectorfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate713TwistorSphereSocketBundleAndSO3Action(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SelectorFirewallInherited || a.Inherited.K7PlusDimension != 4 || a.Inherited.K7MinusDimension != 3 || !a.Inherited.UnitDirectionSelectsJH || !a.Inherited.FamilyValuedU2Socket || a.Inherited.NativeSelectorFound || a.Inherited.CanonicalJHSelected {
		t.Fatalf("bad Gate712 inheritance: %+v", a.Inherited)
	}
	if a.Twistor.SphereDimension != 2 || a.Twistor.ComplexProjectiveDimension != 1 || !a.Twistor.FamilyNativeObject || a.Twistor.SinglePointSelected || !strings.Contains(a.Twistor.EquivalentDescription, "CP1") || !strings.Contains(a.Twistor.Verdict, StatusK7PlusHiggsSocketTwistorSphereFamily) {
		t.Fatalf("bad twistor sphere audit: %+v", a.Twistor)
	}
	if a.SocketBundle.FiberDimension != 4 || a.SocketBundle.CommutantDimension != 3 || a.SocketBundle.SpanJHDimension != 1 || a.SocketBundle.SingleSocketPromoted || !a.SocketBundle.FamilyValuedSocketBundle || !strings.Contains(a.SocketBundle.Verdict, StatusU2SocketBundleOverS2) {
		t.Fatalf("bad socket bundle audit: %+v", a.SocketBundle)
	}
	if !a.SO3Action.ActsTransitivelyOnS2 || !a.SO3Action.PreservesFanoData || a.SO3Action.PreferredPoint || a.SO3Action.CanonicalAxisSelected || !strings.Contains(a.SO3Action.Verdict, StatusNoNativeTwistorPointSelector) {
		t.Fatalf("bad SO3 action audit: %+v", a.SO3Action)
	}
}

func TestGate713DataSeparationVacuumAndPhysicalFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.DataSeparation.DependentCount != 3 || a.DataSeparation.InvariantCount != 5 || !a.DataSeparation.SeparationValid || !strings.Contains(a.DataSeparation.Verdict, StatusSelectorDependentAndInvariantDataSeparated) {
		t.Fatalf("bad data separation audit: %+v", a.DataSeparation)
	}
	if !strings.Contains(a.VacuumFirewall.MissingPoint, "n_*") || a.VacuumFirewall.NativeSelectorCertified || !a.VacuumFirewall.EnvironmentalSealAllowed || len(a.VacuumFirewall.SealNames) != 3 || !strings.Contains(a.VacuumFirewall.Verdict, StatusSingleHiggsSocketRequiresSelectorOrSeal) {
		t.Fatalf("bad vacuum selector firewall: %+v", a.VacuumFirewall)
	}
	f := a.PhysicalFirewall
	if f.TwistorBundlePhysicalElectroweak || f.ChosenJHPhysicalHiggsStructure || f.SpanJHHypercharge || f.CommutantPhysicalSU2L || f.K7MinusFlavorHierarchy || len(f.MissingMaps) != 4 || !strings.Contains(f.Verdict, StatusInternalSocketBundleNotPhysicalElectroweak) || !strings.Contains(f.Verdict, StatusNoHyperchargeAssignmentOrNormalization) {
		t.Fatalf("physical firewall violated: %+v", f)
	}
	res := Generation2K7TwistorSphereHiggsSocketBundleAndVacuumSelectorFirewallAuditTheorem().Verify()
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
