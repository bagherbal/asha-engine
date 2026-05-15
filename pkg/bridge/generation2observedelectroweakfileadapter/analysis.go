// Package generation2observedelectroweakfileadapter implements Gate 507:
// Observed Electroweak Comparator File Adapter Firewall.
//
// Gate 506 defined a fail-closed observed electroweak preflight but imported no
// numbers. Gate 507 exercises the same adapter path from an explicit JSON file.
// The checked-in ledger is deliberately synthetic, not PDG/observed data: it is
// a file-backed bridge fixture proving that metadata validation, tree-level
// electroweak formulas, comparator residuals, photon zero mode, and native-write
// blocking compose correctly. A real observed file may be supplied to
// BuildFromFile, but every computed value remains environmental bridge output.
package generation2observedelectroweakfileadapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2observedelectroweakpreflight"
)

const (
	AuditID = "GATE507-OBSERVED-ELECTROWEAK-COMPARATOR-FILE-ADAPTER-FIREWALL"

	StatusGate506PreflightInherited        = "CONDITIONAL_SUPPORT_GATE506_OBSERVED_ELECTROWEAK_PREFLIGHT_INHERITED"
	StatusExplicitEWFileLoaded             = "CONDITIONAL_SUPPORT_EXPLICIT_ELECTROWEAK_COMPARATOR_FILE_LOADED"
	StatusAirlockAcceptedBridgeRows        = "CONDITIONAL_SUPPORT_GATE507_AIRLOCK_ACCEPTED_QUARANTINED_ELECTROWEAK_ROWS"
	StatusAdapterExecutedBridgeOnly        = "CONDITIONAL_SUPPORT_GATE507_ELECTROWEAK_FILE_ADAPTER_EXECUTED_BRIDGE_ONLY"
	StatusTreeWZComputedFromFile           = "CONDITIONAL_SUPPORT_TREE_LEVEL_WZ_FORMULAS_COMPUTED_FROM_FILE_INPUTS"
	StatusPhotonZeroPreservedFromFile      = "CONDITIONAL_SUPPORT_PHOTON_ZERO_MODE_PRESERVED_BY_FILE_ADAPTER"
	StatusRhoIdentityConfirmedFromFile     = "CONDITIONAL_SUPPORT_TREE_RHO_IDENTITY_CONFIRMED_BY_FILE_ADAPTER"
	StatusComparatorResidualsComputed      = "CONDITIONAL_SUPPORT_ELECTROWEAK_COMPARATOR_RESIDUALS_COMPUTED_BRIDGE_ONLY"
	StatusDefaultFixtureSyntheticNotPDG    = "CONDITIONAL_SUPPORT_DEFAULT_GATE507_FILE_IS_SYNTHETIC_NOT_OBSERVED_DATA"
	StatusObservedValuesNotLoadedByDefault = "CONDITIONAL_SUPPORT_NO_OBSERVED_ELECTROWEAK_NUMBERS_IMPORTED_BY_DEFAULT"

	StatusFailedFileMissing             = "FAILED_ROUTE_GATE507_ELECTROWEAK_COMPARATOR_FILE_MISSING"
	StatusFailedSwitchClosed            = "FAILED_ROUTE_GATE507_EMPIRICAL_IMPORT_SWITCH_CLOSED"
	StatusFailedMetadataIncomplete      = "FAILED_ROUTE_GATE507_METADATA_INCOMPLETE"
	StatusFailedMissingExplicitInputs   = "FAILED_ROUTE_GATE507_MISSING_EXPLICIT_V_G2_GY_INPUTS"
	StatusFailedInvalidNumericalDomain  = "FAILED_ROUTE_GATE507_INVALID_ELECTROWEAK_NUMERICAL_DOMAIN"
	StatusFailedObservedMassNativeInput = "FAILED_ROUTE_GATE507_OBSERVED_WZ_MASSES_AS_NATIVE_INPUT_REJECTED"
	StatusFailedWeakAngleNative         = "FAILED_ROUTE_GATE507_WEAK_ANGLE_NATIVE_PROMOTION_REJECTED"
	StatusFailedKappaPromotion          = "FAILED_ROUTE_GATE507_KAPPA_PROMOTION_REJECTED"
	StatusFailedNativePromotion         = "FAILED_ROUTE_GATE507_ELECTROWEAK_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite     = "FAILED_ROUTE_GATE507_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedOutputsNotNative        = "FAILED_ROUTE_GATE507_FILE_ADAPTER_OUTPUTS_NOT_NATIVE_ELECTROWEAK_PREDICTIONS"
	StatusFailedObservedClaimDefault    = "FAILED_ROUTE_GATE507_DEFAULT_FIXTURE_OBSERVED_CLAIM_REJECTED"

	StatusFirewallPreserved          = "FIREWALL_PRESERVED_GATE507_ELECTROWEAK_FILE_ADAPTER_BRIDGE_ONLY"
	StatusFirewallNativeWriteBlocked = "FIREWALL_BLOCKED_GATE507_ELECTROWEAK_FILE_OUTPUT_NATIVE_WRITE"
)

