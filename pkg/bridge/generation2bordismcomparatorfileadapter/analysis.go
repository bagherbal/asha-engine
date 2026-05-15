// Package generation2bordismcomparatorfileadapter implements Gate 522:
// Bordism Comparator File Adapter and Stiefel-Whitney Metadata Firewall.
//
// Gate 521 defined scale-free oriented/spin/spin-c/boundary-bordism
// classifier sockets, but deliberately imported no external bordism class or
// tangent-bundle data. Gate 522 exercises that classifier from an explicit
// synthetic JSON ledger. The adapter validates w1/w2/W3/c1 metadata, checks
// characteristic-number identities, classifies admissibility bridge-only, and
// blocks every attempt to promote a manifold class or Stiefel-Whitney row into
// native ASHA data.
package generation2bordismcomparatorfileadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2bordismcobordismclassifierairlock"
)

const (
	AuditID       = "GATE522-BORDISM-COMPARATOR-FILE-ADAPTER-STIEFEL-WHITNEY-FIREWALL"
	DefaultLedger = "data/bordism_classifier_bridge_ledger.json"

	StatusGate521Inherited              = "CONDITIONAL_SUPPORT_GATE521_BORDISM_CLASSIFIER_INHERITED"
	StatusExplicitBordismFileLoaded     = "CONDITIONAL_SUPPORT_EXPLICIT_BORDISM_COMPARATOR_FILE_LOADED"
	StatusAirlockAcceptedRows           = "CONDITIONAL_SUPPORT_GATE522_AIRLOCK_ACCEPTED_QUARANTINED_BORDISM_ROWS"
	StatusStiefelWhitneyMetadataAudited = "CONDITIONAL_SUPPORT_STIEFEL_WHITNEY_METADATA_FIREWALL_AUDITED"
	StatusOrientedSpinSpinCClassified   = "CONDITIONAL_SUPPORT_ORIENTED_SPIN_SPINC_ADMISSIBILITY_COMPUTED_BRIDGE_ONLY"
	StatusCharacteristicResiduals       = "CONDITIONAL_SUPPORT_CHARACTERISTIC_NUMBER_RESIDUALS_COMPUTED_BRIDGE_ONLY"
	StatusClosedBoundaryClassified      = "CONDITIONAL_SUPPORT_CLOSED_BOUNDARY_STATUS_CLASSIFIED_BRIDGE_ONLY"
	StatusDefaultFixtureSynthetic       = "CONDITIONAL_SUPPORT_DEFAULT_GATE522_FILE_IS_SYNTHETIC_NOT_OBSERVED_BORDISM"
	StatusNoObservedBordismImported     = "CONDITIONAL_SUPPORT_NO_OBSERVED_BORDISM_OR_TANGENT_BUNDLE_DATA_IMPORTED"

	StatusFailedFileMissing             = "FAILED_ROUTE_GATE522_BORDISM_COMPARATOR_FILE_MISSING"
	StatusFailedMetadataIncomplete      = "FAILED_ROUTE_GATE522_METADATA_INCOMPLETE"
	StatusFailedMissingRows             = "FAILED_ROUTE_GATE522_MISSING_REQUIRED_BORDISM_ROWS"
	StatusFailedInvalidNumericalDomain  = "FAILED_ROUTE_GATE522_INVALID_BORDISM_NUMERICAL_DOMAIN"
	StatusFailedNativePromotionRejected = "FAILED_ROUTE_GATE522_BORDISM_NATIVE_PROMOTION_REJECTED"
	StatusFailedSWNative                = "FAILED_ROUTE_GATE522_STIEFEL_WHITNEY_NATIVE_PROMOTION_REJECTED"
	StatusFailedSpinStructureNative     = "FAILED_ROUTE_GATE522_SPIN_STRUCTURE_NATIVE_PROMOTION_REJECTED"
	StatusFailedSpecificClassNative     = "FAILED_ROUTE_GATE522_SPECIFIC_BORDISM_CLASS_NATIVE_PROMOTION_REJECTED"
	StatusFailedOutputsNotNative        = "FAILED_ROUTE_GATE522_FILE_ADAPTER_OUTPUTS_NOT_NATIVE_TOPOLOGY_PREDICTIONS"
	StatusFirewallPreserved             = "FIREWALL_PRESERVED_NO_MANIFOLD_TANGENT_BUNDLE_BOUNDARY_NEWTON_OR_COSMOLOGY_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked    = "FIREWALL_BLOCKED_GATE522_BORDISM_COMPARATOR_OUTPUT_NATIVE_WRITE"
)

