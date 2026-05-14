// Package lietrialitypullback implements Gate 252:
// Lie Algebra Triality Pullback / Hermitian Q_8vC Neutral 3-Plane Audit.
//
// Gate 251 opened the complex weight-space route: odd-dimensional complex
// neutral kernels are allowed in principle, but the physical Hermitian
// electroweak matrices Q_8vC and Z_8vC were not derived. Gate 252 audits the
// next proposed bridge: use infinitesimal Spin(8) triality to transport the
// known electroweak action from the spinor/Fock side to the vector carrier.
//
// The audit is deliberately conservative. Infinitesimal triality is the right
// kind of representation-theoretic mechanism, but it does not become a
// theorem unless the engine has explicit so(8) generator coordinates on the
// spinor side, an explicit triality automorphism on so(8), and a
// real-structure-compatible action on the complexified vector carrier.
package lietrialitypullback

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/complexweightspacekernel"
)

const (
	AuditID = "GATE252-LIE-TRIALITY-PULLBACK-Q8VC-NEUTRAL-3PLANE-AUDIT"

	StatusInfinitesimalTrialityPreflight = "CONDITIONAL_SUPPORT_INFINITESIMAL_TRIALITY_PREFLIGHT"
	StatusSpinorEWBridgeKnown            = "CONDITIONAL_SUPPORT_SPINOR_EW_BRIDGE_REPRESENTATIONS_INHERITED"
	StatusHermitianCapacityInherited     = "CONDITIONAL_SUPPORT_COMPLEX_HERMITIAN_WEIGHT_CAPACITY_INHERITED"
	StatusSpinorSO8CoordsMissing         = "FAILED_ROUTE_SPINOR_EW_GENERATORS_NOT_SO8_BIVECTOR_COORDINATES"
	StatusTrialityMapMissing             = "FAILED_ROUTE_EXPLICIT_LIE_TRIALITY_AUTOMORPHISM_DERIVATION"
	StatusVectorMatricesMissing          = "FAILED_ROUTE_LIE_TRIALITY_Q8VC_MATRIX_DERIVATION"
	StatusNeutral3PlaneMissing           = "FAILED_ROUTE_LIE_TRIALITY_NEUTRAL_3PLANE_DERIVATION"
	StatusJCompatibilityMissing          = "FAILED_ROUTE_REAL_STRUCTURE_COMPATIBLE_TRIALITY_DERIVATION"
	StatusVTauStillBlocked               = "FAILED_ROUTE_V_TAU_CONSTRUCTION"
	StatusYukawaStillBlocked             = "FAILED_ROUTE_YUKAWA_TEXTURE_DERIVATION"
)

type InheritedGate251Audit struct {
	Complex8VKnown              bool
	HermitianWeightCapacity     bool
	OddComplexKernelCapacity    bool
	NativeHermitianMatrices     bool
	ComplexNeutralKernelDerived bool
	NeutralKernelDim3           bool
	ComplexTrialityArena        bool
	CanonicalTrialityMap        bool
	RealStructureCompatible     bool
	VTauConstructed             bool
	TrialityUnblocked           bool
	YukawaTextureDerived        bool
	TruthStatement              string
}

type InfinitesimalTrialityAudit struct {
	Spin8LieAlgebra              string
	LieAlgebraDimension          int
	TrialityOuterAutomorphism    string
	CanPermuteRepresentations    bool
	ActsOnLieAlgebra             bool
	RequiresExplicitAutomorphism bool
	ExplicitAutomorphismDerived  bool
	CanonicalWithoutChoice       bool
	Verdict                      string
}

type SpinorGeneratorAudit struct {
	RequiredGenerators               []string
	BridgeRepresentationsKnown       bool
	SpinorFockActionKnown            bool
	ScalarBundleActionKnown          bool
	AsSO8BivectorCoordinates         bool
	AsSkewHermitianSpin8Generators   bool
	SuitableForInfinitesimalTriality bool
	Obstruction                      string
	Verdict                          string
}

type TranslationAudit struct {
	InputSpinorGeneratorsAvailable bool
	InfinitesimalTrialityMapKnown  bool
	CanPushT3To8V                  bool
	CanPushYTo8V                   bool
	T3VectorMatrixDerived          bool
	YVectorMatrixDerived           bool
	ManualDictionaryRejected       bool
	Obstruction                    string
	Verdict                        string
}

type HermitianQAudit struct {
	ComplexCarrier             string
	HermitianRule              string
	T3VectorMatrixDerived      bool
	YVectorMatrixDerived       bool
	HT3Constructed             bool
	HYConstructed              bool
	Q8vCConstructed            bool
	Z8vCConstructed            bool
	HermitianMatricesAvailable bool
	Verdict                    string
}

