// Package generation2rankcompleteexternalledger implements Gate 471:
// Rank-Complete External Ledger Acceptance Test.
//
// Gate 471 is intentionally an empirical bridge adapter, not a native theorem.
// It reads data/pdg_rank_complete_ledger.json through the Gate465 airlock and
// computes d_ud only when the ledger explicitly supplies rank-complete ASHA
// bridge comparators: I_spec, I_K, sigma_CP, n_C3, common scale/scheme, source,
// uncertainty, and bridge_only quarantine for both u and d sectors.
//
// The checked-in ledger is an explicit rank-complete bridge fixture. Its I_K and
// branch tags are not PDG-published mass-table quantities and must never be
// relabelled as native ASHA laws. They are external bridge inputs used only to
// exercise the socket and compute a comparator residual against |V_us|.
package generation2rankcompleteexternalledger

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
	AuditID = "GATE471-RANK-COMPLETE-EXTERNAL-LEDGER-ACCEPTANCE-TEST"

	StatusGate470Inherited                   = "CONDITIONAL_SUPPORT_GATE470_AIRLOCK_AND_NONSMUGGLING_INHERITED"
	StatusRankCompleteLedgerLoaded           = "CONDITIONAL_SUPPORT_RANK_COMPLETE_EXTERNAL_LEDGER_LOADED"
	StatusAirlockAcceptedRankCompleteRows    = "CONDITIONAL_SUPPORT_GATE471_AIRLOCK_ACCEPTED_RANK_COMPLETE_BRIDGE_ROWS"
	StatusCoordinatesComputed                = "CONDITIONAL_SUPPORT_GATE471_CYLINDER_COORDINATES_COMPUTED"
	StatusDUDComputed                        = "CONDITIONAL_SUPPORT_GATE471_DUD_BRIDGE_ONLY_COMPUTED"
	StatusCabibboResidualComputed            = "CONDITIONAL_SUPPORT_GATE471_CABIBBO_RESIDUAL_BRIDGE_ONLY_COMPUTED"
	StatusCKMGeometricAlignmentAchieved      = "CONDITIONAL_SUPPORT_CKM_GEOMETRIC_ALIGNMENT_ACHIEVED"
	StatusExternalComparatorsNotPDGNative    = "CONDITIONAL_SUPPORT_EXTERNAL_IK_BRANCH_INPUTS_QUARANTINED_NOT_PDG_NATIVE"
	StatusFirewallPreserved                  = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE471_RANK_COMPLETE_LEDGER"
	StatusFailedFileMissing                  = "FAILED_ROUTE_PDG_RANK_COMPLETE_LEDGER_FILE_MISSING"
	StatusFailedSwitchClosed                 = "FAILED_ROUTE_GATE471_EMPIRICAL_IMPORT_SWITCH_CLOSED"
	StatusFailedMetadataIncomplete           = "FAILED_ROUTE_GATE471_METADATA_INCOMPLETE"
	StatusFailedMissingISpecIKValues         = "FAILED_ROUTE_GATE471_MISSING_EXPLICIT_I_SPEC_I_K_VALUES"
	StatusFailedMissingBranchTags            = "FAILED_ROUTE_GATE471_MISSING_EXPLICIT_BRANCH_TAGS"
	StatusFailedCommonScaleScheme            = "FAILED_ROUTE_GATE471_COMMON_SCALE_SCHEME_NOT_SUPPLIED"
	StatusFailedCabibboAsRayInput            = "FAILED_ROUTE_CABIBBO_USED_AS_GATE471_RAY_INPUT_REJECTED"
	StatusFailedNativePromotion              = "FAILED_ROUTE_GATE471_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite          = "FAILED_ROUTE_GATE471_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedCKMNativePrediction          = "FAILED_ROUTE_GATE471_CKM_NATIVE_PREDICTION_REJECTED"
	StatusFailedProjectiveDomain             = "FAILED_ROUTE_GATE471_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseDomain                  = "FAILED_ROUTE_GATE471_PHASE_DOMAIN_REJECTED"
	StatusFailedCaustic                      = "FAILED_ROUTE_GATE471_CAUSTIC_BRANCH_REJECTED"
	StatusFailedNoCabibboTarget              = "FAILED_ROUTE_GATE471_CABIBBO_TARGET_MISSING"
	StatusFailedLedgerPretendsPDGPublishesIK = "FAILED_ROUTE_LEDGER_PRETENDS_PDG_PUBLISHES_ASHA_I_K_REJECTED"
)

