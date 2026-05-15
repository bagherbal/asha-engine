// Package generation2observedtopologyboundaryfileadapter implements Gate 520:
// Observed Topology and Boundary File Adapter Firewall.
//
// Gate 519 defined the fail-closed topology/boundary comparator preflight and
// imported no global topology or boundary spectrum. Gate 520 exercises the same
// airlock from an explicit JSON ledger. The checked-in ledger is deliberately
// synthetic, not observed cosmological topology: it proves that file loading,
// provenance validation, APS/signature residual arithmetic, and native-write
// blocking compose correctly without promoting external manifold data into an
// ASHA theorem.
package generation2observedtopologyboundaryfileadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2observedtopologyboundarypreflight"
)

const (
	AuditID       = "GATE520-OBSERVED-TOPOLOGY-BOUNDARY-FILE-ADAPTER-FIREWALL"
	DefaultLedger = "data/topology_boundary_observed_bridge_ledger.json"

	StatusGate519PreflightInherited      = "CONDITIONAL_SUPPORT_GATE519_TOPOLOGY_BOUNDARY_PREFLIGHT_INHERITED"
	StatusExplicitTopologyFileLoaded     = "CONDITIONAL_SUPPORT_EXPLICIT_TOPOLOGY_BOUNDARY_COMPARATOR_FILE_LOADED"
	StatusAirlockAcceptedQuarantinedRows = "CONDITIONAL_SUPPORT_GATE520_AIRLOCK_ACCEPTED_QUARANTINED_TOPOLOGY_BOUNDARY_ROWS"
	StatusFileAdapterExecutedBridgeOnly  = "CONDITIONAL_SUPPORT_GATE520_TOPOLOGY_BOUNDARY_FILE_ADAPTER_EXECUTED_BRIDGE_ONLY"
	StatusAPSFormulaComputedFromFile     = "CONDITIONAL_SUPPORT_APS_INDEX_FORMULA_COMPUTED_FROM_FILE_INPUTS"
	StatusSignatureResidualComputed      = "CONDITIONAL_SUPPORT_SIGNATURE_PONTRYAGIN_RESIDUAL_COMPUTED_BRIDGE_ONLY"
	StatusClosedOpenBoundaryClassified   = "CONDITIONAL_SUPPORT_CLOSED_VERSUS_BOUNDARY_STATUS_CLASSIFIED_BRIDGE_ONLY"
	StatusComparatorResidualsComputed    = "CONDITIONAL_SUPPORT_TOPOLOGY_BOUNDARY_COMPARATOR_RESIDUALS_COMPUTED_BRIDGE_ONLY"
	StatusDefaultFixtureSynthetic        = "CONDITIONAL_SUPPORT_DEFAULT_GATE520_FILE_IS_SYNTHETIC_NOT_OBSERVED_TOPOLOGY"
	StatusNoObservedTopologyLoaded       = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_OR_BOUNDARY_NUMBERS_IMPORTED_BY_DEFAULT"

	StatusFailedFileMissing             = "FAILED_ROUTE_GATE520_TOPOLOGY_BOUNDARY_COMPARATOR_FILE_MISSING"
	StatusFailedMetadataIncomplete      = "FAILED_ROUTE_GATE520_METADATA_INCOMPLETE"
	StatusFailedMissingRequiredRows     = "FAILED_ROUTE_GATE520_MISSING_REQUIRED_TOPOLOGY_OR_BOUNDARY_ROWS"
	StatusFailedInvalidNumericalDomain  = "FAILED_ROUTE_GATE520_INVALID_TOPOLOGY_BOUNDARY_NUMERICAL_DOMAIN"
	StatusFailedNativePromotionRejected = "FAILED_ROUTE_GATE520_TOPOLOGY_BOUNDARY_NATIVE_PROMOTION_REJECTED"
	StatusFailedGlobalIndexNative       = "FAILED_ROUTE_GATE520_GLOBAL_APS_INDEX_NATIVE_PROMOTION_REJECTED"
	StatusFailedEtaNative               = "FAILED_ROUTE_GATE520_ETA_NATIVE_PROMOTION_REJECTED"
	StatusFailedBoundarySpectrumNative  = "FAILED_ROUTE_GATE520_BOUNDARY_SPECTRUM_NATIVE_PROMOTION_REJECTED"
	StatusFailedOutputsNotNative        = "FAILED_ROUTE_GATE520_FILE_ADAPTER_OUTPUTS_NOT_NATIVE_TOPOLOGY_PREDICTIONS"
	StatusFailedObservedClaimDefault    = "FAILED_ROUTE_GATE520_DEFAULT_FIXTURE_OBSERVED_TOPOLOGY_CLAIM_REJECTED"
	StatusFirewallPreserved             = "FIREWALL_PRESERVED_GATE520_TOPOLOGY_BOUNDARY_FILE_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked    = "FIREWALL_BLOCKED_GATE520_TOPOLOGY_BOUNDARY_FILE_OUTPUT_NATIVE_WRITE"
)

