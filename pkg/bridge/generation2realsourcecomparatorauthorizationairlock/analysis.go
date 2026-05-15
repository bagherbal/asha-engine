// Package generation2realsourcecomparatorauthorizationairlock implements Gate 542:
// Real Source Comparator Authorization Manifest Airlock.
//
// Gate 541 proved that a real-looking Schwinger source is rejected when the
// real-source switch is off, operator intent is absent, and provenance is
// incomplete. Gate 542 defines the authorization manifest that would be required
// before any future non-synthetic Schwinger comparator may run. The manifest is
// bridge-quarantine only: it authorizes at most comparator staging and never
// native registry writes or physical-dynamics promotion.
package generation2realsourcecomparatorauthorizationairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2reallookingschwingersourcenegativecontroladapter"
)

const (
	AuditID = "GATE542-REAL-SOURCE-COMPARATOR-AUTHORIZATION-MANIFEST-AIRLOCK"

	StatusGate541NegativeControlInherited         = "CONDITIONAL_SUPPORT_GATE541_NEGATIVE_CONTROL_INHERITED"
	StatusAuthorizationManifestAirlockDefined     = "CONDITIONAL_SUPPORT_REAL_SOURCE_COMPARATOR_AUTHORIZATION_MANIFEST_AIRLOCK_DEFINED"
	StatusAuthorizationManifestRowsEnumerated     = "CONDITIONAL_SUPPORT_AUTHORIZATION_MANIFEST_SCHEMA_ROWS_ENUMERATED"
	StatusBridgeQuarantineOnlyAuthorization       = "CONDITIONAL_SUPPORT_COMPARATOR_AUTHORIZATION_LIMITED_TO_BRIDGE_QUARANTINE"
	StatusExplicitOperatorIntentManifestRequired  = "CONDITIONAL_SUPPORT_EXPLICIT_OPERATOR_INTENT_MANIFEST_REQUIRED"
	StatusNativeWriteLockManifestRequired         = "CONDITIONAL_SUPPORT_NATIVE_WRITE_LOCK_MANIFEST_REQUIRED"
	StatusComparatorAuthorizationBlockedPreflight = "CONDITIONAL_SUPPORT_REAL_SOURCE_COMPARATOR_AUTHORIZATION_BLOCKED_IN_PREFLIGHT"
	StatusNoRealSourceImported                    = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE542"
	StatusNativePromotionRejected                 = "CONDITIONAL_SUPPORT_AUTHORIZATION_MANIFEST_NATIVE_PROMOTION_REJECTED"

	StatusFailedManifestNotSchwinger        = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedManifestNotOSProof          = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedManifestNotWick             = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedManifestNotHilbert          = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedManifestNotHamiltonian      = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedManifestNotUnitary          = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedManifestNotGlobal           = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedManifestNotArrow            = "FAILED_ROUTE_AUTHORIZATION_MANIFEST_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedNoLiveComparatorInPreflight = "FAILED_ROUTE_REAL_SOURCE_COMPARATOR_NOT_AUTHORIZED_IN_GATE542_PREFLIGHT"
	StatusFailedNoSourceManifestImported    = "FAILED_ROUTE_NO_NON_SYNTHETIC_SOURCE_AUTHORIZATION_MANIFEST_IMPORTED_IN_GATE542"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_GATE542_AUTHORIZATION_MANIFEST_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked        = "FIREWALL_BLOCKED_GATE542_COMPARATOR_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate541NegativeControlExecuted bool
	Gate541ChecksumVerified        bool
	Gate541SwitchOffRejected       bool
	Gate541NoIntentRejected        bool
	Gate541InsufficientProvenance  bool
	Gate541ComparatorBlocked       bool
	Gate541NoNativeWrite           bool
	Gate541RedirectsToGate542      bool

	Verdict, Reason string
}

