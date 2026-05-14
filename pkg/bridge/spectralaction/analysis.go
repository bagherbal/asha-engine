// Package spectralaction implements Gate 163: finite spectral action principle /
// spectral triple construction audit.
//
// Gate 162 proved that the seven-root contact zeta ledger is exact, rational,
// Galois-invariant, branch-free, and pole-free. Gate 163 asks the stronger
// question: can those spectral invariants be promoted to a genuine finite
// spectral action in the noncommutative-geometric sense?
//
// The answer is deliberately conservative. The gate separates available exact
// spectral pre-data from the missing spectral-triple ingredients. It audits
// finite Dirac-like candidates and spectral-action ansatzes, but refuses to
// interpret any zeta scalar as a coupling, threshold row, mass, or physical
// constant without a representation-complete finite algebra/Hilbert/Dirac/real
// structure/grading/cutoff/gauge-map chain.
package spectralaction

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/contactzeta"
)

type Ingredient struct {
	Name               string
	Source             string
	Available          bool
	Canonical          bool
	RequiredForAction  bool
	BlocksConstruction bool
	Verdict            string
}

type DiracCandidate struct {
	Name                       string
	Formula                    string
	ExactOverQ                 bool
	SelfAdjoint                bool
	FiniteSpectrum             bool
	GaloisInvariant            bool
	BranchFree                 bool
	UsesObservedInput          bool
	UsesBranchChoice           bool
	RequiresIndividualQuartic  bool
	NeedsAlgebraRepresentation bool
	NeedsRealStructure         bool
	NeedsGrading               bool
	NeedsOrderOneCalculus      bool
	NeedsGaugeFluctuationMap   bool
	OrderOneVerified           bool
	OrientabilityVerified      bool
	PoincareDualityVerified    bool
	KOCompatibilityVerified    bool
	GaugeKineticMapDerived     bool
	ActionCoefficientCanonical bool
	PromotableToSpectralTriple bool
	BetaRowsAllowed            int
	PhysicalConstantsDerived   bool
	Verdict                    string
}

type ActionAnsatz struct {
	Name                        string
	Formula                     string
	UsesZetaLedger              bool
	ExactOverQ                  bool
	GaloisInvariant             bool
	BranchFree                  bool
	UsesObservedInput           bool
	UsesBranchChoice            bool
	RequiresCutoffFunction      bool
	RequiresTestFunction        bool
	RequiresDiracOperator       bool
	RequiresGaugeRepresentation bool
	CoefficientsCanonical       bool
	GaugeKineticRows            int
	BoundaryConstraintsDerived  int
	ThresholdBetaRows           int
	PhysicalConstantsDerived    bool
	Verdict                     string
}

type AxiomAudit struct {
	IngredientsAudited         int
	AvailableIngredients       int
	CanonicalIngredients       int
	RequiredIngredients        int
	MissingRequiredCanonical   int
	FiniteAlgebraAvailable     bool
	FiniteHilbertCandidate     bool
	ContactZetaLedgerAvailable bool
	ContactOverlapAvailable    bool
	AlgebraRepresentationReady bool
	FiniteDiracSelected        bool
	RealStructureSelected      bool
	GradingSelected            bool
	OrderOneCalculusVerified   bool
	OrientabilityVerified      bool
	PoincareDualityVerified    bool
	KOCompatibilityVerified    bool
	CanonicalCutoffSelected    bool
	GaugeFluctuationMapDerived bool
	SpectralTripleComplete     bool
	Verdict                    string
}

type DiracAudit struct {
	CandidatesAudited           int
	ExactCandidates             int
	SelfAdjointCandidates       int
	FiniteSpectrumCandidates    int
	GaloisInvariantCandidates   int
	BranchFreeCandidates        int
	ObservedInputsUsed          bool
	BranchChoicesUsed           int
	NeedRepresentation          int
	NeedRealStructure           int
	NeedGrading                 int
	NeedOrderOneCalculus        int
	NeedGaugeFluctuationMap     int
	OrderOneVerified            int
	PromotableCandidates        int
	GaugeKineticMapsDerived     int
	CanonicalActionCoefficients int
	BetaRowsAllowed             int
	PhysicalConstantsDerived    bool
	Verdict                     string
}

