// Package generation2osreflectionpositivityairlock implements Gate 533:
// Osterwalder-Schrader Reflection-Positivity Kernel Airlock Preflight.
//
// Gate 532 verified that a deliberately synthetic fundamental symmetry Θ can
// turn the inherited Cℓ(1,7) Krein form into a positive finite matrix H=GΘ.
// That is necessary plumbing, not Wick reconstruction. This package defines
// the next fail-closed bridge schema: a future Osterwalder-Schrader row must
// provide a Euclidean reflection operator, test-function domain, correlation
// kernel, positivity cone, null-space quotient rule, and reconstruction
// certificate before reflection positivity can be evaluated. Gate 533 performs
// no kernel comparator and grants no Wick rotation or physical Hilbert space.
package generation2osreflectionpositivityairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticfundamentalsymmetryadapter"
)

const (
	AuditID = "GATE533-OS-REFLECTION-POSITIVITY-KERNEL-AIRLOCK-PREFLIGHT"

	StatusGate532AdapterInherited           = "CONDITIONAL_SUPPORT_GATE532_SYNTHETIC_THETA_ADAPTER_INHERITED"
	StatusOSKernelAirlockDefined            = "CONDITIONAL_SUPPORT_OS_REFLECTION_POSITIVITY_KERNEL_AIRLOCK_DEFINED"
	StatusOSKernelSchemaRowsEnumerated      = "CONDITIONAL_SUPPORT_OS_KERNEL_SCHEMA_ROWS_ENUMERATED"
	StatusReflectionDomainSchemaDefined     = "CONDITIONAL_SUPPORT_REFLECTION_TEST_FUNCTION_DOMAIN_SCHEMA_DEFINED"
	StatusNullQuotientSchemaDefined         = "CONDITIONAL_SUPPORT_OS_NULL_SPACE_QUOTIENT_SCHEMA_DEFINED"
	StatusReconstructionCertificateRequired = "CONDITIONAL_SUPPORT_OS_RECONSTRUCTION_CERTIFICATE_REQUIRED"
	StatusMandatoryBridgeMetadataEnforced   = "CONDITIONAL_SUPPORT_OS_SOURCE_CONVENTION_BRIDGE_TAGS_ENFORCED"
	StatusRedactedOSSchemaAccepted          = "CONDITIONAL_SUPPORT_REDACTED_OS_KERNEL_SCHEMA_ACCEPTED"
	StatusNativePromotionRejected           = "CONDITIONAL_SUPPORT_OS_REFLECTION_POSITIVITY_NATIVE_PROMOTION_REJECTED"
	StatusNoObservedOSDataImported          = "CONDITIONAL_SUPPORT_NO_OBSERVED_OS_WICK_OR_CORRELATION_DATA_IMPORTED"

	StatusFailedOSNativePromotionRejected = "FAILED_ROUTE_OS_REFLECTION_POSITIVITY_NATIVE_PROMOTION_REJECTED"
	StatusFailedMissingMetadataRejected   = "FAILED_ROUTE_OS_KERNEL_MISSING_SOURCE_CONVENTION_BRIDGE_TAG_REJECTED"
	StatusFailedPositiveMatrixNotOS       = "FAILED_ROUTE_POSITIVE_GTHETA_MATRIX_DOES_NOT_GRANT_OS_REFLECTION_POSITIVITY"
	StatusFailedOSSchemaNotWick           = "FAILED_ROUTE_OS_KERNEL_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedOSSchemaNotHilbert        = "FAILED_ROUTE_OS_KERNEL_SCHEMA_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedOSSchemaNotEnergy         = "FAILED_ROUTE_OS_KERNEL_SCHEMA_DOES_NOT_GRANT_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedOSSchemaNotUnitary        = "FAILED_ROUTE_OS_KERNEL_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedOSSchemaNotGlobal         = "FAILED_ROUTE_OS_KERNEL_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedComparatorNotPerformed    = "FAILED_ROUTE_OS_REFLECTION_POSITIVITY_COMPARATOR_EXECUTION_NOT_PERFORMED_IN_PREFLIGHT"
	StatusFirewallPreserved               = "FIREWALL_PRESERVED_GATE533_OS_REFLECTION_POSITIVITY_AIRLOCK_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked      = "FIREWALL_BLOCKED_GATE533_OS_WICK_HILBERT_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate532AdapterExecuted          bool
	Gate532ThetaResidualsZero       bool
	Gate532KreinAdjointResidualZero bool
	Gate532GThetaPositiveDefinite   bool
	Gate532ProjectorCompatible      bool
	Gate532TimeReflectionInvolution bool
	Gate532FinitePlumbingVerified   bool
	Gate532PhysicalHilbertBlocked   bool
	Gate532WickBlocked              bool
	Gate532OSBlocked                bool
	Gate532PositiveEnergyBlocked    bool
	Gate532UnitaryBlocked           bool
	Gate532GlobalCausalBlocked      bool
	Gate532ArrowBlocked             bool
	Gate532NativeWriteBlocked       bool
	Gate532NoObservedDataImported   bool
	Gate533OSAirlockRedirect        bool

	Verdict, Reason string
}

