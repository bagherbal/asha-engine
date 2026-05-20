package generation2flavorboundaryreadoutandstresspullorientationsealfactorizationaudit

import (
	"strings"
	"testing"
)

func TestGate787FactorizationAndDegreeRule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate786.Inherited || !strings.Contains(a.Gate786.PackageSeal, "Theta_ext") || a.Gate786.PriorNative {
		t.Fatalf("bad Gate786 inheritance: %+v", a.Gate786)
	}
	if !a.Factorization.Recorded || !containsAll(a.Factorization.Subseals, []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"}) || !a.Factorization.DegreeZeroCanonical || !a.Factorization.ThreeNontrivialSubobjects {
		t.Fatalf("bad factorization: %+v", a.Factorization)
	}
	if !a.DegreeRule.Audited || !strings.Contains(a.DegreeRule.Rule, "Lambda^(n-1)B_boundary") || !a.DegreeRule.ExplainsCubicStop || a.DegreeRule.Native || !a.DegreeRule.ThetaExtSealed || !a.DegreeRule.ProjectorPowersContinue {
		t.Fatalf("bad degree rule audit: %+v", a.DegreeRule)
	}
}

func TestGate787FlavorReadoutAndBoundaryKappa(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FlavorReadout.Audited || !closeRel(a.FlavorReadout.KappaERed, kappaERedSnapshot, 1e-15) || !closeRel(a.FlavorReadout.KappaOrient+a.FlavorReadout.KappaBoundary, kappaERedSnapshot, 1e-15) || !a.FlavorReadout.MixedReadout || a.FlavorReadout.NativeFromBoundary || a.FlavorReadout.NativeFlavorTheorem {
		t.Fatalf("bad flavor-boundary readout: %+v", a.FlavorReadout)
	}
	if !containsAll(a.FlavorReadout.TermTypes, []string{"PMNS", "CKM", "hypercharge", "boundary-stress"}) {
		t.Fatalf("missing term types: %+v", a.FlavorReadout.TermTypes)
	}
	if !a.KappaBoundary.Audited || !a.KappaBoundary.StrongSourceType || a.KappaBoundary.FullKappaESource || !strings.Contains(a.KappaBoundary.MainNonNativePart, "J_CKM") || a.KappaBoundary.NativeCoupling {
		t.Fatalf("bad boundary kappa audit: %+v", a.KappaBoundary)
	}
	if !closeRel(a.KappaBoundary.KappaBoundary, (-5.0/3.0+xiBoundarySnapshot*pK7Snapshot)*sSplitSnapshot*sSplitSnapshot, 1e-18) {
		t.Fatalf("bad kappa_boundary value: %.18g", a.KappaBoundary.KappaBoundary)
	}
}

func TestGate787BoundaryAxesStressPullAndRuntime(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Axis.Audited || !a.Axis.SplitAxisCandidate || a.Axis.SplitAxisSourcesKappaE || !a.Axis.MidpointAxisCandidate || a.Axis.MidpointAxisSourcesKappaE || a.Axis.BoundaryAxesReplaceFlavor {
		t.Fatalf("bad axis candidates: %+v", a.Axis)
	}
	if !a.StressPull.Audited || !closeRel(a.StressPull.Magnitude, 2*pK7Snapshot, 1e-15) || !a.StressPull.NegativeSign || a.StressPull.NativeNegativeSign || a.StressPull.MatchingSignNative || !containsAll(a.StressPull.CandidateSignSources, []string{"ordered boundary orientation", "restorative stress-pull"}) {
		t.Fatalf("bad stress-pull audit: %+v", a.StressPull)
	}
	if !a.Runtime.Audited || a.Runtime.ContainsForbidden || !a.Runtime.FormulaLevelIndependent || a.Runtime.TheoremLevelIndependent || !containsAll(a.Runtime.ForbiddenDirectVariables, []string{"lambda_runtime", "m_H_tree", "C_Higgs", "G_F", "v"}) {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
}

func TestGate787MinimalClassificationPropagationFirewallsAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Minimal.Recorded || !containsAll(a.Minimal.Seals, []string{"DegreeRuleSeal", "FlavorBoundaryReadoutSeal", "BoundaryStressPullOrientationSeal"}) || a.Minimal.Native || !strings.Contains(a.Minimal.Mapping["BoundaryStressPullOrientationSeal"], "negative sign") {
		t.Fatalf("bad minimal classification: %+v", a.Minimal)
	}
	if !a.Propagation.Recorded || !strings.Contains(a.Propagation.FWall3, "Level B+") || !strings.Contains(a.Propagation.CHistory, "Level B") || !strings.Contains(a.Propagation.CHiggs, "not Level C") {
		t.Fatalf("bad propagation: %+v", a.Propagation)
	}
	if !a.Firewalls.Enforced || a.Firewalls.KappaENativeFlavorTheorem || a.Firewalls.KappaBoundaryFullFlavorTheorem || a.Firewalls.KappaOrientPMNSCKMTheorem || a.Firewalls.SplitMidpointAxesKappaETheorem || a.Firewalls.TwoPMagnitudeFullSignTheorem || a.Firewalls.NegativeCubicSignNative || a.Firewalls.ResponsePackageNativeGenerating || a.Firewalls.FWallNative || a.Firewalls.KappaLambdaNative || a.Firewalls.CHistoryIndependent || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not make the exterior response package native") || !strings.Contains(a.FinalStatement, "three sharper missing subobjects") || !strings.Contains(a.FinalStatement, "FlavorBoundaryReadoutSeal") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate787TheoremStatuses(t *testing.T) {
	res := Generation2FlavorBoundaryReadoutAndStressPullOrientationSealFactorizationAuditTheorem().Verify()
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
