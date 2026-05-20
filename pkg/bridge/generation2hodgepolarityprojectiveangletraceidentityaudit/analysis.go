// Package generation2hodgepolarityprojectiveangletraceidentityaudit implements
// Gate 642: HodgePolarity ProjectiveAngle TraceIdentity Audit.
//
// Gate 641 organized the compact/split obstruction as a projective angle with
// sin(theta)=4*sqrt(3)/sqrt(217) and cos(theta)=13/sqrt(217). Gate 642 audits
// whether the full contraction pair (13, 4*sqrt(3)) can be reduced to native
// Frobenius/projector trace expressions involving the K7 Hodge polarity blocks.
// It treats the 13/48/217 skeleton as a structured obstruction candidate, while
// preserving the firewall that no native trace identity, split-G2 carrier,
// boundary assignment, or physical angle is certified.
package generation2hodgepolarityprojectiveangletraceidentityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate641 "github.com/bagherbal/asha-engine/pkg/bridge/generation2twistresidualcomplementanglesourceaudit"
)

const (
	AuditID = "GATE642-HODGE-POLARITY-PROJECTIVE-ANGLE-TRACE-IDENTITY-AUDIT"

	StatusGate641AngleInherited      = "PASS_GATE641_PROJECTIVE_ANGLE_INHERITED"
	StatusRawContractionsComputed    = "PASS_RAW_FROBENIUS_CONTRACTIONS_COMPUTED"
	StatusHodgeSectorBlocksComputed  = "PASS_HODGE_SECTOR_BLOCK_DECOMPOSITION_COMPUTED"
	StatusProjectivePairAudited      = "PASS_PROJECTIVE_PAIR_13_AND_4SQRT3_AUDITED"
	StatusHodgePolarityBlockSkeleton = "CONDITIONAL_SUPPORT_13_AND_48_HAVE_HODGE_POLARITY_BLOCK_SKELETON"
	StatusOffSectorObstructionBlock  = "CONDITIONAL_SUPPORT_48_AS_OFF_SECTOR_OBSTRUCTION_BLOCK_CANDIDATE"
	StatusTraceIdentityCandidates    = "PASS_TRACE_IDENTITY_CANDIDATES_FOR_PROJECTIVE_ANGLE_AUDITED"
	StatusNoNativeTraceIdentity      = "FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_PROJECTIVE_ANGLE_YET"
	StatusNoCertifiedSplitG2         = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress           = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem       = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport    = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalAngle            = "FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_ANGLE"
	StatusNoPhysicalMetric           = "FAILED_ROUTE_PROJECTIVE_ANGLE_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNoHiggsFlavorGauge         = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusTraceIdentityArtifact      = "FAILED_ROUTE_PROJECTIVE_ANGLE_TRACE_IDENTITY_IS_NUMERICAL_ARTIFACT"
	StatusGate642Boundary            = "FIREWALL_PRESERVED_GATE642_PROJECTIVE_ANGLE_IS_INTERNAL_OBSTRUCTION_ONLY"
)

const (
	k7Dim                  = 7
	k7PlusDim              = 4
	k7MinusDim             = 3
	traceSK                = 1
	octonionicDim          = 14
	lambda4SelfDualDim     = 35
	alignmentRoot          = 13
	failureNumerator       = 48
	alignmentNumerator     = 169
	angleDenominator       = 217
	traceIdentityTolerance = 1e-12
)

type Gate641Inheritance struct {
	SinTheta                        float64
	CosTheta                        float64
	TanTheta                        float64
	SinSquared                      float64
	CosSquared                      float64
	FailureNumerator                int
	AlignmentRoot                   int
	AlignmentNumerator              int
	Denominator                     int
	ComplementIdentified            bool
	ProjectiveAngleAudited          bool
	ThirteenSourcesAudited          bool
	TraceIdentityCertifiedByGate641 bool
	SplitG2CertifiedByGate641       bool
	BoundaryStressByGate641         bool
	SevenOver72TheoremByGate641     bool
	ScalarFlavorByGate641           bool
	PhysicalAngleByGate641          bool
	Gate641FirewallPreserved        bool
	Verdict                         string
}

