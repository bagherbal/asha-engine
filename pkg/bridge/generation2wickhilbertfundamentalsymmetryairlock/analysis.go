// Package generation2wickhilbertfundamentalsymmetryairlock implements Gate 531:
// Wick/Hilbert Fundamental-Symmetry Airlock Preflight.
//
// Gate 530 proved that a synthetic 3+1 projector can pass the dimensional
// socket residual checks, but it deliberately left the Lorentzian quantum
// obligations untouched. This package defines the fail-closed preflight schema
// for importing a Krein-to-Hilbert fundamental symmetry, Wick/reflection data,
// and time-orientation conventions as bridge data. It performs no Hilbert
// reconstruction and grants no Wick rotation, positive energy, unitarity, or
// physical state-space theorem.
package generation2wickhilbertfundamentalsymmetryairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2synthetic3plus1projectionadapter"
)

const (
	AuditID = "GATE531-WICK-HILBERT-FUNDAMENTAL-SYMMETRY-AIRLOCK-PREFLIGHT"

	StatusGate530AdapterInherited            = "CONDITIONAL_SUPPORT_GATE530_SYNTHETIC_3PLUS1_ADAPTER_INHERITED"
	StatusFundamentalSymmetryAirlockDefined  = "CONDITIONAL_SUPPORT_FUNDAMENTAL_SYMMETRY_AIRLOCK_DEFINED"
	StatusKreinHilbertSchemaRowsEnumerated   = "CONDITIONAL_SUPPORT_KREIN_TO_HILBERT_SCHEMA_ROWS_ENUMERATED"
	StatusWickReflectionSchemaDefined        = "CONDITIONAL_SUPPORT_WICK_REFLECTION_POSITIVITY_SCHEMA_DEFINED"
	StatusProjectorCompatibilityGuardDefined = "CONDITIONAL_SUPPORT_PROJECTOR_COMPATIBILITY_GUARD_DEFINED"
	StatusMandatoryBridgeMetadataEnforced    = "CONDITIONAL_SUPPORT_MANDATORY_SOURCE_CONVENTION_BRIDGE_TAGS_ENFORCED"
	StatusRedactedFundamentalSchemaAccepted  = "CONDITIONAL_SUPPORT_REDACTED_FUNDAMENTAL_SYMMETRY_SCHEMA_ACCEPTED"
	StatusNativePromotionRejected            = "CONDITIONAL_SUPPORT_FUNDAMENTAL_SYMMETRY_NATIVE_PROMOTION_REJECTED"
	StatusNoObservedHilbertDataImported      = "CONDITIONAL_SUPPORT_NO_OBSERVED_HILBERT_WICK_OR_BOUNDARY_DATA_IMPORTED"

	StatusFailedThetaNativePromotionRejected = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_NATIVE_PROMOTION_REJECTED"
	StatusFailedMissingMetadataRejected      = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_MISSING_SOURCE_CONVENTION_BRIDGE_TAG_REJECTED"
	StatusFailedThetaDoesNotGrantHilbert     = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_POSITIVE_HILBERT_SPACE"
	StatusFailedThetaDoesNotGrantWick        = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedThetaDoesNotGrantReflection  = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_REFLECTION_POSITIVITY"
	StatusFailedThetaDoesNotGrantEnergy      = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_POSITIVE_ENERGY"
	StatusFailedThetaDoesNotGrantUnitary     = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedThetaDoesNotGrantGlobal      = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedComparatorNotPerformed       = "FAILED_ROUTE_FUNDAMENTAL_SYMMETRY_COMPARATOR_EXECUTION_NOT_PERFORMED_IN_PREFLIGHT"
	StatusFirewallPreserved                  = "FIREWALL_PRESERVED_GATE531_WICK_HILBERT_AIRLOCK_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked         = "FIREWALL_BLOCKED_GATE531_WICK_HILBERT_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate530AdapterExecuted            bool
	Gate530ProjectorResidualsZero     bool
	Gate530Rank44Confirmed            bool
	Gate530ExternalSignature13        bool
	Gate530WickBlocked                bool
	Gate530HilbertBlocked             bool
	Gate530UnitaryBlocked             bool
	Gate530InternalGaugeBlocked       bool
	Gate530NoObservedDimensionData    bool
	Gate530NativeWriteBlocked         bool
	Gate530ReopenedSealedFirewalls    bool
	Gate531FundamentalAirlockRedirect bool

	Verdict, Reason string
}

