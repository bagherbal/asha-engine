// Package generation2syntheticevidenceboardadapter implements Gate 550:
// Synthetic Evidence Board Adapter Dry Run.
//
// Gate 549 defined the physical-correlation evidence-board airlock. Gate 550
// loads a checksum-protected synthetic evidence-board manifest through that
// airlock, validates all 17 rows, verifies citation/uncertainty/reproducibility/
// revocation/versioning/downstream-policy/post-board-audit metadata and a
// zero-native-delta manifest, then rejects the fixture as real bridge evidence
// because the source chain and released evidence remain synthetic.
package generation2syntheticevidenceboardadapter

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

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2physicalcorrelationevidenceboardairlock"
)

const (
	AuditID       = "GATE550-SYNTHETIC-EVIDENCE-BOARD-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_evidence_board_manifest_gate550.json"

	StatusGate549EvidenceBoardInherited      = "CONDITIONAL_SUPPORT_GATE549_EVIDENCE_BOARD_AIRLOCK_INHERITED"
	StatusSyntheticEvidenceBoardLoaded       = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_MANIFEST_LOADED"
	StatusSyntheticEvidenceBoardExecuted     = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_ADAPTER_EXECUTED"
	StatusSyntheticEvidence17RowsAccepted    = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_17_SCHEMA_ROWS_ACCEPTED"
	StatusSyntheticEvidenceChecksumVerified  = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_CHECKSUM_VERIFIED"
	StatusSyntheticEvidenceMetadataEnforced  = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_METADATA_SIEVE_ENFORCED"
	StatusSyntheticEvidenceCitationParsed    = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_CITATION_SCOPE_PARSED"
	StatusSyntheticEvidenceUncertaintyParsed = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_UNCERTAINTY_METADATA_PARSED"
	StatusSyntheticEvidenceReproParsed       = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_REPRODUCIBILITY_METADATA_PARSED"
	StatusSyntheticEvidenceRevocationParsed  = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_REVOCATION_HOOKS_PARSED"
	StatusSyntheticEvidenceVersionParsed     = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_VERSIONED_INDEX_PARSED"
	StatusSyntheticEvidenceNativeDeltaZero   = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_NATIVE_DELTA_ZERO_VERIFIED"
	StatusSyntheticEvidenceBlocked           = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_BLOCKED_AS_REAL_BRIDGE_EVIDENCE"
	StatusNoBridgeEvidenceBoarded            = "CONDITIONAL_SUPPORT_NO_BRIDGE_EVIDENCE_BOARDED_IN_GATE550"
	StatusNoRealSourceImported               = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE550"
	StatusNativePromotionRejected            = "CONDITIONAL_SUPPORT_SYNTHETIC_EVIDENCE_BOARD_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing         = "FAILED_ROUTE_GATE550_SYNTHETIC_EVIDENCE_BOARD_MANIFEST_MISSING"
	StatusFailedRowsIncomplete        = "FAILED_ROUTE_GATE550_SYNTHETIC_EVIDENCE_BOARD_ROWS_INCOMPLETE"
	StatusFailedMetadataIncomplete    = "FAILED_ROUTE_GATE550_SYNTHETIC_EVIDENCE_BOARD_METADATA_INCOMPLETE"
	StatusFailedChecksumMismatch      = "FAILED_ROUTE_GATE550_SYNTHETIC_EVIDENCE_BOARD_CHECKSUM_MISMATCH"
	StatusFailedBoardLeaked           = "FAILED_ROUTE_GATE550_SYNTHETIC_EVIDENCE_BOARD_LEAKED_BRIDGE_EVIDENCE"
	StatusFailedSyntheticNotEvidence  = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_CANNOT_BOARD_SYNTHETIC_OUTPUT_AS_BRIDGE_EVIDENCE"
	StatusFailedSyntheticNotReal      = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_CANNOT_AUTHENTICATE_REAL_SOURCE_CHAIN"
	StatusFailedNotSchwinger          = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedNotOS                 = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedNotWick               = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedNotHilbert            = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedNotHamilton           = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedNotUnitary            = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedNotGlobal             = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedNotArrow              = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedOutputStillQuarantine = "FAILED_ROUTE_SYNTHETIC_EVIDENCE_BOARD_OUTPUT_REMAINS_QUARANTINED"
	StatusFirewallPreserved           = "FIREWALL_PRESERVED_GATE550_SYNTHETIC_EVIDENCE_BOARD_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked  = "FIREWALL_BLOCKED_GATE550_SYNTHETIC_EVIDENCE_BOARD_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate549AirlockDefined      bool
	Gate549SchemaRows          int
	Gate549CitationScope       bool
	Gate549Uncertainty         bool
	Gate549Reproducibility     bool
	Gate549Revocation          bool
	Gate549NativeDeltaRequired bool
	Gate549NoBoardEvidence     bool
	Gate549NativeWriteLocked   bool
	Gate549NoPhysicalClaims    bool
	Gate549RedirectsToGate550  bool

	Verdict string
	Reason  string
}