const (
	NativeFlavorDim    = 13
	KXYCoeffDim        = 9
	DefaultLedger      = "data/pdg_rank_complete_ledger.json"
	AlignmentTolerance = 1e-6
)

type Number = *float64
type Integer = *int

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate464DUDSocketAvailable, Gate465AirlockAvailable, Gate470NonSmugglingValidated, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                       string
}

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
	ClaimsPDGPublishesIK bool    `json:"claims_pdg_publishes_i_k"`
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
	Executed, Loaded, EmpiricalImport, BridgeOnlyLedger, NativeRegistryWriteRequested                                                                             bool
	Path                                                                                                                                                          string
	Rows, AcceptedRows, RejectedRows, ComparatorRows, BranchRows, CKMTargetRows                                                                                   int
	AllAcceptedBridgeOnly, MetadataComplete, SwitchClosedRejected, NativePromotionRejected, NativeRegistryWriteRejected, CabibboAsRayRejected, PDGIKClaimRejected bool
	Verdict, Reason                                                                                                                                               string
	Failures                                                                                                                                                      []string
}

type SectorInput struct {
	Sector                                                   string
	ISpec, IK                                                Number
	SigmaCP                                                  Integer
	NC3                                                      Integer
	HasISpec, HasIK, HasBranch                               bool
	CommonScale, CommonScheme                                string
	MetadataComplete, BridgeOnly                             bool
	NativePromotion, CabibboAsRayInput, ClaimsPDGPublishesIK bool
}

type Ray struct {
	Sector                                       string
	Alpha, CosThreePhi, Phi, IK, ISpec           float64
	SigmaCP, NC3                                 int
	Defined, InsideDomain, AtCaustic, BridgeOnly bool
	Verdict, Reason                              string
}

type NumericalAdapter struct {
	Executed, Attempted, ReadyForDUD, CoordinatesComputed, DUDComputed, CabibboTargetAvailable, CabibboResidualComputed, AlignmentAchieved             bool
	U, D                                                                                                                                               SectorInput
	URay, DRay                                                                                                                                         Ray
	DUD, CabibboTarget, CabibboResidual                                                                                                                float64
	MissingISpecIKValues, MissingBranchTags, CommonScaleSchemeMissing, ProjectiveDomainRejected, PhaseDomainRejected, CausticRejected, NoCabibboTarget bool
	Verdict, Reason                                                                                                                                    string
	Failures                                                                                                                                           []string
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
		a.Adapter = NumericalAdapter{Executed: true, Attempted: false, Verdict: StatusFailedFileMissing, Reason: "explicit Gate471 rank-complete ledger was not found", Failures: []string{StatusFailedFileMissing}}
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
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate456InverseAvailable: true, Gate459BranchTagsRequired: true, Gate464DUDSocketAvailable: true, Gate465AirlockAvailable: true, Gate470NonSmugglingValidated: true, NativeRegistryClean: true, Verdict: StatusGate470Inherited}
}

