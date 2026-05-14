// Package fourcyclechernweil implements Gate 181: finite oriented four-cycle /
// Chern-Weil carrier construction search.
//
// Gate 180 proved that exact finite spectra are not enough to produce
// heat-kernel coefficients, instanton normalization, or threshold corrections.
// The first missing geometric object is an oriented four-dimensional carrier (or
// finite four-cycle surrogate) on which a Chern-Weil pairing can be defined.
//
// This gate searches the currently derived finite objects for such a carrier. It
// is deliberately conservative: a 4-dimensional vector space, a grade-4 chamber,
// or a topological scalar seal is not sufficient. A candidate must provide a
// boundaryless oriented fundamental class, an integration functional, a gauge
// bundle/connection map, a curvature two-form pairing tr(F∧F), a trace
// normalization, and an integer topological-charge map.
package fourcyclechernweil

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/continuumdecouplingbridge"
)

type CandidateClass string

const (
	VectorSpaceCarrier CandidateClass = "vector-space-carrier"
	ExteriorChamber    CandidateClass = "exterior-grade-chamber"
	IncidenceComplex   CandidateClass = "boolean-incidence-complex"
	ContactCarrier     CandidateClass = "contact-carrier"
	InternalHilbert    CandidateClass = "internal-hilbert-space"
	TopologicalScalar  CandidateClass = "topological-scalar"
	CombinatorialPlane CandidateClass = "combinatorial-incidence-plane"
)

type FourCycleCandidate struct {
	Name                         string
	Source                       string
	Class                        CandidateClass
	Dimension                    int
	ExactFiniteData              bool
	FourDimensionalCarrier       bool
	OrientedBasisAvailable       bool
	BoundaryOperatorAvailable    bool
	BoundarylessCycleVerified    bool
	CanonicalFundamentalClass    bool
	IntegrationFunctionalDerived bool
	GaugeBundleMapDerived        bool
	ConnectionOneFormDerived     bool
	CurvatureTwoFormDerived      bool
	WedgePairingFAndFDerived     bool
	TraceNormalizationDerived    bool
	IntegerTopologicalChargeMap  bool
	HochschildOrientabilityCycle bool
	ChernWeilCarrierComplete     bool
	FailureReason                string
}

type RequirementAudit struct {
	Name                  string
	RequiredForFourCycle  bool
	RequiredForChernWeil  bool
	RequiredForInstanton  bool
	DerivedByAnyCandidate bool
	Blocking              bool
	Comment               string
}

type ConstructionSearchAudit struct {
	CandidatesAudited           int
	ExactFiniteCandidates       int
	FourDimensionalCandidates   int
	BoundarylessCycleCandidates int
	CanonicalFundamentalClasses int
	IntegrationFunctionals      int
	GaugeBundleMaps             int
	CurvaturePairings           int
	TraceNormalizations         int
	IntegerChargeMaps           int
	HochschildCycles            int
	CompleteChernWeilCarriers   int
	ExistingCandidatePromoted   bool
	FiniteFourCycleDerived      bool
	ChernWeilCarrierDerived     bool
	InstantonBridgePromoted     bool
	Verdict                     string
}

type HomologyOrientationAudit struct {
	GradeFourDataAvailable         bool
	DimensionFourVectorSpacesExist bool
	BoundaryOperatorDerived        bool
	NonzeroClosedFourCycleDerived  bool
	CanonicalRepresentativeDerived bool
	OrientationSignDerived         bool
	AutomorphismInvariantSelector  bool
	HochschildCycleRealizesGamma   bool
	Verdict                        string
}

type ChernWeilAudit struct {
	GaugeAlgebraClosed             bool
	RepresentationTraceRatioClosed bool
	TopologicalSealAvailable       bool
	PrincipalBundleDerived         bool
	ConnectionOnFourCarrierDerived bool
	CurvatureTwoFormDerived        bool
	TracePairingDerived            bool
	IntegralOfTrFedgeFDerived      bool
	IntegerInstantonNumberDerived  bool
	ContinuumNormalizationPromoted bool
	Verdict                        string
}