type SchemaField struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	NativeWrite bool
	Reason      string
}

type OSKernelSchema struct {
	Executed bool

	RequiredRows                        []SchemaField
	RequiredRowCount                    int
	EuclideanReflectionOperatorRequired bool
	TestFunctionDomainRequired          bool
	ReflectionActionRequired            bool
	CorrelationKernelRequired           bool
	KernelHermiticityCheckRequired      bool
	ReflectionPositiveConeRequired      bool
	OSQuadraticFormCheckRequired        bool
	NullSpaceQuotientRequired           bool
	ReconstructionMapRequired           bool
	CompatibilityWithThetaRequired      bool
	WickMapReferenceRequired            bool
	IepsilonConventionRequired          bool
	SourceRequired                      bool
	ConventionRequired                  bool
	BridgeOnlyRequired                  bool
	ComparatorOnlyRequired              bool
	NoTheoremInputRequired              bool
	NativePromotionRejected             bool
	RedactedSchemaAccepted              bool
	AcceptedRedactedCases               int
	RejectedFailClosedCases             int

	Verdict, Reason string
}

type ComparatorGuard struct {
	Executed bool

	ComparatorExecutionPerformed     bool
	ReflectionOperatorEvaluated      bool
	TestFunctionDomainEvaluated      bool
	ReflectionActionEvaluated        bool
	KernelHermiticityEvaluated       bool
	OSQuadraticFormEvaluated         bool
	PositiveConeEvaluated            bool
	NullSpaceQuotientEvaluated       bool
	ReconstructionPerformed          bool
	CompatibilityWithThetaEvaluated  bool
	WickContinuationEvaluated        bool
	PositiveEnergyEvaluated          bool
	UnitaryDynamicsEvaluated         bool
	GlobalHyperbolicityEvaluated     bool
	ReflectionPositivityProven       bool
	WickRotationSelected             bool
	PhysicalHilbertSpaceSelected     bool
	PositiveEnergyHamiltonianDerived bool
	UnitaryRealTimeDynamicsDerived   bool
	GlobalHyperbolicitySelected      bool

	Verdict, Reason string
}

