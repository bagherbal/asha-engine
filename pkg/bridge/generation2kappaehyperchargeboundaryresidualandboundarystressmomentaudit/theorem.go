package generation2kappaehyperchargeboundaryresidualandboundarystressmomentaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2KappaEHyperchargeBoundaryResidualAndBoundaryStressMomentAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 748 — Kappa_e Hypercharge-Boundary Residual and Boundary-Stress Moment Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate748 kappa_e boundary-stress moment audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate747 hypercharge boundary-square audit", Passed: a.Gate747.Inherited && a.Gate747.HyperBoundaryCorrectionClose && a.Gate747.ResidualNotZero && a.Gate747.FlavorFirewallKept && strings.Contains(a.Gate747.Verdict, StatusGate747KappaEHyperchargeBoundarySquareInherited), Detail: FormatGate747(a.Gate747)},
			{Name: "compute Gate747 residual over K7 second raw moment", Passed: Near(a.Residual.Residual, 8.149544918e-9, 1e-17) && Near(a.Residual.M2Wall, 1.624013231638281e-7, 1e-19) && Near(a.Residual.Ratio, 0.0501815179795, 1e-12) && a.Residual.SecondMomentScale && strings.Contains(a.Residual.Verdict, StatusGate747ResidualOverM2WallComputed), Detail: FormatResidual(a.Residual)},
			{Name: "audit boundary-stress coefficient candidates", Passed: Near(a.Stress.XiBoundary, 0.0503471644870914, 1e-15) && a.Stress.BestCandidate == "xi_boundary midpoint" && len(a.Stress.Candidates) == 3 && strings.Contains(a.Stress.Verdict, StatusBoundaryStressCandidatesAudited), Detail: FormatStress(a.Stress)},
			{Name: "define xi_boundary M2_wall correction", Passed: Near(a.Correction.StressMoment, 8.176446130250547e-9, 1e-18) && Near(a.Correction.KappaEHyperStress, 0.005503554218475772, 1e-18) && Near(a.Correction.ResidualAfterCorrection, -2.6901212160646004e-11, 1e-18) && a.Correction.CompressionFactor > 250 && a.Correction.CorrectionNotExact && strings.Contains(a.Correction.Verdict, StatusBoundaryStressMomentCorrectionDefined), Detail: FormatCorrection(a.Correction)},
			{Name: "test scalar-runtime replacement with boundary-stress moment correction", Passed: Near(a.Replacement.RuntimeHyperStressShift, 1.3369860774048448e-13, 5e-15) && math.Abs(a.Replacement.RuntimeHyperStressShift) < math.Abs(a.Replacement.RuntimeHyperBoundaryShift) && a.Replacement.StressImprovementOverOrient > 1e5 && a.Replacement.StressImprovementOverHyper > 100 && a.Replacement.ReplacementNotNative && strings.Contains(a.Replacement.Verdict, StatusScalarRuntimeReplacementTested), Detail: FormatReplacement(a.Replacement)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceType.Expression, "xi_boundary") && len(a.SourceType.Terms) == 4 && strings.Contains(a.SourceType.Verdict, StatusSourceTypeInterpretationRecorded), Detail: FormatSourceType(a.SourceType)},
			{Name: "enforce noncircularity and physical firewalls", Passed: a.Firewall.Theta13EmpiricalBridgeInput && a.Firewall.JCKMEmpiricalBridgeInput && a.Firewall.FiveThirdsMatureButUncoupled && a.Firewall.XiBoundaryBridgeStressQuantity && a.Firewall.M2WallBoundaryMomentNotFlavor && !a.Firewall.DerivesFlavorTheorem && !a.Firewall.DerivesScalarRuntime && !a.Firewall.DerivesHiggsMass && !a.Firewall.DerivesYukawa && strings.Contains(a.Firewall.Verdict, StatusGate748Boundary), Detail: FormatFirewall(a.Firewall)},
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
