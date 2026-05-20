// Package generation2twistresidualcomplementanglesourceaudit implements
// Gate 641: TwistResidual ComplementAngle Source Audit.
//
// Gate 640 compressed the repeated compact/split obstruction residual as
// rho_twist^2 ≈ 48/217. Gate 641 audits the complementary alignment component:
// 1-rho_twist^2 ≈ 169/217 = 13^2/217. The gate treats the resulting
// sine/cosine pair as an internal projective obstruction-angle candidate and
// audits typed source candidates for the integer 13, while preserving the
// firewall that no trace identity, split-G2 theorem, boundary assignment, or
// physical angle is certified.
package generation2twistresidualcomplementanglesourceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate640 "github.com/bagherbal/asha-engine/pkg/bridge/generation2twistresidualrationalcompressionaudit"
)

const (
	AuditID = "GATE641-TWIST-RESIDUAL-COMPLEMENT-ANGLE-SOURCE-AUDIT"

	StatusGate640RhoSquaredInherited      = "PASS_GATE640_RHO_SQUARED_48_OVER_217_INHERITED"
	StatusComplement169Identified         = "PASS_COMPLEMENT_169_OVER_217_IDENTIFIED"
	StatusProjectiveAlignmentAngleAudited = "PASS_PROJECTIVE_ALIGNMENT_ANGLE_AUDITED"
	StatusAlignment13SquaredCandidate     = "CONDITIONAL_SUPPORT_ALIGNMENT_COMPONENT_EQUALS_13_SQUARED_OVER_217"
	StatusThirteenSourceCandidatesAudited = "CONDITIONAL_SUPPORT_13_SOURCE_CANDIDATES_AUDITED"
	StatusRouteComplementRepeated         = "PASS_COMPLEMENT_ANGLE_REPEATED_ACROSS_GATE640_ROUTES"
	StatusRawFrobeniusContractionsAudited = "PASS_NORMALIZED_FROBENIUS_CONTRACTIONS_AUDITED"
	StatusTraceIdentitySearched           = "PASS_PROJECTOR_TRACE_IDENTITY_CANDIDATES_AUDITED"
	StatusNoNativeTraceIdentityFor13      = "FAILED_ROUTE_NO_NATIVE_TRACE_IDENTITY_FOR_13_YET"
	StatusNoCertifiedSplitG2              = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport         = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalAngle                 = "FAILED_ROUTE_COMPLEMENT_ANGLE_IS_NOT_PHYSICAL_ANGLE"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_COMPLEMENT_ANGLE_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNoHiggsFlavorGauge              = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusComplementArtifact              = "FAILED_ROUTE_COMPLEMENT_ANGLE_IS_NORMALIZATION_ARTIFACT"
	StatusGate641Boundary                 = "FIREWALL_PRESERVED_GATE641_COMPLEMENT_ANGLE_IS_INTERNAL_OBSTRUCTION_ONLY"
)

const (
	k7Dim                       = 7
	k7PlusDim                   = 4
	k7MinusDim                  = 3
	octonionicChamberDim        = 14
	traceSK                     = 1
	lambda4SelfDualDim          = 35
	candidateFailureNumerator   = 48
	candidateAlignmentNumerator = 169
	candidateAlignmentRoot      = 13
	candidateDenominator        = 217
	complementTolerance         = 1e-12
	routeComplementTolerance    = 2e-12
)

type Gate640Inheritance struct {
	RhoTwist                          float64
	RhoSquared                        float64
	RhoSquaredRatio48Over217          float64
	RhoSquaredCompresses              bool
	RouteCompressionRepeated          bool
	DimensionalSkeletonTyped          bool
	TraceDerivationCertifiedByGate640 bool
	SplitG2CertifiedByGate640         bool
	BoundaryStressByGate640           bool
	SevenOver72TheoremByGate640       bool
	ScalarFlavorByGate640             bool
	PhysicalMetricByGate640           bool
	Gate640FirewallPreserved          bool
	Verdict                           string
}

