// Package generation2syntheticreleasereviewmanifestadapter implements Gate 547:
// Synthetic Release-Review Manifest Adapter Dry Run.
//
// Gate 546 defined the release-review airlock required before a quarantined
// comparator output can be cited as bridge evidence. Gate 547 loads a synthetic
// release-review manifest through that airlock, validates the 15-row parser,
// review/reproducibility/source-chain/citation metadata, checksum, rollback,
// and zero-native-write manifest, then rejects release because the underlying
// output remains synthetic and unauthenticated as physical bridge evidence.
package generation2syntheticreleasereviewmanifestadapter

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

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2comparatoroutputreleaseairlock"
)

const (
	AuditID       = "GATE547-SYNTHETIC-RELEASE-REVIEW-MANIFEST-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_release_review_manifest_gate547.json"

	StatusGate546ReleaseAirlockInherited        = "CONDITIONAL_SUPPORT_GATE546_RELEASE_AIRLOCK_INHERITED"
	StatusSyntheticReleaseManifestLoaded        = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_MANIFEST_LOADED"
	StatusSyntheticReleaseAdapterExecuted       = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_MANIFEST_ADAPTER_EXECUTED"
	StatusSyntheticRelease15RowsAccepted        = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_15_SCHEMA_ROWS_ACCEPTED"
	StatusSyntheticReleaseChecksumVerified      = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_CHECKSUM_VERIFIED"
	StatusSyntheticReleaseMetadataEnforced      = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_METADATA_SIEVE_ENFORCED"
	StatusSyntheticReleaseHumanReviewParsed     = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_HUMAN_REVIEW_METADATA_PARSED"
	StatusSyntheticReleaseReproducibilityParsed = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_REPRODUCIBILITY_METADATA_PARSED"
	StatusSyntheticReleaseSourceChainParsed     = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_SOURCE_CHAIN_METADATA_PARSED"
	StatusSyntheticReleaseCitationScopeParsed   = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_CITATION_SCOPE_PARSED"
	StatusSyntheticReleaseNativeDeltaZero       = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_NATIVE_WRITE_DELTA_ZERO_VERIFIED"
	StatusSyntheticReleaseBlocked               = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_BLOCKED_FOR_SYNTHETIC_OUTPUT"
	StatusNoBridgeEvidenceReleased              = "CONDITIONAL_SUPPORT_NO_BRIDGE_EVIDENCE_RELEASED_IN_GATE547"
	StatusNoRealSourceImported                  = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE547"
	StatusNativePromotionRejected               = "CONDITIONAL_SUPPORT_SYNTHETIC_RELEASE_REVIEW_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing          = "FAILED_ROUTE_GATE547_SYNTHETIC_RELEASE_REVIEW_MANIFEST_MISSING"
	StatusFailedRowsIncomplete         = "FAILED_ROUTE_GATE547_SYNTHETIC_RELEASE_REVIEW_ROWS_INCOMPLETE"
	StatusFailedMetadataIncomplete     = "FAILED_ROUTE_GATE547_SYNTHETIC_RELEASE_REVIEW_METADATA_INCOMPLETE"
	StatusFailedChecksumMismatch       = "FAILED_ROUTE_GATE547_SYNTHETIC_RELEASE_REVIEW_CHECKSUM_MISMATCH"
	StatusFailedReleaseLeaked          = "FAILED_ROUTE_GATE547_SYNTHETIC_RELEASE_REVIEW_LEAKED_BRIDGE_EVIDENCE"
	StatusFailedSyntheticNotEvidence   = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_CANNOT_RELEASE_SYNTHETIC_OUTPUT_AS_BRIDGE_EVIDENCE"
	StatusFailedSyntheticNotReal       = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_CANNOT_AUTHENTICATE_REAL_SOURCE"
	StatusFailedSyntheticNotSchwinger  = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticNotOS         = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedSyntheticNotWick       = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticNotHilbert    = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticNotHamilton   = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticNotUnitary    = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticNotGlobal     = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticNotArrow      = "FAILED_ROUTE_SYNTHETIC_RELEASE_MANIFEST_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFailedOutputStillQuarantined = "FAILED_ROUTE_SYNTHETIC_RELEASE_REVIEW_OUTPUT_REMAINS_QUARANTINED"
	StatusFirewallPreserved            = "FIREWALL_PRESERVED_GATE547_SYNTHETIC_RELEASE_REVIEW_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked   = "FIREWALL_BLOCKED_GATE547_SYNTHETIC_RELEASE_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate546AirlockDefined     bool
	Gate546SchemaRows         int
	Gate546QuarantinedPresent bool
	Gate546ReleaseBlocked     bool
	Gate546NoBridgeEvidence   bool
	Gate546NativeWriteLocked  bool
	Gate546AbortSynthetic     bool
	Gate546NoRealSource       bool
	Gate546NoPhysicalClaims   bool
	Gate546RedirectsToGate547 bool

	Verdict string
	Reason  string
}

