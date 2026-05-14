// Package canonicalfinitediracselector implements Gate 269:
// Canonical Finite Dirac Selector / Order-One Spectral Triple Completion Audit.
//
// Gate 268 showed that raw spectral moments of the formal odd self-adjoint
// finite Dirac family D_F(M) depend on the unselected singular spectrum of M.
// Gate 269 tests whether the Noncommutative Geometry order-one condition can
// canonically select M from the currently derived data.
//
// The gate deliberately separates two levels. On the mode-level native
// C ⊕ M3(C) preflight, a toy order-one calculation can be performed and it
// reduces a generic 4×4 complex block M to the block-commutant form
// diag(x, y I3). That is real progress: temporal/spatial leakage and internal
// color anisotropy are removed. But this is not yet a physical finite Dirac
// theorem. The project still lacks a faithful representation on the doubled
// S_C carrier, a physical opposite action through J, and a non-vacuous one-form
// calculus. Moreover the surviving two-parameter family has variable raw
// Tr(D²)/Tr(D⁴), so no canonical D_F or Higgs ratio is derived.
package canonicalfinitediracselector

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitespectralactionreattempt"
)

const (
	AuditID = "GATE269-CANONICAL-FINITE-DIRAC-ORDER-ONE-SELECTOR-AUDIT"

	StatusGate268Inherited          = "CONDITIONAL_SUPPORT_GATE268_SPECTRAL_ACTION_REATTEMPT_INHERITED"
	StatusOrderOneDefined           = "CONDITIONAL_SUPPORT_ORDER_ONE_CONDITION_FORMALLY_DEFINED"
	StatusModeLevelAlgebraPreflight = "CONDITIONAL_SUPPORT_MODE_LEVEL_C_PLUS_M3C_ORDER_ONE_PREFLIGHT"
	StatusOrderOneSieveReduced      = "CONDITIONAL_SUPPORT_ORDER_ONE_SIEVE_REDUCES_GENERIC_M"
	StatusAllowedMomentsReevaluated = "CONDITIONAL_SUPPORT_ORDER_ONE_ALLOWED_MOMENTS_REEVALUATED"
	StatusFailedFaithfulRep         = "FAILED_ROUTE_FAITHFUL_TOTAL_SC_ALGEBRA_REPRESENTATION_MISSING"
	StatusFailedOppositeAction      = "FAILED_ROUTE_PHYSICAL_OPPOSITE_ALGEBRA_ACTION_MISSING"
	StatusFailedNonVacuousCalculus  = "FAILED_ROUTE_NON_VACUOUS_ORDER_ONE_CALCULUS_NOT_DERIVED"
	StatusFailedCanonicalDF         = "FAILED_ROUTE_ORDER_ONE_DOES_NOT_SELECT_UNIQUE_CANONICAL_DF"
	StatusFailedMomentRatio         = "FAILED_ROUTE_ORDER_ONE_ALLOWED_TRACE_RATIO_STILL_AMPLITUDE_DEPENDENT"
	StatusFailedHiggsRatio          = "FAILED_ROUTE_HIGGS_RATIO_STILL_NOT_DERIVED"
	StatusEmpiricalSealPreserved    = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate268Inheritance struct {
	ScaffoldRetrieved       bool
	FormalDFFamilyAvailable bool
	RawMomentsEvaluated     bool
	MomentDependenceExposed bool
	CanonicalDFDerived      bool
	HiggsRatioDerived       bool
	FirewallPreserved       bool
	RecommendedNextGate     string
	Verdict                 string
}

type OrderOneDefinition struct {
	Formula                string
	AlgebraSymbol          string
	DiracSymbol            string
	RealStructureSymbol    string
	RequiresRepresentation bool
	RequiresOppositeAction bool
	RequiresAllAAndB       bool
	Defined                bool
	Verdict                string
}

type AlgebraRepresentationAudit struct {
	NativeAlgebraName             string
	ModeCarrier                   string
	ModeDimension                 int
	ModeLevelCPlusM3Available     bool
	FullSCRepresentationDerived   bool
	LeftRepresentationDerived     bool
	RightRepresentationDerived    bool
	OppositeRepresentationDerived bool
	PhysicalJDerived              bool
	NonVacuousOneFormsAvailable   bool
	ToyModePreflightAllowed       bool
	ImportedConnesAlgebra         bool
	Verdict                       string
}

type GenericDiracBlockAudit struct {
	Carrier                        string
	GenericMatrixShape             string
	InitialComplexParameters       int
	InitialRealParameters          int
	OrderOneToyConstraint          string
	AllowedComplexParameters       int
	AllowedRealParameters          int
	EliminatedComplexParameters    int
	TemporalSpatialLeakageRemoved  bool
	ColorAnisotropyRemoved         bool
	AllowedFamilyFormula           string
	SieveNontrivial                bool
	SievePhysicalOnFullSC          bool
	OneFormsVanishForAllowedFamily bool
	CanonicalBlockSelected         bool
	Verdict                        string
}

type ConstraintRow struct {
	Name      string
	Before    string
	After     string
	Reason    string
	Satisfied bool
}

type MomentRow struct {
	Name            string
	X               float64
	Y               float64
	SingularValues  []float64
	TraceD2         float64
	TraceD4         float64
	RawRatio        float64
	OrderOneAllowed bool
	Canonical       bool
	Comment         string
}

type SpectralMomentReevaluation struct {
	Rows                          []MomentRow
	MomentsRecomputed             bool
	AllRowsOrderOneAllowed        bool
	RawRatioStableAcrossAllowedDF bool
	DependsOnSurvivingAmplitudes  bool
	SeeleyDeWittMapDerived        bool
	HiggsRatioDerived             bool
	Verdict                       string
}

type CanonicalDFVerdict struct {
	OrderOneSieveNontrivial       bool
	UniqueDFSelected              bool
	SurvivingFamilyDimensionC     int
	RequiresAdditionalSelector    bool
	CouldUseNormalizationOnly     bool
	NormalizationDerived          bool
	GaugeProjectionDerived        bool
	ScalarFluctuationMapDerived   bool
	PromotableFiniteDiracOperator bool
	Verdict                       string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	NoConnesAlgebraImported         bool
	NoYukawaFitUsed                 bool
	ToySieveNotPromoted             bool
	NoHiggsPredictionClaim          bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureObligation struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Obligations             []FutureObligation
	NeedFaithfulSCRep       bool
	NeedPhysicalOppositeJ   bool
	NeedNonVacuousOneForms  bool
	NeedQuaternionicOrWeakH bool
	NeedCanonicalAmplitude  bool
	NeedHeatKernelMap       bool
	RecommendedNextGate     string
	Verdict                 string
}

type Summary struct {
	Gate268Inherited         bool
	OrderOneDefined          bool
	ModeAlgebraPreflight     bool
	OrderOneSieveReduced     bool
	CanonicalDFDerived       bool
	AllowedMomentRatioStable bool
	HiggsRatioDerived        bool
	FirewallPreserved        bool
	Status                   string
	NextGate                 string
	Comment                  string
}

type Analysis struct {
	PreviousGate268 finitespectralactionreattempt.Analysis
	Inheritance     Gate268Inheritance
	Definition      OrderOneDefinition
	Algebra         AlgebraRepresentationAudit
	Constraints     []ConstraintRow
	Sieve           GenericDiracBlockAudit
	Moments         SpectralMomentReevaluation
	Canonical       CanonicalDFVerdict
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
		prev, err := finitespectralactionreattempt.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 268 predecessor: %w", err)
			return
		}
		inh := inheritGate268(prev)
		def := defineOrderOne()
		alg := auditAlgebraRepresentation()
		constraints, sieve := applyModeLevelOrderOneSieve(def, alg)
		moments := reevaluateAllowedMoments(sieve)
		canonical := auditCanonicalDF(sieve, moments)
		firewall := auditFirewall(inh, alg, sieve, canonical)
		future := defineFutureMap(alg, sieve, moments, canonical)
		summary := summarize(inh, def, alg, sieve, moments, canonical, firewall, future)
		truth := buildTruth(inh, alg, sieve, moments, canonical, firewall)
		defaultA = Analysis{PreviousGate268: prev, Inheritance: inh, Definition: def, Algebra: alg, Constraints: constraints, Sieve: sieve, Moments: moments, Canonical: canonical, Firewall: firewall, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate268(prev finitespectralactionreattempt.Analysis) Gate268Inheritance {
	return Gate268Inheritance{
		ScaffoldRetrieved:       prev.Summary.ScaffoldRetrieved,
		FormalDFFamilyAvailable: prev.Summary.FormalDFFamilyAvailable,
		RawMomentsEvaluated:     prev.Summary.RawMomentsEvaluated,
		MomentDependenceExposed: prev.Summary.MomentDependenceExposed,
		CanonicalDFDerived:      prev.Dirac.CanonicalBlockSelected,
		HiggsRatioDerived:       prev.Summary.HiggsRatioDerived,
		FirewallPreserved:       prev.Summary.FirewallPreserved,
		RecommendedNextGate:     prev.Future.RecommendedNextGate,
		Verdict:                 StatusGate268Inherited + "; Gate 268 supplies the formal D_F family and proves that canonical selection is the next obstruction",
	}
}

func defineOrderOne() OrderOneDefinition {
	return OrderOneDefinition{
		Formula:                "[[D_F, rho(a)], J rho(b*) J^{-1}] = 0 for all a,b in A_F",
		AlgebraSymbol:          "A_F = C ⊕ M3(C)",
		DiracSymbol:            "D_F(M) = [[0,M],[M†,0]]",
		RealStructureSymbol:    "J candidate / opposite algebra action",
		RequiresRepresentation: true,
		RequiresOppositeAction: true,
		RequiresAllAAndB:       true,
		Defined:                true,
		Verdict:                StatusOrderOneDefined + "; the formula is exact, but it is only meaningful after rho and the opposite action are fixed",
	}
}

func auditAlgebraRepresentation() AlgebraRepresentationAudit {
	return AlgebraRepresentationAudit{
		NativeAlgebraName:             "C ⊕ M3(C)",
		ModeCarrier:                   "W = C·N0 ⊕ C·{N1,N2,N3}",
		ModeDimension:                 4,
		ModeLevelCPlusM3Available:     true,
		FullSCRepresentationDerived:   false,
		LeftRepresentationDerived:     false,
		RightRepresentationDerived:    false,
		OppositeRepresentationDerived: false,
		PhysicalJDerived:              false,
		NonVacuousOneFormsAvailable:   false,
		ToyModePreflightAllowed:       true,
		ImportedConnesAlgebra:         false,
		Verdict:                       StatusModeLevelAlgebraPreflight + "; only the mode-level 1⊕3 algebra can be audited without importing the full Connes representation",
	}
}

func applyModeLevelOrderOneSieve(def OrderOneDefinition, alg AlgebraRepresentationAudit) ([]ConstraintRow, GenericDiracBlockAudit) {
	rows := []ConstraintRow{
		{Name: "temporal-to-spatial row", Before: "r ∈ Hom(C³,C)", After: "r=0", Reason: "varying λ∈C and B∈M3(C) makes rB-λr fail the double commutator unless r vanishes", Satisfied: true},
		{Name: "spatial-to-temporal column", Before: "c ∈ Hom(C,C³)", After: "c=0", Reason: "varying λ and B makes cλ-Bc fail the double commutator unless c vanishes", Satisfied: true},
		{Name: "color internal block", Before: "D ∈ M3(C)", After: "D=y I3", Reason: "the only color block whose commutators commute with all M3(C) probes is the scalar color center", Satisfied: true},
		{Name: "temporal scalar block", Before: "x ∈ C", After: "x free", Reason: "the temporal C summand is central on the one-dimensional temporal slot", Satisfied: true},
	}
	return rows, GenericDiracBlockAudit{
		Carrier:                        "mode-level W preflight, not the full doubled S_C spectral Hilbert space",
		GenericMatrixShape:             "M = [[x,r],[c,D]] with x∈C, r∈C^{1×3}, c∈C^{3×1}, D∈M3(C)",
		InitialComplexParameters:       16,
		InitialRealParameters:          32,
		OrderOneToyConstraint:          "[[M,a],b]=0 for all a,b∈C⊕M3(C) under the same-side mode representation",
		AllowedComplexParameters:       2,
		AllowedRealParameters:          4,
		EliminatedComplexParameters:    14,
		TemporalSpatialLeakageRemoved:  true,
		ColorAnisotropyRemoved:         true,
		AllowedFamilyFormula:           "M_order1(x,y)=diag(x,y,y,y)",
		SieveNontrivial:                def.Defined && alg.ModeLevelCPlusM3Available,
		SievePhysicalOnFullSC:          false,
		OneFormsVanishForAllowedFamily: true,
		CanonicalBlockSelected:         false,
		Verdict:                        StatusOrderOneSieveReduced + "; the mode-level toy calculus removes leakage but leaves a two-complex-parameter commutant family and produces no non-vacuous one-forms",
	}
}

func reevaluateAllowedMoments(sieve GenericDiracBlockAudit) SpectralMomentReevaluation {
	rows := []MomentRow{
		momentRow("order-one unit commutant", 1, 1, true, false, "x=y=1: the most symmetric allowed representative, not a selected theorem"),
		momentRow("order-one lepton-weight deformation", 2, 1, true, false, "x=2,y=1 is still allowed by the mode-level order-one sieve"),
		momentRow("order-one color-weight deformation", 1, 2, true, false, "x=1,y=2 is also allowed; the raw ratio changes again"),
	}
	stable := true
	for i := 1; i < len(rows); i++ {
		if !approx(rows[0].RawRatio, rows[i].RawRatio, 1e-12) {
			stable = false
		}
	}
	return SpectralMomentReevaluation{
		Rows:                          rows,
		MomentsRecomputed:             sieve.SieveNontrivial,
		AllRowsOrderOneAllowed:        true,
		RawRatioStableAcrossAllowedDF: stable,
		DependsOnSurvivingAmplitudes:  !stable,
		SeeleyDeWittMapDerived:        false,
		HiggsRatioDerived:             false,
		Verdict:                       StatusAllowedMomentsReevaluated + "; " + StatusFailedMomentRatio + "; even after the toy order-one sieve, Tr(D²)/Tr(D⁴) depends on the surviving x:y amplitude ratio",
	}
}

func momentRow(name string, x, y float64, allowed, canonical bool, comment string) MomentRow {
	sig := []float64{math.Abs(x), math.Abs(y), math.Abs(y), math.Abs(y)}
	var s2, s4 float64
	for _, s := range sig {
		s2 += s * s
		s4 += s * s * s * s
	}
	tr2 := 2 * s2
	tr4 := 2 * s4
	ratio := math.Inf(1)
	if tr4 != 0 {
		ratio = tr2 / tr4
	}
	return MomentRow{Name: name, X: x, Y: y, SingularValues: sig, TraceD2: tr2, TraceD4: tr4, RawRatio: ratio, OrderOneAllowed: allowed, Canonical: canonical, Comment: comment}
}

func auditCanonicalDF(sieve GenericDiracBlockAudit, moments SpectralMomentReevaluation) CanonicalDFVerdict {
	return CanonicalDFVerdict{
		OrderOneSieveNontrivial:       sieve.SieveNontrivial,
		UniqueDFSelected:              false,
		SurvivingFamilyDimensionC:     sieve.AllowedComplexParameters,
		RequiresAdditionalSelector:    true,
		CouldUseNormalizationOnly:     false,
		NormalizationDerived:          false,
		GaugeProjectionDerived:        false,
		ScalarFluctuationMapDerived:   false,
		PromotableFiniteDiracOperator: false,
		Verdict:                       StatusFailedCanonicalDF + "; the order-one preflight leaves M=diag(x,yI3), and neither the relative amplitude x:y nor a physical full-S_C representation is selected",
	}
}

func auditFirewall(inh Gate268Inheritance, alg AlgebraRepresentationAudit, sieve GenericDiracBlockAudit, canonical CanonicalDFVerdict) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    true,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          true,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		NoConnesAlgebraImported:         !alg.ImportedConnesAlgebra,
		NoYukawaFitUsed:                 true,
		ToySieveNotPromoted:             sieve.SieveNontrivial && !sieve.SievePhysicalOnFullSC && !canonical.PromotableFiniteDiracOperator,
		NoHiggsPredictionClaim:          true,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalSealPreserved + "; the mode-level order-one calculation is retained as a preflight sieve and not promoted into a physical spectral triple",
	}
}