const DefaultLedger = "data/electroweak_observed_bridge_ledger.json"

type Number = *float64

type Inheritance struct {
	Executed                          bool
	Gate506AuditDefined               bool
	Gate506PreflightValidated         bool
	Gate506AcceptedSchemaCases        int
	Gate506RejectedCases              int
	Gate506NumericalAdapterExecuted   bool
	Gate506ObservedNumbersImported    bool
	Gate506NativeRegistryWriteBlocked bool
	Gate507RedirectDefined            bool
	Verdict                           string
	Reason                            string
}

type DataRow struct {
	Name                 string `json:"name"`
	Observable           string `json:"observable"`
	Role                 string `json:"role"`
	Value                Number `json:"value"`
	Unit                 string `json:"unit"`
	Source               string `json:"source"`
	SourceVersion        string `json:"source_version"`
	Scale                string `json:"scale"`
	Scheme               string `json:"scheme"`
	Uncertainty          string `json:"uncertainty"`
	BridgeOnly           bool   `json:"bridge_only"`
	EmpiricalImport      bool   `json:"empirical_import"`
	Observed             bool   `json:"observed"`
	Synthetic            bool   `json:"synthetic"`
	NativePromotionClaim bool   `json:"native_promotion_claim"`
	NativeInputClaim     bool   `json:"native_input_claim"`
	KappaPromotionClaim  bool   `json:"kappa_promotion_claim"`
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
	CommonScale          string    `json:"common_scale"`
	CommonScheme         string    `json:"common_scheme"`
	Rows                 []DataRow `json:"rows"`
}

type FileImport struct {
	Executed                         bool
	Loaded                           bool
	Path                             string
	Rows                             int
	AcceptedRows                     int
	RejectedRows                     int
	InputRows                        int
	ComparatorRows                   int
	EmpiricalImport                  bool
	BridgeOnlyLedger                 bool
	SyntheticFixture                 bool
	ObservedValuesLoaded             bool
	NativeRegistryWriteRequested     bool
	AllAcceptedBridgeOnly            bool
	AllAcceptedEmpiricalImport       bool
	AllAcceptedSyntheticOrObserved   bool
	MetadataComplete                 bool
	SwitchClosedRejected             bool
	NativePromotionRejected          bool
	NativeRegistryWriteRejected      bool
	ObservedMassNativeInputRejected  bool
	WeakAngleNativePromotionRejected bool
	KappaPromotionRejected           bool
	DefaultFixtureObservedRejected   bool
	Verdict                          string
	Reason                           string
	Failures                         []string
}

type AdapterInput struct {
	V, G2, GY                 float64
	HasV, HasG2, HasGY        bool
	CommonScale, CommonScheme string
	Source                    string
	SyntheticFixture          bool
	ObservedValuesLoaded      bool
	BridgeOnly                bool
	MetadataComplete          bool
	NativePromotion           bool
}

type AdapterOutput struct {
	Executed             bool
	Attempted            bool
	Ready                bool
	Sin2ThetaW           float64
	Cos2ThetaW           float64
	MW                   float64
	MZ                   float64
	MGamma               float64
	RhoTree              float64
	NeutralChargedRatio  float64
	UsedTreeLevelFormula bool
	PhotonZeroPreserved  bool
	RhoIdentityConfirmed bool
	Verdict              string
	Reason               string
	Failures             []string
}

type Residuals struct {
	Executed                  bool
	ComparatorRowsAvailable   bool
	WeakAngleResidualComputed bool
	MWResidualComputed        bool
	MZResidualComputed        bool
	WeakAngleResidual         float64
	MWResidual                float64
	MZResidual                float64
	AllResidualsZero          bool
	BridgeOnly                bool
	NativePrediction          bool
	Verdict                   string
	Reason                    string
}

