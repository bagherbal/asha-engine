package generation2flavorreducedscalarhiggsnormalformandkappaesubstitutionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FlavorReducedScalarHiggsNormalFormAndKappaESubstitutionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 752 — Flavor-Reduced Scalar-Higgs Normal Form and Kappa_e Substitution Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate752 flavor-reduced normal-form audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate751 typed normal form", Passed: a.Gate751.Inherited && a.Gate751.TypedNormalFormReady && a.Gate751.IllegalTermsRejected && a.Gate751.KappaInsertionRecorded, Detail: FormatGate751(a.Gate751)},
			{Name: "inherit Gate748 kappa_e source form", Passed: strings.Contains(a.KappaERed.Formula, "xi_boundary") && strings.Contains(a.KappaERed.Formula, "5/3") && strings.Contains(a.KappaERed.Verdict, StatusKappaEReducedCandidateDefined), Detail: FormatKappaEReduced(a.KappaERed)},
			{Name: "compute kappa_e_red residual", Passed: math.Abs(a.KappaERed.KappaEReduced-0.005503554218475772) < 1e-18 && math.Abs(a.KappaERed.Residual+2.6901212160646004e-11) < 1e-20, Detail: FormatKappaEReduced(a.KappaERed)},
			{Name: "define reduced cubic wall polynomial", Passed: strings.Contains(a.FWallRed.MapType, "Q_boundary -> Q_history") && strings.Contains(a.FWallRed.Polynomial, "kappa_e_red") && math.Abs(a.FWallRed.Difference-4.363913744254155e-18) < 1e-20, Detail: FormatFWallRed(a.FWallRed)},
			{Name: "write reduced scalar-Higgs normal form", Passed: strings.Contains(a.RuntimeRed.Formula, "kappa_e_red") && strings.Contains(a.RuntimeRed.Expanded, "Tr_K7+") && math.Abs(a.RuntimeRed.RuntimeShift-1.3369860774048448e-13) < 5e-16, Detail: FormatRuntimeRed(a.RuntimeRed)},
			{Name: "audit numerical residual", Passed: a.Residual.RuntimeShiftAbs < 2e-13 && strings.Contains(a.Residual.Verdict, StatusNumericalResidualAudited), Detail: FormatResidual(a.Residual)},
			{Name: "audit double insertion sensitivity", Passed: math.Abs(a.Sensitivity.PK7SSquared-1.624013231638281e-7) < 1e-20 && math.Abs(a.Sensitivity.Agreement) < 1e-16 && strings.Contains(a.Sensitivity.Formula, "1-p_K7"), Detail: FormatSensitivity(a.Sensitivity)},
			{Name: "classify reduction status", Passed: a.Reduction.BareSealReduced && !a.Reduction.NativeFlavorTheorem && len(a.Reduction.Components) == 4 && strings.Contains(a.Reduction.Verdict, StatusKappaESealPartiallyReduced), Detail: FormatReduction(a.Reduction)},
			{Name: "preserve physical firewalls", Passed: a.Firewalls.KappaERedNativeBlocked && a.Firewalls.PMNSCKMBlocked && a.Firewalls.FlavorDeficitBlocked && a.Firewalls.YukawaBlocked && a.Firewalls.RuntimePredictionBlocked && a.Firewalls.HiggsMassBlocked, Detail: FormatFirewalls(a.Firewalls)},
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
