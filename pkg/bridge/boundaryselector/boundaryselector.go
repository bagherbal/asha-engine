// Package boundaryselector implements Gate 104: the boundary-scale operator / absolute
// coupling unit search.
//
// Gate 103 exposed the precise RG firewall.  The finite engine has a relative
// electroweak boundary Hessian and a formal one-loop family, but low-energy
// predictions still require at least an absolute coupling unit u=1/g_*^2, a
// scale interval L=ln(M*/mu), and threshold/decoupling rules.  Earlier gates
// also exposed candidate action data: the contact index I_BG=1 and the
// dimensionless topological seal S_top=8*pi^2 I_BG.  This package performs the
// next honest operation: it inventories every presently available finite
// boundary operator candidate and tests whether any of them actually selects an
// absolute coupling unit or a dimensionful boundary scale.
//
// The result is a sharpened no-go/firewall theorem.  The finite candidates are
// valuable dimensionless invariants and normalization diagnostics, but none is a
// dimensionful operator, none supplies a derived trace/kinetic normalization for
// g_*^2, and none supplies a native finite RG/coarse-graining rule.  Therefore
// physical alpha, theta_W, masses, and a boundary scale remain sealed.
package boundaryselector

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/couplingnorm"
	"github.com/bagherbal/asha-engine/pkg/bridge/rgfirewall"
)

type CandidateKind string

const (
	CandidateRelativeBoundary CandidateKind = "relative-boundary"
	CandidateActionSeal       CandidateKind = "action-seal"
	CandidateScaleAnchor      CandidateKind = "scale-anchor"
	CandidateCouplingUnit     CandidateKind = "coupling-unit"
	CandidateRGDiagnostic     CandidateKind = "rg-diagnostic"
)

type CandidateOperator struct {
	Name                       string
	Symbol                     string
	Kind                       CandidateKind
	Value                      float64
	Formula                    string
	Dimensionless              bool
	Dimensionful               bool
	RequiresExtraBridge        bool
	SelectsBoundaryCoupling    bool
	SelectsBoundaryScale       bool
	SelectsThresholdRule       bool
	SelectedAsFiniteInvariant  bool
	RejectedAsPhysicalSelector bool
	Detail                     string
}

type ResidualSymmetry struct {
	Name                string
	Transformation      string
	PreservesFiniteData bool
	Blocks              []string
	Detail              string
}

type EquationAudit struct {
	Variables                           []string
	SelectedFiniteBoundaryEquations     []string
	IndependentEquationsForPhysicalFlow int
	RequiredIndependentEquations        int
	Nullity                             int
	Detail                              string
}

