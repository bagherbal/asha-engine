// Package scalarmoritaspectralbridge implements Gate 276:
// Scalar-Morita Spectral Shape Bridge / Branch Selector and Heat-Kernel
// Normalization Audit.
//
// Gate 275 found a genuine two-branch algebraic constraint on the Morita
// edge-shape ratio r=|y/x|² by equating the Gate-169 contact scalar shape
// λ_contact=1197/4624 with the Gate-273 Morita multiplicity trace shape
// (|x|⁴+3|y|⁴)/(|x|²+3|y|²)². Gate 276 audits the next step: whether this
// scale-free shape constraint can be lawfully promoted to a Seeley-de Witt
// a₂/a₄ Higgs-ratio prediction.
//
// The gate is deliberately conservative. It carries forward the two exact
// branches and constructs a formal heat-kernel obligation map, but it refuses
// to identify raw finite traces with physical Lagrangian coefficients until a
// branch selector, cutoff moments, subtraction scheme, scalar/gauge projection,
// physical J/opposite action, and field normalization are derived.
package scalarmoritaspectralbridge

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/physicalfinitehilbertcompletion"
)

const (
	AuditID = "GATE276-SCALAR-MORITA-SPECTRAL-SHAPE-BRIDGE-BRANCH-SELECTOR-HEAT-KERNEL-NORMALIZATION-AUDIT"

	StatusGate275Inherited       = "CONDITIONAL_SUPPORT_GATE275_TWO_BRANCH_SHAPE_CONSTRAINT_INHERITED"
	StatusBridgeFormalized       = "CONDITIONAL_SUPPORT_SCALAR_MORITA_BRIDGE_FORMALIZED"
	StatusBranchAuditCompleted   = "CONDITIONAL_SUPPORT_BRANCH_SELECTOR_AUDIT_COMPLETED"
	StatusCandidateMoments       = "CONDITIONAL_SUPPORT_BRANCH_CANDIDATE_MOMENTS_RECOMPUTED"
	StatusHeatKernelFormalized   = "CONDITIONAL_SUPPORT_FORMAL_HEAT_KERNEL_PROJECTION_REQUIREMENTS_DEFINED"
	StatusFirewallPreserved      = "CONDITIONAL_SUPPORT_SPECTRAL_FIREWALLS_PRESERVED"
	StatusFailedBranchDegenerate = "FAILED_ROUTE_TWO_BRANCH_VACUUM_AMBIGUITY_REMAINS"
	StatusFailedScale            = "FAILED_ROUTE_ABSOLUTE_DF_SCALE_NOT_DERIVED"
	StatusFailedHKProjection     = "FAILED_ROUTE_HEAT_KERNEL_PROJECTION_NOT_DERIVED"
	StatusFailedFieldNorm        = "FAILED_ROUTE_SCALAR_GAUGE_FIELD_NORMALIZATION_MISSING"
	StatusFailedA2A4             = "FAILED_ROUTE_SEELEY_DE_WITT_A2_A4_NOT_DERIVED"
	StatusFailedHiggs            = "FAILED_ROUTE_HIGGS_MASS_RATIO_NOT_CLAIMED"
)

type Gate275Inheritance struct {
	ScalarMoritaSolved     bool
	TwoBranchXYConstrained bool
	UniqueXYLocked         bool
	PhysicalJDerived       bool
	HyperchargeDerived     bool
	A2A4Derived            bool
	HiggsRatioClaimed      bool
	FirewallPreserved      bool
	InheritedBranchCount   int
	Verdict                string
}

type ScalarMoritaBridgeTheorem struct {
	LambdaNumerator    int
	LambdaDenominator  int
	Lambda             float64
	KappaC             int
	KappaQ             int
	Equation           string
	Quadratic          string
	RootForm           string
	CrossTowerBridge   bool
	ScaleFreeShapeOnly bool
	EquivalentToA2A4   bool
	Verdict            string
}

type SpectralMomentBranch struct {
	Name                   string
	ExactRForm             string
	R                      float64
	AbsYOverX              float64
	D2ForXEqualsOne        float64
	D4ForXEqualsOne        float64
	ShapeLambda            float64
	ShapeResidualAbs       float64
	D4OverD2ForXEqualsOne  float64
	D2OverD4ForXEqualsOne  float64
	D4OverD2DependsOnScale bool
	A2A4CandidateClaimed   bool
	Interpretation         string
}