type SchemaField struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	NativeWrite bool
	Reason      string
}

type FundamentalSymmetrySchema struct {
	Executed bool

	RequiredRows                        []SchemaField
	RequiredRowCount                    int
	KreinMetricMatrixRequired           bool
	FundamentalSymmetryMatrixRequired   bool
	ThetaInvolutionCheckRequired        bool
	ThetaKreinSelfAdjointCheckRequired  bool
	PositiveHilbertFormCheckRequired    bool
	ProjectorCompatibilityCheckRequired bool
	TimeReflectionOperatorRequired      bool
	WickMapRequired                     bool
	IepsilonPrescriptionRequired        bool
	ReflectionPositivityProofRequired   bool
	PositiveEnergySpectrumRequired      bool
	GlobalHyperbolicityDataRequired     bool
	SourceRequired                      bool
	ConventionRequired                  bool
	BridgeOnlyRequired                  bool
	NoTheoremInputRequired              bool
	NativePromotionRejected             bool
	RedactedSchemaAccepted              bool
	AcceptedRedactedCases               int
	RejectedFailClosedCases             int

	Verdict, Reason string
}

type AlgebraicObligationGuard struct {
	Executed bool

	ComparatorExecutionPerformed     bool
	ThetaSquaredIdentityEvaluated    bool
	ThetaKreinSelfAdjointEvaluated   bool
	HilbertFormPositiveEvaluated     bool
	ProjectorCommutationEvaluated    bool
	TimeReflectionEvaluated          bool
	WickContinuationEvaluated        bool
	ReflectionPositivityEvaluated    bool
	PositiveEnergyEvaluated          bool
	UnitaryDynamicsEvaluated         bool
	GlobalHyperbolicityEvaluated     bool
	PositiveHilbertProductGranted    bool
	PhysicalStateSpaceSelected       bool
	WickRotationSelected             bool
	ReflectionPositivityProven       bool
	PositiveEnergyHamiltonianDerived bool
	UnitaryRealTimeDynamicsDerived   bool
	GlobalHyperbolicitySelected      bool

	Verdict, Reason string
}

type NativeRejection struct {
	Executed bool

	NativeFundamentalSymmetryWrite bool
	NativeHilbertProductWrite      bool
	NativePhysicalStateSpaceWrite  bool
	NativeTimeReflectionWrite      bool
	NativeWickWrite                bool
	NativeReflectionWrite          bool
	NativePositiveEnergyWrite      bool
	NativeUnitaryDynamicsWrite     bool
	NativeGlobalCausalWrite        bool
	NativeProjectorUpgradeWrite    bool
	ComparatorExecutionPerformed   bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedHilbertDataImported     bool
	ObservedWickDataImported        bool
	ObservedBoundaryDataImported    bool
	ObservedHamiltonianDataImported bool
	NativeFundamentalSymmetryWrite  bool
	NativeHilbertProductWrite       bool
	NativePhysicalStateSpaceWrite   bool
	NativeWickWrite                 bool
	NativeReflectionWrite           bool
	NativePositiveEnergyWrite       bool
	NativeUnitaryDynamicsWrite      bool
	NativeGlobalCausalWrite         bool
	Native3Plus1UpgradeWrite        bool
	ReopenedFlavorFirewall          bool
	ReopenedEWScaleFirewall         bool
	ReopenedGravityFirewall         bool
	ReopenedTopologyFirewall        bool
	NativeRegistryWritten           bool

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
	Schema      FundamentalSymmetrySchema
	Guard       AlgebraicObligationGuard
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
	g530, err := generation2synthetic3plus1projectionadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate530 synthetic projection adapter: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g530)
	a.Schema = buildSchema(a.Inheritance)
	a.Guard = buildGuard(a.Schema)
	a.Rejection = buildNativeRejection(a.Schema, a.Guard)
	a.Firewall = buildFirewall(a.Inheritance, a.Schema, a.Guard, a.Rejection)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g530 generation2synthetic3plus1projectionadapter.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate530AdapterExecuted:            g530.Output.Executed && g530.Output.Ready && g530.Output.CliffordCompatible,
		Gate530ProjectorResidualsZero:     g530.Output.AllResidualsZero,
		Gate530Rank44Confirmed:            g530.Output.ProjectorRankValid && g530.Output.ComplementRankValid && g530.Output.ProjectorRank == 4 && g530.Output.ComplementRank == 4,
		Gate530ExternalSignature13:        g530.Output.ExternalSignatureOK && g530.Output.ExternalPositive == 1 && g530.Output.ExternalNegative == 3,
		Gate530WickBlocked:                !g530.Firewall.WickRotationGranted,
		Gate530HilbertBlocked:             !g530.Firewall.PositiveHilbertGranted,
		Gate530UnitaryBlocked:             !g530.Firewall.UnitaryRealTimeGranted,
		Gate530InternalGaugeBlocked:       !g530.Firewall.InternalGaugeNativeIdentification,
		Gate530NoObservedDimensionData:    !g530.Firewall.ObservedDimensionImported,
		Gate530NativeWriteBlocked:         !g530.Firewall.NativeRegistryWritten,
		Gate530ReopenedSealedFirewalls:    g530.Firewall.ReopenedFlavorFirewall || g530.Firewall.ReopenedEWScaleFirewall || g530.Firewall.ReopenedGravityFirewall || g530.Firewall.ReopenedTopologyFirewall,
		Gate531FundamentalAirlockRedirect: g530.Next.Gate == 531,
		Verdict:                           StatusGate530AdapterInherited,
		Reason:                            "Gate531 inherits Gate530's result: the synthetic 3+1 projector adapter closes dimensional socket residuals, but Wick rotation, Hilbert positivity, reflection positivity, positive energy, unitary dynamics, and global hyperbolicity remain untouched bridge obligations.",
	}
}