type ManifestRow struct {
	Name        string
	Required    bool
	BridgeOnly  bool
	Comparator  bool
	Quarantine  bool
	NativeWrite bool
	Reason      string
}

type ManifestSchema struct {
	Executed bool

	Rows                  []ManifestRow
	RequiredRows          int
	BridgeOnlyRows        int
	ComparatorRows        int
	QuarantineRows        int
	NativeWriteRows       int
	OperatorIntentRow     bool
	SourceIdentityRow     bool
	AuthenticityLedgerRow bool
	Gate536AlignmentRow   bool
	Gate540SwitchRow      bool
	AccessGrantRow        bool
	ChecksumProofRow      bool
	ProvenanceReportRow   bool
	ComparatorScopeRow    bool
	QuarantineTargetRow   bool
	ModeDeclarationRow    bool
	NativeWriteLockRow    bool
	RollbackTraceRow      bool
	HumanReviewRow        bool

	Verdict, Reason string
}

type AuthorizationState struct {
	Executed bool

	ManifestImported              bool
	ComparatorLiveAuthorization   bool
	ComparatorDryRunAuthorization bool
	ExplicitOperatorIntentPresent bool
	AuthenticatedSourceIdentity   bool
	AuthenticityLedgerBound       bool
	Gate536SchemaAlignmentBound   bool
	Gate540SwitchEnableBound      bool
	LicenseAccessGrantBound       bool
	ChecksumProofBound            bool
	ProvenanceIntegrityBound      bool
	ComparatorScopeBound          bool
	QuarantineTargetBound         bool
	NativeWriteLocked             bool
	RollbackTraceReady            bool
	HumanReviewPresent            bool
	BridgeQuarantineOnly          bool
	NativeWriteAuthorization      bool
	RealSourceLoaded              bool
	ObservedCorrelationLoaded     bool
	ConstructiveMeasureLoaded     bool
	PhysicalOSCertificateLoaded   bool
	PhysicalWickMapLoaded         bool
	PhysicalHamiltonianLoaded     bool
	ComparatorExecutionPerformed  bool

	Verdict, Reason string
}

type Guard struct {
	Executed bool

	ComparatorCanRunInPreflight bool
	AuthorizationCanWriteNative bool
	RealSchwingerImported       bool
	PhysicalSchwingerDerived    bool
	OSPositivityProven          bool
	WickRotationSelected        bool
	PhysicalHilbertSelected     bool
	HamiltonianDerived          bool
	UnitaryDynamicsDerived      bool
	GlobalCausalitySelected     bool
	ArrowOfTimeSelected         bool

	Verdict, Reason string
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
	Inheritance   Inheritance
	Schema        ManifestSchema
	Authorization AuthorizationState
	Guard         Guard
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
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g541, err := generation2reallookingschwingersourcenegativecontroladapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate541 real-looking negative-control adapter: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g541)}
	a.Schema = buildSchema()
	a.Authorization = buildAuthorization(a.Schema)
	a.Guard = buildGuard(a.Authorization)
	a.Firewall = buildFirewall(a.Authorization, a.Guard)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2reallookingschwingersourcenegativecontroladapter.Analysis) Inheritance {
	return Inheritance{
		Executed:                       true,
		Gate541NegativeControlExecuted: g.Import.Executed && g.Rejection.Executed,
		Gate541ChecksumVerified:        g.Import.ChecksumVerified,
		Gate541SwitchOffRejected:       g.Rejection.SwitchOff && g.Rejection.RejectedAsPhysicalSource,
		Gate541NoIntentRejected:        g.Rejection.NoExplicitOperatorIntent,
		Gate541InsufficientProvenance:  g.Rejection.InsufficientProvenance,
		Gate541ComparatorBlocked:       !g.Rejection.ComparatorExecutionPerformed && !g.Firewall.ComparatorExecutionPerformed,
		Gate541NoNativeWrite:           !g.Firewall.NativeRegistryWritten,
		Gate541RedirectsToGate542:      g.Next.Gate == 542,
		Verdict:                        StatusGate541NegativeControlInherited,
		Reason:                         "Gate542 inherits Gate541's proof that real-looking source metadata and checksum plumbing are rejected before comparator execution when authorization is absent.",
	}
}

