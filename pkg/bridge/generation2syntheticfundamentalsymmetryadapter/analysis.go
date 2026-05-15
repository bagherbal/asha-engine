// Package generation2syntheticfundamentalsymmetryadapter implements Gate 532:
// Synthetic Fundamental-Symmetry Ledger Adapter and Positivity Residual Dry Run.
//
// Gate 531 defined the fail-closed Wick/Hilbert fundamental-symmetry airlock.
// This package performs the first synthetic, file-backed dry run through that
// airlock.  It evaluates only finite algebraic plumbing: Θ²=I, Krein
// self-adjointness, positivity of the matrix H=GΘ, and commutation with the
// accepted Gate 530 3+1 projector.  It deliberately does not infer Wick
// rotation, Osterwalder-Schrader reflection positivity, positive energy,
// unitary dynamics, global hyperbolicity, or a physical state-space theorem.
package generation2syntheticfundamentalsymmetryadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2synthetic3plus1projectionadapter"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2wickhilbertfundamentalsymmetryairlock"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID       = "GATE532-SYNTHETIC-FUNDAMENTAL-SYMMETRY-LEDGER-ADAPTER-POSITIVITY-DRY-RUN"
	DefaultLedger = "data/synthetic_fundamental_symmetry_ledger_gate532.json"

	StatusGate531AirlockInherited              = "CONDITIONAL_SUPPORT_GATE531_WICK_HILBERT_AIRLOCK_INHERITED"
	StatusSyntheticLedgerLoaded                = "CONDITIONAL_SUPPORT_SYNTHETIC_FUNDAMENTAL_SYMMETRY_LEDGER_LOADED"
	StatusSyntheticThetaRowAccepted            = "CONDITIONAL_SUPPORT_GATE532_AIRLOCK_ACCEPTED_SYNTHETIC_THETA_ROW"
	StatusThetaComparatorExecuted              = "CONDITIONAL_SUPPORT_SYNTHETIC_THETA_COMPARATOR_EXECUTED"
	StatusThetaInvolutionResidualZero          = "CONDITIONAL_SUPPORT_THETA_INVOLUTION_RESIDUAL_ZERO"
	StatusThetaKreinSelfAdjointResidualZero    = "CONDITIONAL_SUPPORT_THETA_KREIN_SELF_ADJOINT_RESIDUAL_ZERO"
	StatusGThetaPositiveDefinite               = "CONDITIONAL_SUPPORT_GTHETA_POSITIVE_DEFINITE_MATRIX_VERIFIED"
	StatusProjectorCompatibilityResidualZero   = "CONDITIONAL_SUPPORT_GATE530_PROJECTOR_COMPATIBILITY_RESIDUAL_ZERO"
	StatusFiniteKreinHilbertPlumbingVerified   = "CONDITIONAL_SUPPORT_FINITE_KREIN_TO_HILBERT_MATRIX_PLUMBING_VERIFIED"
	StatusNoObservedHilbertDataImportedDefault = "CONDITIONAL_SUPPORT_NO_OBSERVED_HILBERT_WICK_OR_BOUNDARY_DATA_IMPORTED_BY_DEFAULT"

	StatusFailedLedgerMissing              = "FAILED_ROUTE_GATE532_SYNTHETIC_THETA_LEDGER_MISSING"
	StatusFailedMetadataIncomplete         = "FAILED_ROUTE_GATE532_FUNDAMENTAL_SYMMETRY_METADATA_INCOMPLETE"
	StatusFailedInvalidMatrixDomain        = "FAILED_ROUTE_GATE532_INVALID_FUNDAMENTAL_SYMMETRY_MATRIX_DOMAIN"
	StatusFailedThetaInvolutionNonzero     = "FAILED_ROUTE_GATE532_THETA_INVOLUTION_RESIDUAL_NONZERO"
	StatusFailedThetaKreinAdjointNonzero   = "FAILED_ROUTE_GATE532_THETA_KREIN_SELF_ADJOINT_RESIDUAL_NONZERO"
	StatusFailedGThetaNotPositive          = "FAILED_ROUTE_GATE532_GTHETA_NOT_POSITIVE_DEFINITE"
	StatusFailedProjectorCompatibility     = "FAILED_ROUTE_GATE532_PROJECTOR_COMPATIBILITY_RESIDUAL_NONZERO"
	StatusFailedSyntheticThetaNative       = "FAILED_ROUTE_SYNTHETIC_THETA_NATIVE_PROMOTION_REJECTED"
	StatusFailedPositiveMatrixNotHilbert   = "FAILED_ROUTE_POSITIVE_GTHETA_MATRIX_DOES_NOT_SELECT_PHYSICAL_HILBERT_SPACE"
	StatusFailedPositiveMatrixNotWick      = "FAILED_ROUTE_POSITIVE_GTHETA_DOES_NOT_GRANT_WICK_ROTATION"
	StatusFailedPositiveMatrixNotOS        = "FAILED_ROUTE_POSITIVE_GTHETA_DOES_NOT_GRANT_OS_REFLECTION_POSITIVITY"
	StatusFailedPositiveMatrixNotEnergy    = "FAILED_ROUTE_POSITIVE_GTHETA_DOES_NOT_GRANT_POSITIVE_ENERGY_HAMILTONIAN"
	StatusFailedPositiveMatrixNotUnitary   = "FAILED_ROUTE_POSITIVE_GTHETA_DOES_NOT_GRANT_UNITARY_REAL_TIME_DYNAMICS"
	StatusFailedPositiveMatrixNotGlobal    = "FAILED_ROUTE_POSITIVE_GTHETA_DOES_NOT_GRANT_GLOBAL_HYPERBOLICITY"
	StatusFailedSyntheticThetaNotTimeArrow = "FAILED_ROUTE_SYNTHETIC_THETA_DOES_NOT_SELECT_ARROW_OF_TIME"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_GATE532_SYNTHETIC_FUNDAMENTAL_SYMMETRY_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked       = "FIREWALL_BLOCKED_GATE532_HILBERT_WICK_NATIVE_WRITE"
)

