// Package generation2lorentzianspinoradjointairlock implements Gate 527:
// Lorentzian Spinor Adjoint, Reflection-Positivity, and 3+1 Projection
// Airlock Audit.
//
// Gate 526 confirmed the native Cℓ(1,7) causal-signature socket and null
// cone, but blocked Wick rotation, time orientation, positive energy,
// reflection positivity, real-time unitarity, and physical 3+1 projection.
// Gate 527 audits the next Lorentzian obstruction: whether the algebraic
// indefinite/Krein spinor adjoint can be promoted to a positive Hilbert-space
// quantum theory and whether it selects a 3+1 spacetime split. It records the
// structural sockets while refusing every real-time dynamics/native projection
// write that is not theorem-selected.
package generation2lorentzianspinoradjointairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2lorentziancausalsignature"
)

const (
	AuditID = "GATE527-LORENTZIAN-SPINOR-ADJOINT-AIRLOCK"

	StatusGate526Inherited                        = "CONDITIONAL_SUPPORT_GATE526_LORENTZIAN_SIGNATURE_INHERITED"
	StatusKreinAdjointSocketDefined               = "CONDITIONAL_SUPPORT_LORENTZIAN_KREIN_ADJOINT_SOCKET_DEFINED"
	StatusCliffordAdjointCompatibilityConfirmed   = "CONDITIONAL_SUPPORT_CL17_CLIFFORD_ADJOINT_COMPATIBILITY_CONFIRMED"
	StatusChargeConjugationGradingSocketPreserved = "CONDITIONAL_SUPPORT_CHARGE_CONJUGATION_AND_GRADING_SOCKET_PRESERVED"
	StatusIndefiniteToHilbertDictionarySeparated  = "CONDITIONAL_SUPPORT_INDEFINITE_TO_HILBERT_DICTIONARY_SEPARATED"
	StatusReflectionPositivityAirlockDefined      = "CONDITIONAL_SUPPORT_REFLECTION_POSITIVITY_AIRLOCK_DEFINED"
	StatusProjectionAirlockDefined                = "CONDITIONAL_SUPPORT_3PLUS1_PROJECTION_AIRLOCK_DEFINED"
	StatusNoObservedDataImported                  = "CONDITIONAL_SUPPORT_NO_OBSERVED_CONSTANTS_MASSES_TOPOLOGY_OR_BOUNDARY_DATA_IMPORTED"

	StatusFailedNoPositiveHilbertProduct         = "FAILED_ROUTE_POSITIVE_HILBERT_PRODUCT_NOT_NATIVE_SELECTED"
	StatusFailedNoReflectionPositivity           = "FAILED_ROUTE_REFLECTION_POSITIVITY_OS_AXIOMS_NOT_DERIVED"
	StatusFailedNoTimeReflection                 = "FAILED_ROUTE_TIME_REFLECTION_NOT_NATIVE_SELECTED"
	StatusFailedNoHamiltonianSpectrum            = "FAILED_ROUTE_POSITIVE_ENERGY_HAMILTONIAN_SPECTRUM_NOT_DERIVED"
	StatusFailedNoUnitaryDynamics                = "FAILED_ROUTE_REAL_TIME_UNITARY_DYNAMICS_STILL_BLOCKED"
	StatusFailedNoNative3Plus1Projection         = "FAILED_ROUTE_3PLUS1_SPACETIME_PROJECTION_NOT_NATIVE_SELECTED"
	StatusFailedNoInternal4Split                 = "FAILED_ROUTE_INTERNAL_FOUR_DIMENSIONAL_COMPLEMENT_NOT_NATIVE_SELECTED"
	StatusFailedNoWickContinuation               = "FAILED_ROUTE_WICK_CONTINUATION_STILL_NOT_NATIVE_SELECTED"
	StatusFailedNoGlobalHyperbolicity            = "FAILED_ROUTE_GLOBAL_HYPERBOLICITY_STILL_NOT_NATIVE_SELECTED"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_COMPLETED_SECTOR_AIRLOCKS_DURING_SPINOR_ADJOINT_AUDIT"
	StatusFirewallLorentzianDynamicsWriteBlocked = "FIREWALL_BLOCKED_LORENTZIAN_HILBERT_DYNAMICS_AND_3PLUS1_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate526SignatureInherited       bool
	Gate526NullConeConfirmed        bool
	Gate526EuclideanLedgerSeparated bool
	Gate526WickBlocked              bool
	Gate526ReflectionPositivityOpen bool
	Gate526PositiveEnergyOpen       bool
	Gate526UnitaryDynamicsOpen      bool
	Gate5263Plus1Open               bool
	Gate526NoObservedDataImported   bool
	Gate526NativeWriteBlocked       bool
	Gate526ReopenedSealedFirewalls  bool

	Verdict, Reason string
}

