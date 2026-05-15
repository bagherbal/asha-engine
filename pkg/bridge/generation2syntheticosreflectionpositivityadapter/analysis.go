// Package generation2syntheticosreflectionpositivityadapter implements Gate 534:
// Synthetic OS Reflection-Positivity Kernel Adapter Dry Run.
//
// Gate 533 defined the Osterwalder-Schrader kernel/test-domain airlock and
// deliberately refused to infer reflection positivity from finite H=GΘ matrix
// positivity. This package loads a deliberately synthetic, source-tagged OS
// kernel fixture and evaluates the finite reflection-positivity plumbing:
// reflection involution, kernel symmetry, positive-time domain closure, OS Gram
// positivity, explicit test-vector quadratic forms, null-quotient metadata, and
// Gate 532 Θ compatibility. Passing this adapter is still bridge-only. It does
// not derive physical Schwinger functions, Wick rotation, a Hamiltonian,
// unitarity, global hyperbolicity, or the arrow of time.
package generation2syntheticosreflectionpositivityadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2osreflectionpositivityairlock"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID       = "GATE534-SYNTHETIC-OS-REFLECTION-POSITIVITY-KERNEL-ADAPTER-DRY-RUN"
	DefaultLedger = "data/synthetic_os_reflection_positivity_ledger_gate534.json"

	StatusGate533AirlockInherited           = "CONDITIONAL_SUPPORT_GATE533_OS_KERNEL_AIRLOCK_INHERITED"
	StatusSyntheticOSLedgerLoaded           = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_KERNEL_LEDGER_LOADED"
	StatusSyntheticOSRowAccepted            = "CONDITIONAL_SUPPORT_GATE534_AIRLOCK_ACCEPTED_SYNTHETIC_OS_ROW"
	StatusSyntheticOSAdapterExecuted        = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_ADAPTER_EXECUTED"
	StatusReflectionInvolutionResidualZero  = "CONDITIONAL_SUPPORT_EUCLIDEAN_REFLECTION_INVOLUTION_RESIDUAL_ZERO"
	StatusKernelSymmetryResidualZero        = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_KERNEL_SYMMETRY_RESIDUAL_ZERO"
	StatusReflectionCovarianceResidualZero  = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_REFLECTION_COVARIANCE_RESIDUAL_ZERO"
	StatusPositiveTimeDomainClosed          = "CONDITIONAL_SUPPORT_SYNTHETIC_POSITIVE_TIME_DOMAIN_CLOSED"
	StatusOSGramPositiveDefinite            = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_GRAM_POSITIVE_DEFINITE"
	StatusOSQuadraticFormPositive           = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_QUADRATIC_FORM_NONNEGATIVE"
	StatusNullQuotientMetadataConsistent    = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_NULL_QUOTIENT_METADATA_CONSISTENT"
	StatusGate532ThetaCompatibilityDeclared = "CONDITIONAL_SUPPORT_SYNTHETIC_OS_GATE532_THETA_COMPATIBILITY_DECLARED"
	StatusNoObservedOSDataImportedDefault   = "CONDITIONAL_SUPPORT_NO_OBSERVED_OS_WICK_OR_CORRELATION_DATA_IMPORTED_BY_DEFAULT"

	StatusFailedLedgerMissing               = "FAILED_ROUTE_GATE534_SYNTHETIC_OS_LEDGER_MISSING"
	StatusFailedMetadataIncomplete          = "FAILED_ROUTE_GATE534_OS_KERNEL_METADATA_INCOMPLETE"
	StatusFailedInvalidMatrixDomain         = "FAILED_ROUTE_GATE534_INVALID_OS_KERNEL_MATRIX_DOMAIN"
	StatusFailedReflectionInvolutionNonzero = "FAILED_ROUTE_GATE534_REFLECTION_INVOLUTION_RESIDUAL_NONZERO"
	StatusFailedKernelSymmetryNonzero       = "FAILED_ROUTE_GATE534_OS_KERNEL_SYMMETRY_RESIDUAL_NONZERO"
	StatusFailedReflectionCovarianceNonzero = "FAILED_ROUTE_GATE534_OS_REFLECTION_COVARIANCE_RESIDUAL_NONZERO"
	StatusFailedPositiveTimeDomainNotClosed = "FAILED_ROUTE_GATE534_POSITIVE_TIME_DOMAIN_CLOSURE_FAILED"
	StatusFailedOSGramNotPositive           = "FAILED_ROUTE_GATE534_OS_GRAM_NOT_POSITIVE_DEFINITE"
	StatusFailedOSQuadraticFormNegative     = "FAILED_ROUTE_GATE534_OS_QUADRATIC_FORM_NEGATIVE_ON_SYNTHETIC_DOMAIN"
	StatusFailedSyntheticOSNative           = "FAILED_ROUTE_SYNTHETIC_OS_KERNEL_NATIVE_PROMOTION_REJECTED"
	StatusFailedSyntheticOSNotSchwinger     = "FAILED_ROUTE_SYNTHETIC_OS_KERNEL_DOES_NOT_DERIVE_PHYSICAL_SCHWINGER_FUNCTIONS"
	StatusFailedSyntheticOSNotWick          = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedSyntheticOSNotHilbert       = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedSyntheticOSNotHamiltonian   = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_DERIVE_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedSyntheticOSNotUnitary       = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedSyntheticOSNotGlobal        = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticOSNotArrow         = "FAILED_ROUTE_SYNTHETIC_OS_POSITIVITY_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_GATE534_SYNTHETIC_OS_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked        = "FIREWALL_BLOCKED_GATE534_OS_WICK_HILBERT_HAMILTONIAN_NATIVE_WRITE"
)

const tolerance = 1e-12