type NeutralKernelAudit struct {
	Definition             string
	Q8vCConstructed        bool
	EigensystemComputed    bool
	KernelDimensionKnown   bool
	KernelComplexDimension int
	ExactlyThree           bool
	ThreePlaneDerived      bool
	DependsOnMissingQ      bool
	Verdict                string
}

type TrialityTransportAudit struct {
	ComplexTrialityArenaKnown      bool
	Neutral3PlaneAvailable         bool
	Canonical8vCTo8sCMapDerived    bool
	NeutralPlaneImageInSpinorKnown bool
	RealStructureJKnownOnSpinor    bool
	RealStructureJKnownOnVector    bool
	CommutesWithJ                  bool
	TransportPhysicallyMeaningful  bool
	Obstruction                    string
	Verdict                        string
}

type VTauAudit struct {
	TauEta                 []int
	NeedsNeutral3Plane     bool
	Neutral3PlaneAvailable bool
	NeedsScalarSlotFrame   bool
	ScalarSlotFrameDerived bool
	Constructed            bool
	TrialityTransportReady bool
	YukawaTextureDerived   bool
	RejectedBecause        string
	Verdict                string
}

type FirewallAudit struct {
	InventedSO8Coordinates bool
	InventedLieTrialityMap bool
	InventedT3VectorMatrix bool
	InventedYVectorMatrix  bool
	InventedQ8vC           bool
	ForcedKernelDim3       bool
	IgnoredJCompatibility  bool
	ConstructedVTauByHand  bool
	InsertedYukawaTexture  bool
	ClaimedCKMPMNS         bool
	PollutedFiniteCore     bool
	Verdict                string
}

