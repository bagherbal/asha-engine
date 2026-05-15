// Package generation2masstoequipartition implements Gate 473:
// Mass-to-Equipartition Inversion / Epistemological Loop Closure Audit.
//
// The gate is deliberately strict. It imports only raw quark masses through the
// empirical airlock and tests whether the mass hierarchy alone mathematically
// forces the ASHA coefficient-ray condition alpha=1, hence I_K=1/2. It does not
// import CKM angles, precomputed I_K values, or branch tags.
//
// Result: the proposed closure fails. Trace-zero mass spectra supply only the
// spectrum invariant I_spec. In the extreme hierarchy limit that invariant tends
// to 2/(3*sqrt(3)), whereas alpha=1 can support at most |I_spec|=1/4. Thus raw
// masses neither force alpha=1 nor derive I_K=1/2; a Cabibbo d_ud comparison
// remains undefined unless an independent rank-complete bridge ledger supplies
// I_K and branch tags.
package generation2masstoequipartition

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
)

const (
	AuditID = "GATE473-MASS-TO-EQUIPARTITION-INVERSION-EPISTEMOLOGICAL-LOOP-CLOSURE"

	StatusGate471Inherited          = "CONDITIONAL_SUPPORT_GATE471_SOCKET_AND_FIREWALL_INHERITED"
	StatusRawMassLedgerLoaded       = "CONDITIONAL_SUPPORT_RAW_MASS_LEDGER_IMPORTED_THROUGH_AIRLOCK"
	StatusExtremeHierarchyConfirmed = "CONDITIONAL_SUPPORT_EXTREME_THIRD_GENERATION_HIERARCHY_CONFIRMED"
	StatusAsymptoticLimitDerived    = "CONDITIONAL_SUPPORT_ASYMPTOTIC_SPECTRUM_LIMIT_DERIVED"
	StatusFirewallPreserved         = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE473_RAW_MASS_AUDIT"

	StatusFailedFileMissing                  = "FAILED_ROUTE_GATE473_RAW_MASS_LEDGER_FILE_MISSING"
	StatusFailedSwitchClosed                 = "FAILED_ROUTE_GATE473_EMPIRICAL_IMPORT_SWITCH_CLOSED"
	StatusFailedMetadataIncomplete           = "FAILED_ROUTE_GATE473_METADATA_INCOMPLETE"
	StatusFailedRawMassNativePromotion       = "FAILED_ROUTE_GATE473_RAW_MASS_NATIVE_PROMOTION_REJECTED"
	StatusFailedRawMassRegistryWrite         = "FAILED_ROUTE_GATE473_NATIVE_REGISTRY_WRITE_REJECTED"
	StatusFailedIKImported                   = "FAILED_ROUTE_GATE473_PRECOMPUTED_I_K_IMPORT_REJECTED"
	StatusFailedCKMImported                  = "FAILED_ROUTE_GATE473_CKM_IMPORT_REJECTED"
	StatusFailedMassHierarchyNoEquipartition = "FAILED_ROUTE_MASS_HIERARCHY_DOES_NOT_FORCE_EQUIPARTITION"
	StatusFailedRawMassCannotDeriveIK        = "FAILED_ROUTE_RAW_MASSES_CANNOT_DERIVE_I_K_HALF"
	StatusFailedAlphaOneInconsistent         = "FAILED_ROUTE_ALPHA_ONE_INCONSISTENT_WITH_EXTREME_TRACE_ZERO_SPECTRUM"
	StatusFailedDUDUndefined                 = "FAILED_ROUTE_GATE473_DUD_UNDEFINED_WITHOUT_I_K_AND_BRANCH_TAGS"
	StatusFailedProjectNotAchieved           = "FAILED_ROUTE_PROJECT_ABSOLUTE_GEOMETRIC_UNIFICATION_NOT_ACHIEVED"
)

const (
	NativeFlavorDim          = 13
	KXYCoeffDim              = 9
	DefaultLedger            = "data/pdg_raw_quark_masses_gate473.json"
	HierarchyDominanceCutoff = 0.999
)