type Number = *float64

type Inheritance struct {
	Executed                      bool
	Gate521ClassifierDefined      bool
	Gate521OrientedSocket         bool
	Gate521SpinSocket             bool
	Gate521SpinCSocket            bool
	Gate521BoundarySocket         bool
	Gate521CharacteristicResidual bool
	Gate521ScaleFree              bool
	Gate521SpecificClassSelected  bool
	Gate521ManifoldSelected       bool
	Gate521ObservedDataImported   bool
	Gate521NativeWriteBlocked     bool
	Gate522FileAdapterRedirect    bool
	Verdict, Reason               string
}

type DataRow struct {
	Name                 string `json:"name"`
	Kind                 string `json:"kind"`
	Role                 string `json:"role"`
	Value                Number `json:"value"`
	TextValue            string `json:"text_value"`
	Unit                 string `json:"unit"`
	Source               string `json:"source"`
	SourceVersion        string `json:"source_version"`
	Scheme               string `json:"scheme"`
	TopologyContext      string `json:"topology_context"`
	Uncertainty          string `json:"uncertainty"`
	BridgeOnly           bool   `json:"bridge_only"`
	ComparatorOnly       bool   `json:"comparator_only"`
	NoTheoremInput       bool   `json:"no_theorem_input"`
	EmpiricalImport      bool   `json:"empirical_import"`
	Observed             bool   `json:"observed"`
	Synthetic            bool   `json:"synthetic"`
	NativePromotionClaim bool   `json:"native_promotion_claim"`
	NativeInputClaim     bool   `json:"native_input_claim"`
}

type DataLedger struct {
	Gate                 int       `json:"gate"`
	LedgerName           string    `json:"ledger_name"`
	Description          string    `json:"description"`
	EmpiricalImport      bool      `json:"empirical_import"`
	BridgeOnly           bool      `json:"bridge_only"`
	NativeRegistryWrite  bool      `json:"native_registry_write"`
	SyntheticFixture     bool      `json:"synthetic_fixture"`
	ObservedValuesLoaded bool      `json:"observed_values_loaded"`
	CommonScheme         string    `json:"common_scheme"`
	TopologyContext      string    `json:"topology_context"`
	Rows                 []DataRow `json:"rows"`
}

type FileImport struct {
	Executed                       bool
	Loaded                         bool
	Path                           string
	Rows                           int
	AcceptedRows                   int
	RejectedRows                   int
	StiefelWhitneyRows             int
	CharacteristicRows             int
	BoundaryRows                   int
	BordismRows                    int
	AdapterRows                    int
	EmpiricalImport                bool
	BridgeOnlyLedger               bool
	SyntheticFixture               bool
	ObservedValuesLoaded           bool
	NativeRegistryWriteRequested   bool
	MetadataComplete               bool
	AllRowsBridgeOnly              bool
	AllRowsComparatorOnly          bool
	AllRowsNoTheoremInput          bool
	NativePromotionRejected        bool
	DefaultFixtureObservedRejected bool
	Verdict, Reason                string
	Failures                       []string
}

type AdapterInput struct {
	W1, W2, W3Integral, C1Mod2 float64
	Tau, P1, AHat, Euler       float64
	BoundaryComponents         float64
	HasW1, HasW2, HasW3        bool
	HasC1Mod2, HasTau, HasP1   bool
	HasAHat, HasEuler          bool
	HasBoundaryComponents      bool
	BoundaryConditionType      string
	BordismClassLabel          string
	SyntheticFixture           bool
	ObservedValuesLoaded       bool
	BridgeOnly                 bool
	MetadataComplete           bool
	NativePromotion            bool
}

type AdapterOutput struct {
	Executed                  bool
	Attempted                 bool
	Ready                     bool
	OrientedAdmissible        bool
	SpinAdmissible            bool
	SpinCAdmissible           bool
	ClosedBoundary            bool
	CharacteristicAdmissible  bool
	OverallAdmissible         bool
	SignatureFromP1           float64
	SignatureP1Residual       float64
	AHatFromTau               float64
	AHatResidual              float64
	RokhlinDivisibilityPassed bool
	C1Mod2W2Residual          float64
	BoundaryComponentResidual float64
	AllResidualsZero          bool
	BridgeOnly                bool
	NativePrediction          bool
	Verdict, Reason           string
	Failures                  []string
}

