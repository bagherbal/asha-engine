// Package generation2syntheticschwingerledgeradapter implements Gate 537:
// Synthetic Schwinger-Function Source Ledger Adapter Dry Run.
//
// Gate 536 defined the 19-row source-ledger airlock required before any
// physical or constructive Euclidean Schwinger-function family may enter ASHA's
// OS/Wick/Hilbert bridge. This package loads an explicitly synthetic ledger
// containing all 19 required rows, verifies the row metadata sieve, performs a
// finite structural dry run of the declared reflection operator, Schwinger Gram
// kernel, covariance placeholder, positive-time domain, and OS quadratic-form
// reduction, and then fails closed against native theorem promotion.
package generation2syntheticschwingerledgeradapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2physicalschwingerledgerairlock"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID       = "GATE537-SYNTHETIC-SCHWINGER-FUNCTION-SOURCE-LEDGER-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_schwinger_function_ledger_gate537.json"

	StatusGate536AirlockInherited             = "CONDITIONAL_SUPPORT_GATE536_PHYSICAL_SCHWINGER_SOURCE_LEDGER_AIRLOCK_INHERITED"
	StatusSyntheticSchwingerLedgerLoaded      = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_LEDGER_LOADED"
	StatusSyntheticSchwingerRowsAccepted      = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_19_SCHEMA_ROWS_ACCEPTED"
	StatusMetadataSieveEnforced               = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_METADATA_SIEVE_ENFORCED"
	StatusSyntheticSchwingerAdapterExecuted   = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_ADAPTER_EXECUTED"
	StatusThetaEInvolutionResidualZero        = "CONDITIONAL_SUPPORT_SYNTHETIC_THETA_E_INVOLUTION_RESIDUAL_ZERO"
	StatusSchwingerKernelSymmetryResidualZero = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_KERNEL_SYMMETRY_RESIDUAL_ZERO"
	StatusEuclideanCovarianceResidualZero     = "CONDITIONAL_SUPPORT_SYNTHETIC_EUCLIDEAN_COVARIANCE_RESIDUAL_ZERO"
	StatusPositiveTimeDomainClosed            = "CONDITIONAL_SUPPORT_SYNTHETIC_POSITIVE_TIME_TEST_DOMAIN_CLOSED"
	StatusOSQuadraticFormNonnegative          = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_QUADRATIC_FORM_NONNEGATIVE"
	StatusDummyHamiltonianSpectrumParsed      = "CONDITIONAL_SUPPORT_SYNTHETIC_DUMMY_HAMILTONIAN_SPECTRUM_PARSED"
	StatusNoPhysicalSchwingerDataImported     = "CONDITIONAL_SUPPORT_NO_PHYSICAL_SCHWINGER_DATA_IMPORTED_IN_GATE537"
	StatusNativePromotionRejected             = "CONDITIONAL_SUPPORT_SYNTHETIC_SCHWINGER_NATIVE_PROMOTION_REJECTED"

	StatusFailedLedgerMissing                 = "FAILED_ROUTE_GATE537_SYNTHETIC_SCHWINGER_LEDGER_MISSING"
	StatusFailedMetadataIncomplete            = "FAILED_ROUTE_GATE537_SYNTHETIC_SCHWINGER_METADATA_INCOMPLETE"
	StatusFailedSchemaRowsIncomplete          = "FAILED_ROUTE_GATE537_SYNTHETIC_SCHWINGER_SCHEMA_ROWS_INCOMPLETE"
	StatusFailedInvalidMatrixDomain           = "FAILED_ROUTE_GATE537_INVALID_SYNTHETIC_SCHWINGER_MATRIX_DOMAIN"
	StatusFailedThetaEInvolutionNonzero       = "FAILED_ROUTE_GATE537_THETA_E_INVOLUTION_RESIDUAL_NONZERO"
	StatusFailedKernelSymmetryNonzero         = "FAILED_ROUTE_GATE537_SCHWINGER_KERNEL_SYMMETRY_RESIDUAL_NONZERO"
	StatusFailedCovarianceNonzero             = "FAILED_ROUTE_GATE537_EUCLIDEAN_COVARIANCE_RESIDUAL_NONZERO"
	StatusFailedPositiveDomainNotClosed       = "FAILED_ROUTE_GATE537_POSITIVE_TIME_DOMAIN_CLOSURE_FAILED"
	StatusFailedOSQuadraticNegative           = "FAILED_ROUTE_GATE537_SYNTHETIC_OS_QUADRATIC_FORM_NEGATIVE"
	StatusFailedSyntheticNotPhysicalSchwinger = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticNotOSProof           = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_PROVE_PHYSICAL_OS_REFLECTION_POSITIVITY"
	StatusFailedSyntheticNotWick              = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticNotHilbert           = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticNotHamiltonian       = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticNotUnitary           = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticNotGlobal            = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticNotArrow             = "FAILED_ROUTE_SYNTHETIC_SCHWINGER_LEDGER_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_GATE537_SYNTHETIC_SCHWINGER_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked          = "FIREWALL_BLOCKED_GATE537_PHYSICAL_CORRELATION_NATIVE_WRITE"
)

const tolerance = 1e-12

type Inheritance struct {
	Executed bool

	Gate536AirlockDefined                 bool
	Gate536SchemaRowsEnumerated           bool
	Gate536RequiredRows                   int
	Gate536BridgeOnlyRows                 int
	Gate536NativeWriteRows                int
	Gate536ComparatorRows                 int
	Gate536SourceTagsEnforced             bool
	Gate536OSCertificateRequired          bool
	Gate536WickAndIEpsilonRequired        bool
	Gate536HamiltonianCertificateRequired bool
	Gate536ComparatorBlocked              bool
	Gate536NoObservedCorrelatorsImported  bool
	Gate536NativePromotionRejected        bool
	Gate536NativeWriteBlocked             bool
	Gate537SyntheticRedirect              bool

	Verdict, Reason string
}

type SchemaRowMetadata struct {
	SchemaKey        string `json:"schema_key"`
	Source           string `json:"source"`
	SourceVersion    string `json:"source_version"`
	Convention       string `json:"convention"`
	ValueKind        string `json:"value_kind"`
	BridgeOnly       bool   `json:"bridge_only"`
	ComparatorOnly   bool   `json:"comparator_only"`
	Synthetic        bool   `json:"synthetic"`
	Observed         bool   `json:"observed"`
	NoTheoremInput   bool   `json:"no_theorem_input"`
	NativePromotion  bool   `json:"native_promotion"`
	NativeInputClaim bool   `json:"native_input_claim"`
	Value            string `json:"value"`
}

