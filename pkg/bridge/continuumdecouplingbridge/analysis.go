// Package continuumdecouplingbridge implements Gate 180: continuum
// decoupling bridge axiom inventory / finite heat-kernel matching preflight.
//
// Gates 177-179 showed that non-universal threshold deformations are not
// available from the current finite threshold candidates. Gate 180 does not try
// to compute another fitted Δb vector. Instead it audits the exact axioms that
// would be required to promote existing finite spectra into continuum
// heat-kernel/decoupling data.
//
// The gate is intentionally strict. A finite spectrum or zeta ledger is not yet
// a heat-kernel contribution unless the engine also supplies a local continuum
// carrier, bundle/trace normalization, a Laplace-type operator, mass dimensions,
// activation predicates, and a matching law.
package continuumdecouplingbridge

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdorigindichotomy"
)

type AxiomCategory string

const (
	GeometricCarrierAxiom AxiomCategory = "geometric-carrier"
	BundleTraceAxiom      AxiomCategory = "bundle-trace-normalization"
	OperatorAxiom         AxiomCategory = "operator-heat-kernel"
	MassMatchingAxiom     AxiomCategory = "mass-threshold-matching"
	RepresentationAxiom   AxiomCategory = "representation-beta-index"
	CompatibilityAxiom    AxiomCategory = "compatibility-firewall"
)

type BridgeAxiom struct {
	Name                      string
	Category                  AxiomCategory
	RequiredForHeatKernel     bool
	RequiredForThresholds     bool
	FinitePredataAvailable    bool
	CanonicalForBridge        bool
	MissingReason             string
	BlocksHeatKernelPromotion bool
	BlocksThresholdPromotion  bool
	CanBeImportedAsConvention bool
	ImportWouldReduceNullity  bool
}

type AnchorHeatKernelAudit struct {
	Name                          string
	Source                        string
	CandidateFromGate179          bool
	ExactFiniteSpectrum           bool
	GaloisOrBranchSafe            bool
	LaplaceTypeOperator           bool
	PositiveEllipticContinuum     bool
	LocalityOrSectionFunctor      bool
	FourDimensionalCarrier        bool
	GaugeBundleAction             bool
	PhysicalMassDimension         bool
	CutoffMomentConvention        bool
	SeeleyDeWittCoefficients      bool
	ThresholdActivationPredicate  bool
	DecouplingMatchingLaw         bool
	RepresentationBetaRow         bool
	CanContributeToHeatKernel     bool
	CanGenerateNonUniversalDeltaB bool
	Verdict                       string
}

type HeatKernelPreflightAudit struct {
	AxiomsAudited                    int
	RequiredHeatKernelAxioms         int
	RequiredThresholdAxioms          int
	FinitePredataAxioms              int
	CanonicalBridgeAxioms            int
	MissingHeatKernelAxioms          int
	MissingThresholdAxioms           int
	AnchorsAudited                   int
	ExactFiniteAnchors               int
	PromotableHeatKernelAnchors      int
	PromotableThresholdAnchors       int
	A0CoefficientDerived             bool
	A2CoefficientDerived             bool
	A4GaugeCoefficientDerived        bool
	FiniteHeatKernelMatchingDerived  bool
	ContinuumDecouplingBridgeDerived bool
	Gate177RepairPromoted            bool
	Verdict                          string
}

type ChernWeilTraceAudit struct {
	TopologicalSealAvailable       bool
	RepresentationTraceRatioClosed bool
	ChernWeilFormDerived           bool
	OrientedFourCycleDerived       bool
	PrincipalBundleDerived         bool
	TraceNormalizationDerived      bool
	InstantonNumberMapDerived      bool
	AbsoluteCouplingPromoted       bool
	Verdict                        string
}

type DecouplingLawAudit struct {
	MassUnitDerived            bool
	ActivationPredicateDerived bool
	HeavyLightSplitDerived     bool
	MatchingScaleDerived       bool
	SchemeConventionDerived    bool
	ThresholdLogLawDerived     bool
	NonUniversalDeltaBDerived  bool
	Verdict                    string
}