type Number = *float64

type Inheritance struct {
	Executed                      bool
	Gate519PreflightDefined       bool
	Gate519TopologyRows           int
	Gate519BoundaryRows           int
	Gate519RequiresBridgeOnly     bool
	Gate519RejectsNativePromotion bool
	Gate519ComparatorExecuted     bool
	Gate519ObservedDataImported   bool
	Gate519NativeRegistryBlocked  bool
	Gate520FileAdapterRedirect    bool
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
	TopologyRows                   int
	BoundaryRows                   int
	AdapterRows                    int
	ComparatorRows                 int
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
	NativeRegistryWriteRejected    bool
	DefaultFixtureObservedRejected bool
	Verdict, Reason                string
	Failures                       []string
}

type AdapterInput struct {
	EulerCharacteristic        float64
	PontryaginP1               float64
	SignatureTau               float64
	GlobalAPSIndex             float64
	LocalIndexIntegral         float64
	EtaInvariant               float64
	KernelDimensionH           float64
	BoundaryComponentCount     float64
	HasEuler, HasP1, HasTau    bool
	HasGlobalAPS, HasLocal     bool
	HasEta, HasH, HasBoundary  bool
	BoundaryConditionType      string
	BoundarySpectrumDescriptor string
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
	BoundaryCorrection        float64
	ComputedAPSIndex          float64
	ComputedSignatureFromP1   float64
	APSResidual               float64
	SignatureResidual         float64
	ClosedManifold            bool
	BoundaryMode              bool
	UsesAPSBoundaryCorrection bool
	AllResidualsZero          bool
	BridgeOnly                bool
	NativePrediction          bool
	Verdict, Reason           string
	Failures                  []string
}

type Firewall struct {
	Executed                         bool
	ObservedTopologyImported         bool
	ObservedBoundaryDataImported     bool
	ObservedBoundarySpectrumImported bool
	SyntheticFixtureOnly             bool
	FileRowsNative                   bool
	AdapterOutputsNative             bool
	EulerNativePrediction            bool
	PontryaginNativePrediction       bool
	SignatureNativePrediction        bool
	GlobalAPSIndexNativePrediction   bool
	EtaNativePrediction              bool
	BoundarySpectrumNativePrediction bool
	BoundaryConditionNativeSelected  bool
	NativeRegistryWritten            bool
	NewtonPlanckCosmologyImported    bool
	Verdict, Reason                  string
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
		a.Output = AdapterOutput{Executed: true, Verdict: StatusFailedFileMissing, Reason: "explicit Gate520 topology/boundary comparator file was not found", Failures: []string{StatusFailedFileMissing}}
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
	g519, err := generation2observedtopologyboundarypreflight.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedFileMissing, Reason: fmt.Sprintf("could not inherit Gate519 preflight: %v", err)}
	}
	return Inheritance{
		Executed:                      true,
		Gate519PreflightDefined:       g519.Topology.Executed && g519.Boundary.Executed && g519.Policy.Executed,
		Gate519TopologyRows:           g519.Topology.RequiredRows,
		Gate519BoundaryRows:           g519.Boundary.RequiredRows,
		Gate519RequiresBridgeOnly:     g519.Policy.RequiresBridgeOnlyTrue,
		Gate519RejectsNativePromotion: g519.Policy.RequiresNativePromotionFalse && g519.Rejection.TopologyNativePredictionBlocked,
		Gate519ComparatorExecuted:     g519.Firewall.ComparatorExecuted,
		Gate519ObservedDataImported:   g519.Firewall.ObservedTopologyImported || g519.Firewall.ObservedBoundaryDataImported,
		Gate519NativeRegistryBlocked:  !g519.Firewall.NativeTopologyWrite && !g519.Firewall.NativeBoundaryWrite && !g519.Firewall.NativeGlobalIndexWrite,
		Gate520FileAdapterRedirect:    g519.Next.Gate == 520,
		Verdict:                       StatusGate519PreflightInherited,
		Reason:                        "Gate520 inherits Gate519's fail-closed topology/boundary schema, bridge-only metadata policy, and explicit native-write rejection.",
	}
}

