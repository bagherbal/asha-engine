// Package contactmoduleaction implements Gate 183: contact-module to
// Fock/scalar representation action search.
//
// Gate 182 derived finite algebraic locality on the contact carrier itself:
// the complexified contact spectral algebra C[Ω_contact] is a seven-point
// finite base and K7 is its regular projective module. Gate 183 asks whether
// that contact-local base acts canonically on the physical Fock/spinor carrier
// H_Fock or on the active scalar carrier H_Φ.
//
// The gate intentionally avoids arbitrary maps C^7 -> M_16(C) or C^7 -> M_4(C).
// It audits only geometrically constrained routes already suggested by the
// engine: Clifford multiplication, the quartic primary ideal, and the projected
// finite connection. The result is a partial positive theorem: several exact
// pre-actions exist, but none yet becomes a canonical physical module action of
// C[Ω_contact] on H_Fock or H_Φ.
package contactmoduleaction

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitebundlemap"
)

type Route string

const (
	CliffordSpinorRoute Route = "clifford-spinor-action"
	QuarticScalarRoute  Route = "quartic-scalar-ideal"
	ConnectionRoute     Route = "connection-induced-action"
)

type CliffordSpinorAudit struct {
	Route                               Route
	ContactCarrierDim                   int
	FockSpinorDim                       int
	CliffordBookkeepingAvailable        bool
	K7VectorActionCanonical             bool
	ActionOnSpinorsCanonical            bool
	OddChiralityExchangeCompatible      bool
	LinearK7ToEndFockMapDerived         bool
	MultiplicativeContactAlgebraHom     bool
	CommutativeSpectralIdempotentAction bool
	OmegaIntertwiningLawDerived         bool
	RequiresContactEigenvectorBranch    bool
	InducesFockProjectiveModule         bool
	InducesPhysicalSpinorBundle         bool
	Verdict                             string
}

type QuarticScalarIdealAudit struct {
	Route                            Route
	QuarticPrimaryDim                int
	ScalarCarrierDim                 int
	QuarticPolynomial                string
	GaloisSafePrimaryIdeal           bool
	AbstractRankOneModuleOverQuartic bool
	CompanionRepresentationAvailable bool
	BranchFreeQuarticBlock           bool
	ScalarOperatorWithQuarticMinimal bool
	CanonicalHphiIdentification      bool
	ProjectiveScalarModuleDerived    bool
	PhysicalScalarBundleDerived      bool
	Verdict                          string
}

type ConnectionInducedActionAudit struct {
	Route                               Route
	ProjectedConnectionAvailable        bool
	OffDiagonalBlockConnectionAvailable bool
	SecondFundamentalCurvatureAvailable bool
	CompressedConnectionCanonical       bool
	AdjointActionCandidate              bool
	CommutatorActionCandidate           bool
	ClosesOnContactSpectralAlgebra      bool
	PullbackToFockDerived               bool
	PullbackToScalarDerived             bool
	FockDiracCommutatorClosed           bool
	GaugeCovariantModuleActionDerived   bool
	Verdict                             string
}

type ActionCandidate struct {
	Name                   string
	Route                  Route
	Domain                 string
	TargetCarrier          string
	TargetDim              int
	CanonicalPredata       bool
	BranchFree             bool
	AlgebraHomomorphism    bool
	ProjectiveModuleAction bool
	PhysicalCarrierAction  bool
	ChernWeilReady         bool
	Verdict                string
}

type RouteSummary struct {
	CandidatesAudited          int
	CanonicalPreactions        int
	BranchFreeCandidates       int
	AlgebraHomomorphisms       int
	ProjectiveModuleActions    int
	PhysicalFockActions        int
	PhysicalScalarActions      int
	ChernWeilReadyActions      int
	ArbitraryMapsUsed          int
	CompletePhysicalBundleMaps int
	Comment                    string
}

