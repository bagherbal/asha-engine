// Package generation2syntheticsourceauthenticityadapter implements Gate 539:
// Synthetic Source-Authenticity Ledger Adapter Rejection Dry Run.
//
// Gate 538 defined the provenance/integrity sieve required before any
// non-synthetic Schwinger source may be treated as an auditable bridge input.
// Gate 539 loads a fully synthetic 13-row authenticity ledger, verifies the
// checksum, provenance rows, source/convention tags, and quarantine flags, and
// then deliberately rejects the fixture as physical evidence. This gate tests
// source-authentication plumbing; it does not authenticate a universe.
package generation2syntheticsourceauthenticityadapter

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

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2schwingersourceauthenticityairlock"
)

const (
	AuditID       = "GATE539-SYNTHETIC-SOURCE-AUTHENTICITY-LEDGER-ADAPTER-REJECTION-DRY-RUN"
	DefaultLedger = "data/synthetic_source_authenticity_ledger_gate539.json"

	StatusGate538AirlockInherited              = "CONDITIONAL_SUPPORT_GATE538_SOURCE_AUTHENTICITY_AIRLOCK_INHERITED"
	StatusSyntheticAuthenticityLedgerLoaded    = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_LOADED"
	StatusAuthenticityRowsAccepted             = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_13_SCHEMA_ROWS_ACCEPTED"
	StatusChecksumIntegrityVerified            = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_CHECKSUM_VERIFIED"
	StatusProvenanceRowsParsed                 = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_PROVENANCE_ROWS_PARSED"
	StatusMetadataSieveEnforced                = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_METADATA_SIEVE_ENFORCED"
	StatusSyntheticAuthenticityAdapterExecuted = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_ADAPTER_EXECUTED"
	StatusSyntheticSourceRejectedAsPhysical    = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_FIXTURE_REJECTED_AS_PHYSICAL_SOURCE"
	StatusQuarantineTagsPreserved              = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_QUARANTINE_TAGS_PRESERVED"
	StatusNoRealSourceImported                 = "CONDITIONAL_SUPPORT_NO_REAL_SCHWINGER_SOURCE_IMPORTED_IN_GATE539"
	StatusNativePromotionRejected              = "CONDITIONAL_SUPPORT_SYNTHETIC_SOURCE_AUTHENTICITY_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing              = "FAILED_ROUTE_GATE539_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_MISSING"
	StatusFailedMetadataIncomplete         = "FAILED_ROUTE_GATE539_SYNTHETIC_SOURCE_AUTHENTICITY_METADATA_INCOMPLETE"
	StatusFailedSchemaRowsIncomplete       = "FAILED_ROUTE_GATE539_SYNTHETIC_SOURCE_AUTHENTICITY_SCHEMA_ROWS_INCOMPLETE"
	StatusFailedChecksumMismatch           = "FAILED_ROUTE_GATE539_SYNTHETIC_SOURCE_AUTHENTICITY_CHECKSUM_MISMATCH"
	StatusFailedSyntheticNotPhysicalSource = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_CANNOT_AUTHENTICATE_PHYSICAL_SOURCE"
	StatusFailedSyntheticNotSchwinger      = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticNotOSProof        = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedSyntheticNotWick           = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticNotHilbert        = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticNotHamiltonian    = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticNotUnitary        = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticNotGlobal         = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticNotArrow          = "FAILED_ROUTE_SYNTHETIC_SOURCE_AUTHENTICITY_LEDGER_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_GATE539_SYNTHETIC_SOURCE_AUTHENTICITY_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked       = "FIREWALL_BLOCKED_GATE539_REAL_SOURCE_NATIVE_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate538AirlockDefined       bool
	Gate538SchemaRowsEnumerated bool
	Gate538RequiredRows         int
	Gate538BridgeOnlyRows       int
	Gate538ComparatorRows       int
	Gate538NativeWriteRows      int
	Gate538DiscriminatorDefined bool
	Gate538SyntheticRejected    bool
	Gate538ComparatorBlocked    bool
	Gate538NoRealSourceImported bool
	Gate538NativeWriteBlocked   bool
	Gate539SyntheticRedirect    bool

	Verdict, Reason string
}

