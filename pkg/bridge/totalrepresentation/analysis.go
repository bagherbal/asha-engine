// Package totalrepresentation implements Gate 165: finite algebra
// representation on the total spectral Hilbert space / faithful action
// obstruction audit.
//
// Gate 164 proved that contact-only Dirac candidates are order-one-safe only
// vacuously, while nontrivial finite Dirac candidates cannot even be lawfully
// tested until a canonical total algebra representation exists. Gate 165 audits
// exactly that missing representation layer. It separates carrier-level exact
// actions from a single faithful action on the total spectral Hilbert space.
//
// The result is conservative: several canonical representations exist on their
// own carriers, but all are block-local. No currently available map glues the
// contact, Fock, scalar, and Clifford/projector sectors into one faithful,
// nontrivial total representation. Therefore no finite Dirac operator, spectral
// triple, gauge fluctuation map, beta row, mass, scale, or physical constant is
// opened.
package totalrepresentation

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/diracorderone"
)

type CarrierCandidate struct {
	Name                  string
	Dimension             int
	Source                string
	Available             bool
	Canonical             bool
	PhysicalCarrier       bool
	AuxiliaryCarrier      bool
	HasInnerProduct       bool
	HasContactSector      bool
	HasMatterSector       bool
	HasScalarSector       bool
	TotalHilbertCandidate bool
	CanonicalTotalHilbert bool
	RequiresObservedInput bool
	RequiresBranchChoice  bool
	Verdict               string
}

type AlgebraActionCandidate struct {
	Name                      string
	Algebra                   string
	Carrier                   string
	Source                    string
	Available                 bool
	CanonicalOnOwnCarrier     bool
	FaithfulOnOwnCarrier      bool
	ActsOnContact             bool
	ActsOnMatter              bool
	ActsOnScalar              bool
	ActsOnAuxiliaryExterior   bool
	FaithfulOnTotalHilbert    bool
	NontrivialAcrossSectors   bool
	Commutative               bool
	GeneratesNonzeroOneForms  bool
	GaugeFluctuationTrivial   bool
	CompatibleWithKnownJ      bool
	CompatibleWithKnownGamma  bool
	RequiresSectorIntertwiner bool
	RequiresRealStructure     bool
	RequiresGrading           bool
	RequiresImportedSMAlgebra bool
	UsesObservedInput         bool
	UsesBranchChoice          bool
	Verdict                   string
}

type GlueMapCandidate struct {
	Name                 string
	From                 string
	To                   string
	Source               string
	Available            bool
	Canonical            bool
	BranchFree           bool
	Injective            bool
	SurjectiveOntoTarget bool
	Intertwining         bool
	Isometric            bool
	GaloisSafe           bool
	UsesObservedInput    bool
	RequiresBranchChoice bool
	Verdict              string
}

type AssemblyCandidate struct {
	Name                        string
	Formula                     string
	Available                   bool
	Canonical                   bool
	BranchFree                  bool
	UsesObservedInput           bool
	UsesBranchChoice            bool
	ContainsContact             bool
	ContainsMatter              bool
	ContainsScalar              bool
	ContainsAuxiliaryExterior   bool
	TotalCarrierComplete        bool
	SingleFiniteAlgebra         bool
	FaithfulTotalRepresentation bool
	NontrivialSectorMixing      bool
	NonzeroOneForms             bool
	CompatibleWithDiracSearch   bool
	CompatibleWithRealStructure bool
	CompatibleWithGrading       bool
	PromotableToSpectralTriple  bool
	GaugeFluctuationMapRows     int
	GaugeKineticMapRows         int
	BoundaryConstraintsDerived  int
	ThresholdBetaRows           int
	PhysicalConstantsDerived    bool
	RequiresImportedSMAlgebra   bool
	Verdict                     string
}

type CarrierAudit struct {
	CandidatesAudited      int
	AvailableCandidates    int
	CanonicalCandidates    int
	PhysicalCarriers       int
	AuxiliaryCarriers      int
	TotalHilbertCandidates int
	CanonicalTotalHilberts int
	ObservedInputsUsed     bool
	BranchChoicesUsed      int
	Verdict                string
}

type AlgebraActionAudit struct {
	CandidatesAudited            int
	AvailableCandidates          int
	CanonicalOwnCarrierActions   int
	FaithfulOwnCarrierActions    int
	ContactActions               int
	MatterActions                int
	ScalarActions                int
	AuxiliaryExteriorActions     int
	FaithfulTotalRepresentations int
	NontrivialCrossSectorActions int
	CommutativeActions           int
	NonzeroOneFormActions        int
	GaugeTrivialActions          int
	RequireSectorIntertwiner     int
	RequireRealStructure         int
	RequireGrading               int
	RequireImportedSMAlgebra     int
	ObservedInputsUsed           bool
	BranchChoicesUsed            int
	Verdict                      string
}

type GlueAudit struct {
	CandidatesAudited      int
	AvailableCandidates    int
	CanonicalCandidates    int
	BranchFreeCandidates   int
	InjectiveCandidates    int
	IntertwiningCandidates int
	IsometricCandidates    int
	GaloisSafeCandidates   int
	ObservedInputsUsed     bool
	BranchChoicesUsed      int
	Verdict                string
}