type RawFrobeniusContraction struct {
	RouteName                    string
	NormalizedInnerProduct       float64
	NormalizedInnerProductSquare float64
	NormalizedFailureSquare      float64
	IntegerInnerProductSquare    int
	IntegerFailureSquare         int
	IntegerProductNormSquare     int
	RatioStatement               string
	ProjectivePairMatches        bool
	Comment                      string
}

type RawFrobeniusContractionAudit struct {
	Contractions             []RawFrobeniusContraction
	CandidateCosSquared      float64
	CandidateSinSquared      float64
	MaxCosSquaredDelta       float64
	MaxSinSquaredDelta       float64
	IntegerRatioVerified     bool
	NativeTraceIdentityFound bool
	Verdict                  string
}

type HodgeSectorBlock struct {
	Name              string
	Carrier           string
	Dimension         int
	SignedTrace       int
	Contribution      int
	Formula           string
	Typed             bool
	NativeContraction bool
	Reason            string
}

type HodgeSectorBlockAudit struct {
	PDim                    int
	QDim                    int
	AlignmentAmplitude      int
	FailureAmplitudeSquared int
	FailureAmplitudeText    string
	Denominator             int
	Blocks                  []HodgeSectorBlock
	BlockSkeletonMatches    bool
	NativeTraceIdentity     bool
	Verdict                 string
}

type ProjectivePairAudit struct {
	AlignmentAmplitude             int
	FailureAmplitudeSquared        int
	FailureAmplitudeText           string
	Denominator                    int
	PythagoreanIntegerResidual     int
	TanSquaredNumerator            int
	TanSquaredDenominator          int
	Formula                        string
	PairMatches                    bool
	DerivedFromNativeTraceIdentity bool
	Verdict                        string
}

type TraceIdentityCandidate struct {
	Name                string
	Formula             string
	Value               float64
	Target              float64
	Residual            float64
	MatchesTarget       bool
	NativeTraceIdentity bool
	Reason              string
}

type TraceIdentityAudit struct {
	Candidates               []TraceIdentityCandidate
	BestCandidateName        string
	BestCandidateResidual    float64
	NativeTraceIdentityFound bool
	Verdict                  string
}

type ClassificationAudit struct {
	ProjectiveAngleInherited     bool
	RawContractionsComputed      bool
	HodgeSectorBlocksComputed    bool
	BlockSkeletonSupported       bool
	NativeTraceIdentityCertified bool
	ObstructionOnly              bool
	Interpretation               string
	Verdict                      string
}

type Firewalls struct {
	ClaimsNativeTraceIdentity bool
	ClaimsSplitG2             bool
	ClaimsBoundaryStress      bool
	ClaimsSevenOver72Theorem  bool
	ClaimsScalarFlavor        bool
	ClaimsPhysicalAngle       bool
	ClaimsPhysicalMetric      bool
	ClaimsFlavor              bool
	ClaimsHiggsMass           bool
	ClaimsCKMPMNS             bool
	ClaimsGaugeUnification    bool
	Verdict                   string
}

