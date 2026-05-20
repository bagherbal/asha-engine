package generation2historydefectreadoutfunctionalselectionaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate697InheritanceAndHistoryClosureWall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.BoundaryQuotientPayoffInherited || !strings.Contains(a.Inherited.BoundaryPayoff, "sigma_boundary") || !nearly(a.Inherited.BoundaryPayoffValue, lambdaLambda12+r3Minus1) || !a.Inherited.BoundaryPayoffDescends || a.Inherited.ResponseOperator != "R_split=sigma_boundary(b)P_K7" || a.Inherited.Rho72 != "rho_72=I_H72/72" || !a.Inherited.NoNativePayoffCoupling || !a.Inherited.NoNativeHistoryResponse || !a.Inherited.NoNativeSevenOver72 {
		t.Fatalf("bad Gate697 inheritance: %+v", a.Inherited)
	}
	if a.ClosureWall.WallEquation != "kappa_lambda+kappa_e+lambda=0" || !strings.Contains(a.ClosureWall.EquivalentPositiveClosure, "lambda<0") || !nearly(a.ClosureWall.KappaLambda, kappaLambda) || !nearly(a.ClosureWall.KappaE, kappaE) || !nearly(a.ClosureWall.Lambda, lambdaLambda12) || !a.ClosureWall.LambdaNegative || !nearly(a.ClosureWall.ClosureDefect, a.Readout.DBase) || a.ClosureWall.WallSatisfied {
		t.Fatalf("bad closure wall: %+v", a.ClosureWall)
	}
}

func TestSigmaHistoryReadout(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Readout
	if !strings.Contains(r.SigmaDefinition, "kappa_lambda+kappa_e+lambda") || r.InputSpace != "H_history=span(kappa_lambda,kappa_e,lambda)" || r.InputVector != [3]float64{kappaLambda, kappaE, lambdaLambda12} || r.DBaseExpression != "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)" || !nearly(r.DBase, kappaLambda+kappaE+lambdaLambda12) || !r.MatchesInheritedDBase || !r.MeasuresClosureDefect {
		t.Fatalf("bad readout: %+v", r)
	}
	for _, want := range []string{StatusSigmaHistoryReadoutDefined, StatusDBaseIdentifiedAsHistoryDefectQuotient, StatusDBaseCanonicalHistoryClosureDefectReadout} {
		if !strings.Contains(r.Verdict, want) {
			t.Fatalf("missing readout verdict %s in %q", want, r.Verdict)
		}
	}
}

func TestAlternativeHistoryReadoutsAndSignedWall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Alternatives.Alternatives) != 6 || !a.Alternatives.KSumRejected || !a.Alternatives.LambdaOnlyRejected || !a.Alternatives.KappaLambdaOnlyRejected || !a.Alternatives.KappaEOnlyRejected || !a.Alternatives.AbsoluteFormEquivalent || !a.Alternatives.SignedFormAccepted || !a.Alternatives.AllAudited {
		t.Fatalf("bad alternatives audit: %+v", a.Alternatives)
	}
	byName := map[string]HistoryReadoutAlternative{}
	for _, x := range a.Alternatives.Alternatives {
		byName[x.Name] = x
	}
	if byName["K_sum"].Active || byName["K_sum"].IncludesSignedLambda || byName["lambda-only"].IncludesKappaLambda || byName["lambda-only"].IncludesKappaE {
		t.Fatalf("bad rejected alternatives: %+v", byName)
	}
	if !byName["absolute closure form"].EquivalentWhenLambdaN || byName["absolute closure form"].PreservesOrientation || byName["absolute closure form"].Active {
		t.Fatalf("absolute form misclassified: %+v", byName["absolute closure form"])
	}
	if !byName["signed history form"].Active || !byName["signed history form"].IncludesSignedLambda || !byName["signed history form"].PreservesOrientation || !nearly(byName["signed history form"].Value, a.Readout.DBase) {
		t.Fatalf("signed form not accepted: %+v", byName["signed history form"])
	}
	if !a.SignAudit.LambdaNegative || a.SignAudit.SignedForm != "kappa_lambda+kappa_e+lambda" || a.SignAudit.AbsoluteForm != "kappa_lambda+kappa_e-|lambda|" || !nearly(a.SignAudit.DBaseSigned, a.SignAudit.DBaseAbsolute) || !a.SignAudit.FormsEquivalentNumerically || !a.SignAudit.SignedFormPreferred || !a.SignAudit.OrientationPreserved {
		t.Fatalf("bad sign audit: %+v", a.SignAudit)
	}
}

func TestBridgeReconstructionResidualAndSourceTypes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Bridge.LeftSide, "sigma_history") || !strings.Contains(a.Bridge.RightSide, "sigma_boundary") || !strings.Contains(a.Bridge.ExpandedEquation, "I_H72/72") || !nearly(a.Bridge.SigmaHistory, a.Readout.DBase) || !nearly(a.Bridge.SigmaBoundaryExpectation, a.Inherited.Expectation) || math.Abs(a.Bridge.ResidualE1-a.Inherited.ResidualE1) > 1e-17 || !a.Bridge.MatchesInheritedResidual || !a.Bridge.RelatesQuotientToPayoff {
		t.Fatalf("bad bridge reconstruction: %+v", a.Bridge)
	}
	if math.Abs(a.Residual.ResidualE1-a.Inherited.ResidualE1) > 1e-17 || !a.Residual.QuadraticClueRetained || a.Residual.QuadraticCluePromoted || a.Residual.NativeResidualExplanation {
		t.Fatalf("bad residual status: %+v", a.Residual)
	}
	if !strings.Contains(a.SourceTypes.KappaLambdaRole, "scalar") || !strings.Contains(a.SourceTypes.KappaERole, "flavor") || !strings.Contains(a.SourceTypes.LambdaRole, "signed") || !strings.Contains(a.SourceTypes.DBaseRole, "closure-defect") || !strings.Contains(a.SourceTypes.SSplitRole, "boundary") || !strings.Contains(a.SourceTypes.Rho72Role, "maximum-entropy") || !strings.Contains(a.SourceTypes.PK7Role, "Boolean-octonionic") {
		t.Fatalf("bad source types: %+v", a.SourceTypes)
	}
}

func TestMissingTheoremFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Missing.Candidates) != 6 || !strings.Contains(a.Missing.PreciseGap, "sigma_history") || !strings.Contains(a.Missing.PreciseGap, "sigma_boundary") || !strings.Contains(a.Missing.Verdict, StatusNoNativeExpectedK7PayoffEqualsHistoryDefect) || !strings.Contains(a.Missing.Verdict, StatusNoNativeHistoryBoundaryResponseTheorem) || !strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Firewall.ClaimsExpectedK7PayoffEqualsHistoryDefectNatively || a.Firewall.ClaimsNativeHistoryBoundaryResponseTheorem || a.Firewall.ClaimsNativeSevenOver72Theorem || a.Firewall.ClaimsNativePayoffTheorem || a.Firewall.ClaimsBoundaryStressDerived || a.Firewall.ClaimsScalarRGMatching || a.Firewall.ClaimsHiggsMass || a.Firewall.ClaimsGaugeUnification || a.Firewall.ClaimsFlavorDerivation || a.Firewall.ClaimsCKMPMNS || a.Firewall.Verdict != StatusGate698HistoryDefectReadoutBoundary {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
	res := Generation2HistoryDefectReadoutFunctionalSelectionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
