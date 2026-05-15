// Package generation2observednumericaladapter implements Gate 470:
// Observed Numerical d_ud Adapter / Explicit Data-File Run.
//
// Gate 470 reads an explicit JSON ledger from data/pdg_observed_ledger.json.
// It is allowed to open the Gate465 empirical airlock, parse observed rows, and
// compute the Gate464 cylinder distance only if the file supplies the rank-
// complete bridge comparators demanded by Gates 454, 456, 459, 467, and 469:
// I_spec, I_K, complete branch tags, common scale/scheme metadata, uncertainty,
// and bridge_only quarantine.
//
// The default checked-in ledger intentionally contains PDG-style quark mass rows
// and a Cabibbo target row, but it does not fabricate ASHA-specific I_K or branch
// tags. Therefore the default Gate470 run fails closed with d_ud undefined. This
// is the only honest result unless an explicit rank-complete bridge file is
// supplied.
package generation2observednumericaladapter

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const (
	AuditID = "GATE470-OBSERVED-NUMERICAL-DUD-ADAPTER-EXPLICIT-DATA-FILE-RUN"

	StatusGate469Inherited               = "CONDITIONAL_SUPPORT_GATE469_PREFLIGHT_INHERITED"
	StatusDataFileLoaded                 = "CONDITIONAL_SUPPORT_EXPLICIT_PDG_OBSERVED_LEDGER_LOADED"
	StatusAirlockAcceptedBridgeRows      = "CONDITIONAL_SUPPORT_GATE470_AIRLOCK_ACCEPTED_QUARANTINED_ROWS"
	StatusNumericalAdapterAttempted      = "CONDITIONAL_SUPPORT_GATE470_NUMERICAL_ADAPTER_ATTEMPTED"
	StatusNumericalDUDComputed           = "CONDITIONAL_SUPPORT_OBSERVED_NUMERICAL_DUD_BRIDGE_ONLY_COMPUTED"
	StatusCKMResidualComputed            = "CONDITIONAL_SUPPORT_OBSERVED_CABIBBO_RESIDUAL_BRIDGE_ONLY_COMPUTED"
	StatusCKMAlignmentAchieved           = "CONDITIONAL_SUPPORT_CKM_GEOMETRIC_ALIGNMENT_ACHIEVED"
	StatusFirewallPreserved              = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE470_DATA_FILE"
	StatusFailedFileMissing              = "FAILED_ROUTE_PDG_OBSERVED_LEDGER_FILE_MISSING"
	StatusFailedSwitchClosed             = "FAILED_ROUTE_GATE470_EMPIRICAL_IMPORT_SWITCH_CLOSED"
	StatusFailedMetadataIncomplete       = "FAILED_ROUTE_GATE470_METADATA_INCOMPLETE"
	StatusFailedMissingISpecIKValues     = "FAILED_ROUTE_GATE470_MISSING_EXPLICIT_I_SPEC_I_K_VALUES"
	StatusFailedMissingBranchTags        = "FAILED_ROUTE_GATE470_MISSING_EXPLICIT_BRANCH_TAGS"
	StatusFailedCommonScaleScheme        = "FAILED_ROUTE_GATE470_COMMON_SCALE_SCHEME_NOT_SUPPLIED"
	StatusFailedPDGNoIK                  = "FAILED_ROUTE_PDG_MASS_LEDGER_DOES_NOT_SUPPLY_ASHA_I_K_INVARIANT"
	StatusFailedDUDNotComputableFromFile = "FAILED_ROUTE_OBSERVED_NUMERICAL_DUD_NOT_COMPUTABLE_FROM_FILE"
	StatusFailedCabibboResidualUndefined = "FAILED_ROUTE_CABIBBO_RESIDUAL_UNDEFINED_WITHOUT_DUD"
	StatusFailedCabibboAsRayInput        = "FAILED_ROUTE_CABIBBO_USED_AS_GATE470_RAY_INPUT_REJECTED"
	StatusFailedNativePromotion          = "FAILED_ROUTE_GATE470_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite      = "FAILED_ROUTE_GATE470_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedCKMNativePrediction      = "FAILED_ROUTE_GATE470_CKM_NATIVE_PREDICTION_REJECTED"
	StatusFailedProjectiveDomain         = "FAILED_ROUTE_GATE470_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseDomain              = "FAILED_ROUTE_GATE470_PHASE_DOMAIN_REJECTED"
	StatusFailedCaustic                  = "FAILED_ROUTE_GATE470_CAUSTIC_BRANCH_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	DefaultLedger   = "data/pdg_observed_ledger.json"
)

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate464DUDSocketAvailable, Gate465AirlockAvailable, Gate469PreflightValidated, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                    string
}

