package generation2chistorynativesourceindependenceandtransportlawaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CHistoryNativeSourceIndependenceAndTransportLawAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 781 — C_History Native-Source Independence and Transport-Law Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate781CHistoryNativeSourceBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 780 prediction independence audit", Passed: a.Gate780.Inherited && a.Gate780.CHiggsFormula == "C_Higgs=C_Yukawa C_History" && a.Gate780.CHistoryFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && strings.Contains(a.Gate780.PredictionStatus, "Level A/B") && closeRel(a.Gate780.CHistory, 1.038025177923625, 1e-15) && closeRel(a.Gate780.DeltaHistory, 0.03802517792362492, 1e-15) && closeRel(a.Gate780.EpsilonYukawa, 0.0007751811187991509, 1e-15) && a.Gate780.DominantCorrection, Detail: FormatGate780(a.Gate780)},
			{Name: "expand C_History dependency cluster", Passed: a.Cluster.Expanded && a.Cluster.Formula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && closeRel(a.Cluster.LHopf, 0.039788735772973836, 1e-15) && closeRel(a.Cluster.KappaLambdaRed, 0.04432304306956136, 1e-15) && closeRel(a.Cluster.Complement, 0.9556769569304386, 1e-15) && closeRel(a.Cluster.ComputedCHistory, 1.038025177923625, 1e-15) && a.Cluster.MatchesSnapshot && strings.Contains(a.Cluster.KappaLambdaRedFormula, "F_wall_3_red") && containsAll(a.Cluster.KappaLambdaRedComponents, []string{"|lambda(Lambda12)| boundary scalar wall coordinate", "F_wall_3_red(s) cubic boundary response polynomial", "kappa_e_red reduced flavor-wall input"}) && strings.Contains(a.Cluster.ClusterQuestion, "without scalar-runtime"), Detail: FormatCluster(a.Cluster)},
			{Name: "audit L_Hopf radial-Hessian Hopf source", Passed: a.LHopf.Audited && strings.Contains(a.LHopf.Formula, "supp(H_V(x0))") && a.LHopf.MaximumEntropyStateRequired && a.LHopf.HessianSupportProjectorRequired && a.LHopf.PhaseLoopPayoffRequired && !a.LHopf.HistoryEvaluatesEventTheorem && !a.LHopf.NativeTheorem && a.LHopf.ConditionalSourceTyping && closeRel(a.LHopf.EventWeight, 0.25, 1e-15) && closeRel(a.LHopf.PhasePayoff, 1/(2*3.141592653589793), 1e-15) && closeRel(a.LHopf.LHopf, 0.039788735772973836, 1e-15) && len(a.LHopf.MissingIngredients) == 3, Detail: FormatLHopf(a.LHopf)},
			{Name: "audit transport law form", Passed: a.Transport.Audited && a.Transport.TransportFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && strings.Contains(a.Transport.BaselineInterpretation, "untransported scalar baseline") && strings.Contains(a.Transport.LHopfInterpretation, "radial-Hessian Hopf") && a.Transport.BracketFormula == "1-kappa_lambda_red" && strings.Contains(a.Transport.BracketInterpretation, "scalar matching complement") && !a.Transport.NativeTransportLaw && a.Transport.BridgeReconstruction, Detail: FormatTransport(a.Transport)},
			{Name: "audit kappa_lambda_red runtime independence", Passed: a.Runtime.Audited && strings.Contains(a.Runtime.KappaLambdaRedFormula, "F_wall_3_red") && strings.Contains(a.Runtime.BoundaryWallCoordinate, "boundary scalar wall") && strings.Contains(a.Runtime.BoundaryResponsePolynomial, "deficit-closure") && strings.Contains(a.Runtime.FlavorWallReducedInput, "flavor") && a.Runtime.UsesLambdaRuntimeTarget && !a.Runtime.UsesTreeMassTarget && !a.Runtime.UsesPoleMassTarget && a.Runtime.UsesHiggsTargetClosure && !a.Runtime.CanBeEvaluatedWithoutRuntimeClosure && a.Runtime.ReducedButNotRuntimeIndependent && len(a.Runtime.IndependentComponents) >= 3 && len(a.Runtime.RuntimeDependentOrSealedComponents) >= 4, Detail: FormatRuntime(a.Runtime)},
			{Name: "record branch outcomes", Passed: a.Branches.Recorded && strings.Contains(a.Branches.StrongSuccess, "independent") && strings.Contains(a.Branches.PartialSuccess, "bottleneck") && strings.Contains(a.Branches.Failure, "bridge consistency") && a.Branches.SelectedOutcome == "Outcome 2 — partial success" && a.Branches.NextGate == "Gate 782 — Boundary-Flavor Scalar Matching Complement Independence Audit" && strings.Contains(a.Branches.Reason, "kappa_lambda_red"), Detail: FormatBranches(a.Branches)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.CHistoryNativePredictionComponent && !a.Firewalls.LHopfNativeHistoryLoopTheorem && !a.Firewalls.TransportLawNativeTheorem && !a.Firewalls.KappaLambdaNativeScalarTheorem && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNativeTheorem && a.Firewalls.Verdict == StatusGate781CHistoryNativeSourceBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
