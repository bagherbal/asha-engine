// Package generation2physicalcorrelationevidenceboardairlock implements Gate 549:
// Physical Correlation Evidence Board Airlock.
//
// Gate 548 closed the physical-correlation import/release pipeline as a
// bridge-only frontier. Gate 549 defines the next non-promotional layer: an
// evidence board for organizing future released bridge evidence with citations,
// uncertainty, reproducibility, environmental classification, revocation hooks,
// and a mandatory zero-native-delta check. No bridge evidence is currently
// available, no board entry is admitted, and no native theorem is written.
package generation2physicalcorrelationevidenceboardairlock

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2physicalcorrelationreleaseclosureledger"
)

const (
	AuditID = "GATE549-PHYSICAL-CORRELATION-EVIDENCE-BOARD-AIRLOCK"

	StatusGate548ClosureInherited          = "CONDITIONAL_SUPPORT_GATE548_PHYSICAL_CORRELATION_CLOSURE_INHERITED"
	StatusEvidenceBoardAirlockDefined      = "CONDITIONAL_SUPPORT_PHYSICAL_CORRELATION_EVIDENCE_BOARD_AIRLOCK_DEFINED"
	StatusEvidenceBoardRowsEnumerated      = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_SCHEMA_ROWS_ENUMERATED"
	StatusCitationScopeDefined             = "CONDITIONAL_SUPPORT_BRIDGE_EVIDENCE_CITATION_SCOPE_SCHEMA_DEFINED"
	StatusUncertaintyReproducibilitySchema = "CONDITIONAL_SUPPORT_UNCERTAINTY_AND_REPRODUCIBILITY_BOARD_SCHEMA_DEFINED"
	StatusEnvironmentalClassification      = "CONDITIONAL_SUPPORT_ENVIRONMENTAL_CLASSIFICATION_SCHEMA_DEFINED"
	StatusRevocationRollbackHooks          = "CONDITIONAL_SUPPORT_REVOCATION_AND_ROLLBACK_HOOKS_DEFINED"
	StatusNativeDeltaZeroRequired          = "CONDITIONAL_SUPPORT_NATIVE_DELTA_ZERO_CHECK_REQUIRED"
	StatusEvidenceBoardReleaseBlocked      = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_RELEASE_BLOCKED_IN_PREFLIGHT"
	StatusNoBridgeEvidenceBoarded          = "CONDITIONAL_SUPPORT_NO_BRIDGE_EVIDENCE_BOARDED_IN_GATE549"
	StatusEvidenceBoardNativeRejected      = "CONDITIONAL_SUPPORT_EVIDENCE_BOARD_NATIVE_PROMOTION_REJECTED"

	StatusFailedBoardDoesNotDeriveSchwinger = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedBoardDoesNotProveOS         = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedBoardDoesNotGrantWick       = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedBoardDoesNotSelectHilbert   = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedBoardDoesNotDeriveHamilton  = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedBoardDoesNotGrantUnitary    = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedBoardDoesNotGrantGlobal     = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedBoardDoesNotSelectArrow     = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedNoReleasedEvidence          = "FAILED_ROUTE_NO_RELEASED_BRIDGE_EVIDENCE_AVAILABLE_IN_GATE549"
	StatusFailedBoardNotNativeLaw           = "FAILED_ROUTE_EVIDENCE_BOARD_SCHEMA_DOES_NOT_GRANT_NATIVE_LAW"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_GATE549_EVIDENCE_BOARD_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked        = "FIREWALL_BLOCKED_GATE549_EVIDENCE_BOARD_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate548ClosureEmitted      bool
	Gate548RowsClosed          int
	Gate548NativeFrozen        bool
	Gate548BridgeMapped        bool
	Gate548EnvironmentalMapped bool
	Gate548NoRealSource        bool
	Gate548NoBridgeEvidence    bool
	Gate548NativeWriteLocked   bool
	Gate548FirewallComplete    bool
	Gate548RedirectsToGate549  bool

	Verdict string
	Reason  string
}