type FirewallAudit struct {
	UsesObservedInputForDerivation bool
	ContinuumBridgeDerived         bool
	HeatKernelMatchingDerived      bool
	ThresholdCorrectedBetaDerived  bool
	NonUniversalDeltaBDerived      bool
	AbsoluteCouplingDerived        bool
	PhysicalConstantsDerived       bool
	BoundaryScaleDerived           bool
	StrictNullityBefore            int
	StrictNullityAfter             int
	ConditionalNullityBefore       int
	ConditionalNullityAfter        int
	SealedClaims                   []string
	OpenRequirements               []string
	RecommendedNextGate            string
	Verdict                        string
}

type Analysis struct {
	PreviousGate179 thresholdorigindichotomy.Analysis
	Axioms          []BridgeAxiom
	Anchors         []AnchorHeatKernelAudit
	Preflight       HeatKernelPreflightAudit
	ChernWeilTrace  ChernWeilTraceAudit
	DecouplingLaw   DecouplingLawAudit
	Firewall        FirewallAudit
	TruthStatement  string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := thresholdorigindichotomy.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev thresholdorigindichotomy.Analysis) (Analysis, error) {
	if !prev.Dichotomy.DichotomyCompleteAtCurrentStage || prev.Dichotomy.ThresholdOriginDerived || prev.Dichotomy.Gate177RepairPromoted {
		return Analysis{}, fmt.Errorf("Gate 180 requires Gate 179 to leave threshold origins open and underived")
	}
	if prev.Firewall.StrictNullityAfter != 3 || prev.Firewall.ConditionalNullityAfter != 2 {
		return Analysis{}, fmt.Errorf("Gate 180 requires Gate 179 nullity ledger 3/2")
	}
	if prev.Firewall.UsesObservedInputForDerivation || prev.Firewall.NonUniversalDeltaBDerived || prev.Firewall.ThresholdCorrectedBetaDerived {
		return Analysis{}, fmt.Errorf("Gate 180 refuses non-quarantined threshold input")
	}

	axioms := buildAxioms()
	anchors := buildAnchors(prev)
	preflight := auditPreflight(axioms, anchors)
	chern := ChernWeilTraceAudit{
		TopologicalSealAvailable:       true,
		RepresentationTraceRatioClosed: true,
		ChernWeilFormDerived:           false,
		OrientedFourCycleDerived:       false,
		PrincipalBundleDerived:         false,
		TraceNormalizationDerived:      false,
		InstantonNumberMapDerived:      false,
		AbsoluteCouplingPromoted:       false,
		Verdict:                        "topological and representation-trace predata exist, but no Chern-Weil four-cycle/bundle/trace map promotes them into continuum normalization",
	}
	decoupling := DecouplingLawAudit{
		MassUnitDerived:            false,
		ActivationPredicateDerived: false,
		HeavyLightSplitDerived:     false,
		MatchingScaleDerived:       false,
		SchemeConventionDerived:    false,
		ThresholdLogLawDerived:     false,
		NonUniversalDeltaBDerived:  false,
		Verdict:                    "no finite mass unit, activation predicate, heavy/light split, matching scale, or threshold logarithm is derived",
	}
	firewall := FirewallAudit{
		UsesObservedInputForDerivation: false,
		ContinuumBridgeDerived:         false,
		HeatKernelMatchingDerived:      false,
		ThresholdCorrectedBetaDerived:  false,
		NonUniversalDeltaBDerived:      false,
		AbsoluteCouplingDerived:        false,
		PhysicalConstantsDerived:       false,
		BoundaryScaleDerived:           false,
		StrictNullityBefore:            prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:             prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:       prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:        prev.Firewall.ConditionalNullityAfter,
		SealedClaims: []string{
			"finite zeta data alone are heat-kernel coefficients",
			"contact spectra have a physical mass dimension without a scale map",
			"Gate 177 threshold fits are derived decoupling laws",
			"the topological action seal fixes continuum Chern-Weil normalization without a four-cycle and bundle map",
		},
		OpenRequirements: []string{
			"oriented four-dimensional continuum carrier or finite four-cycle surrogate",
			"principal gauge bundle / connection map",
			"Chern-Weil and trace normalization bridge",
			"Laplace-type operator with heat-kernel coefficient extraction",
			"finite mass unit, activation predicate, and decoupling/matching law",
		},
		RecommendedNextGate: "Gate 181 — finite oriented four-cycle / Chern-Weil carrier construction search",
		Verdict:             "Gate 180 derives no threshold corrections. It proves the continuum decoupling bridge is axiomatically under-specified: exact finite spectra exist, but the heat-kernel and matching machinery needed to turn them into Δb_i is missing.",
	}
	truth := "Gate 180 is a preflight theorem. It inventories the axioms required to turn existing finite spectra into continuum heat-kernel or decoupling data and finds that no existing anchor is promotable. The failure is not spectral; exact finite spectra and zeta ledgers exist. The missing layer is geometric and analytic: oriented carrier, bundle/trace/Chern-Weil map, Laplace-type operator, mass dimension, activation predicate, and matching law. No nullity reduction is allowed."
	return Analysis{PreviousGate179: prev, Axioms: axioms, Anchors: anchors, Preflight: preflight, ChernWeilTrace: chern, DecouplingLaw: decoupling, Firewall: firewall, TruthStatement: truth}, nil
}

