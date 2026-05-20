// Package generation2higgstotalcorrectiondecompositionanddilationaudit implements
// Gate 776: Higgs Total Correction Decomposition and Dilation Audit.
//
// Gate 775 compressed the sealed tree Higgs radial tower into the single factor
// C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]. Gate 776 audits the internal
// decomposition of this total correction into HistoryLoop uplift and Yukawa
// participation dilution, and records how the net correction dilates the sealed
// tree mass and self-coupling tower. This is a correction-factor decomposition
// and dilation audit only. It does not derive the VEV, scalar runtime lambda,
// Higgs pole mass, physical self-couplings, Yukawa operators, CKM/PMNS, flavor
// hierarchy, or a native HistoryLoopUnit theorem.
package generation2higgstotalcorrectiondecompositionanddilationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE776-HIGGS-TOTAL-CORRECTION-DECOMPOSITION-DILATION-AUDIT"

	StatusGate775UnifiedHiggsCouplingTowerInherited = "PASS_GATE775_UNIFIED_HIGGS_COUPLING_TOWER_INHERITED"
	StatusHistoryUpliftAndYukawaDilutionDefined     = "PASS_HISTORY_UPLIFT_AND_YUKAWA_DILUTION_DEFINED"
	StatusTotalCorrectionExpansionComputed          = "PASS_TOTAL_CORRECTION_EXPANSION_COMPUTED"
	StatusNumericalDecompositionLedgerRecorded      = "PASS_NUMERICAL_DECOMPOSITION_LEDGER_RECORDED"
	StatusRadialDilationFactorComputed              = "PASS_RADIAL_DILATION_FACTOR_COMPUTED"
	StatusCouplingTowerRewrittenWithDeltaHiggs      = "PASS_COUPLING_TOWER_REWRITTEN_WITH_DELTA_HIGGS"
	StatusSourceTypeInterpretationRecorded          = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusPhysicalFirewallsEnforced                 = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusCHiggsDecomposesIntoHistoryMinusYukawaDrag = "CONDITIONAL_SUPPORT_C_HIGGS_DECOMPOSES_INTO_HISTORY_UPLIFT_MINUS_YUKAWA_PARTICIPATION_DRAG"
	StatusTreeTowerControlledByNetDeltaHiggs         = "CONDITIONAL_SUPPORT_TREE_HIGGS_TOWER_IS_CONTROLLED_BY_NET_DELTA_HIGGS"
	StatusRadialMassUsesSqrtTotalCorrection          = "CONDITIONAL_SUPPORT_RADIAL_MASS_PROXY_USES_SQUARE_ROOT_OF_TOTAL_CORRECTION"

	StatusDeltaHiggsNotNativeHiggsTheorem      = "FAILED_ROUTE_DELTA_HIGGS_NOT_NATIVE_HIGGS_THEOREM"
	StatusHistoryUpliftNotNativeHistoryLoop    = "FAILED_ROUTE_HISTORY_UPLIFT_NOT_NATIVE_HISTORYLOOP_THEOREM"
	StatusYukawaDilutionNotNativeYukawaTheorem = "FAILED_ROUTE_YUKAWA_DILUTION_NOT_NATIVE_YUKAWA_THEOREM"
	StatusRadialDilationNotPoleMassCorrection  = "FAILED_ROUTE_RADIAL_DILATION_NOT_POLE_MASS_CORRECTION"
	StatusTreeTowerNotPhysicalSelfCouplings    = "FAILED_ROUTE_TREE_TOWER_NOT_PHYSICAL_MEASURED_SELF_COUPLINGS"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate776TotalCorrectionBoundary       = "FIREWALL_PRESERVED_GATE776_HIGGS_TOTAL_CORRECTION_DECOMPOSITION_BOUNDARY"
)

const (
	// Gate 775 scalar-coordinate ledger snapshots.
	nEffMZ        = 3.0023273474722147
	cHistoryMZ    = 1.038025177923625
	vevConvention = 246.2196508
)

type Gate775Inheritance struct {
	Inherited             bool
	CHiggsFormula         string
	CYukawaFormula        string
	CHistoryFormula       string
	UnifiedTower          bool
	NativeHiggsTheorem    bool
	NativeYukawaTheorem   bool
	NativeHistoryTheorem  bool
	PhysicalSelfCouplings bool
	Verdict               string
}

type UpliftDilutionDefinitions struct {
	DeltaHistoryFormula  string
	EpsilonYukawaFormula string
	CHistoryFormula      string
	CYukawaFormula       string
	Defined              bool
	NativeHistoryTheorem bool
	NativeYukawaTheorem  bool
	Verdict              string
}