type Inheritance struct {
	Executed bool

	Gate533AirlockDefined             bool
	Gate533SchemaRowsEnumerated       bool
	Gate533RequiresReflection         bool
	Gate533RequiresTestDomain         bool
	Gate533RequiresKernel             bool
	Gate533RequiresQuadraticForm      bool
	Gate533RequiresNullQuotient       bool
	Gate533RequiresThetaCompatibility bool
	Gate533RequiresWickIEpsilon       bool
	Gate533RequiresSourceConvention   bool
	Gate533ComparatorBlocked          bool
	Gate533OSWickHilbertBlocked       bool
	Gate533NativeWriteBlocked         bool
	Gate533NoObservedDataImported     bool
	Gate534SyntheticRedirect          bool

	Verdict, Reason string
}

type OSKernelRow struct {
	Name                             string        `json:"name"`
	EuclideanReflectionOperator      [][]float64   `json:"euclidean_reflection_operator"`
	CorrelationKernelGramMatrix      [][]float64   `json:"correlation_kernel_gram_matrix"`
	PositiveTimeSupport              []int         `json:"positive_time_support"`
	PositiveTimeTranslationOperators [][][]float64 `json:"positive_time_translation_operators"`
	TestVectors                      [][]float64   `json:"test_vectors"`
	NullSpaceQuotientRule            string        `json:"null_space_quotient_rule"`
	ReconstructionMapCertificate     string        `json:"reconstruction_map_certificate"`
	Gate532ThetaCompatibility        string        `json:"gate532_theta_compatibility"`
	WickMapReference                 string        `json:"wick_map_reference"`
	IepsilonConvention               string        `json:"i_epsilon_convention"`
	KernelSymmetryConvention         string        `json:"kernel_symmetry_convention"`
	ReflectionPositiveCone           string        `json:"reflection_positive_cone"`
	OSQuadraticFormDefinition        string        `json:"os_quadratic_form_definition"`
	Source                           string        `json:"source"`
	SourceVersion                    string        `json:"source_version"`
	Convention                       string        `json:"convention"`
	Uncertainty                      string        `json:"uncertainty"`
	BridgeOnly                       bool          `json:"bridge_only"`
	ComparatorOnly                   bool          `json:"comparator_only"`
	Synthetic                        bool          `json:"synthetic"`
	Observed                         bool          `json:"observed"`
	NoTheoremInput                   bool          `json:"no_theorem_input"`
	NativePromotion                  bool          `json:"native_promotion"`
	NativeInputClaim                 bool          `json:"native_input_claim"`
}

type OSKernelLedger struct {
	Gate                      int           `json:"gate"`
	LedgerName                string        `json:"ledger_name"`
	Description               string        `json:"description"`
	Gate533AirlockReference   string        `json:"gate533_airlock_reference"`
	Gate532ThetaReference     string        `json:"gate532_theta_reference"`
	BridgeOnly                bool          `json:"bridge_only"`
	NativeRegistryWrite       bool          `json:"native_registry_write"`
	SyntheticFixture          bool          `json:"synthetic_fixture"`
	ObservedOSLoaded          bool          `json:"observed_os_loaded"`
	ObservedWickLoaded        bool          `json:"observed_wick_loaded"`
	ObservedCorrelationLoaded bool          `json:"observed_correlation_loaded"`
	ObservedHamiltonianLoaded bool          `json:"observed_hamiltonian_loaded"`
	Source                    string        `json:"source"`
	SourceVersion             string        `json:"source_version"`
	Convention                string        `json:"convention"`
	Rows                      []OSKernelRow `json:"rows"`
}

type FileImport struct {
	Executed                     bool
	Loaded                       bool
	Path                         string
	Rows                         int
	AcceptedRows                 int
	RejectedRows                 int
	BridgeOnlyLedger             bool
	SyntheticFixture             bool
	ObservedOSLoaded             bool
	ObservedWickLoaded           bool
	ObservedCorrelationLoaded    bool
	ObservedHamiltonianLoaded    bool
	NativeRegistryWriteRequested bool
	Gate533ReferenceComplete     bool
	Gate532ReferenceComplete     bool
	MetadataComplete             bool
	AllRowsBridgeOnly            bool
	AllRowsComparatorOnly        bool
	AllRowsSynthetic             bool
	AllRowsNoTheoremInput        bool
	AnyObservedClaim             bool
	NativePromotionRejected      bool
	Verdict, Reason              string
	Failures                     []string
}

type AdapterInput struct {
	Rows                         int
	Reflection                   linear.Matrix
	Kernel                       linear.Matrix
	PositiveProjection           linear.Matrix
	PositiveTimeSupport          []int
	Translations                 []linear.Matrix
	TestVectors                  []linear.Matrix
	NullSpaceQuotientRule        string
	ReconstructionMapCertificate string
	Gate532ThetaCompatibility    string
	WickMapReference             string
	IepsilonConvention           string
	KernelSymmetryConvention     string
	ReflectionPositiveCone       string
	OSQuadraticFormDefinition    string
	BridgeOnly                   bool
	SyntheticFixture             bool
	ObservedOSLoaded             bool
	ObservedWickLoaded           bool
	ObservedCorrelationLoaded    bool
	ObservedHamiltonianLoaded    bool
	NativePromotion              bool
	MetadataComplete             bool
}