func buildSchema(in Inheritance) FundamentalSymmetrySchema {
	rows := []SchemaField{
		{Name: "krein_metric_matrix", Required: true, BridgeOnly: true, Reason: "defines the indefinite form against which the adjoint and fundamental symmetry are audited"},
		{Name: "fundamental_symmetry_matrix", Required: true, BridgeOnly: true, Reason: "candidate Θ or J_F with Θ²=I used only to test a positive form, not to declare a physical Hilbert space"},
		{Name: "projector_reference", Required: true, BridgeOnly: true, Reason: "binds the Hilbert/Wick candidate to an already accepted 3+1 bridge projector row"},
		{Name: "time_reflection_operator", Required: true, BridgeOnly: true, Reason: "required before any Osterwalder-Schrader or reflection-positivity comparator can be meaningful"},
		{Name: "wick_map_convention", Required: true, BridgeOnly: true, Reason: "records the Euclidean-to-Lorentzian continuation convention without selecting it natively"},
		{Name: "i_epsilon_prescription", Required: true, BridgeOnly: true, Reason: "separates analytic-continuation boundary conditions from finite Clifford algebra"},
		{Name: "reflection_positivity_certificate", Required: true, BridgeOnly: true, Reason: "must be supplied or proven before positive Hilbert reconstruction can be claimed"},
		{Name: "positive_energy_certificate", Required: true, BridgeOnly: true, Reason: "prevents a positive metric candidate from being smuggled into a Hamiltonian spectrum theorem"},
		{Name: "global_causal_boundary_data", Required: true, BridgeOnly: true, Reason: "global hyperbolicity and causal boundary data are separate from local matrix positivity"},
		{Name: "source", Required: true, BridgeOnly: true, Reason: "airlock rows must be source-tagged"},
		{Name: "source_version", Required: true, BridgeOnly: true, Reason: "airlock rows must be reproducibly versioned"},
		{Name: "convention", Required: true, BridgeOnly: true, Reason: "signature, adjoint, Wick, and reflection conventions must be explicit"},
		{Name: "bridge_only", Required: true, BridgeOnly: true, Reason: "all rows remain bridge data"},
		{Name: "no_theorem_input", Required: true, BridgeOnly: true, Reason: "preflight schemas and future fixtures cannot be native theorem inputs"},
		{Name: "native_promotion", Required: true, BridgeOnly: true, Reason: "must be false; true is rejected fail-closed"},
	}
	return FundamentalSymmetrySchema{
		Executed:                            true,
		RequiredRows:                        rows,
		RequiredRowCount:                    len(rows),
		KreinMetricMatrixRequired:           true,
		FundamentalSymmetryMatrixRequired:   true,
		ThetaInvolutionCheckRequired:        true,
		ThetaKreinSelfAdjointCheckRequired:  true,
		PositiveHilbertFormCheckRequired:    true,
		ProjectorCompatibilityCheckRequired: in.Gate530AdapterExecuted && in.Gate530Rank44Confirmed,
		TimeReflectionOperatorRequired:      true,
		WickMapRequired:                     true,
		IepsilonPrescriptionRequired:        true,
		ReflectionPositivityProofRequired:   true,
		PositiveEnergySpectrumRequired:      true,
		GlobalHyperbolicityDataRequired:     true,
		SourceRequired:                      true,
		ConventionRequired:                  true,
		BridgeOnlyRequired:                  true,
		NoTheoremInputRequired:              true,
		NativePromotionRejected:             true,
		RedactedSchemaAccepted:              true,
		AcceptedRedactedCases:               1,
		RejectedFailClosedCases:             3,
		Verdict:                             strings.Join([]string{StatusFundamentalSymmetryAirlockDefined, StatusKreinHilbertSchemaRowsEnumerated, StatusWickReflectionSchemaDefined, StatusProjectorCompatibilityGuardDefined, StatusMandatoryBridgeMetadataEnforced, StatusRedactedFundamentalSchemaAccepted, StatusNativePromotionRejected}, ";"),
		Reason:                              "Gate531 defines the required source-tagged bridge schema for a future fundamental-symmetry/Wick/Hilbert ledger. The schema requires Θ²=I, Krein self-adjointness, positivity of GΘ, projector compatibility, time reflection, Wick convention, iε prescription, OS/reflection positivity, positive-energy, and global-causal certificates, while rejecting native promotion.",
	}
}