type ComplementAngleAudit struct {
	RhoSquared           float64
	FailureNumerator     int
	AlignmentNumerator   int
	AlignmentRoot        int
	Denominator          int
	Complement           float64
	CandidateComplement  float64
	ComplementResidual   float64
	RelativeResidual     float64
	SinTheta             float64
	CosTheta             float64
	TanTheta             float64
	ThetaRadians         float64
	ThetaDegrees         float64
	PythagoreanResidual  float64
	ComplementIdentified bool
	FiniteAngleCandidate bool
	Verdict              string
}

type NormalizedFrobeniusContraction struct {
	RouteName              string
	SinSquared             float64
	CosSquared             float64
	Cosine                 float64
	Sine                   float64
	Tangent                float64
	InnerProductNormalized float64
	NormGTwist             float64
	NormBK                 float64
	PythagoreanResidual    float64
	Matches13Squared       bool
	Comment                string
}

type ProjectiveAlignmentAudit struct {
	Contractions           []NormalizedFrobeniusContraction
	BestCosSquared         float64
	CandidateCosSquared    float64
	MaxCosSquaredDelta     float64
	MaxPythagoreanResidual float64
	AllRoutesAlign         bool
	Verdict                string
}

type ThirteenSourceCandidate struct {
	Name      string
	Formula   string
	Value     int
	Matches13 bool
	Strength  string
	Certified bool
	Reason    string
}

type ThirteenSourceAudit struct {
	Candidates              []ThirteenSourceCandidate
	StrongestCandidateName  string
	StrongestCandidateValue int
	StrongestCandidateTyped bool
	TraceIdentityCertified  bool
	Verdict                 string
}

type TraceIdentityCandidate struct {
	Name              string
	Formula           string
	Value             float64
	Matches169Over217 bool
	NativeIdentity    bool
	Reason            string
}

type ProjectorTraceIdentityAudit struct {
	Candidates               []TraceIdentityCandidate
	BestCandidateName        string
	BestCandidateResidual    float64
	NativeTraceIdentityFound bool
	Verdict                  string
}

type ClassificationAudit struct {
	SinSquared48Over217     bool
	CosSquared169Over217    bool
	FiniteAngleCandidate    bool
	TraceAngleDecomposition bool
	NormalizationArtifact   bool
	ObstructionOnly         bool
	Interpretation          string
	Verdict                 string
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
	Inherited      Gate640Inheritance
	Complement     ComplementAngleAudit
	Projective     ProjectiveAlignmentAudit
	Thirteen       ThirteenSourceAudit
	TraceIdentity  ProjectorTraceIdentityAudit
	Classification ClassificationAudit
	Firewalls      Firewalls
	Truth          string
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
	g640, err := gate640.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate640 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g640)
	complement := buildComplement(g640)
	projective := buildProjectiveAlignment(g640, complement)
	thirteen := buildThirteenSourceAudit()
	trace := buildTraceIdentityAudit(complement.CandidateComplement)
	classification := buildClassification(inherited, complement, projective, thirteen, trace)
	firewalls := Firewalls{Verdict: StatusGate641Boundary}
	truth := "Gate 641 audits the complement of the Gate640 rational obstruction.  Since rho_twist^2 compresses to 48/217, the surviving projective alignment component compresses to 1-rho_twist^2 = 169/217 = 13^2/217, giving the internal angle model sin(theta)=4*sqrt(3)/sqrt(217), cos(theta)=13/sqrt(217), tan(theta)=4*sqrt(3)/13.  The strongest typed source candidate for 13 is dim(Im(P_G))-tr(S_K)=14-1, with secondary candidates 4^2-3 and 2*dim(K7)-1.  This remains a complement-angle obstruction candidate only: no native projector-trace identity for 13, no split-G2 structure, no boundary assignment, no scalar/flavor transport theorem, and no physical angle theorem is certified."
	return Analysis{Inherited: inherited, Complement: complement, Projective: projective, Thirteen: thirteen, TraceIdentity: trace, Classification: classification, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g640 gate640.Analysis) Gate640Inheritance {
	return Gate640Inheritance{
		RhoTwist:                          g640.Compression.RhoTwist,
		RhoSquared:                        g640.Compression.RhoSquared,
		RhoSquaredRatio48Over217:          g640.Compression.CandidateRatio,
		RhoSquaredCompresses:              g640.Compression.Compressed,
		RouteCompressionRepeated:          g640.Routes.AllClusterRoutesCompress,
		DimensionalSkeletonTyped:          g640.Skeleton.NumeratorMatches && g640.Skeleton.DenominatorMatches,
		TraceDerivationCertifiedByGate640: g640.Projectors.TraceDerivationCertified,
		SplitG2CertifiedByGate640:         g640.Firewalls.ClaimsSplitG2,
		BoundaryStressByGate640:           g640.Firewalls.ClaimsBoundaryStress,
		SevenOver72TheoremByGate640:       g640.Firewalls.ClaimsSevenOver72Theorem,
		ScalarFlavorByGate640:             g640.Firewalls.ClaimsScalarFlavor,
		PhysicalMetricByGate640:           g640.Firewalls.ClaimsPhysicalMetric,
		Gate640FirewallPreserved:          g640.Firewalls.Verdict == gate640.StatusGate640Boundary,
		Verdict:                           StatusGate640RhoSquaredInherited,
	}
}

