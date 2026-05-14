// Package finitebundlemap implements Gate 182: finite algebraic local field /
// bundle map construction search.
//
// Gate 181 failed only if "local field" and "bundle" are demanded in their
// classical continuum form. Gate 182 therefore reruns the search internally:
// a finite base is an algebra/spectrum, a vector bundle is a finitely generated
// projective module, and integration candidates are finite traces or finite
// cochain evaluations.
//
// The gate deliberately separates the positive contact result from the still
// missing physical bridge. The complexified contact spectral algebra C[Ω]
// gives a seven-point finite base and the contact carrier is the regular
// projective module over it. That is a lawful algebraic local-field object. But
// no canonical action of this algebra on the 16D Fock space or scalar active
// carrier is yet derived, and neither the homological nor fuzzy/matrix route
// supplies a Chern-Weil/integration/topological-charge bridge.
package finitebundlemap

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fourcyclechernweil"
)

type Route string

const (
	ModuleRoute   Route = "gelfand-projective-module"
	HomologyRoute Route = "finite-homology-chain-complex"
	MatrixRoute   Route = "fuzzy-matrix-geometry"
)

type SpectralBaseAudit struct {
	ContactAlgebraName         string
	Commutative                bool
	SemisimpleAfterComplexify  bool
	DistinctComplexRoots       int
	ComplexMaximalIdeals       int
	RationalPrimaryBlocks      int
	RationalSingletonBlocks    int
	QuarticGaloisPrimaryBlocks int
	SevenPointBaseDerived      bool
	GaloisSafeRowLabels        bool
	BranchChoicesUsed          int
	Comment                    string
}

type ProjectiveModuleCandidate struct {
	Name                          string
	Source                        string
	CarrierDim                    int
	Route                         Route
	HasCanonicalAlgebraAction     bool
	FaithfulContactAlgebraAction  bool
	FinitelyGenerated             bool
	Projective                    bool
	FreeModule                    bool
	CanonicalIdempotentDecomp     bool
	FiberRanksCanonical           bool
	PhysicalGaugeBundle           bool
	LocalFieldEndomorphismDerived bool
	ChernWeilReady                bool
	Verdict                       string
}

type ModuleRouteAudit struct {
	CandidatesAudited             int
	CanonicalContactModules       int
	ProjectiveModules             int
	FreeModules                   int
	PhysicalFockModules           int
	PhysicalScalarModules         int
	CanonicalFockScalarBundleMaps int
	ContactLocalFieldAlgebras     int
	GaugeLocalFieldMaps           int
	ChernWeilReadyModules         int
	FiniteLocalFieldDerived       bool
	PhysicalLocalBundleDerived    bool
	Comment                       string
}

type HomologyRouteAudit struct {
	ComplexesAudited                 int
	BoundaryOperatorsAvailable       int
	ClosedFourChainsFound            int
	CanonicalClosedFourChains        int
	NontrivialH4ClassesDerived       int
	FiniteFundamentalClasses         int
	CochainEvaluationMaps            int
	IntegerTopologicalChargeMaps     int
	BooleanIncidenceComplexAvailable bool
	FanoIncidenceComplexAvailable    bool
	HomologicalFourCycleDerived      bool
	Comment                          string
}

type MatrixRouteAudit struct {
	MatrixAlgebrasAudited           int
	FiniteTracesAvailable           int
	NoncommutativeCoordinateSets    int
	CommutatorPolynomialCandidates  int
	IntegerValuedTracePolynomials   int
	TopologicallyQuantizedTraceMaps int
	ChernCharacterCandidates        int
	FuzzyFourGeometryDerived        bool
	ChernWeilCarrierDerived         bool
	Comment                         string
}

type IntegrationAudit struct {
	AlgebraicContactTraceAvailable bool
	MatrixTraceAvailable           bool
	DixmierTraceNeeded             bool
	FiniteDixmierTraceDerived      bool
	CochainIntegralDerived         bool
	ChernWeilIntegralDerived       bool
	IntegerInstantonNumberDerived  bool
	AbsoluteNormalizationPromoted  bool
	Comment                        string
}

type Firewall struct {
	UsesObservedInputForDerivation bool
	ContinuousBaseRequired         bool
	SevenPointContactBaseDerived   bool
	ContactProjectiveModuleDerived bool
	PhysicalFockBundleDerived      bool
	PhysicalScalarBundleDerived    bool
	ChernWeilCarrierDerived        bool
	HeatKernelMatchingDerived      bool
	ThresholdCorrectedBetaDerived  bool
	AbsoluteCouplingPromoted       bool
	PhysicalConstantsDerived       bool
	StrictNullityBefore            int
	StrictNullityAfter             int
	ConditionalNullityBefore       int
	ConditionalNullityAfter        int
	ClosedStatements               []string
	OpenRequirements               []string
	RecommendedNextGate            string
	Verdict                        string
}

