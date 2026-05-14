// Package coarsegrain implements Gate 105: the native finite coarse-graining /
// threshold activation operator search.
//
// Gate 104 proved that the finite boundary data are only relative until an
// absolute coupling unit, a boundary scale, and threshold activation rules are
// derived.  This package searches for the next missing object: a native finite
// RG/coarse-graining operator.  Such an operator would have to do more than list
// spectra.  It must define a canonical finite flow step, compose as a semigroup,
// select a scale/log parameter, and decide when finite modes become continuum
// active or decouple.
//
// The result is deliberately hygienic.  Several useful finite diagnostics exist:
// projection maps, quotient maps, spectral orderings, contact-overlap anchors,
// action weights, and continuum beta slopes.  But none of them currently forms a
// native RG operator.  Therefore threshold corrections and physical running
// constants remain sealed.
package coarsegrain

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/boundaryselector"
	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdactivation"
)

type OperatorKind string

const (
	OperatorProjection          OperatorKind = "projection"
	OperatorQuotient            OperatorKind = "quotient"
	OperatorSpectralOrdering    OperatorKind = "spectral-ordering"
	OperatorActionWeight        OperatorKind = "action-weight"
	OperatorContinuumBeta       OperatorKind = "continuum-beta-diagnostic"
	OperatorThresholdClassifier OperatorKind = "threshold-classifier"
)

type CandidateOperator struct {
	Name   string
	Symbol string
	Kind   OperatorKind

	FiniteDataAvailable          bool
	EndomorphismLike             bool
	IdempotentOrStatic           bool
	SemigroupLawDerived          bool
	ScaleParameterDerived        bool
	FixedPointSelected           bool
	ThresholdPredicateDerived    bool
	DecouplingRuleDerived        bool
	AbsoluteCouplingRenormalized bool
	NativeFiniteRGStep           bool

	RejectedAsRGStep bool
	Detail           string
}

type Requirement struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type ActivationScheduleWitness struct {
	Name                       string
	Rule                       string
	CompatibleWithData         bool
	ThresholdCorrectionAllowed bool
	Detail                     string
}

type Analysis struct {
	Boundary   boundaryselector.Analysis
	Thresholds thresholdactivation.Analysis

	CandidateOperators []CandidateOperator
	OperatorCount      int
	Requirements       []Requirement

	FiniteBoundarySeedInherited bool
	BoundaryKY                  float64
	BoundarySin2                float64
	RelativeHessianTrace        float64
	ThresholdInventoryInherited bool
	ThresholdDecisionCount      int

	ProjectionOperatorsFound            bool
	SpectralAnchorsFound                bool
	StaticClassifiersFound              bool
	NativeCoarseGrainingFound           bool
	SemigroupLawDerived                 bool
	ScaleLogParameterDerived            bool
	FlowFixedPointSelected              bool
	ThresholdActivationPredicateDerived bool
	DecouplingMatchingRuleDerived       bool
	AbsoluteCouplingRunningDerived      bool

	OpenThresholdModes           int
	VacuumOnlyModes              int
	ContinuumCandidates          int
	ActivationSchedules          []ActivationScheduleWitness
	NonUniqueActivationWitnessed bool

	ResidualNullityBefore  int
	ResidualNullityAfter   int
	ResidualSymmetryBroken bool

	PhysicalWeakAngleDerived bool
	FineStructureDerived     bool
	PhysicalMassesDerived    bool
	HiddenObservedInputUsed  bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		boundary, err := boundaryselector.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		thresholds, err := thresholdactivation.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(boundary, thresholds)
	})
	return defaultValue, defaultErr
}

