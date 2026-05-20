package generation2augmentedchamberdefecttraceresponseaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate673Inheritance) string {
	return fmt.Sprintf("inherited=%t splitLine=%t baseLine=%t coeff=%t fullK7Failed=%t noStress=%t no7=%t noBoundary=%t firewall=%t dBase=%.15g sSplit=%.15g qPull=%.15g residual=%.15g verdict=%q", x.InheritedLinePullback, x.BoundarySplitLineDefined, x.BaseDefectLineDefined, x.PullbackCoefficientComputed, x.FullK7BoundaryMapFailed, x.NoNativeStressSplitTheorem, x.NoNativeSevenOver72Theorem, x.NoBoundaryStressDerivation, x.FirewallPreserved, x.DBase, x.SSplit, x.QPull, x.SevenOver72Residual, x.Verdict)
}

func FormatChamber(x AugmentedChamberAudit) string {
	return fmt.Sprintf("lambda4=%d boundary=%d total=%d pair=%q split=%q traceWeight=%.15g verdict=%q", x.Lambda4Dimension, x.BoundaryDimension, x.TotalDimension, x.BoundaryPair, x.SplitLine, x.TraceWeight, x.Verdict)
}

func FormatRankSeven(x RankSevenDefectSourceAudit) string {
	return fmt.Sprintf("dimK7=%d ker=%d coker=%d fano=%d numerator=%d firewall=%q sources=[%s] verdict=%q", x.DimK7, x.DimKernelA, x.DimCokernelA, x.FanoHitchinCarrierDimension, x.NumeratorCandidate, x.FanoHitchinBoundaryFirewall, strings.Join(x.CandidateSources, "; "), x.Verdict)
}

func FormatTrace(x ScalarTraceResponseAudit) string {
	return fmt.Sprintf("dBase=%.15g sSplit=%.15g qPull=%.15g qTrace=%.15g pullback=%.15g residual=%.15g requiresVector=%t requiresScalarTrace=%t interpretation=%q verdict=%q", x.DBase, x.SSplit, x.QPull, x.QTrace, x.TracePullback, x.TraceResidual, x.RequiresVectorMap, x.RequiresScalarTraceMap, x.Interpretation, x.Verdict)
}

func FormatAlternative(x DenominatorAlternative) string {
	return fmt.Sprintf("%s weight=%.15g pullback=%.15g residual=%.15g abs=%.15g typing=%q class=%q", x.Name, x.Weight, x.Pullback, x.Residual, x.AbsResidual, x.Typing, x.Classification)
}

func FormatAlternatives(x DenominatorAlternativeAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		parts = append(parts, FormatAlternative(a))
	}
	return fmt.Sprintf("best=%q bestWeight=%.15g bestResidual=%.15g alternatives=[%s] verdict=%q", x.BestName, x.BestWeight, x.BestResidual, strings.Join(parts, "; "), x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("targets=[%s] missing=[%s] support=[%s] verdict=%q", strings.Join(x.NativeTheoremTargets, "; "), strings.Join(x.MissingTheorems, "; "), strings.Join(x.AllowedSupport, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsTrace=%t claimsStress=%t claims7=%t claimsK7Map=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeTraceResponse, x.ClaimsNativeStressSplitPullback, x.ClaimsNativeSevenOver72, x.ClaimsFullK7BoundaryMap, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