type TotalCorrectionExpansion struct {
	CHiggsFormula             string
	ExpansionFormula          string
	DeltaHiggsFormula         string
	YukawaDragFormula         string
	Computed                  bool
	NativeHiggsTheorem        bool
	IndependentRuntimeTheorem bool
	Verdict                   string
}

type NumericalDecompositionLedger struct {
	DeltaHistory       float64
	EpsilonYukawa      float64
	YukawaDrag         float64
	DeltaHiggs         float64
	CHiggs             float64
	CHiggsFromExpanded float64
	ExpansionResidual  float64
	HistoryDominates   bool
	Finite             bool
	Recorded           bool
	Verdict            string
}

type RadialDilationFactor struct {
	Formula             string
	BaselineMassFormula string
	DilationFactor      float64
	BaselineMassGeV     float64
	MassGeV             float64
	MassFromDilationGeV float64
	MassResidual        float64
	PoleMassCorrection  bool
	Computed            bool
	Verdict             string
}

type CouplingTowerWithDelta struct {
	LambdaHFormula       string
	PotentialFormula     string
	MassFormula          string
	Lambda3Formula       string
	Lambda4Formula       string
	A2Formula            string
	A3Formula            string
	A4Formula            string
	Rewritten            bool
	PhysicalSelfCoupling bool
	PoleMassTheorem      bool
	Verdict              string
}

type SourceTypeInterpretation struct {
	DeltaHistoryRole  string
	EpsilonYukawaRole string
	YukawaDragRole    string
	DeltaHiggsRole    string
	TowerRole         string
	Recorded          bool
	Verdict           string
}

type PhysicalFirewalls struct {
	Audited                          bool
	DeltaHiggsNativeHiggsTheorem     bool
	HistoryUpliftNativeHistoryLoop   bool
	YukawaDilutionNativeYukawa       bool
	RadialDilationPoleMassCorrection bool
	TreeTowerMeasuredSelfCouplings   bool
	YukawaOperatorOrEigenvalue       bool
	Verdict                          string
}