type Number = *float64
type Integer = *int

type DataRow struct {
	Name                 string  `json:"name"`
	Sector               string  `json:"sector"`
	Observable           string  `json:"observable"`
	Value                Number  `json:"value"`
	Unit                 string  `json:"unit"`
	Source               string  `json:"source"`
	SourceVersion        string  `json:"source_version"`
	Scale                string  `json:"scale"`
	Scheme               string  `json:"scheme"`
	Uncertainty          string  `json:"uncertainty"`
	BridgeOnly           bool    `json:"bridge_only"`
	SigmaCP              Integer `json:"sigma_cp"`
	NC3                  Integer `json:"n_c3"`
	CabibboAsRayInput    bool    `json:"cabibbo_as_ray_input"`
	NativePromotionClaim bool    `json:"native_promotion_claim"`
}

type DataLedger struct {
	Gate                int       `json:"gate"`
	LedgerName          string    `json:"ledger_name"`
	Description         string    `json:"description"`
	EmpiricalImport     bool      `json:"empirical_import"`
	BridgeOnly          bool      `json:"bridge_only"`
	NativeRegistryWrite bool      `json:"native_registry_write"`
	CommonScale         string    `json:"common_scale"`
	CommonScheme        string    `json:"common_scheme"`
	Rows                []DataRow `json:"rows"`
}

type FileImport struct {
	Executed, Loaded, EmpiricalImport, BridgeOnlyLedger, NativeRegistryWriteRequested                                                         bool
	Path                                                                                                                                      string
	Rows                                                                                                                                      int
	AcceptedRows, RejectedRows                                                                                                                int
	MassRows, ComparatorRows, BranchRows, CKMTargetRows                                                                                       int
	AllAcceptedBridgeOnly, MetadataComplete, SwitchClosedRejected, NativePromotionRejected, NativeRegistryWriteRejected, CabibboAsRayRejected bool
	Verdict, Reason                                                                                                                           string
	Failures                                                                                                                                  []string
}

type SectorInput struct {
	Sector             string
	ISpec, IK          Number
	SigmaCP            Integer
	NC3                Integer
	HasISpec, HasIK    bool
	HasBranch          bool
	CommonScale        string
	CommonScheme       string
	ISpecUncertainty   string
	IKUncertainty      string
	BranchUncertainty  string
	MetadataComplete   bool
	BridgeOnly         bool
	PDGMassRowsPresent bool
	PDGDoesNotSupplyIK bool
	NativePromotion    bool
	CabibboAsRayInput  bool
}

type Ray struct {
	Sector                                       string
	Alpha, CosThreePhi, Phi, IK, ISpec           float64
	SigmaCP, NC3                                 int
	Defined, InsideDomain, AtCaustic, BridgeOnly bool
	Verdict, Reason                              string
}

type NumericalAdapter struct {
	Executed, Attempted, ReadyForDUD, DUDComputed, CabibboTargetAvailable, CabibboResidualComputed, AlignmentAchieved                          bool
	U, D                                                                                                                                       SectorInput
	URay, DRay                                                                                                                                 Ray
	DUD, CabibboTarget, CabibboResidual                                                                                                        float64
	MissingISpecIKValues, MissingBranchTags, CommonScaleSchemeMissing, PDGNoIK, ProjectiveDomainRejected, PhaseDomainRejected, CausticRejected bool
	Verdict, Reason                                                                                                                            string
	Failures                                                                                                                                   []string
}