type SchemaRow struct {
	Name        string
	Purpose     string
	Required    bool
	NativeWrite string
}

type EvidenceBoardSchema struct {
	Executed bool
	Rows     []SchemaRow

	RowCount                      int
	AllRowsRequired               bool
	CitationScopeDefined          bool
	UncertaintyBudgetRequired     bool
	ReproducibilityRecordRequired bool
	EnvironmentalClassRequired    bool
	RevocationHooksRequired       bool
	NativeDeltaZeroRequired       bool
	DownstreamUsagePolicyRequired bool
	PostBoardAuditRequired        bool
	BridgeOnly                    bool
	NativePromotionRejected       bool
	Verdict                       string
	Reason                        string
}

type BoardState struct {
	Executed bool

	ReleasedBridgeEvidenceAvailable bool
	EvidenceBoardManifestImported   bool
	EvidenceEntriesAccepted         int
	CitationScopeAccepted           bool
	UncertaintyAccepted             bool
	ReproducibilityAccepted         bool
	EnvironmentalClassAccepted      bool
	RevocationHooksAccepted         bool
	NativeDeltaZeroVerified         bool
	BoardReleased                   bool
	BoardedAsBridgeEvidence         bool
	NativeWriteLocked               bool
	NativeWriteAuthorization        bool
	NativeRegistryWrite             bool
	PrefightOnly                    bool

	Verdict  string
	Reason   string
	Failures []string
}

type Firewall struct {
	Executed bool

	PhysicalSchwingerFunctionsLoaded bool
	PhysicalOSCertificateLoaded      bool
	PhysicalWickMapLoaded            bool
	PhysicalHilbertSpaceLoaded       bool
	PhysicalHamiltonianLoaded        bool
	UnitaryDynamicsLoaded            bool
	GlobalCausalityLoaded            bool
	TimeArrowLoaded                  bool
	ReleasedBridgeEvidenceLoaded     bool
	EvidenceBoardEntryWritten        bool
	NativeSchwingerFunctionWrite     bool
	NativeOSPositivityWrite          bool
	NativeWickWrite                  bool
	NativeHilbertWrite               bool
	NativeHamiltonianWrite           bool
	NativeUnitaryDynamicsWrite       bool
	NativeGlobalCausalWrite          bool
	NativeTimeArrowWrite             bool
	NativeRegistryWritten            bool

	Verdict string
	Reason  string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      EvidenceBoardSchema
	State       BoardState
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
	g548, err := generation2physicalcorrelationreleaseclosureledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate548 physical-correlation closure ledger: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g548)}
	a.Schema = buildSchema()
	a.State = buildState(a.Inheritance, a.Schema)
	a.Firewall = buildFirewall(a.State)
	a.Registry = buildRegistry()
	a.Next = buildNext()
	a.Truth = truth()
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2physicalcorrelationreleaseclosureledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                   true,
		Gate548ClosureEmitted:      g.Closure.Executed && g.Closure.RowCount == 12,
		Gate548RowsClosed:          g.Closure.RowCount,
		Gate548NativeFrozen:        g.Closure.NativeFrontierFrozen,
		Gate548BridgeMapped:        g.Closure.BridgeFrontierMapped,
		Gate548EnvironmentalMapped: g.Closure.EnvironmentalMapped,
		Gate548NoRealSource:        !g.Guard.AuthenticatedNonSyntheticSource && !g.Firewall.PhysicalSchwingerFunctionsLoaded,
		Gate548NoBridgeEvidence:    !g.Guard.BridgeEvidenceReleased && !g.Firewall.ReleasedBridgeEvidence,
		Gate548NativeWriteLocked:   g.Guard.NativeWriteLocked && !g.Guard.NativeWriteAuthorization && !g.Firewall.NativeRegistryWritten,
		Gate548FirewallComplete:    g.Closure.SchwingerBlockClosed && g.Closure.AuthenticityClosed && g.Closure.ImportSwitchClosed && g.Closure.ComparatorClosed && g.Closure.ReleaseClosed,
		Gate548RedirectsToGate549:  g.Next.Gate == 549,
		Verdict:                    StatusGate548ClosureInherited,
		Reason:                     "Gate549 inherits Gate548's closure ledger: Gates 536-547 are mapped as bridge-only, no authenticated non-synthetic source exists, no bridge evidence has been released, and native writes remain locked.",
	}
}

