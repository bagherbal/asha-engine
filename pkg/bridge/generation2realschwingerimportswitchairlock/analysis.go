// Package generation2realschwingerimportswitchairlock implements Gate 540:
// Real Schwinger Source Import Switch Preflight.
//
// Gate 539 proved that ASHA can parse a complete synthetic source-authenticity
// ledger, verify its checksum, and reject the fixture as physical evidence.
// Gate 540 defines the explicit fail-closed switch that must be turned on before
// any non-synthetic Schwinger source/authenticity comparator can execute. The
// switch is off by default and no real or constructive correlator is imported.
package generation2realschwingerimportswitchairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticsourceauthenticityadapter"
)

const (
	AuditID = "GATE540-REAL-SCHWINGER-SOURCE-IMPORT-SWITCH-PREFLIGHT"

	StatusGate539SyntheticAuthenticityInherited = "CONDITIONAL_SUPPORT_GATE539_SYNTHETIC_AUTHENTICITY_ADAPTER_INHERITED"
	StatusRealImportSwitchAirlockDefined        = "CONDITIONAL_SUPPORT_REAL_SCHWINGER_SOURCE_IMPORT_SWITCH_AIRLOCK_DEFINED"
	StatusRealImportSwitchSchemaEnumerated      = "CONDITIONAL_SUPPORT_REAL_IMPORT_SWITCH_SCHEMA_ROWS_ENUMERATED"
	StatusRealImportSwitchDefaultOff            = "CONDITIONAL_SUPPORT_REAL_IMPORT_SWITCH_DEFAULT_OFF"
	StatusExplicitOperatorIntentRequired        = "CONDITIONAL_SUPPORT_EXPLICIT_OPERATOR_INTENT_REQUIRED_FOR_REAL_SOURCE_IMPORT"
	StatusComparatorAuthorizationBlocked        = "CONDITIONAL_SUPPORT_REAL_SOURCE_COMPARATOR_AUTHORIZATION_BLOCKED_BY_DEFAULT"
	StatusNoRealSourceImported                  = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE540"
	StatusNativePromotionRejected               = "CONDITIONAL_SUPPORT_REAL_SOURCE_IMPORT_SWITCH_NATIVE_PROMOTION_REJECTED"

	StatusFailedSwitchSchemaNotSchwinger   = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSwitchSchemaNotOSProof     = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedSwitchSchemaNotWick        = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSwitchSchemaNotHilbert     = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSwitchSchemaNotHamiltonian = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSwitchSchemaNotUnitary     = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSwitchSchemaNotGlobal      = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSwitchSchemaNotArrow       = "FAILED_ROUTE_REAL_IMPORT_SWITCH_SCHEMA_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedImportSwitchOff            = "FAILED_ROUTE_REAL_SCHWINGER_SOURCE_IMPORT_SWITCH_OFF_BY_DEFAULT"
	StatusFailedNoExplicitSourceIntent     = "FAILED_ROUTE_NO_EXPLICIT_OPERATOR_INTENT_FOR_REAL_SOURCE_IMPORT_IN_GATE540"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_GATE540_REAL_SOURCE_IMPORT_SWITCH_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked       = "FIREWALL_BLOCKED_GATE540_REAL_SOURCE_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate539AdapterExecuted       bool
	Gate539RowsAccepted          bool
	Gate539ChecksumVerified      bool
	Gate539MetadataSieveEnforced bool
	Gate539FixtureRejected       bool
	Gate539NoRealSourceImported  bool
	Gate539NativeWriteBlocked    bool
	Gate540SwitchRedirect        bool

	Verdict, Reason string
}

type SwitchRow struct {
	Name        string
	Required    bool
	DefaultOff  bool
	BridgeOnly  bool
	Comparator  bool
	NativeWrite bool
	Reason      string
}

type SwitchSchema struct {
	Executed bool

	Rows                  []SwitchRow
	RequiredRows          int
	DefaultOffRows        int
	BridgeOnlyRows        int
	ComparatorRows        int
	NativeWriteRows       int
	OperatorIntentRow     bool
	SourceURIRow          bool
	AuthenticityLedgerRow bool
	ChecksumProofRow      bool
	LicenseAccessRow      bool
	Gate536AlignmentRow   bool
	QuarantineRow         bool
	NativeWriteLockRow    bool

	Verdict, Reason string
}

