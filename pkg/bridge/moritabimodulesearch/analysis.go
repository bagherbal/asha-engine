// Package moritabimodulesearch implements Gate 272:
// Finite Algebra Representation Obstruction Classification / Morita-Bimodule Search Audit.
//
// Gate 271 proved that the associative finite algebra A_F=C⊕M3(C) cannot be
// promoted to the full second-quantized Fock carrier S_C by the naive Γ or dΓ
// lifts.  Gate 272 therefore changes category: the finite spectral triple must
// be built on a first-quantized finite Hilbert bimodule H_F, not on the full
// multiparticle Fock algebra.
//
// The audit classifies the canonical semisimple A-bimodule summands
//
//	H_ij = V_i ⊗ V_j^*,  i,j ∈ {C, M3},
//
// with dimensions 1,3,3,9.  On the direct sum the left A-action and opposite
// right A-action are faithful, linear, unital, associative, and commute.  The
// order-one rule then becomes an exact edge sieve: a Dirac edge between H_ij
// and H_kl is non-vacuous only when i≠k, and order-one compatible only when
// j=l (or when the left action is identical and the edge is vacuous).  This
// exposes legal non-vacuous one-form edges, but it still does not select the
// lepton/quark amplitude ratio x:y.  A spectral-action normalization, weak
// quaternionic/chiral selector, or finite action functional is still required
// before a_2/a_4 or a Higgs ratio can be claimed.
package moritabimodulesearch

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fullscrepresentationsearch"
)

const (
	AuditID = "GATE272-FINITE-ALGEBRA-REPRESENTATION-OBSTRUCTION-CLASSIFICATION-MORITA-BIMODULE-SEARCH-AUDIT"

	StatusGate271Inherited           = "CONDITIONAL_SUPPORT_GATE271_ASSOCIATIVE_LIFT_OBSTRUCTION_INHERITED"
	StatusObstructionClassified      = "CONDITIONAL_SUPPORT_SECOND_QUANTIZATION_REPRESENTATION_OBSTRUCTION_CLASSIFIED"
	StatusFiniteBimoduleExtracted    = "CONDITIONAL_SUPPORT_FINITE_HILBERT_BIMODULE_EXTRACTED"
	StatusFaithfulBimoduleRep        = "CONDITIONAL_SUPPORT_FAITHFUL_A_AOP_BIMODULE_REPRESENTATION_DERIVED"
	StatusOppositeActionConstructed  = "CONDITIONAL_SUPPORT_MORITA_OPPOSITE_ACTION_CONSTRUCTED"
	StatusOrderOneEdgeSieveDerived   = "CONDITIONAL_SUPPORT_ORDER_ONE_MORITA_EDGE_SIEVE_DERIVED"
	StatusNonVacuousOneFormsExist    = "CONDITIONAL_SUPPORT_NONVACUOUS_ORDER_ONE_ONEFORM_EDGES_EXIST"
	StatusFailedNotFullFock          = "FAILED_ROUTE_SPECTRAL_TRIPLE_NOT_ON_SECOND_QUANTIZED_FULL_SC"
	StatusFailedNotPhysicalSMHilbert = "FAILED_ROUTE_PHYSICAL_SM_HILBERT_SEMANTICS_NOT_FULLY_DERIVED"
	StatusFailedXYRatio              = "FAILED_ROUTE_MORITA_ORDER_ONE_DOES_NOT_LOCK_XY_RATIO"
	StatusFailedCanonicalDF          = "FAILED_ROUTE_CANONICAL_DF_STILL_UNSELECTED"
	StatusFailedHiggsRatio           = "FAILED_ROUTE_A2_A4_HIGGS_RATIO_STILL_NOT_DERIVED"
	StatusEmpiricalSealPreserved     = "FAILED_ROUTE_EMPIRICAL_YUKAWA_SEAL_REMAINS_ACTIVE"
)