type AssemblyAudit struct {
	CandidatesAudited            int
	AvailableCandidates          int
	CanonicalCandidates          int
	BranchFreeCandidates         int
	ObservedInputsUsed           bool
	BranchChoicesUsed            int
	ContactContaining            int
	MatterContaining             int
	ScalarContaining             int
	TotalCarrierComplete         int
	SingleFiniteAlgebra          int
	FaithfulTotalRepresentations int
	NontrivialSectorMixing       int
	NonzeroOneForms              int
	DiracCompatible              int
	RealStructureCompatible      int
	GradingCompatible            int
	PromotableSpectralTriples    int
	GaugeFluctuationMapRows      int
	GaugeKineticMapRows          int
	BoundaryConstraintsDerived   int
	ThresholdBetaRows            int
	PhysicalConstantsDerived     bool
	RequiresImportedSMAlgebra    bool
	Verdict                      string
}

type FirewallAudit struct {
	Gate164Inherited             bool
	FiniteSpectralPreData        bool
	CanonicalCarrierActions      int
	FaithfulTotalRepresentations int
	CanonicalGlueMaps            int
	PromotableSpectralTriples    int
	SpectralTripleComplete       bool
	FiniteDiracSelected          bool
	RealStructureSelected        bool
	GradingSelected              bool
	OrderOneCalculusVerified     bool
	GaugeFluctuationMapDerived   bool
	GaugeKineticMapRows          int
	IndividualQuarticRows        int
	CanonicalQuarticBranches     int
	GaugeRepresentationRows      int
	LocalFieldRows               int
	MassActivationRows           int
	DecouplingRows               int
	DynkinIndexRows              int
	ThresholdBetaRows            int
	ProvenZeroRows               int
	BoundaryConstraintsDerived   int
	PhysicalConstantsDerived     bool
	BetaPermissionFirewallClosed bool
	Verdict                      string
}

type Summary struct {
	ContactRows                  int
	ZetaValuesComputed           int
	CarrierCandidatesAudited     int
	CanonicalCarriers            int
	CanonicalTotalHilberts       int
	AlgebraActionsAudited        int
	CanonicalOwnCarrierActions   int
	FaithfulOwnCarrierActions    int
	FaithfulTotalRepresentations int
	GlueMapsAudited              int
	CanonicalGlueMaps            int
	AssembliesAudited            int
	PromotableSpectralTriples    int
	GaugeKineticMapRows          int
	BoundaryConstraintsDerived   int
	ThresholdBetaRows            int
	ResidualNullityBefore        int
	ResidualNullityAfter         int
}

type Analysis struct {
	Previous diracorderone.Analysis

	Carriers       []CarrierCandidate
	AlgebraActions []AlgebraActionCandidate
	GlueMaps       []GlueMapCandidate
	Assemblies     []AssemblyCandidate

	CarrierAudit       CarrierAudit
	AlgebraActionAudit AlgebraActionAudit
	GlueAudit          GlueAudit
	AssemblyAudit      AssemblyAudit
	Firewall           FirewallAudit
	Summary            Summary

	ContactRows                  int
	ContactZetaValues            int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	FiniteSpectralPreData        bool
	CanonicalCarrierActions      int
	FaithfulOwnCarrierActions    int
	FaithfulTotalRepresentations int
	CanonicalGlueMaps            int
	SpectralTripleComplete       bool
	FiniteDiracSelected          bool
	RealStructureSelected        bool
	GradingSelected              bool
	OrderOneCalculusVerified     bool
	OrientabilityVerified        bool
	PoincareDualityVerified      bool
	KOCompatibilityVerified      bool
	CanonicalCutoffSelected      bool
	GaugeFluctuationMapDerived   bool
	SpectralActionPrincipleReady bool
	GaugeKineticMapRows          int
	IndividualQuarticRows        int
	CanonicalQuarticBranches     int
	GaugeRepresentationRows      int
	SpinStatisticsRows           int
	LocalFieldRows               int
	MassActivationRows           int
	DecouplingRows               int
	DynkinIndexRows              int
	ThresholdBetaRows            int
	ProvenZeroRows               int
	BoundaryConstraintsDerived   int
	ContactBetaRowsAllowed       int
	ContactZeroRowsProved        int
	BetaPermissionFirewallClosed bool
	ThresholdCorrectedBeta       bool
	FullBetaMatchingTensor       bool
	ResidualNullityBefore        int
	ResidualNullityAfter         int
	HiddenObservedInputUsed      bool
	PhysicalWeakAngleDerived     bool
	FineStructureDerived         bool
	PhysicalMassesDerived        bool
	PhysicalScaleDerived         bool

	TruthStatement      string
	RejectedClaims      []string
	RemainingUnknowns   []string
	RecommendedNextGate string
}