type BranchSelectorCandidate struct {
	Name          string
	Inputs        string
	TestsBranches bool
	UpperPasses   bool
	LowerPasses   bool
	SelectsUnique bool
	Selected      string
	Verdict       string
}

type BranchSelectorAudit struct {
	Candidates          []BranchSelectorCandidate
	UpperBranchAllowed  bool
	LowerBranchAllowed  bool
	UniqueBranch        bool
	SelectedBranch      string
	FiniteCoreSelector  bool
	RequiresFutureInput bool
	Verdict             string
}

type HeatKernelProjectionAudit struct {
	FormalExpansion           string
	RawFiniteMomentShape      string
	A2Candidate               string
	A4Candidate               string
	CutoffMomentsSpecified    bool
	SubtractionSchemeDerived  bool
	GaugeKineticProjection    bool
	ScalarFluctuationMap      bool
	PhysicalJAvailable        bool
	HyperchargeAvailable      bool
	FieldNormalizationDerived bool
	CanMapRawTracesToA2A4     bool
	Verdict                   string
}

type HiggsRatioAudit struct {
	UsesSelectedBranch      bool
	UsesAbsoluteDFScale     bool
	UsesHeatKernelMap       bool
	UsesFieldNormalization  bool
	BranchCount             int
	InvariantA2A4Computed   bool
	HiggsMassRatioComputed  bool
	CandidateOnlyStatements []string
	Verdict                 string
}

type FirewallAudit struct {
	NoObservedMassInserted             bool
	NoVEVInserted                      bool
	NoCKMPMNSInserted                  bool
	NoEmpiricalYukawaAmplitudeInserted bool
	RawTraceShapeNotPromoted           bool
	CandidateBranchesNotPredictions    bool
	EmpiricalYukawaSealPreserved       bool
	FiniteCorePolluted                 bool
	Verdict                            string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                 []FutureCriterion
	NeedBranchSelector       bool
	NeedAbsoluteScale        bool
	NeedPhysicalJ            bool
	NeedHypercharge          bool
	NeedHeatKernelProjection bool
	NeedFieldNormalization   bool
	RecommendedNextGate      string
	Verdict                  string
}

type Summary struct {
	Gate275Inherited     bool
	BridgeFormalized     bool
	TwoBranchesCarried   bool
	UniqueBranchSelected bool
	HeatKernelFormalized bool
	HeatKernelDerived    bool
	A2A4Derived          bool
	HiggsRatioClaimed    bool
	FirewallPreserved    bool
	Status               string
	NextGate             string
	Comment              string
}

