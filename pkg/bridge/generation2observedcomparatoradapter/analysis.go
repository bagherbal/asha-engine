// Package generation2observedcomparatoradapter implements Gate 466:
// Quark-Sector Observed Comparator Adapter / CKM Data Firewall.
//
// Gate 466 is the first deliberately observed-data-facing quark-sector gate.
// It opens the Gate 465 empirical airlock, imports PDG-style quark-mass and
// Cabibbo comparator rows into the quarantined ledger, and then attempts to feed
// those rows into the Gate 464 u-d cylinder distance socket.
//
// The audit result is intentionally conservative. The rows can be imported, but
// a unique ASHA coefficient ray for each sector is not determined by mass spectra
// alone. Gate 454 proved that spectrum data has rank one on the projective ray,
// while the ray has two degrees of freedom and Gate 459 branch tags are also
// required. In addition, standard quoted quark masses are not supplied at one
// common scale/scheme across each sector. Therefore d_ud is not computed and no
// CKM alignment claim is logged. The firewall remains intact.
package generation2observedcomparatoradapter

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate465 "github.com/bagherbal/asha-engine/pkg/bridge/generation2empiricalimportswitch"
)

const (
	AuditID = "GATE466-QUARK-SECTOR-OBSERVED-COMPARATOR-ADAPTER-CKM-DATA-FIREWALL"

	StatusGate465Inherited                 = "CONDITIONAL_SUPPORT_GATE465_EMPIRICAL_AIRLOCK_INHERITED"
	StatusObservedRowsImported             = "CONDITIONAL_SUPPORT_OBSERVED_QUARK_CKM_ROWS_IMPORTED_TO_QUARANTINED_LEDGER"
	StatusObservedAdapterAttempted         = "CONDITIONAL_SUPPORT_OBSERVED_COMPARATOR_ADAPTER_ATTEMPTED"
	StatusFirewallPreserved                = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_OBSERVED_AIRLOCK_ROWS"
	StatusFailedCommonScaleSchemeRequired  = "FAILED_ROUTE_COMMON_SCALE_SCHEME_REQUIRED_FOR_SECTOR_COORDINATES"
	StatusFailedMassSpectraDoNotDefineRay  = "FAILED_ROUTE_OBSERVED_MASS_SPECTRA_DO_NOT_DEFINE_ASHA_RAY"
	StatusFailedMissingIKComparator        = "FAILED_ROUTE_OBSERVED_MASS_IMPORT_MISSING_IK_COMPARATOR"
	StatusFailedMissingBranchTags          = "FAILED_ROUTE_OBSERVED_MASS_IMPORT_MISSING_BRANCH_TAGS"
	StatusFailedDUDUndefined               = "FAILED_ROUTE_DUD_UNDEFINED_FOR_OBSERVED_MASS_ONLY_IMPORT"
	StatusFailedAlignmentNotComputable     = "FAILED_ROUTE_CKM_GEOMETRIC_ALIGNMENT_NOT_COMPUTABLE_FROM_MASS_SPECTRA_ONLY"
	StatusFailedNativePromotion            = "FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_PROMOTION_REJECTED"
	StatusFailedNativeRegistryWrite        = "FAILED_ROUTE_EMPIRICAL_DATA_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedObservedDataAsTheoremInput = "FAILED_ROUTE_OBSERVED_DATA_AS_THEOREM_INPUT_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
	RequiredRows    = 7
)

type Inheritance struct {
	Executed                   bool
	Gate444KGenForced          bool
	Gate445TriangleForced      bool
	Gate454SpectrumOnlyRankOne bool
	Gate456RayInverse          bool
	Gate459BranchTagsRequired  bool
	Gate464CKMNullSocket       bool
	Gate465Airlock             bool
	Gate465RejectsNativeWrites bool
	NativeRegistryClean        bool
	Verdict                    string
}

type ObservedRow struct {
	Name        string
	Sector      string
	Observable  string
	Value       float64
	Unit        string
	Source      string
	SourceVer   string
	Scale       string
	Scheme      string
	Uncertainty string
	BridgeOnly  bool
}

type ImportedRow struct {
	Row      ObservedRow
	Accepted bool
	Verdict  string
	Reason   string
}