var defaultOnce sync.Once
var defaultValue Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := diracorderone.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev diracorderone.Analysis) (Analysis, error) {
	if !prev.FiniteSpectralPreData || prev.ContactRows != 7 || prev.ContactZetaValues != 5 || !prev.BetaPermissionFirewallClosed {
		return Analysis{}, fmt.Errorf("Gate 165 requires Gate 164 finite spectral pre-data and closed firewall")
	}
	if prev.SpectralTripleComplete || prev.FiniteDiracSelected || prev.RealStructureSelected || prev.GradingSelected || prev.OrderOneCalculusVerified || prev.GaugeFluctuationMapDerived || prev.GaugeKineticMapRows != 0 || prev.ThresholdBetaRows != 0 || prev.BoundaryConstraintsDerived != 0 || prev.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 165 requires Gate 164 to have no selected spectral triple, Dirac, gauge map, beta row, boundary constraint, or observed input")
	}

	carriers := carrierCandidates()
	actions := algebraActionCandidates()
	glue := glueMapCandidates()
	assemblies := assemblyCandidates()

	carrierAudit := auditCarriers(carriers)
	actionAudit := auditActions(actions)
	glueAudit := auditGlue(glue)
	assemblyAudit := auditAssemblies(assemblies)

	truth := "Gate 165 audits the finite-algebra representation problem on the total spectral Hilbert space. Canonical actions exist on separate carriers: the contact spectral algebra acts on K7, the Boolean/G2 projectors act on Lambda^4 R8, Fock charge operators act on the 16-state matter carrier, and the scalar electroweak action acts on the active scalar carrier. None of these is a faithful representation of one finite algebra on the total spectral Hilbert space, and the formal direct-sum assembly is only blockwise bookkeeping: it does not derive a canonical sector-intertwining representation, real structure, grading, nonzero one-forms, or gauge fluctuations. Therefore the representation prerequisite for a nontrivial finite Dirac operator and spectral action remains missing."

	firewall := FirewallAudit{
		Gate164Inherited:             true,
		FiniteSpectralPreData:        true,
		CanonicalCarrierActions:      actionAudit.CanonicalOwnCarrierActions,
		FaithfulTotalRepresentations: actionAudit.FaithfulTotalRepresentations,
		CanonicalGlueMaps:            glueAudit.CanonicalCandidates,
		PromotableSpectralTriples:    assemblyAudit.PromotableSpectralTriples,
		SpectralTripleComplete:       false,
		FiniteDiracSelected:          false,
		RealStructureSelected:        false,
		GradingSelected:              false,
		OrderOneCalculusVerified:     false,
		GaugeFluctuationMapDerived:   false,
		GaugeKineticMapRows:          0,
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		GaugeRepresentationRows:      0,
		LocalFieldRows:               0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		DynkinIndexRows:              0,
		ThresholdBetaRows:            0,
		ProvenZeroRows:               0,
		BoundaryConstraintsDerived:   0,
		PhysicalConstantsDerived:     false,
		BetaPermissionFirewallClosed: true,
		Verdict:                      "faithful total representation missing; beta firewall remains closed",
	}

	summary := Summary{
		ContactRows:                  prev.ContactRows,
		ZetaValuesComputed:           prev.ContactZetaValues,
		CarrierCandidatesAudited:     carrierAudit.CandidatesAudited,
		CanonicalCarriers:            carrierAudit.CanonicalCandidates,
		CanonicalTotalHilberts:       carrierAudit.CanonicalTotalHilberts,
		AlgebraActionsAudited:        actionAudit.CandidatesAudited,
		CanonicalOwnCarrierActions:   actionAudit.CanonicalOwnCarrierActions,
		FaithfulOwnCarrierActions:    actionAudit.FaithfulOwnCarrierActions,
		FaithfulTotalRepresentations: actionAudit.FaithfulTotalRepresentations,
		GlueMapsAudited:              glueAudit.CandidatesAudited,
		CanonicalGlueMaps:            glueAudit.CanonicalCandidates,
		AssembliesAudited:            assemblyAudit.CandidatesAudited,
		PromotableSpectralTriples:    assemblyAudit.PromotableSpectralTriples,
		GaugeKineticMapRows:          0,
		BoundaryConstraintsDerived:   0,
		ThresholdBetaRows:            0,
		ResidualNullityBefore:        3,
		ResidualNullityAfter:         3,
	}

	return Analysis{
		Previous:                     prev,
		Carriers:                     carriers,
		AlgebraActions:               actions,
		GlueMaps:                     glue,
		Assemblies:                   assemblies,
		CarrierAudit:                 carrierAudit,
		AlgebraActionAudit:           actionAudit,
		GlueAudit:                    glueAudit,
		AssemblyAudit:                assemblyAudit,
		Firewall:                     firewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ContactZetaValues:            prev.ContactZetaValues,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		FiniteSpectralPreData:        true,
		CanonicalCarrierActions:      actionAudit.CanonicalOwnCarrierActions,
		FaithfulOwnCarrierActions:    actionAudit.FaithfulOwnCarrierActions,
		FaithfulTotalRepresentations: actionAudit.FaithfulTotalRepresentations,
		CanonicalGlueMaps:            glueAudit.CanonicalCandidates,
		SpectralTripleComplete:       false,
		FiniteDiracSelected:          false,
		RealStructureSelected:        false,
		GradingSelected:              false,
		OrderOneCalculusVerified:     false,
		OrientabilityVerified:        false,
		PoincareDualityVerified:      false,
		KOCompatibilityVerified:      false,
		CanonicalCutoffSelected:      false,
		GaugeFluctuationMapDerived:   false,
		SpectralActionPrincipleReady: false,
		GaugeKineticMapRows:          0,
		IndividualQuarticRows:        0,
		CanonicalQuarticBranches:     0,
		GaugeRepresentationRows:      0,
		SpinStatisticsRows:           0,
		LocalFieldRows:               0,
		MassActivationRows:           0,
		DecouplingRows:               0,
		DynkinIndexRows:              0,
		ThresholdBetaRows:            0,
		ProvenZeroRows:               0,
		BoundaryConstraintsDerived:   0,
		ContactBetaRowsAllowed:       0,
		ContactZeroRowsProved:        0,
		BetaPermissionFirewallClosed: true,
		ThresholdCorrectedBeta:       false,
		FullBetaMatchingTensor:       false,
		ResidualNullityBefore:        3,
		ResidualNullityAfter:         3,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"do not identify the contact spectral algebra with the full finite algebra",
			"do not import the Connes Standard Model algebra C+H+M3(C) as a derived result",
			"do not treat a formal direct-sum action as a canonical interacting representation",
			"do not infer gauge kinetic rows or threshold beta rows without a faithful total representation",
		},
		RemainingUnknowns: []string{
			"canonical total finite algebra A_total",
			"faithful representation rho: A_total -> End(H_total)",
			"canonical sector-intertwining maps between contact, scalar, and Fock carriers",
			"global real structure J and grading gamma compatible with rho",
			"nontrivial one-form calculus and gauge fluctuation map",
		},
		RecommendedNextGate: "Gate 166 — sector-intertwiner reconstruction / total representation glue-map search",
	}, nil
}