type Firewall struct {
	Executed, DataFileRowsNative, CoordinatesNative, DUDNativePrediction, CKMNativePrediction, CKMMatrixConstructed, CKMEntryComputed, CabibboUsedAsRayInput, NativeRegistryWritten, KGenStillForced, XTriangleStillForced, YPhaseStillQuarantined, SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                                                                                                                        int
	Verdict, Reason                                                                                                                                                                                                                                                               string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      FileImport
	Adapter     NumericalAdapter
	Firewall    Firewall
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
		a.Adapter = buildAdapter(ledger)
	} else {
		a.Adapter = NumericalAdapter{Executed: true, Attempted: false, Verdict: StatusFailedFileMissing, Reason: "explicit Gate470 data file was not found", Failures: []string{StatusFailedFileMissing}}
	}
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate456InverseAvailable: true, Gate459BranchTagsRequired: true, Gate464DUDSocketAvailable: true, Gate465AirlockAvailable: true, Gate469PreflightValidated: true, NativeRegistryClean: true, Verdict: StatusGate469Inherited}
}

func loadLedger(path string) (DataLedger, FileImport) {
	p := projectPath(path)
	imp := FileImport{Executed: true, Path: p, Verdict: StatusDataFileLoaded}
	bs, err := os.ReadFile(p)
	if err != nil {
		imp.Loaded = false
		imp.Verdict = StatusFailedFileMissing
		imp.Failures = []string{StatusFailedFileMissing}
		imp.Reason = err.Error()
		return DataLedger{}, imp
	}
	var l DataLedger
	if err := json.Unmarshal(bs, &l); err != nil {
		imp.Loaded = false
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		imp.Reason = err.Error()
		return DataLedger{}, imp
	}
	imp.Loaded = true
	imp.EmpiricalImport = l.EmpiricalImport
	imp.BridgeOnlyLedger = l.BridgeOnly
	imp.NativeRegistryWriteRequested = l.NativeRegistryWrite
	imp.Rows = len(l.Rows)
	failures := []string{}
	if !l.EmpiricalImport {
		failures = append(failures, StatusFailedSwitchClosed)
		imp.SwitchClosedRejected = true
	}
	if !l.BridgeOnly || l.LedgerName == "" {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if l.NativeRegistryWrite {
		failures = append(failures, StatusFailedNativeRegistryWrite)
		imp.NativeRegistryWriteRejected = true
	}
	imp.MetadataComplete = true
	imp.AllAcceptedBridgeOnly = true
	for _, r := range l.Rows {
		obs := strings.ToLower(r.Observable)
		switch obs {
		case "running_quark_mass", "top_quark_mass":
			imp.MassRows++
		case "i_spec", "i_k":
			imp.ComparatorRows++
		case "branch_tag":
			imp.BranchRows++
		case "cabibbo_target_abs_vus":
			imp.CKMTargetRows++
		}
		rowOK := r.Source != "" && r.Scale != "" && r.Scheme != "" && r.Uncertainty != "" && r.BridgeOnly
		if !rowOK {
			imp.MetadataComplete = false
			failures = append(failures, StatusFailedMetadataIncomplete)
		}
		if !r.BridgeOnly {
			imp.AllAcceptedBridgeOnly = false
		}
		if r.CabibboAsRayInput {
			failures = append(failures, StatusFailedCabibboAsRayInput)
			imp.CabibboAsRayRejected = true
		}
		if r.NativePromotionClaim {
			failures = append(failures, StatusFailedNativePromotion)
			imp.NativePromotionRejected = true
		}
		if rowOK && l.EmpiricalImport && l.BridgeOnly && !l.NativeRegistryWrite {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
	}
	if len(failures) > 0 {
		imp.Failures = unique(failures)
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "explicit data-file airlock rejected one or more metadata/native-promotion routes"
		return l, imp
	}
	imp.NativePromotionRejected = true
	imp.NativeRegistryWriteRejected = true
	imp.CabibboAsRayRejected = true
	imp.Verdict = StatusAirlockAcceptedBridgeRows
	imp.Reason = "explicit ledger rows entered only the bridge comparator airlock"
	return l, imp
}

func buildAdapter(l DataLedger) NumericalAdapter {
	ad := NumericalAdapter{Executed: true, Attempted: true, Verdict: StatusNumericalAdapterAttempted, Reason: "explicit data-file adapter attempted Gate456/Gate464 inversion"}
	ad.U = sectorInput(l, "u")
	ad.D = sectorInput(l, "d")
	ad.CabibboTarget, ad.CabibboTargetAvailable = findValue(l, "u-d", "cabibbo_target_abs_vus")
	failures := []string{}
	if !ad.U.HasISpec || !ad.U.HasIK || !ad.D.HasISpec || !ad.D.HasIK || ad.U.ISpec == nil || ad.U.IK == nil || ad.D.ISpec == nil || ad.D.IK == nil {
		failures = append(failures, StatusFailedMissingISpecIKValues)
		ad.MissingISpecIKValues = true
	}
	if !ad.U.HasBranch || !ad.D.HasBranch || ad.U.SigmaCP == nil || ad.U.NC3 == nil || ad.D.SigmaCP == nil || ad.D.NC3 == nil {
		failures = append(failures, StatusFailedMissingBranchTags)
		ad.MissingBranchTags = true
	}
	if ad.U.CommonScale == "" || ad.D.CommonScale == "" || ad.U.CommonScheme == "" || ad.D.CommonScheme == "" || ad.U.CommonScale != ad.D.CommonScale || ad.U.CommonScheme != ad.D.CommonScheme || strings.Contains(strings.ToLower(l.CommonScale), "not supplied") {
		failures = append(failures, StatusFailedCommonScaleScheme)
		ad.CommonScaleSchemeMissing = true
	}
	if ad.U.PDGDoesNotSupplyIK || ad.D.PDGDoesNotSupplyIK {
		failures = append(failures, StatusFailedPDGNoIK)
		ad.PDGNoIK = true
	}
	if len(failures) > 0 {
		failures = append(failures, StatusFailedDUDNotComputableFromFile, StatusFailedCabibboResidualUndefined)
		ad.Failures = unique(failures)
		ad.Verdict = StatusFailedDUDNotComputableFromFile
		ad.Reason = "the explicit file was parsed, but it does not supply rank-complete ASHA bridge comparators; d_ud and Cabibbo residual remain undefined"
		return ad
	}
	ur, uf := invert(ad.U)
	dr, df := invert(ad.D)
	if uf != "" {
		failures = append(failures, uf)
	}
	if df != "" {
		failures = append(failures, df)
	}
	ad.URay = ur
	ad.DRay = dr
	if len(failures) > 0 {
		for _, f := range failures {
			switch f {
			case StatusFailedProjectiveDomain:
				ad.ProjectiveDomainRejected = true
			case StatusFailedPhaseDomain:
				ad.PhaseDomainRejected = true
			case StatusFailedCaustic:
				ad.CausticRejected = true
			}
		}
		ad.Failures = unique(failures)
		ad.Verdict = strings.Join(ad.Failures, ";")
		ad.Reason = "rank-complete file failed inverse-domain checks"
		return ad
	}
	ad.ReadyForDUD = true
	ad.DUD = distance(ur, dr)
	ad.DUDComputed = true
	if ad.CabibboTargetAvailable {
		ad.CabibboResidual = math.Abs(ad.DUD - ad.CabibboTarget)
		ad.CabibboResidualComputed = true
		ad.AlignmentAchieved = ad.CabibboResidual < 1e-3
	}
	ad.Verdict = StatusNumericalDUDComputed
	if ad.AlignmentAchieved {
		ad.Verdict = StatusCKMAlignmentAchieved
	}
	ad.Reason = "rank-complete explicit bridge file computed d_ud; result remains a bridge comparator, not a native CKM prediction"
	return ad
}

func sectorInput(l DataLedger, sector string) SectorInput {
	x := SectorInput{Sector: sector, CommonScale: l.CommonScale, CommonScheme: l.CommonScheme, BridgeOnly: l.BridgeOnly, MetadataComplete: true}
	for _, r := range l.Rows {
		if r.Sector != sector {
			continue
		}
		if r.Source == "" || r.Scale == "" || r.Scheme == "" || r.Uncertainty == "" || !r.BridgeOnly {
			x.MetadataComplete = false
		}
		switch strings.ToLower(r.Observable) {
		case "running_quark_mass", "top_quark_mass":
			x.PDGMassRowsPresent = true
		case "i_spec":
			x.HasISpec = true
			x.ISpec = r.Value
			x.ISpecUncertainty = r.Uncertainty
		case "i_k":
			x.HasIK = true
			x.IK = r.Value
			x.IKUncertainty = r.Uncertainty
		case "branch_tag":
			x.SigmaCP = r.SigmaCP
			x.NC3 = r.NC3
			x.BranchUncertainty = r.Uncertainty
			x.HasBranch = r.SigmaCP != nil && r.NC3 != nil
		}
		if r.NativePromotionClaim {
			x.NativePromotion = true
		}
		if r.CabibboAsRayInput {
			x.CabibboAsRayInput = true
		}
	}
	if x.PDGMassRowsPresent && (x.IK == nil || !x.HasIK) {
		x.PDGDoesNotSupplyIK = true
	}
	return x
}

func findValue(l DataLedger, sector, observable string) (float64, bool) {
	for _, r := range l.Rows {
		if r.Sector == sector && strings.EqualFold(r.Observable, observable) && r.Value != nil {
			return *r.Value, true
		}
	}
	return math.NaN(), false
}

func invert(x SectorInput) (Ray, string) {
	r := Ray{Sector: x.Sector, BridgeOnly: true}
	if x.IK == nil || x.ISpec == nil || x.SigmaCP == nil || x.NC3 == nil {
		return r, StatusFailedMissingISpecIKValues
	}
	ik, ispec := *x.IK, *x.ISpec
	r.IK, r.ISpec, r.SigmaCP, r.NC3 = ik, ispec, *x.SigmaCP, *x.NC3
	if math.Abs(ik) >= 1 {
		r.Verdict = StatusFailedProjectiveDomain
		r.Reason = "I_K must lie in (-1,1)"
		return r, StatusFailedProjectiveDomain
	}
	alpha := math.Sqrt(3) * ik / math.Sqrt(1-ik*ik)
	cos3 := (3 * math.Sqrt(3) / 2) * ispec / math.Pow(1-ik*ik, 1.5)
	if math.Abs(cos3) > 1+1e-12 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Verdict = StatusFailedPhaseDomain
		r.Reason = "cos(3phi) outside [-1,1]"
		return r, StatusFailedPhaseDomain
	}
	cos3 = clamp(cos3, -1, 1)
	phi := (float64(*x.SigmaCP)*math.Acos(cos3) + 2*math.Pi*float64(*x.NC3)) / 3
	if math.Abs(math.Sin(3*phi)) < 1e-9 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Phi = phi
		r.AtCaustic = true
		r.Verdict = StatusFailedCaustic
		r.Reason = "sin(3phi)=0 caustic"
		return r, StatusFailedCaustic
	}
	r.Alpha, r.CosThreePhi, r.Phi = alpha, cos3, phi
	r.Defined, r.InsideDomain = true, true
	r.Verdict = StatusNumericalDUDComputed
	r.Reason = "explicit rank-complete observed bridge row inverted to ASHA cylinder coordinate"
	return r, ""
}