type FirewallAudit struct {
	UsesObservedInputForDerivation bool
	FourCycleDerived               bool
	ChernWeilCarrierDerived        bool
	InstantonTraceBridgeDerived    bool
	AbsoluteCouplingPromoted       bool
	HeatKernelMatchingDerived      bool
	ThresholdCorrectedBetaDerived  bool
	PhysicalConstantsDerived       bool
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
	PreviousGate180 continuumdecouplingbridge.Analysis
	Candidates      []FourCycleCandidate
	Requirements    []RequirementAudit
	Search          ConstructionSearchAudit
	Homology        HomologyOrientationAudit
	ChernWeil       ChernWeilAudit
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
		prev, err := continuumdecouplingbridge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev continuumdecouplingbridge.Analysis) (Analysis, error) {
	if prev.Firewall.StrictNullityAfter != 3 || prev.Firewall.ConditionalNullityAfter != 2 {
		return Analysis{}, fmt.Errorf("Gate 181 requires Gate 180 nullity ledger 3/2")
	}
	if prev.Firewall.UsesObservedInputForDerivation || prev.Firewall.ContinuumBridgeDerived || prev.Firewall.HeatKernelMatchingDerived || prev.Firewall.NonUniversalDeltaBDerived || prev.Firewall.AbsoluteCouplingDerived {
		return Analysis{}, fmt.Errorf("Gate 181 requires Gate 180 bridge firewall to remain closed")
	}
	if prev.Preflight.PromotableHeatKernelAnchors != 0 || prev.Preflight.PromotableThresholdAnchors != 0 || prev.ChernWeilTrace.OrientedFourCycleDerived || prev.ChernWeilTrace.ChernWeilFormDerived {
		return Analysis{}, fmt.Errorf("Gate 181 requires Gate 180 to leave the four-cycle/Chern-Weil carrier missing")
	}

	candidates := buildCandidates()
	requirements := auditRequirements(candidates)
	search := auditSearch(candidates)
	homology := HomologyOrientationAudit{
		GradeFourDataAvailable:         true,
		DimensionFourVectorSpacesExist: true,
		BoundaryOperatorDerived:        true,
		NonzeroClosedFourCycleDerived:  false,
		CanonicalRepresentativeDerived: false,
		OrientationSignDerived:         false,
		AutomorphismInvariantSelector:  false,
		HochschildCycleRealizesGamma:   false,
		Verdict:                        "grade-4 and 4-dimensional predata exist, but no canonical boundaryless oriented fundamental class or Hochschild orientability cycle is derived",
	}
	chern := ChernWeilAudit{
		GaugeAlgebraClosed:             true,
		RepresentationTraceRatioClosed: true,
		TopologicalSealAvailable:       true,
		PrincipalBundleDerived:         false,
		ConnectionOnFourCarrierDerived: false,
		CurvatureTwoFormDerived:        false,
		TracePairingDerived:            false,
		IntegralOfTrFedgeFDerived:      false,
		IntegerInstantonNumberDerived:  false,
		ContinuumNormalizationPromoted: false,
		Verdict:                        "gauge algebra, representation trace, and S_top are available, but no four-carrier bundle/connection/Chern-Weil integral promotes them to instanton normalization",
	}
	firewall := FirewallAudit{
		UsesObservedInputForDerivation: false,
		FourCycleDerived:               search.FiniteFourCycleDerived,
		ChernWeilCarrierDerived:        search.ChernWeilCarrierDerived,
		InstantonTraceBridgeDerived:    false,
		AbsoluteCouplingPromoted:       false,
		HeatKernelMatchingDerived:      false,
		ThresholdCorrectedBetaDerived:  false,
		PhysicalConstantsDerived:       false,
		StrictNullityBefore:            prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:             prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:       prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:        prev.Firewall.ConditionalNullityAfter,
		SealedClaims: []string{
			"Λ⁴R⁸ by itself is an oriented spacetime four-cycle",
			"a 4-dimensional internal vector space is a Chern-Weil integration carrier",
			"the finite topological scalar S_top supplies an instanton number map without a bundle and trace convention",
			"the Fock spectral triple supplies a Hochschild four-cycle for continuum orientation",
		},
		OpenRequirements: []string{
			"canonical nonzero boundaryless four-cycle or oriented four-dimensional base carrier",
			"integration functional / fundamental class",
			"principal gauge bundle and connection map on that carrier",
			"curvature two-form and tr(F∧F) pairing",
			"integer topological-charge and trace-normalization bridge",
		},
		RecommendedNextGate: "Gate 182 — finite local field/bundle map construction search",
		Verdict:             "no currently derived object is a complete finite oriented four-cycle or Chern-Weil carrier; the instanton/heat-kernel bridge remains geometric rather than spectral",
	}
	truth := "Gate 181 searches all currently available finite carriers for the first missing continuum bridge object: an oriented four-cycle / Chern-Weil carrier. Several objects contain suggestive predata: Λ⁴R⁸ has grade-four basis elements, the scalar/Lorentzian sectors have 4-dimensional vector spaces, the Fock spectral triple has J and γ, and S_top=8π² is a topological scalar. None supplies the complete chain: boundaryless oriented fundamental class, integration functional, gauge bundle/connection, curvature two-form pairing tr(F∧F), trace normalization, and integer instanton map. Therefore the finite-to-continuum normalization bridge is not promoted and nullity is unchanged."
	return Analysis{PreviousGate180: prev, Candidates: candidates, Requirements: requirements, Search: search, Homology: homology, ChernWeil: chern, Firewall: firewall, TruthStatement: truth}, nil
}