func carrierCandidates() []CarrierCandidate {
	return []CarrierCandidate{
		{Name: "contact vacuum carrier K7", Dimension: 7, Source: "Gate 5 / exact contact overlap tower", Available: true, Canonical: true, PhysicalCarrier: true, HasInnerProduct: true, HasContactSector: true, TotalHilbertCandidate: false, Verdict: "canonical contact carrier only; not a total Hilbert space"},
		{Name: "four-mode Fock matter carrier", Dimension: 16, Source: "Gates 12-17", Available: true, Canonical: true, PhysicalCarrier: true, HasInnerProduct: true, HasMatterSector: true, TotalHilbertCandidate: false, Verdict: "canonical matter bookkeeping carrier only; no contact action"},
		{Name: "active scalar carrier H_phi", Dimension: 4, Source: "Gates 20 and 84-103", Available: true, Canonical: true, PhysicalCarrier: true, HasInnerProduct: true, HasScalarSector: true, TotalHilbertCandidate: false, Verdict: "scalar electroweak carrier only; no full finite algebra action"},
		{Name: "middle exterior chamber Lambda^4 R8", Dimension: 70, Source: "Gates 1-5", Available: true, Canonical: true, AuxiliaryCarrier: true, HasInnerProduct: true, HasContactSector: true, TotalHilbertCandidate: false, Verdict: "exact construction chamber; too large and auxiliary for total spectral Hilbert space"},
		{Name: "formal H_contact plus H_Fock tensor H_phi", Dimension: 71, Source: "formal Gate 165 assembly", Available: true, Canonical: false, PhysicalCarrier: true, HasInnerProduct: true, HasContactSector: true, HasMatterSector: true, HasScalarSector: true, TotalHilbertCandidate: true, CanonicalTotalHilbert: false, Verdict: "contains the sectors but is only a formal direct sum; no canonical representation glue"},
		{Name: "doubled NCG-style H plus JH", Dimension: 142, Source: "spectral-triple target", Available: false, Canonical: false, PhysicalCarrier: true, TotalHilbertCandidate: true, CanonicalTotalHilbert: false, RequiresBranchChoice: true, Verdict: "requires a global real structure and grading not yet derived"},
	}
}

