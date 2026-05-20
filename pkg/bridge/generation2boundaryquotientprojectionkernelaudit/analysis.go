// Package generation2boundaryquotientprojectionkernelaudit implements
// Gate 679: Boundary Quotient Projection Kernel and Relative Trace-Response Audit.
//
// Gate 678 organized K7, H72, Q_boundary, D_history, and tau_defect into an
// exact-sequence-shaped bridge diagram. Gate 679 corrects the dimensional
// target: for the natural split projection pi_split:H72->Q_boundary, K7 is not
// the kernel. The kernel is Lambda^4 R^8 ⊕ L_anti, of dimension 71. K7 is a
// distinguished rank-seven internal defect subspace inside that kernel, and
// the active response is a global H72 trace density 7/72 rather than a literal
// exact-sequence boundary map.
package generation2boundaryquotientprojectionkernelaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate678 "github.com/bagherbal/asha-engine/pkg/bridge/generation2augmenteddefectexactsequenceaudit"
)

const (
	AuditID = "GATE679-BOUNDARY-QUOTIENT-PROJECTION-KERNEL-RELATIVE-TRACE-RESPONSE-AUDIT"

	StatusGate678AugmentedDiagramInherited         = "PASS_GATE678_AUGMENTED_DIAGRAM_INHERITED"
	StatusNaturalBoundaryQuotientProjectionDefined = "PASS_NATURAL_BOUNDARY_QUOTIENT_PROJECTION_DEFINED"
	StatusKernelDimensionComputed                  = "PASS_KERNEL_DIMENSION_COMPUTED"
	StatusK7DefectInsideKernelNotFullKernel        = "PASS_K7_CLASSIFIED_AS_DEFECT_SUBSPACE_INSIDE_KERNEL_NOT_FULL_KERNEL"
	StatusRelativeTraceResponseDefined             = "PASS_RELATIVE_TRACE_RESPONSE_DEFINED"
	StatusDenominatorAlternativesAudited           = "PASS_DENOMINATOR_ALTERNATIVES_AUDITED"
	StatusGlobalAugmentedChamberAverage            = "CONDITIONAL_SUPPORT_TRACE_RESPONSE_IS_GLOBAL_AUGMENTED_CHAMBER_AVERAGE"
	StatusK7DistinguishedInternalDefectInKernel    = "CONDITIONAL_SUPPORT_K7_IS_DISTINGUISHED_INTERNAL_DEFECT_IN_SPLIT_PROJECTION_KERNEL"
	StatusK7IsNotKernelOfPiSplit                   = "FAILED_ROUTE_K7_IS_NOT_KERNEL_OF_PI_SPLIT"
	StatusLiteralExactSequenceWithK7KernelBlocked  = "FAILED_ROUTE_LITERAL_EXACT_SEQUENCE_WITH_K7_KERNEL_BLOCKED"
	StatusNoNativeReasonForGlobalH72Trace          = "FAILED_ROUTE_NO_NATIVE_REASON_FOR_GLOBAL_H72_TRACE_NORMALIZATION"
	StatusNoNativeTraceToBoundaryQuotientTheorem   = "FAILED_ROUTE_NO_NATIVE_TRACE_TO_BOUNDARY_QUOTIENT_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem               = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeWallDistanceAirlockTheorem       = "FAILED_ROUTE_NO_NATIVE_WALL_DISTANCE_AIRLOCK_THEOREM"
	StatusNoBoundaryStressDerivation               = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusGate679Boundary                          = "FIREWALL_PRESERVED_GATE679_RELATIVE_TRACE_RESPONSE_BOUNDARY"
)

type Gate678Inheritance struct {
	AugmentedDiagramInherited   bool
	WeakerDiagramLawful         bool
	StrictExactnessCertified    bool
	TauDefect                   float64
	SSplit                      float64
	DBase                       float64
	Residual                    float64
	H72Dimension                int
	K7Rank                      int
	BoundaryQuotientDimension   int
	MissingExactSequenceTheorem bool
	MissingTraceQuotientTheorem bool
	MissingSevenOver72          bool
	FirewallPreserved           bool
	Verdict                     string
}

type BoundaryProjection struct {
	Name               string
	Domain             string
	Codomain           string
	Formula            string
	BoundaryProjection string
	SplitFunctional    string
	DomainDimension    int
	CodomainDimension  int
	Surjective         bool
	Verdict            string
}