type Analysis struct {
	Inherited       Gate641Inheritance
	RawContractions RawFrobeniusContractionAudit
	SectorBlocks    HodgeSectorBlockAudit
	ProjectivePair  ProjectivePairAudit
	TraceIdentity   TraceIdentityAudit
	Classification  ClassificationAudit
	Firewalls       Firewalls
	Truth           string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g641, err := gate641.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate641 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g641)
	raw := buildRawContractions(g641)
	blocks := buildSectorBlocks()
	pair := buildProjectivePair()
	trace := buildTraceIdentityAudit()
	classification := buildClassification(inherited, raw, blocks, pair, trace)
	firewalls := Firewalls{Verdict: StatusGate642Boundary}
	truth := "Gate 642 audits the full projective contraction pair behind the Gate641 angle.  The normalized comparison has cos(theta)=13/sqrt(217) and sin(theta)=4*sqrt(3)/sqrt(217), so the integer skeleton is <g_twist,B_K>^2 : ||g_twist||^2||B_K||^2 = 169 : 217 and failure^2 = 48 : 217.  The same pair is organized by the K7 Hodge polarity dimensions p=4 and q=3 through 13=p^2-q and 48=p^2*q.  This gives a strong Hodge-polarity block skeleton for the obstruction angle, but it is not yet a native Frobenius/projector trace identity: no certified formula derives the raw contractions from P_{K7+}, P_{K7-}, S_K, B_K, and the compact Omega tensor.  The angle remains an internal obstruction only, not split-G2, boundary stress, scalar/flavor transport, or physical geometry."
	return Analysis{Inherited: inherited, RawContractions: raw, SectorBlocks: blocks, ProjectivePair: pair, TraceIdentity: trace, Classification: classification, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g641 gate641.Analysis) Gate641Inheritance {
	return Gate641Inheritance{
		SinTheta:                        g641.Complement.SinTheta,
		CosTheta:                        g641.Complement.CosTheta,
		TanTheta:                        g641.Complement.TanTheta,
		SinSquared:                      g641.Complement.RhoSquared,
		CosSquared:                      g641.Complement.Complement,
		FailureNumerator:                g641.Complement.FailureNumerator,
		AlignmentRoot:                   g641.Complement.AlignmentRoot,
		AlignmentNumerator:              g641.Complement.AlignmentNumerator,
		Denominator:                     g641.Complement.Denominator,
		ComplementIdentified:            g641.Complement.ComplementIdentified,
		ProjectiveAngleAudited:          g641.Projective.AllRoutesAlign,
		ThirteenSourcesAudited:          g641.Thirteen.StrongestCandidateTyped && !g641.Thirteen.TraceIdentityCertified,
		TraceIdentityCertifiedByGate641: g641.TraceIdentity.NativeTraceIdentityFound,
		SplitG2CertifiedByGate641:       g641.Firewalls.ClaimsSplitG2,
		BoundaryStressByGate641:         g641.Firewalls.ClaimsBoundaryStress,
		SevenOver72TheoremByGate641:     g641.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorByGate641:           g641.Firewalls.ClaimsScalarFlavor,
		PhysicalAngleByGate641:          g641.Firewalls.ClaimsPhysicalAngle,
		Gate641FirewallPreserved:        g641.Firewalls.Verdict == gate641.StatusGate641Boundary,
		Verdict:                         StatusGate641AngleInherited,
	}
}

func buildRawContractions(g641 gate641.Analysis) RawFrobeniusContractionAudit {
	cos2Target := float64(alignmentNumerator) / float64(angleDenominator)
	sin2Target := float64(failureNumerator) / float64(angleDenominator)
	maxCosDelta := 0.0
	maxSinDelta := 0.0
	all := true
	contractions := []RawFrobeniusContraction{}
	for _, c := range g641.Projective.Contractions {
		cos2 := c.CosSquared
		sin2 := c.SinSquared
		cosDelta := math.Abs(cos2 - cos2Target)
		sinDelta := math.Abs(sin2 - sin2Target)
		if cosDelta > maxCosDelta {
			maxCosDelta = cosDelta
		}
		if sinDelta > maxSinDelta {
			maxSinDelta = sinDelta
		}
		matches := cosDelta < traceIdentityTolerance && sinDelta < traceIdentityTolerance
		if !matches {
			all = false
		}
		contractions = append(contractions, RawFrobeniusContraction{
			RouteName:                    c.RouteName,
			NormalizedInnerProduct:       c.Cosine,
			NormalizedInnerProductSquare: cos2,
			NormalizedFailureSquare:      sin2,
			IntegerInnerProductSquare:    alignmentNumerator,
			IntegerFailureSquare:         failureNumerator,
			IntegerProductNormSquare:     angleDenominator,
			RatioStatement:               "<g_twist,B_K>_F^2 : ||g_twist||_F^2||B_K||_F^2 = 169 : 217; orthogonal failure square = 48 : 217 after projective normalization",
			ProjectivePairMatches:        matches,
			Comment:                      "this records the normalized projective contraction skeleton; the raw unnormalized matrices are still not reduced to a native trace identity",
		})
	}
	verdict := join(StatusRawContractionsComputed, StatusProjectivePairAudited)
	if !all || len(contractions) < 3 {
		verdict = StatusTraceIdentityArtifact
	}
	return RawFrobeniusContractionAudit{Contractions: contractions, CandidateCosSquared: cos2Target, CandidateSinSquared: sin2Target, MaxCosSquaredDelta: maxCosDelta, MaxSinSquaredDelta: maxSinDelta, IntegerRatioVerified: all && len(contractions) >= 3, NativeTraceIdentityFound: false, Verdict: verdict}
}

