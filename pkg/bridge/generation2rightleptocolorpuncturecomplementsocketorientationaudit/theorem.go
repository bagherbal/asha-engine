package generation2rightleptocolorpuncturecomplementsocketorientationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "generation2.right.leptocolor.puncture.complement.socket.orientation.audit"
	theoremName = "Gate 841 — Right LeptoColor Puncture Complement and Socket-Orientation Audit"
)

func Generation2RightLeptoColorPunctureComplementSocketOrientationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 840 punctured right rectangle", Passed: a.Carrier.Gate838BodyInherited && a.Carrier.Gate839CompressionInherited && a.Carrier.Gate840PunctureInherited && a.Carrier.RightRectangleRank == RightRectangleRank && containsAll(a.Carrier.Supports, []string{SupportGate840PunctureInherited}), Detail: FormatCarrier(a.Carrier)},
			{Name: "certify 8=7+1 puncture complement anatomy", Passed: a.Complement.SupportAnatomyCertified && !a.Complement.CompressionTheoremCertified && a.Complement.ActiveRank == 7 && a.Complement.PunctureRank == 1 && a.Complement.FullRank == 8 && a.Complement.DominantColorRank == 3 && a.Complement.RestQuartetRank == 4 && a.Complement.Orthogonal && a.Complement.Complete && a.Complement.EightEqualsSevenPlusOne && containsAll(a.Complement.Supports, []string{SupportRightRectangleComplement, SupportActiveSupportRankSeven, SupportPunctureSingletonRankOne}) && containsAll(a.Complement.Failures, []string{FailureCompressionCandidateNotTheorem, FailureNoTypedCompressionMap, FailureNoAggregateCompressionMap}), Detail: FormatComplement(a.Complement)},
			{Name: "certify B-L compensating puncture pattern", Passed: a.BMinusL.DominantColorTrace == 1 && a.BMinusL.RestQuartetTrace == 0 && a.BMinusL.ActiveTrace == 1 && a.BMinusL.PunctureTrace == -1 && a.BMinusL.FullTrace == 0 && a.BMinusL.ActivePlusPunctureCancel && a.BMinusL.FullNeutral && a.BMinusL.CompensatingPuncturePattern && containsAll(a.BMinusL.Supports, []string{SupportBMinusLActivePlusOne, SupportBMinusLPunctureMinusOne, SupportBMinusLFullRightRectangleNeutral, SupportPunctureIsCompensatingSingleton}), Detail: FormatBMinusL(a.BMinusL)},
			{Name: "audit sterile/null-edge puncture without promotion", Passed: a.Sterile.RightSocket && a.Sterile.Leptonic && a.Sterile.Colorless && a.Sterile.ExcludedFromActive && a.Sterile.Rank == 1 && a.Sterile.BMinusLTrace == -1 && !a.Sterile.DFEdgeDataAvailable && !a.Sterile.NullEdgeCertified && !a.Sterile.SterilePunctureCertified && !a.Sterile.PhysicalParticleAssignmentCertified && containsAll(a.Sterile.Supports, []string{SupportSterileNullEdgeCandidate}) && containsAll(a.Sterile.Failures, []string{FailureNoDFEdgeData, FailureNoNullEdgeTheorem, FailureNoSterilePunctureTheorem, FailurePunctureNotPhysicalParticle, FailureNoRightNeutrinoTheorem}), Detail: FormatSterile(a.Sterile)},
			{Name: "audit dominant/rest socket orientation without certification", Passed: a.Orientation.DominantWouldSourceI3 && a.Orientation.RestWouldCarryW && a.Orientation.SameSocketContainsDominantColorAndPuncture && !a.Orientation.DominantOrientationCertified && !a.Orientation.RestOrientationCertified && !a.Orientation.DFOrHiggsSelectorCertified && !a.Orientation.BoundaryRestSelectorCertified && !a.Orientation.OrientationMapCertified && containsAll(a.Orientation.Supports, []string{SupportDominantColorOrientationCandidate, SupportRestQuartetOrientationCandidate}) && containsAll(a.Orientation.Failures, []string{FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoDFOrHiggsOrientationSelector, FailureNoBoundaryRestOrientationSelector, FailureNoTypedSocketOrientationMap}), Detail: FormatOrientation(a.Orientation)},
			{Name: "classify aggregate finite-body location as candidate only", Passed: a.Location.FiniteBodyLocationCandidate && a.Location.TopBlockLocatedIfOrientation && a.Location.RestBlockLocatedIfOrientation && !a.Location.CompressionMapCertified && !a.Location.AlphaDerivedByCompression && !a.Location.TraceMagnitudeReadoutCertified && !a.Location.R3 && !a.Location.R4 && containsAll(a.Location.Failures, []string{FailureNoTypedCompressionMap, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoAggregateCompressionMap, FailureNotR3, FailureNotR4}), Detail: FormatLocation(a.Location)},
			{Name: "preserve ledger freeze and R2++ classification", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlus && !a.Ledger.R3 && !a.Ledger.R4 && !a.Ledger.AlphaIsNative && a.Impact.PunctureComplementLawCertified && a.Impact.BMinusLCompensationFound && a.Impact.SterileNullEdgeStillUncertified && a.Impact.OrientationStillMissing && a.Impact.CompressionMapStillMissing && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && containsText([]string{a.Impact.Classification}, "8=7+1"), Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 841 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.RightCharacterSplitStillSeal && a.Firewalls.NoFullRhoFActionLedger && a.Firewalls.NoDFEdgeData && a.Firewalls.NoNullEdgeTheorem && a.Firewalls.NoSterilePunctureTheorem && a.Firewalls.PunctureNotPhysicalParticle && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoDominantOrientationTheorem && a.Firewalls.NoRestOrientationTheorem && a.Firewalls.NoDFOrHiggsOrientationSelector && a.Firewalls.NoBoundaryRestOrientationSelector && a.Firewalls.NoTypedSocketOrientationMap && a.Firewalls.NoTypedCompressionMap && a.Firewalls.CompressionCandidateNotTheorem && a.Firewalls.NoAlphaDerivation && a.Firewalls.AlphaSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.CompressionNotYukawaMagnitude && a.Firewalls.NoAggregateCompressionMap && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NoObservedYukawaFit && a.Firewalls.NoThreeGeneration && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate841, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatCarrier(a.Carrier), FormatComplement(a.Complement), FormatBMinusL(a.BMinusL), FormatSterile(a.Sterile), FormatOrientation(a.Orientation), FormatLocation(a.Location), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
