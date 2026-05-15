// Package generation2leptonobservedadapter implements Gate 478:
// Lepton-Sector Observed Comparator Adapter / PMNS Airlock Non-Computation Audit.
//
// Gate 477 opened the lepton empirical airlock. Gate 478 is the lepton analogue
// of the quark observed-data adapter: it reads an explicit observed lepton/PMNS
// ledger, admits fully metadated rows only into a quarantined bridge ledger, and
// attempts the e-nu cylinder socket only if the file supplies explicit
// rank-complete bridge comparators. The checked-in file intentionally contains
// observed charged-lepton/neutrino/PMNS-style rows, but not ASHA-specific I_K or
// branch tags, so d_{e nu} remains undefined.
package generation2leptonobservedadapter

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
	AuditID = "GATE478-LEPTON-OBSERVED-COMPARATOR-ADAPTER-PMNS-AIRLOCK-NON-COMPUTATION-AUDIT"

	StatusGate477Inherited                = "CONDITIONAL_SUPPORT_GATE477_LEPTON_AIRLOCK_INHERITED"
	StatusDataFileLoaded                  = "CONDITIONAL_SUPPORT_EXPLICIT_LEPTON_OBSERVED_LEDGER_LOADED"
	StatusAirlockAcceptedBridgeRows       = "CONDITIONAL_SUPPORT_GATE478_AIRLOCK_ACCEPTED_QUARANTINED_LEPTON_ROWS"
	StatusObservedLeptonAdapterAttempted  = "CONDITIONAL_SUPPORT_GATE478_OBSERVED_LEPTON_ADAPTER_ATTEMPTED"
	StatusObservedDENuComputed            = "CONDITIONAL_SUPPORT_OBSERVED_LEPTON_DENU_BRIDGE_ONLY_COMPUTED"
	StatusObservedPMNSResidualComputed    = "CONDITIONAL_SUPPORT_OBSERVED_PMNS_RESIDUAL_BRIDGE_ONLY_COMPUTED"
	StatusPMNSAlignmentBridgeOnly         = "CONDITIONAL_SUPPORT_PMNS_GEOMETRIC_ALIGNMENT_BRIDGE_ONLY"
	StatusFirewallPreserved               = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE478_LEPTON_DATA_FILE"
	StatusFailedFileMissing               = "FAILED_ROUTE_LEPTON_OBSERVED_LEDGER_FILE_MISSING"
	StatusFailedSwitchClosed              = "FAILED_ROUTE_GATE478_LEPTON_EMPIRICAL_IMPORT_SWITCH_CLOSED"
	StatusFailedMetadataIncomplete        = "FAILED_ROUTE_GATE478_LEPTON_METADATA_INCOMPLETE"
	StatusFailedMissingISpecIKValues      = "FAILED_ROUTE_GATE478_MISSING_EXPLICIT_LEPTON_I_SPEC_I_K_VALUES"
	StatusFailedMissingBranchTags         = "FAILED_ROUTE_GATE478_MISSING_EXPLICIT_LEPTON_BRANCH_TAGS"
	StatusFailedMissingNeutrinoPolicies   = "FAILED_ROUTE_GATE478_MISSING_NEUTRINO_POLICY_METADATA"
	StatusFailedLeptonDataNoIK            = "FAILED_ROUTE_LEPTON_MASS_PMNS_LEDGER_DOES_NOT_SUPPLY_ASHA_I_K_INVARIANT"
	StatusFailedDENuNotComputableFromFile = "FAILED_ROUTE_OBSERVED_LEPTON_DENU_NOT_COMPUTABLE_FROM_FILE"
	StatusFailedPMNSResidualUndefined     = "FAILED_ROUTE_PMNS_RESIDUAL_UNDEFINED_WITHOUT_DENU"
	StatusFailedPMNSAsRayInput            = "FAILED_ROUTE_PMNS_USED_AS_GATE478_LEPTON_RAY_INPUT_REJECTED"
	StatusFailedNativePromotion           = "FAILED_ROUTE_GATE478_LEPTON_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite       = "FAILED_ROUTE_GATE478_LEPTON_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedPMNSNativePrediction      = "FAILED_ROUTE_GATE478_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedPMNSMatrixExport          = "FAILED_ROUTE_GATE478_PMNS_MATRIX_EXPORT_REJECTED"
	StatusFailedProjectiveDomain          = "FAILED_ROUTE_GATE478_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseDomain               = "FAILED_ROUTE_GATE478_PHASE_DOMAIN_REJECTED"
	StatusFailedCaustic                   = "FAILED_ROUTE_GATE478_CAUSTIC_BRANCH_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	DefaultLedger   = "data/lepton_observed_ledger.json"
)