func buildAxioms() []BridgeAxiom {
	return []BridgeAxiom{
		{Name: "oriented four-dimensional carrier / finite four-cycle", Category: GeometricCarrierAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, MissingReason: "no continuum base or finite orientability cycle is derived for Chern-Weil integration", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "principal SU(3)×SU(2)×U(1) bundle / connection map", Category: BundleTraceAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, MissingReason: "finite gauge generators close, but no principal bundle or field-strength integration carrier is constructed", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "Chern-Weil normalization and instanton-number map", Category: BundleTraceAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, FinitePredataAvailable: true, MissingReason: "S_top=8π² is available, but its map to ∫tr(F∧F) is not", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "continuum trace / kinetic normalization convention", Category: BundleTraceAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, FinitePredataAvailable: true, MissingReason: "representation-trace ratios are closed, but absolute trace normalization is not promoted", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "local field-map from finite anchors to continuum sections", Category: GeometricCarrierAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, MissingReason: "contact/B-sector/scalar anchors are not mapped to local field bundles", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "Laplace-type positive elliptic operator", Category: OperatorAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, MissingReason: "finite spectra are not yet spectra of a continuum Laplace-type operator", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "heat-kernel cutoff/test function and moment convention", Category: OperatorAxiom, RequiredForHeatKernel: true, RequiredForThresholds: false, MissingReason: "no canonical cutoff function or spectral-action moments are selected for threshold matching", BlocksHeatKernelPromotion: true},
		{Name: "Seeley-DeWitt coefficient extraction a0/a2/a4", Category: OperatorAxiom, RequiredForHeatKernel: true, RequiredForThresholds: true, MissingReason: "finite zeta values are known, but no local coefficient expansion is derived", BlocksHeatKernelPromotion: true, BlocksThresholdPromotion: true},
		{Name: "physical mass dimension / scale map", Category: MassMatchingAxiom, RequiredForHeatKernel: false, RequiredForThresholds: true, MissingReason: "finite eigenvalues are dimensionless without a mass unit or boundary scale", BlocksThresholdPromotion: true},
		{Name: "heavy-light activation predicate", Category: MassMatchingAxiom, RequiredForHeatKernel: false, RequiredForThresholds: true, MissingReason: "no finite rule decides which modes are active, decoupled, constrained, or baseline", BlocksThresholdPromotion: true},
		{Name: "decoupling and matching logarithm law", Category: MassMatchingAxiom, RequiredForHeatKernel: false, RequiredForThresholds: true, MissingReason: "no finite theorem supplies threshold logs or matching-scale dependence", BlocksThresholdPromotion: true},
		{Name: "gauge representation rows for threshold modes", Category: RepresentationAxiom, RequiredForHeatKernel: false, RequiredForThresholds: true, FinitePredataAvailable: true, MissingReason: "baseline representation rows exist, but open finite threshold anchors lack complete rows", BlocksThresholdPromotion: true},
		{Name: "anomaly/vacuum compatibility for any heavy sector", Category: CompatibilityAxiom, RequiredForHeatKernel: false, RequiredForThresholds: true, MissingReason: "new-sector threshold origins would need compatibility checks not yet supplied", BlocksThresholdPromotion: true},
	}
}

