package generation2strongthresholdsignfieldcontentviabilityaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(i InheritedGate608) string {
	return fmt.Sprintf("Lambda12=%s Lambda13=%s Lambda23=%s deltaU=%s deltaAlpha=%s DeltaB3=%s b3SM=%s b3Eff=%s relative=%s verdict=%q", f64(i.Lambda12GeV), f64(i.Lambda13GeV), f64(i.Lambda23GeV), f64(i.Delta3ThresholdRequired), f64(i.DeltaAlpha3InvRequired), f64(i.DeltaB3Required), f64(i.B3SM), f64(i.B3EffectiveDiagnostic), f64(i.RelativeB3Deformation), i.Verdict)
}

func FormatSignConventionRow(r SignConventionRow) string {
	return fmt.Sprintf("statement=%q equation=%q requiredSign=%q interpretation=%q verdict=%q", r.Statement, r.Equation, r.RequiredSign, r.Interpretation, r.Verdict)
}
func FormatSignConventions(rows []SignConventionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSignConventionRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatWrongSignMatter(w WrongSignMatterAudit) string {
	return fmt.Sprintf("ordinaryMatterSign=%q reason=%q required=%q conclusion=%q verdict=%q", w.OrdinaryMatterContributionSign, w.Reason, w.RequiredSign, w.Conclusion, w.Verdict)
}

func FormatCorrectionOriginRow(r CorrectionOriginViabilityRow) string {
	return fmt.Sprintf("origin=%q class=%q expectedSign=%q signCompatible=%t size=%q native=%t certified=%t interpretation=%q verdict=%q", r.Origin, r.Class, r.ExpectedSign, r.SignCompatible, r.SizeComment, r.Native, r.Certified, r.Interpretation, r.Verdict)
}
func FormatCorrectionOrigins(rows []CorrectionOriginViabilityRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCorrectionOriginRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatBoundaryThreshold(b BoundaryThresholdSlotAudit) string {
	return fmt.Sprintf("slot=%q requiredDeltaU=%s requiredDeltaAlpha=%s signCompatible=%t betaEquivalent=%s interpretation=%q verdict=%q", b.SlotName, f64(b.RequiredDeltaU), f64(b.RequiredDeltaAlpha), b.SignCompatible, f64(b.UniformBetaEquivalent), b.Interpretation, b.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("strongThreshold=%t colorKinetic=%t extraColored=%t gaugeExtension=%t unification=%t statement=%q verdict=%q", n.HasNativeStrongThresholdTheorem, n.HasNativeColorKineticBoundary, n.HasNativeExtraColoredSpectrum, n.HasNativeGaugeSectorExtension, n.ClaimsUnification, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("newParticles=%t thresholdExistence=%t unification=%t altersAF=%t endpoint=%t verdict=%q", f.IntroducesNewParticles, f.ClaimsThresholdExistence, f.ClaimsGaugeUnification, f.AltersAF, f.DerivesEndpoint, f.Verdict)
}