func buildGuard(s FundamentalSymmetrySchema) AlgebraicObligationGuard {
	return AlgebraicObligationGuard{
		Executed:                         true,
		ComparatorExecutionPerformed:     false,
		ThetaSquaredIdentityEvaluated:    false,
		ThetaKreinSelfAdjointEvaluated:   false,
		HilbertFormPositiveEvaluated:     false,
		ProjectorCommutationEvaluated:    false,
		TimeReflectionEvaluated:          false,
		WickContinuationEvaluated:        false,
		ReflectionPositivityEvaluated:    false,
		PositiveEnergyEvaluated:          false,
		UnitaryDynamicsEvaluated:         false,
		GlobalHyperbolicityEvaluated:     false,
		PositiveHilbertProductGranted:    false,
		PhysicalStateSpaceSelected:       false,
		WickRotationSelected:             false,
		ReflectionPositivityProven:       false,
		PositiveEnergyHamiltonianDerived: false,
		UnitaryRealTimeDynamicsDerived:   false,
		GlobalHyperbolicitySelected:      false,
		Verdict:                          strings.Join([]string{StatusFailedComparatorNotPerformed, StatusFailedThetaDoesNotGrantHilbert, StatusFailedThetaDoesNotGrantWick, StatusFailedThetaDoesNotGrantReflection, StatusFailedThetaDoesNotGrantEnergy, StatusFailedThetaDoesNotGrantUnitary, StatusFailedThetaDoesNotGrantGlobal}, ";"),
		Reason:                           fmt.Sprintf("Gate531 is a preflight schema only. It enumerates %d required rows but performs no Θ², G-self-adjointness, positivity, OS, Wick, Hamiltonian-spectrum, unitary, or global-causal comparator execution.", s.RequiredRowCount),
	}
}

