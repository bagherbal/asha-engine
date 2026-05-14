package orientationtruechirality

import (
	"fmt"
	"strings"
)

func FormatOrientation(a OrientationOperatorAudit) string {
	return fmt.Sprintf("candidate=%q source=%q actsOnSC=%t dimC=%d dimR=%d volume=%t equivGamma=%t commuteGamma=%t antiGamma=%t distinct=%t distinctEig=%t eig(+/-)=%d/%d adjusted=%t verdict=%s", a.CandidateName, a.Source, a.ActsOnSC, a.DimensionComplex, a.DimensionReal, a.VolumeElementAvailable, a.EquivalentToGamma, a.CommutesWithGamma, a.AntiCommutesWithGamma, a.DistinctFromGamma, a.DistinctEigenspaces, a.EvenEigenDimC, a.OddEigenDimC, a.ManualSignAdjusted, a.Verdict)
}

func FormatTauEta(a TauEtaPullbackAudit) string {
	return fmt.Sprintf("tauEta=%v flipped=%v scalarFunctional=%t endomorphismSC=%t pullback=%t gaugeProjection=%t chiralityOperator=%t verdict=%s", a.TauEtaDegrees, a.OrientationFlippedDegrees, a.FunctionalOnScalarBundle, a.EndomorphismOnSC, a.CanonicalPullbackDerived, a.GaugeProjectionMapDerived, a.CanActAsChiralityOperator, a.Verdict)
}

func FormatComparison(a GradingComparisonAudit) string {
	return fmt.Sprintf("gamma=%q chi=%q gamma(e/o)=%d/%d chi(+/-)=%d/%d sameSpectrum=%t sameEig=%t commute=%t anti=%t physical=%t verdict=%s", a.GammaName, a.ChiName, a.GammaEvenDimC, a.GammaOddDimC, a.ChiPlusDimC, a.ChiMinusDimC, a.SameSpectrum, a.SameEigenspaces, a.OperatorsCommute, a.OperatorsAntiCommute, a.PhysicalChiralityDerived, a.Verdict)
}

func FormatPlanes(rows []PlaneChiAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s class=%s modes=%v D(+/-)=%d/%d S(+/-)=%d/%d uniformD=%t uniformS=%t preservesChi=%t oneChi=%t selected=%t", p.Plane, p.PlaneClass, p.ModeIndices, p.DoubletPlusDimC, p.DoubletMinusDimC, p.SingletPlusDimC, p.SingletMinusDimC, p.DoubletsUniformChi, p.SingletsUniformChi, p.SU2PreservesChi, p.SU2ActsOnlyOnOneChi, p.SelectedByChi))
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(a ChiPlaneSieveAudit) string {
	return fmt.Sprintf("planes=%d uniformD=%d uniformS=%d selected=%v chiBreaks=%t allSame=%t temporalSpatialClasses=%t verdict=%s", a.CandidatePlanes, a.UniformDoubletPlanes, a.UniformSingletPlanes, a.SelectedPlanes, a.ChiBreaksDegeneracy, a.AllPlanesSameCounts, a.TemporalSpatialClasses, a.Verdict)
}

func FormatWeak(a WeakOutcomeAudit) string {
	return fmt.Sprintf("gate238Failed=%t chiImproves=%t tauEtaOperator=%t uniquePlane=%t physicalLeft=%t globalH=%t verdict=%s", a.Gate238GammaSelectorFailed, a.VolumeChiImprovesGate238, a.TauEtaSuppliesOperator, a.UniqueWeakPlaneSelected, a.PhysicalLeftHandedActionDerived, a.GlobalHSummandDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("adjustedChi=%t importedGamma5=%t importedConnesChi=%t forcedPlane=%t promotedTauEta=%t claimedLeft=%t claimedH=%t polluted=%t verdict=%s", a.AdjustedChiSignsToFit, a.ImportedSMGamma5, a.ImportedConnesChirality, a.ForcedWeakPlane, a.PromotedTauEtaToOperator, a.ClaimedLeftHandedAction, a.ClaimedGlobalH, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("volume=%t chiDistinct=%t tauPullback=%t uniformDoublets=%t chiSelects=%t physical=%t globalH=%t status=%s next=%q comment=%q", a.VolumeOrientationAvailable, a.ChiDistinctFromGamma, a.TauEtaPullbackDerived, a.UniformChiDoublets, a.ChiSelectsPlane, a.PhysicalChiralityDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