func buildSchema() ManifestSchema {
	rows := []ManifestRow{
		{"operator_intent_signature", true, true, true, true, false, "Explicit human/operator intent must be bound before any real-source comparator may arm."},
		{"authenticated_source_identity", true, true, true, true, false, "The source identity must be immutable, non-synthetic, and externally authenticated."},
		{"authenticity_ledger_reference", true, true, true, true, false, "The Gate538/Gate539 source-authenticity sieve must be referenced by hash or artifact ID."},
		{"gate536_schema_alignment_report", true, true, true, true, false, "The candidate source must align with every Gate536 Schwinger-function row."},
		{"gate540_switch_enable_record", true, true, true, true, false, "The default-off switch must have an explicit enable record, not an implicit parser side effect."},
		{"license_and_access_grant", true, true, true, true, false, "Access rights and redistribution limits must be captured before import."},
		{"checksum_or_proof_hash_verification", true, true, true, true, false, "The exact payload/proof hash must be verified before comparator staging."},
		{"provenance_integrity_report", true, true, true, true, false, "Construction, measure, regulator, renormalization, and reproducibility provenance must be complete."},
		{"comparator_scope_declaration", true, true, true, true, false, "The comparator must declare exactly which Schwinger/OS/Wick/Hamiltonian checks it may attempt."},
		{"quarantine_output_target", true, true, true, true, false, "All outputs must land in bridge quarantine, never in native law-space."},
		{"dry_run_or_live_comparator_mode", true, true, true, true, false, "The manifest must distinguish dry-run staging from live source comparator execution."},
		{"native_write_lock", true, true, true, true, false, "Native registry writes must remain locked even if a bridge comparator later passes."},
		{"rollback_audit_trace", true, true, true, true, false, "Every staged import must have a reversible audit trail."},
		{"human_review_attestation", true, true, true, true, false, "A human review row is required because source authenticity is an epistemic, not algebraic, claim."},
	}
	s := ManifestSchema{Executed: true, Rows: rows, Verdict: StatusAuthorizationManifestRowsEnumerated, Reason: "Gate542 enumerates the bridge-only authorization manifest for future real-source comparator staging."}
	for _, r := range rows {
		if r.Required {
			s.RequiredRows++
		}
		if r.BridgeOnly {
			s.BridgeOnlyRows++
		}
		if r.Comparator {
			s.ComparatorRows++
		}
		if r.Quarantine {
			s.QuarantineRows++
		}
		if r.NativeWrite {
			s.NativeWriteRows++
		}
		switch r.Name {
		case "operator_intent_signature":
			s.OperatorIntentRow = true
		case "authenticated_source_identity":
			s.SourceIdentityRow = true
		case "authenticity_ledger_reference":
			s.AuthenticityLedgerRow = true
		case "gate536_schema_alignment_report":
			s.Gate536AlignmentRow = true
		case "gate540_switch_enable_record":
			s.Gate540SwitchRow = true
		case "license_and_access_grant":
			s.AccessGrantRow = true
		case "checksum_or_proof_hash_verification":
			s.ChecksumProofRow = true
		case "provenance_integrity_report":
			s.ProvenanceReportRow = true
		case "comparator_scope_declaration":
			s.ComparatorScopeRow = true
		case "quarantine_output_target":
			s.QuarantineTargetRow = true
		case "dry_run_or_live_comparator_mode":
			s.ModeDeclarationRow = true
		case "native_write_lock":
			s.NativeWriteLockRow = true
		case "rollback_audit_trace":
			s.RollbackTraceRow = true
		case "human_review_attestation":
			s.HumanReviewRow = true
		}
	}
	return s
}

