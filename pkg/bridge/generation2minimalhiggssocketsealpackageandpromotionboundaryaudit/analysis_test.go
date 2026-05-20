package generation2minimalhiggssocketsealpackageandpromotionboundaryaudit

import (
	"strings"
	"testing"
)

func TestGate721SealPackageAssemblyAndAvailability(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate720.Inherited || !a.Gate720.NAndQTypeDistinct || !a.Gate720.NRequiresSelectorSeal || !a.Gate720.QRequiresNormalizationSeal || !a.Gate720.ConditionalSocketReadyButNotNative || !a.Gate720.NoNativeN || !a.Gate720.NoNativeQ || !a.Gate720.NoPhysicalHiggs || !a.Gate720.NoHiggsRuntime || !a.Gate720.NoYukawa {
		t.Fatalf("bad Gate720 inheritance: %+v", a.Gate720)
	}
	if a.Package.SealCount != 2 || !a.Package.SuppliesN || !a.Package.SuppliesQ || !a.Package.Minimal || a.Package.Native {
		t.Fatalf("bad minimal seal package: %+v", a.Package)
	}
	if a.Package.TwistorSelectorSeal.Name != "TwistorSelectorSeal" || !a.Package.TwistorSelectorSeal.Required || a.Package.TwistorSelectorSeal.Native || !strings.Contains(a.Package.TwistorSelectorSeal.Output, "K7+_J") {
		t.Fatalf("bad twistor seal role: %+v", a.Package.TwistorSelectorSeal)
	}
	if a.Package.HyperchargeNormalizationSeal.Name != "HyperchargeNormalizationSeal" || !a.Package.HyperchargeNormalizationSeal.Required || a.Package.HyperchargeNormalizationSeal.Native || !strings.Contains(a.Package.HyperchargeNormalizationSeal.Output, "qJ_H") {
		t.Fatalf("bad hypercharge seal role: %+v", a.Package.HyperchargeNormalizationSeal)
	}
	if !a.Assembly.SelectedComplexCarrier || !a.Assembly.InternalU2Socket || !a.Assembly.SU2Compatibility || !a.Assembly.U1PhaseCompatibility || !a.Assembly.FullIntertwinerCandidate {
		t.Fatalf("bad sealed socket assembly: %+v", a.Assembly)
	}
	if len(a.Available.Structures) != 5 || !a.Available.AllAvailable {
		t.Fatalf("bad available structures: %+v", a.Available)
	}
}

func TestGate721MinimalityIndependenceAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Blocked.Verdict) == 0 || a.Blocked.WhyNSelected || a.Blocked.WhyQHasValue || a.Blocked.PhysicalEqualitySU2U1 || a.Blocked.ScalarPotential || a.Blocked.QuarticRuntimeLambda || a.Blocked.HiggsPoleMass || a.Blocked.YukawaOperatorConstruction || a.Blocked.FlavorHierarchy || a.Blocked.CKMPMNS || a.Blocked.BlockedCount != 9 {
		t.Fatalf("blocked physics firewall failed: %+v", a.Blocked)
	}
	if len(a.Minimality.Removals) != 4 || !a.Minimality.RemoveNBreaks || !a.Minimality.RemoveQBreaks || !a.Minimality.RemoveCBreaks || !a.Minimality.RemoveK7PlusBreaks || !a.Minimality.PairMinimal {
		t.Fatalf("minimality audit failed: %+v", a.Minimality)
	}
	if !a.Independence.TypeDistinct || !a.Independence.NotMutuallyDerivable || a.Independence.QFromSevenOver72 || a.Independence.NFromScalarBridgeData || a.Independence.NFromPK7 || a.Independence.QFromAbsN || len(a.Independence.ForbiddenShortcuts) != 4 {
		t.Fatalf("independence audit failed: %+v", a.Independence)
	}
	if a.Physical.TwistorSelectorSealNativeVacuumTheorem || a.Physical.HyperchargeSealNativeDerivation || a.Physical.SealedSocketFullPhysicalHiggsTheorem || a.Physical.SealedSocketHiggsMassTheorem || a.Physical.SealedSocketYukawaTheorem || a.Physical.SealedSocketCKMPMNSTtheorem || !a.Physical.NoScalarPotentialOrRuntimeLambda || !a.Physical.NoHiggsMassTheorem || !a.Physical.NoYukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("physical firewall failed: %+v", a.Physical)
	}
	res := Generation2MinimalHiggsSocketSealPackageAndPromotionBoundaryAuditTheorem().Verify()
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