func algebraActionCandidates() []AlgebraActionCandidate {
	return []AlgebraActionCandidate{
		{Name: "contact spectral algebra", Algebra: "Q[Omega_contact]", Carrier: "K7", Source: "Gates 149-162", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: true, ActsOnContact: true, Commutative: true, GaugeFluctuationTrivial: true, Verdict: "canonical exact contact action, but commutative and total-gauge-trivial"},
		{Name: "Boolean/G2 projector algebra", Algebra: "Alg(P_B,P_G)", Carrier: "Lambda^4 R8", Source: "Gates 3-5", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: true, ActsOnContact: true, ActsOnAuxiliaryExterior: true, GaugeFluctuationTrivial: true, Verdict: "canonical construction algebra on the exterior chamber, not a total physical representation"},
		{Name: "Clifford bookkeeping action", Algebra: "Cl(1,7)", Carrier: "exterior/Witt bookkeeping", Source: "Gates 2 and 12-13", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: false, ActsOnAuxiliaryExterior: true, RequiresSectorIntertwiner: true, RequiresRealStructure: true, RequiresGrading: true, Verdict: "bookkeeping action, not a faithful physical total representation"},
		{Name: "Fock charge-number algebra", Algebra: "Q[N0,N1,N2,N3,B-L]", Carrier: "H_Fock", Source: "Gates 12-17", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: true, ActsOnMatter: true, Commutative: true, GaugeFluctuationTrivial: true, Verdict: "exact matter charge action, but it does not act on contact or scalar sectors"},
		{Name: "scalar electroweak action", Algebra: "su(2)+u(1) scalar representation", Carrier: "H_phi", Source: "Gates 20 and 84-103", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: true, ActsOnScalar: true, RequiresRealStructure: true, RequiresGrading: true, Verdict: "variational scalar action, not a full finite algebra representation"},
		{Name: "matter-scalar tensor block", Algebra: "Fock charges tensor scalar EW action", Carrier: "H_Fock tensor H_phi", Source: "Gate 17 plus electroweak ladder", Available: true, CanonicalOnOwnCarrier: true, FaithfulOnOwnCarrier: false, ActsOnMatter: true, ActsOnScalar: true, RequiresSectorIntertwiner: true, RequiresRealStructure: true, RequiresGrading: true, Verdict: "useful block action, still missing the contact sector and total J/gamma"},
		{Name: "formal block-direct-sum action", Algebra: "A_contact direct-sum A_Fock direct-sum A_phi", Carrier: "K7 plus H_Fock tensor H_phi", Source: "formal Gate 165 assembly", Available: true, CanonicalOnOwnCarrier: false, FaithfulOnOwnCarrier: false, ActsOnContact: true, ActsOnMatter: true, ActsOnScalar: true, RequiresSectorIntertwiner: true, RequiresRealStructure: true, RequiresGrading: true, Verdict: "blockwise bookkeeping only; not one derived faithful finite algebra"},
		{Name: "imported Connes Standard Model algebra", Algebra: "C plus H plus M3(C)", Carrier: "NCG finite Hilbert space", Source: "external NCG template", Available: false, CanonicalOnOwnCarrier: false, FaithfulOnOwnCarrier: false, RequiresImportedSMAlgebra: true, RequiresSectorIntertwiner: true, RequiresRealStructure: true, RequiresGrading: true, UsesObservedInput: false, Verdict: "not derived from the finite Boolean-Octonionic engine; importing it would violate the gate"},
	}
}

func glueMapCandidates() []GlueMapCandidate {
	return []GlueMapCandidate{
		{Name: "Fock-contact kernel", From: "H_Fock", To: "K7", Source: "Gates 14-15 and 138-139", Available: true, Canonical: false, BranchFree: true, Injective: false, Intertwining: false, Isometric: false, GaloisSafe: true, Verdict: "diagnostic bridge exists, but kernel/intertwiner is non-unique"},
		{Name: "scalar-contact block connection", From: "H_phi", To: "K7", Source: "Gates 11 and 160", Available: true, Canonical: false, BranchFree: true, Injective: false, Intertwining: false, Isometric: false, GaloisSafe: true, Verdict: "reaches protected/contact data only as a variational diagnostic; no canonical total intertwiner"},
		{Name: "broken-generator contact projection", From: "broken scalar orbit", To: "quartic contact block", Source: "Gates 96-100 and 160", Available: true, Canonical: false, BranchFree: true, Injective: false, Intertwining: false, Isometric: false, GaloisSafe: true, Verdict: "Gate 160 showed it does not split or glue the quartic block"},
		{Name: "Clifford-Witt-to-Fock identification", From: "Cl(1,7) bookkeeping", To: "H_Fock", Source: "Gates 12-13", Available: true, Canonical: false, BranchFree: true, Injective: false, Intertwining: false, Isometric: false, GaloisSafe: true, Verdict: "bookkeeping identification, not a representation-preserving functor"},
		{Name: "global sector functor", From: "finite construction tower", To: "H_total", Source: "Gate 165 target", Available: false, Canonical: false, BranchFree: false, Injective: false, Intertwining: false, Isometric: false, GaloisSafe: false, RequiresBranchChoice: true, Verdict: "the missing object: a canonical total-sector representation functor"},
	}
}

func assemblyCandidates() []AssemblyCandidate {
	return []AssemblyCandidate{
		{Name: "contact-only spectral representation", Formula: "rho(a)=a(Omega) on K7", Available: true, Canonical: true, BranchFree: true, ContainsContact: true, SingleFiniteAlgebra: true, Verdict: "canonical but not total and produces no nontrivial gauge calculus"},
		{Name: "matter-scalar tensor representation", Formula: "rho_Fock tensor rho_phi on H_Fock tensor H_phi", Available: true, Canonical: true, BranchFree: true, ContainsMatter: true, ContainsScalar: true, SingleFiniteAlgebra: false, Verdict: "captures charge/scalar bookkeeping but omits contact spectral carrier"},
		{Name: "exterior projector master action", Formula: "rho(P_B,P_G) on Lambda^4 R8", Available: true, Canonical: true, BranchFree: true, ContainsContact: true, ContainsAuxiliaryExterior: true, SingleFiniteAlgebra: true, Verdict: "exact finite construction action but not the physical total spectral Hilbert representation"},
		{Name: "formal direct-sum total action", Formula: "rho_K direct-sum rho_Fock direct-sum rho_phi", Available: true, Canonical: false, BranchFree: true, ContainsContact: true, ContainsMatter: true, ContainsScalar: true, TotalCarrierComplete: true, SingleFiniteAlgebra: false, Verdict: "sector listing, not a canonical interacting representation"},
		{Name: "imported NCG Standard Model representation", Formula: "rho(C+H+M3(C))", Available: false, Canonical: false, BranchFree: false, UsesObservedInput: false, UsesBranchChoice: true, ContainsMatter: true, ContainsScalar: true, RequiresImportedSMAlgebra: true, Verdict: "external template not derived by the finite engine"},
		{Name: "Asha total faithful representation", Formula: "rho: A_total -> End(H_total)", Available: false, Canonical: false, BranchFree: false, ContainsContact: true, ContainsMatter: true, ContainsScalar: true, TotalCarrierComplete: true, SingleFiniteAlgebra: true, Verdict: "target object remains unconstructed"},
	}
}

