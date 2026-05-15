// Package generation2syntheticcomparatorharnessadapter implements Gate 545:
// Synthetic Comparator-Harness Result Adapter Dry Run.
//
// Gate 544 defined the real-source comparator execution harness but deliberately
// executed no comparator. Gate 545 loads a fake result bundle through that
// harness, writes only a quarantine-output record, verifies abort/rollback and
// native-write-lock metadata, and rejects any promotion to physical Schwinger,
// OS, Wick, Hilbert, Hamiltonian, unitary, causal, or time-arrow law.
package generation2syntheticcomparatorharnessadapter

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

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2realsourcecomparatorharnessairlock"
)

const (
	AuditID       = "GATE545-SYNTHETIC-COMPARATOR-HARNESS-RESULT-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_comparator_harness_result_bundle_gate545.json"

	StatusGate544HarnessInherited             = "CONDITIONAL_SUPPORT_GATE544_COMPARATOR_HARNESS_AIRLOCK_INHERITED"
	StatusSyntheticComparatorBundleLoaded     = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_HARNESS_RESULT_BUNDLE_LOADED"
	StatusSyntheticComparatorAdapterExecuted  = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_HARNESS_ADAPTER_EXECUTED"
	StatusSyntheticComparator16RowsAccepted   = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_HARNESS_16_SCHEMA_ROWS_ACCEPTED"
	StatusSyntheticComparatorChecksumVerified = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_RESULT_BUNDLE_CHECKSUM_VERIFIED"
	StatusSyntheticComparatorMetadataEnforced = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_RESULT_METADATA_SIEVE_ENFORCED"
	StatusSyntheticDryRunExecutedQuarantine   = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_DRY_RUN_EXECUTED_IN_BRIDGE_QUARANTINE"
	StatusSyntheticOutputsParsed              = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_WICK_HILBERT_HAMILTONIAN_OUTPUTS_PARSED"
	StatusSyntheticQuarantineOutputWritten    = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_QUARANTINE_OUTPUT_WRITTEN"
	StatusSyntheticAbortRollbackVerified      = "CONDITIONAL_SUPPORT_SYNTHETIC_ABORT_ROLLBACK_METADATA_VERIFIED"
	StatusSyntheticNativeWriteLockVerified    = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_NATIVE_WRITE_LOCK_VERIFIED"
	StatusNoRealSourceImported                = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_NO_REAL_SOURCE_IMPORTED"
	StatusNativePromotionRejected             = "CONDITIONAL_SUPPORT_SYNTHETIC_COMPARATOR_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing         = "FAILED_ROUTE_GATE545_SYNTHETIC_COMPARATOR_RESULT_BUNDLE_MISSING"
	StatusFailedRowsIncomplete        = "FAILED_ROUTE_GATE545_SYNTHETIC_COMPARATOR_HARNESS_ROWS_INCOMPLETE"
	StatusFailedMetadataIncomplete    = "FAILED_ROUTE_GATE545_SYNTHETIC_COMPARATOR_METADATA_INCOMPLETE"
	StatusFailedChecksumMismatch      = "FAILED_ROUTE_GATE545_SYNTHETIC_COMPARATOR_CHECKSUM_MISMATCH"
	StatusFailedResultBundleLeaked    = "FAILED_ROUTE_GATE545_SYNTHETIC_COMPARATOR_RESULT_BUNDLE_LEAKED_PHYSICAL_CLAIM"
	StatusFailedSyntheticNotReal      = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_OUTPUT_CANNOT_AUTHENTICATE_REAL_SOURCE"
	StatusFailedSyntheticNotSchwinger = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticNotOS        = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedSyntheticNotWick      = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticNotHilbert   = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticNotHamilton  = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticNotUnitary   = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticNotGlobal    = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticNotArrow     = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_RESULT_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedOutputQuarantined     = "FAILED_ROUTE_SYNTHETIC_COMPARATOR_OUTPUT_REMAINS_QUARANTINED"
	StatusFirewallPreserved           = "FIREWALL_PRESERVED_GATE545_SYNTHETIC_COMPARATOR_HARNESS_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked  = "FIREWALL_BLOCKED_GATE545_SYNTHETIC_COMPARATOR_OUTPUT_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate544HarnessDefined     bool
	Gate544RowsEnumerated     bool
	Gate544RequiredRows       int
	Gate544QuarantineSchema   bool
	Gate544AbortConditions    bool
	Gate544NativeWriteLocked  bool
	Gate544ComparatorBlocked  bool
	Gate544NoRealSource       bool
	Gate544NoQuarantineOutput bool
	Gate544RedirectsToGate545 bool

	Verdict, Reason string
}