const tolerance = 1e-12

type Inheritance struct {
	Executed bool

	Gate531AirlockDefined           bool
	Gate531SchemaRowsEnumerated     bool
	Gate531RequiresKreinMetric      bool
	Gate531RequiresTheta            bool
	Gate531RequiresProjectorCompat  bool
	Gate531RequiresSourceConvention bool
	Gate531RequiresBridgeOnly       bool
	Gate531RejectsNativePromotion   bool
	Gate531ComparatorBlocked        bool
	Gate531HilbertWickOSBlocked     bool
	Gate531NoObservedDataImported   bool
	Gate531NativeWriteBlocked       bool
	Gate532SyntheticRedirect        bool

	Verdict, Reason string
}

type FundamentalSymmetryRow struct {
	Name                            string      `json:"name"`
	KreinMetricMatrix               [][]float64 `json:"krein_metric_matrix"`
	FundamentalSymmetryMatrix       [][]float64 `json:"fundamental_symmetry_matrix"`
	TimeReflectionOperator          [][]float64 `json:"time_reflection_operator"`
	WickMapConvention               string      `json:"wick_map_convention"`
	IepsilonPrescription            string      `json:"i_epsilon_prescription"`
	ReflectionPositivityCertificate string      `json:"reflection_positivity_certificate"`
	PositiveEnergyCertificate       string      `json:"positive_energy_certificate"`
	GlobalCausalBoundaryData        string      `json:"global_causal_boundary_data"`
	Source                          string      `json:"source"`
	SourceVersion                   string      `json:"source_version"`
	Convention                      string      `json:"convention"`
	Uncertainty                     string      `json:"uncertainty"`
	BridgeOnly                      bool        `json:"bridge_only"`
	ComparatorOnly                  bool        `json:"comparator_only"`
	MatrixPositivityOnly            bool        `json:"matrix_positivity_only"`
	NoTheoremInput                  bool        `json:"no_theorem_input"`
	Synthetic                       bool        `json:"synthetic"`
	Observed                        bool        `json:"observed"`
	NativePromotion                 bool        `json:"native_promotion"`
	NativeInputClaim                bool        `json:"native_input_claim"`
}

type FundamentalSymmetryLedger struct {
	Gate                   int                      `json:"gate"`
	LedgerName             string                   `json:"ledger_name"`
	Description            string                   `json:"description"`
	ProjectorReference     string                   `json:"projector_reference"`
	BridgeOnly             bool                     `json:"bridge_only"`
	NativeRegistryWrite    bool                     `json:"native_registry_write"`
	SyntheticFixture       bool                     `json:"synthetic_fixture"`
	ObservedHilbertLoaded  bool                     `json:"observed_hilbert_loaded"`
	ObservedWickLoaded     bool                     `json:"observed_wick_loaded"`
	ObservedBoundaryLoaded bool                     `json:"observed_boundary_loaded"`
	Source                 string                   `json:"source"`
	SourceVersion          string                   `json:"source_version"`
	Convention             string                   `json:"convention"`
	Rows                   []FundamentalSymmetryRow `json:"rows"`
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
	ObservedHilbertLoaded        bool
	ObservedWickLoaded           bool
	ObservedBoundaryLoaded       bool
	NativeRegistryWriteRequested bool
	ProjectorReferenceComplete   bool
	MetadataComplete             bool
	AllRowsBridgeOnly            bool
	AllRowsComparatorOnly        bool
	AllRowsMatrixPositivityOnly  bool
	AllRowsNoTheoremInput        bool
	AllRowsSynthetic             bool
	AnyObservedClaim             bool
	NativePromotionRejected      bool
	Verdict, Reason              string
	Failures                     []string
}

type AdapterInput struct {
	Rows                            int
	G                               linear.Matrix
	Theta                           linear.Matrix
	TimeReflection                  linear.Matrix
	Projector                       linear.Matrix
	ProjectorReference              string
	WickMapConvention               string
	IepsilonPrescription            string
	ReflectionPositivityCertificate string
	PositiveEnergyCertificate       string
	GlobalCausalBoundaryData        string
	BridgeOnly                      bool
	SyntheticFixture                bool
	ObservedHilbertLoaded           bool
	ObservedWickLoaded              bool
	ObservedBoundaryLoaded          bool
	NativePromotion                 bool
	MetadataComplete                bool
}

type AdapterOutput struct {
	Executed  bool
	Attempted bool
	Ready     bool

	Dimension                        int
	ComparatorOnly                   bool
	BridgeOnly                       bool
	NativePrediction                 bool
	ThetaTrace                       float64
	ThetaSquaredIdentityResidual     float64
	ThetaKreinSelfAdjointResidual    float64
	GThetaSymmetryResidual           float64
	GThetaEigenMin                   float64
	GThetaEigenMax                   float64
	GThetaPositiveEigenvalues        int
	GThetaNegativeEigenvalues        int
	GThetaZeroEigenvalues            int
	GThetaPositiveDefinite           bool
	ProjectorCompatibilityResidual   float64
	TimeReflectionInvolutionResidual float64
	FiniteMatrixPlumbingVerified     bool
	PositiveHilbertMatrixVerified    bool
	PhysicalHilbertSpaceGranted      bool
	WickRotationGranted              bool
	ReflectionPositivityGranted      bool
	PositiveEnergyGranted            bool
	UnitaryRealTimeGranted           bool
	GlobalHyperbolicityGranted       bool
	ArrowOfTimeSelected              bool
	Verdict, Reason                  string
	Failures                         []string
}

