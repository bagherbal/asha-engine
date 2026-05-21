package generation2dominantsocketleptocolortracecompressionmapaudit

import (
	"strings"
	"testing"
)

func TestGate839FiniteBodyAndWInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Body.Gate838Inherited || !a.Body.CoarseLedgerExists || a.Body.HPartDim != 16 || a.Body.HFDim != 32 || a.Body.EDim != 4 || a.Body.WDim != 4 {
		t.Fatalf("bad inherited body: %s", FormatBody(a.Body))
	}
	if a.W.P1Rank != 1 || a.W.P3Rank != 3 || a.W.Dim != 4 || !a.W.BMinusLTraceZero || !a.W.BMinusLRestActionOnW {
		t.Fatalf("bad W carrier: %s", FormatW(a.W))
	}
}

func TestGate839SocketCompressionRankAnatomy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Compression.TopRank != 3 || a.Compression.RestRank != 4 || a.Compression.AggregateRank != 7 || !a.Compression.MatchesI3PlusW || !a.Compression.NonCircular || a.Compression.UsesObservedData || a.Compression.IsTheorem {
		t.Fatalf("bad compression candidate: %s", FormatCompression(a.Compression))
	}
	if a.Compression.TopSelectorCertified || a.Compression.RestSelectorCertified || a.Compression.CompressionMapCertified {
		t.Fatalf("compression over-certified: %s", FormatCompression(a.Compression))
	}
	if !containsAll(a.Compression.Failures, []string{FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureNoSocketPairCompressionMap, FailureCompressionCandidateNotTheorem}) {
		t.Fatalf("missing compression failures: %s", strings.Join(a.Compression.Failures, ","))
	}
}

func TestGate839FineSocketSelectorsRemainMissing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sockets.RankOneSocketProjectorsPossible || a.Sockets.FineSocketProjectorsCertified || !a.Sockets.SocketAtomsBasisDependentWithoutSelector || a.Sockets.DominantSelectorCertified || a.Sockets.RestSelectorCertified || a.Sockets.EtErCanonical {
		t.Fatalf("socket selector audit invalid: %s", FormatSockets(a.Sockets))
	}
	if !containsAll(a.Sockets.Failures, []string{FailureNoFineSocketProjectors, FailureSocketAtomsBasisDependent, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureEtErNotCanonical}) {
		t.Fatalf("missing socket failures: %s", strings.Join(a.Sockets.Failures, ","))
	}
}

func TestGate839ShadowImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Shadow.TopIsIdentityI3Candidate || !a.Shadow.RestUsesBMinusLTransferOnW || a.Shadow.AlphaDerivedByCompression || a.Shadow.TraceMagnitudeReadoutCertified || a.Shadow.AggregateOperatorIsSectorLedger || a.Shadow.R3 || a.Shadow.R4 {
		t.Fatalf("shadow over-certified: %s", FormatShadow(a.Shadow))
	}
	if !a.Impact.CompressionCandidateFormulated || !a.Impact.FiniteBodyLocationForAggregateSuggested || !a.Impact.RankAnatomyExplainedConditionally || !a.Impact.SelectorsMissing || !a.Impact.CompressionMapMissing || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact invalid: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoFineSocketProjectors || !a.Firewalls.NoDominantSelector || !a.Firewalls.NoRestSelector || !a.Firewalls.NoCompressionMap || !a.Firewalls.SevenNotK7 || !a.Firewalls.NoTraceMagnitudeReadout || !a.Firewalls.AlphaSealed || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate839 {
		t.Fatalf("firewall invalid: %+v", a.Firewalls)
	}
	res := Generation2DominantSocketLeptoColorTraceCompressionMapAuditTheorem().Verify()
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