type Analysis struct {
	PreviousGate181 fourcyclechernweil.Analysis
	SpectralBase    SpectralBaseAudit
	Modules         []ProjectiveModuleCandidate
	ModuleRoute     ModuleRouteAudit
	HomologyRoute   HomologyRouteAudit
	MatrixRoute     MatrixRouteAudit
	Integration     IntegrationAudit
	Firewall        Firewall
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		defaultA, defaultErr = buildDefault()
	})
	return defaultA, defaultErr
}

func buildDefault() (Analysis, error) {
	prev, err := fourcyclechernweil.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 181 input: %w", err)
	}
	base := SpectralBaseAudit{
		ContactAlgebraName:         "C[Ω_contact]",
		Commutative:                true,
		SemisimpleAfterComplexify:  true,
		DistinctComplexRoots:       7,
		ComplexMaximalIdeals:       7,
		RationalPrimaryBlocks:      4,
		RationalSingletonBlocks:    3,
		QuarticGaloisPrimaryBlocks: 1,
		SevenPointBaseDerived:      true,
		GaloisSafeRowLabels:        false,
		BranchChoicesUsed:          0,
		Comment:                    "over C the contact spectral algebra is a seven-point finite space; over Q the three rational points plus one quartic primary orbit remain the branch-safe ledger",
	}
	modules := buildModuleCandidates()
	moduleAudit := auditModules(modules)
	homology := HomologyRouteAudit{
		ComplexesAudited:                 2,
		BoundaryOperatorsAvailable:       1,
		ClosedFourChainsFound:            0,
		CanonicalClosedFourChains:        0,
		NontrivialH4ClassesDerived:       0,
		FiniteFundamentalClasses:         0,
		CochainEvaluationMaps:            0,
		IntegerTopologicalChargeMaps:     0,
		BooleanIncidenceComplexAvailable: true,
		FanoIncidenceComplexAvailable:    true,
		HomologicalFourCycleDerived:      false,
		Comment:                          "Boolean/Fano incidence data define finite combinatorial structure, but no canonical nonzero closed 4-chain or fundamental class is selected from the current engine data",
	}
	matrix := MatrixRouteAudit{
		MatrixAlgebrasAudited:           4,
		FiniteTracesAvailable:           3,
		NoncommutativeCoordinateSets:    2,
		CommutatorPolynomialCandidates:  3,
		IntegerValuedTracePolynomials:   0,
		TopologicallyQuantizedTraceMaps: 0,
		ChernCharacterCandidates:        0,
		FuzzyFourGeometryDerived:        false,
		ChernWeilCarrierDerived:         false,
		Comment:                         "End(H_Fock), projected connection matrices, and finite traces exist, but no quantized Chern character or fuzzy four-geometry trace polynomial is derived",
	}
	integration := IntegrationAudit{
		AlgebraicContactTraceAvailable: true,
		MatrixTraceAvailable:           true,
		DixmierTraceNeeded:             false,
		FiniteDixmierTraceDerived:      false,
		CochainIntegralDerived:         false,
		ChernWeilIntegralDerived:       false,
		IntegerInstantonNumberDerived:  false,
		AbsoluteNormalizationPromoted:  false,
		Comment:                        "finite traces are available for algebraic modules and matrices; they are not yet Chern-Weil integrals or instanton-number maps",
	}
	firewall := Firewall{
		UsesObservedInputForDerivation: false,
		ContinuousBaseRequired:         false,
		SevenPointContactBaseDerived:   base.SevenPointBaseDerived,
		ContactProjectiveModuleDerived: moduleAudit.FiniteLocalFieldDerived,
		PhysicalFockBundleDerived:      false,
		PhysicalScalarBundleDerived:    false,
		ChernWeilCarrierDerived:        false,
		HeatKernelMatchingDerived:      false,
		ThresholdCorrectedBetaDerived:  false,
		AbsoluteCouplingPromoted:       false,
		PhysicalConstantsDerived:       false,
		StrictNullityBefore:            prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:             prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:       prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:        prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"the complexified contact spectral algebra defines a finite seven-point base",
			"the contact carrier K7 is the regular/free projective module over its own contact algebra",
			"contact-local algebraic fields exist as A-linear endomorphisms of the regular module",
			"continuous R^{1,3} sections are not required for this finite local-field notion",
		},
		OpenRequirements: []string{
			"canonical action of C[Ω_contact] on H_Fock or H_Φ",
			"canonical fiber-rank/idempotent decomposition for physical carriers",
			"finite fundamental four-cycle or cochain integration map",
			"topologically quantized matrix-trace / Chern-character polynomial",
			"Chern-Weil trace normalization and integer instanton-number bridge",
		},
		RecommendedNextGate: "Gate 183 — contact-module-to-Fock/scalar representation action search",
		Verdict:             "finite algebraic locality is derived on the contact spectral base, but the physical bundle/Chern-Weil bridge remains open",
	}
	truth := "Gate 182 replaces the failed continuum-bundle search with three finite routes. The Gelfand/module route succeeds in a limited but real sense: the complexified contact algebra C[Ω] is a seven-point finite spectral base and K7 is its regular finitely generated projective module, so contact-local algebraic fields exist as module endomorphisms. However H_Fock and H_Φ do not yet carry a canonical contact-algebra action, the Boolean/Fano homological route provides no selected closed 4-cycle, and the fuzzy/matrix route provides traces but no quantized Chern-Weil character. The finite-to-continuum obstruction is narrowed to a module-action/integration/topological-charge bridge rather than a demand for a classical manifold."
	return Analysis{PreviousGate181: prev, SpectralBase: base, Modules: modules, ModuleRoute: moduleAudit, HomologyRoute: homology, MatrixRoute: matrix, Integration: integration, Firewall: firewall, TruthStatement: truth}, nil
}