type ImportSwitch struct {
	Executed bool

	DefaultOff                          bool
	RealSourceImportEnabled             bool
	ExplicitOperatorIntentProvided      bool
	NonSyntheticSourceURIProvided       bool
	AuthenticityLedgerReferenceProvided bool
	ChecksumOrProofHashProvided         bool
	LicenseAndAccessGrantProvided       bool
	ComparatorPlanAuthorized            bool
	DryRunOnly                          bool
	QuarantineTargetDefined             bool
	NativeWriteAuthorized               bool
	NativeRegistryWriteRequested        bool
	ComparatorExecutionAllowed          bool
	NonSyntheticSourceLoaded            bool
	ObservedCorrelationLoaded           bool
	ConstructiveMeasureLoaded           bool
	PhysicalOSCertificateLoaded         bool
	PhysicalWickMapLoaded               bool
	PhysicalHamiltonianLoaded           bool

	Verdict, Reason string
}

type Guard struct {
	Executed bool

	ComparatorExecutionPerformed bool
	PhysicalSchwingerImported    bool
	ConstructiveMeasureImported  bool
	ObservedCorrelationImported  bool
	PhysicalSchwingerDerived     bool
	OSPositivityProven           bool
	WickRotationSelected         bool
	PhysicalHilbertSpaceSelected bool
	PositiveHamiltonianDerived   bool
	UnitaryDynamicsDerived       bool
	GlobalCausalitySelected      bool
	ArrowOfTimeSelected          bool

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	RealSchwingerSourceImported    bool
	ObservedCorrelationImported    bool
	ConstructiveMeasureImported    bool
	PhysicalOSCertificateImported  bool
	PhysicalWickMapImported        bool
	PhysicalHamiltonianImported    bool
	RealImportSwitchOff            bool
	ComparatorAuthorizationBlocked bool
	NativeSchwingerFunctionWrite   bool
	NativeEuclideanMeasureWrite    bool
	NativeOSPositivityWrite        bool
	NativeWickWrite                bool
	NativeHilbertWrite             bool
	NativeHamiltonianWrite         bool
	NativeUnitaryDynamicsWrite     bool
	NativeGlobalCausalWrite        bool
	NativeTimeArrowWrite           bool
	ReopenedFlavorFirewall         bool
	ReopenedEWScaleFirewall        bool
	ReopenedGravityScaleFirewall   bool
	ReopenedTopologyFirewall       bool
	ReopenedDimensionalFirewall    bool
	ReopenedKreinHilbertFirewall   bool
	NativeRegistryWritten          bool

	Verdict, Reason string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      SwitchSchema
	Switch      ImportSwitch
	Guard       Guard
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
	g539, err := generation2syntheticsourceauthenticityadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate539 synthetic source-authenticity adapter: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g539)}
	a.Schema = buildSchema()
	a.Switch = buildSwitch()
	a.Guard = buildGuard(a.Switch)
	a.Firewall = buildFirewall(a.Switch, a.Guard)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2syntheticsourceauthenticityadapter.Analysis) Inheritance {
	return Inheritance{
		Executed:                     true,
		Gate539AdapterExecuted:       g.Output.Executed && g.Output.Ready && g.Output.AuthenticityPlumbingVerified,
		Gate539RowsAccepted:          g.Import.Loaded && g.Import.Rows == 13 && g.Import.AcceptedRows == 13 && g.Import.RequiredSchemaRowsMatched,
		Gate539ChecksumVerified:      g.Import.ChecksumVerified && g.Output.ChecksumVerified,
		Gate539MetadataSieveEnforced: g.Import.AllRowsBridgeOnly && g.Import.AllRowsNoTheoremInput && g.Import.AllRowsSynthetic && !g.Import.AnyPhysicalClaim && !g.Import.AnyNativePromotionClaim,
		Gate539FixtureRejected:       g.Output.SyntheticFixtureRejectedAsPhysical && !g.Output.PhysicalSourceAuthenticated,
		Gate539NoRealSourceImported:  !g.Import.RealSchwingerSourceImported && !g.Import.ObservedCorrelationLoaded && !g.Import.ConstructiveMeasureLoaded && !g.Import.PhysicalOSCertificateLoaded && !g.Import.PhysicalWickMapLoaded && !g.Import.PhysicalHamiltonianLoaded,
		Gate539NativeWriteBlocked:    !g.Firewall.NativeRegistryWritten && !g.Firewall.NativeSchwingerFunctionWrite && !g.Firewall.NativeOSPositivityWrite && !g.Firewall.NativeWickWrite && !g.Firewall.NativeHamiltonianWrite,
		Gate540SwitchRedirect:        g.Next.Gate == 540,
		Verdict:                      StatusGate539SyntheticAuthenticityInherited,
		Reason:                       "Gate540 inherits Gate539's verified synthetic source-authenticity parser, checksum path, rejection verdict, and real-source native firewall.",
	}
}

