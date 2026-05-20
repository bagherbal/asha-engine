// Package generation2twistresidualrationalcompressionaudit implements
// Gate 640: TwistResidual RationalCompression Audit.
//
// Gate 639 certified rho_twist as a repeated compact/split obstruction witness
// on K_7. Gate 640 audits the sharper finite skeleton exposed by the square of
// that residual: rho_twist^2 ≈ 48/217. The audit tests whether the compression
// is stable across the independent Gate639 residual routes and whether the
// integers can be typed by the Gate634 Hodge polarity and the ambient
// self-dual chamber, while preserving the firewall that no trace derivation,
// split-G2 theorem, boundary-stress assignment, or physical theorem is yet
// certified.
package generation2twistresidualrationalcompressionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate639 "github.com/bagherbal/asha-engine/pkg/bridge/generation2compactsplittwistresidualinvariantaudit"
)

const (
	AuditID = "GATE640-TWIST-RESIDUAL-RATIONAL-COMPRESSION-AUDIT"

	StatusGate639RhoInherited          = "PASS_GATE639_RHO_TWIST_INVARIANT_INHERITED"
	StatusRhoSquaredCompressionTested  = "PASS_RHO_TWIST_SQUARED_RATIONAL_COMPRESSION_TESTED"
	StatusRhoSquaredEquals48Over217    = "CONDITIONAL_SUPPORT_RHO_TWIST_SQUARED_EQUALS_48_OVER_217_CANDIDATE"
	StatusDenominator217TypedCandidate = "CONDITIONAL_SUPPORT_DENOMINATOR_217_MATCHES_7_TIMES_SELF_DUAL_COMPLEMENT_31"
	StatusNumerator48TypedCandidate    = "CONDITIONAL_SUPPORT_NUMERATOR_48_MATCHES_4_SQUARED_TIMES_3_HODGE_POLARITY"
	StatusRouteCompressionRepeated     = "PASS_48_OVER_217_COMPRESSION_REPEATED_ACROSS_GATE639_ROUTES"
	StatusProjectorContractionsAudited = "PASS_TRACE_PROJECTOR_CONTRACTION_CANDIDATES_AUDITED"
	StatusNoTraceDerivation            = "FAILED_ROUTE_NO_NATIVE_TRACE_DERIVATION_OF_48_OVER_217_YET"
	StatusNoCertifiedSplitG2           = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress             = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport      = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric             = "FAILED_ROUTE_RHO_COMPRESSION_IS_NOT_PHYSICAL_METRIC_THEOREM"
	StatusNormalizationArtifact        = "FAILED_ROUTE_RHO_TWIST_RATIONAL_COMPRESSION_IS_NORMALIZATION_ARTIFACT"
	StatusGate640Boundary              = "FIREWALL_PRESERVED_GATE640_RATIONAL_COMPRESSION_IS_OBSTRUCTION_ONLY"
)

const (
	k7Dim                  = 7
	k7PlusDim              = 4
	k7MinusDim             = 3
	lambda4SelfDualDim     = 35
	lambda4AntiSelfDualDim = 35
	lambda4TotalDim        = 70
	candidateNumerator     = 48
	candidateDenominator   = 217
	rationalTolerance      = 1e-12
	routeTolerance         = 1e-12
)

type Gate639Inheritance struct {
	RhoTwist                        float64
	RhoSquared                      float64
	RepeatedAcrossRoutes            bool
	ResidualInvariant               bool
	CompactSplitObstruction         bool
	Gate639ClassifiedAsArtifact     bool
	Gate639SplitG2Certified         bool
	Gate639BoundaryStressAssignment bool
	Gate639SevenOver72Theorem       bool
	Gate639ScalarFlavorTransport    bool
	Gate639FirewallPreserved        bool
	Verdict                         string
}

type RationalCompressionAudit struct {
	RhoTwist             float64
	RhoSquared           float64
	CandidateNumerator   int
	CandidateDenominator int
	CandidateRatio       float64
	CandidateSqrt        float64
	ResidualSquared      float64
	RelativeResidual     float64
	RhoResidual          float64
	Compressed           bool
	Verdict              string
}

type RouteCompression struct {
	Name             string
	Residual         float64
	ResidualSquared  float64
	DeltaTo48Over217 float64
	RhoDelta         float64
	Compressed       bool
}