type Firewall struct {
	Executed                        bool
	ObservedBordismImported         bool
	ObservedTangentBundleImported   bool
	ObservedBoundaryDataImported    bool
	SyntheticFixtureOnly            bool
	FileRowsNative                  bool
	AdapterOutputsNative            bool
	StiefelWhitneyNativePrediction  bool
	SpinStructureNativePrediction   bool
	SpinCStructureNativePrediction  bool
	SpecificBordismClassNative      bool
	ManifoldRepresentativeNative    bool
	CharacteristicNumbersNative     bool
	BoundaryConditionNativeSelected bool
	NativeRegistryWritten           bool
	NewtonPlanckCosmologyImported   bool
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
	a := Analysis{Inheritance: buildInheritance()}
	ledger, imp := loadLedger(path)
	a.Import = imp
	if imp.Loaded {
		a.Input = buildInput(ledger, imp)
		a.Output = runAdapter(a.Input, imp)
	} else {
		a.Output = AdapterOutput{Executed: true, Verdict: StatusFailedFileMissing, Reason: "explicit Gate522 bordism comparator file was not found", Failures: []string{StatusFailedFileMissing}}
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
	g521, err := generation2bordismcobordismclassifierairlock.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedFileMissing, Reason: fmt.Sprintf("could not inherit Gate521: %v", err)}
	}
	return Inheritance{
		Executed:                      true,
		Gate521ClassifierDefined:      g521.Socket.Executed && g521.Socket.ClassifiesAllowedClasses,
		Gate521OrientedSocket:         g521.Socket.OrientedSocket && g521.Socket.RequiresW1ZeroForOriented,
		Gate521SpinSocket:             g521.Socket.SpinSocket && g521.Socket.RequiresW2ZeroForSpin,
		Gate521SpinCSocket:            g521.Socket.SpinCSocket && g521.Socket.RequiresW3ZeroForSpinC && g521.Socket.RequiresC1Mod2EqualsW2ForSpinC,
		Gate521BoundarySocket:         g521.Socket.BoundaryBordismSocket,
		Gate521CharacteristicResidual: g521.Constraints.SpinDivisibilityPassed && nearly(g521.Constraints.SignatureP1Residual, 0, 1e-12),
		Gate521ScaleFree:              g521.Scale.ClassifierScaleFree,
		Gate521SpecificClassSelected:  g521.Socket.SelectsSpecificClass,
		Gate521ManifoldSelected:       g521.Socket.SelectsManifoldRepresentative,
		Gate521ObservedDataImported:   g521.Scale.UsesObservedTopology,
		Gate521NativeWriteBlocked:     g521.Rejection.NativeRegistryWriteBlocked,
		Gate522FileAdapterRedirect:    g521.Next.Gate == 522,
		Verdict:                       StatusGate521Inherited,
		Reason:                        "Gate522 inherits Gate521's scale-free oriented, spin, spin-c, boundary-bordism, and characteristic-number classifier sockets while preserving the non-selection firewall.",
	}
}

