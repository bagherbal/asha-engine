package generation2boundaryexteriordegreeresponsemapandcubicstopsourceaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryExteriorDegreeResponseMapAndCubicStopSourceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 784 — Boundary Exterior-Degree Response Map and Cubic Stop Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate783 raw-moment generating function", Passed: a.Gate783.Inherited && strings.Contains(a.Gate783.ResponseFunction, "f_3") && strings.Contains(a.Gate783.BoundarySubBottleneck, "F_wall_3_red"), Detail: a.Gate783.ResponseFunction},
			{Name: "type boundary pair object", Passed: a.BoundaryPair.Typed && a.BoundaryPair.Dim == 2 && containsAll(a.BoundaryPair.Basis, []string{"b_lambda", "b_R"}) && a.BoundaryPair.TwoDimensionalBridge && !a.BoundaryPair.IsSpacetimeCarrier && !a.BoundaryPair.IsFlavorCarrier && !a.BoundaryPair.IsK7Carrier, Detail: FormatBoundaryPair(a.BoundaryPair)},
			{Name: "audit exterior-degree candidate", Passed: a.Exterior.Audited && a.Exterior.Lambda3Zero && a.Exterior.CubicStopCandidate && !a.Exterior.DimensionAloneProvesStop && !a.Exterior.NativeThetaExtMapCertified && strings.Contains(a.Exterior.DegreeAssignments["M4"], "Lambda^3"), Detail: FormatExterior(a.Exterior)},
			{Name: "identify required Theta_ext map", Passed: a.ThetaExt.Identified && a.ThetaExt.MapName == "Theta_ext" && strings.Contains(a.ThetaExt.Codomain, "Lambda^(n-1)") && !a.ThetaExt.Certified && a.ThetaExt.DimensionShortcutRejected && containsAll(a.ThetaExt.Assignments, []string{"M1", "M2", "M3", "M4"}), Detail: FormatTheta(a.ThetaExt)},
			{Name: "audit cubic coefficient with degree-two source", Passed: a.Cubic.Audited && closeRel(a.Cubic.TwoP, 7.0/36.0, 1e-15) && a.Cubic.CompatibleWithDegree2 && strings.Contains(a.Cubic.MagnitudeSource, "dim(B_boundary)") && !a.Cubic.SignSourceCertified && !a.Cubic.StressPullTheorem, Detail: FormatCubic(a.Cubic)},
			{Name: "record degree-by-degree response table", Passed: a.Table.Recorded && strings.Contains(a.Table.Degree0, "M1") && strings.Contains(a.Table.Degree1, "kappa_e_red M2") && strings.Contains(a.Table.Degree2, "-2p M3") && strings.Contains(a.Table.Degree3, "Lambda^3") && a.Table.KappaERedNaturallyDegree1 && !a.Table.FlavorBoundaryModulationNative, Detail: strings.Join([]string{a.Table.Degree0, a.Table.Degree1, a.Table.Degree2, a.Table.Degree3}, "; ")},
			{Name: "preserve projector idempotence firewall", Passed: a.Projector.Preserved && strings.Contains(a.Projector.PowerLaw, "s^n P_7") && !a.Projector.ProjectorStopsExpansion && a.Projector.NeedsExteriorMap, Detail: a.Projector.PowerLaw},
			{Name: "reaudit M4 rejection", Passed: a.M4.Audited && a.M4.UntypedM4FitRejected && a.M4.BlockedIfThetaExt && !a.M4.NativeThetaExtCertified && !a.M4.TypedCoefficientSource && !a.M4.NativeCubicStopTheorem && a.M4.M4 > 0, Detail: a.M4.M4Degree},
			{Name: "record generating-function candidate", Passed: a.Generating.Recorded && strings.Contains(a.Generating.Form, "Tr") && strings.Contains(a.Generating.Truncation, "x+kappa_e_red") && strings.Contains(a.Generating.ExteriorDegreeReading, "degree-2") && !a.Generating.NativeFunctionTheorem && !a.Generating.HigherTermsProvedVanish, Detail: a.Generating.Truncation},
			{Name: "record relation to kappa_lambda_red", Passed: a.Relation.Recorded && strings.Contains(a.Relation.KappaLambdaRelation, "F_wall_3_red") && a.Relation.WouldUpgradeFWall && !a.Relation.KappaLambdaNative && !a.Relation.CHistoryIndependent && strings.Contains(a.Relation.CHiggsStatus, "not Level C"), Detail: FormatRelation(a.Relation)},
			{Name: "record prediction-level classification", Passed: a.Prediction.Recorded && strings.Contains(a.Prediction.FWall3Level, "Level B+") && strings.Contains(a.Prediction.KappaLambdaLevel, "Level B") && strings.Contains(a.Prediction.CHistoryLevel, "Level B") && strings.Contains(a.Prediction.CHiggsLevel, "not Level C") && !a.Prediction.FWallLevelC, Detail: a.Prediction.FWall3Level + "; " + a.Prediction.CHiggsLevel},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.BoundaryExteriorDegreeNative && !a.Firewalls.DimensionTwoProof && !a.Firewalls.TwoPCoefficientNative && !a.Firewalls.NegativeSignNative && !a.Firewalls.FWallNative && !a.Firewalls.KappaLambdaNative && !a.Firewalls.CHistoryIndependent && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNative && a.Firewalls.Verdict == StatusFirewallPreservedGate784, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