type ActionAudit struct {
	AnsatzesAudited            int
	UsingZetaLedger            int
	ExactAnsatzes              int
	GaloisInvariantAnsatzes    int
	BranchFreeAnsatzes         int
	ObservedInputsUsed         bool
	BranchChoicesUsed          int
	NeedCutoffFunction         int
	NeedTestFunction           int
	NeedDiracOperator          int
	NeedGaugeRepresentation    int
	CanonicalCoefficients      int
	GaugeKineticRows           int
	BoundaryConstraintsDerived int
	ThresholdBetaRows          int
	PhysicalConstantsDerived   bool
	Verdict                    string
}

type FirewallAudit struct {
	Gate162Inherited             bool
	SpectralTripleComplete       bool
	FiniteDiracSelected          bool
	RealStructureSelected        bool
	GradingSelected              bool
	CanonicalCutoffSelected      bool
	GaugeFluctuationMapDerived   bool
	GaugeKineticMapRows          int
	IndividualQuarticRows        int
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
	ContactRows                 int
	ZetaValuesComputed          int
	IngredientsAudited          int
	MissingRequiredCanonical    int
	DiracCandidatesAudited      int
	PromotableDiracCandidates   int
	ActionAnsatzesAudited       int
	CanonicalActionCoefficients int
	GaugeKineticRows            int
	BoundaryConstraintsDerived  int
	ThresholdBetaRows           int
	ResidualNullityBefore       int
	ResidualNullityAfter        int
}

type Analysis struct {
	Previous contactzeta.Analysis

	Ingredients     []Ingredient
	DiracCandidates []DiracCandidate
	ActionAnsatzes  []ActionAnsatz
	AxiomAudit      AxiomAudit
	DiracAudit      DiracAudit
	ActionAudit     ActionAudit
	Firewall        FirewallAudit
	Summary         Summary

	ContactRows                  int
	ContactZetaValues            int
	ExactRationalOverlapMatrix   bool
	ExactCharacteristicCertified bool
	ExactRootIsolationCertified  bool
	FiniteZetaPoleCount          int
	AnalyticContinuationNeeded   bool
	PositiveNonzeroSpectrumRows  int
	GaloisInvariantOrbits        int
	RationalSingletonRows        int
	QuarticOrbitRows             int
	QuarticCollectiveBlocks      int
	FiniteSpectralPreData        bool
	SpectralTripleComplete       bool
	FiniteAlgebraAvailable       bool
	FiniteHilbertCandidate       bool
	AlgebraRepresentationReady   bool
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
		prev, err := contactzeta.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev)
	})
	return defaultValue, defaultErr
}