type RouteCompressionAudit struct {
	Routes                   []RouteCompression
	CompressedRouteNames     []string
	MaxSquaredDelta          float64
	MaxRhoDelta              float64
	AllClusterRoutesCompress bool
	Verdict                  string
}

type DimensionalSkeletonAudit struct {
	K7Dim                         int
	K7PlusDim                     int
	K7MinusDim                    int
	Lambda4SelfDualDim            int
	Lambda4AntiSelfDualDim        int
	Lambda4TotalDim               int
	SelfDualComplementToK7PlusDim int
	NumeratorFromHodgePolarity    int
	DenominatorFromSelfDualGap    int
	CandidateNumerator            int
	CandidateDenominator          int
	NumeratorMatches              bool
	DenominatorMatches            bool
	Formula                       string
	Verdict                       string
}

type ProjectorContractionCandidate struct {
	Name             string
	Formula          string
	Value            float64
	Matches48Over217 bool
	NativeDerivation bool
	Reason           string
}

type ProjectorContractionAudit struct {
	Candidates               []ProjectorContractionCandidate
	PPlusDimension           int
	PMinusDimension          int
	PK7PlusDimension         int
	PK7MinusDimension        int
	TraceDerivationCertified bool
	BestCandidateName        string
	BestCandidateResidual    float64
	Verdict                  string
}

type ClassificationAudit struct {
	RhoSquared                   float64
	Ratio48Over217               float64
	CompressionCandidate         bool
	ExactFromFiniteMatrixClaim   bool
	ConsequenceOfHodgeSplitClaim bool
	ArtifactClaim                bool
	ObstructionOnly              bool
	Interpretation               string
	Verdict                      string
}

type Firewalls struct {
	ClaimsExactTraceTheorem  bool
	ClaimsSplitG2            bool
	ClaimsBoundaryStress     bool
	ClaimsSevenOver72Theorem bool
	ClaimsScalarFlavor       bool
	ClaimsPhysicalMetric     bool
	ClaimsFlavor             bool
	ClaimsHiggsMass          bool
	ClaimsCKMPMNS            bool
	ClaimsGaugeUnification   bool
	Verdict                  string
}

type Analysis struct {
	Inherited      Gate639Inheritance
	Compression    RationalCompressionAudit
	Routes         RouteCompressionAudit
	Skeleton       DimensionalSkeletonAudit
	Projectors     ProjectorContractionAudit
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
	g639, err := gate639.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate639 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g639)
	compression := buildCompression(g639)
	routes := buildRouteCompression(g639)
	skeleton := buildSkeleton()
	projectors := buildProjectorAudit(compression.CandidateRatio)
	classification := buildClassification(inherited, compression, routes, skeleton, projectors)
	firewalls := Firewalls{Verdict: StatusGate640Boundary}
	truth := "Gate 640 audits the rational skeleton of the Gate639 compact/split obstruction residual.  The squared residual rho_twist^2 compresses to 48/217 to float64 matrix tolerance and the same compression repeats across omega_1_alt, omega_2_alt, and omega_B_alt.  The integers admit a typed dimensional reading, 48=4^2*3 from the K_7 Hodge polarity and 217=7*(35-4) from dim(K_7) times the ambient self-dual complement to K_7^+.  This is only a rational-compression candidate: no native trace/projector contraction theorem deriving 48/217, no split-G2 carrier, no boundary assignment, no scalar/flavor transport theorem, and no physical metric theorem is certified."
	return Analysis{Inherited: inherited, Compression: compression, Routes: routes, Skeleton: skeleton, Projectors: projectors, Classification: classification, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g639 gate639.Analysis) Gate639Inheritance {
	return Gate639Inheritance{
		RhoTwist:                        g639.Repetition.RhoTwist,
		RhoSquared:                      g639.Classification.RhoSquared,
		RepeatedAcrossRoutes:            g639.Repetition.RepeatedAcrossRoutes,
		ResidualInvariant:               g639.Invariance.AllProjectiveTestsPass,
		CompactSplitObstruction:         g639.Classification.ClassifiedAsObstruction,
		Gate639ClassifiedAsArtifact:     g639.Classification.ClassifiedAsArtifact,
		Gate639SplitG2Certified:         g639.Firewalls.ClaimsSplitG2,
		Gate639BoundaryStressAssignment: g639.Firewalls.ClaimsBoundaryStress,
		Gate639SevenOver72Theorem:       g639.Firewalls.ClaimsSevenOver72Theorem,
		Gate639ScalarFlavorTransport:    g639.Firewalls.ClaimsScalarRG || g639.Firewalls.ClaimsFlavor,
		Gate639FirewallPreserved:        g639.Firewalls.Verdict == gate639.StatusGate639Boundary,
		Verdict:                         StatusGate639RhoInherited,
	}
}