func buildAuthorization(s ManifestSchema) AuthorizationState {
	return AuthorizationState{
		Executed:                      true,
		ManifestImported:              false,
		ComparatorLiveAuthorization:   false,
		ComparatorDryRunAuthorization: false,
		ExplicitOperatorIntentPresent: false,
		AuthenticatedSourceIdentity:   false,
		AuthenticityLedgerBound:       false,
		Gate536SchemaAlignmentBound:   false,
		Gate540SwitchEnableBound:      false,
		LicenseAccessGrantBound:       false,
		ChecksumProofBound:            false,
		ProvenanceIntegrityBound:      false,
		ComparatorScopeBound:          false,
		QuarantineTargetBound:         false,
		NativeWriteLocked:             s.NativeWriteLockRow && s.NativeWriteRows == 0,
		RollbackTraceReady:            false,
		HumanReviewPresent:            false,
		BridgeQuarantineOnly:          s.QuarantineRows == len(s.Rows) && s.NativeWriteRows == 0,
		NativeWriteAuthorization:      false,
		RealSourceLoaded:              false,
		ObservedCorrelationLoaded:     false,
		ConstructiveMeasureLoaded:     false,
		PhysicalOSCertificateLoaded:   false,
		PhysicalWickMapLoaded:         false,
		PhysicalHamiltonianLoaded:     false,
		ComparatorExecutionPerformed:  false,
		Verdict:                       StatusComparatorAuthorizationBlockedPreflight,
		Reason:                        "Gate542 defines the manifest but imports no authorization document; comparator execution remains blocked in preflight.",
	}
}

func buildGuard(a AuthorizationState) Guard {
	return Guard{
		Executed:                    true,
		ComparatorCanRunInPreflight: false,
		AuthorizationCanWriteNative: false,
		RealSchwingerImported:       a.RealSourceLoaded,
		PhysicalSchwingerDerived:    false,
		OSPositivityProven:          false,
		WickRotationSelected:        false,
		PhysicalHilbertSelected:     false,
		HamiltonianDerived:          false,
		UnitaryDynamicsDerived:      false,
		GlobalCausalitySelected:     false,
		ArrowOfTimeSelected:         false,
		Verdict:                     StatusNativePromotionRejected,
		Reason:                      "An authorization manifest can only permit future bridge-quarantine comparison; it cannot write Schwinger/Wick/Hilbert/Hamiltonian claims into native ASHA law.",
	}
}