type AdapterOutput struct {
	Executed  bool
	Attempted bool
	Ready     bool

	Dimension                         int
	PositiveDomainDimension           int
	ComparatorOnly                    bool
	BridgeOnly                        bool
	NativePrediction                  bool
	ReflectionInvolutionResidual      float64
	KernelSymmetryResidual            float64
	ReflectionCovarianceResidual      float64
	PositiveTimeDomainClosureResidual float64
	OSGramSymmetryResidual            float64
	OSGramEigenMin                    float64
	OSGramEigenMax                    float64
	OSGramPositiveEigenvalues         int
	OSGramNegativeEigenvalues         int
	OSGramZeroEigenvalues             int
	OSGramPositiveDefinite            bool
	QuadraticMinimum                  float64
	QuadraticMaximum                  float64
	NonzeroTestVectors                int
	NullTestVectors                   int
	PositiveQuadraticVectors          int
	NullQuadraticVectors              int
	NegativeQuadraticVectors          int
	AllSyntheticQuadraticsNonnegative bool
	NullQuotientMetadataConsistent    bool
	Gate532ThetaCompatibilityDeclared bool
	FiniteOSPlumbingVerified          bool
	SyntheticOSPositivityVerified     bool
	PhysicalSchwingerFunctionsDerived bool
	PhysicalHilbertSpaceSelected      bool
	WickRotationGranted               bool
	PositiveEnergyHamiltonianDerived  bool
	UnitaryRealTimeDynamicsDerived    bool
	GlobalHyperbolicityGranted        bool
	ArrowOfTimeSelected               bool
	Verdict, Reason                   string
	Failures                          []string
}

type Firewall struct {
	Executed bool

	ObservedOSDataImported          bool
	ObservedWickDataImported        bool
	ObservedCorrelationDataImported bool
	ObservedHamiltonianDataImported bool
	SyntheticFixtureOnly            bool
	FileRowsNative                  bool
	AdapterOutputsNative            bool
	NativeOSKernelWrite             bool
	NativeSchwingerFunctionWrite    bool
	NativeHilbertProductWrite       bool
	NativePhysicalStateSpaceWrite   bool
	NativeWickWrite                 bool
	NativeHamiltonianWrite          bool
	NativeUnitaryDynamicsWrite      bool
	NativeGlobalCausalWrite         bool
	NativeTimeArrowWrite            bool
	ReopenedFlavorFirewall          bool
	ReopenedEWScaleFirewall         bool
	ReopenedGravityFirewall         bool
	ReopenedTopologyFirewall        bool
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
	g533, err := generation2osreflectionpositivityairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate533 OS kernel airlock: %w", err)
	}
	a := Analysis{Inheritance: buildInheritance(g533)}
	ledger, imp := loadLedger(path)
	a.Import = imp
	if imp.Loaded && imp.MetadataComplete && imp.AcceptedRows == 1 {
		in, err := buildInput(ledger, imp)
		if err != nil {
			a.Input = AdapterInput{Rows: len(ledger.Rows), MetadataComplete: imp.MetadataComplete, BridgeOnly: ledger.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture, ObservedOSLoaded: ledger.ObservedOSLoaded, ObservedWickLoaded: ledger.ObservedWickLoaded, ObservedCorrelationLoaded: ledger.ObservedCorrelationLoaded, ObservedHamiltonianLoaded: ledger.ObservedHamiltonianLoaded, NativePromotion: ledger.NativeRegistryWrite}
			a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
		} else {
			a.Input = in
			a.Output = runAdapter(in)
		}
	} else if !imp.Loaded {
		a.Output = AdapterOutput{Executed: true, Attempted: false, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedLedgerMissing, Reason: "explicit Gate534 synthetic OS ledger was not found", Failures: []string{StatusFailedLedgerMissing}}
	} else {
		a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedMetadataIncomplete, Reason: "Gate534 OS kernel file did not satisfy the Gate533 airlock metadata domain", Failures: []string{StatusFailedMetadataIncomplete}}
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

func buildInheritance(g533 generation2osreflectionpositivityairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate533AirlockDefined:             g533.Schema.Executed && g533.Schema.EuclideanReflectionOperatorRequired,
		Gate533SchemaRowsEnumerated:       g533.Schema.RequiredRowCount >= 19,
		Gate533RequiresReflection:         g533.Schema.EuclideanReflectionOperatorRequired && g533.Schema.ReflectionActionRequired,
		Gate533RequiresTestDomain:         g533.Schema.TestFunctionDomainRequired && g533.Schema.ReflectionPositiveConeRequired,
		Gate533RequiresKernel:             g533.Schema.CorrelationKernelRequired && g533.Schema.KernelHermiticityCheckRequired,
		Gate533RequiresQuadraticForm:      g533.Schema.OSQuadraticFormCheckRequired,
		Gate533RequiresNullQuotient:       g533.Schema.NullSpaceQuotientRequired,
		Gate533RequiresThetaCompatibility: g533.Schema.CompatibilityWithThetaRequired,
		Gate533RequiresWickIEpsilon:       g533.Schema.WickMapReferenceRequired && g533.Schema.IepsilonConventionRequired,
		Gate533RequiresSourceConvention:   g533.Schema.SourceRequired && g533.Schema.ConventionRequired && g533.Schema.BridgeOnlyRequired && g533.Schema.ComparatorOnlyRequired && g533.Schema.NoTheoremInputRequired,
		Gate533ComparatorBlocked:          !g533.Guard.ComparatorExecutionPerformed,
		Gate533OSWickHilbertBlocked:       !g533.Guard.ReflectionPositivityProven && !g533.Guard.WickRotationSelected && !g533.Guard.PhysicalHilbertSpaceSelected && !g533.Guard.PositiveEnergyHamiltonianDerived && !g533.Guard.UnitaryRealTimeDynamicsDerived && !g533.Guard.GlobalHyperbolicitySelected,
		Gate533NativeWriteBlocked:         !g533.Firewall.NativeRegistryWritten && !g533.Firewall.NativeOSKernelWrite && !g533.Firewall.NativeWickWrite && !g533.Firewall.NativeHilbertProductWrite,
		Gate533NoObservedDataImported:     !g533.Firewall.ObservedOSDataImported && !g533.Firewall.ObservedWickDataImported && !g533.Firewall.ObservedCorrelationDataImported && !g533.Firewall.ObservedHamiltonianDataImported,
		Gate534SyntheticRedirect:          g533.Next.Gate == 534,
		Verdict:                           StatusGate533AirlockInherited,
		Reason:                            "Gate534 inherits Gate533's OS kernel/test-domain airlock and now executes only a synthetic bridge-only kernel dry run. The inherited firewall forbids native OS/Wick/Hilbert promotion.",
	}
}

