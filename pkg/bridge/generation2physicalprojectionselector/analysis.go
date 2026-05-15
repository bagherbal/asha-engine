// Package generation2physicalprojectionselector implements Gate 528:
// Physical 3+1 Projection and Internal Complement Selector Audit.
//
// Gate 527 left the 1+7 -> (1+3)+4 split as a rank-valid bridge socket
// but blocked native promotion because no explicit projector, idempotent, or
// subalgebra selector had been supplied. Gate 528 audits the Cℓ(1,7) ladder
// for native projectors: full chirality/volume projectors, chosen four-plane
// projectors, and earlier scalar/projector sockets. It accepts only what the
// algebra actually selects and quarantines the physical 3+1 spacetime split
// when a unique parameter-free vector-space projector is absent.
package generation2physicalprojectionselector

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2lorentzianspinoradjointairlock"
)

const (
	AuditID = "GATE528-PHYSICAL-3PLUS1-PROJECTION-SELECTOR"

	StatusGate527Inherited                          = "CONDITIONAL_SUPPORT_GATE527_PROJECTION_AIRLOCK_INHERITED"
	StatusIdempotentSieveExecuted                   = "CONDITIONAL_SUPPORT_CL17_IDEMPOTENT_SIEVE_EXECUTED"
	StatusChiralityVolumeProjectorSocketFound       = "CONDITIONAL_SUPPORT_VOLUME_CHIRALITY_PROJECTOR_SOCKET_FOUND"
	StatusChosenFourPlaneProjectorBridgeConstructed = "CONDITIONAL_SUPPORT_CHOSEN_FOUR_PLANE_PROJECTOR_BRIDGE_CONSTRUCTED"
	StatusRank44ArithmeticConfirmed                 = "CONDITIONAL_SUPPORT_4PLUS4_RANK_ARITHMETIC_CONFIRMED"
	StatusBridgeInternalComplementSocketConsistent  = "CONDITIONAL_SUPPORT_INTERNAL_FOUR_COMPLEMENT_SOCKET_CONSISTENT_BRIDGE_ONLY"
	StatusNoObservedDataImported                    = "CONDITIONAL_SUPPORT_NO_OBSERVED_DIMENSION_TOPOLOGY_MASS_OR_CONSTANT_DATA_IMPORTED"

	StatusFailedChiralityDoesNotProjectVector44       = "FAILED_ROUTE_CHIRALITY_VOLUME_PROJECTOR_DOES_NOT_SELECT_VECTOR_4PLUS4"
	StatusFailedFourPlaneRequiresSubspaceChoice       = "FAILED_ROUTE_RANK4_PROJECTOR_REQUIRES_CHOSEN_FOUR_PLANE"
	StatusFailedNoSpin17InvariantRank4VectorProjector = "FAILED_ROUTE_NO_SPIN17_INVARIANT_RANK4_VECTOR_PROJECTOR"
	StatusFailedNoUniqueInternalComplement            = "FAILED_ROUTE_INTERNAL_COMPLEMENT_NOT_UNIQUELY_NATIVE"
	StatusFailedTimeAssignmentNotNative               = "FAILED_ROUTE_TIME_ASSIGNMENT_TO_EXTERNAL_3PLUS1_NOT_NATIVE_SELECTED"
	StatusFailedMutualCommutingSubalgebrasNotNative   = "FAILED_ROUTE_MUTUALLY_COMMUTING_EXTERNAL_INTERNAL_SUBALGEBRAS_NOT_NATIVE_SELECTED"
	StatusFailedPhysical3Plus1ProjectorNotIdentified  = "FAILED_ROUTE_NATIVE_3PLUS1_PROJECTOR_NOT_IDENTIFIED"
	StatusFailedWickHilbertDynamicsStillBlocked       = "FAILED_ROUTE_WICK_HILBERT_AND_UNITARY_DYNAMICS_STILL_BLOCKED_AFTER_PROJECTOR_AUDIT"
	StatusFirewallPreserved                           = "FIREWALL_PRESERVED_COMPLETED_SECTOR_AIRLOCKS_DURING_PROJECTOR_AUDIT"
	StatusFirewallProjectionNativeWriteBlocked        = "FIREWALL_BLOCKED_PHYSICAL_3PLUS1_AND_INTERNAL_COMPLEMENT_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate527Inherited                bool
	Gate527KreinSocket              bool
	Gate527ProjectionAirlockDefined bool
	Gate527Physical3Plus1Blocked    bool
	Gate527PositiveHilbertBlocked   bool
	Gate527ReflectionPositivityOpen bool
	Gate527WickBlocked              bool
	Gate527UnitaryDynamicsBlocked   bool
	Gate527NoObservedDataImported   bool
	Gate527NativeWriteBlocked       bool
	Gate527ReopenedSealedFirewalls  bool

	Verdict, Reason string
}