type EvidenceBoardRow struct {
	SchemaKey         string `json:"schema_key"`
	Source            string `json:"source"`
	SourceVersion     string `json:"source_version"`
	Convention        string `json:"convention"`
	ValueKind         string `json:"value_kind"`
	Value             string `json:"value"`
	BridgeOnly        bool   `json:"bridge_only"`
	EvidenceBoardOnly bool   `json:"evidence_board_only"`
	QuarantineOnly    bool   `json:"quarantine_only"`
	DryRunOnly        bool   `json:"dry_run_only"`
	Synthetic         bool   `json:"synthetic"`
	NoTheoremInput    bool   `json:"no_theorem_input"`
	NativePromotion   bool   `json:"native_promotion"`
	NativeWrite       bool   `json:"native_write"`
	PhysicalClaim     bool   `json:"physical_claim"`
	BridgeEvidence    bool   `json:"bridge_evidence_claim"`
	Observed          bool   `json:"observed"`
}

type Ledger struct {
	Gate                                    int                `json:"gate"`
	LedgerName                              string             `json:"ledger_name"`
	Description                             string             `json:"description"`
	Gate549EvidenceBoardReference           string             `json:"gate549_evidence_board_reference"`
	Gate548ClosureReference                 string             `json:"gate548_closure_reference"`
	BridgeOnly                              bool               `json:"bridge_only"`
	SyntheticFixture                        bool               `json:"synthetic_fixture"`
	EvidenceBoardManifestImported           bool               `json:"evidence_board_manifest_imported"`
	BoardEntryCandidate                     bool               `json:"board_entry_candidate"`
	CandidateEntries                        int                `json:"candidate_entries"`
	ReleasedBridgeEvidenceAvailable         bool               `json:"released_bridge_evidence_available"`
	AuthenticatedNonSyntheticSourceChain    bool               `json:"authenticated_non_synthetic_source_chain"`
	AuthenticatedNonSyntheticBridgeEvidence bool               `json:"authenticated_non_synthetic_bridge_evidence"`
	SyntheticReleasedEvidenceReference      bool               `json:"synthetic_released_evidence_reference"`
	CitationScopeParsed                     bool               `json:"citation_scope_parsed"`
	EnvironmentalClassificationParsed       bool               `json:"environmental_classification_parsed"`
	UncertaintyBudgetParsed                 bool               `json:"uncertainty_budget_parsed"`
	ResidualThresholdRecordParsed           bool               `json:"residual_threshold_record_parsed"`
	IndependentReproducibilityRecordParsed  bool               `json:"independent_reproducibility_record_parsed"`
	CertificateMapParsed                    bool               `json:"certificate_map_parsed"`
	NativeDeltaZero                         bool               `json:"native_delta_zero"`
	RevocationAndRollbackHooksParsed        bool               `json:"revocation_and_rollback_hooks_parsed"`
	VersionedEvidenceIndexParsed            bool               `json:"versioned_evidence_index_parsed"`
	HumanCurationAttestationParsed          bool               `json:"human_curation_attestation_parsed"`
	DownstreamUsagePolicyParsed             bool               `json:"downstream_usage_policy_parsed"`
	PostBoardAuditLogParsed                 bool               `json:"post_board_audit_log_parsed"`
	EvidenceBoardAcceptanceAllowed          bool               `json:"evidence_board_acceptance_allowed"`
	EvidenceEntriesAccepted                 int                `json:"evidence_entries_accepted"`
	BoardedAsBridgeEvidence                 bool               `json:"boarded_as_bridge_evidence"`
	RealSchwingerSourceImported             bool               `json:"real_schwinger_source_imported"`
	PhysicalSchwingerFunctionsLoaded        bool               `json:"physical_schwinger_functions_loaded"`
	PhysicalOSCertificateLoaded             bool               `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded                   bool               `json:"physical_wick_map_loaded"`
	PhysicalHilbertSpaceLoaded              bool               `json:"physical_hilbert_space_loaded"`
	PhysicalHamiltonianLoaded               bool               `json:"physical_hamiltonian_loaded"`
	UnitaryDynamicsLoaded                   bool               `json:"unitary_dynamics_loaded"`
	GlobalCausalityLoaded                   bool               `json:"global_causality_loaded"`
	TimeArrowLoaded                         bool               `json:"time_arrow_loaded"`
	NativeWriteLock                         bool               `json:"native_write_lock"`
	NativeWriteAuthorization                bool               `json:"native_write_authorization"`
	NativeRegistryWrite                     bool               `json:"native_registry_write"`
	Source                                  string             `json:"source"`
	SourceVersion                           string             `json:"source_version"`
	Convention                              string             `json:"convention"`
	CanonicalPayload                        map[string]any     `json:"canonical_payload"`
	CanonicalPayloadSHA256                  string             `json:"canonical_payload_sha256"`
	Rows                                    []EvidenceBoardRow `json:"rows"`
}

type Import struct {
	Executed bool
	Loaded   bool
	Path     string

	Rows             int
	AcceptedRows     int
	RejectedRows     int
	MissingRows      []string
	DuplicateRows    []string
	ChecksumVerified bool
	ChecksumExpected string
	ChecksumActual   string

	BridgeOnly       bool
	SyntheticFixture bool
	ManifestImported bool
	CandidateEntries int
	BoardCandidate   bool

	ReleasedEvidenceAvailable   bool
	AuthenticatedSourceChain    bool
	AuthenticatedBridgeEvidence bool
	SyntheticReleasedReference  bool
	AcceptanceAllowed           bool
	EntriesAccepted             int
	BoardedAsBridgeEvidence     bool
	NativeDeltaZero             bool
	NativeWriteLock             bool
	NativeWriteAuthorization    bool
	NativeRegistryWrite         bool

	AllBridgeOnly        bool
	AllEvidenceBoardOnly bool
	AllQuarantineOnly    bool
	AllDryRunOnly        bool
	AllSynthetic         bool
	AllNoTheorem         bool
	AllSourceTagged      bool
	AllConventionTagged  bool
	AnyNativePromotion   bool
	AnyNativeWrite       bool
	AnyPhysicalClaim     bool
	AnyBridgeEvidence    bool
	AnyObservedClaim     bool

	Verdict string
	Reason  string
}

type BoardResult struct {
	Executed bool

	ManifestParsed              bool
	CitationScopeParsed         bool
	EnvironmentalClassParsed    bool
	UncertaintyBudgetParsed     bool
	ResidualThresholdParsed     bool
	ReproducibilityParsed       bool
	CertificateMapParsed        bool
	NativeDeltaZero             bool
	RevocationHooksParsed       bool
	VersionedIndexParsed        bool
	HumanCurationParsed         bool
	DownstreamUsageParsed       bool
	PostBoardAuditParsed        bool
	SyntheticUnderlyingEvidence bool
	AuthenticatedSourceChain    bool
	AuthenticatedBridgeEvidence bool
	AcceptanceAllowed           bool
	EvidenceEntriesAccepted     int
	BoardedAsBridgeEvidence     bool
	NativeWriteLocked           bool
	NativeWriteAuthorization    bool
	NativeRegistryWrite         bool
	BlockedBecauseSynthetic     bool

	Verdict string
	Reason  string
}

type Firewall struct {
	Executed bool

	SyntheticBoardManifestPresent    bool
	BridgeEvidenceBoarded            bool
	RealSchwingerSourceImported      bool
	AuthenticatedRealSource          bool
	PhysicalSchwingerFunctionsLoaded bool
	OSPositivityCertificateLoaded    bool
	WickMapLoaded                    bool
	HilbertSpaceReconstructed        bool
	HamiltonianSpectrumLoaded        bool
	UnitaryDynamicsLoaded            bool
	GlobalCausalityLoaded            bool
	TimeArrowLoaded                  bool
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
	Import      Import
	Board       BoardResult
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
	g549, err := generation2physicalcorrelationevidenceboardairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate549 evidence-board airlock: %w", err)
	}
	ledger, resolved, err := loadLedger(path)
	a := Analysis{Inheritance: buildInheritance(g549)}
	if err != nil {
		a.Import = Import{Executed: true, Loaded: false, Path: path, Verdict: StatusFailedLedgerMissing, Reason: err.Error()}
		a.Board = BoardResult{Executed: true, Verdict: StatusFailedLedgerMissing, Reason: "No synthetic evidence-board ledger could be loaded."}
		a.Firewall = buildFirewall(a.Board, a.Import)
		a.Registry = buildRegistry()
		a.Next = buildNext(a)
		a.Truth = truth(a)
		return a, err
	}
	a.Import = buildImport(ledger, resolved)
	a.Board = buildBoard(a.Import, ledger)
	a.Firewall = buildFirewall(a.Board, a.Import)
	a.Registry = buildRegistry()
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2physicalcorrelationevidenceboardairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                   true,
		Gate549AirlockDefined:      g.Schema.Executed && g.Schema.RowCount == 17 && g.Schema.BridgeOnly && g.Schema.NativePromotionRejected,
		Gate549SchemaRows:          g.Schema.RowCount,
		Gate549CitationScope:       g.Schema.CitationScopeDefined,
		Gate549Uncertainty:         g.Schema.UncertaintyBudgetRequired,
		Gate549Reproducibility:     g.Schema.ReproducibilityRecordRequired,
		Gate549Revocation:          g.Schema.RevocationHooksRequired,
		Gate549NativeDeltaRequired: g.Schema.NativeDeltaZeroRequired,
		Gate549NoBoardEvidence:     !g.State.ReleasedBridgeEvidenceAvailable && !g.State.BoardedAsBridgeEvidence && g.State.EvidenceEntriesAccepted == 0,
		Gate549NativeWriteLocked:   g.State.NativeWriteLocked && !g.State.NativeWriteAuthorization && !g.State.NativeRegistryWrite,
		Gate549NoPhysicalClaims:    !g.Firewall.PhysicalSchwingerFunctionsLoaded && !g.Firewall.PhysicalOSCertificateLoaded && !g.Firewall.PhysicalWickMapLoaded && !g.Firewall.PhysicalHilbertSpaceLoaded && !g.Firewall.PhysicalHamiltonianLoaded,
		Gate549RedirectsToGate550:  g.Next.Gate == 550,
		Verdict:                    StatusGate549EvidenceBoardInherited,
		Reason:                     "Gate550 inherits Gate549's evidence-board airlock: 17 board rows are defined, citation and zero-native-delta rules are required, no bridge evidence is available, and native writes remain locked.",
	}
}

func requiredRows() []string {
	return []string{"evidence_board_identifier", "released_bridge_evidence_reference", "authenticated_source_chain_reference", "comparator_result_reference", "release_review_reference", "citation_scope_and_claim_boundaries", "environmental_classification", "uncertainty_budget", "residual_threshold_record", "independent_reproducibility_record", "certificate_map_os_wick_hilbert_hamiltonian", "native_delta_zero_manifest", "revocation_and_rollback_hooks", "versioned_evidence_index", "human_curation_attestation", "downstream_usage_policy", "post_board_audit_log"}
}

func buildImport(l Ledger, resolved string) Import {
	req := requiredRows()
	seen := map[string]int{}
	imp := Import{Executed: true, Loaded: true, Path: resolved, Rows: len(l.Rows), BridgeOnly: l.BridgeOnly, SyntheticFixture: l.SyntheticFixture, ManifestImported: l.EvidenceBoardManifestImported, CandidateEntries: l.CandidateEntries, BoardCandidate: l.BoardEntryCandidate, ReleasedEvidenceAvailable: l.ReleasedBridgeEvidenceAvailable, AuthenticatedSourceChain: l.AuthenticatedNonSyntheticSourceChain, AuthenticatedBridgeEvidence: l.AuthenticatedNonSyntheticBridgeEvidence, SyntheticReleasedReference: l.SyntheticReleasedEvidenceReference, AcceptanceAllowed: l.EvidenceBoardAcceptanceAllowed, EntriesAccepted: l.EvidenceEntriesAccepted, BoardedAsBridgeEvidence: l.BoardedAsBridgeEvidence, NativeDeltaZero: l.NativeDeltaZero, NativeWriteLock: l.NativeWriteLock, NativeWriteAuthorization: l.NativeWriteAuthorization, NativeRegistryWrite: l.NativeRegistryWrite, AllBridgeOnly: true, AllEvidenceBoardOnly: true, AllQuarantineOnly: true, AllDryRunOnly: true, AllSynthetic: true, AllNoTheorem: true, AllSourceTagged: true, AllConventionTagged: true, Verdict: StatusSyntheticEvidenceBoardExecuted, Reason: "Synthetic evidence-board manifest loaded, parsed, checksum-verified, and rejected as real bridge evidence because the source chain and released evidence remain synthetic."}
	for _, r := range l.Rows {
		seen[r.SchemaKey]++
		if r.Source == "" {
			imp.AllSourceTagged = false
		}
		if r.Convention == "" {
			imp.AllConventionTagged = false
		}
		if !r.BridgeOnly {
			imp.AllBridgeOnly = false
		}
		if !r.EvidenceBoardOnly {
			imp.AllEvidenceBoardOnly = false
		}
		if !r.QuarantineOnly {
			imp.AllQuarantineOnly = false
		}
		if !r.DryRunOnly {
			imp.AllDryRunOnly = false
		}
		if !r.Synthetic {
			imp.AllSynthetic = false
		}
		if !r.NoTheoremInput {
			imp.AllNoTheorem = false
		}
		if r.NativePromotion {
			imp.AnyNativePromotion = true
		}
		if r.NativeWrite {
			imp.AnyNativeWrite = true
		}
		if r.PhysicalClaim {
			imp.AnyPhysicalClaim = true
		}
		if r.BridgeEvidence {
			imp.AnyBridgeEvidence = true
		}
		if r.Observed {
			imp.AnyObservedClaim = true
		}
	}
	for _, k := range req {
		if seen[k] == 0 {
			imp.MissingRows = append(imp.MissingRows, k)
		}
		if seen[k] > 1 {
			imp.DuplicateRows = append(imp.DuplicateRows, k)
		}
	}
	sort.Strings(imp.MissingRows)
	sort.Strings(imp.DuplicateRows)
	imp.AcceptedRows = len(req) - len(imp.MissingRows)
	imp.RejectedRows = len(imp.MissingRows) + len(imp.DuplicateRows)
	imp.ChecksumExpected = l.CanonicalPayloadSHA256
	imp.ChecksumActual = checksum(l.CanonicalPayload)
	imp.ChecksumVerified = imp.ChecksumExpected != "" && imp.ChecksumExpected == imp.ChecksumActual
	if len(imp.MissingRows) > 0 || len(imp.DuplicateRows) > 0 || imp.AcceptedRows != 17 || len(l.Rows) != 17 {
		imp.Verdict = StatusFailedRowsIncomplete
		imp.Reason = "Synthetic evidence-board manifest does not exactly cover the 17 Gate549 rows."
	}
	if !imp.ChecksumVerified {
		imp.Verdict = StatusFailedChecksumMismatch
		imp.Reason = "Synthetic evidence-board manifest checksum does not match canonical payload."
	}
	if !imp.AllBridgeOnly || !imp.AllEvidenceBoardOnly || !imp.AllQuarantineOnly || !imp.AllDryRunOnly || !imp.AllSynthetic || !imp.AllNoTheorem || !imp.AllSourceTagged || !imp.AllConventionTagged || imp.AnyNativePromotion || imp.AnyNativeWrite || imp.AnyPhysicalClaim || imp.AnyBridgeEvidence || imp.AnyObservedClaim {
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = "Synthetic evidence-board manifest metadata sieve failed."
	}
	if imp.ReleasedEvidenceAvailable || imp.AuthenticatedSourceChain || imp.AuthenticatedBridgeEvidence || imp.AcceptanceAllowed || imp.EntriesAccepted != 0 || imp.BoardedAsBridgeEvidence || !imp.NativeDeltaZero || !imp.NativeWriteLock || imp.NativeWriteAuthorization || imp.NativeRegistryWrite {
		imp.Verdict = StatusFailedBoardLeaked
		imp.Reason = "Synthetic evidence-board manifest leaked bridge evidence or native write authority."
	}
	return imp
}

func buildBoard(imp Import, l Ledger) BoardResult {
	b := BoardResult{Executed: true,
		ManifestParsed:              imp.Loaded && imp.ManifestImported,
		CitationScopeParsed:         l.CitationScopeParsed,
		EnvironmentalClassParsed:    l.EnvironmentalClassificationParsed,
		UncertaintyBudgetParsed:     l.UncertaintyBudgetParsed,
		ResidualThresholdParsed:     l.ResidualThresholdRecordParsed,
		ReproducibilityParsed:       l.IndependentReproducibilityRecordParsed,
		CertificateMapParsed:        l.CertificateMapParsed,
		NativeDeltaZero:             l.NativeDeltaZero,
		RevocationHooksParsed:       l.RevocationAndRollbackHooksParsed,
		VersionedIndexParsed:        l.VersionedEvidenceIndexParsed,
		HumanCurationParsed:         l.HumanCurationAttestationParsed,
		DownstreamUsageParsed:       l.DownstreamUsagePolicyParsed,
		PostBoardAuditParsed:        l.PostBoardAuditLogParsed,
		SyntheticUnderlyingEvidence: l.SyntheticReleasedEvidenceReference,
		AuthenticatedSourceChain:    l.AuthenticatedNonSyntheticSourceChain,
		AuthenticatedBridgeEvidence: l.AuthenticatedNonSyntheticBridgeEvidence,
		AcceptanceAllowed:           l.EvidenceBoardAcceptanceAllowed,
		EvidenceEntriesAccepted:     l.EvidenceEntriesAccepted,
		BoardedAsBridgeEvidence:     l.BoardedAsBridgeEvidence,
		NativeWriteLocked:           l.NativeWriteLock,
		NativeWriteAuthorization:    l.NativeWriteAuthorization,
		NativeRegistryWrite:         l.NativeRegistryWrite,
		BlockedBecauseSynthetic:     l.SyntheticFixture && !l.AuthenticatedNonSyntheticBridgeEvidence && !l.ReleasedBridgeEvidenceAvailable,
		Verdict:                     StatusSyntheticEvidenceBlocked,
		Reason:                      "The synthetic board manifest validates evidence-board plumbing but is blocked from boarding because the release chain is synthetic and no authenticated non-synthetic bridge evidence exists.",
	}
	if !b.ManifestParsed || !b.CitationScopeParsed || !b.EnvironmentalClassParsed || !b.UncertaintyBudgetParsed || !b.ResidualThresholdParsed || !b.ReproducibilityParsed || !b.CertificateMapParsed || !b.NativeDeltaZero || !b.RevocationHooksParsed || !b.VersionedIndexParsed || !b.HumanCurationParsed || !b.DownstreamUsageParsed || !b.PostBoardAuditParsed {
		b.Verdict = StatusFailedMetadataIncomplete
		b.Reason = "Evidence-board metadata plumbing incomplete."
	}
	if !b.SyntheticUnderlyingEvidence || b.AuthenticatedSourceChain || b.AuthenticatedBridgeEvidence || b.AcceptanceAllowed || b.EvidenceEntriesAccepted != 0 || b.BoardedAsBridgeEvidence || !b.NativeWriteLocked || b.NativeWriteAuthorization || b.NativeRegistryWrite || !b.BlockedBecauseSynthetic {
		b.Verdict = StatusFailedBoardLeaked
		b.Reason = "Evidence-board guard leaked synthetic evidence into bridge evidence or native writes."
	}
	return b
}

func buildFirewall(b BoardResult, imp Import) Firewall {
	return Firewall{Executed: true,
		SyntheticBoardManifestPresent:    imp.Loaded,
		BridgeEvidenceBoarded:            b.BoardedAsBridgeEvidence,
		RealSchwingerSourceImported:      false,
		AuthenticatedRealSource:          false,
		PhysicalSchwingerFunctionsLoaded: false,
		OSPositivityCertificateLoaded:    false,
		WickMapLoaded:                    false,
		HilbertSpaceReconstructed:        false,
		HamiltonianSpectrumLoaded:        false,
		UnitaryDynamicsLoaded:            false,
		GlobalCausalityLoaded:            false,
		TimeArrowLoaded:                  false,
		NativeSchwingerFunctionWrite:     false,
		NativeOSPositivityWrite:          false,
		NativeWickWrite:                  false,
		NativeHilbertWrite:               false,
		NativeHamiltonianWrite:           false,
		NativeUnitaryDynamicsWrite:       false,
		NativeGlobalCausalWrite:          false,
		NativeTimeArrowWrite:             false,
		NativeRegistryWritten:            b.NativeRegistryWrite,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "Gate550 admits only a synthetic evidence-board parser fixture; no real source, bridge-evidence board entry, physical certificate, dynamics, or native registry write is admitted.",
	}
}

func buildRegistry() RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No native law is written at Gate550.", "The synthetic evidence-board adapter verifies zero-native-delta metadata and cannot mutate Cℓ(1,7), Schwinger, OS, Wick, Hilbert, Hamiltonian, unitary, causal, or time-arrow theorems."},
		BridgeEntries:        []string{"Gate550 validates the 17-row evidence-board parser using a checksum-protected synthetic fixture.", "Citation scope, environmental classification, uncertainty, residual thresholds, reproducibility, certificate mapping, revocation hooks, versioning, curation, downstream policy, post-board audit, and native-delta-zero plumbing are accepted only in bridge quarantine."},
		EnvironmentalEntries: []string{"The fixture remains synthetic environmental metadata, not released bridge evidence.", "Real source chains, real bridge evidence, physical certificates, reproducibility reports, uncertainty budgets, and human curation remain external obligations."},
		FailedRoutes:         []string{StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedNotSchwinger, StatusFailedNotOS, StatusFailedNotWick, StatusFailedNotHilbert, StatusFailedNotHamilton, StatusFailedNotUnitary, StatusFailedNotGlobal, StatusFailedNotArrow, StatusFailedOutputStillQuarantine, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"A future non-synthetic evidence-board adapter may admit only already released bridge evidence with authenticated source chain, reproducibility, curation, and native-delta-zero proof; it still cannot become native law."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 551, Title: "Physical Evidence Board Sector Closure Ledger", Reason: "Gate550 proves the synthetic evidence-board parser and rejection path. The next safe gate is a closure ledger for the whole physical-correlation evidence-board layer, freezing what can be cited as bridge-only evidence and what remains native-forbidden.", PrimaryTask: "Emit a closure/frontier map for Gates 536-550, preserving the rule that no physical correlation board entry, released bridge evidence, or native dynamics theorem exists without authenticated non-synthetic sources and review."}
}

func truth(a Analysis) string {
	return "Gate550 executes the synthetic evidence-board adapter: all 17 evidence-board rows parse, checksum and governance metadata pass, native-delta-zero is verified, but boarding remains blocked because the evidence chain is synthetic and unauthenticated as non-synthetic bridge evidence."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate549AirlockDefined || a.Inheritance.Gate549SchemaRows != 17 || !a.Inheritance.Gate549CitationScope || !a.Inheritance.Gate549Uncertainty || !a.Inheritance.Gate549Reproducibility || !a.Inheritance.Gate549Revocation || !a.Inheritance.Gate549NativeDeltaRequired || !a.Inheritance.Gate549NoBoardEvidence || !a.Inheritance.Gate549NativeWriteLocked || !a.Inheritance.Gate549NoPhysicalClaims || !a.Inheritance.Gate549RedirectsToGate550 {
		bad = append(bad, "Gate549 inheritance incomplete")
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 17 || len(a.Import.MissingRows) > 0 || a.Import.RejectedRows != 0 || len(a.Import.DuplicateRows) > 0 || !a.Import.ChecksumVerified {
		bad = append(bad, "evidence-board manifest schema/checksum incomplete")
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllEvidenceBoardOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || !a.Import.AllSourceTagged || !a.Import.AllConventionTagged || a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyBridgeEvidence || a.Import.AnyObservedClaim {
		bad = append(bad, "metadata sieve failed")
	}
	if !a.Board.ManifestParsed || !a.Board.CitationScopeParsed || !a.Board.EnvironmentalClassParsed || !a.Board.UncertaintyBudgetParsed || !a.Board.ResidualThresholdParsed || !a.Board.ReproducibilityParsed || !a.Board.CertificateMapParsed || !a.Board.NativeDeltaZero || !a.Board.RevocationHooksParsed || !a.Board.VersionedIndexParsed || !a.Board.HumanCurationParsed || !a.Board.DownstreamUsageParsed || !a.Board.PostBoardAuditParsed || !a.Board.SyntheticUnderlyingEvidence || a.Board.AuthenticatedSourceChain || a.Board.AuthenticatedBridgeEvidence || a.Board.AcceptanceAllowed || a.Board.EvidenceEntriesAccepted != 0 || a.Board.BoardedAsBridgeEvidence || !a.Board.NativeWriteLocked || a.Board.NativeWriteAuthorization || a.Board.NativeRegistryWrite || !a.Board.BlockedBecauseSynthetic {
		bad = append(bad, "board guard leaked")
	}
	if a.Firewall.BridgeEvidenceBoarded || a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate550 validation failed: %s", strings.Join(bad, "; "))
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
	return fmt.Sprintf("%s: rows=%d airlock=%t citation=%t uncertainty=%t reproducibility=%t revocation=%t native_delta=%t no_board=%t native_locked=%t no_physical=%t redirects=%t; %s", x.Verdict, x.Gate549SchemaRows, x.Gate549AirlockDefined, x.Gate549CitationScope, x.Gate549Uncertainty, x.Gate549Reproducibility, x.Gate549Revocation, x.Gate549NativeDeltaRequired, x.Gate549NoBoardEvidence, x.Gate549NativeWriteLocked, x.Gate549NoPhysicalClaims, x.Gate549RedirectsToGate550, x.Reason)
}

func FormatImport(x Import) string {
	return fmt.Sprintf("%s;%s;%s;%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%s duplicates=%s checksum=%t expected=%s actual=%s candidate=%t entries=%d released=%t auth_chain=%t auth_bridge=%t synthetic_ref=%t allowed=%t accepted_entries=%d boarded=%t native_delta=%t native_lock=%t native_auth=%t native_write=%t bridge=%t board_only=%t quarantine=%t dryrun=%t synthetic=%t no_theorem=%t; %s", StatusSyntheticEvidenceBoardLoaded, StatusSyntheticEvidence17RowsAccepted, StatusSyntheticEvidenceChecksumVerified, x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, strings.Join(x.MissingRows, ","), strings.Join(x.DuplicateRows, ","), x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.BoardCandidate, x.CandidateEntries, x.ReleasedEvidenceAvailable, x.AuthenticatedSourceChain, x.AuthenticatedBridgeEvidence, x.SyntheticReleasedReference, x.AcceptanceAllowed, x.EntriesAccepted, x.BoardedAsBridgeEvidence, x.NativeDeltaZero, x.NativeWriteLock, x.NativeWriteAuthorization, x.NativeRegistryWrite, x.AllBridgeOnly, x.AllEvidenceBoardOnly, x.AllQuarantineOnly, x.AllDryRunOnly, x.AllSynthetic, x.AllNoTheorem, x.Reason)
}

func FormatBoard(x BoardResult) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: manifest=%t citation=%t env=%t uncertainty=%t threshold=%t repro=%t cert=%t delta_zero=%t revocation=%t version=%t curation=%t downstream=%t post_audit=%t synthetic=%t auth_chain=%t auth_bridge=%t allowed=%t accepted=%d boarded=%t native_lock=%t native_auth=%t registry=%t blocked_synthetic=%t; %s", StatusSyntheticEvidenceMetadataEnforced, StatusSyntheticEvidenceCitationParsed, StatusSyntheticEvidenceUncertaintyParsed, StatusSyntheticEvidenceReproParsed, StatusSyntheticEvidenceRevocationParsed, StatusSyntheticEvidenceVersionParsed, StatusSyntheticEvidenceNativeDeltaZero, StatusSyntheticEvidenceBlocked, StatusNoBridgeEvidenceBoarded, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedOutputStillQuarantine, x.Verdict, x.ManifestParsed, x.CitationScopeParsed, x.EnvironmentalClassParsed, x.UncertaintyBudgetParsed, x.ResidualThresholdParsed, x.ReproducibilityParsed, x.CertificateMapParsed, x.NativeDeltaZero, x.RevocationHooksParsed, x.VersionedIndexParsed, x.HumanCurationParsed, x.DownstreamUsageParsed, x.PostBoardAuditParsed, x.SyntheticUnderlyingEvidence, x.AuthenticatedSourceChain, x.AuthenticatedBridgeEvidence, x.AcceptanceAllowed, x.EvidenceEntriesAccepted, x.BoardedAsBridgeEvidence, x.NativeWriteLocked, x.NativeWriteAuthorization, x.NativeRegistryWrite, x.BlockedBecauseSynthetic, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: manifest=%t boarded=%t real=%t auth_real=%t schwinger=%t os=%t wick=%t hilbert=%t ham=%t unitary=%t global=%t arrow=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t registry=%t; %s", StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedNotSchwinger, StatusFailedNotOS, StatusFailedNotWick, StatusFailedNotHilbert, StatusFailedNotHamilton, StatusFailedNotUnitary, StatusFailedNotGlobal, StatusFailedNotArrow, StatusFailedOutputStillQuarantine, StatusNativePromotionRejected, x.SyntheticBoardManifestPresent, x.BridgeEvidenceBoarded, x.RealSchwingerSourceImported, x.AuthenticatedRealSource, x.PhysicalSchwingerFunctionsLoaded, x.OSPositivityCertificateLoaded, x.WickMapLoaded, x.HilbertSpaceReconstructed, x.HamiltonianSpectrumLoaded, x.UnitaryDynamicsLoaded, x.GlobalCausalityLoaded, x.TimeArrowLoaded, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeRegistryWritten, x.Reason)
}

func Statuses() []string {
	return []string{StatusGate549EvidenceBoardInherited, StatusSyntheticEvidenceBoardLoaded, StatusSyntheticEvidenceBoardExecuted, StatusSyntheticEvidence17RowsAccepted, StatusSyntheticEvidenceChecksumVerified, StatusSyntheticEvidenceMetadataEnforced, StatusSyntheticEvidenceCitationParsed, StatusSyntheticEvidenceUncertaintyParsed, StatusSyntheticEvidenceReproParsed, StatusSyntheticEvidenceRevocationParsed, StatusSyntheticEvidenceVersionParsed, StatusSyntheticEvidenceNativeDeltaZero, StatusSyntheticEvidenceBlocked, StatusNoBridgeEvidenceBoarded, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedNotSchwinger, StatusFailedNotOS, StatusFailedNotWick, StatusFailedNotHilbert, StatusFailedNotHamilton, StatusFailedNotUnitary, StatusFailedNotGlobal, StatusFailedNotArrow, StatusFailedOutputStillQuarantine, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
