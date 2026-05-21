package generation2hopfcl17phaseanchorbridgeaudit

import (
	"strings"
	"testing"
)

func TestGate902HopfAndCL17AreStrongestButNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HopfS1.HasOrientedPhaseCircle || !a.HopfS1.MatchesLambdaBarShape || !a.HopfS1.CanReadAirlockIfSealed || a.HopfS1.NativeSocketOrderMapCertified {
		t.Fatalf("bad Hopf route: %s", FormatHopfS1(a.HopfS1))
	}
	if !a.CL17Chirality.OmegaSquaredMinusOne || !a.CL17Chirality.RequiresComplexChirality || !a.CL17Chirality.CanOrientIOverMinusI || a.CL17Chirality.TypedToRightCharacterPhase || a.CL17Chirality.NativeSocketOrderMapCertified {
		t.Fatalf("bad CL17 route: %s", FormatCL17Chirality(a.CL17Chirality))
	}
}

func TestGate902HopfChiralityBridgeIsObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.HopfChiralityBridge
	if !b.HopfAndChiralityCompatible || !b.PointsToSamePhaseWound || b.NativeBridgeCertified || b.CanAnchorRightCharacters {
		t.Fatalf("bridge incorrectly promoted: %s", FormatHopfChiralityBridge(b))
	}
	if !containsAll(b.Failures, []string{FailureNoNativeHopfChiralityBridge}) {
		t.Fatalf("missing bridge failure: %s", FormatHopfChiralityBridge(b))
	}
}

func TestGate902JBoundarySpectralDoNotSelectPhase(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.JKO.RelevantToConjugacy || a.JKO.SelectsLambda || a.JKO.BreaksZ2 || a.JKO.NativeTheorem {
		t.Fatalf("bad J/KO route: %s", FormatJKO(a.JKO))
	}
	if !a.BoundaryPair.HasExteriorOrientation || !a.BoundaryPair.SelectsDegreeOrder || a.BoundaryPair.SelectsRightCharacterPhase {
		t.Fatalf("bad boundary route: %s", FormatBoundaryPair(a.BoundaryPair))
	}
	if !a.SpectralOrientation.DeepCandidate || a.SpectralOrientation.CycleCertified || a.SpectralOrientation.MapsToSocketOrder || a.SpectralOrientation.NativeTheorem {
		t.Fatalf("bad spectral route: %s", FormatSpectralOrientation(a.SpectralOrientation))
	}
}

func TestGate902RankingAndFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.Ranking.StrongestCandidates, []string{HopfS1PhaseOrientation, CL17ComplexChirality}) || a.Ranking.BridgeCandidate != HopfChiralityBridge || a.Ranking.NativeSourceFound {
		t.Fatalf("bad ranking: %s", FormatRanking(a.Ranking))
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || !near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("bad freeze: %s", FormatFreeze(a.Freeze))
	}
}

func TestGate902Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate902Theorem(t *testing.T) {
	res := Generation2HopfCL17PhaseAnchorBridgeAuditTheorem().Verify()
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