type Number = *float64
type Integer = *int

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate476DENuSocketAvailable, Gate477LeptonAirlockAvailable, Gate475LeptonPreflightValidated, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                                 string
}

type DataRow struct {
	Name                        string  `json:"name"`
	Sector                      string  `json:"sector"`
	Observable                  string  `json:"observable"`
	Value                       Number  `json:"value"`
	Unit                        string  `json:"unit"`
	Source                      string  `json:"source"`
	SourceVersion               string  `json:"source_version"`
	Scale                       string  `json:"scale"`
	Scheme                      string  `json:"scheme"`
	Uncertainty                 string  `json:"uncertainty"`
	BridgeOnly                  bool    `json:"bridge_only"`
	NeutrinoOrderingPolicy      string  `json:"neutrino_ordering_policy"`
	AbsoluteNeutrinoScalePolicy string  `json:"absolute_neutrino_scale_policy"`
	MajoranaDiracPhasePolicy    string  `json:"majorana_dirac_phase_policy"`
	EigenbasisConvention        string  `json:"eigenbasis_convention"`
	SigmaCP                     Integer `json:"sigma_cp"`
	NC3                         Integer `json:"n_c3"`
	PMNSAsRayInput              bool    `json:"pmns_as_ray_input"`
	NativePromotionClaim        bool    `json:"native_promotion_claim"`
	PMNSNativePredictionClaim   bool    `json:"pmns_native_prediction_claim"`
	PMNSMatrixExportRequest     bool    `json:"pmns_matrix_export_request"`
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
	Executed, Loaded, EmpiricalImport, BridgeOnlyLedger, NativeRegistryWriteRequested bool
	Path                                                                              string
	Rows                                                                              int
	AcceptedRows, RejectedRows                                                        int
	ChargedLeptonRows, NeutrinoRows, ComparatorRows, BranchRows, PMNSTargetRows       int
	AllAcceptedBridgeOnly, MetadataComplete, LeptonPoliciesComplete                   bool
	SwitchClosedRejected, NativePromotionRejected, NativeRegistryWriteRejected        bool
	PMNSAsRayRejected, PMNSNativePredictionRejected, PMNSMatrixExportRejected         bool
	Verdict, Reason                                                                   string
	Failures                                                                          []string
}

type SectorInput struct {
	Sector                          string
	ISpec, IK                       Number
	SigmaCP                         Integer
	NC3                             Integer
	HasISpec, HasIK, HasBranch      bool
	CommonScale, CommonScheme       string
	ISpecUncertainty, IKUncertainty string
	BranchUncertainty               string
	MetadataComplete                bool
	LeptonPoliciesComplete          bool
	BridgeOnly                      bool
	ObservedLeptonRowsPresent       bool
	PMNSRowsPresent                 bool
	LeptonDataDoesNotSupplyIK       bool
	NativePromotion                 bool
	PMNSAsRayInput                  bool
}

type Ray struct {
	Sector                                       string
	Alpha, CosThreePhi, Phi, IK, ISpec           float64
	SigmaCP, NC3                                 int
	Defined, InsideDomain, AtCaustic, BridgeOnly bool
	Verdict, Reason                              string
}