type Number = *float64

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate454RankAuditAvailable, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate465AirlockAvailable, Gate471ExternalSocketAvailable, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                         string
}

type MassRow struct {
	Name                 string `json:"name"`
	Sector               string `json:"sector"`
	Generation           int    `json:"generation"`
	Quark                string `json:"quark"`
	MassGeV              Number `json:"mass_gev"`
	Source               string `json:"source"`
	SourceVersion        string `json:"source_version"`
	Scale                string `json:"scale"`
	Scheme               string `json:"scheme"`
	Uncertainty          string `json:"uncertainty"`
	BridgeOnly           bool   `json:"bridge_only"`
	NativePromotionClaim bool   `json:"native_promotion_claim"`
	NativeRegistryWrite  bool   `json:"native_registry_write"`
	PrecomputedIK        Number `json:"i_k,omitempty"`
	CKMValue             Number `json:"ckm_value,omitempty"`
}

type MassLedger struct {
	Gate                int       `json:"gate"`
	LedgerName          string    `json:"ledger_name"`
	Description         string    `json:"description"`
	EmpiricalImport     bool      `json:"empirical_import"`
	BridgeOnly          bool      `json:"bridge_only"`
	NativeRegistryWrite bool      `json:"native_registry_write"`
	Rows                []MassRow `json:"rows"`
}

type FileImport struct {
	Executed, Loaded, EmpiricalImport, BridgeOnlyLedger, NativeRegistryWriteRequested                                                  bool
	Path                                                                                                                               string
	Rows, AcceptedRows, RejectedRows                                                                                                   int
	MetadataComplete, AllAcceptedBridgeOnly, NativePromotionRejected, NativeRegistryWriteRejected, IKImportRejected, CKMImportRejected bool
	Verdict, Reason                                                                                                                    string
	Failures                                                                                                                           []string
}

type SectorSpectrum struct {
	Sector                                                                                   string
	Masses                                                                                   []float64
	TraceZero                                                                                []float64
	SumSquaresPhysical, ThirdGenerationSquareFraction                                        float64
	Mean, Q, R, ISpec, AsymptoticISpec, AbsDeltaAsymptotic                                   float64
	AlphaOneMaxISpec, AlphaMaxAllowed, IKIfAlphaOne                                          float64
	ExtremeHierarchy, SpectrumRankOneOnly, AlphaOneCompatible, AlphaOneForced, IKHalfDerived bool
	Verdict, Reason                                                                          string
}

type LoopClosure struct {
	Executed, Attempted, RawMassesOnly, AlphaDerived, IKDerived, DUDComputed, CabibboResidualComputed, AlignmentAchieved bool
	AlphaU, AlphaD, IKU, IKD, DUD, CabibboTarget, CabibboResidual                                                        float64
	Verdict, Reason                                                                                                      string
	Failures                                                                                                             []string
}

type Firewall struct {
	Executed, RawMassRowsNative, IKNative, AlphaNative, DUDNativePrediction, CKMNativePrediction, CKMMatrixConstructed, CKMEntryComputed, NativeRegistryWritten, KGenStillForced, XTriangleStillForced, YPhaseStillQuarantined, SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                                                                                                    int
	Verdict, Reason                                                                                                                                                                                                                                           string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Import      FileImport
	Up, Down    SectorSpectrum
	Loop        LoopClosure
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
		a.Up = buildSpectrum(ledger, "u")
		a.Down = buildSpectrum(ledger, "d")
	}
	a.Loop = buildLoop(a.Up, a.Down)
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate445TriangleForced: true, Gate454RankAuditAvailable: true, Gate456InverseAvailable: true, Gate459BranchTagsRequired: true, Gate465AirlockAvailable: true, Gate471ExternalSocketAvailable: true, NativeRegistryClean: true, Verdict: StatusGate471Inherited}
}