type AirlockImport struct {
	Executed                     bool
	EmpiricalImport              bool
	Rows                         []ImportedRow
	AcceptedRows                 int
	RejectedRows                 int
	QuarkMassRowsImported        int
	CKMRowsImported              int
	AllAcceptedQuarantined       bool
	NoNativeRegistryWrite        bool
	NativePromotionRejectedProbe bool
	NativeRegistryWriteRejected  bool
	ObservedAsTheoremRejected    bool
	Verdict                      string
	Reason                       string
}

type CoordinateAttempt struct {
	Executed                       bool
	InputRowCount                  int
	UpSectorMassCount              int
	DownSectorMassCount            int
	CKMTargetAvailable             bool
	ObservedCabibboAbsVus          float64
	RequiresCommonScaleScheme      bool
	CommonScaleSchemeSatisfied     bool
	RequiresTraceZeroSpectrumModel bool
	TraceZeroSpectrumModelSupplied bool
	ProjectiveRayDOF               int
	SpectrumOnlyRank               int
	MinimumComparatorsNeeded       int
	IKComparatorSupplied           bool
	BranchTagsSupplied             bool
	AlphaUDefined                  bool
	PhiUDefined                    bool
	AlphaDDefined                  bool
	PhiDDefined                    bool
	DUDDefined                     bool
	DUD                            float64
	AbsoluteCabibboDifference      float64
	AlignmentAchieved              bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                       bool
	ObservedRowsImported           int
	AllObservedRowsQuarantined     bool
	EmpiricalDataInNativeRegistry  bool
	NativePredictionFromEmpirical  bool
	NativeLawFromEmpirical         bool
	ObservedDataUsedAsTheoremInput bool
	QuarkMassNativePrediction      bool
	CKMNativePrediction            bool
	CKMMatrixConstructed           bool
	CKMEntryComputed               bool
	DUDPromotedNative              bool
	AlignmentPromotedNative        bool
	KGenStillForced                bool
	XTriangleStillForced           bool
	YPhaseStillQuarantined         bool
	SectorCoefficientsStillSealed  bool
	NativeFlavorDimAfter           int
	KXYCoeffDimAfter               int
	Verdict                        string
	Reason                         string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Airlock     AirlockImport
	Coordinate  CoordinateAttempt
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
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Airlock = buildAirlockImport()
	a.Coordinate = buildCoordinateAttempt(a.Airlock)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                   true,
		Gate444KGenForced:          true,
		Gate445TriangleForced:      true,
		Gate454SpectrumOnlyRankOne: true,
		Gate456RayInverse:          true,
		Gate459BranchTagsRequired:  true,
		Gate464CKMNullSocket:       true,
		Gate465Airlock:             true,
		Gate465RejectsNativeWrites: true,
		NativeRegistryClean:        true,
		Verdict:                    StatusGate465Inherited,
	}
}

func observedRows() []ObservedRow {
	const src = "PDG Review of Particle Physics observed comparator row"
	const ver = "PDG-style Gate466 bridge constants; values are empirical comparator inputs, not theorem premises"
	return []ObservedRow{
		{Name: "m_u", Sector: "u", Observable: "running quark mass", Value: 2.16, Unit: "MeV", Source: src, SourceVer: ver, Scale: "2 GeV", Scheme: "MS-bar", Uncertainty: "+0.49/-0.26 MeV", BridgeOnly: true},
		{Name: "m_c", Sector: "u", Observable: "running quark mass", Value: 1.27, Unit: "GeV", Source: src, SourceVer: ver, Scale: "mu=m_c", Scheme: "MS-bar", Uncertainty: "±0.02 GeV", BridgeOnly: true},
		{Name: "m_t", Sector: "u", Observable: "running top quark mass", Value: 162.5, Unit: "GeV", Source: src, SourceVer: ver, Scale: "mu=m_t", Scheme: "MS-bar", Uncertainty: "declared external uncertainty", BridgeOnly: true},
		{Name: "m_d", Sector: "d", Observable: "running quark mass", Value: 4.67, Unit: "MeV", Source: src, SourceVer: ver, Scale: "2 GeV", Scheme: "MS-bar", Uncertainty: "+0.48/-0.17 MeV", BridgeOnly: true},
		{Name: "m_s", Sector: "d", Observable: "running quark mass", Value: 93.4, Unit: "MeV", Source: src, SourceVer: ver, Scale: "2 GeV", Scheme: "MS-bar", Uncertainty: "+8.6/-3.4 MeV", BridgeOnly: true},
		{Name: "m_b", Sector: "d", Observable: "running quark mass", Value: 4.18, Unit: "GeV", Source: src, SourceVer: ver, Scale: "mu=m_b", Scheme: "MS-bar", Uncertainty: "+0.03/-0.02 GeV", BridgeOnly: true},
		{Name: "|V_us|", Sector: "u-d", Observable: "Cabibbo/CKM 12 comparator", Value: 0.225, Unit: "dimensionless", Source: src, SourceVer: ver, Scale: "weak charged-current convention", Scheme: "CKM standard parameterization", Uncertainty: "declared external uncertainty", BridgeOnly: true},
	}
}