type Firewall struct {
	UsesObservedInputForDerivation     bool
	ArbitraryLinearMapUsed             bool
	ContactBaseInherited               bool
	ContactRegularModuleInherited      bool
	CliffordSpinorPreactionDerived     bool
	QuarticAbstractScalarModuleDerived bool
	ConnectionPreactionAudited         bool
	CanonicalFockActionDerived         bool
	CanonicalScalarActionDerived       bool
	PhysicalBundleMapDerived           bool
	ChernWeilCarrierDerived            bool
	HeatKernelMatchingDerived          bool
	ThresholdCorrectedBetaDerived      bool
	AbsoluteCouplingPromoted           bool
	PhysicalConstantsDerived           bool
	StrictNullityBefore                int
	StrictNullityAfter                 int
	ConditionalNullityBefore           int
	ConditionalNullityAfter            int
	ClosedStatements                   []string
	OpenRequirements                   []string
	RecommendedNextGate                string
	Verdict                            string
}

type Analysis struct {
	PreviousGate182 finitebundlemap.Analysis
	CliffordSpinor  CliffordSpinorAudit
	QuarticScalar   QuarticScalarIdealAudit
	Connection      ConnectionInducedActionAudit
	Candidates      []ActionCandidate
	Summary         RouteSummary
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
	prev, err := finitebundlemap.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 182 input: %w", err)
	}
	cliff := CliffordSpinorAudit{
		Route:                               CliffordSpinorRoute,
		ContactCarrierDim:                   7,
		FockSpinorDim:                       16,
		CliffordBookkeepingAvailable:        true,
		K7VectorActionCanonical:             true,
		ActionOnSpinorsCanonical:            true,
		OddChiralityExchangeCompatible:      true,
		LinearK7ToEndFockMapDerived:         true,
		MultiplicativeContactAlgebraHom:     false,
		CommutativeSpectralIdempotentAction: false,
		OmegaIntertwiningLawDerived:         false,
		RequiresContactEigenvectorBranch:    true,
		InducesFockProjectiveModule:         false,
		InducesPhysicalSpinorBundle:         false,
		Verdict:                             "Clifford multiplication gives a canonical K7-vector action on the 16D spinor/Fock carrier, but vector Clifford action is not a multiplicative representation of the commutative contact spectral algebra C[Ω] and does not provide spectral idempotent fibers on H_Fock.",
	}
	quartic := QuarticScalarIdealAudit{
		Route:                            QuarticScalarRoute,
		QuarticPrimaryDim:                4,
		ScalarCarrierDim:                 4,
		QuarticPolynomial:                "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		GaloisSafePrimaryIdeal:           true,
		AbstractRankOneModuleOverQuartic: true,
		CompanionRepresentationAvailable: true,
		BranchFreeQuarticBlock:           true,
		ScalarOperatorWithQuarticMinimal: false,
		CanonicalHphiIdentification:      false,
		ProjectiveScalarModuleDerived:    false,
		PhysicalScalarBundleDerived:      false,
		Verdict:                          "The quartic primary ideal is a canonical 4D Galois-safe algebra and has an abstract rank-one module/companion representation, but H_Φ is not canonically identified with that ideal because no scalar operator on H_Φ has the quartic minimal polynomial or contact-idempotent action.",
	}
	conn := ConnectionInducedActionAudit{
		Route:                               ConnectionRoute,
		ProjectedConnectionAvailable:        true,
		OffDiagonalBlockConnectionAvailable: true,
		SecondFundamentalCurvatureAvailable: true,
		CompressedConnectionCanonical:       true,
		AdjointActionCandidate:              true,
		CommutatorActionCandidate:           true,
		ClosesOnContactSpectralAlgebra:      false,
		PullbackToFockDerived:               false,
		PullbackToScalarDerived:             false,
		FockDiracCommutatorClosed:           false,
		GaugeCovariantModuleActionDerived:   false,
		Verdict:                             "The Gate-11 connection supplies canonical block/off-diagonal predata and second-fundamental curvature, but its commutator/adjoint actions do not close as a C[Ω] module action on H_Fock or H_Φ.",
	}
	candidates := buildCandidates(cliff, quartic, conn)
	summary := auditCandidates(candidates)
	firewall := Firewall{
		UsesObservedInputForDerivation:     false,
		ArbitraryLinearMapUsed:             false,
		ContactBaseInherited:               prev.Firewall.SevenPointContactBaseDerived,
		ContactRegularModuleInherited:      prev.Firewall.ContactProjectiveModuleDerived,
		CliffordSpinorPreactionDerived:     cliff.LinearK7ToEndFockMapDerived,
		QuarticAbstractScalarModuleDerived: quartic.AbstractRankOneModuleOverQuartic,
		ConnectionPreactionAudited:         conn.ProjectedConnectionAvailable,
		CanonicalFockActionDerived:         false,
		CanonicalScalarActionDerived:       false,
		PhysicalBundleMapDerived:           false,
		ChernWeilCarrierDerived:            false,
		HeatKernelMatchingDerived:          false,
		ThresholdCorrectedBetaDerived:      false,
		AbsoluteCouplingPromoted:           false,
		PhysicalConstantsDerived:           false,
		StrictNullityBefore:                prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                 prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:           prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:            prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"arbitrary maps C^7 -> End(H_Fock) and C^7 -> End(H_Φ) are not used",
			"K7 has a canonical Clifford vector action on the 16D Fock/spinor carrier",
			"the quartic contact primary ideal is a canonical 4D Galois-safe abstract module candidate",
			"the projected connection is valid predata but not yet a C[Ω]-module pullback",
		},
		OpenRequirements: []string{
			"multiplicative representation of C[Ω_contact] on H_Fock",
			"spectral idempotent/fiber decomposition of H_Fock compatible with Clifford action",
			"scalar operator on H_Φ with the quartic contact minimal polynomial or equivalent canonical ideal action",
			"connection-induced commutator/adjoint action closing on the contact spectral algebra",
			"Chern-Weil-ready physical bundle map with integration/topological-charge pairing",
		},
		RecommendedNextGate: "Gate 184 — Clifford-contact spectral idempotent/commutant obstruction or construction",
		Verdict:             "finite module-action predata exists, but no canonical physical C[Ω_contact] action on H_Fock or H_Φ is derived",
	}
	truth := "Gate 183 tests three constrained bridges from the contact spectral base to physical carriers. Clifford multiplication gives a real canonical K7-vector action on the 16D Fock/spinor space, and the quartic primary ideal gives a canonical 4D abstract scalar-module candidate. However neither object is yet a multiplicative C[Ω_contact] representation on H_Fock or H_Φ, and the projected connection does not close as a contact-algebra pullback. The finite-bundle obstruction is therefore narrowed from arbitrary dimensional mismatch to a precise module-action problem: derive a contact spectral idempotent action on the spinor/scalar carriers, or prove that no such action can be canonical."
	return Analysis{PreviousGate182: prev, CliffordSpinor: cliff, QuarticScalar: quartic, Connection: conn, Candidates: candidates, Summary: summary, Firewall: firewall, TruthStatement: truth}, nil
}

