package generation2k7splitsignaturehodgebilinearaudit

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

func FormatInherited(i Gate635Inheritance) string {
	return fmt.Sprintf("K7=%d plus=%d minus=%d trace=%s det=%s carrierFirewall=%t noTheta=%t no1plus3=%t traceNoLine=%t noBoundary=%t no7over72=%t verdict=%q", i.K7Dimension, i.PlusDimension, i.MinusDimension, f64(i.Trace), f64(i.Determinant), i.CarrierFirewallPreserved, i.NoK7ToFockSelectorMap, i.NoOnePlusThreeRefinement, i.TraceNotDistinguishedLine, i.NoBoundaryAssignment, i.NoSevenOver72Theorem, i.Verdict)
}

func FormatDefinition(d HodgeBilinearDefinition) string {
	return fmt.Sprintf("formula=%q metric=%q matrix=%q dim=%d rows=%d cols=%d symmetric=%t nondegenerate=%t inherited=%t verdict=%q", d.Formula, d.MetricFormula, d.MatrixRepresentative, d.Dimension, d.Rows, d.Cols, d.Symmetric, d.Nondegenerate, d.InheritedFromSK, d.Verdict)
}

func FormatSignature(s SignatureCertificate) string {
	return fmt.Sprintf("inertia=(%d,%d,%d) trace=%s det=%s detSign=%d eig=%s nullCone=%t pos=%t neg=%t split=%t notation=%q verdict=%q", s.InertiaPlus, s.InertiaMinus, s.InertiaZero, f64(s.Trace), f64(s.Determinant), s.DeterminantSign, f64s(s.Eigenvalues), s.NullConeExists, s.PositiveDefinite, s.NegativeDefinite, s.SplitIndefinite, s.SignatureNotation, s.Verdict)
}

func FormatMetricConversion(m MetricConversionAudit) string {
	return fmt.Sprintf("g=%q relation=%q op=%q orth=%t sym=%t inv=%t symResidual=%s orthResidual=%s invResidual=%s equals=%t verdict=%q", m.EuclideanMetric, m.BilinearRelation, m.ConversionOperator, m.SKOrthogonal, m.SKSymmetric, m.SKInvolutive, f64(m.SymmetryResidual), f64(m.OrthogonalityResidual), f64(m.InvolutionResidual), m.BEqualsGComposedWithSK, m.Verdict)
}

func FormatOrthogonality(o PlusMinusOrthogonalityAudit) string {
	return fmt.Sprintf("dims=%d|%d gOrth=%t bOrth=%t bPlus=%q bMinus=%q crossZero=%t projectorOrth=%s complementarity=%s verdict=%q", o.PlusDimension, o.MinusDimension, o.GOrthogonal, o.BOrthogonal, o.BRestrictedToPlus, o.BRestrictedToMinus, o.CrossTermZero, f64(o.ProjectorOrthogonality), f64(o.ProjectorComplementarity), o.Verdict)
}

func FormatOctonionic(o OctonionicCompatibilityAudit) string {
	return fmt.Sprintf("lane=%q splitMatch=%t omega=%t crossProduct=%t calibration=%t splitG2=%t preserves=%t verdict=%q reason=%q", o.CandidateLane, o.SplitSignatureMatchesDimension, o.OmegaKThreeFormCertified, o.CrossProductCertified, o.CalibrationCertified, o.G2SplitStructureCertified, o.PreservationByNativeG2Operator, o.Verdict, o.Reason)
}

func FormatStabilizer(s StabilizerAudit) string {
	return fmt.Sprintf("bilinear=%q orientation=%q splitG2=%q stabilizerCertified=%t splitG2Certified=%t needsOmega=%t physicalMetric=%t verdict=%q", s.BilinearStabilizerCandidate, s.OrientationPreservingCandidate, s.SplitG2CandidateSubgroup, s.StabilizerCertified, s.SplitG2Certified, s.NeedsOmegaK, s.PhysicalMetricClaimed, s.Verdict)
}

func FormatFirewalls(f SelectorBoundaryFirewall) string {
	return fmt.Sprintf("k7Fock=%t onePlusThree=%t boundary=%t sevenOver72=%t physicalMetric=%t scalar=%t higgs=%t flavor=%t ckmPmns=%t gauge=%t verdict=%q", f.K7ToFockMapCertified, f.OnePlusThreeSelectorDerived, f.BoundaryStressAssigned, f.SevenOver72Promoted, f.PhysicalSpacetimeMetric, f.ScalarRGMatchingClaimed, f.HiggsMassClaimed, f.FlavorClaimed, f.CKMPMNSClaimed, f.GaugeUnificationClaimed, f.Verdict)
}

func FormatMissingObject(m MissingObjectAudit) string {
	return fmt.Sprintf("previous=%q current=%q why=%q splitG2=%t boundary=%t omegaVerdict=%q boundaryVerdict=%q", m.PreviousMissingObject, m.CurrentMissingObject, m.WhySharper, m.CanSupportSplitG2, m.CanSupportBoundary, m.VerdictOmega, m.VerdictBoundary)
}