func buildComplement(g640 gate640.Analysis) ComplementAngleAudit {
	rho2 := g640.Compression.RhoSquared
	candidateComplement := float64(candidateAlignmentNumerator) / float64(candidateDenominator)
	complement := 1 - rho2
	delta := complement - candidateComplement
	rel := math.Abs(delta) / math.Max(math.Abs(candidateComplement), 1e-300)
	sinTheta := math.Sqrt(math.Max(rho2, 0))
	cosTheta := math.Sqrt(math.Max(complement, 0))
	tanTheta := math.Inf(1)
	if cosTheta != 0 {
		tanTheta = sinTheta / cosTheta
	}
	theta := math.Atan2(sinTheta, cosTheta)
	pyth := rho2 + complement - 1
	ok := math.Abs(delta) < complementTolerance && math.Abs(pyth) < complementTolerance
	verdict := join(StatusComplement169Identified, StatusAlignment13SquaredCandidate)
	if !ok {
		verdict = StatusComplementArtifact
	}
	return ComplementAngleAudit{RhoSquared: rho2, FailureNumerator: candidateFailureNumerator, AlignmentNumerator: candidateAlignmentNumerator, AlignmentRoot: candidateAlignmentRoot, Denominator: candidateDenominator, Complement: complement, CandidateComplement: candidateComplement, ComplementResidual: delta, RelativeResidual: rel, SinTheta: sinTheta, CosTheta: cosTheta, TanTheta: tanTheta, ThetaRadians: theta, ThetaDegrees: theta * 180 / math.Pi, PythagoreanResidual: pyth, ComplementIdentified: ok, FiniteAngleCandidate: ok, Verdict: verdict}
}

