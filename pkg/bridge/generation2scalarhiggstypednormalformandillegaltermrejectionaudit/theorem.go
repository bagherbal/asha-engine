package generation2scalarhiggstypednormalformandillegaltermrejectionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarHiggsTypedNormalFormAndIllegalTermRejectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 751 — Scalar-Higgs Typed Normal Form and Illegal-Term Rejection Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate751 typed normal-form audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate750 typed ledger", Passed: a.Gate750.Inherited && a.Gate750.TypedLedgerReady && a.Gate750.TraceAirlocksReady && a.Gate750.FirewallsPreserved && strings.Contains(a.Gate750.Verdict, StatusGate750TypeLedgerInherited), Detail: FormatGate750(a.Gate750)},
			{Name: "type scalar-Higgs domains", Passed: strings.Contains(a.Domains.Lambda4Chamber, "Lambda^4") && strings.Contains(a.Domains.BoundaryQuotientLine, "Q_boundary") && strings.Contains(a.Domains.ScalarRuntimeLine, "Q_runtime"), Detail: FormatDomains(a.Domains)},
			{Name: "type boundary quotient coordinate", Passed: a.Boundary.CoordinateName == "S_split" && a.Boundary.LivesIn == "Q_boundary" && math.Abs(a.Boundary.SValue-0.0012924448188162962) < 1e-18, Detail: FormatBoundary(a.Boundary)},
			{Name: "type K7 response operator in H72", Passed: strings.Contains(a.K7Response.LiftedProjector, "End(H72)") && a.K7Response.NotTensorProduct && a.K7Response.NotBoundaryMap && strings.Contains(a.K7Response.Verdict, StatusK7ResponseOperatorTyped), Detail: FormatK7Response(a.K7Response)},
			{Name: "type raw moment map", Passed: math.Abs(a.Moments.EventWeight-pK7) < 1e-18 && math.Abs(a.Moments.M1-pK7*a.Boundary.SValue) < 1e-20 && strings.Contains(a.Moments.Formula, "p_K7"), Detail: FormatMoments(a.Moments)},
			{Name: "define F_wall_3 as Q_boundary to Q_history scalar response", Passed: a.Cubic.NotOperatorOnK7 && strings.Contains(a.Cubic.MapType, "Q_boundary -> Q_history") && strings.Contains(a.Cubic.Verdict, StatusFWall3QBoundaryQHistoryScalarResponse), Detail: FormatCubic(a.Cubic)},
			{Name: "type L_Hopf as K7+ trace expectation", Passed: a.Hopf.ScalarAfterTrace && strings.Contains(a.Hopf.HopfOperator, "End(K7+)") && math.Abs(a.Hopf.LHopf-1/(8*math.Pi)) < 1e-18, Detail: FormatHopf(a.Hopf)},
			{Name: "write scalar-Higgs typed normal form", Passed: a.NormalForm.AllOperatorCollapsed && strings.Contains(a.NormalForm.ExpandedFormula, "Tr_K7+") && strings.Contains(a.NormalForm.Verdict, StatusRuntimeAfterTraceCollapse), Detail: FormatNormalForm(a.NormalForm)},
			{Name: "audit legal operations", Passed: len(a.Legal.LawfulAdditions) == 3 && len(a.Legal.LawfulProducts) == 4 && len(a.Legal.TraceMaps) == 3 && len(a.Legal.RuntimeScalars) == 4 && strings.Contains(a.Legal.Verdict, StatusLegalOperationAuditCompleted), Detail: FormatLegal(a.Legal)},
			{Name: "reject illegal cross-type terms", Passed: len(a.Illegal.RejectedTerms) == 9 && a.Illegal.K7BoundaryBlocked && a.Illegal.FWallOperatorBlocked && a.Illegal.LBoundaryCoeffBlocked && a.Illegal.SevenLoopBlocked && a.Illegal.TreePoleBlocked && a.Illegal.PredictionBlocked && strings.Contains(a.Illegal.Verdict, StatusSevenOver72NotSourceOfOneOver8Pi), Detail: FormatIllegal(a.Illegal)},
			{Name: "record kappa_e insertion as optional non-native bridge substitution", Passed: a.KappaE.InsideFWall3 && a.KappaE.OutsideRuntimeTransport && !a.KappaE.NativeFlavorTheorem && strings.Contains(a.KappaE.CandidateFormula, "xi_boundary") && strings.Contains(a.KappaE.Verdict, StatusKappaESubstitutionNotNativeFlavorTheorem), Detail: FormatKappaE(a.KappaE)},
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