type AuthenticityRowMetadata struct {
	SchemaKey        string `json:"schema_key"`
	Source           string `json:"source"`
	SourceVersion    string `json:"source_version"`
	Convention       string `json:"convention"`
	ValueKind        string `json:"value_kind"`
	BridgeOnly       bool   `json:"bridge_only"`
	ComparatorOnly   bool   `json:"comparator_only"`
	Synthetic        bool   `json:"synthetic"`
	PhysicalClaim    bool   `json:"physical_claim"`
	Observed         bool   `json:"observed"`
	NoTheoremInput   bool   `json:"no_theorem_input"`
	NativePromotion  bool   `json:"native_promotion"`
	NativeInputClaim bool   `json:"native_input_claim"`
	Value            string `json:"value"`
}

type SyntheticSourceAuthenticityLedger struct {
	Gate                                    int                       `json:"gate"`
	LedgerName                              string                    `json:"ledger_name"`
	Description                             string                    `json:"description"`
	Gate538AirlockReference                 string                    `json:"gate538_airlock_reference"`
	Gate537SchwingerFixtureReference        string                    `json:"gate537_schwinger_fixture_reference"`
	BridgeOnly                              bool                      `json:"bridge_only"`
	NativeRegistryWrite                     bool                      `json:"native_registry_write"`
	SyntheticFixture                        bool                      `json:"synthetic_fixture"`
	PhysicalSourceClaim                     bool                      `json:"physical_source_claim"`
	NonSyntheticPhysicalOrConstructiveClaim bool                      `json:"non_synthetic_physical_or_constructive_claim"`
	RealSchwingerSourceImported             bool                      `json:"real_schwinger_source_imported"`
	ObservedCorrelationLoaded               bool                      `json:"observed_correlation_loaded"`
	ConstructiveMeasureLoaded               bool                      `json:"constructive_measure_loaded"`
	PhysicalOSCertificateLoaded             bool                      `json:"physical_os_certificate_loaded"`
	PhysicalWickMapLoaded                   bool                      `json:"physical_wick_map_loaded"`
	PhysicalHamiltonianLoaded               bool                      `json:"physical_hamiltonian_loaded"`
	Source                                  string                    `json:"source"`
	SourceVersion                           string                    `json:"source_version"`
	Convention                              string                    `json:"convention"`
	CanonicalPayload                        map[string]any            `json:"canonical_payload"`
	CanonicalPayloadSHA256                  string                    `json:"canonical_payload_sha256"`
	Rows                                    []AuthenticityRowMetadata `json:"rows"`
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

	BridgeOnlyLedger             bool
	SyntheticFixture             bool
	PhysicalSourceClaim          bool
	NonSyntheticClaim            bool
	RealSchwingerSourceImported  bool
	ObservedCorrelationLoaded    bool
	ConstructiveMeasureLoaded    bool
	PhysicalOSCertificateLoaded  bool
	PhysicalWickMapLoaded        bool
	PhysicalHamiltonianLoaded    bool
	NativeRegistryWriteRequested bool
	Gate538ReferenceComplete     bool
	Gate537ReferenceComplete     bool
	SourceTaggedLedger           bool
	ConventionTaggedLedger       bool
	MetadataComplete             bool
	RequiredSchemaRowsMatched    bool
	AllRowsBridgeOnly            bool
	AllRowsNoTheoremInput        bool
	AllRowsSynthetic             bool
	AllRowsSourceTagged          bool
	AllRowsConventionTagged      bool
	AnyPhysicalClaim             bool
	AnyObservedClaim             bool
	AnyNativePromotionClaim      bool
	ChecksumExpected             string
	ChecksumActual               string
	ChecksumVerified             bool

	Verdict, Reason string
	Failures        []string
}

type AdapterOutput struct {
	Executed  bool
	Attempted bool
	Ready     bool

	RowsParsed                         int
	RequiredRowsMatched                int
	BridgeOnlyRows                     int
	NoTheoremInputRows                 int
	SourceTaggedRows                   int
	ConventionTaggedRows               int
	SyntheticRows                      int
	PhysicalClaimRows                  int
	ObservedRows                       int
	NativePromotionRows                int
	ComparatorRows                     int
	ChecksumVerified                   bool
	ImmutableSourceParsed              bool
	LicenseParsed                      bool
	ConstructionProvenanceParsed       bool
	RenormalizationParsed              bool
	Gate536AlignmentParsed             bool
	CovarianceProvenanceParsed         bool
	OSCertificateProvenanceParsed      bool
	WickIEpsilonParsed                 bool
	HamiltonianDomainParsed            bool
	UncertaintyLedgerParsed            bool
	QuarantineParsed                   bool
	AuthenticityPlumbingVerified       bool
	SyntheticFixtureRejectedAsPhysical bool
	PhysicalSourceAuthenticated        bool
	PhysicalSchwingerFunctionsDerived  bool
	PhysicalOSPositivityProven         bool
	WickRotationGranted                bool
	PhysicalHilbertSpaceSelected       bool
	PositiveEnergyHamiltonianDerived   bool
	UnitaryRealTimeDynamicsDerived     bool
	GlobalHyperbolicityGranted         bool
	ArrowOfTimeSelected                bool

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
	SyntheticFixtureOnly          bool
	SourceAuthenticatedAsPhysical bool
	FileRowsNative                bool
	AdapterOutputsNative          bool
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
	Inheritance Inheritance
	Import      FileImport
	Output      AdapterOutput
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
	cache.Once.Do(func() { cache.a, cache.err = BuildFromFile(DefaultLedger) })
	return cache.a, cache.err
}
func Build() (Analysis, error) { return BuildFromFile(DefaultLedger) }

