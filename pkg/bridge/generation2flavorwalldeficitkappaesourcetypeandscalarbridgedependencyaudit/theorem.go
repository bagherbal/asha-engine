package generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FlavorWallDeficitKappaESourceTypeAndScalarBridgeDependencyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 746 — Flavor-Wall Deficit Kappa_e Source-Type and Scalar-Bridge Dependency Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate746 kappa_e source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate745 pole diagnostic boundary", Passed: a.Gate745.Inherited && a.Gate745.DiagnosticAllowed && a.Gate745.ExternalPoleIsNotASHA && a.Gate745.NoIndependentPoleMass && a.Gate745.NoYukawaTheorem && strings.Contains(a.Gate745.Verdict, StatusGate745PoleDiagnosticBoundaryInherited), Detail: FormatGate745(a.Gate745)},
			{Name: "audit kappa_e scalar bridge dependency", Passed: Near(a.Dependency.KappaE, 0.00550355419157456, 1e-18) && a.Dependency.AppearsInBoundaryPolynomial && a.Dependency.AppearsInRuntimeTransport && a.Dependency.StructurallyActive && strings.Contains(a.Dependency.Verdict, StatusKappaEScalarBridgeDependencyAudited), Detail: FormatDependency(a.Dependency)},
			{Name: "compute orientation candidate", Passed: Near(a.Orientation.KappaEOrient, 0.00550633006471245, 1e-18) && Near(a.Orientation.DeltaKappaE, -2.77587313789e-6, 1e-15) && a.Orientation.CloseButNotExact && a.Orientation.TypedPMNSLeakage && a.Orientation.TypedCKMCorrection && !a.Orientation.NativeTheorem && strings.Contains(a.Orientation.Verdict, StatusKappaEOrientationCandidateComputed), Detail: FormatOrientation(a.Orientation)},
			{Name: "test kappa_e_orient replacement in scalar bridge", Passed: math.Abs(a.Replacement.RuntimeShift) > 1e-8 && math.Abs(a.Replacement.RuntimeShift) < 2e-8 && Near(a.Replacement.RuntimeOrientResidual, -1.3795353970280644e-8, 5e-14) && a.Replacement.ApproximationOnly && strings.Contains(a.Replacement.Verdict, StatusKappaEOrientReplacementTested), Detail: FormatReplacement(a.Replacement)},
			{Name: "audit residual source candidates", Passed: len(a.Residual.Candidates) == 4 && !a.Residual.NativeSourceCertified && strings.Contains(a.Residual.Verdict, StatusDeltaKappaESourceCandidatesAudited) && strings.Contains(a.Residual.Verdict, StatusNoNativeKappaEOrientationResidualSource), Detail: FormatResidual(a.Residual)},
			{Name: "enforce flavor and scalar-runtime firewalls", Passed: !a.Firewall.DerivesPMNS && !a.Firewall.DerivesCKM && !a.Firewall.DerivesFlavorHierarchy && !a.Firewall.DerivesYukawaEigenvalues && !a.Firewall.DerivesScalarRuntime && !a.Firewall.DerivesHiggsMass && a.Firewall.KappaEStillBridgeSeal && strings.Contains(a.Firewall.Verdict, StatusFlavorFirewallEnforced), Detail: FormatFirewall(a.Firewall)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