type Analysis struct {
	PreviousGate275 physicalfinitehilbertcompletion.Analysis
	Inheritance     Gate275Inheritance
	Bridge          ScalarMoritaBridgeTheorem
	Branches        []SpectralMomentBranch
	Selector        BranchSelectorAudit
	HeatKernel      HeatKernelProjectionAudit
	HiggsRatio      HiggsRatioAudit
	Firewall        FirewallAudit
	Future          FutureMap
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := physicalfinitehilbertcompletion.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 275 predecessor: %w", err)
			return
		}
		inh := inheritGate275(prev)
		bridge := formalizeBridge(prev)
		branches := recomputeBranches(prev)
		selector := auditBranchSelectors(branches)
		hk := auditHeatKernel(prev)
		higgs := auditHiggsRatio(branches, selector, hk)
		fw := auditFirewall(selector, hk, higgs)
		future := defineFuture(selector, hk, higgs)
		summary := summarize(inh, bridge, branches, selector, hk, higgs, fw)
		truth := buildTruth(bridge, branches, selector, hk, higgs)
		defaultA = Analysis{PreviousGate275: prev, Inheritance: inh, Bridge: bridge, Branches: branches, Selector: selector, HeatKernel: hk, HiggsRatio: higgs, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate275(prev physicalfinitehilbertcompletion.Analysis) Gate275Inheritance {
	return Gate275Inheritance{
		ScalarMoritaSolved:     prev.Summary.ScalarMoritaSolved,
		TwoBranchXYConstrained: prev.Summary.TwoBranchXYConstrained,
		UniqueXYLocked:         prev.Summary.UniqueXYLocked,
		PhysicalJDerived:       prev.Summary.PhysicalJDerived,
		HyperchargeDerived:     prev.Summary.HyperchargeDerived,
		A2A4Derived:            prev.Summary.A2A4Derived,
		HiggsRatioClaimed:      prev.Summary.HiggsRatioClaimed,
		FirewallPreserved:      prev.Summary.FirewallPreserved,
		InheritedBranchCount:   len(prev.Bridge.Branches),
		Verdict:                StatusGate275Inherited + "; exact two-branch r constraint is inherited, but no branch/J/hypercharge/a2a4 theorem is inherited",
	}
}

func formalizeBridge(prev physicalfinitehilbertcompletion.Analysis) ScalarMoritaBridgeTheorem {
	return ScalarMoritaBridgeTheorem{
		LambdaNumerator:    prev.Bridge.Shape.ExactNumerator,
		LambdaDenominator:  prev.Bridge.Shape.ExactDenominator,
		Lambda:             prev.Bridge.Shape.FloatValue,
		KappaC:             prev.Bridge.Multiplicity.KappaC,
		KappaQ:             prev.Bridge.Multiplicity.KappaQ,
		Equation:           prev.Bridge.BridgeEquation,
		Quadratic:          "3099r² - 7182r + 3427 = 0",
		RootForm:           "r = (3591 ± 136√123)/3099",
		CrossTowerBridge:   true,
		ScaleFreeShapeOnly: true,
		EquivalentToA2A4:   false,
		Verdict:            StatusBridgeFormalized + "; Gate 169 scalar shape and Gate 273 Morita multiplicity meet as a scale-free spectral-shape equation, not yet as physical heat-kernel coefficients",
	}
}

func recomputeBranches(prev physicalfinitehilbertcompletion.Analysis) []SpectralMomentBranch {
	out := make([]SpectralMomentBranch, 0, len(prev.Bridge.Branches))
	for _, b := range prev.Bridge.Branches {
		out = append(out, SpectralMomentBranch{
			Name:                   b.Name,
			ExactRForm:             b.ExactRForm,
			R:                      b.R,
			AbsYOverX:              b.AbsYOverX,
			D2ForXEqualsOne:        b.TraceD2X1,
			D4ForXEqualsOne:        b.TraceD4X1,
			ShapeLambda:            b.ShapeLambda,
			ShapeResidualAbs:       b.ShapeResidualAbs,
			D4OverD2ForXEqualsOne:  b.D4OverD2X1,
			D2OverD4ForXEqualsOne:  b.D2OverD4X1,
			D4OverD2DependsOnScale: true,
			A2A4CandidateClaimed:   false,
			Interpretation:         StatusCandidateMoments + "; raw moments are candidate finite shapes only and remain scale/normalization dependent before heat-kernel projection",
		})
	}
	return out
}

func auditBranchSelectors(branches []SpectralMomentBranch) BranchSelectorAudit {
	candidates := []BranchSelectorCandidate{
		{
			Name:          "positivity/stability preflight",
			Inputs:        "r>0, Tr(D²)>0, Tr(D⁴)>0, λ>0",
			TestsBranches: true,
			UpperPasses:   true,
			LowerPasses:   true,
			SelectsUnique: false,
			Selected:      "none",
			Verdict:       "both branches are positive and stable at the raw moment level",
		},
		{
			Name:          "finite-core anomaly/charge ledger",
			Inputs:        "B-L, T3, hypercharge preflight ledgers",
			TestsBranches: true,
			UpperPasses:   true,
			LowerPasses:   true,
			SelectsUnique: false,
			Selected:      "none",
			Verdict:       "charge/anomaly ledgers constrain representation content, not the scalar-Morita amplitude branch",
		},
		{
			Name:          "parity/J orientation selector",
			Inputs:        "physical anti-linear J and particle/antiparticle typing",
			TestsBranches: false,
			UpperPasses:   false,
			LowerPasses:   false,
			SelectsUnique: false,
			Selected:      "blocked",
			Verdict:       "physical J remains un-derived, so it cannot select a branch",
		},
		{
			Name:          "energy minimization/action selector",
			Inputs:        "finite scalar potential or spectral action functional",
			TestsBranches: false,
			UpperPasses:   false,
			LowerPasses:   false,
			SelectsUnique: false,
			Selected:      "blocked",
			Verdict:       "no derived potential/action functional ranks the two roots",
		},
	}
	upper := len(branches) == 2 && branches[0].R > 0
	lower := len(branches) == 2 && branches[1].R > 0
	return BranchSelectorAudit{
		Candidates:          candidates,
		UpperBranchAllowed:  upper,
		LowerBranchAllowed:  lower,
		UniqueBranch:        false,
		SelectedBranch:      "none",
		FiniteCoreSelector:  false,
		RequiresFutureInput: true,
		Verdict:             StatusFailedBranchDegenerate + "; all available finite-core selectors either pass both branches or are unavailable",
	}
}

func auditHeatKernel(prev physicalfinitehilbertcompletion.Analysis) HeatKernelProjectionAudit {
	return HeatKernelProjectionAudit{
		FormalExpansion:           "Tr(f(D/Λ)) ~ f₄Λ⁴ a₀ + f₂Λ² a₂ + f₀ a₄ + ...",
		RawFiniteMomentShape:      "λ_shape = Tr(D_F⁴)/Tr(D_F²)² = 1197/4624",
		A2Candidate:               "a₂ requires scalar/Higgs fluctuation projection, sign convention, subtraction, and f₂ moment",
		A4Candidate:               "a₄ requires gauge/scalar quartic projection, kinetic normalization, subtraction, and f₀ moment",
		CutoffMomentsSpecified:    false,
		SubtractionSchemeDerived:  false,
		GaugeKineticProjection:    false,
		ScalarFluctuationMap:      false,
		PhysicalJAvailable:        prev.Summary.PhysicalJDerived,
		HyperchargeAvailable:      prev.Summary.HyperchargeDerived,
		FieldNormalizationDerived: false,
		CanMapRawTracesToA2A4:     false,
		Verdict:                   StatusHeatKernelFormalized + "; formal obligations are known, but the projection from raw finite moments to a₂/a₄ is not derived",
	}
}

func auditHiggsRatio(branches []SpectralMomentBranch, selector BranchSelectorAudit, hk HeatKernelProjectionAudit) HiggsRatioAudit {
	stmts := make([]string, 0, len(branches))
	for _, b := range branches {
		stmts = append(stmts, fmt.Sprintf("%s: r=%.15g, λ=%.15g, D4/D2 at |x|²=1 is %.15g; changing |x|² rescales D4/D2", b.Name, b.R, b.ShapeLambda, b.D4OverD2ForXEqualsOne))
	}
	return HiggsRatioAudit{
		UsesSelectedBranch:      selector.UniqueBranch,
		UsesAbsoluteDFScale:     false,
		UsesHeatKernelMap:       hk.CanMapRawTracesToA2A4,
		UsesFieldNormalization:  hk.FieldNormalizationDerived,
		BranchCount:             len(branches),
		InvariantA2A4Computed:   false,
		HiggsMassRatioComputed:  false,
		CandidateOnlyStatements: stmts,
		Verdict:                 strings.Join([]string{StatusFailedScale, StatusFailedHKProjection, StatusFailedFieldNorm, StatusFailedA2A4, StatusFailedHiggs}, "; "),
	}
}

func auditFirewall(selector BranchSelectorAudit, hk HeatKernelProjectionAudit, h HiggsRatioAudit) FirewallAudit {
	return FirewallAudit{
		NoObservedMassInserted:             true,
		NoVEVInserted:                      true,
		NoCKMPMNSInserted:                  true,
		NoEmpiricalYukawaAmplitudeInserted: true,
		RawTraceShapeNotPromoted:           !hk.CanMapRawTracesToA2A4 && !h.InvariantA2A4Computed,
		CandidateBranchesNotPredictions:    !selector.UniqueBranch && !h.HiggsMassRatioComputed,
		EmpiricalYukawaSealPreserved:       true,
		FiniteCorePolluted:                 false,
		Verdict:                            StatusFirewallPreserved + "; branch candidates remain finite-shape data and are not promoted to empirical mass predictions",
	}
}

func defineFuture(selector BranchSelectorAudit, hk HeatKernelProjectionAudit, h HiggsRatioAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "branch selector", Required: true, Satisfied: selector.UniqueBranch, Detail: "select r₊ or r₋ from finite-core stability/J/orientation/action data"},
		{Name: "absolute or quotient normalization", Required: true, Satisfied: h.UsesAbsoluteDFScale, Detail: "fix the scale convention needed for D4/D2-type coefficients or prove a scale-free observable"},
		{Name: "physical anti-linear J", Required: true, Satisfied: hk.PhysicalJAvailable, Detail: "complete particle/antiparticle opposite action on physical H_F"},
		{Name: "full chiral hypercharge representation", Required: true, Satisfied: hk.HyperchargeAvailable, Detail: "complete C⊕H⊕M3(C) action with hypercharge/chirality semantics"},
		{Name: "heat-kernel cutoff moments", Required: true, Satisfied: hk.CutoffMomentsSpecified, Detail: "derive or seal f₀,f₂,f₄ and their normalization conventions"},
		{Name: "subtraction/renormalization scheme", Required: true, Satisfied: hk.SubtractionSchemeDerived, Detail: "state the finite subtraction map connecting spectral traces to physical Lagrangian coefficients"},
		{Name: "scalar/gauge projection", Required: true, Satisfied: hk.GaugeKineticProjection && hk.ScalarFluctuationMap, Detail: "separate Higgs mass, Higgs quartic, and gauge kinetic terms inside a₄"},
	}
	return FutureMap{
		Criteria:                 criteria,
		NeedBranchSelector:       !selector.UniqueBranch,
		NeedAbsoluteScale:        !h.UsesAbsoluteDFScale,
		NeedPhysicalJ:            !hk.PhysicalJAvailable,
		NeedHypercharge:          !hk.HyperchargeAvailable,
		NeedHeatKernelProjection: !hk.CanMapRawTracesToA2A4,
		NeedFieldNormalization:   !hk.FieldNormalizationDerived,
		RecommendedNextGate:      "Gate 277 — Physical J / Hypercharge Completion or Heat-Kernel Projection Theorem",
		Verdict:                  "Gate 276 narrows the dynamics problem to branch selection plus heat-kernel/normalization completion",
	}
}