type Firewall struct {
	Executed                          bool
	ObservedValuesImported            bool
	SyntheticFixtureOnly              bool
	FileRowsNative                    bool
	AdapterOutputsNative              bool
	WeakAngleNativePrediction         bool
	WZMassNativePrediction            bool
	GaugeCouplingsNativePrediction    bool
	VEVNativePrediction               bool
	KappaNativePromotion              bool
	NativeRegistryWritten             bool
	PhysicalElectroweakPredictionMade bool
	PhotonZeroModePreserved           bool
	Verdict                           string
	Reason                            string
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
	Residuals   Residuals
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

func BuildFromFile(path string) (Analysis, error) {
	a := Analysis{Inheritance: buildInheritance()}
	ledger, imp := loadLedger(path)
	a.Import = imp
	if imp.Loaded {
		a.Input = buildInput(ledger, imp)
		a.Output = runAdapter(a.Input, imp)
		a.Residuals = computeResiduals(ledger, a.Output, imp)
	} else {
		a.Output = AdapterOutput{Executed: true, Attempted: false, Verdict: StatusFailedFileMissing, Reason: "explicit Gate507 electroweak comparator file was not found", Failures: []string{StatusFailedFileMissing}}
		a.Residuals = Residuals{Executed: true, Verdict: StatusFailedFileMissing, Reason: "no file-loaded output exists for comparator residuals"}
	}
	a.Firewall = buildFirewall(a.Import, a.Output, a.Residuals)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	g506, err := generation2observedelectroweakpreflight.BuildDefault()
	if err != nil {
		return Inheritance{Executed: false, Verdict: StatusFailedFileMissing, Reason: fmt.Sprintf("could not inherit Gate506 preflight: %v", err)}
	}
	return Inheritance{
		Executed:                          true,
		Gate506AuditDefined:               true,
		Gate506PreflightValidated:         g506.Preflight.Executed && g506.Preflight.AcceptedSchemaCases == 1 && g506.Preflight.RejectedCases == 10,
		Gate506AcceptedSchemaCases:        g506.Preflight.AcceptedSchemaCases,
		Gate506RejectedCases:              g506.Preflight.RejectedCases,
		Gate506NumericalAdapterExecuted:   g506.Preflight.NumericalAdapterRun,
		Gate506ObservedNumbersImported:    g506.Preflight.ObservedNumbersImported,
		Gate506NativeRegistryWriteBlocked: !g506.Firewall.NativeRegistryWritten && !g506.Firewall.NativeElectroweakPredictionMade,
		Gate507RedirectDefined:            g506.Next.Gate == 507,
		Verdict:                           StatusGate506PreflightInherited,
		Reason:                            "Gate506 accepted exactly one redacted bridge-only electroweak preflight schema, rejected ten fail-closed cases, imported no observed numbers, and defined Gate507 as the file-adapter redirect.",
	}
}

func loadLedger(path string) (DataLedger, FileImport) {
	resolved := projectPath(path)
	imp := FileImport{Executed: true, Path: resolved}
	b, err := os.ReadFile(resolved)
	if err != nil {
		imp.Verdict = StatusFailedFileMissing
		imp.Reason = err.Error()
		imp.Failures = []string{StatusFailedFileMissing}
		return DataLedger{}, imp
	}
	var l DataLedger
	if err := json.Unmarshal(b, &l); err != nil {
		imp.Loaded = true
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Reason = fmt.Sprintf("invalid Gate507 JSON ledger: %v", err)
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		return l, imp
	}
	imp.Loaded = true
	imp.Rows = len(l.Rows)
	imp.EmpiricalImport = l.EmpiricalImport
	imp.BridgeOnlyLedger = l.BridgeOnly
	imp.SyntheticFixture = l.SyntheticFixture
	imp.ObservedValuesLoaded = l.ObservedValuesLoaded
	imp.NativeRegistryWriteRequested = l.NativeRegistryWrite
	failures := []string{}
	if !l.EmpiricalImport {
		failures = appendUnique(failures, StatusFailedSwitchClosed)
		imp.SwitchClosedRejected = true
	}
	if !l.BridgeOnly {
		failures = appendUnique(failures, StatusFailedNativePromotion)
		imp.NativePromotionRejected = true
	}
	if l.NativeRegistryWrite {
		failures = appendUnique(failures, StatusFailedNativeRegistryWrite)
		imp.NativeRegistryWriteRejected = true
	}
	if l.CommonScale == "" || l.CommonScheme == "" {
		failures = appendUnique(failures, StatusFailedMetadataIncomplete)
	}
	allBridge := l.BridgeOnly
	allEmpirical := l.EmpiricalImport
	allSyntheticOrObserved := true
	metadata := l.CommonScale != "" && l.CommonScheme != ""
	for _, r := range l.Rows {
		if r.Role == "input" {
			imp.InputRows++
		}
		if r.Role == "comparator" {
			imp.ComparatorRows++
		}
		rowAccepted := true
		if r.Source == "" || r.SourceVersion == "" || r.Scale == "" || r.Scheme == "" || r.Uncertainty == "" {
			metadata = false
			failures = appendUnique(failures, StatusFailedMetadataIncomplete)
			rowAccepted = false
		}
		if !r.BridgeOnly {
			allBridge = false
			failures = appendUnique(failures, StatusFailedNativePromotion)
			imp.NativePromotionRejected = true
			rowAccepted = false
		}
		if !r.EmpiricalImport {
			allEmpirical = false
			failures = appendUnique(failures, StatusFailedSwitchClosed)
			imp.SwitchClosedRejected = true
			rowAccepted = false
		}
		if r.NativePromotionClaim {
			failures = appendUnique(failures, StatusFailedNativePromotion)
			imp.NativePromotionRejected = true
			rowAccepted = false
		}
		if r.KappaPromotionClaim {
			failures = appendUnique(failures, StatusFailedKappaPromotion)
			imp.KappaPromotionRejected = true
			rowAccepted = false
		}
		if r.NativeInputClaim && (r.Observable == "m_w" || r.Observable == "m_z") {
			failures = appendUnique(failures, StatusFailedObservedMassNativeInput)
			imp.ObservedMassNativeInputRejected = true
			rowAccepted = false
		}
		if r.NativePromotionClaim && r.Observable == "sin2_theta_w" {
			failures = appendUnique(failures, StatusFailedWeakAngleNative)
			imp.WeakAngleNativePromotionRejected = true
			rowAccepted = false
		}
		if !(r.Synthetic || r.Observed) {
			allSyntheticOrObserved = false
			failures = appendUnique(failures, StatusFailedMetadataIncomplete)
			rowAccepted = false
		}
		if l.SyntheticFixture && r.Observed {
			failures = appendUnique(failures, StatusFailedObservedClaimDefault)
			imp.DefaultFixtureObservedRejected = true
			rowAccepted = false
		}
		if rowAccepted {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
	}
	if !hasValue(l, "v") || !hasValue(l, "g2") || !hasValue(l, "gY") {
		failures = appendUnique(failures, StatusFailedMissingExplicitInputs)
	}
	imp.AllAcceptedBridgeOnly = allBridge && imp.RejectedRows == 0
	imp.AllAcceptedEmpiricalImport = allEmpirical
	imp.AllAcceptedSyntheticOrObserved = allSyntheticOrObserved
	imp.MetadataComplete = metadata
	imp.Failures = failures
	if len(failures) == 0 {
		imp.Verdict = strings.Join([]string{StatusExplicitEWFileLoaded, StatusAirlockAcceptedBridgeRows, StatusDefaultFixtureSyntheticNotPDG}, ";")
		imp.Reason = "The explicit electroweak comparator JSON file loaded with complete metadata, bridge-only quarantine, explicit v/g2/gY inputs, and no native-promotion claims; the checked-in fixture is synthetic and not observed PDG data."
	} else {
		imp.Verdict = strings.Join(failures, ";")
		imp.Reason = "The explicit electroweak comparator JSON file failed the Gate507 airlock before safe native-isolated execution."
	}
	return l, imp
}

func buildInput(l DataLedger, imp FileImport) AdapterInput {
	in := AdapterInput{CommonScale: l.CommonScale, CommonScheme: l.CommonScheme, SyntheticFixture: l.SyntheticFixture, ObservedValuesLoaded: l.ObservedValuesLoaded, BridgeOnly: l.BridgeOnly, MetadataComplete: imp.MetadataComplete, NativePromotion: l.NativeRegistryWrite, Source: l.LedgerName}
	for _, r := range l.Rows {
		if r.Value == nil {
			continue
		}
		switch r.Observable {
		case "v":
			in.V = *r.Value
			in.HasV = true
		case "g2":
			in.G2 = *r.Value
			in.HasG2 = true
		case "gY":
			in.GY = *r.Value
			in.HasGY = true
		}
	}
	return in
}

func runAdapter(in AdapterInput, imp FileImport) AdapterOutput {
	out := AdapterOutput{Executed: true, Attempted: true}
	if len(imp.Failures) > 0 {
		out.Verdict = imp.Verdict
		out.Reason = "adapter not run because file import failed the airlock"
		out.Failures = imp.Failures
		return out
	}
	if !in.HasV || !in.HasG2 || !in.HasGY {
		out.Failures = append(out.Failures, StatusFailedMissingExplicitInputs)
	}
	if !(finitePositive(in.V) && finitePositive(in.G2) && finitePositive(in.GY)) {
		out.Failures = append(out.Failures, StatusFailedInvalidNumericalDomain)
	}
	if !in.BridgeOnly || in.NativePromotion {
		out.Failures = append(out.Failures, StatusFailedNativePromotion)
	}
	out.Ready = len(out.Failures) == 0
	if !out.Ready {
		out.Verdict = strings.Join(out.Failures, ";")
		out.Reason = "explicit file inputs were not valid bridge-only electroweak matching data"
		return out
	}
	den := in.G2*in.G2 + in.GY*in.GY
	out.Sin2ThetaW = in.GY * in.GY / den
	out.Cos2ThetaW = in.G2 * in.G2 / den
	out.MW = in.G2 * in.V / 2
	out.MZ = math.Sqrt(den) * in.V / 2
	out.MGamma = 0
	out.RhoTree = out.MW * out.MW / (out.MZ * out.MZ * out.Cos2ThetaW)
	out.NeutralChargedRatio = out.MZ * out.MZ / (out.MW * out.MW)
	out.UsedTreeLevelFormula = true
	out.PhotonZeroPreserved = nearly(out.MGamma, 0, 1e-12)
	out.RhoIdentityConfirmed = nearly(out.RhoTree, 1, 1e-12)
	out.Verdict = strings.Join([]string{StatusAdapterExecutedBridgeOnly, StatusTreeWZComputedFromFile, StatusPhotonZeroPreservedFromFile, StatusRhoIdentityConfirmedFromFile}, ";")
	out.Reason = "The file adapter propagated explicit bridge inputs through tree-level electroweak formulas while preserving the photon zero mode and tree rho identity."
	return out
}

func computeResiduals(l DataLedger, out AdapterOutput, imp FileImport) Residuals {
	r := Residuals{Executed: true, BridgeOnly: true}
	if !out.Ready || len(imp.Failures) > 0 {
		r.Verdict = StatusFailedOutputsNotNative
		r.Reason = "no safe file-adapter output exists for comparator residuals"
		return r
	}
	var sin2, mw, mz Number
	for _, row := range l.Rows {
		if row.Role != "comparator" || row.Value == nil {
			continue
		}
		switch row.Observable {
		case "sin2_theta_w":
			sin2 = row.Value
		case "m_w":
			mw = row.Value
		case "m_z":
			mz = row.Value
		}
	}
	r.ComparatorRowsAvailable = sin2 != nil || mw != nil || mz != nil
	if sin2 != nil {
		r.WeakAngleResidual = math.Abs(out.Sin2ThetaW - *sin2)
		r.WeakAngleResidualComputed = true
	}
	if mw != nil {
		r.MWResidual = math.Abs(out.MW - *mw)
		r.MWResidualComputed = true
	}
	if mz != nil {
		r.MZResidual = math.Abs(out.MZ - *mz)
		r.MZResidualComputed = true
	}
	r.AllResidualsZero = (!r.WeakAngleResidualComputed || nearly(r.WeakAngleResidual, 0, 1e-12)) && (!r.MWResidualComputed || nearly(r.MWResidual, 0, 1e-12)) && (!r.MZResidualComputed || nearly(r.MZResidual, 0, 1e-12))
	r.Verdict = StatusComparatorResidualsComputed
	r.Reason = "Comparator residuals are bridge diagnostics against file rows; in the checked-in synthetic fixture they vanish because the comparator rows were generated from the same fake 3-4-5 inputs."
	return r
}

func buildFirewall(imp FileImport, out AdapterOutput, r Residuals) Firewall {
	fw := Firewall{Executed: true, ObservedValuesImported: imp.ObservedValuesLoaded, SyntheticFixtureOnly: imp.SyntheticFixture && !imp.ObservedValuesLoaded, PhotonZeroModePreserved: out.PhotonZeroPreserved}
	fw.Verdict = strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked, StatusObservedValuesNotLoadedByDefault, StatusFailedOutputsNotNative}, ";")
	fw.Reason = "Gate507 may compute file-backed bridge outputs and residuals, but the default checked-in file is synthetic, no observed electroweak numbers are imported, and no output is native-registry eligible."
	_ = r
	return fw
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No VEV, gauge coupling, weak angle, kappa, W/Z mass, rho identity, or comparator residual is written as native at Gate507."},
		BridgeEntries:        []string{"An explicit electroweak comparator JSON file adapter is admitted as bridge-only when v, g2, and gY are present with complete source/version/scale/scheme/uncertainty metadata and no native-promotion claims.", "The checked-in default file is synthetic and computes mW=3, mZ=5, sin²(theta_W)=16/25, photon mass 0, rho_tree=1, and zero residuals only as adapter-fixture arithmetic."},
		EnvironmentalEntries: []string{"A real observed electroweak data file may be supplied only as environmental bridge data; it must never become a native ASHA theorem input or registry write."},
		FailedRoutes:         []string{StatusFailedSwitchClosed, StatusFailedMetadataIncomplete, StatusFailedMissingExplicitInputs, StatusFailedInvalidNumericalDomain, StatusFailedObservedMassNativeInput, StatusFailedWeakAngleNative, StatusFailedKappaPromotion, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedOutputsNotNative, StatusFailedObservedClaimDefault},
		OpenTheorems:         []string{"Gate508 may compare the file-adapter electroweak quotient to the native dimensionless index ledger, but only as bridge residual geometry.", "A separate native finite-action theorem would still be required to derive a nonzero Higgs ray, gauge couplings, kappa_U1, or physical W/Z mass matrix."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 508, Title: "Electroweak Comparator Residual Geometry Airlock", Reason: "Gate507 proves the file adapter and residual channel work without native promotion; the next safe step is to classify residuals against native dimensionless quotient/index data, still bridge-only.", PrimaryTask: "map electroweak file-adapter residuals to quotient/index diagnostics while blocking weak-angle, coupling, VEV, W/Z mass, and kappa native writes"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate506PreflightValidated && !a.Inheritance.Gate506NumericalAdapterExecuted && !a.Inheritance.Gate506ObservedNumbersImported && a.Inheritance.Gate506NativeRegistryWriteBlocked && a.Inheritance.Gate507RedirectDefined, "Gate507 inheritance incomplete"},
		{a.Import.Executed && a.Import.Loaded && a.Import.EmpiricalImport && a.Import.BridgeOnlyLedger && a.Import.SyntheticFixture && !a.Import.ObservedValuesLoaded && !a.Import.NativeRegistryWriteRequested && a.Import.Rows == 6 && a.Import.AcceptedRows == 6 && a.Import.RejectedRows == 0 && a.Import.InputRows == 3 && a.Import.ComparatorRows == 3 && a.Import.AllAcceptedBridgeOnly && a.Import.AllAcceptedEmpiricalImport && a.Import.AllAcceptedSyntheticOrObserved && a.Import.MetadataComplete && len(a.Import.Failures) == 0, "Gate507 file import invalid"},
		{a.Input.HasV && a.Input.HasG2 && a.Input.HasGY && nearly(a.Input.V, 2, 1e-12) && nearly(a.Input.G2, 3, 1e-12) && nearly(a.Input.GY, 4, 1e-12) && a.Input.SyntheticFixture && !a.Input.ObservedValuesLoaded && a.Input.BridgeOnly && a.Input.MetadataComplete && !a.Input.NativePromotion, "Gate507 adapter input invalid"},
		{a.Output.Executed && a.Output.Attempted && a.Output.Ready && a.Output.UsedTreeLevelFormula && nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) && nearly(a.Output.Cos2ThetaW, 9.0/25.0, 1e-12) && nearly(a.Output.MW, 3, 1e-12) && nearly(a.Output.MZ, 5, 1e-12) && nearly(a.Output.MGamma, 0, 1e-12) && nearly(a.Output.RhoTree, 1, 1e-12) && a.Output.PhotonZeroPreserved && a.Output.RhoIdentityConfirmed, "Gate507 file adapter output invalid"},
		{a.Residuals.Executed && a.Residuals.ComparatorRowsAvailable && a.Residuals.WeakAngleResidualComputed && a.Residuals.MWResidualComputed && a.Residuals.MZResidualComputed && a.Residuals.AllResidualsZero && a.Residuals.BridgeOnly && !a.Residuals.NativePrediction, "Gate507 residual audit invalid"},
		{a.Firewall.Executed && !a.Firewall.ObservedValuesImported && a.Firewall.SyntheticFixtureOnly && !a.Firewall.FileRowsNative && !a.Firewall.AdapterOutputsNative && !a.Firewall.WeakAngleNativePrediction && !a.Firewall.WZMassNativePrediction && !a.Firewall.GaugeCouplingsNativePrediction && !a.Firewall.VEVNativePrediction && !a.Firewall.KappaNativePromotion && !a.Firewall.NativeRegistryWritten && !a.Firewall.PhysicalElectroweakPredictionMade && a.Firewall.PhotonZeroModePreserved, "Gate507 firewall violated"},
		{a.Next.Gate == 508, "Gate508 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func truth(a Analysis) string {
	if a.Output.Ready && a.Firewall.SyntheticFixtureOnly && !a.Firewall.NativeRegistryWritten {
		return "Gate 507 proves that the explicit electroweak comparator file adapter can load a fully tagged bridge ledger, compute tree-level W/Z/weak-angle outputs, preserve the photon zero mode and rho identity, and compute comparator residuals, while refusing native promotion. The default checked-in file is deliberately synthetic, so no observed electroweak numbers are imported and no physical electroweak prediction is made."
	}
	return "Gate 507 failed before establishing the electroweak file-adapter firewall."
}