func BuildFromFile(path string) (Analysis, error) {
	g538, err := generation2schwingersourceauthenticityairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate538 Schwinger source-authenticity airlock: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g538)}
	ledger, imp := loadLedger(path, g538)
	a.Import = imp
	if imp.Loaded && imp.MetadataComplete && imp.RequiredSchemaRowsMatched && imp.ChecksumVerified {
		a.Output = runAdapter(ledger, imp)
	} else if !imp.Loaded {
		a.Output = AdapterOutput{Executed: true, Attempted: false, Verdict: StatusFailedLedgerMissing, Reason: "explicit Gate539 synthetic source-authenticity ledger was not found", Failures: []string{StatusFailedLedgerMissing}}
	} else {
		failure := StatusFailedMetadataIncomplete
		if !imp.RequiredSchemaRowsMatched {
			failure = StatusFailedSchemaRowsIncomplete
		}
		if !imp.ChecksumVerified {
			failure = StatusFailedChecksumMismatch
		}
		a.Output = AdapterOutput{Executed: true, Attempted: true, RowsParsed: imp.Rows, Verdict: failure, Reason: "Gate539 synthetic source-authenticity ledger did not satisfy the Gate538 schema/checksum/quarantine domain", Failures: []string{failure}}
	}
	a.Firewall = buildFirewall(a.Import, a.Output)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2schwingersourceauthenticityairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                    true,
		Gate538AirlockDefined:       g.Schema.Executed && g.Schema.ImmutableSourceRequired && g.Schema.ChecksumRequired,
		Gate538SchemaRowsEnumerated: len(g.Schema.Rows) == 13 && g.Schema.RequiredRows == 13,
		Gate538RequiredRows:         g.Schema.RequiredRows,
		Gate538BridgeOnlyRows:       g.Schema.BridgeOnlyRows,
		Gate538ComparatorRows:       g.Schema.ComparatorRows,
		Gate538NativeWriteRows:      g.Schema.NativeWriteRows,
		Gate538DiscriminatorDefined: g.Discriminator.Executed && g.Discriminator.SyntheticLedgerRecognized,
		Gate538SyntheticRejected:    !g.Discriminator.SyntheticLedgerAcceptedAsPhysical && !g.Discriminator.PhysicalSchwingerAuthenticated,
		Gate538ComparatorBlocked:    !g.Guard.ComparatorExecutionPerformed,
		Gate538NoRealSourceImported: !g.Discriminator.NonSyntheticSourceLoaded && !g.Discriminator.ObservedCorrelationLoaded && !g.Discriminator.ConstructiveMeasureLoaded,
		Gate538NativeWriteBlocked:   !g.Firewall.NativeRegistryWritten && !g.Firewall.NativeSchwingerWrite && !g.Firewall.NativeOSProofWrite && !g.Firewall.NativeWickWrite && !g.Firewall.NativeHamiltonianWrite,
		Gate539SyntheticRedirect:    g.Next.Gate == 539,
		Verdict:                     StatusGate538AirlockInherited,
		Reason:                      "Gate539 inherits Gate538's 13-row source-authenticity sieve and executes only a synthetic rejection dry run.",
	}
}