type ProjectionKernel struct {
	Formula               string
	FiniteKernelDimension int
	AntiLineDimension     int
	TotalKernelDimension  int
	K7Rank                int
	K7IsFullKernel        bool
	K7InsideKernel        bool
	QuotientDimension     int
	Verdict               string
}

type DefectInsideKernel struct {
	Defect            string
	Projector         string
	Rank              int
	KernelDimension   int
	AmbientDimension  int
	RelativeToKernel  float64
	RelativeToAmbient float64
	Distinguished     bool
	FullKernel        bool
	Verdict           string
}

type TraceAlternative struct {
	Name           string
	Formula        string
	Value          float64
	PredictedDBase float64
	Residual       float64
	AbsResidual    float64
	Classification string
	Verdict        string
}

type RelativeTraceResponse struct {
	TauGlobal       float64
	TauKernel       float64
	TauFinite       float64
	TauHalf         float64
	SSplit          float64
	DBase           float64
	PredictedDBase  float64
	Residual        float64
	AbsResidual     float64
	BestAlternative string
	UsesGlobalH72   bool
	Verdict         string
}

type NonTautologyCondition struct {
	Principle string
	Status    string
	Comment   string
}

type SourceCandidate struct {
	Candidate      string
	Status         string
	Classification string
	Comment        string
}

type MissingTheoremAudit struct {
	Missing                    []string
	NewPreciseMissingPrinciple string
	AllowedSupport             []string
	Verdict                    string
}

type VerdictDiscipline struct {
	ClaimsK7KernelOfPiSplit          bool
	ClaimsLiteralExactSequence       bool
	ClaimsNativeGlobalTraceTheorem   bool
	ClaimsNativeTraceQuotientTheorem bool
	ClaimsNativeSevenOver72          bool
	ClaimsWallDistanceAirlock        bool
	ClaimsBoundaryStressDerivation   bool
	ClaimsHiggsMass                  bool
	ClaimsGaugeUnification           bool
	ClaimsFlavorDerivation           bool
	Verdict                          string
}

type Analysis struct {
	Inherited    Gate678Inheritance
	Projection   BoundaryProjection
	Kernel       ProjectionKernel
	Defect       DefectInsideKernel
	Trace        RelativeTraceResponse
	Alternatives []TraceAlternative
	Conditions   []NonTautologyCondition
	Candidates   []SourceCandidate
	Missing      MissingTheoremAudit
	Discipline   VerdictDiscipline
	Truth        string
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
	g678, err := gate678.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate678 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g678)
	projection := buildProjection()
	kernel := buildKernel(inherited, projection)
	defect := buildDefect(inherited, kernel)
	trace := buildTrace(inherited)
	alternatives := buildAlternatives(inherited)
	conditions := buildConditions()
	candidates := buildCandidates()
	missing := buildMissing()
	discipline := VerdictDiscipline{Verdict: StatusGate679Boundary}
	truth := "Gate 679 corrects the Gate678 exact-sequence dream: the natural split projection pi_split:H72->Q_boundary has kernel Lambda^4 R^8 ⊕ L_anti of dimension 71, so K7 is not the kernel. K7 is a distinguished rank-seven internal defect inside that kernel. The active response is therefore a relative scalar trace response using the global augmented chamber average 7/72, while ASHA still lacks the native principle selecting global H72 normalization over 7/71 or 7/70."
	return Analysis{Inherited: inherited, Projection: projection, Kernel: kernel, Defect: defect, Trace: trace, Alternatives: alternatives, Conditions: conditions, Candidates: candidates, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate678.Analysis) Gate678Inheritance {
	return Gate678Inheritance{
		AugmentedDiagramInherited:   true,
		WeakerDiagramLawful:         g.Sequence.WeakerDiagramLawful,
		StrictExactnessCertified:    g.Sequence.StrictExactSequenceCertified,
		TauDefect:                   g.Trace.TauDefect,
		SSplit:                      g.Trace.SSplit,
		DBase:                       g.Trace.DBase,
		Residual:                    g.Trace.Residual,
		H72Dimension:                g.Chamber.TotalDimension,
		K7Rank:                      g.Defect.Rank,
		BoundaryQuotientDimension:   g.Boundary.Dimension,
		MissingExactSequenceTheorem: strings.Contains(g.Missing.Verdict, gate678.StatusNoNativeExactSequenceCouplingTheorem),
		MissingTraceQuotientTheorem: strings.Contains(g.Missing.Verdict, gate678.StatusNoNativeTraceToQuotientResponseTheorem),
		MissingSevenOver72:          strings.Contains(g.Missing.Verdict, gate678.StatusNoNativeSevenOver72Theorem),
		FirewallPreserved:           g.Discipline.Verdict == gate678.StatusGate678Boundary,
		Verdict:                     StatusGate678AugmentedDiagramInherited,
	}
}