func Build(prev contactzeta.Analysis) (Analysis, error) {
	if !prev.BetaPermissionFirewallClosed || prev.ContactBetaRowsAllowed != 0 || prev.ContactZeroRowsProved != 0 || prev.BoundaryConstraintsDerived != 0 {
		return Analysis{}, fmt.Errorf("Gate 163 requires Gate 162 finite-zeta firewall")
	}
	if prev.SpectralTripleComplete || prev.FiniteDiracSelected || prev.RealStructureSelected || prev.GradingSelected || prev.CanonicalCutoffSelected || prev.GaugeKineticMapRows != 0 {
		return Analysis{}, fmt.Errorf("Gate 163 expects Gate 162 to leave spectral-action structure unselected")
	}
	if prev.HiddenObservedInputUsed || prev.PhysicalWeakAngleDerived || prev.FineStructureDerived || prev.PhysicalMassesDerived || prev.PhysicalScaleDerived {
		return Analysis{}, fmt.Errorf("Gate 163 refuses hidden observed physical input")
	}

	ingredients := spectralTripleIngredients()
	diracCandidates := finiteDiracCandidates()
	actionAnsatzes := spectralActionAnsatzes()

	axiomAudit := auditAxioms(ingredients)
	diracAudit := auditDirac(diracCandidates)
	actionAudit := auditAction(actionAnsatzes)
	firewall := FirewallAudit{
		Gate162Inherited:             true,
		SpectralTripleComplete:       false,
		FiniteDiracSelected:          false,
		RealStructureSelected:        false,
		GradingSelected:              false,
		CanonicalCutoffSelected:      false,
		GaugeFluctuationMapDerived:   false,
		GaugeKineticMapRows:          0,
		IndividualQuarticRows:        0,
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
		Verdict:                      "finite spectral pre-data exist, but the spectral triple and gauge fluctuation chain are incomplete; beta and physical constants remain sealed",
	}

	summary := Summary{
		ContactRows:                 prev.ContactRows,
		ZetaValuesComputed:          prev.ContactZetaValues,
		IngredientsAudited:          len(ingredients),
		MissingRequiredCanonical:    axiomAudit.MissingRequiredCanonical,
		DiracCandidatesAudited:      len(diracCandidates),
		PromotableDiracCandidates:   diracAudit.PromotableCandidates,
		ActionAnsatzesAudited:       len(actionAnsatzes),
		CanonicalActionCoefficients: actionAudit.CanonicalCoefficients,
		GaugeKineticRows:            actionAudit.GaugeKineticRows,
		BoundaryConstraintsDerived:  actionAudit.BoundaryConstraintsDerived,
		ThresholdBetaRows:           actionAudit.ThresholdBetaRows,
		ResidualNullityBefore:       prev.ResidualNullityAfter,
		ResidualNullityAfter:        prev.ResidualNullityAfter,
	}

	truth := "Gate 163 audits whether the exact seven-root zeta ledger can be promoted into a finite spectral action principle. The finite contact overlap, characteristic data, positive nonzero spectrum, and zeta values are usable spectral pre-data. However, no canonical finite spectral triple is completed: the representation on a spectral Hilbert space, finite Dirac operator, real structure, grading, order-one calculus, orientability/Poincare-duality checks, cutoff/test function, and gauge fluctuation map remain unselected. Consequently no spectral-action coefficient, gauge kinetic row, threshold beta row, boundary constraint, mass, scale, or physical constant is derived."

	return Analysis{
		Previous:                     prev,
		Ingredients:                  ingredients,
		DiracCandidates:              diracCandidates,
		ActionAnsatzes:               actionAnsatzes,
		AxiomAudit:                   axiomAudit,
		DiracAudit:                   diracAudit,
		ActionAudit:                  actionAudit,
		Firewall:                     firewall,
		Summary:                      summary,
		ContactRows:                  prev.ContactRows,
		ContactZetaValues:            prev.ContactZetaValues,
		ExactRationalOverlapMatrix:   prev.ExactRationalOverlapMatrix,
		ExactCharacteristicCertified: prev.ExactCharacteristicCertified,
		ExactRootIsolationCertified:  prev.ExactRootIsolationCertified,
		FiniteZetaPoleCount:          prev.FiniteZetaPoleCount,
		AnalyticContinuationNeeded:   prev.AnalyticContinuationNeeded,
		PositiveNonzeroSpectrumRows:  prev.PositiveNonzeroSpectrumRows,
		GaloisInvariantOrbits:        prev.GaloisInvariantOrbits,
		RationalSingletonRows:        prev.RationalSingletonRows,
		QuarticOrbitRows:             prev.QuarticOrbitRows,
		QuarticCollectiveBlocks:      prev.QuarticCollectiveBlocks,
		FiniteSpectralPreData:        true,
		SpectralTripleComplete:       false,
		FiniteAlgebraAvailable:       axiomAudit.FiniteAlgebraAvailable,
		FiniteHilbertCandidate:       axiomAudit.FiniteHilbertCandidate,
		AlgebraRepresentationReady:   false,
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
		ResidualNullityBefore:        prev.ResidualNullityAfter,
		ResidualNullityAfter:         prev.ResidualNullityAfter,
		HiddenObservedInputUsed:      false,
		PhysicalWeakAngleDerived:     false,
		FineStructureDerived:         false,
		PhysicalMassesDerived:        false,
		PhysicalScaleDerived:         false,
		TruthStatement:               truth,
		RejectedClaims: []string{
			"the finite zeta ledger by itself is a spectral action principle",
			"the contact overlap Omega is automatically the finite Dirac operator of a spectral triple",
			"a formal heat/zeta expansion has canonical coefficients without a cutoff/test function",
			"spectral-action scalars derive kappa_U1, 5/3 normalization, weak-angle data, or physical constants",
			"finite spectral pre-data permit threshold beta rows without a gauge fluctuation and representation-complete map",
		},
		RemainingUnknowns: []string{
			"canonical representation of the finite algebra on the spectral Hilbert carrier",
			"finite Dirac-like operator satisfying the spectral-triple axioms",
			"real structure J and grading gamma compatible with the quartic Galois firewall",
			"order-one calculus, orientability, Poincare duality, and KO-compatibility checks",
			"canonical cutoff/test function and coefficient ledger for Tr(f(D/Lambda))",
			"gauge fluctuation map from spectral action terms to gauge kinetic rows",
		},
		RecommendedNextGate: "Gate 164 — finite Dirac candidate construction / order-one axiom obstruction audit",
	}, nil
}