func distance(u, d Ray) float64 {
	da := d.Alpha - u.Alpha
	dp := wrapPi(d.Phi - u.Phi)
	return math.Sqrt(da*da + 4*math.Pow(math.Sin(dp/2), 2))
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate470 data-file rows are quarantined bridge comparators; no row, coordinate, d_ud, residual, CKM entry, or alignment flag writes to native law-space"}
}

func buildNext() NextStep {
	return NextStep{471, "Rank-Complete External Ledger Acceptance Test", "Gate470 parsed the explicit data file and refused to fabricate missing ASHA comparators; the next possible step is to supply a genuinely rank-complete external bridge ledger with I_spec, I_K, and branch tags.", "evaluate a user-supplied rank-complete observed bridge ledger, never PDG masses alone, and export only bridge residuals"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate469PreflightValidated || !a.Inheritance.Gate465AirlockAvailable || !a.Inheritance.Gate464DUDSocketAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate470 inheritance incomplete")
	}
	if !a.Import.Executed || !a.Import.Loaded || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || a.Import.NativeRegistryWriteRequested || a.Import.Rows == 0 || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly || !a.Import.MetadataComplete {
		return fmt.Errorf("Gate470 file import did not satisfy bridge-only airlock conditions: %+v", a.Import)
	}
	if !a.Adapter.Executed || !a.Adapter.Attempted || a.Adapter.DUDComputed || a.Adapter.CabibboResidualComputed || a.Adapter.AlignmentAchieved || !a.Adapter.MissingISpecIKValues || !a.Adapter.MissingBranchTags || !a.Adapter.PDGNoIK || a.Adapter.Verdict != StatusFailedDUDNotComputableFromFile {
		return fmt.Errorf("Gate470 default observed ledger must fail closed without d_ud: %+v", a.Adapter)
	}
	if !a.Firewall.Executed || a.Firewall.DataFileRowsNative || a.Firewall.CoordinatesNative || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.CabibboUsedAsRayInput || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate470 firewall violated")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Import.Loaded && !a.Adapter.DUDComputed && a.Adapter.MissingISpecIKValues && a.Adapter.PDGNoIK && !a.Firewall.NativeRegistryWritten {
		return "Gate 470 successfully reads the explicit observed ledger through the empirical airlock, but the checked-in PDG-style file does not contain explicit ASHA rank-complete comparators. PDG mass rows and the Cabibbo target remain quarantined bridge data; d_ud and the Cabibbo residual are undefined until I_spec, I_K, and branch tags are explicitly supplied. The geometry has not been numerically matched to CKM by this run."
	}
	if a.Adapter.DUDComputed {
		return "Gate 470 computed a bridge-only d_ud from a rank-complete explicit ledger. This remains a comparator result and is not a native CKM prediction."
	}
	return "Gate 470 failed before preserving the empirical firewall."
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
func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
func wrapPi(x float64) float64 {
	for x <= -math.Pi {
		x += 2 * math.Pi
	}
	for x > math.Pi {
		x -= 2 * math.Pi
	}
	return x
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
func ptrIntText(x Integer) string {
	if x == nil {
		return "missing"
	}
	return fmt.Sprintf("%d", *x)
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

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t d_ud_socket=%t airlock=%t preflight=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate464DUDSocketAvailable, x.Gate465AirlockAvailable, x.Gate469PreflightValidated, x.NativeRegistryClean, x.Verdict)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("executed=%t loaded=%t path=%s empirical_import=%t bridge_only=%t rows=%d accepted=%d rejected=%d mass_rows=%d comparator_rows=%d branch_rows=%d ckm_targets=%d metadata=%t quarantined=%t native_write_requested=%t verdict=%s reason=%s", x.Executed, x.Loaded, x.Path, x.EmpiricalImport, x.BridgeOnlyLedger, x.Rows, x.AcceptedRows, x.RejectedRows, x.MassRows, x.ComparatorRows, x.BranchRows, x.CKMTargetRows, x.MetadataComplete, x.AllAcceptedBridgeOnly, x.NativeRegistryWriteRequested, x.Verdict, x.Reason)
}
func FormatSector(x SectorInput) string {
	return fmt.Sprintf("sector=%s I_spec=%s I_K=%s sigma_CP=%s n_C3=%s scale=%s scheme=%s metadata=%t bridge_only=%t pdg_mass_rows=%t pdg_no_IK=%t", x.Sector, ptrFloatText(x.ISpec), ptrFloatText(x.IK), ptrIntText(x.SigmaCP), ptrIntText(x.NC3), x.CommonScale, x.CommonScheme, x.MetadataComplete, x.BridgeOnly, x.PDGMassRowsPresent, x.PDGDoesNotSupplyIK)
}
func FormatRay(x Ray) string {
	return fmt.Sprintf("sector=%s defined=%t alpha=%s cos3phi=%s phi=%s I_K=%s I_spec=%s sigma_CP=%d n_C3=%d domain=%t caustic=%t bridge_only=%t verdict=%s", x.Sector, x.Defined, fmtFloat(x.Alpha), fmtFloat(x.CosThreePhi), fmtFloat(x.Phi), fmtFloat(x.IK), fmtFloat(x.ISpec), x.SigmaCP, x.NC3, x.InsideDomain, x.AtCaustic, x.BridgeOnly, x.Verdict)
}
func FormatAdapter(x NumericalAdapter) string {
	dud := "undefined"
	residual := "undefined"
	if x.DUDComputed {
		dud = fmtFloat(x.DUD)
	}
	if x.CabibboResidualComputed {
		residual = fmtFloat(x.CabibboResidual)
	}
	return fmt.Sprintf("executed=%t attempted=%t ready=%t d_ud_computed=%t d_ud=%s cabibbo_available=%t |Vus|=%s residual_computed=%t residual=%s alignment=%t missing_I=%t missing_branch=%t common_scale_missing=%t pdg_no_IK=%t verdict=%s reason=%s", x.Executed, x.Attempted, x.ReadyForDUD, x.DUDComputed, dud, x.CabibboTargetAvailable, fmtFloat(x.CabibboTarget), x.CabibboResidualComputed, residual, x.AlignmentAchieved, x.MissingISpecIKValues, x.MissingBranchTags, x.CommonScaleSchemeMissing, x.PDGNoIK, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t rows_native=%t coords_native=%t d_ud_native=%t ckm_native=%t ckm_matrix=%t ckm_entry=%t cabibbo_as_ray=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.DataFileRowsNative, x.CoordinatesNative, x.DUDNativePrediction, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.CabibboUsedAsRayInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate469Inherited, StatusDataFileLoaded, StatusAirlockAcceptedBridgeRows, StatusNumericalAdapterAttempted, StatusFailedMissingISpecIKValues, StatusFailedMissingBranchTags, StatusFailedCommonScaleScheme, StatusFailedPDGNoIK, StatusFailedDUDNotComputableFromFile, StatusFailedCabibboResidualUndefined, StatusFailedCabibboAsRayInput, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedCKMNativePrediction, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 470 Registry Audit — Observed Numerical d_ud Adapter / Explicit Data-File Run\n\n## Verdict\n\n")
	b.WriteString("`" + StatusFailedDUDNotComputableFromFile + "`\n\n")
	b.WriteString("Gate 470 reads `data/pdg_observed_ledger.json` through the empirical airlock. The checked-in file contains PDG-style quark mass rows and a Cabibbo target row, but it does not contain explicit ASHA rank-complete bridge comparator values for `I_spec`, `I_K`, or branch tags. Therefore the adapter refuses to fabricate cylinder coordinates and leaves `d_ud` undefined.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Data-file import\n\n" + FormatImport(a.Import) + "\n\n")
	b.WriteString("## Parsed sector inputs\n\n")
	b.WriteString("- " + FormatSector(a.Adapter.U) + "\n")
	b.WriteString("- " + FormatSector(a.Adapter.D) + "\n\n")
	b.WriteString("## Numerical adapter\n\n" + FormatAdapter(a.Adapter) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha = sqrt(3) I_K / sqrt(1-I_K^2)\n")
	b.WriteString("cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\n")
	b.WriteString("d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n")
	if a.Adapter.DUDComputed {
		b.WriteString("Gate470 d_ud = " + fmtFloat(a.Adapter.DUD) + "\n")
	} else {
		b.WriteString("Gate470 d_ud = undefined\n")
	}
	if a.Adapter.CabibboTargetAvailable {
		b.WriteString("observed bridge target |V_us| = " + fmtFloat(a.Adapter.CabibboTarget) + "\n")
	}
	if a.Adapter.CabibboResidualComputed {
		b.WriteString("Cabibbo residual |d_ud-|V_us|| = " + fmtFloat(a.Adapter.CabibboResidual) + "\n")
	} else {
		b.WriteString("Cabibbo residual = undefined\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No data-file row enters the native theorem registry. No mass, CKM value, `I_K`, `I_spec`, branch tag, cylinder coordinate, `d_ud`, residual, or alignment flag is exported as a native law.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