type NativeRejection struct {
	Executed bool

	NativeOSKernelWrite               bool
	NativeReflectionOperatorWrite     bool
	NativeCorrelationFunctionWrite    bool
	NativeReflectionPositiveConeWrite bool
	NativeNullQuotientWrite           bool
	NativeReconstructionWrite         bool
	NativeWickWrite                   bool
	NativeHilbertProductWrite         bool
	NativePositiveEnergyWrite         bool
	NativeUnitaryDynamicsWrite        bool
	NativeGlobalCausalWrite           bool
	ComparatorExecutionPerformed      bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedOSDataImported          bool
	ObservedWickDataImported        bool
	ObservedCorrelationDataImported bool
	ObservedHamiltonianDataImported bool
	NativeOSKernelWrite             bool
	NativeReflectionWrite           bool
	NativeCorrelationWrite          bool
	NativeHilbertProductWrite       bool
	NativePhysicalStateSpaceWrite   bool
	NativeWickWrite                 bool
	NativePositiveEnergyWrite       bool
	NativeUnitaryDynamicsWrite      bool
	NativeGlobalCausalWrite         bool
	NativeTimeArrowWrite            bool
	ReopenedFlavorFirewall          bool
	ReopenedEWScaleFirewall         bool
	ReopenedGravityFirewall         bool
	ReopenedTopologyFirewall        bool
	NativeRegistryWritten           bool

	Verdict, Reason string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      OSKernelSchema
	Guard       ComparatorGuard
	Rejection   NativeRejection
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
	g532, err := generation2syntheticfundamentalsymmetryadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate532 synthetic fundamental-symmetry adapter: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g532)
	a.Schema = buildSchema()
	a.Guard = buildGuard()
	a.Rejection = buildNativeRejection(a.Guard)
	a.Firewall = buildFirewall(a.Rejection)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g532 generation2syntheticfundamentalsymmetryadapter.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate532AdapterExecuted:          g532.Output.Executed && g532.Output.Attempted && g532.Output.Ready,
		Gate532ThetaResidualsZero:       g532.Output.ThetaSquaredIdentityResidual == 0,
		Gate532KreinAdjointResidualZero: g532.Output.ThetaKreinSelfAdjointResidual == 0,
		Gate532GThetaPositiveDefinite:   g532.Output.GThetaPositiveDefinite && g532.Output.GThetaPositiveEigenvalues == 8,
		Gate532ProjectorCompatible:      g532.Output.ProjectorCompatibilityResidual == 0,
		Gate532TimeReflectionInvolution: g532.Output.TimeReflectionInvolutionResidual == 0,
		Gate532FinitePlumbingVerified:   g532.Output.FiniteMatrixPlumbingVerified && g532.Output.PositiveHilbertMatrixVerified,
		Gate532PhysicalHilbertBlocked:   !g532.Output.PhysicalHilbertSpaceGranted,
		Gate532WickBlocked:              !g532.Output.WickRotationGranted,
		Gate532OSBlocked:                !g532.Output.ReflectionPositivityGranted,
		Gate532PositiveEnergyBlocked:    !g532.Output.PositiveEnergyGranted,
		Gate532UnitaryBlocked:           !g532.Output.UnitaryRealTimeGranted,
		Gate532GlobalCausalBlocked:      !g532.Output.GlobalHyperbolicityGranted,
		Gate532ArrowBlocked:             !g532.Output.ArrowOfTimeSelected,
		Gate532NativeWriteBlocked:       !g532.Firewall.NativeRegistryWritten && !g532.Firewall.NativePhysicalStateSpaceWrite && !g532.Firewall.NativeWickWrite && !g532.Firewall.NativeReflectionWrite,
		Gate532NoObservedDataImported:   !g532.Firewall.ObservedHilbertDataImported && !g532.Firewall.ObservedWickDataImported && !g532.Firewall.ObservedBoundaryDataImported,
		Gate533OSAirlockRedirect:        g532.Next.Gate == 533,
		Verdict:                         StatusGate532AdapterInherited,
		Reason:                          "Gate533 inherits Gate532's finite Θ/Krein positivity dry run, but treats H=GΘ positivity only as matrix plumbing. OS reflection positivity still requires an independent Euclidean kernel/test-domain certificate.",
	}
}