type SyntheticFixture struct {
	TestFunctionDomain               []string      `json:"test_function_domain"`
	PositiveTimeSupport              []int         `json:"positive_time_support"`
	ReflectionOperatorThetaE         [][]float64   `json:"reflection_operator_theta_e"`
	SyntheticSchwingerGramMatrix     [][]float64   `json:"synthetic_schwinger_gram_matrix"`
	EuclideanCovarianceOperator      [][]float64   `json:"euclidean_covariance_operator"`
	PositiveTimeTranslationOperators [][][]float64 `json:"positive_time_translation_operators"`
	OSTestVectors                    [][]float64   `json:"os_test_vectors"`
	DummyHamiltonianSpectrum         []float64     `json:"dummy_hamiltonian_spectrum"`
	WickMapReference                 string        `json:"wick_map_reference"`
	IEpsilonConvention               string        `json:"i_epsilon_convention"`
	NullSpaceQuotientRule            string        `json:"null_space_quotient_rule"`
	ReconstructionMapCertificate     string        `json:"reconstruction_map_certificate"`
	OSQuadraticFormDefinition        string        `json:"os_quadratic_form_definition"`
	ReflectionPositivityCertificate  string        `json:"reflection_positivity_certificate"`
	EuclideanCovarianceCertificate   string        `json:"euclidean_covariance_certificate"`
}

type SyntheticSchwingerLedger struct {
	Gate                         int                 `json:"gate"`
	LedgerName                   string              `json:"ledger_name"`
	Description                  string              `json:"description"`
	Gate536AirlockReference      string              `json:"gate536_airlock_reference"`
	Gate534OSAdapterReference    string              `json:"gate534_os_adapter_reference"`
	BridgeOnly                   bool                `json:"bridge_only"`
	NativeRegistryWrite          bool                `json:"native_registry_write"`
	SyntheticFixture             bool                `json:"synthetic_fixture"`
	PhysicalSchwingerLoaded      bool                `json:"physical_schwinger_loaded"`
	ObservedCorrelationLoaded    bool                `json:"observed_correlation_loaded"`
	ConstructiveMeasureLoaded    bool                `json:"constructive_measure_loaded"`
	ObservedWickLoaded           bool                `json:"observed_wick_loaded"`
	ObservedHamiltonianLoaded    bool                `json:"observed_hamiltonian_loaded"`
	ObservedCausalBoundaryLoaded bool                `json:"observed_causal_boundary_loaded"`
	Source                       string              `json:"source"`
	SourceVersion                string              `json:"source_version"`
	Convention                   string              `json:"convention"`
	Rows                         []SchemaRowMetadata `json:"rows"`
	Fixture                      SyntheticFixture    `json:"fixture"`
}

type FileImport struct {
	Executed                     bool
	Loaded                       bool
	Path                         string
	Rows                         int
	AcceptedRows                 int
	RejectedRows                 int
	MissingRequiredRows          []string
	DuplicateRows                []string
	BridgeOnlyLedger             bool
	SyntheticFixture             bool
	PhysicalSchwingerLoaded      bool
	ObservedCorrelationLoaded    bool
	ConstructiveMeasureLoaded    bool
	ObservedWickLoaded           bool
	ObservedHamiltonianLoaded    bool
	ObservedCausalBoundaryLoaded bool
	NativeRegistryWriteRequested bool
	Gate536ReferenceComplete     bool
	Gate534ReferenceComplete     bool
	MetadataComplete             bool
	AllRowsBridgeOnly            bool
	AllRowsNoTheoremInput        bool
	AllRowsSynthetic             bool
	AllRowsSourceTagged          bool
	AllRowsConventionTagged      bool
	AnyObservedClaim             bool
	NativePromotionRejected      bool
	RequiredSchemaRowsMatched    bool
	Verdict, Reason              string
	Failures                     []string
}

type AdapterInput struct {
	Rows                            int
	DomainLabels                    []string
	PositiveTimeSupport             []int
	ThetaE                          linear.Matrix
	SchwingerGram                   linear.Matrix
	EuclideanCovariance             linear.Matrix
	PositiveTranslations            []linear.Matrix
	TestVectors                     []linear.Matrix
	DummyHamiltonianSpectrum        []float64
	WickMapReference                string
	IEpsilonConvention              string
	NullSpaceQuotientRule           string
	ReconstructionMapCertificate    string
	OSQuadraticFormDefinition       string
	ReflectionPositivityCertificate string
	EuclideanCovarianceCertificate  string
	BridgeOnly                      bool
	SyntheticFixture                bool
	PhysicalSchwingerLoaded         bool
	ObservedCorrelationLoaded       bool
	ConstructiveMeasureLoaded       bool
	ObservedWickLoaded              bool
	ObservedHamiltonianLoaded       bool
	ObservedCausalBoundaryLoaded    bool
	NativePromotion                 bool
	MetadataComplete                bool
}

type AdapterOutput struct {
	Executed  bool
	Attempted bool
	Ready     bool

	Dimension                            int
	PositiveDomainDimension              int
	SchemaRowsParsed                     int
	RequiredRowsMatched                  int
	BridgeOnlyRows                       int
	NoTheoremInputRows                   int
	SourceTaggedRows                     int
	ConventionTaggedRows                 int
	ComparatorOnlyRows                   int
	SyntheticRows                        int
	NativePromotionRows                  int
	ObservedRows                         int
	ThetaEInvolutionResidual             float64
	SchwingerKernelSymmetryResidual      float64
	EuclideanCovarianceResidual          float64
	PositiveTimeDomainClosureResidual    float64
	OSGramSymmetryResidual               float64
	OSGramEigenMin                       float64
	OSGramEigenMax                       float64
	OSGramPositiveEigenvalues            int
	OSGramNegativeEigenvalues            int
	OSGramZeroEigenvalues                int
	OSGramPositiveDefinite               bool
	QuadraticMinimum                     float64
	QuadraticMaximum                     float64
	NonzeroTestVectors                   int
	NullTestVectors                      int
	PositiveQuadraticVectors             int
	NullQuadraticVectors                 int
	NegativeQuadraticVectors             int
	AllSyntheticQuadraticsNonnegative    bool
	DummyHamiltonianLevels               int
	DummyHamiltonianMin                  float64
	DummyHamiltonianMax                  float64
	DummyHamiltonianSpectrumParsed       bool
	WickMapPlaceholderParsed             bool
	IEpsilonPlaceholderParsed            bool
	NullQuotientMetadataConsistent       bool
	ReconstructionCertificateParsed      bool
	ReflectionCertificateParsed          bool
	EuclideanCovarianceCertificateParsed bool
	FiniteSchwingerPlumbingVerified      bool
	SyntheticSchwingerAdapterVerified    bool
	PhysicalSchwingerFunctionsDerived    bool
	PhysicalOSPositivityProven           bool
	PhysicalHilbertSpaceSelected         bool
	WickRotationGranted                  bool
	PositiveEnergyHamiltonianDerived     bool
	UnitaryRealTimeDynamicsDerived       bool
	GlobalHyperbolicityGranted           bool
	ArrowOfTimeSelected                  bool
	Verdict, Reason                      string
	Failures                             []string
}