func spectralTripleIngredients() []Ingredient {
	return []Ingredient{
		ingredient("finite Boolean-octonionic algebraic carrier", "Gates 1-5,149", true, true, true, false, "exact finite carrier exists"),
		ingredient("contact overlap spectral operator Omega", "Gates 149-162", true, true, true, false, "exact rational overlap and spectral ledger exist"),
		ingredient("finite zeta ledger", "Gate 162", true, true, true, false, "branch-free seven-root zeta values exist"),
		ingredient("spectral Hilbert representation of the algebra", "not yet constructed", true, false, true, true, "candidate finite Hilbert sectors exist, but no spectral-triple representation map is canonical"),
		ingredient("finite Dirac operator D", "not yet selected", false, false, true, true, "Omega-like candidates are only pre-Dirac diagnostics until axioms are checked"),
		ingredient("real structure J", "not yet selected", false, false, true, true, "no canonical charge-conjugation/KO structure survives the quartic firewall"),
		ingredient("grading gamma", "Gate 159 obstruction", false, false, true, true, "nontrivial quartic ghost/parity grading is not canonical"),
		ingredient("order-one calculus", "not yet verified", false, false, true, true, "commutator calculus requires representation and D"),
		ingredient("orientability and Poincare duality", "not yet verified", false, false, true, true, "intersection-form data are not defined without J/gamma/representation"),
		ingredient("canonical cutoff/test function", "not yet selected", false, false, true, true, "finite action coefficients are arbitrary without a cutoff rule"),
		ingredient("gauge fluctuation to kinetic-row map", "not yet derived", false, false, true, true, "no representation-complete map from spectral terms to beta/coupling rows"),
	}
}

func ingredient(name, source string, available, canonical, required, blocks bool, verdict string) Ingredient {
	return Ingredient{Name: name, Source: source, Available: available, Canonical: canonical, RequiredForAction: required, BlocksConstruction: blocks, Verdict: verdict}
}

