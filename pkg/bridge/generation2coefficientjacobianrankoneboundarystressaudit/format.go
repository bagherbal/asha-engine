package generation2coefficientjacobianrankoneboundarystressaudit

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

func FormatInherited(h Inherited) string {
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s lambda12=%s xi=%s delta3=%s deltaLambda=%s eta3=%s EStress=%s EOverXi=%s verdict=%q", f64(h.Lambda12GeV), f64(h.R3MinusOne), f64(h.LambdaLambda12), f64(h.XiBoundary), f64(h.Delta3ColorBoundary), f64(h.DeltaLambdaBoundary), f64(h.Eta3), f64(h.BoundaryResidual), f64(h.ResidualOverXi), h.Verdict)
}

func FormatDependency(r CoefficientDependency) string {
	return fmt.Sprintf("source=%q dG=%q dS=%q dependency=%q certified=%t bridge=%t native=%t obstruction=%q verdict=%q", r.Source, r.AffectsGauge, r.AffectsScalar, r.Dependency, r.Certified, r.Bridge, r.Native, r.Obstruction, r.Verdict)
}
func FormatDependencyGraph(rows []CoefficientDependency) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDependency(r))
	}
	return strings.Join(parts, " | ")
}

func FormatShadowMap(s NormalizedShadowMap) string {
	return fmt.Sprintf("preferred=%q alternate=%q raw=%q rawSafe=%t preferredSafe=%t color=%s scalar=%s xi=%s statement=%q verdict=%q", s.PreferredPair, s.AlternatePair, s.RawPair, s.RawPairTypeSafe, s.PreferredTypeSafe, f64(s.ColorShadow), f64(s.ScalarShadow), f64(s.XiBoundary), s.Statement, s.Verdict)
}

func FormatJacobianEntry(j JacobianEntry) string {
	return fmt.Sprintf("source=%q dColor=%q dScalar=%q exact=%t certified=%t comment=%q verdict=%q", j.Source, j.DColor, j.DScalar, j.Exact, j.Certified, j.Comment, j.Verdict)
}
func FormatJacobian(rows []JacobianEntry) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatJacobianEntry(r))
	}
	return strings.Join(parts, " | ")
}

func FormatRankOneCandidate(r RankOneSourceCandidate) string {
	return fmt.Sprintf("source=%q colorPositive=%t scalarNegative=%t bridgeRank1=%t native=%t extraSeal=%t class=%q obstruction=%q verdict=%q", r.Source, r.ProducesColorPositive, r.ProducesScalarNegative, r.RankOneBridgeDefinable, r.Native, r.RequiresExtraSeal, r.Classification, r.Obstruction, r.Verdict)
}
func FormatRankOneCandidates(rows []RankOneSourceCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatRankOneCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatRankClassification(r RankClassification) string {
	return fmt.Sprintf("nativeRank1=%t bridgeRank1=%t rankTwo=%t insufficient=%t best=%q statement=%q verdict=%q", r.NativeRankOneFound, r.BridgeRankOneDefinable, r.RankTwoIndependentSlots, r.GrammarInsufficient, r.BestClassification, r.Statement, r.Verdict)
}

func FormatAntiAlignment(a AntiAlignmentTest) string {
	return fmt.Sprintf("candidate=%q canForce=%t native=%t residual=%s residualOverXi=%s statement=%q verdict=%q", a.Candidate, a.CanForceAntiAlignment, a.Native, f64(a.StressResidual), f64(a.ResidualOverXi), a.Statement, a.Verdict)
}

func FormatCanonicalNormalization(c CanonicalNormalizationAudit) string {
	return fmt.Sprintf("runtime=%q KphiKnown=%t canonicalLedger=%t auditBeforeAfter=%t statement=%q verdict=%q", c.RuntimeLambdaConvention, c.KPhiKnown, c.CanonicalScalarLedgerKnown, c.CanAuditLambdaBeforeAfterK, c.Statement, c.Verdict)
}

func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("sectorF0=%t qStress=%t C3Lambda=%t scalarNorm=%t threshold=%t nativeXi=%t statement=%q verdict=%q", n.SectorSplitF0, n.NativeQStress, n.C3LambdaRelation, n.ScalarNormalization, n.ThresholdMatching, n.NativeXi, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("xiNative=%t lambdaZero=%t higgsMass=%t higgsStability=%t unification=%t threshold=%t nativeCorrection=%t verdict=%q", f.ClaimsXiNative, f.ClaimsLambdaZero, f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsGaugeUnification, f.ClaimsThresholdExistence, f.ClaimsNativeCorrection, f.Verdict)
}