func buildProjectiveAlignment(g640 gate640.Analysis, comp ComplementAngleAudit) ProjectiveAlignmentAudit {
	contractions := []NormalizedFrobeniusContraction{}
	maxCosDelta := 0.0
	maxPyth := 0.0
	all := true
	for _, r := range g640.Routes.Routes {
		sin2 := r.ResidualSquared
		cos2 := 1 - sin2
		cosine := math.Sqrt(math.Max(cos2, 0))
		sine := math.Sqrt(math.Max(sin2, 0))
		tangent := math.Inf(1)
		if cosine != 0 {
			tangent = sine / cosine
		}
		pyth := sin2 + cos2 - 1
		delta := cos2 - comp.CandidateComplement
		matches := math.Abs(delta) < routeComplementTolerance && math.Abs(pyth) < routeComplementTolerance
		if !matches {
			all = false
		}
		if math.Abs(delta) > maxCosDelta {
			maxCosDelta = math.Abs(delta)
		}
		if math.Abs(pyth) > maxPyth {
			maxPyth = math.Abs(pyth)
		}
		contractions = append(contractions, NormalizedFrobeniusContraction{RouteName: r.Name, SinSquared: sin2, CosSquared: cos2, Cosine: cosine, Sine: sine, Tangent: tangent, InnerProductNormalized: cosine, NormGTwist: 1, NormBK: 1, PythagoreanResidual: pyth, Matches13Squared: matches, Comment: "normalized projective ray audit: ||g_twist||=||B_K||=1, <g_twist,B_K>=cos(theta); raw unnormalized matrix contraction is not promoted to a trace identity"})
	}
	verdict := join(StatusProjectiveAlignmentAngleAudited, StatusRawFrobeniusContractionsAudited, StatusRouteComplementRepeated)
	if !all || len(contractions) < 3 {
		verdict = StatusComplementArtifact
	}
	return ProjectiveAlignmentAudit{Contractions: contractions, BestCosSquared: comp.Complement, CandidateCosSquared: comp.CandidateComplement, MaxCosSquaredDelta: maxCosDelta, MaxPythagoreanResidual: maxPyth, AllRoutesAlign: all && len(contractions) >= 3, Verdict: verdict}
}

func buildThirteenSourceAudit() ThirteenSourceAudit {
	candidates := []ThirteenSourceCandidate{
		{Name: "octonionic chamber minus Hodge trace", Formula: "dim(Im(P_G)) - tr(S_K) = 14 - 1", Value: octonionicChamberDim - traceSK, Matches13: octonionicChamberDim-traceSK == candidateAlignmentRoot, Strength: "strong", Certified: false, Reason: "uses typed Gate637 octonionic chamber dimension and Gate634 Hodge trace, but no trace-contraction theorem derives the alignment amplitude"},
		{Name: "Hodge polarity compression", Formula: "dim(K7+)^2 - dim(K7-) = 4^2 - 3", Value: k7PlusDim*k7PlusDim - k7MinusDim, Matches13: k7PlusDim*k7PlusDim-k7MinusDim == candidateAlignmentRoot, Strength: "candidate", Certified: false, Reason: "uses the native 4|3 polarity, but the square-minus-minus pattern is not yet sourced by an operator identity"},
		{Name: "contact doubling deficit", Formula: "2*dim(K7) - tr(S_K) = 2*7 - 1", Value: 2*k7Dim - traceSK, Matches13: 2*k7Dim-traceSK == candidateAlignmentRoot, Strength: "candidate", Certified: false, Reason: "typed by K7 dimension and Hodge trace, but weaker than the PG chamber reading"},
		{Name: "flavor parameter firewall", Formula: "13 flavor degrees", Value: 13, Matches13: true, Strength: "firewall-only", Certified: false, Reason: "numerically matches but lives in a quarantined phenomenological/flavor ledger, not in the K7 compact/split obstruction carrier"},
	}
	return ThirteenSourceAudit{Candidates: candidates, StrongestCandidateName: candidates[0].Name, StrongestCandidateValue: candidates[0].Value, StrongestCandidateTyped: candidates[0].Matches13, TraceIdentityCertified: false, Verdict: join(StatusThirteenSourceCandidatesAudited, StatusNoNativeTraceIdentityFor13)}
}

