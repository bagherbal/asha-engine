// Package complexweightspacekernel implements Gate 251:
// Complex Weight-Space Decomposition / 8_{vC} Neutral Kernel Audit.
//
// Gate 250 proved that a real Clifford-bivector adjoint action on 8_v is
// skew-adjoint and therefore cannot yield an exact three-dimensional real
// kernel. Gate 251 tests the correct quantum-mechanical pivot: complexify the
// vector carrier, convert skew derivations into Hermitian weight operators, and
// ask whether a complex neutral three-plane and a Spin(8) vector-to-spinor
// transport become available.
//
// The audit is deliberately strict. Complexification removes the even-rank
// obstruction in principle, but it does not manufacture the missing physical
// electroweak matrices Q_8vC and Z_8vC. It also records a subtle triality fact:
// over C, the three eight-dimensional Spin(8) modules are related by the outer
// triality automorphism, but a concrete vector-to-spinor map is not canonical
// without an explicit automorphism, basis, and real-structure compatibility.
package complexweightspacekernel

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/adjointbivectoraction"
)

const (
	AuditID = "GATE251-COMPLEX-WEIGHT-SPACE-8VC-NEUTRAL-KERNEL-AUDIT"

	Status8VComplexification         = "CONDITIONAL_SUPPORT_8V_COMPLEXIFICATION_PREFLIGHT"
	StatusHermitianWeightCapacity    = "CONDITIONAL_SUPPORT_HERMITIAN_WEIGHT_SPACE_CAPACITY"
	StatusOddComplexKernelCapacity   = "CONDITIONAL_SUPPORT_ODD_COMPLEX_KERNEL_CAPACITY"
	StatusNativeHermitianUnavailable = "FAILED_ROUTE_NATIVE_HERMITIAN_Q8VC_MATRICES_UNAVAILABLE"
	StatusNeutral3PlaneBlocked       = "FAILED_ROUTE_COMPLEX_NEUTRAL_3PLANE_DERIVATION"
	StatusComplexTrialityPreflight   = "CONDITIONAL_SUPPORT_COMPLEX_TRIALITY_ARENA_PREFLIGHT"
	StatusCanonicalTrialityBlocked   = "FAILED_ROUTE_CANONICAL_COMPLEX_TRIALITY_ISOMORPHISM"
	StatusRealStructureBlocked       = "FAILED_ROUTE_REAL_STRUCTURE_COMPATIBILITY_DERIVATION"
	StatusVTauStillBlocked           = "FAILED_ROUTE_COMPLEX_WEIGHT_V_TAU_CONSTRUCTION"
	StatusYukawaStillBlocked         = "FAILED_ROUTE_YUKAWA_TEXTURE_DERIVATION"
)

type InheritedGate250Audit struct {
	Carrier8VKnown              bool
	CliffordAdjointAvailable    bool
	CandidateMatricesComputable bool
	EWBivectorsRetrieved        bool
	Q8VMatrixDerived            bool
	NeutralKernelDerived        bool
	NeutralKernelDim3           bool
	RealBivector3KernelPossible bool
	VTauConstructed             bool
	TrialityUnblocked           bool
	YukawaTextureDerived        bool
	TruthStatement              string
}

type ComplexificationAudit struct {
	RealCarrierName         string
	ComplexCarrierName      string
	RealDimension           int
	ComplexDimension        int
	UnderlyingRealDimension int
	ComplexificationNative  bool
	EvenRankObstructionLift bool
	Verdict                 string
}

type HermitianPreflightAudit struct {
	RealSkewGeneratorAvailable           bool
	HermitianConversion                  string
	HermitianOperatorsHaveRealSpectrum   bool
	OddWeightSpacesAllowed               bool
	CandidateSimpleBlade                 string
	CandidateSimpleBladeKernelComplexDim int
	PhysicalQHermitianDerived            bool
	PhysicalZHermitianDerived            bool
	Verdict                              string
}

type CartanWeightAudit struct {
	RequiredOperators              []string
	CartanCommutingPairDerived     bool
	Q8vCMatrixDerived              bool
	Z8vCMatrixDerived              bool
	SimultaneouslyDiagonal         bool
	WeightSpectrumDerived          bool
	ManualChargeAssignmentRejected bool
	Obstruction                    string
	Verdict                        string
}