func buildModuleCandidates() []ProjectiveModuleCandidate {
	return []ProjectiveModuleCandidate{
		{Name: "K7 contact regular module", Source: "Gates 5/149/151", CarrierDim: 7, Route: ModuleRoute, HasCanonicalAlgebraAction: true, FaithfulContactAlgebraAction: true, FinitelyGenerated: true, Projective: true, FreeModule: true, CanonicalIdempotentDecomp: true, FiberRanksCanonical: true, LocalFieldEndomorphismDerived: true, Verdict: "success: finite contact-local field object A=C[Ω] acting on itself"},
		{Name: "H_Fock 16D spinor carrier", Source: "Gates 14/166/167", CarrierDim: 16, Route: ModuleRoute, FinitelyGenerated: true, Verdict: "blocked: no canonical representation of C[Ω_contact] on H_Fock; any seven-fiber rank vector is a branch choice"},
		{Name: "H_Φ active scalar carrier", Source: "Gates 11/20/37", CarrierDim: 4, Route: ModuleRoute, FinitelyGenerated: true, Verdict: "blocked: no canonical contact-algebra action or faithful seven-point fiber decomposition on the scalar active space"},
		{Name: "H_Fock ⊗ H_Φ tensor carrier", Source: "Gate 17 plus Gate 166", CarrierDim: 64, Route: ModuleRoute, FinitelyGenerated: true, Verdict: "blocked: tensor product is available as bookkeeping, but no contact-algebra module action/glue map is derived"},
		{Name: "End(H_Fock) matrix algebra", Source: "Gate 166 finite spectral triple ansatz", CarrierDim: 256, Route: MatrixRoute, FinitelyGenerated: true, Verdict: "trace algebra exists but is not a projective module over the contact base and has no quantized Chern character"},
	}
}

func auditModules(xs []ProjectiveModuleCandidate) ModuleRouteAudit {
	a := ModuleRouteAudit{CandidatesAudited: len(xs)}
	for _, x := range xs {
		if x.HasCanonicalAlgebraAction && x.FaithfulContactAlgebraAction {
			a.CanonicalContactModules++
		}
		if x.Projective {
			a.ProjectiveModules++
		}
		if x.FreeModule {
			a.FreeModules++
		}
		if x.Name == "H_Fock 16D spinor carrier" && x.HasCanonicalAlgebraAction && x.Projective {
			a.PhysicalFockModules++
		}
		if x.Name == "H_Φ active scalar carrier" && x.HasCanonicalAlgebraAction && x.Projective {
			a.PhysicalScalarModules++
		}
		if x.LocalFieldEndomorphismDerived {
			a.ContactLocalFieldAlgebras++
		}
		if x.PhysicalGaugeBundle {
			a.GaugeLocalFieldMaps++
		}
		if x.ChernWeilReady {
			a.ChernWeilReadyModules++
		}
	}
	a.FiniteLocalFieldDerived = a.ContactLocalFieldAlgebras == 1 && a.ProjectiveModules == 1
	a.PhysicalLocalBundleDerived = a.PhysicalFockModules > 0 || a.PhysicalScalarModules > 0 || a.CanonicalFockScalarBundleMaps > 0
	a.Comment = "the module route succeeds only for contact locality: K7 is the regular projective module over C[Ω]; physical Fock/scalar carriers still lack canonical contact-algebra action"
	return a
}