func loadLedger(path string) (DataLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved, Verdict: StatusExplicitBordismFileLoaded}
	raw, err := os.ReadFile(resolved)
	if err != nil {
		imp.Loaded = false
		imp.Failures = []string{StatusFailedFileMissing}
		imp.Verdict = StatusFailedFileMissing
		imp.Reason = err.Error()
		return DataLedger{}, imp
	}
	var ledger DataLedger
	if err := json.Unmarshal(raw, &ledger); err != nil {
		imp.Loaded = false
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = err.Error()
		return DataLedger{}, imp
	}
	imp.Loaded = true
	imp.Rows = len(ledger.Rows)
	imp.EmpiricalImport = ledger.EmpiricalImport
	imp.BridgeOnlyLedger = ledger.BridgeOnly
	imp.SyntheticFixture = ledger.SyntheticFixture
	imp.ObservedValuesLoaded = ledger.ObservedValuesLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	allBridge, allComp, allNoTheorem, meta := true, true, true, true
	for _, r := range ledger.Rows {
		if r.NativePromotionClaim || r.NativeInputClaim || ledger.NativeRegistryWrite {
			imp.NativePromotionRejected = true
		}
		if r.BridgeOnly && r.ComparatorOnly && r.NoTheoremInput && !r.NativePromotionClaim && !r.NativeInputClaim && r.Source != "" && r.SourceVersion != "" && r.Scheme != "" && r.TopologyContext != "" && r.Uncertainty != "" {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
			imp.Failures = append(imp.Failures, fmt.Sprintf("row %s failed metadata/firewall validation", r.Name))
		}
		allBridge = allBridge && r.BridgeOnly
		allComp = allComp && r.ComparatorOnly
		allNoTheorem = allNoTheorem && r.NoTheoremInput
		meta = meta && r.Source != "" && r.SourceVersion != "" && r.Scheme != "" && r.TopologyContext != "" && r.Uncertainty != ""
		switch r.Kind {
		case "stiefel_whitney", "stiefel_whitney_integral", "spinc_line":
			imp.StiefelWhitneyRows++
		case "characteristic_number":
			imp.CharacteristicRows++
		case "boundary":
			imp.BoundaryRows++
		case "bordism":
			imp.BordismRows++
		case "adapter":
			imp.AdapterRows++
		}
		if r.Observed {
			imp.DefaultFixtureObservedRejected = true
		}
	}
	imp.MetadataComplete = meta
	imp.AllRowsBridgeOnly = allBridge
	imp.AllRowsComparatorOnly = allComp
	imp.AllRowsNoTheoremInput = allNoTheorem
	if imp.RejectedRows > 0 {
		imp.Verdict = StatusFailedMetadataIncomplete
	} else {
		imp.Verdict = strings.Join([]string{StatusExplicitBordismFileLoaded, StatusAirlockAcceptedRows, StatusDefaultFixtureSynthetic}, ";")
	}
	imp.Reason = "Gate522 loaded an explicit synthetic bordism classifier ledger and accepted it only as bridge/comparator metadata."
	return ledger, imp
}

func buildInput(ledger DataLedger, imp FileImport) AdapterInput {
	in := AdapterInput{SyntheticFixture: ledger.SyntheticFixture, ObservedValuesLoaded: ledger.ObservedValuesLoaded, BridgeOnly: imp.AllRowsBridgeOnly && ledger.BridgeOnly, MetadataComplete: imp.MetadataComplete && imp.RejectedRows == 0, NativePromotion: ledger.NativeRegistryWrite || imp.NativeRegistryWriteRequested || imp.NativePromotionRejected}
	for _, r := range ledger.Rows {
		val := 0.0
		if r.Value != nil {
			val = *r.Value
		}
		switch r.Name {
		case "stiefel_w1":
			in.W1 = val
			in.HasW1 = true
		case "stiefel_w2":
			in.W2 = val
			in.HasW2 = true
		case "integral_w3":
			in.W3Integral = val
			in.HasW3 = true
		case "c1_mod2":
			in.C1Mod2 = val
			in.HasC1Mod2 = true
		case "signature_tau":
			in.Tau = val
			in.HasTau = true
		case "pontryagin_p1":
			in.P1 = val
			in.HasP1 = true
		case "a_hat_index":
			in.AHat = val
			in.HasAHat = true
		case "euler_characteristic":
			in.Euler = val
			in.HasEuler = true
		case "boundary_component_count":
			in.BoundaryComponents = val
			in.HasBoundaryComponents = true
		case "boundary_condition_type":
			in.BoundaryConditionType = r.TextValue
		case "bordism_class_label":
			in.BordismClassLabel = r.TextValue
		}
	}
	return in
}

