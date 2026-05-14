// Package liecarrierprojection implements Gate 245:
// Lie Algebra Isomorphism / Scalar-to-Spatial Carrier Projection Audit.
//
// Gate 244 traced tau_eta=(2,-2,1) to scalar-bundle trace records
// tau_eta(Q^TQ), tau_eta(Z^TZ), and tau_eta(T3L^T Y_phi). Gate 245 asks
// whether those scalar observables can be lawfully decomposed back to native
// contact-preserving derivations, then to spatial Fock modes e1,e2,e3.  The
// audit is intentionally strict: Q and Z decompose into the two-dimensional
// neutral electroweak plane spanned by T3 and Y_phi, not into the three su(2)
// basis generators. The contact su(2) itself still lacks a canonical lift to a
// selected spatial two-plane/axis basis. Therefore the chained carrier
// projection theorem is not derived and omega_tau remains blocked.
package liecarrierprojection

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/characteristicpullback"
)

const (
	AuditID = "GATE245-LIE-ALGEBRA-ISOMORPHISM-SCALAR-TO-SPATIAL-CARRIER-PROJECTION-AUDIT"

	StatusOperatorDecompositionTraced   = "CONDITIONAL_SUPPORT_EW_OPERATOR_DECOMPOSITION_TRACED"
	StatusNeutralPlaneObstruction       = "FAILED_ROUTE_TAU_ETA_SLOTS_NOT_SU2_BASIS"
	StatusDerivationBladePreflight      = "CONDITIONAL_SUPPORT_SU2_BIVECTOR_CAPACITY_PREFLIGHT"
	StatusDerivationBladeObstruction    = "FAILED_ROUTE_NATIVE_SU2_TO_SPATIAL_AXIS_ISOMORPHISM"
	StatusCarrierProjectionObstruction  = "FAILED_ROUTE_SCALAR_TO_SPATIAL_CARRIER_PROJECTION"
	StatusExteriorRepresentativeBlocked = "FAILED_ROUTE_LIE_PULLBACK_EXTERIOR_FORM_REPRESENTATIVE"
	StatusWeakPlaneBlocked              = "FAILED_ROUTE_LIE_PULLBACK_WEAK_PLANE_SELECTION"
	StatusGenerationTextureBlocked      = "FAILED_ROUTE_LIE_PULLBACK_GENERATION_TEXTURE"
	StatusGlobalHStillUnselected        = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type InheritedGate244Audit struct {
	TauEtaOriginKnown             bool
	SourceOperatorsKnown          bool
	OperatorModeAlignmentDerived  bool
	ExteriorRepresentativeDerived bool
	WeakPlaneDerived              bool
	GenerationTextureDerived      bool
	TruthStatement                string
}

type EWSourceOperator struct {
	Slot                   int
	TraceExpression        string
	TauValue               int
	SourceObservable       string
	Decomposition          string
	UsesT1                 bool
	UsesT2                 bool
	UsesT3                 bool
	UsesYPhi               bool
	QuadraticObservable    bool
	LieBasisElement        bool
	NeutralPlaneElement    bool
	SpatialAxisLabel       string
	SpatialAxisDerived     bool
	RepresentativeIfForced string
	ForcedMapRejected      bool
	Verdict                string
}

type OperatorDecompositionAudit struct {
	SourceGate                     string
	LieAlgebraAvailable            string
	Records                        []EWSourceOperator
	TauSequence                    []int
	EWDecompositionTraced          bool
	NeutralEWPlaneDimension        int
	FullContactLieBasisDimension   int
	ThreeTauSlots                  bool
	SlotsAreThreeSU2BasisElements  bool
	SlotsAreQuadraticScalarRecords bool
	QZMixT3AndYPhi                 bool
	MissingT1T2SlotOrigins         bool
	Verdict                        string
}

type DerivationToBladeAudit struct {
	ContactSU2Available               bool
	U1Available                       bool
	SpatialCarrierModes               []string
	CandidateSpatialBivectors         []string
	CandidateSU2Capacity              bool
	ExplicitContactGeneratorMatrices  bool
	CanonicalWeakPlaneDerived         bool
	CanonicalSpatialAxisBasisDerived  bool
	OneToOneDerivationAxisMap         bool
	SpatialBivectorsFormSU2Abstractly bool
	BivectorToFockModePullbackDerived bool
	Verdict                           string
}

type CarrierProjectionAudit struct {
	ScalarObservableToDerivationMap   bool
	DerivationToBladeMap              bool
	BladeToFockAxisMap                bool
	ChainedProjectionDerived          bool
	HypotheticalProjection            string
	HypotheticalProjectionRejected    bool
	ProjectionFailure                 string
	ExteriorRepresentativeConstructed bool
	CandidateExteriorRepresentative   string
	WeakAxisIfProjectionExisted       string
	WeakPlaneIfProjectionExisted      string
	Verdict                           string
}

type GenerationProjectionAudit struct {
	TauSequence                  []int
	GenerationBreakingCapacity   bool
	ScalarToGenerationMapDerived bool
	TrialityCarrierMapDerived    bool
	LieAlgebraChainRelevant      bool
	GenerationOperatorDerived    bool
	GenerationTextureDerived     bool
	Verdict                      string
}

type FirewallAudit struct {
	ForcedQZT3ToAxes             bool
	ForcedSU2ToSpatialAxes       bool
	ForcedExteriorRepresentative bool
	ForcedTrialityMap            bool
	ImportedWeakPlane            bool
	ImportedConnesAlgebra        bool
	ClaimedPhysicalChirality     bool
	ClaimedGlobalH               bool
	ClaimedGenerationTexture     bool
	ClaimedCKMPMNS               bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	OperatorDecompositionTraced   bool
	TauSlotsAreSU2Basis           bool
	DerivationBladeCapacity       bool
	DerivationAxisMapDerived      bool
	CarrierProjectionDerived      bool
	ExteriorRepresentativeDerived bool
	WeakPlaneConditionallyVisible bool
	WeakPlaneDerived              bool
	GenerationBreakingCapacity    bool
	GenerationTextureDerived      bool
	GlobalHDerived                bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate244       InheritedGate244Audit
	OperatorDecomposition OperatorDecompositionAudit
	DerivationBlade       DerivationToBladeAudit
	CarrierProjection     CarrierProjectionAudit
	GenerationProjection  GenerationProjectionAudit
	Firewall              FirewallAudit
	Summary               Summary
	TruthStatement        string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prevRaw, err := characteristicpullback.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate244(prevRaw)
		op := auditOperatorDecomposition()
		blade := auditDerivationToBlade()
		proj := auditCarrierProjection(op, blade)
		gen := auditGenerationProjection(op)
		fw := auditFirewall()
		sum := summarize(op, blade, proj, gen)
		truth := buildTruth(op, blade, proj, gen)
		defaultA = Analysis{PreviousGate244: prev, OperatorDecomposition: op, DerivationBlade: blade, CarrierProjection: proj, GenerationProjection: gen, Firewall: fw, Summary: sum, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate244(a characteristicpullback.Analysis) InheritedGate244Audit {
	return InheritedGate244Audit{
		TauEtaOriginKnown:             a.Summary.TauEtaOriginTraced,
		SourceOperatorsKnown:          a.Origin.ExactOperatorOriginsRecovered,
		OperatorModeAlignmentDerived:  a.Summary.OperatorModeAlignmentDerived,
		ExteriorRepresentativeDerived: a.Summary.ExteriorRepresentativeDerived,
		WeakPlaneDerived:              a.Summary.WeakPlaneDerived,
		GenerationTextureDerived:      a.Summary.GenerationTextureDerived,
		TruthStatement:                a.TruthStatement,
	}
}

func auditOperatorDecomposition() OperatorDecompositionAudit {
	recs := []EWSourceOperator{
		{Slot: 0, TraceExpression: "tau_eta(Q^T Q)", TauValue: 2, SourceObservable: "Q", Decomposition: "Q = T3L + Y_phi", UsesT1: false, UsesT2: false, UsesT3: true, UsesYPhi: true, QuadraticObservable: true, LieBasisElement: false, NeutralPlaneElement: true, SpatialAxisLabel: "", SpatialAxisDerived: false, RepresentativeIfForced: "2 e_1", ForcedMapRejected: true, Verdict: "Q is an electromagnetic neutral combination in span{T3L,Y_phi}; Q^TQ is a scalar quadratic observable, not an su(2) basis generator or spatial axis label"},
		{Slot: 1, TraceExpression: "tau_eta(Z^T Z)", TauValue: -2, SourceObservable: "Z", Decomposition: "Z = T3L - Y_phi", UsesT1: false, UsesT2: false, UsesT3: true, UsesYPhi: true, QuadraticObservable: true, LieBasisElement: false, NeutralPlaneElement: true, SpatialAxisLabel: "", SpatialAxisDerived: false, RepresentativeIfForced: "-2 e_2", ForcedMapRejected: true, Verdict: "Z is the orthogonal neutral combination in span{T3L,Y_phi}; Z^TZ is a scalar quadratic observable, not an su(2) basis generator or spatial axis label"},
		{Slot: 2, TraceExpression: "tau_eta(T3L^T Y_phi)", TauValue: 1, SourceObservable: "T3L/Y_phi mixed pairing", Decomposition: "T3L paired with Y_phi", UsesT1: false, UsesT2: false, UsesT3: true, UsesYPhi: true, QuadraticObservable: true, LieBasisElement: false, NeutralPlaneElement: true, SpatialAxisLabel: "", SpatialAxisDerived: false, RepresentativeIfForced: "e_3", ForcedMapRejected: true, Verdict: "the mixed trace is a bilinear scalar-pairing between T3L and Y_phi, not a third independent derivation or a spatial axis label"},
	}
	return OperatorDecompositionAudit{
		SourceGate:                     "Gate 244 + electroweak scalar/boundary bridge",
		LieAlgebraAvailable:            "contact-preserving su(2)⊕u(1) with basis {T1,T2,T3L,Y_phi}; neutral observables use Q=T3L+Y_phi and Z=T3L-Y_phi",
		Records:                        recs,
		TauSequence:                    []int{2, -2, 1},
		EWDecompositionTraced:          true,
		NeutralEWPlaneDimension:        2,
		FullContactLieBasisDimension:   4,
		ThreeTauSlots:                  true,
		SlotsAreThreeSU2BasisElements:  false,
		SlotsAreQuadraticScalarRecords: true,
		QZMixT3AndYPhi:                 true,
		MissingT1T2SlotOrigins:         true,
		Verdict:                        "The tau_eta slots decompose to scalar quadratic records in the two-dimensional neutral electroweak plane span{T3L,Y_phi}. They are not the three su(2) basis generators {T1,T2,T3}; T1 and T2 do not appear as tau_eta source slots.",
	}
}

func auditDerivationToBlade() DerivationToBladeAudit {
	return DerivationToBladeAudit{
		ContactSU2Available:               true,
		U1Available:                       true,
		SpatialCarrierModes:               []string{"a†_1", "a†_2", "a†_3"},
		CandidateSpatialBivectors:         []string{"e_1∧e_2", "e_2∧e_3", "e_3∧e_1"},
		CandidateSU2Capacity:              true,
		ExplicitContactGeneratorMatrices:  false,
		CanonicalWeakPlaneDerived:         false,
		CanonicalSpatialAxisBasisDerived:  false,
		OneToOneDerivationAxisMap:         false,
		SpatialBivectorsFormSU2Abstractly: true,
		BivectorToFockModePullbackDerived: false,
		Verdict:                           "Spatial bivectors have the abstract capacity to realize an su(2)-like rotation algebra, but the engine still has no derived map from the contact-preserving su(2) generators to a canonical ordered spatial axis/bivector basis on W.",
	}
}

func auditCarrierProjection(op OperatorDecompositionAudit, blade DerivationToBladeAudit) CarrierProjectionAudit {
	derived := op.SlotsAreThreeSU2BasisElements && blade.OneToOneDerivationAxisMap && blade.BivectorToFockModePullbackDerived
	return CarrierProjectionAudit{
		ScalarObservableToDerivationMap:   op.EWDecompositionTraced && !op.SlotsAreThreeSU2BasisElements,
		DerivationToBladeMap:              blade.OneToOneDerivationAxisMap,
		BladeToFockAxisMap:                blade.BivectorToFockModePullbackDerived,
		ChainedProjectionDerived:          derived,
		HypotheticalProjection:            "tau_eta(Q^TQ,Z^TZ,T3L^T Y_phi) -> (2e_1,-2e_2,e_3)",
		HypotheticalProjectionRejected:    true,
		ProjectionFailure:                 "the first link maps tau_eta slots to neutral scalar observables in span{T3L,Y_phi}, not to the three su(2) generators; the second link lacks a canonical su(2)->spatial-axis map",
		ExteriorRepresentativeConstructed: false,
		CandidateExteriorRepresentative:   "omega_tau ?= 2 e_1 - 2 e_2 + e_3",
		WeakAxisIfProjectionExisted:       "a†_3",
		WeakPlaneIfProjectionExisted:      "U={a†_1,a†_2}",
		Verdict:                           "The chained carrier projection theorem fails. The conditional weak-plane roadmap remains visible, but no lawful exterior representative is constructed.",
	}
}

func auditGenerationProjection(op OperatorDecompositionAudit) GenerationProjectionAudit {
	return GenerationProjectionAudit{
		TauSequence:                  op.TauSequence,
		GenerationBreakingCapacity:   true,
		ScalarToGenerationMapDerived: false,
		TrialityCarrierMapDerived:    false,
		LieAlgebraChainRelevant:      false,
		GenerationOperatorDerived:    false,
		GenerationTextureDerived:     false,
		Verdict:                      "The Lie-algebra-to-spatial carrier chain does not supply the separate tau_eta->triality-generation pullback. The signed 1+1+1 capacity survives, but no generation texture is derived.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedQZT3ToAxes:             false,
		ForcedSU2ToSpatialAxes:       false,
		ForcedExteriorRepresentative: false,
		ForcedTrialityMap:            false,
		ImportedWeakPlane:            false,
		ImportedConnesAlgebra:        false,
		ClaimedPhysicalChirality:     false,
		ClaimedGlobalH:               false,
		ClaimedGenerationTexture:     false,
		ClaimedCKMPMNS:               false,
		FiniteCorePolluted:           false,
		Verdict:                      "Gate 245 performs decomposition and isomorphism audits without forcing scalar trace slots onto spatial axes or triality generations.",
	}
}

func summarize(op OperatorDecompositionAudit, blade DerivationToBladeAudit, proj CarrierProjectionAudit, gen GenerationProjectionAudit) Summary {
	return Summary{
		OperatorDecompositionTraced:   op.EWDecompositionTraced,
		TauSlotsAreSU2Basis:           op.SlotsAreThreeSU2BasisElements,
		DerivationBladeCapacity:       blade.CandidateSU2Capacity && blade.SpatialBivectorsFormSU2Abstractly,
		DerivationAxisMapDerived:      blade.OneToOneDerivationAxisMap,
		CarrierProjectionDerived:      proj.ChainedProjectionDerived,
		ExteriorRepresentativeDerived: proj.ExteriorRepresentativeConstructed,
		WeakPlaneConditionallyVisible: proj.WeakPlaneIfProjectionExisted != "",
		WeakPlaneDerived:              false,
		GenerationBreakingCapacity:    gen.GenerationBreakingCapacity,
		GenerationTextureDerived:      false,
		GlobalHDerived:                false,
		Status:                        strings.Join([]string{StatusOperatorDecompositionTraced, StatusNeutralPlaneObstruction, StatusDerivationBladePreflight, StatusDerivationBladeObstruction, StatusCarrierProjectionObstruction, StatusExteriorRepresentativeBlocked, StatusWeakPlaneBlocked, StatusGenerationTextureBlocked, StatusGlobalHStillUnselected}, "\n"),
		NextGate:                      "Gate 246 — electroweak scalar-to-Fock representation functor or explicit physical-chirality seal audit",
		Comment:                       "Gate 245 decomposes the scalar trace origins but finds that they live in the neutral electroweak scalar-observable plane, not in a three-generator su(2) or spatial-axis basis. The missing theorem is now sharper: derive an H_Phi-to-W representation functor, not merely a Lie-algebra analogy.",
	}
}

func buildTruth(op OperatorDecompositionAudit, blade DerivationToBladeAudit, proj CarrierProjectionAudit, gen GenerationProjectionAudit) string {
	return fmt.Sprintf("Gate 245 decomposes tau_eta=%v to the neutral electroweak scalar observables Q=T3L+Y_phi, Z=T3L-Y_phi, and T3L·Y_phi. This proves the source labels are structured, but also proves they are not the three su(2) generators or spatial Fock axes. Spatial bivectors have su(2) capacity (%t), yet no native contact-su(2)->axis map is derived. Therefore the carrier projection %q is rejected, the conditional weak plane %q remains unselected, and the signed generation capacity remains without a triality pullback.", op.TauSequence, blade.CandidateSU2Capacity, proj.CandidateExteriorRepresentative, proj.WeakPlaneIfProjectionExisted)
}