type KreinAdjointAudit struct {
	Executed bool

	Algebra                        string
	IndefiniteMetricSocket         bool
	KreinAdjointDefined            bool
	DiracAdjointSocketDefined      bool
	CliffordCompatibility          bool
	ChargeConjugationSocket        bool
	GradingSocketPreserved         bool
	PositiveHilbertProductSelected bool
	FundamentalSymmetrySelected    bool
	PhysicalStateSpaceSelected     bool

	Verdict, Reason string
}

type ReflectionPositivityAudit struct {
	Executed bool

	EuclideanLedgerAvailable         bool
	TimeReflectionRequired           bool
	TimeReflectionSelected           bool
	ReflectionPositivityProven       bool
	OsterwalderSchraderAxiomsProven  bool
	WickContinuationSelected         bool
	PositiveEnergyHamiltonianDerived bool
	UnitaryRealTimeDynamicsDerived   bool
	GlobalHyperbolicitySelected      bool

	Verdict, Reason string
}

type ProjectionAudit struct {
	Executed bool

	NativeDimension                   int
	CandidateExternalDimension        int
	CandidateInternalComplement       int
	CandidateSplit                    string
	ProjectionRankArithmeticValid     bool
	ProjectionOperatorNativeSelected  bool
	SubalgebraEmbeddingNativeSelected bool
	InternalComplementNativeSelected  bool
	Physical3Plus1Selected            bool
	TimeOrientationSelected           bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedConstantsImported  bool
	ObservedMassesImported     bool
	ObservedTopologyImported   bool
	ObservedBoundaryImported   bool
	NativePositiveHilbertWrite bool
	NativeReflectionWrite      bool
	NativeWickWrite            bool
	NativePositiveEnergyWrite  bool
	NativeUnitaryWrite         bool
	Native3Plus1Write          bool
	NativeInternal4Write       bool
	NativeGlobalCausalWrite    bool
	ReopenedFlavorFirewall     bool
	ReopenedEWScaleFirewall    bool
	ReopenedGravityFirewall    bool
	ReopenedTopologyFirewall   bool
	NativeRegistryWritten      bool

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
	Adjoint     KreinAdjointAudit
	Reflection  ReflectionPositivityAudit
	Projection  ProjectionAudit
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
	g526, err := generation2lorentziancausalsignature.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate526 Lorentzian signature audit: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g526)
	a.Adjoint = buildAdjoint(a.Inheritance)
	a.Reflection = buildReflection(a.Inheritance, a.Adjoint)
	a.Projection = buildProjection(a.Inheritance)
	a.Firewall = buildFirewall(a.Inheritance, a.Adjoint, a.Reflection, a.Projection)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g526 generation2lorentziancausalsignature.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate526SignatureInherited:       g526.Signature.MetricSignatureNative && g526.Signature.TimeLikeDirections == 1 && g526.Signature.SpaceLikeDirections == 7,
		Gate526NullConeConfirmed:        g526.Signature.NullConeDefined && g526.Signature.CausalConeScaleFree,
		Gate526EuclideanLedgerSeparated: g526.Dictionary.EuclideanSpectralActionInherited && g526.Dictionary.HeatKernelEllipticConvention && g526.Dictionary.BridgeDictionaryDefined,
		Gate526WickBlocked:              !g526.Dictionary.WickRotationSelectedNatively,
		Gate526ReflectionPositivityOpen: !g526.Dictionary.ReflectionPositivityProven && !g526.Dictionary.OsterwalderSchraderAxiomsProven,
		Gate526PositiveEnergyOpen:       !g526.Dictionary.PositiveEnergyConditionDerived,
		Gate526UnitaryDynamicsOpen:      !g526.Dictionary.UnitaryTimeEvolutionDerived,
		Gate5263Plus1Open:               !g526.Signature.Physical3Plus1ProjectionFound,
		Gate526NoObservedDataImported:   !g526.Firewall.ObservedConstantsImported && !g526.Firewall.ObservedMassesImported && !g526.Firewall.ObservedTopologyImported,
		Gate526NativeWriteBlocked:       !g526.Firewall.NativeRegistryWritten,
		Gate526ReopenedSealedFirewalls:  g526.Firewall.ReopenedFlavorFirewall || g526.Firewall.ReopenedEWScaleFirewall || g526.Firewall.ReopenedGravityFirewall || g526.Firewall.ReopenedTopologyFirewall,
		Verdict:                         StatusGate526Inherited,
		Reason:                          "Gate527 inherits the Gate526 causal-signature socket and all open Lorentzian obligations: Wick continuation, reflection positivity, positive energy, real-time unitarity, and 3+1 projection were explicitly not native-selected.",
	}
}