func buildAnchors(prev thresholdorigindichotomy.Analysis) []AnchorHeatKernelAudit {
	anchors := make([]AnchorHeatKernelAudit, 0, len(prev.Continuum.CandidateExistingAnchors)+2)
	for _, name := range prev.Continuum.CandidateExistingAnchors {
		anchors = append(anchors, AnchorHeatKernelAudit{
			Name:                 name,
			Source:               "Gate 178 / Gate 179 existing finite anchor",
			CandidateFromGate179: true,
			ExactFiniteSpectrum:  true,
			GaloisOrBranchSafe:   name != "seven contact partial-overlap modes",
			Verdict:              "exact finite spectral anchor, but no local heat-kernel carrier, mass unit, activation predicate, or matching law",
		})
	}
	anchors = append(anchors,
		AnchorHeatKernelAudit{Name: "Gate 167 representation-trace gauge ratio", Source: "Gate 167", ExactFiniteSpectrum: false, GaloisOrBranchSafe: true, GaugeBundleAction: true, RepresentationBetaRow: true, Verdict: "closes gauge ratios, but is not a finite threshold spectrum or heat-kernel coefficient extraction"},
		AnchorHeatKernelAudit{Name: "Gate 174 topological action seal", Source: "Gate 174", ExactFiniteSpectrum: false, GaloisOrBranchSafe: true, Verdict: "topological predata, but no Chern-Weil carrier/trace bridge or decoupling law"},
	)
	for i := range anchors {
		anchors[i].LaplaceTypeOperator = false
		anchors[i].PositiveEllipticContinuum = false
		anchors[i].LocalityOrSectionFunctor = false
		anchors[i].FourDimensionalCarrier = false
		anchors[i].PhysicalMassDimension = false
		anchors[i].CutoffMomentConvention = false
		anchors[i].SeeleyDeWittCoefficients = false
		anchors[i].ThresholdActivationPredicate = false
		anchors[i].DecouplingMatchingLaw = false
		anchors[i].CanContributeToHeatKernel = false
		anchors[i].CanGenerateNonUniversalDeltaB = false
	}
	return anchors
}

func auditPreflight(axioms []BridgeAxiom, anchors []AnchorHeatKernelAudit) HeatKernelPreflightAudit {
	p := HeatKernelPreflightAudit{AxiomsAudited: len(axioms), AnchorsAudited: len(anchors)}
	for _, a := range axioms {
		if a.RequiredForHeatKernel {
			p.RequiredHeatKernelAxioms++
			if !a.CanonicalForBridge {
				p.MissingHeatKernelAxioms++
			}
		}
		if a.RequiredForThresholds {
			p.RequiredThresholdAxioms++
			if !a.CanonicalForBridge {
				p.MissingThresholdAxioms++
			}
		}
		if a.FinitePredataAvailable {
			p.FinitePredataAxioms++
		}
		if a.CanonicalForBridge {
			p.CanonicalBridgeAxioms++
		}
	}
	for _, h := range anchors {
		if h.ExactFiniteSpectrum {
			p.ExactFiniteAnchors++
		}
		if h.CanContributeToHeatKernel {
			p.PromotableHeatKernelAnchors++
		}
		if h.CanGenerateNonUniversalDeltaB {
			p.PromotableThresholdAnchors++
		}
	}
	p.A0CoefficientDerived = false
	p.A2CoefficientDerived = false
	p.A4GaugeCoefficientDerived = false
	p.FiniteHeatKernelMatchingDerived = false
	p.ContinuumDecouplingBridgeDerived = false
	p.Gate177RepairPromoted = false
	p.Verdict = "all current finite anchors fail the heat-kernel/threshold promotion test; exact spectra are predata, not Seeley-DeWitt or decoupling data"
	return p
}