func buildCandidates() []FourCycleCandidate {
	return []FourCycleCandidate{
		{Name: "Λ⁴R⁸ middle exterior chamber", Source: "Gates 1/149", Class: ExteriorChamber, Dimension: 70, ExactFiniteData: true, OrientedBasisAvailable: true, FailureReason: "grade-four forms are finite algebra carriers, not a selected oriented four-dimensional integration cycle"},
		{Name: "Boolean Λ³→Λ⁴ incidence complex", Source: "Gate 3", Class: IncidenceComplex, Dimension: 70, ExactFiniteData: true, BoundaryOperatorAvailable: true, FailureReason: "incidence/coboundary structure exists, but no nonzero canonical closed four-chain or fundamental class is selected"},
		{Name: "Lorentzian H_base ≅ R^{1,3}", Source: "R-involution / symmetry-breaking layer", Class: VectorSpaceCarrier, Dimension: 4, ExactFiniteData: true, FourDimensionalCarrier: true, OrientedBasisAvailable: true, FailureReason: "4D signature data exist, but no compact/boundaryless cycle, integration functional, or gauge bundle over it is constructed"},
		{Name: "active scalar/Higgs 4-space H_Φ", Source: "Gates 11/20/37", Class: VectorSpaceCarrier, Dimension: 4, ExactFiniteData: true, FourDimensionalCarrier: true, OrientedBasisAvailable: true, FailureReason: "internal scalar carrier, not a spacetime four-cycle or Chern-Weil base"},
		{Name: "contact vacuum K₇", Source: "Gate 5", Class: ContactCarrier, Dimension: 7, ExactFiniteData: true, FailureReason: "7D contact carrier admits many 4-subsets, but no automorphism-invariant oriented 4-cycle selector"},
		{Name: "Fano incidence plane", Source: "Gates 115-123", Class: CombinatorialPlane, Dimension: 7, ExactFiniteData: true, FailureReason: "finite incidence geometry is transitive and lower-dimensional; quotienting erases, rather than selects, orientation data"},
		{Name: "16D Fock-spinor spectral Hilbert space", Source: "Gates 14/166/167", Class: InternalHilbert, Dimension: 16, ExactFiniteData: true, OrientedBasisAvailable: true, FailureReason: "internal Hilbert representation has charges/J/γ but no total algebra Hochschild four-cycle or spacetime product triple"},
		{Name: "Gate 174 topological action seal S_top=8π²", Source: "Gate 174", Class: TopologicalScalar, Dimension: 0, ExactFiniteData: true, FailureReason: "topological scalar is an action value, not an integration carrier or integer Chern-Weil charge map"},
		{Name: "collective contact/zeta spectral ledger", Source: "Gates 161/162", Class: TopologicalScalar, Dimension: 7, ExactFiniteData: true, FailureReason: "Galois-invariant spectral functionals are branch-safe but contain no local four-cycle or curvature pairing"},
	}
}