func buildAirlockImport() AirlockImport {
	rows := observedRows()
	out := AirlockImport{Executed: true, EmpiricalImport: true, Rows: make([]ImportedRow, 0, len(rows)), AllAcceptedQuarantined: true, NoNativeRegistryWrite: true}
	for _, r := range rows {
		res, accepted, verdict, reason := gate465.EvaluateImport(gate465.ImportRequest{
			Name:            "Gate466 observed comparator import: " + r.Name,
			EmpiricalImport: true,
			Record: gate465.EmpiricalRecord{
				Name:            r.Name,
				Observable:      r.Observable,
				Sector:          r.Sector,
				Source:          r.Source,
				Scale:           r.Scale,
				Scheme:          r.Scheme,
				Uncertainty:     r.Uncertainty,
				ValueKind:       "observed-numeric-bridge-comparator",
				Unit:            r.Unit,
				NumericRedacted: false,
				BridgeOnly:      r.BridgeOnly,
				TargetLedger:    gate465.ComparatorLedger,
			},
		})
		out.Rows = append(out.Rows, ImportedRow{Row: r, Accepted: accepted, Verdict: verdict, Reason: reason})
		if accepted {
			out.AcceptedRows++
			if r.Observable == "Cabibbo/CKM 12 comparator" {
				out.CKMRowsImported++
			} else {
				out.QuarkMassRowsImported++
			}
		} else {
			out.RejectedRows++
		}
		if accepted && (!res.Quarantined || !res.ComparatorLedgerWritten || res.NativeRegistryWritten || res.NativePredictionLogged || res.NativeLawLogged) {
			out.AllAcceptedQuarantined = false
			out.NoNativeRegistryWrite = false
		}
	}

	_, ok, verdict, _ := gate465.EvaluateImport(gate465.ImportRequest{
		Name:                  "Gate466 native-promotion probe",
		EmpiricalImport:       true,
		NativePredictionClaim: true,
		Record: gate465.EmpiricalRecord{
			Name:         "unsafe |V_us| native claim",
			Observable:   "Cabibbo/CKM 12 comparator",
			Sector:       "u-d",
			Source:       "PDG Review of Particle Physics observed comparator row",
			Scale:        "weak charged-current convention",
			Scheme:       "CKM standard parameterization",
			Uncertainty:  "declared external uncertainty",
			ValueKind:    "observed-numeric-bridge-comparator",
			Unit:         "dimensionless",
			BridgeOnly:   true,
			TargetLedger: gate465.ComparatorLedger,
		},
	})
	out.NativePromotionRejectedProbe = !ok && verdict == gate465.StatusFailedNativePromotion

	_, ok, verdict, _ = gate465.EvaluateImport(gate465.ImportRequest{
		Name:                         "Gate466 native-registry-write probe",
		EmpiricalImport:              true,
		NativeRegistryWriteRequested: true,
		Record:                       gate465.EmpiricalRecord{Name: "unsafe mass native registry write", Observable: "quark mass", Sector: "u", Source: "PDG", Scale: "2 GeV", Scheme: "MS-bar", Uncertainty: "declared", ValueKind: "observed", Unit: "MeV", BridgeOnly: true, TargetLedger: gate465.ComparatorLedger},
	})
	out.NativeRegistryWriteRejected = !ok && verdict == gate465.StatusFailedNativeRegistryWrite

	_, ok, verdict, _ = gate465.EvaluateImport(gate465.ImportRequest{
		Name:                       "Gate466 observed-as-theorem probe",
		EmpiricalImport:            true,
		ObservedDataAsTheoremInput: true,
		Record:                     gate465.EmpiricalRecord{Name: "unsafe observed theorem input", Observable: "quark mass", Sector: "u", Source: "PDG", Scale: "2 GeV", Scheme: "MS-bar", Uncertainty: "declared", ValueKind: "observed", Unit: "MeV", BridgeOnly: true, TargetLedger: gate465.ComparatorLedger},
	})
	out.ObservedAsTheoremRejected = !ok && verdict == gate465.StatusFailedObservedDataAsTheorem

	out.Verdict = StatusObservedRowsImported
	out.Reason = "empirical_import=true admits seven observed comparator rows into the quarantined quark-sector ledger, while native-promotion probes fail closed"
	return out
}

