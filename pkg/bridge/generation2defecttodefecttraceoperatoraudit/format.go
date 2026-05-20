package generation2defecttodefecttraceoperatoraudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate676Inheritance) string {
	return fmt.Sprintf("quotient=%t canonical=%t traceOnQuotient=%t lessTautological=%t tau=%.15g sSplit=%.15g dBase=%.15g residual=%.15g noTraceQuotient=%t no7=%t noWall=%t firewall=%t verdict=%q", x.BoundaryQuotientInherited, x.SplitIsCanonicalBoundaryQuotient, x.TraceActsOnBoundaryQuotient, x.LessTautologicalRoute, x.TauDefect, x.SSplit, x.DBase, x.Residual, x.NoNativeTraceBoundaryQuotient, x.NoNativeSevenOver72Theorem, x.NoNativeWallAirlockTheorem, x.FirewallPreserved, x.Verdict)
}

func FormatDomain(x DomainDefectLine) string {
	return fmt.Sprintf("name=%q space=%q quotient=%q coord=%q dim=%d sSplit=%.15g canonical=%t interp=%q verdict=%q", x.Name, x.Space, x.Quotient, x.Coordinate, x.Dimension, x.SSplit, x.CanonicalFromGate676, x.Interpretation, x.Verdict)
}

func FormatCodomain(x CodomainDefectLine) string {
	return fmt.Sprintf("name=%q space=%q coord=%q dim=%d dBase=%.15g interp=%q verdict=%q", x.Name, x.Space, x.Coordinate, x.Dimension, x.DBase, x.Interpretation, x.Verdict)
}

func FormatOperator(x TraceResponseOperator) string {
	return fmt.Sprintf("name=%q domain=%q codomain=%q formula=%q tau=%.15g trace=%d/%d linear=%t scalarFunctional=%t vectorMap=%t verdict=%q", x.Name, x.Domain, x.Codomain, x.Formula, x.TauDefect, x.TraceNumerator, x.TraceDenominator, x.Linear, x.ScalarFunctionalOnly, x.RequiresVectorMap, x.Verdict)
}

func FormatTest(x OperatorTest) string {
	return fmt.Sprintf("sSplit=%.15g dBase=%.15g tau=%.15g pred=%.15g residual=%.15g abs=%.15g relDBase=%.15g verdict=%q", x.SSplit, x.DBase, x.TauDefect, x.PredictedDBase, x.Residual, x.AbsResidual, x.RelativeToDBase, x.Verdict)
}

func FormatRequirement(x NonTautologyRequirement) string {
	return fmt.Sprintf("%s source=%q status=%q comment=%q", x.Requirement, x.SourceGate, x.Status, x.Comment)
}

func FormatRequirements(xs []NonTautologyRequirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatRequirement(x))
	}
	return strings.Join(parts, "; ")
}

func FormatCoupler(x CouplerCandidate) string {
	return fmt.Sprintf("%s status=%q class=%q comment=%q", x.Candidate, x.Status, x.Classification, x.Comment)
}

func FormatCouplers(xs []CouplerCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCoupler(x))
	}
	return strings.Join(parts, "; ")
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("targets=[%s] missing=[%s] support=[%s] verdict=%q", strings.Join(x.NativeTheoremTargets, "; "), strings.Join(x.MissingTheorems, "; "), strings.Join(x.AllowedSupport, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsCoupling=%t claimsOperator=%t claims7=%t claimsWall=%t claimsBoundary=%t claimsK7Map=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeTraceCouplesDefects, x.ClaimsNativeTraceOperator, x.ClaimsNativeSevenOver72, x.ClaimsNativeWallAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsFullK7BoundaryMap, x.ClaimsHiggsMassPrediction, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