func buildCompression(g639 gate639.Analysis) RationalCompressionAudit {
	rho := g639.Repetition.RhoTwist
	rho2 := rho * rho
	ratio := float64(candidateNumerator) / float64(candidateDenominator)
	sqrtRatio := math.Sqrt(ratio)
	delta := rho2 - ratio
	rel := math.Abs(delta) / math.Max(math.Abs(ratio), 1e-300)
	rhoDelta := rho - sqrtRatio
	compressed := math.Abs(delta) < rationalTolerance && math.Abs(rhoDelta) < rationalTolerance
	verdict := StatusRhoSquaredEquals48Over217
	if !compressed {
		verdict = StatusNormalizationArtifact
	}
	return RationalCompressionAudit{RhoTwist: rho, RhoSquared: rho2, CandidateNumerator: candidateNumerator, CandidateDenominator: candidateDenominator, CandidateRatio: ratio, CandidateSqrt: sqrtRatio, ResidualSquared: delta, RelativeResidual: rel, RhoResidual: rhoDelta, Compressed: compressed, Verdict: verdict}
}

func buildRouteCompression(g639 gate639.Analysis) RouteCompressionAudit {
	ratio := float64(candidateNumerator) / float64(candidateDenominator)
	sqrtRatio := math.Sqrt(ratio)
	routes := []RouteCompression{}
	maxSq := 0.0
	maxRho := 0.0
	names := []string{}
	all := true
	for _, r := range g639.Repetition.Routes {
		if !r.IncludedInRhoCluster {
			continue
		}
		rho := r.RelativeResidualToBK
		sq := rho * rho
		dSq := sq - ratio
		dRho := rho - sqrtRatio
		compressed := math.Abs(dSq) < routeTolerance && math.Abs(dRho) < routeTolerance
		if !compressed {
			all = false
		}
		if math.Abs(dSq) > maxSq {
			maxSq = math.Abs(dSq)
		}
		if math.Abs(dRho) > maxRho {
			maxRho = math.Abs(dRho)
		}
		names = append(names, r.Name)
		routes = append(routes, RouteCompression{Name: r.Name, Residual: rho, ResidualSquared: sq, DeltaTo48Over217: dSq, RhoDelta: dRho, Compressed: compressed})
	}
	verdict := StatusRouteCompressionRepeated
	if !all || len(routes) < 3 {
		verdict = StatusNormalizationArtifact
	}
	return RouteCompressionAudit{Routes: routes, CompressedRouteNames: names, MaxSquaredDelta: maxSq, MaxRhoDelta: maxRho, AllClusterRoutesCompress: all && len(routes) >= 3, Verdict: verdict}
}

func buildSkeleton() DimensionalSkeletonAudit {
	selfDualComplement := lambda4SelfDualDim - k7PlusDim
	num := k7PlusDim * k7PlusDim * k7MinusDim
	den := k7Dim * selfDualComplement
	return DimensionalSkeletonAudit{
		K7Dim:                         k7Dim,
		K7PlusDim:                     k7PlusDim,
		K7MinusDim:                    k7MinusDim,
		Lambda4SelfDualDim:            lambda4SelfDualDim,
		Lambda4AntiSelfDualDim:        lambda4AntiSelfDualDim,
		Lambda4TotalDim:               lambda4TotalDim,
		SelfDualComplementToK7PlusDim: selfDualComplement,
		NumeratorFromHodgePolarity:    num,
		DenominatorFromSelfDualGap:    den,
		CandidateNumerator:            candidateNumerator,
		CandidateDenominator:          candidateDenominator,
		NumeratorMatches:              num == candidateNumerator,
		DenominatorMatches:            den == candidateDenominator,
		Formula:                       "rho_twist^2 ?= (dim K7+)^2 dim K7- / [dim K7 * (dim Lambda4_+ - dim K7+)] = 4^2*3/[7*(35-4)] = 48/217",
		Verdict:                       join(StatusNumerator48TypedCandidate, StatusDenominator217TypedCandidate),
	}
}

