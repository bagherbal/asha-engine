package generation2basisinvariance

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.BasisAudit.GeneralUnitaryRejected || !a.BasisAudit.AllNativeAllowedPreserve13 || a.BasisAudit.AnyNativeAllowedDeletes13 {
		t.Fatalf("basis audit did not reject fake suppression: %s", FormatBasisAudit(a.BasisAudit))
	}
	if a.Support.CanRephaseToNN || !a.Support.SupportPatternInvariant {
		t.Fatalf("support audit allowed a fake NN rephase: %s", FormatSupportAudit(a.Support))
	}
	if a.Spectral.SameInvariantClass {
		t.Fatalf("spectral audit incorrectly identified triangle with NN chain: %s", FormatSpectralAudit(a.Spectral))
	}
	if !a.Firewall.GSTFritzschRelationsQuarantined {
		t.Fatalf("GST/Fritzsch relation must remain quarantined: %s", FormatFirewall(a.Firewall))
	}
}

func TestCommutatorNorms(t *testing.T) {
	full := commutatorNormSquared(support(true, true, true))
	nn := commutatorNormSquared(support(true, true, false))
	if full != 12 {
		t.Fatalf("full triangle commutator norm^2 got %d want 12", full)
	}
	if nn != 4 {
		t.Fatalf("nearest-neighbor commutator norm^2 got %d want 4", nn)
	}
}

func TestDeterminantClasses(t *testing.T) {
	full := determinantEpsilon3Coeff(support(true, true, true))
	nn := determinantEpsilon3Coeff(support(true, true, false))
	if full != 2 {
		t.Fatalf("full triangle det coefficient got %d want 2", full)
	}
	if nn != 0 {
		t.Fatalf("nearest-neighbor det coefficient got %d want 0", nn)
	}
}

func TestNativeAllowedTransformationsCannotDelete13(t *testing.T) {
	b := buildBasisAudit()
	for _, tr := range b.Transformations {
		if tr.AllowedNativeGauge && tr.CanDelete13Edge {
			t.Fatalf("native-allowed transform deletes 1-3: %s", FormatBasisTransformation(tr))
		}
	}
}

func TestTheoremPassesAsFailedRouteAudit(t *testing.T) {
	res := Generation2FamilyBasisInvarianceTextureGaugeArtifactAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
	if string(res.Status) != "FAILED_ROUTE" {
		t.Fatalf("Gate 452 should be a failed-route audit, got %s", res.Status)
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedNoBasisSuppression, StatusFailedNNTextureNotGaugeEquivalent, StatusFailedGeneralFamilyRotationBreaksKAddress, "centralizer_U(3)(K_gen)=U(1)^3", "||[K,X_triangle]||_F^2=12"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func TestAlmostEqualHelper(t *testing.T) {
	if !almostEqual(1, 1+1e-13) {
		t.Fatal("almostEqual unexpectedly false")
	}
}

func TestTextListHelper(t *testing.T) {
	if got := textList([]string{"a", "b"}); got != "a; b" {
		t.Fatalf("textList got %q", got)
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