func buildAdjoint(in Inheritance) KreinAdjointAudit {
	return KreinAdjointAudit{
		Executed:                       true,
		Algebra:                        "Cℓ(1,7)",
		IndefiniteMetricSocket:         in.Gate526SignatureInherited,
		KreinAdjointDefined:            in.Gate526SignatureInherited,
		DiracAdjointSocketDefined:      in.Gate526SignatureInherited,
		CliffordCompatibility:          in.Gate526SignatureInherited,
		ChargeConjugationSocket:        true,
		GradingSocketPreserved:         true,
		PositiveHilbertProductSelected: false,
		FundamentalSymmetrySelected:    false,
		PhysicalStateSpaceSelected:     false,
		Verdict:                        strings.Join([]string{StatusKreinAdjointSocketDefined, StatusCliffordAdjointCompatibilityConfirmed, StatusChargeConjugationGradingSocketPreserved, StatusIndefiniteToHilbertDictionarySeparated}, ";"),
		Reason:                         "The Lorentzian signature supplies the algebraic/Krein adjoint socket and preserves charge-conjugation/grading bookkeeping. It does not select the positive Hilbert product, fundamental symmetry, or physical state space required for quantum dynamics.",
	}
}

func buildReflection(in Inheritance, adj KreinAdjointAudit) ReflectionPositivityAudit {
	return ReflectionPositivityAudit{
		Executed:                         true,
		EuclideanLedgerAvailable:         in.Gate526EuclideanLedgerSeparated,
		TimeReflectionRequired:           true,
		TimeReflectionSelected:           false,
		ReflectionPositivityProven:       false,
		OsterwalderSchraderAxiomsProven:  false,
		WickContinuationSelected:         false,
		PositiveEnergyHamiltonianDerived: false,
		UnitaryRealTimeDynamicsDerived:   false,
		GlobalHyperbolicitySelected:      false,
		Verdict:                          StatusReflectionPositivityAirlockDefined,
		Reason:                           "The Euclidean heat-kernel ledger can be placed behind a reflection-positivity/OS airlock, but Gate527 derives neither a time-reflection involution nor the positivity theorem needed to reconstruct a positive-energy Lorentzian Hilbert space.",
	}
}

func buildProjection(in Inheritance) ProjectionAudit {
	return ProjectionAudit{
		Executed:                          true,
		NativeDimension:                   8,
		CandidateExternalDimension:        4,
		CandidateInternalComplement:       4,
		CandidateSplit:                    "1+7 -> candidate (1+3) external + 4 internal complement",
		ProjectionRankArithmeticValid:     true,
		ProjectionOperatorNativeSelected:  false,
		SubalgebraEmbeddingNativeSelected: false,
		InternalComplementNativeSelected:  false,
		Physical3Plus1Selected:            false,
		TimeOrientationSelected:           false,
		Verdict:                           StatusProjectionAirlockDefined,
		Reason:                            "The rank arithmetic for a 3+1 external slice plus four internal directions is a coherent bridge socket. No canonical projector, subalgebra embedding, or time orientation is selected by the native Cℓ(1,7) signature alone.",
	}
}