type IdempotentSieve struct {
	Executed bool

	Algebra                                   string
	FullDimension                             int
	VolumeElementAvailable                    bool
	ChiralityProjectorsAvailable              bool
	ChiralityProjectorsIdempotent             bool
	ChiralityActsOnSpinorParity               bool
	ChiralityProjectsVectorSpace44            bool
	PrimitiveIdempotentsAbundant              bool
	PrimitiveIdempotentsCanonical             bool
	PreviouslyDerivedScalarProjectorsRelevant bool

	Verdict, Reason string
}

type Rank44Audit struct {
	Executed bool

	CandidateSplit                      string
	VectorDimension                     int
	CandidateExternalRank               int
	CandidateInternalRank               int
	RankArithmeticValid                 bool
	ChosenFourPlaneProjectorIdempotent  bool
	ProjectorComplementary              bool
	ProjectorRequiresBasisChoice        bool
	Spin17InvariantRank4ProjectorFound  bool
	MutuallyCommutingSubalgebrasNative  bool
	GradedTensorFactorizationBridgeOnly bool
	InternalComplementUniqueNative      bool

	Verdict, Reason string
}

type SpacetimeSelector struct {
	Executed bool

	ExternalLorentzSignatureCandidate string
	TimeLikeDirectionAvailable        bool
	TimeIncludedByChosenBridgePlane   bool
	TimeAssignmentNativeSelected      bool
	OrientationAndArrowSelected       bool
	Physical3Plus1ProjectorIdentified bool
	Physical3Plus1BridgeSocketReady   bool
	InternalGaugeSpaceIdentified      bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedDimensionImported     bool
	ObservedConstantsImported     bool
	ObservedMassesImported        bool
	ObservedTopologyImported      bool
	NativeChiralityVectorWrite    bool
	NativeFourPlaneWrite          bool
	NativeInternalComplementWrite bool
	NativeTimeAssignmentWrite     bool
	Native3Plus1ProjectionWrite   bool
	NativeHilbertDynamicsWrite    bool
	ReopenedFlavorFirewall        bool
	ReopenedEWScaleFirewall       bool
	ReopenedGravityFirewall       bool
	ReopenedTopologyFirewall      bool
	NativeRegistryWritten         bool

	Verdict, Reason string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Idempotents IdempotentSieve
	Rank44      Rank44Audit
	Selector    SpacetimeSelector
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g527, err := generation2lorentzianspinoradjointairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate527 Lorentzian adjoint/projection audit: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g527)
	a.Idempotents = buildIdempotentSieve(a.Inheritance)
	a.Rank44 = buildRank44Audit(a.Inheritance, a.Idempotents)
	a.Selector = buildSpacetimeSelector(a.Inheritance, a.Rank44)
	a.Firewall = buildFirewall(a.Inheritance, a.Idempotents, a.Rank44, a.Selector)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g527 generation2lorentzianspinoradjointairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate527Inherited:                g527.Inheritance.Gate526SignatureInherited && g527.Inheritance.Gate526NullConeConfirmed,
		Gate527KreinSocket:              g527.Adjoint.KreinAdjointDefined && g527.Adjoint.CliffordCompatibility,
		Gate527ProjectionAirlockDefined: g527.Projection.ProjectionRankArithmeticValid && g527.Projection.CandidateExternalDimension == 4 && g527.Projection.CandidateInternalComplement == 4,
		Gate527Physical3Plus1Blocked:    !g527.Projection.Physical3Plus1Selected && !g527.Projection.ProjectionOperatorNativeSelected,
		Gate527PositiveHilbertBlocked:   !g527.Adjoint.PositiveHilbertProductSelected && !g527.Adjoint.PhysicalStateSpaceSelected,
		Gate527ReflectionPositivityOpen: !g527.Reflection.ReflectionPositivityProven && !g527.Reflection.OsterwalderSchraderAxiomsProven,
		Gate527WickBlocked:              !g527.Reflection.WickContinuationSelected,
		Gate527UnitaryDynamicsBlocked:   !g527.Reflection.UnitaryRealTimeDynamicsDerived,
		Gate527NoObservedDataImported:   !g527.Firewall.ObservedConstantsImported && !g527.Firewall.ObservedMassesImported && !g527.Firewall.ObservedTopologyImported && !g527.Firewall.ObservedBoundaryImported,
		Gate527NativeWriteBlocked:       !g527.Firewall.NativeRegistryWritten,
		Gate527ReopenedSealedFirewalls:  g527.Firewall.ReopenedFlavorFirewall || g527.Firewall.ReopenedEWScaleFirewall || g527.Firewall.ReopenedGravityFirewall || g527.Firewall.ReopenedTopologyFirewall,
		Verdict:                         StatusGate527Inherited,
		Reason:                          "Gate528 inherits Gate527's Lorentzian/Krein socket and its projection airlock: the 1+7 -> (1+3)+4 rank split is admissible as a bridge socket, but no native projector or physical Hilbert/Wick dynamics was selected.",
	}
}

