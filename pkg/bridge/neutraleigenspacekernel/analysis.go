// Package neutraleigenspacekernel implements Gate 249:
// Neutral Eigenspace Kernel / Invariant 3-Plane Isomorphism Audit.
//
// Gate 248 proved that the Spin(8) vector carrier 8_v is native, but that the
// neutral scalar trace triple tau_eta=(2,-2,1) cannot be assigned to arbitrary
// Gamma-basis directions without a coordinate-free scalar-to-vector map.
// Gate 249 tests a more lawful candidate for that map: derive an invariant
// neutral kernel inside 8_v from the native electroweak derivations Q and Z.
// The result is intentionally type-strict. The neutral-kernel strategy is the
// right kind of coordinate-free map, but the project does not yet derive an
// explicit Q/Z representation on 8_v, so the kernel, 3-plane, v_tau, and
// triality pullback all remain blocked.
package neutraleigenspacekernel

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/vectorrepresentative8v"
)

const (
	AuditID = "GATE249-NEUTRAL-EIGENSPACE-KERNEL-INVARIANT-3PLANE-ISOMORPHISM-AUDIT"

	Status8VCarrierInherited        = "CONDITIONAL_SUPPORT_8V_CARRIER_INHERITED"
	StatusNeutralKernelStrategy     = "CONDITIONAL_SUPPORT_NEUTRAL_KERNEL_STRATEGY_PREFLIGHT"
	StatusScalarTraceSlotsInherited = "CONDITIONAL_SUPPORT_NEUTRAL_SCALAR_TRACE_SLOTS_INHERITED"
	StatusDerivationActionBlocked   = "FAILED_ROUTE_EW_DERIVATION_ACTION_ON_8V"
	StatusNeutralKernelBlocked      = "FAILED_ROUTE_NEUTRAL_KERNEL_3PLANE_DERIVATION"
	StatusScalarNeutralPlaneBlocked = "FAILED_ROUTE_SCALAR_TO_NEUTRAL_3PLANE_ISOMORPHISM"
	StatusVTauBlocked               = "FAILED_ROUTE_NEUTRAL_KERNEL_V_TAU_CONSTRUCTION"
	StatusTrialityStillBlocked      = "FAILED_ROUTE_TRIALITY_PREFLIGHT_WITHOUT_NEUTRAL_8V_VECTOR"
	StatusYukawaTextureStillBlocked = "FAILED_ROUTE_YUKAWA_TEXTURE_DERIVATION"
)

type InheritedGate248Audit struct {
	Basis8VKnown            bool
	ScalarTraceOriginKnown  bool
	DimensionallyEmbeddable bool
	ScalarTo8VMapDerived    bool
	VTauConstructed         bool
	TrialityUnblocked       bool
	YukawaTextureDerived    bool
	CKMPMNSDerived          bool
	FermionMassesDerived    bool
	TruthStatement          string
}

type VectorCarrierAudit struct {
	RepresentationName       string
	Dimension                int
	NativeCarrierKnown       bool
	BasisLabels              []string
	RealOctonionicSplitKnown bool
	Verdict                  string
}

type EWDerivationActionAudit struct {
	Operators                    []string
	Source                       string
	QKnownAsScalarObservable     bool
	ZKnownAsScalarObservable     bool
	QMatrixOn8VDerived           bool
	ZMatrixOn8VDerived           bool
	SimultaneouslyDiagonal       bool
	ChargeSpectrumKnown          bool
	ManualRepresentation         string
	ManualRepresentationRejected bool
	Obstruction                  string
	Verdict                      string
}

type NeutralKernelAudit struct {
	KernelDefinition         string
	Computed                 bool
	DimensionKnown           bool
	Dimension                int
	ExactlyThreeDimensional  bool
	BasisIndependent         bool
	InvariantSubspaceDerived bool
	DependsOnMissingQMatrix  bool
	Verdict                  string
}

type ScalarPlaneIsomorphismAudit struct {
	ScalarTraceSlots          []string
	TauEta                    []int
	ScalarSlotCount           int
	NeutralKernelDimension    int
	DimensionMatch            bool
	CanonicalIsomorphism      bool
	BasisIndependentPairing   bool
	QZTYToNeutralBasisDerived bool
	Obstruction               string
	Verdict                   string
}