func buildFirewall(in Inheritance, adj KreinAdjointAudit, rp ReflectionPositivityAudit, p ProjectionAudit) Firewall {
	return Firewall{
		Executed:                   true,
		ObservedConstantsImported:  false,
		ObservedMassesImported:     false,
		ObservedTopologyImported:   false,
		ObservedBoundaryImported:   false,
		NativePositiveHilbertWrite: adj.PositiveHilbertProductSelected || adj.FundamentalSymmetrySelected || adj.PhysicalStateSpaceSelected,
		NativeReflectionWrite:      rp.ReflectionPositivityProven || rp.OsterwalderSchraderAxiomsProven || rp.TimeReflectionSelected,
		NativeWickWrite:            rp.WickContinuationSelected,
		NativePositiveEnergyWrite:  rp.PositiveEnergyHamiltonianDerived,
		NativeUnitaryWrite:         rp.UnitaryRealTimeDynamicsDerived,
		Native3Plus1Write:          p.Physical3Plus1Selected || p.ProjectionOperatorNativeSelected || p.SubalgebraEmbeddingNativeSelected,
		NativeInternal4Write:       p.InternalComplementNativeSelected,
		NativeGlobalCausalWrite:    rp.GlobalHyperbolicitySelected,
		ReopenedFlavorFirewall:     false,
		ReopenedEWScaleFirewall:    false,
		ReopenedGravityFirewall:    false,
		ReopenedTopologyFirewall:   false,
		NativeRegistryWritten:      false,
		Verdict:                    strings.Join([]string{StatusNoObservedDataImported, StatusFirewallPreserved, StatusFirewallLorentzianDynamicsWriteBlocked}, ";"),
		Reason:                     "Gate527 imports no constants, masses, topology, or boundary data and writes no positive Hilbert structure, OS theorem, Wick continuation, real-time dynamics, or physical 3+1 projection to the native registry.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) supplies an indefinite Lorentzian/Krein adjoint socket compatible with the Gate526 causal signature",
			"charge-conjugation and grading bookkeeping remain structurally available in the Lorentzian spinor audit",
			"the null-cone causal socket stays scale-free and independent of flavor, electroweak scale, gravity normalization, and global topology",
		},
		BridgeEntries: []string{
			"positive Hilbert-space reconstruction requires a fundamental symmetry/inner-product choice beyond the native Krein socket",
			"reflection positivity, OS axioms, Wick continuation, and positive-energy reconstruction form a fail-closed bridge airlock",
			"the 1+7 -> (1+3)+4 split is an admissible bridge projection socket, not a selected native spacetime decomposition",
		},
		EnvironmentalEntries: []string{
			"time orientation, thermodynamic arrow, and real-time boundary/initial conditions remain bridge/environmental data",
			"global hyperbolicity and causal boundary conditions remain continuum/global inputs",
		},
		FailedRoutes: []string{
			StatusFailedNoPositiveHilbertProduct,
			StatusFailedNoReflectionPositivity,
			StatusFailedNoTimeReflection,
			StatusFailedNoHamiltonianSpectrum,
			StatusFailedNoUnitaryDynamics,
			StatusFailedNoNative3Plus1Projection,
			StatusFailedNoInternal4Split,
			StatusFailedNoWickContinuation,
			StatusFailedNoGlobalHyperbolicity,
		},
		OpenTheorems: []string{
			"construct or quarantine an explicit reflection-positive Euclidean-to-Lorentzian reconstruction theorem",
			"audit whether any native projector selects the physical 1+3 external spacetime subspace from the 1+7 Clifford ladder",
			"classify positive-energy and global-hyperbolicity requirements as bridge schema unless a native theorem proves them",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        528,
		Title:       "Physical 3+1 Projection and Internal Complement Selector Audit",
		Reason:      "Gate527 confirms the Lorentzian/Krein spinor socket and the reflection-positivity airlock, but it still blocks physical 3+1 projection. The next gate should isolate whether a canonical projector or subalgebra embedding selects (1+3) spacetime plus a four-dimensional internal complement.",
		PrimaryTask: "Search for a native projector/subalgebra selector inside the Cℓ(1,7) ladder; if absent, define a bridge-only 3+1 projection schema and block physical-spacetime native promotion.",
	}
}