func buildSectorBlocks() HodgeSectorBlockAudit {
	p := k7PlusDim
	q := k7MinusDim
	alignment := p*p - q
	failure := p * p * q
	den := alignment*alignment + failure
	blocks := []HodgeSectorBlock{
		{Name: "positive-positive Hodge sector", Carrier: "K7+ x K7+", Dimension: p, SignedTrace: p, Contribution: p * p, Formula: "p^2 = dim(K7+)^2 = 4^2 = 16", Typed: true, NativeContraction: false, Reason: "typed by Gate634 Hodge-positive sector, but not yet derived as a Frobenius contraction block"},
		{Name: "negative Hodge subtraction", Carrier: "K7-", Dimension: q, SignedTrace: -q, Contribution: -q, Formula: "-q = -dim(K7-) = -3", Typed: true, NativeContraction: false, Reason: "typed by Gate634 Hodge-negative sector and gives 13=p^2-q, but the subtraction rule is not yet an operator theorem"},
		{Name: "off-sector obstruction block", Carrier: "K7+ squared coupled to K7-", Dimension: q, SignedTrace: 0, Contribution: failure, Formula: "p^2*q = 4^2*3 = 48", Typed: true, NativeContraction: false, Reason: "matches the obstruction numerator and suggests a plus-sector/negative-sector coupling block; no tensor contraction identity is certified"},
		{Name: "projective norm closure", Carrier: "angle denominator", Dimension: den, SignedTrace: 0, Contribution: den, Formula: "217=(p^2-q)^2+p^2*q", Typed: true, NativeContraction: false, Reason: "closes the finite angle pair but remains a rational skeleton rather than a raw trace theorem"},
	}
	matches := alignment == alignmentRoot && failure == failureNumerator && den == angleDenominator
	verdict := join(StatusHodgeSectorBlocksComputed, StatusHodgePolarityBlockSkeleton, StatusOffSectorObstructionBlock)
	if !matches {
		verdict = StatusTraceIdentityArtifact
	}
	return HodgeSectorBlockAudit{PDim: p, QDim: q, AlignmentAmplitude: alignment, FailureAmplitudeSquared: failure, FailureAmplitudeText: "4*sqrt(3)", Denominator: den, Blocks: blocks, BlockSkeletonMatches: matches, NativeTraceIdentity: false, Verdict: verdict}
}

func buildProjectivePair() ProjectivePairAudit {
	alignment := k7PlusDim*k7PlusDim - k7MinusDim
	failure := k7PlusDim * k7PlusDim * k7MinusDim
	den := alignment*alignment + failure
	residual := den - angleDenominator
	matches := alignment == alignmentRoot && failure == failureNumerator && den == angleDenominator && residual == 0
	verdict := StatusHodgePolarityBlockSkeleton
	if !matches {
		verdict = StatusTraceIdentityArtifact
	}
	return ProjectivePairAudit{AlignmentAmplitude: alignment, FailureAmplitudeSquared: failure, FailureAmplitudeText: "4*sqrt(3)", Denominator: den, PythagoreanIntegerResidual: residual, TanSquaredNumerator: failure, TanSquaredDenominator: alignment * alignment, Formula: "tan^2(theta)=p^2*q/(p^2-q)^2 with p=dim(K7+)=4 and q=dim(K7-)=3", PairMatches: matches, DerivedFromNativeTraceIdentity: false, Verdict: verdict}
}

