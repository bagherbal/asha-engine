// Package generation2syntheticauthorizationmanifestadapter implements Gate 543:
// Synthetic Comparator Authorization Manifest Adapter Dry Run.
//
// Gate 542 defined the authorization manifest required before any future
// non-synthetic Schwinger source comparator may be staged. Gate 543 loads a
// complete synthetic 14-row manifest, verifies checksum, row coverage,
// source/convention metadata, dry-run comparator scope, quarantine target, and
// native-write lock. The manifest can arm only synthetic bridge-quarantine
// staging; it cannot import a real source or promote physical dynamics.
package generation2syntheticauthorizationmanifestadapter

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2realsourcecomparatorauthorizationairlock"
)

const (
	AuditID       = "GATE543-SYNTHETIC-COMPARATOR-AUTHORIZATION-MANIFEST-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_authorization_manifest_ledger_gate543.json"

	StatusGate542AirlockInherited           = "CONDITIONAL_SUPPORT_GATE542_AUTHORIZATION_MANIFEST_AIRLOCK_INHERITED"
	StatusSyntheticManifestLoaded           = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_MANIFEST_LOADED"
	StatusSyntheticManifestAdapterExecuted  = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_MANIFEST_ADAPTER_EXECUTED"
	StatusSyntheticManifest14RowsAccepted   = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_MANIFEST_14_SCHEMA_ROWS_ACCEPTED"
	StatusSyntheticManifestChecksumVerified = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_MANIFEST_CHECKSUM_VERIFIED"
	StatusSyntheticManifestMetadataEnforced = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_MANIFEST_METADATA_SIEVE_ENFORCED"
	StatusSyntheticDryRunArmedQuarantine    = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_DRY_RUN_ARMED_FOR_BRIDGE_QUARANTINE"
	StatusSyntheticLiveComparatorBlocked    = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_LIVE_COMPARATOR_BLOCKED"
	StatusNoRealSourceImported              = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE543"
	StatusNativePromotionRejected           = "CONDITIONAL_SUPPORT_SYNTHETIC_AUTHORIZATION_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing           = "FAILED_ROUTE_GATE543_SYNTHETIC_AUTHORIZATION_MANIFEST_LEDGER_MISSING"
	StatusFailedSchemaRowsIncomplete    = "FAILED_ROUTE_GATE543_SYNTHETIC_AUTHORIZATION_MANIFEST_SCHEMA_ROWS_INCOMPLETE"
	StatusFailedMetadataIncomplete      = "FAILED_ROUTE_GATE543_SYNTHETIC_AUTHORIZATION_MANIFEST_METADATA_INCOMPLETE"
	StatusFailedChecksumMismatch        = "FAILED_ROUTE_GATE543_SYNTHETIC_AUTHORIZATION_MANIFEST_CHECKSUM_MISMATCH"
	StatusFailedSyntheticNotRealAuth    = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_CANNOT_AUTHORIZE_REAL_SOURCE_IMPORT"
	StatusFailedSyntheticNotSchwinger   = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticNotOSProof     = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedSyntheticNotWick        = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticNotHilbert     = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticNotHamiltonian = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticNotUnitary     = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticNotGlobal      = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticNotArrow       = "FAILED_ROUTE_SYNTHETIC_AUTHORIZATION_MANIFEST_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved             = "FIREWALL_PRESERVED_GATE543_SYNTHETIC_AUTHORIZATION_MANIFEST_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked    = "FIREWALL_BLOCKED_GATE543_SYNTHETIC_AUTHORIZATION_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate542AirlockDefined     bool
	Gate542RowsEnumerated     bool
	Gate542RequiredRows       int
	Gate542BridgeOnlyRows     int
	Gate542ComparatorRows     int
	Gate542NativeWriteRows    int
	Gate542ComparatorBlocked  bool
	Gate542NativeWriteLocked  bool
	Gate542NoRealSource       bool
	Gate542RedirectsToGate543 bool

	Verdict, Reason string
}

type ManifestRow struct {
	SchemaKey       string `json:"schema_key"`
	Source          string `json:"source"`
	SourceVersion   string `json:"source_version"`
	Convention      string `json:"convention"`
	ValueKind       string `json:"value_kind"`
	Value           string `json:"value"`
	BridgeOnly      bool   `json:"bridge_only"`
	ComparatorOnly  bool   `json:"comparator_only"`
	QuarantineOnly  bool   `json:"quarantine_only"`
	DryRunOnly      bool   `json:"dry_run_only"`
	Synthetic       bool   `json:"synthetic"`
	NoTheoremInput  bool   `json:"no_theorem_input"`
	NativePromotion bool   `json:"native_promotion"`
	NativeWrite     bool   `json:"native_write"`
	PhysicalClaim   bool   `json:"physical_claim"`
	Observed        bool   `json:"observed"`
}

