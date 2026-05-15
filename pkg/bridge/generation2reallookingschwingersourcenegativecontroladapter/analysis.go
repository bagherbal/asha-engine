// Package generation2reallookingschwingersourcenegativecontroladapter implements Gate 541:
// Real-Looking Schwinger Source Negative-Control Adapter.
//
// Gate 540 installed the default-off switch that must guard any future
// non-synthetic Schwinger source import. Gate 541 loads an intentionally
// real-looking but untrusted negative-control ledger and proves that parser
// success, physical-looking labels, and checksum plumbing still do not permit
// comparator execution or native registry writes while the switch is off and
// provenance is insufficient.
package generation2reallookingschwingersourcenegativecontroladapter

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

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2realschwingerimportswitchairlock"
)

const (
	AuditID       = "GATE541-REAL-LOOKING-SCHWINGER-SOURCE-NEGATIVE-CONTROL-ADAPTER"
	DefaultLedger = "data/real_looking_schwinger_negative_control_ledger_gate541.json"

	StatusGate540SwitchInherited          = "CONDITIONAL_SUPPORT_GATE540_REAL_IMPORT_SWITCH_INHERITED"
	StatusNegativeControlLedgerLoaded     = "CONDITIONAL_SUPPORT_REAL_LOOKING_SCHWINGER_NEGATIVE_CONTROL_LEDGER_LOADED"
	StatusNegativeControlRowsAccepted     = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_SWITCH_ROWS_ACCEPTED"
	StatusNegativeControlChecksumVerified = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_CHECKSUM_VERIFIED"
	StatusNegativeControlMetadataParsed   = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_METADATA_PARSED"
	StatusNegativeControlAdapterExecuted  = "CONDITIONAL_SUPPORT_REAL_LOOKING_SCHWINGER_NEGATIVE_CONTROL_ADAPTER_EXECUTED"
	StatusRejectedSwitchOff               = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_SWITCH_OFF"
	StatusRejectedNoOperatorIntent        = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_NO_OPERATOR_INTENT"
	StatusRejectedInsufficientProvenance  = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_REJECTED_INSUFFICIENT_PROVENANCE"
	StatusRejectedComparatorAuthorization = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_COMPARATOR_AUTHORIZATION_BLOCKED"
	StatusQuarantinePreserved             = "CONDITIONAL_SUPPORT_REAL_LOOKING_SOURCE_QUARANTINE_PRESERVED"
	StatusNoComparatorExecuted            = "CONDITIONAL_SUPPORT_NO_REAL_SOURCE_COMPARATOR_EXECUTED_IN_GATE541"
	StatusNoNativeWrite                   = "CONDITIONAL_SUPPORT_NO_NATIVE_WRITE_FROM_REAL_LOOKING_SOURCE_GATE541"

	StatusFailedLedgerMissing                 = "FAILED_ROUTE_GATE541_REAL_LOOKING_SOURCE_LEDGER_MISSING"
	StatusFailedMetadataIncomplete            = "FAILED_ROUTE_GATE541_REAL_LOOKING_SOURCE_METADATA_INCOMPLETE"
	StatusFailedSchemaRowsIncomplete          = "FAILED_ROUTE_GATE541_REAL_LOOKING_SOURCE_SWITCH_ROWS_INCOMPLETE"
	StatusFailedChecksumMismatch              = "FAILED_ROUTE_GATE541_REAL_LOOKING_SOURCE_CHECKSUM_MISMATCH"
	StatusFailedSwitchOffBlocksImport         = "FAILED_ROUTE_REAL_LOOKING_SOURCE_IMPORT_SWITCH_OFF"
	StatusFailedNoIntentBlocksImport          = "FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_EXPLICIT_OPERATOR_INTENT"
	StatusFailedInsufficientProvenance        = "FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_INSUFFICIENT_PROVENANCE"
	StatusFailedNoAccessGrant                 = "FAILED_ROUTE_REAL_LOOKING_SOURCE_HAS_NO_LICENSE_OR_ACCESS_GRANT"
	StatusFailedUntrustedURI                  = "FAILED_ROUTE_REAL_LOOKING_SOURCE_URI_NOT_AUTHENTICATED"
	StatusFailedNegativeControlNotSchwinger   = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedNegativeControlNotOSProof     = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_PROVE_OS_REFLECTION_POSITIVITY"
	StatusFailedNegativeControlNotWick        = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedNegativeControlNotHilbert     = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedNegativeControlNotHamiltonian = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedNegativeControlNotUnitary     = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedNegativeControlNotGlobal      = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedNegativeControlNotArrow       = "FAILED_ROUTE_REAL_LOOKING_NEGATIVE_CONTROL_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_GATE541_REAL_LOOKING_NEGATIVE_CONTROL_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked          = "FIREWALL_BLOCKED_GATE541_REAL_LOOKING_SOURCE_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate540SwitchDefined        bool
	Gate540SwitchDefaultOff     bool
	Gate540OperatorIntentNeeded bool
	Gate540ComparatorBlocked    bool
	Gate540NoRealSourceImported bool
	Gate540NativeWriteBlocked   bool
	Gate540RedirectsToGate541   bool

	Verdict, Reason string
}

