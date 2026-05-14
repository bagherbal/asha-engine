// Package fullscrepresentationsearch implements Gate 271:
// Full S_C Finite Algebra Representation Search / Opposite-Action Construction Audit.
//
// Gate 270 exposed the central NCG tension on a small chiral diagnostic:
// non-vacuous one-forms require a left/right mismatch, but the toy opposite
// action failed the order-one double commutator.  Gate 271 therefore searches
// for the missing object on the native full Fock carrier S_C=Λ*(C^4).  The
// audit deliberately distinguishes three mathematically different lifts:
//
//	Γ(A)   : exterior functor lift; multiplicative but not linear/additive.
//	dΓ(A)  : creation-annihilation bilinear lift; linear/Lie-like but not a
//	         unital associative *-representation of C⊕M3(C).
//	W⊂S_C  : one-particle faithful action; faithful on the seed sector but not
//	         a canonical full-carrier representation.
//
// This distinction is the gate result.  The engine can write native Fock
// operators on the full 16-state carrier, but it still lacks a faithful
// associative representation of the finite algebra on the full doubled S_C
// space together with a physical anti-linear J/opposite action.  Therefore the
// order-one condition cannot yet be promoted into a completed spectral triple,
// and the x:y/Higgs-ratio problem remains sealed.
package fullscrepresentationsearch

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/faithfuloppositeactionrep"
)