func buildTraceIdentityAudit() TraceIdentityAudit {
	cos2 := float64(alignmentNumerator) / float64(angleDenominator)
	sin2 := float64(failureNumerator) / float64(angleDenominator)
	candidates := []TraceIdentityCandidate{
		{Name: "Hodge-polarity block skeleton", Formula: "(p^2-q)^2 / [(p^2-q)^2+p^2*q]", Value: float64((k7PlusDim*k7PlusDim-k7MinusDim)*(k7PlusDim*k7PlusDim-k7MinusDim)) / float64((k7PlusDim*k7PlusDim-k7MinusDim)*(k7PlusDim*k7PlusDim-k7MinusDim)+k7PlusDim*k7PlusDim*k7MinusDim), Target: cos2, NativeTraceIdentity: false, Reason: "matches cos^2(theta), but it is a dimension-block skeleton, not a contraction of P_{K7+}, P_{K7-}, S_K, B_K, and Omega"},
		{Name: "off-sector obstruction skeleton", Formula: "p^2*q / [(p^2-q)^2+p^2*q]", Value: float64(k7PlusDim*k7PlusDim*k7MinusDim) / float64((k7PlusDim*k7PlusDim-k7MinusDim)*(k7PlusDim*k7PlusDim-k7MinusDim)+k7PlusDim*k7PlusDim*k7MinusDim), Target: sin2, NativeTraceIdentity: false, Reason: "matches sin^2(theta), but the off-sector block norm is not yet produced by a certified tensor identity"},
		{Name: "octonionic-minus-trace alignment shadow", Formula: "(dim(Im(P_G))-tr(S_K))^2 / 217", Value: float64((octonionicDim-traceSK)*(octonionicDim-traceSK)) / float64(angleDenominator), Target: cos2, NativeTraceIdentity: false, Reason: "matches the alignment component but is ambiguous with the Hodge-polarity expression 4^2-3"},
		{Name: "raw projector trace placeholder", Formula: "Tr(F(P_K+,P_K-,S_K,Omega_0)) / Tr(G(P_K+,P_K-,S_K,Omega_0))", Value: math.NaN(), Target: cos2, NativeTraceIdentity: false, Reason: "this is the missing theorem; no explicit native contraction F/G has been certified"},
	}
	bestName := ""
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		if math.IsNaN(c.Value) {
			continue
		}
		res := math.Abs(c.Value - c.Target)
		c.MatchesTarget = res < traceIdentityTolerance
		if res < bestResidual {
			bestResidual = res
			bestName = c.Name
		}
	}
	// Fill the matches flag after the range copy above.
	for i := range candidates {
		if math.IsNaN(candidates[i].Value) {
			continue
		}
		candidates[i].Residual = math.Abs(candidates[i].Value - candidates[i].Target)
		candidates[i].MatchesTarget = candidates[i].Residual < traceIdentityTolerance
	}
	return TraceIdentityAudit{Candidates: candidates, BestCandidateName: bestName, BestCandidateResidual: bestResidual, NativeTraceIdentityFound: false, Verdict: join(StatusTraceIdentityCandidates, StatusNoNativeTraceIdentity)}
}

func buildClassification(inh Gate641Inheritance, raw RawFrobeniusContractionAudit, blocks HodgeSectorBlockAudit, pair ProjectivePairAudit, trace TraceIdentityAudit) ClassificationAudit {
	candidate := inh.ComplementIdentified && inh.ProjectiveAngleAudited && raw.IntegerRatioVerified && blocks.BlockSkeletonMatches && pair.PairMatches && !trace.NativeTraceIdentityFound && inh.Gate641FirewallPreserved
	artifact := !candidate || inh.TraceIdentityCertifiedByGate641
	interp := "the projective compact/split angle has a typed Hodge-polarity block skeleton: with p=dim(K7+)=4 and q=dim(K7-)=3, the alignment amplitude is p^2-q=13 and the failure block is p^2*q=48, giving cos^2=13^2/217 and sin^2=48/217.  This is not yet a native trace identity because no certified contraction of the projectors and Omega tensor produces these raw integers."
	verdict := StatusHodgePolarityBlockSkeleton
	if artifact {
		verdict = StatusTraceIdentityArtifact
	}
	return ClassificationAudit{ProjectiveAngleInherited: inh.ComplementIdentified && inh.ProjectiveAngleAudited, RawContractionsComputed: raw.IntegerRatioVerified, HodgeSectorBlocksComputed: blocks.BlockSkeletonMatches, BlockSkeletonSupported: candidate, NativeTraceIdentityCertified: trace.NativeTraceIdentityFound, ObstructionOnly: candidate, Interpretation: interp, Verdict: verdict}
}

func Statuses() []string {
	return []string{StatusGate641AngleInherited, StatusRawContractionsComputed, StatusHodgeSectorBlocksComputed, StatusProjectivePairAudited, StatusHodgePolarityBlockSkeleton, StatusOffSectorObstructionBlock, StatusTraceIdentityCandidates, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate642Boundary}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
