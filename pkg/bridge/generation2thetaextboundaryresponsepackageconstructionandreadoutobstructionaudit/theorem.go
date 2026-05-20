package generation2thetaextboundaryresponsepackageconstructionandreadoutobstructionaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ThetaExtBoundaryResponsePackageConstructionAndReadoutObstructionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 785 — ThetaExt Boundary Response Package Construction and Readout Obstruction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate784 exterior-degree response boundary", Passed: a.Gate784.Inherited && strings.Contains(a.Gate784.MissingObject, "Theta_ext") && strings.Contains(a.Gate784.PriorLevel, "Level B+"), Detail: a.Gate784.MissingObject},
			{Name: "separate lift and readout problems", Passed: a.Separation.Separated && strings.Contains(a.Separation.ThetaExtCodomain, "Lambda^(n-1)") && strings.Contains(a.Separation.ChiExtDomain, "Lambda^0") && strings.Contains(a.Separation.PolynomialRepresentation, "chi_ext") && !a.Separation.ThetaExtAloneSufficient, Detail: FormatSeparation(a.Separation)},
			{Name: "type exterior response algebra", Passed: a.Algebra.Typed && strings.Contains(a.Algebra.Algebra, "Lambda^0") && containsAll(a.Algebra.Basis, []string{"b_lambda", "b_R"}) && strings.Contains(a.Algebra.VolumeForm, "omega_B") && a.Algebra.RequiresDegreeOneAxis && !a.Algebra.HasNativeDegreeOneAxis && a.Algebra.ConditionalLabelledBasis, Detail: a.Algebra.Algebra},
			{Name: "construct conditional Theta_ext response package", Passed: a.Package.Constructed && strings.Contains(a.Package.ThetaM1, "1_B") && strings.Contains(a.Package.ThetaM2, "beta_B") && strings.Contains(a.Package.ThetaM3, "omega_B") && strings.Contains(a.Package.ThetaHigher, "M_n>=4") && closeRel(a.Package.Chi0, 1, 1e-15) && closeRel(a.Package.Chi1, a.Ledger.KappaE, 1e-15) && closeRel(a.Package.Chi2, -2*a.Ledger.P, 1e-15) && a.Package.MatchesFWall3 && !a.Package.ReadoutNative, Detail: FormatPackage(a.Package)},
			{Name: "complete naturality audit", Passed: a.Naturality.Completed && a.Naturality.OnlyDimensionTwoNative && !a.Naturality.CanonicalNonzeroVectorFromDimension && !a.Naturality.CanonicalNonzeroCovectorFromDimension && a.Naturality.LabelledBasisConditional, Detail: a.Naturality.Verdict},
			{Name: "separate magnitude and sign", Passed: a.MagnitudeSign.Separated && closeRel(a.MagnitudeSign.TwoP, 7.0/36.0, 1e-15) && strings.Contains(a.MagnitudeSign.MagnitudeSource, "dim(B_boundary)") && strings.Contains(a.MagnitudeSign.Sign, "negative") && !a.MagnitudeSign.OrientationSignNative, Detail: a.MagnitudeSign.Sign},
			{Name: "audit cubic stop under conditional package", Passed: a.CubicStop.Audited && a.CubicStop.M4 > 0 && strings.Contains(a.CubicStop.M4Degree, "Lambda^3") && a.CubicStop.BlockedIfDegreeRule && !a.CubicStop.DegreeRuleNative && !a.CubicStop.CubicStopNative, Detail: a.CubicStop.M4Degree},
			{Name: "audit exterior exponential shortcut", Passed: a.Exponential.Audited && strings.Contains(a.Exponential.SingleVectorExp, "1+beta") && !a.Exponential.SingleVectorProducesDegreeTwo && strings.Contains(a.Exponential.DegreeTwoRequirement, "two distinct boundary legs"), Detail: a.Exponential.DegreeTwoRequirement},
			{Name: "reclassify prediction status", Passed: a.Prediction.Reclassified && strings.Contains(a.Prediction.FWall3Status, "Level B+") && strings.Contains(a.Prediction.KappaLambda, "Level B") && strings.Contains(a.Prediction.CHistory, "Level B") && strings.Contains(a.Prediction.CHiggs, "not Level C") && !a.Prediction.FWallLevelC, Detail: FormatPrediction(a.Prediction)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.ThetaExtPackageNative && !a.Firewalls.ConditionalBetaNative && !a.Firewalls.ChiExtNative && !a.Firewalls.DimensionTwoProof && !a.Firewalls.OmegaBSignDerived && !a.Firewalls.FWallNative && !a.Firewalls.KappaLambdaNative && !a.Firewalls.CHistoryIndependent && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNative && a.Firewalls.Verdict == StatusFirewallPreservedGate785, Detail: a.Firewalls.Verdict},
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