func loadLedger(path string, g538 generation2schwingersourceauthenticityairlock.Analysis) (SyntheticSourceAuthenticityLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved, Verdict: StatusFailedLedgerMissing, Reason: "ledger not loaded", Failures: []string{StatusFailedLedgerMissing}}
	b, err := os.ReadFile(resolved)
	if err != nil {
		return SyntheticSourceAuthenticityLedger{}, imp
	}
	var ledger SyntheticSourceAuthenticityLedger
	if err := json.Unmarshal(b, &ledger); err != nil {
		imp.Loaded = true
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		return ledger, imp
	}
	imp.Loaded = true
	imp.Rows = len(ledger.Rows)
	imp.BridgeOnlyLedger = ledger.BridgeOnly
	imp.SyntheticFixture = ledger.SyntheticFixture
	imp.PhysicalSourceClaim = ledger.PhysicalSourceClaim
	imp.NonSyntheticClaim = ledger.NonSyntheticPhysicalOrConstructiveClaim
	imp.RealSchwingerSourceImported = ledger.RealSchwingerSourceImported
	imp.ObservedCorrelationLoaded = ledger.ObservedCorrelationLoaded
	imp.ConstructiveMeasureLoaded = ledger.ConstructiveMeasureLoaded
	imp.PhysicalOSCertificateLoaded = ledger.PhysicalOSCertificateLoaded
	imp.PhysicalWickMapLoaded = ledger.PhysicalWickMapLoaded
	imp.PhysicalHamiltonianLoaded = ledger.PhysicalHamiltonianLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.Gate538ReferenceComplete = strings.TrimSpace(ledger.Gate538AirlockReference) != ""
	imp.Gate537ReferenceComplete = strings.TrimSpace(ledger.Gate537SchwingerFixtureReference) != ""
	imp.SourceTaggedLedger = strings.TrimSpace(ledger.Source) != "" && strings.TrimSpace(ledger.SourceVersion) != ""
	imp.ConventionTaggedLedger = strings.TrimSpace(ledger.Convention) != ""
	imp.ChecksumExpected = strings.TrimSpace(ledger.CanonicalPayloadSHA256)
	imp.ChecksumActual = checksumPayload(ledger.CanonicalPayload)
	imp.ChecksumVerified = imp.ChecksumExpected != "" && imp.ChecksumActual == imp.ChecksumExpected
	imp.AllRowsBridgeOnly = true
	imp.AllRowsNoTheoremInput = true
	imp.AllRowsSynthetic = true
	imp.AllRowsSourceTagged = true
	imp.AllRowsConventionTagged = true

	required := requiredRowSet(g538)
	seen := map[string]int{}
	for _, row := range ledger.Rows {
		if validateRowMetadata(row) {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
		seen[row.SchemaKey]++
		imp.AllRowsBridgeOnly = imp.AllRowsBridgeOnly && row.BridgeOnly
		imp.AllRowsNoTheoremInput = imp.AllRowsNoTheoremInput && row.NoTheoremInput
		imp.AllRowsSynthetic = imp.AllRowsSynthetic && row.Synthetic
		imp.AllRowsSourceTagged = imp.AllRowsSourceTagged && strings.TrimSpace(row.Source) != "" && strings.TrimSpace(row.SourceVersion) != ""
		imp.AllRowsConventionTagged = imp.AllRowsConventionTagged && strings.TrimSpace(row.Convention) != ""
		imp.AnyPhysicalClaim = imp.AnyPhysicalClaim || row.PhysicalClaim
		imp.AnyObservedClaim = imp.AnyObservedClaim || row.Observed
		imp.AnyNativePromotionClaim = imp.AnyNativePromotionClaim || row.NativePromotion || row.NativeInputClaim
	}
	for name := range required {
		if seen[name] == 0 {
			imp.MissingRequiredRows = append(imp.MissingRequiredRows, name)
		}
		if seen[name] > 1 {
			imp.DuplicateRows = append(imp.DuplicateRows, name)
		}
	}
	sort.Strings(imp.MissingRequiredRows)
	sort.Strings(imp.DuplicateRows)
	imp.RequiredSchemaRowsMatched = len(required) == 13 && imp.Rows == 13 && imp.AcceptedRows == 13 && len(imp.MissingRequiredRows) == 0 && len(imp.DuplicateRows) == 0
	imp.MetadataComplete = ledger.Gate == 539 && imp.RequiredSchemaRowsMatched && imp.BridgeOnlyLedger && imp.SyntheticFixture && !imp.PhysicalSourceClaim && !imp.NonSyntheticClaim && !imp.RealSchwingerSourceImported && !imp.ObservedCorrelationLoaded && !imp.ConstructiveMeasureLoaded && !imp.PhysicalOSCertificateLoaded && !imp.PhysicalWickMapLoaded && !imp.PhysicalHamiltonianLoaded && !imp.NativeRegistryWriteRequested && imp.Gate538ReferenceComplete && imp.Gate537ReferenceComplete && imp.SourceTaggedLedger && imp.ConventionTaggedLedger && imp.AllRowsBridgeOnly && imp.AllRowsNoTheoremInput && imp.AllRowsSynthetic && imp.AllRowsSourceTagged && imp.AllRowsConventionTagged && !imp.AnyPhysicalClaim && !imp.AnyObservedClaim && !imp.AnyNativePromotionClaim
	if imp.MetadataComplete && imp.ChecksumVerified {
		imp.Verdict = strings.Join([]string{StatusSyntheticAuthenticityLedgerLoaded, StatusAuthenticityRowsAccepted, StatusChecksumIntegrityVerified, StatusMetadataSieveEnforced}, ";")
		imp.Reason = "Gate539 synthetic source-authenticity ledger loaded with exactly the Gate538 13-row schema, verified canonical-payload checksum, source/convention tags, bridge_only=true, no_theorem_input=true, synthetic=true everywhere, and no physical/native-promotion claims."
		imp.Failures = nil
	} else {
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = "ledger metadata, row flags, or checksum violate the Gate538 source-authenticity airlock"
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		if !imp.RequiredSchemaRowsMatched {
			imp.Failures = append(imp.Failures, StatusFailedSchemaRowsIncomplete)
		}
		if !imp.ChecksumVerified {
			imp.Failures = append(imp.Failures, StatusFailedChecksumMismatch)
		}
	}
	return ledger, imp
}

