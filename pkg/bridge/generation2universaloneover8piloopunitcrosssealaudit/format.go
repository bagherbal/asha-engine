package generation2universaloneover8piloopunitcrosssealaudit

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

func FormatScalarInherited(s ScalarInherited) string {
	return fmt.Sprintf("proxy=%s runtime=%s delta=%s rho=%s L=%s ansatz=%s ansatzMinusRuntime=%s relAnsatz=%s previousRelative=%q previousDiagnostic=%q verdict=%q", f64(s.LambdaProxy), f64(s.LambdaRuntime), f64(s.DeltaLambdaMatch), f64(s.RhoLambdaMatch), f64(s.LoopUnit), f64(s.LambdaAnsatz), f64(s.AnsatzMinusRuntime), f64(s.RelativeAnsatzResidual), s.PreviousRelativeVerdict, s.PreviousDiagnosticVerdict, s.Verdict)
}

func FormatFlavorInherited(f FlavorInherited) string {
	return fmt.Sprintf("epsilon=%s L=%s kappa=%s orientationCandidate=%s orientationResidual=%s epsRaw=%s epsOrient=%s rawResidual=%s orientResidual=%s verdict=%q", f64(f.EpsilonE), f64(f.LoopUnit), f64(f.KappaE), f64(f.OrientationCandidate), f64(f.OrientationResidual), f64(f.EpsilonRawL), f64(f.EpsilonOrientation), f64(f.RawLResidual), f64(f.OrientationEpsilonResidual), f.Verdict)
}

func FormatNormalForm(n SharedLoopUnitNormalForm) string {
	return fmt.Sprintf("L=%s flavor=%q scalar=%q kappaE=%s kappaLambda=%s verdict=%q", f64(n.LoopUnit), n.FlavorEquation, n.ScalarEquation, f64(n.FlavorKappaE), f64(n.ScalarKappaLambda), n.Verdict)
}

func FormatKappaCandidate(c KappaCandidate) string {
	return fmt.Sprintf("name=%q value=%s residual=%s relResidual=%s typed=%t native=%t comment=%q", c.Name, f64(c.Value), f64(c.Residual), f64(c.RelativeResidual), c.Typed, c.NativeCertified, c.Comment)
}

func FormatKappaCandidates(rows []KappaCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatKappaCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatKappaComparison(k KappaComparisonTable) string {
	return fmt.Sprintf("kappaE=%s kappaLambda=%s delta=%s closest=%q closestDelta=%s candidates=[%s] verdict=%q", f64(k.KappaE), f64(k.KappaLambda), f64(k.Delta), k.ClosestName, f64(k.ClosestDelta), FormatKappaCandidates(k.Candidates), k.Verdict)
}

func FormatScalarQuality(s ScalarAnsatzQuality) string {
	return fmt.Sprintf("proxy=%s L=%s ansatz=%s runtime=%s ansatzMinusRuntime=%s rel=%s massAnsatz=%s massRuntime=%s deltaMass=%s diagnosticOnly=%t verdict=%q", f64(s.LambdaProxy), f64(s.LoopUnit), f64(s.LambdaAnsatz), f64(s.LambdaRuntime), f64(s.AnsatzMinusRuntime), f64(s.RelativeRuntimeResidual), f64(s.MassAnsatzGeV), f64(s.MassRuntimeGeV), f64(s.DeltaMassGeV), s.DiagnosticOnly, s.Verdict)
}

func FormatFlavorQuality(f FlavorAnsatzQuality) string {
	return fmt.Sprintf("L=%s epsilon=%s raw=%s rawResidual=%s rawRel=%s orientationCandidate=%s epsOrient=%s orientResidual=%s orientRel=%s improvement=%s verdict=%q", f64(f.LoopUnit), f64(f.EpsilonE), f64(f.EpsilonRawL), f64(f.RawResidual), f64(f.RawRelativeResidual), f64(f.OrientationCandidate), f64(f.EpsilonOrientation), f64(f.OrientationResidual), f64(f.OrientationRelativeResidual), f64(f.ResidualImprovementFactor), f.Verdict)
}

func FormatSignAndRole(s SignAndRoleAudit) string {
	return fmt.Sprintf("flavorBelowL=%t scalarAboveProxy=%t oppositeSigns=%t native=%t statement=%q verdict=%q", s.FlavorUsesBelowL, s.ScalarUsesAboveProxy, s.OppositeSigns, s.NativeTheoremClaimed, s.Statement, s.Verdict)
}

func FormatCrossSealType(c CrossSealTypeAudit) string {
	return fmt.Sprintf("type=%q roles=%q disallowed=%q bridgeOnly=%t verdict=%q", c.LikelyType, strings.Join(c.AllowedRoles, ", "), strings.Join(c.DisallowedPromotions, ", "), c.BridgeOnly, c.Verdict)
}

func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("L=%t scalarMatching=%t koideWall=%t crossSeal=%t orientationBalance=%t higgsPole=%t statement=%q verdict=%q", n.NativeOneOver8PiTheorem, n.NativeScalarMatchingTheorem, n.NativeKoideWallTheorem, n.NativeCrossSealTheorem, n.NativeOrientationBalance, n.NativeHiggsPoleTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("koide=%t higgsMass=%t scalarStability=%t pmnsCkm=%t unification=%t nativeLoop=%t verdict=%q", f.ClaimsKoideDerived, f.ClaimsHiggsMassDerived, f.ClaimsScalarStability, f.ClaimsPMNSCKMDerived, f.ClaimsGaugeUnification, f.ClaimsNativeLoopTheorem, f.Verdict)
}