func loadLedger(path string) (DataLedger, FileImport) {
	resolved := resolvePath(path)
	imp := FileImport{Executed: true, Path: resolved}
	b, err := os.ReadFile(resolved)
	if err != nil {
		imp.Verdict = StatusFailedFileMissing
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedFileMissing}
		return DataLedger{}, imp
	}
	var ledger DataLedger
	if err := json.Unmarshal(b, &ledger); err != nil {
		imp.Loaded = true
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		return ledger, imp
	}
	imp.Loaded = true
	imp.Rows = len(ledger.Rows)
	imp.EmpiricalImport = ledger.EmpiricalImport
	imp.BridgeOnlyLedger = ledger.BridgeOnly
	imp.SyntheticFixture = ledger.SyntheticFixture
	imp.ObservedValuesLoaded = ledger.ObservedValuesLoaded
	imp.NativeRegistryWriteRequested = ledger.NativeRegistryWrite
	imp.NativeRegistryWriteRejected = !ledger.NativeRegistryWrite
	required := map[string]bool{
		"euler_characteristic": false, "pontryagin_p1": false, "signature_tau": false, "global_aps_index": false,
		"local_index_integral": false, "manifold_dimension": false, "orientation_and_closedness": false, "topology_model_id": false,
		"boundary_condition_type": false, "eta_invariant_value": false, "kernel_dimension_h": false, "boundary_spectrum_descriptor": false,
		"boundary_orientation": false, "boundary_component_count": false, "boundary_model_id": false,
	}
	metadataComplete := ledger.LedgerName != "" && ledger.CommonScheme != "" && ledger.TopologyContext != ""
	allBridge := ledger.BridgeOnly
	allComparator := true
	allNoTheorem := true
	accepted := 0
	failures := []string{}
	for _, r := range ledger.Rows {
		if _, ok := required[r.Name]; ok {
			required[r.Name] = true
		}
		switch r.Kind {
		case "topology":
			imp.TopologyRows++
		case "boundary":
			imp.BoundaryRows++
		case "adapter":
			imp.AdapterRows++
		case "comparator":
			imp.ComparatorRows++
		}
		rowMeta := r.Name != "" && r.Kind != "" && r.Role != "" && r.Source != "" && r.SourceVersion != "" && r.Scheme != "" && r.TopologyContext != "" && r.Uncertainty != ""
		if !rowMeta {
			metadataComplete = false
		}
		if !r.BridgeOnly {
			allBridge = false
		}
		if !r.ComparatorOnly {
			allComparator = false
		}
		if !r.NoTheoremInput {
			allNoTheorem = false
		}
		if r.NativePromotionClaim || r.NativeInputClaim {
			failures = append(failures, StatusFailedNativePromotionRejected)
		}
		if rowMeta && r.BridgeOnly && r.ComparatorOnly && r.NoTheoremInput && !r.NativePromotionClaim && !r.NativeInputClaim {
			accepted++
		}
	}
	missing := false
	for _, seen := range required {
		if !seen {
			missing = true
		}
	}
	if missing {
		failures = append(failures, StatusFailedMissingRequiredRows)
	}
	if !metadataComplete {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if !ledger.BridgeOnly || ledger.NativeRegistryWrite {
		failures = append(failures, StatusFailedNativePromotionRejected)
	}
	if ledger.ObservedValuesLoaded && ledger.SyntheticFixture {
		failures = append(failures, StatusFailedObservedClaimDefault)
	}
	imp.AcceptedRows = accepted
	imp.RejectedRows = imp.Rows - accepted
	imp.MetadataComplete = metadataComplete && !missing
	imp.AllRowsBridgeOnly = allBridge
	imp.AllRowsComparatorOnly = allComparator
	imp.AllRowsNoTheoremInput = allNoTheorem
	imp.NativePromotionRejected = !ledger.NativeRegistryWrite
	imp.DefaultFixtureObservedRejected = ledger.SyntheticFixture && !ledger.ObservedValuesLoaded
	imp.Failures = unique(failures)
	if len(imp.Failures) == 0 {
		imp.Verdict = strings.Join([]string{StatusExplicitTopologyFileLoaded, StatusAirlockAcceptedQuarantinedRows, StatusDefaultFixtureSynthetic, StatusNoObservedTopologyLoaded}, ";")
		imp.Reason = "Explicit topology/boundary file loaded as a synthetic, bridge-only comparator fixture with complete Gate519 metadata and no native-promotion request."
	} else {
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "Gate520 file failed one or more topology/boundary airlock requirements."
	}
	return ledger, imp
}

