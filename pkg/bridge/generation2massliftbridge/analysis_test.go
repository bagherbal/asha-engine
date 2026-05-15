package generation2massliftbridge

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Collapse.ForcesClosedTriangle || !a.Collapse.ForcesXGenSupport {
		t.Fatalf("expected closed triangle/X_gen support: %+v", a.Collapse)
	}
	if !a.Axiom.GeometricallyForcedTopology || !a.Axiom.LiftsGeneration2Zero {
		t.Fatalf("expected forced mass-lift topology: %+v", a.Axiom)
	}
	if !a.Firewall.BridgeAmplitudeSealed || a.Firewall.KXYCoeffDimStillFree != KXYCoeffDim {
		t.Fatalf("firewall changed incorrectly: %+v", a.Firewall)
	}
}

func TestDeterminantIdentityAndOpenChainFailure(t *testing.T) {
	open := makeCandidate(EdgeWeights{A: 1, B: 1, C: 0})
	if !open.EndpointBalanced || !open.OpenChain || open.DeterminantNonZero {
		t.Fatalf("balanced open chain must preserve zero determinant: %s", FormatCandidate(open))
	}
	tri := makeCandidate(EdgeWeights{A: 1, B: 1, C: 1})
	if !tri.EndpointBalanced || !tri.ClosedTriangle || !tri.DeterminantNonZero || tri.C3 != 2 {
		t.Fatalf("balanced triangle must lift at cubic order: %s", FormatCandidate(tri))
	}
	unbalanced := makeCandidate(EdgeWeights{A: 1, B: 0, C: 1})
	if !unbalanced.UnbalancedLift || unbalanced.EndpointBalanced {
		t.Fatalf("unbalanced lift must be rejected as lopsided source: %s", FormatCandidate(unbalanced))
	}
}

func TestEnumerationTopologyCounts(t *testing.T) {
	s := enumerate(1)
	if !s.UniqueUnsignedTopology || len(s.BalancedLiftCandidates) != 8 || len(s.OpenChainFailures) != 4 || s.SignedVariants != 8 {
		t.Fatalf("unexpected topology counts: %s", FormatSieve(s))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2MassLiftBridgeStructuralZeroCompatibilitySieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusTriangleBridgeTopologyForced, StatusXGenSupportSelectedAsTopology, StatusFailedAmplitudeNotPredicted, "det(K_gen + ε B_lift) = 2 ε^3"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