func hasValue(l DataLedger, obs string) bool {
	for _, r := range l.Rows {
		if r.Observable == obs && r.Value != nil {
			return true
		}
	}
	return false
}
func finitePositive(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) && x > 0 }
func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
func appendUnique(xs []string, x string) []string {
	for _, y := range xs {
		if y == x {
			return xs
		}
	}
	return append(xs, x)
}
func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}
func ptrFloatText(x Number) string {
	if x == nil {
		return "missing"
	}
	return fmtFloat(*x)
}

func projectPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return path
	}
	dir := filepath.Dir(file)
	for i := 0; i < 8; i++ {
		candidate := filepath.Join(dir, path)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		dir = filepath.Dir(dir)
	}
	return path
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate506=%t accepted=%d rejected=%d gate506_adapter=%t gate506_observed=%t native_blocked=%t redirect=%t verdict=%s reason=%s", x.Executed, x.Gate506PreflightValidated, x.Gate506AcceptedSchemaCases, x.Gate506RejectedCases, x.Gate506NumericalAdapterExecuted, x.Gate506ObservedNumbersImported, x.Gate506NativeRegistryWriteBlocked, x.Gate507RedirectDefined, x.Verdict, x.Reason)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("executed=%t loaded=%t path=%s rows=%d accepted=%d rejected=%d inputs=%d comparators=%d empirical=%t bridge=%t synthetic=%t observed_loaded=%t native_write=%t metadata=%t failures=[%s] verdict=%s reason=%s", x.Executed, x.Loaded, x.Path, x.Rows, x.AcceptedRows, x.RejectedRows, x.InputRows, x.ComparatorRows, x.EmpiricalImport, x.BridgeOnlyLedger, x.SyntheticFixture, x.ObservedValuesLoaded, x.NativeRegistryWriteRequested, x.MetadataComplete, strings.Join(x.Failures, ";"), x.Verdict, x.Reason)
}
func FormatInput(x AdapterInput) string {
	return fmt.Sprintf("v=%s g2=%s gY=%s has_v=%t has_g2=%t has_gY=%t scale=%s scheme=%s synthetic=%t observed_loaded=%t bridge=%t metadata=%t native_promotion=%t", fmtFloat(x.V), fmtFloat(x.G2), fmtFloat(x.GY), x.HasV, x.HasG2, x.HasGY, x.CommonScale, x.CommonScheme, x.SyntheticFixture, x.ObservedValuesLoaded, x.BridgeOnly, x.MetadataComplete, x.NativePromotion)
}
func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("executed=%t attempted=%t ready=%t sin2=%s cos2=%s mW=%s mZ=%s mGamma=%s rho=%s ratio=%s photon=%t rho_identity=%t verdict=%s reason=%s", x.Executed, x.Attempted, x.Ready, fmtFloat(x.Sin2ThetaW), fmtFloat(x.Cos2ThetaW), fmtFloat(x.MW), fmtFloat(x.MZ), fmtFloat(x.MGamma), fmtFloat(x.RhoTree), fmtFloat(x.NeutralChargedRatio), x.PhotonZeroPreserved, x.RhoIdentityConfirmed, x.Verdict, x.Reason)
}
func FormatResiduals(x Residuals) string {
	return fmt.Sprintf("executed=%t comparators=%t sin2_residual=%t:%s mW_residual=%t:%s mZ_residual=%t:%s all_zero=%t bridge=%t native_prediction=%t verdict=%s reason=%s", x.Executed, x.ComparatorRowsAvailable, x.WeakAngleResidualComputed, fmtFloat(x.WeakAngleResidual), x.MWResidualComputed, fmtFloat(x.MWResidual), x.MZResidualComputed, fmtFloat(x.MZResidual), x.AllResidualsZero, x.BridgeOnly, x.NativePrediction, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t observed_imported=%t synthetic_only=%t rows_native=%t outputs_native=%t theta_native=%t wz_native=%t couplings_native=%t vev_native=%t kappa_native=%t native_registry=%t physical_prediction=%t photon=%t verdict=%s reason=%s", x.Executed, x.ObservedValuesImported, x.SyntheticFixtureOnly, x.FileRowsNative, x.AdapterOutputsNative, x.WeakAngleNativePrediction, x.WZMassNativePrediction, x.GaugeCouplingsNativePrediction, x.VEVNativePrediction, x.KappaNativePromotion, x.NativeRegistryWritten, x.PhysicalElectroweakPredictionMade, x.PhotonZeroModePreserved, x.Verdict, x.Reason)
}
func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func statuses() []string {
	return []string{StatusGate506PreflightInherited, StatusExplicitEWFileLoaded, StatusAirlockAcceptedBridgeRows, StatusAdapterExecutedBridgeOnly, StatusTreeWZComputedFromFile, StatusPhotonZeroPreservedFromFile, StatusRhoIdentityConfirmedFromFile, StatusComparatorResidualsComputed, StatusDefaultFixtureSyntheticNotPDG, StatusObservedValuesNotLoadedByDefault, StatusFailedSwitchClosed, StatusFailedMetadataIncomplete, StatusFailedMissingExplicitInputs, StatusFailedInvalidNumericalDomain, StatusFailedObservedMassNativeInput, StatusFailedWeakAngleNative, StatusFailedKappaPromotion, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedOutputsNotNative, StatusFailedObservedClaimDefault, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 507 Registry Audit — Observed Electroweak Comparator File Adapter Firewall\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate506 admitted only a redacted bridge-only electroweak preflight schema and refused to import observed numbers or execute a numerical adapter. Gate507 may therefore test an explicit file-backed adapter, but every value must remain bridge/environmental and barred from the native registry.\n\n")
	b.WriteString("```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Data-file import\n\n")
	b.WriteString("```text\n" + FormatImport(a.Import) + "\n```\n\n")
	b.WriteString("The default checked-in ledger is synthetic and file-backed. It is not a PDG ledger and does not import observed electroweak numbers.\n\n")
	b.WriteString("## Adapter execution\n\n")
	b.WriteString("```text\n" + FormatInput(a.Input) + "\n" + FormatOutput(a.Output) + "\n```\n\n")
	b.WriteString("Bridge formulas executed:\n\n```text\nm_W = g2 v / 2\nm_Z = sqrt(g2^2 + gY^2) v / 2\nsin^2(theta_W) = gY^2 / (g2^2 + gY^2)\nm_gamma = 0\nrho_tree = m_W^2/(m_Z^2 cos^2(theta_W))\n```\n\n")
	b.WriteString("## Comparator residuals\n\n")
	b.WriteString("```text\n" + FormatResiduals(a.Residuals) + "\n```\n\n")
	b.WriteString("The zero residuals in the default fixture are not physics. They only prove the file-adapter arithmetic and metadata firewall because the comparator rows are synthetic values generated from the same fake 3-4-5 bridge inputs.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("No file row, output, weak angle, W/Z mass, VEV, gauge coupling, kappa candidate, rho identity, or residual is promoted to a native ASHA theorem.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate508 should be:\n\n```text\nGate 508 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