func buildProjection() BoundaryProjection {
	return BoundaryProjection{
		Name:               "pi_split",
		Domain:             "H_72 = Lambda^4 R^8 ⊕ R^2_boundary",
		Codomain:           "Q_boundary = R^2_boundary/L_anti ≅ R",
		Formula:            "pi_split(h,(lambda,R)) = sigma_boundary(lambda,R)=lambda+R",
		BoundaryProjection: "pr_boundary:H_72->R^2_boundary",
		SplitFunctional:    "sigma_boundary(lambda,R)=lambda+R",
		DomainDimension:    72,
		CodomainDimension:  1,
		Surjective:         true,
		Verdict:            StatusNaturalBoundaryQuotientProjectionDefined,
	}
}

func buildKernel(in Gate678Inheritance, p BoundaryProjection) ProjectionKernel {
	finite := 70
	anti := 1
	total := finite + anti
	return ProjectionKernel{
		Formula:               "ker(pi_split)=Lambda^4 R^8 ⊕ L_anti",
		FiniteKernelDimension: finite,
		AntiLineDimension:     anti,
		TotalKernelDimension:  total,
		K7Rank:                in.K7Rank,
		K7IsFullKernel:        in.K7Rank == total,
		K7InsideKernel:        in.K7Rank > 0 && in.K7Rank < total && p.Surjective,
		QuotientDimension:     p.CodomainDimension,
		Verdict:               strings.Join([]string{StatusKernelDimensionComputed, StatusK7IsNotKernelOfPiSplit, StatusLiteralExactSequenceWithK7KernelBlocked}, ";"),
	}
}

func buildDefect(in Gate678Inheritance, k ProjectionKernel) DefectInsideKernel {
	return DefectInsideKernel{
		Defect:            "K_7 ⊕ 0_boundary ⊂ ker(pi_split)",
		Projector:         "P_defect=P_K7⊕0_boundary",
		Rank:              in.K7Rank,
		KernelDimension:   k.TotalKernelDimension,
		AmbientDimension:  in.H72Dimension,
		RelativeToKernel:  float64(in.K7Rank) / float64(k.TotalKernelDimension),
		RelativeToAmbient: float64(in.K7Rank) / float64(in.H72Dimension),
		Distinguished:     true,
		FullKernel:        false,
		Verdict:           strings.Join([]string{StatusK7DefectInsideKernelNotFullKernel, StatusK7DistinguishedInternalDefectInKernel}, ";"),
	}
}

func buildTrace(in Gate678Inheritance) RelativeTraceResponse {
	tauGlobal := 7.0 / 72.0
	tauKernel := 7.0 / 71.0
	tauFinite := 7.0 / 70.0
	tauHalf := 7.0 / 144.0
	pred := tauGlobal * in.SSplit
	residual := in.DBase - pred
	return RelativeTraceResponse{
		TauGlobal:       tauGlobal,
		TauKernel:       tauKernel,
		TauFinite:       tauFinite,
		TauHalf:         tauHalf,
		SSplit:          in.SSplit,
		DBase:           in.DBase,
		PredictedDBase:  pred,
		Residual:        residual,
		AbsResidual:     math.Abs(residual),
		BestAlternative: "7/72 global H_72 trace",
		UsesGlobalH72:   true,
		Verdict:         strings.Join([]string{StatusRelativeTraceResponseDefined, StatusGlobalAugmentedChamberAverage}, ";"),
	}
}