func buildSchema() EvidenceBoardSchema {
	rows := []SchemaRow{
		{Name: "evidence_board_identifier", Purpose: "stable identifier for a future bridge-evidence board", Required: true, NativeWrite: "forbidden"},
		{Name: "released_bridge_evidence_reference", Purpose: "reference to a Gate546/547-style released bridge-evidence object", Required: true, NativeWrite: "forbidden"},
		{Name: "authenticated_source_chain_reference", Purpose: "links evidence back to authenticated non-synthetic source-chain metadata", Required: true, NativeWrite: "forbidden"},
		{Name: "comparator_result_reference", Purpose: "links the board row to a quarantined comparator result and checksum", Required: true, NativeWrite: "forbidden"},
		{Name: "release_review_reference", Purpose: "records human review, release authorization, and release scope", Required: true, NativeWrite: "forbidden"},
		{Name: "citation_scope_and_claim_boundaries", Purpose: "states exactly what may be cited as bridge evidence and what may not be claimed", Required: true, NativeWrite: "forbidden"},
		{Name: "environmental_classification", Purpose: "classifies the evidence as environmental/source data rather than native law", Required: true, NativeWrite: "forbidden"},
		{Name: "uncertainty_budget", Purpose: "records residuals, tolerances, uncertainty intervals, and invalidity domains", Required: true, NativeWrite: "forbidden"},
		{Name: "residual_threshold_record", Purpose: "stores the comparator threshold policy used for evidence acceptance", Required: true, NativeWrite: "forbidden"},
		{Name: "independent_reproducibility_record", Purpose: "requires independent rerun or reproducibility metadata", Required: true, NativeWrite: "forbidden"},
		{Name: "certificate_map_os_wick_hilbert_hamiltonian", Purpose: "maps which OS/Wick/Hilbert/Hamiltonian certificates are attached and scoped", Required: true, NativeWrite: "forbidden"},
		{Name: "native_delta_zero_manifest", Purpose: "proves the evidence board changes no native theorem registry entry", Required: true, NativeWrite: "forbidden"},
		{Name: "revocation_and_rollback_hooks", Purpose: "defines how evidence is withdrawn if source, checksum, license, or reproducibility fails", Required: true, NativeWrite: "forbidden"},
		{Name: "versioned_evidence_index", Purpose: "keeps board entries versioned, auditable, and reproducible", Required: true, NativeWrite: "forbidden"},
		{Name: "human_curation_attestation", Purpose: "records final curator review without granting native law", Required: true, NativeWrite: "forbidden"},
		{Name: "downstream_usage_policy", Purpose: "prevents downstream native promotion or overclaiming", Required: true, NativeWrite: "forbidden"},
		{Name: "post_board_audit_log", Purpose: "records post-board checks, revocation state, and citation history", Required: true, NativeWrite: "forbidden"},
	}
	s := EvidenceBoardSchema{Executed: true, Rows: rows, RowCount: len(rows), AllRowsRequired: true, CitationScopeDefined: true, UncertaintyBudgetRequired: true, ReproducibilityRecordRequired: true, EnvironmentalClassRequired: true, RevocationHooksRequired: true, NativeDeltaZeroRequired: true, DownstreamUsagePolicyRequired: true, PostBoardAuditRequired: true, BridgeOnly: true, NativePromotionRejected: true, Verdict: StatusEvidenceBoardAirlockDefined, Reason: "A future evidence board must organize only released bridge evidence, with citations, uncertainty, reproducibility, environmental classification, revocation hooks, usage policy, post-board audit, and a mandatory zero-native-delta manifest."}
	for _, r := range rows {
		if !r.Required || r.NativeWrite != "forbidden" {
			s.AllRowsRequired = false
			s.NativePromotionRejected = false
		}
	}
	return s
}

