package chiralweakselector

import (
	"fmt"
	"strings"
)

func FormatGamma(a GammaAudit) string {
	return fmt.Sprintf("grading=%q evenDimC=%d oddDimC=%d fromGate233=%t equatedSMChirality=%t physicalChirality=%t verdict=%s", a.Grading, a.EvenStateDimC, a.OddStateDimC, a.RetrievedFromGate233, a.EquatedToSMChirality, a.PhysicalChiralityDerived, a.Verdict)
}

func FormatPlanes(rows []PlaneChiralityAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s class=%s modes=%v D(e/o)=%d/%d S(e/o)=%d/%d uniformD=%t uniformS=%t preservesGamma=%t oneParity=%t", p.Plane, p.PlaneClass, p.ModeIndices, p.DoubletEvenDimC, p.DoubletOddDimC, p.SingletEvenDimC, p.SingletOddDimC, p.DoubletsUniformParity, p.SingletsUniformParity, p.SU2PreservesGamma, p.SU2ActsOnlyOnOneParity))
	}
	return strings.Join(parts, " | ")
}

func FormatSieve(a DegeneracySieveAudit) string {
	return fmt.Sprintf("planes=%d uniformD=%d uniformS=%d selected=%v gammaBreaks=%t allSameCounts=%t verdict=%s", a.CandidatePlanes, a.UniformDoubletPlanes, a.UniformSingletPlanes, a.ChiralSelectedPlanes, a.GammaBreaksDegeneracy, a.AllPlanesSameCounts, a.Verdict)
}

func FormatTemporal(a TemporalSpatialAudit) string {
	return fmt.Sprintf("temporalSpatial=%d %v pureSpatial=%d %v classDistinction=%t uniquePlane=%t verdict=%s", a.TemporalSpatialPlaneCount, a.TemporalSpatialPlanes, a.PureSpatialPlaneCount, a.PureSpatialPlanes, a.ClassDistinctionExists, a.UniquePlaneSelected, a.Verdict)
}

func FormatWeak(a WeakActionAudit) string {
	return fmt.Sprintf("localHInherited=%t gammaSelector=%t temporalSelector=%t contactMap=%t hyperchargeColor=%t globalH=%t physicalLeft=%t verdict=%s", a.CandidateLocalHSupportInherited, a.GammaSelectorWorks, a.TemporalSpatialSelectorWorks, a.ContactSU2PlaneMapDerived, a.HyperchargeColorAttachment, a.GlobalHSummandDerived, a.PhysicalLeftHandedActionDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedLeft=%t forcedPlane=%t importedSMChirality=%t importedPauli=%t importedConnes=%t claimedH=%t claimedOrderOne=%t polluted=%t verdict=%s", a.ForcedLeftHandedAssignment, a.ForcedWeakPlane, a.ImportedSMChirality, a.ImportedPauliMatrices, a.ImportedConnesAlgebra, a.ClaimedGlobalH, a.ClaimedOrderOne, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gamma=%t uniformDoublets=%t gammaSelects=%t temporalClasses=%t uniqueWeakPlane=%t physicalLeft=%t globalH=%t status=%s next=%q comment=%q", a.GammaParityAvailable, a.UniformChiralDoublets, a.GammaSelectsPlane, a.TemporalSpatialClasses, a.UniqueWeakPlaneDerived, a.PhysicalLeftActionDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