func buildCandidates(cliff CliffordSpinorAudit, quartic QuarticScalarIdealAudit, conn ConnectionInducedActionAudit) []ActionCandidate {
	return []ActionCandidate{
		{Name: "Clifford K7 vector action on H_Fock", Route: cliff.Route, Domain: "K7 ⊂ Cl(1,7)", TargetCarrier: "H_Fock", TargetDim: cliff.FockSpinorDim, CanonicalPredata: cliff.LinearK7ToEndFockMapDerived, BranchFree: false, AlgebraHomomorphism: cliff.MultiplicativeContactAlgebraHom, ProjectiveModuleAction: cliff.InducesFockProjectiveModule, PhysicalCarrierAction: cliff.InducesPhysicalSpinorBundle, Verdict: cliff.Verdict},
		{Name: "Contact spectral idempotents on H_Fock", Route: cliff.Route, Domain: "C[Ω_contact]", TargetCarrier: "H_Fock", TargetDim: cliff.FockSpinorDim, CanonicalPredata: cliff.K7VectorActionCanonical, BranchFree: false, AlgebraHomomorphism: false, ProjectiveModuleAction: false, PhysicalCarrierAction: false, Verdict: "blocked: no branch-free spectral-idempotent fiber decomposition of the 16D spinor carrier is derived"},
		{Name: "Quartic primary ideal abstract module", Route: quartic.Route, Domain: "Q[x]/q_4(x)", TargetCarrier: "abstract 4D module", TargetDim: quartic.QuarticPrimaryDim, CanonicalPredata: quartic.AbstractRankOneModuleOverQuartic, BranchFree: quartic.BranchFreeQuarticBlock, AlgebraHomomorphism: quartic.CompanionRepresentationAvailable, ProjectiveModuleAction: true, PhysicalCarrierAction: false, Verdict: "success only abstractly: the quartic ideal has a branch-free rank-one module, but this is not yet H_Φ"},
		{Name: "Quartic primary ideal action on H_Φ", Route: quartic.Route, Domain: "Q[x]/q_4(x)", TargetCarrier: "H_Φ", TargetDim: quartic.ScalarCarrierDim, CanonicalPredata: quartic.GaloisSafePrimaryIdeal, BranchFree: quartic.BranchFreeQuarticBlock, AlgebraHomomorphism: false, ProjectiveModuleAction: quartic.ProjectiveScalarModuleDerived, PhysicalCarrierAction: quartic.PhysicalScalarBundleDerived, Verdict: quartic.Verdict},
		{Name: "Gate-11 connection adjoint/commutator pullback", Route: conn.Route, Domain: "projected connection algebra", TargetCarrier: "H_Fock ⊕ H_Φ", TargetDim: 20, CanonicalPredata: conn.ProjectedConnectionAvailable, BranchFree: true, AlgebraHomomorphism: conn.ClosesOnContactSpectralAlgebra, ProjectiveModuleAction: false, PhysicalCarrierAction: conn.GaugeCovariantModuleActionDerived, Verdict: conn.Verdict},
	}
}