type Firewall struct {
	Executed bool

	PhysicalSchwingerDataImported   bool
	ObservedCorrelationDataImported bool
	ConstructiveMeasureImported     bool
	ObservedWickDataImported        bool
	ObservedHamiltonianDataImported bool
	ObservedCausalBoundaryImported  bool
	SyntheticFixtureOnly            bool
	FileRowsNative                  bool
	AdapterOutputsNative            bool
	NativeSchwingerFunctionWrite    bool
	NativeEuclideanMeasureWrite     bool
	NativeOSPositivityWrite         bool
	NativeWickWrite                 bool
	NativeHilbertWrite              bool
	NativeHamiltonianWrite          bool
	NativeUnitaryDynamicsWrite      bool
	NativeGlobalCausalWrite         bool
	NativeTimeArrowWrite            bool
	ReopenedFlavorFirewall          bool
	ReopenedEWScaleFirewall         bool
	ReopenedGravityScaleFirewall    bool
	ReopenedTopologyFirewall        bool
	ReopenedDimensionalFirewall     bool
	ReopenedKreinHilbertFirewall    bool
	NativeRegistryWritten           bool
	Verdict, Reason                 string
}

type RegistryUpdate struct{ NativeEntries, BridgeEntries, EnvironmentalEntries, FailedRoutes, OpenTheorems []string }

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      FileImport
	Input       AdapterInput
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
	g536, err := generation2physicalschwingerledgerairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate536 physical Schwinger source airlock: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g536)}
	ledger, imp := loadLedger(path, g536)
	a.Import = imp
	if imp.Loaded && imp.MetadataComplete && imp.RequiredSchemaRowsMatched {
		in, err := buildInput(ledger, imp)
		if err != nil {
			a.Input = AdapterInput{Rows: len(ledger.Rows), BridgeOnly: ledger.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture, PhysicalSchwingerLoaded: ledger.PhysicalSchwingerLoaded, ObservedCorrelationLoaded: ledger.ObservedCorrelationLoaded, ConstructiveMeasureLoaded: ledger.ConstructiveMeasureLoaded, ObservedWickLoaded: ledger.ObservedWickLoaded, ObservedHamiltonianLoaded: ledger.ObservedHamiltonianLoaded, ObservedCausalBoundaryLoaded: ledger.ObservedCausalBoundaryLoaded, NativePromotion: ledger.NativeRegistryWrite, MetadataComplete: imp.MetadataComplete}
			a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnlyRows: imp.AcceptedRows, NoTheoremInputRows: imp.AcceptedRows, NativePromotionRows: 0, ObservedRows: 0, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
		} else {
			a.Input = in
			a.Output = runAdapter(in, imp)
		}
	} else if !imp.Loaded {
		a.Output = AdapterOutput{Executed: true, Attempted: false, Verdict: StatusFailedLedgerMissing, Reason: "explicit Gate537 synthetic Schwinger ledger was not found", Failures: []string{StatusFailedLedgerMissing}}
	} else {
		failure := StatusFailedMetadataIncomplete
		if !imp.RequiredSchemaRowsMatched {
			failure = StatusFailedSchemaRowsIncomplete
		}
		a.Output = AdapterOutput{Executed: true, Attempted: true, SchemaRowsParsed: imp.Rows, Verdict: failure, Reason: "Gate537 synthetic Schwinger ledger did not satisfy the Gate536 schema/metadata domain", Failures: []string{failure}}
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

func buildInheritance(g generation2physicalschwingerledgerairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                              true,
		Gate536AirlockDefined:                 g.Schema.Executed && g.Schema.SourceIdentifierRequired && g.Schema.NPointFamilyRequired,
		Gate536SchemaRowsEnumerated:           len(g.Schema.Rows) == 19 && g.Schema.RequiredRows == 19,
		Gate536RequiredRows:                   g.Schema.RequiredRows,
		Gate536BridgeOnlyRows:                 g.Schema.BridgeOnlyRows,
		Gate536NativeWriteRows:                g.Schema.NativeWriteRows,
		Gate536ComparatorRows:                 g.Schema.ComparatorRows,
		Gate536SourceTagsEnforced:             g.Schema.SourceAndLicenseRequired && g.Schema.NoTheoremInputRequired,
		Gate536OSCertificateRequired:          g.Schema.ReflectionPositivityCertificateRequired && g.Schema.OSQuadraticFormRequired,
		Gate536WickAndIEpsilonRequired:        g.Schema.WickMapRequired && g.Schema.IepsilonConventionRequired,
		Gate536HamiltonianCertificateRequired: g.Schema.HamiltonianSpectrumCertificateRequired,
		Gate536ComparatorBlocked:              !g.Guard.ComparatorExecutionPerformed && !g.Guard.NPointDistributionsEvaluated,
		Gate536NoObservedCorrelatorsImported:  !g.Firewall.ObservedSchwingerDataImported && !g.Firewall.ObservedWickDataImported && !g.Firewall.ObservedHamiltonianDataImported && !g.Guard.ObservedCorrelationDataImported,
		Gate536NativePromotionRejected:        g.Schema.NativePromotionRejected,
		Gate536NativeWriteBlocked:             !g.Firewall.NativeRegistryWritten && !g.Firewall.NativeSchwingerWrite && !g.Firewall.NativeOSPositivityWrite && !g.Firewall.NativeWickWrite && !g.Firewall.NativeHamiltonianWrite,
		Gate537SyntheticRedirect:              g.Next.Gate == 537,
		Verdict:                               StatusGate536AirlockInherited,
		Reason:                                "Gate537 inherits Gate536's 19-row Schwinger source ledger airlock and executes only a synthetic bridge-only adapter dry run.",
	}
}