func FormatAxiom(a BridgeAxiom) string {
	return fmt.Sprintf("%s[%s](heat=%t,threshold=%t,predata=%t,canonical=%t,blocksHeat=%t,blocksThreshold=%t,importConv=%t): %s", a.Name, a.Category, a.RequiredForHeatKernel, a.RequiredForThresholds, a.FinitePredataAvailable, a.CanonicalForBridge, a.BlocksHeatKernelPromotion, a.BlocksThresholdPromotion, a.CanBeImportedAsConvention, a.MissingReason)
}

func FormatAxioms(xs []BridgeAxiom) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatAxiom(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatAnchor(a AnchorHeatKernelAudit) string {
	return fmt.Sprintf("%s(src=%s,gate179=%t,spec=%t,laplace=%t,local=%t,4d=%t,bundle=%t,mass=%t,aSDW=%t,activation=%t,decouple=%t,rep=%t,heat=%t,Δb=%t): %s", a.Name, a.Source, a.CandidateFromGate179, a.ExactFiniteSpectrum, a.LaplaceTypeOperator, a.LocalityOrSectionFunctor, a.FourDimensionalCarrier, a.GaugeBundleAction, a.PhysicalMassDimension, a.SeeleyDeWittCoefficients, a.ThresholdActivationPredicate, a.DecouplingMatchingLaw, a.RepresentationBetaRow, a.CanContributeToHeatKernel, a.CanGenerateNonUniversalDeltaB, a.Verdict)
}

func FormatAnchors(xs []AnchorHeatKernelAudit) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatAnchor(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatPreflight(p HeatKernelPreflightAudit) string {
	return fmt.Sprintf("axioms=%d heatReq=%d thresholdReq=%d predata=%d canonical=%d missingHeat=%d missingThreshold=%d anchors=%d exactAnchors=%d heatAnchors=%d thresholdAnchors=%d a0=%t a2=%t a4Gauge=%t heatKernel=%t decoupling=%t promotes=%t: %s", p.AxiomsAudited, p.RequiredHeatKernelAxioms, p.RequiredThresholdAxioms, p.FinitePredataAxioms, p.CanonicalBridgeAxioms, p.MissingHeatKernelAxioms, p.MissingThresholdAxioms, p.AnchorsAudited, p.ExactFiniteAnchors, p.PromotableHeatKernelAnchors, p.PromotableThresholdAnchors, p.A0CoefficientDerived, p.A2CoefficientDerived, p.A4GaugeCoefficientDerived, p.FiniteHeatKernelMatchingDerived, p.ContinuumDecouplingBridgeDerived, p.Gate177RepairPromoted, p.Verdict)
}

func FormatChernWeilTrace(c ChernWeilTraceAudit) string {
	return fmt.Sprintf("topSeal=%t repRatio=%t cernWeil=%t fourCycle=%t bundle=%t trace=%t instantonMap=%t coupling=%t: %s", c.TopologicalSealAvailable, c.RepresentationTraceRatioClosed, c.ChernWeilFormDerived, c.OrientedFourCycleDerived, c.PrincipalBundleDerived, c.TraceNormalizationDerived, c.InstantonNumberMapDerived, c.AbsoluteCouplingPromoted, c.Verdict)
}

func FormatDecouplingLaw(d DecouplingLawAudit) string {
	return fmt.Sprintf("mass=%t activation=%t heavyLight=%t scale=%t scheme=%t logLaw=%t Δb=%t: %s", d.MassUnitDerived, d.ActivationPredicateDerived, d.HeavyLightSplitDerived, d.MatchingScaleDerived, d.SchemeConventionDerived, d.ThresholdLogLawDerived, d.NonUniversalDeltaBDerived, d.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("obs=%t continuum=%t heatKernel=%t beta=%t Δb=%t coupling=%t constants=%t scale=%t strict=%d->%d conditional=%d->%d sealed=[%s] open=[%s] next=%s: %s", f.UsesObservedInputForDerivation, f.ContinuumBridgeDerived, f.HeatKernelMatchingDerived, f.ThresholdCorrectedBetaDerived, f.NonUniversalDeltaBDerived, f.AbsoluteCouplingDerived, f.PhysicalConstantsDerived, f.BoundaryScaleDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalNullityBefore, f.ConditionalNullityAfter, strings.Join(f.SealedClaims, "; "), strings.Join(f.OpenRequirements, "; "), f.RecommendedNextGate, f.Verdict)
}