func finiteDiracCandidates() []DiracCandidate {
	return []DiracCandidate{
		diracCandidate("contact overlap", "D = Omega_contact", true, true, true, true, true),
		diracCandidate("inverse contact overlap", "D = Omega_contact^{-1}", true, true, true, true, true),
		diracCandidate("centered contact overlap", "D = Omega_contact - Tr(Omega)/7 I", true, true, true, true, true),
		diracCandidate("zeta-normalized contact overlap", "D = 7 Omega_contact / Tr(Omega)", true, true, true, true, true),
		diracCandidate("quartic collective scalar block", "D_q = spectral scalar on the quartic primary block", true, true, true, true, true),
	}
}

func diracCandidate(name, formula string, exact, selfAdjoint, finiteSpectrum, galois, branchFree bool) DiracCandidate {
	return DiracCandidate{
		Name:                       name,
		Formula:                    formula,
		ExactOverQ:                 exact,
		SelfAdjoint:                selfAdjoint,
		FiniteSpectrum:             finiteSpectrum,
		GaloisInvariant:            galois,
		BranchFree:                 branchFree,
		UsesObservedInput:          false,
		UsesBranchChoice:           false,
		RequiresIndividualQuartic:  false,
		NeedsAlgebraRepresentation: true,
		NeedsRealStructure:         true,
		NeedsGrading:               true,
		NeedsOrderOneCalculus:      true,
		NeedsGaugeFluctuationMap:   true,
		OrderOneVerified:           false,
		OrientabilityVerified:      false,
		PoincareDualityVerified:    false,
		KOCompatibilityVerified:    false,
		GaugeKineticMapDerived:     false,
		ActionCoefficientCanonical: false,
		PromotableToSpectralTriple: false,
		BetaRowsAllowed:            0,
		PhysicalConstantsDerived:   false,
		Verdict:                    "exact finite operator candidate, but not a spectral-triple Dirac operator until representation, J, gamma, order-one calculus, and gauge fluctuation map are derived",
	}
}

func spectralActionAnsatzes() []ActionAnsatz {
	return []ActionAnsatz{
		actionAnsatz("finite heat expansion", "Tr f(D/Lambda) = sum_k f_k Tr(D^k)", true),
		actionAnsatz("finite zeta expansion", "sum_s c_s zeta_contact(s)", true),
		actionAnsatz("quartic collective spectral action", "sum_s c_s zeta_q(s)", true),
		actionAnsatz("determinant action", "c log det(D^2)", true),
		actionAnsatz("gauge-fluctuated action", "Tr f((D+A+JAJ^{-1})/Lambda)", true),
	}
}

func actionAnsatz(name, formula string, usesZeta bool) ActionAnsatz {
	return ActionAnsatz{
		Name:                        name,
		Formula:                     formula,
		UsesZetaLedger:              usesZeta,
		ExactOverQ:                  true,
		GaloisInvariant:             true,
		BranchFree:                  true,
		UsesObservedInput:           false,
		UsesBranchChoice:            false,
		RequiresCutoffFunction:      true,
		RequiresTestFunction:        true,
		RequiresDiracOperator:       true,
		RequiresGaugeRepresentation: true,
		CoefficientsCanonical:       false,
		GaugeKineticRows:            0,
		BoundaryConstraintsDerived:  0,
		ThresholdBetaRows:           0,
		PhysicalConstantsDerived:    false,
		Verdict:                     "formal spectral-action ansatz; coefficients and physics interpretation remain non-canonical until the finite spectral triple and gauge map are constructed",
	}
}