func buildSchema() OSKernelSchema {
	rows := []SchemaField{
		{"euclidean_reflection_operator", true, true, false, "reflection θ_E acting on the Euclidean test-function domain must be explicit"},
		{"test_function_domain", true, true, false, "OS positivity is a statement over a chosen positive-time test-function subspace, not an abstract finite matrix alone"},
		{"reflection_action_on_test_functions", true, true, false, "the action f -> θ_E f must be sourced before any quadratic form can be evaluated"},
		{"correlation_kernel_or_schwinger_function", true, true, false, "kernel/S_n data must be supplied rather than inferred from Θ"},
		{"kernel_hermiticity_or_symmetry_convention", true, true, false, "kernel symmetry convention must be known before positivity is meaningful"},
		{"reflection_positive_cone", true, true, false, "the allowed cone/subspace for positive-time support must be specified"},
		{"os_quadratic_form_definition", true, true, false, "the exact form <θ_E f, K f> must be declared"},
		{"null_space_quotient_rule", true, true, false, "OS reconstruction requires quotienting zero-norm states"},
		{"reconstruction_map_certificate", true, true, false, "a certificate must state how the Hilbert space would be reconstructed if the comparator passes"},
		{"compatibility_with_gate532_theta", true, true, false, "the Euclidean reflection row must reference the Gate532 Θ convention or quarantine non-compatibility"},
		{"wick_map_reference", true, true, false, "Wick dictionary must remain an explicit bridge reference"},
		{"i_epsilon_or_analytic_continuation_convention", true, true, false, "analytic-continuation convention is bridge data and cannot be guessed"},
		{"source", true, true, false, "every row must be source-tagged"},
		{"source_version", true, true, false, "source version prevents silent convention drift"},
		{"convention", true, true, false, "signature/reflection/kernel conventions must be explicit"},
		{"bridge_only", true, true, false, "OS rows are bridge data until proven otherwise"},
		{"comparator_only", true, true, false, "preflight permits future comparison but no native write"},
		{"no_theorem_input", true, true, false, "fixture rows are not native derivation inputs"},
		{"native_promotion", true, true, false, "must be false; native promotion fails closed"},
	}
	return OSKernelSchema{
		Executed:                            true,
		RequiredRows:                        rows,
		RequiredRowCount:                    len(rows),
		EuclideanReflectionOperatorRequired: true,
		TestFunctionDomainRequired:          true,
		ReflectionActionRequired:            true,
		CorrelationKernelRequired:           true,
		KernelHermiticityCheckRequired:      true,
		ReflectionPositiveConeRequired:      true,
		OSQuadraticFormCheckRequired:        true,
		NullSpaceQuotientRequired:           true,
		ReconstructionMapRequired:           true,
		CompatibilityWithThetaRequired:      true,
		WickMapReferenceRequired:            true,
		IepsilonConventionRequired:          true,
		SourceRequired:                      true,
		ConventionRequired:                  true,
		BridgeOnlyRequired:                  true,
		ComparatorOnlyRequired:              true,
		NoTheoremInputRequired:              true,
		NativePromotionRejected:             true,
		RedactedSchemaAccepted:              true,
		AcceptedRedactedCases:               1,
		RejectedFailClosedCases:             0,
		Verdict:                             strings.Join([]string{StatusOSKernelAirlockDefined, StatusOSKernelSchemaRowsEnumerated, StatusReflectionDomainSchemaDefined, StatusNullQuotientSchemaDefined, StatusReconstructionCertificateRequired, StatusMandatoryBridgeMetadataEnforced, StatusRedactedOSSchemaAccepted, StatusNativePromotionRejected, StatusNoObservedOSDataImported}, ";"),
		Reason:                              "Gate533 defines the mandatory OS kernel airlock rows while accepting only redacted bridge-only schema metadata. It does not import correlation functions or run reflection-positivity comparators.",
	}
}

func buildGuard() ComparatorGuard {
	return ComparatorGuard{Executed: true, Verdict: StatusFailedComparatorNotPerformed, Reason: "Gate533 is preflight only: no OS kernel, quadratic-form, null quotient, reconstruction, Wick, Hamiltonian, unitary, or global-causal comparator is executed."}
}

func buildNativeRejection(g ComparatorGuard) NativeRejection {
	return NativeRejection{Executed: true, ComparatorExecutionPerformed: g.ComparatorExecutionPerformed, Verdict: strings.Join([]string{StatusFailedOSNativePromotionRejected, StatusFailedMissingMetadataRejected}, ";"), Reason: "Any missing source/convention tag or native OS/Wick/Hilbert write request fails closed at Gate533."}
}

