package resolventcubictagselector

import (
	"fmt"
	"strings"
)

func FormatRoots(rs []QuarticRoot) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, fmt.Sprintf("%s interval=%s approx=%.10f mapped=%t sector=%q verdict=%s", r.Label, r.Interval, r.Approx, r.SectorMapped, r.Sector, r.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatResolvent(a ResolventAudit) string {
	return fmt.Sprintf("quartic=%q resolvent=%q coeffs=%v roots={%s} branches=%d 2+2=%t irreducible=%t priorSelected=%t verdict=%s", a.QuarticPolynomial, a.ResolventIntegerPolynomial, a.ResolventMonicCoefficients, FormatRoots(a.QuarticRoots), len(a.Branches), a.EncodesTwoPlusTwo, a.IrreducibleOverQ, a.CanonicalRootPreviouslySelected, a.Verdict)
}

func FormatTag(t TopologicalTag) string {
	return fmt.Sprintf("%s source=%s actsOn=%q data=%q rule=%q sealed=%t rootOperator=%t verdict=%s", t.Name, t.SourceGate, t.ActsOn, t.NativeData, t.SelectionRule, t.IsSealed, t.DerivedAsOperatorOnQuarticRoots, t.Verdict)
}

func FormatTags(a TagAudit) string {
	parts := make([]string, 0, len(a.Tags))
	for _, t := range a.Tags {
		parts = append(parts, FormatTag(t))
	}
	return fmt.Sprintf("tauUD=%t bGapNu=%t sectorReach=%t rootReach=%t observed=%t tags={%s} verdict=%s", a.TauEtaBindsUD, a.BGapTagsNeutrino, a.TagsReachSectorLabels, a.TagsReachQuarticRoots, a.UsesObservedMasses, strings.Join(parts, "; "), a.Verdict)
}

func FormatPairing(b ResolventBranch) string {
	return fmt.Sprintf("%s roots=%s sectors=%s selected=%t elimTau=%t elimB=%t rootMap=%t contactRoot=%t verdict=%s", b.Label, b.PairingOnRoots, b.PairingOnSectors, b.SelectedByTags, b.EliminatedByTauEta, b.EliminatedByBGap, b.RequiresRootSectorMap, b.ContactRootSelected, b.Verdict)
}

func FormatSieve(a PairingSieve) string {
	parts := make([]string, 0, len(a.CandidatePairings))
	for _, b := range a.CandidatePairings {
		parts = append(parts, FormatPairing(b))
	}
	return fmt.Sprintf("total=%d survivors=%d eliminated=%d selected=%s uniqueSector=%t uniqueContactRoot=%t pairings={%s} verdict=%s", a.TotalCandidates, a.SurvivingSectorPairings, a.EliminatedCandidates, a.SelectedSectorPairing, a.UniqueSectorPairing, a.UniqueContactRoot, strings.Join(parts, "; "), a.Verdict)
}

func FormatGate275Branch(b Gate275Branch) string {
	return fmt.Sprintf("%s r=%s≈%.15g |y/x|≈%.15g map=%q selected=%t verdict=%s", b.Name, b.ExactR, b.R, b.AbsYOverX, b.ContactPairingMap, b.Selected, b.Verdict)
}

func FormatProjection(a BranchProjectionAudit) string {
	parts := make([]string, 0, len(a.BranchesInherited))
	for _, b := range a.BranchesInherited {
		parts = append(parts, FormatGate275Branch(b))
	}
	return fmt.Sprintf("selectedSector=%s rootSelected=%t rootToR=%t uniqueR=%t selectedR=%s branches={%s} verdict=%s", a.SelectedSectorPairing, a.ResolventRootSelected, a.ResolventRootToRBranchMap, a.UniqueRBranchSelected, a.SelectedRBranch, strings.Join(parts, "; "), a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("mass=%t ckm=%t empiricalY=%t arbitraryRootMap=%t sectorNoOver=%t contactNoOver=%t rNoOver=%t higgsNoClaim=%t polluted=%t verdict=%s", !a.NoObservedMassesUsed, !a.NoCKMPMNSUsed, !a.NoEmpiricalYukawaInserted, !a.NoArbitraryRootSectorMap, a.SectorPairingNotOverpromoted, a.ContactRootNotOverpromoted, a.RBranchNotOverpromoted, a.HiggsRatioNotClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	parts := make([]string, 0, len(a.Criteria))
	for _, c := range a.Criteria {
		parts = append(parts, fmt.Sprintf("%s[required=%t satisfied=%t detail=%s]", c.Name, c.Required, c.Satisfied, c.Detail))
	}
	return fmt.Sprintf("rootMap=%t projectors=%t resolventToR=%t branch=%t heat=%t criteria={%s} next=%s verdict=%s", a.NeedRootSectorMap, a.NeedContactProjectors, a.NeedResolventToRMap, a.NeedBranchSelector, a.NeedHeatKernelMap, strings.Join(parts, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("resolvent=%t tags=%t sector=%t contact=%t r=%t higgs=%t firewall=%t next=%s status=%s comment=%q", a.ResolventRetrieved, a.TagsApplied, a.UniqueSectorPairing, a.UniqueContactRoot, a.UniqueAmplitudeBranch, a.HiggsRatioDerived, a.FirewallPreserved, a.NextGate, a.Status, a.Comment)
}