func loadLedger(path string) (OSKernelLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved, Verdict: StatusFailedLedgerMissing, Reason: "ledger not loaded", Failures: []string{StatusFailedLedgerMissing}}
	b, err := os.ReadFile(resolved)
	if err != nil {
		return OSKernelLedger{}, imp
	}
	var ledger OSKernelLedger
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
	imp.ObservedOSLoaded = ledger.ObservedOSLoaded
	imp.ObservedWickLoaded = ledger.ObservedWickLoaded
	imp.ObservedCorrelationLoaded = ledger.ObservedCorrelationLoaded
	imp.ObservedHamiltonianLoaded = ledger.ObservedHamiltonianLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.Gate533ReferenceComplete = strings.TrimSpace(ledger.Gate533AirlockReference) != ""
	imp.Gate532ReferenceComplete = strings.TrimSpace(ledger.Gate532ThetaReference) != ""
	imp.AllRowsBridgeOnly = true
	imp.AllRowsComparatorOnly = true
	imp.AllRowsSynthetic = true
	imp.AllRowsNoTheoremInput = true
	for _, row := range ledger.Rows {
		rowOK := validateRowMetadata(row)
		if rowOK {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
		imp.AllRowsBridgeOnly = imp.AllRowsBridgeOnly && row.BridgeOnly
		imp.AllRowsComparatorOnly = imp.AllRowsComparatorOnly && row.ComparatorOnly
		imp.AllRowsSynthetic = imp.AllRowsSynthetic && row.Synthetic
		imp.AllRowsNoTheoremInput = imp.AllRowsNoTheoremInput && row.NoTheoremInput
		imp.AnyObservedClaim = imp.AnyObservedClaim || row.Observed
		if row.NativePromotion || row.NativeInputClaim {
			imp.NativePromotionRejected = true
		}
	}
	imp.MetadataComplete = ledger.Gate == 534 && imp.Rows == 1 && imp.AcceptedRows == 1 && imp.BridgeOnlyLedger && !imp.NativeRegistryWriteRequested && imp.SyntheticFixture && !imp.ObservedOSLoaded && !imp.ObservedWickLoaded && !imp.ObservedCorrelationLoaded && !imp.ObservedHamiltonianLoaded && imp.Gate533ReferenceComplete && imp.Gate532ReferenceComplete && imp.AllRowsBridgeOnly && imp.AllRowsComparatorOnly && imp.AllRowsSynthetic && imp.AllRowsNoTheoremInput && !imp.AnyObservedClaim && !imp.NativePromotionRejected && strings.TrimSpace(ledger.Source) != "" && strings.TrimSpace(ledger.SourceVersion) != "" && strings.TrimSpace(ledger.Convention) != ""
	if imp.MetadataComplete {
		imp.Verdict = strings.Join([]string{StatusSyntheticOSLedgerLoaded, StatusSyntheticOSRowAccepted}, ";")
		imp.Reason = "Gate534 synthetic OS kernel ledger loaded with complete bridge-only source/convention metadata and no observed or native-promotion claims."
		imp.Failures = nil
	} else {
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = "ledger metadata or row flags violate the Gate533 OS kernel airlock"
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		if imp.NativePromotionRejected || imp.NativeRegistryWriteRequested {
			imp.Failures = append(imp.Failures, StatusFailedSyntheticOSNative)
		}
	}
	return ledger, imp
}

func validateRowMetadata(row OSKernelRow) bool {
	return row.BridgeOnly && row.ComparatorOnly && row.Synthetic && row.NoTheoremInput && !row.Observed && !row.NativePromotion && !row.NativeInputClaim && strings.TrimSpace(row.Name) != "" && strings.TrimSpace(row.NullSpaceQuotientRule) != "" && strings.TrimSpace(row.ReconstructionMapCertificate) != "" && strings.TrimSpace(row.Gate532ThetaCompatibility) != "" && strings.TrimSpace(row.WickMapReference) != "" && strings.TrimSpace(row.IepsilonConvention) != "" && strings.TrimSpace(row.KernelSymmetryConvention) != "" && strings.TrimSpace(row.ReflectionPositiveCone) != "" && strings.TrimSpace(row.OSQuadraticFormDefinition) != "" && strings.TrimSpace(row.Source) != "" && strings.TrimSpace(row.SourceVersion) != "" && strings.TrimSpace(row.Convention) != "" && len(row.PositiveTimeSupport) > 0 && len(row.TestVectors) > 0
}