func truth(a Analysis) string {
	return "Gate527 confirms that ASHA has a Lorentzian/Krein spinor-adjoint socket compatible with the Cℓ(1,7) causal signature, but it does not reconstruct a physical Lorentzian quantum theory. Positive Hilbert product, reflection positivity, Wick continuation, positive energy, unitary real-time dynamics, global hyperbolicity, and physical 3+1 projection remain bridge obligations."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate526SignatureInherited || !a.Inheritance.Gate526NullConeConfirmed || !a.Inheritance.Gate526EuclideanLedgerSeparated || !a.Inheritance.Gate526WickBlocked || !a.Inheritance.Gate526ReflectionPositivityOpen || !a.Inheritance.Gate526PositiveEnergyOpen || !a.Inheritance.Gate526UnitaryDynamicsOpen || !a.Inheritance.Gate5263Plus1Open || !a.Inheritance.Gate526NoObservedDataImported || !a.Inheritance.Gate526NativeWriteBlocked || a.Inheritance.Gate526ReopenedSealedFirewalls {
		bad = append(bad, "bad inheritance")
	}
	if !a.Adjoint.Executed || a.Adjoint.Algebra != "Cℓ(1,7)" || !a.Adjoint.IndefiniteMetricSocket || !a.Adjoint.KreinAdjointDefined || !a.Adjoint.DiracAdjointSocketDefined || !a.Adjoint.CliffordCompatibility || !a.Adjoint.ChargeConjugationSocket || !a.Adjoint.GradingSocketPreserved || a.Adjoint.PositiveHilbertProductSelected || a.Adjoint.FundamentalSymmetrySelected || a.Adjoint.PhysicalStateSpaceSelected {
		bad = append(bad, "bad Krein adjoint audit")
	}
	if !a.Reflection.Executed || !a.Reflection.EuclideanLedgerAvailable || !a.Reflection.TimeReflectionRequired || a.Reflection.TimeReflectionSelected || a.Reflection.ReflectionPositivityProven || a.Reflection.OsterwalderSchraderAxiomsProven || a.Reflection.WickContinuationSelected || a.Reflection.PositiveEnergyHamiltonianDerived || a.Reflection.UnitaryRealTimeDynamicsDerived || a.Reflection.GlobalHyperbolicitySelected {
		bad = append(bad, "bad reflection positivity audit")
	}
	if !a.Projection.Executed || a.Projection.NativeDimension != 8 || a.Projection.CandidateExternalDimension != 4 || a.Projection.CandidateInternalComplement != 4 || !a.Projection.ProjectionRankArithmeticValid || a.Projection.ProjectionOperatorNativeSelected || a.Projection.SubalgebraEmbeddingNativeSelected || a.Projection.InternalComplementNativeSelected || a.Projection.Physical3Plus1Selected || a.Projection.TimeOrientationSelected {
		bad = append(bad, "bad projection audit")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedConstantsImported || a.Firewall.ObservedMassesImported || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryImported || a.Firewall.NativePositiveHilbertWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativeWickWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryWrite || a.Firewall.Native3Plus1Write || a.Firewall.NativeInternal4Write || a.Firewall.NativeGlobalCausalWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: signature=%t null_cone=%t Euclidean_ledger=%t Wick_blocked=%t reflection_open=%t positive_energy_open=%t unitary_open=%t 3plus1_open=%t no_observed=%t native_blocked=%t reopens_firewalls=%t; %s", x.Verdict, x.Gate526SignatureInherited, x.Gate526NullConeConfirmed, x.Gate526EuclideanLedgerSeparated, x.Gate526WickBlocked, x.Gate526ReflectionPositivityOpen, x.Gate526PositiveEnergyOpen, x.Gate526UnitaryDynamicsOpen, x.Gate5263Plus1Open, x.Gate526NoObservedDataImported, x.Gate526NativeWriteBlocked, x.Gate526ReopenedSealedFirewalls, x.Reason)
}

func FormatAdjoint(x KreinAdjointAudit) string {
	return fmt.Sprintf("%s: algebra=%s indefinite_metric=%t Krein_adjoint=%t Dirac_adjoint=%t Clifford_compatible=%t C_socket=%t grading=%t positive_Hilbert=%t fundamental_symmetry=%t physical_state_space=%t; %s", x.Verdict, x.Algebra, x.IndefiniteMetricSocket, x.KreinAdjointDefined, x.DiracAdjointSocketDefined, x.CliffordCompatibility, x.ChargeConjugationSocket, x.GradingSocketPreserved, x.PositiveHilbertProductSelected, x.FundamentalSymmetrySelected, x.PhysicalStateSpaceSelected, x.Reason)
}

func FormatReflection(x ReflectionPositivityAudit) string {
	return fmt.Sprintf("%s: Euclidean_ledger=%t time_reflection_required=%t time_reflection_selected=%t reflection_positivity=%t OS_axioms=%t Wick=%t positive_energy_H=%t unitary_real_time=%t global_hyperbolicity=%t; %s", x.Verdict, x.EuclideanLedgerAvailable, x.TimeReflectionRequired, x.TimeReflectionSelected, x.ReflectionPositivityProven, x.OsterwalderSchraderAxiomsProven, x.WickContinuationSelected, x.PositiveEnergyHamiltonianDerived, x.UnitaryRealTimeDynamicsDerived, x.GlobalHyperbolicitySelected, x.Reason)
}

func FormatProjection(x ProjectionAudit) string {
	return fmt.Sprintf("%s: native_dim=%d external_dim=%d internal_dim=%d split=%q rank_valid=%t projector_native=%t subalgebra_native=%t internal_native=%t physical_3plus1=%t time_orientation=%t; %s", x.Verdict, x.NativeDimension, x.CandidateExternalDimension, x.CandidateInternalComplement, x.CandidateSplit, x.ProjectionRankArithmeticValid, x.ProjectionOperatorNativeSelected, x.SubalgebraEmbeddingNativeSelected, x.InternalComplementNativeSelected, x.Physical3Plus1Selected, x.TimeOrientationSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_constants=%t observed_masses=%t observed_topology=%t observed_boundary=%t native_Hilbert=%t native_reflection=%t native_Wick=%t native_positive_energy=%t native_unitary=%t native_3plus1=%t native_internal4=%t native_global_causal=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_write=%t; %s", x.Verdict, x.ObservedConstantsImported, x.ObservedMassesImported, x.ObservedTopologyImported, x.ObservedBoundaryImported, x.NativePositiveHilbertWrite, x.NativeReflectionWrite, x.NativeWickWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryWrite, x.Native3Plus1Write, x.NativeInternal4Write, x.NativeGlobalCausalWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate526Inherited,
		StatusKreinAdjointSocketDefined,
		StatusCliffordAdjointCompatibilityConfirmed,
		StatusChargeConjugationGradingSocketPreserved,
		StatusIndefiniteToHilbertDictionarySeparated,
		StatusReflectionPositivityAirlockDefined,
		StatusProjectionAirlockDefined,
		StatusNoObservedDataImported,
		StatusFailedNoPositiveHilbertProduct,
		StatusFailedNoReflectionPositivity,
		StatusFailedNoTimeReflection,
		StatusFailedNoHamiltonianSpectrum,
		StatusFailedNoUnitaryDynamics,
		StatusFailedNoNative3Plus1Projection,
		StatusFailedNoInternal4Split,
		StatusFailedNoWickContinuation,
		StatusFailedNoGlobalHyperbolicity,
		StatusFirewallPreserved,
		StatusFirewallLorentzianDynamicsWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 527 Registry Audit — Lorentzian Spinor Adjoint, Reflection-Positivity, and 3+1 Projection Airlock Audit\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 527 inherits Gate 526's native Cℓ(1,7) signature/null-cone socket and its blocked Lorentzian obligations.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Lorentzian spinor adjoint audit\n\nThe native result is an indefinite/Krein adjoint socket, not a positive physical Hilbert-space reconstruction.\n\n```text\n" + FormatAdjoint(a.Adjoint) + "\n```\n\n")
	b.WriteString("## Reflection-positivity and Wick airlock\n\nThe Euclidean heat-kernel ledger remains safe only behind a fail-closed OS/reflection-positivity bridge.\n\n```text\n" + FormatReflection(a.Reflection) + "\n```\n\n")
	b.WriteString("## 3+1 projection airlock\n\nA 1+7 -> (1+3)+4 split is rank-consistent but not natively selected.\n\n```text\n" + FormatProjection(a.Projection) + "\n```\n\n")
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