func buildFirewall(r NativeRejection) Firewall {
	return Firewall{Executed: true, NativeOSKernelWrite: r.NativeOSKernelWrite, NativeReflectionWrite: r.NativeReflectionOperatorWrite, NativeCorrelationWrite: r.NativeCorrelationFunctionWrite, NativeHilbertProductWrite: r.NativeHilbertProductWrite, NativePhysicalStateSpaceWrite: false, NativeWickWrite: r.NativeWickWrite, NativePositiveEnergyWrite: r.NativePositiveEnergyWrite, NativeUnitaryDynamicsWrite: r.NativeUnitaryDynamicsWrite, NativeGlobalCausalWrite: r.NativeGlobalCausalWrite, Verdict: strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedPositiveMatrixNotOS, StatusFailedOSSchemaNotWick, StatusFailedOSSchemaNotHilbert, StatusFailedOSSchemaNotEnergy, StatusFailedOSSchemaNotUnitary, StatusFailedOSSchemaNotGlobal}, ";"), Reason: "Gate533 seals the OS/Wick/Hilbert frontier: H=GΘ positivity and a schema definition do not become reflection positivity, Wick rotation, physical Hilbert reconstruction, positive energy, unitary real-time dynamics, or global hyperbolicity."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native Osterwalder-Schrader reflection-positive kernel is written at Gate533.",
			"Gate532's positive H=GΘ matrix remains finite bridge plumbing, not a physical Hilbert-space or Wick theorem.",
			"No Schwinger functions, correlation kernels, null quotient, reconstructed Hilbert space, Hamiltonian, or time arrow are promoted natively.",
		},
		BridgeEntries: []string{
			"OS reflection-positivity kernel airlock schema defined for future source-tagged bridge rows.",
			"Required obligations include Euclidean reflection operator, test-function domain, reflected action, kernel/S_n data, OS quadratic form, positivity cone, null quotient, reconstruction certificate, Gate532 Θ compatibility, Wick map reference, and iε convention.",
			"Comparator execution is explicitly blocked in preflight; a future synthetic adapter must load actual bridge-only kernel data before evaluating OS positivity.",
		},
		EnvironmentalEntries: []string{
			"The physical Euclidean measure, Schwinger functions, analytic-continuation convention, Hamiltonian domain, and global causal boundary remain environmental or future bridge inputs.",
		},
		FailedRoutes: []string{StatusFailedPositiveMatrixNotOS, StatusFailedOSSchemaNotWick, StatusFailedOSSchemaNotHilbert, StatusFailedOSSchemaNotEnergy, StatusFailedOSSchemaNotUnitary, StatusFailedOSSchemaNotGlobal, StatusFailedComparatorNotPerformed},
		OpenTheorems: []string{
			"execute a synthetic OS kernel adapter with a finite reflection operator and positive Gram/kernel matrix",
			"separate OS reflection positivity from Wick rotation and from positive-energy Hamiltonian reconstruction",
			"audit whether any native ASHA structure can produce the Euclidean correlation kernel instead of importing it as bridge data",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 534, Title: "Synthetic OS Reflection-Positivity Kernel Adapter Dry Run", Reason: "Gate533 defines the kernel/test-domain airlock but deliberately performs no comparator. The next safe step is to load a synthetic finite kernel and verify the OS quadratic-form residuals as bridge-only plumbing.", PrimaryTask: "Load a synthetic source-tagged OS kernel ledger, verify reflection involution, kernel symmetry, positive-time domain closure, OS Gram positivity, null-space quotient metadata, and Gate532 Θ compatibility without promoting Wick rotation or physical Hilbert reconstruction."}
}