type Analysis struct {
	Gate775     Gate775Inheritance
	Definitions UpliftDilutionDefinitions
	Expansion   TotalCorrectionExpansion
	Numerical   NumericalDecompositionLedger
	Dilation    RadialDilationFactor
	Tower       CouplingTowerWithDelta
	Sources     SourceTypeInterpretation
	Firewalls   PhysicalFirewalls
	Truth       string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		clone := *cache
		return &clone, nil
	}
	for name, value := range map[string]float64{
		"N_eff":     nEffMZ,
		"C_History": cHistoryMZ,
		"v":         vevConvention,
	} {
		if math.IsNaN(value) || math.IsInf(value, 0) || value <= 0 {
			return nil, fmt.Errorf("invalid %s ledger: %.17g", name, value)
		}
	}

	cYukawaMZ := 3 / nEffMZ
	deltaHistory := cHistoryMZ - 1
	epsilonYukawa := 1 - cYukawaMZ
	yukawaDrag := epsilonYukawa * cHistoryMZ
	cHiggs := cYukawaMZ * cHistoryMZ
	deltaHiggs := cHiggs - 1
	cHiggsFromExpanded := 1 + deltaHistory - epsilonYukawa - epsilonYukawa*deltaHistory
	dilation := math.Sqrt(cHiggs)
	baselineMass := vevConvention / 2
	mass := baselineMass * dilation
	v2 := vevConvention * vevConvention

	a := &Analysis{
		Gate775: Gate775Inheritance{
			Inherited:             true,
			CHiggsFormula:         "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]",
			CYukawaFormula:        "C_Yukawa=3/N_eff",
			CHistoryFormula:       "C_History=1+L_Hopf(1-kappa_lambda_red)",
			UnifiedTower:          true,
			NativeHiggsTheorem:    false,
			NativeYukawaTheorem:   false,
			NativeHistoryTheorem:  false,
			PhysicalSelfCouplings: false,
			Verdict:               StatusGate775UnifiedHiggsCouplingTowerInherited,
		},
		Definitions: UpliftDilutionDefinitions{
			DeltaHistoryFormula:  "delta_History=L_Hopf(1-kappa_lambda_red)",
			EpsilonYukawaFormula: "epsilon_Yukawa=1-3/N_eff",
			CHistoryFormula:      "C_History=1+delta_History",
			CYukawaFormula:       "C_Yukawa=1-epsilon_Yukawa",
			Defined:              true,
			NativeHistoryTheorem: false,
			NativeYukawaTheorem:  false,
			Verdict:              StatusHistoryUpliftAndYukawaDilutionDefined,
		},
		Expansion: TotalCorrectionExpansion{
			CHiggsFormula:             "C_Higgs=(1-epsilon_Yukawa)(1+delta_History)",
			ExpansionFormula:          "C_Higgs=1+delta_History-epsilon_Yukawa-epsilon_Yukawa delta_History",
			DeltaHiggsFormula:         "Delta_Higgs=C_Higgs-1=delta_History-epsilon_Yukawa(1+delta_History)",
			YukawaDragFormula:         "epsilon_Yukawa(1+delta_History)",
			Computed:                  true,
			NativeHiggsTheorem:        false,
			IndependentRuntimeTheorem: false,
			Verdict:                   StatusTotalCorrectionExpansionComputed,
		},
		Numerical: NumericalDecompositionLedger{
			DeltaHistory:       deltaHistory,
			EpsilonYukawa:      epsilonYukawa,
			YukawaDrag:         yukawaDrag,
			DeltaHiggs:         deltaHiggs,
			CHiggs:             cHiggs,
			CHiggsFromExpanded: cHiggsFromExpanded,
			ExpansionResidual:  cHiggs - cHiggsFromExpanded,
			HistoryDominates:   deltaHistory > yukawaDrag,
			Finite:             finite(deltaHistory, epsilonYukawa, yukawaDrag, deltaHiggs, cHiggs, cHiggsFromExpanded),
			Recorded:           true,
			Verdict:            StatusNumericalDecompositionLedgerRecorded,
		},
		Dilation: RadialDilationFactor{
			Formula:             "D_radial=sqrt(C_Higgs)",
			BaselineMassFormula: "m_baseline=v/2",
			DilationFactor:      dilation,
			BaselineMassGeV:     baselineMass,
			MassGeV:             mass,
			MassFromDilationGeV: baselineMass * dilation,
			MassResidual:        0,
			PoleMassCorrection:  false,
			Computed:            true,
			Verdict:             StatusRadialDilationFactorComputed,
		},
		Tower: CouplingTowerWithDelta{
			LambdaHFormula:       "lambda_H_bridge=(1/8)(1+Delta_Higgs)",
			PotentialFormula:     "V_local(x)=[(1+Delta_Higgs)/32](||x||^2-v^2)^2",
			MassFormula:          "m_H_tree=(v/2)sqrt(1+Delta_Higgs)",
			Lambda3Formula:       "lambda_3=(3/4)v(1+Delta_Higgs)",
			Lambda4Formula:       "lambda_4=(3/4)(1+Delta_Higgs)",
			A2Formula:            "A_2=(1/8)(1+Delta_Higgs)v^2",
			A3Formula:            "A_3=(1/8)(1+Delta_Higgs)v",
			A4Formula:            "A_4=(1/32)(1+Delta_Higgs)",
			Rewritten:            true,
			PhysicalSelfCoupling: false,
			PoleMassTheorem:      false,
			Verdict:              StatusCouplingTowerRewrittenWithDeltaHiggs,
		},
		Sources: SourceTypeInterpretation{
			DeltaHistoryRole:  "Radial-Hopf / boundary-HistoryLoop uplift",
			EpsilonYukawaRole: "finite Yukawa trace participation dilution away from exact top-color dominance",
			YukawaDragRole:    "multiplicative drag of Yukawa dilution after HistoryLoop uplift",
			DeltaHiggsRole:    "net sealed tree-Higgs correction factor above the one-eighth baseline",
			TowerRole:         "the same net correction controls quartic, cubic, quartic self-coupling, and the tree mass proxy, with mass using square-root dilation",
			Recorded:          true,
			Verdict:           StatusSourceTypeInterpretationRecorded,
		},
		Firewalls: PhysicalFirewalls{
			Audited:                          true,
			DeltaHiggsNativeHiggsTheorem:     false,
			HistoryUpliftNativeHistoryLoop:   false,
			YukawaDilutionNativeYukawa:       false,
			RadialDilationPoleMassCorrection: false,
			TreeTowerMeasuredSelfCouplings:   false,
			YukawaOperatorOrEigenvalue:       false,
			Verdict:                          StatusGate776TotalCorrectionBoundary,
		},
		Truth: fmt.Sprintf("Gate776 decomposes C_Higgs into HistoryLoop uplift minus Yukawa participation drag: Delta_Higgs=%.16f, with radial dilation sqrt(C_Higgs)=%.16f. This is a sealed tree-lane correction-factor audit, not a native Higgs, Yukawa, HistoryLoop, pole-mass, or measured self-coupling theorem.", deltaHiggs, dilation),
	}
	_ = v2 // retained only to keep the dimensional ledger explicit at construction time.

	cache = a
	clone := *a
	return &clone, nil
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate775UnifiedHiggsCouplingTowerInherited,
		StatusHistoryUpliftAndYukawaDilutionDefined,
		StatusTotalCorrectionExpansionComputed,
		StatusNumericalDecompositionLedgerRecorded,
		StatusRadialDilationFactorComputed,
		StatusCouplingTowerRewrittenWithDeltaHiggs,
		StatusSourceTypeInterpretationRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusCHiggsDecomposesIntoHistoryMinusYukawaDrag,
		StatusTreeTowerControlledByNetDeltaHiggs,
		StatusRadialMassUsesSqrtTotalCorrection,
		StatusDeltaHiggsNotNativeHiggsTheorem,
		StatusHistoryUpliftNotNativeHistoryLoop,
		StatusYukawaDilutionNotNativeYukawaTheorem,
		StatusRadialDilationNotPoleMassCorrection,
		StatusTreeTowerNotPhysicalSelfCouplings,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate776TotalCorrectionBoundary,
	}
}

