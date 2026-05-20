package generation2boundaryhistoryresidualsecondmomentandruntimetransportaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryHistoryResidualSecondMomentAndRuntimePropagationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 729 — Boundary-History Residual Second-Moment and Runtime Propagation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate729 boundary-history residual second-moment audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate728 dual event-expectation runtime assembly", Passed: a.Gate728.Inherited && near(a.Gate728.P_K7, 7.0/72.0, 1e-18) && near(a.Gate728.L, 1/(8*math.Pi), 1e-18) && a.Gate728.DualEventExpectationClosure && a.Gate728.AssembledRuntimeNotIndependent && a.Gate728.PremisesNotNative, Detail: FormatGate728(a.Gate728)},
			{Name: "define boundary uplift response operator", Passed: near(a.Uplift.LeadingExpectation, a.Gate728.P_K7*a.Gate728.SSplit, 1e-18) && a.Uplift.MatchesGate700Leading && strings.Contains(a.Uplift.Verdict, StatusBoundaryUpliftResponseOperatorDefined), Detail: FormatUplift(a.Uplift)},
			{Name: "compute second raw moment and wall residual coefficient", Passed: near(a.Moment.M2Wall, a.Gate728.P_K7*a.Gate728.SSplit*a.Gate728.SSplit, 1e-18) && near(a.Moment.C2Wall, a.Gate728.EWall/a.Moment.M2Wall, 1e-18) && a.Moment.SecondOrderSuppressed && a.Moment.ResidualMuchSmallerThanMoment && strings.Contains(a.Moment.Verdict, StatusSecondRawMomentComputed), Detail: FormatMoment(a.Moment)},
			{Name: "audit typed coefficient candidates", Passed: a.Coefficients.KappaEClosestSmall && a.Coefficients.NotExact && a.Coefficients.ClosestName == "kappa_e" && strings.Contains(a.Coefficients.Verdict, StatusKappaECloseToSecondOrderCoefficient), Detail: FormatCoefficients(a.Coefficients)},
			{Name: "test kappa_e second-order correction", Passed: near(a.KappaECorr.KappaEM2, kappaE*a.Moment.M2Wall, 1e-18) && near(a.KappaECorr.ResidualAfterCorrection, a.Gate728.EWall-a.KappaECorr.KappaEM2, 1e-18) && a.KappaECorr.ImprovesRawResidual && a.KappaECorr.NotExact && a.KappaECorr.NotIndependentlyCertified && strings.Contains(a.KappaECorr.Verdict, StatusKappaESecondOrderCorrectionTested), Detail: FormatKappaECorrection(a.KappaECorr)},
			{Name: "audit variance control scale", Passed: near(a.Variance.VarianceWall, a.Gate728.P_K7*(1-a.Gate728.P_K7)*a.Gate728.SSplit*a.Gate728.SSplit, 1e-18) && a.Variance.RelevantTypedScale && !a.Variance.SelectedActiveCorrection && strings.Contains(a.Variance.Verdict, StatusVarianceFormNotYetSelected), Detail: FormatVariance(a.Variance)},
			{Name: "propagate wall residual compression into scalar runtime residual", Passed: near(a.Runtime.RawRuntimeResidual, lambdaProxyMZ*a.Gate728.L*a.Gate728.EWall, 1e-18) && near(a.Runtime.CorrectedRuntimeResidual, lambdaProxyMZ*a.Gate728.L*a.KappaECorr.ResidualAfterCorrection, 1e-18) && a.Runtime.CompressionFollowsWallResidual && strings.Contains(a.Runtime.Verdict, StatusRuntimeResidualPropagationAudited), Detail: FormatRuntime(a.Runtime)},
			{Name: "preserve noncircularity firewall", Passed: a.NonCircular.DBaseContainsKappaE && a.NonCircular.KappaEUsedAsCoefficient && !a.NonCircular.IndependentTheorem && a.NonCircular.PartiallyDependent && strings.Contains(a.NonCircular.Verdict, StatusKappaEResidualCoefficientDependent), Detail: FormatNoncircularity(a.NonCircular)},
			{Name: "preserve physical theorem firewalls", Passed: !a.Firewall.ClaimsNativeBoundaryHistory && !a.Firewall.ClaimsNativeSecondOrderExpansion && !a.Firewall.ClaimsNativeScalarRuntime && !a.Firewall.ClaimsHiggsMassTheorem && !a.Firewall.ClaimsYukawaOperatorTheorem && !a.Firewall.ClaimsCKMPMNSTheorem && strings.Contains(a.Firewall.Verdict, StatusGate729Boundary), Detail: FormatFirewall(a.Firewall)},
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