func buildSchema() SwitchSchema {
	rows := []SwitchRow{
		{"real_source_import_switch", true, true, true, true, false, "single explicit switch that must be on before any non-synthetic Schwinger source can be loaded"},
		{"explicit_operator_intent", true, true, true, true, false, "human/operator intent must be recorded separately from parser success"},
		{"non_synthetic_source_uri", true, true, true, true, false, "future comparator needs an immutable non-synthetic source location or reference"},
		{"authenticity_ledger_reference", true, true, true, true, false, "future source must reference the Gate538/Gate539 authenticity schema"},
		{"checksum_or_proof_hash_reference", true, true, true, true, false, "future source must carry an integrity hash or proof object"},
		{"license_and_access_grant_reference", true, true, true, true, false, "future source must be legally/auditably accessible"},
		{"source_class_non_synthetic_assertion", true, true, true, true, false, "future source must explicitly declare whether it is physical, constructive, lattice, or theorem-derived"},
		{"gate536_schema_alignment_reference", true, true, true, true, false, "future source must map back to the 19-row Schwinger ledger schema"},
		{"comparator_execution_plan", true, true, true, true, false, "comparator execution must be planned and source-bounded before any run"},
		{"quarantine_output_target", true, true, true, true, false, "all comparator outputs must go to bridge quarantine, not native registry"},
		{"native_write_lock", true, true, true, false, false, "native registry writes remain locked even when the import switch is later enabled"},
		{"rollback_audit_trace", true, true, true, true, false, "future import attempts must be reversible and audit-traceable"},
	}
	s := SwitchSchema{Executed: true, Rows: rows, Verdict: StatusRealImportSwitchAirlockDefined, Reason: "Gate540 enumerates the fail-closed switch rows required before non-synthetic Schwinger source import can even attempt comparator execution."}
	for _, r := range rows {
		if r.Required {
			s.RequiredRows++
		}
		if r.DefaultOff {
			s.DefaultOffRows++
		}
		if r.BridgeOnly {
			s.BridgeOnlyRows++
		}
		if r.Comparator {
			s.ComparatorRows++
		}
		if r.NativeWrite {
			s.NativeWriteRows++
		}
		s.OperatorIntentRow = s.OperatorIntentRow || r.Name == "explicit_operator_intent"
		s.SourceURIRow = s.SourceURIRow || r.Name == "non_synthetic_source_uri"
		s.AuthenticityLedgerRow = s.AuthenticityLedgerRow || r.Name == "authenticity_ledger_reference"
		s.ChecksumProofRow = s.ChecksumProofRow || r.Name == "checksum_or_proof_hash_reference"
		s.LicenseAccessRow = s.LicenseAccessRow || r.Name == "license_and_access_grant_reference"
		s.Gate536AlignmentRow = s.Gate536AlignmentRow || r.Name == "gate536_schema_alignment_reference"
		s.QuarantineRow = s.QuarantineRow || r.Name == "quarantine_output_target"
		s.NativeWriteLockRow = s.NativeWriteLockRow || r.Name == "native_write_lock"
	}
	return s
}