func auditCandidates(xs []ActionCandidate) RouteSummary {
	s := RouteSummary{CandidatesAudited: len(xs)}
	for _, x := range xs {
		if x.CanonicalPredata {
			s.CanonicalPreactions++
		}
		if x.BranchFree {
			s.BranchFreeCandidates++
		}
		if x.AlgebraHomomorphism {
			s.AlgebraHomomorphisms++
		}
		if x.ProjectiveModuleAction {
			s.ProjectiveModuleActions++
		}
		if x.TargetCarrier == "H_Fock" && x.PhysicalCarrierAction {
			s.PhysicalFockActions++
		}
		if x.TargetCarrier == "H_Φ" && x.PhysicalCarrierAction {
			s.PhysicalScalarActions++
		}
		if x.ChernWeilReady {
			s.ChernWeilReadyActions++
		}
	}
	s.CompletePhysicalBundleMaps = s.PhysicalFockActions + s.PhysicalScalarActions
	s.Comment = "constrained pre-actions exist, and the quartic ideal has an abstract projective module, but no physical C[Ω]-module action on H_Fock or H_Φ is selected"
	return s
}

func FormatCliffordSpinor(a CliffordSpinorAudit) string {
	return fmt.Sprintf("route=%s Kdim=%d FockDim=%d cl=%t kAction=%t spinAction=%t oddGamma=%t linearMap=%t algebraHom=%t idempotents=%t omegaIntertwine=%t branch=%t fockModule=%t spinorBundle=%t (%s)", a.Route, a.ContactCarrierDim, a.FockSpinorDim, a.CliffordBookkeepingAvailable, a.K7VectorActionCanonical, a.ActionOnSpinorsCanonical, a.OddChiralityExchangeCompatible, a.LinearK7ToEndFockMapDerived, a.MultiplicativeContactAlgebraHom, a.CommutativeSpectralIdempotentAction, a.OmegaIntertwiningLawDerived, a.RequiresContactEigenvectorBranch, a.InducesFockProjectiveModule, a.InducesPhysicalSpinorBundle, a.Verdict)
}

