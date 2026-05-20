package generation2compactsplittwistresidualinvariantaudit

import (
	"strings"
	"testing"
)

func TestGate639Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.GOmegaAlignedWithGK || !a.Inherited.BKAsScaledGOmegaSK || a.Inherited.CompactOmegaAndBKFused || a.Inherited.NativeSplitCompatibleTwist || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || !a.Inherited.Gate638FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Repetition.RepeatedAcrossRoutes || a.Repetition.ClusterCount < 3 || a.Repetition.Spread > repeatedTolerance || a.Repetition.RhoTwist < 0.47 || a.Repetition.RhoTwist > 0.471 {
		t.Fatalf("bad repetition audit: %+v", a.Repetition)
	}
	seen := map[string]bool{}
	for _, name := range a.Repetition.ClusterRouteNames {
		seen[name] = true
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("rho cluster missing %s: %+v", want, a.Repetition.ClusterRouteNames)
		}
	}
	if !a.Invariance.AllProjectiveTestsPass || a.Invariance.MaxDrift > invarianceTolerance || !a.Invariance.BasisChangeInvariant || !a.Invariance.OmegaRescaleInvariant || !a.Invariance.TargetSignInvariant || !a.Invariance.SKOrientationInvariant || !a.Invariance.DeterminantVolumeStable || !a.Invariance.TraceFreeStable {
		t.Fatalf("bad invariance audit: %+v", a.Invariance)
	}
	if a.SourceSweep.CompactSourcesRemoveRho || a.SourceSweep.BestCompactSourceResidual <= a.SourceSweep.BestSplitTwistResidual {
		t.Fatalf("bad source sweep: %+v", a.SourceSweep)
	}
	if a.Classification.ClassifiedAsArtifact || !a.Classification.ClassifiedAsOrbitDistance || !a.Classification.ClassifiedAsObstruction || a.Classification.Verdict != StatusCompactSplitObstruction {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewalls.ClaimsPhysicalSpacetime || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarRG || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsSplitG2 {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2CompactSplitTwistResidualInvariantAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate638UnfusedInherited, StatusTwistResidualRepeated, StatusResidualInvarianceTests, StatusResidualNotNormalization, StatusProjectiveResidualAudited, StatusCompactSplitObstruction, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusGate639Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