type Summary struct {
	InfinitesimalTrialityCapacity bool
	SpinorEWBridgeKnown           bool
	SpinorSO8Coordinates          bool
	ExplicitLieTrialityMap        bool
	VectorEWMatriciesDerived      bool
	Q8vCConstructed               bool
	Neutral3PlaneDerived          bool
	JCompatibleTransport          bool
	VTauConstructed               bool
	TrialityUnblocked             bool
	YukawaTextureDerived          bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate251 InheritedGate251Audit
	Triality        InfinitesimalTrialityAudit
	SpinorInput     SpinorGeneratorAudit
	Translation     TranslationAudit
	HermitianQ      HermitianQAudit
	NeutralKernel   NeutralKernelAudit
	Transport       TrialityTransportAudit
	VTau            VTauAudit
	Firewall        FirewallAudit
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
		prevRaw, err := complexweightspacekernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate251(prevRaw)
		triality := auditInfinitesimalTriality(prev)
		spinor := auditSpinorGenerators()
		translation := auditTranslation(triality, spinor)
		herm := auditHermitianQ(translation)
		kernel := auditNeutralKernel(herm)
		transport := auditTrialityTransport(prev, kernel)
		vtau := auditVTau(kernel, transport)
		firewall := auditFirewall()
		summary := summarize(prev, triality, spinor, translation, herm, kernel, transport, vtau)
		truth := buildTruth(prev, triality, spinor, translation, herm, kernel, transport, vtau)
		defaultA = Analysis{PreviousGate251: prev, Triality: triality, SpinorInput: spinor, Translation: translation, HermitianQ: herm, NeutralKernel: kernel, Transport: transport, VTau: vtau, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate251(a complexweightspacekernel.Analysis) InheritedGate251Audit {
	return InheritedGate251Audit{
		Complex8VKnown:              a.Summary.Complex8VKnown,
		HermitianWeightCapacity:     a.Summary.HermitianWeightCapacity,
		OddComplexKernelCapacity:    a.Summary.OddComplexKernelCapacity,
		NativeHermitianMatrices:     a.Summary.NativeHermitianMatrices,
		ComplexNeutralKernelDerived: a.Summary.ComplexNeutralKernelDerived,
		NeutralKernelDim3:           a.Summary.NeutralKernelDim3,
		ComplexTrialityArena:        a.Summary.ComplexTrialityArena,
		CanonicalTrialityMap:        a.Summary.CanonicalTrialityMap,
		RealStructureCompatible:     a.Summary.RealStructureCompatible,
		VTauConstructed:             a.Summary.VTauConstructed,
		TrialityUnblocked:           a.Summary.TrialityUnblocked,
		YukawaTextureDerived:        a.Summary.YukawaTextureDerived,
		TruthStatement:              a.TruthStatement,
	}
}

func auditInfinitesimalTriality(prev InheritedGate251Audit) InfinitesimalTrialityAudit {
	return InfinitesimalTrialityAudit{
		Spin8LieAlgebra:              "so(8) = Λ²R⁸",
		LieAlgebraDimension:          28,
		TrialityOuterAutomorphism:    "Out(Spin(8)) ≅ S3; infinitesimal action permutes representation realizations of so(8)",
		CanPermuteRepresentations:    prev.ComplexTrialityArena,
		ActsOnLieAlgebra:             true,
		RequiresExplicitAutomorphism: true,
		ExplicitAutomorphismDerived:  false,
		CanonicalWithoutChoice:       false,
		Verdict:                      "Infinitesimal triality is the right representation-theoretic target, but the engine has not derived explicit automorphism matrices or a canonical representative choice.",
	}
}

func auditSpinorGenerators() SpinorGeneratorAudit {
	return SpinorGeneratorAudit{
		RequiredGenerators:               []string{"T3L", "Y_phi"},
		BridgeRepresentationsKnown:       true,
		SpinorFockActionKnown:            true,
		ScalarBundleActionKnown:          true,
		AsSO8BivectorCoordinates:         false,
		AsSkewHermitianSpin8Generators:   false,
		SuitableForInfinitesimalTriality: false,
		Obstruction:                      "T3L and Y_phi are known in bridge/scalar/Fock representation ledgers, but not as explicit coordinates in so(8)=Λ²R⁸ acting on 8_s/8_c.",
		Verdict:                          "The input generators are physically named, but not typed as Spin(8) Lie-algebra elements suitable for infinitesimal-triality transport.",
	}
}

func auditTranslation(t InfinitesimalTrialityAudit, s SpinorGeneratorAudit) TranslationAudit {
	inputOK := s.SuitableForInfinitesimalTriality
	mapOK := t.ExplicitAutomorphismDerived && t.CanonicalWithoutChoice
	canPush := inputOK && mapOK
	return TranslationAudit{
		InputSpinorGeneratorsAvailable: inputOK,
		InfinitesimalTrialityMapKnown:  mapOK,
		CanPushT3To8V:                  canPush,
		CanPushYTo8V:                   canPush,
		T3VectorMatrixDerived:          false,
		YVectorMatrixDerived:           false,
		ManualDictionaryRejected:       true,
		Obstruction:                    "The chain requires both spinor-side so(8) coordinates and an explicit infinitesimal-triality automorphism; both remain missing.",
		Verdict:                        "No vector electroweak matrices are derived; translating by representation names alone would be a type error.",
	}
}

func auditHermitianQ(t TranslationAudit) HermitianQAudit {
	matrices := t.T3VectorMatrixDerived && t.YVectorMatrixDerived
	return HermitianQAudit{
		ComplexCarrier:             "8_vC = 8_v ⊗_R C",
		HermitianRule:              "H = iA for a real skew generator A",
		T3VectorMatrixDerived:      t.T3VectorMatrixDerived,
		YVectorMatrixDerived:       t.YVectorMatrixDerived,
		HT3Constructed:             false,
		HYConstructed:              false,
		Q8vCConstructed:            false,
		Z8vCConstructed:            false,
		HermitianMatricesAvailable: matrices,
		Verdict:                    "Hermitian construction is lawful in principle, but Q_8vC and Z_8vC cannot be built without the translated vector matrices.",
	}
}

func auditNeutralKernel(h HermitianQAudit) NeutralKernelAudit {
	computed := h.Q8vCConstructed
	return NeutralKernelAudit{
		Definition:             "ker(Q_8vC) = {v ∈ 8_vC | Q_8vC v = 0}",
		Q8vCConstructed:        h.Q8vCConstructed,
		EigensystemComputed:    computed,
		KernelDimensionKnown:   computed,
		KernelComplexDimension: 0,
		ExactlyThree:           false,
		ThreePlaneDerived:      false,
		DependsOnMissingQ:      true,
		Verdict:                "The complex neutral kernel is still not computed; the dimensionality remains an unverified target, not a theorem.",
	}
}

func auditTrialityTransport(prev InheritedGate251Audit, k NeutralKernelAudit) TrialityTransportAudit {
	return TrialityTransportAudit{
		ComplexTrialityArenaKnown:      prev.ComplexTrialityArena,
		Neutral3PlaneAvailable:         k.ThreePlaneDerived,
		Canonical8vCTo8sCMapDerived:    false,
		NeutralPlaneImageInSpinorKnown: false,
		RealStructureJKnownOnSpinor:    true,
		RealStructureJKnownOnVector:    false,
		CommutesWithJ:                  false,
		TransportPhysicallyMeaningful:  false,
		Obstruction:                    "Even after complexification, triality transport needs explicit maps and compatibility with the real structures on vector and spinor carriers.",
		Verdict:                        "No canonical J-compatible transport of a neutral vector plane into the Fock/spinor carrier is derived.",
	}
}

func auditVTau(k NeutralKernelAudit, t TrialityTransportAudit) VTauAudit {
	return VTauAudit{
		TauEta:                 []int{2, -2, 1},
		NeedsNeutral3Plane:     true,
		Neutral3PlaneAvailable: k.ThreePlaneDerived,
		NeedsScalarSlotFrame:   true,
		ScalarSlotFrameDerived: false,
		Constructed:            false,
		TrialityTransportReady: t.TransportPhysicallyMeaningful,
		YukawaTextureDerived:   false,
		RejectedBecause:        "v_tau requires a derived neutral 3-plane, a canonical scalar-slot frame in that plane, and J-compatible triality transport; none are derived.",
		Verdict:                "tau_eta retains generation-texture capacity but is not promoted into a vector or spinor operator.",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		InventedSO8Coordinates: false,
		InventedLieTrialityMap: false,
		InventedT3VectorMatrix: false,
		InventedYVectorMatrix:  false,
		InventedQ8vC:           false,
		ForcedKernelDim3:       false,
		IgnoredJCompatibility:  false,
		ConstructedVTauByHand:  false,
		InsertedYukawaTexture:  false,
		ClaimedCKMPMNS:         false,
		PollutedFiniteCore:     false,
		Verdict:                "No Spin(8) coordinates, triality maps, vector electroweak matrices, neutral kernels, or Yukawa textures are inserted by hand.",
	}
}

func summarize(prev InheritedGate251Audit, tr InfinitesimalTrialityAudit, s SpinorGeneratorAudit, x TranslationAudit, h HermitianQAudit, k NeutralKernelAudit, tt TrialityTransportAudit, v VTauAudit) Summary {
	return Summary{
		InfinitesimalTrialityCapacity: tr.ActsOnLieAlgebra && tr.CanPermuteRepresentations,
		SpinorEWBridgeKnown:           s.BridgeRepresentationsKnown,
		SpinorSO8Coordinates:          s.AsSO8BivectorCoordinates,
		ExplicitLieTrialityMap:        tr.ExplicitAutomorphismDerived,
		VectorEWMatriciesDerived:      x.T3VectorMatrixDerived && x.YVectorMatrixDerived,
		Q8vCConstructed:               h.Q8vCConstructed,
		Neutral3PlaneDerived:          k.ThreePlaneDerived,
		JCompatibleTransport:          tt.TransportPhysicallyMeaningful,
		VTauConstructed:               v.Constructed,
		TrialityUnblocked:             v.Constructed && tt.TransportPhysicallyMeaningful,
		YukawaTextureDerived:          v.YukawaTextureDerived,
		Status:                        strings.Join([]string{StatusInfinitesimalTrialityPreflight, StatusSpinorEWBridgeKnown, StatusHermitianCapacityInherited, StatusSpinorSO8CoordsMissing, StatusTrialityMapMissing, StatusVectorMatricesMissing, StatusNeutral3PlaneMissing, StatusJCompatibilityMissing, StatusVTauStillBlocked, StatusYukawaStillBlocked}, ";"),
		NextGate:                      "derive explicit Spin(8)/so(8) coordinates for the electroweak generators or a faithful action functor from the Fock/scalar bridge into so(8)",
		Comment:                       "Gate 252 validates infinitesimal triality as the correct kind of bridge, but the bridge cannot be used until its domain data and explicit automorphism are derived.",
	}
}

func buildTruth(prev InheritedGate251Audit, tr InfinitesimalTrialityAudit, s SpinorGeneratorAudit, x TranslationAudit, h HermitianQAudit, k NeutralKernelAudit, tt TrialityTransportAudit, v VTauAudit) string {
	return fmt.Sprintf("Gate 252 tests the proposed infinitesimal Spin(8) triality pullback. The route is well typed only if T3L/Y_phi are explicit so(8) spinor generators and an explicit triality automorphism transports them to 8_vC. Current data provide complex/Hermitian capacity and bridge-level electroweak names, but not the required so(8) coordinates or J-compatible triality map. Therefore Q_8vC, its neutral 3-plane, v_tau, and Yukawa textures remain un-derived. prevComplex=%t trialityCapacity=%t spinorSO8=%t map=%t Q=%t kernel3=%t transport=%t vtau=%t", prev.Complex8VKnown, tr.ActsOnLieAlgebra, s.AsSO8BivectorCoordinates, x.InfinitesimalTrialityMapKnown, h.Q8vCConstructed, k.ExactlyThree, tt.TransportPhysicallyMeaningful, v.Constructed)
}