type Firewall struct {
	Executed bool

	ObservedHilbertDataImported    bool
	ObservedWickDataImported       bool
	ObservedBoundaryDataImported   bool
	SyntheticFixtureOnly           bool
	FileRowsNative                 bool
	AdapterOutputsNative           bool
	NativeFundamentalSymmetryWrite bool
	NativeHilbertProductWrite      bool
	NativePhysicalStateSpaceWrite  bool
	NativeWickWrite                bool
	NativeReflectionWrite          bool
	NativePositiveEnergyWrite      bool
	NativeUnitaryDynamicsWrite     bool
	NativeGlobalCausalWrite        bool
	NativeTimeArrowWrite           bool
	ReopenedFlavorFirewall         bool
	ReopenedEWScaleFirewall        bool
	ReopenedGravityFirewall        bool
	ReopenedTopologyFirewall       bool
	NativeRegistryWritten          bool
	Verdict, Reason                string
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
	a := Analysis{Inheritance: buildInheritance()}
	ledger, imp := loadLedger(path)
	a.Import = imp
	if imp.Loaded && imp.MetadataComplete && imp.AcceptedRows == 1 {
		in, err := buildInput(ledger, imp)
		if err != nil {
			a.Input = AdapterInput{Rows: len(ledger.Rows), MetadataComplete: imp.MetadataComplete, BridgeOnly: ledger.BridgeOnly, SyntheticFixture: ledger.SyntheticFixture, ObservedHilbertLoaded: ledger.ObservedHilbertLoaded, ObservedWickLoaded: ledger.ObservedWickLoaded, ObservedBoundaryLoaded: ledger.ObservedBoundaryLoaded, NativePromotion: ledger.NativeRegistryWrite}
			a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, ComparatorOnly: true, NativePrediction: false, Verdict: StatusFailedInvalidMatrixDomain, Reason: err.Error(), Failures: []string{StatusFailedInvalidMatrixDomain}}
		} else {
			a.Input = in
			a.Output = runAdapter(in)
		}
	} else if !imp.Loaded {
		a.Output = AdapterOutput{Executed: true, Attempted: false, BridgeOnly: true, ComparatorOnly: true, NativePrediction: false, Verdict: StatusFailedLedgerMissing, Reason: "explicit Gate532 synthetic fundamental-symmetry ledger was not found", Failures: []string{StatusFailedLedgerMissing}}
	} else {
		a.Output = AdapterOutput{Executed: true, Attempted: true, BridgeOnly: true, ComparatorOnly: true, NativePrediction: false, Verdict: StatusFailedMetadataIncomplete, Reason: "Gate532 fundamental-symmetry file did not satisfy the Gate531 airlock metadata domain", Failures: []string{StatusFailedMetadataIncomplete}}
	}
	a.Firewall = buildFirewall(a.Import, a.Output)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	g531, err := generation2wickhilbertfundamentalsymmetryairlock.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedMetadataIncomplete, Reason: fmt.Sprintf("could not inherit Gate531 Wick/Hilbert airlock: %v", err)}
	}
	return Inheritance{
		Executed:                        true,
		Gate531AirlockDefined:           g531.Schema.Executed && g531.Schema.KreinMetricMatrixRequired && g531.Schema.FundamentalSymmetryMatrixRequired,
		Gate531SchemaRowsEnumerated:     g531.Schema.RequiredRowCount >= 15,
		Gate531RequiresKreinMetric:      g531.Schema.KreinMetricMatrixRequired,
		Gate531RequiresTheta:            g531.Schema.FundamentalSymmetryMatrixRequired && g531.Schema.ThetaInvolutionCheckRequired && g531.Schema.ThetaKreinSelfAdjointCheckRequired,
		Gate531RequiresProjectorCompat:  g531.Schema.ProjectorCompatibilityCheckRequired,
		Gate531RequiresSourceConvention: g531.Schema.SourceRequired && g531.Schema.ConventionRequired,
		Gate531RequiresBridgeOnly:       g531.Schema.BridgeOnlyRequired && g531.Schema.NoTheoremInputRequired,
		Gate531RejectsNativePromotion:   g531.Schema.NativePromotionRejected && !g531.Rejection.NativeFundamentalSymmetryWrite,
		Gate531ComparatorBlocked:        !g531.Guard.ComparatorExecutionPerformed,
		Gate531HilbertWickOSBlocked:     !g531.Guard.PositiveHilbertProductGranted && !g531.Guard.WickRotationSelected && !g531.Guard.ReflectionPositivityProven && !g531.Guard.PositiveEnergyHamiltonianDerived && !g531.Guard.UnitaryRealTimeDynamicsDerived && !g531.Guard.GlobalHyperbolicitySelected,
		Gate531NoObservedDataImported:   !g531.Firewall.ObservedHilbertDataImported && !g531.Firewall.ObservedWickDataImported && !g531.Firewall.ObservedBoundaryDataImported && !g531.Firewall.ObservedHamiltonianDataImported,
		Gate531NativeWriteBlocked:       !g531.Firewall.NativeRegistryWritten && !g531.Firewall.NativeFundamentalSymmetryWrite && !g531.Firewall.NativeHilbertProductWrite,
		Gate532SyntheticRedirect:        g531.Next.Gate == 532,
		Verdict:                         StatusGate531AirlockInherited,
		Reason:                          "Gate532 inherits Gate531's fail-closed Wick/Hilbert airlock: a Θ row may be checked only as sourced, synthetic, bridge-only finite matrix plumbing, without native state-space, Wick, OS, positive-energy, unitary, or global-causal promotion.",
	}
}