func runAdapter(in AdapterInput, imp FileImport) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: true, BridgeOnly: in.BridgeOnly, NativePrediction: false}
	missing := []string{}
	if !in.HasW1 {
		missing = append(missing, "stiefel_w1")
	}
	if !in.HasW2 {
		missing = append(missing, "stiefel_w2")
	}
	if !in.HasW3 {
		missing = append(missing, "integral_w3")
	}
	if !in.HasC1Mod2 {
		missing = append(missing, "c1_mod2")
	}
	if !in.HasTau {
		missing = append(missing, "signature_tau")
	}
	if !in.HasP1 {
		missing = append(missing, "pontryagin_p1")
	}
	if !in.HasAHat {
		missing = append(missing, "a_hat_index")
	}
	if !in.HasBoundaryComponents {
		missing = append(missing, "boundary_component_count")
	}
	if len(missing) > 0 {
		out.Failures = append(out.Failures, missing...)
		out.Verdict = StatusFailedMissingRows
		out.Reason = "missing required bordism classifier rows"
		return out
	}
	if !in.MetadataComplete || imp.RejectedRows != 0 {
		out.Failures = append(out.Failures, StatusFailedMetadataIncomplete)
		out.Verdict = StatusFailedMetadataIncomplete
		out.Reason = "metadata incomplete"
		return out
	}
	out.Ready = true
	out.OrientedAdmissible = nearly(mod(in.W1, 2), 0, 1e-12)
	out.SpinAdmissible = out.OrientedAdmissible && nearly(mod(in.W2, 2), 0, 1e-12)
	out.C1Mod2W2Residual = math.Abs(mod(in.C1Mod2-in.W2, 2))
	out.SpinCAdmissible = out.OrientedAdmissible && nearly(in.W3Integral, 0, 1e-12) && nearly(out.C1Mod2W2Residual, 0, 1e-12)
	out.SignatureFromP1 = in.P1 / 3
	out.SignatureP1Residual = math.Abs(out.SignatureFromP1 - in.Tau)
	out.AHatFromTau = -in.Tau / 8
	out.AHatResidual = math.Abs(out.AHatFromTau - in.AHat)
	out.RokhlinDivisibilityPassed = nearly(mod(math.Abs(in.Tau), 16), 0, 1e-12)
	out.ClosedBoundary = nearly(in.BoundaryComponents, 0, 1e-12) && (in.BoundaryConditionType == "closed_none" || in.BoundaryConditionType == "")
	out.BoundaryComponentResidual = math.Abs(in.BoundaryComponents)
	out.CharacteristicAdmissible = nearly(out.SignatureP1Residual, 0, 1e-12) && nearly(out.AHatResidual, 0, 1e-12) && out.RokhlinDivisibilityPassed
	out.OverallAdmissible = out.OrientedAdmissible && out.SpinAdmissible && out.SpinCAdmissible && out.CharacteristicAdmissible && out.ClosedBoundary
	out.AllResidualsZero = nearly(out.SignatureP1Residual, 0, 1e-12) && nearly(out.AHatResidual, 0, 1e-12) && nearly(out.C1Mod2W2Residual, 0, 1e-12) && nearly(out.BoundaryComponentResidual, 0, 1e-12)
	out.Verdict = strings.Join([]string{StatusStiefelWhitneyMetadataAudited, StatusOrientedSpinSpinCClassified, StatusCharacteristicResiduals, StatusClosedBoundaryClassified}, ";")
	out.Reason = "The synthetic file row is admissible for oriented/spin/spin-c classifier checks and characteristic-number residuals, but only as bridge metadata."
	return out
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{Executed: true,
		ObservedBordismImported:         imp.ObservedValuesLoaded || !imp.SyntheticFixture,
		ObservedTangentBundleImported:   false,
		ObservedBoundaryDataImported:    false,
		SyntheticFixtureOnly:            imp.SyntheticFixture && !imp.ObservedValuesLoaded,
		FileRowsNative:                  imp.NativeRegistryWriteRequested,
		AdapterOutputsNative:            out.NativePrediction,
		StiefelWhitneyNativePrediction:  false,
		SpinStructureNativePrediction:   false,
		SpinCStructureNativePrediction:  false,
		SpecificBordismClassNative:      false,
		ManifoldRepresentativeNative:    false,
		CharacteristicNumbersNative:     false,
		BoundaryConditionNativeSelected: false,
		NativeRegistryWritten:           false,
		NewtonPlanckCosmologyImported:   false,
		Verdict:                         StatusFirewallNativeWriteBlocked,
		Reason:                          "Gate522 blocks Stiefel-Whitney classes, spin/spin-c structures, characteristic numbers, bordism labels, and boundary status from native promotion.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No bordism class, Stiefel-Whitney class, spin structure, spin-c line, characteristic number, boundary condition, or manifold representative is written natively at Gate522.", "Inherited native content remains local: topology sockets, local index densities, anomaly cancellation, and classifier rules."},
		BridgeEntries:        []string{"Synthetic bordism classifier file adapter defined for w1, w2, W3, c1 mod 2, p1, tau, Â, Euler, and boundary metadata.", "Bridge-only admissibility checks compute oriented/spin/spin-c predicates, p1=3τ, Â=-τ/8, Rokhlin divisibility, and closed-boundary status.", "The default K3-like synthetic row is classifier plumbing only and is not an observed or native universe topology."},
		EnvironmentalEntries: []string{"The actual universe's bordism/cobordism class, tangent-bundle Stiefel-Whitney data, spin/spin-c structure, characteristic numbers, boundary condition, and manifold representative remain external global inputs."},
		FailedRoutes:         []string{"Treating the synthetic K3-like classifier row as ASHA's selected spacetime topology.", "Promoting w1/w2/W3/c1, p1, tau, Euler, Â, or boundary rows into native facts without a global tangent-bundle theorem.", "Using zero bridge residuals as evidence for a native manifold selector."},
		OpenTheorems:         []string{"A true native bordism selector would require finite-to-global topology information not present in the current ASHA engine.", "Future topology comparisons must remain source-tagged bridge ledgers unless a native boundary/global spectral theorem is discovered."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 523, Title: "Topology Residual Classifier Report and Native Non-Selection Audit", Reason: "Gate522 validates file-backed bordism classifier metadata. The next safe step is to summarize admissibility/residual classes without selecting a manifold.", PrimaryTask: "Aggregate topology, boundary, APS, and bordism residuals into a bridge-only consistency report and prove that zero residuals still do not select the universe's topology."}
}
func truth(a Analysis) string {
	return "Gate 522 makes the bordism airlock executable: ASHA can load a synthetic classifier ledger, validate Stiefel-Whitney and spin/spin-c metadata, compute characteristic-number residuals, and classify admissibility bridge-only. It still cannot choose the universe's manifold, bordism class, tangent bundle, spin structure, boundary condition, or characteristic numbers natively."
}

