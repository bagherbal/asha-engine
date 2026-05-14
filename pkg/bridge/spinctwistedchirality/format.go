package spinctwistedchirality

import (
	"fmt"
	"strings"
)

func FormatU1(a U1GeneratorAudit) string {
	return fmt.Sprintf("name=%q source=%q actsOnSC=%t weights=%v temporal=%.6g spatial=%.6g spatialDegenerate=%t importedSMHypercharge=%t nativeAsHypercharge=%t verdict=%s", a.Name, a.Source, a.ActsOnSC, a.ModeWeights, a.TemporalWeight, a.SpatialWeight, a.SpatialWeightsDegenerate, a.ImportedSMHypercharge, a.NativeContactU1AsHypercharge, a.Verdict)
}

func FormatTwist(a TwistOperatorAudit) string {
	return fmt.Sprintf("name=%q formula=%q commutes=%t diagonal=%t involution=%t zeroEig=%t distinctGamma=%t physicalChirality=%t manualFit=%t verdict=%s", a.Name, a.Formula, a.GammaCommutesWithY, a.IsDiagonalOnFockBasis, a.IsInvolution, a.HasZeroEigenvalues, a.DistinctFromGamma, a.PhysicalChiralityDerived, a.ManualHyperchargeFit, a.Verdict)
}

func FormatPlanes(rows []PlaneTwistAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s class=%s modes=%v weights=%v preservesY=%t comm=%.6g Dtwist=%v Stwist=%v uniformD=%t uniformS=%t selected=%t", p.Plane, p.PlaneClass, p.ModeIndices, p.PlaneModeWeights, p.SU2PreservesY, p.U1CommutatorResidual, p.DoubletTwistEigenvalues, p.SingletTwistEigenvalues, p.DoubletsUniformTwist, p.SingletsUniformTwist, p.SelectedByTwist))
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(a TwistedSieveAudit) string {
	return fmt.Sprintf("planes=%d preserving=%v rejected=%v uniformD=%v selected=%v temporalRejected=%t pureSpatialRemain=%d twistBreaks=%t verdict=%s", a.CandidatePlanes, a.U1PreservingPlanes, a.U1RejectedPlanes, a.UniformDoubletPlanes, a.SelectedPlanes, a.TemporalSpatialRejected, a.PureSpatialDegeneracy, a.TwistBreaksDegeneracy, a.Verdict)
}

func FormatWeak(a WeakOutcomeAudit) string {
	return fmt.Sprintf("gate239Failed=%t classSieveImproved=%t uniquePlane=%t physicalLeft=%t globalH=%t orderOne=%t verdict=%s", a.Gate239ChiFailed, a.U1TwistImprovesClassSieve, a.UniqueWeakPlaneSelected, a.PhysicalLeftHandedDerived, a.GlobalHSummandDerived, a.OrderOneReady, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("importedSMY=%t tunedWeights=%t forcedPlane=%t forcedLeft=%t importedSpinC=%t claimedH=%t claimedOrderOne=%t polluted=%t verdict=%s", a.ImportedSMHypercharge, a.TunedU1Weights, a.ForcedWeakPlane, a.ForcedLeftHandedAction, a.ImportedSpinCStructure, a.ClaimedGlobalH, a.ClaimedOrderOne, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("nativeU1=%t twist=%t temporalRejected=%t pureSpatialRemain=%d uniformDoublets=%t uniquePlane=%t physicalChirality=%t globalH=%t status=%s next=%q comment=%q", a.NativeU1Available, a.TwistConstructed, a.U1RejectsTemporalPlanes, a.PureSpatialPlanesRemain, a.UniformTwistedDoublets, a.UniqueWeakPlaneDerived, a.PhysicalChiralityDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