func requiredRowSet(g538 generation2schwingersourceauthenticityairlock.Analysis) map[string]bool {
	out := map[string]bool{}
	for _, row := range g538.Schema.Rows {
		out[row.Name] = true
	}
	return out
}

func validateRowMetadata(row AuthenticityRowMetadata) bool {
	return strings.TrimSpace(row.SchemaKey) != "" && strings.TrimSpace(row.Source) != "" && strings.TrimSpace(row.SourceVersion) != "" && strings.TrimSpace(row.Convention) != "" && row.BridgeOnly && row.Synthetic && !row.PhysicalClaim && !row.Observed && row.NoTheoremInput && !row.NativePromotion && !row.NativeInputClaim
}

func checksumPayload(payload map[string]any) string {
	b, err := json.Marshal(canonicalize(payload))
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func canonicalize(v any) any { return v }

func runAdapter(ledger SyntheticSourceAuthenticityLedger, imp FileImport) AdapterOutput {
	rowMap := map[string]AuthenticityRowMetadata{}
	bridge, noTheorem, source, convention, synthetic, physical, observed, native, comparator := 0, 0, 0, 0, 0, 0, 0, 0, 0
	for _, row := range ledger.Rows {
		rowMap[row.SchemaKey] = row
		if row.BridgeOnly {
			bridge++
		}
		if row.NoTheoremInput {
			noTheorem++
		}
		if strings.TrimSpace(row.Source) != "" && strings.TrimSpace(row.SourceVersion) != "" {
			source++
		}
		if strings.TrimSpace(row.Convention) != "" {
			convention++
		}
		if row.Synthetic {
			synthetic++
		}
		if row.PhysicalClaim {
			physical++
		}
		if row.Observed {
			observed++
		}
		if row.NativePromotion || row.NativeInputClaim {
			native++
		}
		if row.ComparatorOnly {
			comparator++
		}
	}
	out := AdapterOutput{
		Executed: true, Attempted: true, Ready: true,
		RowsParsed: imp.Rows, RequiredRowsMatched: imp.AcceptedRows,
		BridgeOnlyRows: bridge, NoTheoremInputRows: noTheorem, SourceTaggedRows: source, ConventionTaggedRows: convention, SyntheticRows: synthetic, PhysicalClaimRows: physical, ObservedRows: observed, NativePromotionRows: native, ComparatorRows: comparator,
		ChecksumVerified:                   imp.ChecksumVerified,
		ImmutableSourceParsed:              strings.TrimSpace(rowMap["immutable_source_identifier"].Value) != "",
		LicenseParsed:                      strings.TrimSpace(rowMap["license_and_access_rights"].Value) != "",
		ConstructionProvenanceParsed:       strings.TrimSpace(rowMap["construction_or_measure_provenance"].Value) != "",
		RenormalizationParsed:              strings.TrimSpace(rowMap["renormalization_and_regulator_provenance"].Value) != "",
		Gate536AlignmentParsed:             strings.TrimSpace(rowMap["field_label_alignment_with_gate536"].Value) != "",
		CovarianceProvenanceParsed:         strings.TrimSpace(rowMap["euclidean_covariance_certificate_provenance"].Value) != "",
		OSCertificateProvenanceParsed:      strings.TrimSpace(rowMap["os_reflection_positivity_certificate_provenance"].Value) != "",
		WickIEpsilonParsed:                 strings.TrimSpace(rowMap["wick_and_i_epsilon_provenance"].Value) != "",
		HamiltonianDomainParsed:            strings.TrimSpace(rowMap["hamiltonian_spectrum_domain_certificate"].Value) != "",
		UncertaintyLedgerParsed:            strings.TrimSpace(rowMap["uncertainty_validity_and_reproducibility_ledger"].Value) != "",
		QuarantineParsed:                   rowMap["bridge_only_no_native_write_quarantine"].BridgeOnly && rowMap["bridge_only_no_native_write_quarantine"].NoTheoremInput && !rowMap["bridge_only_no_native_write_quarantine"].NativePromotion,
		SyntheticFixtureRejectedAsPhysical: ledger.SyntheticFixture && !ledger.PhysicalSourceClaim && !ledger.NonSyntheticPhysicalOrConstructiveClaim,
		PhysicalSourceAuthenticated:        false,
	}
	out.AuthenticityPlumbingVerified = out.ChecksumVerified && out.RowsParsed == 13 && out.RequiredRowsMatched == 13 && out.BridgeOnlyRows == 13 && out.NoTheoremInputRows == 13 && out.SourceTaggedRows == 13 && out.ConventionTaggedRows == 13 && out.SyntheticRows == 13 && out.PhysicalClaimRows == 0 && out.ObservedRows == 0 && out.NativePromotionRows == 0 && out.ComparatorRows == 12 && out.ImmutableSourceParsed && out.LicenseParsed && out.ConstructionProvenanceParsed && out.RenormalizationParsed && out.Gate536AlignmentParsed && out.CovarianceProvenanceParsed && out.OSCertificateProvenanceParsed && out.WickIEpsilonParsed && out.HamiltonianDomainParsed && out.UncertaintyLedgerParsed && out.QuarantineParsed && out.SyntheticFixtureRejectedAsPhysical && !out.PhysicalSourceAuthenticated
	if out.AuthenticityPlumbingVerified {
		out.Verdict = strings.Join([]string{StatusSyntheticAuthenticityAdapterExecuted, StatusChecksumIntegrityVerified, StatusProvenanceRowsParsed, StatusSyntheticSourceRejectedAsPhysical, StatusQuarantineTagsPreserved, StatusNoRealSourceImported, StatusNativePromotionRejected}, ";")
		out.Reason = "The synthetic Gate539 ledger proves the source-authenticity parser, checksum, provenance, and quarantine plumbing, then rejects the fixture as physical source data."
	} else {
		out.Verdict = StatusFailedMetadataIncomplete
		out.Reason = "synthetic authenticity adapter plumbing failed"
		out.Failures = []string{StatusFailedMetadataIncomplete}
	}
	return out
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{
		Executed:                      true,
		RealSchwingerSourceImported:   imp.RealSchwingerSourceImported,
		ObservedCorrelationImported:   imp.ObservedCorrelationLoaded,
		ConstructiveMeasureImported:   imp.ConstructiveMeasureLoaded,
		PhysicalOSCertificateImported: imp.PhysicalOSCertificateLoaded,
		PhysicalWickMapImported:       imp.PhysicalWickMapLoaded,
		PhysicalHamiltonianImported:   imp.PhysicalHamiltonianLoaded,
		SyntheticFixtureOnly:          imp.SyntheticFixture && !imp.PhysicalSourceClaim && !imp.NonSyntheticClaim,
		SourceAuthenticatedAsPhysical: out.PhysicalSourceAuthenticated,
		Verdict:                       strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticNotPhysicalSource, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow}, ";"),
		Reason:                        "A synthetic authenticity fixture may validate parser/provenance/checksum plumbing, but it cannot authenticate a real source or write Schwinger functions, OS proof, Wick continuation, Hilbert reconstruction, Hamiltonian, dynamics, causality, or time orientation into native ASHA law.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		BridgeEntries: []string{StatusGate538AirlockInherited, StatusSyntheticAuthenticityAdapterExecuted, StatusChecksumIntegrityVerified, StatusProvenanceRowsParsed, StatusQuarantineTagsPreserved},
		FailedRoutes:  []string{StatusFailedSyntheticNotPhysicalSource, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallNativeWriteBlocked},
		OpenTheorems:  []string{"real Schwinger source import", "source-authenticity comparator execution", "physical OS/Wick/Hilbert/Hamiltonian reconstruction"},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 540, Title: "Real Schwinger Source Import Switch Preflight", Reason: "Gate539 proves the synthetic authenticity parser and rejection path. The next safe boundary is a switch that can detect whether a non-synthetic source is intentionally supplied before any authenticity comparator can execute.", PrimaryTask: "Define the explicit import switch for real/constructive Schwinger sources and keep it off by default."}
}