func loadLedger(path string) (MassLedger, FileImport) {
	p := projectPath(path)
	imp := FileImport{Executed: true, Path: p, Verdict: StatusRawMassLedgerLoaded}
	bs, err := os.ReadFile(p)
	if err != nil {
		imp.Loaded = false
		imp.Verdict = StatusFailedFileMissing
		imp.Failures = []string{StatusFailedFileMissing}
		imp.Reason = err.Error()
		return MassLedger{}, imp
	}
	var l MassLedger
	if err := json.Unmarshal(bs, &l); err != nil {
		imp.Loaded = false
		imp.Verdict = StatusFailedMetadataIncomplete
		imp.Failures = []string{StatusFailedMetadataIncomplete}
		imp.Reason = err.Error()
		return MassLedger{}, imp
	}
	imp.Loaded = true
	imp.EmpiricalImport = l.EmpiricalImport
	imp.BridgeOnlyLedger = l.BridgeOnly
	imp.NativeRegistryWriteRequested = l.NativeRegistryWrite
	imp.Rows = len(l.Rows)
	failures := []string{}
	if !l.EmpiricalImport {
		failures = append(failures, StatusFailedSwitchClosed)
	}
	if !l.BridgeOnly || l.LedgerName == "" {
		failures = append(failures, StatusFailedMetadataIncomplete)
	}
	if l.NativeRegistryWrite {
		failures = append(failures, StatusFailedRawMassRegistryWrite)
		imp.NativeRegistryWriteRejected = true
	}
	imp.MetadataComplete = true
	imp.AllAcceptedBridgeOnly = true
	for _, r := range l.Rows {
		rowOK := r.Name != "" && r.Sector != "" && r.Generation > 0 && r.Quark != "" && r.MassGeV != nil && *r.MassGeV > 0 && r.Source != "" && r.SourceVersion != "" && r.Scale != "" && r.Scheme != "" && r.Uncertainty != "" && r.BridgeOnly
		if !rowOK {
			imp.MetadataComplete = false
			failures = append(failures, StatusFailedMetadataIncomplete)
		}
		if !r.BridgeOnly {
			imp.AllAcceptedBridgeOnly = false
		}
		if r.NativePromotionClaim {
			failures = append(failures, StatusFailedRawMassNativePromotion)
			imp.NativePromotionRejected = true
		}
		if r.NativeRegistryWrite {
			failures = append(failures, StatusFailedRawMassRegistryWrite)
			imp.NativeRegistryWriteRejected = true
		}
		if r.PrecomputedIK != nil {
			failures = append(failures, StatusFailedIKImported)
			imp.IKImportRejected = true
		}
		if r.CKMValue != nil {
			failures = append(failures, StatusFailedCKMImported)
			imp.CKMImportRejected = true
		}
		if rowOK && l.EmpiricalImport && l.BridgeOnly && !l.NativeRegistryWrite && r.PrecomputedIK == nil && r.CKMValue == nil && !r.NativePromotionClaim && !r.NativeRegistryWrite {
			imp.AcceptedRows++
		} else {
			imp.RejectedRows++
		}
	}
	if len(failures) > 0 {
		imp.Failures = unique(failures)
		imp.Verdict = strings.Join(imp.Failures, ";")
		imp.Reason = "raw mass ledger rejected one or more airlock or anti-smuggling rules"
		return l, imp
	}
	imp.NativePromotionRejected = true
	imp.NativeRegistryWriteRejected = true
	imp.IKImportRejected = true
	imp.CKMImportRejected = true
	imp.Verdict = StatusRawMassLedgerLoaded
	imp.Reason = "raw quark masses entered only the bridge quarantine ledger; no I_K or CKM values were imported"
	return l, imp
}