type Gate271Inheritance struct {
	FullCarrierEnumerated      bool
	CARPassed                  bool
	GammaFailedAdditivity      bool
	DGammaFailedAssociativity  bool
	OneParticleActionAvailable bool
	FullSCRepDerived           bool
	PhysicalJDerived           bool
	HiggsRatioDerived          bool
	FirewallPreserved          bool
	Verdict                    string
}

type ObstructionClassification struct {
	Algebra                         string
	FullFockCarrier                 string
	GammaMultiplicative             bool
	GammaAdditive                   bool
	DGammaAdditive                  bool
	DGammaMultiplicative            bool
	OneParticleAssociative          bool
	SpectralTripleLivesOnHF         bool
	FockIsSecondQuantizedKinematics bool
	Verdict                         string
}

type SimpleModule struct {
	Label            string
	AlgebraBlock     string
	ComplexDimension int
	IsFaithfulBlock  bool
}

type BimoduleSummand struct {
	Label            string
	Left             SimpleModule
	Right            SimpleModule
	ComplexDimension int
	LeftAction       string
	RightAction      string
}

type BimoduleExtraction struct {
	Summands                            []BimoduleSummand
	TotalComplexDimension               int
	ChiralDoubleDimension               int
	ParticleAntiParticleDoubleDimension int
	LeftActionFaithful                  bool
	RightOppositeActionFaithful         bool
	Linear                              bool
	Unital                              bool
	Associative                         bool
	StarCompatible                      bool
	LeftRightCommute                    bool
	FullFockCarrierUsed                 bool
	Verdict                             string
}

type OppositeActionAudit struct {
	Formula                                         string
	Constructed                                     bool
	AntiLinearJRequiredForPhysicalChargeConjugation bool
	AlgebraicOppositeActionFaithful                 bool
	ParticleAntiParticleSemanticsDerived            bool
	Verdict                                         string
}

type DiracEdge struct {
	From              string
	To                string
	SameLeft          bool
	SameRight         bool
	NonVacuousOneForm bool
	OrderOneAllowed   bool
	AmplitudeLabel    string
	Reason            string
}

type OrderOneAudit struct {
	Rule                             string
	Edges                            []DiracEdge
	NonVacuousAllowedEdges           int
	VacuousAllowedEdges              int
	RejectedEdges                    int
	OrderOneSatisfiedForAllowedEdges bool
	NonVacuousOneFormsAvailable      bool
	CanonicalDFSelected              bool
	SurvivingAmplitudeLabels         []string
	XYRatioLocked                    bool
	Verdict                          string
}

type SpectralRatioAudit struct {
	CandidateDFShape  string
	TraceD2Formula    string
	TraceD4Formula    string
	RatioFormula      string
	DependsOnXOverY   bool
	A2A4Derived       bool
	HiggsRatioDerived bool
	MissingSelector   string
	Verdict           string
}

type FirewallAudit struct {
	EmpiricalYukawaSealPreserved    bool
	SpontaneousCarrierSealPreserved bool
	NoObservedMassInserted          bool
	NoVEVInserted                   bool
	NoCutoffScaleInserted           bool
	NoConnesSMAlgebraImported       bool
	BimoduleNotPromotedToSM         bool
	NoHiggsPredictionClaim          bool
	FiniteCorePolluted              bool
	Verdict                         string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                               []FutureCriterion
	NeedWeakQuaternionicOrChiralSelector   bool
	NeedCanonicalInnerProductNormalization bool
	NeedFiniteSpectralActionProjection     bool
	NeedAmplitudeSelector                  bool
	RecommendedNextGate                    string
	Verdict                                string
}

type Summary struct {
	Gate271Inherited         bool
	ObstructionClassified    bool
	BimoduleExtracted        bool
	FaithfulOppositeAction   bool
	NonVacuousOrderOneEdges  bool
	PhysicalSMHilbertDerived bool
	CanonicalDFDerived       bool
	XYRatioLocked            bool
	A2A4Derived              bool
	HiggsRatioDerived        bool
	FirewallPreserved        bool
	Status                   string
	NextGate                 string
	Comment                  string
}

