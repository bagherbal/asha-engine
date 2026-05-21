package generation2rightcharacterphaseorientationsourceaudit

import (
	"strings"
	"testing"
)

func TestGate899InheritedZ2RequiresPhaseOrientation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	i := a.Inherited
	if !i.Z2Family || !i.RequiresPhaseOrientation || i.CanSelectSigmaPlusNatively {
		t.Fatalf("bad inherited audit: %s", FormatInherited(i))
	}
	if !containsAll(i.Characters, []string{RightCharacterPlus, RightCharacterMinus}) || !containsAll(i.Airlocks, []string{PuncturePlus, PunctureMinus}) || !containsAll(i.Failures, []string{FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation}) {
		t.Fatalf("missing inherited statuses: %s", FormatInherited(i))
	}
}

func TestGate899HopfAndChiralityAreStrongestCandidatesButNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	h := a.HopfS1
	if !h.PhaseOrientation || !h.CanOrientLambdaBar || !h.StrongestCandidate || h.NativeTheorem || !h.SelectsPlusIfSealed {
		t.Fatalf("bad Hopf route: %s", FormatHopfS1(h))
	}
	if !containsAll(h.Supports, []string{SupportHopfStrongestCandidate, SupportHopfCanSourceIfSealed}) || !containsAll(h.Failures, []string{FailureNoHopfToSocketOrderTheorem, FailurePhaseOrientationSealNotNative}) {
		t.Fatalf("missing Hopf statuses: %s", FormatHopfS1(h))
	}
	c := a.ComplexChirality
	if !c.OmegaSquaredMinusOne || !c.RequiresComplexAirlock || !c.CanOrientIOverMinusI || c.TypedToRightCharacters || c.NativeSocketOrderTheorem {
		t.Fatalf("bad chirality route: %s", FormatComplexChirality(c))
	}
	if !containsAll(c.Supports, []string{SupportChiralityCandidate, SupportCL17FirewallCandidate}) || !containsAll(c.Failures, []string{FailureNoCL17ChiralityToRightOrderMap, FailureComplexChiralityNotSocketOrderTheorem}) {
		t.Fatalf("missing chirality statuses: %s", FormatComplexChirality(c))
	}
}

func TestGate899JBoundaryAndSpectralRoutesDoNotSelectSocketPhase(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.JKO.RelevantToConjugation || a.JKO.KOSignCertified || a.JKO.SelectsPlus || a.JKO.NativeTheorem {
		t.Fatalf("J/KO leaked: %s", FormatJKO(a.JKO))
	}
	if !a.BoundaryPair.ExteriorOrientation || !a.BoundaryPair.SelectsDegreeOrder || a.BoundaryPair.SelectsSocketPhase {
		t.Fatalf("boundary leaked: %s", FormatBoundaryPair(a.BoundaryPair))
	}
	if !a.SpectralOrientation.DeepCandidate || a.SpectralOrientation.CycleCertified || a.SpectralOrientation.MapsToSocketOrder || a.SpectralOrientation.NativeTheorem {
		t.Fatalf("spectral leaked: %s", FormatSpectralOrientation(a.SpectralOrientation))
	}
}

func TestGate899RankingReducesWoundToPhaseOrientationSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Ranking
	if !containsAll(r.StrongestCandidates, []string{HopfS1PhaseOrientation, ComplexChiralityAirlock}) || r.PhaseSealName != RightCharacterPhaseSeal || r.SocketSelectorName != SocketOrderPhaseSelector || r.NativeSourceFound || r.NextFrontier != NextFrontier {
		t.Fatalf("bad ranking: %s", FormatRanking(r))
	}
	if !containsAll(r.Supports, []string{SupportHopfAndChiralityStrongest, SupportSocketOrderReducedToPhaseSeal}) || !containsAll(r.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus}) {
		t.Fatalf("missing ranking statuses: %s", FormatRanking(r))
	}
}

func TestGate899FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate899Theorem(t *testing.T) {
	res := Generation2RightCharacterPhaseOrientationSourceAuditTheorem().Verify()
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