type ComplexNeutralKernelAudit struct {
	Definition               string
	Computed                 bool
	DimensionKnown           bool
	Dimension                int
	ExactlyThreeComplexDim   bool
	OddDimAllowedInPrinciple bool
	DependsOnMissingQ8vC     bool
	BindingFailure           string
	Verdict                  string
}

type ComplexTrialityAudit struct {
	Spin8TrialityOverC                bool
	Modules                           []string
	SameComplexDimension              bool
	OuterAutomorphismRequired         bool
	CanonicalUntwistedIsomorphism     bool
	NeutralKernelAvailable            bool
	MapNeutralKernelToSpinor          bool
	RealStructureCompatibilityChecked bool
	CompatibleWithJ                   bool
	Obstruction                       string
	Verdict                           string
}

type VTauAudit struct {
	TauEta                 []int
	NeedsNeutral3Plane     bool
	Neutral3PlaneAvailable bool
	NeedsScalarSlotFrame   bool
	ScalarSlotFrameDerived bool
	Constructed            bool
	WouldFeedTriality      bool
	RejectedBecause        string
	Verdict                string
}

type FirewallAudit struct {
	InventedQ8vC                 bool
	InventedZ8vC                 bool
	AssignedComplexWeightsByHand bool
	ForcedKernelDim3             bool
	InventedTrialityIsomorphism  bool
	IgnoredRealStructure         bool
	ConstructedVTauByHand        bool
	InsertedYukawaTexture        bool
	ClaimedCKMPMNS               bool
	PollutedFiniteCore           bool
	Verdict                      string
}