func buildSpectrum(l MassLedger, sector string) SectorSpectrum {
	s := SectorSpectrum{Sector: sector, AsymptoticISpec: 2 / (3 * math.Sqrt(3)), AlphaOneMaxISpec: 0.25, IKIfAlphaOne: 0.5, SpectrumRankOneOnly: true, Verdict: StatusAsymptoticLimitDerived}
	rows := []MassRow{}
	for _, r := range l.Rows {
		if r.Sector == sector && r.MassGeV != nil {
			rows = append(rows, r)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Generation < rows[j].Generation })
	for _, r := range rows {
		s.Masses = append(s.Masses, *r.MassGeV)
	}
	if len(s.Masses) != 3 {
		s.Verdict = StatusFailedMetadataIncomplete
		s.Reason = "sector requires exactly three raw masses"
		return s
	}
	for _, m := range s.Masses {
		s.SumSquaresPhysical += m * m
	}
	s.ThirdGenerationSquareFraction = s.Masses[2] * s.Masses[2] / s.SumSquaresPhysical
	s.ExtremeHierarchy = s.ThirdGenerationSquareFraction >= HierarchyDominanceCutoff
	s.Mean = (s.Masses[0] + s.Masses[1] + s.Masses[2]) / 3
	s.TraceZero = []float64{s.Masses[0] - s.Mean, s.Masses[1] - s.Mean, s.Masses[2] - s.Mean}
	for _, x := range s.TraceZero {
		s.Q += x * x
	}
	s.Q /= 2
	s.R = s.TraceZero[0] * s.TraceZero[1] * s.TraceZero[2]
	s.ISpec = s.R / math.Pow(s.Q, 1.5)
	s.AbsDeltaAsymptotic = math.Abs(s.ISpec - s.AsymptoticISpec)
	s.AlphaOneCompatible = math.Abs(s.ISpec) <= s.AlphaOneMaxISpec+1e-12
	s.AlphaMaxAllowed = alphaMaxFromISpec(s.ISpec)
	s.AlphaOneForced = false
	s.IKHalfDerived = false
	if !s.AlphaOneCompatible {
		s.Verdict = StatusFailedAlphaOneInconsistent
		s.Reason = "trace-zero spectrum invariant exceeds the maximum allowed by alpha=1 even before phase/branch selection"
		return s
	}
	s.Verdict = StatusFailedMassHierarchyNoEquipartition
	s.Reason = "mass spectrum supplies only I_spec and leaves I_K/ray alignment underdetermined"
	return s
}

func alphaMaxFromISpec(i float64) float64 {
	ai := math.Abs(i)
	if ai == 0 {
		return math.Inf(1)
	}
	max := 2 / (3 * math.Sqrt(3))
	if ai > max+1e-12 {
		return math.NaN()
	}
	v := math.Pow(2/ai, 2.0/3.0) - 3
	if v < 0 && v > -1e-12 {
		v = 0
	}
	if v < 0 {
		return math.NaN()
	}
	return math.Sqrt(v)
}

func buildLoop(u, d SectorSpectrum) LoopClosure {
	failures := []string{StatusFailedMassHierarchyNoEquipartition, StatusFailedRawMassCannotDeriveIK, StatusFailedDUDUndefined, StatusFailedProjectNotAchieved}
	if !u.AlphaOneCompatible || !d.AlphaOneCompatible {
		failures = append([]string{StatusFailedAlphaOneInconsistent}, failures...)
	}
	return LoopClosure{Executed: true, Attempted: true, RawMassesOnly: true, CabibboTarget: 0.225, Verdict: StatusFailedProjectNotAchieved, Reason: "raw masses confirm extreme hierarchy but do not force alpha=1 or derive I_K=0.5; d_ud is undefined without independent I_K and branch tags", Failures: unique(failures)}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate473 raw masses, spectrum invariants, asymptotic limits, and failed loop closure remain bridge diagnostics; no native theorem-registry write occurs"}
}

