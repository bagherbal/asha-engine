package generation2chistorynativesourceindependenceandtransportlawaudit

import (
	"strings"
	"testing"
)

func TestGate781InheritanceAndCluster(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate780.Inherited || !a.Gate780.DominantCorrection || !closeRel(a.Gate780.CHistory, 1.038025177923625, 1e-15) || !strings.Contains(a.Gate780.PredictionStatus, "Level A/B") {
		t.Fatalf("bad Gate780 inheritance: %+v", a.Gate780)
	}
	if !a.Cluster.Expanded || a.Cluster.Formula != "C_History=1+L_Hopf(1-kappa_lambda_red)" || !a.Cluster.MatchesSnapshot || !closeRel(a.Cluster.Complement, 0.9556769569304386, 1e-15) || !strings.Contains(a.Cluster.KappaLambdaRedFormula, "kappa_e_red") || !containsAll(a.Cluster.KappaLambdaRedComponents, []string{"|lambda(Lambda12)| boundary scalar wall coordinate", "F_wall_3_red(s) cubic boundary response polynomial", "kappa_e_red reduced flavor-wall input"}) {
		t.Fatalf("bad cluster: %+v", a.Cluster)
	}
}

func TestGate781LHopfAndTransport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LHopf.Audited || !strings.Contains(a.LHopf.Formula, "1/(8*pi)") || !a.LHopf.MaximumEntropyStateRequired || !a.LHopf.HessianSupportProjectorRequired || !a.LHopf.PhaseLoopPayoffRequired || a.LHopf.HistoryEvaluatesEventTheorem || a.LHopf.NativeTheorem || !a.LHopf.ConditionalSourceTyping || len(a.LHopf.MissingIngredients) != 3 {
		t.Fatalf("bad L_Hopf audit: %+v", a.LHopf)
	}
	if !a.Transport.Audited || a.Transport.TransportFormula != "C_History=1+L_Hopf(1-kappa_lambda_red)" || a.Transport.NativeTransportLaw || !a.Transport.BridgeReconstruction || !strings.Contains(a.Transport.BracketInterpretation, "scalar matching complement") {
		t.Fatalf("bad transport audit: %+v", a.Transport)
	}
}

func TestGate781RuntimeBranchesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Runtime.Audited || !a.Runtime.UsesLambdaRuntimeTarget || a.Runtime.UsesTreeMassTarget || a.Runtime.UsesPoleMassTarget || !a.Runtime.UsesHiggsTargetClosure || a.Runtime.CanBeEvaluatedWithoutRuntimeClosure || !a.Runtime.ReducedButNotRuntimeIndependent || len(a.Runtime.RuntimeDependentOrSealedComponents) < 4 {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
	if !a.Branches.Recorded || a.Branches.SelectedOutcome != "Outcome 2 — partial success" || a.Branches.NextGate != "Gate 782 — Boundary-Flavor Scalar Matching Complement Independence Audit" {
		t.Fatalf("bad branch outcomes: %+v", a.Branches)
	}
	if !a.Firewalls.Enforced || a.Firewalls.CHistoryNativePredictionComponent || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.TransportLawNativeTheorem || a.Firewalls.KappaLambdaNativeScalarTheorem || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNativeTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate781TheoremStatuses(t *testing.T) {
	res := Generation2CHistoryNativeSourceIndependenceAndTransportLawAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
