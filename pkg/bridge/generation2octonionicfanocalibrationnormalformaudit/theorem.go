package generation2octonionicfanocalibrationnormalformaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2OctonionicFanoCalibrationNormalFormIdentityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 652 — Octonionic Fano Calibration Normal-Form Identity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate652 octonionic normal-form audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate651 channel calibration and firewall", Passed: a.Inherited.CalibrationInherited && a.Inherited.AAAUnit && a.Inherited.AABEqualNegativeUnits && a.Inherited.ReconstructsPPlusMinus3 && !a.Inherited.FullSymbolicCalibration && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate651FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit B as negative-sector volume form", Passed: a.BVolume.BIsVolumeForm && a.BVolume.Beta == 1 && a.BVolume.ResidualAgainstVolume < tol && len(a.BVolume.Basis) == minusDim, Detail: FormatBVolume(a.BVolume)},
			{Name: "extract A as three calibrated two-form wedges", Passed: a.AExtract.AllExtracted && a.AExtract.OrthogonalTriple && a.AExtract.EqualNorms && a.AExtract.WedgeOrthonormal && a.AExtract.Residual < tol && len(a.AExtract.Rows) == minusDim, Detail: FormatAExtract(a.AExtract)},
			{Name: "audit quaternionic/Fano two-form triple", Passed: a.Quaternionic.FormsDefineEndomorphisms && a.Quaternionic.WedgeIdentityPasses && a.Quaternionic.QuaternionicIdentities && a.Quaternionic.IdentityResidual < tol, Detail: FormatQuaternionic(a.Quaternionic)},
			{Name: "derive AAA positive channel from two-form triple", Passed: a.AAA.CPositive == 1 && a.AAA.ScalarMultipleOfP && a.AAA.AnisotropyResidual < tol, Detail: FormatAAA(a.AAA)},
			{Name: "derive AAB/ABA/BAA negative channels from volume and two-form triple", Passed: a.AAB.EqualToMinusPositive && a.AAB.CombinedCoefficient == -3 && a.AAB.CombinedResidual < tol && len(a.AAB.Rows) == 3, Detail: FormatAAB(a.AAB)},
			{Name: "audit equal-unit source without symbolic promotion", Passed: a.EqualUnit.SameAlphaBetaNormalization && a.EqualUnit.FanoIncidenceSymmetry && a.EqualUnit.QuaternionicNormalization && !a.EqualUnit.RouteSpecificOnly && !a.EqualUnit.BasisFreeProofCertified, Detail: FormatEqualUnit(a.EqualUnit)},
			{Name: "audit route-universal normal form after normalization", Passed: a.Routes.AllRoutesReduce && a.Routes.SameNormalFormAfterNorm && !a.Routes.RouteDependentScale && len(a.Routes.Rows) == 3, Detail: FormatRoutes(a.Routes)},
			{Name: "state normal-form theorem target while preserving theorem gap", Passed: a.Theorem.FiniteNormalFormIdentitiesPass && !a.Theorem.FullSymbolicOctonionicTheorem, Detail: FormatTheorem(a.Theorem)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicOctonionicTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate652Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate652 turns the Gate651 finite equal-unit calibration into a Fano normal-form theorem target while keeping the basis-free proof gap explicit.")
		if a.Theorem.FullSymbolicOctonionicTheorem || a.Firewalls.ClaimsFullSymbolicOctonionicTheorem {
			notes = append(notes, "WARNING_SYMBOLIC_OCTONIONIC_THEOREM_PROMOTION_BLOCKED")
		}
		if !strings.Contains(a.EqualUnit.Verdict, StatusNoFullSymbolicOctonionicTheorem) {
			notes = append(notes, "WARNING_MISSING_OCTONIONIC_THEOREM_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