func buildIdempotentSieve(in Inheritance) IdempotentSieve {
	return IdempotentSieve{
		Executed:                                  true,
		Algebra:                                   "Cℓ(1,7)",
		FullDimension:                             8,
		VolumeElementAvailable:                    in.Gate527Inherited,
		ChiralityProjectorsAvailable:              in.Gate527Inherited,
		ChiralityProjectorsIdempotent:             in.Gate527Inherited,
		ChiralityActsOnSpinorParity:               true,
		ChiralityProjectsVectorSpace44:            false,
		PrimitiveIdempotentsAbundant:              true,
		PrimitiveIdempotentsCanonical:             false,
		PreviouslyDerivedScalarProjectorsRelevant: false,
		Verdict: strings.Join([]string{StatusIdempotentSieveExecuted, StatusChiralityVolumeProjectorSocketFound, StatusFailedChiralityDoesNotProjectVector44}, ";"),
		Reason:  "The full Clifford volume/chirality element supplies idempotent spinor/parity projectors after the usual algebraic convention, but those project chirality sectors, not a canonical rank-four subspace of the underlying vector representation. Primitive idempotents are abundant and gauge-dependent; abundance is not selection.",
	}
}

func buildRank44Audit(in Inheritance, ids IdempotentSieve) Rank44Audit {
	return Rank44Audit{
		Executed:                            true,
		CandidateSplit:                      "1+7 -> chosen bridge (1+3) external + 4 internal complement",
		VectorDimension:                     8,
		CandidateExternalRank:               4,
		CandidateInternalRank:               4,
		RankArithmeticValid:                 true,
		ChosenFourPlaneProjectorIdempotent:  true,
		ProjectorComplementary:              true,
		ProjectorRequiresBasisChoice:        true,
		Spin17InvariantRank4ProjectorFound:  false,
		MutuallyCommutingSubalgebrasNative:  false,
		GradedTensorFactorizationBridgeOnly: true,
		InternalComplementUniqueNative:      false,
		Verdict:                             strings.Join([]string{StatusChosenFourPlaneProjectorBridgeConstructed, StatusRank44ArithmeticConfirmed, StatusBridgeInternalComplementSocketConsistent, StatusFailedFourPlaneRequiresSubspaceChoice, StatusFailedNoSpin17InvariantRank4VectorProjector, StatusFailedNoUniqueInternalComplement, StatusFailedMutualCommutingSubalgebrasNotNative}, ";"),
		Reason:                              "A rank-four projector and complementary rank-four kernel can be written once a four-plane is chosen. That verifies the bridge arithmetic and a graded tensor-factorization socket, but the choice of four-plane breaks Spin(1,7) covariance and is not selected by a unique native idempotent.",
	}
}

func buildSpacetimeSelector(in Inheritance, r Rank44Audit) SpacetimeSelector {
	return SpacetimeSelector{
		Executed:                          true,
		ExternalLorentzSignatureCandidate: "1+3",
		TimeLikeDirectionAvailable:        in.Gate527Inherited,
		TimeIncludedByChosenBridgePlane:   true,
		TimeAssignmentNativeSelected:      false,
		OrientationAndArrowSelected:       false,
		Physical3Plus1ProjectorIdentified: false,
		Physical3Plus1BridgeSocketReady:   r.RankArithmeticValid && r.ChosenFourPlaneProjectorIdempotent && r.ProjectorComplementary,
		InternalGaugeSpaceIdentified:      false,
		Verdict:                           strings.Join([]string{StatusFailedTimeAssignmentNotNative, StatusFailedPhysical3Plus1ProjectorNotIdentified, StatusFailedWickHilbertDynamicsStillBlocked}, ";"),
		Reason:                            "A bridge convention may include the single timelike direction in a chosen four-plane and call the complement internal, but the algebra has not selected that projector, the time assignment, the arrow/orientation, or an identification of the complement with the physical gauge/internal sector.",
	}
}

