package generation2defecttodefecttraceoperatoraudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2DefectToDefectTraceCouplingOperatorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 677 — Defect-to-Defect Trace Coupling Operator Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 677 — Defect-to-Defect Trace Coupling Operator Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate676 boundary quotient", Passed: a.Inherited.BoundaryQuotientInherited && a.Inherited.SplitIsCanonicalBoundaryQuotient && a.Inherited.TraceActsOnBoundaryQuotient && a.Inherited.LessTautologicalRoute && a.Inherited.FirewallPreserved && a.Inherited.Verdict == StatusGate676BoundaryQuotientInherited, Detail: FormatInherited(a.Inherited)},
			{Name: "define domain defect line", Passed: a.Domain.Dimension == 1 && a.Domain.CanonicalFromGate676 && math.Abs(a.Domain.SSplit-a.Inherited.SSplit) < 1e-15 && a.Domain.Verdict == StatusDomainDefectLineDefined, Detail: FormatDomain(a.Domain)},
			{Name: "define codomain defect line", Passed: a.Codomain.Dimension == 1 && math.Abs(a.Codomain.DBase-a.Inherited.DBase) < 1e-15 && a.Codomain.Verdict == StatusCodomainDefectLineDefined, Detail: FormatCodomain(a.Codomain)},
			{Name: "define trace response operator", Passed: a.Operator.Linear && a.Operator.ScalarFunctionalOnly && !a.Operator.RequiresVectorMap && math.Abs(a.Operator.TauDefect-7.0/72.0) < 1e-15 && strings.Contains(a.Operator.Verdict, StatusTraceResponseOperatorDefined) && strings.Contains(a.Operator.Verdict, StatusDefectToDefectLinearResponseForm), Detail: FormatOperator(a.Operator)},
			{Name: "compute operator residual", Passed: math.Abs(a.Test.PredictedDBase-((7.0/72.0)*a.Test.SSplit)) < 1e-15 && math.Abs(a.Test.Residual-8.52583439801e-10) < 1e-14 && a.Test.AbsResidual < 1e-8 && a.Test.Verdict == StatusOperatorResidualComputed, Detail: FormatTest(a.Test)},
			{Name: "restate non-tautology requirements", Passed: len(a.Requirements) == 5 && a.Requirements[0].Status == "supplied" && a.Requirements[3].Status == "missing theorem" && a.Requirements[4].Status == "partially supplied", Detail: FormatRequirements(a.Requirements)},
			{Name: "audit coupler candidates", Passed: len(a.Couplers) == 6 && a.Couplers[0].Candidate == "augmented chamber trace-response" && a.Couplers[1].Status == "missing theorem" && a.Couplers[5].Status == "conditional support", Detail: FormatCouplers(a.Couplers)},
			{Name: "record missing theorem targets", Passed: len(a.Missing.NativeTheoremTargets) == 3 && len(a.Missing.MissingTheorems) == 4 && strings.Contains(a.Missing.Verdict, StatusNoNativeTraceCouplesDefects) && strings.Contains(a.Missing.Verdict, StatusNoNativeTraceResponseOperatorTheorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeTraceCouplesDefects && !a.Discipline.ClaimsNativeTraceOperator && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsNativeWallAirlock && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsFullK7BoundaryMap && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate677Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 677 — Defect-to-Defect Trace Coupling Operator Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
