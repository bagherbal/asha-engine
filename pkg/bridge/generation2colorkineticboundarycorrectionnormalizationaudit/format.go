package generation2colorkineticboundarycorrectionnormalizationaudit

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

func FormatInherited(i InheritedGate609) string {
	return fmt.Sprintf("Lambda12=%s gStar=%s g3=%s uStar=%s u3=%s delta=%s deltaAlpha=%s DeltaB3=%s b3Eff=%s wrongSignMatter=%t verdict=%q", f64(i.Lambda12GeV), f64(i.GStar), f64(i.G3Runtime), f64(i.UStar), f64(i.U3Runtime), f64(i.Delta3Required), f64(i.DeltaAlpha3Inv), f64(i.DeltaB3Required), f64(i.B3EffectiveDiagnostic), i.ExtraColoredMatterWrongSign, i.Verdict)
}

func FormatBoundaryCorrectionRow(r BoundaryKineticCorrectionRow) string {
	return fmt.Sprintf("quantity=%q formula=%q value=%s interpretation=%q verdict=%q", r.Quantity, r.Formula, f64(r.Value), r.Interpretation, r.Verdict)
}
func FormatBoundaryCorrections(rows []BoundaryKineticCorrectionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatBoundaryCorrectionRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatFractionalCorrection(f FractionalCorrectionAudit) string {
	return fmt.Sprintf("etaStar=%s etaRuntime=%s percentStar=%s percentRuntime=%s alphaStarInv=%s alpha3Inv=%s deltaAlpha=%s interpretation=%q verdict=%q", f64(f.EtaAgainstUStar), f64(f.EtaAgainstU3), f64(f.PercentAgainstUStar), f64(f.PercentAgainstU3), f64(f.AlphaStarInv), f64(f.Alpha3InvRuntime), f64(f.DeltaAlphaInv), f.Interpretation, f.Verdict)
}

func FormatGaugeCoefficientAudit(g SpectralActionGaugeCoefficientAudit) string {
	return fmt.Sprintf("lane=%q shift=%q fraction=%s signCompatible=%t native=%t certified=%t interpretation=%q verdict=%q", g.SymbolicLane, g.BoundaryShift, f64(g.RequiredFraction), g.SignCompatible, g.Native, g.Certified, g.Interpretation, g.Verdict)
}

func FormatTraceNormalizationRow(r TraceNormalizationRow) string {
	return fmt.Sprintf("object=%q nativeStatus=%q canColorOnly=%t interpretation=%q verdict=%q", r.Object, r.NativeStatus, r.CanSupplyColorOnlyCorrection, r.Interpretation, r.Verdict)
}
func FormatTraceNormalizations(rows []TraceNormalizationRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatTraceNormalizationRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatFSAStatus(s FiniteSpectralActionStatus) string {
	return fmt.Sprintf("independentColor=%t sectorSplitF0=%t su3Only=%t algebraExtension=%t bSector=%t statement=%q verdict=%q", s.HasIndependentColorKineticCoefficient, s.HasSectorSplitF0Moment, s.HasSU3OnlyBoundaryCorrection, s.HasFiniteAlgebraExtension, s.HasBSectorColorKineticTheorem, s.Statement, s.Verdict)
}

func FormatThresholdLocalized(t ThresholdLocalizedInterpretation) string {
	return fmt.Sprintf("slot=%q deltaU=%s deltaAlpha=%s signCompatible=%t betaEquivalent=%s cleaner=%t interpretation=%q verdict=%q", t.SlotName, f64(t.RequiredDeltaU), f64(t.RequiredDeltaAlphaInv), t.SignCompatible, f64(t.FullIntervalBetaEquivalent), t.CleanerThanFullIntervalMatter, t.Interpretation, t.Verdict)
}

func FormatTwoLoopSchemeCaution(c TwoLoopSchemeCaution) string {
	return fmt.Sprintf("twoLoop=%t scheme=%t alphaS=%t closure=%t statement=%q verdict=%q", c.TwoLoopCouldShiftResidual, c.SchemeCouldShiftResidual, c.AlphaSUncertaintyRelevant, c.ClosureCertified, c.Statement, c.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("colorKinetic=%t threshold=%t unification=%t altersAF=%t addsColored=%t statement=%q verdict=%q", n.ProvesColorKineticBoundaryCorrection, n.ProvesThresholdSpectrum, n.ProvesFullGaugeUnification, n.AltersAF, n.AddsColoredStates, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("correctionExists=%t unification=%t altersAF=%t newColored=%t endpoint=%t verdict=%q", f.ClaimsCorrectionExists, f.ClaimsGaugeUnification, f.AltersFiniteAlgebra, f.AddsNewColoredStates, f.DerivesEndpoint, f.Verdict)
}
