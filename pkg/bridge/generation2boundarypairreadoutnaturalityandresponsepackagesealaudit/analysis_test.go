package generation2boundarypairreadoutnaturalityandresponsepackagesealaudit

import (
	"strings"
	"testing"
)

func TestGate786InventoryBasisAndDegreeOneAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate785.Inherited || !strings.Contains(a.Gate785.ConditionalPackage, "chi_ext") || a.Gate785.PriorNative {
		t.Fatalf("bad Gate785 inheritance: %+v", a.Gate785)
	}
	if !a.Inventory.Recorded || !a.Inventory.RolesSeparated || a.Inventory.AutoDefinesPackage || !strings.Contains(a.Inventory.BoundaryCarrier, "B_boundary") || !containsAll(a.Inventory.BoundaryAxes, []string{"b_lambda", "b_R"}) || !containsAll(a.Inventory.ExteriorPackage, []string{"Theta_ext", "chi_ext", "negative stress-pull sign"}) {
		t.Fatalf("bad inventory: %+v", a.Inventory)
	}
	if !a.Basis.Audited || !a.Basis.LabelledBridgeBasis || a.Basis.NativeInvariantBasis || a.Basis.NativeVerdict != StatusLabelledBridgeBasisNotNativeInvariantBasis {
		t.Fatalf("bad basis audit: %+v", a.Basis)
	}
	if !a.DegreeOne.Audited || !a.DegreeOne.SplitAxisCandidate || a.DegreeOne.SplitAxisSourcesKappaE || !a.DegreeOne.MidpointAxisCandidate || a.DegreeOne.MidpointSourcesKappaE || !a.DegreeOne.FlavorReadoutCoefficient || a.DegreeOne.KappaESourcedByBoundary || a.DegreeOne.NativeReadoutTheorem {
		t.Fatalf("bad degree-one audit: %+v", a.DegreeOne)
	}
}

func TestGate786DegreeTwoChiExtNaturalityAndSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DegreeTwo.Audited || !a.DegreeTwo.Lambda2Exists || !strings.Contains(a.DegreeTwo.VolumeForm, "omega_B") || !a.DegreeTwo.RequiresOrderedBasis || !closeRel(a.DegreeTwo.TwoP, 7.0/36.0, 1e-15) || a.DegreeTwo.NativeOrderedOrientation || a.DegreeTwo.NativeNegativeStressPull {
		t.Fatalf("bad degree-two audit: %+v", a.DegreeTwo)
	}
	if !a.ChiExt.Audited || !a.ChiExt.DegreeZeroCanonical || !closeRel(a.ChiExt.Chi0, 1, 1e-15) || !a.ChiExt.DegreeOneRequiresFlavor || !closeRel(a.ChiExt.Chi1, kappaERedSnapshot, 1e-15) || !a.ChiExt.DegreeTwoRequiresK7AndSign || !closeRel(a.ChiExt.Chi2, -2*pK7Snapshot, 1e-15) || a.ChiExt.NativeFromBoundaryPairAlone {
		t.Fatalf("bad chi_ext audit: %+v", a.ChiExt)
	}
	if !a.Naturality.Audited || !a.Naturality.AbstractGL2Freedom || a.Naturality.CanonicalBetaUnderGL2 || a.Naturality.CanonicalOmegaSignUnderGL2 || a.Naturality.CanonicalChiUnderGL2 || !a.Naturality.LabelledPairReducesSymmetry || !a.Naturality.LabelledPackageBridgeSealed {
		t.Fatalf("bad naturality audit: %+v", a.Naturality)
	}
	if !a.Seal.Defined || a.Seal.Name != "BoundaryExteriorResponsePackageSeal" || !a.Seal.Minimal || a.Seal.Native || !a.Seal.MatchesFWall || !containsAll(a.Seal.Components, []string{"Theta_ext", "chi_ext", "ordered boundary orientation"}) || !containsAll(a.Seal.Supplies, []string{"chi_ext(beta_B)=kappa_e_red", "chi_ext(omega_B)=-2p", "Theta_ext(M_n>=4)=0"}) {
		t.Fatalf("bad response-package seal: %+v", a.Seal)
	}
}

func TestGate786ImpactPropagationFirewallsAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.Audited || !a.Impact.RepresentableBySeal || a.Impact.NativeGeneratingFunction || !strings.Contains(a.Impact.Status, "not a native") {
		t.Fatalf("bad F_wall impact: %+v", a.Impact)
	}
	if !a.Propagation.Recorded || !strings.Contains(a.Propagation.FWall3, "Level B+") || !strings.Contains(a.Propagation.KappaLambda, "Level B") || !strings.Contains(a.Propagation.CHistory, "Level B") || !strings.Contains(a.Propagation.CHiggs, "not Level C") {
		t.Fatalf("bad propagation: %+v", a.Propagation)
	}
	if !a.Firewalls.Enforced || a.Firewalls.LabelledAxesNative || a.Firewalls.SplitAxisKappaETheorem || a.Firewalls.MidpointKappaETheorem || a.Firewalls.OmegaSignTheorem || a.Firewalls.TwoPMagnitudeFullTheorem || a.Firewalls.SealNativeGeneratingFunc || a.Firewalls.FWallNative || a.Firewalls.KappaLambdaNative || a.Firewalls.CHistoryIndependent || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNative {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not source the package natively") || !strings.Contains(a.FinalStatement, "BoundaryExteriorResponsePackageSeal") || !strings.Contains(a.FinalStatement, "next bottleneck") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate786TheoremStatuses(t *testing.T) {
	res := Generation2BoundaryPairReadoutNaturalityAndResponsePackageSealAuditTheorem().Verify()
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