func buildInput(ledger OSKernelLedger, imp FileImport) (AdapterInput, error) {
	if len(ledger.Rows) != 1 {
		return AdapterInput{}, fmt.Errorf("Gate534 adapter expects exactly one synthetic OS row; got %d", len(ledger.Rows))
	}
	row := ledger.Rows[0]
	R, err := linear.FromRows(row.EuclideanReflectionOperator)
	if err != nil {
		return AdapterInput{}, fmt.Errorf("reflection operator: %w", err)
	}
	K, err := linear.FromRows(row.CorrelationKernelGramMatrix)
	if err != nil {
		return AdapterInput{}, fmt.Errorf("kernel matrix: %w", err)
	}
	if R.Rows() != R.Cols() || K.Rows() != K.Cols() || R.Rows() != K.Rows() {
		return AdapterInput{}, fmt.Errorf("reflection and kernel must be square over the same finite domain: R=%dx%d K=%dx%d", R.Rows(), R.Cols(), K.Rows(), K.Cols())
	}
	n := R.Rows()
	P := linear.NewMatrix(n, n)
	seen := map[int]bool{}
	for _, idx := range row.PositiveTimeSupport {
		if idx < 0 || idx >= n || seen[idx] {
			return AdapterInput{}, fmt.Errorf("invalid positive-time support index %d", idx)
		}
		seen[idx] = true
		P.Set(idx, idx, 1)
	}
	translations := make([]linear.Matrix, 0, len(row.PositiveTimeTranslationOperators))
	for i, raw := range row.PositiveTimeTranslationOperators {
		T, err := linear.FromRows(raw)
		if err != nil {
			return AdapterInput{}, fmt.Errorf("positive-time translation %d: %w", i, err)
		}
		if T.Rows() != n || T.Cols() != n {
			return AdapterInput{}, fmt.Errorf("positive-time translation %d dimension mismatch: %dx%d expected %dx%d", i, T.Rows(), T.Cols(), n, n)
		}
		translations = append(translations, T)
	}
	if len(translations) == 0 {
		return AdapterInput{}, fmt.Errorf("at least one positive-time translation operator is required")
	}
	testVectors := make([]linear.Matrix, 0, len(row.TestVectors))
	for i, raw := range row.TestVectors {
		if len(raw) != n {
			return AdapterInput{}, fmt.Errorf("test vector %d dimension mismatch: got %d expected %d", i, len(raw), n)
		}
		v := linear.NewMatrix(n, 1)
		for r, x := range raw {
			v.Set(r, 0, x)
		}
		testVectors = append(testVectors, v)
	}
	return AdapterInput{Rows: len(ledger.Rows), Reflection: R, Kernel: K, PositiveProjection: P, PositiveTimeSupport: append([]int(nil), row.PositiveTimeSupport...), Translations: translations, TestVectors: testVectors, NullSpaceQuotientRule: row.NullSpaceQuotientRule, ReconstructionMapCertificate: row.ReconstructionMapCertificate, Gate532ThetaCompatibility: row.Gate532ThetaCompatibility, WickMapReference: row.WickMapReference, IepsilonConvention: row.IepsilonConvention, KernelSymmetryConvention: row.KernelSymmetryConvention, ReflectionPositiveCone: row.ReflectionPositiveCone, OSQuadraticFormDefinition: row.OSQuadraticFormDefinition, BridgeOnly: ledger.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture, ObservedOSLoaded: ledger.ObservedOSLoaded, ObservedWickLoaded: ledger.ObservedWickLoaded, ObservedCorrelationLoaded: ledger.ObservedCorrelationLoaded, ObservedHamiltonianLoaded: ledger.ObservedHamiltonianLoaded, NativePromotion: ledger.NativeRegistryWrite, MetadataComplete: imp.MetadataComplete}, nil
}