func loadLedger(path string) (FundamentalSymmetryLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved}
	b, err := os.ReadFile(resolved)
	if err != nil {
		imp.Verdict = StatusFailedLedgerMissing
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedLedgerMissing}
		return FundamentalSymmetryLedger{}, imp
	}
	var ledger FundamentalSymmetryLedger
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
	imp.ObservedHilbertLoaded = ledger.ObservedHilbertLoaded
	imp.ObservedWickLoaded = ledger.ObservedWickLoaded
	imp.ObservedBoundaryLoaded = ledger.ObservedBoundaryLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.ProjectorReferenceComplete = ledger.ProjectorReference != ""
	imp.NativePromotionRejected = !ledger.NativeRegistryWrite
	metadata := ledger.Gate == 532 && ledger.LedgerName != "" && ledger.ProjectorReference != "" && ledger.Source != "" && ledger.SourceVersion != "" && ledger.Convention != "" && len(ledger.Rows) == 1
	allBridge := ledger.BridgeOnly
	allComparator := true
	allMatrixOnly := true
	allNoTheorem := true
	allSynthetic := ledger.SyntheticFixture
	anyObserved := ledger.ObservedHilbertLoaded || ledger.ObservedWickLoaded || ledger.ObservedBoundaryLoaded
	accepted := 0
	failures := []string{}
	for _, r := range ledger.Rows {
		rowMeta := r.Name != "" && r.Source != "" && r.SourceVersion != "" && r.Convention != "" && r.Uncertainty != "" && r.WickMapConvention != "" && r.IepsilonPrescription != "" && r.ReflectionPositivityCertificate != "" && r.PositiveEnergyCertificate != "" && r.GlobalCausalBoundaryData != ""
		shapeOK := matrixShapeOK(r.KreinMetricMatrix, 8, 8) && matrixShapeOK(r.FundamentalSymmetryMatrix, 8, 8) && matrixShapeOK(r.TimeReflectionOperator, 8, 8)
		if !rowMeta {
			metadata = false
		}
		if !shapeOK {
			failures = append(failures, StatusFailedInvalidMatrixDomain)
		}
		if !r.BridgeOnly {
			allBridge = false
		}
		if !r.ComparatorOnly {
			allComparator = false
		}
		if !r.MatrixPositivityOnly {
			allMatrixOnly = false
		}
		if !r.NoTheoremInput {
			allNoTheorem = false
		}
		if !r.Synthetic {
			allSynthetic = false
		}
		if r.Observed {
			anyObserved = true
		}
		if r.NativePromotion || r.NativeInputClaim {
			failures = append(failures, StatusFailedSyntheticThetaNative)
		}
		if rowMeta && shapeOK && r.BridgeOnly && r.ComparatorOnly && r.MatrixPositivityOnly && r.NoTheoremInput && r.Synthetic && !r.Observed && !r.NativePromotion && !r.NativeInputClaim {
			accepted++
		}
	}
	if !metadata {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if !ledger.BridgeOnly || ledger.NativeRegistryWrite {
		failures = append(failures, StatusFailedSyntheticThetaNative)
	}
	if anyObserved {
		failures = append(failures, StatusFailedPositiveMatrixNotHilbert)
	}
	imp.AcceptedRows = accepted
	imp.RejectedRows = imp.Rows - accepted
	imp.MetadataComplete = metadata && accepted == 1 && len(failures) == 0
	imp.AllRowsBridgeOnly = allBridge
	imp.AllRowsComparatorOnly = allComparator
	imp.AllRowsMatrixPositivityOnly = allMatrixOnly
	imp.AllRowsNoTheoremInput = allNoTheorem
	imp.AllRowsSynthetic = allSynthetic
	imp.AnyObservedClaim = anyObserved
	imp.Failures = unique(failures)
	if len(imp.Failures) == 0 {
		imp.Verdict = strings.Join([]string{StatusSyntheticLedgerLoaded, StatusSyntheticThetaRowAccepted, StatusNoObservedHilbertDataImportedDefault}, ";")
		imp.Reason = "Gate532 loaded a deliberately synthetic, source-tagged fundamental-symmetry row through the Gate531 airlock with bridge_only=true, matrix_positivity_only=true, no_theorem_input=true, and native_promotion=false."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Gate532 file import failed one or more Gate531 fundamental-symmetry airlock requirements."
	}
	return ledger, imp
}