func auditRequirements(candidates []FourCycleCandidate) []RequirementAudit {
	reqs := []RequirementAudit{
		{Name: "oriented four-dimensional carrier", RequiredForFourCycle: true, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "some 4D vector spaces exist, but none is promoted to an integration base"},
		{Name: "boundaryless nonzero cycle", RequiredForFourCycle: true, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "no candidate verifies ∂C₄=0 with a canonical nonzero representative"},
		{Name: "canonical fundamental class / orientation sign", RequiredForFourCycle: true, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "available orientations are basis conventions, not invariant fundamental classes"},
		{Name: "integration functional", RequiredForFourCycle: true, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "no finite ∫_C map is derived"},
		{Name: "principal gauge bundle / connection", RequiredForFourCycle: false, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "gauge algebra closes, but not as a bundle over a four-carrier"},
		{Name: "curvature two-form and wedge square", RequiredForFourCycle: false, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "finite commutators are not yet continuum curvature 2-forms"},
		{Name: "trace normalization", RequiredForFourCycle: false, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "representation ratios are closed, absolute Chern-Weil trace scale is not"},
		{Name: "integer topological charge map", RequiredForFourCycle: false, RequiredForChernWeil: true, RequiredForInstanton: true, Comment: "S_top exists, but no map to k∈Z is derived"},
	}
	for i := range reqs {
		reqs[i].DerivedByAnyCandidate = requirementDerived(reqs[i].Name, candidates)
		reqs[i].Blocking = !reqs[i].DerivedByAnyCandidate
	}
	return reqs
}

func requirementDerived(name string, candidates []FourCycleCandidate) bool {
	switch name {
	case "oriented four-dimensional carrier":
		// Deliberately false: a 4D vector space is not enough unless it is also an oriented integration carrier.
		return false
	case "boundaryless nonzero cycle":
		for _, c := range candidates {
			if c.BoundarylessCycleVerified {
				return true
			}
		}
	case "canonical fundamental class / orientation sign":
		for _, c := range candidates {
			if c.CanonicalFundamentalClass {
				return true
			}
		}
	case "integration functional":
		for _, c := range candidates {
			if c.IntegrationFunctionalDerived {
				return true
			}
		}
	case "principal gauge bundle / connection":
		for _, c := range candidates {
			if c.GaugeBundleMapDerived && c.ConnectionOneFormDerived {
				return true
			}
		}
	case "curvature two-form and wedge square":
		for _, c := range candidates {
			if c.CurvatureTwoFormDerived && c.WedgePairingFAndFDerived {
				return true
			}
		}
	case "trace normalization":
		for _, c := range candidates {
			if c.TraceNormalizationDerived {
				return true
			}
		}
	case "integer topological charge map":
		for _, c := range candidates {
			if c.IntegerTopologicalChargeMap {
				return true
			}
		}
	}
	return false
}

func auditSearch(candidates []FourCycleCandidate) ConstructionSearchAudit {
	a := ConstructionSearchAudit{CandidatesAudited: len(candidates)}
	for _, c := range candidates {
		if c.ExactFiniteData {
			a.ExactFiniteCandidates++
		}
		if c.FourDimensionalCarrier {
			a.FourDimensionalCandidates++
		}
		if c.BoundarylessCycleVerified {
			a.BoundarylessCycleCandidates++
		}
		if c.CanonicalFundamentalClass {
			a.CanonicalFundamentalClasses++
		}
		if c.IntegrationFunctionalDerived {
			a.IntegrationFunctionals++
		}
		if c.GaugeBundleMapDerived {
			a.GaugeBundleMaps++
		}
		if c.CurvatureTwoFormDerived && c.WedgePairingFAndFDerived {
			a.CurvaturePairings++
		}
		if c.TraceNormalizationDerived {
			a.TraceNormalizations++
		}
		if c.IntegerTopologicalChargeMap {
			a.IntegerChargeMaps++
		}
		if c.HochschildOrientabilityCycle {
			a.HochschildCycles++
		}
		if c.ChernWeilCarrierComplete {
			a.CompleteChernWeilCarriers++
		}
	}
	a.ExistingCandidatePromoted = a.CompleteChernWeilCarriers > 0
	a.FiniteFourCycleDerived = a.BoundarylessCycleCandidates > 0 && a.CanonicalFundamentalClasses > 0 && a.IntegrationFunctionals > 0
	a.ChernWeilCarrierDerived = a.CompleteChernWeilCarriers > 0
	a.InstantonBridgePromoted = a.ChernWeilCarrierDerived && a.IntegerChargeMaps > 0 && a.TraceNormalizations > 0
	a.Verdict = "current finite candidates contain grade-four, 4D, spectral, and topological predata, but no complete oriented four-cycle/Chern-Weil carrier"
	return a
}