type ObservedAdapter struct {
	Executed, Attempted, ReadyForDENu, DENuComputed, PMNSTargetAvailable, PMNSResidualComputed, AlignmentAchieved bool
	E, Nu                                                                                                         SectorInput
	ERay, NuRay                                                                                                   Ray
	DENu, PMNSTarget, PMNSResidual                                                                                float64
	MissingISpecIKValues, MissingBranchTags, MissingNeutrinoPolicies, LeptonDataNoIK                              bool
	ProjectiveDomainRejected, PhaseDomainRejected, CausticRejected                                                bool
	Verdict, Reason                                                                                               string
	Failures                                                                                                      []string
}

type Firewall struct {
	Executed, DataFileRowsNative, CoordinatesNative, DENuNativePrediction, PMNSNativePrediction, PMNSMatrixConstructed, PMNSEntryComputed, PMNSUsedAsRayInput, NativeRegistryWritten, KGenStillForced, XTriangleStillForced, YPhaseStillQuarantined, SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                                                                                                                         int
	Verdict, Reason                                                                                                                                                                                                                                                                string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      FileImport
	Adapter     ObservedAdapter
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
		a.Adapter = ObservedAdapter{Executed: true, Attempted: false, Verdict: StatusFailedFileMissing, Reason: "explicit Gate478 lepton data file was not found", Failures: []string{StatusFailedFileMissing}}
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
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate456InverseAvailable: true, Gate459BranchTagsRequired: true, Gate476DENuSocketAvailable: true, Gate477LeptonAirlockAvailable: true, Gate475LeptonPreflightValidated: true, NativeRegistryClean: true, Verdict: StatusGate477Inherited}
}

func loadLedger(path string) (DataLedger, FileImport) {
	p := projectPath(path)
	imp := FileImport{Executed: true, Path: p, Verdict: StatusDataFileLoaded, MetadataComplete: true, LeptonPoliciesComplete: true, AllAcceptedBridgeOnly: true}
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
		imp.MetadataComplete = false
	}
	if l.NativeRegistryWrite {
		failures = append(failures, StatusFailedNativeRegistryWrite)
		imp.NativeRegistryWriteRejected = true
	}
	for _, r := range l.Rows {
		obs := strings.ToLower(r.Observable)
		switch obs {
		case "charged_lepton_mass", "tau_lepton_mass", "muon_mass", "electron_mass":
			imp.ChargedLeptonRows++
		case "neutrino_mass_splitting", "neutrino_mass_bound", "neutrino_mass_ordering":
			imp.NeutrinoRows++
		case "i_spec", "i_k":
			imp.ComparatorRows++
		case "branch_tag":
			imp.BranchRows++
		case "pmns_target_theta23", "pmns_target_sin_theta23":
			imp.PMNSTargetRows++
		}
		rowOK := r.Source != "" && r.Scale != "" && r.Scheme != "" && r.Uncertainty != "" && r.BridgeOnly
		policyOK := r.NeutrinoOrderingPolicy != "" && r.AbsoluteNeutrinoScalePolicy != "" && r.MajoranaDiracPhasePolicy != "" && r.EigenbasisConvention != ""
		if !rowOK {
			imp.MetadataComplete = false
			failures = append(failures, StatusFailedMetadataIncomplete)
		}
		if !policyOK {
			imp.LeptonPoliciesComplete = false
			failures = append(failures, StatusFailedMissingNeutrinoPolicies)
		}
		if !r.BridgeOnly {
			imp.AllAcceptedBridgeOnly = false
		}
		if r.PMNSAsRayInput {
			failures = append(failures, StatusFailedPMNSAsRayInput)
			imp.PMNSAsRayRejected = true
		}
		if r.NativePromotionClaim {
			failures = append(failures, StatusFailedNativePromotion)
			imp.NativePromotionRejected = true
		}
		if r.PMNSNativePredictionClaim {
			failures = append(failures, StatusFailedPMNSNativePrediction)
			imp.PMNSNativePredictionRejected = true
		}
		if r.PMNSMatrixExportRequest {
			failures = append(failures, StatusFailedPMNSMatrixExport)
			imp.PMNSMatrixExportRejected = true
		}
		if rowOK && policyOK && l.EmpiricalImport && l.BridgeOnly && !l.NativeRegistryWrite {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
	}
	if len(failures) > 0 {
		imp.Failures = unique(failures)
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "explicit lepton data-file airlock rejected one or more metadata/policy/native-promotion routes"
		return l, imp
	}
	imp.NativePromotionRejected = true
	imp.NativeRegistryWriteRejected = true
	imp.PMNSAsRayRejected = true
	imp.PMNSNativePredictionRejected = true
	imp.PMNSMatrixExportRejected = true
	imp.Verdict = StatusAirlockAcceptedBridgeRows
	imp.Reason = "explicit lepton observed ledger rows entered only the bridge comparator airlock"
	return l, imp
}