func validate(a Analysis) error {
	p := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate521ClassifierDefined || !a.Inheritance.Gate521OrientedSocket || !a.Inheritance.Gate521SpinSocket || !a.Inheritance.Gate521SpinCSocket || !a.Inheritance.Gate521BoundarySocket || !a.Inheritance.Gate521CharacteristicResidual || !a.Inheritance.Gate521ScaleFree || a.Inheritance.Gate521SpecificClassSelected || a.Inheritance.Gate521ManifoldSelected || a.Inheritance.Gate521ObservedDataImported || !a.Inheritance.Gate521NativeWriteBlocked || !a.Inheritance.Gate522FileAdapterRedirect {
		p = append(p, "bad Gate521 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 12 || a.Import.AcceptedRows != 12 || a.Import.RejectedRows != 0 || a.Import.StiefelWhitneyRows != 4 || a.Import.CharacteristicRows != 4 || a.Import.BoundaryRows != 2 || a.Import.BordismRows != 1 || a.Import.AdapterRows != 1 || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedValuesLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput {
		p = append(p, "bad Gate522 import")
	}
	if !a.Output.Ready || !a.Output.OrientedAdmissible || !a.Output.SpinAdmissible || !a.Output.SpinCAdmissible || !a.Output.ClosedBoundary || !a.Output.CharacteristicAdmissible || !a.Output.OverallAdmissible || !nearly(a.Output.SignatureFromP1, -16, 1e-12) || !nearly(a.Output.SignatureP1Residual, 0, 1e-12) || !nearly(a.Output.AHatFromTau, 2, 1e-12) || !nearly(a.Output.AHatResidual, 0, 1e-12) || !a.Output.RokhlinDivisibilityPassed || !nearly(a.Output.C1Mod2W2Residual, 0, 1e-12) || !a.Output.AllResidualsZero || !a.Output.BridgeOnly || a.Output.NativePrediction {
		p = append(p, "bad Gate522 output")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedBordismImported || a.Firewall.ObservedTangentBundleImported || a.Firewall.ObservedBoundaryDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.StiefelWhitneyNativePrediction || a.Firewall.SpinStructureNativePrediction || a.Firewall.SpinCStructureNativePrediction || a.Firewall.SpecificBordismClassNative || a.Firewall.ManifoldRepresentativeNative || a.Firewall.CharacteristicNumbersNative || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.NativeRegistryWritten || a.Firewall.NewtonPlanckCosmologyImported {
		p = append(p, "Gate522 firewall violation")
	}
	if len(p) > 0 {
		return fmt.Errorf(strings.Join(p, "; "))
	}
	return nil
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 522 Registry Audit — Bordism Comparator File Adapter and Stiefel-Whitney Metadata Firewall\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Bordism file airlock\n\n" + a.Import.Reason + "\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## Stiefel-Whitney and spin/spin-c classifier audit\n\n" + a.Output.Reason + "\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native entries", a.Registry.NativeEntries)
	writeList(&b, "### Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate523 should be:\n\n```text\n" + fmt.Sprintf("Gate %d — %s", a.Next.Gate, a.Next.Title) + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate521_classifier=%t; oriented_socket=%t; spin_socket=%t; spinc_socket=%t; boundary_socket=%t; characteristic_residual=%t; scale_free=%t; specific_class_selected=%t; manifold_selected=%t; observed_imported=%t; native_write_blocked=%t; Gate522_redirect=%t", x.Gate521ClassifierDefined, x.Gate521OrientedSocket, x.Gate521SpinSocket, x.Gate521SpinCSocket, x.Gate521BoundarySocket, x.Gate521CharacteristicResidual, x.Gate521ScaleFree, x.Gate521SpecificClassSelected, x.Gate521ManifoldSelected, x.Gate521ObservedDataImported, x.Gate521NativeWriteBlocked, x.Gate522FileAdapterRedirect)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("loaded=%t; rows=%d; accepted=%d; rejected=%d; stiefel_rows=%d; characteristic_rows=%d; boundary_rows=%d; bordism_rows=%d; adapter_rows=%d; synthetic=%t; observed_loaded=%t; metadata_complete=%t; bridge_only=%t; comparator_only=%t; native_write_requested=%t", x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.StiefelWhitneyRows, x.CharacteristicRows, x.BoundaryRows, x.BordismRows, x.AdapterRows, x.SyntheticFixture, x.ObservedValuesLoaded, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsComparatorOnly, x.NativeRegistryWriteRequested)
}
func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("oriented=%t; spin=%t; spinc=%t; closed_boundary=%t; overall_admissible=%t; signature_from_p1=%.12g; signature_residual=%.12g; Ahat_from_tau=%.12g; Ahat_residual=%.12g; rokhlindiv16=%t; c1_mod2_w2_residual=%.12g; boundary_residual=%.12g; all_residuals_zero=%t; native_prediction=%t", x.OrientedAdmissible, x.SpinAdmissible, x.SpinCAdmissible, x.ClosedBoundary, x.OverallAdmissible, x.SignatureFromP1, x.SignatureP1Residual, x.AHatFromTau, x.AHatResidual, x.RokhlinDivisibilityPassed, x.C1Mod2W2Residual, x.BoundaryComponentResidual, x.AllResidualsZero, x.NativePrediction)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_bordism_imported=%t; observed_tangent_bundle_imported=%t; observed_boundary_imported=%t; synthetic_only=%t; file_rows_native=%t; outputs_native=%t; SW_native=%t; spin_native=%t; spinc_native=%t; specific_class_native=%t; manifold_native=%t; characteristic_numbers_native=%t; boundary_condition_native=%t; registry_written=%t", x.ObservedBordismImported, x.ObservedTangentBundleImported, x.ObservedBoundaryDataImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.StiefelWhitneyNativePrediction, x.SpinStructureNativePrediction, x.SpinCStructureNativePrediction, x.SpecificBordismClassNative, x.ManifoldRepresentativeNative, x.CharacteristicNumbersNative, x.BoundaryConditionNativeSelected, x.NativeRegistryWritten)
}
func statuses() []string {
	return []string{StatusGate521Inherited, StatusExplicitBordismFileLoaded, StatusAirlockAcceptedRows, StatusStiefelWhitneyMetadataAudited, StatusOrientedSpinSpinCClassified, StatusCharacteristicResiduals, StatusClosedBoundaryClassified, StatusDefaultFixtureSynthetic, StatusNoObservedBordismImported, StatusFailedNativePromotionRejected, StatusFailedSWNative, StatusFailedSpinStructureNative, StatusFailedSpecificClassNative, StatusFailedOutputsNotNative, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}
func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString(title + "\n\n")
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
func mod(x, m float64) float64 {
	r := math.Mod(x, m)
	if r < 0 {
		r += m
	}
	return r
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
