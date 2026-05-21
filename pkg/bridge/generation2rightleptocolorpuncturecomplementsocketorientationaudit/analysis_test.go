package generation2rightleptocolorpuncturecomplementsocketorientationaudit

import (
	"strings"
	"testing"
)

func TestGate841ComplementRanks(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Complement.SupportAnatomyCertified || a.Complement.CompressionTheoremCertified {
		t.Fatalf("bad complement theorem status: %s", FormatComplement(a.Complement))
	}
	if a.Complement.ActiveRank != 7 || a.Complement.PunctureRank != 1 || a.Complement.FullRank != 8 || !a.Complement.EightEqualsSevenPlusOne || !a.Complement.Orthogonal || !a.Complement.Complete {
		t.Fatalf("bad complement ranks: %s", FormatComplement(a.Complement))
	}
}

func TestGate841BMinusLCompensatingPuncture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.BMinusL.DominantColorTrace != 1 || a.BMinusL.RestQuartetTrace != 0 || a.BMinusL.ActiveTrace != 1 || a.BMinusL.PunctureTrace != -1 || a.BMinusL.FullTrace != 0 {
		t.Fatalf("bad B-L compensation: %s", FormatBMinusL(a.BMinusL))
	}
	if !a.BMinusL.ActivePlusPunctureCancel || !a.BMinusL.FullNeutral || !a.BMinusL.CompensatingPuncturePattern {
		t.Fatalf("missing compensation flags: %s", FormatBMinusL(a.BMinusL))
	}
}

func TestGate841SterileNullEdgeAndOrientationBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sterile.RightSocket || !a.Sterile.Leptonic || !a.Sterile.Colorless || !a.Sterile.ExcludedFromActive || a.Sterile.Rank != 1 {
		t.Fatalf("bad sterile candidate anatomy: %s", FormatSterile(a.Sterile))
	}
	if a.Sterile.DFEdgeDataAvailable || a.Sterile.NullEdgeCertified || a.Sterile.SterilePunctureCertified || a.Sterile.PhysicalParticleAssignmentCertified {
		t.Fatalf("sterile puncture over-certified: %s", FormatSterile(a.Sterile))
	}
	if a.Orientation.DominantOrientationCertified || a.Orientation.RestOrientationCertified || a.Orientation.OrientationMapCertified || a.Orientation.DFOrHiggsSelectorCertified || a.Orientation.BoundaryRestSelectorCertified {
		t.Fatalf("orientation over-certified: %s", FormatOrientation(a.Orientation))
	}
	if !containsAll(a.Sterile.Failures, []string{FailureNoDFEdgeData, FailureNoNullEdgeTheorem, FailureNoSterilePunctureTheorem, FailureNoRightNeutrinoTheorem}) {
		t.Fatalf("missing sterile failures: %s", strings.Join(a.Sterile.Failures, ","))
	}
}

func TestGate841FirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.PunctureComplementLawCertified || !a.Impact.BMinusLCompensationFound || !a.Impact.SterileNullEdgeStillUncertified || !a.Impact.OrientationStillMissing || !a.Impact.CompressionMapStillMissing || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact invalid: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoDFEdgeData || !a.Firewalls.NoSterilePunctureTheorem || !a.Firewalls.PunctureNotPhysicalParticle || !a.Firewalls.NoDominantOrientationTheorem || !a.Firewalls.NoRestOrientationTheorem || !a.Firewalls.NoTypedCompressionMap || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.AlphaSealed || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate841 {
		t.Fatalf("firewall invalid: %+v", a.Firewalls)
	}
	res := Generation2RightLeptoColorPunctureComplementSocketOrientationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