const (
	AuditID = "GATE271-FULL-SC-FINITE-ALGEBRA-REPRESENTATION-OPPOSITE-ACTION-CONSTRUCTION-AUDIT"

	StatusGate270Inherited           = "CONDITIONAL_SUPPORT_GATE270_NONVACUOUS_TARGET_INHERITED"
	StatusFullFockCarrierEnumerated  = "CONDITIONAL_SUPPORT_FULL_SC_FOCK_CARRIER_ENUMERATED"
	StatusCARPreflightPassed         = "CONDITIONAL_SUPPORT_CREATION_ANNIHILATION_CAR_PREFLIGHT_PASSED"
	StatusExteriorLiftAudited        = "CONDITIONAL_SUPPORT_EXTERIOR_FUNCTOR_LIFT_AUDITED"
	StatusDGammaLiftAudited          = "CONDITIONAL_SUPPORT_SECOND_QUANTIZED_DGAMMA_LIFT_AUDITED"
	StatusOneParticleFaithful        = "CONDITIONAL_SUPPORT_ONE_PARTICLE_SECTOR_FAITHFUL_ACTION_AVAILABLE"
	StatusOppositeActionAudited      = "CONDITIONAL_SUPPORT_OPPOSITE_ACTION_REQUIREMENTS_AUDITED"
	StatusFailedNoFullRep            = "FAILED_ROUTE_FULL_SC_ASSOCIATIVE_ALGEBRA_REPRESENTATION_NOT_DERIVED"
	StatusFailedGammaNonlinear       = "FAILED_ROUTE_EXTERIOR_GAMMA_LIFT_NOT_ADDITIVE"
	StatusFailedDGammaNonAssociative = "FAILED_ROUTE_DGAMMA_LIFT_NOT_UNITAL_ASSOCIATIVE_REPRESENTATION"
	StatusFailedOneParticleOnly      = "FAILED_ROUTE_ONE_PARTICLE_ACTION_DOES_NOT_DEFINE_FULL_SC_REPRESENTATION"
	StatusFailedPhysicalJ            = "FAILED_ROUTE_PHYSICAL_J_OPPOSITE_ACTION_STILL_MISSING"
	StatusFailedOrderOne             = "FAILED_ROUTE_FULL_SC_ORDER_ONE_NOT_REEVALUATED_AS_THEOREM"
	StatusFailedXYRatio              = "FAILED_ROUTE_XY_RATIO_STILL_UNCONSTRAINED"
	StatusFailedHiggsRatio           = "FAILED_ROUTE_INVARIANT_HIGGS_RATIO_NOT_DERIVED"
	StatusEmpiricalSealPreserved     = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate270Inheritance struct {
	CandidateOneFormsExposed bool
	CandidateOrderOnePasses  bool
	FullSCRepresentation     bool
	PhysicalOppositeAction   bool
	XYRatioSelected          bool
	HiggsRatioDerived        bool
	FirewallPreserved        bool
	RecommendedNextGate      string
	Verdict                  string
}

type FockCarrierAudit struct {
	ModeCount                      int
	BaseComplexDimension           int
	DoubledComplexDimension        int
	BasisMasksEnumerated           int
	GradeHistogram                 map[int]int
	ParityEvenStates               int
	ParityOddStates                int
	CreationOperatorsAvailable     bool
	AnnihilationOperatorsAvailable bool
	CARMaxResidual                 float64
	CARPassed                      bool
	Verdict                        string
}

type AlgebraProbe struct {
	Name string
	Diag [4]float64
}

type LiftCandidate struct {
	Name                     string
	Formula                  string
	ActsOnFullSC             bool
	UsesCreationAnnihilation bool
	FaithfulOnOneParticle    bool
	LinearAdditive           bool
	Multiplicative           bool
	Unital                   bool
	StarCompatible           bool
	AssociativeAlgebraRep    bool
	DiagnosticDefect         float64
	DefectDetail             string
	Verdict                  string
}

type RepresentationSearchAudit struct {
	FiniteAlgebra                string
	TargetCarrier                string
	Candidates                   []LiftCandidate
	ValidFullAssociativeRepFound bool
	FullSCPromotionBlocked       bool
	BestNativeOperatorCalculus   string
	Verdict                      string
}

type OppositeActionAudit struct {
	RequiresValidLeftRepresentation  bool
	CandidateJFormula                string
	CandidateJAntiLinear             bool
	CandidateJPhysicalSemantics      bool
	OppositeActionConstructed        bool
	OrderOneCanBeEvaluatedPhysically bool
	Verdict                          string
}

type OrderOneReevaluationAudit struct {
	DiracFamilyFormula           string
	FullSCLeftRepAvailable       bool
	PhysicalOppositeRepAvailable bool
	NonVacuousOneFormsDerived    bool
	OrderOneSatisfied            bool
	ReevaluatedAsSpectralTriple  bool
	Gate270ToyResidualInherited  float64
	Verdict                      string
}

type RatioAudit struct {
	XYRatioSelected                bool
	TraceRatioStable               bool
	GaugeProjectionDerived         bool
	ScalarFluctuationMapDerived    bool
	HeatKernelNormalizationDerived bool
	HiggsRatioDerived              bool
	Verdict                        string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	NoConnesModelImported           bool
	NoCandidatePromoted             bool
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
	Obligations                   []FutureObligation
	NeedAssociativeFullSCRep      bool
	NeedPhysicalJ                 bool
	NeedOrderOnePassingNonVacuous bool
	NeedCanonicalXYSelector       bool
	NeedSpectralActionProjection  bool
	RecommendedNextGate           string
	Verdict                       string
}

type Summary struct {
	Gate270Inherited           bool
	FullCarrierEnumerated      bool
	CARPreflightPassed         bool
	NativeOperatorLiftsAudited bool
	ValidFullSCRepDerived      bool
	PhysicalJDerived           bool
	FullOrderOneProved         bool
	NonVacuousOneFormsProved   bool
	XYRatioSelected            bool
	HiggsRatioDerived          bool
	FirewallPreserved          bool
	Status                     string
	NextGate                   string
	Comment                    string
}

type Analysis struct {
	PreviousGate270 faithfuloppositeactionrep.Analysis
	Inheritance     Gate270Inheritance
	Carrier         FockCarrierAudit
	Representation  RepresentationSearchAudit
	Opposite        OppositeActionAudit
	OrderOne        OrderOneReevaluationAudit
	Ratio           RatioAudit
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
		prev, err := faithfuloppositeactionrep.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 270 predecessor: %w", err)
			return
		}
		inh := inheritGate270(prev)
		carrier := auditFockCarrier()
		rep := searchRepresentations(carrier)
		opp := auditOppositeAction(rep)
		orderOne := reevaluateOrderOne(prev, rep, opp)
		ratio := auditRatio(orderOne)
		firewall := auditFirewall(rep, opp, orderOne, ratio)
		future := defineFutureMap(rep, opp, orderOne, ratio)
		summary := summarize(inh, carrier, rep, opp, orderOne, ratio, firewall, future)
		truth := buildTruth(inh, carrier, rep, opp, orderOne, ratio, firewall)
		defaultA = Analysis{PreviousGate270: prev, Inheritance: inh, Carrier: carrier, Representation: rep, Opposite: opp, OrderOne: orderOne, Ratio: ratio, Firewall: firewall, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate270(prev faithfuloppositeactionrep.Analysis) Gate270Inheritance {
	return Gate270Inheritance{
		CandidateOneFormsExposed: prev.Summary.CandidateOneFormsNonzero,
		CandidateOrderOnePasses:  prev.Summary.CandidateOrderOnePasses,
		FullSCRepresentation:     prev.Summary.FullSCRepresentation,
		PhysicalOppositeAction:   prev.Summary.PhysicalOppositeAction,
		XYRatioSelected:          prev.Ratio.XToYSelected,
		HiggsRatioDerived:        prev.Summary.HiggsRatioDerived,
		FirewallPreserved:        prev.Summary.FirewallPreserved,
		RecommendedNextGate:      prev.Future.RecommendedNextGate,
		Verdict:                  StatusGate270Inherited + "; Gate 270 exposes the non-vacuous-one-form target but not a physical full-S_C spectral triple",
	}
}

func auditFockCarrier() FockCarrierAudit {
	basis := fockMasks(4)
	hist := map[int]int{}
	even, odd := 0, 0
	for _, m := range basis {
		g := popcount(m)
		hist[g]++
		if g%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	maxResidual := carMaxResidual(4)
	passed := maxResidual < 1e-12
	return FockCarrierAudit{
		ModeCount:                      4,
		BaseComplexDimension:           len(basis),
		DoubledComplexDimension:        2 * len(basis),
		BasisMasksEnumerated:           len(basis),
		GradeHistogram:                 hist,
		ParityEvenStates:               even,
		ParityOddStates:                odd,
		CreationOperatorsAvailable:     true,
		AnnihilationOperatorsAvailable: true,
		CARMaxResidual:                 maxResidual,
		CARPassed:                      passed,
		Verdict:                        StatusFullFockCarrierEnumerated + "; " + StatusCARPreflightPassed + "; native a†_k,a_k operators correctly generate the 16-state S_C basis and satisfy the CAR preflight",
	}
}

func searchRepresentations(carrier FockCarrierAudit) RepresentationSearchAudit {
	identity := AlgebraProbe{Name: "I4", Diag: [4]float64{1, 1, 1, 1}}
	twoI := AlgebraProbe{Name: "2I4", Diag: [4]float64{2, 2, 2, 2}}
	D := AlgebraProbe{Name: "D=diag(1,2,0,0)", Diag: [4]float64{1, 2, 0, 0}}
	D2 := squareDiag(D)

	gammaAddDefect := maxAbsDiff(gammaEigen(twoI), addEigen(gammaEigen(identity), gammaEigen(identity)))
	dGammaUnitalDefect := maxAbsDiff(dGammaEigen(identity), ones(16))
	dGammaMultDefect := maxAbsDiff(dGammaEigen(D2), mulEigen(dGammaEigen(D), dGammaEigen(D)))

	candidates := []LiftCandidate{
		{
			Name:                     "Γ exterior functor lift",
			Formula:                  "Γ(A)|_{Λ^k W}=Λ^k A",
			ActsOnFullSC:             true,
			UsesCreationAnnihilation: false,
			FaithfulOnOneParticle:    true,
			LinearAdditive:           false,
			Multiplicative:           true,
			Unital:                   true,
			StarCompatible:           true,
			AssociativeAlgebraRep:    false,
			DiagnosticDefect:         gammaAddDefect,
			DefectDetail:             "additivity defect max|Γ(2I)-Γ(I)-Γ(I)| on S_C = " + fmtFloat(gammaAddDefect),
			Verdict:                  StatusExteriorLiftAudited + "; " + StatusFailedGammaNonlinear + "; Γ is a group/functorial lift, not a linear algebra representation of C⊕M3(C)",
		},
		{
			Name:                     "dΓ creation-annihilation bilinear lift",
			Formula:                  "dΓ(A)=Σ_{ij} A_ij a†_i a_j",
			ActsOnFullSC:             true,
			UsesCreationAnnihilation: true,
			FaithfulOnOneParticle:    true,
			LinearAdditive:           true,
			Multiplicative:           false,
			Unital:                   false,
			StarCompatible:           true,
			AssociativeAlgebraRep:    false,
			DiagnosticDefect:         math.Max(dGammaUnitalDefect, dGammaMultDefect),
			DefectDetail:             "unital defect max|dΓ(I)-I|=" + fmtFloat(dGammaUnitalDefect) + "; multiplicative defect max|dΓ(D²)-dΓ(D)²|=" + fmtFloat(dGammaMultDefect),
			Verdict:                  StatusDGammaLiftAudited + "; " + StatusFailedDGammaNonAssociative + "; dΓ is the native bilinear operator calculus, but it is not a unital associative representation of the finite algebra",
		},
		{
			Name:                     "one-particle sector inclusion",
			Formula:                  "ρ_W(λ,B)=diag(λ,B) on W=C⊕C^3⊂Λ¹W",
			ActsOnFullSC:             false,
			UsesCreationAnnihilation: false,
			FaithfulOnOneParticle:    true,
			LinearAdditive:           true,
			Multiplicative:           true,
			Unital:                   true,
			StarCompatible:           true,
			AssociativeAlgebraRep:    false,
			DiagnosticDefect:         0,
			DefectDetail:             "faithful only after projecting to Λ¹W; no canonical action on Λ⁰,Λ²,Λ³,Λ⁴ sectors is derived",
			Verdict:                  StatusOneParticleFaithful + "; " + StatusFailedOneParticleOnly + "; the seed action is faithful but not a full S_C spectral representation",
		},
	}
	return RepresentationSearchAudit{
		FiniteAlgebra:                "A_F = C ⊕ M3(C)",
		TargetCarrier:                "S_C=Λ*(C^4), dim_C=16; doubled carrier S_C⊕S_C*, dim_C=32",
		Candidates:                   candidates,
		ValidFullAssociativeRepFound: false,
		FullSCPromotionBlocked:       true,
		BestNativeOperatorCalculus:   "dΓ(A)=Σ A_ij a†_i a_j gives a lawful Fock-operator/Lie calculus, but not the required associative algebra representation",
		Verdict:                      StatusFailedNoFullRep + "; native Fock operators exist, but every audited lift fails one of linearity, multiplicativity, unitality, or full-carrier canonicity",
	}
}

func auditOppositeAction(rep RepresentationSearchAudit) OppositeActionAudit {
	return OppositeActionAudit{
		RequiresValidLeftRepresentation:  true,
		CandidateJFormula:                "candidate doubled conjugation J(ψ,φ)=(conj φ,conj ψ); physical charge-conjugation semantics not derived on S_C",
		CandidateJAntiLinear:             true,
		CandidateJPhysicalSemantics:      false,
		OppositeActionConstructed:        false,
		OrderOneCanBeEvaluatedPhysically: false,
		Verdict:                          StatusOppositeActionAudited + "; " + StatusFailedPhysicalJ + "; without a valid full-S_C left representation and particle/antiparticle semantics, Jρ(a*)J^-1 is not a physical opposite action",
	}
}

func reevaluateOrderOne(prev faithfuloppositeactionrep.Analysis, rep RepresentationSearchAudit, opp OppositeActionAudit) OrderOneReevaluationAudit {
	return OrderOneReevaluationAudit{
		DiracFamilyFormula:           "D_F(M)=((0,M),(M†,0)), inherited M_order1=diag(x,y,y,y) only on mode seed W",
		FullSCLeftRepAvailable:       rep.ValidFullAssociativeRepFound,
		PhysicalOppositeRepAvailable: opp.OppositeActionConstructed,
		NonVacuousOneFormsDerived:    false,
		OrderOneSatisfied:            false,
		ReevaluatedAsSpectralTriple:  false,
		Gate270ToyResidualInherited:  prev.Residual.FrobeniusNormSq,
		Verdict:                      StatusFailedOrderOne + "; the full order-one condition is not re-evaluated as a spectral-triple theorem because the required full representation/opposite action remains missing; Gate 270 toy residual ||·||²=" + fmtFloat(prev.Residual.FrobeniusNormSq) + " remains diagnostic only",
	}
}

func auditRatio(orderOne OrderOneReevaluationAudit) RatioAudit {
	return RatioAudit{
		XYRatioSelected:                false,
		TraceRatioStable:               false,
		GaugeProjectionDerived:         false,
		ScalarFluctuationMapDerived:    false,
		HeatKernelNormalizationDerived: false,
		HiggsRatioDerived:              false,
		Verdict:                        StatusFailedXYRatio + "; " + StatusFailedHiggsRatio + "; no invariant spectral ratio can be extracted before a physical full spectral triple, scalar/gauge projection, and heat-kernel normalization exist",
	}
}

func auditFirewall(rep RepresentationSearchAudit, opp OppositeActionAudit, orderOne OrderOneReevaluationAudit, ratio RatioAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    true,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          true,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		NoConnesModelImported:           true,
		NoCandidatePromoted:             !rep.ValidFullAssociativeRepFound && !opp.OppositeActionConstructed && !orderOne.ReevaluatedAsSpectralTriple,
		NoHiggsPredictionClaim:          !ratio.HiggsRatioDerived,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalSealPreserved + "; the Γ and dΓ calculations are retained as representation diagnostics and not promoted to physical NCG data",
	}
}

func defineFutureMap(rep RepresentationSearchAudit, opp OppositeActionAudit, orderOne OrderOneReevaluationAudit, ratio RatioAudit) FutureMap {
	obligations := []FutureObligation{
		{Name: "linear unital associative *-representation of C⊕M3(C) on full S_C", Required: true, Satisfied: rep.ValidFullAssociativeRepFound, Detail: "Γ is not additive; dΓ is not unital/multiplicative; Λ¹W is not the full carrier."},
		{Name: "physical anti-linear J with particle/antiparticle semantics", Required: true, Satisfied: opp.CandidateJPhysicalSemantics && opp.OppositeActionConstructed, Detail: "A swap-conjugation formula is not enough without derived opposite action semantics."},
		{Name: "non-vacuous one-forms satisfying order-one", Required: true, Satisfied: orderOne.NonVacuousOneFormsDerived && orderOne.OrderOneSatisfied, Detail: "Gate 270 exposed nonzero toy one-forms but they fail order-one; Gate 271 lacks the full representation needed to retry."},
		{Name: "canonical x:y selector for M=diag(x,yI3)", Required: true, Satisfied: ratio.XYRatioSelected, Detail: "The lepton/quark amplitude ratio remains free."},
		{Name: "gauge/scalar fluctuation projection", Required: true, Satisfied: ratio.GaugeProjectionDerived && ratio.ScalarFluctuationMapDerived, Detail: "Higgs/gauge coefficients require a projection map from one-forms to physical fields."},
		{Name: "heat-kernel and subtraction normalization", Required: true, Satisfied: ratio.HeatKernelNormalizationDerived, Detail: "Raw finite traces cannot be interpreted as Seeley-de Witt coefficients without normalization."},
	}
	return FutureMap{
		Obligations:                   obligations,
		NeedAssociativeFullSCRep:      true,
		NeedPhysicalJ:                 true,
		NeedOrderOnePassingNonVacuous: true,
		NeedCanonicalXYSelector:       true,
		NeedSpectralActionProjection:  true,
		RecommendedNextGate:           "Gate 272 — Finite Algebra Representation Obstruction Classification / Morita-Bimodule Search Audit",
		Verdict:                       "Gate 271 closes the naive full-S_C lift attempt and points to a representation-classification/Morita-bimodule search rather than another ad hoc chiral action",
	}
}

func summarize(inh Gate270Inheritance, carrier FockCarrierAudit, rep RepresentationSearchAudit, opp OppositeActionAudit, orderOne OrderOneReevaluationAudit, ratio RatioAudit, fw FirewallAudit, future FutureMap) Summary {
	return Summary{
		Gate270Inherited:           inh.CandidateOneFormsExposed && !inh.FullSCRepresentation,
		FullCarrierEnumerated:      carrier.BaseComplexDimension == 16 && carrier.DoubledComplexDimension == 32,
		CARPreflightPassed:         carrier.CARPassed,
		NativeOperatorLiftsAudited: len(rep.Candidates) == 3,
		ValidFullSCRepDerived:      rep.ValidFullAssociativeRepFound,
		PhysicalJDerived:           opp.OppositeActionConstructed && opp.CandidateJPhysicalSemantics,
		FullOrderOneProved:         orderOne.ReevaluatedAsSpectralTriple && orderOne.OrderOneSatisfied,
		NonVacuousOneFormsProved:   orderOne.NonVacuousOneFormsDerived,
		XYRatioSelected:            ratio.XYRatioSelected,
		HiggsRatioDerived:          ratio.HiggsRatioDerived,
		FirewallPreserved:          !fw.FiniteCorePolluted && fw.NoCandidatePromoted && fw.NoHiggsPredictionClaim,
		Status:                     StatusFullFockCarrierEnumerated + "; " + StatusFailedNoFullRep + "; " + StatusFailedPhysicalJ + "; " + StatusFailedHiggsRatio,
		NextGate:                   future.RecommendedNextGate,
		Comment:                    "The full Fock carrier and native operator calculus exist, but the required associative full-S_C algebra representation and physical opposite action remain missing.",
	}
}

func buildTruth(inh Gate270Inheritance, carrier FockCarrierAudit, rep RepresentationSearchAudit, opp OppositeActionAudit, orderOne OrderOneReevaluationAudit, ratio RatioAudit, fw FirewallAudit) string {
	lines := []string{
		"Gate 271 audits the full S_C representation problem rather than importing the standard NCG representation.",
		fmt.Sprintf("The native Fock carrier is present: dim_C(S_C)=%d and dim_C(S_C⊕S_C*)=%d with CAR residual %.3g.", carrier.BaseComplexDimension, carrier.DoubledComplexDimension, carrier.CARMaxResidual),
		"The exterior functor Γ(A) acts on all grades but is not additive, while dΓ(A)=ΣA_ij a†_i a_j is native and linear but not unital/multiplicative as an associative algebra representation.",
		"Therefore the engine still lacks a faithful full-S_C representation of C⊕M3(C) and cannot lawfully construct the physical opposite action Jρ(a*)J^-1.",
		"The order-one condition, non-vacuous one-forms, x:y selector, and Higgs spectral ratio remain blocked without polluting the finite core.",
	}
	return strings.Join(lines, " ")
}

// --- finite Fock diagnostics ---

func fockMasks(n int) []int {
	out := make([]int, 1<<n)
	for i := range out {
		out[i] = i
	}
	return out
}

func popcount(x int) int {
	c := 0
	for x != 0 {
		c += x & 1
		x >>= 1
	}
	return c
}

func fermionSign(mask, mode int) float64 {
	below := mask & ((1 << mode) - 1)
	if popcount(below)%2 == 0 {
		return 1
	}
	return -1
}

func creationMatrix(n, mode int) [][]float64 {
	dim := 1 << n
	m := zeroMatrix(dim)
	for mask := 0; mask < dim; mask++ {
		if mask&(1<<mode) != 0 {
			continue
		}
		to := mask | (1 << mode)
		m[to][mask] = fermionSign(mask, mode)
	}
	return m
}

func annihilationMatrix(n, mode int) [][]float64 {
	dim := 1 << n
	m := zeroMatrix(dim)
	for mask := 0; mask < dim; mask++ {
		if mask&(1<<mode) == 0 {
			continue
		}
		to := mask &^ (1 << mode)
		m[to][mask] = fermionSign(mask, mode)
	}
	return m
}

func carMaxResidual(n int) float64 {
	dim := 1 << n
	I := identityMatrix(dim)
	max := 0.0
	for i := 0; i < n; i++ {
		ai := annihilationMatrix(n, i)
		ci := creationMatrix(n, i)
		for j := 0; j < n; j++ {
			aj := annihilationMatrix(n, j)
			cj := creationMatrix(n, j)
			anti := addMatrix(mulMatrix(ai, cj), mulMatrix(cj, ai))
			target := zeroMatrix(dim)
			if i == j {
				target = I
			}
			d := matrixMaxAbsDiff(anti, target)
			if d > max {
				max = d
			}
			antiCreate := addMatrix(mulMatrix(ci, cj), mulMatrix(cj, ci))
			if d := matrixMaxAbsDiff(antiCreate, zeroMatrix(dim)); d > max {
				max = d
			}
			antiAnnih := addMatrix(mulMatrix(ai, aj), mulMatrix(aj, ai))
			if d := matrixMaxAbsDiff(antiAnnih, zeroMatrix(dim)); d > max {
				max = d
			}
		}
	}
	return max
}

func gammaEigen(p AlgebraProbe) []float64 {
	out := make([]float64, 16)
	for mask := 0; mask < 16; mask++ {
		prod := 1.0
		for k := 0; k < 4; k++ {
			if mask&(1<<k) != 0 {
				prod *= p.Diag[k]
			}
		}
		out[mask] = prod
	}
	return out
}

func dGammaEigen(p AlgebraProbe) []float64 {
	out := make([]float64, 16)
	for mask := 0; mask < 16; mask++ {
		sum := 0.0
		for k := 0; k < 4; k++ {
			if mask&(1<<k) != 0 {
				sum += p.Diag[k]
			}
		}
		out[mask] = sum
	}
	return out
}

func squareDiag(p AlgebraProbe) AlgebraProbe {
	q := AlgebraProbe{Name: p.Name + " squared"}
	for i := range p.Diag {
		q.Diag[i] = p.Diag[i] * p.Diag[i]
	}
	return q
}

func addEigen(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] + b[i]
	}
	return out
}
func mulEigen(a, b []float64) []float64 {
	out := make([]float64, len(a))
	for i := range a {
		out[i] = a[i] * b[i]
	}
	return out
}
func ones(n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		out[i] = 1
	}
	return out
}
func maxAbsDiff(a, b []float64) float64 {
	m := 0.0
	for i := range a {
		d := math.Abs(a[i] - b[i])
		if d > m {
			m = d
		}
	}
	return m
}

func zeroMatrix(n int) [][]float64 {
	m := make([][]float64, n)
	for i := range m {
		m[i] = make([]float64, n)
	}
	return m
}
func identityMatrix(n int) [][]float64 {
	m := zeroMatrix(n)
	for i := 0; i < n; i++ {
		m[i][i] = 1
	}
	return m
}
func addMatrix(a, b [][]float64) [][]float64 {
	n := len(a)
	m := zeroMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m[i][j] = a[i][j] + b[i][j]
		}
	}
	return m
}
func mulMatrix(a, b [][]float64) [][]float64 {
	n := len(a)
	m := zeroMatrix(n)
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if a[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return m
}
func matrixMaxAbsDiff(a, b [][]float64) float64 {
	m := 0.0
	for i := range a {
		for j := range a[i] {
			d := math.Abs(a[i][j] - b[i][j])
			if d > m {
				m = d
			}
		}
	}
	return m
}

func fmtFloat(x float64) string { return fmt.Sprintf("%.12g", x) }
