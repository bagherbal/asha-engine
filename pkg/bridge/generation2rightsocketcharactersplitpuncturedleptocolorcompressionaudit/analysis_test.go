package generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit

import (
	"strings"
	"testing"
)

func TestGate840RightCharacterSplitSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RightSplit.CharacterSplitSealAudited || !a.RightSplit.CharacterProjectorsSourceTypedBySeal || !a.RightSplit.UnorderedPairCertified || !a.RightSplit.CharacterPairOrthogonal || !a.RightSplit.CharacterPairComplete {
		t.Fatalf("right character split not source-typed: %s", FormatRightSplit(a.RightSplit))
	}
	if a.RightSplit.NativeDerivationCertified || a.RightSplit.ExplicitRhoRMatrixCertified || a.RightSplit.FullRhoFActionLedgerCertified || a.RightSplit.DominantRestOrientationCertified {
		t.Fatalf("right character split over-certified: %s", FormatRightSplit(a.RightSplit))
	}
	if !containsAll(a.RightSplit.Failures, []string{FailureRightCharacterSplitSealNotNative, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureDominantRestOrientationMissing}) {
		t.Fatalf("missing split failures: %s", strings.Join(a.RightSplit.Failures, ","))
	}
}

func TestGate840PuncturedRectangleRanks(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Puncture.FullRank != 8 || a.Puncture.TopRank != 3 || a.Puncture.RestRank != 4 || a.Puncture.SelectedRank != 7 || a.Puncture.ExcludedRank != 1 || !a.Puncture.RanksCloseRightRectangle {
		t.Fatalf("bad puncture ranks: %s", FormatPuncture(a.Puncture))
	}
	if !a.Puncture.UsesRightCharacterSockets || a.Puncture.OrientationCertified || a.Puncture.CompressionMapCertified || a.Puncture.IsTheorem {
		t.Fatalf("puncture over-certified: %s", FormatPuncture(a.Puncture))
	}
}

func TestGate840BMinusLPunctureConservation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.BMinusL.TopTrace != 1 || a.BMinusL.RestTrace != 0 || a.BMinusL.SelectedTrace != 1 || a.BMinusL.ExcludedTrace != -1 || a.BMinusL.FullRectangleTrace != 0 {
		t.Fatalf("bad B-L balance: %s", FormatBMinusL(a.BMinusL))
	}
	if !a.BMinusL.SelectedExcludedCancel || !a.BMinusL.FullRectangleNeutral || !a.BMinusL.PunctureConservationPattern {
		t.Fatalf("missing conservation pattern: %s", FormatBMinusL(a.BMinusL))
	}
}

func TestGate840ShadowImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.FineSocketProblemPartiallyResolvedBySeal || !a.Impact.OrientationStillMissing || !a.Impact.CompressionMapStillMissing || !a.Impact.BMinusLConservationPatternFound || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact invalid: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.CharacterSplitSealNotNative || !a.Firewalls.OrientationMissing || !a.Firewalls.NoDominantSelector || !a.Firewalls.NoRestSelector || !a.Firewalls.NoCompressionMap || !a.Firewalls.ExcludedSingletonNotParticle || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.AlphaSealed || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate840 {
		t.Fatalf("firewall invalid: %+v", a.Firewalls)
	}
	res := Generation2RightSocketCharacterSplitPuncturedLeptoColorCompressionAuditTheorem().Verify()
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