func buildInput(ledger DataLedger, imp FileImport) AdapterInput {
	in := AdapterInput{SyntheticFixture: ledger.SyntheticFixture, ObservedValuesLoaded: ledger.ObservedValuesLoaded, BridgeOnly: ledger.BridgeOnly, MetadataComplete: imp.MetadataComplete, NativePromotion: ledger.NativeRegistryWrite}
	for _, r := range ledger.Rows {
		v := 0.0
		has := false
		if r.Value != nil {
			v = *r.Value
			has = true
		}
		switch r.Name {
		case "euler_characteristic":
			in.EulerCharacteristic, in.HasEuler = v, has
		case "pontryagin_p1":
			in.PontryaginP1, in.HasP1 = v, has
		case "signature_tau":
			in.SignatureTau, in.HasTau = v, has
		case "global_aps_index":
			in.GlobalAPSIndex, in.HasGlobalAPS = v, has
		case "local_index_integral":
			in.LocalIndexIntegral, in.HasLocal = v, has
		case "eta_invariant_value":
			in.EtaInvariant, in.HasEta = v, has
		case "kernel_dimension_h":
			in.KernelDimensionH, in.HasH = v, has
		case "boundary_component_count":
			in.BoundaryComponentCount, in.HasBoundary = v, has
		case "boundary_condition_type":
			in.BoundaryConditionType = r.TextValue
		case "boundary_spectrum_descriptor":
			in.BoundarySpectrumDescriptor = r.TextValue
		}
	}
	return in
}