func Build(boundary boundaryselector.Analysis, thresholds thresholdactivation.Analysis) (Analysis, error) {
	if !boundary.FiniteBoundarySeedSelected || !boundary.RelativeKineticNormalizationComplete {
		return Analysis{}, fmt.Errorf("Gate 105 requires Gate 104 finite boundary seed")
	}
	if boundary.BoundaryCouplingDerived || boundary.BoundaryScaleDerived || boundary.ThresholdRuleDerived || boundary.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 105 refuses already-physicalized boundary data or hidden observed input")
	}
	if thresholds.CandidateCount == 0 {
		return Analysis{}, fmt.Errorf("Gate 105 requires threshold inventory")
	}
	if thresholds.HiddenScaleInserted {
		return Analysis{}, fmt.Errorf("Gate 105 refuses hidden threshold scale insertion")
	}

	operators := []CandidateOperator{
		{
			Name:                "finite active-sector projection",
			Symbol:              "P_active",
			Kind:                OperatorProjection,
			FiniteDataAvailable: true,
			EndomorphismLike:    true,
			IdempotentOrStatic:  true,
			RejectedAsRGStep:    true,
			Detail:              "projects onto the active scalar/contact carrier; useful sector selection, but an idempotent projection is not an iterated scale flow",
		},
		{
			Name:                "gauge quotient / protected-broken quotient",
			Symbol:              "q: carrier -> carrier/O(3)",
			Kind:                OperatorQuotient,
			FiniteDataAvailable: true,
			EndomorphismLike:    true,
			IdempotentOrStatic:  true,
			RejectedAsRGStep:    true,
			Detail:              "removes gauge redundancy and fixes orientation classes; it is a kinematic quotient, not a running/coarse-graining semigroup",
		},
		{
			Name:                "B-sector spectral truncation family",
			Symbol:              "T_ε(B)",
			Kind:                OperatorSpectralOrdering,
			FiniteDataAvailable: true,
			EndomorphismLike:    true,
			IdempotentOrStatic:  false,
			RejectedAsRGStep:    true,
			Detail:              "spectral ordering exists, but the cutoff ε or scale parameter is arbitrary until a native rule selects it",
		},
		{
			Name:                "contact partial-overlap ordering",
			Symbol:              "spec(P_contact)",
			Kind:                OperatorSpectralOrdering,
			FiniteDataAvailable: true,
			EndomorphismLike:    false,
			IdempotentOrStatic:  true,
			RejectedAsRGStep:    true,
			Detail:              "orders finite overlap modes; does not decide whether they are thresholds, regulators, or vacuum-frustration modes",
		},
		{
			Name:                "topological action weight",
			Symbol:              "exp(-8π²I_BG)",
			Kind:                OperatorActionWeight,
			FiniteDataAvailable: true,
			EndomorphismLike:    false,
			IdempotentOrStatic:  true,
			RejectedAsRGStep:    true,
			Detail:              "dimensionless action weight; no composition law, no scale derivative, and no threshold crossing predicate",
		},
		{
			Name:                "continuum one-loop beta vector",
			Symbol:              "b=(41/10,-19/6,-7)",
			Kind:                OperatorContinuumBeta,
			FiniteDataAvailable: true,
			EndomorphismLike:    false,
			IdempotentOrStatic:  false,
			RejectedAsRGStep:    true,
			Detail:              "a continuum diagnostic attached to the finite inventory; it supplies slopes after a continuum assumption, not a native finite RG map",
		},
		{
			Name:                "threshold activation classifier",
			Symbol:              "A(mode)∈{candidate,open,vacuum-only}",
			Kind:                OperatorThresholdClassifier,
			FiniteDataAvailable: true,
			EndomorphismLike:    false,
			IdempotentOrStatic:  true,
			RejectedAsRGStep:    true,
			Detail:              "classifies present knowledge state; it does not derive activation times, matching scales, or Δb_i corrections",
		},
	}

	projectionFound := false
	spectralFound := false
	staticFound := false
	nativeFound := false
	semigroup := false
	scaleLog := false
	fixedPoint := false
	thresholdPredicate := false
	decoupling := false
	couplingRunning := false
	for _, op := range operators {
		if op.Kind == OperatorProjection || op.Kind == OperatorQuotient {
			projectionFound = true
		}
		if op.Kind == OperatorSpectralOrdering {
			spectralFound = true
		}
		if op.IdempotentOrStatic {
			staticFound = true
		}
		if op.NativeFiniteRGStep {
			nativeFound = true
		}
		if op.SemigroupLawDerived {
			semigroup = true
		}
		if op.ScaleParameterDerived {
			scaleLog = true
		}
		if op.FixedPointSelected {
			fixedPoint = true
		}
		if op.ThresholdPredicateDerived {
			thresholdPredicate = true
		}
		if op.DecouplingRuleDerived {
			decoupling = true
		}
		if op.AbsoluteCouplingRenormalized {
			couplingRunning = true
		}
	}

	open, vacuumOnly, continuum := 0, 0, 0
	for _, d := range thresholds.Decisions {
		switch d.Status {
		case thresholdactivation.ThresholdOpen, thresholdactivation.IntegratedOutOpen, thresholdactivation.Unclassified:
			open++
		case thresholdactivation.VacuumFrustrationOnly:
			vacuumOnly++
		case thresholdactivation.ContinuumFieldCandidate:
			continuum++
		}
	}

	schedules := []ActivationScheduleWitness{
		{
			Name:                       "minimal inactive-open schedule",
			Rule:                       "only already-established continuum candidates remain active; all open threshold modes are withheld from Δb_i",
			CompatibleWithData:         true,
			ThresholdCorrectionAllowed: false,
			Detail:                     "matches the current firewall because no open mode has a derived activation predicate",
		},
		{
			Name:                       "hypothetical ordered-overlap schedule",
			Rule:                       "activate open contact/B spectral modes by descending finite eigenvalue after a future cutoff law is supplied",
			CompatibleWithData:         true,
			ThresholdCorrectionAllowed: false,
			Detail:                     "also compatible with current finite data, but cannot be used yet because the cutoff law and representation/matching maps are not derived",
		},
	}

	requirements := []Requirement{
		{Name: "finite endomap or functor on the carrier/state algebra", Required: true, Satisfied: projectionFound, Detail: "projection/quotient candidates exist"},
		{Name: "nontrivial composable RG semigroup", Required: true, Satisfied: semigroup, Detail: "no C_a∘C_b=C_ab or additive-L law is derived; current maps are static/idempotent or diagnostic"},
		{Name: "canonical scale/log parameter", Required: true, Satisfied: scaleLog, Detail: "no finite rule selects L, ε, shell number, or boundary scale"},
		{Name: "fixed point or stationary boundary condition", Required: true, Satisfied: fixedPoint, Detail: "K_*=diag(1,1,1,5/3) is a boundary seed, not a fixed point of a finite flow"},
		{Name: "threshold crossing predicate", Required: true, Satisfied: thresholdPredicate, Detail: "open modes have no derived activation predicate"},
		{Name: "decoupling/matching contribution Δb_i", Required: true, Satisfied: decoupling, Detail: "no open threshold mode is allowed to correct beta coefficients"},
		{Name: "absolute coupling renormalization", Required: true, Satisfied: couplingRunning, Detail: "relative Hessian is preserved under gauge-action prefactor rescaling"},
	}

	nonUnique := len(schedules) >= 2 && schedules[0].CompatibleWithData && schedules[1].CompatibleWithData
	trace := 0.0
	if t, err := boundary.RG.EmbeddedBoundaryHessian.Trace(); err == nil {
		trace = t
	}
	truth := "Gate 105 searched for a native finite coarse-graining/RG-step operator. Projection, quotient, spectral-ordering, contact-overlap, action-weight, continuum-beta, and threshold-classifier candidates are all real pieces of the finite engine. But none currently supplies the required semigroup law, canonical scale parameter, fixed point, threshold crossing predicate, decoupling map, or absolute coupling renormalization. Therefore the project has not failed; it has isolated the exact next missing structure: a finite RG operator, not another boundary normalization."

	return Analysis{
		Boundary:                            boundary,
		Thresholds:                          thresholds,
		CandidateOperators:                  operators,
		OperatorCount:                       len(operators),
		Requirements:                        requirements,
		FiniteBoundarySeedInherited:         true,
		BoundaryKY:                          boundary.BoundaryKY,
		BoundarySin2:                        boundary.BoundarySin2,
		RelativeHessianTrace:                trace,
		ThresholdInventoryInherited:         thresholds.CandidateCount > 0,
		ThresholdDecisionCount:              len(thresholds.Decisions),
		ProjectionOperatorsFound:            projectionFound,
		SpectralAnchorsFound:                spectralFound,
		StaticClassifiersFound:              staticFound,
		NativeCoarseGrainingFound:           nativeFound,
		SemigroupLawDerived:                 semigroup,
		ScaleLogParameterDerived:            scaleLog,
		FlowFixedPointSelected:              fixedPoint,
		ThresholdActivationPredicateDerived: thresholdPredicate,
		DecouplingMatchingRuleDerived:       decoupling,
		AbsoluteCouplingRunningDerived:      couplingRunning,
		OpenThresholdModes:                  open,
		VacuumOnlyModes:                     vacuumOnly,
		ContinuumCandidates:                 continuum,
		ActivationSchedules:                 schedules,
		NonUniqueActivationWitnessed:        nonUnique,
		ResidualNullityBefore:               boundary.EquationAudit.Nullity,
		ResidualNullityAfter:                boundary.EquationAudit.Nullity,
		ResidualSymmetryBroken:              false,
		PhysicalWeakAngleDerived:            false,
		FineStructureDerived:                false,
		PhysicalMassesDerived:               false,
		HiddenObservedInputUsed:             false,
		TruthStatement:                      truth,
		RejectedClaims: []string{
			"a spectral ordering is already a finite RG flow",
			"an idempotent projection or quotient is enough to define running couplings",
			"contact partial-overlap modes may correct beta coefficients before an activation predicate exists",
			"the continuum one-loop beta vector is itself the native finite coarse-graining theorem",
			"different admissible threshold schedules can be collapsed by choice rather than derivation",
		},
		RemainingUnknowns: []string{
			"U-25A-COARSE-GRAINING-FUNCTOR: define the finite state/carrier map C_s and its domain/codomain",
			"U-25B-SEMIGROUP-LAW: prove C_s∘C_t=C_{s+t} or an equivalent finite shell-composition law",
			"U-25C-SCALE-PARAMETER: derive the native finite variable that becomes L=ln(M*/μ) or replaces it",
			"U-25D-THRESHOLD-PREDICATE: derive when a finite mode is active, decoupled, regulator-only, or vacuum-only",
			"U-25E-MATCHING-MAP: derive each active threshold contribution to Δb_1, Δb_2, Δb_3 without observed scale fitting",
			"U-25F-PREFactor-FLOW: derive whether the absolute gauge-action prefactor runs or remains convention-only",
		},
		RecommendedNextGate: "Gate 106 — finite shell functor / semigroup construction attempt",
	}, nil
}

func FormatOperators(xs []CandidateOperator) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		verdict := "not-rg-step"
		if x.NativeFiniteRGStep {
			verdict = "native-rg-step"
		}
		parts = append(parts, fmt.Sprintf("%s:%s [%s, endomap=%t, semigroup=%t, scale=%t, threshold=%t, %s]", x.Symbol, x.Name, x.Kind, x.EndomorphismLike, x.SemigroupLawDerived, x.ScaleParameterDerived, x.ThresholdPredicateDerived, verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatRequirements(xs []Requirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%t (%s)", x.Name, x.Satisfied, x.Detail))
	}
	return strings.Join(parts, "; ")
}

func FormatSchedules(xs []ActivationScheduleWitness) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: compatible=%t, threshold-correction-allowed=%t, rule=%s", x.Name, x.CompatibleWithData, x.ThresholdCorrectionAllowed, x.Rule))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