func buildInput(ledger FundamentalSymmetryLedger, imp FileImport) (AdapterInput, error) {
	row := ledger.Rows[0]
	g, err := linear.FromRows(row.KreinMetricMatrix)
	if err != nil {
		return AdapterInput{}, err
	}
	theta, err := linear.FromRows(row.FundamentalSymmetryMatrix)
	if err != nil {
		return AdapterInput{}, err
	}
	timeReflection, err := linear.FromRows(row.TimeReflectionOperator)
	if err != nil {
		return AdapterInput{}, err
	}
	g530, err := generation2synthetic3plus1projectionadapter.BuildDefault()
	if err != nil {
		return AdapterInput{}, fmt.Errorf("could not inherit Gate530 projector: %w", err)
	}
	return AdapterInput{
		Rows: len(ledger.Rows), G: g, Theta: theta, TimeReflection: timeReflection, Projector: g530.Input.P,
		ProjectorReference: ledger.ProjectorReference,
		WickMapConvention:  row.WickMapConvention, IepsilonPrescription: row.IepsilonPrescription,
		ReflectionPositivityCertificate: row.ReflectionPositivityCertificate, PositiveEnergyCertificate: row.PositiveEnergyCertificate, GlobalCausalBoundaryData: row.GlobalCausalBoundaryData,
		BridgeOnly:             ledger.BridgeOnly && row.BridgeOnly,
		SyntheticFixture:       ledger.SyntheticFixture && row.Synthetic,
		ObservedHilbertLoaded:  ledger.ObservedHilbertLoaded || row.Observed,
		ObservedWickLoaded:     ledger.ObservedWickLoaded,
		ObservedBoundaryLoaded: ledger.ObservedBoundaryLoaded,
		NativePromotion:        ledger.NativeRegistryWrite || row.NativePromotion || row.NativeInputClaim,
		MetadataComplete:       imp.MetadataComplete,
	}, nil
}