func runAdapter(in AdapterInput, imp FileImport) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: imp.Loaded, BridgeOnly: true, NativePrediction: false}
	ready := imp.Loaded && imp.MetadataComplete && imp.AcceptedRows == imp.Rows && imp.Rows >= 15 && in.HasP1 && in.HasTau && in.HasGlobalAPS && in.HasLocal && in.HasEta && in.HasH && in.HasBoundary && !in.NativePromotion && in.BridgeOnly
	out.Ready = ready
	if !ready {
		out.Verdict = strings.Join([]string{StatusFailedMetadataIncomplete, StatusFailedMissingRequiredRows}, ";")
		out.Reason = "Topology/boundary file adapter did not meet the complete bridge-only input domain."
		out.Failures = []string{StatusFailedMetadataIncomplete, StatusFailedMissingRequiredRows}
		return out
	}
	if in.BoundaryComponentCount < 0 || math.IsNaN(in.LocalIndexIntegral) || math.IsNaN(in.EtaInvariant) || math.IsNaN(in.KernelDimensionH) {
		out.Verdict = StatusFailedInvalidNumericalDomain
		out.Reason = "Invalid topology/boundary numerical domain."
		out.Failures = []string{StatusFailedInvalidNumericalDomain}
		return out
	}
	out.BoundaryCorrection = (in.EtaInvariant + in.KernelDimensionH) / 2.0
	out.ComputedAPSIndex = in.LocalIndexIntegral - out.BoundaryCorrection
	out.ComputedSignatureFromP1 = in.PontryaginP1 / 3.0
	out.APSResidual = math.Abs(in.GlobalAPSIndex - out.ComputedAPSIndex)
	out.SignatureResidual = math.Abs(in.SignatureTau - out.ComputedSignatureFromP1)
	out.ClosedManifold = in.BoundaryComponentCount == 0
	out.BoundaryMode = !out.ClosedManifold
	out.UsesAPSBoundaryCorrection = out.BoundaryMode
	out.AllResidualsZero = nearly(out.APSResidual, 0, 1e-12) && nearly(out.SignatureResidual, 0, 1e-12)
	out.Verdict = strings.Join([]string{StatusFileAdapterExecutedBridgeOnly, StatusAPSFormulaComputedFromFile, StatusSignatureResidualComputed, StatusClosedOpenBoundaryClassified, StatusComparatorResidualsComputed}, ";")
	out.Reason = "Gate520 computed bridge-only APS and signature/Pontryagin residuals from a quarantined file fixture; zero residuals validate adapter plumbing, not native manifold prediction."
	return out
}

func buildFirewall(imp FileImport, out AdapterOutput) Firewall {
	return Firewall{
		Executed:                         true,
		ObservedTopologyImported:         imp.ObservedValuesLoaded,
		ObservedBoundaryDataImported:     imp.ObservedValuesLoaded,
		ObservedBoundarySpectrumImported: false,
		SyntheticFixtureOnly:             imp.SyntheticFixture && !imp.ObservedValuesLoaded,
		FileRowsNative:                   false,
		AdapterOutputsNative:             out.NativePrediction,
		EulerNativePrediction:            false,
		PontryaginNativePrediction:       false,
		SignatureNativePrediction:        false,
		GlobalAPSIndexNativePrediction:   false,
		EtaNativePrediction:              false,
		BoundarySpectrumNativePrediction: false,
		BoundaryConditionNativeSelected:  false,
		NativeRegistryWritten:            imp.NativeRegistryWriteRequested,
		NewtonPlanckCosmologyImported:    false,
		Verdict:                          strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusFailedOutputsNotNative, StatusFailedGlobalIndexNative, StatusFailedEtaNative, StatusFailedBoundarySpectrumNative}, ";"),
		Reason:                           "Gate520 file rows and residual outputs remain bridge-only comparator artifacts. No topology integer, eta invariant, boundary spectrum, boundary condition, or gravity/cosmology scale is written natively.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new native manifold topology, boundary eta, global APS index, Euler characteristic, signature, Pontryagin number, or boundary condition is written at Gate520.",
			"Inherited native content remains limited to local index-density sockets, characteristic-class sockets, APS formula socket, anomaly-inflow socket, and mixed gravitational trace cancellation.",
		},
		BridgeEntries: []string{
			"File-backed topology/boundary comparator adapter implemented with Gate519 metadata validation.",
			"Bridge-only APS residual and signature/Pontryagin residual computations executed on an explicit synthetic fixture.",
			"Closed-versus-boundary status classified as comparator metadata only.",
		},
		EnvironmentalEntries: []string{
			"Actual χ(M), p_i(M), τ(M), global APS index, eta invariant, h, boundary spectrum, boundary condition, and global manifold topology remain environmental/global input data.",
		},
		FailedRoutes: []string{
			"Promoting file-loaded topology or boundary rows into native manifold selection.",
			"Treating zero bridge residuals as proof that ASHA predicted the global topology or boundary spectrum.",
			"Using APS/global index rows as theorem inputs instead of comparator targets.",
		},
		OpenTheorems: []string{
			"A real observed-topology comparator may be supplied later, but must remain bridge-only and fully source-tagged.",
			"A native bordism/manifold selector or native boundary spectral theorem would be required before any global topology write.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 521, Title: "Bordism and Cobordism Classifier Airlock", Reason: "Gate520 validates comparator plumbing but still cannot select global topology. The next native-compatible question is whether finite ASHA data constrain only bordism classes rather than specific manifolds.", PrimaryTask: "Audit oriented/spin/spin-c bordism sockets and determine whether ASHA can classify allowed topology classes without selecting a specific universe topology."}
}

