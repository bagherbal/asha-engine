// Package generation2physicalschwingerledgerairlock implements Gate 536:
// Physical Schwinger-Function Source Ledger Airlock.
//
// Gate 535 closed the synthetic OS/Wick/Hilbert sector as a frontier map. This
// package opens the next honest bridge boundary: a fail-closed schema for
// importing sourced Euclidean Schwinger-function or constructive-QFT correlation
// data. It performs no kernel comparison, imports no observed/physical
// correlators, and grants no Wick rotation, Hamiltonian, unitarity, global
// causality, or time-arrow theorem.
package generation2physicalschwingerledgerairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2oswickhilbertsectorclosureledger"
)

const (
	AuditID = "GATE536-PHYSICAL-SCHWINGER-FUNCTION-SOURCE-LEDGER-AIRLOCK"

	StatusGate535SectorClosureInherited         = "CONDITIONAL_SUPPORT_GATE535_OS_WICK_HILBERT_SECTOR_CLOSURE_INHERITED"
	StatusSchwingerSourceLedgerAirlockDefined   = "CONDITIONAL_SUPPORT_PHYSICAL_SCHWINGER_SOURCE_LEDGER_AIRLOCK_DEFINED"
	StatusSchwingerSchemaRowsEnumerated         = "CONDITIONAL_SUPPORT_SCHWINGER_FUNCTION_SCHEMA_ROWS_ENUMERATED"
	StatusEuclideanCorrelationProvenanceReady   = "CONDITIONAL_SUPPORT_EUCLIDEAN_CORRELATION_PROVENANCE_SCHEMA_READY"
	StatusOSCompatibilityCertificateRequired    = "CONDITIONAL_SUPPORT_OS_COMPATIBILITY_CERTIFICATE_REQUIRED"
	StatusConstructiveQFTSourceTagsEnforced     = "CONDITIONAL_SUPPORT_CONSTRUCTIVE_QFT_SOURCE_TAGS_ENFORCED"
	StatusComparatorExecutionBlockedInPreflight = "CONDITIONAL_SUPPORT_SCHWINGER_COMPARATOR_EXECUTION_BLOCKED_IN_PREFLIGHT"
	StatusNoObservedCorrelatorsImported         = "CONDITIONAL_SUPPORT_NO_OBSERVED_OR_PHYSICAL_SCHWINGER_DATA_IMPORTED"
	StatusNativePromotionRejected               = "CONDITIONAL_SUPPORT_PHYSICAL_SCHWINGER_NATIVE_PROMOTION_REJECTED"

	StatusFailedSchemaNotSchwinger    = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSchemaNotOSPositivity = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedSchemaNotWick         = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSchemaNotHilbert      = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSchemaNotHamiltonian  = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSchemaNotUnitary      = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSchemaNotGlobal       = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSchemaNotArrow        = "FAILED_ROUTE_SCHWINGER_SOURCE_SCHEMA_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedObservedImportAbsent  = "FAILED_ROUTE_NO_PHYSICAL_SCHWINGER_LEDGER_IMPORTED_IN_GATE536_PREFLIGHT"
	StatusFirewallPreserved           = "FIREWALL_PRESERVED_GATE536_PHYSICAL_SCHWINGER_LEDGER_AIRLOCK_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked  = "FIREWALL_BLOCKED_GATE536_PHYSICAL_CORRELATION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate535ClosureLedgerEmitted     bool
	Gate535NativeFrontierFrozen     bool
	Gate535BridgeFrontierMapped     bool
	Gate535EnvironmentalMapped      bool
	Gate535OSBlockClosed            bool
	Gate535DynamicsMapped           bool
	Gate535SchwingerBlocked         bool
	Gate535WickBlocked              bool
	Gate535HilbertBlocked           bool
	Gate535HamiltonianBlocked       bool
	Gate535UnitaryBlocked           bool
	Gate535GlobalBlocked            bool
	Gate535ArrowBlocked             bool
	Gate535NoObservedDynamics       bool
	Gate535NativeWriteBlocked       bool
	Gate536SchwingerAirlockRedirect bool

	Verdict, Reason string
}

type SchemaRow struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	Comparator  bool
	NativeWrite bool
	Reason      string
}