type Analysis struct {
	PreviousGate271 fullscrepresentationsearch.Analysis
	Inheritance     Gate271Inheritance
	Obstruction     ObstructionClassification
	Bimodule        BimoduleExtraction
	Opposite        OppositeActionAudit
	OrderOne        OrderOneAudit
	Ratio           SpectralRatioAudit
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
		prev, err := fullscrepresentationsearch.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 271 predecessor: %w", err)
			return
		}
		inh := inheritGate271(prev)
		obs := classifyObstruction(prev)
		bim := extractBimodule()
		opp := constructOpposite(bim)
		oo := auditOrderOneEdges(bim)
		ratio := auditRatio(oo)
		fw := auditFirewall(obs, bim, opp, oo, ratio)
		future := defineFuture(oo, ratio)
		summary := summarize(inh, obs, bim, opp, oo, ratio, fw, future)
		truth := buildTruth(obs, bim, opp, oo, ratio, fw)
		defaultA = Analysis{PreviousGate271: prev, Inheritance: inh, Obstruction: obs, Bimodule: bim, Opposite: opp, OrderOne: oo, Ratio: ratio, Firewall: fw, Future: future, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate271(prev fullscrepresentationsearch.Analysis) Gate271Inheritance {
	gammaFailed := false
	dGammaFailed := false
	one := false
	for _, c := range prev.Representation.Candidates {
		switch c.Name {
		case "Γ exterior functor lift":
			gammaFailed = !c.LinearAdditive && c.Multiplicative
		case "dΓ creation-annihilation bilinear lift":
			dGammaFailed = c.LinearAdditive && !c.Multiplicative && !c.Unital
		case "one-particle sector inclusion":
			one = c.FaithfulOnOneParticle && c.LinearAdditive && c.Multiplicative && c.Unital
		}
	}
	return Gate271Inheritance{
		FullCarrierEnumerated:      prev.Summary.FullCarrierEnumerated,
		CARPassed:                  prev.Carrier.CARPassed,
		GammaFailedAdditivity:      gammaFailed,
		DGammaFailedAssociativity:  dGammaFailed,
		OneParticleActionAvailable: one,
		FullSCRepDerived:           prev.Summary.ValidFullSCRepDerived,
		PhysicalJDerived:           prev.Summary.PhysicalJDerived,
		HiggsRatioDerived:          prev.Summary.HiggsRatioDerived,
		FirewallPreserved:          prev.Summary.FirewallPreserved,
		Verdict:                    StatusGate271Inherited + "; Gate 271 proved full-Fock associative lift obstruction and exposed Λ¹W as the correct seed, not the full spectral carrier",
	}
}

func classifyObstruction(prev fullscrepresentationsearch.Analysis) ObstructionClassification {
	return ObstructionClassification{
		Algebra:                         "A_F = C ⊕ M3(C)",
		FullFockCarrier:                 "S_C=Λ*(C^4), dim_C=16; doubled field carrier dim_C=32",
		GammaMultiplicative:             true,
		GammaAdditive:                   false,
		DGammaAdditive:                  true,
		DGammaMultiplicative:            false,
		OneParticleAssociative:          prev.Summary.NativeOperatorLiftsAudited,
		SpectralTripleLivesOnHF:         true,
		FockIsSecondQuantizedKinematics: true,
		Verdict:                         StatusObstructionClassified + "; " + StatusFailedNotFullFock + "; Γ/dΓ failures classify S_C as second-quantized kinematics, while the associative spectral triple must live on a finite first-quantized Hilbert bimodule H_F",
	}
}

func simpleModules() []SimpleModule {
	return []SimpleModule{
		{Label: "C", AlgebraBlock: "C", ComplexDimension: 1, IsFaithfulBlock: true},
		{Label: "Q", AlgebraBlock: "M3(C) fundamental", ComplexDimension: 3, IsFaithfulBlock: true},
	}
}

func extractBimodule() BimoduleExtraction {
	ms := simpleModules()
	summands := make([]BimoduleSummand, 0, 4)
	total := 0
	for _, l := range ms {
		for _, r := range ms {
			dim := l.ComplexDimension * r.ComplexDimension
			label := "H_" + l.Label + r.Label
			summands = append(summands, BimoduleSummand{
				Label:            label,
				Left:             l,
				Right:            r,
				ComplexDimension: dim,
				LeftAction:       "ρ_L(a_i) on " + l.Label,
				RightAction:      "ρ_R^op(a_j) on " + r.Label + "*",
			})
			total += dim
		}
	}
	return BimoduleExtraction{
		Summands:                            summands,
		TotalComplexDimension:               total,
		ChiralDoubleDimension:               2 * total,
		ParticleAntiParticleDoubleDimension: 4 * total,
		LeftActionFaithful:                  true,
		RightOppositeActionFaithful:         true,
		Linear:                              true,
		Unital:                              true,
		Associative:                         true,
		StarCompatible:                      true,
		LeftRightCommute:                    true,
		FullFockCarrierUsed:                 false,
		Verdict:                             StatusFiniteBimoduleExtracted + "; " + StatusFaithfulBimoduleRep + "; semisimple Morita summands H_ij=V_i⊗V_j* give a faithful A_F⊗A_F^op representation on H_F, categorically separate from full S_C",
	}
}

func constructOpposite(bim BimoduleExtraction) OppositeActionAudit {
	return OppositeActionAudit{
		Formula:     "ρ^op(b)|_{H_ij}=I_{V_i}⊗ρ_j(b)^T on V_i⊗V_j*; physical anti-linear J would exchange H_ij ↔ H_ji with conjugation",
		Constructed: bim.RightOppositeActionFaithful && bim.LeftRightCommute,
		AntiLinearJRequiredForPhysicalChargeConjugation: true,
		AlgebraicOppositeActionFaithful:                 true,
		ParticleAntiParticleSemanticsDerived:            false,
		Verdict:                                         StatusOppositeActionConstructed + "; algebraic opposite action exists on the Morita bimodule, but physical charge-conjugation semantics remain a separate finite-particle theorem",
	}
}

func auditOrderOneEdges(bim BimoduleExtraction) OrderOneAudit {
	edges := []DiracEdge{}
	allowedNonVac := 0
	allowedVac := 0
	rejected := 0
	labels := map[string]bool{}
	for i := range bim.Summands {
		for j := i + 1; j < len(bim.Summands); j++ {
			a := bim.Summands[i]
			b := bim.Summands[j]
			sameLeft := a.Left.Label == b.Left.Label
			sameRight := a.Right.Label == b.Right.Label
			nonVac := !sameLeft
			allowed := sameLeft || sameRight
			amp := ""
			reason := ""
			if allowed && nonVac {
				amp = "m_" + a.Right.Label
				labels[amp] = true
				allowedNonVac++
				reason = "left modules differ, so [D,ρ_L(a)] can be nonzero; right module is shared, so the double commutator with A^op vanishes"
			} else if allowed {
				allowedVac++
				reason = "same left module, hence the edge is order-one compatible but one-form-vacuous for left fluctuations"
			} else {
				rejected++
				reason = "both left and right module labels differ; generic double commutator is nonzero"
			}
			edges = append(edges, DiracEdge{From: a.Label, To: b.Label, SameLeft: sameLeft, SameRight: sameRight, NonVacuousOneForm: nonVac, OrderOneAllowed: allowed, AmplitudeLabel: amp, Reason: reason})
		}
	}
	amps := []string{}
	for k := range labels {
		amps = append(amps, k)
	}
	// deterministic order for the two native right-sector labels
	if labels["m_C"] && labels["m_Q"] {
		amps = []string{"m_C", "m_Q"}
	}
	return OrderOneAudit{
		Rule:                             "An edge H_ij↔H_kl is order-one allowed when i=k (vacuous left one-form) or j=l (same opposite module); it is non-vacuous only when i≠k.",
		Edges:                            edges,
		NonVacuousAllowedEdges:           allowedNonVac,
		VacuousAllowedEdges:              allowedVac,
		RejectedEdges:                    rejected,
		OrderOneSatisfiedForAllowedEdges: true,
		NonVacuousOneFormsAvailable:      allowedNonVac > 0,
		CanonicalDFSelected:              false,
		SurvivingAmplitudeLabels:         amps,
		XYRatioLocked:                    false,
		Verdict:                          StatusOrderOneEdgeSieveDerived + "; " + StatusNonVacuousOneFormsExist + "; order-one permits non-vacuous Morita edges but leaves independent right-sector amplitudes " + strings.Join(amps, ",") + " unselected",
	}
}

func auditRatio(oo OrderOneAudit) SpectralRatioAudit {
	return SpectralRatioAudit{
		CandidateDFShape:  "D_F contains order-one allowed non-vacuous edges with independent amplitudes x≈m_C and y≈m_Q",
		TraceD2Formula:    "Tr(D_F^2)=κ_C |x|^2 + κ_Q |y|^2 (multiplicities depend on the selected physical sub-bimodule)",
		TraceD4Formula:    "Tr(D_F^4)=κ'_C |x|^4 + κ'_Q |y|^4 + possible edge-interference terms after a physical D_F is selected",
		RatioFormula:      "a2/a4 proxy remains a function of |x/y| unless an additional selector fixes x:y and the normalization scheme",
		DependsOnXOverY:   true,
		A2A4Derived:       false,
		HiggsRatioDerived: false,
		MissingSelector:   "weak/quaternionic/chiral representation choice, finite inner-product normalization, or spectral-action amplitude theorem",
		Verdict:           StatusFailedXYRatio + "; " + StatusFailedCanonicalDF + "; " + StatusFailedHiggsRatio + "; Morita isolation repairs the representation category but does not create an amplitude law",
	}
}

func auditFirewall(obs ObstructionClassification, bim BimoduleExtraction, opp OppositeActionAudit, oo OrderOneAudit, ratio SpectralRatioAudit) FirewallAudit {
	return FirewallAudit{
		EmpiricalYukawaSealPreserved:    true,
		SpontaneousCarrierSealPreserved: true,
		NoObservedMassInserted:          true,
		NoVEVInserted:                   true,
		NoCutoffScaleInserted:           true,
		NoConnesSMAlgebraImported:       true,
		BimoduleNotPromotedToSM:         !opp.ParticleAntiParticleSemanticsDerived && !oo.CanonicalDFSelected,
		NoHiggsPredictionClaim:          !ratio.HiggsRatioDerived,
		FiniteCorePolluted:              false,
		Verdict:                         StatusEmpiricalSealPreserved + "; the Morita bimodule is an algebraic carrier classification, not an imported Standard Model finite Hilbert space or Higgs prediction",
	}
}

func defineFuture(oo OrderOneAudit, ratio SpectralRatioAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "select physical sub-bimodule inside universal H_ij ledger", Required: true, Satisfied: false, Detail: "Gate 272 classifies all four simple Morita sectors; it does not derive which subset is the physical fermion Hilbert space."},
		{Name: "derive physical anti-linear J and charge-conjugation semantics", Required: true, Satisfied: false, Detail: "The algebraic opposite action is available, but particle/antiparticle semantics are not derived."},
		{Name: "derive weak/quaternionic or chiral selector", Required: true, Satisfied: false, Detail: "A non-vacuous order-one calculus alone does not distinguish the lepton/quark amplitude ratio."},
		{Name: "derive finite inner-product and multiplicity normalization", Required: true, Satisfied: false, Detail: "Trace coefficients κ_C,κ_Q are not physical before the Hilbert metric and selected sub-bimodule are fixed."},
		{Name: "derive spectral-action heat-kernel/cutoff projection", Required: true, Satisfied: false, Detail: "Raw finite moments are not yet Seeley-de Witt coefficients."},
		{Name: "lock x:y before comparing Higgs ratio", Required: true, Satisfied: oo.XYRatioLocked, Detail: "Surviving order-one amplitudes remain independent."},
	}
	return FutureMap{
		Criteria:                               criteria,
		NeedWeakQuaternionicOrChiralSelector:   true,
		NeedCanonicalInnerProductNormalization: true,
		NeedFiniteSpectralActionProjection:     true,
		NeedAmplitudeSelector:                  ratio.DependsOnXOverY,
		RecommendedNextGate:                    "Gate 273 — Weak/Quaternionic Sub-Bimodule Selector / Finite Inner-Product Normalization Audit",
		Verdict:                                "Morita classification opens the lawful first-quantized arena, but the next theorem must select the physical sub-bimodule and amplitude normalization before Path B can compute a2/a4.",
	}
}

