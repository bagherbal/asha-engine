// Package characteristicpullback implements Gate 244:
// Characteristic Class / Operator-to-Mode Pullback Audit.
//
// Gate 243 proved that the Clifford action c:Λ*(W)->End(S_C) exists, but
// tau_eta=(2,-2,1) is not in the action domain because it is a scalar-bundle
// eta-graded trace sequence, not an exterior form representative.  Gate 244
// traces the exact source operators behind the three tau_eta entries and asks
// whether those source labels lawfully attach the entries to the spatial Fock
// modes e1,e2,e3.  The result is intentionally strict: the origin operators are
// exact and stable, but they live in the sealed scalar curvature-observable
// algebra.  They do not carry spatial-mode labels, basis-blade coefficients, or
// a characteristic-class representative in Λ*(W).  Therefore no exterior form
// representative is constructed and the weak/generation sieves remain blocked.
package characteristicpullback

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE244-CHARACTERISTIC-CLASS-OPERATOR-TO-MODE-PULLBACK-AUDIT"

	StatusTauEtaOriginTraced          = "CONDITIONAL_SUPPORT_TAU_ETA_OPERATOR_ORIGIN_TRACED"
	StatusNativeTraceSequenceStable   = "CONDITIONAL_SUPPORT_NATIVE_TRACE_SEQUENCE_STABLE"
	StatusFailedOperatorModeAlignment = "FAILED_ROUTE_SOURCE_OPERATORS_NOT_SPATIAL_FOCK_MODES"
	StatusFailedExteriorFormRep       = "FAILED_ROUTE_EXTERIOR_FORM_REPRESENTATIVE_DERIVATION"
	StatusFailedCharacteristicClass   = "FAILED_ROUTE_CHARACTERISTIC_CLASS_REPRESENTATIVE_DERIVATION"
	StatusFailedWeakPlane             = "FAILED_ROUTE_CHARACTERISTIC_PULLBACK_WEAK_PLANE_SELECTION"
	StatusFailedGenerationTexture     = "FAILED_ROUTE_CHARACTERISTIC_PULLBACK_GENERATION_TEXTURE"
	StatusGlobalHStillUnselected      = "FAILED_ROUTE_GLOBAL_H_SUMMAND_STILL_UNSELECTED"
)

type InheritedGate243Audit struct {
	CliffordActionAvailable bool
	TauEtaInCliffordDomain  bool
	EndomorphismConstructed bool
	TruthStatement          string
}

type InheritedScalarClassAudit struct {
	SourceGate           string
	EtaGradedFunctional  string
	EtaEvenDomain        string
	StableNativeDegrees  bool
	ScalarBundleOnly     bool
	ContinuumFormDerived bool
}

type TraceOriginRecord struct {
	Slot                         int
	Name                         string
	Expression                   string
	Value                        int
	ExpectedRational             string
	StableInteger                bool
	SourceOperatorFamily         string
	LivesOnScalarBundle          bool
	LivesOnFockCarrierW          bool
	IsExteriorFormCoefficient    bool
	SpatialAxisLabelDerived      bool
	TrialityGenerationLabelKnown bool
	CandidateAxisIfForced        string
	ForcedCandidateRejected      bool
	Verdict                      string
}

type OperatorOriginTraceAudit struct {
	SourceGate                        string
	FunctionalExpression              string
	FunctionalDomain                  string
	Records                           []TraceOriginRecord
	Sequence                          []int
	StableNativeDegrees               bool
	ExactOperatorOriginsRecovered     bool
	OperatorsAreCurvatureObservables  bool
	OperatorsAreSpatialModeProjectors bool
	OperatorsAreBasisBlades           bool
	Verdict                           string
}

type SpatialModeAlignmentAudit struct {
	SpatialAxes                         []string
	TraceSlots                          []string
	ModeSlotCountCompatible             bool
	NativeOperatorToModeMapDerived      bool
	QZT3YInherentlyLinkToSpatialAxes    bool
	OperatorDefinitionsUseFockModes     bool
	ScalarBundleToFockProjectionDerived bool
	ManualMapCandidate                  string
	ManualMapRejected                   bool
	AlignmentVerdict                    string
}

type CharacteristicRepresentativeAudit struct {
	CharacteristicClassLanguageAvailable bool
	FiniteEtaTraceFunctional             bool
	ChernCharacterRepresentativeDerived  bool
	PontryaginFormRepresentativeDerived  bool
	ExteriorGradeKnown                   bool
	BasisBladeLabelsKnown                bool
	NormalizationDerived                 bool
	RepresentativeConstructed            bool
	HypotheticalRepresentative           string
	HypotheticalRepresentativeRejected   bool
	RejectionReason                      string
	Verdict                              string
}