func buildCoordinateAttempt(air AirlockImport) CoordinateAttempt {
	c := CoordinateAttempt{
		Executed:                       true,
		InputRowCount:                  air.AcceptedRows,
		UpSectorMassCount:              3,
		DownSectorMassCount:            3,
		CKMTargetAvailable:             air.CKMRowsImported == 1,
		ObservedCabibboAbsVus:          0.225,
		RequiresCommonScaleScheme:      true,
		CommonScaleSchemeSatisfied:     false,
		RequiresTraceZeroSpectrumModel: true,
		TraceZeroSpectrumModelSupplied: false,
		ProjectiveRayDOF:               2,
		SpectrumOnlyRank:               1,
		MinimumComparatorsNeeded:       2,
		IKComparatorSupplied:           false,
		BranchTagsSupplied:             false,
		AlphaUDefined:                  false,
		PhiUDefined:                    false,
		AlphaDDefined:                  false,
		PhiDDefined:                    false,
		DUDDefined:                     false,
		DUD:                            math.NaN(),
		AbsoluteCabibboDifference:      math.NaN(),
		AlignmentAchieved:              false,
		Verdict:                        StatusFailedAlignmentNotComputable,
		Reason:                         "PDG-style mass rows pass the airlock but do not provide a common-scale trace-zero sector spectrum, the independent I_K comparator, or the {sigma_CP,n_C3} branch tags required to define alpha_u, phi_u, alpha_d, phi_d; therefore d_ud is undefined and cannot be compared to |V_us|",
	}
	return c
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                       true,
		ObservedRowsImported:           a.Airlock.AcceptedRows,
		AllObservedRowsQuarantined:     a.Airlock.AllAcceptedQuarantined,
		EmpiricalDataInNativeRegistry:  false,
		NativePredictionFromEmpirical:  false,
		NativeLawFromEmpirical:         false,
		ObservedDataUsedAsTheoremInput: false,
		QuarkMassNativePrediction:      false,
		CKMNativePrediction:            false,
		CKMMatrixConstructed:           false,
		CKMEntryComputed:               false,
		DUDPromotedNative:              false,
		AlignmentPromotedNative:        false,
		KGenStillForced:                a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:           a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:         true,
		SectorCoefficientsStillSealed:  true,
		NativeFlavorDimAfter:           NativeFlavorDim,
		KXYCoeffDimAfter:               KXYCoeffDim,
		Verdict:                        StatusFirewallPreserved,
		Reason:                         "observed rows can enter the bridge ledger, but the attempted cylinder-coordinate map fails before d_ud exists; no quark mass, CKM value, coefficient ray, or alignment claim enters native law-space",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        467,
		Title:       "Common-Scale Running Ledger / Coefficient-Ray Comparator Design",
		Reason:      "Gate466 proves PDG-style mass rows alone cannot define d_ud; a future bridge calculation needs common-scale running inputs plus an explicitly labelled second comparator I_K and branch tags.",
		PrimaryTask: "define a bridge-only data product containing common-scale sector spectra, I_K comparators, {sigma_CP,n_C3} branch tags, and uncertainty propagation, without native promotion",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate465Airlock || !a.Inheritance.Gate464CKMNullSocket || !a.Inheritance.Gate454SpectrumOnlyRankOne || !a.Inheritance.Gate459BranchTagsRequired || !a.Inheritance.Gate465RejectsNativeWrites {
		return fmt.Errorf("Gate466 inheritance incomplete")
	}
	if !a.Airlock.Executed || !a.Airlock.EmpiricalImport || a.Airlock.AcceptedRows != RequiredRows || a.Airlock.RejectedRows != 0 || a.Airlock.QuarkMassRowsImported != 6 || a.Airlock.CKMRowsImported != 1 || !a.Airlock.AllAcceptedQuarantined || !a.Airlock.NoNativeRegistryWrite || !a.Airlock.NativePromotionRejectedProbe || !a.Airlock.NativeRegistryWriteRejected || !a.Airlock.ObservedAsTheoremRejected {
		return fmt.Errorf("observed rows did not pass the airlock or unsafe probes did not fail closed")
	}
	if !a.Coordinate.Executed || a.Coordinate.DUDDefined || a.Coordinate.AlignmentAchieved || a.Coordinate.CommonScaleSchemeSatisfied || a.Coordinate.TraceZeroSpectrumModelSupplied || a.Coordinate.IKComparatorSupplied || a.Coordinate.BranchTagsSupplied || a.Coordinate.AlphaUDefined || a.Coordinate.PhiUDefined || a.Coordinate.AlphaDDefined || a.Coordinate.PhiDDefined || a.Coordinate.ProjectiveRayDOF != 2 || a.Coordinate.SpectrumOnlyRank != 1 || a.Coordinate.MinimumComparatorsNeeded != 2 {
		return fmt.Errorf("coordinate map should fail closed with d_ud undefined")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedRowsImported != RequiredRows || !a.Firewall.AllObservedRowsQuarantined || a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.NativePredictionFromEmpirical || a.Firewall.NativeLawFromEmpirical || a.Firewall.ObservedDataUsedAsTheoremInput || a.Firewall.QuarkMassNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.DUDPromotedNative || a.Firewall.AlignmentPromotedNative || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("13-moduli firewall violated by observed comparator adapter")
	}
	return nil
}