func runAdapter(in AdapterInput) AdapterOutput {
	n := in.Reflection.Rows()
	failures := []string{}
	R2, _ := in.Reflection.Mul(in.Reflection)
	reflectionResidual, _ := R2.MaxAbsDiff(linear.Identity(n))
	kernelSymResidual, _ := in.Kernel.MaxAbsDiff(in.Kernel.Transpose())
	RK, _ := in.Reflection.Mul(in.Kernel)
	RKR, _ := RK.Mul(in.Reflection)
	reflectionCovarianceResidual, _ := RKR.MaxAbsDiff(in.Kernel)
	closureResidual := positiveTimeClosureResidual(in.PositiveProjection, in.Translations)
	osGram, err := osGram(in.Reflection, in.Kernel, in.PositiveTimeSupport)
	if err != nil {
		return AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
	}
	gramSymResidual, _ := osGram.MaxAbsDiff(osGram.Transpose())
	eigs, err := linear.SymmetricEigenJacobi(osGram, tolerance, 1000)
	if err != nil {
		return AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, NativePrediction: false, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
	}
	minEig, maxEig := extrema(eigs.Values)
	pos, neg, zero := inertia(eigs.Values)
	qmin, qmax := math.Inf(1), math.Inf(-1)
	nonzeroVecs, nullVecs, posVecs, zeroQ, negQ := 0, 0, 0, 0, 0
	for _, f := range in.TestVectors {
		q := osQuadratic(in.Reflection, in.Kernel, f)
		if q < qmin {
			qmin = q
		}
		if q > qmax {
			qmax = q
		}
		if f.FrobeniusNorm() <= tolerance {
			nullVecs++
		} else {
			nonzeroVecs++
		}
		switch {
		case q > tolerance:
			posVecs++
		case math.Abs(q) <= tolerance:
			zeroQ++
		default:
			negQ++
		}
	}
	if len(in.TestVectors) == 0 {
		qmin, qmax = 0, 0
	}
	if reflectionResidual > tolerance {
		failures = append(failures, StatusFailedReflectionInvolutionNonzero)
	}
	if kernelSymResidual > tolerance {
		failures = append(failures, StatusFailedKernelSymmetryNonzero)
	}
	if reflectionCovarianceResidual > tolerance {
		failures = append(failures, StatusFailedReflectionCovarianceNonzero)
	}
	if closureResidual > tolerance {
		failures = append(failures, StatusFailedPositiveTimeDomainNotClosed)
	}
	if !(gramSymResidual <= tolerance && pos == len(in.PositiveTimeSupport) && neg == 0 && zero == 0 && minEig > tolerance) {
		failures = append(failures, StatusFailedOSGramNotPositive)
	}
	if negQ > 0 || posVecs != nonzeroVecs || zeroQ != nullVecs {
		failures = append(failures, StatusFailedOSQuadraticFormNegative)
	}
	nullOK := strings.Contains(strings.ToLower(in.NullSpaceQuotientRule), "zero") && zeroQ == nullVecs
	thetaOK := strings.Contains(in.Gate532ThetaCompatibility, "Gate532") || strings.Contains(in.Gate532ThetaCompatibility, "gate532")
	if !nullOK {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if !thetaOK {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	ok := len(failures) == 0
	verdicts := []string{StatusSyntheticOSAdapterExecuted}
	if reflectionResidual == 0 {
		verdicts = append(verdicts, StatusReflectionInvolutionResidualZero)
	}
	if kernelSymResidual == 0 {
		verdicts = append(verdicts, StatusKernelSymmetryResidualZero)
	}
	if reflectionCovarianceResidual == 0 {
		verdicts = append(verdicts, StatusReflectionCovarianceResidualZero)
	}
	if closureResidual == 0 {
		verdicts = append(verdicts, StatusPositiveTimeDomainClosed)
	}
	if pos == len(in.PositiveTimeSupport) && neg == 0 && zero == 0 && minEig > tolerance {
		verdicts = append(verdicts, StatusOSGramPositiveDefinite)
	}
	if negQ == 0 && posVecs == nonzeroVecs && zeroQ == nullVecs {
		verdicts = append(verdicts, StatusOSQuadraticFormPositive)
	}
	if nullOK {
		verdicts = append(verdicts, StatusNullQuotientMetadataConsistent)
	}
	if thetaOK {
		verdicts = append(verdicts, StatusGate532ThetaCompatibilityDeclared)
	}
	if !ok {
		verdicts = failures
	}
	return AdapterOutput{Executed: true, Attempted: true, Ready: ok, Dimension: n, PositiveDomainDimension: len(in.PositiveTimeSupport), ComparatorOnly: true, BridgeOnly: in.BridgeOnly, NativePrediction: false, ReflectionInvolutionResidual: clean(reflectionResidual), KernelSymmetryResidual: clean(kernelSymResidual), ReflectionCovarianceResidual: clean(reflectionCovarianceResidual), PositiveTimeDomainClosureResidual: clean(closureResidual), OSGramSymmetryResidual: clean(gramSymResidual), OSGramEigenMin: clean(minEig), OSGramEigenMax: clean(maxEig), OSGramPositiveEigenvalues: pos, OSGramNegativeEigenvalues: neg, OSGramZeroEigenvalues: zero, OSGramPositiveDefinite: pos == len(in.PositiveTimeSupport) && neg == 0 && zero == 0 && minEig > tolerance, QuadraticMinimum: clean(qmin), QuadraticMaximum: clean(qmax), NonzeroTestVectors: nonzeroVecs, NullTestVectors: nullVecs, PositiveQuadraticVectors: posVecs, NullQuadraticVectors: zeroQ, NegativeQuadraticVectors: negQ, AllSyntheticQuadraticsNonnegative: negQ == 0 && posVecs == nonzeroVecs && zeroQ == nullVecs, NullQuotientMetadataConsistent: nullOK, Gate532ThetaCompatibilityDeclared: thetaOK, FiniteOSPlumbingVerified: ok, SyntheticOSPositivityVerified: ok, PhysicalSchwingerFunctionsDerived: false, PhysicalHilbertSpaceSelected: false, WickRotationGranted: false, PositiveEnergyHamiltonianDerived: false, UnitaryRealTimeDynamicsDerived: false, GlobalHyperbolicityGranted: false, ArrowOfTimeSelected: false, Verdict: strings.Join(verdicts, ";"), Reason: "Gate534 executes only a synthetic finite OS kernel plumbing check. Positivity of this fixture does not derive physical Schwinger functions, Wick rotation, Hamiltonian dynamics, unitarity, global hyperbolicity, or time orientation.", Failures: failures}
}

func positiveTimeClosureResidual(P linear.Matrix, translations []linear.Matrix) float64 {
	n := P.Rows()
	Q, _ := linear.Identity(n).Sub(P)
	max := 0.0
	for _, T := range translations {
		TP, _ := T.Mul(P)
		QTP, _ := Q.Mul(TP)
		if v := QTP.FrobeniusNorm(); v > max {
			max = v
		}
	}
	return clean(max)
}

func osGram(R, K linear.Matrix, support []int) (linear.Matrix, error) {
	RTK, err := R.Transpose().Mul(K)
	if err != nil {
		return linear.Matrix{}, err
	}
	out := linear.NewMatrix(len(support), len(support))
	for i, r := range support {
		for j, c := range support {
			out.Set(i, j, RTK.At(r, c))
		}
	}
	return out, nil
}

func osQuadratic(R, K, f linear.Matrix) float64 {
	thetaF, _ := R.Mul(f)
	kf, _ := K.Mul(f)
	q, _ := thetaF.Transpose().Mul(kf)
	return clean(q.At(0, 0))
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{Executed: true, ObservedOSDataImported: imp.ObservedOSLoaded, ObservedWickDataImported: imp.ObservedWickLoaded, ObservedCorrelationDataImported: imp.ObservedCorrelationLoaded, ObservedHamiltonianDataImported: imp.ObservedHamiltonianLoaded, SyntheticFixtureOnly: imp.SyntheticFixture && !imp.AnyObservedClaim, FileRowsNative: imp.NativePromotionRejected || imp.NativeRegistryWriteRequested, AdapterOutputsNative: out.NativePrediction || out.PhysicalSchwingerFunctionsDerived || out.PhysicalHilbertSpaceSelected || out.WickRotationGranted || out.PositiveEnergyHamiltonianDerived || out.UnitaryRealTimeDynamicsDerived || out.GlobalHyperbolicityGranted || out.ArrowOfTimeSelected, NativeRegistryWritten: false, Verdict: strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticOSNotSchwinger, StatusFailedSyntheticOSNotWick, StatusFailedSyntheticOSNotHilbert, StatusFailedSyntheticOSNotHamiltonian, StatusFailedSyntheticOSNotUnitary, StatusFailedSyntheticOSNotGlobal, StatusFailedSyntheticOSNotArrow}, ";"), Reason: "Gate534 proves only that the OS adapter plumbing accepts a synthetic reflection-positive fixture. It writes no native Schwinger functions, Wick rotation, Hilbert reconstruction, Hamiltonian, unitary dynamics, global-causal structure, or arrow of time."}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native Osterwalder-Schrader kernel, Schwinger function, Wick map, Hamiltonian, or physical Hilbert space is written at Gate534.",
			"The synthetic positive OS Gram matrix is bridge plumbing only; it is not a derivation of the universe's correlation functions.",
			"The Lorentzian quantum-dynamics firewall remains closed after the dry run.",
		},
		BridgeEntries: []string{
			"Synthetic OS kernel ledger adapter loaded a source-tagged bridge-only finite fixture.",
			"The adapter verified reflection involution, kernel symmetry, reflection covariance, positive-time domain closure, OS Gram positive definiteness, test-vector quadratic nonnegativity, null-quotient metadata, and Gate532 Θ compatibility.",
			"The result confirms the OS socket can carry a mathematically clean finite reflection-positive fixture without promoting it to native physics.",
		},
		EnvironmentalEntries: []string{
			"Physical Schwinger functions, Euclidean measure, analytic continuation, iε prescription, positive-energy Hamiltonian domain, global causality, and time orientation remain environmental or future sourced bridge data.",
		},
		FailedRoutes: []string{StatusFailedSyntheticOSNotSchwinger, StatusFailedSyntheticOSNotWick, StatusFailedSyntheticOSNotHilbert, StatusFailedSyntheticOSNotHamiltonian, StatusFailedSyntheticOSNotUnitary, StatusFailedSyntheticOSNotGlobal, StatusFailedSyntheticOSNotArrow},
		OpenTheorems: []string{
			"replace the synthetic OS fixture with a sourced physical or constructive Euclidean correlation kernel only through the bridge airlock",
			"audit whether any native ASHA spectral/action data can generate a reflection-positive Schwinger kernel without empirical input",
			"keep Wick rotation and Hamiltonian reconstruction as separate comparators even when OS positivity passes",
		},
	}
}