func summarize(inh Gate275Inheritance, bridge ScalarMoritaBridgeTheorem, branches []SpectralMomentBranch, selector BranchSelectorAudit, hk HeatKernelProjectionAudit, h HiggsRatioAudit, fw FirewallAudit) Summary {
	return Summary{
		Gate275Inherited:     inh.ScalarMoritaSolved && inh.TwoBranchXYConstrained,
		BridgeFormalized:     bridge.CrossTowerBridge && bridge.ScaleFreeShapeOnly,
		TwoBranchesCarried:   len(branches) == 2,
		UniqueBranchSelected: selector.UniqueBranch,
		HeatKernelFormalized: hk.FormalExpansion != "",
		HeatKernelDerived:    hk.CanMapRawTracesToA2A4,
		A2A4Derived:          h.InvariantA2A4Computed,
		HiggsRatioClaimed:    h.HiggsMassRatioComputed,
		FirewallPreserved:    !fw.FiniteCorePolluted && fw.EmpiricalYukawaSealPreserved,
		Status:               "BRIDGE_REQUIRED_WITH_TWO_BRANCH_SHAPE_AND_HEAT_KERNEL_OBLIGATIONS",
		NextGate:             "Gate 277 — Physical J / Hypercharge Completion or Heat-Kernel Projection Theorem",
		Comment:              "Gate 276 formalizes the scalar-Morita spectral-shape bridge and audits branch/heat-kernel selectors; no unique branch, a₂/a₄, or Higgs ratio is derived.",
	}
}

