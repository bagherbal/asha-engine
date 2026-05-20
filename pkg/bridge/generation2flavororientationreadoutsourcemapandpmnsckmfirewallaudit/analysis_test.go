package generation2flavororientationreadoutsourcemapandpmnsckmfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate788KappaOrientDecomposition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate787.Inherited || !a.Gate787.KappaOrientIsFocus || !strings.Contains(a.Gate787.CompositeSeal, "FlavorBoundaryReadoutSeal") {
		t.Fatalf("bad Gate787 inheritance: %+v", a.Gate787)
	}
	if !a.Decomposition.Recorded || a.Decomposition.Formula != "sin^2(theta13)/4 - J_CKM" || !a.Decomposition.ShapeTyped || a.Decomposition.Native {
		t.Fatalf("bad kappa_orient decomposition: %+v", a.Decomposition)
	}
	if !strings.Contains(a.Decomposition.PMNSTerm, "PMNS") || !strings.Contains(a.Decomposition.CKMTerm, "J_CKM") || !strings.Contains(a.Decomposition.NegativeSign, "subtraction") {
		t.Fatalf("bad term typing: %+v", a.Decomposition)
	}
}

func TestGate788PMNSAndCKMFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PMNS.Audited || a.PMNS.Theta13Native || !a.PMNS.QuarterResonance || a.PMNS.TypedMapFromK7Quarter || !a.PMNS.RemainsFlavorSealInput {
		t.Fatalf("bad PMNS audit: %+v", a.PMNS)
	}
	if !containsAll(a.PMNS.QuarterResonanceSources, []string{"rho_plus=I_K7+/4", "Tr(rho_plus P_rad)=1/4"}) {
		t.Fatalf("missing quarter resonance sources: %+v", a.PMNS.QuarterResonanceSources)
	}
	if !a.CKM.Audited || a.CKM.JCKMNative || !a.CKM.NegativeSignCandidate || a.CKM.NativeSignTheorem || !a.CKM.RemainsFlavorSeal {
		t.Fatalf("bad CKM audit: %+v", a.CKM)
	}
	if !strings.Contains(a.CKM.SourceType, "orientation area") {
		t.Fatalf("bad CKM source type: %s", a.CKM.SourceType)
	}
}

func TestGate788BoundaryCorrectionAndGeometryCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BoundaryOnly.Audited || !closeRel(a.BoundaryOnly.KappaOrient, kappaOrientSnapshot, 2e-15) || !closeRel(a.BoundaryOnly.KappaBoundary, kappaBoundarySnapshot, 2e-15) || a.BoundaryOnly.AbsRatioBoundaryToOrient >= 1e-3 || !a.BoundaryOnly.BoundaryPartSmallCorrection || a.BoundaryOnly.BoundaryReplacesOrient {
		t.Fatalf("bad boundary-only audit: %+v", a.BoundaryOnly)
	}
	if !a.Geometry.Audited || a.Geometry.K7HodgeDerivesPMNSCKM || a.Geometry.HiggsRadialDerivesTheta13 || a.Geometry.BoundaryPairDerivesFlavorMixing || a.Geometry.NEffDerivesMixingAngles || a.Geometry.GenerationMixingOperatorFound {
		t.Fatalf("bad geometry candidates: %+v", a.Geometry)
	}
	for _, key := range []string{"K7 Hodge polarity 4|3", "K7+ radial/Higgs event", "Boundary pair", "N_eff"} {
		if _, ok := a.Geometry.Candidates[key]; !ok {
			t.Fatalf("missing geometry candidate %s", key)
		}
	}
}

func TestGate788SealRefinementRuntimePropagationAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SealRefinement.Recorded || !containsAll(a.SealRefinement.RefinedSeals, []string{"FlavorOrientationReadoutSeal", "BoundaryGaugeCorrectionSeal"}) || !a.SealRefinement.BoundaryGaugeStronglyTyped || !a.SealRefinement.OrientationTrueObstruction || a.SealRefinement.OrientationSealNative {
		t.Fatalf("bad seal refinement: %+v", a.SealRefinement)
	}
	if !a.Runtime.Audited || a.Runtime.ContainsForbidden || !a.Runtime.FormulaLevelIndependent || a.Runtime.TheoremLevelIndependent {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
	if !a.Propagation.Recorded || !strings.Contains(a.Propagation.KappaOrient, "FlavorOrientationReadoutSeal") || !strings.Contains(a.Propagation.CHistory, "Level B") || !strings.Contains(a.Propagation.CHiggs, "not Level C") {
		t.Fatalf("bad propagation: %+v", a.Propagation)
	}
	if !a.Firewalls.Enforced || a.Firewalls.KappaOrientNativeFlavorTheorem || a.Firewalls.Theta13DerivedPMNSTheorem || a.Firewalls.JCKMDerivedCKMTheorem || a.Firewalls.QuarterResonanceProof || a.Firewalls.K7QuarterTheta13SourceTheorem || a.Firewalls.NEffMixingAngleTheorem || a.Firewalls.BoundaryPairFlavorMixingTheorem || a.Firewalls.KappaBoundaryFullKappaETheorem || a.Firewalls.FWallNativeBoundaryTheorem || a.Firewalls.CHistoryFullPrediction || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate788FinalStatementAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not source kappa_orient natively") || !strings.Contains(a.FinalStatement, "FlavorOrientationReadoutSeal") || !strings.Contains(a.FinalStatement, "native generation/mixing operator") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2FlavorOrientationReadoutSourceMapAndPMNSCKMFirewallAuditTheorem().Verify()
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