func FormatCandidate(c FourCycleCandidate) string {
	return fmt.Sprintf("%s[%s](src=%s,dim=%d,exact=%t,4d=%t,orient=%t,boundary=%t,closed=%t,fund=%t,int=%t,bundle=%t,conn=%t,curv=%t,F∧F=%t,trace=%t,kZ=%t,hoch=%t,complete=%t): %s", c.Name, c.Class, c.Source, c.Dimension, c.ExactFiniteData, c.FourDimensionalCarrier, c.OrientedBasisAvailable, c.BoundaryOperatorAvailable, c.BoundarylessCycleVerified, c.CanonicalFundamentalClass, c.IntegrationFunctionalDerived, c.GaugeBundleMapDerived, c.ConnectionOneFormDerived, c.CurvatureTwoFormDerived, c.WedgePairingFAndFDerived, c.TraceNormalizationDerived, c.IntegerTopologicalChargeMap, c.HochschildOrientabilityCycle, c.ChernWeilCarrierComplete, c.FailureReason)
}

func FormatCandidates(xs []FourCycleCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCandidate(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatRequirement(r RequirementAudit) string {
	return fmt.Sprintf("%s(four=%t,chern=%t,instanton=%t,derived=%t,blocking=%t): %s", r.Name, r.RequiredForFourCycle, r.RequiredForChernWeil, r.RequiredForInstanton, r.DerivedByAnyCandidate, r.Blocking, r.Comment)
}

func FormatRequirements(xs []RequirementAudit) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatRequirement(x))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatSearch(a ConstructionSearchAudit) string {
	return fmt.Sprintf("candidates=%d exact=%d fourD=%d closed=%d fundamental=%d integrals=%d bundles=%d curvature=%d trace=%d kZ=%d hochschild=%d complete=%d promoted=%t fourCycle=%t chern=%t instanton=%t: %s", a.CandidatesAudited, a.ExactFiniteCandidates, a.FourDimensionalCandidates, a.BoundarylessCycleCandidates, a.CanonicalFundamentalClasses, a.IntegrationFunctionals, a.GaugeBundleMaps, a.CurvaturePairings, a.TraceNormalizations, a.IntegerChargeMaps, a.HochschildCycles, a.CompleteChernWeilCarriers, a.ExistingCandidatePromoted, a.FiniteFourCycleDerived, a.ChernWeilCarrierDerived, a.InstantonBridgePromoted, a.Verdict)
}

func FormatHomology(h HomologyOrientationAudit) string {
	return fmt.Sprintf("grade4=%t dim4=%t boundary=%t closed4=%t rep=%t orient=%t autSel=%t hochGamma=%t: %s", h.GradeFourDataAvailable, h.DimensionFourVectorSpacesExist, h.BoundaryOperatorDerived, h.NonzeroClosedFourCycleDerived, h.CanonicalRepresentativeDerived, h.OrientationSignDerived, h.AutomorphismInvariantSelector, h.HochschildCycleRealizesGamma, h.Verdict)
}

func FormatChernWeil(c ChernWeilAudit) string {
	return fmt.Sprintf("gauge=%t repTrace=%t Stop=%t bundle=%t conn=%t curv=%t trace=%t intTrFF=%t kZ=%t norm=%t: %s", c.GaugeAlgebraClosed, c.RepresentationTraceRatioClosed, c.TopologicalSealAvailable, c.PrincipalBundleDerived, c.ConnectionOnFourCarrierDerived, c.CurvatureTwoFormDerived, c.TracePairingDerived, c.IntegralOfTrFedgeFDerived, c.IntegerInstantonNumberDerived, c.ContinuumNormalizationPromoted, c.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("observed=%t fourCycle=%t chern=%t instanton=%t abs=%t heat=%t thresholds=%t constants=%t strict=%d->%d conditional=%d->%d sealed=[%s] open=[%s] next=%s: %s", f.UsesObservedInputForDerivation, f.FourCycleDerived, f.ChernWeilCarrierDerived, f.InstantonTraceBridgeDerived, f.AbsoluteCouplingPromoted, f.HeatKernelMatchingDerived, f.ThresholdCorrectedBetaDerived, f.PhysicalConstantsDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.ConditionalNullityBefore, f.ConditionalNullityAfter, strings.Join(f.SealedClaims, "; "), strings.Join(f.OpenRequirements, "; "), f.RecommendedNextGate, f.Verdict)
}