func buildState(i Inheritance, s EvidenceBoardSchema) BoardState {
	st := BoardState{Executed: true,
		ReleasedBridgeEvidenceAvailable: false,
		EvidenceBoardManifestImported:   false,
		EvidenceEntriesAccepted:         0,
		CitationScopeAccepted:           false,
		UncertaintyAccepted:             false,
		ReproducibilityAccepted:         false,
		EnvironmentalClassAccepted:      false,
		RevocationHooksAccepted:         false,
		NativeDeltaZeroVerified:         false,
		BoardReleased:                   false,
		BoardedAsBridgeEvidence:         false,
		NativeWriteLocked:               true,
		NativeWriteAuthorization:        false,
		NativeRegistryWrite:             false,
		PrefightOnly:                    true,
		Verdict:                         StatusNoBridgeEvidenceBoarded,
		Reason:                          "The evidence-board schema is defined, but no released bridge evidence exists; therefore no evidence-board manifest is imported, no board row is accepted, and native writes remain locked.",
	}
	if !i.Gate548ClosureEmitted || i.Gate548RowsClosed != 12 || !i.Gate548NoRealSource || !i.Gate548NoBridgeEvidence || !i.Gate548NativeWriteLocked || !i.Gate548RedirectsToGate549 {
		st.Failures = append(st.Failures, "Gate548 inheritance incomplete")
	}
	if !s.Executed || s.RowCount != 17 || !s.AllRowsRequired || !s.NativeDeltaZeroRequired || !s.NativePromotionRejected || !s.BridgeOnly {
		st.Failures = append(st.Failures, "evidence-board schema incomplete")
	}
	if st.ReleasedBridgeEvidenceAvailable || st.EvidenceBoardManifestImported || st.EvidenceEntriesAccepted != 0 || st.BoardedAsBridgeEvidence || st.NativeRegistryWrite || !st.NativeWriteLocked {
		st.Failures = append(st.Failures, "evidence-board state leaked")
	}
	if len(st.Failures) > 0 {
		st.Verdict = strings.Join(st.Failures, ";")
		st.Reason = "Gate549 evidence-board airlock failed."
	}
	return st
}

func buildFirewall(st BoardState) Firewall {
	return Firewall{Executed: true,
		PhysicalSchwingerFunctionsLoaded: false,
		PhysicalOSCertificateLoaded:      false,
		PhysicalWickMapLoaded:            false,
		PhysicalHilbertSpaceLoaded:       false,
		PhysicalHamiltonianLoaded:        false,
		UnitaryDynamicsLoaded:            false,
		GlobalCausalityLoaded:            false,
		TimeArrowLoaded:                  false,
		ReleasedBridgeEvidenceLoaded:     st.ReleasedBridgeEvidenceAvailable,
		EvidenceBoardEntryWritten:        st.BoardedAsBridgeEvidence,
		NativeSchwingerFunctionWrite:     false,
		NativeOSPositivityWrite:          false,
		NativeWickWrite:                  false,
		NativeHilbertWrite:               false,
		NativeHamiltonianWrite:           false,
		NativeUnitaryDynamicsWrite:       false,
		NativeGlobalCausalWrite:          false,
		NativeTimeArrowWrite:             false,
		NativeRegistryWritten:            st.NativeRegistryWrite,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "The evidence board is a citation/organization airlock only; no released evidence, physical certificates, dynamics, or native registry writes are admitted.",
	}
}