type SourceLedgerSchema struct {
	Executed bool

	Rows                                    []SchemaRow
	RequiredRows                            int
	BridgeOnlyRows                          int
	ComparatorRows                          int
	NativeWriteRows                         int
	SourceIdentifierRequired                bool
	ConstructiveDefinitionRequired          bool
	EuclideanDomainRequired                 bool
	FieldAlgebraRequired                    bool
	TestFunctionDomainRequired              bool
	NPointFamilyRequired                    bool
	SymmetryConventionRequired              bool
	DistributionRegularityRequired          bool
	EuclideanCovarianceRequired             bool
	ReflectionOperatorRequired              bool
	OSQuadraticFormRequired                 bool
	ReflectionPositivityCertificateRequired bool
	NullQuotientRuleRequired                bool
	ReconstructionMapRequired               bool
	WickMapRequired                         bool
	IepsilonConventionRequired              bool
	HamiltonianSpectrumCertificateRequired  bool
	RenormalizationSchemeRequired           bool
	UncertaintyLedgerRequired               bool
	SourceAndLicenseRequired                bool
	NoTheoremInputRequired                  bool
	NativePromotionRejected                 bool
	RedactedSchemaAccepted                  bool

	Verdict, Reason string
}

type Guard struct {
	Executed bool

	ComparatorExecutionPerformed     bool
	PhysicalSchwingerFunctionsLoaded bool
	ObservedCorrelationDataImported  bool
	ConstructiveMeasureImported      bool
	NPointDistributionsEvaluated     bool
	ReflectionPositivityEvaluated    bool
	OSNullQuotientComputed           bool
	WickContinuationEvaluated        bool
	HilbertReconstructionEvaluated   bool
	HamiltonianSpectrumEvaluated     bool
	UnitaryDynamicsEvaluated         bool
	GlobalHyperbolicityEvaluated     bool
	TimeArrowEvaluated               bool
	PhysicalSchwingerDerived         bool
	OSPositivityProven               bool
	WickRotationSelected             bool
	PhysicalHilbertSpaceSelected     bool
	PositiveEnergyHamiltonianDerived bool
	UnitaryDynamicsDerived           bool
	GlobalCausalStructureSelected    bool
	ArrowOfTimeSelected              bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ObservedSchwingerDataImported   bool
	ObservedWickDataImported        bool
	ObservedHamiltonianDataImported bool
	ObservedCausalBoundaryImported  bool
	NativeSchwingerWrite            bool
	NativeEuclideanMeasureWrite     bool
	NativeOSPositivityWrite         bool
	NativeWickWrite                 bool
	NativeHilbertWrite              bool
	NativeHamiltonianWrite          bool
	NativeUnitaryWrite              bool
	NativeGlobalCausalWrite         bool
	NativeTimeArrowWrite            bool
	ReopenedFlavorFirewall          bool
	ReopenedEWScaleFirewall         bool
	ReopenedGravityScaleFirewall    bool
	ReopenedTopologyFirewall        bool
	ReopenedDimensionalFirewall     bool
	ReopenedKreinHilbertFirewall    bool
	NativeRegistryWritten           bool

	Verdict, Reason string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      SourceLedgerSchema
	Guard       Guard
	Firewall    Firewall
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
	g535, err := generation2oswickhilbertsectorclosureledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate535 sector closure ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g535)
	a.Schema = buildSchema(a.Inheritance)
	a.Guard = buildGuard(a.Schema)
	a.Firewall = buildFirewall(a.Inheritance, a.Schema, a.Guard)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2oswickhilbertsectorclosureledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate535ClosureLedgerEmitted:     g.Ledger.Executed && g.Ledger.FrontierConsistent,
		Gate535NativeFrontierFrozen:     strings.Contains(g.Ledger.Verdict, generation2oswickhilbertsectorclosureledger.StatusNativeFrontierFrozen),
		Gate535BridgeFrontierMapped:     strings.Contains(g.Ledger.Verdict, generation2oswickhilbertsectorclosureledger.StatusBridgeCompatibilityFrontierMapped),
		Gate535EnvironmentalMapped:      strings.Contains(g.Ledger.Verdict, generation2oswickhilbertsectorclosureledger.StatusEnvironmentalFrontierMapped),
		Gate535OSBlockClosed:            g.Ledger.OSRowsClosed,
		Gate535DynamicsMapped:           g.Ledger.DynamicsRowsMapped,
		Gate535SchwingerBlocked:         strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotSchwinger),
		Gate535WickBlocked:              strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotWick),
		Gate535HilbertBlocked:           strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotHilbert),
		Gate535HamiltonianBlocked:       strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotHamiltonian),
		Gate535UnitaryBlocked:           strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotUnitary),
		Gate535GlobalBlocked:            strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotGlobal),
		Gate535ArrowBlocked:             strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFailedClosureNotArrow),
		Gate535NoObservedDynamics:       strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusNoObservedDynamicsImported),
		Gate535NativeWriteBlocked:       !g.Firewall.NativeSchwingerWrite && !g.Firewall.NativeHamiltonianWrite && strings.Contains(g.Firewall.Verdict, generation2oswickhilbertsectorclosureledger.StatusFirewallNativeWriteBlocked),
		Gate536SchwingerAirlockRedirect: true,
		Verdict:                         StatusGate535SectorClosureInherited,
		Reason:                          "Gate536 inherits the closed OS/Wick/Hilbert frontier and opens only a schema airlock for sourced physical Schwinger data.",
	}
}