func defineFutureMap(alg AlgebraRepresentationAudit, sieve GenericDiracBlockAudit, moments SpectralMomentReevaluation, canonical CanonicalDFVerdict) FutureMap {
	obligations := []FutureObligation{
		{Name: "faithful representation of C⊕M3(C) on doubled S_C", Required: true, Satisfied: alg.FullSCRepresentationDerived, Detail: "Mode-level W preflight must be lifted to the full finite Hilbert carrier."},
		{Name: "physical opposite action through J", Required: true, Satisfied: alg.OppositeRepresentationDerived && alg.PhysicalJDerived, Detail: "Jb*J^{-1} must be a derived anti-linear opposite representation, not a placeholder."},
		{Name: "left/right chiral representation split", Required: true, Satisfied: alg.LeftRepresentationDerived && alg.RightRepresentationDerived, Detail: "Order-one constraints require actual rho_L and rho_R actions on the chosen D_F block."},
		{Name: "non-vacuous one-form calculus", Required: true, Satisfied: alg.NonVacuousOneFormsAvailable && !sieve.OneFormsVanishForAllowedFamily, Detail: "Allowed D_F must generate nonzero inner fluctuations rather than commuting with the algebra."},
		{Name: "canonical x:y amplitude selector", Required: true, Satisfied: canonical.UniqueDFSelected, Detail: "The surviving diag(x,yI3) family still has an unselected relative amplitude."},
		{Name: "heat-kernel / Seeley-de Witt map", Required: true, Satisfied: moments.SeeleyDeWittMapDerived, Detail: "Raw Tr(D²), Tr(D⁴) must be promoted by a cutoff-moment and subtraction scheme."},
		{Name: "gauge and scalar fluctuation projections", Required: true, Satisfied: canonical.GaugeProjectionDerived && canonical.ScalarFluctuationMapDerived, Detail: "Higgs and gauge coefficients require separate projection maps."},
	}
	return FutureMap{
		Obligations:             obligations,
		NeedFaithfulSCRep:       true,
		NeedPhysicalOppositeJ:   true,
		NeedNonVacuousOneForms:  true,
		NeedQuaternionicOrWeakH: true,
		NeedCanonicalAmplitude:  true,
		NeedHeatKernelMap:       true,
		RecommendedNextGate:     "Gate 270 — Faithful Opposite-Action Representation / Non-Vacuous One-Form Calculus Audit",
		Verdict:                 "Gate 269 shows the order-one equation can act as a sieve, but the next theorem must make the finite algebra representation and opposite action non-vacuous on S_C.",
	}
}

