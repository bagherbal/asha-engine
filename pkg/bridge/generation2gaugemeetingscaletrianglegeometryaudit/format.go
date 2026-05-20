package generation2gaugemeetingscaletrianglegeometryaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(a InheritedGate607) string {
	return fmt.Sprintf("Lambda12=%s Lambda13=%s Lambda23=%s gStar=%s g3Lambda12=%s delta3Threshold=%s deltaAlphaInv=%s deltaB3=%s verdict=%q", f64(a.Lambda12GeV), f64(a.Lambda13GeV), f64(a.Lambda23GeV), f64(a.GStar), f64(a.G3Lambda12), f64(a.Delta3ThresholdRequired), f64(a.DeltaAlpha3InvRequired), f64(a.DeltaB3Required), a.Verdict)
}

func FormatPairwiseRow(r PairwiseMeetingScaleRow) string {
	return fmt.Sprintf("pair=%q scaleGeV=%s log10Scale=%s t=%s gMeet=%s exact=%q verdict=%q", r.Pair, f64(r.ScaleGeV), f64(r.Log10Scale), f64(r.LogMuOverMZ), f64(r.CouplingAtMeeting), r.ExactPair, r.Verdict)
}
func FormatPairwiseTable(rows []PairwiseMeetingScaleRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPairwiseRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatLogGeometry(g LogTriangleGeometry) string {
	keys := make([]string, 0, len(g.DistancesFromMean))
	for k := range g.DistancesFromMean {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	dist := make([]string, 0, len(keys))
	for _, k := range keys {
		dist = append(dist, fmt.Sprintf("%s:%s", k, f64(g.DistancesFromMean[k])))
	}
	return fmt.Sprintf("ratio13/12=%s ratio23/13=%s ratio23/12=%s log13/12=%s log23/13=%s spread=%s geomGeV=%s logGeom=%s distances={%s} statement=%q verdict=%q", f64(g.Ratio13Over12), f64(g.Ratio23Over13), f64(g.Ratio23Over12), f64(g.Log10Ratio13Over12), f64(g.Log10Ratio23Over13), f64(g.SpreadDecades), f64(g.GeometricMeanGeV), f64(g.Log10GeometricMean), strings.Join(dist, ","), g.SkewStatement, g.Verdict)
}

func FormatBoundaryChoiceRow(r BoundaryChoiceResidualRow) string {
	return fmt.Sprintf("choice=%q scaleGeV=%s exact=%q g=[%s,%s,%s] u=[%s,%s,%s] maxDeltaU=%s residual=%q verdict=%q", r.ChoiceScale, f64(r.ScaleGeV), r.ExactPair, f64(r.G1), f64(r.G2), f64(r.G3), f64(r.U1), f64(r.U2), f64(r.U3), f64(r.MaxDeltaU), r.ResidualSummary, r.Verdict)
}
func FormatBoundaryChoices(rows []BoundaryChoiceResidualRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatBoundaryChoiceRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatBetaRow(r BetaDeformationVectorRow) string {
	return fmt.Sprintf("strategy=%q LambdaU=%s targetU=%s DeltaB=[%s,%s,%s] norm=%s formula=%q interpretation=%q verdict=%q", r.Strategy, f64(r.LambdaUGeV), f64(r.TargetU), f64(r.DeltaB1), f64(r.DeltaB2), f64(r.DeltaB3), f64(r.Norm), r.Formula, r.Interpretation, r.Verdict)
}
func FormatBetaDeformations(rows []BetaDeformationVectorRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatBetaRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatThresholdSlotRow(r ThresholdOriginSlotRow) string {
	return fmt.Sprintf("candidate=%q kind=%q wouldMove=%q status=%q native=%q verdict=%q", r.Candidate, r.Kind, r.WouldMove, r.CurrentStatus, r.NativeStatus, r.Verdict)
}
func FormatThresholdOriginSlots(rows []ThresholdOriginSlotRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatThresholdSlotRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatNativeStatus(n ASHANativeStatus) string {
	return fmt.Sprintf("thresholdSpectrum=%t bColored=%t algebraExtension=%t colorCorrection=%t lambdaUSelection=%t unification=%t statement=%q verdict=%q", n.ProvidesNativeThresholdSpectrum, n.ProvidesBColoredDeformation, n.ProvidesFiniteAlgebraExtension, n.ProvidesBoundaryColorCorrection, n.ProvidesNativeLambdaUSelection, n.ClaimsUnification, n.Statement, n.Verdict)
}

func FormatScalarRelation(s ScalarTriangleRelation) string {
	return fmt.Sprintf("zeroCrossing=%s Lambda12=%s Lambda13=%s Lambda23=%s statement=%q verdict=%q", f64(s.ZeroCrossingGeV), f64(s.Lambda12GeV), f64(s.Lambda13GeV), f64(s.Lambda23GeV), s.Statement, s.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("unification=%t newFields=%t thresholdsFit=%t LambdaUPromoted=%t endpoint=%t scalarClose=%t verdict=%q", f.ClaimsUnification, f.IntroducesNewFields, f.FitsThresholds, f.PromotesLambdaU, f.DerivesEndpoint, f.UsesScalarToClose, f.Verdict)
}