func auditAxioms(ingredients []Ingredient) AxiomAudit {
	a := AxiomAudit{IngredientsAudited: len(ingredients), Verdict: "exact spectral pre-data exist, but required canonical spectral-triple axioms are missing"}
	for _, item := range ingredients {
		if item.Available {
			a.AvailableIngredients++
		}
		if item.Canonical {
			a.CanonicalIngredients++
		}
		if item.RequiredForAction {
			a.RequiredIngredients++
			if !item.Available || !item.Canonical {
				a.MissingRequiredCanonical++
			}
		}
	}
	a.FiniteAlgebraAvailable = true
	a.FiniteHilbertCandidate = true
	a.ContactZetaLedgerAvailable = true
	a.ContactOverlapAvailable = true
	a.AlgebraRepresentationReady = false
	a.FiniteDiracSelected = false
	a.RealStructureSelected = false
	a.GradingSelected = false
	a.OrderOneCalculusVerified = false
	a.OrientabilityVerified = false
	a.PoincareDualityVerified = false
	a.KOCompatibilityVerified = false
	a.CanonicalCutoffSelected = false
	a.GaugeFluctuationMapDerived = false
	a.SpectralTripleComplete = false
	return a
}

func auditDirac(candidates []DiracCandidate) DiracAudit {
	a := DiracAudit{CandidatesAudited: len(candidates), Verdict: "all audited finite operator candidates are exact diagnostics, but none is promoted to a spectral-triple Dirac operator"}
	for _, c := range candidates {
		if c.ExactOverQ {
			a.ExactCandidates++
		}
		if c.SelfAdjoint {
			a.SelfAdjointCandidates++
		}
		if c.FiniteSpectrum {
			a.FiniteSpectrumCandidates++
		}
		if c.GaloisInvariant {
			a.GaloisInvariantCandidates++
		}
		if c.BranchFree {
			a.BranchFreeCandidates++
		}
		if c.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if c.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
		if c.NeedsAlgebraRepresentation {
			a.NeedRepresentation++
		}
		if c.NeedsRealStructure {
			a.NeedRealStructure++
		}
		if c.NeedsGrading {
			a.NeedGrading++
		}
		if c.NeedsOrderOneCalculus {
			a.NeedOrderOneCalculus++
		}
		if c.NeedsGaugeFluctuationMap {
			a.NeedGaugeFluctuationMap++
		}
		if c.OrderOneVerified {
			a.OrderOneVerified++
		}
		if c.PromotableToSpectralTriple {
			a.PromotableCandidates++
		}
		if c.GaugeKineticMapDerived {
			a.GaugeKineticMapsDerived++
		}
		if c.ActionCoefficientCanonical {
			a.CanonicalActionCoefficients++
		}
		a.BetaRowsAllowed += c.BetaRowsAllowed
		if c.PhysicalConstantsDerived {
			a.PhysicalConstantsDerived = true
		}
	}
	return a
}

func auditAction(ansatzes []ActionAnsatz) ActionAudit {
	a := ActionAudit{AnsatzesAudited: len(ansatzes), Verdict: "formal spectral-action ansatzes exist, but none has canonical coefficients or a gauge-kinetic interpretation"}
	for _, c := range ansatzes {
		if c.UsesZetaLedger {
			a.UsingZetaLedger++
		}
		if c.ExactOverQ {
			a.ExactAnsatzes++
		}
		if c.GaloisInvariant {
			a.GaloisInvariantAnsatzes++
		}
		if c.BranchFree {
			a.BranchFreeAnsatzes++
		}
		if c.UsesObservedInput {
			a.ObservedInputsUsed = true
		}
		if c.UsesBranchChoice {
			a.BranchChoicesUsed++
		}
		if c.RequiresCutoffFunction {
			a.NeedCutoffFunction++
		}
		if c.RequiresTestFunction {
			a.NeedTestFunction++
		}
		if c.RequiresDiracOperator {
			a.NeedDiracOperator++
		}
		if c.RequiresGaugeRepresentation {
			a.NeedGaugeRepresentation++
		}
		if c.CoefficientsCanonical {
			a.CanonicalCoefficients++
		}
		a.GaugeKineticRows += c.GaugeKineticRows
		a.BoundaryConstraintsDerived += c.BoundaryConstraintsDerived
		a.ThresholdBetaRows += c.ThresholdBetaRows
		if c.PhysicalConstantsDerived {
			a.PhysicalConstantsDerived = true
		}
	}
	return a
}