func buildSchema(in Inheritance) SourceLedgerSchema {
	rows := []SchemaRow{
		{Name: "source_identifier_and_version", Required: true, BridgeOnly: true, Comparator: false, Reason: "Every physical correlation row must be attributable to a source, construction, publication, dataset, or explicit synthetic fixture."},
		{Name: "constructive_definition_or_measure", Required: true, BridgeOnly: true, Comparator: false, Reason: "A Schwinger family cannot be guessed from a finite matrix; its Euclidean measure or constructive definition must be supplied."},
		{Name: "euclidean_domain_and_dimension", Required: true, BridgeOnly: true, Comparator: false, Reason: "The Euclidean coordinate domain, boundary condition, and dimension are external bridge conventions."},
		{Name: "field_algebra_and_labels", Required: true, BridgeOnly: true, Comparator: false, Reason: "The fields whose correlations are being imported must be typed before any OS check can act."},
		{Name: "test_function_domain", Required: true, BridgeOnly: true, Comparator: true, Reason: "Reflection positivity is only meaningful on a declared positive-time test-function domain."},
		{Name: "n_point_schwinger_family", Required: true, BridgeOnly: true, Comparator: true, Reason: "The comparator needs an explicit family S_n or finite Gram reduction, not a symbolic claim."},
		{Name: "symmetry_and_permutation_convention", Required: true, BridgeOnly: true, Comparator: true, Reason: "Bosonic/fermionic ordering and Euclidean symmetry conventions must be explicit."},
		{Name: "distribution_regularization_class", Required: true, BridgeOnly: true, Comparator: true, Reason: "Schwinger functions are generally distributions; regularity and regulator metadata must be auditable."},
		{Name: "euclidean_covariance_certificate", Required: true, BridgeOnly: true, Comparator: true, Reason: "OS reconstruction requires the Euclidean covariance lane to be source-tagged, not presumed."},
		{Name: "reflection_operator_theta_E", Required: true, BridgeOnly: true, Comparator: true, Reason: "The Euclidean time reflection used by the quadratic form must be declared."},
		{Name: "os_quadratic_form_definition", Required: true, BridgeOnly: true, Comparator: true, Reason: "The exact sesquilinear/quadratic convention must be known before testing positivity."},
		{Name: "reflection_positivity_certificate", Required: true, BridgeOnly: true, Comparator: true, Reason: "A certificate, proof, or finite comparator result is required before any OS claim can pass."},
		{Name: "null_space_quotient_rule", Required: true, BridgeOnly: true, Comparator: true, Reason: "Zero-norm states must be quotiented by an explicit rule before Hilbert reconstruction."},
		{Name: "reconstruction_map_certificate", Required: true, BridgeOnly: true, Comparator: true, Reason: "Hilbert reconstruction is downstream of OS data and cannot be inferred from schema presence."},
		{Name: "wick_map_and_i_epsilon_convention", Required: true, BridgeOnly: true, Comparator: false, Reason: "Analytic continuation and contour choices remain separate bridge conventions."},
		{Name: "hamiltonian_spectrum_certificate", Required: true, BridgeOnly: true, Comparator: false, Reason: "Positive energy requires a separate spectral/domain certificate."},
		{Name: "renormalization_scheme_and_scale", Required: true, BridgeOnly: true, Comparator: false, Reason: "Physical correlators require normalization and scheme metadata."},
		{Name: "uncertainty_and_validity_domain", Required: true, BridgeOnly: true, Comparator: false, Reason: "Approximation order, lattice spacing, cutoff, or theorem domain must be bounded."},
		{Name: "bridge_only_no_theorem_input_tags", Required: true, BridgeOnly: true, Comparator: false, Reason: "Imported physical rows must not be treated as native ASHA derivations."},
	}
	nativeRows := 0
	bridgeRows := 0
	comparatorRows := 0
	for _, row := range rows {
		if row.NativeWrite {
			nativeRows++
		}
		if row.BridgeOnly {
			bridgeRows++
		}
		if row.Comparator {
			comparatorRows++
		}
	}
	return SourceLedgerSchema{
		Executed:                                true,
		Rows:                                    rows,
		RequiredRows:                            len(rows),
		BridgeOnlyRows:                          bridgeRows,
		ComparatorRows:                          comparatorRows,
		NativeWriteRows:                         nativeRows,
		SourceIdentifierRequired:                true,
		ConstructiveDefinitionRequired:          true,
		EuclideanDomainRequired:                 true,
		FieldAlgebraRequired:                    true,
		TestFunctionDomainRequired:              true,
		NPointFamilyRequired:                    true,
		SymmetryConventionRequired:              true,
		DistributionRegularityRequired:          true,
		EuclideanCovarianceRequired:             true,
		ReflectionOperatorRequired:              true,
		OSQuadraticFormRequired:                 true,
		ReflectionPositivityCertificateRequired: true,
		NullQuotientRuleRequired:                true,
		ReconstructionMapRequired:               true,
		WickMapRequired:                         true,
		IepsilonConventionRequired:              true,
		HamiltonianSpectrumCertificateRequired:  true,
		RenormalizationSchemeRequired:           true,
		UncertaintyLedgerRequired:               true,
		SourceAndLicenseRequired:                true,
		NoTheoremInputRequired:                  true,
		NativePromotionRejected:                 true,
		RedactedSchemaAccepted:                  in.Gate535ClosureLedgerEmitted,
		Verdict:                                 strings.Join([]string{StatusSchwingerSourceLedgerAirlockDefined, StatusSchwingerSchemaRowsEnumerated, StatusEuclideanCorrelationProvenanceReady, StatusOSCompatibilityCertificateRequired, StatusConstructiveQFTSourceTagsEnforced, StatusNativePromotionRejected}, ";"),
		Reason:                                  "Gate536 enumerates the minimum source-tagged fields required before a physical Schwinger-function family can replace the synthetic OS kernel.",
	}
}