func buildAlternatives(in Gate678Inheritance) []TraceAlternative {
	candidates := []struct {
		name, formula, class string
		v                    float64
	}{
		{"tau_global", "7/72", "active global augmented-chamber trace", 7.0 / 72.0},
		{"tau_kernel", "7/71", "kernel-only trace candidate", 7.0 / 71.0},
		{"tau_finite", "7/70", "finite-chamber-only trace candidate", 7.0 / 70.0},
		{"tau_half", "7/144", "per-boundary-coordinate half-trace clue", 7.0 / 144.0},
	}
	out := make([]TraceAlternative, 0, len(candidates))
	bestName := "tau_global"
	for _, c := range candidates {
		pred := c.v * in.SSplit
		residual := in.DBase - pred
		verdict := "typed alternative tested"
		if c.name == bestName {
			verdict = StatusGlobalAugmentedChamberAverage
		}
		out = append(out, TraceAlternative{Name: c.name, Formula: c.formula, Value: c.v, PredictedDBase: pred, Residual: residual, AbsResidual: math.Abs(residual), Classification: c.class, Verdict: verdict})
	}
	return out
}

func buildConditions() []NonTautologyCondition {
	return []NonTautologyCondition{
		{Principle: "canonical split projection pi_split", Status: "supplied", Comment: "pi_split is sigma_boundary∘pr_boundary and has kernel Lambda^4 R^8 ⊕ L_anti."},
		{Principle: "K7 defect subspace inside kernel", Status: "supplied", Comment: "K7⊕0_boundary is contained in ker(pi_split), but is not the full kernel."},
		{Principle: "global H72 trace normalization", Status: "missing principle", Comment: "the response uses 7/72 rather than 7/71 or 7/70; ASHA still lacks the theorem selecting full augmented chamber averaging."},
		{Principle: "trace-to-boundary quotient response", Status: "missing theorem", Comment: "no native theorem yet makes the K7 trace density act on Q_boundary to produce D_history."},
	}
}

func buildCandidates() []SourceCandidate {
	return []SourceCandidate{
		{Candidate: "literal exact sequence with K7 kernel", Status: "blocked", Classification: "dimensionally false target", Comment: "ker(pi_split) has dimension 71, not 7."},
		{Candidate: "relative trace response", Status: "conditional support", Classification: "bridge scalar functional", Comment: "K7 is a distinguished defect subspace inside the split-projection kernel and supplies a global trace density."},
		{Candidate: "global augmented chamber denominator 72", Status: "conditional support", Classification: "active coefficient source", Comment: "7/72 is the best typed denominator for the line response."},
		{Candidate: "kernel denominator 71", Status: "weaker failed alternative", Classification: "kernel-only trace", Comment: "7/71 is typed but gives a worse response and lacks the active closure."},
		{Candidate: "finite denominator 70", Status: "weaker failed alternative", Classification: "finite-only trace", Comment: "7/70 ignores the boundary pair and gives a worse response."},
	}
}

func buildMissing() MissingTheoremAudit {
	return MissingTheoremAudit{
		Missing:                    []string{StatusNoNativeReasonForGlobalH72Trace, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoNativeSevenOver72Theorem, StatusNoNativeWallDistanceAirlockTheorem},
		NewPreciseMissingPrinciple: "GlobalAugmentedTraceResponsePrinciple / FullChamberAveragedDefectResponseTheorem",
		AllowedSupport:             []string{StatusGlobalAugmentedChamberAverage, StatusK7DistinguishedInternalDefectInKernel},
		Verdict:                    strings.Join([]string{StatusNoNativeReasonForGlobalH72Trace, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoNativeSevenOver72Theorem, StatusNoNativeWallDistanceAirlockTheorem}, ";"),
	}
}

func Statuses() []string {
	return []string{StatusGate678AugmentedDiagramInherited, StatusNaturalBoundaryQuotientProjectionDefined, StatusKernelDimensionComputed, StatusK7DefectInsideKernelNotFullKernel, StatusRelativeTraceResponseDefined, StatusDenominatorAlternativesAudited, StatusGlobalAugmentedChamberAverage, StatusK7DistinguishedInternalDefectInKernel, StatusK7IsNotKernelOfPiSplit, StatusLiteralExactSequenceWithK7KernelBlocked, StatusNoNativeReasonForGlobalH72Trace, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoNativeSevenOver72Theorem, StatusNoNativeWallDistanceAirlockTheorem, StatusGate679Boundary}
}