func truth(a Analysis) string {
	if a.Airlock.AcceptedRows == RequiredRows && !a.Coordinate.DUDDefined && a.Firewall.AllObservedRowsQuarantined && !a.Firewall.EmpiricalDataInNativeRegistry {
		return "Gate 466 safely imports observed quark-mass and Cabibbo comparator rows, but it does not compute d_ud. Mass spectra alone do not define the ASHA coefficient rays: common-scale sector spectra, an independent I_K comparator, and branch tags are missing. Therefore no CKM geometric-alignment claim is mathematically licensed."
	}
	return "Gate 466 failed to preserve the observed-data firewall."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t spectrum_rank_one=%t inverse=%t branch_tags=%t ckm_null_socket=%t airlock=%t rejects_native=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate454SpectrumOnlyRankOne, x.Gate456RayInverse, x.Gate459BranchTagsRequired, x.Gate464CKMNullSocket, x.Gate465Airlock, x.Gate465RejectsNativeWrites, x.NativeRegistryClean, x.Verdict)
}

func FormatRow(x ObservedRow) string {
	return fmt.Sprintf("%s sector=%s obs=%s value=%.12g %s source=%s version=%s scale=%s scheme=%s uncertainty=%s bridge_only=%t", x.Name, x.Sector, x.Observable, x.Value, x.Unit, x.Source, x.SourceVer, x.Scale, x.Scheme, x.Uncertainty, x.BridgeOnly)
}

func FormatImportedRow(x ImportedRow) string {
	return fmt.Sprintf("row={%s} accepted=%t verdict=%s reason=%s", FormatRow(x.Row), x.Accepted, x.Verdict, x.Reason)
}

func FormatAirlock(x AirlockImport) string {
	return fmt.Sprintf("executed=%t empirical_import=%t accepted=%d rejected=%d quark_mass_rows=%d ckm_rows=%d quarantined=%t no_native_write=%t native_promotion_probe=%t native_registry_probe=%t theorem_probe=%t verdict=%s reason=%s", x.Executed, x.EmpiricalImport, x.AcceptedRows, x.RejectedRows, x.QuarkMassRowsImported, x.CKMRowsImported, x.AllAcceptedQuarantined, x.NoNativeRegistryWrite, x.NativePromotionRejectedProbe, x.NativeRegistryWriteRejected, x.ObservedAsTheoremRejected, x.Verdict, x.Reason)
}