func buildGuard(schema SourceLedgerSchema) Guard {
	return Guard{
		Executed: true,
		Verdict:  strings.Join([]string{StatusComparatorExecutionBlockedInPreflight, StatusFailedSchemaNotSchwinger, StatusFailedSchemaNotOSPositivity, StatusFailedSchemaNotWick, StatusFailedSchemaNotHilbert, StatusFailedSchemaNotHamiltonian, StatusFailedSchemaNotUnitary, StatusFailedSchemaNotGlobal, StatusFailedSchemaNotArrow, StatusFailedObservedImportAbsent}, ";"),
		Reason:   "Gate536 is a schema preflight only; no physical correlator, OS comparator, Wick map, Hilbert reconstruction, Hamiltonian spectrum, or causal/time-orientation test is executed.",
	}
}

func buildFirewall(in Inheritance, schema SourceLedgerSchema, guard Guard) Firewall {
	return Firewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusNoObservedCorrelatorsImported, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSchemaNotSchwinger, StatusFailedSchemaNotOSPositivity, StatusFailedSchemaNotWick, StatusFailedSchemaNotHilbert, StatusFailedSchemaNotHamiltonian, StatusFailedSchemaNotUnitary, StatusFailedSchemaNotGlobal, StatusFailedSchemaNotArrow}, ";"),
		Reason:   "The source-ledger airlock imports no physical Schwinger data and writes no native Euclidean measure, OS positivity, Wick, Hilbert, Hamiltonian, unitary, global-causal, or time-arrow theorem.",
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 537, Title: "Synthetic Schwinger-Function Source Ledger Adapter Dry Run", Reason: "Gate536 defines the physical Schwinger source schema. The safe next step is a synthetic file-backed dry run that verifies the schema and finite comparator plumbing without importing observed or constructive physical correlators.", PrimaryTask: "Load a deliberately synthetic Schwinger family ledger, check schema completeness and finite OS-compatibility reductions, and preserve every Wick/Hamiltonian/native firewall."}
}