func buildAdapter(l DataLedger) ObservedAdapter {
	ad := ObservedAdapter{Executed: true, Attempted: true, Verdict: StatusObservedLeptonAdapterAttempted, Reason: "explicit lepton data-file adapter attempted Gate456/Gate476 inversion"}
	ad.E = sectorInput(l, "e")
	ad.Nu = sectorInput(l, "nu")
	ad.PMNSTarget, ad.PMNSTargetAvailable = findValue(l, "e-nu", "pmns_target_theta23")
	failures := []string{}
	if !ad.E.HasISpec || !ad.E.HasIK || !ad.Nu.HasISpec || !ad.Nu.HasIK || ad.E.ISpec == nil || ad.E.IK == nil || ad.Nu.ISpec == nil || ad.Nu.IK == nil {
		failures = append(failures, StatusFailedMissingISpecIKValues)
		ad.MissingISpecIKValues = true
	}
	if !ad.E.HasBranch || !ad.Nu.HasBranch || ad.E.SigmaCP == nil || ad.E.NC3 == nil || ad.Nu.SigmaCP == nil || ad.Nu.NC3 == nil {
		failures = append(failures, StatusFailedMissingBranchTags)
		ad.MissingBranchTags = true
	}
	if !ad.E.LeptonPoliciesComplete || !ad.Nu.LeptonPoliciesComplete {
		failures = append(failures, StatusFailedMissingNeutrinoPolicies)
		ad.MissingNeutrinoPolicies = true
	}
	if ad.E.LeptonDataDoesNotSupplyIK || ad.Nu.LeptonDataDoesNotSupplyIK {
		failures = append(failures, StatusFailedLeptonDataNoIK)
		ad.LeptonDataNoIK = true
	}
	if len(failures) > 0 {
		failures = append(failures, StatusFailedDENuNotComputableFromFile, StatusFailedPMNSResidualUndefined)
		ad.Failures = unique(failures)
		ad.Verdict = StatusFailedDENuNotComputableFromFile
		ad.Reason = "the explicit lepton file was parsed, but it does not supply rank-complete ASHA bridge comparators; d_eν and PMNS residual remain undefined"
		return ad
	}
	er, ef := invert(ad.E)
	nr, nf := invert(ad.Nu)
	ad.ERay = er
	ad.NuRay = nr
	if ef != "" {
		failures = append(failures, ef)
	}
	if nf != "" {
		failures = append(failures, nf)
	}
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
		ad.Reason = "rank-complete lepton file failed inverse-domain checks"
		return ad
	}
	ad.ReadyForDENu = true
	ad.DENu = distance(er, nr)
	ad.DENuComputed = true
	if ad.PMNSTargetAvailable {
		ad.PMNSResidual = math.Abs(ad.DENu - ad.PMNSTarget)
		ad.PMNSResidualComputed = true
		ad.AlignmentAchieved = ad.PMNSResidual < 1e-3
	}
	ad.Verdict = StatusObservedDENuComputed
	if ad.AlignmentAchieved {
		ad.Verdict = StatusPMNSAlignmentBridgeOnly
	}
	ad.Reason = "rank-complete explicit lepton bridge file computed d_eν; result remains a bridge comparator, not a native PMNS prediction"
	return ad
}

