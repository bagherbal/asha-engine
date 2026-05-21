package generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "generation2.right.socket.character.split.punctured.leptocolor.compression.audit"
	theoremName = "Gate 840 — RightSocket Character Split and Punctured LeptoColor Compression Audit"
)

func Generation2RightSocketCharacterSplitPuncturedLeptoColorCompressionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 839 finite-body compression candidate", Passed: a.Body.Gate839Inherited && a.Body.Gate838BodyInherited && a.Body.HPartDim == HPartDim && a.Body.HFDim == HFDim && a.Body.RightSlotDim == RightSlotDim && containsAll(a.Body.Supports, []string{SupportGate839CompressionInherited}), Detail: FormatBody(a.Body)},
			{Name: "source-type right socket character split as seal", Passed: a.RightSplit.CharacterSplitSealAudited && a.RightSplit.CharacterProjectorsSourceTypedBySeal && a.RightSplit.UnorderedPairCertified && a.RightSplit.CharacterPairOrthogonal && a.RightSplit.CharacterPairComplete && !a.RightSplit.NativeDerivationCertified && !a.RightSplit.ExplicitRhoRMatrixCertified && !a.RightSplit.FullRhoFActionLedgerCertified && !a.RightSplit.DominantRestOrientationCertified && containsAll(a.RightSplit.Supports, []string{SupportRightCharacterSplitSeal, SupportEPlusEMinusAsCharacterProjectors, SupportRankOneSocketsNotArbitraryIfCharacters}) && containsAll(a.RightSplit.Failures, []string{FailureRightCharacterSplitSealNotNative, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureDominantRestOrientationMissing}), Detail: FormatRightSplit(a.RightSplit)},
			{Name: "reverify W carrier and B-L neutral quartet", Passed: a.W.Dim == WDim && a.W.P1Rank == LeptonBlockDim && a.W.P3Rank == ColorBlockDim && a.W.P1P3Orthogonal && a.W.P1PlusP3CompletesW && a.W.BMinusLTraceZeroOnW, Detail: FormatW(a.W)},
			{Name: "audit punctured 2x4 right lepto-color rectangle", Passed: a.Puncture.FullRank == 8 && a.Puncture.TopRank == 3 && a.Puncture.RestRank == 4 && a.Puncture.SelectedRank == 7 && a.Puncture.ExcludedRank == 1 && a.Puncture.RanksCloseRightRectangle && a.Puncture.UsesRightCharacterSockets && !a.Puncture.OrientationCertified && !a.Puncture.CompressionMapCertified && !a.Puncture.IsTheorem && containsAll(a.Puncture.Supports, []string{SupportPuncturedRectangleCandidate, SupportSelectedRankSeven, SupportExcludedSingletonRankOne}) && containsAll(a.Puncture.Failures, []string{FailureDominantRestOrientationMissing, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailurePunctureNotCompressionTheorem, FailureNoTypedCompressionMap}), Detail: FormatPuncture(a.Puncture)},
			{Name: "certify B-L puncture conservation pattern", Passed: a.BMinusL.TopTraceMatches && a.BMinusL.RestTraceMatches && a.BMinusL.SelectedTrace == 1 && a.BMinusL.ExcludedTrace == -1 && a.BMinusL.SelectedExcludedCancel && a.BMinusL.FullRectangleNeutral && a.BMinusL.PunctureConservationPattern && containsAll(a.BMinusL.Supports, []string{SupportBMinusLSelectedPlusOne, SupportBMinusLExcludedMinusOne, SupportFullRightRectangleNeutral}), Detail: FormatBMinusL(a.BMinusL)},
			{Name: "classify aggregate operator as oriented compression shadow only", Passed: a.Shadow.TopBlockLocatedInRightColorSocket && a.Shadow.RestBlockLocatedInRightLeptoColorSocket && a.Shadow.OrientationNeeded && !a.Shadow.CompressionMapCertified && !a.Shadow.AlphaDerivedByCompression && !a.Shadow.TraceMagnitudeReadoutCertified && !a.Shadow.R3 && !a.Shadow.R4 && containsAll(a.Shadow.Failures, []string{FailureNoDFOrHiggsOrientation, FailureNoBoundaryRestOrientation, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4}), Detail: FormatShadow(a.Shadow)},
			{Name: "preserve ledger freeze and R2++ classification", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaIsNative && a.Impact.FineSocketProblemPartiallyResolvedBySeal && a.Impact.OrientationStillMissing && a.Impact.CompressionMapStillMissing && a.Impact.BMinusLConservationPatternFound && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && strings.Contains(a.Impact.Classification, "punctured"), Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 840 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.CharacterSplitSealNotNative && a.Firewalls.NoExplicitRhoRProof && a.Firewalls.NoFullRhoFActionLedger && a.Firewalls.OrientationMissing && a.Firewalls.NoDominantSelector && a.Firewalls.NoRestSelector && a.Firewalls.PunctureNotCompressionTheorem && a.Firewalls.NoCompressionMap && a.Firewalls.ExcludedSingletonNotParticle && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoDFOrHiggsOrientation && a.Firewalls.NoBoundaryRestOrientation && a.Firewalls.NoAlphaDerivation && a.Firewalls.AlphaSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.CompressionNotYukawaMagnitude && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoThreeGeneration && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate840, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatBody(a.Body), FormatRightSplit(a.RightSplit), FormatW(a.W), FormatPuncture(a.Puncture), FormatBMinusL(a.BMinusL), FormatShadow(a.Shadow), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