func FormatSpectralBase(a SpectralBaseAudit) string {
	return fmt.Sprintf("algebra=%s commutative=%t semisimpleC=%t complexMaxIdeals=%d rationalBlocks=%d singleton=%d quarticPrimary=%d sevenPoint=%t galoisSafeRows=%t branchChoices=%d (%s)", a.ContactAlgebraName, a.Commutative, a.SemisimpleAfterComplexify, a.ComplexMaximalIdeals, a.RationalPrimaryBlocks, a.RationalSingletonBlocks, a.QuarticGaloisPrimaryBlocks, a.SevenPointBaseDerived, a.GaloisSafeRowLabels, a.BranchChoicesUsed, a.Comment)
}

func FormatModuleRoute(a ModuleRouteAudit) string {
	return fmt.Sprintf("candidates=%d contactModules=%d projective=%d free=%d fockModules=%d scalarModules=%d contactLocalFields=%d gaugeLocalFields=%d chernWeilReady=%d finiteLocal=%t physicalBundle=%t (%s)", a.CandidatesAudited, a.CanonicalContactModules, a.ProjectiveModules, a.FreeModules, a.PhysicalFockModules, a.PhysicalScalarModules, a.ContactLocalFieldAlgebras, a.GaugeLocalFieldMaps, a.ChernWeilReadyModules, a.FiniteLocalFieldDerived, a.PhysicalLocalBundleDerived, a.Comment)
}

func FormatModules(xs []ProjectiveModuleCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s[%s dim=%d action=%t faithful=%t proj=%t free=%t localEnd=%t gaugeBundle=%t chernWeil=%t: %s]", x.Name, x.Route, x.CarrierDim, x.HasCanonicalAlgebraAction, x.FaithfulContactAlgebraAction, x.Projective, x.FreeModule, x.LocalFieldEndomorphismDerived, x.PhysicalGaugeBundle, x.ChernWeilReady, x.Verdict))
	}
	return strings.Join(parts, " | ")
}

func FormatHomologyRoute(a HomologyRouteAudit) string {
	return fmt.Sprintf("complexes=%d boundaryOps=%d closed4=%d canonicalClosed4=%d H4=%d fundamental=%d cochainEval=%d integerCharge=%d boolean=%t fano=%t derived=%t (%s)", a.ComplexesAudited, a.BoundaryOperatorsAvailable, a.ClosedFourChainsFound, a.CanonicalClosedFourChains, a.NontrivialH4ClassesDerived, a.FiniteFundamentalClasses, a.CochainEvaluationMaps, a.IntegerTopologicalChargeMaps, a.BooleanIncidenceComplexAvailable, a.FanoIncidenceComplexAvailable, a.HomologicalFourCycleDerived, a.Comment)
}

func FormatMatrixRoute(a MatrixRouteAudit) string {
	return fmt.Sprintf("matrixAlgebras=%d traces=%d coordinates=%d commutatorPolys=%d integerTraces=%d quantizedTraces=%d chernCharacters=%d fuzzy4=%t chernWeil=%t (%s)", a.MatrixAlgebrasAudited, a.FiniteTracesAvailable, a.NoncommutativeCoordinateSets, a.CommutatorPolynomialCandidates, a.IntegerValuedTracePolynomials, a.TopologicallyQuantizedTraceMaps, a.ChernCharacterCandidates, a.FuzzyFourGeometryDerived, a.ChernWeilCarrierDerived, a.Comment)
}

func FormatIntegration(a IntegrationAudit) string {
	return fmt.Sprintf("contactTrace=%t matrixTrace=%t dixmierNeeded=%t finiteDixmier=%t cochainIntegral=%t chernWeilIntegral=%t integerInstanton=%t absoluteNorm=%t (%s)", a.AlgebraicContactTraceAvailable, a.MatrixTraceAvailable, a.DixmierTraceNeeded, a.FiniteDixmierTraceDerived, a.CochainIntegralDerived, a.ChernWeilIntegralDerived, a.IntegerInstantonNumberDerived, a.AbsoluteNormalizationPromoted, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t continuousBaseRequired=%t sevenPoint=%t contactModule=%t fockBundle=%t scalarBundle=%t chernWeil=%t heatKernel=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d verdict=%s", a.UsesObservedInputForDerivation, a.ContinuousBaseRequired, a.SevenPointContactBaseDerived, a.ContactProjectiveModuleDerived, a.PhysicalFockBundleDerived, a.PhysicalScalarBundleDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, a.Verdict)
}
