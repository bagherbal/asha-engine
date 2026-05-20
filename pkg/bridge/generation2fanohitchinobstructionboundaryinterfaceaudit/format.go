package generation2fanohitchinobstructionboundaryinterfaceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate654Inheritance) string {
	return fmt.Sprintf("sourced=%t pgForces=%t gauge=%t basisFree=%t ray=%q cos=%.15g rho2=%.15g splitG2=%t boundary=%t seven=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.InternalMechanismSourced, x.PGForcesFanoNormalForm, x.GaugeControlledSource, x.BasisFreeSourceTheorem, x.HitchinRay, x.CosTheta, x.RhoSquared, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.Gate654FirewallPreserved, x.Verdict)
}

func FormatInvariants(x InternalInvariantLedger) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s=%g [%s] gauge=%t orient=%t", r.Name, r.Value, r.Expression, r.GaugeInvariant, r.OrientationBound))
	}
	return fmt.Sprintf("traceSK=%.3g traceG=%.3g norm2SK=%.3g norm2G=%.3g detG=%.3g inner=%.15g rho2=%.15g ranks=%d|%d|%d so3=%d fano=%d channels=%q native=%t gaugeClass=%t boundaryData=%t verdict=%q rows=%s", x.TraceSK, x.TraceGUn, x.Norm2SK, x.Norm2GUn, x.DetGUn, x.ProjectiveInner, x.ObstructionSquare, x.RankPlus, x.RankMinus, x.RankK7, x.SO3GaugeDim, x.FanoTripleCount, x.ChannelCount, x.AllNativeFinite, x.AllGaugeClassified, x.BoundaryDataPresent, x.Verdict, strings.Join(parts, "; "))
}

func FormatSevenOver72(x SevenOver72InterfaceAudit) string {
	return fmt.Sprintf("weight=%.15g numerator=%s denominator=%q addsBeyondNumerator=%t boundaryPair=%t traceMap=%t structures7=%t theorem=%t verdict=%q", x.CandidateWeight, strings.Join(x.NumeratorSources, ","), x.DenominatorCandidate, x.FanoAddsBeyondNumerator, x.BoundaryPairSupplied, x.TraceMapSupplied, x.StructuresNumerator7, x.CertifiedSevenOver72Theorem, x.Verdict)
}

func FormatBoundaryStress(x BoundaryStressInterfaceAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s=%.15g closest=%s residual=%.3g rel=%.3g class=%q", r.Candidate, r.Value, r.ClosestSeal, r.AbsResidual, r.RelResidual, r.Classification))
	}
	return fmt.Sprintf("certified=%t nearOnly=%t noArbitrary=%t verdict=%q rows=%s", x.CertifiedBoundaryStressSource, x.NearBridgeClueOnly, x.NoArbitrarySearch, x.Verdict, strings.Join(parts, "; "))
}

func FormatHistoryLoop(x HistoryLoopUnitInterfaceAudit) string {
	return fmt.Sprintf("L=%.15g piS1=%t heat=%t angular=%t finiteOnly=%t certified=%t verdict=%q", x.TargetL, x.SuppliesPiOrS1, x.SuppliesHeatKernel, x.SuppliesAngularMeasure, x.FiniteAlgebraicOnly, x.CertifiedSource, x.Verdict)
}

func FormatFlavor(x FlavorOrientationInterfaceAudit) string {
	return fmt.Sprintf("targets=%s usesFlavor=%t intertwiner=%t angleMap=%t rejectsProximity=%t certified=%t verdict=%q", strings.Join(x.Targets, ","), x.UsesFlavorData, x.TypedIntertwinerSupplied, x.ObstructionAngleMappedToFlavor, x.RejectsNumericalProximityWithoutMap, x.CertifiedFlavorMap, x.Verdict)
}

func FormatBoundaryMap(x BoundaryMapObstructionAudit) string {
	return fmt.Sprintf("missingPsi=%q missingTau=%q hasPsi=%t hasTau=%t boundaryPair=%t seven=%t scalarFlavor=%t verdict=%q", x.MissingPsi, x.MissingTau, x.HasPsi, x.HasTau, x.CanAssignBoundaryPair, x.CanAssignSevenOver72, x.CanAssignScalarFlavor, x.Verdict)
}

func FormatSeal(x FanoHitchinObstructionSeal) string {
	return fmt.Sprintf("name=%q carrier=%q split=%q source=%q normalForm=%q ray=%q cos=%q rho2=%q boundary=%q internalOnly=%t verdict=%q", x.Name, x.Carrier, x.Split, x.Source, x.NormalForm, x.HitchinMetricRay, x.CosTheta, x.ResidualSquare, x.BoundaryStatus, x.InternalOnly, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("splitG2=%t boundary=%t seven=%t scalarFlavor=%t history=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsHistoryLoopUnit, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