type VTauNeutralVectorAudit struct {
	Candidate            string
	Constructed          bool
	LawfulRepresentative bool
	Coefficients         []int
	HostSubspace         string
	WouldHaveNormSquared int
	WouldFeedTriality    bool
	RejectedBecause      string
	Verdict              string
}

type TrialityPreflightAudit struct {
	Requires8VVector            bool
	Neutral8VVectorAvailable    bool
	TrialityCanRun              bool
	DiagonalTextureConstructed  bool
	GenerationBreakingCapacity  bool
	NonCommutingTextureCapacity bool
	CKMPMNSDerived              bool
	FermionMassesDerived        bool
	Verdict                     string
}

type FirewallAudit struct {
	InventedQActionOn8V        bool
	ForcedNeutralKernelDim3    bool
	AssignedScalarSlotsToBasis bool
	ConstructedVTauByHand      bool
	InventedTrialityMatrices   bool
	InsertedYukawaTexture      bool
	ImportedObservedMasses     bool
	ImportedCKMPMNS            bool
	ClaimedFiniteFlavorTheorem bool
	PollutedFiniteCore         bool
	Verdict                    string
}

type Summary struct {
	Basis8VKnown              bool
	EWDerivationActionDerived bool
	NeutralKernelDerived      bool
	NeutralKernelDim3         bool
	ScalarNeutralIsomorphism  bool
	VTauConstructed           bool
	TrialityUnblocked         bool
	YukawaTextureDerived      bool
	CKMPMNSDerived            bool
	FermionMassesDerived      bool
	Status                    string
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	PreviousGate248 InheritedGate248Audit
	VectorCarrier   VectorCarrierAudit
	EWAction        EWDerivationActionAudit
	NeutralKernel   NeutralKernelAudit
	ScalarPlane     ScalarPlaneIsomorphismAudit
	VTau            VTauNeutralVectorAudit
	Triality        TrialityPreflightAudit
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
		prevRaw, err := vectorrepresentative8v.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate248(prevRaw)
		carrier := auditVectorCarrier(prevRaw)
		ew := auditEWDerivationAction()
		kernel := auditNeutralKernel(ew)
		scalar := auditScalarPlane(kernel)
		vtau := auditVTau(scalar, kernel)
		triality := auditTriality(vtau)
		firewall := auditFirewall()
		summary := summarize(prev, carrier, ew, kernel, scalar, vtau, triality)
		truth := buildTruth(prev, carrier, ew, kernel, scalar, vtau, triality)
		defaultA = Analysis{PreviousGate248: prev, VectorCarrier: carrier, EWAction: ew, NeutralKernel: kernel, ScalarPlane: scalar, VTau: vtau, Triality: triality, Firewall: firewall, Summary: summary, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate248(a vectorrepresentative8v.Analysis) InheritedGate248Audit {
	return InheritedGate248Audit{
		Basis8VKnown:            a.Summary.Basis8VKnown,
		ScalarTraceOriginKnown:  a.Summary.ScalarTraceOriginKnown,
		DimensionallyEmbeddable: a.Summary.DimensionallyEmbeddable,
		ScalarTo8VMapDerived:    a.Summary.ScalarTo8VMapDerived,
		VTauConstructed:         a.Summary.VTauConstructed,
		TrialityUnblocked:       a.Summary.TrialityUnblocked,
		YukawaTextureDerived:    a.Summary.YukawaTextureDerived,
		CKMPMNSDerived:          a.Summary.CKMPMNSDerived,
		FermionMassesDerived:    a.Summary.FermionMassesDerived,
		TruthStatement:          a.TruthStatement,
	}
}

func auditVectorCarrier(prevRaw vectorrepresentative8v.Analysis) VectorCarrierAudit {
	return VectorCarrierAudit{
		RepresentationName:       "Spin(8) vector representation 8_v",
		Dimension:                prevRaw.VectorBasis.Dimension,
		NativeCarrierKnown:       prevRaw.VectorBasis.NativeCarrierKnown,
		BasisLabels:              append([]string(nil), prevRaw.VectorBasis.BasisLabels...),
		RealOctonionicSplitKnown: prevRaw.VectorBasis.RealOctonionicSplitKnown,
		Verdict:                  "the 8_v carrier is available as the only lawful host for a neutral-kernel vector representative, but no electroweak operator action on it is inherited from Gate 248",
	}
}

func auditEWDerivationAction() EWDerivationActionAudit {
	return EWDerivationActionAudit{
		Operators:                    []string{"Q = T3L + Y_phi", "Z = T3L - Y_phi", "T3L^T Y_phi"},
		Source:                       "neutral electroweak scalar trace ledger inherited from H_Phi",
		QKnownAsScalarObservable:     true,
		ZKnownAsScalarObservable:     true,
		QMatrixOn8VDerived:           false,
		ZMatrixOn8VDerived:           false,
		SimultaneouslyDiagonal:       false,
		ChargeSpectrumKnown:          false,
		ManualRepresentation:         "Q_8v ?= diag(q_0,...,q_7), Z_8v ?= diag(z_0,...,z_7)",
		ManualRepresentationRejected: true,
		Obstruction:                  "the project has neutral scalar observables Q and Z on H_Phi/fermion charge ledgers, but no derived representation matrices Q_8v and Z_8v acting on the Spin(8) vector carrier",
		Verdict:                      "the neutral-kernel strategy is well-typed only after Q and Z act on 8_v; that representation is not derived yet",
	}
}

func auditNeutralKernel(ew EWDerivationActionAudit) NeutralKernelAudit {
	computed := ew.QMatrixOn8VDerived
	dimensionKnown := computed
	dim := -1
	exact3 := false
	if computed {
		// No current path reaches this branch; it is left explicit so future
		// gates have a single place to add the true kernel computation.
		dim = 0
		exact3 = dim == 3
	}
	return NeutralKernelAudit{
		KernelDefinition:         "ker(Q_8v) = {v in 8_v | Q_8v v = 0}",
		Computed:                 computed,
		DimensionKnown:           dimensionKnown,
		Dimension:                dim,
		ExactlyThreeDimensional:  exact3,
		BasisIndependent:         computed,
		InvariantSubspaceDerived: exact3,
		DependsOnMissingQMatrix:  !computed,
		Verdict:                  "the neutral eigenspace would be a coordinate-free 3-plane if Q_8v were derived and dim ker(Q_8v)=3; currently the kernel is not computable",
	}
}

func auditScalarPlane(kernel NeutralKernelAudit) ScalarPlaneIsomorphismAudit {
	tau := []int{2, -2, 1}
	slots := []string{"tau_eta(Q^TQ)=2", "tau_eta(Z^TZ)=-2", "tau_eta(T3L^T Y_phi)=1"}
	dimMatch := kernel.DimensionKnown && kernel.Dimension == len(tau)
	return ScalarPlaneIsomorphismAudit{
		ScalarTraceSlots:          slots,
		TauEta:                    tau,
		ScalarSlotCount:           len(tau),
		NeutralKernelDimension:    kernel.Dimension,
		DimensionMatch:            dimMatch,
		CanonicalIsomorphism:      dimMatch && kernel.BasisIndependent && kernel.InvariantSubspaceDerived,
		BasisIndependentPairing:   false,
		QZTYToNeutralBasisDerived: false,
		Obstruction:               "even a three-dimensional neutral kernel would still require a basis-independent pairing between the three scalar trace slots and a canonical basis/frame of ker(Q_8v); neither is derived",
		Verdict:                   "the scalar trace triple cannot yet be identified with a neutral 8_v 3-plane",
	}
}

func auditVTau(s ScalarPlaneIsomorphismAudit, k NeutralKernelAudit) VTauNeutralVectorAudit {
	lawful := s.CanonicalIsomorphism && s.BasisIndependentPairing && s.QZTYToNeutralBasisDerived
	return VTauNeutralVectorAudit{
		Candidate:            "v_tau ?= 2 n_1 - 2 n_2 + n_3, with {n_i} a canonical frame of ker(Q_8v)",
		Constructed:          lawful,
		LawfulRepresentative: lawful,
		Coefficients:         append([]int(nil), s.TauEta...),
		HostSubspace:         k.KernelDefinition,
		WouldHaveNormSquared: 9,
		WouldFeedTriality:    lawful,
		RejectedBecause:      "the neutral kernel and its scalar-slot frame are not derived; constructing v_tau would force a basis assignment",
		Verdict:              "v_tau remains blocked until Q_8v, dim ker(Q_8v)=3, and a scalar-slot frame are derived",
	}
}

func auditTriality(v VTauNeutralVectorAudit) TrialityPreflightAudit {
	return TrialityPreflightAudit{
		Requires8VVector:            true,
		Neutral8VVectorAvailable:    v.LawfulRepresentative,
		TrialityCanRun:              v.LawfulRepresentative,
		DiagonalTextureConstructed:  false,
		GenerationBreakingCapacity:  true,
		NonCommutingTextureCapacity: true,
		CKMPMNSDerived:              false,
		FermionMassesDerived:        false,
		Verdict:                     "Spin(8) triality remains blocked because the required neutral 8_v representative has not been constructed",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		InventedQActionOn8V:        false,
		ForcedNeutralKernelDim3:    false,
		AssignedScalarSlotsToBasis: false,
		ConstructedVTauByHand:      false,
		InventedTrialityMatrices:   false,
		InsertedYukawaTexture:      false,
		ImportedObservedMasses:     false,
		ImportedCKMPMNS:            false,
		ClaimedFiniteFlavorTheorem: false,
		PollutedFiniteCore:         false,
		Verdict:                    "Gate 249 preserves the coordinate-free firewall: no Q_8v representation, neutral kernel, v_tau, or Yukawa texture is invented",
	}
}

func summarize(prev InheritedGate248Audit, c VectorCarrierAudit, ew EWDerivationActionAudit, k NeutralKernelAudit, sp ScalarPlaneIsomorphismAudit, v VTauNeutralVectorAudit, tr TrialityPreflightAudit) Summary {
	status := strings.Join([]string{
		Status8VCarrierInherited,
		StatusNeutralKernelStrategy,
		StatusScalarTraceSlotsInherited,
		StatusDerivationActionBlocked,
		StatusNeutralKernelBlocked,
		StatusScalarNeutralPlaneBlocked,
		StatusVTauBlocked,
		StatusTrialityStillBlocked,
		StatusYukawaTextureStillBlocked,
	}, "\n")
	return Summary{
		Basis8VKnown:              prev.Basis8VKnown && c.NativeCarrierKnown && c.Dimension == 8,
		EWDerivationActionDerived: ew.QMatrixOn8VDerived && ew.ZMatrixOn8VDerived,
		NeutralKernelDerived:      k.Computed && k.InvariantSubspaceDerived,
		NeutralKernelDim3:         k.ExactlyThreeDimensional,
		ScalarNeutralIsomorphism:  sp.CanonicalIsomorphism,
		VTauConstructed:           v.LawfulRepresentative,
		TrialityUnblocked:         tr.TrialityCanRun,
		YukawaTextureDerived:      false,
		CKMPMNSDerived:            false,
		FermionMassesDerived:      false,
		Status:                    status,
		NextGate:                  "derive explicit electroweak derivation matrices on 8_v, or prove no such action exists in the current finite geometry",
		Comment:                   "The neutral-kernel route is coordinate-free and physically well-motivated, but it cannot run until Q and Z act on 8_v. Gate 249 therefore blocks v_tau construction without falsifying the broader tau_eta texture capacity.",
	}
}

func buildTruth(prev InheritedGate248Audit, c VectorCarrierAudit, ew EWDerivationActionAudit, k NeutralKernelAudit, sp ScalarPlaneIsomorphismAudit, v VTauNeutralVectorAudit, tr TrialityPreflightAudit) string {
	parts := []string{
		fmt.Sprintf("Gate248 basis8v=%t scalarOrigin=%t map=%t", prev.Basis8VKnown, prev.ScalarTraceOriginKnown, prev.ScalarTo8VMapDerived),
		fmt.Sprintf("8v dim=%d native=%t", c.Dimension, c.NativeCarrierKnown),
		fmt.Sprintf("Q8v=%t Z8v=%t", ew.QMatrixOn8VDerived, ew.ZMatrixOn8VDerived),
		fmt.Sprintf("neutralKernelComputed=%t dimKnown=%t dim=%d exact3=%t", k.Computed, k.DimensionKnown, k.Dimension, k.ExactlyThreeDimensional),
		fmt.Sprintf("scalarPlaneIso=%t vtau=%t triality=%t", sp.CanonicalIsomorphism, v.LawfulRepresentative, tr.TrialityCanRun),
	}
	return strings.Join(parts, " | ")
}