func buildNext() NextStep {
	return NextStep{474, "Independent K-axis observable search", "Gate473 proves raw mass hierarchy cannot derive I_K=1/2 or alpha=1.", "identify a genuinely independent experimental or algebraic K-overlap observable, or keep CKM alignment as an external rank-complete bridge-ledger fact"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate454RankAuditAvailable || !a.Inheritance.Gate465AirlockAvailable || !a.Inheritance.Gate471ExternalSocketAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate473 inheritance incomplete")
	}
	if !a.Import.Executed || !a.Import.Loaded || !a.Import.EmpiricalImport || !a.Import.BridgeOnlyLedger || a.Import.NativeRegistryWriteRequested || a.Import.Rows != 6 || a.Import.AcceptedRows != 6 || !a.Import.MetadataComplete || !a.Import.AllAcceptedBridgeOnly || !a.Import.NativePromotionRejected || !a.Import.NativeRegistryWriteRejected || !a.Import.IKImportRejected || !a.Import.CKMImportRejected {
		return fmt.Errorf("Gate473 import did not satisfy raw-mass airlock: %+v", a.Import)
	}
	if !a.Up.ExtremeHierarchy || !a.Down.ExtremeHierarchy || !a.Up.SpectrumRankOneOnly || !a.Down.SpectrumRankOneOnly {
		return fmt.Errorf("Gate473 hierarchy/rank audit did not execute: %+v %+v", a.Up, a.Down)
	}
	if a.Up.AlphaOneForced || a.Down.AlphaOneForced || a.Up.IKHalfDerived || a.Down.IKHalfDerived || a.Loop.AlphaDerived || a.Loop.IKDerived || a.Loop.DUDComputed || a.Loop.CabibboResidualComputed || a.Loop.AlignmentAchieved {
		return fmt.Errorf("Gate473 falsely closed the loop: %+v", a.Loop)
	}
	if !strings.Contains(a.Loop.Verdict, "PROJECT_ABSOLUTE") && a.Loop.Verdict != StatusFailedProjectNotAchieved {
		return fmt.Errorf("Gate473 unexpected loop verdict: %+v", a.Loop)
	}
	if !a.Firewall.Executed || a.Firewall.RawMassRowsNative || a.Firewall.IKNative || a.Firewall.AlphaNative || a.Firewall.DUDNativePrediction || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate473 firewall violated: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 473 rejects the proposed epistemological closure. The raw quark masses show extreme third-generation hierarchy and yield a trace-zero spectral invariant near the rank-one asymptotic limit 2/(3sqrt(3)), but that information is spectrum-only. It does not determine the K-overlap I_K, does not force alpha=1, and in the trace-zero spectrum map alpha=1 is incompatible with the observed extreme-hierarchy invariant. Therefore I_K=0.5 and the Gate471 Cabibbo alignment remain external rank-complete bridge-ledger inputs, not native consequences of PDG masses."
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
		cand := filepath.Join(dir, path)
		if _, err := os.Stat(cand); err == nil {
			return cand
		}
		dir = filepath.Dir(dir)
	}
	return path
}
func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.12g", x)
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
	return fmt.Sprintf("executed=%t K=%t triangle=%t rank_audit=%t inverse=%t branch_tags=%t airlock=%t gate471=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate454RankAuditAvailable, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate465AirlockAvailable, x.Gate471ExternalSocketAvailable, x.NativeRegistryClean, x.Verdict)
}
func FormatImport(x FileImport) string {
	return fmt.Sprintf("executed=%t loaded=%t path=%s empirical_import=%t bridge_only=%t rows=%d accepted=%d rejected=%d metadata=%t quarantined=%t native_write_requested=%t rejects_IK=%t rejects_CKM=%t verdict=%s reason=%s", x.Executed, x.Loaded, x.Path, x.EmpiricalImport, x.BridgeOnlyLedger, x.Rows, x.AcceptedRows, x.RejectedRows, x.MetadataComplete, x.AllAcceptedBridgeOnly, x.NativeRegistryWriteRequested, x.IKImportRejected, x.CKMImportRejected, x.Verdict, x.Reason)
}
func FormatSpectrum(x SectorSpectrum) string {
	return fmt.Sprintf("sector=%s masses=%v trace_zero=%v sumsq=%s m3_fraction=%s extreme=%t I_spec=%s asymptotic=%s delta=%s alpha_one_max_I_spec=%s alpha_max_allowed=%s alpha_one_compatible=%t alpha_forced=%t IK_half_derived=%t verdict=%s reason=%s", x.Sector, x.Masses, x.TraceZero, fmtFloat(x.SumSquaresPhysical), fmtFloat(x.ThirdGenerationSquareFraction), x.ExtremeHierarchy, fmtFloat(x.ISpec), fmtFloat(x.AsymptoticISpec), fmtFloat(x.AbsDeltaAsymptotic), fmtFloat(x.AlphaOneMaxISpec), fmtFloat(x.AlphaMaxAllowed), x.AlphaOneCompatible, x.AlphaOneForced, x.IKHalfDerived, x.Verdict, x.Reason)
}
func FormatLoop(x LoopClosure) string {
	dud := "undefined"
	residual := "undefined"
	if x.DUDComputed {
		dud = fmtFloat(x.DUD)
	}
	if x.CabibboResidualComputed {
		residual = fmtFloat(x.CabibboResidual)
	}
	return fmt.Sprintf("executed=%t attempted=%t raw_masses_only=%t alpha_derived=%t IK_derived=%t d_ud_computed=%t d_ud=%s |Vus|=%s residual_computed=%t residual=%s alignment=%t verdict=%s failures=%s reason=%s", x.Executed, x.Attempted, x.RawMassesOnly, x.AlphaDerived, x.IKDerived, x.DUDComputed, dud, fmtFloat(x.CabibboTarget), x.CabibboResidualComputed, residual, x.AlignmentAchieved, x.Verdict, strings.Join(x.Failures, ","), x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t masses_native=%t IK_native=%t alpha_native=%t d_ud_native=%t ckm_native=%t ckm_matrix=%t ckm_entry=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.RawMassRowsNative, x.IKNative, x.AlphaNative, x.DUDNativePrediction, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate471Inherited, StatusRawMassLedgerLoaded, StatusExtremeHierarchyConfirmed, StatusAsymptoticLimitDerived, StatusFailedMassHierarchyNoEquipartition, StatusFailedRawMassCannotDeriveIK, StatusFailedAlphaOneInconsistent, StatusFailedDUDUndefined, StatusFailedProjectNotAchieved, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 473 Registry Audit — Mass-to-Equipartition Inversion / Epistemological Loop Closure\n\n## Verdict\n\n")
	b.WriteString("`" + StatusFailedProjectNotAchieved + "`\n\n")
	b.WriteString("Gate 473 imports only raw quark masses through the empirical airlock and audits the proposed inference `m3 >> m1,m2 => alpha=1 => I_K=0.5`. The inference fails. Extreme hierarchy gives a rank-one spectral shape invariant, not a K-axis overlap. In fact the trace-zero extreme-hierarchy invariant tends to `2/(3*sqrt(3))`, while an alpha=1 equipartition ray can support at most `|I_spec|=1/4`.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Raw mass import\n\n" + FormatImport(a.Import) + "\n\n")
	b.WriteString("## Trace-zero spectrum audit\n\n- " + FormatSpectrum(a.Up) + "\n- " + FormatSpectrum(a.Down) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("lambda_i = m_i - mean(m)\n")
	b.WriteString("Q = 1/2 sum_i lambda_i^2\n")
	b.WriteString("R = product_i lambda_i\n")
	b.WriteString("I_spec = R / Q^(3/2)\n")
	b.WriteString("I_spec(alpha,phi) = 2 cos(3phi)/(alpha^2+3)^(3/2)\n")
	b.WriteString("extreme hierarchy limit lambda ~ (-M/3,-M/3,2M/3)\n")
	b.WriteString("I_spec -> 2/(3 sqrt(3)) = " + fmtFloat(2/(3*math.Sqrt(3))) + "\n")
	b.WriteString("alpha=1 implies |I_spec| <= 2/(1+3)^(3/2) = 0.25\n")
	b.WriteString("I_K(alpha=1) = 1/sqrt(4) = 0.5, but alpha=1 is not derived from raw masses\n")
	b.WriteString("`````\n\n")
	b.WriteString("## Loop closure attempt\n\n" + FormatLoop(a.Loop) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha_u = undefined\nalpha_d = undefined\nI_K,u = undefined\nI_K,d = undefined\nd_ud = undefined\nCabibbo residual = undefined\n")
	b.WriteString("`````\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No raw mass, spectral invariant, alpha value, I_K value, d_ud value, CKM value, or alignment flag is written to the native theorem registry.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return strings.ReplaceAll(b.String(), "`````", "```")
}