type ReleaseManifestRow struct {
	SchemaKey       string `json:"schema_key"`
	Source          string `json:"source"`
	SourceVersion   string `json:"source_version"`
	Convention      string `json:"convention"`
	ValueKind       string `json:"value_kind"`
	Value           string `json:"value"`
	BridgeOnly      bool   `json:"bridge_only"`
	ReleaseOnly     bool   `json:"release_only"`
	QuarantineOnly  bool   `json:"quarantine_only"`
	DryRunOnly      bool   `json:"dry_run_only"`
	Synthetic       bool   `json:"synthetic"`
	NoTheoremInput  bool   `json:"no_theorem_input"`
	NativePromotion bool   `json:"native_promotion"`
	NativeWrite     bool   `json:"native_write"`
	PhysicalClaim   bool   `json:"physical_claim"`
	BridgeEvidence  bool   `json:"bridge_evidence_claim"`
	Observed        bool   `json:"observed"`
}

type Ledger struct {
	Gate                             int                  `json:"gate"`
	LedgerName                       string               `json:"ledger_name"`
	Description                      string               `json:"description"`
	Gate546ReleaseAirlockReference   string               `json:"gate546_release_airlock_reference"`
	Gate545QuarantineResultReference string               `json:"gate545_quarantine_result_reference"`
	BridgeOnly                       bool                 `json:"bridge_only"`
	SyntheticFixture                 bool                 `json:"synthetic_fixture"`
	ReleaseManifestImported          bool                 `json:"release_manifest_imported"`
	OperatorReleaseIntent            bool                 `json:"operator_release_intent"`
	HumanReviewAttestation           bool                 `json:"human_review_attestation"`
	IndependentReproducibilityReport bool                 `json:"independent_reproducibility_report"`
	AuthenticatedSourceChain         bool                 `json:"authenticated_source_chain"`
	ResidualThresholdAccepted        bool                 `json:"residual_threshold_accepted"`
	PhysicalClaimDiscriminator       bool                 `json:"physical_claim_discriminator"`
	BridgeEvidenceCitationScope      string               `json:"bridge_evidence_citation_scope"`
	ReleaseTargetQuarantineOnly      bool                 `json:"release_target_quarantine_only"`
	BridgeEvidenceReleaseAllowed     bool                 `json:"bridge_evidence_release_allowed"`
	BridgeEvidenceReleased           bool                 `json:"bridge_evidence_released"`
	RealSchwingerSourceImported      bool                 `json:"real_schwinger_source_imported"`
	AuthenticatedNonSyntheticSource  bool                 `json:"authenticated_non_synthetic_source"`
	PhysicalSchwingerFunctionsLoaded bool                 `json:"physical_schwinger_functions_loaded"`
	PhysicalOSCertificateLoaded      bool                 `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded            bool                 `json:"physical_wick_map_loaded"`
	PhysicalHilbertSpaceLoaded       bool                 `json:"physical_hilbert_space_loaded"`
	PhysicalHamiltonianLoaded        bool                 `json:"physical_hamiltonian_loaded"`
	UnitaryDynamicsLoaded            bool                 `json:"unitary_dynamics_loaded"`
	GlobalCausalityLoaded            bool                 `json:"global_causality_loaded"`
	TimeArrowLoaded                  bool                 `json:"time_arrow_loaded"`
	NativeWriteLock                  bool                 `json:"native_write_lock"`
	NativeWriteDeltaZero             bool                 `json:"native_write_delta_zero"`
	NativeRegistryWrite              bool                 `json:"native_registry_write"`
	RollbackAndRevocationPlan        string               `json:"rollback_and_revocation_plan"`
	PostReleaseAuditLog              string               `json:"post_release_audit_log"`
	Source                           string               `json:"source"`
	SourceVersion                    string               `json:"source_version"`
	Convention                       string               `json:"convention"`
	CanonicalPayload                 map[string]any       `json:"canonical_payload"`
	CanonicalPayloadSHA256           string               `json:"canonical_payload_sha256"`
	Rows                             []ReleaseManifestRow `json:"rows"`
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
	ManifestImported    bool
	OperatorIntent      bool
	HumanReview         bool
	Reproducibility     bool
	SourceChain         bool
	ResidualThreshold   bool
	Discriminator       bool
	CitationScope       bool
	QuarantineOnly      bool
	ReleaseAllowed      bool
	Released            bool
	RealSource          bool
	AuthenticatedReal   bool
	SchwingerLoaded     bool
	OSCertLoaded        bool
	WickMapLoaded       bool
	HilbertLoaded       bool
	HamiltonianLoaded   bool
	UnitaryLoaded       bool
	GlobalLoaded        bool
	ArrowLoaded         bool
	NativeWriteLock     bool
	NativeDeltaZero     bool
	NativeWrite         bool
	RollbackPlan        bool
	PostReleaseAudit    bool
	AllBridgeOnly       bool
	AllReleaseOnly      bool
	AllQuarantineOnly   bool
	AllDryRunOnly       bool
	AllSynthetic        bool
	AllNoTheorem        bool
	AllSourceTagged     bool
	AllConventionTagged bool
	AnyNativePromotion  bool
	AnyNativeWrite      bool
	AnyPhysicalClaim    bool
	AnyBridgeEvidence   bool
	AnyObservedClaim    bool
	ChecksumExpected    string
	ChecksumActual      string
	ChecksumVerified    bool
	Verdict             string
	Reason              string
	Failures            []string
}

type ReviewResult struct {
	Executed bool

	ManifestParsed                bool
	HumanReviewMetadataParsed     bool
	ReproducibilityMetadataParsed bool
	SourceChainMetadataParsed     bool
	ResidualThresholdPolicyParsed bool
	PhysicalClaimDiscriminator    bool
	CitationScopeQuarantineOnly   bool
	NativeWriteDeltaZero          bool
	RollbackPlanParsed            bool
	PostReleaseAuditParsed        bool
	SyntheticUnderlyingOutput     bool
	AuthenticatedSourceChain      bool
	ReleaseAllowed                bool
	BridgeEvidenceReleased        bool
	NativeWriteLocked             bool
	NativeWriteAuthorization      bool
	NativeRegistryWrite           bool
	BlockedBecauseSynthetic       bool

	Verdict  string
	Reason   string
	Failures []string
}

type Firewall struct {
	Executed bool

	SyntheticReleaseManifestPresent  bool
	ComparatorOutputReleased         bool
	BridgeEvidenceClaimReleased      bool
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
	Review      ReviewResult
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
	g546, err := generation2comparatoroutputreleaseairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate546 release airlock: %w", err)
	}
	ledger, resolved, err := loadLedger(path)
	if err != nil {
		return Analysis{}, fmt.Errorf("%s: %w", StatusFailedLedgerMissing, err)
	}
	a := Analysis{Inheritance: buildInheritance(g546)}
	a.Import = buildImport(ledger, resolved)
	a.Review = buildReview(a.Inheritance, a.Import)
	a.Firewall = buildFirewall(a.Import, a.Review)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2comparatoroutputreleaseairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                  true,
		Gate546AirlockDefined:     g.Guard.AirlockDefined,
		Gate546SchemaRows:         g.Schema.RequiredRows,
		Gate546QuarantinedPresent: g.Guard.QuarantinedComparatorPresent,
		Gate546ReleaseBlocked:     !g.Guard.BridgeEvidenceReleaseAllowed && !g.Guard.BridgeEvidenceReleased,
		Gate546NoBridgeEvidence:   !g.Firewall.BridgeEvidenceClaimReleased,
		Gate546NativeWriteLocked:  g.Guard.NativeWriteLocked && !g.Guard.NativeWriteAuthorization && !g.Guard.NativeRegistryWrite,
		Gate546AbortSynthetic:     g.Guard.AbortTriggeredBySynthetic,
		Gate546NoRealSource:       !g.Firewall.RealSchwingerSourceImported && !g.Firewall.AuthenticatedRealSource,
		Gate546NoPhysicalClaims:   !g.Firewall.PhysicalSchwingerFunctionsLoaded && !g.Firewall.OSPositivityCertificateLoaded && !g.Firewall.WickMapLoaded && !g.Firewall.HilbertSpaceReconstructed && !g.Firewall.HamiltonianSpectrumLoaded && !g.Firewall.UnitaryDynamicsLoaded && !g.Firewall.GlobalCausalityLoaded && !g.Firewall.TimeArrowLoaded,
		Gate546RedirectsToGate547: g.Next.Gate == 547,
		Verdict:                   StatusGate546ReleaseAirlockInherited,
		Reason:                    "Gate547 inherits Gate546's release airlock, 15-row review schema, quarantined Gate545 output, release block, and native-write lock.",
	}
}

func requiredRows() []string {
	return []string{
		"quarantine_result_reference",
		"comparator_result_checksum_reference",
		"authenticated_source_chain_reference",
		"operator_release_intent",
		"human_review_attestation",
		"independent_reproducibility_report",
		"residual_threshold_policy",
		"os_wick_hilbert_hamiltonian_certificate_map",
		"physical_claim_discriminator",
		"environmental_boundary_statement",
		"bridge_evidence_citation_scope",
		"native_write_delta_manifest",
		"release_target_quarantine_only",
		"rollback_and_revocation_plan",
		"post_release_audit_log",
	}
}

func buildImport(l Ledger, p string) Import {
	imp := Import{Executed: true, Loaded: true, Path: p, Rows: len(l.Rows), BridgeOnly: l.BridgeOnly, SyntheticFixture: l.SyntheticFixture, ManifestImported: l.ReleaseManifestImported, OperatorIntent: l.OperatorReleaseIntent, HumanReview: l.HumanReviewAttestation, Reproducibility: l.IndependentReproducibilityReport, SourceChain: l.AuthenticatedSourceChain, ResidualThreshold: l.ResidualThresholdAccepted, Discriminator: l.PhysicalClaimDiscriminator, CitationScope: strings.Contains(l.BridgeEvidenceCitationScope, "quarantine"), QuarantineOnly: l.ReleaseTargetQuarantineOnly, ReleaseAllowed: l.BridgeEvidenceReleaseAllowed, Released: l.BridgeEvidenceReleased, RealSource: l.RealSchwingerSourceImported, AuthenticatedReal: l.AuthenticatedNonSyntheticSource, SchwingerLoaded: l.PhysicalSchwingerFunctionsLoaded, OSCertLoaded: l.PhysicalOSCertificateLoaded, WickMapLoaded: l.PhysicalWickMapLoaded, HilbertLoaded: l.PhysicalHilbertSpaceLoaded, HamiltonianLoaded: l.PhysicalHamiltonianLoaded, UnitaryLoaded: l.UnitaryDynamicsLoaded, GlobalLoaded: l.GlobalCausalityLoaded, ArrowLoaded: l.TimeArrowLoaded, NativeWriteLock: l.NativeWriteLock, NativeDeltaZero: l.NativeWriteDeltaZero, NativeWrite: l.NativeRegistryWrite, RollbackPlan: l.RollbackAndRevocationPlan != "", PostReleaseAudit: l.PostReleaseAuditLog != "", ChecksumExpected: l.CanonicalPayloadSHA256}
	seen := map[string]bool{}
	req := map[string]bool{}
	for _, r := range requiredRows() {
		req[r] = false
	}
	imp.AllBridgeOnly, imp.AllReleaseOnly, imp.AllQuarantineOnly, imp.AllDryRunOnly, imp.AllSynthetic, imp.AllNoTheorem, imp.AllSourceTagged, imp.AllConventionTagged = true, true, true, true, true, true, true, true
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
		imp.AllReleaseOnly = imp.AllReleaseOnly && r.ReleaseOnly
		imp.AllQuarantineOnly = imp.AllQuarantineOnly && r.QuarantineOnly
		imp.AllDryRunOnly = imp.AllDryRunOnly && r.DryRunOnly
		imp.AllSynthetic = imp.AllSynthetic && r.Synthetic
		imp.AllNoTheorem = imp.AllNoTheorem && r.NoTheoremInput
		imp.AllSourceTagged = imp.AllSourceTagged && r.Source != "" && r.SourceVersion != ""
		imp.AllConventionTagged = imp.AllConventionTagged && r.Convention != ""
		imp.AnyNativePromotion = imp.AnyNativePromotion || r.NativePromotion
		imp.AnyNativeWrite = imp.AnyNativeWrite || r.NativeWrite
		imp.AnyPhysicalClaim = imp.AnyPhysicalClaim || r.PhysicalClaim
		imp.AnyBridgeEvidence = imp.AnyBridgeEvidence || r.BridgeEvidence
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
	if !imp.AllBridgeOnly || !imp.AllReleaseOnly || !imp.AllQuarantineOnly || !imp.AllDryRunOnly || !imp.AllSynthetic || !imp.AllNoTheorem || !imp.AllSourceTagged || !imp.AllConventionTagged || imp.AnyNativePromotion || imp.AnyNativeWrite || imp.AnyPhysicalClaim || imp.AnyBridgeEvidence || imp.AnyObservedClaim {
		imp.Failures = append(imp.Failures, StatusFailedMetadataIncomplete)
	}
	if !imp.ChecksumVerified {
		imp.Failures = append(imp.Failures, StatusFailedChecksumMismatch)
	}
	if imp.ReleaseAllowed || imp.Released || imp.RealSource || imp.AuthenticatedReal || imp.SchwingerLoaded || imp.OSCertLoaded || imp.WickMapLoaded || imp.HilbertLoaded || imp.HamiltonianLoaded || imp.UnitaryLoaded || imp.GlobalLoaded || imp.ArrowLoaded || imp.NativeWrite {
		imp.Failures = append(imp.Failures, StatusFailedReleaseLeaked)
	}
	if len(imp.Failures) == 0 {
		imp.Verdict = StatusSyntheticReleaseAdapterExecuted
		imp.Reason = "Synthetic release-review manifest parsed all 15 Gate546 rows, verified checksum, and remained quarantine-only with zero native-write delta."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Synthetic release-review manifest failed schema, metadata, checksum, or release firewall validation."
	}
	return imp
}

func buildReview(i Inheritance, imp Import) ReviewResult {
	r := ReviewResult{Executed: true,
		ManifestParsed:                imp.Loaded && imp.ManifestImported && imp.ChecksumVerified && len(imp.Failures) == 0,
		HumanReviewMetadataParsed:     imp.HumanReview && imp.OperatorIntent,
		ReproducibilityMetadataParsed: imp.Reproducibility && imp.ResidualThreshold,
		SourceChainMetadataParsed:     true,
		ResidualThresholdPolicyParsed: imp.ResidualThreshold,
		PhysicalClaimDiscriminator:    imp.Discriminator,
		CitationScopeQuarantineOnly:   imp.CitationScope && imp.QuarantineOnly,
		NativeWriteDeltaZero:          imp.NativeWriteLock && imp.NativeDeltaZero && !imp.NativeWrite && !imp.AnyNativeWrite,
		RollbackPlanParsed:            imp.RollbackPlan,
		PostReleaseAuditParsed:        imp.PostReleaseAudit,
		SyntheticUnderlyingOutput:     i.Gate546QuarantinedPresent && imp.SyntheticFixture,
		AuthenticatedSourceChain:      imp.SourceChain && imp.AuthenticatedReal,
		ReleaseAllowed:                false,
		BridgeEvidenceReleased:        false,
		NativeWriteLocked:             imp.NativeWriteLock,
		NativeWriteAuthorization:      false,
		NativeRegistryWrite:           false,
		BlockedBecauseSynthetic:       i.Gate546AbortSynthetic && imp.SyntheticFixture && !imp.AuthenticatedReal,
		Verdict:                       StatusSyntheticReleaseBlocked,
		Reason:                        "The synthetic release manifest exercises review plumbing, but release remains blocked because the underlying comparator output is synthetic and has no authenticated non-synthetic source chain.",
	}
	if !r.ManifestParsed || !r.HumanReviewMetadataParsed || !r.ReproducibilityMetadataParsed || !r.SourceChainMetadataParsed || !r.ResidualThresholdPolicyParsed || !r.PhysicalClaimDiscriminator || !r.CitationScopeQuarantineOnly || !r.NativeWriteDeltaZero || !r.RollbackPlanParsed || !r.PostReleaseAuditParsed {
		r.Failures = append(r.Failures, StatusFailedMetadataIncomplete)
	}
	if r.AuthenticatedSourceChain || r.ReleaseAllowed || r.BridgeEvidenceReleased || !r.NativeWriteLocked || r.NativeWriteAuthorization || r.NativeRegistryWrite || !r.BlockedBecauseSynthetic {
		r.Failures = append(r.Failures, StatusFailedReleaseLeaked)
	}
	if len(r.Failures) > 0 {
		r.Verdict = strings.Join(r.Failures, ";")
		r.Reason = "Synthetic release-review result leaked authorization or failed review metadata."
	}
	return r
}

func buildFirewall(i Import, r ReviewResult) Firewall {
	return Firewall{Executed: true,
		SyntheticReleaseManifestPresent:  i.ManifestImported,
		ComparatorOutputReleased:         r.BridgeEvidenceReleased,
		BridgeEvidenceClaimReleased:      r.ReleaseAllowed,
		RealSchwingerSourceImported:      i.RealSource,
		AuthenticatedRealSource:          i.AuthenticatedReal,
		PhysicalSchwingerFunctionsLoaded: i.SchwingerLoaded,
		OSPositivityCertificateLoaded:    i.OSCertLoaded,
		WickMapLoaded:                    i.WickMapLoaded,
		HilbertSpaceReconstructed:        i.HilbertLoaded,
		HamiltonianSpectrumLoaded:        i.HamiltonianLoaded,
		UnitaryDynamicsLoaded:            i.UnitaryLoaded,
		GlobalCausalityLoaded:            i.GlobalLoaded,
		TimeArrowLoaded:                  i.ArrowLoaded,
		NativeSchwingerFunctionWrite:     false,
		NativeOSPositivityWrite:          false,
		NativeWickWrite:                  false,
		NativeHilbertWrite:               false,
		NativeHamiltonianWrite:           false,
		NativeUnitaryDynamicsWrite:       false,
		NativeGlobalCausalWrite:          false,
		NativeTimeArrowWrite:             false,
		NativeRegistryWritten:            false,
		Verdict:                          StatusFirewallPreserved,
		Reason:                           "Gate547 parses a synthetic release-review manifest but releases no bridge evidence and writes no native physics.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"unchanged: Gate547 writes no native Schwinger, OS, Wick, Hilbert, Hamiltonian, unitarity, global-causal, time-arrow, or release theorem"},
		BridgeEntries:        []string{"synthetic release-review manifest adapter", "checksum-verified 15-row release manifest dry run", "human-review, reproducibility, source-chain, citation-scope, rollback, post-release audit, and zero-native-write metadata parsed"},
		EnvironmentalEntries: []string{"authenticated non-synthetic source chain", "real comparator result", "actual bridge-evidence release", "physical OS/Wick/Hilbert/Hamiltonian certificates"},
		FailedRoutes:         []string{StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputStillQuarantined, StatusFirewallNativeWriteBlocked},
		OpenTheorems:         []string{"Gate548 candidate: release-sector closure ledger mapping parser, comparator, release, and source authenticity boundaries before any non-synthetic external evidence can be accepted."},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 548, Title: "Physical Correlation Import/Release Sector Closure Ledger", Reason: "Gate547 proves the synthetic release manifest parser and rejection path. The next safe gate is a closure ledger for the whole Schwinger/source/comparator/release pipeline, freezing what is native, bridge-only, and environmental.", PrimaryTask: "Emit a closure/frontier map for Gates 536-547, preserving the rule that no physical correlation data, released bridge evidence, or native dynamics theorem exists without authenticated non-synthetic sources and review."}
}

func truth(a Analysis) string {
	return "Gate547 executes the synthetic release-review manifest adapter: all 15 release rows parse, checksum and review/reproducibility metadata pass, zero-native-write delta is verified, but release remains blocked because the comparator output is synthetic and unauthenticated as physical bridge evidence."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Gate546AirlockDefined || a.Inheritance.Gate546SchemaRows != 15 || !a.Inheritance.Gate546QuarantinedPresent || !a.Inheritance.Gate546ReleaseBlocked || !a.Inheritance.Gate546NoBridgeEvidence || !a.Inheritance.Gate546NativeWriteLocked || !a.Inheritance.Gate546AbortSynthetic || !a.Inheritance.Gate546NoRealSource || !a.Inheritance.Gate546NoPhysicalClaims || !a.Inheritance.Gate546RedirectsToGate547 {
		bad = append(bad, "Gate546 inheritance incomplete")
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 15 || len(a.Import.MissingRows) > 0 || a.Import.RejectedRows != 0 || len(a.Import.DuplicateRows) > 0 || !a.Import.ChecksumVerified {
		bad = append(bad, "release manifest schema/checksum incomplete")
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllReleaseOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || !a.Import.AllSourceTagged || !a.Import.AllConventionTagged || a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyBridgeEvidence || a.Import.AnyObservedClaim {
		bad = append(bad, "metadata sieve failed")
	}
	if !a.Review.ManifestParsed || !a.Review.HumanReviewMetadataParsed || !a.Review.ReproducibilityMetadataParsed || !a.Review.SourceChainMetadataParsed || !a.Review.ResidualThresholdPolicyParsed || !a.Review.PhysicalClaimDiscriminator || !a.Review.CitationScopeQuarantineOnly || !a.Review.NativeWriteDeltaZero || !a.Review.RollbackPlanParsed || !a.Review.PostReleaseAuditParsed || !a.Review.SyntheticUnderlyingOutput || a.Review.AuthenticatedSourceChain || a.Review.ReleaseAllowed || a.Review.BridgeEvidenceReleased || !a.Review.NativeWriteLocked || a.Review.NativeWriteAuthorization || a.Review.NativeRegistryWrite || !a.Review.BlockedBecauseSynthetic {
		bad = append(bad, "review guard leaked")
	}
	if a.Firewall.ComparatorOutputReleased || a.Firewall.BridgeEvidenceClaimReleased || a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate547 validation failed: %s", strings.Join(bad, "; "))
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
	return fmt.Sprintf("%s: schema_rows=%d airlock=%t quarantined=%t blocked=%t no_bridge=%t native_locked=%t abort_synthetic=%t no_real=%t no_physical=%t redirects=%t; %s", x.Verdict, x.Gate546SchemaRows, x.Gate546AirlockDefined, x.Gate546QuarantinedPresent, x.Gate546ReleaseBlocked, x.Gate546NoBridgeEvidence, x.Gate546NativeWriteLocked, x.Gate546AbortSynthetic, x.Gate546NoRealSource, x.Gate546NoPhysicalClaims, x.Gate546RedirectsToGate547, x.Reason)
}

func FormatImport(x Import) string {
	return fmt.Sprintf("%s;%s;%s;%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%s duplicates=%s checksum=%t expected=%s actual=%s manifest=%t intent=%t review=%t reproducibility=%t source_chain=%t threshold=%t discriminator=%t citation=%t quarantine=%t release_allowed=%t released=%t native_lock=%t delta_zero=%t native_write=%t real=%t auth_real=%t bridge=%t release_only=%t quarantine_only=%t dryrun_only=%t synthetic=%t no_theorem=%t; %s", StatusSyntheticReleaseManifestLoaded, StatusSyntheticRelease15RowsAccepted, StatusSyntheticReleaseChecksumVerified, x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, strings.Join(x.MissingRows, ","), strings.Join(x.DuplicateRows, ","), x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.ManifestImported, x.OperatorIntent, x.HumanReview, x.Reproducibility, x.SourceChain, x.ResidualThreshold, x.Discriminator, x.CitationScope, x.QuarantineOnly, x.ReleaseAllowed, x.Released, x.NativeWriteLock, x.NativeDeltaZero, x.NativeWrite, x.RealSource, x.AuthenticatedReal, x.AllBridgeOnly, x.AllReleaseOnly, x.AllQuarantineOnly, x.AllDryRunOnly, x.AllSynthetic, x.AllNoTheorem, x.Reason)
}

func FormatReview(x ReviewResult) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: manifest=%t human=%t repro=%t source_chain_meta=%t threshold=%t discriminator=%t citation_quarantine=%t delta_zero=%t rollback=%t post_audit=%t synthetic_output=%t auth_chain=%t release_allowed=%t released=%t native_locked=%t native_auth=%t registry=%t blocked_synthetic=%t; %s", StatusSyntheticReleaseHumanReviewParsed, StatusSyntheticReleaseReproducibilityParsed, StatusSyntheticReleaseSourceChainParsed, StatusSyntheticReleaseCitationScopeParsed, StatusSyntheticReleaseNativeDeltaZero, StatusSyntheticReleaseBlocked, StatusNoBridgeEvidenceReleased, StatusNoRealSourceImported, StatusNativePromotionRejected, x.Verdict, x.ManifestParsed, x.HumanReviewMetadataParsed, x.ReproducibilityMetadataParsed, x.SourceChainMetadataParsed, x.ResidualThresholdPolicyParsed, x.PhysicalClaimDiscriminator, x.CitationScopeQuarantineOnly, x.NativeWriteDeltaZero, x.RollbackPlanParsed, x.PostReleaseAuditParsed, x.SyntheticUnderlyingOutput, x.AuthenticatedSourceChain, x.ReleaseAllowed, x.BridgeEvidenceReleased, x.NativeWriteLocked, x.NativeWriteAuthorization, x.NativeRegistryWrite, x.BlockedBecauseSynthetic, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s: manifest=%t released=%t bridge_claim=%t real=%t auth_real=%t schwinger=%t os=%t wick=%t hilbert=%t ham=%t unitary=%t global=%t arrow=%t native_s=%t native_os=%t native_wick=%t native_hilbert=%t native_ham=%t native_unitary=%t native_global=%t native_arrow=%t registry=%t; %s", StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputStillQuarantined, StatusNativePromotionRejected, x.SyntheticReleaseManifestPresent, x.ComparatorOutputReleased, x.BridgeEvidenceClaimReleased, x.RealSchwingerSourceImported, x.AuthenticatedRealSource, x.PhysicalSchwingerFunctionsLoaded, x.OSPositivityCertificateLoaded, x.WickMapLoaded, x.HilbertSpaceReconstructed, x.HamiltonianSpectrumLoaded, x.UnitaryDynamicsLoaded, x.GlobalCausalityLoaded, x.TimeArrowLoaded, x.NativeSchwingerFunctionWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.NativeRegistryWritten, x.Reason)
}

func Statuses() []string {
	return []string{StatusGate546ReleaseAirlockInherited, StatusSyntheticReleaseManifestLoaded, StatusSyntheticReleaseAdapterExecuted, StatusSyntheticRelease15RowsAccepted, StatusSyntheticReleaseChecksumVerified, StatusSyntheticReleaseMetadataEnforced, StatusSyntheticReleaseHumanReviewParsed, StatusSyntheticReleaseReproducibilityParsed, StatusSyntheticReleaseSourceChainParsed, StatusSyntheticReleaseCitationScopeParsed, StatusSyntheticReleaseNativeDeltaZero, StatusSyntheticReleaseBlocked, StatusNoBridgeEvidenceReleased, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotEvidence, StatusFailedSyntheticNotReal, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOS, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamilton, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFailedOutputStillQuarantined, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