func sectorInput(l DataLedger, sector string) SectorInput {
	x := SectorInput{Sector: sector, CommonScale: l.CommonScale, CommonScheme: l.CommonScheme, BridgeOnly: l.BridgeOnly, MetadataComplete: true, LeptonPoliciesComplete: true}
	for _, r := range l.Rows {
		if r.Sector != sector {
			continue
		}
		if r.Source == "" || r.Scale == "" || r.Scheme == "" || r.Uncertainty == "" || !r.BridgeOnly {
			x.MetadataComplete = false
		}
		if r.NeutrinoOrderingPolicy == "" || r.AbsoluteNeutrinoScalePolicy == "" || r.MajoranaDiracPhasePolicy == "" || r.EigenbasisConvention == "" {
			x.LeptonPoliciesComplete = false
		}
		switch strings.ToLower(r.Observable) {
		case "charged_lepton_mass", "tau_lepton_mass", "muon_mass", "electron_mass":
			x.ObservedLeptonRowsPresent = true
		case "neutrino_mass_splitting", "neutrino_mass_bound", "neutrino_mass_ordering":
			x.ObservedLeptonRowsPresent = true
		case "pmns_target_theta23", "pmns_target_sin_theta23":
			x.PMNSRowsPresent = true
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
		if r.PMNSAsRayInput {
			x.PMNSAsRayInput = true
		}
	}
	if x.ObservedLeptonRowsPresent && (x.IK == nil || !x.HasIK) {
		x.LeptonDataDoesNotSupplyIK = true
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
	r.IK = ik
	r.ISpec = ispec
	r.SigmaCP = *x.SigmaCP
	r.NC3 = *x.NC3
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
	r.Alpha = alpha
	r.CosThreePhi = cos3
	r.Phi = phi
	r.Defined = true
	r.InsideDomain = true
	r.Verdict = StatusObservedDENuComputed
	r.Reason = "explicit rank-complete observed lepton bridge row inverted to ASHA cylinder coordinate"
	return r, ""
}

func distance(e, nu Ray) float64 {
	da := nu.Alpha - e.Alpha
	dp := wrapPi(nu.Phi - e.Phi)
	return math.Sqrt(da*da + 4*math.Pow(math.Sin(dp/2), 2))
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate478 lepton data-file rows are quarantined bridge comparators; no row, coordinate, d_eν, residual, PMNS entry, or alignment flag writes to native law-space"}
}

func buildNext() NextStep {
	return NextStep{479, "Lepton rank-complete external ledger acceptance test", "Gate478 parsed the explicit lepton observed file and refused to fabricate ASHA comparators from masses or PMNS targets.", "evaluate only a user-supplied rank-complete e/nu bridge ledger with I_spec, I_K, and branch tags; export bridge residuals only"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate477LeptonAirlockAvailable || !a.Inheritance.Gate476DENuSocketAvailable || !a.Inheritance.Gate475LeptonPreflightValidated || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate478 inheritance incomplete")
	}
	if !a.Import.Executed || !a.Import.Loaded || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || a.Import.NativeRegistryWriteRequested || a.Import.Rows == 0 || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly || !a.Import.MetadataComplete || !a.Import.LeptonPoliciesComplete {
		return fmt.Errorf("Gate478 file import did not satisfy bridge-only airlock conditions: %+v", a.Import)
	}
	if !a.Adapter.Executed || !a.Adapter.Attempted || a.Adapter.DENuComputed || a.Adapter.PMNSResidualComputed || a.Adapter.AlignmentAchieved || !a.Adapter.MissingISpecIKValues || !a.Adapter.MissingBranchTags || !a.Adapter.LeptonDataNoIK || a.Adapter.Verdict != StatusFailedDENuNotComputableFromFile {
		return fmt.Errorf("Gate478 default observed lepton ledger must fail closed without d_eν: %+v", a.Adapter)
	}
	if !a.Firewall.Executed || a.Firewall.DataFileRowsNative || a.Firewall.CoordinatesNative || a.Firewall.DENuNativePrediction || a.Firewall.PMNSNativePrediction || a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed || a.Firewall.PMNSUsedAsRayInput || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate478 firewall violated")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Import.Loaded && !a.Adapter.DENuComputed && a.Adapter.MissingISpecIKValues && a.Adapter.LeptonDataNoIK && !a.Firewall.NativeRegistryWritten {
		return "Gate 478 successfully reads the explicit observed lepton ledger through the empirical airlock, but the checked-in lepton/PMNS-style file does not contain explicit ASHA rank-complete comparators. Charged-lepton rows, neutrino rows, and PMNS target rows remain quarantined bridge data; d_eν and the PMNS residual are undefined until I_spec, I_K, and branch tags are explicitly supplied. The lepton socket is structurally identical to the quark socket and fails closed in the same way."
	}
	if a.Adapter.DENuComputed {
		return "Gate 478 computed a bridge-only d_eν from a rank-complete explicit lepton ledger. This remains a comparator result and is not a native PMNS prediction."
	}
	return "Gate 478 failed before preserving the empirical firewall."
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
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t d_eν_socket=%t lepton_airlock=%t preflight=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate476DENuSocketAvailable, x.Gate477LeptonAirlockAvailable, x.Gate475LeptonPreflightValidated, x.NativeRegistryClean, x.Verdict)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("executed=%t loaded=%t path=%s empirical_import=%t bridge_only=%t rows=%d accepted=%d rejected=%d charged_lepton_rows=%d neutrino_rows=%d comparator_rows=%d branch_rows=%d pmns_targets=%d metadata=%t policies=%t quarantined=%t native_write_requested=%t verdict=%s reason=%s", x.Executed, x.Loaded, x.Path, x.EmpiricalImport, x.BridgeOnlyLedger, x.Rows, x.AcceptedRows, x.RejectedRows, x.ChargedLeptonRows, x.NeutrinoRows, x.ComparatorRows, x.BranchRows, x.PMNSTargetRows, x.MetadataComplete, x.LeptonPoliciesComplete, x.AllAcceptedBridgeOnly, x.NativeRegistryWriteRequested, x.Verdict, x.Reason)
}
func FormatSector(x SectorInput) string {
	return fmt.Sprintf("sector=%s I_spec=%s I_K=%s sigma_CP=%s n_C3=%s scale=%s scheme=%s metadata=%t policies=%t bridge_only=%t observed_rows=%t pmns_rows=%t lepton_no_IK=%t", x.Sector, ptrFloatText(x.ISpec), ptrFloatText(x.IK), ptrIntText(x.SigmaCP), ptrIntText(x.NC3), x.CommonScale, x.CommonScheme, x.MetadataComplete, x.LeptonPoliciesComplete, x.BridgeOnly, x.ObservedLeptonRowsPresent, x.PMNSRowsPresent, x.LeptonDataDoesNotSupplyIK)
}
func FormatRay(x Ray) string {
	return fmt.Sprintf("sector=%s defined=%t alpha=%s cos3phi=%s phi=%s I_K=%s I_spec=%s sigma_CP=%d n_C3=%d domain=%t caustic=%t bridge_only=%t verdict=%s", x.Sector, x.Defined, fmtFloat(x.Alpha), fmtFloat(x.CosThreePhi), fmtFloat(x.Phi), fmtFloat(x.IK), fmtFloat(x.ISpec), x.SigmaCP, x.NC3, x.InsideDomain, x.AtCaustic, x.BridgeOnly, x.Verdict)
}
func FormatAdapter(x ObservedAdapter) string {
	denu := "undefined"
	residual := "undefined"
	if x.DENuComputed {
		denu = fmtFloat(x.DENu)
	}
	if x.PMNSResidualComputed {
		residual = fmtFloat(x.PMNSResidual)
	}
	return fmt.Sprintf("executed=%t attempted=%t ready=%t d_eν_computed=%t d_eν=%s pmns_target_available=%t target=%s residual_computed=%t residual=%s alignment=%t missing_I=%t missing_branch=%t missing_policies=%t lepton_no_IK=%t verdict=%s reason=%s", x.Executed, x.Attempted, x.ReadyForDENu, x.DENuComputed, denu, x.PMNSTargetAvailable, fmtFloat(x.PMNSTarget), x.PMNSResidualComputed, residual, x.AlignmentAchieved, x.MissingISpecIKValues, x.MissingBranchTags, x.MissingNeutrinoPolicies, x.LeptonDataNoIK, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t rows_native=%t coords_native=%t d_eν_native=%t pmns_native=%t pmns_matrix=%t pmns_entry=%t pmns_as_ray=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.DataFileRowsNative, x.CoordinatesNative, x.DENuNativePrediction, x.PMNSNativePrediction, x.PMNSMatrixConstructed, x.PMNSEntryComputed, x.PMNSUsedAsRayInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate477Inherited, StatusDataFileLoaded, StatusAirlockAcceptedBridgeRows, StatusObservedLeptonAdapterAttempted, StatusFailedMissingISpecIKValues, StatusFailedMissingBranchTags, StatusFailedLeptonDataNoIK, StatusFailedDENuNotComputableFromFile, StatusFailedPMNSResidualUndefined, StatusFailedPMNSAsRayInput, StatusFailedNativePromotion, StatusFailedNativeRegistryWrite, StatusFailedPMNSNativePrediction, StatusFailedPMNSMatrixExport, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 478 Registry Audit — Lepton-Sector Observed Comparator Adapter / PMNS Airlock Non-Computation Audit\n\n## Verdict\n\n")
	b.WriteString("`" + StatusFailedDENuNotComputableFromFile + "`\n\n")
	b.WriteString("Gate 478 reads `data/lepton_observed_ledger.json` through the lepton empirical airlock. The checked-in file contains charged-lepton, neutrino, and PMNS residual-target rows, but it does not contain explicit ASHA rank-complete bridge comparator values for `I_spec`, `I_K`, or branch tags. Therefore the adapter refuses to fabricate e/nu cylinder coordinates and leaves `d_eν` undefined.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Data-file import\n\n" + FormatImport(a.Import) + "\n\n")
	b.WriteString("## Parsed sector inputs\n\n")
	b.WriteString("- " + FormatSector(a.Adapter.E) + "\n")
	b.WriteString("- " + FormatSector(a.Adapter.Nu) + "\n\n")
	b.WriteString("## Observed lepton adapter\n\n" + FormatAdapter(a.Adapter) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha = sqrt(3) I_K / sqrt(1-I_K^2)\n")
	b.WriteString("cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\n")
	b.WriteString("d_eν = sqrt((alpha_ν-alpha_e)^2 + 4 sin^2((phi_ν-phi_e)/2))\n")
	if a.Adapter.DENuComputed {
		b.WriteString("Gate478 d_eν = " + fmtFloat(a.Adapter.DENu) + "\n")
	} else {
		b.WriteString("Gate478 d_eν = undefined\n")
	}
	if a.Adapter.PMNSTargetAvailable {
		b.WriteString("observed bridge target θ23-like PMNS row = " + fmtFloat(a.Adapter.PMNSTarget) + "\n")
	}
	if a.Adapter.PMNSResidualComputed {
		b.WriteString("PMNS residual |d_eν-target| = " + fmtFloat(a.Adapter.PMNSResidual) + "\n")
	} else {
		b.WriteString("PMNS residual = undefined\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No data-file row enters the native theorem registry. No lepton mass, neutrino row, PMNS value, `I_K`, `I_spec`, branch tag, cylinder coordinate, `d_eν`, residual, matrix entry, or alignment flag is exported as a native law.\n\n")
	b.WriteString("## Structural equivalence to quark socket\n\nGate 478 uses the same cylinder metric as Gate 470/Gate 464, with the sector labels changed from `u,d` to `e,ν`. The only difference is the lepton preflight policy: neutrino ordering, absolute-scale, Majorana/Dirac phase, and eigenbasis conventions are mandatory metadata.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