type WeakPlaneOutcome struct {
	InheritedConditionalAxis      string
	InheritedConditionalPlane     string
	ExteriorRepresentativeDerived bool
	SpatialModeAlignmentDerived   bool
	AxisTagged                    string
	WeakPlaneSelected             string
	S3DegeneracyBroken            bool
	PhysicalWeakPlaneDerived      bool
	GlobalHSummandDerived         bool
	Verdict                       string
}

type GenerationOutcome struct {
	Sequence                     []int
	DistinctEigenvalueCapacity   bool
	CharacteristicRepresentative bool
	TrialityCarrierMapDerived    bool
	GenerationOperatorDerived    bool
	GenerationTextureDerived     bool
	CKMPMNSDerived               bool
	Verdict                      string
}

type FirewallAudit struct {
	ForcedOperatorToModeMap      bool
	ForcedExteriorRepresentative bool
	ForcedCharacteristicClass    bool
	ForcedTrialityMap            bool
	ImportedWeakPlane            bool
	ImportedGenerationTexture    bool
	PromotedScalarTraceToMatrix  bool
	ClaimedPhysicalChirality     bool
	ClaimedGlobalH               bool
	ClaimedCKMPMNS               bool
	ClaimedFermionMasses         bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type Summary struct {
	TauEtaOriginTraced            bool
	NativeSequenceStable          bool
	OperatorModeAlignmentDerived  bool
	ExteriorRepresentativeDerived bool
	CharacteristicClassDerived    bool
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
	PreviousCliffordPullback InheritedGate243Audit
	ScalarFundamentalClass   InheritedScalarClassAudit
	Origin                   OperatorOriginTraceAudit
	SpatialAlignment         SpatialModeAlignmentAudit
	CharacteristicRep        CharacteristicRepresentativeAudit
	WeakPlane                WeakPlaneOutcome
	Generation               GenerationOutcome
	Firewall                 FirewallAudit
	Summary                  Summary
	TruthStatement           string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev := inheritedGate243()
		sfc := inheritedScalarClass()
		origin := auditOrigin(sfc)
		spatial := auditSpatialAlignment(origin)
		char := auditCharacteristicRepresentative(sfc)
		weak := auditWeak(spatial, char)
		gen := auditGeneration(origin, char)
		fw := auditFirewall()
		summary := summarize(origin, spatial, char, weak, gen)
		truth := buildTruth(origin, char, weak, gen)
		defaultA = Analysis{PreviousCliffordPullback: prev, ScalarFundamentalClass: sfc, Origin: origin, SpatialAlignment: spatial, CharacteristicRep: char, WeakPlane: weak, Generation: gen, Firewall: fw, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritedGate243() InheritedGate243Audit {
	return InheritedGate243Audit{
		CliffordActionAvailable: true,
		TauEtaInCliffordDomain:  false,
		EndomorphismConstructed: false,
		TruthStatement:          "Gate 243 confirmed that Clifford multiplication c:Λ*(W)->End(S_C) exists, but tau_eta=(2,-2,1) remains a scalar trace functional outside the Clifford-action domain.",
	}
}

func inheritedScalarClass() InheritedScalarClassAudit {
	return InheritedScalarClassAudit{
		SourceGate:           "Gate 193 — scalar fundamental class / eta-graded finite trace",
		EtaGradedFunctional:  "tau_eta(O)=Tr_HPhi(eta O)",
		EtaEvenDomain:        "eta-even scalar curvature-observable algebra generated by Q^TQ, Z^TZ, T_a^T T_b, and T3L^T Y_phi",
		StableNativeDegrees:  true,
		ScalarBundleOnly:     true,
		ContinuumFormDerived: false,
	}
}

func auditOrigin(sfc InheritedScalarClassAudit) OperatorOriginTraceAudit {
	recs := []TraceOriginRecord{
		{Slot: 0, Name: "neutral electromagnetic split", Expression: "tau_eta(Q^T Q)", Value: 2, ExpectedRational: "2", StableInteger: true, SourceOperatorFamily: "neutral electromagnetic scalar-curvature observable Q^TQ on H_Phi", LivesOnScalarBundle: true, LivesOnFockCarrierW: false, IsExteriorFormCoefficient: false, SpatialAxisLabelDerived: false, TrialityGenerationLabelKnown: false, CandidateAxisIfForced: "a†_1", ForcedCandidateRejected: true, Verdict: "exact eta-graded scalar-bundle trace record; not a spatial Fock-mode projector, exterior blade coefficient, or generation label"},
		{Slot: 1, Name: "neutral Z split", Expression: "tau_eta(Z^T Z)", Value: -2, ExpectedRational: "-2", StableInteger: true, SourceOperatorFamily: "neutral Z scalar-curvature observable Z^TZ on H_Phi", LivesOnScalarBundle: true, LivesOnFockCarrierW: false, IsExteriorFormCoefficient: false, SpatialAxisLabelDerived: false, TrialityGenerationLabelKnown: false, CandidateAxisIfForced: "a†_2", ForcedCandidateRejected: true, Verdict: "exact eta-graded scalar-bundle trace record; not a spatial Fock-mode projector, exterior blade coefficient, or generation label"},
		{Slot: 2, Name: "neutral mixed pairing", Expression: "tau_eta(T3L^T Y_phi)", Value: 1, ExpectedRational: "1", StableInteger: true, SourceOperatorFamily: "neutral mixed scalar pairing T3L^T Y_phi on H_Phi", LivesOnScalarBundle: true, LivesOnFockCarrierW: false, IsExteriorFormCoefficient: false, SpatialAxisLabelDerived: false, TrialityGenerationLabelKnown: false, CandidateAxisIfForced: "a†_3", ForcedCandidateRejected: true, Verdict: "exact eta-graded scalar-bundle trace record; not a spatial Fock-mode projector, exterior blade coefficient, or generation label"},
	}
	seq := []int{2, -2, 1}
	return OperatorOriginTraceAudit{
		SourceGate:                        sfc.SourceGate,
		FunctionalExpression:              sfc.EtaGradedFunctional,
		FunctionalDomain:                  sfc.EtaEvenDomain,
		Records:                           recs,
		Sequence:                          seq,
		StableNativeDegrees:               sfc.StableNativeDegrees,
		ExactOperatorOriginsRecovered:     true,
		OperatorsAreCurvatureObservables:  true,
		OperatorsAreSpatialModeProjectors: false,
		OperatorsAreBasisBlades:           false,
		Verdict:                           "The exact origin of tau_eta=(2,-2,1) is recovered: eta-graded traces of Q^TQ, Z^TZ, and T3L^T Y_phi on the sealed scalar bundle. These are curvature-observable labels, not spatial Fock-mode labels.",
	}
}

func auditSpatialAlignment(origin OperatorOriginTraceAudit) SpatialModeAlignmentAudit {
	return SpatialModeAlignmentAudit{
		SpatialAxes:                         []string{"a†_1", "a†_2", "a†_3"},
		TraceSlots:                          []string{"Q^TQ", "Z^TZ", "T3L^T Y_phi"},
		ModeSlotCountCompatible:             len(origin.Sequence) == 3,
		NativeOperatorToModeMapDerived:      false,
		QZT3YInherentlyLinkToSpatialAxes:    false,
		OperatorDefinitionsUseFockModes:     false,
		ScalarBundleToFockProjectionDerived: false,
		ManualMapCandidate:                  "Q^TQ -> a†_1, Z^TZ -> a†_2, T3L^T Y_phi -> a†_3",
		ManualMapRejected:                   true,
		AlignmentVerdict:                    "The three trace slots and the three spatial modes have matching cardinality only. Their source definitions live in the scalar bundle, so aligning them with a†_1,a†_2,a†_3 would be hand-matching.",
	}
}

func auditCharacteristicRepresentative(sfc InheritedScalarClassAudit) CharacteristicRepresentativeAudit {
	return CharacteristicRepresentativeAudit{
		CharacteristicClassLanguageAvailable: true,
		FiniteEtaTraceFunctional:             sfc.ScalarBundleOnly && sfc.StableNativeDegrees,
		ChernCharacterRepresentativeDerived:  false,
		PontryaginFormRepresentativeDerived:  false,
		ExteriorGradeKnown:                   false,
		BasisBladeLabelsKnown:                false,
		NormalizationDerived:                 false,
		RepresentativeConstructed:            false,
		HypotheticalRepresentative:           "omega_tau ?= 2 e_1 - 2 e_2 + e_3 or its dual",
		HypotheticalRepresentativeRejected:   true,
		RejectionReason:                      "The scalar trace records do not derive an exterior grade, basis blade labels, scalar-to-Fock projection, or characteristic-class form representative. The displayed omega_tau would import exactly the missing slot labelling.",
		Verdict:                              "A finite eta-graded functional exists, but a Chern/Pontryagin/exterior representative of tau_eta in Lambda*(W) is not derived. Therefore tau_eta still cannot enter the Clifford action domain.",
	}
}

func auditWeak(sp SpatialModeAlignmentAudit, ch CharacteristicRepresentativeAudit) WeakPlaneOutcome {
	return WeakPlaneOutcome{
		InheritedConditionalAxis:      "a†_3",
		InheritedConditionalPlane:     "U={a†_1,a†_2}",
		ExteriorRepresentativeDerived: ch.RepresentativeConstructed,
		SpatialModeAlignmentDerived:   sp.NativeOperatorToModeMapDerived,
		AxisTagged:                    "",
		WeakPlaneSelected:             "",
		S3DegeneracyBroken:            false,
		PhysicalWeakPlaneDerived:      false,
		GlobalHSummandDerived:         false,
		Verdict:                       "If omega_tau were lawfully derived, the inherited |tau_eta|=(2,2,1) logic would tag a†_3 and select U={a†_1,a†_2}. Gate 244 does not derive omega_tau or the operator-to-mode map, so no weak plane is selected.",
	}
}

func auditGeneration(origin OperatorOriginTraceAudit, ch CharacteristicRepresentativeAudit) GenerationOutcome {
	return GenerationOutcome{
		Sequence:                     append([]int(nil), origin.Sequence...),
		DistinctEigenvalueCapacity:   true,
		CharacteristicRepresentative: ch.RepresentativeConstructed,
		TrialityCarrierMapDerived:    false,
		GenerationOperatorDerived:    false,
		GenerationTextureDerived:     false,
		CKMPMNSDerived:               false,
		Verdict:                      "The signed sequence still has the exact 1+1+1 capacity for generation breaking, but no characteristic representative or triality carrier map is derived, so no generation operator or texture is produced.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedOperatorToModeMap:      false,
		ForcedExteriorRepresentative: false,
		ForcedCharacteristicClass:    false,
		ForcedTrialityMap:            false,
		ImportedWeakPlane:            false,
		ImportedGenerationTexture:    false,
		PromotedScalarTraceToMatrix:  false,
		ClaimedPhysicalChirality:     false,
		ClaimedGlobalH:               false,
		ClaimedCKMPMNS:               false,
		ClaimedFermionMasses:         false,
		FiniteCorePolluted:           false,
		Verdict:                      "Gate 244 traces tau_eta origins without assigning scalar trace slots to spatial or triality carriers by hand.",
	}
}

func summarize(origin OperatorOriginTraceAudit, sp SpatialModeAlignmentAudit, ch CharacteristicRepresentativeAudit, weak WeakPlaneOutcome, gen GenerationOutcome) Summary {
	return Summary{
		TauEtaOriginTraced:            origin.ExactOperatorOriginsRecovered,
		NativeSequenceStable:          origin.StableNativeDegrees,
		OperatorModeAlignmentDerived:  sp.NativeOperatorToModeMapDerived,
		ExteriorRepresentativeDerived: ch.RepresentativeConstructed,
		CharacteristicClassDerived:    ch.ChernCharacterRepresentativeDerived || ch.PontryaginFormRepresentativeDerived,
		WeakPlaneConditionallyVisible: weak.InheritedConditionalPlane != "",
		WeakPlaneDerived:              weak.PhysicalWeakPlaneDerived,
		GenerationBreakingCapacity:    gen.DistinctEigenvalueCapacity,
		GenerationTextureDerived:      gen.GenerationTextureDerived,
		GlobalHDerived:                false,
		Status: strings.Join([]string{
			StatusTauEtaOriginTraced,
			StatusNativeTraceSequenceStable,
			StatusFailedOperatorModeAlignment,
			StatusFailedExteriorFormRep,
			StatusFailedCharacteristicClass,
			StatusFailedWeakPlane,
			StatusFailedGenerationTexture,
			StatusGlobalHStillUnselected,
		}, "; "),
		NextGate: "Gate 245 — scalar-to-Fock tensor lift / H_Phi to W carrier projection audit",
		Comment:  "Gate 244 recovers the exact source operators behind tau_eta, but those operators are scalar curvature observables rather than spatial Fock modes. A lawful carrier projection is still the missing theorem.",
	}
}

func buildTruth(origin OperatorOriginTraceAudit, ch CharacteristicRepresentativeAudit, weak WeakPlaneOutcome, gen GenerationOutcome) string {
	return fmt.Sprintf("Gate 244 traces tau_eta=%v to eta-graded scalar-bundle traces of Q^TQ, Z^TZ, and T3L^T Y_phi. The sequence is exact and stable, but the origin operators are curvature observables on H_Phi, not spatial Fock-mode projectors or exterior blades. Therefore the candidate representative %q is rejected as hand-labelled, the weak plane %q remains conditional only, and the generation spectrum %v remains a capacity rather than a derived texture.", origin.Sequence, ch.HypotheticalRepresentative, weak.InheritedConditionalPlane, gen.Sequence)
}