func buildSwitch() ImportSwitch {
	return ImportSwitch{
		Executed:                            true,
		DefaultOff:                          true,
		RealSourceImportEnabled:             false,
		ExplicitOperatorIntentProvided:      false,
		NonSyntheticSourceURIProvided:       false,
		AuthenticityLedgerReferenceProvided: false,
		ChecksumOrProofHashProvided:         false,
		LicenseAndAccessGrantProvided:       false,
		ComparatorPlanAuthorized:            false,
		DryRunOnly:                          true,
		QuarantineTargetDefined:             true,
		NativeWriteAuthorized:               false,
		NativeRegistryWriteRequested:        false,
		ComparatorExecutionAllowed:          false,
		NonSyntheticSourceLoaded:            false,
		ObservedCorrelationLoaded:           false,
		ConstructiveMeasureLoaded:           false,
		PhysicalOSCertificateLoaded:         false,
		PhysicalWickMapLoaded:               false,
		PhysicalHamiltonianLoaded:           false,
		Verdict:                             StatusRealImportSwitchDefaultOff,
		Reason:                              "real/constructive Schwinger source import is intentionally off by default; no source URI, explicit intent, checksum/proof hash, access grant, or comparator authorization is present in this preflight.",
	}
}

func buildGuard(sw ImportSwitch) Guard {
	return Guard{Executed: true, ComparatorExecutionPerformed: false, PhysicalSchwingerImported: false, ConstructiveMeasureImported: false, ObservedCorrelationImported: false, PhysicalSchwingerDerived: false, OSPositivityProven: false, WickRotationSelected: false, PhysicalHilbertSpaceSelected: false, PositiveHamiltonianDerived: false, UnitaryDynamicsDerived: false, GlobalCausalitySelected: false, ArrowOfTimeSelected: false, Verdict: StatusComparatorAuthorizationBlocked, Reason: "with the real-source switch off, every Schwinger/OS/Wick/Hamiltonian comparator remains blocked."}
}