func runAdapter(in AdapterInput) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: true, Dimension: 8, BridgeOnly: true, ComparatorOnly: true, NativePrediction: false}
	ready := in.MetadataComplete && in.BridgeOnly && in.SyntheticFixture && !in.ObservedHilbertLoaded && !in.ObservedWickLoaded && !in.ObservedBoundaryLoaded && !in.NativePromotion && in.G.Rows() == 8 && in.G.Cols() == 8 && in.Theta.Rows() == 8 && in.Theta.Cols() == 8 && in.Projector.Rows() == 8 && in.Projector.Cols() == 8
	out.Ready = ready
	if !ready {
		out.Verdict = StatusFailedMetadataIncomplete
		out.Reason = "Synthetic fundamental-symmetry adapter input did not meet the complete bridge-only domain."
		out.Failures = []string{StatusFailedMetadataIncomplete}
		return out
	}
	theta2, _ := in.Theta.Mul(in.Theta)
	theta2MinusI, _ := theta2.Sub(linear.Identity(8))
	out.ThetaSquaredIdentityResidual = theta2MinusI.FrobeniusNorm()
	out.ThetaTrace, _ = in.Theta.Trace()

	// Krein self-adjointness for involutive diagonal G: Θ†_G=G Θᵀ G.  Residual is ΘᵀG - GΘ, equivalent for G²=I.
	thetaTG, _ := in.Theta.Transpose().Mul(in.G)
	gTheta, _ := in.G.Mul(in.Theta)
	kreinAdjointResidual, _ := thetaTG.Sub(gTheta)
	out.ThetaKreinSelfAdjointResidual = kreinAdjointResidual.FrobeniusNorm()

	gThetaT := gTheta.Transpose()
	gThetaSymResidual, _ := gTheta.Sub(gThetaT)
	out.GThetaSymmetryResidual = gThetaSymResidual.FrobeniusNorm()
	if gTheta.IsSymmetric(100 * tolerance) {
		eig, err := linear.SymmetricEigenJacobi(gTheta, tolerance, 200)
		if err == nil {
			out.GThetaEigenMin, out.GThetaEigenMax = eig.Values[0], eig.Values[0]
			for _, v := range eig.Values {
				if v < out.GThetaEigenMin {
					out.GThetaEigenMin = v
				}
				if v > out.GThetaEigenMax {
					out.GThetaEigenMax = v
				}
				switch {
				case v > tolerance:
					out.GThetaPositiveEigenvalues++
				case v < -tolerance:
					out.GThetaNegativeEigenvalues++
				default:
					out.GThetaZeroEigenvalues++
				}
			}
		}
	}
	out.GThetaPositiveDefinite = out.GThetaPositiveEigenvalues == 8 && out.GThetaNegativeEigenvalues == 0 && out.GThetaZeroEigenvalues == 0 && out.GThetaEigenMin > tolerance
	comm, _ := linear.Commutator(in.Theta, in.Projector)
	out.ProjectorCompatibilityResidual = comm.FrobeniusNorm()
	r2, _ := in.TimeReflection.Mul(in.TimeReflection)
	r2MinusI, _ := r2.Sub(linear.Identity(8))
	out.TimeReflectionInvolutionResidual = r2MinusI.FrobeniusNorm()

	out.FiniteMatrixPlumbingVerified = nearly(out.ThetaSquaredIdentityResidual, 0, tolerance) && nearly(out.ThetaKreinSelfAdjointResidual, 0, tolerance) && nearly(out.GThetaSymmetryResidual, 0, tolerance) && out.GThetaPositiveDefinite && nearly(out.ProjectorCompatibilityResidual, 0, tolerance) && nearly(out.TimeReflectionInvolutionResidual, 0, tolerance)
	out.PositiveHilbertMatrixVerified = out.FiniteMatrixPlumbingVerified
	out.PhysicalHilbertSpaceGranted = false
	out.WickRotationGranted = false
	out.ReflectionPositivityGranted = false
	out.PositiveEnergyGranted = false
	out.UnitaryRealTimeGranted = false
	out.GlobalHyperbolicityGranted = false
	out.ArrowOfTimeSelected = false

	failures := []string{}
	if !nearly(out.ThetaSquaredIdentityResidual, 0, tolerance) {
		failures = append(failures, StatusFailedThetaInvolutionNonzero)
	}
	if !nearly(out.ThetaKreinSelfAdjointResidual, 0, tolerance) || !nearly(out.GThetaSymmetryResidual, 0, tolerance) {
		failures = append(failures, StatusFailedThetaKreinAdjointNonzero)
	}
	if !out.GThetaPositiveDefinite {
		failures = append(failures, StatusFailedGThetaNotPositive)
	}
	if !nearly(out.ProjectorCompatibilityResidual, 0, tolerance) {
		failures = append(failures, StatusFailedProjectorCompatibility)
	}
	out.Failures = unique(failures)
	if len(out.Failures) == 0 {
		out.Verdict = strings.Join([]string{StatusThetaComparatorExecuted, StatusThetaInvolutionResidualZero, StatusThetaKreinSelfAdjointResidualZero, StatusGThetaPositiveDefinite, StatusProjectorCompatibilityResidualZero, StatusFiniteKreinHilbertPlumbingVerified}, ";")
		out.Reason = "The synthetic Θ fixture passes finite bridge-only algebra: Θ²=I, ΘᵀG=GΘ, H=GΘ is symmetric positive definite, [Θ,P_530]=0, and the time-reflection operator squares to identity. This verifies matrix plumbing only."
	} else {
		out.Verdict = strings.Join(out.Failures, ";")
		out.Reason = "Gate532 adapter found nonzero Θ/Krein/projector residuals or a non-positive H=GΘ matrix."
	}
	return out
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{
		Executed:                       true,
		ObservedHilbertDataImported:    imp.ObservedHilbertLoaded || imp.AnyObservedClaim,
		ObservedWickDataImported:       imp.ObservedWickLoaded,
		ObservedBoundaryDataImported:   imp.ObservedBoundaryLoaded,
		SyntheticFixtureOnly:           imp.SyntheticFixture && !imp.ObservedHilbertLoaded && !imp.ObservedWickLoaded && !imp.ObservedBoundaryLoaded && !imp.AnyObservedClaim,
		FileRowsNative:                 false,
		AdapterOutputsNative:           out.NativePrediction,
		NativeFundamentalSymmetryWrite: false,
		NativeHilbertProductWrite:      false,
		NativePhysicalStateSpaceWrite:  out.PhysicalHilbertSpaceGranted,
		NativeWickWrite:                out.WickRotationGranted,
		NativeReflectionWrite:          out.ReflectionPositivityGranted,
		NativePositiveEnergyWrite:      out.PositiveEnergyGranted,
		NativeUnitaryDynamicsWrite:     out.UnitaryRealTimeGranted,
		NativeGlobalCausalWrite:        out.GlobalHyperbolicityGranted,
		NativeTimeArrowWrite:           out.ArrowOfTimeSelected,
		ReopenedFlavorFirewall:         false,
		ReopenedEWScaleFirewall:        false,
		ReopenedGravityFirewall:        false,
		ReopenedTopologyFirewall:       false,
		NativeRegistryWritten:          imp.NativeRegistryWriteRequested,
		Verdict:                        strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedSyntheticThetaNative, StatusFailedPositiveMatrixNotHilbert, StatusFailedPositiveMatrixNotWick, StatusFailedPositiveMatrixNotOS, StatusFailedPositiveMatrixNotEnergy, StatusFailedPositiveMatrixNotUnitary, StatusFailedPositiveMatrixNotGlobal, StatusFailedSyntheticThetaNotTimeArrow}, ";"),
		Reason:                         "Gate532 confirms only a synthetic finite positive matrix form H=GΘ. The result is bridge plumbing and does not become a native fundamental symmetry, physical Hilbert space, Wick rotation, OS reconstruction, positive-energy Hamiltonian, unitary dynamics, global causal structure, or arrow of time.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native fundamental symmetry Θ is written at Gate532.",
			"Cℓ(1,7) still contributes only the native indefinite/Krein causal socket; the positive H=GΘ form is a synthetic bridge comparator output.",
			"No physical Hilbert state space, Wick map, reflection-positive Euclidean theory, Hamiltonian spectrum, unitary dynamics, global hyperbolicity, or time arrow is promoted natively.",
		},
		BridgeEntries: []string{
			"File-backed synthetic fundamental-symmetry adapter implemented against the Gate531 Wick/Hilbert airlock.",
			"Bridge residuals verify Θ²−I=0, ΘᵀG−GΘ=0, H=GΘ positive-definite, [Θ,P_530]=0, and R_time²−I=0 for the default synthetic fixture.",
			"The finite positive matrix can be used as a safe dry-run for the Krein-to-Hilbert socket, but only with bridge_only=true, matrix_positivity_only=true, no_theorem_input=true, and native_promotion=false.",
		},
		EnvironmentalEntries: []string{
			"The actual physical fundamental symmetry, time reflection, Wick/iε dictionary, thermodynamic arrow, Hamiltonian domain, and global causal boundary remain environmental or future bridge data.",
			"Finite matrix positivity is necessary plumbing for Hilbert reconstruction but is not sufficient for OS reflection positivity or Lorentzian quantum dynamics.",
		},
		FailedRoutes: []string{
			StatusFailedSyntheticThetaNative,
			StatusFailedPositiveMatrixNotHilbert,
			StatusFailedPositiveMatrixNotWick,
			StatusFailedPositiveMatrixNotOS,
			StatusFailedPositiveMatrixNotEnergy,
			StatusFailedPositiveMatrixNotUnitary,
			StatusFailedPositiveMatrixNotGlobal,
			StatusFailedSyntheticThetaNotTimeArrow,
		},
		OpenTheorems: []string{
			"define an Osterwalder-Schrader/reflection-positivity kernel airlock instead of inferring it from H=GΘ positivity",
			"audit whether a native algebraic mechanism can select a non-synthetic fundamental symmetry Θ rather than importing it by bridge ledger",
			"separate time-reflection involution, thermodynamic time arrow, positive-energy spectrum, and global hyperbolicity as independent gates",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        533,
		Title:       "Osterwalder-Schrader Reflection-Positivity Kernel Airlock Preflight",
		Reason:      "Gate532 verifies that a synthetic Θ can turn the imported Krein matrix into a positive finite form, but positive H=GΘ still does not prove reflection positivity or Wick reconstruction. The next safe step is to define the OS kernel/correlation-data airlock.",
		PrimaryTask: "Define source-tagged bridge requirements for a Euclidean reflection operator, test-function domain, two-point/kernel data, positivity cone, and reconstruction certificate without promoting any Wick rotation or physical Hilbert space natively.",
	}
}