func FormatIngredient(i Ingredient) string {
	return fmt.Sprintf("%s available=%t canonical=%t required=%t blocks=%t source=%s (%s)", i.Name, i.Available, i.Canonical, i.RequiredForAction, i.BlocksConstruction, i.Source, i.Verdict)
}

func FormatIngredientList(items []Ingredient) string {
	parts := make([]string, 0, len(items))
	for _, item := range items {
		parts = append(parts, fmt.Sprintf("%s:available=%t/canonical=%t", item.Name, item.Available, item.Canonical))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatDiracCandidate(c DiracCandidate) string {
	return fmt.Sprintf("%s: %s exactQ=%t selfAdjoint=%t finite=%t galois=%t branchFree=%t observed=%t branchChoice=%t needs(rep=%t,J=%t,gamma=%t,order1=%t,gauge=%t) order1=%t promotable=%t gaugeMap=%t coeff=%t beta=%d physical=%t (%s)", c.Name, c.Formula, c.ExactOverQ, c.SelfAdjoint, c.FiniteSpectrum, c.GaloisInvariant, c.BranchFree, c.UsesObservedInput, c.UsesBranchChoice, c.NeedsAlgebraRepresentation, c.NeedsRealStructure, c.NeedsGrading, c.NeedsOrderOneCalculus, c.NeedsGaugeFluctuationMap, c.OrderOneVerified, c.PromotableToSpectralTriple, c.GaugeKineticMapDerived, c.ActionCoefficientCanonical, c.BetaRowsAllowed, c.PhysicalConstantsDerived, c.Verdict)
}

func FormatDiracList(candidates []DiracCandidate) string {
	parts := make([]string, 0, len(candidates))
	for _, c := range candidates {
		parts = append(parts, fmt.Sprintf("%s(promotable=%t)", c.Name, c.PromotableToSpectralTriple))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatActionAnsatz(c ActionAnsatz) string {
	return fmt.Sprintf("%s: %s zeta=%t exactQ=%t galois=%t branchFree=%t observed=%t branchChoice=%t needs(cutoff=%t,test=%t,D=%t,gaugeRep=%t) coeff=%t gaugeRows=%d constraints=%d beta=%d physical=%t (%s)", c.Name, c.Formula, c.UsesZetaLedger, c.ExactOverQ, c.GaloisInvariant, c.BranchFree, c.UsesObservedInput, c.UsesBranchChoice, c.RequiresCutoffFunction, c.RequiresTestFunction, c.RequiresDiracOperator, c.RequiresGaugeRepresentation, c.CoefficientsCanonical, c.GaugeKineticRows, c.BoundaryConstraintsDerived, c.ThresholdBetaRows, c.PhysicalConstantsDerived, c.Verdict)
}

func FormatActionAnsatzList(ansatzes []ActionAnsatz) string {
	parts := make([]string, 0, len(ansatzes))
	for _, c := range ansatzes {
		parts = append(parts, fmt.Sprintf("%s(coeff=%t,gaugeRows=%d)", c.Name, c.CoefficientsCanonical, c.GaugeKineticRows))
	}
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func FormatAxiomAudit(a AxiomAudit) string {
	return fmt.Sprintf("ingredients=%d available=%d canonical=%d required=%d missingCanonical=%d algebra=%t H=%t zeta=%t Omega=%t rep=%t D=%t J=%t gamma=%t order1=%t orient=%t duality=%t KO=%t cutoff=%t gauge=%t complete=%t (%s)", a.IngredientsAudited, a.AvailableIngredients, a.CanonicalIngredients, a.RequiredIngredients, a.MissingRequiredCanonical, a.FiniteAlgebraAvailable, a.FiniteHilbertCandidate, a.ContactZetaLedgerAvailable, a.ContactOverlapAvailable, a.AlgebraRepresentationReady, a.FiniteDiracSelected, a.RealStructureSelected, a.GradingSelected, a.OrderOneCalculusVerified, a.OrientabilityVerified, a.PoincareDualityVerified, a.KOCompatibilityVerified, a.CanonicalCutoffSelected, a.GaugeFluctuationMapDerived, a.SpectralTripleComplete, a.Verdict)
}

func FormatDiracAudit(a DiracAudit) string {
	return fmt.Sprintf("candidates=%d exact=%d selfAdjoint=%d finite=%d galois=%d branchFree=%d observed=%t branchChoices=%d need(rep=%d,J=%d,gamma=%d,order1=%d,gauge=%d) order1=%d promotable=%d gaugeMaps=%d coeffs=%d beta=%d physical=%t (%s)", a.CandidatesAudited, a.ExactCandidates, a.SelfAdjointCandidates, a.FiniteSpectrumCandidates, a.GaloisInvariantCandidates, a.BranchFreeCandidates, a.ObservedInputsUsed, a.BranchChoicesUsed, a.NeedRepresentation, a.NeedRealStructure, a.NeedGrading, a.NeedOrderOneCalculus, a.NeedGaugeFluctuationMap, a.OrderOneVerified, a.PromotableCandidates, a.GaugeKineticMapsDerived, a.CanonicalActionCoefficients, a.BetaRowsAllowed, a.PhysicalConstantsDerived, a.Verdict)
}

func FormatActionAudit(a ActionAudit) string {
	return fmt.Sprintf("ansatzes=%d zeta=%d exact=%d galois=%d branchFree=%d observed=%t branchChoices=%d need(cutoff=%d,test=%d,D=%d,gaugeRep=%d) coeffs=%d gaugeRows=%d constraints=%d beta=%d physical=%t (%s)", a.AnsatzesAudited, a.UsingZetaLedger, a.ExactAnsatzes, a.GaloisInvariantAnsatzes, a.BranchFreeAnsatzes, a.ObservedInputsUsed, a.BranchChoicesUsed, a.NeedCutoffFunction, a.NeedTestFunction, a.NeedDiracOperator, a.NeedGaugeRepresentation, a.CanonicalCoefficients, a.GaugeKineticRows, a.BoundaryConstraintsDerived, a.ThresholdBetaRows, a.PhysicalConstantsDerived, a.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate162=%t triple=%t D=%t J=%t gamma=%t cutoff=%t gauge=%t gaugeRows=%d individualQuartic=%d gaugeRep=%d local=%d mass=%d decoupling=%d dynkin=%d beta=%d zero=%d constraints=%d physical=%t closed=%t (%s)", f.Gate162Inherited, f.SpectralTripleComplete, f.FiniteDiracSelected, f.RealStructureSelected, f.GradingSelected, f.CanonicalCutoffSelected, f.GaugeFluctuationMapDerived, f.GaugeKineticMapRows, f.IndividualQuarticRows, f.GaugeRepresentationRows, f.LocalFieldRows, f.MassActivationRows, f.DecouplingRows, f.DynkinIndexRows, f.ThresholdBetaRows, f.ProvenZeroRows, f.BoundaryConstraintsDerived, f.PhysicalConstantsDerived, f.BetaPermissionFirewallClosed, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("contactRows=%d zeta=%d ingredients=%d missingCanonical=%d D=%d promotableD=%d actions=%d coeffs=%d gaugeRows=%d constraints=%d beta=%d nullity=%d→%d", s.ContactRows, s.ZetaValuesComputed, s.IngredientsAudited, s.MissingRequiredCanonical, s.DiracCandidatesAudited, s.PromotableDiracCandidates, s.ActionAnsatzesAudited, s.CanonicalActionCoefficients, s.GaugeKineticRows, s.BoundaryConstraintsDerived, s.ThresholdBetaRows, s.ResidualNullityBefore, s.ResidualNullityAfter)
}