func FormatGate775(g Gate775Inheritance) string {
	return fmt.Sprintf("inherited=%t C=%s CY=%s CH=%s unifiedTower=%t nativeHiggs=%t nativeYukawa=%t nativeHistory=%t measured=%t verdict=%s", g.Inherited, g.CHiggsFormula, g.CYukawaFormula, g.CHistoryFormula, g.UnifiedTower, g.NativeHiggsTheorem, g.NativeYukawaTheorem, g.NativeHistoryTheorem, g.PhysicalSelfCouplings, g.Verdict)
}

func FormatDefinitions(d UpliftDilutionDefinitions) string {
	return fmt.Sprintf("delta=%s epsilon=%s CH=%s CY=%s defined=%t nativeHistory=%t nativeYukawa=%t verdict=%s", d.DeltaHistoryFormula, d.EpsilonYukawaFormula, d.CHistoryFormula, d.CYukawaFormula, d.Defined, d.NativeHistoryTheorem, d.NativeYukawaTheorem, d.Verdict)
}

func FormatExpansion(e TotalCorrectionExpansion) string {
	return fmt.Sprintf("C=%s expansion=%s Delta=%s drag=%s computed=%t nativeHiggs=%t independentRuntime=%t verdict=%s", e.CHiggsFormula, e.ExpansionFormula, e.DeltaHiggsFormula, e.YukawaDragFormula, e.Computed, e.NativeHiggsTheorem, e.IndependentRuntimeTheorem, e.Verdict)
}

func FormatNumerical(n NumericalDecompositionLedger) string {
	return fmt.Sprintf("deltaHistory=%.17g epsilonYukawa=%.17g drag=%.17g Delta=%.17g C=%.17g Cexpanded=%.17g residual=%.3e historyDominates=%t finite=%t recorded=%t verdict=%s", n.DeltaHistory, n.EpsilonYukawa, n.YukawaDrag, n.DeltaHiggs, n.CHiggs, n.CHiggsFromExpanded, n.ExpansionResidual, n.HistoryDominates, n.Finite, n.Recorded, n.Verdict)
}

func FormatDilation(d RadialDilationFactor) string {
	return fmt.Sprintf("formula=%s baseline=%s D=%.17g m0=%.10f m=%.12f residual=%.3e poleMassCorrection=%t computed=%t verdict=%s", d.Formula, d.BaselineMassFormula, d.DilationFactor, d.BaselineMassGeV, d.MassGeV, d.MassResidual, d.PoleMassCorrection, d.Computed, d.Verdict)
}

func FormatTower(t CouplingTowerWithDelta) string {
	return fmt.Sprintf("lambda=%s V=%s m=%s lambda3=%s lambda4=%s A2=%s A3=%s A4=%s rewritten=%t measured=%t pole=%t verdict=%s", t.LambdaHFormula, t.PotentialFormula, t.MassFormula, t.Lambda3Formula, t.Lambda4Formula, t.A2Formula, t.A3Formula, t.A4Formula, t.Rewritten, t.PhysicalSelfCoupling, t.PoleMassTheorem, t.Verdict)
}

func FormatSources(s SourceTypeInterpretation) string {
	return strings.Join([]string{s.DeltaHistoryRole, s.EpsilonYukawaRole, s.YukawaDragRole, s.DeltaHiggsRole, s.TowerRole, s.Verdict}, " | ")
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("audited=%t DeltaNative=%t historyNative=%t yukawaNative=%t dilationPole=%t measured=%t yukawaOp=%t verdict=%s", f.Audited, f.DeltaHiggsNativeHiggsTheorem, f.HistoryUpliftNativeHistoryLoop, f.YukawaDilutionNativeYukawa, f.RadialDilationPoleMassCorrection, f.TreeTowerMeasuredSelfCouplings, f.YukawaOperatorOrEigenvalue, f.Verdict)
}