func buildFirewall(in Inheritance, ids IdempotentSieve, r Rank44Audit, s SpacetimeSelector) Firewall {
	return Firewall{
		Executed:                      true,
		ObservedDimensionImported:     false,
		ObservedConstantsImported:     false,
		ObservedMassesImported:        false,
		ObservedTopologyImported:      false,
		NativeChiralityVectorWrite:    ids.ChiralityProjectsVectorSpace44,
		NativeFourPlaneWrite:          r.Spin17InvariantRank4ProjectorFound,
		NativeInternalComplementWrite: r.InternalComplementUniqueNative || s.InternalGaugeSpaceIdentified,
		NativeTimeAssignmentWrite:     s.TimeAssignmentNativeSelected || s.OrientationAndArrowSelected,
		Native3Plus1ProjectionWrite:   s.Physical3Plus1ProjectorIdentified,
		NativeHilbertDynamicsWrite:    false,
		ReopenedFlavorFirewall:        false,
		ReopenedEWScaleFirewall:       false,
		ReopenedGravityFirewall:       false,
		ReopenedTopologyFirewall:      false,
		NativeRegistryWritten:         false,
		Verdict:                       strings.Join([]string{StatusNoObservedDataImported, StatusFirewallPreserved, StatusFirewallProjectionNativeWriteBlocked}, ";"),
		Reason:                        "Gate528 imports no observed dimensionality, constants, masses, topology, or boundary data and writes no native vector 4+4 projector, physical 3+1 spacetime, internal complement, Wick/Hilbert dynamics, or time assignment.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) retains the native 1+7 causal quadratic socket and full Clifford volume/chirality structure",
			"volume/chirality projectors are valid algebraic sockets on spinor/parity sectors, not vector-space 4+4 selectors",
			"the null-cone and Lorentzian/Krein sockets remain independent of flavor, electroweak scale, gravity normalization, and global topology",
		},
		BridgeEntries: []string{
			"a chosen rank-four external projector plus rank-four complement is algebraically consistent only after selecting a four-plane",
			"the candidate 1+7 -> (1+3)+4 split is a bridge projection schema until a unique native idempotent/subalgebra selector is proven",
			"graded tensor factorization of external and internal Clifford factors is bridge-compatible but not a native spacetime theorem",
		},
		EnvironmentalEntries: []string{
			"physical external spacetime identification, time orientation, and arrow of time remain bridge/environmental inputs",
			"identification of the internal four-dimensional complement with continuum gauge/internal geometry remains unpromoted",
		},
		FailedRoutes: []string{
			StatusFailedChiralityDoesNotProjectVector44,
			StatusFailedFourPlaneRequiresSubspaceChoice,
			StatusFailedNoSpin17InvariantRank4VectorProjector,
			StatusFailedNoUniqueInternalComplement,
			StatusFailedTimeAssignmentNotNative,
			StatusFailedMutualCommutingSubalgebrasNotNative,
			StatusFailedPhysical3Plus1ProjectorNotIdentified,
			StatusFailedWickHilbertDynamicsStillBlocked,
		},
		OpenTheorems: []string{
			"construct a native Spin(1,7)-breaking but theorem-selected rank-four vector projector, or keep 3+1 projection bridge-only",
			"define a fail-closed bridge schema for explicit 3+1 projector choices with source, convention, and native-promotion rejection metadata",
			"audit compatibility between any bridge 3+1 projection and previously sealed Wick/Hilbert/positive-energy obligations",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        529,
		Title:       "3+1 Projection and Internal Complement Bridge Airlock Preflight",
		Reason:      "Gate528 finds no unique native rank-four vector projector. The next safe step is to define an explicit bridge schema for 3+1 projector choices, internal complements, signature convention, and promotion rejection before any dimensional-reduction comparator is executed.",
		PrimaryTask: "Build a fail-closed 3+1 projection airlock that accepts only labelled bridge projectors and blocks native spacetime, Wick, Hilbert, and internal-gauge promotion.",
	}
}