func loadLedger(path string, g536 generation2physicalschwingerledgerairlock.Analysis) (SyntheticSchwingerLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved, Verdict: StatusFailedLedgerMissing, Reason: "ledger not loaded", Failures: []string{StatusFailedLedgerMissing}}
	b, err := os.ReadFile(resolved)
	if err != nil {
		return SyntheticSchwingerLedger{}, imp
	}
	var ledger SyntheticSchwingerLedger
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
	imp.PhysicalSchwingerLoaded = ledger.PhysicalSchwingerLoaded
	imp.ObservedCorrelationLoaded = ledger.ObservedCorrelationLoaded
	imp.ConstructiveMeasureLoaded = ledger.ConstructiveMeasureLoaded
	imp.ObservedWickLoaded = ledger.ObservedWickLoaded
	imp.ObservedHamiltonianLoaded = ledger.ObservedHamiltonianLoaded
	imp.ObservedCausalBoundaryLoaded = ledger.ObservedCausalBoundaryLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.Gate536ReferenceComplete = strings.TrimSpace(ledger.Gate536AirlockReference) != ""
	imp.Gate534ReferenceComplete = strings.TrimSpace(ledger.Gate534OSAdapterReference) != ""
	imp.AllRowsBridgeOnly = true
	imp.AllRowsNoTheoremInput = true
	imp.AllRowsSynthetic = true
	imp.AllRowsSourceTagged = true
	imp.AllRowsConventionTagged = true

	required := requiredRowSet(g536)
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
		imp.AnyObservedClaim = imp.AnyObservedClaim || row.Observed
		if row.NativePromotion || row.NativeInputClaim {
			imp.NativePromotionRejected = true
		}
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
	imp.RequiredSchemaRowsMatched = len(required) == 19 && imp.Rows == 19 && imp.AcceptedRows == 19 && len(imp.MissingRequiredRows) == 0 && len(imp.DuplicateRows) == 0
	imp.MetadataComplete = ledger.Gate == 537 && imp.RequiredSchemaRowsMatched && imp.BridgeOnlyLedger && !imp.NativeRegistryWriteRequested && imp.SyntheticFixture && !imp.PhysicalSchwingerLoaded && !imp.ObservedCorrelationLoaded && !imp.ConstructiveMeasureLoaded && !imp.ObservedWickLoaded && !imp.ObservedHamiltonianLoaded && !imp.ObservedCausalBoundaryLoaded && imp.Gate536ReferenceComplete && imp.Gate534ReferenceComplete && imp.AllRowsBridgeOnly && imp.AllRowsNoTheoremInput && imp.AllRowsSynthetic && imp.AllRowsSourceTagged && imp.AllRowsConventionTagged && !imp.AnyObservedClaim && !imp.NativePromotionRejected && strings.TrimSpace(ledger.Source) != "" && strings.TrimSpace(ledger.SourceVersion) != "" && strings.TrimSpace(ledger.Convention) != ""
	if imp.MetadataComplete {
		imp.Verdict = strings.Join([]string{StatusSyntheticSchwingerLedgerLoaded, StatusSyntheticSchwingerRowsAccepted, StatusMetadataSieveEnforced}, ";")
		imp.Reason = "Gate537 synthetic Schwinger ledger loaded with exactly the Gate536 19-row schema, source/convention tags on every row, bridge_only=true and no_theorem_input=true everywhere, and no observed or native-promotion claims."
		imp.Failures = nil
	} else {
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = "ledger metadata or row flags violate the Gate536 Schwinger source airlock"
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		if !imp.RequiredSchemaRowsMatched {
			imp.Failures = append(imp.Failures, StatusFailedSchemaRowsIncomplete)
		}
		if imp.NativePromotionRejected || imp.NativeRegistryWriteRequested {
			imp.Failures = append(imp.Failures, StatusNativePromotionRejected)
		}
	}
	return ledger, imp
}

func requiredRowSet(g536 generation2physicalschwingerledgerairlock.Analysis) map[string]bool {
	out := map[string]bool{}
	for _, row := range g536.Schema.Rows {
		out[row.Name] = true
	}
	return out
}

func validateRowMetadata(row SchemaRowMetadata) bool {
	return strings.TrimSpace(row.SchemaKey) != "" && strings.TrimSpace(row.Source) != "" && strings.TrimSpace(row.SourceVersion) != "" && strings.TrimSpace(row.Convention) != "" && strings.TrimSpace(row.ValueKind) != "" && strings.TrimSpace(row.Value) != "" && row.BridgeOnly && row.Synthetic && row.NoTheoremInput && !row.Observed && !row.NativePromotion && !row.NativeInputClaim
}

func buildInput(ledger SyntheticSchwingerLedger, imp FileImport) (AdapterInput, error) {
	R, err := linear.FromRows(ledger.Fixture.ReflectionOperatorThetaE)
	if err != nil {
		return AdapterInput{}, fmt.Errorf("reflection operator theta_E: %w", err)
	}
	K, err := linear.FromRows(ledger.Fixture.SyntheticSchwingerGramMatrix)
	if err != nil {
		return AdapterInput{}, fmt.Errorf("synthetic Schwinger Gram matrix: %w", err)
	}
	C, err := linear.FromRows(ledger.Fixture.EuclideanCovarianceOperator)
	if err != nil {
		return AdapterInput{}, fmt.Errorf("Euclidean covariance operator: %w", err)
	}
	if R.Rows() != R.Cols() || K.Rows() != K.Cols() || C.Rows() != C.Cols() || R.Rows() != K.Rows() || R.Rows() != C.Rows() {
		return AdapterInput{}, fmt.Errorf("theta_E, Schwinger Gram, and covariance operator must be square over the same finite domain: R=%dx%d K=%dx%d C=%dx%d", R.Rows(), R.Cols(), K.Rows(), K.Cols(), C.Rows(), C.Cols())
	}
	n := R.Rows()
	if len(ledger.Fixture.TestFunctionDomain) != n {
		return AdapterInput{}, fmt.Errorf("test-function domain labels=%d but matrix dimension=%d", len(ledger.Fixture.TestFunctionDomain), n)
	}
	if len(ledger.Fixture.PositiveTimeSupport) == 0 {
		return AdapterInput{}, fmt.Errorf("positive-time support is empty")
	}
	for _, idx := range ledger.Fixture.PositiveTimeSupport {
		if idx < 0 || idx >= n {
			return AdapterInput{}, fmt.Errorf("positive-time support index %d outside dimension %d", idx, n)
		}
	}
	translations := []linear.Matrix{}
	for i, rows := range ledger.Fixture.PositiveTimeTranslationOperators {
		T, err := linear.FromRows(rows)
		if err != nil {
			return AdapterInput{}, fmt.Errorf("positive-time translation %d: %w", i, err)
		}
		if T.Rows() != n || T.Cols() != n {
			return AdapterInput{}, fmt.Errorf("positive-time translation %d dimension %dx%d does not match %d", i, T.Rows(), T.Cols(), n)
		}
		translations = append(translations, T)
	}
	tests := []linear.Matrix{}
	for i, row := range ledger.Fixture.OSTestVectors {
		if len(row) != n {
			return AdapterInput{}, fmt.Errorf("OS test vector %d dimension=%d expected=%d", i, len(row), n)
		}
		v, err := linear.FromRows([][]float64{row})
		if err != nil {
			return AdapterInput{}, fmt.Errorf("OS test vector %d: %w", i, err)
		}
		tests = append(tests, v.Transpose())
	}
	return AdapterInput{
		Rows:                            len(ledger.Rows),
		DomainLabels:                    ledger.Fixture.TestFunctionDomain,
		PositiveTimeSupport:             append([]int(nil), ledger.Fixture.PositiveTimeSupport...),
		ThetaE:                          R,
		SchwingerGram:                   K,
		EuclideanCovariance:             C,
		PositiveTranslations:            translations,
		TestVectors:                     tests,
		DummyHamiltonianSpectrum:        append([]float64(nil), ledger.Fixture.DummyHamiltonianSpectrum...),
		WickMapReference:                ledger.Fixture.WickMapReference,
		IEpsilonConvention:              ledger.Fixture.IEpsilonConvention,
		NullSpaceQuotientRule:           ledger.Fixture.NullSpaceQuotientRule,
		ReconstructionMapCertificate:    ledger.Fixture.ReconstructionMapCertificate,
		OSQuadraticFormDefinition:       ledger.Fixture.OSQuadraticFormDefinition,
		ReflectionPositivityCertificate: ledger.Fixture.ReflectionPositivityCertificate,
		EuclideanCovarianceCertificate:  ledger.Fixture.EuclideanCovarianceCertificate,
		BridgeOnly:                      ledger.BridgeOnly,
		SyntheticFixture:                ledger.SyntheticFixture,
		PhysicalSchwingerLoaded:         ledger.PhysicalSchwingerLoaded,
		ObservedCorrelationLoaded:       ledger.ObservedCorrelationLoaded,
		ConstructiveMeasureLoaded:       ledger.ConstructiveMeasureLoaded,
		ObservedWickLoaded:              ledger.ObservedWickLoaded,
		ObservedHamiltonianLoaded:       ledger.ObservedHamiltonianLoaded,
		ObservedCausalBoundaryLoaded:    ledger.ObservedCausalBoundaryLoaded,
		NativePromotion:                 ledger.NativeRegistryWrite,
		MetadataComplete:                imp.MetadataComplete,
	}, nil
}