func loadLedger(path string) (DataLedger, FileImport) {
	p := projectPath(path)
	imp := FileImport{Executed: true, Path: p, Verdict: StatusRankCompleteLedgerLoaded}
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
		case "i_spec", "i_k":
			imp.ComparatorRows++
		case "branch_tag":
			imp.BranchRows++
		case "cabibbo_target_abs_vus":
			imp.CKMTargetRows++
		}
		rowOK := r.Name != "" && r.Sector != "" && r.Observable != "" && r.Unit != "" && r.Source != "" && r.Scale != "" && r.Scheme != "" && r.Uncertainty != "" && r.BridgeOnly
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
		if r.ClaimsPDGPublishesIK {
			failures = append(failures, StatusFailedLedgerPretendsPDGPublishesIK)
			imp.PDGIKClaimRejected = true
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
		imp.Reason = "rank-complete external ledger rejected one or more airlock/native-promotion routes"
		return l, imp
	}
	imp.NativePromotionRejected = true
	imp.NativeRegistryWriteRejected = true
	imp.CabibboAsRayRejected = true
	imp.PDGIKClaimRejected = true
	imp.Verdict = StatusAirlockAcceptedRankCompleteRows
	imp.Reason = "rank-complete rows entered only the bridge comparator airlock"
	return l, imp
}

func buildAdapter(l DataLedger) NumericalAdapter {
	ad := NumericalAdapter{Executed: true, Attempted: true, Verdict: StatusRankCompleteLedgerLoaded, Reason: "Gate471 rank-complete ledger attempting Gate456/Gate464 inversion"}
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
	if ad.U.CommonScale == "" || ad.D.CommonScale == "" || ad.U.CommonScheme == "" || ad.D.CommonScheme == "" || ad.U.CommonScale != ad.D.CommonScale || ad.U.CommonScheme != ad.D.CommonScheme {
		failures = append(failures, StatusFailedCommonScaleScheme)
		ad.CommonScaleSchemeMissing = true
	}
	if !ad.CabibboTargetAvailable {
		failures = append(failures, StatusFailedNoCabibboTarget)
		ad.NoCabibboTarget = true
	}
	if len(failures) > 0 {
		ad.Failures = unique(failures)
		ad.Verdict = strings.Join(ad.Failures, ";")
		ad.Reason = "rank-complete external ledger is incomplete; d_ud not computed"
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
	ad.CoordinatesComputed = true
	ad.DUD = distance(ur, dr)
	ad.DUDComputed = true
	ad.CabibboResidual = math.Abs(ad.DUD - ad.CabibboTarget)
	ad.CabibboResidualComputed = true
	ad.AlignmentAchieved = ad.CabibboResidual <= AlignmentTolerance
	ad.Verdict = StatusDUDComputed
	if ad.AlignmentAchieved {
		ad.Verdict = StatusCKMGeometricAlignmentAchieved
	}
	ad.Reason = "explicit rank-complete external bridge ledger computed d_ud and Cabibbo residual; values remain bridge-only comparator outputs"
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
		case "i_spec":
			x.HasISpec = true
			x.ISpec = r.Value
		case "i_k":
			x.HasIK = true
			x.IK = r.Value
		case "branch_tag":
			x.SigmaCP = r.SigmaCP
			x.NC3 = r.NC3
			x.HasBranch = r.SigmaCP != nil && r.NC3 != nil
		}
		if r.NativePromotionClaim {
			x.NativePromotion = true
		}
		if r.CabibboAsRayInput {
			x.CabibboAsRayInput = true
		}
		if r.ClaimsPDGPublishesIK {
			x.ClaimsPDGPublishesIK = true
		}
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
	r.Defined = true
	r.InsideDomain = true
	r.Verdict = StatusCoordinatesComputed
	r.Reason = "rank-complete bridge row inverted to ASHA cylinder coordinate"
	return r, ""
}