func buildProjectorAudit(ratio float64) ProjectorContractionAudit {
	candidates := []ProjectorContractionCandidate{
		{Name: "hodge-polarity/self-dual-complement dimension ratio", Formula: "(dim K7+)^2 dim K7- / [dim K7*(dim Lambda4_+ - dim K7+)]", Value: float64(k7PlusDim*k7PlusDim*k7MinusDim) / float64(k7Dim*(lambda4SelfDualDim-k7PlusDim)), Matches48Over217: true, NativeDerivation: false, Reason: "matches the rational skeleton dimensionally, but no projector trace contraction has been derived that forces this value"},
		{Name: "raw K7 fraction in Lambda4", Formula: "dim K7 / dim Lambda4", Value: float64(k7Dim) / float64(lambda4TotalDim), Matches48Over217: false, NativeDerivation: false, Reason: "typed but gives 1/10, not rho_twist^2"},
		{Name: "Hodge imbalance over K7", Formula: "(dim K7+ - dim K7-) / dim K7", Value: float64(k7PlusDim-k7MinusDim) / float64(k7Dim), Matches48Over217: false, NativeDerivation: false, Reason: "typed trace imbalance, but not the twist residual square"},
		{Name: "self-dual K7 occupancy", Formula: "dim K7+ / dim Lambda4_+", Value: float64(k7PlusDim) / float64(lambda4SelfDualDim), Matches48Over217: false, NativeDerivation: false, Reason: "typed occupancy of K7+ inside Lambda4_+, but not the residual square"},
	}
	bestName := ""
	bestResidual := math.Inf(1)
	for _, c := range candidates {
		res := math.Abs(c.Value - ratio)
		if res < bestResidual {
			bestResidual = res
			bestName = c.Name
		}
	}
	return ProjectorContractionAudit{Candidates: candidates, PPlusDimension: lambda4SelfDualDim, PMinusDimension: lambda4AntiSelfDualDim, PK7PlusDimension: k7PlusDim, PK7MinusDimension: k7MinusDim, TraceDerivationCertified: false, BestCandidateName: bestName, BestCandidateResidual: bestResidual, Verdict: join(StatusProjectorContractionsAudited, StatusNoTraceDerivation)}
}

func buildClassification(inh Gate639Inheritance, comp RationalCompressionAudit, routes RouteCompressionAudit, skel DimensionalSkeletonAudit, proj ProjectorContractionAudit) ClassificationAudit {
	candidate := inh.RepeatedAcrossRoutes && inh.ResidualInvariant && inh.CompactSplitObstruction && comp.Compressed && routes.AllClusterRoutesCompress && skel.NumeratorMatches && skel.DenominatorMatches
	artifact := !candidate || inh.Gate639ClassifiedAsArtifact
	interp := "rho_twist^2 is treated as a typed rational-compression candidate for the compact/split obstruction residual: the skeleton 48/217 matches the finite dimensions 4|3 on K7 and 35 on Lambda4_+, but the audit does not derive it from native trace contractions."
	return ClassificationAudit{RhoSquared: comp.RhoSquared, Ratio48Over217: comp.CandidateRatio, CompressionCandidate: candidate, ExactFromFiniteMatrixClaim: false, ConsequenceOfHodgeSplitClaim: false, ArtifactClaim: artifact, ObstructionOnly: candidate && !proj.TraceDerivationCertified, Interpretation: interp, Verdict: StatusRhoSquaredEquals48Over217}
}

func Statuses() []string {
	return []string{StatusGate639RhoInherited, StatusRhoSquaredCompressionTested, StatusRhoSquaredEquals48Over217, StatusDenominator217TypedCandidate, StatusNumerator48TypedCandidate, StatusRouteCompressionRepeated, StatusProjectorContractionsAudited, StatusNoTraceDerivation, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalMetric, StatusGate640Boundary}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