func FormatQuarticScalar(a QuarticScalarIdealAudit) string {
	return fmt.Sprintf("route=%s qdim=%d hphi=%d q=%q galois=%t abstractModule=%t companion=%t branchFree=%t scalarMinPoly=%t hphiId=%t scalarModule=%t physicalBundle=%t (%s)", a.Route, a.QuarticPrimaryDim, a.ScalarCarrierDim, a.QuarticPolynomial, a.GaloisSafePrimaryIdeal, a.AbstractRankOneModuleOverQuartic, a.CompanionRepresentationAvailable, a.BranchFreeQuarticBlock, a.ScalarOperatorWithQuarticMinimal, a.CanonicalHphiIdentification, a.ProjectiveScalarModuleDerived, a.PhysicalScalarBundleDerived, a.Verdict)
}

func FormatConnectionAction(a ConnectionInducedActionAudit) string {
	return fmt.Sprintf("route=%s projected=%t offdiag=%t secondFund=%t compressed=%t adjoint=%t commutator=%t closesCΩ=%t pullFock=%t pullScalar=%t diracClosed=%t gaugeModule=%t (%s)", a.Route, a.ProjectedConnectionAvailable, a.OffDiagonalBlockConnectionAvailable, a.SecondFundamentalCurvatureAvailable, a.CompressedConnectionCanonical, a.AdjointActionCandidate, a.CommutatorActionCandidate, a.ClosesOnContactSpectralAlgebra, a.PullbackToFockDerived, a.PullbackToScalarDerived, a.FockDiracCommutatorClosed, a.GaugeCovariantModuleActionDerived, a.Verdict)
}

func FormatCandidate(a ActionCandidate) string {
	return fmt.Sprintf("%s[%s: %s -> %s dim=%d predata=%t branchFree=%t hom=%t module=%t physical=%t chernWeil=%t: %s]", a.Name, a.Route, a.Domain, a.TargetCarrier, a.TargetDim, a.CanonicalPredata, a.BranchFree, a.AlgebraHomomorphism, a.ProjectiveModuleAction, a.PhysicalCarrierAction, a.ChernWeilReady, a.Verdict)
}

func FormatCandidates(xs []ActionCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCandidate(x))
	}
	return strings.Join(parts, " | ")
}

func FormatSummary(a RouteSummary) string {
	return fmt.Sprintf("candidates=%d preactions=%d branchFree=%d hom=%d projectiveModules=%d fockActions=%d scalarActions=%d chernWeilReady=%d arbitraryMaps=%d completePhysical=%d (%s)", a.CandidatesAudited, a.CanonicalPreactions, a.BranchFreeCandidates, a.AlgebraHomomorphisms, a.ProjectiveModuleActions, a.PhysicalFockActions, a.PhysicalScalarActions, a.ChernWeilReadyActions, a.ArbitraryMapsUsed, a.CompletePhysicalBundleMaps, a.Comment)
}

func FormatFirewall(a Firewall) string {
	return fmt.Sprintf("observed=%t arbitraryMap=%t base=%t regularModule=%t cliffordPre=%t quarticAbstract=%t connection=%t fockAction=%t scalarAction=%t bundle=%t chernWeil=%t heat=%t thresholds=%t absolute=%t constants=%t strict=%d->%d conditional=%d->%d verdict=%s", a.UsesObservedInputForDerivation, a.ArbitraryLinearMapUsed, a.ContactBaseInherited, a.ContactRegularModuleInherited, a.CliffordSpinorPreactionDerived, a.QuarticAbstractScalarModuleDerived, a.ConnectionPreactionAudited, a.CanonicalFockActionDerived, a.CanonicalScalarActionDerived, a.PhysicalBundleMapDerived, a.ChernWeilCarrierDerived, a.HeatKernelMatchingDerived, a.ThresholdCorrectedBetaDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, a.Verdict)
}
