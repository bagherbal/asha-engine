package generation2kappaeorientationresidualandhyperchargenormalizedboundarysquareaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2KappaEOrientationResidualAndHyperchargeNormalizedBoundarySquareAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 747 — Kappa_e Orientation Residual and Hypercharge-Normalized Boundary-Square Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate747 kappa_e hypercharge-boundary audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate746 kappa_e source audit", Passed: a.Gate746.Inherited && a.Gate746.KappaEActiveInput && a.Gate746.OrientationClose && a.Gate746.OrientationNotExact && a.Gate746.FlavorFirewallKept && strings.Contains(a.Gate746.Verdict, StatusGate746KappaESourceAuditInherited), Detail: FormatGate746(a.Gate746)},
			{Name: "compute Delta_kappa_e/S_split^2 and audit typed ratios", Passed: Near(a.Ratio.DeltaKappaE, -2.77587313789e-6, 1e-15) && Near(a.Ratio.SSplitSquared, 1.6704136096850888e-6, 1e-18) && Near(a.Ratio.Ratio, -1.6617879079741393, 1e-12) && a.Ratio.CloseToMinusFiveThirds && a.Ratio.BestCandidate == "-5/3 hypercharge/gauge normalization" && strings.Contains(a.Ratio.Verdict, StatusDeltaKappaEOverSSplitSquaredComputed), Detail: FormatRatio(a.Ratio)},
			{Name: "define hypercharge-normalized boundary-square correction", Passed: Near(a.Correction.Correction, -2.7840226828084814e-6, 1e-18) && Near(a.Correction.KappaEHyperBoundary, 0.005503546042029642, 1e-18) && Near(a.Correction.ResidualAfterCorrection, 8.149544918367644e-9, 1e-18) && a.Correction.CompressionFactor > 330 && a.Correction.CompressionFactor < 350 && a.Correction.CorrectionNotExact && strings.Contains(a.Correction.Verdict, StatusHyperchargeBoundarySquareCorrectionDefined), Detail: FormatCorrection(a.Correction)},
			{Name: "test scalar-runtime replacement with hypercharge boundary correction", Passed: math.Abs(a.Replacement.RuntimeOrientShift) > 1e-8 && Near(a.Replacement.RuntimeHyperBoundaryShift, -4.050107471620379e-11, 5e-14) && a.Replacement.ImprovementFactor > 330 && a.Replacement.ImprovementFactor < 350 && a.Replacement.ReplacementNotNative && strings.Contains(a.Replacement.Verdict, StatusScalarRuntimeReplacementTested), Detail: FormatReplacement(a.Replacement)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceType.Expression, "sin^2(theta13)/4") && len(a.SourceType.Terms) == 3 && strings.Contains(a.SourceType.Verdict, StatusSourceTypeInterpretationRecorded), Detail: FormatSourceType(a.SourceType)},
			{Name: "enforce noncircularity and physical firewalls", Passed: a.Firewall.Theta13EmpiricalBridgeInput && a.Firewall.JCKMEmpiricalBridgeInput && a.Firewall.FiveThirdsMatureButUncoupled && a.Firewall.SSplitBoundaryNotFlavorOperator && !a.Firewall.DerivesFlavorTheorem && !a.Firewall.DerivesScalarRuntime && !a.Firewall.DerivesHiggsMass && !a.Firewall.DerivesYukawa && strings.Contains(a.Firewall.Verdict, StatusGate747Boundary), Detail: FormatFirewall(a.Firewall)},
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