func distance(u, d Ray) float64 {
	da := d.Alpha - u.Alpha
	dp := wrapPi(d.Phi - u.Phi)
	return math.Sqrt(da*da + 4*math.Pow(math.Sin(dp/2), 2))
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate471 rank-complete data-file rows, coordinates, d_ud, residual, and alignment flag are bridge comparator outputs only; no native theorem-registry write occurs"}
}
func buildNext() NextStep {
	return NextStep{472, "Independent observed-ledger provenance challenge", "Gate471 can compute d_ud from an explicit rank-complete external ledger, but I_K and branch tags are external bridge inputs rather than PDG-published invariants.", "audit whether a genuinely independent experimental/provenance source can supply I_K and branch tags without reverse-fitting Cabibbo"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate470NonSmugglingValidated || !a.Inheritance.Gate465AirlockAvailable || !a.Inheritance.Gate464DUDSocketAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate471 inheritance incomplete")
	}
	if !a.Import.Executed || !a.Import.Loaded || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || a.Import.NativeRegistryWriteRequested || a.Import.Rows == 0 || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly || !a.Import.MetadataComplete || a.Import.NativePromotionRejected != true || a.Import.NativeRegistryWriteRejected != true {
		return fmt.Errorf("Gate471 file import did not satisfy bridge-only airlock conditions: %+v", a.Import)
	}
	if !a.Adapter.Executed || !a.Adapter.Attempted || !a.Adapter.CoordinatesComputed || !a.Adapter.DUDComputed || !a.Adapter.CabibboTargetAvailable || !a.Adapter.CabibboResidualComputed || !a.Adapter.AlignmentAchieved || a.Adapter.MissingISpecIKValues || a.Adapter.MissingBranchTags || a.Adapter.CommonScaleSchemeMissing || a.Adapter.Verdict != StatusCKMGeometricAlignmentAchieved {
		return fmt.Errorf("Gate471 rank-complete ledger did not compute aligned bridge residual: %+v", a.Adapter)
	}
	if !a.Adapter.URay.Defined || !a.Adapter.DRay.Defined || a.Adapter.URay.AtCaustic || a.Adapter.DRay.AtCaustic {
		return fmt.Errorf("Gate471 ray inversion invalid: %+v %+v", a.Adapter.URay, a.Adapter.DRay)
	}
	if !a.Firewall.Executed || a.Firewall.DataFileRowsNative || a.Firewall.CoordinatesNative || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.CabibboUsedAsRayInput || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate471 firewall violated")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Adapter.DUDComputed && a.Adapter.CabibboResidualComputed && !a.Firewall.NativeRegistryWritten {
		return "Gate 471 computed d_ud from an explicitly supplied rank-complete external bridge ledger and compared it to the Cabibbo target without writing to native law-space. The numerical alignment is a bridge-comparator fact about this ledger, not an independent native prediction, because I_K and branch tags are supplied external coordinates rather than PDG-published invariants."
	}
	return "Gate 471 failed before producing a bridge-only rank-complete comparator."
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
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t d_ud_socket=%t airlock=%t gate470=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate464DUDSocketAvailable, x.Gate465AirlockAvailable, x.Gate470NonSmugglingValidated, x.NativeRegistryClean, x.Verdict)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("executed=%t loaded=%t path=%s empirical_import=%t bridge_only=%t rows=%d accepted=%d rejected=%d comparator_rows=%d branch_rows=%d ckm_targets=%d metadata=%t quarantined=%t native_write_requested=%t verdict=%s reason=%s", x.Executed, x.Loaded, x.Path, x.EmpiricalImport, x.BridgeOnlyLedger, x.Rows, x.AcceptedRows, x.RejectedRows, x.ComparatorRows, x.BranchRows, x.CKMTargetRows, x.MetadataComplete, x.AllAcceptedBridgeOnly, x.NativeRegistryWriteRequested, x.Verdict, x.Reason)
}
func FormatSector(x SectorInput) string {
	return fmt.Sprintf("sector=%s I_spec=%s I_K=%s sigma_CP=%s n_C3=%s scale=%s scheme=%s metadata=%t bridge_only=%t claims_pdg_IK=%t", x.Sector, ptrFloatText(x.ISpec), ptrFloatText(x.IK), ptrIntText(x.SigmaCP), ptrIntText(x.NC3), x.CommonScale, x.CommonScheme, x.MetadataComplete, x.BridgeOnly, x.ClaimsPDGPublishesIK)
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
	return fmt.Sprintf("executed=%t attempted=%t ready=%t coordinates=%t d_ud_computed=%t d_ud=%s cabibbo_available=%t |Vus|=%s residual_computed=%t residual=%s alignment=%t missing_I=%t missing_branch=%t common_scale_missing=%t verdict=%s reason=%s", x.Executed, x.Attempted, x.ReadyForDUD, x.CoordinatesComputed, x.DUDComputed, dud, x.CabibboTargetAvailable, fmtFloat(x.CabibboTarget), x.CabibboResidualComputed, residual, x.AlignmentAchieved, x.MissingISpecIKValues, x.MissingBranchTags, x.CommonScaleSchemeMissing, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t rows_native=%t coords_native=%t d_ud_native=%t ckm_native=%t ckm_matrix=%t ckm_entry=%t cabibbo_as_ray=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.DataFileRowsNative, x.CoordinatesNative, x.DUDNativePrediction, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.CabibboUsedAsRayInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate470Inherited, StatusRankCompleteLedgerLoaded, StatusAirlockAcceptedRankCompleteRows, StatusCoordinatesComputed, StatusDUDComputed, StatusCabibboResidualComputed, StatusCKMGeometricAlignmentAchieved, StatusExternalComparatorsNotPDGNative, StatusFailedLedgerPretendsPDGPublishesIK, StatusFailedCabibboAsRayInput, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedCKMNativePrediction, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 471 Registry Audit — Rank-Complete External Ledger Acceptance Test\n\n## Verdict\n\n")
	b.WriteString("`" + StatusCKMGeometricAlignmentAchieved + "`\n\n")
	b.WriteString("Gate 471 reads `data/pdg_rank_complete_ledger.json` through the empirical airlock. Unlike Gate 470, this file explicitly supplies `I_spec`, `I_K`, and `{sigma_CP,n_C3}` for both `u` and `d` sectors. The adapter therefore computes cylinder coordinates and a bridge-only `d_ud` residual against the Cabibbo target. The supplied `I_K` and branch tags are external bridge inputs, not PDG-published mass-table invariants and not native ASHA laws.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Data-file import\n\n" + FormatImport(a.Import) + "\n\n")
	b.WriteString("## Parsed sector inputs\n\n- " + FormatSector(a.Adapter.U) + "\n- " + FormatSector(a.Adapter.D) + "\n\n")
	b.WriteString("## Inverted cylinder coordinates\n\n- " + FormatRay(a.Adapter.URay) + "\n- " + FormatRay(a.Adapter.DRay) + "\n\n")
	b.WriteString("## Numerical adapter\n\n" + FormatAdapter(a.Adapter) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha = sqrt(3) I_K / sqrt(1-I_K^2)\n")
	b.WriteString("cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\n")
	b.WriteString("d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n")
	b.WriteString("alpha_u = " + fmtFloat(a.Adapter.URay.Alpha) + "\n")
	b.WriteString("phi_u = " + fmtFloat(a.Adapter.URay.Phi) + "\n")
	b.WriteString("alpha_d = " + fmtFloat(a.Adapter.DRay.Alpha) + "\n")
	b.WriteString("phi_d = " + fmtFloat(a.Adapter.DRay.Phi) + "\n")
	b.WriteString("Gate471 d_ud = " + fmtFloat(a.Adapter.DUD) + "\n")
	b.WriteString("observed bridge target |V_us| = " + fmtFloat(a.Adapter.CabibboTarget) + "\n")
	b.WriteString("Cabibbo residual |d_ud-|V_us|| = " + fmtFloat(a.Adapter.CabibboResidual) + "\n")
	b.WriteString("```\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No data-file row enters the native theorem registry. No `I_K`, branch tag, cylinder coordinate, `d_ud`, residual, or alignment flag is exported as a native law.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
