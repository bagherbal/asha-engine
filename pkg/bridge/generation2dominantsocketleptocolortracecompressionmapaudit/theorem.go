package generation2dominantsocketleptocolortracecompressionmapaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "generation2.dominant.socket.leptocolor.trace.compression.map.audit"
	theoremName = "Gate 839 — DominantSocket LeptoColor Trace-Compression Map Audit"
)

func Generation2DominantSocketLeptoColorTraceCompressionMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 838 sealed finite-sector body", Passed: a.Body.Gate838Inherited && a.Body.CoarseLedgerExists && a.Body.HPartDim == HPartDim && a.Body.HFDim == HFDim && containsAll(a.Body.Supports, []string{SupportGate838SectorBodyInherited, SupportSectorBodyBeforeCompression}), Detail: FormatBody(a.Body)},
			{Name: "reverify W carrier and B-L rest action", Passed: a.W.Dim == WDim && a.W.P1Rank == LeptonBlockDim && a.W.P3Rank == ColorBlockDim && a.W.P1P3Orthogonal && a.W.P1PlusP3CompletesW && a.W.BMinusLTraceZero && a.W.BMinusLRestActionOnW && containsAll(a.W.Supports, []string{SupportBMinusLActsOnRestW}), Detail: FormatW(a.W)},
			{Name: "formulate socket compression candidate with 3+4 rank anatomy", Passed: a.Compression.SocketRank == 1 && a.Compression.TopRank == 3 && a.Compression.RestRank == 4 && a.Compression.AggregateRank == 7 && a.Compression.MatchesI3PlusW && a.Compression.NonCircular && !a.Compression.UsesObservedData && !a.Compression.IsTheorem && containsAll(a.Compression.Supports, []string{SupportSocketCompressionCandidate, SupportTopAsEtTensorP3IfSelectorExists, SupportRestAsErTensorWIfSelectorExists, SupportRanksThreePlusFour, SupportSevenCompressionNotK7}) && containsAll(a.Compression.Failures, []string{FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureNoSocketPairCompressionMap, FailureCompressionCandidateNotTheorem}), Detail: FormatCompression(a.Compression)},
			{Name: "audit missing fine socket selectors", Passed: a.Sockets.RankOneSocketProjectorsPossible && !a.Sockets.FineSocketProjectorsCertified && a.Sockets.SocketAtomsBasisDependentWithoutSelector && !a.Sockets.DominantSelectorCertified && !a.Sockets.RestSelectorCertified && !a.Sockets.EtErCanonical && len(a.Sockets.PotentialSourcesAudited) == 5 && containsAll(a.Sockets.Failures, []string{FailureNoFineSocketProjectors, FailureSocketAtomsBasisDependent, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureEtErNotCanonical, FailureNoDFOrHiggsSocketSelector, FailureNoBoundaryRestSocketSelector}), Detail: FormatSockets(a.Sockets)},
			{Name: "classify aggregate operator as possible trace shadow only", Passed: a.Shadow.TopIsIdentityI3Candidate && a.Shadow.RestUsesBMinusLTransferOnW && !a.Shadow.AlphaDerivedByCompression && !a.Shadow.TraceMagnitudeReadoutCertified && !a.Shadow.AggregateOperatorIsSectorLedger && !a.Shadow.R3 && !a.Shadow.R4 && containsAll(a.Shadow.Supports, []string{SupportAggregateAsTraceShadowIfSelectors, SupportBMinusLActsOnRestW}) && containsAll(a.Shadow.Failures, []string{FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureAlphaStillSealed, FailureNoAlphaDerivation, FailureNotR3, FailureNotR4}), Detail: FormatShadow(a.Shadow)},
			{Name: "preserve ledger freeze and R2++ classification", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaIsNative && a.Impact.CompressionCandidateFormulated && a.Impact.SelectorsMissing && a.Impact.CompressionMapMissing && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.Classification, "socket-compression") && containsAll(a.Impact.Failures, []string{FailureNoSocketPairCompressionMap, FailureNoTraceMagnitudeReadout, FailureAlphaStillSealed, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoParticleAssignment, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4}), Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 839 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NoFineSocketProjectors && a.Firewalls.SocketAtomsBasisDependent && a.Firewalls.NoDominantSelector && a.Firewalls.NoRestSelector && a.Firewalls.NoCompressionMap && a.Firewalls.CompressionCandidateNotTheorem && a.Firewalls.NoDFOrHiggsSelector && a.Firewalls.NoBoundaryRestSelector && a.Firewalls.AggregateNotCompressionTheorem && a.Firewalls.SevenNotK7 && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.CompressionNotYukawaMagnitude && a.Firewalls.AlphaSealed && a.Firewalls.NoAlphaDerivation && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoParticleAssignment && a.Firewalls.NoThreeGeneration && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate839, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatBody(a.Body), FormatW(a.W), FormatSockets(a.Sockets), FormatCompression(a.Compression), FormatShadow(a.Shadow), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
