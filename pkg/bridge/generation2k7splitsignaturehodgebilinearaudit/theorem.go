package generation2k7splitsignaturehodgebilinearaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7SplitSignatureHodgeBilinearAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 636 — K7 Split-Signature Hodge Bilinear Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate636 K7 split-signature Hodge bilinear audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate635 Hodge-polarity carrier firewall", Passed: a.Inherited.Verdict == StatusGate635Inherited && a.Inherited.K7Dimension == 7 && a.Inherited.PlusDimension == 4 && a.Inherited.MinusDimension == 3 && math.Abs(a.Inherited.Trace-1) < 1e-10 && math.Abs(a.Inherited.Determinant+1) < 1e-8 && a.Inherited.CarrierFirewallPreserved && a.Inherited.NoK7ToFockSelectorMap && a.Inherited.NoOnePlusThreeRefinement && a.Inherited.TraceNotDistinguishedLine && a.Inherited.NoBoundaryAssignment && a.Inherited.NoSevenOver72Theorem, Detail: FormatInherited(a.Inherited)},
			{Name: "define Hodge bilinear B_K on K7", Passed: a.Definition.Verdict == StatusBKHodgeBilinearDefined && a.Definition.Formula == "B_K(x,y)=<x,S_*y>|_{K_7}" && a.Definition.Dimension == 7 && a.Definition.Rows == 7 && a.Definition.Cols == 7 && a.Definition.Symmetric && a.Definition.Nondegenerate && a.Definition.InheritedFromSK, Detail: FormatDefinition(a.Definition)},
			{Name: "certify B_K signature (4,3)", Passed: a.Signature.Verdict == StatusBKSignatureCertified && a.Signature.InertiaPlus == 4 && a.Signature.InertiaMinus == 3 && a.Signature.InertiaZero == 0 && a.Signature.DeterminantSign == -1 && math.Abs(a.Signature.Trace-1) < 1e-10 && math.Abs(a.Signature.Determinant+1) < 1e-8 && a.Signature.NullConeExists && !a.Signature.PositiveDefinite && !a.Signature.NegativeDefinite && a.Signature.SplitIndefinite, Detail: FormatSignature(a.Signature)},
			{Name: "certify S_K as metric-conversion operator", Passed: a.MetricConversion.Verdict == StatusBKMetricConversionCertified && a.MetricConversion.SKOrthogonal && a.MetricConversion.SKSymmetric && a.MetricConversion.SKInvolutive && a.MetricConversion.SymmetryResidual < 1e-12 && a.MetricConversion.OrthogonalityResidual < 1e-12 && a.MetricConversion.InvolutionResidual < 1e-12 && a.MetricConversion.BEqualsGComposedWithSK, Detail: FormatMetricConversion(a.MetricConversion)},
			{Name: "certify plus/minus orthogonality under g_K and B_K", Passed: a.Orthogonality.Verdict == StatusPlusMinusOrthogonalityCertified && a.Orthogonality.PlusDimension == 4 && a.Orthogonality.MinusDimension == 3 && a.Orthogonality.GOrthogonal && a.Orthogonality.BOrthogonal && a.Orthogonality.CrossTermZero && a.Orthogonality.ProjectorOrthogonality < 1e-12 && a.Orthogonality.ProjectorComplementarity < 1e-12, Detail: FormatOrthogonality(a.Orthogonality)},
			{Name: "audit split-octonionic compatibility as missing Omega_K", Passed: strings.Contains(a.Octonionic.Verdict, StatusNativeSplitSignature) && strings.Contains(a.Octonionic.Verdict, StatusNoSplitG2Yet) && strings.Contains(a.Octonionic.Verdict, StatusNoOmegaK) && a.Octonionic.SplitSignatureMatchesDimension && !a.Octonionic.OmegaKThreeFormCertified && !a.Octonionic.CrossProductCertified && !a.Octonionic.CalibrationCertified && !a.Octonionic.G2SplitStructureCertified && !a.Octonionic.PreservationByNativeG2Operator, Detail: FormatOctonionic(a.Octonionic)},
			{Name: "audit bilinear stabilizer without split-G2 promotion", Passed: strings.Contains(a.Stabilizer.Verdict, StatusStabilizerCandidateAudited) && strings.Contains(a.Stabilizer.Verdict, StatusNoSplitG2Yet) && a.Stabilizer.BilinearStabilizerCandidate == "O(4,3)" && a.Stabilizer.OrientationPreservingCandidate == "SO(4,3)" && a.Stabilizer.StabilizerCertified && !a.Stabilizer.SplitG2Certified && a.Stabilizer.NeedsOmegaK && !a.Stabilizer.PhysicalMetricClaimed, Detail: FormatStabilizer(a.Stabilizer)},
			{Name: "preserve selector, physical, boundary, and 7/72 firewalls", Passed: !a.Firewalls.K7ToFockMapCertified && !a.Firewalls.OnePlusThreeSelectorDerived && !a.Firewalls.BoundaryStressAssigned && !a.Firewalls.SevenOver72Promoted && !a.Firewalls.PhysicalSpacetimeMetric && !a.Firewalls.ScalarRGMatchingClaimed && !a.Firewalls.HiggsMassClaimed && !a.Firewalls.FlavorClaimed && !a.Firewalls.CKMPMNSClaimed && !a.Firewalls.GaugeUnificationClaimed && a.Firewalls.Verdict == StatusGate636Boundary, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "sharpen missing object to Omega_K", Passed: a.MissingObject.VerdictOmega == StatusNoOmegaK && a.MissingObject.VerdictBoundary == StatusNoBoundaryStressAssignment && strings.Contains(a.MissingObject.CurrentMissingObject, "Omega_K") && !a.MissingObject.CanSupportSplitG2 && !a.MissingObject.CanSupportBoundary, Detail: FormatMissingObject(a.MissingObject)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed posture: B_K is a native split-signature bilinear form with inertia (4,3); the next missing object is Omega_K, not a Fock selector, boundary map, or physical spacetime metric.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
