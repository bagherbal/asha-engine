package generation2boundaryflavorscalarmatchingcomplementindependenceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryFlavorScalarMatchingComplementIndependenceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 782 — Boundary-Flavor Scalar Matching Complement Independence Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusNoNativeBoundaryFlavorScalarMatchingComplementBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 781 C_History macro audit", Passed: a.Gate781.Inherited && a.Gate781.CHistoryFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && closeRel(a.Gate781.LHopf, 0.039788735772973836, 1e-15) && closeRel(a.Gate781.KappaLambdaRed, 0.04432304306956136, 1e-15) && strings.Contains(a.Gate781.SelectedBottleneck, "kappa_lambda_red") && strings.Contains(a.Gate781.SelectedBranch, "Outcome 2"), Detail: FormatGate781(a.Gate781)},
			{Name: "rewrite scalar matching complement as boundary-flavor response form", Passed: a.Rewrite.SelectedBottleneck && strings.Contains(a.Rewrite.BaseFormula, "F_wall_3_red") && strings.Contains(a.Rewrite.WallPolynomial, "kappa_e_red p s^2") && strings.Contains(a.Rewrite.FactoredFormula, "-(1-p s^2)kappa_e_red") && strings.Contains(a.Rewrite.ExpandedFormula, "sin^2(theta13)/4-J_CKM") && strings.Contains(a.Rewrite.SignIdentity, "-(1-p s^2)") && a.Rewrite.SignsChecked && a.Rewrite.NoRuntimeSymbols, Detail: FormatRewrite(a.Rewrite)},
			{Name: "recompute numerical boundary-flavor complement ledger", Passed: closeRel(a.Ledger.M1, 0.0001256543573849177, 1e-14) && closeRel(a.Ledger.M2, 1.624013231638281e-07, 1e-14) && closeRel(a.Ledger.M3, 2.0989474869200057e-10, 1e-14) && closeRel(a.Ledger.FWall3Red, 0.00012565521035653708, 1e-14) && closeRel(a.Ledger.KappaLambdaRed, 0.04432304306956136, 1e-14) && closeRel(a.Ledger.Complement, 0.9556769569304386, 1e-14) && closeRel(a.Ledger.CHistory, 1.038025177923625, 1e-14) && a.Ledger.MatchesFWall && a.Ledger.MatchesKappa && a.Ledger.MatchesComplement && a.Ledger.MatchesCHistory && a.Ledger.MaxAbsDiscrepancy <= 1e-15, Detail: FormatLedger(a.Ledger)},
			{Name: "type every term by wall and layer", Passed: a.Typing.Typed && len(a.Typing.TermTypes) == 8 && len(a.Typing.LayerTypes) == 6 && !a.Typing.RuntimeTargetInFinalFormula && strings.Contains(a.Typing.TermTypes["|lambda|"], "scalar wall") && strings.Contains(a.Typing.TermTypes["1-p s^2"], "scalar matching multiplier") && strings.Contains(a.Typing.LayerTypes["runtime scalar target"], "absent"), Detail: FormatTyping(a.Typing)},
			{Name: "audit K7 role", Passed: a.K7.Audited && a.K7.AppearsOnlyAsPK7 && containsAll(a.K7.RawMoments, []string{"M1=p s", "M2=p s^2", "M3=p s^3"}) && a.K7.NativeSupportOnly && !a.K7.BoundaryVector && !a.K7.FlavorOperator && !a.K7.ScalarWallCoordinate && !a.K7.SourceOfLHopf && !a.K7.HyperchargeNormalization && !a.K7.YukawaTheorem, Detail: FormatK7(a.K7)},
			{Name: "audit boundary raw-moment response polynomial", Passed: a.RawMoment.Audited && a.RawMoment.Formula == "F_wall_3_red=M1+kappa_e_red M2-2pM3" && strings.Contains(a.RawMoment.M1Interpretation, "leading") && strings.Contains(a.RawMoment.M2Interpretation, "flavor-wall") && strings.Contains(a.RawMoment.M3Interpretation, "double-K7") && a.RawMoment.BridgeLayer && !a.RawMoment.NativeGeneratingFunction && !a.RawMoment.NativeRawMomentCoordinate && !a.RawMoment.NativeCubicStop && a.RawMoment.M4ForbiddenWithoutTypedSource && closeRel(a.RawMoment.ComputedFWall, 0.00012565521035653708, 1e-14), Detail: FormatRawMoment(a.RawMoment)},
			{Name: "audit kappa_e reduced flavor wall", Passed: a.Flavor.Audited && strings.Contains(a.Flavor.Formula, "sin^2(theta13)/4") && strings.Contains(a.Flavor.Classification, "external flavor bridge") && !a.Flavor.Theta13Native && !a.Flavor.JCKMNative && !a.Flavor.NativeKappaETheorem && !a.Flavor.NativePMNSOrCKMTheorem && !a.Flavor.NativeYukawaTheorem && !a.Flavor.CanCompareOlderKappaE && math.IsNaN(a.Flavor.OlderKappaEResidual), Detail: FormatFlavor(a.Flavor)},
			{Name: "audit formula-level runtime target absence", Passed: a.Runtime.Audited && strings.Contains(a.Runtime.FinalFormula, "xi_boundary p s^2") && !a.Runtime.UsesLambdaRuntime && !a.Runtime.UsesLambdaRuntimeEff && !a.Runtime.UsesTreeMass && !a.Runtime.UsesPoleMass && !a.Runtime.UsesCHiggs && !a.Runtime.UsesGF && !a.Runtime.UsesVEV && !a.Runtime.UsesHiggsPoleObservable && a.Runtime.FormulaLevelRuntimeTargetAbsence && a.Runtime.EvaluableWithoutDirectHiggsRuntimeVariables && strings.Contains(a.Runtime.FormulaLevelIndependence, "Level B") && strings.Contains(a.Runtime.TheoremLevelIndependence, "not Level C"), Detail: FormatRuntime(a.Runtime)},
			{Name: "audit theorem-level independence firewall", Passed: !a.Runtime.NativeDerivation && !a.Runtime.RawBoundaryResponseIndependentlyProved && !a.Runtime.FlavorInputsIndependentlyProved && !a.Runtime.BoundaryCoordinatesNative, Detail: FormatRuntime(a.Runtime)},
			{Name: "rewrite C_History with full boundary-flavor complement", Passed: a.CHistory.Audited && a.CHistory.Formula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && strings.Contains(a.CHistory.ExpandedForm, "1-|lambda|-p s+2p^2s^3") && strings.Contains(a.CHistory.ExpandedForm, "sin^2(theta13)/4-J_CKM") && closeRel(a.CHistory.Complement, 0.9556769569304386, 1e-14) && closeRel(a.CHistory.CHistory, 1.038025177923625, 1e-14) && strings.Contains(a.CHistory.Classification, "Radial-Hessian Hopf") && !a.CHistory.FullIndependentPredictionComponent, Detail: FormatCHistory(a.CHistory)},
			{Name: "record prediction level classification", Passed: a.Prediction.Recorded && strings.Contains(a.Prediction.KappaLambdaRedLevel, "Level B") && strings.Contains(a.Prediction.CHistoryLevel, "Level B") && strings.Contains(a.Prediction.CHiggsLevel, "not Level C") && strings.Contains(a.Prediction.LevelC, "native ASHA") && strings.Contains(a.Prediction.NextBottleneck, "F_wall_3_red"), Detail: FormatPrediction(a.Prediction)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.KappaLambdaNativeScalarTheorem && !a.Firewalls.KappaLambdaFullyTheoremIndependent && !a.Firewalls.BoundaryResponseNativeTheorem && !a.Firewalls.RawMomentNativeTheorem && !a.Firewalls.CubicStopNativeTheorem && !a.Firewalls.KappaENativeTheorem && !a.Firewalls.PMNSOrCKMNativeTheorem && !a.Firewalls.CHistoryFullIndependentPrediction && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNativeTheorem && a.Firewalls.Verdict == StatusNoNativeBoundaryFlavorScalarMatchingComplementBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