func truth(a Analysis) string {
	return "Gate536 opens the physical Schwinger-function source ledger only as an airlock: ASHA now knows exactly what sourced Euclidean correlation data would be required, but no physical Schwinger functions, OS positivity proof, Wick map, Hilbert space, Hamiltonian, unitary dynamics, global causality, or arrow of time is derived natively."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate535ClosureLedgerEmitted || !a.Inheritance.Gate535NativeFrontierFrozen || !a.Inheritance.Gate535BridgeFrontierMapped || !a.Inheritance.Gate535EnvironmentalMapped || !a.Inheritance.Gate535OSBlockClosed || !a.Inheritance.Gate535DynamicsMapped || !a.Inheritance.Gate535SchwingerBlocked || !a.Inheritance.Gate535WickBlocked || !a.Inheritance.Gate535HilbertBlocked || !a.Inheritance.Gate535HamiltonianBlocked || !a.Inheritance.Gate535UnitaryBlocked || !a.Inheritance.Gate535GlobalBlocked || !a.Inheritance.Gate535ArrowBlocked || !a.Inheritance.Gate535NoObservedDynamics || !a.Inheritance.Gate535NativeWriteBlocked || !a.Inheritance.Gate536SchwingerAirlockRedirect {
		bad = append(bad, "bad Gate535 inheritance")
	}
	if !a.Schema.Executed || len(a.Schema.Rows) != 19 || a.Schema.RequiredRows != 19 || a.Schema.BridgeOnlyRows != 19 || a.Schema.NativeWriteRows != 0 || a.Schema.ComparatorRows < 9 || !a.Schema.SourceIdentifierRequired || !a.Schema.ConstructiveDefinitionRequired || !a.Schema.EuclideanDomainRequired || !a.Schema.FieldAlgebraRequired || !a.Schema.TestFunctionDomainRequired || !a.Schema.NPointFamilyRequired || !a.Schema.SymmetryConventionRequired || !a.Schema.DistributionRegularityRequired || !a.Schema.EuclideanCovarianceRequired || !a.Schema.ReflectionOperatorRequired || !a.Schema.OSQuadraticFormRequired || !a.Schema.ReflectionPositivityCertificateRequired || !a.Schema.NullQuotientRuleRequired || !a.Schema.ReconstructionMapRequired || !a.Schema.WickMapRequired || !a.Schema.IepsilonConventionRequired || !a.Schema.HamiltonianSpectrumCertificateRequired || !a.Schema.RenormalizationSchemeRequired || !a.Schema.UncertaintyLedgerRequired || !a.Schema.SourceAndLicenseRequired || !a.Schema.NoTheoremInputRequired || !a.Schema.NativePromotionRejected || !a.Schema.RedactedSchemaAccepted {
		bad = append(bad, "bad Schwinger source schema")
	}
	if !a.Guard.Executed || a.Guard.ComparatorExecutionPerformed || a.Guard.PhysicalSchwingerFunctionsLoaded || a.Guard.ObservedCorrelationDataImported || a.Guard.ConstructiveMeasureImported || a.Guard.NPointDistributionsEvaluated || a.Guard.ReflectionPositivityEvaluated || a.Guard.OSNullQuotientComputed || a.Guard.WickContinuationEvaluated || a.Guard.HilbertReconstructionEvaluated || a.Guard.HamiltonianSpectrumEvaluated || a.Guard.UnitaryDynamicsEvaluated || a.Guard.GlobalHyperbolicityEvaluated || a.Guard.TimeArrowEvaluated || a.Guard.PhysicalSchwingerDerived || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSpaceSelected || a.Guard.PositiveEnergyHamiltonianDerived || a.Guard.UnitaryDynamicsDerived || a.Guard.GlobalCausalStructureSelected || a.Guard.ArrowOfTimeSelected {
		bad = append(bad, "bad preflight guard")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedSchwingerDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.ObservedCausalBoundaryImported || a.Firewall.NativeSchwingerWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate536 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: closure=%t native_frontier=%t bridge=%t environmental=%t OS_closed=%t dynamics=%t Schwinger_blocked=%t Wick_blocked=%t Hilbert_blocked=%t Hamiltonian_blocked=%t unitary_blocked=%t global_blocked=%t arrow_blocked=%t no_observed=%t native_blocked=%t redirect=%t; %s", x.Verdict, x.Gate535ClosureLedgerEmitted, x.Gate535NativeFrontierFrozen, x.Gate535BridgeFrontierMapped, x.Gate535EnvironmentalMapped, x.Gate535OSBlockClosed, x.Gate535DynamicsMapped, x.Gate535SchwingerBlocked, x.Gate535WickBlocked, x.Gate535HilbertBlocked, x.Gate535HamiltonianBlocked, x.Gate535UnitaryBlocked, x.Gate535GlobalBlocked, x.Gate535ArrowBlocked, x.Gate535NoObservedDynamics, x.Gate535NativeWriteBlocked, x.Gate536SchwingerAirlockRedirect, x.Reason)
}

func FormatSchema(x SourceLedgerSchema) string {
	names := []string{}
	for _, row := range x.Rows {
		names = append(names, row.Name)
	}
	return fmt.Sprintf("%s: rows=%d required=%d bridge_only=%d comparator=%d native_write=%d source=%t measure=%t domain=%t field_algebra=%t test_domain=%t n_point=%t symmetry=%t distribution=%t covariance=%t theta_E=%t os_form=%t os_cert=%t quotient=%t reconstruction=%t wick=%t iepsilon=%t hamiltonian=%t scheme=%t uncertainty=%t source_license=%t no_theorem_input=%t native_rejected=%t redacted=%t names=[%s]; %s", x.Verdict, len(x.Rows), x.RequiredRows, x.BridgeOnlyRows, x.ComparatorRows, x.NativeWriteRows, x.SourceIdentifierRequired, x.ConstructiveDefinitionRequired, x.EuclideanDomainRequired, x.FieldAlgebraRequired, x.TestFunctionDomainRequired, x.NPointFamilyRequired, x.SymmetryConventionRequired, x.DistributionRegularityRequired, x.EuclideanCovarianceRequired, x.ReflectionOperatorRequired, x.OSQuadraticFormRequired, x.ReflectionPositivityCertificateRequired, x.NullQuotientRuleRequired, x.ReconstructionMapRequired, x.WickMapRequired, x.IepsilonConventionRequired, x.HamiltonianSpectrumCertificateRequired, x.RenormalizationSchemeRequired, x.UncertaintyLedgerRequired, x.SourceAndLicenseRequired, x.NoTheoremInputRequired, x.NativePromotionRejected, x.RedactedSchemaAccepted, strings.Join(names, ","), x.Reason)
}

func FormatGuard(x Guard) string {
	return fmt.Sprintf("%s: comparator=%t physical_loaded=%t observed_corr=%t constructive_measure=%t npoint_eval=%t OS_eval=%t quotient=%t Wick_eval=%t Hilbert_eval=%t Hamiltonian_eval=%t unitary_eval=%t global_eval=%t arrow_eval=%t Schwinger_derived=%t OS_proven=%t Wick_selected=%t Hilbert_selected=%t Hamiltonian_derived=%t unitary_derived=%t global_selected=%t arrow_selected=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.PhysicalSchwingerFunctionsLoaded, x.ObservedCorrelationDataImported, x.ConstructiveMeasureImported, x.NPointDistributionsEvaluated, x.ReflectionPositivityEvaluated, x.OSNullQuotientComputed, x.WickContinuationEvaluated, x.HilbertReconstructionEvaluated, x.HamiltonianSpectrumEvaluated, x.UnitaryDynamicsEvaluated, x.GlobalHyperbolicityEvaluated, x.TimeArrowEvaluated, x.PhysicalSchwingerDerived, x.OSPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSpaceSelected, x.PositiveEnergyHamiltonianDerived, x.UnitaryDynamicsDerived, x.GlobalCausalStructureSelected, x.ArrowOfTimeSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_Schwinger=%t observed_Wick=%t observed_Hamiltonian=%t observed_causal=%t native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopened_flavor=%t reopened_EW=%t reopened_gravity=%t reopened_topology=%t reopened_dimension=%t reopened_Krein=%t native_registry=%t; %s", x.Verdict, x.ObservedSchwingerDataImported, x.ObservedWickDataImported, x.ObservedHamiltonianDataImported, x.ObservedCausalBoundaryImported, x.NativeSchwingerWrite, x.NativeEuclideanMeasureWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate535SectorClosureInherited, StatusSchwingerSourceLedgerAirlockDefined, StatusSchwingerSchemaRowsEnumerated, StatusEuclideanCorrelationProvenanceReady, StatusOSCompatibilityCertificateRequired, StatusConstructiveQFTSourceTagsEnforced, StatusComparatorExecutionBlockedInPreflight, StatusNoObservedCorrelatorsImported, StatusNativePromotionRejected, StatusFailedSchemaNotSchwinger, StatusFailedSchemaNotOSPositivity, StatusFailedSchemaNotWick, StatusFailedSchemaNotHilbert, StatusFailedSchemaNotHamiltonian, StatusFailedSchemaNotUnitary, StatusFailedSchemaNotGlobal, StatusFailedSchemaNotArrow, StatusFailedObservedImportAbsent, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