func buildNext(a Analysis) NextStep {
	return NextStep{Gate: 535, Title: "OS/Wick/Hilbert Sector Closure Ledger and Frontier Map", Reason: "Gate534 completes the synthetic dry run of the OS kernel socket. The next safe step is a consolidation ledger that marks the native/bridge/environmental frontier instead of opening another physics-promotion shortcut.", PrimaryTask: "Produce a sector-closing theorem map for the dimensional, Wick, Hilbert, OS, Hamiltonian, unitarity, global-causal, and time-arrow firewalls, listing exactly what is native, bridge-compatible, environmental, and still open."}
}

func truth(a Analysis) string {
	return "Gate534 proves that the Gate533 OS reflection-positivity socket has working finite bridge plumbing: a synthetic reflection, positive-time domain, and positive kernel fixture can pass the quadratic-form checks. It does not prove physical Schwinger functions, Wick rotation, a real Hilbert space of the universe, a positive-energy Hamiltonian, unitary real-time dynamics, global hyperbolicity, or an arrow of time."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate533AirlockDefined || !a.Inheritance.Gate533SchemaRowsEnumerated || !a.Inheritance.Gate533RequiresReflection || !a.Inheritance.Gate533RequiresTestDomain || !a.Inheritance.Gate533RequiresKernel || !a.Inheritance.Gate533RequiresQuadraticForm || !a.Inheritance.Gate533RequiresNullQuotient || !a.Inheritance.Gate533RequiresThetaCompatibility || !a.Inheritance.Gate533RequiresWickIEpsilon || !a.Inheritance.Gate533RequiresSourceConvention || !a.Inheritance.Gate533ComparatorBlocked || !a.Inheritance.Gate533OSWickHilbertBlocked || !a.Inheritance.Gate533NativeWriteBlocked || !a.Inheritance.Gate533NoObservedDataImported || !a.Inheritance.Gate534SyntheticRedirect {
		bad = append(bad, "bad Gate533 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 1 || a.Import.AcceptedRows != 1 || a.Import.RejectedRows != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedOSLoaded || a.Import.ObservedWickLoaded || a.Import.ObservedCorrelationLoaded || a.Import.ObservedHamiltonianLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.Gate533ReferenceComplete || !a.Import.Gate532ReferenceComplete || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsSynthetic || !a.Import.AllRowsNoTheoremInput || a.Import.AnyObservedClaim || a.Import.NativePromotionRejected {
		bad = append(bad, "bad synthetic OS ledger import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || a.Output.Dimension != 4 || a.Output.PositiveDomainDimension != 2 || !a.Output.ComparatorOnly || !a.Output.BridgeOnly || a.Output.NativePrediction || a.Output.ReflectionInvolutionResidual != 0 || a.Output.KernelSymmetryResidual != 0 || a.Output.ReflectionCovarianceResidual != 0 || a.Output.PositiveTimeDomainClosureResidual != 0 || a.Output.OSGramSymmetryResidual != 0 || !a.Output.OSGramPositiveDefinite || a.Output.OSGramPositiveEigenvalues != 2 || a.Output.OSGramNegativeEigenvalues != 0 || a.Output.OSGramZeroEigenvalues != 0 || a.Output.NegativeQuadraticVectors != 0 || !a.Output.AllSyntheticQuadraticsNonnegative || !a.Output.NullQuotientMetadataConsistent || !a.Output.Gate532ThetaCompatibilityDeclared || !a.Output.FiniteOSPlumbingVerified || !a.Output.SyntheticOSPositivityVerified || a.Output.PhysicalSchwingerFunctionsDerived || a.Output.PhysicalHilbertSpaceSelected || a.Output.WickRotationGranted || a.Output.PositiveEnergyHamiltonianDerived || a.Output.UnitaryRealTimeDynamicsDerived || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected {
		bad = append(bad, "bad synthetic OS adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedOSDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedCorrelationDataImported || a.Firewall.ObservedHamiltonianDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.NativeOSKernelWrite || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "bad firewall")
	}
	if len(bad) > 0 {
		return fmt.Errorf("Gate534 validation failed: %s", strings.Join(bad, "; "))
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

func clean(x float64) float64 {
	if math.Abs(x) < tolerance {
		return 0
	}
	return x
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: airlock=%t rows=%t reflection=%t domain=%t kernel=%t quadratic=%t null_quotient=%t theta_compat=%t Wick_iε=%t source_convention=%t comparator_blocked=%t OS_Wick_Hilbert_blocked=%t native_blocked=%t no_observed=%t gate534_redirect=%t; %s", x.Verdict, x.Gate533AirlockDefined, x.Gate533SchemaRowsEnumerated, x.Gate533RequiresReflection, x.Gate533RequiresTestDomain, x.Gate533RequiresKernel, x.Gate533RequiresQuadraticForm, x.Gate533RequiresNullQuotient, x.Gate533RequiresThetaCompatibility, x.Gate533RequiresWickIEpsilon, x.Gate533RequiresSourceConvention, x.Gate533ComparatorBlocked, x.Gate533OSWickHilbertBlocked, x.Gate533NativeWriteBlocked, x.Gate533NoObservedDataImported, x.Gate534SyntheticRedirect, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d bridge_only=%t synthetic=%t observed_OS=%t observed_Wick=%t observed_corr=%t observed_Hamiltonian=%t native_write=%t gate533_ref=%t gate532_ref=%t metadata=%t rows_bridge=%t rows_comparator=%t rows_synthetic=%t rows_no_theorem=%t observed_claim=%t native_rejected=%t path=%s; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.BridgeOnlyLedger, x.SyntheticFixture, x.ObservedOSLoaded, x.ObservedWickLoaded, x.ObservedCorrelationLoaded, x.ObservedHamiltonianLoaded, x.NativeRegistryWriteRequested, x.Gate533ReferenceComplete, x.Gate532ReferenceComplete, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsComparatorOnly, x.AllRowsSynthetic, x.AllRowsNoTheoremInput, x.AnyObservedClaim, x.NativePromotionRejected, x.Path, x.Reason)
}

func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("%s: ready=%t dim=%d positive_dim=%d comparator=%t bridge=%t native=%t R2-I=%g Ksym=%g RKR-K=%g domain_closure=%g OSGram_sym=%g eig_min=%g eig_max=%g eig_pos=%d eig_neg=%d eig_zero=%d OSGram_PD=%t q_min=%g q_max=%g nonzero_vectors=%d null_vectors=%d q_positive=%d q_zero=%d q_negative=%d quadratics_nonnegative=%t null_quotient=%t theta_compat=%t finite_plumbing=%t synthetic_OS=%t physical_Schwinger=%t physical_Hilbert=%t Wick=%t Hamiltonian=%t unitary=%t global=%t arrow=%t; %s", x.Verdict, x.Ready, x.Dimension, x.PositiveDomainDimension, x.ComparatorOnly, x.BridgeOnly, x.NativePrediction, x.ReflectionInvolutionResidual, x.KernelSymmetryResidual, x.ReflectionCovarianceResidual, x.PositiveTimeDomainClosureResidual, x.OSGramSymmetryResidual, x.OSGramEigenMin, x.OSGramEigenMax, x.OSGramPositiveEigenvalues, x.OSGramNegativeEigenvalues, x.OSGramZeroEigenvalues, x.OSGramPositiveDefinite, x.QuadraticMinimum, x.QuadraticMaximum, x.NonzeroTestVectors, x.NullTestVectors, x.PositiveQuadraticVectors, x.NullQuadraticVectors, x.NegativeQuadraticVectors, x.AllSyntheticQuadraticsNonnegative, x.NullQuotientMetadataConsistent, x.Gate532ThetaCompatibilityDeclared, x.FiniteOSPlumbingVerified, x.SyntheticOSPositivityVerified, x.PhysicalSchwingerFunctionsDerived, x.PhysicalHilbertSpaceSelected, x.WickRotationGranted, x.PositiveEnergyHamiltonianDerived, x.UnitaryRealTimeDynamicsDerived, x.GlobalHyperbolicityGranted, x.ArrowOfTimeSelected, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_OS=%t observed_Wick=%t observed_corr=%t observed_Hamiltonian=%t synthetic_only=%t file_native=%t adapter_native=%t native_OS=%t native_Schwinger=%t native_Hilbert=%t native_state=%t native_Wick=%t native_Hamiltonian=%t native_unitary=%t native_global=%t native_arrow=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_registry=%t; %s", x.Verdict, x.ObservedOSDataImported, x.ObservedWickDataImported, x.ObservedCorrelationDataImported, x.ObservedHamiltonianDataImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.NativeOSKernelWrite, x.NativeSchwingerFunctionWrite, x.NativeHilbertProductWrite, x.NativePhysicalStateSpaceWrite, x.NativeWickWrite, x.NativeHamiltonianWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{StatusGate533AirlockInherited, StatusSyntheticOSLedgerLoaded, StatusSyntheticOSRowAccepted, StatusSyntheticOSAdapterExecuted, StatusReflectionInvolutionResidualZero, StatusKernelSymmetryResidualZero, StatusReflectionCovarianceResidualZero, StatusPositiveTimeDomainClosed, StatusOSGramPositiveDefinite, StatusOSQuadraticFormPositive, StatusNullQuotientMetadataConsistent, StatusGate532ThetaCompatibilityDeclared, StatusNoObservedOSDataImportedDefault, StatusFailedSyntheticOSNotSchwinger, StatusFailedSyntheticOSNotWick, StatusFailedSyntheticOSNotHilbert, StatusFailedSyntheticOSNotHamiltonian, StatusFailedSyntheticOSNotUnitary, StatusFailedSyntheticOSNotGlobal, StatusFailedSyntheticOSNotArrow, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 534 Registry Audit — Synthetic OS Reflection-Positivity Kernel Adapter Dry Run\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 534 inherits Gate 533's OS kernel airlock and executes only a synthetic bridge-only comparator fixture.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic ledger import\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## OS quadratic-form dry run\n\nThe synthetic finite fixture evaluates `Q_OS(f)=<θ_E f, K f>` over the declared positive-time domain.\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
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
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
