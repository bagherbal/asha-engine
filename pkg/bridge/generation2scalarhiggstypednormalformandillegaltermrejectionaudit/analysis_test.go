package generation2scalarhiggstypednormalformandillegaltermrejectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate751TypedNormalFormConstruction(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate750.Inherited || !a.Gate750.TypedLedgerReady || !a.Gate750.TraceAirlocksReady || !a.Gate750.FirewallsPreserved {
		t.Fatalf("bad Gate750 inheritance: %+v", a.Gate750)
	}
	if !strings.Contains(a.Domains.BoundaryQuotientLine, "Q_boundary") || !strings.Contains(a.Domains.HistoryReadoutLine, "Q_history") || !strings.Contains(a.Domains.ScalarRuntimeLine, "Q_runtime") {
		t.Fatalf("bad typed domains: %+v", a.Domains)
	}
	if a.Boundary.CoordinateName != "S_split" || a.Boundary.LivesIn != "Q_boundary" || math.Abs(a.Boundary.SValue-0.0012924448188162962) > 1e-18 {
		t.Fatalf("bad boundary coordinate: %+v", a.Boundary)
	}
	if !a.K7Response.NotTensorProduct || !a.K7Response.NotBoundaryMap || !strings.Contains(a.K7Response.ResponseOperator, "sP_7") {
		t.Fatalf("bad K7 response operator: %+v", a.K7Response)
	}
	if math.Abs(a.Moments.EventWeight-pK7) > 1e-18 || math.Abs(a.Moments.M1-pK7*a.Boundary.SValue) > 1e-20 || !strings.Contains(a.Moments.Formula, "s^n") {
		t.Fatalf("bad raw moment map: %+v", a.Moments)
	}
	if !a.Cubic.NotOperatorOnK7 || !strings.Contains(a.Cubic.MapType, "Q_boundary -> Q_history") || !strings.Contains(a.Cubic.Polynomial, "-2p_K7^2") {
		t.Fatalf("bad cubic normal form: %+v", a.Cubic)
	}
	if !a.Hopf.ScalarAfterTrace || math.Abs(a.Hopf.LHopf-1/(8*math.Pi)) > 1e-18 || !strings.Contains(a.Hopf.OperatorLane, "K7+") {
		t.Fatalf("bad Hopf loop factor: %+v", a.Hopf)
	}
	if !a.NormalForm.AllOperatorCollapsed || !strings.Contains(a.NormalForm.RuntimeFormula, "lambda_proxy") || !strings.Contains(a.NormalForm.ExpandedFormula, "F_wall_3") {
		t.Fatalf("bad typed normal form: %+v", a.NormalForm)
	}
}

func TestGate751LegalIllegalAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Legal.LawfulAdditions) != 3 || len(a.Legal.LawfulProducts) != 4 || len(a.Legal.TraceMaps) != 3 || len(a.Legal.RuntimeScalars) != 4 {
		t.Fatalf("bad legal operation audit: %+v", a.Legal)
	}
	if len(a.Illegal.RejectedTerms) != 9 || !a.Illegal.K7BoundaryBlocked || !a.Illegal.FWallOperatorBlocked || !a.Illegal.LBoundaryCoeffBlocked || !a.Illegal.SevenLoopBlocked || !a.Illegal.TreePoleBlocked || !a.Illegal.PredictionBlocked {
		t.Fatalf("bad illegal term rejection: %+v", a.Illegal)
	}
	joinedRejected := strings.Join(a.Illegal.RejectedTerms, "\n")
	for _, want := range []string{"K7 + boundary vector", "P_K7 + S_split", "F_wall_3 as native operator on K7", "7/72 as source of 1/(8*pi)", "tree proxy as pole mass"} {
		if !strings.Contains(joinedRejected, want) {
			t.Fatalf("missing rejected term %q in %s", want, joinedRejected)
		}
	}
	if !a.KappaE.InsideFWall3 || !a.KappaE.OutsideRuntimeTransport || a.KappaE.NativeFlavorTheorem || !strings.Contains(a.KappaE.CandidateFormula, "sin²") {
		t.Fatalf("bad kappa_e insertion status: %+v", a.KappaE)
	}
	res := Generation2ScalarHiggsTypedNormalFormAndIllegalTermRejectionAuditTheorem().Verify()
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