type Ledger struct {
	Gate                          int            `json:"gate"`
	LedgerName                    string         `json:"ledger_name"`
	Description                   string         `json:"description"`
	Gate542AirlockReference       string         `json:"gate542_airlock_reference"`
	Gate540SwitchReference        string         `json:"gate540_switch_reference"`
	BridgeOnly                    bool           `json:"bridge_only"`
	SyntheticFixture              bool           `json:"synthetic_fixture"`
	DryRunComparatorAuthorization bool           `json:"dry_run_comparator_authorization"`
	LiveComparatorAuthorization   bool           `json:"live_comparator_authorization"`
	ComparatorExecutionPerformed  bool           `json:"comparator_execution_performed"`
	QuarantineOutputTarget        string         `json:"quarantine_output_target"`
	NativeWriteLock               bool           `json:"native_write_lock"`
	NativeRegistryWrite           bool           `json:"native_registry_write"`
	RealSchwingerSourceImported   bool           `json:"real_schwinger_source_imported"`
	ObservedCorrelationLoaded     bool           `json:"observed_correlation_loaded"`
	ConstructiveMeasureLoaded     bool           `json:"constructive_measure_loaded"`
	PhysicalOSCertificateLoaded   bool           `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded         bool           `json:"physical_wick_map_loaded"`
	PhysicalHamiltonianLoaded     bool           `json:"physical_hamiltonian_loaded"`
	Source                        string         `json:"source"`
	SourceVersion                 string         `json:"source_version"`
	Convention                    string         `json:"convention"`
	CanonicalPayload              map[string]any `json:"canonical_payload"`
	CanonicalPayloadSHA256        string         `json:"canonical_payload_sha256"`
	Rows                          []ManifestRow  `json:"rows"`
}

type Import struct {
	Executed bool
	Loaded   bool
	Path     string

	Rows                int
	AcceptedRows        int
	RejectedRows        int
	MissingRows         []string
	DuplicateRows       []string
	BridgeOnly          bool
	SyntheticFixture    bool
	DryRunAuthorized    bool
	LiveAuthorized      bool
	ComparatorRan       bool
	QuarantineTarget    bool
	NativeWriteLock     bool
	NativeWrite         bool
	RealSource          bool
	ObservedLoaded      bool
	MeasureLoaded       bool
	OSCertLoaded        bool
	WickMapLoaded       bool
	HamiltonianLoaded   bool
	AllBridgeOnly       bool
	AllComparatorOnly   bool
	AllQuarantineOnly   bool
	AllDryRunOnly       bool
	AllSynthetic        bool
	AllNoTheorem        bool
	AllSourceTagged     bool
	AllConventionTagged bool
	AnyNativePromotion  bool
	AnyNativeWrite      bool
	AnyPhysicalClaim    bool
	AnyObservedClaim    bool
	ChecksumExpected    string
	ChecksumActual      string
	ChecksumVerified    bool
	Verdict, Reason     string
	Failures            []string
}

type Authorization struct {
	Executed bool

	DryRunAuthorizationArmed     bool
	LiveComparatorAuthorization  bool
	ComparatorExecutionPerformed bool
	BridgeQuarantineOnly         bool
	QuarantineOutputTarget       bool
	NativeWriteLocked            bool
	NativeWriteAuthorization     bool
	SyntheticManifestOnly        bool
	CanImportRealSource          bool
	RealSourceLoaded             bool
	ObservedCorrelationLoaded    bool
	ConstructiveMeasureLoaded    bool
	PhysicalOSCertificateLoaded  bool
	PhysicalWickMapLoaded        bool
	PhysicalHamiltonianLoaded    bool
	PhysicalSchwingerDerived     bool
	OSPositivityProven           bool
	WickRotationSelected         bool
	PhysicalHilbertSelected      bool
	HamiltonianDerived           bool
	UnitaryDynamicsDerived       bool
	GlobalCausalitySelected      bool
	ArrowOfTimeSelected          bool

	Verdict, Reason string
	Failures        []string
}

type Firewall struct {
	Executed bool

	RealSchwingerSourceImported   bool
	ObservedCorrelationImported   bool
	ConstructiveMeasureImported   bool
	PhysicalOSCertificateImported bool
	PhysicalWickMapImported       bool
	PhysicalHamiltonianImported   bool
	ComparatorExecutionPerformed  bool
	LiveComparatorAuthorized      bool
	NativeSchwingerFunctionWrite  bool
	NativeEuclideanMeasureWrite   bool
	NativeOSPositivityWrite       bool
	NativeWickWrite               bool
	NativeHilbertWrite            bool
	NativeHamiltonianWrite        bool
	NativeUnitaryDynamicsWrite    bool
	NativeGlobalCausalWrite       bool
	NativeTimeArrowWrite          bool
	ReopenedFlavorFirewall        bool
	ReopenedEWScaleFirewall       bool
	ReopenedGravityScaleFirewall  bool
	ReopenedTopologyFirewall      bool
	ReopenedDimensionalFirewall   bool
	ReopenedKreinHilbertFirewall  bool
	NativeRegistryWritten         bool

	Verdict, Reason string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance   Inheritance
	Import        Import
	Authorization Authorization
	Firewall      Firewall
	Registry      RegistryUpdate
	Next          NextStep
	Truth         string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build(DefaultLedger) })
	return cache.a, cache.err
}

func Build(path string) (Analysis, error) {
	g542, err := generation2realsourcecomparatorauthorizationairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate542 authorization manifest airlock: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g542)}
	ledger, p, err := loadLedger(path)
	if err != nil {
		a.Import = Import{Executed: true, Path: p, Verdict: StatusFailedLedgerMissing, Reason: err.Error(), Failures: []string{StatusFailedLedgerMissing}}
		return a, err
	}
	a.Import = buildImport(ledger, p)
	a.Authorization = buildAuthorization(a.Import)
	a.Firewall = buildFirewall(a.Authorization)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2realsourcecomparatorauthorizationairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                  true,
		Gate542AirlockDefined:     g.Schema.Executed,
		Gate542RowsEnumerated:     len(g.Schema.Rows) == 14 && g.Schema.RequiredRows == 14,
		Gate542RequiredRows:       g.Schema.RequiredRows,
		Gate542BridgeOnlyRows:     g.Schema.BridgeOnlyRows,
		Gate542ComparatorRows:     g.Schema.ComparatorRows,
		Gate542NativeWriteRows:    g.Schema.NativeWriteRows,
		Gate542ComparatorBlocked:  !g.Authorization.ComparatorExecutionPerformed,
		Gate542NativeWriteLocked:  g.Authorization.NativeWriteLocked && !g.Authorization.NativeWriteAuthorization,
		Gate542NoRealSource:       !g.Authorization.RealSourceLoaded && !g.Firewall.RealSchwingerSourceImported,
		Gate542RedirectsToGate543: g.Next.Gate == 543,
		Verdict:                   StatusGate542AirlockInherited,
		Reason:                    "Gate543 inherits Gate542's 14-row real-source comparator authorization manifest and its bridge-quarantine native-write lock.",
	}
}

func requiredRows() []string {
	return []string{"operator_intent_signature", "authenticated_source_identity", "authenticity_ledger_reference", "gate536_schema_alignment_report", "gate540_switch_enable_record", "license_and_access_grant", "checksum_or_proof_hash_verification", "provenance_integrity_report", "comparator_scope_declaration", "quarantine_output_target", "dry_run_or_live_comparator_mode", "native_write_lock", "rollback_audit_trace", "human_review_attestation"}
}

func buildImport(l Ledger, p string) Import {
	imp := Import{Executed: true, Loaded: true, Path: p, Rows: len(l.Rows), BridgeOnly: l.BridgeOnly, SyntheticFixture: l.SyntheticFixture, DryRunAuthorized: l.DryRunComparatorAuthorization, LiveAuthorized: l.LiveComparatorAuthorization, ComparatorRan: l.ComparatorExecutionPerformed, QuarantineTarget: l.QuarantineOutputTarget != "" && strings.Contains(l.QuarantineOutputTarget, "bridge"), NativeWriteLock: l.NativeWriteLock, NativeWrite: l.NativeRegistryWrite, RealSource: l.RealSchwingerSourceImported, ObservedLoaded: l.ObservedCorrelationLoaded, MeasureLoaded: l.ConstructiveMeasureLoaded, OSCertLoaded: l.PhysicalOSCertificateLoaded, WickMapLoaded: l.PhysicalWickMapLoaded, HamiltonianLoaded: l.PhysicalHamiltonianLoaded, ChecksumExpected: l.CanonicalPayloadSHA256}
	seen := map[string]bool{}
	req := map[string]bool{}
	for _, r := range requiredRows() {
		req[r] = false
	}
	imp.AllBridgeOnly, imp.AllComparatorOnly, imp.AllQuarantineOnly, imp.AllDryRunOnly, imp.AllSynthetic, imp.AllNoTheorem, imp.AllSourceTagged, imp.AllConventionTagged = true, true, true, true, true, true, true, true
	for _, r := range l.Rows {
		if seen[r.SchemaKey] {
			imp.DuplicateRows = append(imp.DuplicateRows, r.SchemaKey)
		}
		seen[r.SchemaKey] = true
		if _, ok := req[r.SchemaKey]; ok {
			req[r.SchemaKey] = true
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
		imp.AllBridgeOnly = imp.AllBridgeOnly && r.BridgeOnly
		imp.AllComparatorOnly = imp.AllComparatorOnly && r.ComparatorOnly
		imp.AllQuarantineOnly = imp.AllQuarantineOnly && r.QuarantineOnly
		imp.AllDryRunOnly = imp.AllDryRunOnly && r.DryRunOnly
		imp.AllSynthetic = imp.AllSynthetic && r.Synthetic
		imp.AllNoTheorem = imp.AllNoTheorem && r.NoTheoremInput
		imp.AllSourceTagged = imp.AllSourceTagged && r.Source != "" && r.SourceVersion != ""
		imp.AllConventionTagged = imp.AllConventionTagged && r.Convention != ""
		imp.AnyNativePromotion = imp.AnyNativePromotion || r.NativePromotion
		imp.AnyNativeWrite = imp.AnyNativeWrite || r.NativeWrite
		imp.AnyPhysicalClaim = imp.AnyPhysicalClaim || r.PhysicalClaim
		imp.AnyObservedClaim = imp.AnyObservedClaim || r.Observed
	}
	for _, r := range requiredRows() {
		if !req[r] {
			imp.MissingRows = append(imp.MissingRows, r)
		}
	}
	sort.Strings(imp.MissingRows)
	sort.Strings(imp.DuplicateRows)
	imp.ChecksumActual = checksum(l.CanonicalPayload)
	imp.ChecksumVerified = imp.ChecksumActual == imp.ChecksumExpected
	if len(imp.MissingRows) > 0 || imp.RejectedRows > 0 || len(imp.DuplicateRows) > 0 {
		imp.Failures = append(imp.Failures, StatusFailedSchemaRowsIncomplete)
	}
	if !imp.AllBridgeOnly || !imp.AllComparatorOnly || !imp.AllQuarantineOnly || !imp.AllDryRunOnly || !imp.AllSynthetic || !imp.AllNoTheorem || !imp.AllSourceTagged || !imp.AllConventionTagged || imp.AnyNativePromotion || imp.AnyNativeWrite || imp.AnyPhysicalClaim || imp.AnyObservedClaim {
		imp.Failures = append(imp.Failures, StatusFailedMetadataIncomplete)
	}
	if !imp.ChecksumVerified {
		imp.Failures = append(imp.Failures, StatusFailedChecksumMismatch)
	}
	if len(imp.Failures) == 0 {
		imp.Verdict = StatusSyntheticManifestAdapterExecuted
		imp.Reason = "Synthetic authorization manifest parsed all Gate542 rows, verified checksum, and preserved dry-run bridge quarantine."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Synthetic authorization manifest failed schema, metadata, or checksum validation."
	}
	return imp
}

func buildAuthorization(i Import) Authorization {
	a := Authorization{Executed: true,
		DryRunAuthorizationArmed:     i.Loaded && i.DryRunAuthorized && i.QuarantineTarget && i.NativeWriteLock && i.ChecksumVerified && len(i.Failures) == 0,
		LiveComparatorAuthorization:  false,
		ComparatorExecutionPerformed: false,
		BridgeQuarantineOnly:         i.AllBridgeOnly && i.AllQuarantineOnly && i.QuarantineTarget,
		QuarantineOutputTarget:       i.QuarantineTarget,
		NativeWriteLocked:            i.NativeWriteLock && !i.NativeWrite && !i.AnyNativeWrite,
		NativeWriteAuthorization:     false,
		SyntheticManifestOnly:        i.SyntheticFixture && i.AllSynthetic,
		CanImportRealSource:          false,
		RealSourceLoaded:             i.RealSource,
		ObservedCorrelationLoaded:    i.ObservedLoaded,
		ConstructiveMeasureLoaded:    i.MeasureLoaded,
		PhysicalOSCertificateLoaded:  i.OSCertLoaded,
		PhysicalWickMapLoaded:        i.WickMapLoaded,
		PhysicalHamiltonianLoaded:    i.HamiltonianLoaded,
		Verdict:                      StatusSyntheticDryRunArmedQuarantine,
		Reason:                       "The synthetic manifest arms only a bridge-quarantine dry-run authorization state; live comparator execution and native writes remain blocked.",
	}
	if !a.DryRunAuthorizationArmed {
		a.Failures = append(a.Failures, StatusFailedMetadataIncomplete)
	}
	if i.LiveAuthorized || i.ComparatorRan || i.RealSource || i.ObservedLoaded || i.MeasureLoaded || i.OSCertLoaded || i.WickMapLoaded || i.HamiltonianLoaded || i.NativeWrite {
		a.Failures = append(a.Failures, StatusFailedSyntheticNotRealAuth)
	}
	if len(a.Failures) > 0 {
		a.Verdict = strings.Join(a.Failures, ";")
		a.Reason = "Synthetic manifest attempted to exceed quarantine-only dry-run authorization."
	}
	return a
}

func buildFirewall(a Authorization) Firewall {
	return Firewall{Executed: true,
		RealSchwingerSourceImported:   a.RealSourceLoaded,
		ObservedCorrelationImported:   a.ObservedCorrelationLoaded,
		ConstructiveMeasureImported:   a.ConstructiveMeasureLoaded,
		PhysicalOSCertificateImported: a.PhysicalOSCertificateLoaded,
		PhysicalWickMapImported:       a.PhysicalWickMapLoaded,
		PhysicalHamiltonianImported:   a.PhysicalHamiltonianLoaded,
		ComparatorExecutionPerformed:  a.ComparatorExecutionPerformed,
		LiveComparatorAuthorized:      a.LiveComparatorAuthorization,
		NativeSchwingerFunctionWrite:  false,
		NativeEuclideanMeasureWrite:   false,
		NativeOSPositivityWrite:       false,
		NativeWickWrite:               false,
		NativeHilbertWrite:            false,
		NativeHamiltonianWrite:        false,
		NativeUnitaryDynamicsWrite:    false,
		NativeGlobalCausalWrite:       false,
		NativeTimeArrowWrite:          false,
		ReopenedFlavorFirewall:        false,
		ReopenedEWScaleFirewall:       false,
		ReopenedGravityScaleFirewall:  false,
		ReopenedTopologyFirewall:      false,
		ReopenedDimensionalFirewall:   false,
		ReopenedKreinHilbertFirewall:  false,
		NativeRegistryWritten:         false,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate543 validates authorization-manifest plumbing while keeping source import, live comparator execution, and every native registry write closed.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"unchanged: Gate543 writes no native Schwinger, OS, Wick, Hilbert, Hamiltonian, unitarity, causal, or time-arrow theorem"},
		BridgeEntries:        []string{"synthetic 14-row comparator authorization manifest adapter", "quarantine-only dry-run authorization state", "checksum-verified manifest parser with native-write lock"},
		EnvironmentalEntries: []string{"real non-synthetic source authorization", "live comparator execution", "physical Schwinger/OS/Wick/Hamiltonian certificates"},
		FailedRoutes:         []string{StatusFailedSyntheticNotRealAuth, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"Gate544 candidate: real-source comparator execution harness preflight that defines comparator input/output contracts without loading a source."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 544, Title: "Real-Source Comparator Execution Harness Preflight", Reason: "Gate543 proves a complete synthetic authorization manifest can arm only quarantine dry-run metadata. The next safe boundary is the comparator harness contract itself, still with no source loaded and no native writes.", PrimaryTask: "Define OS/Wick/Hamiltonian comparator input/output contracts, quarantine result schema, and abort conditions without executing a real-source comparator."}
}

func truth(a Analysis) string {
	return "Gate543 verifies the synthetic authorization manifest adapter: all 14 Gate542 rows parse, checksum and metadata pass, quarantine-only dry-run authorization is armed, but real source import, live comparator execution, physical dynamics, and native writes remain blocked."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate542RowsEnumerated || !a.Inheritance.Gate542ComparatorBlocked || !a.Inheritance.Gate542NativeWriteLocked || !a.Inheritance.Gate542RedirectsToGate543 {
		bad = append(bad, "Gate542 inheritance incomplete")
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 14 || len(a.Import.MissingRows) > 0 || a.Import.RejectedRows != 0 || len(a.Import.DuplicateRows) > 0 || !a.Import.ChecksumVerified {
		bad = append(bad, "manifest schema/checksum incomplete")
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllComparatorOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || !a.Import.AllSourceTagged || !a.Import.AllConventionTagged || a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyObservedClaim {
		bad = append(bad, "manifest metadata sieve failed")
	}
	if !a.Authorization.DryRunAuthorizationArmed || a.Authorization.LiveComparatorAuthorization || a.Authorization.ComparatorExecutionPerformed || !a.Authorization.BridgeQuarantineOnly || !a.Authorization.NativeWriteLocked || a.Authorization.NativeWriteAuthorization || !a.Authorization.SyntheticManifestOnly || a.Authorization.CanImportRealSource {
		bad = append(bad, "authorization state leaked")
	}
	if a.Firewall.RealSchwingerSourceImported || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || a.Firewall.ComparatorExecutionPerformed || a.Firewall.LiveComparatorAuthorized || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate543 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func loadLedger(path string) (Ledger, string, error) {
	p, err := resolve(path)
	if err != nil {
		return Ledger{}, path, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return Ledger{}, p, err
	}
	var l Ledger
	if err := json.Unmarshal(b, &l); err != nil {
		return Ledger{}, p, err
	}
	return l, p, nil
}

func resolve(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("runtime caller unavailable")
	}
	dir := filepath.Dir(file)
	for i := 0; i < 6; i++ {
		cand := filepath.Join(dir, path)
		if _, err := os.Stat(cand); err == nil {
			return cand, nil
		}
		dir = filepath.Dir(dir)
	}
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	return "", fmt.Errorf("ledger not found: %s", path)
}

func checksum(v map[string]any) string {
	b, _ := json.Marshal(v)
	s := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(s[:])
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: rows=%d bridge=%d comparator=%d native_write=%d blocked=%t locked=%t no_real=%t redirects=%t; %s", x.Verdict, x.Gate542RequiredRows, x.Gate542BridgeOnlyRows, x.Gate542ComparatorRows, x.Gate542NativeWriteRows, x.Gate542ComparatorBlocked, x.Gate542NativeWriteLocked, x.Gate542NoRealSource, x.Gate542RedirectsToGate543, x.Reason)
}
func FormatImport(x Import) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%s duplicates=%s checksum=%t expected=%s actual=%s bridge=%t comparator=%t quarantine=%t dryrun=%t synthetic=%t no_theorem=%t native_promotion=%t native_write=%t physical=%t observed=%t; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, strings.Join(x.MissingRows, ","), strings.Join(x.DuplicateRows, ","), x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.AllBridgeOnly, x.AllComparatorOnly, x.AllQuarantineOnly, x.AllDryRunOnly, x.AllSynthetic, x.AllNoTheorem, x.AnyNativePromotion, x.AnyNativeWrite, x.AnyPhysicalClaim, x.AnyObservedClaim, x.Reason)
}
func FormatAuthorization(x Authorization) string {
	return fmt.Sprintf("%s: dryrun=%t live=%t comparator=%t quarantine=%t target=%t native_locked=%t native_auth=%t synthetic_only=%t can_import_real=%t real=%t observed=%t measure=%t os_cert=%t wick=%t ham=%t; %s", x.Verdict, x.DryRunAuthorizationArmed, x.LiveComparatorAuthorization, x.ComparatorExecutionPerformed, x.BridgeQuarantineOnly, x.QuarantineOutputTarget, x.NativeWriteLocked, x.NativeWriteAuthorization, x.SyntheticManifestOnly, x.CanImportRealSource, x.RealSourceLoaded, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.PhysicalOSCertificateLoaded, x.PhysicalWickMapLoaded, x.PhysicalHamiltonianLoaded, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: real=%t observed=%t measure=%t os_cert=%t wick=%t ham=%t comparator=%t live=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t registry=%t; %s", StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, x.RealSchwingerSourceImported, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.ComparatorExecutionPerformed, x.LiveComparatorAuthorized, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate542AirlockInherited, StatusSyntheticManifestLoaded, StatusSyntheticManifestAdapterExecuted, StatusSyntheticManifest14RowsAccepted, StatusSyntheticManifestChecksumVerified, StatusSyntheticManifestMetadataEnforced, StatusSyntheticDryRunArmedQuarantine, StatusSyntheticLiveComparatorBlocked, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotRealAuth, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