// RequiresImportedSMAlgebra is intentionally repeated on AssemblyCandidate through
// embedding in the verdict only; keep the struct smaller and the gate focused.

func auditCarriers(items []CarrierCandidate) CarrierAudit {
	a := CarrierAudit{CandidatesAudited: len(items), Verdict: "canonical local carriers exist; no canonical total Hilbert space exists"}
	for _, item := range items {
		if item.Available {
			a.AvailableCandidates++
		}
		if item.Canonical {
			a.CanonicalCandidates++
		}
		if item.PhysicalCarrier {
			a.PhysicalCarriers++
		}
		if item.AuxiliaryCarrier {
			a.AuxiliaryCarriers++
		}
		if item.TotalHilbertCandidate {
			a.TotalHilbertCandidates++
		}
		if item.CanonicalTotalHilbert {
			a.CanonicalTotalHilberts++
		}
		if item.RequiresObservedInput {
			a.ObservedInputsUsed = true
		}
		if item.RequiresBranchChoice {
			a.BranchChoicesUsed++
		}
	}
	return a
}

func auditActions(items []AlgebraActionCandidate) AlgebraActionAudit {
	a := AlgebraActionAudit{CandidatesAudited: len(items), Verdict: "own-carrier actions exist, but no faithful total representation exists"}
	for _, item := range items {
		if item.Available {
			a.AvailableCandidates++
		}
		if item.CanonicalOnOwnCarrier {
			a.CanonicalOwnCarrierActions++
		}
		if item.FaithfulOnOwnCarrier {
			a.FaithfulOwnCarrierActions++
		}
		if item.ActsOnContact {
			a.ContactActions++
		}
		if item.ActsOnMatter {
			a.MatterActions++
		}
		if item.ActsOnScalar {
			a.ScalarActions++
		}
		if item.ActsOnAuxiliaryExterior {
			a.AuxiliaryExteriorActions++
		}
		if item.FaithfulOnTotalHilbert {
			a.FaithfulTotalRepresentations++
		}
		if item.NontrivialAcrossSectors {
			a.NontrivialCrossSectorActions++
		}
		if item.Commutative {
			a.CommutativeActions++
		}
		if item.GeneratesNonzeroOneForms {
			a.NonzeroOneFormActions++
		}
		if item.GaugeFluctuationTrivial {
			a.GaugeTrivialActions++
		}
		if item.RequiresSectorIntertwiner {
			a.RequireSectorIntertwiner++
		}
		if item.RequiresRealStructure {
			a.RequireRealStructure++
		}
		if item.RequiresGrading {
			a.RequireGrading++
		}
		if item.RequiresImportedSMAlgebra {
			a.RequireImportedSMAlgebra++
		}
		if item.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if item.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
	}
	return a
}

func auditGlue(items []GlueMapCandidate) GlueAudit {
	a := GlueAudit{CandidatesAudited: len(items), Verdict: "no candidate is a canonical sector-intertwining glue map"}
	for _, item := range items {
		if item.Available {
			a.AvailableCandidates++
		}
		if item.Canonical {
			a.CanonicalCandidates++
		}
		if item.BranchFree {
			a.BranchFreeCandidates++
		}
		if item.Injective {
			a.InjectiveCandidates++
		}
		if item.Intertwining {
			a.IntertwiningCandidates++
		}
		if item.Isometric {
			a.IsometricCandidates++
		}
		if item.GaloisSafe {
			a.GaloisSafeCandidates++
		}
		if item.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if item.RequiresBranchChoice {
			a.BranchChoicesUsed++
		}
	}
	return a
}

func auditAssemblies(items []AssemblyCandidate) AssemblyAudit {
	a := AssemblyAudit{CandidatesAudited: len(items), Verdict: "no assembly is a faithful canonical total representation"}
	for _, item := range items {
		if item.Available {
			a.AvailableCandidates++
		}
		if item.Canonical {
			a.CanonicalCandidates++
		}
		if item.BranchFree {
			a.BranchFreeCandidates++
		}
		if item.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if item.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
		if item.ContainsContact {
			a.ContactContaining++
		}
		if item.ContainsMatter {
			a.MatterContaining++
		}
		if item.ContainsScalar {
			a.ScalarContaining++
		}
		if item.TotalCarrierComplete {
			a.TotalCarrierComplete++
		}
		if item.SingleFiniteAlgebra {
			a.SingleFiniteAlgebra++
		}
		if item.FaithfulTotalRepresentation {
			a.FaithfulTotalRepresentations++
		}
		if item.NontrivialSectorMixing {
			a.NontrivialSectorMixing++
		}
		if item.NonzeroOneForms {
			a.NonzeroOneForms++
		}
		if item.CompatibleWithDiracSearch {
			a.DiracCompatible++
		}
		if item.CompatibleWithRealStructure {
			a.RealStructureCompatible++
		}
		if item.CompatibleWithGrading {
			a.GradingCompatible++
		}
		if item.PromotableToSpectralTriple {
			a.PromotableSpectralTriples++
		}
		a.GaugeFluctuationMapRows += item.GaugeFluctuationMapRows
		a.GaugeKineticMapRows += item.GaugeKineticMapRows
		a.BoundaryConstraintsDerived += item.BoundaryConstraintsDerived
		a.ThresholdBetaRows += item.ThresholdBetaRows
		if item.PhysicalConstantsDerived {
			a.PhysicalConstantsDerived = true
		}
	}
	return a
}