type Summary struct {
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
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	PreviousGate250 InheritedGate250Audit
	ComplexCarrier  ComplexificationAudit
	Hermitian       HermitianPreflightAudit
	Cartan          CartanWeightAudit
	NeutralKernel   ComplexNeutralKernelAudit
	Triality        ComplexTrialityAudit
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
		prevRaw, err := adjointbivectoraction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate250(prevRaw)
		carrier := auditComplexification(prev)
		herm := auditHermitianPreflight(prev)
		cartan := auditCartanWeight(herm)
		kernel := auditComplexNeutralKernel(cartan)
		triality := auditComplexTriality(kernel)
		vtau := auditVTau(kernel, triality)
		firewall := auditFirewall()
		summary := summarize(carrier, herm, cartan, kernel, triality, vtau)
		truth := buildTruth(prev, carrier, herm, cartan, kernel, triality, vtau)
		defaultA = Analysis{PreviousGate250: prev, ComplexCarrier: carrier, Hermitian: herm, Cartan: cartan, NeutralKernel: kernel, Triality: triality, VTau: vtau, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate250(a adjointbivectoraction.Analysis) InheritedGate250Audit {
	return InheritedGate250Audit{
		Carrier8VKnown:              a.Carrier.Grade1CarrierKnown && a.Carrier.VectorDimension == 8,
		CliffordAdjointAvailable:    a.Summary.CliffordAdjointAvailable,
		CandidateMatricesComputable: a.Summary.CandidateMatricesComputable,
		EWBivectorsRetrieved:        a.Summary.EWBivectorsRetrieved,
		Q8VMatrixDerived:            a.Summary.Q8VMatrixDerived,
		NeutralKernelDerived:        a.Summary.NeutralKernelDerived,
		NeutralKernelDim3:           a.Summary.NeutralKernelDim3,
		RealBivector3KernelPossible: a.Summary.RealBivector3KernelPossible,
		VTauConstructed:             a.Summary.VTauConstructed,
		TrialityUnblocked:           a.Summary.TrialityUnblocked,
		YukawaTextureDerived:        a.Summary.YukawaTextureDerived,
		TruthStatement:              a.TruthStatement,
	}
}

func auditComplexification(prev InheritedGate250Audit) ComplexificationAudit {
	return ComplexificationAudit{
		RealCarrierName:         "8_v",
		ComplexCarrierName:      "8_vC = 8_v ⊗_R C",
		RealDimension:           8,
		ComplexDimension:        8,
		UnderlyingRealDimension: 16,
		ComplexificationNative:  prev.Carrier8VKnown,
		EvenRankObstructionLift: true,
		Verdict:                 "complexifying the native Spin(8) vector carrier is lawful and moves the neutral-kernel question from real skew kernels to complex Hermitian weight spaces",
	}
}

func auditHermitianPreflight(prev InheritedGate250Audit) HermitianPreflightAudit {
	return HermitianPreflightAudit{
		RealSkewGeneratorAvailable:           prev.CliffordAdjointAvailable && prev.CandidateMatricesComputable,
		HermitianConversion:                  "for a real skew-adjoint generator A, H = i A is Hermitian on the complexified carrier",
		HermitianOperatorsHaveRealSpectrum:   true,
		OddWeightSpacesAllowed:               true,
		CandidateSimpleBlade:                 "i R(e1∧e2) has real weights (+2,-2,0×6) over C in the diagnostic normalization",
		CandidateSimpleBladeKernelComplexDim: 6,
		PhysicalQHermitianDerived:            false,
		PhysicalZHermitianDerived:            false,
		Verdict:                              "complex Hermitian weight spaces can have arbitrary multiplicities, including odd dimensions; however the physical Q and Z Hermitian operators are still not derived",
	}
}

func auditCartanWeight(herm HermitianPreflightAudit) CartanWeightAudit {
	return CartanWeightAudit{
		RequiredOperators:              []string{"H_T3 = i R(T3L)", "H_Y = i R(Y_phi)", "Q_8vC = H_T3 + H_Y", "Z_8vC = H_T3 - H_Y"},
		CartanCommutingPairDerived:     false,
		Q8vCMatrixDerived:              herm.PhysicalQHermitianDerived,
		Z8vCMatrixDerived:              herm.PhysicalZHermitianDerived,
		SimultaneouslyDiagonal:         false,
		WeightSpectrumDerived:          false,
		ManualChargeAssignmentRejected: true,
		Obstruction:                    "Gate 250 did not derive T3L or Y_phi as vector-carrier generators; complexification changes the field of scalars, not the availability of the physical electroweak matrices",
		Verdict:                        "the complex weight-space route is well typed, but the native Cartan matrices on 8_vC remain unavailable",
	}
}

func auditComplexNeutralKernel(c CartanWeightAudit) ComplexNeutralKernelAudit {
	computed := c.Q8vCMatrixDerived
	return ComplexNeutralKernelAudit{
		Definition:               "ker(Q_8vC) = {v ∈ 8_v⊗C | Q_8vC v = 0}",
		Computed:                 computed,
		DimensionKnown:           false,
		Dimension:                0,
		ExactlyThreeComplexDim:   false,
		OddDimAllowedInPrinciple: true,
		DependsOnMissingQ8vC:     !computed,
		BindingFailure:           "Q_8vC is not derived, so the neutral complex kernel and its dimension cannot be computed",
		Verdict:                  "complexification removes the real even-kernel obstruction in principle, but it does not construct the neutral three-plane without the actual Hermitian charge operator",
	}
}

func auditComplexTriality(k ComplexNeutralKernelAudit) ComplexTrialityAudit {
	return ComplexTrialityAudit{
		Spin8TrialityOverC:                true,
		Modules:                           []string{"8_v⊗C", "8_s⊗C", "8_c⊗C"},
		SameComplexDimension:              true,
		OuterAutomorphismRequired:         true,
		CanonicalUntwistedIsomorphism:     false,
		NeutralKernelAvailable:            k.Computed && k.ExactlyThreeComplexDim,
		MapNeutralKernelToSpinor:          false,
		RealStructureCompatibilityChecked: false,
		CompatibleWithJ:                   false,
		Obstruction:                       "complex Spin(8) triality relates the three 8-dimensional modules by outer automorphism, but the project has not derived a canonical vector-to-spinor isomorphism or shown that it commutes with the real structure J",
		Verdict:                           "complex triality is the right arena, but not yet a canonical scalar-to-spinor transport functor",
	}
}

func auditVTau(k ComplexNeutralKernelAudit, t ComplexTrialityAudit) VTauAudit {
	return VTauAudit{
		TauEta:                 []int{2, -2, 1},
		NeedsNeutral3Plane:     true,
		Neutral3PlaneAvailable: k.Computed && k.ExactlyThreeComplexDim,
		NeedsScalarSlotFrame:   true,
		ScalarSlotFrameDerived: false,
		Constructed:            false,
		WouldFeedTriality:      t.MapNeutralKernelToSpinor,
		RejectedBecause:        "no neutral 3-plane, no canonical scalar-slot frame in that plane, and no real-structure-compatible complex triality map are derived",
		Verdict:                "v_tau remains unconstructed; tau_eta is still a scalar trace ledger with texture capacity but no vector representative",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{Verdict: "firewall preserved: no Hermitian charge matrices, complex weights, neutral three-plane, triality isomorphism, v_tau, or Yukawa texture were inserted by hand"}
}

func summarize(c ComplexificationAudit, h HermitianPreflightAudit, cartan CartanWeightAudit, k ComplexNeutralKernelAudit, t ComplexTrialityAudit, vtau VTauAudit) Summary {
	status := strings.Join([]string{
		Status8VComplexification,
		StatusHermitianWeightCapacity,
		StatusOddComplexKernelCapacity,
		StatusNativeHermitianUnavailable,
		StatusNeutral3PlaneBlocked,
		StatusComplexTrialityPreflight,
		StatusCanonicalTrialityBlocked,
		StatusRealStructureBlocked,
		StatusVTauStillBlocked,
		StatusYukawaStillBlocked,
	}, "\n")
	return Summary{
		Complex8VKnown:              c.ComplexificationNative && c.ComplexDimension == 8,
		HermitianWeightCapacity:     h.RealSkewGeneratorAvailable && h.HermitianOperatorsHaveRealSpectrum,
		OddComplexKernelCapacity:    h.OddWeightSpacesAllowed && k.OddDimAllowedInPrinciple,
		NativeHermitianMatrices:     cartan.Q8vCMatrixDerived && cartan.Z8vCMatrixDerived,
		ComplexNeutralKernelDerived: k.Computed,
		NeutralKernelDim3:           k.ExactlyThreeComplexDim,
		ComplexTrialityArena:        t.Spin8TrialityOverC && t.SameComplexDimension,
		CanonicalTrialityMap:        t.CanonicalUntwistedIsomorphism && t.MapNeutralKernelToSpinor,
		RealStructureCompatible:     t.RealStructureCompatibilityChecked && t.CompatibleWithJ,
		VTauConstructed:             vtau.Constructed,
		TrialityUnblocked:           vtau.Constructed && t.MapNeutralKernelToSpinor && t.CompatibleWithJ,
		YukawaTextureDerived:        false,
		Status:                      status,
		NextGate:                    "Gate 252 — native Hermitian electroweak action on 8_vC or representation-functor derivation audit",
		Comment:                     "Gate 251 resolves the real even-rank obstruction only in principle. The actual Q_8vC/Z_8vC matrices, neutral three-plane, real-compatible triality map, and v_tau remain un-derived.",
	}
}

func buildTruth(prev InheritedGate250Audit, c ComplexificationAudit, h HermitianPreflightAudit, cartan CartanWeightAudit, k ComplexNeutralKernelAudit, t ComplexTrialityAudit, vtau VTauAudit) string {
	return fmt.Sprintf("Gate 251 complexifies the known %s carrier to %s and confirms that Hermitian weight-space decompositions can have odd-dimensional eigenspaces, so the Gate 250 real-skew even-kernel obstruction is not a fundamental obstruction to a complex neutral three-plane. However, the physical Hermitian matrices Q_8vC and Z_8vC are still unavailable (%s). Complex Spin(8) triality provides the correct representation arena, but only via a noncanonical outer-automorphism choice whose real-structure compatibility is not derived. Therefore ker(Q_8vC), v_tau, triality transport, and Yukawa texture remain blocked.", c.RealCarrierName, c.ComplexCarrierName, cartan.Obstruction)
}