func truth(a Analysis) string {
	return "Gate532 confirms that the Wick/Hilbert socket can carry a synthetic fundamental symmetry through finite algebraic tests: Θ²=I, Krein self-adjointness, positive H=GΘ, Gate530 projector compatibility, and time-reflection involution all close for the default fixture. This is a successful plumbing check, not a physical-universe selection. The positive matrix does not grant Wick rotation, OS reflection positivity, positive energy, unitary real-time dynamics, global hyperbolicity, or the arrow of time."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate531AirlockDefined || !a.Inheritance.Gate531SchemaRowsEnumerated || !a.Inheritance.Gate531RequiresKreinMetric || !a.Inheritance.Gate531RequiresTheta || !a.Inheritance.Gate531RequiresProjectorCompat || !a.Inheritance.Gate531RequiresSourceConvention || !a.Inheritance.Gate531RequiresBridgeOnly || !a.Inheritance.Gate531RejectsNativePromotion || !a.Inheritance.Gate531ComparatorBlocked || !a.Inheritance.Gate531HilbertWickOSBlocked || !a.Inheritance.Gate531NoObservedDataImported || !a.Inheritance.Gate531NativeWriteBlocked || !a.Inheritance.Gate532SyntheticRedirect {
		bad = append(bad, "bad Gate531 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 1 || a.Import.AcceptedRows != 1 || a.Import.RejectedRows != 0 || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedHilbertLoaded || a.Import.ObservedWickLoaded || a.Import.ObservedBoundaryLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.ProjectorReferenceComplete || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsMatrixPositivityOnly || !a.Import.AllRowsNoTheoremInput || !a.Import.AllRowsSynthetic || a.Import.AnyObservedClaim {
		bad = append(bad, "bad file import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || !a.Output.ComparatorOnly || !a.Output.BridgeOnly || a.Output.NativePrediction || !a.Output.FiniteMatrixPlumbingVerified || !a.Output.PositiveHilbertMatrixVerified || a.Output.PhysicalHilbertSpaceGranted || a.Output.WickRotationGranted || a.Output.ReflectionPositivityGranted || a.Output.PositiveEnergyGranted || a.Output.UnitaryRealTimeGranted || a.Output.GlobalHyperbolicityGranted || a.Output.ArrowOfTimeSelected || !nearly(a.Output.ThetaSquaredIdentityResidual, 0, tolerance) || !nearly(a.Output.ThetaKreinSelfAdjointResidual, 0, tolerance) || !nearly(a.Output.GThetaSymmetryResidual, 0, tolerance) || !a.Output.GThetaPositiveDefinite || a.Output.GThetaPositiveEigenvalues != 8 || a.Output.GThetaNegativeEigenvalues != 0 || a.Output.GThetaZeroEigenvalues != 0 || !nearly(a.Output.ProjectorCompatibilityResidual, 0, tolerance) || !nearly(a.Output.TimeReflectionInvolutionResidual, 0, tolerance) {
		bad = append(bad, "bad adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedHilbertDataImported || a.Firewall.ObservedWickDataImported || a.Firewall.ObservedBoundaryDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.NativeFundamentalSymmetryWrite || a.Firewall.NativeHilbertProductWrite || a.Firewall.NativePhysicalStateSpaceWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeReflectionWrite || a.Firewall.NativePositiveEnergyWrite || a.Firewall.NativeUnitaryDynamicsWrite || a.Firewall.NativeGlobalCausalWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.ReopenedFlavorFirewall || a.Firewall.ReopenedEWScaleFirewall || a.Firewall.ReopenedGravityFirewall || a.Firewall.ReopenedTopologyFirewall || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: airlock=%t schema_rows=%t G_required=%t theta_required=%t projector_compat=%t source_convention=%t bridge_only=%t native_rejected=%t comparator_blocked=%t hilbert_wick_os_blocked=%t no_observed=%t native_blocked=%t gate532_redirect=%t; %s", x.Verdict, x.Gate531AirlockDefined, x.Gate531SchemaRowsEnumerated, x.Gate531RequiresKreinMetric, x.Gate531RequiresTheta, x.Gate531RequiresProjectorCompat, x.Gate531RequiresSourceConvention, x.Gate531RequiresBridgeOnly, x.Gate531RejectsNativePromotion, x.Gate531ComparatorBlocked, x.Gate531HilbertWickOSBlocked, x.Gate531NoObservedDataImported, x.Gate531NativeWriteBlocked, x.Gate532SyntheticRedirect, x.Reason)
}

func FormatImport(x FileImport) string {
	return fmt.Sprintf("%s: loaded=%t rows=%d accepted=%d rejected=%d bridge_only=%t synthetic_fixture=%t observed_Hilbert=%t observed_Wick=%t observed_boundary=%t native_write_requested=%t projector_ref=%t metadata_complete=%t all_bridge_only=%t all_comparator_only=%t matrix_only=%t no_theorem_input=%t all_synthetic=%t observed_claim=%t native_promotion_rejected=%t; %s", x.Verdict, x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.BridgeOnlyLedger, x.SyntheticFixture, x.ObservedHilbertLoaded, x.ObservedWickLoaded, x.ObservedBoundaryLoaded, x.NativeRegistryWriteRequested, x.ProjectorReferenceComplete, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsComparatorOnly, x.AllRowsMatrixPositivityOnly, x.AllRowsNoTheoremInput, x.AllRowsSynthetic, x.AnyObservedClaim, x.NativePromotionRejected, x.Reason)
}

func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("%s: ready=%t dim=%d theta_trace=%.12g theta2_minus_I=%.12g thetaT_G_minus_G_theta=%.12g Gtheta_sym=%.12g eig_min=%.12g eig_max=%.12g eig_pos=%d eig_neg=%d eig_zero=%d Gtheta_positive=%t comm_theta_P530=%.12g Rtime2_minus_I=%.12g finite_plumbing=%t positive_matrix=%t physical_Hilbert=%t Wick=%t OS=%t positive_energy=%t unitary=%t global=%t arrow=%t bridge_only=%t native_prediction=%t; %s", x.Verdict, x.Ready, x.Dimension, x.ThetaTrace, x.ThetaSquaredIdentityResidual, x.ThetaKreinSelfAdjointResidual, x.GThetaSymmetryResidual, x.GThetaEigenMin, x.GThetaEigenMax, x.GThetaPositiveEigenvalues, x.GThetaNegativeEigenvalues, x.GThetaZeroEigenvalues, x.GThetaPositiveDefinite, x.ProjectorCompatibilityResidual, x.TimeReflectionInvolutionResidual, x.FiniteMatrixPlumbingVerified, x.PositiveHilbertMatrixVerified, x.PhysicalHilbertSpaceGranted, x.WickRotationGranted, x.ReflectionPositivityGranted, x.PositiveEnergyGranted, x.UnitaryRealTimeGranted, x.GlobalHyperbolicityGranted, x.ArrowOfTimeSelected, x.BridgeOnly, x.NativePrediction, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_Hilbert=%t observed_Wick=%t observed_boundary=%t synthetic_only=%t file_rows_native=%t outputs_native=%t native_theta=%t native_Hilbert=%t native_state=%t native_Wick=%t native_reflection=%t native_positive_energy=%t native_unitary=%t native_global=%t native_arrow=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t native_registry_written=%t; %s", x.Verdict, x.ObservedHilbertDataImported, x.ObservedWickDataImported, x.ObservedBoundaryDataImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.NativeFundamentalSymmetryWrite, x.NativeHilbertProductWrite, x.NativePhysicalStateSpaceWrite, x.NativeWickWrite, x.NativeReflectionWrite, x.NativePositiveEnergyWrite, x.NativeUnitaryDynamicsWrite, x.NativeGlobalCausalWrite, x.NativeTimeArrowWrite, x.ReopenedFlavorFirewall, x.ReopenedEWScaleFirewall, x.ReopenedGravityFirewall, x.ReopenedTopologyFirewall, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate531AirlockInherited,
		StatusSyntheticLedgerLoaded,
		StatusSyntheticThetaRowAccepted,
		StatusThetaComparatorExecuted,
		StatusThetaInvolutionResidualZero,
		StatusThetaKreinSelfAdjointResidualZero,
		StatusGThetaPositiveDefinite,
		StatusProjectorCompatibilityResidualZero,
		StatusFiniteKreinHilbertPlumbingVerified,
		StatusNoObservedHilbertDataImportedDefault,
		StatusFailedSyntheticThetaNative,
		StatusFailedPositiveMatrixNotHilbert,
		StatusFailedPositiveMatrixNotWick,
		StatusFailedPositiveMatrixNotOS,
		StatusFailedPositiveMatrixNotEnergy,
		StatusFailedPositiveMatrixNotUnitary,
		StatusFailedPositiveMatrixNotGlobal,
		StatusFailedSyntheticThetaNotTimeArrow,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 532 Registry Audit — Synthetic Fundamental-Symmetry Ledger Adapter and Positivity Residual Dry Run\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\nGate 532 inherits Gate 531's Wick/Hilbert airlock. A fundamental-symmetry row may be loaded only as synthetic, source-tagged, matrix-positivity-only bridge data.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic ledger import\n\n" + a.Import.Reason + "\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## Θ / Krein positivity residuals\n\nThe adapter evaluates only finite matrix obligations: `Θ²=I`, `ΘᵀG=GΘ`, symmetry and positivity of `H=GΘ`, compatibility `[Θ,P_530]=0`, and `R_time²=I`.\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
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

func matrixShapeOK(x [][]float64, rows, cols int) bool {
	if len(x) != rows {
		return false
	}
	for _, r := range x {
		if len(r) != cols {
			return false
		}
		for _, v := range r {
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return false
			}
		}
	}
	return true
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}

func unique(xs []string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func resolvePath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "../../.."))
	return filepath.Join(root, path)
}
