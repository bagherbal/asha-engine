package generation2boundaryantialignmentquotienttracecouplingaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate675Inheritance) string {
	return fmt.Sprintf("traceInherited=%t functional=%t tau=%.15g dBase=%.15g sSplit=%.15g residual=%.15g missingActsOnLine=%t noTrace=%t noWall=%t no7=%t firewall=%t verdict=%q", x.TraceResponseCandidateInherited, x.AugmentedTraceFunctionalDefined, x.TauDefect, x.DBase, x.SSplit, x.Residual, x.MissingReasonTraceActsOnLine, x.NoNativeTraceResponseTheorem, x.NoNativeWallAirlockTheorem, x.NoNativeSevenOver72Theorem, x.FirewallPreserved, x.Verdict)
}

func FormatBoundaryPlane(x BoundaryPlaneAudit) string {
	return fmt.Sprintf("plane=%q coords=[%s] vector=(%.15g,%.15g) dim=%d use=%q verdict=%q", x.Plane, strings.Join(x.Coordinates, ","), x.Vector[0], x.Vector[1], x.Dimension, x.WallCoordinateUse, x.Verdict)
}

func FormatAntiAlignment(x AntiAlignmentAudit) string {
	return fmt.Sprintf("constraint=%q line=%q generator=(%.15g,%.15g) antiVector=(%.15g,%.15g) sigmaOnGenerator=%.15g kernel=%t interpretation=%q verdict=%q", x.Constraint, x.AntiAlignmentLine, x.AntiAlignmentGenerator[0], x.AntiAlignmentGenerator[1], x.AntiAlignmentVector[0], x.AntiAlignmentVector[1], x.SigmaOnAntiAlignmentVector, x.IsInKernelOfSigma, x.Interpretation, x.Verdict)
}

func FormatQuotient(x QuotientFunctionalAudit) string {
	return fmt.Sprintf("functional=%q vector=(%.15g,%.15g) kernel=%q quotient=%q sSplit=%.15g sigma(b)=%.15g canonical=%t verdict=%q", x.Functional, x.FunctionalVector[0], x.FunctionalVector[1], x.Kernel, x.QuotientSpace, x.SSplit, x.SigmaBoundaryVector, x.CanonicalCokernelDefect, x.Verdict)
}

func FormatBaseDefect(x BaseDefectLineAudit) string {
	return fmt.Sprintf("dBase=%.15g kappaLambda=%.15g kappaE=%.15g lambda=%.15g equation=%q interpretation=%q verdict=%q", x.DBase, x.KappaLambda, x.KappaE, x.Lambda, x.DefectEquation, x.Interpretation, x.Verdict)
}

func FormatCoupling(x TraceCouplingAudit) string {
	return fmt.Sprintf("tau=%.15g dBase=%.15g sSplit=%.15g pred=%.15g residual=%.15g abs=%.15g qPull=%.15g scalarFunctional=%t vectorMap=%t verdict=%q", x.TauDefect, x.DBase, x.SSplit, x.PredictedDBase, x.Residual, x.AbsResidual, x.QPull, x.RequiresScalarFunctional, x.RequiresVectorBoundaryMap, x.Verdict)
}

func FormatUpgrade(x NonTautologyUpgradeAudit) string {
	return fmt.Sprintf("gate675=%q upgrade=%q stillMissing=%q lessTautological=%t promotable=%t verdict=%q", x.Gate675Problem, x.Gate676Upgrade, x.StillMissing, x.LessTautological, x.PromotableToTheorem, x.Verdict)
}

func FormatSource(x SourceCandidate) string {
	return fmt.Sprintf("%s status=%q class=%q comment=%q", x.Candidate, x.Status, x.Classification, x.Comment)
}

func FormatSources(xs []SourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatSource(x))
	}
	return strings.Join(parts, "; ")
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("targets=[%s] missing=[%s] support=[%s] verdict=%q", strings.Join(x.NativeTheoremTargets, "; "), strings.Join(x.MissingTheorems, "; "), strings.Join(x.AllowedSupport, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsTraceQuotient=%t claims7=%t claimsWall=%t claimsBoundary=%t claimsK7Map=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeTraceBoundaryQuotient, x.ClaimsNativeSevenOver72, x.ClaimsNativeWallAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsFullK7BoundaryMap, x.ClaimsHiggsMassPrediction, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