func summarize(inh Gate268Inheritance, def OrderOneDefinition, alg AlgebraRepresentationAudit, sieve GenericDiracBlockAudit, moments SpectralMomentReevaluation, canonical CanonicalDFVerdict, firewall FirewallAudit, future FutureMap) Summary {
	status := strings.Join([]string{
		StatusGate268Inherited,
		StatusOrderOneDefined,
		StatusModeLevelAlgebraPreflight,
		StatusOrderOneSieveReduced,
		StatusAllowedMomentsReevaluated,
		StatusFailedFaithfulRep,
		StatusFailedOppositeAction,
		StatusFailedNonVacuousCalculus,
		StatusFailedCanonicalDF,
		StatusFailedMomentRatio,
		StatusFailedHiggsRatio,
	}, "; ")
	return Summary{
		Gate268Inherited:         inh.ScaffoldRetrieved && inh.FormalDFFamilyAvailable && inh.MomentDependenceExposed,
		OrderOneDefined:          def.Defined,
		ModeAlgebraPreflight:     alg.ModeLevelCPlusM3Available && alg.ToyModePreflightAllowed,
		OrderOneSieveReduced:     sieve.SieveNontrivial && sieve.AllowedComplexParameters < sieve.InitialComplexParameters,
		CanonicalDFDerived:       canonical.UniqueDFSelected,
		AllowedMomentRatioStable: moments.RawRatioStableAcrossAllowedDF,
		HiggsRatioDerived:        moments.HiggsRatioDerived,
		FirewallPreserved:        firewall.EmpiricalYukawaSealPreserved && firewall.ToySieveNotPromoted && !firewall.FiniteCorePolluted,
		Status:                   status,
		NextGate:                 future.RecommendedNextGate,
		Comment:                  "Gate 269 applies the order-one condition as a mode-level sieve and reduces M to diag(x,yI3), but this remains a vacuous/preflight commutant family with unselected amplitudes and no physical Higgs ratio.",
	}
}

func buildTruth(inh Gate268Inheritance, alg AlgebraRepresentationAudit, sieve GenericDiracBlockAudit, moments SpectralMomentReevaluation, canonical CanonicalDFVerdict, firewall FirewallAudit) string {
	return fmt.Sprintf("Gate 269 truth: inherited268=%t modeCPlusM3=%t sieve=%d→%d complex params physicalFullSC=%t ratioStable=%t canonicalDF=%t higgsRatio=%t firewall=%t", inh.ScaffoldRetrieved, alg.ModeLevelCPlusM3Available, sieve.InitialComplexParameters, sieve.AllowedComplexParameters, sieve.SievePhysicalOnFullSC, moments.RawRatioStableAcrossAllowedDF, canonical.UniqueDFSelected, moments.HiggsRatioDerived, firewall.ToySieveNotPromoted && !firewall.FiniteCorePolluted)
}

func approx(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