func truth(a Analysis) string {
	return "Gate 520 safely loads a topology/boundary comparator file and computes APS/signature residuals only inside the bridge airlock. The adapter can check external global-topology hypotheses against ASHA's local sockets, but no file row, zero residual, eta value, boundary condition, or global index becomes a native ASHA theorem."
}

func validate(a Analysis) error {
	problems := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate519PreflightDefined || !a.Inheritance.Gate519RequiresBridgeOnly || !a.Inheritance.Gate519RejectsNativePromotion || a.Inheritance.Gate519ComparatorExecuted || a.Inheritance.Gate519ObservedDataImported || !a.Inheritance.Gate519NativeRegistryBlocked || !a.Inheritance.Gate520FileAdapterRedirect {
		problems = append(problems, "bad Gate519 inheritance")
	}
	if !a.Import.Loaded || a.Import.Rows != 15 || a.Import.AcceptedRows != 15 || a.Import.RejectedRows != 0 || a.Import.TopologyRows != 7 || a.Import.BoundaryRows != 7 || a.Import.AdapterRows != 1 || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || !a.Import.SyntheticFixture || a.Import.ObservedValuesLoaded || a.Import.NativeRegistryWriteRequested || !a.Import.MetadataComplete || !a.Import.AllRowsBridgeOnly || !a.Import.AllRowsComparatorOnly || !a.Import.AllRowsNoTheoremInput {
		problems = append(problems, "bad file import")
	}
	if !a.Output.Executed || !a.Output.Attempted || !a.Output.Ready || !a.Output.BridgeOnly || a.Output.NativePrediction || !a.Output.UsesAPSBoundaryCorrection || !a.Output.AllResidualsZero || !nearly(a.Output.BoundaryCorrection, 2, 1e-12) || !nearly(a.Output.ComputedAPSIndex, 9, 1e-12) || !nearly(a.Output.ComputedSignatureFromP1, 1, 1e-12) {
		problems = append(problems, "bad adapter output")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedTopologyImported || a.Firewall.ObservedBoundaryDataImported || !a.Firewall.SyntheticFixtureOnly || a.Firewall.FileRowsNative || a.Firewall.AdapterOutputsNative || a.Firewall.EulerNativePrediction || a.Firewall.PontryaginNativePrediction || a.Firewall.SignatureNativePrediction || a.Firewall.GlobalAPSIndexNativePrediction || a.Firewall.EtaNativePrediction || a.Firewall.BoundarySpectrumNativePrediction || a.Firewall.BoundaryConditionNativeSelected || a.Firewall.NativeRegistryWritten || a.Firewall.NewtonPlanckCosmologyImported {
		problems = append(problems, "firewall violation")
	}
	if len(problems) > 0 {
		return fmt.Errorf(strings.Join(problems, "; "))
	}
	return nil
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 520 Registry Audit — Observed Topology and Boundary File Adapter Firewall\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## File-backed topology/boundary import\n\n" + a.Import.Reason + "\n\n```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("## Bridge APS and signature residuals\n\n" + a.Output.Reason + "\n\n```text\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "### Native entries", a.Registry.NativeEntries)
	writeList(&b, "### Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "### Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "### Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "### Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate521 should be:\n\n```text\n" + fmt.Sprintf("Gate %d — %s", a.Next.Gate, a.Next.Title) + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate519 preflight=%t; topology_rows=%d; boundary_rows=%d; bridge_only_required=%t; native_promotion_rejected=%t; comparator_executed=%t; observed_data_imported=%t; native_write_blocked=%t; Gate520_redirect=%t", x.Gate519PreflightDefined, x.Gate519TopologyRows, x.Gate519BoundaryRows, x.Gate519RequiresBridgeOnly, x.Gate519RejectsNativePromotion, x.Gate519ComparatorExecuted, x.Gate519ObservedDataImported, x.Gate519NativeRegistryBlocked, x.Gate520FileAdapterRedirect)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("loaded=%t; rows=%d; accepted=%d; rejected=%d; topology_rows=%d; boundary_rows=%d; adapter_rows=%d; empirical_import=%t; bridge_only=%t; synthetic_fixture=%t; observed_values_loaded=%t; native_write_requested=%t; metadata_complete=%t; all_bridge_only=%t; all_comparator_only=%t; no_theorem_input=%t", x.Loaded, x.Rows, x.AcceptedRows, x.RejectedRows, x.TopologyRows, x.BoundaryRows, x.AdapterRows, x.EmpiricalImport, x.BridgeOnlyLedger, x.SyntheticFixture, x.ObservedValuesLoaded, x.NativeRegistryWriteRequested, x.MetadataComplete, x.AllRowsBridgeOnly, x.AllRowsComparatorOnly, x.AllRowsNoTheoremInput)
}
func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("ready=%t; boundary_correction=(eta+h)/2=%.12g; computed_APS=%.12g; APS_residual=%.12g; computed_signature=p1/3=%.12g; signature_residual=%.12g; closed=%t; boundary_mode=%t; all_residuals_zero=%t; bridge_only=%t; native_prediction=%t", x.Ready, x.BoundaryCorrection, x.ComputedAPSIndex, x.APSResidual, x.ComputedSignatureFromP1, x.SignatureResidual, x.ClosedManifold, x.BoundaryMode, x.AllResidualsZero, x.BridgeOnly, x.NativePrediction)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_topology=%t; observed_boundary=%t; observed_boundary_spectrum=%t; synthetic_only=%t; file_rows_native=%t; outputs_native=%t; chi_native=%t; p_native=%t; signature_native=%t; APS_native=%t; eta_native=%t; boundary_spectrum_native=%t; boundary_condition_native=%t; native_registry_written=%t; Newton_Planck_cosmology_imported=%t", x.ObservedTopologyImported, x.ObservedBoundaryDataImported, x.ObservedBoundarySpectrumImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.EulerNativePrediction, x.PontryaginNativePrediction, x.SignatureNativePrediction, x.GlobalAPSIndexNativePrediction, x.EtaNativePrediction, x.BoundarySpectrumNativePrediction, x.BoundaryConditionNativeSelected, x.NativeRegistryWritten, x.NewtonPlanckCosmologyImported)
}

func statuses() []string {
	return []string{StatusGate519PreflightInherited, StatusExplicitTopologyFileLoaded, StatusAirlockAcceptedQuarantinedRows, StatusFileAdapterExecutedBridgeOnly, StatusAPSFormulaComputedFromFile, StatusSignatureResidualComputed, StatusClosedOpenBoundaryClassified, StatusComparatorResidualsComputed, StatusDefaultFixtureSynthetic, StatusNoObservedTopologyLoaded, StatusFailedNativePromotionRejected, StatusFailedGlobalIndexNative, StatusFailedEtaNative, StatusFailedBoundarySpectrumNative, StatusFailedOutputsNotNative, StatusFailedObservedClaimDefault, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
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