func buildFirewall(a AuthorizationState, g Guard) Firewall {
	return Firewall{
		Executed:                      true,
		RealSchwingerSourceImported:   a.RealSourceLoaded,
		ObservedCorrelationImported:   a.ObservedCorrelationLoaded,
		ConstructiveMeasureImported:   a.ConstructiveMeasureLoaded,
		PhysicalOSCertificateImported: a.PhysicalOSCertificateLoaded,
		PhysicalWickMapImported:       a.PhysicalWickMapLoaded,
		PhysicalHamiltonianImported:   a.PhysicalHamiltonianLoaded,
		ComparatorExecutionPerformed:  a.ComparatorExecutionPerformed || g.ComparatorCanRunInPreflight,
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
		Reason:                        "No source, comparator output, or native theorem write crosses the Gate542 authorization-manifest airlock.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"unchanged: finite Clifford geometry, triality/family structure, anomaly/stability/spectral-action structural law-space"},
		BridgeEntries:        []string{"Gate542 authorization manifest schema for future real-source Schwinger comparator staging", "bridge-quarantine-only comparator authorization boundary", "native-write lock row required before any future comparator run"},
		EnvironmentalEntries: []string{"actual non-synthetic Schwinger source", "operator intent artifact", "license/access grant", "physical Wick/Hamiltonian/OS certificates"},
		FailedRoutes:         []string{StatusFailedManifestNotSchwinger, StatusFailedManifestNotOSProof, StatusFailedManifestNotWick, StatusFailedManifestNotHilbert, StatusFailedManifestNotHamiltonian, StatusFailedManifestNotUnitary, StatusFailedManifestNotGlobal, StatusFailedManifestNotArrow, StatusFailedNoLiveComparatorInPreflight, StatusFailedNoSourceManifestImported},
		OpenTheorems:         []string{"Gate543 candidate: synthetic authorization manifest adapter dry run that parses all manifest rows, arms quarantine-only dry-run metadata, and still blocks native writes."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 543, Title: "Synthetic Comparator Authorization Manifest Adapter Dry Run", Reason: "Gate542 defines the authorization manifest. The next safe step is a synthetic manifest fixture that fills every row and proves authorization can arm only bridge quarantine, not native writes.", PrimaryTask: "Load a fake complete authorization manifest, verify rows/tags, confirm quarantine-only staging, and reject physical import/native promotion."}
}

func truth(a Analysis) string {
	return "Gate542 defines the real-source comparator authorization manifest as a bridge-quarantine-only airlock: future comparator execution needs explicit intent, authenticated source identity, provenance, license/access, checksum proof, Gate536 alignment, scope declaration, quarantine target, rollback trace, human review, and a native-write lock; no real source or physical dynamics is imported in preflight."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate541NegativeControlExecuted || !a.Inheritance.Gate541ChecksumVerified || !a.Inheritance.Gate541SwitchOffRejected || !a.Inheritance.Gate541NoIntentRejected || !a.Inheritance.Gate541InsufficientProvenance || !a.Inheritance.Gate541ComparatorBlocked || !a.Inheritance.Gate541NoNativeWrite || !a.Inheritance.Gate541RedirectsToGate542 {
		bad = append(bad, "Gate541 inheritance incomplete")
	}
	if !a.Schema.Executed || len(a.Schema.Rows) != 14 || a.Schema.RequiredRows != 14 || a.Schema.BridgeOnlyRows != 14 || a.Schema.ComparatorRows != 14 || a.Schema.QuarantineRows != 14 || a.Schema.NativeWriteRows != 0 || !a.Schema.OperatorIntentRow || !a.Schema.SourceIdentityRow || !a.Schema.AuthenticityLedgerRow || !a.Schema.Gate536AlignmentRow || !a.Schema.Gate540SwitchRow || !a.Schema.AccessGrantRow || !a.Schema.ChecksumProofRow || !a.Schema.ProvenanceReportRow || !a.Schema.ComparatorScopeRow || !a.Schema.QuarantineTargetRow || !a.Schema.ModeDeclarationRow || !a.Schema.NativeWriteLockRow || !a.Schema.RollbackTraceRow || !a.Schema.HumanReviewRow {
		bad = append(bad, "authorization manifest schema incomplete")
	}
	if !a.Authorization.Executed || a.Authorization.ManifestImported || a.Authorization.ComparatorLiveAuthorization || a.Authorization.ComparatorDryRunAuthorization || a.Authorization.ExplicitOperatorIntentPresent || a.Authorization.AuthenticatedSourceIdentity || a.Authorization.RealSourceLoaded || a.Authorization.ObservedCorrelationLoaded || a.Authorization.ConstructiveMeasureLoaded || a.Authorization.PhysicalOSCertificateLoaded || a.Authorization.PhysicalWickMapLoaded || a.Authorization.PhysicalHamiltonianLoaded || a.Authorization.ComparatorExecutionPerformed || !a.Authorization.NativeWriteLocked || !a.Authorization.BridgeQuarantineOnly || a.Authorization.NativeWriteAuthorization {
		bad = append(bad, "authorization state did not fail closed")
	}
	if !a.Guard.Executed || a.Guard.ComparatorCanRunInPreflight || a.Guard.AuthorizationCanWriteNative || a.Guard.RealSchwingerImported || a.Guard.PhysicalSchwingerDerived || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected || a.Guard.PhysicalHilbertSelected || a.Guard.HamiltonianDerived || a.Guard.UnitaryDynamicsDerived || a.Guard.GlobalCausalitySelected || a.Guard.ArrowOfTimeSelected {
		bad = append(bad, "guard leaked physical dynamics")
	}
	if !a.Firewall.Executed || a.Firewall.RealSchwingerSourceImported || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || a.Firewall.ComparatorExecutionPerformed || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked comparator/native write")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate542 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: executed=%t checksum=%t switch_rejected=%t intent_rejected=%t provenance_rejected=%t comparator_blocked=%t native_blocked=%t redirects=%t; %s", x.Verdict, x.Gate541NegativeControlExecuted, x.Gate541ChecksumVerified, x.Gate541SwitchOffRejected, x.Gate541NoIntentRejected, x.Gate541InsufficientProvenance, x.Gate541ComparatorBlocked, x.Gate541NoNativeWrite, x.Gate541RedirectsToGate542, x.Reason)
}
func FormatSchema(x ManifestSchema) string {
	names := []string{}
	for _, r := range x.Rows {
		names = append(names, r.Name)
	}
	return fmt.Sprintf("%s: rows=%d required=%d bridge=%d comparator=%d quarantine=%d native_write=%d names=%s; %s", x.Verdict, len(x.Rows), x.RequiredRows, x.BridgeOnlyRows, x.ComparatorRows, x.QuarantineRows, x.NativeWriteRows, strings.Join(names, ","), x.Reason)
}
func FormatAuthorization(x AuthorizationState) string {
	return fmt.Sprintf("%s: imported=%t live=%t dryrun=%t intent=%t source=%t license=%t checksum=%t provenance=%t quarantine_only=%t native_locked=%t comparator_executed=%t real_source=%t native_auth=%t; %s", x.Verdict, x.ManifestImported, x.ComparatorLiveAuthorization, x.ComparatorDryRunAuthorization, x.ExplicitOperatorIntentPresent, x.AuthenticatedSourceIdentity, x.LicenseAccessGrantBound, x.ChecksumProofBound, x.ProvenanceIntegrityBound, x.BridgeQuarantineOnly, x.NativeWriteLocked, x.ComparatorExecutionPerformed, x.RealSourceLoaded, x.NativeWriteAuthorization, x.Reason)
}
func FormatGuard(x Guard) string {
	return fmt.Sprintf("%s: comparator_preflight=%t native_auth=%t real=%t schwinger=%t os=%t wick=%t hilbert=%t hamiltonian=%t unitary=%t global=%t arrow=%t; %s", x.Verdict, x.ComparatorCanRunInPreflight, x.AuthorizationCanWriteNative, x.RealSchwingerImported, x.PhysicalSchwingerDerived, x.OSPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSelected, x.HamiltonianDerived, x.UnitaryDynamicsDerived, x.GlobalCausalitySelected, x.ArrowOfTimeSelected, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: real=%t observed=%t measure=%t os_cert=%t wick=%t ham=%t comparator=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t registry=%t; %s", x.Verdict, x.RealSchwingerSourceImported, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.ComparatorExecutionPerformed, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate541NegativeControlInherited, StatusAuthorizationManifestAirlockDefined, StatusAuthorizationManifestRowsEnumerated, StatusBridgeQuarantineOnlyAuthorization, StatusExplicitOperatorIntentManifestRequired, StatusNativeWriteLockManifestRequired, StatusComparatorAuthorizationBlockedPreflight, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedManifestNotSchwinger, StatusFailedManifestNotOSProof, StatusFailedManifestNotWick, StatusFailedManifestNotHilbert, StatusFailedManifestNotHamiltonian, StatusFailedManifestNotUnitary, StatusFailedManifestNotGlobal, StatusFailedManifestNotArrow, StatusFailedNoLiveComparatorInPreflight, StatusFailedNoSourceManifestImported, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
