package generation2socketcharacteridentificationedgeintertwinerpromotionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SOCKET_CHARACTER_IDENTIFICATION_EDGE_INTERTWINER_PROMOTION_AUDIT"
	theoremName = "Gate 862 — SocketCharacter Identification and Edge-Intertwiner Promotion Audit"
)

func Generation2SocketCharacterIdentificationEdgeIntertwinerPromotionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "audit right and oriented weak character ledgers", Passed: charactersOK(a.Characters) && containsAll(a.Characters[0].Supports, []string{StatusRightCharacterLedger}) && containsAll(a.Characters[2].Supports, []string{StatusWeakCharacterLedger}), Detail: FormatCharacters(a.Characters)},
			{Name: "define C_R to C_H character identification as orientation seal", Passed: a.IDMap.Defined && a.IDMap.OrientationSeal && !a.IDMap.Native && !a.IDMap.OperatorCertified && containsAll(a.IDMap.Supports, []string{StatusCharacterMapFormulated, SupportSocketMatchOrientation}) && containsAll(a.IDMap.Failures, []string{FailureCRToCHNotNative, FailureSocketMatchSeal}), Detail: FormatIdentification(a.IDMap)},
			{Name: "formulate and satisfy active edge intertwiner equations given the seal", Passed: activeIntertwinersOK(a.Edges) && containsAll(a.IDMap.Supports, []string{SupportEdgeIntertwinersGivenID}), Detail: FormatEdges(a.Edges)},
			{Name: "preserve puncture edge zero under character identification", Passed: punctureOK(a.Edges) && a.Kernel.PunctureZero && a.Kernel.PunctureNotForcedByCharacterID && containsAll(a.Kernel.Supports, []string{StatusPunctureZeroReaudited, SupportPunctureNotForced}), Detail: FormatKernel(a.Kernel) + " | " + FormatEdges(a.Edges)},
			{Name: "sharpen stabilizer first-order compatibility but keep it sealed", Passed: a.FirstOrder.Gate861Inherited && a.FirstOrder.ColorCentralityInstalled && a.FirstOrder.CharacterIdentificationNeeded && a.FirstOrder.CharacterIdentificationSealed && a.FirstOrder.StabilizerOperatorCompatibilitySharpened && !a.FirstOrder.StabilizerOperatorCompatibilityNative && containsAll(a.FirstOrder.Supports, []string{StatusGate861Inherited, StatusFirstOrderSharpened, SupportFirstOrderSharpened}) && containsAll(a.FirstOrder.Failures, []string{FailureSocketMatchSeal, FailureNoFullUnbrokenAF}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "block native character theorem, magnitudes, R3/R4, and ledger updates", Passed: !a.Impact.NativeCharacterTheorem && !a.Impact.NativeFiniteTripleProof && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 862 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.CRToCHNotNative && a.Firewalls.SocketMatchSeal && a.Firewalls.NoFullUnbrokenAF && a.Firewalls.AForientNotFullAF && a.Firewalls.NoNativeFiniteTriple && a.Firewalls.NoFullOperatorFirstOrder && a.Firewalls.NoCompleteJProof && a.Firewalls.NoBimoduleProof && a.Firewalls.NoYukawaMagnitudes && a.Firewalls.NoNumericalYukawa && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenTheorem && a.Firewalls.NoR3 && a.Firewalls.NoR4 && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatCharacters(a.Characters), FormatIdentification(a.IDMap), FormatEdges(a.Edges), FormatFirstOrder(a.FirstOrder), FormatKernel(a.Kernel), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func charactersOK(chars []SocketCharacter) bool {
	seen := map[string]bool{}
	for _, c := range chars {
		seen[c.Name] = true
		if !c.Typed || c.Rank != 1 || c.Native {
			return false
		}
		if c.Algebra == "C_H" && !c.OrientationRelative {
			return false
		}
	}
	return seen["chi_R^+"] && seen["chi_R^-"] && seen["chi_H^+"] && seen["chi_H^-"]
}

func activeIntertwinersOK(edges []EdgeIntertwiner) bool {
	seen := map[string]bool{}
	for _, e := range edges {
		seen[e.Name] = true
		if e.PunctureEdge {
			continue
		}
		if !e.HoldsGivenIdentification || e.OperatorCertified || e.NumericalValue || e.YukawaMagnitude {
			return false
		}
		if e.Name == "Y_+3" && (e.DomainCharacter != "chi_R^+" || e.CodomainCharacter != "chi_H^+" || !e.ColorCentral) {
			return false
		}
		if e.Name == "Y_-3" && (e.DomainCharacter != "chi_R^-" || e.CodomainCharacter != "chi_H^-" || !e.ColorCentral) {
			return false
		}
		if e.Name == "Y_-1" && (e.DomainCharacter != "chi_R^-" || e.CodomainCharacter != "chi_H^-" || !e.LeptonTrivial) {
			return false
		}
	}
	return seen["Y_+3"] && seen["Y_-3"] && seen["Y_-1"] && seen["Y_+1"]
}

func punctureOK(edges []EdgeIntertwiner) bool {
	for _, e := range edges {
		if e.PunctureEdge {
			return e.Name == "Y_+1" && e.EdgeExpression == "0" && !e.PunctureForced && e.HoldsGivenIdentification && !e.OperatorCertified
		}
	}
	return false
}
