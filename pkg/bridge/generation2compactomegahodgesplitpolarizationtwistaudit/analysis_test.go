package generation2compactomegahodgesplitpolarizationtwistaudit

import (
	"strings"
	"testing"
)

func TestGate638Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.K7Dimension != 7 || a.Inherited.BKInertia != "(4,3,0)" || a.Inherited.BestOmegaInertia != "(7,0,0)" || !a.Inherited.NativePullbackTensorExists || a.Inherited.CompatibleOmegaKCertified || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || !a.Inherited.CompactOmegaAndBKConflict || !a.Inherited.Gate637FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.MetricAlign.GOmegaInertiaPlus != 7 || a.MetricAlign.GOmegaInertiaMinus != 0 || a.MetricAlign.GOmegaInertiaZero != 0 || !a.MetricAlign.AlignedWithGK || !a.MetricAlign.CompactPositive || a.MetricAlign.RelativeResidualToGK > 1e-8 {
		t.Fatalf("bad metric alignment: %+v", a.MetricAlign)
	}
	if !a.Reconstruction.BKEqualsGKSK || !a.Reconstruction.BKEqualsScaledGOmegaSK || a.Reconstruction.ScaledGOmegaSKResidual > 1e-8 {
		t.Fatalf("bad reconstruction: %+v", a.Reconstruction)
	}
	if !a.SKAction.SKOrthogonalForGOmega || !(a.SKAction.Omega3Inertia == "(7,0,0)" || a.SKAction.Omega3Inertia == "(0,7,0)") || a.SKAction.OrthogonalityResidual > 1e-8 {
		t.Fatalf("bad S_K action: %+v", a.SKAction)
	}
	if len(a.Twists.Candidates) != 4 || a.Twists.AdmissibleAlternatingCandidates != 4 || a.Twists.SplitCompatibleCandidates != 0 || a.Twists.NativeSKTwistMatchesBK {
		t.Fatalf("bad twists: %+v", a.Twists)
	}
	for _, c := range a.Twists.Candidates {
		if c.AntisymmetryResidual > 1e-9 || !c.HitchinMetricComputed || c.SplitCompatibleWithBK {
			t.Fatalf("bad twist candidate: %+v", c)
		}
	}
	if !a.CrossProduct.CompactCrossProductDefined || a.CrossProduct.OmegaBMatchesBK {
		t.Fatalf("bad cross product audit: %+v", a.CrossProduct)
	}
	if !a.Interpretation.GOmegaAlignedWithGK || !a.Interpretation.BKIsHodgePolarizedCompactMetric || a.Interpretation.NativeSplitCompatibleTwistFound || a.Interpretation.CompactOmegaAndBKFused {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if a.Firewalls.ClaimsPhysicalSpacetime || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsScalarRG || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsSplitG2 {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2CompactOmegaHodgeSplitPolarizationTwistAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate637Inherited, StatusGOmegaToGKAlignment, StatusGOmegaAlignedCompactGK, StatusBKEqualsGKSK, StatusSKActionOnOmegaAudited, StatusTwistAdmissibilityAudited, StatusNoSKTwistMatchesBK, StatusCompactOmegaBKDoNotFuse, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusGate638Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
