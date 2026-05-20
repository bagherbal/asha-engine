package generation2historyloopunitsourcetypeaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(i Gate623Inheritance) string {
	return fmt.Sprintf("L=%s flavor=%q scalar=%q kappaE=%s kappaLambda=%s scalarBridge=%t flavorBridge=%t nativeL=%t nativeCrossSeal=%t verdict=%q", f64(i.LoopUnit), i.FlavorNormalForm, i.ScalarNormalForm, f64(i.KappaE), f64(i.KappaLambda), i.ScalarBridgeOnly, i.FlavorBridgeOnly, i.NativeLTheorem, i.NativeCrossSeal, i.Verdict)
}

func FormatDecompositionRow(r LDecompositionRow) string {
	return fmt.Sprintf("expr=%q value=%s lane=%q typedObject=%q role=%q typed=%t native=%t arbitrary=%t", r.Expression, f64(r.Value), r.Lane, r.TypedObject, r.CandidateRole, r.Typed, r.NativeCertified, r.ArbitraryConstant)
}

func FormatDecompositionRows(rows []LDecompositionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDecompositionRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatDecompositions(t LDecompositionTable) string {
	return fmt.Sprintf("L=%s allValuesMatch=%t allTyped=%t noArbitrarySearch=%t rows=[%s] verdict=%q", f64(t.LoopUnit), t.AllValuesMatch, t.AllRowsTyped, t.NoArbitrarySearch, FormatDecompositionRows(t.Rows), t.Verdict)
}

func FormatHopfPhase(h HopfPhaseAudit) string {
	return fmt.Sprintf("gate570Hopf=%t gate570Reeb=%t gate572CP3=%t fiber=%q contact=%q reeb=%q quotient=%q measure=%q circle=%t quarterCandidate=%t quarterCertified=%t mapFlavor=%t mapScalar=%t physicalTime=%t verdict=%q", h.Gate570HopfS7Certified, h.Gate570ReebPhaseCertified, h.Gate572CP3Certified, h.HopfFiber, h.ContactForm, h.ReebAction, h.ProjectiveQuotient, h.NormalizedPhaseMeasure, h.CirclePhaseNormalization, h.QuarterProjectionCandidate, h.QuarterProjectionCertified, h.MapToFlavorWallCertified, h.MapToScalarMatchingCertified, h.PhysicalTimeClaimed, h.Verdict)
}

func FormatWeakQuarterCandidate(c WeakQuarterCandidate) string {
	return fmt.Sprintf("name=%q expr=%q typed=%t native=%t comment=%q", c.Name, c.Expression, c.Typed, c.NativeCertified, c.Comment)
}

func FormatWeakQuarterCandidates(rows []WeakQuarterCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatWeakQuarterCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatWeakQuarter(w WeakQuarterAudit) string {
	return fmt.Sprintf("factor=%s weakTyped=%t pmnsTyped=%t nativeL=%t nativeWeakQuarterLoop=%t candidates=[%s] verdict=%q", f64(w.Factor), w.WeakNormalizationTyped, w.PMNSOverlapTyped, w.NativeConnectionToL, w.NativeWeakQuarterLoopLaw, FormatWeakQuarterCandidates(w.Candidates), w.Verdict)
}

func FormatHeatKernelOperation(op HeatKernelOperation) string {
	return fmt.Sprintf("name=%q input=%q output=%q certified=%t comment=%q", op.Name, op.Input, op.Output, op.Certified, op.Comment)
}

func FormatHeatKernelOperations(rows []HeatKernelOperation) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatHeatKernelOperation(r))
	}
	return strings.Join(parts, " | ")
}

func FormatHeatKernel(h HeatKernelLoopFactorAudit) string {
	return fmt.Sprintf("L=%s fourDLoop=%s boundarySurface=%s anyCertified=%t operations=[%s] verdict=%q", f64(h.LoopUnit), f64(h.FourDLoopUnit), f64(h.BoundarySurfaceUnit), h.AnyCertifiedReduction, FormatHeatKernelOperations(h.Operations), h.Verdict)
}

func FormatScalarCandidate(c ScalarKappaCandidate) string {
	return fmt.Sprintf("name=%q value=%s residual=%s relResidual=%s typed=%t native=%t comment=%q", c.Name, f64(c.Value), f64(c.Residual), f64(c.RelativeResidual), c.Typed, c.NativeCertified, c.Comment)
}

func FormatScalarCandidates(rows []ScalarKappaCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatScalarCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatScalarRole(s ScalarRoleAudit) string {
	return fmt.Sprintf("proxy=%s runtime=%s rho=%s L=%s kappaLambda=%s normalForm=%q closest=%q closestResidual=%s sourceCertified=%t candidates=[%s] verdict=%q", f64(s.LambdaProxy), f64(s.LambdaRuntime), f64(s.RhoLambdaMatch), f64(s.LoopUnit), f64(s.KappaLambda), s.NormalForm, s.ClosestName, f64(s.ClosestResidual), s.KappaSourceCertified, FormatScalarCandidates(s.Candidates), s.Verdict)
}

func FormatFlavorRole(f FlavorRoleAudit) string {
	return fmt.Sprintf("epsilon=%s L=%s kappaE=%s sin2theta13Quarter=%s JCKM=%s orientation=%s epsOrientation=%s residual=%s normalForm=%q class=%q native=%t verdict=%q", f64(f.EpsilonE), f64(f.LoopUnit), f64(f.KappaE), f64(f.Sin2Theta13Quarter), f64(f.JCKM), f64(f.OrientationCandidate), f64(f.EpsilonOrientation), f64(f.Residual), f.NormalForm, f.Classification, f.NativeDerived, f.Verdict)
}

func FormatCrossSealRow(r CrossSealRow) string {
	return fmt.Sprintf("seal=%q base=%q correction=%q sign=%q residual=%q native=%q", r.Seal, r.BaseUnit, r.Correction, r.SignRole, r.Residual, r.NativeStatus)
}

func FormatCrossSealRows(rows []CrossSealRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCrossSealRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatCrossSeal(c CrossSealComparisonTable) string {
	return fmt.Sprintf("sharedLBridge=%t nativeCrossSeal=%t rows=[%s] verdict=%q", c.SharedLBridgeSeal, c.NativeCrossSeal, FormatCrossSealRows(c.Rows), c.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("nativeL=%t hopfFlavor=%t hopfScalar=%t heatKernelL=%t weakQuarter=%t crossSealOrientation=%t statement=%q verdict=%q", n.NativeLTheorem, n.NativeHopfToFlavorWallMap, n.NativeHopfToScalarMatchingMap, n.NativeHeatKernelToLReduction, n.NativeWeakQuarterLoopTheorem, n.NativeCrossSealOrientationLaw, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("koide=%t higgsMass=%t scalarStability=%t pmnsCkm=%t gaugeUnification=%t nativeLoop=%t physicalTime=%t verdict=%q", f.ClaimsKoideDerived, f.ClaimsHiggsMassDerived, f.ClaimsScalarStability, f.ClaimsPMNSCKMDerived, f.ClaimsGaugeUnification, f.ClaimsNativeLoopTheorem, f.ClaimsPhysicalTime, f.Verdict)
}