func buildFirewall(sw ImportSwitch, g Guard) Firewall {
	return Firewall{Executed: true, RealSchwingerSourceImported: sw.NonSyntheticSourceLoaded, ObservedCorrelationImported: sw.ObservedCorrelationLoaded || g.ObservedCorrelationImported, ConstructiveMeasureImported: sw.ConstructiveMeasureLoaded || g.ConstructiveMeasureImported, PhysicalOSCertificateImported: sw.PhysicalOSCertificateLoaded, PhysicalWickMapImported: sw.PhysicalWickMapLoaded, PhysicalHamiltonianImported: sw.PhysicalHamiltonianLoaded, RealImportSwitchOff: !sw.RealSourceImportEnabled, ComparatorAuthorizationBlocked: !sw.ComparatorExecutionAllowed && !g.ComparatorExecutionPerformed, NativeSchwingerFunctionWrite: false, NativeEuclideanMeasureWrite: false, NativeOSPositivityWrite: false, NativeWickWrite: false, NativeHilbertWrite: false, NativeHamiltonianWrite: false, NativeUnitaryDynamicsWrite: false, NativeGlobalCausalWrite: false, NativeTimeArrowWrite: false, ReopenedFlavorFirewall: false, ReopenedEWScaleFirewall: false, ReopenedGravityScaleFirewall: false, ReopenedTopologyFirewall: false, ReopenedDimensionalFirewall: false, ReopenedKreinHilbertFirewall: false, NativeRegistryWritten: false, Verdict: StatusFirewallPreserved, Reason: "the switch airlock stays bridge-only: no real source is loaded and no physical-correlation or dynamics object can write to the native registry."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No new native C\\ell(1,7), Schwinger, OS, Wick, Hilbert, Hamiltonian, unitary, global-causal, or time-arrow theorem is added by Gate540."},
		BridgeEntries:        []string{StatusGate539SyntheticAuthenticityInherited, StatusRealImportSwitchAirlockDefined, StatusRealImportSwitchSchemaEnumerated, StatusRealImportSwitchDefaultOff, StatusExplicitOperatorIntentRequired, StatusComparatorAuthorizationBlocked, StatusNoRealSourceImported, StatusNativePromotionRejected},
		EnvironmentalEntries: []string{"Any future non-synthetic Schwinger source, constructive measure, lattice dataset, physical OS certificate, Wick map, Hamiltonian spectrum, uncertainty ledger, and access/license record remains environmental or sourced bridge data."},
		FailedRoutes:         []string{StatusFailedSwitchSchemaNotSchwinger, StatusFailedSwitchSchemaNotOSProof, StatusFailedSwitchSchemaNotWick, StatusFailedSwitchSchemaNotHilbert, StatusFailedSwitchSchemaNotHamiltonian, StatusFailedSwitchSchemaNotUnitary, StatusFailedSwitchSchemaNotGlobal, StatusFailedSwitchSchemaNotArrow, StatusFailedImportSwitchOff, StatusFailedNoExplicitSourceIntent, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"Future Gate 541 may execute a disabled-switch negative-control adapter using a real-looking but untrusted non-synthetic ledger, confirming rejection before any physical import path is opened."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 541, Title: "Real-Looking Schwinger Source Negative-Control Adapter", Reason: "Gate540 defines the real-source switch and leaves it off. The next safe dry run is a real-looking source fixture that must be rejected while the switch is off and while provenance remains insufficient.", PrimaryTask: "Load an intentionally untrusted non-synthetic-looking ledger and prove the import switch refuses comparator execution and native writes."}
}

func truth(a Analysis) string {
	return "Gate540 defines the explicit real-source import switch and keeps it off: parser success, checksums, and provenance schemas are not permission to load non-synthetic Schwinger data, execute OS/Wick/Hamiltonian comparators, or write physical dynamics into the native ASHA registry."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate539AdapterExecuted || !a.Inheritance.Gate539RowsAccepted || !a.Inheritance.Gate539ChecksumVerified || !a.Inheritance.Gate539MetadataSieveEnforced || !a.Inheritance.Gate539FixtureRejected || !a.Inheritance.Gate539NoRealSourceImported || !a.Inheritance.Gate539NativeWriteBlocked || !a.Inheritance.Gate540SwitchRedirect {
		bad = append(bad, "bad Gate539 inheritance")
	}
	if !a.Schema.Executed || len(a.Schema.Rows) != 12 || a.Schema.RequiredRows != 12 || a.Schema.DefaultOffRows != 12 || a.Schema.BridgeOnlyRows != 12 || a.Schema.ComparatorRows != 11 || a.Schema.NativeWriteRows != 0 || !a.Schema.OperatorIntentRow || !a.Schema.SourceURIRow || !a.Schema.AuthenticityLedgerRow || !a.Schema.ChecksumProofRow || !a.Schema.LicenseAccessRow || !a.Schema.Gate536AlignmentRow || !a.Schema.QuarantineRow || !a.Schema.NativeWriteLockRow {
		bad = append(bad, "bad switch schema")
	}
	if !a.Switch.Executed || !a.Switch.DefaultOff || a.Switch.RealSourceImportEnabled || a.Switch.ExplicitOperatorIntentProvided || a.Switch.NonSyntheticSourceURIProvided || a.Switch.AuthenticityLedgerReferenceProvided || a.Switch.ChecksumOrProofHashProvided || a.Switch.LicenseAndAccessGrantProvided || a.Switch.ComparatorPlanAuthorized || !a.Switch.DryRunOnly || !a.Switch.QuarantineTargetDefined || a.Switch.NativeWriteAuthorized || a.Switch.NativeRegistryWriteRequested || a.Switch.ComparatorExecutionAllowed || a.Switch.NonSyntheticSourceLoaded || a.Switch.ObservedCorrelationLoaded || a.Switch.ConstructiveMeasureLoaded || a.Switch.PhysicalOSCertificateLoaded || a.Switch.PhysicalWickMapLoaded || a.Switch.PhysicalHamiltonianLoaded {
		bad = append(bad, "bad real import switch")
	}
	if !a.Guard.Executed || a.Guard.ComparatorExecutionPerformed || a.Guard.PhysicalSchwingerImported || a.Guard.ConstructiveMeasureImported || a.Guard.ObservedCorrelationImported || a.Guard.PhysicalSchwingerDerived || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSpaceSelected || a.Guard.PositiveHamiltonianDerived || a.Guard.UnitaryDynamicsDerived || a.Guard.GlobalCausalitySelected || a.Guard.ArrowOfTimeSelected {
		bad = append(bad, "bad guard")
	}
	if !a.Firewall.Executed || a.Firewall.RealSchwingerSourceImported || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || !a.Firewall.RealImportSwitchOff || !a.Firewall.ComparatorAuthorizationBlocked || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate540 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: adapter=%t rows=%t checksum=%t metadata=%t rejected=%t real_absent=%t native_blocked=%t redirect=%t; %s", x.Verdict, x.Gate539AdapterExecuted, x.Gate539RowsAccepted, x.Gate539ChecksumVerified, x.Gate539MetadataSieveEnforced, x.Gate539FixtureRejected, x.Gate539NoRealSourceImported, x.Gate539NativeWriteBlocked, x.Gate540SwitchRedirect, x.Reason)
}
func FormatSchema(x SwitchSchema) string {
	return fmt.Sprintf("%s: rows=%d required=%d default_off=%d bridge=%d comparator=%d native=%d operator=%t source_uri=%t authenticity=%t checksum=%t license=%t gate536=%t quarantine=%t native_lock=%t; %s", x.Verdict, len(x.Rows), x.RequiredRows, x.DefaultOffRows, x.BridgeOnlyRows, x.ComparatorRows, x.NativeWriteRows, x.OperatorIntentRow, x.SourceURIRow, x.AuthenticityLedgerRow, x.ChecksumProofRow, x.LicenseAccessRow, x.Gate536AlignmentRow, x.QuarantineRow, x.NativeWriteLockRow, x.Reason)
}
func FormatSwitch(x ImportSwitch) string {
	return fmt.Sprintf("%s: default_off=%t enabled=%t intent=%t uri=%t auth_ref=%t checksum=%t license=%t plan=%t dry_run=%t quarantine=%t native_auth=%t native_request=%t comparator_allowed=%t loaded=%t observed=%t measure=%t OS=%t wick=%t hamiltonian=%t; %s", x.Verdict, x.DefaultOff, x.RealSourceImportEnabled, x.ExplicitOperatorIntentProvided, x.NonSyntheticSourceURIProvided, x.AuthenticityLedgerReferenceProvided, x.ChecksumOrProofHashProvided, x.LicenseAndAccessGrantProvided, x.ComparatorPlanAuthorized, x.DryRunOnly, x.QuarantineTargetDefined, x.NativeWriteAuthorized, x.NativeRegistryWriteRequested, x.ComparatorExecutionAllowed, x.NonSyntheticSourceLoaded, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.PhysicalOSCertificateLoaded, x.PhysicalWickMapLoaded, x.PhysicalHamiltonianLoaded, x.Reason)
}
func FormatGuard(x Guard) string {
	return fmt.Sprintf("%s: comparator=%t Schwinger=%t measure=%t observed=%t derived=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t unitary=%t global=%t arrow=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.PhysicalSchwingerImported, x.ConstructiveMeasureImported, x.ObservedCorrelationImported, x.PhysicalSchwingerDerived, x.OSPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSpaceSelected, x.PositiveHamiltonianDerived, x.UnitaryDynamicsDerived, x.GlobalCausalitySelected, x.ArrowOfTimeSelected, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: real=%t observed=%t measure=%t OS_import=%t wick_import=%t hamiltonian_import=%t switch_off=%t comparator_blocked=%t native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopened_flavor=%t reopened_EW=%t reopened_gravity=%t reopened_topology=%t reopened_dimension=%t reopened_Krein=%t native_registry=%t; %s", x.Verdict, x.RealSchwingerSourceImported, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.RealImportSwitchOff, x.ComparatorAuthorizationBlocked, x.NativeSchwingerFunctionWrite, x.NativeEuclideanMeasureWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate539SyntheticAuthenticityInherited, StatusRealImportSwitchAirlockDefined, StatusRealImportSwitchSchemaEnumerated, StatusRealImportSwitchDefaultOff, StatusExplicitOperatorIntentRequired, StatusComparatorAuthorizationBlocked, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSwitchSchemaNotSchwinger, StatusFailedSwitchSchemaNotOSProof, StatusFailedSwitchSchemaNotWick, StatusFailedSwitchSchemaNotHilbert, StatusFailedSwitchSchemaNotHamiltonian, StatusFailedSwitchSchemaNotUnitary, StatusFailedSwitchSchemaNotGlobal, StatusFailedSwitchSchemaNotArrow, StatusFailedImportSwitchOff, StatusFailedNoExplicitSourceIntent, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