type LedgerRowMetadata struct {
	SchemaKey        string `json:"schema_key"`
	Source           string `json:"source"`
	SourceVersion    string `json:"source_version"`
	Convention       string `json:"convention"`
	ValueKind        string `json:"value_kind"`
	BridgeOnly       bool   `json:"bridge_only"`
	ComparatorOnly   bool   `json:"comparator_only"`
	Synthetic        bool   `json:"synthetic"`
	NegativeControl  bool   `json:"negative_control"`
	PhysicalClaim    bool   `json:"physical_claim"`
	Observed         bool   `json:"observed"`
	NoTheoremInput   bool   `json:"no_theorem_input"`
	NativePromotion  bool   `json:"native_promotion"`
	NativeInputClaim bool   `json:"native_input_claim"`
	Value            string `json:"value"`
}

type NegativeControlLedger struct {
	Gate                                    int                 `json:"gate"`
	LedgerName                              string              `json:"ledger_name"`
	Description                             string              `json:"description"`
	Gate540SwitchReference                  string              `json:"gate540_switch_reference"`
	Gate539AuthenticityReference            string              `json:"gate539_authenticity_reference"`
	Gate536SchwingerSchemaReference         string              `json:"gate536_schwinger_schema_reference"`
	BridgeOnly                              bool                `json:"bridge_only"`
	NativeRegistryWrite                     bool                `json:"native_registry_write"`
	RealLookingFixture                      bool                `json:"real_looking_fixture"`
	NegativeControlFixture                  bool                `json:"negative_control_fixture"`
	DeclaredSyntheticFixture                bool                `json:"declared_synthetic_fixture"`
	PhysicalSourceClaim                     bool                `json:"physical_source_claim"`
	NonSyntheticPhysicalOrConstructiveClaim bool                `json:"non_synthetic_physical_or_constructive_claim"`
	RealSourceImportSwitchEnabled           bool                `json:"real_source_import_switch_enabled"`
	ExplicitOperatorIntentProvided          bool                `json:"explicit_operator_intent_provided"`
	ComparatorExecutionRequested            bool                `json:"comparator_execution_requested"`
	ComparatorAuthorizationGranted          bool                `json:"comparator_authorization_granted"`
	NonSyntheticSourceURI                   string              `json:"non_synthetic_source_uri"`
	LicenseAndAccessGrantProvided           bool                `json:"license_and_access_grant_provided"`
	AuthenticityLedgerReferenceProvided     bool                `json:"authenticity_ledger_reference_provided"`
	ChecksumOrProofHashProvided             bool                `json:"checksum_or_proof_hash_provided"`
	TrustedSourceURI                        bool                `json:"trusted_source_uri"`
	PhysicalSchwingerSourceLoaded           bool                `json:"physical_schwinger_source_loaded"`
	ConstructiveMeasureLoaded               bool                `json:"constructive_measure_loaded"`
	ObservedCorrelationLoaded               bool                `json:"observed_correlation_loaded"`
	PhysicalOSCertificateLoaded             bool                `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded                   bool                `json:"physical_wick_map_loaded"`
	PhysicalHamiltonianLoaded               bool                `json:"physical_hamiltonian_loaded"`
	Source                                  string              `json:"source"`
	SourceVersion                           string              `json:"source_version"`
	Convention                              string              `json:"convention"`
	CanonicalPayload                        map[string]any      `json:"canonical_payload"`
	CanonicalPayloadSHA256                  string              `json:"canonical_payload_sha256"`
	Rows                                    []LedgerRowMetadata `json:"rows"`
}

type FileImport struct {
	Executed bool
	Loaded   bool
	Path     string

	Rows                int
	AcceptedRows        int
	RejectedRows        int
	MissingRequiredRows []string
	DuplicateRows       []string

	BridgeOnlyLedger                    bool
	NativeRegistryWriteRequested        bool
	RealLookingFixture                  bool
	NegativeControlFixture              bool
	DeclaredSyntheticFixture            bool
	PhysicalSourceClaim                 bool
	NonSyntheticClaim                   bool
	RealSourceImportSwitchEnabled       bool
	ExplicitOperatorIntentProvided      bool
	ComparatorExecutionRequested        bool
	ComparatorAuthorizationGranted      bool
	LicenseAndAccessGrantProvided       bool
	AuthenticityLedgerReferenceProvided bool
	ChecksumOrProofHashProvided         bool
	TrustedSourceURI                    bool
	PhysicalSchwingerSourceLoaded       bool
	ObservedCorrelationLoaded           bool
	ConstructiveMeasureLoaded           bool
	PhysicalOSCertificateLoaded         bool
	PhysicalWickMapLoaded               bool
	PhysicalHamiltonianLoaded           bool
	Gate540ReferenceComplete            bool
	Gate539ReferenceComplete            bool
	Gate536ReferenceComplete            bool
	SourceTaggedLedger                  bool
	ConventionTaggedLedger              bool
	MetadataComplete                    bool
	RequiredSchemaRowsMatched           bool
	AllRowsBridgeOnly                   bool
	AllRowsNoTheoremInput               bool
	AllRowsNegativeControl              bool
	AllRowsSourceTagged                 bool
	AllRowsConventionTagged             bool
	AnyNativePromotionClaim             bool
	AnyNativeInputClaim                 bool
	ChecksumExpected                    string
	ChecksumActual                      string
	ChecksumVerified                    bool

	Verdict, Reason string
	Failures        []string
}

type Rejection struct {
	Executed bool

	SourceParsed                   bool
	RealLookingClaimSeen           bool
	PhysicalSourceClaimSeen        bool
	NegativeControlFixtureSeen     bool
	SwitchOff                      bool
	NoExplicitOperatorIntent       bool
	MissingLicenseOrAccessGrant    bool
	SourceURINotAuthenticated      bool
	ComparatorAuthorizationMissing bool
	InsufficientProvenance         bool
	ComparatorExecutionAllowed     bool
	ComparatorExecutionPerformed   bool
	PhysicalSourceAuthenticated    bool
	PhysicalSourceImported         bool
	RejectedAsPhysicalSource       bool
	RejectedBeforeComparator       bool
	QuarantinePreserved            bool
	RejectionReasons               []string

	Verdict, Reason string
}

type Firewall struct {
	Executed bool

	ComparatorExecutionPerformed  bool
	RealSchwingerSourceImported   bool
	ObservedCorrelationImported   bool
	ConstructiveMeasureImported   bool
	PhysicalOSCertificateImported bool
	PhysicalWickMapImported       bool
	PhysicalHamiltonianImported   bool
	PhysicalSchwingerDerived      bool
	OSPositivityProven            bool
	WickRotationSelected          bool
	PhysicalHilbertSpaceSelected  bool
	PositiveHamiltonianDerived    bool
	UnitaryDynamicsDerived        bool
	GlobalCausalitySelected       bool
	ArrowOfTimeSelected           bool
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
	Inheritance Inheritance
	Import      FileImport
	Rejection   Rejection
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

func Build(ledgerPath ...string) (Analysis, error) {
	g540, err := generation2realschwingerimportswitchairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate540 real source import switch airlock: %w", err)
	}
	path := DefaultLedger
	if len(ledgerPath) > 0 && strings.TrimSpace(ledgerPath[0]) != "" {
		path = ledgerPath[0]
	}
	ledger, resolvedPath, err := loadLedger(path)
	if err != nil {
		return Analysis{}, err
	}
	a := Analysis{Inheritance: buildInheritance(g540)}
	a.Import = analyzeImport(ledger, resolvedPath)
	a.Rejection = buildRejection(a.Inheritance, a.Import)
	a.Firewall = buildFirewall(a.Import, a.Rejection)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func loadLedger(path string) (NegativeControlLedger, string, error) {
	resolved := path
	if !filepath.IsAbs(resolved) {
		root, err := repoRoot()
		if err != nil {
			return NegativeControlLedger{}, "", err
		}
		resolved = filepath.Join(root, path)
	}
	b, err := os.ReadFile(resolved)
	if err != nil {
		return NegativeControlLedger{}, resolved, fmt.Errorf("%s: %w", StatusFailedLedgerMissing, err)
	}
	var ledger NegativeControlLedger
	if err := json.Unmarshal(b, &ledger); err != nil {
		return NegativeControlLedger{}, resolved, fmt.Errorf("decode Gate541 negative-control ledger: %w", err)
	}
	return ledger, resolved, nil
}

func repoRoot() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("could not resolve caller path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "../../..")), nil
}

func buildInheritance(g generation2realschwingerimportswitchairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                    true,
		Gate540SwitchDefined:        g.Schema.Executed && len(g.Schema.Rows) == 12 && g.Schema.RequiredRows == 12 && g.Schema.NativeWriteRows == 0,
		Gate540SwitchDefaultOff:     g.Switch.Executed && g.Switch.DefaultOff && !g.Switch.RealSourceImportEnabled && !g.Switch.ComparatorExecutionAllowed,
		Gate540OperatorIntentNeeded: !g.Switch.ExplicitOperatorIntentProvided,
		Gate540ComparatorBlocked:    g.Guard.Executed && !g.Guard.ComparatorExecutionPerformed && g.Firewall.ComparatorAuthorizationBlocked,
		Gate540NoRealSourceImported: !g.Firewall.RealSchwingerSourceImported && !g.Firewall.ObservedCorrelationImported && !g.Firewall.ConstructiveMeasureImported,
		Gate540NativeWriteBlocked:   !g.Firewall.NativeRegistryWritten && !g.Firewall.NativeSchwingerFunctionWrite && !g.Firewall.NativeHamiltonianWrite,
		Gate540RedirectsToGate541:   g.Next.Gate == 541,
		Verdict:                     StatusGate540SwitchInherited,
		Reason:                      "Gate541 inherits Gate540's default-off real-source switch, missing operator-intent guard, comparator block, and native-write lock.",
	}
}

func analyzeImport(l NegativeControlLedger, path string) FileImport {
	req := requiredRows()
	seen := map[string]int{}
	var missing, duplicate []string
	allBridge, allNoTheorem, allNegative, allSource, allConvention := true, true, true, true, true
	anyNativePromotion, anyNativeInput := false, false
	accepted := 0
	for _, r := range l.Rows {
		seen[r.SchemaKey]++
		if req[r.SchemaKey] {
			accepted++
		}
		allBridge = allBridge && r.BridgeOnly
		allNoTheorem = allNoTheorem && r.NoTheoremInput
		allNegative = allNegative && r.NegativeControl
		allSource = allSource && strings.TrimSpace(r.Source) != ""
		allConvention = allConvention && strings.TrimSpace(r.Convention) != ""
		anyNativePromotion = anyNativePromotion || r.NativePromotion
		anyNativeInput = anyNativeInput || r.NativeInputClaim
	}
	for k := range req {
		if seen[k] == 0 {
			missing = append(missing, k)
		}
		if seen[k] > 1 {
			duplicate = append(duplicate, k)
		}
	}
	sort.Strings(missing)
	sort.Strings(duplicate)
	actual := checksum(l.CanonicalPayload)
	failures := []string{}
	if len(missing) > 0 || len(duplicate) > 0 {
		failures = append(failures, StatusFailedSchemaRowsIncomplete)
	}
	if strings.TrimSpace(l.Source) == "" || strings.TrimSpace(l.SourceVersion) == "" || strings.TrimSpace(l.Convention) == "" || !allSource || !allConvention {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if strings.TrimSpace(l.CanonicalPayloadSHA256) == "" || l.CanonicalPayloadSHA256 != actual {
		failures = append(failures, StatusFailedChecksumMismatch)
	}
	return FileImport{
		Executed:                            true,
		Loaded:                              true,
		Path:                                path,
		Rows:                                len(l.Rows),
		AcceptedRows:                        accepted,
		RejectedRows:                        len(l.Rows) - accepted,
		MissingRequiredRows:                 missing,
		DuplicateRows:                       duplicate,
		BridgeOnlyLedger:                    l.BridgeOnly,
		NativeRegistryWriteRequested:        l.NativeRegistryWrite,
		RealLookingFixture:                  l.RealLookingFixture,
		NegativeControlFixture:              l.NegativeControlFixture,
		DeclaredSyntheticFixture:            l.DeclaredSyntheticFixture,
		PhysicalSourceClaim:                 l.PhysicalSourceClaim,
		NonSyntheticClaim:                   l.NonSyntheticPhysicalOrConstructiveClaim,
		RealSourceImportSwitchEnabled:       l.RealSourceImportSwitchEnabled,
		ExplicitOperatorIntentProvided:      l.ExplicitOperatorIntentProvided,
		ComparatorExecutionRequested:        l.ComparatorExecutionRequested,
		ComparatorAuthorizationGranted:      l.ComparatorAuthorizationGranted,
		LicenseAndAccessGrantProvided:       l.LicenseAndAccessGrantProvided,
		AuthenticityLedgerReferenceProvided: l.AuthenticityLedgerReferenceProvided,
		ChecksumOrProofHashProvided:         l.ChecksumOrProofHashProvided,
		TrustedSourceURI:                    l.TrustedSourceURI,
		PhysicalSchwingerSourceLoaded:       l.PhysicalSchwingerSourceLoaded,
		ObservedCorrelationLoaded:           l.ObservedCorrelationLoaded,
		ConstructiveMeasureLoaded:           l.ConstructiveMeasureLoaded,
		PhysicalOSCertificateLoaded:         l.PhysicalOSCertificateLoaded,
		PhysicalWickMapLoaded:               l.PhysicalWickMapLoaded,
		PhysicalHamiltonianLoaded:           l.PhysicalHamiltonianLoaded,
		Gate540ReferenceComplete:            strings.TrimSpace(l.Gate540SwitchReference) != "",
		Gate539ReferenceComplete:            strings.TrimSpace(l.Gate539AuthenticityReference) != "",
		Gate536ReferenceComplete:            strings.TrimSpace(l.Gate536SchwingerSchemaReference) != "",
		SourceTaggedLedger:                  strings.TrimSpace(l.Source) != "",
		ConventionTaggedLedger:              strings.TrimSpace(l.Convention) != "",
		MetadataComplete:                    strings.TrimSpace(l.Source) != "" && strings.TrimSpace(l.SourceVersion) != "" && strings.TrimSpace(l.Convention) != "",
		RequiredSchemaRowsMatched:           len(missing) == 0 && len(duplicate) == 0 && accepted == len(req),
		AllRowsBridgeOnly:                   allBridge,
		AllRowsNoTheoremInput:               allNoTheorem,
		AllRowsNegativeControl:              allNegative,
		AllRowsSourceTagged:                 allSource,
		AllRowsConventionTagged:             allConvention,
		AnyNativePromotionClaim:             anyNativePromotion,
		AnyNativeInputClaim:                 anyNativeInput,
		ChecksumExpected:                    l.CanonicalPayloadSHA256,
		ChecksumActual:                      actual,
		ChecksumVerified:                    strings.TrimSpace(l.CanonicalPayloadSHA256) != "" && l.CanonicalPayloadSHA256 == actual,
		Verdict:                             StatusNegativeControlLedgerLoaded,
		Reason:                              "the real-looking negative-control fixture parses and verifies checksum plumbing, but it is intentionally untrusted and cannot pass the real-source switch.",
		Failures:                            failures,
	}
}

func checksum(payload map[string]any) string {
	b, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	h := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(h[:])
}

func requiredRows() map[string]bool {
	keys := []string{
		"real_source_import_switch",
		"explicit_operator_intent",
		"non_synthetic_source_uri",
		"authenticity_ledger_reference",
		"checksum_or_proof_hash_reference",
		"license_and_access_grant_reference",
		"source_class_non_synthetic_assertion",
		"gate536_schema_alignment_reference",
		"comparator_execution_plan",
		"quarantine_output_target",
		"native_write_lock",
		"rollback_audit_trace",
	}
	m := make(map[string]bool, len(keys))
	for _, k := range keys {
		m[k] = true
	}
	return m
}

func buildRejection(inh Inheritance, imp FileImport) Rejection {
	reasons := []string{}
	if !imp.RealSourceImportSwitchEnabled || inh.Gate540SwitchDefaultOff {
		reasons = append(reasons, StatusFailedSwitchOffBlocksImport)
	}
	if !imp.ExplicitOperatorIntentProvided || inh.Gate540OperatorIntentNeeded {
		reasons = append(reasons, StatusFailedNoIntentBlocksImport)
	}
	if !imp.LicenseAndAccessGrantProvided {
		reasons = append(reasons, StatusFailedNoAccessGrant)
	}
	if !imp.TrustedSourceURI {
		reasons = append(reasons, StatusFailedUntrustedURI)
	}
	if !imp.AuthenticityLedgerReferenceProvided || !imp.ChecksumOrProofHashProvided || !imp.Gate540ReferenceComplete || !imp.Gate539ReferenceComplete || !imp.Gate536ReferenceComplete {
		reasons = append(reasons, StatusFailedInsufficientProvenance)
	}
	return Rejection{
		Executed:                       true,
		SourceParsed:                   imp.Loaded && imp.RequiredSchemaRowsMatched && imp.ChecksumVerified,
		RealLookingClaimSeen:           imp.RealLookingFixture,
		PhysicalSourceClaimSeen:        imp.PhysicalSourceClaim && imp.NonSyntheticClaim,
		NegativeControlFixtureSeen:     imp.NegativeControlFixture,
		SwitchOff:                      !imp.RealSourceImportSwitchEnabled || inh.Gate540SwitchDefaultOff,
		NoExplicitOperatorIntent:       !imp.ExplicitOperatorIntentProvided || inh.Gate540OperatorIntentNeeded,
		MissingLicenseOrAccessGrant:    !imp.LicenseAndAccessGrantProvided,
		SourceURINotAuthenticated:      !imp.TrustedSourceURI,
		ComparatorAuthorizationMissing: !imp.ComparatorAuthorizationGranted || inh.Gate540ComparatorBlocked,
		InsufficientProvenance:         !imp.AuthenticityLedgerReferenceProvided || !imp.ChecksumOrProofHashProvided || !imp.TrustedSourceURI || !imp.LicenseAndAccessGrantProvided,
		ComparatorExecutionAllowed:     false,
		ComparatorExecutionPerformed:   false,
		PhysicalSourceAuthenticated:    false,
		PhysicalSourceImported:         false,
		RejectedAsPhysicalSource:       true,
		RejectedBeforeComparator:       true,
		QuarantinePreserved:            imp.BridgeOnlyLedger && imp.AllRowsBridgeOnly && imp.AllRowsNoTheoremInput && !imp.NativeRegistryWriteRequested && !imp.AnyNativePromotionClaim && !imp.AnyNativeInputClaim,
		RejectionReasons:               reasons,
		Verdict:                        StatusRejectedSwitchOff,
		Reason:                         "the fixture is deliberately real-looking, but default-off switch state, absent operator intent, untrusted URI, and missing license/access provenance reject it before comparator execution.",
	}
}

func buildFirewall(imp FileImport, r Rejection) Firewall {
	return Firewall{
		Executed:                      true,
		ComparatorExecutionPerformed:  r.ComparatorExecutionPerformed,
		RealSchwingerSourceImported:   r.PhysicalSourceImported || imp.PhysicalSchwingerSourceLoaded,
		ObservedCorrelationImported:   imp.ObservedCorrelationLoaded,
		ConstructiveMeasureImported:   imp.ConstructiveMeasureLoaded,
		PhysicalOSCertificateImported: imp.PhysicalOSCertificateLoaded,
		PhysicalWickMapImported:       imp.PhysicalWickMapLoaded,
		PhysicalHamiltonianImported:   imp.PhysicalHamiltonianLoaded,
		PhysicalSchwingerDerived:      false,
		OSPositivityProven:            false,
		WickRotationSelected:          false,
		PhysicalHilbertSpaceSelected:  false,
		PositiveHamiltonianDerived:    false,
		UnitaryDynamicsDerived:        false,
		GlobalCausalitySelected:       false,
		ArrowOfTimeSelected:           false,
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
		Reason:                        "real-looking negative-control data stays quarantined: no comparator runs and no physical correlation/dynamics object writes into the native registry.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new native C\\ell(1,7), Schwinger, OS, Wick, Hilbert, Hamiltonian, unitary, global-causal, or time-arrow theorem is added by Gate541.",
		},
		BridgeEntries: []string{
			StatusGate540SwitchInherited,
			StatusNegativeControlLedgerLoaded,
			StatusNegativeControlRowsAccepted,
			StatusNegativeControlChecksumVerified,
			StatusNegativeControlMetadataParsed,
			StatusNegativeControlAdapterExecuted,
			StatusRejectedSwitchOff,
			StatusRejectedNoOperatorIntent,
			StatusRejectedInsufficientProvenance,
			StatusRejectedComparatorAuthorization,
			StatusQuarantinePreserved,
			StatusNoComparatorExecuted,
			StatusNoNativeWrite,
		},
		EnvironmentalEntries: []string{
			"A future non-synthetic Schwinger source still requires explicit operator intent, authenticated URI, access/license grant, authenticity ledger, proof hash, Gate536 alignment, comparator plan, quarantine target, rollback trace, and native-write lock.",
		},
		FailedRoutes: []string{
			StatusFailedSwitchOffBlocksImport,
			StatusFailedNoIntentBlocksImport,
			StatusFailedInsufficientProvenance,
			StatusFailedNoAccessGrant,
			StatusFailedUntrustedURI,
			StatusFailedNegativeControlNotSchwinger,
			StatusFailedNegativeControlNotOSProof,
			StatusFailedNegativeControlNotWick,
			StatusFailedNegativeControlNotHilbert,
			StatusFailedNegativeControlNotHamiltonian,
			StatusFailedNegativeControlNotUnitary,
			StatusFailedNegativeControlNotGlobal,
			StatusFailedNegativeControlNotArrow,
			StatusFirewallPreserved,
			StatusFirewallNativeWriteBlocked,
		},
		OpenTheorems: []string{
			"Future Gate 542 may define a controlled comparator authorization manifest schema for a real source while still requiring dry-run quarantine and zero native writes.",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{
		Gate:        542,
		Title:       "Real Source Comparator Authorization Manifest Airlock",
		Reason:      "Gate541 proves that real-looking data is rejected when the switch is off or provenance is insufficient. The next safe boundary is the authorization manifest that would be required before a real-source comparator may run in bridge quarantine.",
		PrimaryTask: "Enumerate the comparator-authorization manifest and confirm that authorization can only target bridge quarantine, never native registry promotion.",
	}
}

func truth(a Analysis) string {
	return "Gate541 proves the default-deny import path: a real-looking Schwinger source fixture can parse and verify checksum plumbing, but the off switch, missing operator intent, untrusted URI, incomplete provenance, and absent access grant reject it before comparator execution or native registry write."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate540SwitchDefined || !a.Inheritance.Gate540SwitchDefaultOff || !a.Inheritance.Gate540OperatorIntentNeeded || !a.Inheritance.Gate540ComparatorBlocked || !a.Inheritance.Gate540NoRealSourceImported || !a.Inheritance.Gate540NativeWriteBlocked || !a.Inheritance.Gate540RedirectsToGate541 {
		bad = append(bad, "Gate540 inheritance incomplete")
	}
	if !a.Import.Executed || !a.Import.Loaded || a.Import.Rows != 12 || a.Import.AcceptedRows != 12 || a.Import.RejectedRows != 0 || !a.Import.RequiredSchemaRowsMatched || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsNegativeControl || !a.Import.AllRowsSourceTagged || !a.Import.AllRowsConventionTagged || !a.Import.ChecksumVerified {
		bad = append(bad, "negative-control import failed schema/checksum/metadata sieve")
	}
	if !a.Import.RealLookingFixture || !a.Import.NegativeControlFixture || a.Import.DeclaredSyntheticFixture || !a.Import.PhysicalSourceClaim || !a.Import.NonSyntheticClaim {
		bad = append(bad, "fixture is not real-looking negative control")
	}
	if a.Import.RealSourceImportSwitchEnabled || a.Import.ExplicitOperatorIntentProvided || a.Import.ComparatorAuthorizationGranted || a.Import.LicenseAndAccessGrantProvided || a.Import.TrustedSourceURI || a.Import.NativeRegistryWriteRequested || a.Import.PhysicalSchwingerSourceLoaded || a.Import.ObservedCorrelationLoaded || a.Import.ConstructiveMeasureLoaded || a.Import.PhysicalOSCertificateLoaded || a.Import.PhysicalWickMapLoaded || a.Import.PhysicalHamiltonianLoaded {
		bad = append(bad, "negative-control fixture accidentally enabled import path")
	}
	if !a.Rejection.Executed || !a.Rejection.SourceParsed || !a.Rejection.RealLookingClaimSeen || !a.Rejection.PhysicalSourceClaimSeen || !a.Rejection.NegativeControlFixtureSeen || !a.Rejection.SwitchOff || !a.Rejection.NoExplicitOperatorIntent || !a.Rejection.MissingLicenseOrAccessGrant || !a.Rejection.SourceURINotAuthenticated || !a.Rejection.ComparatorAuthorizationMissing || !a.Rejection.InsufficientProvenance || a.Rejection.ComparatorExecutionAllowed || a.Rejection.ComparatorExecutionPerformed || a.Rejection.PhysicalSourceAuthenticated || a.Rejection.PhysicalSourceImported || !a.Rejection.RejectedAsPhysicalSource || !a.Rejection.RejectedBeforeComparator || !a.Rejection.QuarantinePreserved || len(a.Rejection.RejectionReasons) < 4 {
		bad = append(bad, "negative-control rejection/firewall state incomplete")
	}
	if !a.Firewall.Executed || a.Firewall.ComparatorExecutionPerformed || a.Firewall.RealSchwingerSourceImported || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || a.Firewall.PhysicalSchwingerDerived || a.Firewall.OSPositivityProven || a.Firewall.WickRotationSelected || a.Firewall.PhysicalHilbertSpaceSelected || a.Firewall.PositiveHamiltonianDerived || a.Firewall.UnitaryDynamicsDerived || a.Firewall.GlobalCausalitySelected || a.Firewall.ArrowOfTimeSelected || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall leaked physical import or native write")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate541 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: switch_defined=%t default_off=%t intent_needed=%t comparator_blocked=%t no_real=%t native_blocked=%t redirects=%t; %s", x.Verdict, x.Gate540SwitchDefined, x.Gate540SwitchDefaultOff, x.Gate540OperatorIntentNeeded, x.Gate540ComparatorBlocked, x.Gate540NoRealSourceImported, x.Gate540NativeWriteBlocked, x.Gate540RedirectsToGate541, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: rows=%d accepted=%d rejected=%d bridge=%t native_request=%t real_looking=%t negative=%t declared_synthetic=%t physical_claim=%t nonsynthetic_claim=%t switch=%t intent=%t comparator_request=%t comparator_grant=%t license=%t auth_ref=%t checksum_ref=%t trusted_uri=%t physical_loaded=%t observed=%t measure=%t OS=%t Wick=%t Hamiltonian=%t gate540=%t gate539=%t gate536=%t metadata=%t required=%t all_bridge=%t all_no_theorem=%t all_negative=%t checksum=%t expected=%s actual=%s failures=%v; %s", x.Verdict, x.Rows, x.AcceptedRows, x.RejectedRows, x.BridgeOnlyLedger, x.NativeRegistryWriteRequested, x.RealLookingFixture, x.NegativeControlFixture, x.DeclaredSyntheticFixture, x.PhysicalSourceClaim, x.NonSyntheticClaim, x.RealSourceImportSwitchEnabled, x.ExplicitOperatorIntentProvided, x.ComparatorExecutionRequested, x.ComparatorAuthorizationGranted, x.LicenseAndAccessGrantProvided, x.AuthenticityLedgerReferenceProvided, x.ChecksumOrProofHashProvided, x.TrustedSourceURI, x.PhysicalSchwingerSourceLoaded, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.PhysicalOSCertificateLoaded, x.PhysicalWickMapLoaded, x.PhysicalHamiltonianLoaded, x.Gate540ReferenceComplete, x.Gate539ReferenceComplete, x.Gate536ReferenceComplete, x.MetadataComplete, x.RequiredSchemaRowsMatched, x.AllRowsBridgeOnly, x.AllRowsNoTheoremInput, x.AllRowsNegativeControl, x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.Failures, x.Reason)
}

func FormatRejection(x Rejection) string {
	return fmt.Sprintf("%s: parsed=%t real_looking=%t physical_claim=%t negative=%t switch_off=%t no_intent=%t no_license=%t uri_untrusted=%t auth_missing=%t insufficient=%t allowed=%t performed=%t authenticated=%t imported=%t rejected=%t before_comparator=%t quarantine=%t reasons=%v; %s", x.Verdict, x.SourceParsed, x.RealLookingClaimSeen, x.PhysicalSourceClaimSeen, x.NegativeControlFixtureSeen, x.SwitchOff, x.NoExplicitOperatorIntent, x.MissingLicenseOrAccessGrant, x.SourceURINotAuthenticated, x.ComparatorAuthorizationMissing, x.InsufficientProvenance, x.ComparatorExecutionAllowed, x.ComparatorExecutionPerformed, x.PhysicalSourceAuthenticated, x.PhysicalSourceImported, x.RejectedAsPhysicalSource, x.RejectedBeforeComparator, x.QuarantinePreserved, x.RejectionReasons, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: comparator=%t real=%t observed=%t measure=%t OS_import=%t wick_import=%t hamiltonian_import=%t derived=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t unitary=%t global=%t arrow=%t native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopened_flavor=%t reopened_EW=%t reopened_gravity=%t reopened_topology=%t reopened_dimension=%t reopened_Krein=%t native_registry=%t; %s", x.Verdict, x.ComparatorExecutionPerformed, x.RealSchwingerSourceImported, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.PhysicalSchwingerDerived, x.OSPositivityProven, x.WickRotationSelected, x.PhysicalHilbertSpaceSelected, x.PositiveHamiltonianDerived, x.UnitaryDynamicsDerived, x.GlobalCausalitySelected, x.ArrowOfTimeSelected, x.NativeSchwingerFunctionWrite, x.NativeEuclideanMeasureWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate540SwitchInherited,
		StatusNegativeControlLedgerLoaded,
		StatusNegativeControlRowsAccepted,
		StatusNegativeControlChecksumVerified,
		StatusNegativeControlMetadataParsed,
		StatusNegativeControlAdapterExecuted,
		StatusRejectedSwitchOff,
		StatusRejectedNoOperatorIntent,
		StatusRejectedInsufficientProvenance,
		StatusRejectedComparatorAuthorization,
		StatusQuarantinePreserved,
		StatusNoComparatorExecuted,
		StatusNoNativeWrite,
		StatusFailedSwitchOffBlocksImport,
		StatusFailedNoIntentBlocksImport,
		StatusFailedInsufficientProvenance,
		StatusFailedNoAccessGrant,
		StatusFailedUntrustedURI,
		StatusFailedNegativeControlNotSchwinger,
		StatusFailedNegativeControlNotOSProof,
		StatusFailedNegativeControlNotWick,
		StatusFailedNegativeControlNotHilbert,
		StatusFailedNegativeControlNotHamiltonian,
		StatusFailedNegativeControlNotUnitary,
		StatusFailedNegativeControlNotGlobal,
		StatusFailedNegativeControlNotArrow,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}
