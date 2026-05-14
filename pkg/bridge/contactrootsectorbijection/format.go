package contactrootsectorbijection

import (
	"fmt"
	"strings"
)

func FormatRoot(r ContactRoot) string {
	return fmt.Sprintf("%s interval=%s approx=%.10f rank=%d dist0=%.10f suppressed=%t closest=%s distRat=%.10f magSelect=%t verdict=%s", r.Label, r.Interval, r.Approx, r.OrderRank, r.DistanceToZero, r.NearZeroSuppressed, r.ClosestSimpleRational, r.DistanceToSimpleRational, r.CanBeSelectedByMagnitude, r.Verdict)
}

func FormatMagnitude(a RootMagnitudeAudit) string {
	parts := make([]string, 0, len(a.Roots))
	for _, r := range a.Roots {
		parts = append(parts, FormatRoot(r))
	}
	return fmt.Sprintf("poly=%q roots={%s} O1=%t null=%t orderAvail=%t orderInvariant=%t magMap=%t bgapComparable=%t verdict=%s", a.Polynomial, strings.Join(parts, "; "), a.AllRootsO1, a.AnyNativeNullRoot, a.OrderingAvailable, a.OrderingInvariant, a.MagnitudeBijectionDerived, a.BGapScaleComparableToRoots, a.Verdict)
}

func FormatConstraint(c Constraint) string {
	return fmt.Sprintf("%s source=%s data=%q actsOn=%q labelRoot=%t selectPair=%t extraFunctor=%t verdict=%s", c.Name, c.SourceGate, c.Data, c.ActsOn, c.CanLabelRoot, c.CanSelectPairing, c.RequiresExtraFunctor, c.Verdict)
}

func FormatConstraints(a ConstraintAudit) string {
	parts := make([]string, 0, len(a.Constraints))
	for _, c := range a.Constraints {
		parts = append(parts, FormatConstraint(c))
	}
	return fmt.Sprintf("morita=%t bgap=%t tau=%t sectorReach=%t rootReach=%t observed=%t constraints={%s} verdict=%s", a.MoritaMultiplicityAvailable, a.BGapTagAvailable, a.TauEtaPairingAvailable, a.ConstraintsReachSectors, a.ConstraintsReachRoots, a.UsesObservedMasses, strings.Join(parts, "; "), a.Verdict)
}

func FormatProjectors(a ProjectorAudit) string {
	return fmt.Sprintf("quarticIrred=%t resolventIrred=%t rootProjQ=%t pairProjQ=%t splitting=%t resolventAdj=%t rationalProj=%t verdict=%s", a.QuarticIrreducibleOverQ, a.ResolventIrreducibleOverQ, a.IndividualRootProjectorsOverQ, a.TwoPlusTwoPairProjectorsOverQ, a.RequiresSplittingField, a.RequiresResolventRootAdjunction, a.RationalContactProjectorDerived, a.Verdict)
}

func FormatPairing(p PairingCandidate) string {
	return fmt.Sprintf("%s roots=%s sums=(%.10f,%.10f) products=(%.10f,%.10f) meanGap=%.10f compatible=%t multSelect=%t bgapSelect=%t selected=%t verdict=%s", p.Label, p.RootPairing, p.PairSums[0], p.PairSums[1], p.PairProducts[0], p.PairProducts[1], p.PairMeanGap, p.CompatibleWithUD_ENU, p.SelectedByMultiplicity, p.SelectedByBGap, p.SelectedAsResolventRoot, p.Verdict)
}

func FormatPairings(a PairingAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, p := range a.Candidates {
		parts = append(parts, FormatPairing(p))
	}
	return fmt.Sprintf("total=%d compatible=%d selected=%d unique=%t selectedRoot=%q pairings={%s} verdict=%s", a.TotalCandidates, a.CompatibleWithSectorSplit, a.SelectedPairings, a.UniqueRootPairing, a.SelectedRootPairing, strings.Join(parts, "; "), a.Verdict)
}

func FormatBijection(a BijectionAudit) string {
	return fmt.Sprintf("total=%d afterSector=%d afterBGap=%d afterTau=%d unique=%t assignment=%v verdict=%s", a.TotalRootSectorBijections, a.BijectionsAfterSectorPairing, a.BijectionsAfterBGapNuTag, a.BijectionsAfterUDTauTag, a.UniqueBijection, a.DerivedAssignment, a.Verdict)
}

func FormatBranchProjection(a BranchProjectionAudit) string {
	return fmt.Sprintf("rPlus=%.15g |y/x|Plus=%.15g rMinus=%.15g |y/x|Minus=%.15g resolventRoot=%t map=%t unique=%t selected=%q verdict=%s", a.RPlus, a.AbsYOverXPlus, a.RMinus, a.AbsYOverXMinus, a.ResolventRootSelected, a.RootPairingToRBranchMap, a.UniqueAmplitudeBranch, a.SelectedBranch, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("mass=%t ckm=%t empiricalY=%t orderPromotion=%t arbitraryMap=%t bgapRootMap=%t multAmp=%t higgsNoClaim=%t polluted=%t verdict=%s", !a.NoObservedMassesUsed, !a.NoCKMPMNSUsed, !a.NoEmpiricalYukawaInserted, !a.NoRootOrderingPromotion, !a.NoArbitraryRootSectorMap, !a.NoBGapScaleToRootMap, !a.NoMultiplicityToAmplitude, a.NoHiggsRatioClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	parts := make([]string, 0, len(a.Criteria))
	for _, c := range a.Criteria {
		parts = append(parts, fmt.Sprintf("%s[required=%t satisfied=%t detail=%s]", c.Name, c.Required, c.Satisfied, c.Detail))
	}
	return fmt.Sprintf("projector=%t rootMap=%t rMap=%t JY=%t heat=%t criteria={%s} next=%s verdict=%s", a.NeedContactProjectorAction, a.NeedRootSectorBijection, a.NeedResolventToRBranchMap, a.NeedPhysicalJHypercharge, a.NeedHeatKernelNormalization, strings.Join(parts, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("roots=%t constraints=%t projectors=%t rootPair=%t rootSector=%t r=%t higgs=%t firewall=%t next=%s status=%s comment=%q", a.RootsRetrieved, a.ConstraintsAudited, a.ProjectorSemanticsFound, a.UniqueRootPairing, a.UniqueRootSectorMap, a.AmplitudeBranchLocked, a.HiggsRatioDerived, a.FirewallPreserved, a.NextGate, a.Status, a.Comment)
}