type HarnessRow struct {
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

type ComparatorSubResult struct {
	Performed                 bool    `json:"performed"`
	Synthetic                 bool    `json:"synthetic"`
	Residual                  float64 `json:"residual,omitempty"`
	NonnegativeSampleCount    int     `json:"nonnegative_sample_count,omitempty"`
	IEpsilonConvention        string  `json:"i_epsilon_convention,omitempty"`
	MinimumEigenvalue         float64 `json:"minimum_eigenvalue,omitempty"`
	PhysicalCertificate       bool    `json:"physical_certificate,omitempty"`
	PhysicalWickMap           bool    `json:"physical_wick_map,omitempty"`
	PhysicalHilbertSpace      bool    `json:"physical_hilbert_space,omitempty"`
	PositiveEnergyCertificate bool    `json:"positive_energy_certificate,omitempty"`
	PhysicalUnitary           bool    `json:"physical_unitarity,omitempty"`
	GlobalHyperbolicity       bool    `json:"global_hyperbolicity,omitempty"`
	ArrowSelected             bool    `json:"arrow_selected,omitempty"`
	NullQuotientRule          string  `json:"null_quotient_rule,omitempty"`
}

type ResultBundle struct {
	OSReflectionPositivity ComparatorSubResult `json:"os_reflection_positivity"`
	WickContinuation       ComparatorSubResult `json:"wick_continuation"`
	HilbertReconstruction  ComparatorSubResult `json:"hilbert_reconstruction"`
	HamiltonianSpectrum    ComparatorSubResult `json:"hamiltonian_spectrum"`
	UnitaryDynamics        ComparatorSubResult `json:"unitary_dynamics"`
	GlobalCausality        ComparatorSubResult `json:"global_causality"`
	TimeArrow              ComparatorSubResult `json:"time_arrow"`
}

type Ledger struct {
	Gate                            int            `json:"gate"`
	LedgerName                      string         `json:"ledger_name"`
	Description                     string         `json:"description"`
	Gate544HarnessReference         string         `json:"gate544_harness_reference"`
	BridgeOnly                      bool           `json:"bridge_only"`
	SyntheticFixture                bool           `json:"synthetic_fixture"`
	DryRunComparatorExecution       bool           `json:"dry_run_comparator_execution"`
	LiveComparatorExecution         bool           `json:"live_comparator_execution"`
	RealSchwingerSourceImported     bool           `json:"real_schwinger_source_imported"`
	AuthenticatedNonSyntheticSource bool           `json:"authenticated_non_synthetic_source"`
	ObservedCorrelationLoaded       bool           `json:"observed_correlation_loaded"`
	ConstructiveMeasureLoaded       bool           `json:"constructive_measure_loaded"`
	PhysicalOSCertificateLoaded     bool           `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded           bool           `json:"physical_wick_map_loaded"`
	PhysicalHamiltonianLoaded       bool           `json:"physical_hamiltonian_loaded"`
	QuarantineOutputWritten         bool           `json:"quarantine_output_written"`
	QuarantineOutputTarget          string         `json:"quarantine_output_target"`
	NativeWriteLock                 bool           `json:"native_write_lock"`
	NativeRegistryWrite             bool           `json:"native_registry_write"`
	AbortTriggered                  bool           `json:"abort_triggered"`
	AbortReason                     string         `json:"abort_reason"`
	RollbackAuditTrace              string         `json:"rollback_audit_trace"`
	HumanReviewRequired             bool           `json:"human_review_required"`
	Source                          string         `json:"source"`
	SourceVersion                   string         `json:"source_version"`
	Convention                      string         `json:"convention"`
	CanonicalPayload                map[string]any `json:"canonical_payload"`
	CanonicalPayloadSHA256          string         `json:"canonical_payload_sha256"`
	ResultBundle                    ResultBundle   `json:"result_bundle"`
	Rows                            []HarnessRow   `json:"rows"`
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
	DryRunExecuted      bool
	LiveExecuted        bool
	RealSource          bool
	AuthenticatedReal   bool
	ObservedLoaded      bool
	MeasureLoaded       bool
	OSCertLoaded        bool
	WickMapLoaded       bool
	HamiltonianLoaded   bool
	QuarantineOutput    bool
	QuarantineTarget    bool
	NativeWriteLock     bool
	NativeWrite         bool
	AbortTriggered      bool
	AbortReason         string
	RollbackTrace       bool
	HumanReviewRequired bool
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

type DryRunResult struct {
	Executed bool

	DryRunComparatorExecuted     bool
	LiveComparatorExecuted       bool
	BridgeQuarantineOnly         bool
	QuarantineOutputWritten      bool
	NativeWriteLocked            bool
	NativeWriteAuthorization     bool
	AbortTriggered               bool
	RollbackTracePresent         bool
	HumanReviewRequired          bool
	OSOutputParsed               bool
	WickOutputParsed             bool
	HilbertOutputParsed          bool
	HamiltonianOutputParsed      bool
	UnitaryDynamicsParsed        bool
	GlobalCausalityParsed        bool
	TimeArrowParsed              bool
	SyntheticOSResidualZero      bool
	SyntheticWickResidualZero    bool
	SyntheticHilbertResidualZero bool
	SyntheticHamiltonianPositive bool
	PhysicalOSProof              bool
	PhysicalWickMap              bool
	PhysicalHilbertSpace         bool
	PhysicalHamiltonian          bool
	PhysicalUnitaryDynamics      bool
	PhysicalGlobalCausality      bool
	PhysicalArrowOfTime          bool

	Verdict, Reason string
	Failures        []string
}

type Firewall struct {
	Executed bool

	RealSchwingerSourceImported   bool
	AuthenticatedRealSource       bool
	ObservedCorrelationImported   bool
	ConstructiveMeasureImported   bool
	PhysicalOSCertificateImported bool
	PhysicalWickMapImported       bool
	PhysicalHamiltonianImported   bool
	DryRunComparatorExecuted      bool
	LiveComparatorExecuted        bool
	QuarantineOutputWritten       bool
	NativeSchwingerFunctionWrite  bool
	NativeEuclideanMeasureWrite   bool
	NativeOSPositivityWrite       bool
	NativeWickWrite               bool
	NativeHilbertWrite            bool
	NativeHamiltonianWrite        bool
	NativeUnitaryDynamicsWrite    bool
	NativeGlobalCausalWrite       bool
	NativeTimeArrowWrite          bool
	NativeRegistryWritten         bool

	Verdict, Reason string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      Import
	DryRun      DryRunResult
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
	cache.Once.Do(func() { cache.a, cache.err = Build(DefaultLedger) })
	return cache.a, cache.err
}

func Build(path string) (Analysis, error) {
	g544, err := generation2realsourcecomparatorharnessairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate544 comparator harness airlock: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g544)}
	ledger, p, err := loadLedger(path)
	if err != nil {
		a.Import = Import{Executed: true, Path: p, Verdict: StatusFailedLedgerMissing, Reason: err.Error(), Failures: []string{StatusFailedLedgerMissing}}
		return a, err
	}
	a.Import = buildImport(ledger, p)
	a.DryRun = buildDryRun(a.Import, ledger.ResultBundle)
	a.Firewall = buildFirewall(a.Import, a.DryRun)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2realsourcecomparatorharnessairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                  true,
		Gate544HarnessDefined:     g.Guard.HarnessDefined,
		Gate544RowsEnumerated:     g.Schema.Executed && g.Schema.RequiredRows == 16,
		Gate544RequiredRows:       g.Schema.RequiredRows,
		Gate544QuarantineSchema:   g.Guard.QuarantineOutputSchemaAvailable,
		Gate544AbortConditions:    g.Guard.AbortConditionsDefined,
		Gate544NativeWriteLocked:  g.Guard.NativeWriteLocked && !g.Guard.NativeWriteAuthorization,
		Gate544ComparatorBlocked:  !g.Guard.ComparatorExecutionPerformed,
		Gate544NoRealSource:       !g.Guard.RealSourceLoaded && !g.Firewall.RealSchwingerSourceImported,
		Gate544NoQuarantineOutput: !g.Guard.QuarantineOutputWritten && !g.Firewall.QuarantineOutputWritten,
		Gate544RedirectsToGate545: g.Next.Gate == 545,
		Verdict:                   StatusGate544HarnessInherited,
		Reason:                    "Gate545 inherits Gate544's 16-row comparator harness, quarantine-output schema, abort path, and native-write lock.",
	}
}

func requiredRows() []string {
	return []string{"comparator_run_identifier", "authorization_manifest_reference", "authenticated_source_ledger_reference", "gate536_schema_alignment_reference", "gate538_authenticity_reference", "gate540_switch_reference", "gate542_authorization_reference", "os_reflection_positivity_input_contract", "wick_continuation_input_contract", "hilbert_reconstruction_input_contract", "hamiltonian_spectrum_input_contract", "comparator_quarantine_output_schema", "comparator_abort_conditions", "native_write_lock", "rollback_audit_trace", "human_review_release_gate"}
}

func buildImport(l Ledger, p string) Import {
	imp := Import{Executed: true, Loaded: true, Path: p, Rows: len(l.Rows), BridgeOnly: l.BridgeOnly, SyntheticFixture: l.SyntheticFixture, DryRunExecuted: l.DryRunComparatorExecution, LiveExecuted: l.LiveComparatorExecution, RealSource: l.RealSchwingerSourceImported, AuthenticatedReal: l.AuthenticatedNonSyntheticSource, ObservedLoaded: l.ObservedCorrelationLoaded, MeasureLoaded: l.ConstructiveMeasureLoaded, OSCertLoaded: l.PhysicalOSCertificateLoaded, WickMapLoaded: l.PhysicalWickMapLoaded, HamiltonianLoaded: l.PhysicalHamiltonianLoaded, QuarantineOutput: l.QuarantineOutputWritten, QuarantineTarget: strings.Contains(l.QuarantineOutputTarget, "bridge/quarantine"), NativeWriteLock: l.NativeWriteLock, NativeWrite: l.NativeRegistryWrite, AbortTriggered: l.AbortTriggered, AbortReason: l.AbortReason, RollbackTrace: l.RollbackAuditTrace != "", HumanReviewRequired: l.HumanReviewRequired, ChecksumExpected: l.CanonicalPayloadSHA256}
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
		imp.Failures = append(imp.Failures, StatusFailedRowsIncomplete)
	}
	if !imp.AllBridgeOnly || !imp.AllComparatorOnly || !imp.AllQuarantineOnly || !imp.AllDryRunOnly || !imp.AllSynthetic || !imp.AllNoTheorem || !imp.AllSourceTagged || !imp.AllConventionTagged || imp.AnyNativePromotion || imp.AnyNativeWrite || imp.AnyPhysicalClaim || imp.AnyObservedClaim {
		imp.Failures = append(imp.Failures, StatusFailedMetadataIncomplete)
	}
	if !imp.ChecksumVerified {
		imp.Failures = append(imp.Failures, StatusFailedChecksumMismatch)
	}
	if imp.RealSource || imp.AuthenticatedReal || imp.ObservedLoaded || imp.MeasureLoaded || imp.OSCertLoaded || imp.WickMapLoaded || imp.HamiltonianLoaded || imp.LiveExecuted || imp.NativeWrite {
		imp.Failures = append(imp.Failures, StatusFailedResultBundleLeaked)
	}
	if len(imp.Failures) == 0 {
		imp.Verdict = StatusSyntheticComparatorAdapterExecuted
		imp.Reason = "Synthetic comparator-harness result bundle parsed all Gate544 rows, verified checksum, and stayed inside bridge quarantine."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Synthetic comparator-harness result bundle failed schema, metadata, checksum, or firewall validation."
	}
	return imp
}

func buildDryRun(i Import, b ResultBundle) DryRunResult {
	d := DryRunResult{Executed: true,
		DryRunComparatorExecuted:     i.Loaded && i.DryRunExecuted && i.ChecksumVerified && len(i.Failures) == 0,
		LiveComparatorExecuted:       i.LiveExecuted,
		BridgeQuarantineOnly:         i.AllBridgeOnly && i.AllQuarantineOnly && i.QuarantineTarget,
		QuarantineOutputWritten:      i.QuarantineOutput,
		NativeWriteLocked:            i.NativeWriteLock && !i.NativeWrite && !i.AnyNativeWrite,
		NativeWriteAuthorization:     false,
		AbortTriggered:               i.AbortTriggered && strings.Contains(i.AbortReason, "synthetic_fixture"),
		RollbackTracePresent:         i.RollbackTrace,
		HumanReviewRequired:          i.HumanReviewRequired,
		OSOutputParsed:               b.OSReflectionPositivity.Performed && b.OSReflectionPositivity.Synthetic,
		WickOutputParsed:             b.WickContinuation.Performed && b.WickContinuation.Synthetic,
		HilbertOutputParsed:          b.HilbertReconstruction.Performed && b.HilbertReconstruction.Synthetic,
		HamiltonianOutputParsed:      b.HamiltonianSpectrum.Performed && b.HamiltonianSpectrum.Synthetic,
		UnitaryDynamicsParsed:        b.UnitaryDynamics.Synthetic,
		GlobalCausalityParsed:        b.GlobalCausality.Synthetic,
		TimeArrowParsed:              b.TimeArrow.Synthetic,
		SyntheticOSResidualZero:      b.OSReflectionPositivity.Residual == 0,
		SyntheticWickResidualZero:    b.WickContinuation.Residual == 0,
		SyntheticHilbertResidualZero: b.HilbertReconstruction.Residual == 0,
		SyntheticHamiltonianPositive: b.HamiltonianSpectrum.MinimumEigenvalue > 0,
		PhysicalOSProof:              b.OSReflectionPositivity.PhysicalCertificate,
		PhysicalWickMap:              b.WickContinuation.PhysicalWickMap,
		PhysicalHilbertSpace:         b.HilbertReconstruction.PhysicalHilbertSpace,
		PhysicalHamiltonian:          b.HamiltonianSpectrum.PositiveEnergyCertificate,
		PhysicalUnitaryDynamics:      b.UnitaryDynamics.PhysicalUnitary,
		PhysicalGlobalCausality:      b.GlobalCausality.GlobalHyperbolicity,
		PhysicalArrowOfTime:          b.TimeArrow.ArrowSelected,
		Verdict:                      StatusSyntheticDryRunExecutedQuarantine,
		Reason:                       "The fake comparator result bundle executes only a synthetic bridge-quarantine dry run; it parses OS/Wick/Hilbert/Hamiltonian output fields while aborting physical source promotion.",
	}
	if !d.DryRunComparatorExecuted || !d.BridgeQuarantineOnly || !d.QuarantineOutputWritten || !d.NativeWriteLocked || !d.AbortTriggered || !d.RollbackTracePresent || !d.HumanReviewRequired {
		d.Failures = append(d.Failures, StatusFailedMetadataIncomplete)
	}
	if !d.OSOutputParsed || !d.WickOutputParsed || !d.HilbertOutputParsed || !d.HamiltonianOutputParsed || !d.SyntheticOSResidualZero || !d.SyntheticWickResidualZero || !d.SyntheticHilbertResidualZero || !d.SyntheticHamiltonianPositive {
		d.Failures = append(d.Failures, StatusFailedResultBundleLeaked)
	}
	if d.LiveComparatorExecuted || d.NativeWriteAuthorization || d.PhysicalOSProof || d.PhysicalWickMap || d.PhysicalHilbertSpace || d.PhysicalHamiltonian || d.PhysicalUnitaryDynamics || d.PhysicalGlobalCausality || d.PhysicalArrowOfTime {
		d.Failures = append(d.Failures, StatusFailedResultBundleLeaked)
	}
	if len(d.Failures) > 0 {
		d.Verdict = strings.Join(d.Failures, ";")
		d.Reason = "Synthetic comparator dry run failed quarantine or physical-claim guards."
	}
	return d
}

func buildFirewall(i Import, d DryRunResult) Firewall {
	return Firewall{Executed: true,
		RealSchwingerSourceImported:   i.RealSource,
		AuthenticatedRealSource:       i.AuthenticatedReal,
		ObservedCorrelationImported:   i.ObservedLoaded,
		ConstructiveMeasureImported:   i.MeasureLoaded,
		PhysicalOSCertificateImported: i.OSCertLoaded,
		PhysicalWickMapImported:       i.WickMapLoaded,
		PhysicalHamiltonianImported:   i.HamiltonianLoaded,
		DryRunComparatorExecuted:      d.DryRunComparatorExecuted,
		LiveComparatorExecuted:        d.LiveComparatorExecuted,
		QuarantineOutputWritten:       d.QuarantineOutputWritten,
		NativeSchwingerFunctionWrite:  false,
		NativeEuclideanMeasureWrite:   false,
		NativeOSPositivityWrite:       false,
		NativeWickWrite:               false,
		NativeHilbertWrite:            false,
		NativeHamiltonianWrite:        false,
		NativeUnitaryDynamicsWrite:    false,
		NativeGlobalCausalWrite:       false,
		NativeTimeArrowWrite:          false,
		NativeRegistryWritten:         false,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "Gate545 writes a synthetic comparator result only to bridge quarantine; no real source, physical certificate, live comparator, or native registry write is produced.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"unchanged: Gate545 writes no native Schwinger, OS, Wick, Hilbert, Hamiltonian, unitarity, global-causal, or time-arrow theorem"},
		BridgeEntries:        []string{"synthetic comparator-harness result adapter", "checksum-verified fake OS/Wick/Hilbert/Hamiltonian result bundle", "quarantine output, abort, rollback, human-review, and native-write-lock dry run"},
		EnvironmentalEntries: []string{"authenticated non-synthetic Schwinger source", "real comparator result", "physical OS proof", "physical Wick map", "physical Hilbert reconstruction", "positive-energy Hamiltonian certificate", "unitary real-time dynamics and global causality"},
		FailedRoutes:         []string{StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputQuarantined, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"Gate546 candidate: comparator output release airlock that defines the human-reviewed path from quarantine report to bridge evidence without native promotion."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 546, Title: "Comparator Output Release Airlock Preflight", Reason: "Gate545 proves a synthetic comparator result can be emitted only to bridge quarantine. The next safe boundary is the release airlock defining what review/certification would be required before any quarantined bridge result can be cited as bridge evidence.", PrimaryTask: "Define release criteria, human review, reproducibility, source authenticity linkage, and native-write lock for future comparator outputs without promoting physics natively."}
}

func truth(a Analysis) string {
	return "Gate545 executes the synthetic comparator-harness result adapter: the fake OS/Wick/Hilbert/Hamiltonian bundle parses, checksum passes, quarantine output is written, abort and rollback metadata are verified, and native writes remain blocked."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate544HarnessDefined || !a.Inheritance.Gate544RowsEnumerated || !a.Inheritance.Gate544QuarantineSchema || !a.Inheritance.Gate544AbortConditions || !a.Inheritance.Gate544NativeWriteLocked || !a.Inheritance.Gate544ComparatorBlocked || !a.Inheritance.Gate544NoRealSource || !a.Inheritance.Gate544NoQuarantineOutput || !a.Inheritance.Gate544RedirectsToGate545 {
		bad = append(bad, "Gate544 inheritance incomplete")
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 16 || len(a.Import.MissingRows) > 0 || a.Import.RejectedRows != 0 || len(a.Import.DuplicateRows) > 0 || !a.Import.ChecksumVerified {
		bad = append(bad, "result bundle schema/checksum incomplete")
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllComparatorOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || !a.Import.AllSourceTagged || !a.Import.AllConventionTagged || a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyObservedClaim {
		bad = append(bad, "metadata sieve failed")
	}
	if !a.DryRun.DryRunComparatorExecuted || a.DryRun.LiveComparatorExecuted || !a.DryRun.BridgeQuarantineOnly || !a.DryRun.QuarantineOutputWritten || !a.DryRun.NativeWriteLocked || a.DryRun.NativeWriteAuthorization || !a.DryRun.AbortTriggered || !a.DryRun.RollbackTracePresent || !a.DryRun.HumanReviewRequired || !a.DryRun.OSOutputParsed || !a.DryRun.WickOutputParsed || !a.DryRun.HilbertOutputParsed || !a.DryRun.HamiltonianOutputParsed || !a.DryRun.SyntheticOSResidualZero || !a.DryRun.SyntheticWickResidualZero || !a.DryRun.SyntheticHilbertResidualZero || !a.DryRun.SyntheticHamiltonianPositive || a.DryRun.PhysicalOSProof || a.DryRun.PhysicalWickMap || a.DryRun.PhysicalHilbertSpace || a.DryRun.PhysicalHamiltonian || a.DryRun.PhysicalUnitaryDynamics || a.DryRun.PhysicalGlobalCausality || a.DryRun.PhysicalArrowOfTime {
		bad = append(bad, "dry-run guard leaked")
	}
	if a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || a.Firewall.LiveComparatorExecuted || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate545 validation failed: %s", strings.Join(bad, "; "))
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
	return fmt.Sprintf("%s: rows=%d harness=%t quarantine_schema=%t aborts=%t native_locked=%t comparator_blocked=%t no_real=%t no_output=%t redirects=%t; %s", x.Verdict, x.Gate544RequiredRows, x.Gate544HarnessDefined, x.Gate544QuarantineSchema, x.Gate544AbortConditions, x.Gate544NativeWriteLocked, x.Gate544ComparatorBlocked, x.Gate544NoRealSource, x.Gate544NoQuarantineOutput, x.Gate544RedirectsToGate545, x.Reason)
}
func FormatImport(x Import) string {
	return fmt.Sprintf("%s;%s;%s;%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%s duplicates=%s checksum=%t expected=%s actual=%s dryrun=%t live=%t quarantine_output=%t target=%t native_lock=%t native_write=%t real=%t auth_real=%t observed=%t measure=%t os_cert=%t wick=%t ham=%t bridge=%t comparator=%t quarantine=%t dryrun_only=%t synthetic=%t no_theorem=%t; %s", StatusSyntheticComparatorBundleLoaded, StatusSyntheticComparator16RowsAccepted, StatusSyntheticComparatorChecksumVerified, x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, strings.Join(x.MissingRows, ","), strings.Join(x.DuplicateRows, ","), x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.DryRunExecuted, x.LiveExecuted, x.QuarantineOutput, x.QuarantineTarget, x.NativeWriteLock, x.NativeWrite, x.RealSource, x.AuthenticatedReal, x.ObservedLoaded, x.MeasureLoaded, x.OSCertLoaded, x.WickMapLoaded, x.HamiltonianLoaded, x.AllBridgeOnly, x.AllComparatorOnly, x.AllQuarantineOnly, x.AllDryRunOnly, x.AllSynthetic, x.AllNoTheorem, x.Reason)
}
func FormatDryRun(x DryRunResult) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s: dryrun=%t live=%t quarantine=%t output=%t native_locked=%t native_auth=%t abort=%t rollback=%t review=%t os=%t wick=%t hilbert=%t ham=%t os_res0=%t wick_res0=%t hilbert_res0=%t ham_positive=%t physical_os=%t physical_wick=%t physical_hilbert=%t physical_ham=%t unitary=%t global=%t arrow=%t; %s", StatusSyntheticDryRunExecutedQuarantine, StatusSyntheticOutputsParsed, StatusSyntheticQuarantineOutputWritten, StatusSyntheticAbortRollbackVerified, StatusSyntheticNativeWriteLockVerified, StatusNoRealSourceImported, StatusNativePromotionRejected, x.DryRunComparatorExecuted, x.LiveComparatorExecuted, x.BridgeQuarantineOnly, x.QuarantineOutputWritten, x.NativeWriteLocked, x.NativeWriteAuthorization, x.AbortTriggered, x.RollbackTracePresent, x.HumanReviewRequired, x.OSOutputParsed, x.WickOutputParsed, x.HilbertOutputParsed, x.HamiltonianOutputParsed, x.SyntheticOSResidualZero, x.SyntheticWickResidualZero, x.SyntheticHilbertResidualZero, x.SyntheticHamiltonianPositive, x.PhysicalOSProof, x.PhysicalWickMap, x.PhysicalHilbertSpace, x.PhysicalHamiltonian, x.PhysicalUnitaryDynamics, x.PhysicalGlobalCausality, x.PhysicalArrowOfTime, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: real=%t auth_real=%t observed=%t measure=%t os_cert=%t wick=%t ham=%t dryrun=%t live=%t output=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t registry=%t; %s", StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputQuarantined, StatusNativePromotionRejected, x.RealSchwingerSourceImported, x.AuthenticatedRealSource, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.DryRunComparatorExecuted, x.LiveComparatorExecuted, x.QuarantineOutputWritten, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate544HarnessInherited, StatusSyntheticComparatorBundleLoaded, StatusSyntheticComparatorAdapterExecuted, StatusSyntheticComparator16RowsAccepted, StatusSyntheticComparatorChecksumVerified, StatusSyntheticComparatorMetadataEnforced, StatusSyntheticDryRunExecutedQuarantine, StatusSyntheticOutputsParsed, StatusSyntheticQuarantineOutputWritten, StatusSyntheticAbortRollbackVerified, StatusSyntheticNativeWriteLockVerified, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputQuarantined, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