func truth(a Analysis) string {
	return "Gate539 proves that source-authenticity metadata can be parsed, integrity-checked, and quarantined, but synthetic provenance remains synthetic: no real Schwinger source, constructive measure, physical OS certificate, Wick map, Hilbert space, Hamiltonian, unitarity, global causality, or arrow of time is authenticated or derived."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate538AirlockDefined || !a.Inheritance.Gate538SchemaRowsEnumerated || a.Inheritance.Gate538RequiredRows != 13 || a.Inheritance.Gate538BridgeOnlyRows != 13 || a.Inheritance.Gate538ComparatorRows != 12 || a.Inheritance.Gate538NativeWriteRows != 0 || !a.Inheritance.Gate538DiscriminatorDefined || !a.Inheritance.Gate538SyntheticRejected || !a.Inheritance.Gate538ComparatorBlocked || !a.Inheritance.Gate538NoRealSourceImported || !a.Inheritance.Gate538NativeWriteBlocked || !a.Inheritance.Gate539SyntheticRedirect {
		bad = append(bad, "bad Gate538 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 13 || a.Import.AcceptedRows != 13 || a.Import.RejectedRows != 0 || len(a.Import.MissingRequiredRows) != 0 || len(a.Import.DuplicateRows) != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.PhysicalSourceClaim || a.Import.NonSyntheticClaim || a.Import.RealSchwingerSourceImported || a.Import.ObservedCorrelationLoaded || a.Import.ConstructiveMeasureLoaded || a.Import.PhysicalOSCertificateLoaded || a.Import.PhysicalWickMapLoaded || a.Import.PhysicalHamiltonianLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.Gate538ReferenceComplete || !a.Import.Gate537ReferenceComplete || !a.Import.MetadataComplete || !a.Import.RequiredSchemaRowsMatched || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || !a.Import.AllRowsSourceTagged || !a.Import.AllRowsConventionTagged || a.Import.AnyPhysicalClaim || a.Import.AnyObservedClaim || a.Import.AnyNativePromotionClaim || !a.Import.ChecksumVerified {
		bad = append(bad, "bad synthetic authenticity import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || a.Output.RowsParsed != 13 || a.Output.RequiredRowsMatched != 13 || a.Output.BridgeOnlyRows != 13 || a.Output.NoTheoremInputRows != 13 || a.Output.SourceTaggedRows != 13 || a.Output.ConventionTaggedRows != 13 || a.Output.SyntheticRows != 13 || a.Output.PhysicalClaimRows != 0 || a.Output.ObservedRows != 0 || a.Output.NativePromotionRows != 0 || a.Output.ComparatorRows != 12 || !a.Output.ChecksumVerified || !a.Output.ImmutableSourceParsed || !a.Output.LicenseParsed || !a.Output.ConstructionProvenanceParsed || !a.Output.RenormalizationParsed || !a.Output.Gate536AlignmentParsed || !a.Output.CovarianceProvenanceParsed || !a.Output.OSCertificateProvenanceParsed || !a.Output.WickIEpsilonParsed || !a.Output.HamiltonianDomainParsed || !a.Output.UncertaintyLedgerParsed || !a.Output.QuarantineParsed || !a.Output.AuthenticityPlumbingVerified || !a.Output.SyntheticFixtureRejectedAsPhysical || a.Output.PhysicalSourceAuthenticated || a.Output.PhysicalSchwingerFunctionsDerived || a.Output.PhysicalOSPositivityProven || a.Output.WickRotationGranted || a.Output.PhysicalHilbertSpaceSelected || a.Output.PositiveEnergyHamiltonianDerived || a.Output.UnitaryRealTimeDynamicsDerived || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected {
		bad = append(bad, "bad adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.RealSchwingerSourceImported || a.Firewall.ObservedCorrelationImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.PhysicalOSCertificateImported || a.Firewall.PhysicalWickMapImported || a.Firewall.PhysicalHamiltonianImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.SourceAuthenticatedAsPhysical || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate539 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: airlock=%t rows=%d bridge=%d comparator=%d native_rows=%d discriminator=%t synthetic_rejected=%t comparator_blocked=%t real_absent=%t native_blocked=%t redirect=%t; %s", x.Verdict, x.Gate538AirlockDefined, x.Gate538RequiredRows, x.Gate538BridgeOnlyRows, x.Gate538ComparatorRows, x.Gate538NativeWriteRows, x.Gate538DiscriminatorDefined, x.Gate538SyntheticRejected, x.Gate538ComparatorBlocked, x.Gate538NoRealSourceImported, x.Gate538NativeWriteBlocked, x.Gate539SyntheticRedirect, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%v duplicates=%v bridge=%t synthetic=%t physical_claim=%t nonsynthetic_claim=%t real=%t observed=%t measure=%t OS=%t wick=%t hamiltonian=%t native_request=%t refs=(538:%t,537:%t) metadata=%t checksum=%t expected=%s actual=%s all_bridge=%t all_no_theorem=%t all_synthetic=%t source_tags=%t convention_tags=%t physical_rows=%t observed_rows=%t native_rows=%t; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.MissingRequiredRows, x.DuplicateRows, x.BridgeOnlyLedger, x.SyntheticFixture, x.PhysicalSourceClaim, x.NonSyntheticClaim, x.RealSchwingerSourceImported, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.PhysicalOSCertificateLoaded, x.PhysicalWickMapLoaded, x.PhysicalHamiltonianLoaded, x.NativeRegistryWriteRequested, x.Gate538ReferenceComplete, x.Gate537ReferenceComplete, x.MetadataComplete, x.ChecksumVerified, x.ChecksumExpected, x.ChecksumActual, x.AllRowsBridgeOnly, x.AllRowsNoTheoremInput, x.AllRowsSynthetic, x.AllRowsSourceTagged, x.AllRowsConventionTagged, x.AnyPhysicalClaim, x.AnyObservedClaim, x.AnyNativePromotionClaim, x.Reason)
}

func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("%s: ready=%t rows=%d required=%d bridge=%d no_theorem=%d source=%d convention=%d synthetic=%d physical_claim=%d observed=%d native=%d comparator=%d checksum=%t immutable=%t license=%t construction=%t renormalization=%t gate536=%t covariance=%t OS=%t Wick=%t Hamiltonian=%t uncertainty=%t quarantine=%t plumbing=%t rejected_as_physical=%t authenticated=%t; %s", x.Verdict, x.Ready, x.RowsParsed, x.RequiredRowsMatched, x.BridgeOnlyRows, x.NoTheoremInputRows, x.SourceTaggedRows, x.ConventionTaggedRows, x.SyntheticRows, x.PhysicalClaimRows, x.ObservedRows, x.NativePromotionRows, x.ComparatorRows, x.ChecksumVerified, x.ImmutableSourceParsed, x.LicenseParsed, x.ConstructionProvenanceParsed, x.RenormalizationParsed, x.Gate536AlignmentParsed, x.CovarianceProvenanceParsed, x.OSCertificateProvenanceParsed, x.WickIEpsilonParsed, x.HamiltonianDomainParsed, x.UncertaintyLedgerParsed, x.QuarantineParsed, x.AuthenticityPlumbingVerified, x.SyntheticFixtureRejectedAsPhysical, x.PhysicalSourceAuthenticated, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: real=%t observed=%t measure=%t OS_import=%t wick_import=%t hamiltonian_import=%t synthetic_only=%t authenticated=%t file_native=%t outputs_native=%t native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopened_flavor=%t reopened_EW=%t reopened_gravity=%t reopened_topology=%t reopened_dimension=%t reopened_Krein=%t native_registry=%t; %s", x.Verdict, x.RealSchwingerSourceImported, x.ObservedCorrelationImported, x.ConstructiveMeasureImported, x.PhysicalOSCertificateImported, x.PhysicalWickMapImported, x.PhysicalHamiltonianImported, x.SyntheticFixtureOnly, x.SourceAuthenticatedAsPhysical, x.FileRowsNative, x.AdapterOutputsNative, x.NativeSchwingerFunctionWrite, x.NativeEuclideanMeasureWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate538AirlockInherited, StatusSyntheticAuthenticityLedgerLoaded, StatusAuthenticityRowsAccepted, StatusChecksumIntegrityVerified, StatusProvenanceRowsParsed, StatusMetadataSieveEnforced, StatusSyntheticAuthenticityAdapterExecuted, StatusSyntheticSourceRejectedAsPhysical, StatusQuarantineTagsPreserved, StatusNoRealSourceImported, StatusNativePromotionRejected, StatusFailedSyntheticNotPhysicalSource, StatusFailedSyntheticNotSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func resolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	dir := filepath.Dir(file)
	for i := 0; i < 6; i++ {
		candidate := filepath.Join(dir, path)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		dir = filepath.Dir(dir)
	}
	return path
}