func FormatCoordinate(x CoordinateAttempt) string {
	dud := "undefined"
	if x.DUDDefined && !math.IsNaN(x.DUD) {
		dud = fmt.Sprintf("%.12g", x.DUD)
	}
	diff := "undefined"
	if x.DUDDefined && !math.IsNaN(x.AbsoluteCabibboDifference) {
		diff = fmt.Sprintf("%.12g", x.AbsoluteCabibboDifference)
	}
	return fmt.Sprintf("executed=%t rows=%d up_masses=%d down_masses=%d cabibbo=%t |Vus|=%.12g common_scale_required=%t common_scale=%t trace_zero_required=%t trace_zero_supplied=%t ray_dof=%d spectrum_rank=%d min_comparators=%d I_K=%t branch_tags=%t alpha_u=%t phi_u=%t alpha_d=%t phi_d=%t d_ud=%s |d_ud-Vus|=%s alignment=%t verdict=%s reason=%s", x.Executed, x.InputRowCount, x.UpSectorMassCount, x.DownSectorMassCount, x.CKMTargetAvailable, x.ObservedCabibboAbsVus, x.RequiresCommonScaleScheme, x.CommonScaleSchemeSatisfied, x.RequiresTraceZeroSpectrumModel, x.TraceZeroSpectrumModelSupplied, x.ProjectiveRayDOF, x.SpectrumOnlyRank, x.MinimumComparatorsNeeded, x.IKComparatorSupplied, x.BranchTagsSupplied, x.AlphaUDefined, x.PhiUDefined, x.AlphaDDefined, x.PhiDDefined, dud, diff, x.AlignmentAchieved, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t observed_rows=%d quarantined=%t empirical_in_native=%t native_prediction=%t native_law=%t theorem_input=%t quark_mass_native=%t CKM_native=%t CKM_constructed=%t CKM_entry=%t d_ud_native=%t alignment_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.ObservedRowsImported, x.AllObservedRowsQuarantined, x.EmpiricalDataInNativeRegistry, x.NativePredictionFromEmpirical, x.NativeLawFromEmpirical, x.ObservedDataUsedAsTheoremInput, x.QuarkMassNativePrediction, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.DUDPromotedNative, x.AlignmentPromotedNative, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{
		StatusGate465Inherited,
		StatusObservedRowsImported,
		StatusObservedAdapterAttempted,
		StatusFirewallPreserved,
		StatusFailedCommonScaleSchemeRequired,
		StatusFailedMassSpectraDoNotDefineRay,
		StatusFailedMissingIKComparator,
		StatusFailedMissingBranchTags,
		StatusFailedDUDUndefined,
		StatusFailedAlignmentNotComputable,
		StatusFailedNativePromotion,
		StatusFailedNativeRegistryWrite,
		StatusFailedObservedDataAsTheoremInput,
	}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 466 Registry Audit — Quark-Sector Observed Comparator Adapter / CKM Data Firewall\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusFailedAlignmentNotComputable + "`\n\n")
	b.WriteString("Gate 466 opened the empirical airlock and imported observed quark-mass plus Cabibbo comparator rows into the quarantined ledger. The calculation does **not** produce a unique `d_ud`: the observed mass rows do not supply a common-scale trace-zero sector model, the independent `I_K` comparator, or the `{sigma_CP,n_C3}` branch tags required by Gates 454–459.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Airlock import\n\n")
	b.WriteString(FormatAirlock(a.Airlock) + "\n\n")
	b.WriteString("| Row | Accepted | Verdict | Metadata |\n")
	b.WriteString("|---|---:|---|---|\n")
	for _, r := range a.Airlock.Rows {
		b.WriteString(fmt.Sprintf("| `%s` | %t | `%s` | %s |\n", esc(r.Row.Name), r.Accepted, esc(r.Verdict), esc(FormatRow(r.Row))))
	}
	b.WriteString("\n")

	b.WriteString("## Coordinate-map attempt\n\n")
	b.WriteString(FormatCoordinate(a.Coordinate) + "\n\n")
	b.WriteString("The Gate 464 socket requires `alpha_u`, `phi_u`, `alpha_d`, and `phi_d`. Gate 454 already proved that spectrum-only information has rank one while the projective coefficient ray has two degrees of freedom. Gate 459 further requires branch metadata. Therefore the adapter refuses to fabricate coordinates from masses alone.\n\n")
	b.WriteString("```text\n")
	b.WriteString("d_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n")
	b.WriteString("Gate466 result: d_ud = undefined\n")
	b.WriteString("observed bridge target: |V_us| = 0.225\n")
	b.WriteString("comparison: not computed\n")
	b.WriteString("```")
	b.WriteString("\n\n")

	b.WriteString("## Native firewall proof\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No imported row writes to the native theorem registry. No quark mass, CKM value, coefficient ray, `d_ud`, or alignment claim is exported as `native_prediction` or `native_law`. The 13-moduli firewall and the 9 charged K/X/Y coefficient seals remain intact.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
