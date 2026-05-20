package generation2boundaryflavorscalarmatchingcomplementindependenceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate782RewriteAndLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate781.Inherited || !strings.Contains(a.Gate781.SelectedBranch, "Outcome 2") || !strings.Contains(a.Gate781.SelectedBottleneck, "kappa_lambda_red") {
		t.Fatalf("bad Gate781 inheritance: %+v", a.Gate781)
	}
	if !a.Rewrite.SignsChecked || !strings.Contains(a.Rewrite.FactoredFormula, "-(1-p s^2)kappa_e_red") || !strings.Contains(a.Rewrite.ExpandedFormula, "xi_boundary p s^2") || !a.Rewrite.NoRuntimeSymbols {
		t.Fatalf("bad rewrite: %+v", a.Rewrite)
	}
	if !closeRel(a.Ledger.M1, 0.0001256543573849177, 1e-14) || !closeRel(a.Ledger.M2, 1.624013231638281e-07, 1e-14) || !closeRel(a.Ledger.M3, 2.0989474869200057e-10, 1e-14) {
		t.Fatalf("bad moments: %+v", a.Ledger)
	}
	if !closeRel(a.Ledger.FWall3Red, 0.00012565521035653708, 1e-14) || !closeRel(a.Ledger.KappaLambdaRed, 0.04432304306956136, 1e-14) || !closeRel(a.Ledger.Complement, 0.9556769569304386, 1e-14) || !closeRel(a.Ledger.CHistory, 1.038025177923625, 1e-14) {
		t.Fatalf("bad complement ledger: %+v", a.Ledger)
	}
	if !a.Ledger.MatchesFWall || !a.Ledger.MatchesKappa || !a.Ledger.MatchesComplement || !a.Ledger.MatchesCHistory || !strings.Contains(a.Ledger.DiscrepancyClass, "no material discrepancy") {
		t.Fatalf("bad ledger matches: %+v", a.Ledger)
	}
}

func TestGate782TermTypingK7AndRawMomentFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Typing.Typed || len(a.Typing.TermTypes) != 8 || !strings.Contains(a.Typing.TermTypes["|lambda|"], "scalar wall") || !strings.Contains(a.Typing.LayerTypes["bridge raw-moment response"], "F_wall_3_red") || a.Typing.RuntimeTargetInFinalFormula {
		t.Fatalf("bad typing: %+v", a.Typing)
	}
	if !a.K7.Audited || !a.K7.AppearsOnlyAsPK7 || !a.K7.NativeSupportOnly || a.K7.BoundaryVector || a.K7.FlavorOperator || a.K7.ScalarWallCoordinate || a.K7.SourceOfLHopf || a.K7.HyperchargeNormalization || a.K7.YukawaTheorem {
		t.Fatalf("bad K7 audit: %+v", a.K7)
	}
	if !a.RawMoment.Audited || !a.RawMoment.BridgeLayer || a.RawMoment.NativeGeneratingFunction || a.RawMoment.NativeRawMomentCoordinate || a.RawMoment.NativeCubicStop || !a.RawMoment.M4ForbiddenWithoutTypedSource || !strings.Contains(a.RawMoment.M3Interpretation, "double-K7") {
		t.Fatalf("bad raw moment audit: %+v", a.RawMoment)
	}
}

func TestGate782FlavorRuntimePredictionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flavor.Audited || !strings.Contains(a.Flavor.Classification, "external flavor bridge") || a.Flavor.Theta13Native || a.Flavor.JCKMNative || a.Flavor.NativeKappaETheorem || a.Flavor.NativePMNSOrCKMTheorem || a.Flavor.NativeYukawaTheorem || a.Flavor.CanCompareOlderKappaE || !math.IsNaN(a.Flavor.OlderKappaEResidual) {
		t.Fatalf("bad flavor audit: %+v", a.Flavor)
	}
	if !a.Runtime.Audited || !a.Runtime.FormulaLevelRuntimeTargetAbsence || !a.Runtime.EvaluableWithoutDirectHiggsRuntimeVariables || a.Runtime.UsesLambdaRuntime || a.Runtime.UsesLambdaRuntimeEff || a.Runtime.UsesTreeMass || a.Runtime.UsesPoleMass || a.Runtime.UsesCHiggs || a.Runtime.UsesGF || a.Runtime.UsesVEV || a.Runtime.NativeDerivation || a.Runtime.RawBoundaryResponseIndependentlyProved || a.Runtime.FlavorInputsIndependentlyProved || a.Runtime.BoundaryCoordinatesNative {
		t.Fatalf("bad runtime audit: %+v", a.Runtime)
	}
	if !a.CHistory.Audited || !strings.Contains(a.CHistory.ExpandedForm, "1-|lambda|-p s+2p^2s^3") || a.CHistory.FullIndependentPredictionComponent || !closeRel(a.CHistory.CHistory, 1.038025177923625, 1e-14) {
		t.Fatalf("bad C_History rebuild: %+v", a.CHistory)
	}
	if !a.Prediction.Recorded || !strings.Contains(a.Prediction.KappaLambdaRedLevel, "Level B") || !strings.Contains(a.Prediction.CHistoryLevel, "Level B") || !strings.Contains(a.Prediction.CHiggsLevel, "not Level C") {
		t.Fatalf("bad prediction classification: %+v", a.Prediction)
	}
	if !a.Firewalls.Enforced || a.Firewalls.KappaLambdaNativeScalarTheorem || a.Firewalls.KappaLambdaFullyTheoremIndependent || a.Firewalls.BoundaryResponseNativeTheorem || a.Firewalls.RawMomentNativeTheorem || a.Firewalls.CubicStopNativeTheorem || a.Firewalls.KappaENativeTheorem || a.Firewalls.PMNSOrCKMNativeTheorem || a.Firewalls.CHistoryFullIndependentPrediction || a.Firewalls.TreeProxyPoleMass || a.Firewalls.YukawaNativeTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "removes direct runtime/Higgs target variables") || !strings.Contains(a.FinalStatement, "does not make kappa_lambda_red native") || !strings.Contains(a.FinalStatement, "next bottleneck") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
}

func TestGate782TheoremStatuses(t *testing.T) {
	res := Generation2BoundaryFlavorScalarMatchingComplementIndependenceAuditTheorem().Verify()
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