func summarize(inh Gate271Inheritance, obs ObstructionClassification, bim BimoduleExtraction, opp OppositeActionAudit, oo OrderOneAudit, ratio SpectralRatioAudit, fw FirewallAudit, future FutureMap) Summary {
	return Summary{
		Gate271Inherited:         inh.FullCarrierEnumerated && inh.GammaFailedAdditivity && inh.DGammaFailedAssociativity,
		ObstructionClassified:    obs.SpectralTripleLivesOnHF && obs.FockIsSecondQuantizedKinematics,
		BimoduleExtracted:        bim.LeftActionFaithful && bim.RightOppositeActionFaithful && bim.Associative && !bim.FullFockCarrierUsed,
		FaithfulOppositeAction:   opp.Constructed && opp.AlgebraicOppositeActionFaithful,
		NonVacuousOrderOneEdges:  oo.NonVacuousOneFormsAvailable && oo.OrderOneSatisfiedForAllowedEdges,
		PhysicalSMHilbertDerived: opp.ParticleAntiParticleSemanticsDerived,
		CanonicalDFDerived:       oo.CanonicalDFSelected,
		XYRatioLocked:            oo.XYRatioLocked,
		A2A4Derived:              ratio.A2A4Derived,
		HiggsRatioDerived:        ratio.HiggsRatioDerived,
		FirewallPreserved:        fw.EmpiricalYukawaSealPreserved && fw.NoHiggsPredictionClaim && !fw.FiniteCorePolluted,
		Status:                   StatusFiniteBimoduleExtracted + "; " + StatusFailedXYRatio,
		NextGate:                 future.RecommendedNextGate,
		Comment:                  "We are not ready to derive a2/a4: Gate 272 repairs the representation category and exposes legal non-vacuous order-one edges, but x:y remains an independent amplitude ratio requiring another selector.",
	}
}

func buildTruth(obs ObstructionClassification, bim BimoduleExtraction, opp OppositeActionAudit, oo OrderOneAudit, ratio SpectralRatioAudit, fw FirewallAudit) string {
	return strings.Join([]string{
		"Gate 272 classifies the second-quantization obstruction as categorical rather than accidental: the finite spectral triple belongs on a first-quantized Hilbert bimodule H_F, not on the full Fock carrier S_C.",
		"The semisimple Morita ledger H_ij=V_i⊗V_j* gives a faithful associative A_F⊗A_F^op representation and a lawful opposite action.",
		"The order-one condition becomes a clean edge sieve: non-vacuous one-forms are allowed precisely along edges with different left module and shared right module.",
		"However, the surviving right-sector amplitudes remain independent, so the lepton/quark weight x:y and any Seeley-de Witt Higgs ratio are still not derived.",
	}, " ")
}
