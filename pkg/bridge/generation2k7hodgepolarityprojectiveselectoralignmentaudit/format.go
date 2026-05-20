package generation2k7hodgepolarityprojectiveselectoralignmentaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "NaN"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func f64s(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = f64(x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatInherited(i Gate634Inheritance) string {
	return fmt.Sprintf("K7=%d plus=%d minus=%d trace=%s det=%s stable=%t mixed=%t noBoundary=%t no7over72=%t firewall=%t verdict=%q", i.K7Dimension, i.PlusDimension, i.MinusDimension, f64(i.Trace), f64(i.Determinant), i.HodgeStable, i.MixedHodgePolarity, i.NoBoundaryAssignment, i.NoSevenOver72Theorem, i.Gate634FirewallPreserved, i.Verdict)
}

func FormatK7Subspaces(s K7PolaritySubspaceAudit) string {
	return fmt.Sprintf("carrier=%q plus=%q minus=%q dims=%d|%d sum=%d certified=%t orth=%t native=%t verdict=%q", s.Carrier, s.PlusFormula, s.MinusFormula, s.PlusDimension, s.MinusDimension, s.SumDimension, s.ProjectorsCertified, s.PlusMinusOrthogonal, s.NativeHodgePolarity, s.Verdict)
}

func FormatProjectiveSelector(p ProjectiveSelectorReference) string {
	return fmt.Sprintf("carrier=%q dimC=%d selector=%q coeffs=%s split=%q line=%d spatial=%d cp0cp2=%t onePlusThree=%t stabilizer=%q commutant=%t cp3ToK7=%t verdict=%q", p.Carrier, p.CarrierComplexDimension, p.Selector, f64s(p.SelectorCoefficients), p.SelectorSplit, p.LineComplexDimension, p.SpatialBlockComplexDimension, p.CP0CP2CriticalStrata, p.ProjectiveOnePlusThree, p.Stabilizer, p.MatchesGate555Commutant, p.CP3ToK7FunctorFound, p.Verdict)
}

func FormatAlignment(a PolaritySelectorAlignmentAudit) string {
	return fmt.Sprintf("K+=%d K-=%d Wdim=%d selectorDims=%q fourMatch=%t threeMatch=%t sameCarrier=%t theta=%t k7cp3=%t candidate=%t verdict=%q reason=%q", a.K7PlusDimension, a.K7MinusDimension, a.FockCarrierComplexDimension, a.SelectorLinePlusSpatialDims, a.FourDimensionalMatch, a.ThreeDimensionalMatch, a.SameCarrier, a.TypedThetaMapFound, a.K7ToCP3FunctorFound, a.AlignmentCandidateOnly, a.Verdict, a.Reason)
}

func FormatK7PlusRefinement(r K7PlusRefinementAudit) string {
	return fmt.Sprintf("dim=%d eig=%s identity=%t line=%t threePlane=%t onePlusThree=%t verdict=%q reason=%q", r.K7PlusDimension, f64(r.HodgeEigenvalueOnK7Plus), r.HodgeProjectorActsAsIdentity, r.InternalRankOneLineDerived, r.InternalThreePlaneDerived, r.NativeOnePlusThreeRefinement, r.Verdict, r.Reason)
}

func FormatK7MinusTriplet(r K7MinusTripletAudit) string {
	return fmt.Sprintf("K-=%d selectorSpatial=%d dimMatch=%t typedTriplet=%t usesBL=%t verdict=%q reason=%q", r.K7MinusDimension, r.SelectorSpatialBlockDim, r.DimensionMatch, r.TypedTripletIdentification, r.UsesBMinusLCarrier, r.Verdict, r.Reason)
}

func FormatTraceImbalance(t TraceImbalanceAudit) string {
	return fmt.Sprintf("trace=%s diff=%d det=%s line=%t rankOneProjector=%t needsSelector=%t verdict=%q reason=%q", f64(t.Trace), t.PlusMinusDifference, f64(t.Determinant), t.DistinguishedLineDerived, t.TraceAsRankOneProjector, t.NeedsAdditionalSelector, t.Verdict, t.Reason)
}

func FormatCarrierMap(c CarrierMapAudit) string {
	return fmt.Sprintf("map=%q domain=%q codomain=%q dimResemblance=%t typed=%t CP3ToK7=%t K7ToFock=%t status=%q missing=%q verdict=%q", c.CandidateMap, c.Domain, c.Codomain, c.DimensionResemblance, c.TypedIntertwinerFound, c.FunctorFromProjectiveFockToK7Found, c.FunctorFromK7ToProjectiveFockFound, c.Status, c.MissingObject, c.Verdict)
}

func FormatBoundaryReadiness(b BoundaryReadinessAudit) string {
	return fmt.Sprintf("boundary=%t sevenOver72=%t K7W7=%t K7FockSuffices=%t stillRequired=%q boundaryVerdict=%q sevenVerdict=%q", b.BoundaryStressAssignment, b.SevenOver72Promoted, b.K7ToW7PairingReopened, b.K7ToFockMapWouldSuffice, b.StillRequired, b.VerdictBoundary, b.VerdictSevenOver72)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("K7Fock=%t K7plus1plus3=%t boundary=%t sevenOver72=%t scalarRG=%t higgs=%t flavor=%t ckmPMNS=%t gauge=%t physicalOrientation=%t verdict=%q", f.ClaimsK7ToFockSelectorMap, f.ClaimsK7PlusOnePlusThree, f.ClaimsBoundaryStressAssignment, f.ClaimsSevenOver72Theorem, f.ClaimsScalarRGMatching, f.ClaimsHiggsMassDerivation, f.ClaimsFlavorDerivation, f.ClaimsCKMPMNSDerivation, f.ClaimsGaugeUnification, f.ClaimsPhysicalOrientation, f.Verdict)
}