func buildRegistry() RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native law is written at Gate549.",
			"The native registry remains restricted to previously proved ASHA structural law; evidence-board rows can never mutate native Cℓ(1,7), OS, Wick, Hilbert, Hamiltonian, unitary, causal, or time-arrow theorems.",
		},
		BridgeEntries: []string{
			"Gate549 defines a physical-correlation evidence-board airlock for future released bridge evidence.",
			"The board schema requires citation scope, source-chain linkage, comparator/release references, uncertainty, reproducibility, environmental classification, certificate maps, revocation hooks, downstream usage policy, post-board audit, and native-delta-zero proof.",
		},
		EnvironmentalEntries: []string{
			"Any future board entry remains environmental/source evidence unless a separate theorem proves a native result.",
			"Actual correlation data, source authenticity, OS/Wick/Hilbert/Hamiltonian certificates, uncertainty budgets, reproducibility reports, and human curation remain source-side obligations.",
		},
		FailedRoutes: []string{
			StatusFailedBoardDoesNotDeriveSchwinger,
			StatusFailedBoardDoesNotProveOS,
			StatusFailedBoardDoesNotGrantWick,
			StatusFailedBoardDoesNotSelectHilbert,
			StatusFailedBoardDoesNotDeriveHamilton,
			StatusFailedBoardDoesNotGrantUnitary,
			StatusFailedBoardDoesNotGrantGlobal,
			StatusFailedBoardDoesNotSelectArrow,
			StatusFailedNoReleasedEvidence,
			StatusFailedBoardNotNativeLaw,
			StatusFirewallPreserved,
			StatusFirewallNativeWriteBlocked,
		},
		OpenTheorems: []string{
			"A future non-synthetic evidence-board adapter may only admit already released bridge evidence and must prove native-delta-zero; it still cannot promote evidence into native ASHA law.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 550, Title: "Synthetic Evidence Board Adapter Dry Run", Reason: "Gate549 defines the evidence-board airlock. The next safe step is a synthetic board fixture that verifies citation, uncertainty, reproducibility, revocation, and native-delta-zero plumbing while rejecting synthetic evidence as real bridge evidence.", PrimaryTask: "Load a synthetic evidence-board manifest, verify all 17 rows and zero-native-delta metadata, then block release because no authenticated non-synthetic bridge evidence exists."}
}

