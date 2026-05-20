package generation2hitchinchannelsignequalunitcalibrationaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HitchinChannelSignEqualUnitCalibrationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 651 — Hitchin Channel Sign and Equal-Unit Calibration Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate651 Hitchin channel calibration audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate650 sector-degree selection and calibration gap", Passed: a.Inherited.DegreeSelectionInherited && a.Inherited.SectorLedgerDefined && a.Inherited.PositiveAAAOnly && a.Inherited.NegativeAABOnly && a.Inherited.MixedZeroByDegree && a.Inherited.SignCalibrationGapInherited && a.Inherited.SlotFormulaRecovered && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate650FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit orientation and volume conventions", Passed: a.Orientation.PositiveDim == 4 && a.Orientation.NegativeDim == 3 && a.Orientation.OrientationCompatible && a.Orientation.ConventionDependentSign, Detail: FormatOrientation(a.Orientation)},
			{Name: "compute surviving channel bilinear maps", Passed: a.Maps.MapsComputed && len(a.Maps.Rows) == 4 && contains(a.Maps.OnlySurvivors, "AAA") && contains(a.Maps.OnlySurvivors, "AAB") && contains(a.Maps.OnlySurvivors, "ABA") && contains(a.Maps.OnlySurvivors, "BAA"), Detail: FormatMaps(a.Maps)},
			{Name: "audit AAA positive unit", Passed: a.Positive.CPlus > 0 && a.Positive.ScalarMultipleOfP && a.Positive.AnisotropyResidual < tol && a.Positive.ContributesOnlyPlus, Detail: FormatPositive(a.Positive)},
			{Name: "audit AAB/ABA/BAA negative equal unit", Passed: a.Negative.EqualToMinusCPlus && a.Negative.EachScalarMultipleOfP && a.Negative.CombinedCoefficient == -3*a.Positive.CPlus && a.Negative.CombinedAnisotropy < tol, Detail: FormatNegative(a.Negative)},
			{Name: "classify sign source without symbolic promotion", Passed: a.Sign.FiniteNegativeSignObserved && a.Sign.FiniteEqualUnitObserved && !a.Sign.BasisFreeSourceCertified && a.Sign.RequiresCalibrationIdentityProof, Detail: FormatSign(a.Sign)},
			{Name: "audit route-universal calibration pattern", Passed: a.Routes.AllRoutesPass && a.Routes.SamePatternAfterNorm && !a.Routes.RouteDependentMagnitude && len(a.Routes.Rows) == 3, Detail: FormatRoutes(a.Routes)},
			{Name: "reconstruct P_+ - 3P_- and Gate642 angle", Passed: a.Reconstruction.ReconstructsPPlusMinus3P && a.Reconstruction.RecoversGate642Angle && a.Reconstruction.NormSquared == 31, Detail: FormatReconstruction(a.Reconstruction)},
			{Name: "state calibration theorem target while preserving theorem gap", Passed: a.Theorem.FiniteCalibrationIdentityPasses && !a.Theorem.FullSymbolicCalibrationTheorem, Detail: FormatTheorem(a.Theorem)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicCalibration && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate651Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate651 closes the finite sign/equal-unit audit while preserving the missing basis-free calibration theorem.")
		if a.Theorem.FullSymbolicCalibrationTheorem || a.Firewalls.ClaimsFullSymbolicCalibration {
			notes = append(notes, "WARNING_SYMBOLIC_CALIBRATION_THEOREM_PROMOTION_BLOCKED")
		}
		if !strings.Contains(a.Sign.Verdict, StatusNoFullSymbolicCalibrationTheorem) {
			notes = append(notes, "WARNING_MISSING_CALIBRATION_THEOREM_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func contains(xs []string, want string) bool {
	for _, x := range xs {
		if x == want {
			return true
		}
	}
	return false
}