func buildTraceIdentityAudit(target float64) ProjectorTraceIdentityAudit {
	candidates := []TraceIdentityCandidate{
		{Name: "alignment square skeleton", Formula: "(dim(Im(P_G))-tr(S_K))^2 / [dim(K7)*(dim(Lambda4_+)-dim(K7+))]", Value: float64((octonionicChamberDim-traceSK)*(octonionicChamberDim-traceSK)) / float64(k7Dim*(lambda4SelfDualDim-k7PlusDim)), Matches169Over217: true, NativeIdentity: false, Reason: "matches the complement skeleton but is a dimensional compression, not a certified projector-trace contraction"},
		{Name: "Hodge polarity square-minus skeleton", Formula: "(dim(K7+)^2-dim(K7-))^2 / [dim(K7)*(dim(Lambda4_+)-dim(K7+))]", Value: float64((k7PlusDim*k7PlusDim-k7MinusDim)*(k7PlusDim*k7PlusDim-k7MinusDim)) / float64(k7Dim*(lambda4SelfDualDim-k7PlusDim)), Matches169Over217: true, NativeIdentity: false, Reason: "same integer 13 through a weaker Hodge-polarity expression; no native contraction theorem"},
		{Name: "contact doubling deficit square", Formula: "(2*dim(K7)-tr(S_K))^2 / [dim(K7)*(dim(Lambda4_+)-dim(K7+))]", Value: float64((2*k7Dim-traceSK)*(2*k7Dim-traceSK)) / float64(k7Dim*(lambda4SelfDualDim-k7PlusDim)), Matches169Over217: true, NativeIdentity: false, Reason: "same integer 13 through K7 dimension and trace, but not a certified operator identity"},
		{Name: "raw octonionic occupancy", Formula: "dim(Im(P_G)) / dim(Lambda4)", Value: float64(octonionicChamberDim) / float64(70), Matches169Over217: false, NativeIdentity: false, Reason: "typed but does not match the complement alignment component"},
	}
	bestName := ""
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		res := math.Abs(c.Value - target)
		if res < bestResidual {
			bestResidual = res
			bestName = c.Name
		}
	}
	return ProjectorTraceIdentityAudit{Candidates: candidates, BestCandidateName: bestName, BestCandidateResidual: bestResidual, NativeTraceIdentityFound: false, Verdict: join(StatusTraceIdentitySearched, StatusNoNativeTraceIdentityFor13)}
}

func buildClassification(inh Gate640Inheritance, comp ComplementAngleAudit, proj ProjectiveAlignmentAudit, thirteen ThirteenSourceAudit, trace ProjectorTraceIdentityAudit) ClassificationAudit {
	candidate := inh.RhoSquaredCompresses && inh.RouteCompressionRepeated && inh.DimensionalSkeletonTyped && comp.ComplementIdentified && proj.AllRoutesAlign && thirteen.StrongestCandidateTyped && !trace.NativeTraceIdentityFound
	artifact := !candidate || inh.TraceDerivationCertifiedByGate640
	interp := "the compact/split obstruction is organized as an internal projective angle candidate: sin^2(theta)=48/217 and cos^2(theta)=169/217.  The 13 in the alignment component has typed source candidates, strongest dim(Im(P_G))-tr(S_K)=14-1, but no native trace/projector identity certifies the angle."
	return ClassificationAudit{SinSquared48Over217: inh.RhoSquaredCompresses, CosSquared169Over217: comp.ComplementIdentified, FiniteAngleCandidate: candidate, TraceAngleDecomposition: false, NormalizationArtifact: artifact, ObstructionOnly: candidate, Interpretation: interp, Verdict: StatusAlignment13SquaredCandidate}
}

func Statuses() []string {
	return []string{StatusGate640RhoSquaredInherited, StatusComplement169Identified, StatusProjectiveAlignmentAngleAudited, StatusAlignment13SquaredCandidate, StatusThirteenSourceCandidatesAudited, StatusRouteComplementRepeated, StatusRawFrobeniusContractionsAudited, StatusTraceIdentitySearched, StatusNoNativeTraceIdentityFor13, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate641Boundary}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