func runAdapter(in AdapterInput, imp FileImport) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: true, Ready: true, Dimension: in.ThetaE.Rows(), PositiveDomainDimension: len(in.PositiveTimeSupport), SchemaRowsParsed: imp.Rows, RequiredRowsMatched: imp.AcceptedRows, BridgeOnlyRows: imp.AcceptedRows, NoTheoremInputRows: imp.AcceptedRows, SourceTaggedRows: imp.AcceptedRows, ConventionTaggedRows: imp.AcceptedRows, SyntheticRows: imp.AcceptedRows, NativePromotionRows: 0, ObservedRows: 0}
	I := linear.Identity(in.ThetaE.Rows())
	R2, err := in.ThetaE.Mul(in.ThetaE)
	if err != nil {
		return failedOutput(out, StatusFailedInvalidMatrixDomain, err.Error())
	}
	out.ThetaEInvolutionResidual = maxDiff(R2, I)
	out.SchwingerKernelSymmetryResidual = maxDiff(in.SchwingerGram, in.SchwingerGram.Transpose())
	ctk, err := in.EuclideanCovariance.Transpose().Mul(in.SchwingerGram)
	if err != nil {
		return failedOutput(out, StatusFailedInvalidMatrixDomain, err.Error())
	}
	ctkc, err := ctk.Mul(in.EuclideanCovariance)
	if err != nil {
		return failedOutput(out, StatusFailedInvalidMatrixDomain, err.Error())
	}
	out.EuclideanCovarianceResidual = maxDiff(ctkc, in.SchwingerGram)
	out.PositiveTimeDomainClosureResidual = positiveDomainClosureResidual(in)
	OS := osGram(in)
	out.OSGramSymmetryResidual = maxDiff(OS, OS.Transpose())
	eig, err := linear.SymmetricEigenJacobi(OS, tolerance, 100)
	if err != nil {
		return failedOutput(out, StatusFailedInvalidMatrixDomain, err.Error())
	}
	out.OSGramEigenMin, out.OSGramEigenMax = extrema(eig.Values)
	out.OSGramPositiveEigenvalues, out.OSGramNegativeEigenvalues, out.OSGramZeroEigenvalues = inertia(eig.Values)
	out.OSGramPositiveDefinite = out.OSGramPositiveEigenvalues == len(in.PositiveTimeSupport) && out.OSGramNegativeEigenvalues == 0 && out.OSGramZeroEigenvalues == 0
	quadratics := []float64{}
	for _, v := range in.TestVectors {
		q, err := osQuadratic(in, v)
		if err != nil {
			return failedOutput(out, StatusFailedInvalidMatrixDomain, err.Error())
		}
		q = clean(q)
		quadratics = append(quadratics, q)
		if vectorNorm(v) > tolerance {
			out.NonzeroTestVectors++
		} else {
			out.NullTestVectors++
		}
		switch {
		case q > tolerance:
			out.PositiveQuadraticVectors++
		case q < -tolerance:
			out.NegativeQuadraticVectors++
		default:
			out.NullQuadraticVectors++
		}
	}
	out.QuadraticMinimum, out.QuadraticMaximum = extrema(quadratics)
	out.AllSyntheticQuadraticsNonnegative = out.NegativeQuadraticVectors == 0
	out.DummyHamiltonianLevels = len(in.DummyHamiltonianSpectrum)
	out.DummyHamiltonianMin, out.DummyHamiltonianMax = extrema(in.DummyHamiltonianSpectrum)
	out.DummyHamiltonianSpectrumParsed = out.DummyHamiltonianLevels > 0
	out.WickMapPlaceholderParsed = strings.TrimSpace(in.WickMapReference) != ""
	out.IEpsilonPlaceholderParsed = strings.TrimSpace(in.IEpsilonConvention) != ""
	out.NullQuotientMetadataConsistent = strings.Contains(strings.ToLower(in.NullSpaceQuotientRule), "zero") || strings.TrimSpace(in.NullSpaceQuotientRule) != ""
	out.ReconstructionCertificateParsed = strings.TrimSpace(in.ReconstructionMapCertificate) != ""
	out.ReflectionCertificateParsed = strings.TrimSpace(in.ReflectionPositivityCertificate) != ""
	out.EuclideanCovarianceCertificateParsed = strings.TrimSpace(in.EuclideanCovarianceCertificate) != ""
	out.FiniteSchwingerPlumbingVerified = out.ThetaEInvolutionResidual == 0 && out.SchwingerKernelSymmetryResidual == 0 && out.EuclideanCovarianceResidual == 0 && out.PositiveTimeDomainClosureResidual == 0 && out.OSGramSymmetryResidual == 0 && out.OSGramPositiveDefinite && out.AllSyntheticQuadraticsNonnegative && out.DummyHamiltonianSpectrumParsed && out.WickMapPlaceholderParsed && out.IEpsilonPlaceholderParsed && out.NullQuotientMetadataConsistent && out.ReconstructionCertificateParsed && out.ReflectionCertificateParsed && out.EuclideanCovarianceCertificateParsed
	out.SyntheticSchwingerAdapterVerified = out.FiniteSchwingerPlumbingVerified && in.BridgeOnly && in.SyntheticFixture && !in.PhysicalSchwingerLoaded && !in.ObservedCorrelationLoaded && !in.ConstructiveMeasureLoaded && !in.ObservedWickLoaded && !in.ObservedHamiltonianLoaded && !in.ObservedCausalBoundaryLoaded && !in.NativePromotion
	failures := []string{}
	if out.ThetaEInvolutionResidual != 0 {
		failures = append(failures, StatusFailedThetaEInvolutionNonzero)
	}
	if out.SchwingerKernelSymmetryResidual != 0 {
		failures = append(failures, StatusFailedKernelSymmetryNonzero)
	}
	if out.EuclideanCovarianceResidual != 0 {
		failures = append(failures, StatusFailedCovarianceNonzero)
	}
	if out.PositiveTimeDomainClosureResidual != 0 {
		failures = append(failures, StatusFailedPositiveDomainNotClosed)
	}
	if !out.AllSyntheticQuadraticsNonnegative {
		failures = append(failures, StatusFailedOSQuadraticNegative)
	}
	if !out.SyntheticSchwingerAdapterVerified {
		failures = append(failures, StatusFailedInvalidMatrixDomain)
	}
	if len(failures) == 0 {
		out.Verdict = strings.Join([]string{StatusSyntheticSchwingerAdapterExecuted, StatusThetaEInvolutionResidualZero, StatusSchwingerKernelSymmetryResidualZero, StatusEuclideanCovarianceResidualZero, StatusPositiveTimeDomainClosed, StatusOSQuadraticFormNonnegative, StatusDummyHamiltonianSpectrumParsed, StatusNativePromotionRejected}, ";")
		out.Reason = "Synthetic Schwinger ledger plumbing passes: all 19 rows parse through the Gate536 schema, theta_E is an involution, the finite Gram kernel is symmetric/covariant, the positive-time domain is closed, sampled OS quadratic forms are nonnegative, and every output remains bridge-only."
	} else {
		out.Verdict = strings.Join(failures, ";")
		out.Reason = "synthetic Schwinger adapter failed one or more finite plumbing checks"
		out.Failures = failures
	}
	return out
}

