package generation2strongcouplingthresholdresidualledgeraudit

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

func FormatInherited(a InheritedGate606) string {
	return fmt.Sprintf("gaugeSpine=%t Lambda12=%s gStar=%s g3Lambda=%s Delta3=%s R3=%s verdict=%q", a.GaugeSpinePresent, f64(a.Lambda12GeV), f64(a.GStar), f64(a.G3Lambda), f64(a.Delta3Runtime), f64(a.R3), a.Verdict)
}

func FormatResidualRow(r StrongResidualConversionRow) string {
	return fmt.Sprintf("quantity=%q formula=%q value=%s sign=%q interpretation=%q verdict=%q", r.Quantity, r.Formula, f64(r.Value), r.SignConvention, r.Interpretation, r.Verdict)
}
func FormatResidualTable(rows []StrongResidualConversionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatResidualRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatThresholdSlotRow(r ThresholdCorrectionSlotRow) string {
	return fmt.Sprintf("slot=%q definition=%q required=%s unit=%q sign=%q interpretation=%q verdict=%q", r.Slot, r.Definition, f64(r.RequiredValue), r.Unit, r.SignConvention, r.Interpretation, r.Verdict)
}
func FormatThresholdSlots(rows []ThresholdCorrectionSlotRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatThresholdSlotRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatBetaDeformation(b BetaCoefficientDeformationAudit) string {
	return fmt.Sprintf("t=%s sm_b3=%s deltaU=%s deltaB3=%s effectiveB3=%s fractionAbsSM=%s formula=%q interpretation=%q verdict=%q", f64(b.LogInterval), f64(b.SMb3), f64(b.RequiredDeltaUCorrection), f64(b.DeltaB3Required), f64(b.EffectiveB3), f64(b.FractionOfAbsSMb3), b.Formula, b.Interpretation, b.Verdict)
}

func FormatMeetingScaleRow(r MeetingScaleRow) string {
	return fmt.Sprintf("pair=%q t=%s scaleGeV=%s gMeet=%s formula=%q interpretation=%q verdict=%q", r.Pair, f64(r.LogMuOverMZ), f64(r.ScaleGeV), f64(r.CouplingAtMeeting), r.Formula, r.Interpretation, r.Verdict)
}
func FormatMeetingScales(rows []MeetingScaleRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatMeetingScaleRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatSourceRow(r SourceClassificationRow) string {
	return fmt.Sprintf("candidate=%q type=%q affectsG3=%t status=%q native=%q verdict=%q", r.Candidate, r.Type, r.CouldAffectG3, r.CurrentStatus, r.NativeStatus, r.Verdict)
}
func FormatSources(rows []SourceClassificationRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSourceRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("nativeThreshold=%t extraColored=%t boundaryG3=%t fullUnification=%t statement=%q verdict=%q", n.ProvidesNativeStrongThreshold, n.ProvidesExtraColoredContent, n.ProvidesBoundaryG3Correction, n.ClaimsFullUnification, n.Statement, n.Verdict)
}

func FormatScalarRelation(s ScalarRelation) string {
	return fmt.Sprintf("lambdaLambda12=%s mixed=%t statement=%q verdict=%q", f64(s.LambdaLambda12), s.MixedIntoStrongLedger, s.Statement, s.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("unification=%t newPhysics=%t endpoint=%t wzPhoton=%t thresholdReal=%t scalarMixed=%t verdict=%q", f.ClaimsGaugeUnification, f.ClaimsNewPhysics, f.DerivesEndpoint, f.DerivesWZPhoton, f.ThresholdAssertedReal, f.ScalarMixedIntoStrong, f.Verdict)
}