func buildNativeRejection(s FundamentalSymmetrySchema, g AlgebraicObligationGuard) NativeRejection {
	return NativeRejection{
		Executed:                       true,
		NativeFundamentalSymmetryWrite: false,
		NativeHilbertProductWrite:      false,
		NativePhysicalStateSpaceWrite:  false,
		NativeTimeReflectionWrite:      false,
		NativeWickWrite:                false,
		NativeReflectionWrite:          false,
		NativePositiveEnergyWrite:      false,
		NativeUnitaryDynamicsWrite:     false,
		NativeGlobalCausalWrite:        false,
		NativeProjectorUpgradeWrite:    false,
		ComparatorExecutionPerformed:   g.ComparatorExecutionPerformed,
		Verdict:                        strings.Join([]string{StatusFailedThetaNativePromotionRejected, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                         "Any attempt to write a fundamental symmetry, Hilbert product, physical state space, Wick map, positive-energy Hamiltonian, unitary dynamics, or global causal structure into the native registry is rejected at Gate531.",
	}
}

func buildFirewall(in Inheritance, s FundamentalSymmetrySchema, g AlgebraicObligationGuard, r NativeRejection) Firewall {
	return Firewall{
		Executed:                        true,
		ObservedHilbertDataImported:     false,
		ObservedWickDataImported:        false,
		ObservedBoundaryDataImported:    false,
		ObservedHamiltonianDataImported: false,
		NativeFundamentalSymmetryWrite:  r.NativeFundamentalSymmetryWrite,
		NativeHilbertProductWrite:       r.NativeHilbertProductWrite || g.PositiveHilbertProductGranted,
		NativePhysicalStateSpaceWrite:   r.NativePhysicalStateSpaceWrite || g.PhysicalStateSpaceSelected,
		NativeWickWrite:                 r.NativeWickWrite || g.WickRotationSelected,
		NativeReflectionWrite:           r.NativeReflectionWrite || g.ReflectionPositivityProven,
		NativePositiveEnergyWrite:       r.NativePositiveEnergyWrite || g.PositiveEnergyHamiltonianDerived,
		NativeUnitaryDynamicsWrite:      r.NativeUnitaryDynamicsWrite || g.UnitaryRealTimeDynamicsDerived,
		NativeGlobalCausalWrite:         r.NativeGlobalCausalWrite || g.GlobalHyperbolicitySelected,
		Native3Plus1UpgradeWrite:        r.NativeProjectorUpgradeWrite,
		ReopenedFlavorFirewall:          in.Gate530ReopenedSealedFirewalls,
		ReopenedEWScaleFirewall:         in.Gate530ReopenedSealedFirewalls,
		ReopenedGravityFirewall:         in.Gate530ReopenedSealedFirewalls,
		ReopenedTopologyFirewall:        in.Gate530ReopenedSealedFirewalls,
		NativeRegistryWritten:           false,
		Verdict:                         strings.Join([]string{StatusNoObservedHilbertDataImported, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                          "Gate531 imports no observed Hilbert, Wick, Hamiltonian, or boundary data and performs no comparator execution. Completed flavor, electroweak-scale, gravity-normalization, and topology firewalls remain sealed.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Cℓ(1,7) still contributes only the native indefinite/Krein causal socket inherited from Gates 526–527",
			"Gate530's synthetic dimensional adapter remains bridge plumbing and is not upgraded into a physical state-space theorem",
			"No native fundamental symmetry, positive Hilbert product, Wick map, Hamiltonian, unitary dynamics, or global causal structure is written at Gate531",
		},
		BridgeEntries: []string{
			"Wick/Hilbert airlock schema requires a Krein metric, candidate fundamental symmetry Θ, projector reference, time reflection, Wick map, iε prescription, reflection-positivity certificate, positive-energy certificate, and global-causal boundary data",
			"future comparator rows must check Θ²=I, Θ†_G=Θ, positivity of GΘ, compatibility with the selected 3+1 projector, and OS/reflection positivity before any Hilbert reconstruction claim",
			"all candidate rows must be source-tagged, convention-tagged, bridge_only=true, no_theorem_input=true, and native_promotion=false",
		},
		EnvironmentalEntries: []string{
			"the actual physical time reflection, thermodynamic arrow, analytic-continuation prescription, Hamiltonian domain, and global causal boundary remain environmental or future bridge data",
			"a positive matrix test alone cannot select a physical universe without Wick/reflection, spectrum, and global-causal certificates",
		},
		FailedRoutes: []string{
			StatusFailedThetaNativePromotionRejected,
			StatusFailedMissingMetadataRejected,
			StatusFailedThetaDoesNotGrantHilbert,
			StatusFailedThetaDoesNotGrantWick,
			StatusFailedThetaDoesNotGrantReflection,
			StatusFailedThetaDoesNotGrantEnergy,
			StatusFailedThetaDoesNotGrantUnitary,
			StatusFailedThetaDoesNotGrantGlobal,
			StatusFailedComparatorNotPerformed,
		},
		OpenTheorems: []string{
			"execute a synthetic fundamental-symmetry fixture through the Gate531 schema and compute Θ², G-self-adjointness, GΘ positivity, and projector compatibility residuals",
			"define a separate reflection-positivity/Osterwalder-Schrader comparator rather than inferring it from finite matrix positivity",
			"audit whether positive-energy and unitary real-time dynamics can ever be derived from the finite algebra or must remain continuum/environmental inputs",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        532,
		Title:       "Synthetic Fundamental-Symmetry Ledger Adapter and Positivity Residual Dry Run",
		Reason:      "Gate531 defines the Wick/Hilbert airlock schema but deliberately performs no comparator execution. The next safe step is a synthetic fixture that tests only finite algebraic positivity residuals while keeping OS, Wick, positive-energy, and global-causal firewalls closed.",
		PrimaryTask: "Load a synthetic Θ ledger, compute Θ²=I, Krein self-adjointness, GΘ positive-definiteness, and compatibility with the Gate530 3+1 projector; report the result as bridge plumbing only.",
	}
}

func truth(a Analysis) string {
	return "Gate531 does not solve the Hilbert-space problem; it prevents it from being smuggled in. The gate defines a fail-closed airlock for a future fundamental-symmetry/Wick ledger and states the exact obligations: Θ²=I, Krein self-adjointness, positivity of GΘ, compatibility with the selected 3+1 projector, time reflection, Wick/iε convention, reflection positivity, positive energy, unitary dynamics, and global causal data. None of these are promoted to native ASHA law at this gate."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate530AdapterExecuted || !a.Inheritance.Gate530ProjectorResidualsZero || !a.Inheritance.Gate530Rank44Confirmed || !a.Inheritance.Gate530ExternalSignature13 || !a.Inheritance.Gate530WickBlocked || !a.Inheritance.Gate530HilbertBlocked || !a.Inheritance.Gate530UnitaryBlocked || !a.Inheritance.Gate530InternalGaugeBlocked || !a.Inheritance.Gate530NoObservedDimensionData || !a.Inheritance.Gate530NativeWriteBlocked || a.Inheritance.Gate530ReopenedSealedFirewalls || !a.Inheritance.Gate531FundamentalAirlockRedirect {
		bad = append(bad, "bad inheritance")
	}
	if !a.Schema.Executed || a.Schema.RequiredRowCount < 15 || !a.Schema.KreinMetricMatrixRequired || !a.Schema.FundamentalSymmetryMatrixRequired || !a.Schema.ThetaInvolutionCheckRequired || !a.Schema.ThetaKreinSelfAdjointCheckRequired || !a.Schema.PositiveHilbertFormCheckRequired || !a.Schema.ProjectorCompatibilityCheckRequired || !a.Schema.TimeReflectionOperatorRequired || !a.Schema.WickMapRequired || !a.Schema.IepsilonPrescriptionRequired || !a.Schema.ReflectionPositivityProofRequired || !a.Schema.PositiveEnergySpectrumRequired || !a.Schema.GlobalHyperbolicityDataRequired || !a.Schema.SourceRequired || !a.Schema.ConventionRequired || !a.Schema.BridgeOnlyRequired || !a.Schema.NoTheoremInputRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted || a.Schema.AcceptedRedactedCases != 1 || a.Schema.RejectedFailClosedCases < 3 {
		bad = append(bad, "bad fundamental symmetry schema")
	}
	if !a.Guard.Executed || a.Guard.ComparatorExecutionPerformed || a.Guard.ThetaSquaredIdentityEvaluated || a.Guard.ThetaKreinSelfAdjointEvaluated || a.Guard.HilbertFormPositiveEvaluated || a.Guard.ProjectorCommutationEvaluated || a.Guard.TimeReflectionEvaluated || a.Guard.WickContinuationEvaluated || a.Guard.ReflectionPositivityEvaluated || a.Guard.PositiveEnergyEvaluated || a.Guard.UnitaryDynamicsEvaluated || a.Guard.GlobalHyperbolicityEvaluated || a.Guard.PositiveHilbertProductGranted || a.Guard.PhysicalStateSpaceSelected || a.Guard.WickRotationSelected || a.Guard.ReflectionPositivityProven || a.Guard.PositiveEnergyHamiltonianDerived || a.Guard.UnitaryRealTimeDynamicsDerived || a.Guard.GlobalHyperbolicitySelected {
		bad = append(bad, "bad obligation guard")
	}
	if !a.Rejection.Executed || a.Rejection.NativeFundamentalSymmetryWrite || a.Rejection.NativeHilbertProductWrite || a.Rejection.NativePhysicalStateSpaceWrite || a.Rejection.NativeTimeReflectionWrite || a.Rejection.NativeWickWrite || a.Rejection.NativeReflectionWrite || a.Rejection.NativePositiveEnergyWrite || a.Rejection.NativeUnitaryDynamicsWrite || a.Rejection.NativeGlobalCausalWrite || a.Rejection.NativeProjectorUpgradeWrite || a.Rejection.ComparatorExecutionPerformed {
		bad = append(bad, "bad native rejection")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedHilbertDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedBoundaryDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.NativeFundamentalSymmetryWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.Native3Plus1UpgradeWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: adapter=%t residuals_zero=%t rank44=%t ext_signature_1plus3=%t Wick_blocked=%t Hilbert_blocked=%t unitary_blocked=%t internal_gauge_blocked=%t no_observed_dimension=%t native_blocked=%t reopens_firewalls=%t gate531_redirect=%t; %s", x.Verdict, x.Gate530AdapterExecuted, x.Gate530ProjectorResidualsZero, x.Gate530Rank44Confirmed, x.Gate530ExternalSignature13, x.Gate530WickBlocked, x.Gate530HilbertBlocked, x.Gate530UnitaryBlocked, x.Gate530InternalGaugeBlocked, x.Gate530NoObservedDimensionData, x.Gate530NativeWriteBlocked, x.Gate530ReopenedSealedFirewalls, x.Gate531FundamentalAirlockRedirect, x.Reason)
}

func FormatSchema(x FundamentalSymmetrySchema) string {
	return fmt.Sprintf("%s: rows=%d G_required=%t theta_required=%t theta2=%t Gselfadjoint=%t Gtheta_positive=%t projector_compat=%t time_reflection=%t Wick=%t i_epsilon=%t reflection_cert=%t positive_energy_cert=%t global_causal=%t source=%t convention=%t bridge_only=%t no_theorem=%t native_rejected=%t redacted=%t accepted=%d rejected=%d; %s", x.Verdict, x.RequiredRowCount, x.KreinMetricMatrixRequired, x.FundamentalSymmetryMatrixRequired, x.ThetaInvolutionCheckRequired, x.ThetaKreinSelfAdjointCheckRequired, x.PositiveHilbertFormCheckRequired, x.ProjectorCompatibilityCheckRequired, x.TimeReflectionOperatorRequired, x.WickMapRequired, x.IepsilonPrescriptionRequired, x.ReflectionPositivityProofRequired, x.PositiveEnergySpectrumRequired, x.GlobalHyperbolicityDataRequired, x.SourceRequired, x.ConventionRequired, x.BridgeOnlyRequired, x.NoTheoremInputRequired, x.NativePromotionRejected, x.RedactedSchemaAccepted, x.AcceptedRedactedCases, x.RejectedFailClosedCases, x.Reason)
}

func FormatGuard(x AlgebraicObligationGuard) string {
	return fmt.Sprintf("%s: comparator=%t theta2_eval=%t Gselfadjoint_eval=%t positivity_eval=%t projector_commutation_eval=%t time_reflection_eval=%t Wick_eval=%t reflection_eval=%t positive_energy_eval=%t unitary_eval=%t global_eval=%t Hilbert_granted=%t state_space=%t Wick_selected=%t reflection_proven=%t positive_energy=%t unitary=%t global=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.ThetaSquaredIdentityEvaluated, x.ThetaKreinSelfAdjointEvaluated, x.HilbertFormPositiveEvaluated, x.ProjectorCommutationEvaluated, x.TimeReflectionEvaluated, x.WickContinuationEvaluated, x.ReflectionPositivityEvaluated, x.PositiveEnergyEvaluated, x.UnitaryDynamicsEvaluated, x.GlobalHyperbolicityEvaluated, x.PositiveHilbertProductGranted, x.PhysicalStateSpaceSelected, x.WickRotationSelected, x.ReflectionPositivityProven, x.PositiveEnergyHamiltonianDerived, x.UnitaryRealTimeDynamicsDerived, x.GlobalHyperbolicitySelected, x.Reason)
}

func FormatRejection(x NativeRejection) string {
	return fmt.Sprintf("%s: native_theta=%t native_Hilbert=%t native_state=%t native_time_reflection=%t native_Wick=%t native_reflection=%t native_positive_energy=%t native_unitary=%t native_global=%t projector_upgrade=%t comparator=%t; %s", x.Verdict, x.NativeFundamentalSymmetryWrite, x.NativeHilbertProductWrite, x.NativePhysicalStateSpaceWrite, x.NativeTimeReflectionWrite, x.NativeWickWrite, x.NativeReflectionWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeProjectorUpgradeWrite, x.ComparatorExecutionPerformed, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_Hilbert=%t observed_Wick=%t observed_boundary=%t observed_Hamiltonian=%t native_theta=%t native_Hilbert=%t native_state=%t native_Wick=%t native_reflection=%t native_positive_energy=%t native_unitary=%t native_global=%t native_3plus1_upgrade=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_write=%t; %s", x.Verdict, x.ObservedHilbertDataImported, x.ObservedWickDataImported, x.ObservedBoundaryDataImported, x.ObservedHamiltonianDataImported, x.NativeFundamentalSymmetryWrite, x.NativeHilbertProductWrite, x.NativePhysicalStateSpaceWrite, x.NativeWickWrite, x.NativeReflectionWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.Native3Plus1UpgradeWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate530AdapterInherited,
		StatusFundamentalSymmetryAirlockDefined,
		StatusKreinHilbertSchemaRowsEnumerated,
		StatusWickReflectionSchemaDefined,
		StatusProjectorCompatibilityGuardDefined,
		StatusMandatoryBridgeMetadataEnforced,
		StatusRedactedFundamentalSchemaAccepted,
		StatusNativePromotionRejected,
		StatusNoObservedHilbertDataImported,
		StatusFailedThetaNativePromotionRejected,
		StatusFailedMissingMetadataRejected,
		StatusFailedThetaDoesNotGrantHilbert,
		StatusFailedThetaDoesNotGrantWick,
		StatusFailedThetaDoesNotGrantReflection,
		StatusFailedThetaDoesNotGrantEnergy,
		StatusFailedThetaDoesNotGrantUnitary,
		StatusFailedThetaDoesNotGrantGlobal,
		StatusFailedComparatorNotPerformed,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 531 Registry Audit — Wick/Hilbert Fundamental-Symmetry Airlock Preflight\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 531 inherits Gate 530's synthetic dimensional adapter and its still-closed Wick/Hilbert/unitary firewalls.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Fundamental-symmetry schema\n\nThe airlock defines what a future Krein-to-Hilbert ledger must provide before any positivity comparator can run.\n\n```text\n" + FormatSchema(a.Schema) + "\n```\n\n")
	b.WriteString("### Required schema rows\n\n")
	for _, row := range a.Schema.RequiredRows {
		b.WriteString(fmt.Sprintf("- `%s` — required=%t bridge_only=%t native_write=%t; %s\n", row.Name, row.Required, row.BridgeOnly, row.NativeWrite, row.Reason))
	}
	b.WriteString("\n## Algebraic obligation guard\n\nGate 531 is preflight only: no Θ, Wick, OS, Hamiltonian, unitary, or global-causal comparator is executed.\n\n```text\n" + FormatGuard(a.Guard) + "\n```\n\n")
	b.WriteString("## Native rejection rule\n\n```text\n" + FormatRejection(a.Rejection) + "\n```\n\n")
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
