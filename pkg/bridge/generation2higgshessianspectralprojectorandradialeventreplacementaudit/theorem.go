package generation2higgshessianspectralprojectorandradialeventreplacementaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsHessianSpectralProjectorAndRadialEventReplacementAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 768 — Higgs Hessian Spectral Projector and Radial Event Replacement Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate768HessianSpectralProjectorBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 767 HistoryLoop-Hessian alignment firewall", Passed: a.Gate767.Inherited && a.Gate767.AlignmentSeal == "HistoryLoopHessianRadialAlignmentSeal" && a.Gate767.HistoryProjector == "P_history" && a.Gate767.HessianProjector == "P_hessian" && !a.Gate767.AlignmentNative && !a.Gate767.RankTraceIdentifiesSupport, Detail: FormatGate767(a.Gate767)},
			{Name: "define Hessian spectral support projector", Passed: strings.Contains(a.Spectral.PotentialLane, "U(2)-invariant") && strings.Contains(a.Spectral.HessianFormula, "P_hessian") && a.Spectral.PositiveRadialEigenvalue && a.Spectral.RadialEigenvalueFormula == "2 lambda v^2" && near(a.Spectral.RadialEigenvalueGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV, 1e-9) && len(a.Spectral.AngularEigenvalues) == angularZeroModes && a.Spectral.TraceOfHessianFormula == "Tr(H_V(x_0))=2 lambda v^2" && near(a.Spectral.TraceOfHessianGeV2, a.Spectral.RadialEigenvalueGeV2, 1e-12) && a.Spectral.SupportProjector == "P_Hess=supp(H_V(x_0))" && strings.Contains(a.Spectral.SupportProjectorFormula, "H_V(x_0)/Tr") && a.Spectral.SupportRank == hessianSupportRank && a.Spectral.EqualsHessianProjector && !a.Spectral.NativePotentialTheorem && !a.Spectral.NativeVEVTheorem, Detail: FormatSpectral(a.Spectral)},
			{Name: "replace P_rad by Hessian support inside supplied-potential lane", Passed: strings.Contains(a.Replacement.Before, "independent") && strings.Contains(a.Replacement.After, "P_rad := P_Hess") && strings.Contains(a.Replacement.ReplacementScope, "potential lane only") && a.Replacement.IndependentRadialSymbolReduced && a.Replacement.RequiresSuppliedPotential && a.Replacement.RequiresSuppliedVacuum && a.Replacement.HistoryLoopAlignmentStillRequired && !a.Replacement.NativeAlignmentTheorem, Detail: FormatReplacement(a.Replacement)},
			{Name: "compute HistoryLoop trace with Hessian support", Passed: a.HistoryLoop.State == "rho_plus=I_K7+/4" && strings.Contains(a.HistoryLoop.Projector, "supp") && a.HistoryLoop.Rank == hessianSupportRank && strings.Contains(a.HistoryLoop.TraceWeightFormula, "rank(P_Hess)/4") && near(a.HistoryLoop.TraceWeight, 0.25, 1e-15) && a.HistoryLoop.PhaseLoopPayoff == "1/(2*pi)" && strings.Contains(a.HistoryLoop.LHopfFormula, "supp(H_V") && near(a.HistoryLoop.LHopf, 1/(8*math.Pi), 1e-15) && !a.HistoryLoop.NativeHistoryLoopTheorem, Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "rewrite three-factor form with Hessian support", Passed: strings.Contains(a.ThreeFactor.Formula, "lambda_runtime_eff") && strings.Contains(a.ThreeFactor.SupportForm, "supp(H_V") && near(a.ThreeFactor.NEffective, 3.0023273474722147, 1e-15) && near(a.ThreeFactor.CYukawa, 3/3.0023273474722147, 1e-15) && near(a.ThreeFactor.KappaLambdaRed, 0.04432304306956136, 1e-15) && near(a.ThreeFactor.LHopf, 1/(8*math.Pi), 1e-15) && near(a.ThreeFactor.CHistory, 1.038025177923625, 1e-15) && near(a.ThreeFactor.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && a.ThreeFactor.RewritesOnly && !a.ThreeFactor.IndependentRuntimeTheorem, Detail: FormatThreeFactor(a.ThreeFactor)},
			{Name: "record radial-event source-type upgrade", Passed: a.Upgrade.FromType == "supplied rank-one radial projector" && strings.Contains(a.Upgrade.ToType, "Hessian spectral support") && strings.Contains(a.Upgrade.Upgrade, "no longer arbitrary") && a.Upgrade.StrongerThanGate767 && a.Upgrade.StillBridgeConditional && !a.Upgrade.PotentialAndVacuumNative && !a.Upgrade.HistoryLoopNative, Detail: FormatUpgrade(a.Upgrade)},
			{Name: "record remaining non-reduced obstructions", Passed: !a.Obstruction.PotentialDerived && !a.Obstruction.VacuumDerived && !a.Obstruction.HistoryLoopRuleDerived && !a.Obstruction.HistoryLoopUsesHessianSupportDerived && !a.Obstruction.PoleMassDerived && !a.Obstruction.YukawaDerived && strings.Contains(a.Obstruction.Summary, "does not derive"), Detail: FormatObstruction(a.Obstruction)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.NativePotentialTheorem && !a.Firewalls.NativeVEVTheorem && !a.Firewalls.NativeHistoryLoopHessianAlignment && !a.Firewalls.NativeHistoryLoopUnitTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate768HessianSpectralProjectorBoundary), Detail: FormatFirewalls(a.Firewalls)},
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

func near(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
