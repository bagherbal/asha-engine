package generation2specialbranchselector

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.EdgeAudit.FullTrianglePreserved || a.EdgeAudit.AnyNativeLawSuppresses13 || a.EdgeAudit.NearestNeighborNativelyForced {
		t.Fatalf("Gate 451 must preserve full triangle: %s", FormatEdgeAudit(a.EdgeAudit))
	}
	if a.PhaseAudit.UniqueRayForced || !a.PhaseAudit.ContainsCZeroSurvivor || !a.PhaseAudit.ContainsNonzeroCSurvivor {
		t.Fatalf("Gate 451 must preserve phase continuum: %s", FormatPhaseAudit(a.PhaseAudit))
	}
	if !a.Firewall.GSTFritzschRelationsQuarantined {
		t.Fatalf("GST/Fritzsch relations must stay quarantined: %s", FormatFirewall(a.Firewall))
	}
}

func TestDeterminantSieve(t *testing.T) {
	full := determinantEpsilon3Coeff(support(true, true, true))
	nn := determinantEpsilon3Coeff(support(true, true, false))
	if full != 2 {
		t.Fatalf("full triangle det coefficient got %d want 2", full)
	}
	if nn != 0 {
		t.Fatalf("nearest-neighbor det coefficient got %d want 0", nn)
	}
}

func TestKMSAllowsOneThreeHarmonic(t *testing.T) {
	a := buildEdgeAudit()
	found := false
	for _, e := range a.Edges {
		if e.Name == "13" {
			found = true
			if !e.KMSInteger || !e.Allowed || e.DeltaK != 2 {
				t.Fatalf("1-3 edge should be allowed integer second harmonic: %s", FormatEdge(e))
			}
		}
	}
	if !found {
		t.Fatal("missing 1-3 edge audit")
	}
}

func TestPhaseRayContinuum(t *testing.T) {
	p := buildPhaseAudit()
	if p.SurvivingNonzeroLiftRays < 3 || !p.ContainsCZeroSurvivor || !p.ContainsNonzeroCSurvivor || p.UniqueRayForced {
		t.Fatalf("phase ray was incorrectly collapsed: %s", FormatPhaseAudit(p))
	}
	if !p.PureYRayLiftDegenerate {
		t.Fatalf("pure Y ray should be recorded as lift-degenerate diagnostic: %s", FormatPhaseAudit(p))
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := Generation2TextureZeroSpecialBranchSelectorNecessaryBoundaryAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("Gate 451 should be a failed-route audit, got %s", res.Status)
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedNativeGeometryPreservesFullTriangle, StatusFailedNoNative13EdgeSuppression, StatusFailedNoNativePhaseRaySelector, "det(K+epsilon X_NN)=0", "integer second harmonic"} {
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