func FormatCarrier(c CarrierCandidate) string {
	return fmt.Sprintf("%s dim=%d available=%t canonical=%t physical=%t auxiliary=%t totalCandidate=%t canonicalTotal=%t contact=%t matter=%t scalar=%t observed=%t branch=%t (%s)", c.Name, c.Dimension, c.Available, c.Canonical, c.PhysicalCarrier, c.AuxiliaryCarrier, c.TotalHilbertCandidate, c.CanonicalTotalHilbert, c.HasContactSector, c.HasMatterSector, c.HasScalarSector, c.RequiresObservedInput, c.RequiresBranchChoice, c.Verdict)
}

func FormatAction(c AlgebraActionCandidate) string {
	return fmt.Sprintf("%s algebra=%s carrier=%s available=%t canonicalOwn=%t faithfulOwn=%t acts(K=%t,F=%t,phi=%t,aux=%t) faithfulTotal=%t cross=%t comm=%t oneForms=%t trivial=%t needs(map=%t,J=%t,gamma=%t,SM=%t) observed=%t branch=%t (%s)", c.Name, c.Algebra, c.Carrier, c.Available, c.CanonicalOnOwnCarrier, c.FaithfulOnOwnCarrier, c.ActsOnContact, c.ActsOnMatter, c.ActsOnScalar, c.ActsOnAuxiliaryExterior, c.FaithfulOnTotalHilbert, c.NontrivialAcrossSectors, c.Commutative, c.GeneratesNonzeroOneForms, c.GaugeFluctuationTrivial, c.RequiresSectorIntertwiner, c.RequiresRealStructure, c.RequiresGrading, c.RequiresImportedSMAlgebra, c.UsesObservedInput, c.UsesBranchChoice, c.Verdict)
}

func FormatGlue(c GlueMapCandidate) string {
	return fmt.Sprintf("%s %s->%s available=%t canonical=%t branchFree=%t injective=%t onto=%t intertwining=%t isometric=%t galois=%t observed=%t branch=%t (%s)", c.Name, c.From, c.To, c.Available, c.Canonical, c.BranchFree, c.Injective, c.SurjectiveOntoTarget, c.Intertwining, c.Isometric, c.GaloisSafe, c.UsesObservedInput, c.RequiresBranchChoice, c.Verdict)
}

func FormatAssembly(c AssemblyCandidate) string {
	return fmt.Sprintf("%s: %s available=%t canonical=%t branchFree=%t observed=%t branch=%t contains(K=%t,F=%t,phi=%t,aux=%t) total=%t singleA=%t faithful=%t mixing=%t oneForms=%t D=%t J=%t gamma=%t triple=%t gaugeMap=%d gaugeRows=%d constraints=%d beta=%d physical=%t (%s)", c.Name, c.Formula, c.Available, c.Canonical, c.BranchFree, c.UsesObservedInput, c.UsesBranchChoice, c.ContainsContact, c.ContainsMatter, c.ContainsScalar, c.ContainsAuxiliaryExterior, c.TotalCarrierComplete, c.SingleFiniteAlgebra, c.FaithfulTotalRepresentation, c.NontrivialSectorMixing, c.NonzeroOneForms, c.CompatibleWithDiracSearch, c.CompatibleWithRealStructure, c.CompatibleWithGrading, c.PromotableToSpectralTriple, c.GaugeFluctuationMapRows, c.GaugeKineticMapRows, c.BoundaryConstraintsDerived, c.ThresholdBetaRows, c.PhysicalConstantsDerived, c.Verdict)
}

func FormatCarrierAudit(a CarrierAudit) string {
	return fmt.Sprintf("carriers=%d available=%d canonical=%d physical=%d auxiliary=%d totalCandidates=%d canonicalTotal=%d observed=%t branchChoices=%d (%s)", a.CandidatesAudited, a.AvailableCandidates, a.CanonicalCandidates, a.PhysicalCarriers, a.AuxiliaryCarriers, a.TotalHilbertCandidates, a.CanonicalTotalHilberts, a.ObservedInputsUsed, a.BranchChoicesUsed, a.Verdict)
}