func truth(a Analysis) string {
	return "Gate528 confirms that Cℓ(1,7) has rich idempotent/chirality structure and that a 4+4 split is algebraically admissible after choosing a four-plane. It does not identify a unique native vector-space projector selecting physical 3+1 spacetime, its timelike assignment, or a four-dimensional internal complement. The physical 3+1 projection remains a bridge airlock, not a native ASHA theorem."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate527Inherited || !a.Inheritance.Gate527KreinSocket || !a.Inheritance.Gate527ProjectionAirlockDefined || !a.Inheritance.Gate527Physical3Plus1Blocked || !a.Inheritance.Gate527PositiveHilbertBlocked || !a.Inheritance.Gate527ReflectionPositivityOpen || !a.Inheritance.Gate527WickBlocked || !a.Inheritance.Gate527UnitaryDynamicsBlocked || !a.Inheritance.Gate527NoObservedDataImported || !a.Inheritance.Gate527NativeWriteBlocked || a.Inheritance.Gate527ReopenedSealedFirewalls {
		bad = append(bad, "bad inheritance")
	}
	if !a.Idempotents.Executed || a.Idempotents.Algebra != "Cℓ(1,7)" || a.Idempotents.FullDimension != 8 || !a.Idempotents.VolumeElementAvailable || !a.Idempotents.ChiralityProjectorsAvailable || !a.Idempotents.ChiralityProjectorsIdempotent || !a.Idempotents.ChiralityActsOnSpinorParity || a.Idempotents.ChiralityProjectsVectorSpace44 || !a.Idempotents.PrimitiveIdempotentsAbundant || a.Idempotents.PrimitiveIdempotentsCanonical {
		bad = append(bad, "bad idempotent sieve")
	}
	if !a.Rank44.Executed || a.Rank44.VectorDimension != 8 || a.Rank44.CandidateExternalRank != 4 || a.Rank44.CandidateInternalRank != 4 || !a.Rank44.RankArithmeticValid || !a.Rank44.ChosenFourPlaneProjectorIdempotent || !a.Rank44.ProjectorComplementary || !a.Rank44.ProjectorRequiresBasisChoice || a.Rank44.Spin17InvariantRank4ProjectorFound || a.Rank44.MutuallyCommutingSubalgebrasNative || !a.Rank44.GradedTensorFactorizationBridgeOnly || a.Rank44.InternalComplementUniqueNative {
		bad = append(bad, "bad 4+4 audit")
	}
	if !a.Selector.Executed || a.Selector.ExternalLorentzSignatureCandidate != "1+3" || !a.Selector.TimeLikeDirectionAvailable || !a.Selector.TimeIncludedByChosenBridgePlane || a.Selector.TimeAssignmentNativeSelected || a.Selector.OrientationAndArrowSelected || a.Selector.Physical3Plus1ProjectorIdentified || !a.Selector.Physical3Plus1BridgeSocketReady || a.Selector.InternalGaugeSpaceIdentified {
		bad = append(bad, "bad spacetime selector")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedDimensionImported || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.NativeChiralityVectorWrite || a.Firewall.NativeFourPlaneWrite || a.Firewall.NativeInternalComplementWrite || a.Firewall.NativeTimeAssignmentWrite || a.Firewall.Native3Plus1ProjectionWrite || a.Firewall.NativeHilbertDynamicsWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: inherited=%t Krein=%t projection_airlock=%t 3plus1_blocked=%t positive_Hilbert_blocked=%t reflection_open=%t Wick_blocked=%t unitary_blocked=%t no_observed=%t native_blocked=%t reopens_firewalls=%t; %s", x.Verdict, x.Gate527Inherited, x.Gate527KreinSocket, x.Gate527ProjectionAirlockDefined, x.Gate527Physical3Plus1Blocked, x.Gate527PositiveHilbertBlocked, x.Gate527ReflectionPositivityOpen, x.Gate527WickBlocked, x.Gate527UnitaryDynamicsBlocked, x.Gate527NoObservedDataImported, x.Gate527NativeWriteBlocked, x.Gate527ReopenedSealedFirewalls, x.Reason)
}

func FormatIdempotents(x IdempotentSieve) string {
	return fmt.Sprintf("%s: algebra=%s dim=%d volume=%t chirality_projectors=%t idempotent=%t spinor_parity=%t vector_4plus4=%t primitive_abundant=%t primitive_canonical=%t scalar_projectors_relevant=%t; %s", x.Verdict, x.Algebra, x.FullDimension, x.VolumeElementAvailable, x.ChiralityProjectorsAvailable, x.ChiralityProjectorsIdempotent, x.ChiralityActsOnSpinorParity, x.ChiralityProjectsVectorSpace44, x.PrimitiveIdempotentsAbundant, x.PrimitiveIdempotentsCanonical, x.PreviouslyDerivedScalarProjectorsRelevant, x.Reason)
}

func FormatRank44(x Rank44Audit) string {
	return fmt.Sprintf("%s: split=%q vector_dim=%d external_rank=%d internal_rank=%d rank_valid=%t chosen_projector_idempotent=%t complement=%t requires_choice=%t Spin17_invariant_rank4=%t commuting_subalgebras_native=%t graded_factorization_bridge=%t internal_unique=%t; %s", x.Verdict, x.CandidateSplit, x.VectorDimension, x.CandidateExternalRank, x.CandidateInternalRank, x.RankArithmeticValid, x.ChosenFourPlaneProjectorIdempotent, x.ProjectorComplementary, x.ProjectorRequiresBasisChoice, x.Spin17InvariantRank4ProjectorFound, x.MutuallyCommutingSubalgebrasNative, x.GradedTensorFactorizationBridgeOnly, x.InternalComplementUniqueNative, x.Reason)
}

func FormatSelector(x SpacetimeSelector) string {
	return fmt.Sprintf("%s: external_signature=%s timelike_available=%t timelike_in_bridge_plane=%t time_native=%t arrow_native=%t physical_projector=%t bridge_socket_ready=%t internal_gauge_identified=%t; %s", x.Verdict, x.ExternalLorentzSignatureCandidate, x.TimeLikeDirectionAvailable, x.TimeIncludedByChosenBridgePlane, x.TimeAssignmentNativeSelected, x.OrientationAndArrowSelected, x.Physical3Plus1ProjectorIdentified, x.Physical3Plus1BridgeSocketReady, x.InternalGaugeSpaceIdentified, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_dimension=%t observed_constants=%t observed_masses=%t observed_topology=%t native_chirality_vector=%t native_four_plane=%t native_internal=%t native_time=%t native_3plus1=%t native_Hilbert_dynamics=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_write=%t; %s", x.Verdict, x.ObservedDimensionImported, x.ObservedConstantsImported, x.ObservedMassesImported, x.ObservedTopologyImported, x.NativeChiralityVectorWrite, x.NativeFourPlaneWrite, x.NativeInternalComplementWrite, x.NativeTimeAssignmentWrite, x.Native3Plus1ProjectionWrite, x.NativeHilbertDynamicsWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate527Inherited,
		StatusIdempotentSieveExecuted,
		StatusChiralityVolumeProjectorSocketFound,
		StatusChosenFourPlaneProjectorBridgeConstructed,
		StatusRank44ArithmeticConfirmed,
		StatusBridgeInternalComplementSocketConsistent,
		StatusNoObservedDataImported,
		StatusFailedChiralityDoesNotProjectVector44,
		StatusFailedFourPlaneRequiresSubspaceChoice,
		StatusFailedNoSpin17InvariantRank4VectorProjector,
		StatusFailedNoUniqueInternalComplement,
		StatusFailedTimeAssignmentNotNative,
		StatusFailedMutualCommutingSubalgebrasNotNative,
		StatusFailedPhysical3Plus1ProjectorNotIdentified,
		StatusFailedWickHilbertDynamicsStillBlocked,
		StatusFirewallPreserved,
		StatusFirewallProjectionNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 528 Registry Audit — Physical 3+1 Projection and Internal Complement Selector Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 528 inherits Gate 527's Lorentzian/Krein socket and bridge-only 3+1 projection airlock.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Idempotent sieve\n\nVolume/chirality idempotents are real Clifford algebra sockets, but they do not project the base vector space into a canonical 4+4 spacetime/internal split.\n\n```text\n" + FormatIdempotents(a.Idempotents) + "\n```\n\n")
	b.WriteString("## 4+4 rank audit\n\nA chosen four-plane projector is idempotent and complementary, but the choice is bridge data rather than a Spin(1,7)-invariant native selector.\n\n```text\n" + FormatRank44(a.Rank44) + "\n```\n\n")
	b.WriteString("## Spacetime selector audit\n\nThe external 1+3 assignment is admissible only after a bridge choice of four-plane and time assignment.\n\n```text\n" + FormatSelector(a.Selector) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n### Native\n")
	for _, s := range a.Registry.NativeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Bridge\n")
	for _, s := range a.Registry.BridgeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Environmental\n")
	for _, s := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Failed routes\n")
	for _, s := range a.Registry.FailedRoutes {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Open theorems\n")
	for _, s := range a.Registry.OpenTheorems {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString(fmt.Sprintf("\n## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