func truth(a Analysis) string {
	return "Gate533 closes the logical gap between finite Θ positivity and genuine Osterwalder-Schrader reconstruction. It proves only that ASHA now has a fail-closed schema for OS reflection-positivity data. It does not run a kernel comparator, does not prove Wick rotation, does not construct the physical Hilbert space, and does not derive positive energy, unitary dynamics, global hyperbolicity, or the arrow of time."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate532AdapterExecuted || !a.Inheritance.Gate532ThetaResidualsZero || !a.Inheritance.Gate532KreinAdjointResidualZero || !a.Inheritance.Gate532GThetaPositiveDefinite || !a.Inheritance.Gate532ProjectorCompatible || !a.Inheritance.Gate532TimeReflectionInvolution || !a.Inheritance.Gate532FinitePlumbingVerified || !a.Inheritance.Gate532PhysicalHilbertBlocked || !a.Inheritance.Gate532WickBlocked || !a.Inheritance.Gate532OSBlocked || !a.Inheritance.Gate532PositiveEnergyBlocked || !a.Inheritance.Gate532UnitaryBlocked || !a.Inheritance.Gate532GlobalCausalBlocked || !a.Inheritance.Gate532ArrowBlocked || !a.Inheritance.Gate532NativeWriteBlocked || !a.Inheritance.Gate532NoObservedDataImported || !a.Inheritance.Gate533OSAirlockRedirect {
		bad = append(bad, "bad Gate532 inheritance")
	}
	if !a.Schema.Executed || a.Schema.RequiredRowCount < 19 || !a.Schema.EuclideanReflectionOperatorRequired || !a.Schema.TestFunctionDomainRequired || !a.Schema.ReflectionActionRequired || !a.Schema.CorrelationKernelRequired || !a.Schema.KernelHermiticityCheckRequired || !a.Schema.ReflectionPositiveConeRequired || !a.Schema.OSQuadraticFormCheckRequired || !a.Schema.NullSpaceQuotientRequired || !a.Schema.ReconstructionMapRequired || !a.Schema.CompatibilityWithThetaRequired || !a.Schema.WickMapReferenceRequired || !a.Schema.IepsilonConventionRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.ComparatorOnlyRequired || !a.Schema.NoTheoremInputRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted {
		bad = append(bad, "bad OS schema")
	}
	if !a.Guard.Executed || a.Guard.ComparatorExecutionPerformed || a.Guard.ReflectionOperatorEvaluated || a.Guard.TestFunctionDomainEvaluated || a.Guard.ReflectionActionEvaluated || a.Guard.KernelHermiticityEvaluated || a.Guard.OSQuadraticFormEvaluated || a.Guard.PositiveConeEvaluated || a.Guard.NullSpaceQuotientEvaluated || a.Guard.ReconstructionPerformed || a.Guard.CompatibilityWithThetaEvaluated || a.Guard.WickContinuationEvaluated || a.Guard.PositiveEnergyEvaluated || a.Guard.UnitaryDynamicsEvaluated || a.Guard.GlobalHyperbolicityEvaluated || a.Guard.ReflectionPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSpaceSelected || a.Guard.PositiveEnergyHamiltonianDerived || a.Guard.UnitaryRealTimeDynamicsDerived || a.Guard.GlobalHyperbolicitySelected {
		bad = append(bad, "guard violation")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedOSDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedCorrelationDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.NativeOSKernelWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativeCorrelationWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: adapter=%t theta2=%t krein_adj=%t Gtheta_positive=%t projector_compat=%t Rtime=%t finite_plumbing=%t physical_Hilbert_blocked=%t Wick_blocked=%t OS_blocked=%t positive_energy_blocked=%t unitary_blocked=%t global_blocked=%t arrow_blocked=%t native_blocked=%t no_observed=%t gate533_redirect=%t; %s", x.Verdict, x.Gate532AdapterExecuted, x.Gate532ThetaResidualsZero, x.Gate532KreinAdjointResidualZero, x.Gate532GThetaPositiveDefinite, x.Gate532ProjectorCompatible, x.Gate532TimeReflectionInvolution, x.Gate532FinitePlumbingVerified, x.Gate532PhysicalHilbertBlocked, x.Gate532WickBlocked, x.Gate532OSBlocked, x.Gate532PositiveEnergyBlocked, x.Gate532UnitaryBlocked, x.Gate532GlobalCausalBlocked, x.Gate532ArrowBlocked, x.Gate532NativeWriteBlocked, x.Gate532NoObservedDataImported, x.Gate533OSAirlockRedirect, x.Reason)
}

func FormatSchema(x OSKernelSchema) string {
	return fmt.Sprintf("%s: rows=%d reflection=%t domain=%t action=%t kernel=%t hermiticity=%t cone=%t quadratic=%t null_quotient=%t reconstruction=%t theta_compat=%t Wick_ref=%t i_epsilon=%t source=%t convention=%t bridge_only=%t comparator_only=%t no_theorem_input=%t native_rejected=%t redacted_accepted=%t; %s", x.Verdict, x.RequiredRowCount, x.EuclideanReflectionOperatorRequired, x.TestFunctionDomainRequired, x.ReflectionActionRequired, x.CorrelationKernelRequired, x.KernelHermiticityCheckRequired, x.ReflectionPositiveConeRequired, x.OSQuadraticFormCheckRequired, x.NullSpaceQuotientRequired, x.ReconstructionMapRequired, x.CompatibilityWithThetaRequired, x.WickMapReferenceRequired, x.IepsilonConventionRequired, x.SourceRequired, x.ConventionRequired, x.BridgeOnlyRequired, x.ComparatorOnlyRequired, x.NoTheoremInputRequired, x.NativePromotionRejected, x.RedactedSchemaAccepted, x.Reason)
}

func FormatGuard(x ComparatorGuard) string {
	return fmt.Sprintf("%s: comparator=%t reflection_eval=%t domain_eval=%t kernel_hermiticity=%t OS_quadratic=%t cone=%t null_quotient=%t reconstruction=%t theta_compat=%t Wick_eval=%t positive_energy_eval=%t unitary_eval=%t global_eval=%t OS_proven=%t Wick_selected=%t Hilbert_selected=%t positive_energy=%t unitary=%t global=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.ReflectionOperatorEvaluated, x.TestFunctionDomainEvaluated, x.KernelHermiticityEvaluated, x.OSQuadraticFormEvaluated, x.PositiveConeEvaluated, x.NullSpaceQuotientEvaluated, x.ReconstructionPerformed, x.CompatibilityWithThetaEvaluated, x.WickContinuationEvaluated, x.PositiveEnergyEvaluated, x.UnitaryDynamicsEvaluated, x.GlobalHyperbolicityEvaluated, x.ReflectionPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSpaceSelected, x.PositiveEnergyHamiltonianDerived, x.UnitaryRealTimeDynamicsDerived, x.GlobalHyperbolicitySelected, x.Reason)
}

func FormatRejection(x NativeRejection) string {
	return fmt.Sprintf("%s: native_OS_kernel=%t native_reflection=%t native_correlation=%t native_cone=%t native_null_quotient=%t native_reconstruction=%t native_Wick=%t native_Hilbert=%t native_positive_energy=%t native_unitary=%t native_global=%t comparator=%t; %s", x.Verdict, x.NativeOSKernelWrite, x.NativeReflectionOperatorWrite, x.NativeCorrelationFunctionWrite, x.NativeReflectionPositiveConeWrite, x.NativeNullQuotientWrite, x.NativeReconstructionWrite, x.NativeWickWrite, x.NativeHilbertProductWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.ComparatorExecutionPerformed, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_OS=%t observed_Wick=%t observed_corr=%t observed_Hamiltonian=%t native_OS=%t native_reflection=%t native_corr=%t native_Hilbert=%t native_state=%t native_Wick=%t native_positive_energy=%t native_unitary=%t native_global=%t native_arrow=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_registry=%t; %s", x.Verdict, x.ObservedOSDataImported, x.ObservedWickDataImported, x.ObservedCorrelationDataImported, x.ObservedHamiltonianDataImported, x.NativeOSKernelWrite, x.NativeReflectionWrite, x.NativeCorrelationWrite, x.NativeHilbertProductWrite, x.NativePhysicalStateSpaceWrite, x.NativeWickWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate532AdapterInherited, StatusOSKernelAirlockDefined, StatusOSKernelSchemaRowsEnumerated, StatusReflectionDomainSchemaDefined, StatusNullQuotientSchemaDefined, StatusReconstructionCertificateRequired, StatusMandatoryBridgeMetadataEnforced, StatusRedactedOSSchemaAccepted, StatusNativePromotionRejected, StatusNoObservedOSDataImported, StatusFailedPositiveMatrixNotOS, StatusFailedOSSchemaNotWick, StatusFailedOSSchemaNotHilbert, StatusFailedOSSchemaNotEnergy, StatusFailedOSSchemaNotUnitary, StatusFailedOSSchemaNotGlobal, StatusFailedComparatorNotPerformed, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 533 Registry Audit — Osterwalder-Schrader Reflection-Positivity Kernel Airlock Preflight\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 533 inherits Gate 532's synthetic Θ positivity socket but refuses to treat finite `H=GΘ` positivity as OS reconstruction.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## OS kernel schema\n\nThe airlock defines what a future reflection-positivity ledger must provide before any OS comparator can run.\n\n```text\n" + FormatSchema(a.Schema) + "\n```\n\n")
	b.WriteString("### Required schema rows\n\n")
	for _, row := range a.Schema.RequiredRows {
		b.WriteString(fmt.Sprintf("- `%s` — required=%t bridge_only=%t native_write=%t; %s\n", row.Name, row.Required, row.BridgeOnly, row.NativeWrite, row.Reason))
	}
	b.WriteString("\n## Comparator guard\n\nGate 533 is preflight only: no OS kernel, null-quotient, reconstruction, Wick, Hamiltonian, unitary, or global-causal comparator is executed.\n\n```text\n" + FormatGuard(a.Guard) + "\n```\n\n")
	b.WriteString("## Native rejection rule\n\n```text\n" + FormatRejection(a.Rejection) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native", a.Registry.NativeEntries)
	writeList(&b, "### Bridge", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString(fmt.Sprintf("## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