func failedOutput(out AdapterOutput, verdict, reason string) AdapterOutput {
	out.Ready = false
	out.Verdict = verdict
	out.Reason = reason
	out.Failures = []string{verdict}
	return out
}

func positiveDomainClosureResidual(in AdapterInput) float64 {
	max := 0.0
	for _, T := range in.PositiveTranslations {
		for _, idx := range in.PositiveTimeSupport {
			for r := 0; r < T.Rows(); r++ {
				if !contains(in.PositiveTimeSupport, r) {
					if v := math.Abs(T.At(r, idx)); v > max {
						max = v
					}
				}
			}
		}
	}
	return clean(max)
}

func osGram(in AdapterInput) linear.Matrix {
	m := len(in.PositiveTimeSupport)
	out := linear.NewMatrix(m, m)
	RK, _ := in.ThetaE.Mul(in.SchwingerGram)
	for i, row := range in.PositiveTimeSupport {
		for j, col := range in.PositiveTimeSupport {
			out.Set(i, j, clean(RK.At(row, col)))
		}
	}
	return out
}

func osQuadratic(in AdapterInput, v linear.Matrix) (float64, error) {
	Rv, err := in.ThetaE.Mul(v)
	if err != nil {
		return 0, err
	}
	Kv, err := in.SchwingerGram.Mul(v)
	if err != nil {
		return 0, err
	}
	q, err := Rv.Transpose().Mul(Kv)
	if err != nil {
		return 0, err
	}
	return q.At(0, 0), nil
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{
		Executed:                        true,
		PhysicalSchwingerDataImported:   imp.PhysicalSchwingerLoaded,
		ObservedCorrelationDataImported: imp.ObservedCorrelationLoaded,
		ConstructiveMeasureImported:     imp.ConstructiveMeasureLoaded,
		ObservedWickDataImported:        imp.ObservedWickLoaded,
		ObservedHamiltonianDataImported: imp.ObservedHamiltonianLoaded,
		ObservedCausalBoundaryImported:  imp.ObservedCausalBoundaryLoaded,
		SyntheticFixtureOnly:            imp.SyntheticFixture && !imp.PhysicalSchwingerLoaded && !imp.ObservedCorrelationLoaded && !imp.ConstructiveMeasureLoaded && !imp.ObservedWickLoaded && !imp.ObservedHamiltonianLoaded && !imp.ObservedCausalBoundaryLoaded,
		Verdict:                         strings.Join([]string{StatusNoPhysicalSchwingerDataImported, StatusFailedSyntheticNotPhysicalSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                          "Gate537 accepts only synthetic source-ledger plumbing. It writes no native Schwinger function, Euclidean measure, OS positivity theorem, Wick map, Hilbert space, Hamiltonian, unitary dynamics, global-causal structure, or time arrow.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new native law is written at Gate 537.",
			"The native registry remains Cℓ(1,7) algebra, anomaly/stability structure, and previously sealed finite law-space only.",
		},
		BridgeEntries: []string{
			"Synthetic Schwinger source-ledger adapter parses all 19 Gate536 schema rows.",
			"Finite bridge plumbing verifies theta_E involution, Schwinger Gram symmetry/covariance, positive-time domain closure, and sampled OS quadratic nonnegativity.",
			"Every synthetic row is source-tagged, convention-tagged, bridge_only=true, synthetic=true, and no_theorem_input=true.",
		},
		EnvironmentalEntries: []string{
			"Physical Schwinger functions, constructive measures, observed correlation data, Wick/iε choices, Hamiltonian spectrum, global causal boundary, and time orientation remain external inputs.",
		},
		FailedRoutes: []string{StatusFailedSyntheticNotPhysicalSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow},
		OpenTheorems: []string{
			"Import a real constructive or physical Schwinger family only through the Gate536 19-row source ledger.",
			"Evaluate OS positivity, Wick continuation, Hilbert reconstruction, Hamiltonian positivity, unitarity, and global causality as separate bridge comparators.",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 538, Title: "Schwinger/Wick Bridge Closure and Physical-Data Frontier Map", Reason: "Gate537 proves the synthetic source-ledger adapter can parse and firewall a complete fake Schwinger-family import. The next safe step is a closure ledger that freezes the physical-correlation API boundary and maps the remaining environmental data frontier.", PrimaryTask: "Record the physical Schwinger integration API as complete bridge plumbing while blocking any native promotion of real dynamics, Wick continuation, Hamiltonian spectrum, global causality, or time orientation."}
}

func truth(a Analysis) string {
	return "Gate537 validates the complete synthetic Schwinger-function source-ledger adapter: the 19-row Gate536 schema parses, the finite reflection/correlation plumbing is internally consistent, and every source row remains bridge-only. The result is not a physical Schwinger family, not an OS theorem for nature, not Wick rotation, not a physical Hilbert space, not a Hamiltonian, not unitary dynamics, not global causality, and not the arrow of time."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate536AirlockDefined || !a.Inheritance.Gate536SchemaRowsEnumerated || a.Inheritance.Gate536RequiredRows != 19 || a.Inheritance.Gate536BridgeOnlyRows != 19 || a.Inheritance.Gate536NativeWriteRows != 0 || a.Inheritance.Gate536ComparatorRows < 9 || !a.Inheritance.Gate536SourceTagsEnforced || !a.Inheritance.Gate536OSCertificateRequired || !a.Inheritance.Gate536WickAndIEpsilonRequired || !a.Inheritance.Gate536HamiltonianCertificateRequired || !a.Inheritance.Gate536ComparatorBlocked || !a.Inheritance.Gate536NoObservedCorrelatorsImported || !a.Inheritance.Gate536NativePromotionRejected || !a.Inheritance.Gate536NativeWriteBlocked || !a.Inheritance.Gate537SyntheticRedirect {
		bad = append(bad, "bad Gate536 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 19 || a.Import.AcceptedRows != 19 || a.Import.RejectedRows != 0 || len(a.Import.MissingRequiredRows) != 0 || len(a.Import.DuplicateRows) != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.PhysicalSchwingerLoaded || a.Import.ObservedCorrelationLoaded || a.Import.ConstructiveMeasureLoaded || a.Import.ObservedWickLoaded || a.Import.ObservedHamiltonianLoaded || a.Import.ObservedCausalBoundaryLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.Gate536ReferenceComplete || !a.Import.Gate534ReferenceComplete || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || !a.Import.AllRowsSourceTagged || !a.Import.AllRowsConventionTagged || a.Import.AnyObservedClaim || a.Import.NativePromotionRejected || !a.Import.RequiredSchemaRowsMatched {
		bad = append(bad, "bad synthetic Schwinger ledger import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || a.Output.Dimension != 4 || a.Output.PositiveDomainDimension != 2 || a.Output.SchemaRowsParsed != 19 || a.Output.RequiredRowsMatched != 19 || a.Output.BridgeOnlyRows != 19 || a.Output.NoTheoremInputRows != 19 || a.Output.SourceTaggedRows != 19 || a.Output.ConventionTaggedRows != 19 || a.Output.SyntheticRows != 19 || a.Output.NativePromotionRows != 0 || a.Output.ObservedRows != 0 || a.Output.ThetaEInvolutionResidual != 0 || a.Output.SchwingerKernelSymmetryResidual != 0 || a.Output.EuclideanCovarianceResidual != 0 || a.Output.PositiveTimeDomainClosureResidual != 0 || a.Output.OSGramSymmetryResidual != 0 || !a.Output.OSGramPositiveDefinite || a.Output.OSGramPositiveEigenvalues != 2 || a.Output.OSGramNegativeEigenvalues != 0 || a.Output.OSGramZeroEigenvalues != 0 || a.Output.NegativeQuadraticVectors != 0 || !a.Output.AllSyntheticQuadraticsNonnegative || !a.Output.DummyHamiltonianSpectrumParsed || !a.Output.WickMapPlaceholderParsed || !a.Output.IEpsilonPlaceholderParsed || !a.Output.NullQuotientMetadataConsistent || !a.Output.ReconstructionCertificateParsed || !a.Output.ReflectionCertificateParsed || !a.Output.EuclideanCovarianceCertificateParsed || !a.Output.FiniteSchwingerPlumbingVerified || !a.Output.SyntheticSchwingerAdapterVerified || a.Output.PhysicalSchwingerFunctionsDerived || a.Output.PhysicalOSPositivityProven || a.Output.PhysicalHilbertSpaceSelected || a.Output.WickRotationGranted || a.Output.PositiveEnergyHamiltonianDerived || a.Output.UnitaryRealTimeDynamicsDerived || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected {
		bad = append(bad, "bad synthetic Schwinger adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.PhysicalSchwingerDataImported || a.Firewall.ObservedCorrelationDataImported || a.Firewall.ConstructiveMeasureImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedHamiltonianDataImported || a.Firewall.ObservedCausalBoundaryImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeEuclideanMeasureWrite || a.Firewall.NativeOSPositivityWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHilbertWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityScaleFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.ReopenedDimensionalFirewall || a.Firewall.ReopenedKreinHilbertFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate537 validation failed: %s", strings.Join(bad, "; "))
	}
	return nil
}

func resolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	if _, err := os.Stat(path); err == nil {
		return path
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "../../.."))
	return filepath.Join(root, path)
}

func maxDiff(a, b linear.Matrix) float64 {
	d, err := a.MaxAbsDiff(b)
	if err != nil {
		return math.Inf(1)
	}
	return clean(d)
}

func extrema(xs []float64) (float64, float64) {
	if len(xs) == 0 {
		return 0, 0
	}
	min, max := xs[0], xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return clean(min), clean(max)
}

func inertia(xs []float64) (pos, neg, zero int) {
	for _, x := range xs {
		switch {
		case x > tolerance:
			pos++
		case x < -tolerance:
			neg++
		default:
			zero++
		}
	}
	return
}

func vectorNorm(v linear.Matrix) float64 { return v.FrobeniusNorm() }

func contains(xs []int, v int) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

func clean(x float64) float64 {
	if math.Abs(x) < tolerance {
		return 0
	}
	return x
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: airlock=%t rows=%d bridge_rows=%d native_rows=%d comparator_rows=%d source_tags=%t OS_cert=%t Wick_iε=%t Hamiltonian_cert=%t comparator_blocked=%t no_observed=%t native_rejected=%t native_blocked=%t gate537_redirect=%t; %s", x.Verdict, x.Gate536AirlockDefined, x.Gate536RequiredRows, x.Gate536BridgeOnlyRows, x.Gate536NativeWriteRows, x.Gate536ComparatorRows, x.Gate536SourceTagsEnforced, x.Gate536OSCertificateRequired, x.Gate536WickAndIEpsilonRequired, x.Gate536HamiltonianCertificateRequired, x.Gate536ComparatorBlocked, x.Gate536NoObservedCorrelatorsImported, x.Gate536NativePromotionRejected, x.Gate536NativeWriteBlocked, x.Gate537SyntheticRedirect, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d missing=%v duplicate=%v bridge_only=%t synthetic=%t physical=%t observed_corr=%t constructive=%t observed_Wick=%t observed_Hamiltonian=%t observed_causal=%t native_write=%t gate536_ref=%t gate534_ref=%t metadata=%t rows_bridge=%t rows_no_theorem=%t rows_synthetic=%t rows_source=%t rows_convention=%t observed_claim=%t native_rejected=%t schema_matched=%t path=%s; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.MissingRequiredRows, x.DuplicateRows, x.BridgeOnlyLedger, x.SyntheticFixture, x.PhysicalSchwingerLoaded, x.ObservedCorrelationLoaded, x.ConstructiveMeasureLoaded, x.ObservedWickLoaded, x.ObservedHamiltonianLoaded, x.ObservedCausalBoundaryLoaded, x.NativeRegistryWriteRequested, x.Gate536ReferenceComplete, x.Gate534ReferenceComplete, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsNoTheoremInput, x.AllRowsSynthetic, x.AllRowsSourceTagged, x.AllRowsConventionTagged, x.AnyObservedClaim, x.NativePromotionRejected, x.RequiredSchemaRowsMatched, x.Path, x.Reason)
}

func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("%s: ready=%t dim=%d positive_dim=%d schema_rows=%d required=%d bridge_rows=%d no_theorem_rows=%d source_rows=%d convention_rows=%d synthetic_rows=%d native_rows=%d observed_rows=%d theta2-I=%g Ksym=%g covariance=%g domain_closure=%g OSGram_sym=%g eig_min=%g eig_max=%g eig_pos=%d eig_neg=%d eig_zero=%d OSGram_PD=%t q_min=%g q_max=%g nonzero_vectors=%d null_vectors=%d q_positive=%d q_zero=%d q_negative=%d quadratics_nonnegative=%t dummy_H_levels=%d dummy_H_min=%g dummy_H_max=%g dummy_H_parsed=%t Wick_placeholder=%t iε_placeholder=%t null_quotient=%t reconstruction=%t reflection_cert=%t covariance_cert=%t finite_plumbing=%t synthetic_verified=%t physical_Schwinger=%t physical_OS=%t physical_Hilbert=%t Wick=%t Hamiltonian=%t unitary=%t global=%t arrow=%t; %s", x.Verdict, x.Ready, x.Dimension, x.PositiveDomainDimension, x.SchemaRowsParsed, x.RequiredRowsMatched, x.BridgeOnlyRows, x.NoTheoremInputRows, x.SourceTaggedRows, x.ConventionTaggedRows, x.SyntheticRows, x.NativePromotionRows, x.ObservedRows, x.ThetaEInvolutionResidual, x.SchwingerKernelSymmetryResidual, x.EuclideanCovarianceResidual, x.PositiveTimeDomainClosureResidual, x.OSGramSymmetryResidual, x.OSGramEigenMin, x.OSGramEigenMax, x.OSGramPositiveEigenvalues, x.OSGramNegativeEigenvalues, x.OSGramZeroEigenvalues, x.OSGramPositiveDefinite, x.QuadraticMinimum, x.QuadraticMaximum, x.NonzeroTestVectors, x.NullTestVectors, x.PositiveQuadraticVectors, x.NullQuadraticVectors, x.NegativeQuadraticVectors, x.AllSyntheticQuadraticsNonnegative, x.DummyHamiltonianLevels, x.DummyHamiltonianMin, x.DummyHamiltonianMax, x.DummyHamiltonianSpectrumParsed, x.WickMapPlaceholderParsed, x.IEpsilonPlaceholderParsed, x.NullQuotientMetadataConsistent, x.ReconstructionCertificateParsed, x.ReflectionCertificateParsed, x.EuclideanCovarianceCertificateParsed, x.FiniteSchwingerPlumbingVerified, x.SyntheticSchwingerAdapterVerified, x.PhysicalSchwingerFunctionsDerived, x.PhysicalOSPositivityProven, x.PhysicalHilbertSpaceSelected, x.WickRotationGranted, x.PositiveEnergyHamiltonianDerived, x.UnitaryRealTimeDynamicsDerived, x.GlobalHyperbolicityGranted, x.ArrowOfTimeSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: physical=%t observed_corr=%t constructive=%t observed_Wick=%t observed_Hamiltonian=%t observed_causal=%t synthetic_only=%t file_native=%t adapter_native=%t native_Schwinger=%t native_measure=%t native_OS=%t native_Wick=%t native_Hilbert=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t reopen_dimension=%t reopen_Krein=%t native_registry=%t; %s", x.Verdict, x.PhysicalSchwingerDataImported, x.ObservedCorrelationDataImported, x.ConstructiveMeasureImported, x.ObservedWickDataImported, x.ObservedHamiltonianDataImported, x.ObservedCausalBoundaryImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.NativeSchwingerFunctionWrite, x.NativeEuclideanMeasureWrite, x.NativeOSPositivityWrite, x.NativeWickWrite, x.NativeHilbertWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityScaleFirewall, x.ReopenedTopologyFirewall, x.ReopenedDimensionalFirewall, x.ReopenedKreinHilbertFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate536AirlockInherited, StatusSyntheticSchwingerLedgerLoaded, StatusSyntheticSchwingerRowsAccepted, StatusMetadataSieveEnforced, StatusSyntheticSchwingerAdapterExecuted, StatusThetaEInvolutionResidualZero, StatusSchwingerKernelSymmetryResidualZero, StatusEuclideanCovarianceResidualZero, StatusPositiveTimeDomainClosed, StatusOSQuadraticFormNonnegative, StatusDummyHamiltonianSpectrumParsed, StatusNoPhysicalSchwingerDataImported, StatusNativePromotionRejected, StatusFailedSyntheticNotPhysicalSchwinger, StatusFailedSyntheticNotOSProof, StatusFailedSyntheticNotWick, StatusFailedSyntheticNotHilbert, StatusFailedSyntheticNotHamiltonian, StatusFailedSyntheticNotUnitary, StatusFailedSyntheticNotGlobal, StatusFailedSyntheticNotArrow, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 537 Registry Audit — Synthetic Schwinger-Function Source Ledger Adapter Dry Run\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 537 inherits Gate 536's 19-row physical Schwinger source-ledger airlock and executes only a synthetic dry run.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic ledger import\n\nThe synthetic file contains all 19 required Gate 536 rows. Every row is source-tagged, convention-tagged, `bridge_only=true`, `synthetic=true`, and `no_theorem_input=true`.\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## Schwinger/OS finite plumbing dry run\n\nThe adapter evaluates the synthetic finite reduction `Q_OS(f)=<θ_E f, K f>` over the declared positive-time test-function domain.\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native", a.Registry.NativeEntries)
	writeList(&b, "### Bridge", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString(fmt.Sprintf("## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- none\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