func truth() string {
	return "Gate549 defines the physical-correlation evidence-board airlock: future released bridge evidence can be organized, cited, versioned, scoped, revoked, and audited, but no evidence is currently available and the native registry remains unchanged."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate548ClosureEmitted || a.Inheritance.Gate548RowsClosed != 12 || !a.Inheritance.Gate548NativeFrozen || !a.Inheritance.Gate548BridgeMapped || !a.Inheritance.Gate548EnvironmentalMapped || !a.Inheritance.Gate548NoRealSource || !a.Inheritance.Gate548NoBridgeEvidence || !a.Inheritance.Gate548NativeWriteLocked || !a.Inheritance.Gate548FirewallComplete || !a.Inheritance.Gate548RedirectsToGate549 {
		bad = append(bad, "Gate548 inheritance incomplete")
	}
	if !a.Schema.Executed || a.Schema.RowCount != 17 || !a.Schema.AllRowsRequired || !a.Schema.CitationScopeDefined || !a.Schema.UncertaintyBudgetRequired || !a.Schema.ReproducibilityRecordRequired || !a.Schema.EnvironmentalClassRequired || !a.Schema.RevocationHooksRequired || !a.Schema.NativeDeltaZeroRequired || !a.Schema.DownstreamUsagePolicyRequired || !a.Schema.PostBoardAuditRequired || !a.Schema.BridgeOnly || !a.Schema.NativePromotionRejected {
		bad = append(bad, "evidence-board schema incomplete")
	}
	if a.State.ReleasedBridgeEvidenceAvailable || a.State.EvidenceBoardManifestImported || a.State.EvidenceEntriesAccepted != 0 || a.State.CitationScopeAccepted || a.State.UncertaintyAccepted || a.State.ReproducibilityAccepted || a.State.EnvironmentalClassAccepted || a.State.RevocationHooksAccepted || a.State.NativeDeltaZeroVerified || a.State.BoardReleased || a.State.BoardedAsBridgeEvidence || !a.State.NativeWriteLocked || a.State.NativeWriteAuthorization || a.State.NativeRegistryWrite || !a.State.PrefightOnly {
		bad = append(bad, "evidence-board state leaked")
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidenceLoaded || a.Firewall.EvidenceBoardEntryWritten || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate549 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate548ClosureInherited,
		StatusEvidenceBoardAirlockDefined,
		StatusEvidenceBoardRowsEnumerated,
		StatusCitationScopeDefined,
		StatusUncertaintyReproducibilitySchema,
		StatusEnvironmentalClassification,
		StatusRevocationRollbackHooks,
		StatusNativeDeltaZeroRequired,
		StatusEvidenceBoardReleaseBlocked,
		StatusNoBridgeEvidenceBoarded,
		StatusEvidenceBoardNativeRejected,
		StatusFailedBoardDoesNotDeriveSchwinger,
		StatusFailedBoardDoesNotProveOS,
		StatusFailedBoardDoesNotGrantWick,
		StatusFailedBoardDoesNotSelectHilbert,
		StatusFailedBoardDoesNotDeriveHamilton,
		StatusFailedBoardDoesNotGrantUnitary,
		StatusFailedBoardDoesNotGrantGlobal,
		StatusFailedBoardDoesNotSelectArrow,
		StatusFailedNoReleasedEvidence,
		StatusFailedBoardNotNativeLaw,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("gate548_rows=%d native_frozen=%t bridge_mapped=%t environmental_mapped=%t no_real_source=%t no_bridge_evidence=%t native_locked=%t redirects_gate549=%t",
		i.Gate548RowsClosed, i.Gate548NativeFrozen, i.Gate548BridgeMapped, i.Gate548EnvironmentalMapped, i.Gate548NoRealSource, i.Gate548NoBridgeEvidence, i.Gate548NativeWriteLocked, i.Gate548RedirectsToGate549)
}

func FormatSchema(s EvidenceBoardSchema) string {
	return fmt.Sprintf("rows=%d all_required=%t citation=%t uncertainty=%t reproducibility=%t environmental=%t revocation=%t native_delta_zero=%t bridge_only=%t native_rejected=%t",
		s.RowCount, s.AllRowsRequired, s.CitationScopeDefined, s.UncertaintyBudgetRequired, s.ReproducibilityRecordRequired, s.EnvironmentalClassRequired, s.RevocationHooksRequired, s.NativeDeltaZeroRequired, s.BridgeOnly, s.NativePromotionRejected)
}

func FormatState(st BoardState) string {
	return fmt.Sprintf("released_evidence=%t manifest=%t entries=%d citation=%t uncertainty=%t reproducibility=%t environmental=%t revocation=%t native_delta_zero=%t board_released=%t native_locked=%t native_write=%t preflight=%t failures=%d",
		st.ReleasedBridgeEvidenceAvailable, st.EvidenceBoardManifestImported, st.EvidenceEntriesAccepted, st.CitationScopeAccepted, st.UncertaintyAccepted, st.ReproducibilityAccepted, st.EnvironmentalClassAccepted, st.RevocationHooksAccepted, st.NativeDeltaZeroVerified, st.BoardReleased, st.NativeWriteLocked, st.NativeRegistryWrite, st.PrefightOnly, len(st.Failures))
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("schwinger=%t os=%t wick=%t hilbert=%t hamiltonian=%t unitary=%t global=%t arrow=%t bridge_evidence=%t board_entry=%t native_registry=%t",
		f.PhysicalSchwingerFunctionsLoaded, f.PhysicalOSCertificateLoaded, f.PhysicalWickMapLoaded, f.PhysicalHilbertSpaceLoaded, f.PhysicalHamiltonianLoaded, f.UnitaryDynamicsLoaded, f.GlobalCausalityLoaded, f.TimeArrowLoaded, f.ReleasedBridgeEvidenceLoaded, f.EvidenceBoardEntryWritten, f.NativeRegistryWritten)
}