func FormatActionAudit(a AlgebraActionAudit) string {
	return fmt.Sprintf("actions=%d available=%d canonicalOwn=%d faithfulOwn=%d acts(K=%d,F=%d,phi=%d,aux=%d) faithfulTotal=%d cross=%d comm=%d oneForms=%d trivial=%d needs(map=%d,J=%d,gamma=%d,SM=%d) observed=%t branchChoices=%d (%s)", a.CandidatesAudited, a.AvailableCandidates, a.CanonicalOwnCarrierActions, a.FaithfulOwnCarrierActions, a.ContactActions, a.MatterActions, a.ScalarActions, a.AuxiliaryExteriorActions, a.FaithfulTotalRepresentations, a.NontrivialCrossSectorActions, a.CommutativeActions, a.NonzeroOneFormActions, a.GaugeTrivialActions, a.RequireSectorIntertwiner, a.RequireRealStructure, a.RequireGrading, a.RequireImportedSMAlgebra, a.ObservedInputsUsed, a.BranchChoicesUsed, a.Verdict)
}

func FormatGlueAudit(a GlueAudit) string {
	return fmt.Sprintf("glue=%d available=%d canonical=%d branchFree=%d injective=%d intertwining=%d isometric=%d galois=%d observed=%t branchChoices=%d (%s)", a.CandidatesAudited, a.AvailableCandidates, a.CanonicalCandidates, a.BranchFreeCandidates, a.InjectiveCandidates, a.IntertwiningCandidates, a.IsometricCandidates, a.GaloisSafeCandidates, a.ObservedInputsUsed, a.BranchChoicesUsed, a.Verdict)
}

func FormatAssemblyAudit(a AssemblyAudit) string {
	return fmt.Sprintf("assemblies=%d available=%d canonical=%d branchFree=%d observed=%t branchChoices=%d contains(K=%d,F=%d,phi=%d) total=%d singleA=%d faithful=%d mixing=%d oneForms=%d D=%d J=%d gamma=%d triples=%d gaugeMap=%d gaugeRows=%d constraints=%d beta=%d physical=%t (%s)", a.CandidatesAudited, a.AvailableCandidates, a.CanonicalCandidates, a.BranchFreeCandidates, a.ObservedInputsUsed, a.BranchChoicesUsed, a.ContactContaining, a.MatterContaining, a.ScalarContaining, a.TotalCarrierComplete, a.SingleFiniteAlgebra, a.FaithfulTotalRepresentations, a.NontrivialSectorMixing, a.NonzeroOneForms, a.DiracCompatible, a.RealStructureCompatible, a.GradingCompatible, a.PromotableSpectralTriples, a.GaugeFluctuationMapRows, a.GaugeKineticMapRows, a.BoundaryConstraintsDerived, a.ThresholdBetaRows, a.PhysicalConstantsDerived, a.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate164=%t preData=%t canonicalCarrierActions=%d faithfulTotal=%d glue=%d triples=%d triple=%t D=%t J=%t gamma=%t order1=%t gauge=%t gaugeRows=%d individualQuartic=%d branches=%d gaugeRep=%d local=%d mass=%d decoupling=%d dynkin=%d beta=%d zero=%d constraints=%d physical=%t closed=%t (%s)", f.Gate164Inherited, f.FiniteSpectralPreData, f.CanonicalCarrierActions, f.FaithfulTotalRepresentations, f.CanonicalGlueMaps, f.PromotableSpectralTriples, f.SpectralTripleComplete, f.FiniteDiracSelected, f.RealStructureSelected, f.GradingSelected, f.OrderOneCalculusVerified, f.GaugeFluctuationMapDerived, f.GaugeKineticMapRows, f.IndividualQuarticRows, f.CanonicalQuarticBranches, f.GaugeRepresentationRows, f.LocalFieldRows, f.MassActivationRows, f.DecouplingRows, f.DynkinIndexRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.BoundaryConstraintsDerived, f.PhysicalConstantsDerived, f.BetaPermissionFirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d zeta=%d carriers=%d canonicalCarriers=%d canonicalTotalHilberts=%d actions=%d canonicalOwn=%d faithfulOwn=%d faithfulTotal=%d glue=%d canonicalGlue=%d assemblies=%d triples=%d gaugeRows=%d constraints=%d beta=%d nullity=%d→%d", s.ContactRows, s.ZetaValuesComputed, s.CarrierCandidatesAudited, s.CanonicalCarriers, s.CanonicalTotalHilberts, s.AlgebraActionsAudited, s.CanonicalOwnCarrierActions, s.FaithfulOwnCarrierActions, s.FaithfulTotalRepresentations, s.GlueMapsAudited, s.CanonicalGlueMaps, s.AssembliesAudited, s.PromotableSpectralTriples, s.GaugeKineticMapRows, s.BoundaryConstraintsDerived, s.ThresholdBetaRows, s.ResidualNullityBefore, s.ResidualNullityAfter)
}

func FormatActions(items []AlgebraActionCandidate) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s(total=%t,cross=%t,oneForms=%t)", item.Name, item.FaithfulOnTotalHilbert, item.NontrivialAcrossSectors, item.GeneratesNonzeroOneForms))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatAssemblies(items []AssemblyCandidate) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s(total=%t,faithful=%t,triple=%t)", item.Name, item.TotalCarrierComplete, item.FaithfulTotalRepresentation, item.PromotableToSpectralTriple))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}