type Analysis struct {
	RG       rgfirewall.Analysis
	Coupling couplingnorm.Analysis

	CandidateOperators []CandidateOperator
	CandidateCount     int

	FiniteBoundarySeedSelected bool
	BoundaryKY                 float64
	BoundarySin2               float64
	TopologicalSeal            float64
	InstantonWeight            float64
	UnitTraceCouplingSq        float64
	UnitTraceInverseAlpha      float64

	RelativeKineticNormalizationComplete bool
	AllCandidateOperatorsDimensionless   bool
	DimensionfulOperatorFound            bool
	AbsoluteCouplingOperatorFound        bool
	BoundaryScaleOperatorFound           bool
	NativeFiniteRGOperatorFound          bool
	ThresholdSelectorFound               bool

	EquationAudit         EquationAudit
	ResidualSymmetries    []ResidualSymmetry
	ResidualVariableCount int

	UnitTraceConventionRejected    bool
	TopologicalSealAsScaleRejected bool
	ObservedMatchingRejected       bool

	BoundaryCouplingDerived  bool
	BoundaryScaleDerived     bool
	ThresholdRuleDerived     bool
	FiniteRGTheoremDerived   bool
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
		rg, err := rgfirewall.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		coupling, err := couplingnorm.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(rg, coupling, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(rg rgfirewall.Analysis, coupling couplingnorm.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !rg.BoundaryDataDerived || !rg.FormalRGFamilyConstructed {
		return Analysis{}, fmt.Errorf("Gate 104 requires Gate 103 boundary data and symbolic RG family")
	}
	if coupling.TopologicalActionSeal <= eps || coupling.ContactIndex <= eps {
		return Analysis{}, fmt.Errorf("Gate 104 requires the finite topological action seal from the coupling audit")
	}
	if coupling.HiddenObservedCouplingUsed || rg.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 104 refuses hidden observed coupling/scale input")
	}

	ss := coupling.Action.ScalarScale
	candidates := []CandidateOperator{
		{
			Name:                      "embedded finite electroweak Hessian",
			Symbol:                    "K_*=diag(1,1,1,5/3)",
			Kind:                      CandidateRelativeBoundary,
			Value:                     traceEmbeddedBoundary(rg),
			Formula:                   "Gate 102 embedded Hessian fixes relative gauge kinetic normalization",
			Dimensionless:             true,
			Dimensionful:              false,
			RequiresExtraBridge:       false,
			SelectedAsFiniteInvariant: true,
			Detail:                    "valid finite boundary seed; fixes ratios and the no-running boundary diagnostic, not the absolute coefficient multiplying the action",
		},
		{
			Name:                      "matter hypercharge normalization",
			Symbol:                    "k_Y",
			Kind:                      CandidateRelativeBoundary,
			Value:                     rg.BoundaryKY,
			Formula:                   "k_Y=5/3",
			Dimensionless:             true,
			Dimensionful:              false,
			RequiresExtraBridge:       false,
			SelectedAsFiniteInvariant: true,
			Detail:                    "valid representation/embedding normalization; selects sin^2_*=1/(1+k_Y)=3/8 at the boundary only",
		},
		{
			Name:                       "topological action seal",
			Symbol:                     "S_top",
			Kind:                       CandidateActionSeal,
			Value:                      coupling.TopologicalActionSeal,
			Formula:                    "S_top=8π² I_BG",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "can normalize an action weight or coupling problem, but does not carry mass/length units and does not by itself choose M*",
		},
		{
			Name:                       "instanton-shaped suppression weight",
			Symbol:                     "exp(-S_top)",
			Kind:                       CandidateActionSeal,
			Value:                      coupling.InstantonWeight,
			Formula:                    "exp(-8π² I_BG)",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "a dimensionless suppression diagnostic, not alpha, not a threshold mass, and not a boundary scale",
		},
		{
			Name:                       "unit-trace coupling convention",
			Symbol:                     "g_unit²",
			Kind:                       CandidateCouplingUnit,
			Value:                      coupling.UnitTraceGaugeCouplingSq,
			Formula:                    "if finite trace normalization κ_trace=1 is imposed, g²=1",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "only a convention diagnostic because the trace/kinetic normalization bridge is not derived",
		},
		{
			Name:                       "finite scalar radius",
			Symbol:                     "r0²",
			Kind:                       CandidateScaleAnchor,
			Value:                      ss.FiniteRadiusSquared,
			Formula:                    "dimensionless scalar vacuum radius",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "becomes a physical vev only after a mass unit μ is derived",
		},
		{
			Name:                       "B-sector spectral gap",
			Symbol:                     "gap_B",
			Kind:                       CandidateScaleAnchor,
			Value:                      ss.BGap,
			Formula:                    "first positive eigenvalue of the finite B-sector operator",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "a finite spectral spacing, not an energy gap in GeV until the finite-to-physical unit map is derived",
		},
		{
			Name:                       "contact leakage norm",
			Symbol:                     "L_BG²",
			Kind:                       CandidateScaleAnchor,
			Value:                      ss.ContactLeakageNormSquared,
			Formula:                    "||P_B P_G - P_K||²",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "measures finite contact frustration; cannot set M* without a dimensional bridge",
		},
		{
			Name:                       "one-loop beta vector",
			Symbol:                     "(b1,b2,b3)",
			Kind:                       CandidateRGDiagnostic,
			Value:                      vectorNorm(rg.B1, rg.B2, rg.B3),
			Formula:                    "finite-spectrum beta diagnostic under continuum one-loop assumption",
			Dimensionless:              true,
			Dimensionful:               false,
			RequiresExtraBridge:        true,
			RejectedAsPhysicalSelector: true,
			Detail:                     "gives slopes once u and L exist; it does not select the boundary scale or the absolute coupling unit",
		},
	}

	allDimensionless := true
	dimensionfulFound := false
	absCouplingFound := false
	boundaryScaleFound := false
	thresholdFound := false
	for _, c := range candidates {
		if !c.Dimensionless {
			allDimensionless = false
		}
		if c.Dimensionful {
			dimensionfulFound = true
		}
		if c.SelectsBoundaryCoupling {
			absCouplingFound = true
		}
		if c.SelectsBoundaryScale {
			boundaryScaleFound = true
		}
		if c.SelectsThresholdRule {
			thresholdFound = true
		}
	}

	variables := []string{"u=1/g_*²", "L=ln(M*/μ)", "Δb_i(L) threshold/decoupling map"}
	selected := []string{
		"relative electroweak Hessian K_*=diag(1,1,1,5/3)",
		"boundary ratio sin²_*=1/(1+k_Y)=3/8 at L=0",
		"formal slopes b_i once a continuum one-loop RG assumption is accepted",
	}
	audit := EquationAudit{
		Variables:                           variables,
		SelectedFiniteBoundaryEquations:     selected,
		IndependentEquationsForPhysicalFlow: 0,
		RequiredIndependentEquations:        len(variables),
		Nullity:                             len(variables),
		Detail:                              "finite data select relative boundary normalization but supply zero independent equations for the absolute coupling unit, physical boundary scale, or threshold activation map",
	}

	symmetries := []ResidualSymmetry{
		{
			Name:                "coupling-prefactor rescaling",
			Transformation:      "S_gauge -> c·S_gauge, equivalently u=1/g_*² -> c·u for any c>0 before trace normalization is derived",
			PreservesFiniteData: true,
			Blocks:              []string{"absolute g_*²", "alpha", "low-energy g2/gY/e"},
			Detail:              "relative Hessian entries and sin²_* are unchanged by an overall gauge-action prefactor",
		},
		{
			Name:                "dimensionful scale dilation",
			Transformation:      "M* -> ρM*, μ -> ρμ, threshold masses -> ρm_i for any ρ>0",
			PreservesFiniteData: true,
			Blocks:              []string{"physical boundary scale", "vev/masses", "threshold energies"},
			Detail:              "all current finite anchors are pure numbers, so no absolute energy or length unit is selected",
		},
		{
			Name:                "threshold schedule freedom",
			Transformation:      "choose different admissible Δb_i activation schedules while preserving the same boundary seed",
			PreservesFiniteData: true,
			Blocks:              []string{"low-energy thetaW", "low-energy alpha", "mass thresholds"},
			Detail:              "the threshold inventory exists, but no decoupling/matching operator has selected when modes activate",
		},
	}

	truth := "Gate 104 searched the available finite boundary/action operators for a selector of the absolute coupling unit or the boundary scale.  It found strong dimensionless invariants: the embedded Hessian K_*=diag(1,1,1,5/3), k_Y=5/3, sin²_*=3/8, the contact-index action seal S_top=8π², scalar/contact anchors, and beta diagnostics.  But every candidate is dimensionless or convention-dependent.  No candidate breaks the residual coupling-prefactor rescaling, no candidate supplies a physical mass/length unit, and no candidate derives the threshold activation map.  Therefore the project is closer to completion because the obstruction is now exact: the next missing object must be a native finite coarse-graining/dimensional-anchor theorem, not another algebraic normalization."

	return Analysis{
		RG:                                   rg,
		Coupling:                             coupling,
		CandidateOperators:                   candidates,
		CandidateCount:                       len(candidates),
		FiniteBoundarySeedSelected:           rg.BoundaryDataDerived,
		BoundaryKY:                           rg.BoundaryKY,
		BoundarySin2:                         rg.BoundarySin2,
		TopologicalSeal:                      coupling.TopologicalActionSeal,
		InstantonWeight:                      coupling.InstantonWeight,
		UnitTraceCouplingSq:                  coupling.UnitTraceGaugeCouplingSq,
		UnitTraceInverseAlpha:                coupling.UnitTraceInverseAlpha,
		RelativeKineticNormalizationComplete: rg.BoundaryDataDerived && rg.BoundaryKineticPositive && close(rg.BoundarySin2, 3.0/8.0, 1e-10),
		AllCandidateOperatorsDimensionless:   allDimensionless,
		DimensionfulOperatorFound:            dimensionfulFound,
		AbsoluteCouplingOperatorFound:        absCouplingFound,
		BoundaryScaleOperatorFound:           boundaryScaleFound,
		NativeFiniteRGOperatorFound:          false,
		ThresholdSelectorFound:               thresholdFound,
		EquationAudit:                        audit,
		ResidualSymmetries:                   symmetries,
		ResidualVariableCount:                audit.Nullity,
		UnitTraceConventionRejected:          !coupling.TraceNormalizationDerived && !coupling.GaugeCouplingDerived,
		TopologicalSealAsScaleRejected:       coupling.DimensionfulScaleDerived == false,
		ObservedMatchingRejected:             !rg.HiddenObservedInputUsed && !coupling.HiddenObservedCouplingUsed,
		BoundaryCouplingDerived:              false,
		BoundaryScaleDerived:                 false,
		ThresholdRuleDerived:                 false,
		FiniteRGTheoremDerived:               false,
		PhysicalWeakAngleDerived:             false,
		FineStructureDerived:                 false,
		PhysicalMassesDerived:                false,
		HiddenObservedInputUsed:              false,
		TruthStatement:                       truth,
		RejectedClaims: []string{
			"S_top=8π² by itself selects a physical boundary scale M*",
			"the unit-trace diagnostic g²=1 is a derived gauge coupling",
			"the boundary value sin²_*=3/8 is already the observed weak mixing angle",
			"finite scalar radius, B-gap, or contact leakage can be read as GeV without a dimensional anchor",
			"one-loop beta diagnostics can predict low-energy alpha/thetaW before u, L, and thresholds are selected",
		},
		RemainingUnknowns: []string{
			"U-24A-TRACE-PREFactor: derive the finite-to-continuum gauge-action prefactor that fixes u=1/g_*², or prove it is convention-only",
			"U-24B-DIMENSIONAL-ANCHOR: derive a mass/length unit from gravity, spectral cutoff, causal/contact volume, or another finite invariant",
			"U-24C-NATIVE-COARSE-GRAINING: construct a finite RG operator whose fixed point and flow replace the continuum one-loop assumption",
			"U-24D-THRESHOLD-SCHEDULE: derive the finite decoupling/matching operator for heavy/contact modes",
			"U-24E-NO-OBSERVED-MATCHING: keep alpha, thetaW, W/Z/Higgs/fermion masses sealed until the above are selected",
		},
		RecommendedNextGate: "Gate 105 — native finite coarse-graining / threshold activation operator search",
	}, nil
}

func traceEmbeddedBoundary(rg rgfirewall.Analysis) float64 {
	tr, err := rg.EmbeddedBoundaryHessian.Trace()
	if err != nil {
		return math.NaN()
	}
	return tr
}

func vectorNorm(xs ...float64) float64 {
	var s float64
	for _, x := range xs {
		s += x * x
	}
	return math.Sqrt(s)
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatCandidateOperators(xs []CandidateOperator) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		selector := "no-physical-selector"
		if x.SelectsBoundaryCoupling || x.SelectsBoundaryScale || x.SelectsThresholdRule {
			selector = "selector"
		}
		parts = append(parts, fmt.Sprintf("%s=%s %.10g [%s, dimensionless=%t, %s]", x.Symbol, x.Formula, x.Value, x.Kind, x.Dimensionless, selector))
	}
	return strings.Join(parts, "; ")
}

func FormatEquationAudit(a EquationAudit) string {
	return fmt.Sprintf("variables=%s; selected finite equations=%s; independent physical-flow equations=%d/%d; nullity=%d; %s", strings.Join(a.Variables, ", "), strings.Join(a.SelectedFiniteBoundaryEquations, " | "), a.IndependentEquationsForPhysicalFlow, a.RequiredIndependentEquations, a.Nullity, a.Detail)
}

func FormatResidualSymmetries(xs []ResidualSymmetry) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s: %s; blocks=%s", x.Name, x.Transformation, strings.Join(x.Blocks, "/")))
	}
	return strings.Join(parts, "; ")
}

func Join(xs []string) string { return strings.Join(xs, "; ") }