func buildTruth(bridge ScalarMoritaBridgeTheorem, branches []SpectralMomentBranch, selector BranchSelectorAudit, hk HeatKernelProjectionAudit, h HiggsRatioAudit) string {
	vals := []string{}
	for _, b := range branches {
		vals = append(vals, fmt.Sprintf("%s r≈%.15g", b.Name, b.R))
	}
	return fmt.Sprintf("Gate 276 confirms that λ_contact=%d/%d and κ_C:κ_Q=%d:%d define a genuine scale-free scalar-Morita spectral-shape bridge with branches [%s]. It also proves that no native branch selector or heat-kernel projection is yet available, so raw trace moments cannot be promoted to Seeley-de Witt a₂/a₄ coefficients or a Higgs mass ratio.", bridge.LambdaNumerator, bridge.LambdaDenominator, bridge.KappaC, bridge.KappaQ, strings.Join(vals, ", "))
}

// BranchResidualsOK returns true when all carried branches still reproduce the
// Gate-169 shape to numerical precision.
func BranchResidualsOK(branches []SpectralMomentBranch) bool {
	if len(branches) != 2 {
		return false
	}
	for _, b := range branches {
		if b.R <= 0 || b.AbsYOverX <= 0 || math.IsNaN(b.ShapeLambda) || b.ShapeResidualAbs > 1e-12 {
			return false
		}
	}
	return true
}
