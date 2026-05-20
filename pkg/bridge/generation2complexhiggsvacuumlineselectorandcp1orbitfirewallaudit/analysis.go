// Package generation2complexhiggsvacuumlineselectorandcp1orbitfirewallaudit implements
// Gate 762: Complex Higgs Vacuum Line Selector and CP1 Orbit Firewall Audit.
//
// Gate 761 refined P_rad from a primitive real rank-one projector into a
// gauge-fixed radial representative inside a complex Higgs vacuum line
// Pi_vac_C in K7+_J(n) ~= C^2. Gate 762 audits the deeper object: whether any
// current ASHA structure selects a point of CP1, i.e. a complex rank-one vacuum
// line in the Higgs socket after n supplies J_H(n). This is a selector and orbit
// firewall audit only. It does not derive electroweak symmetry breaking, radial
// gauge fixing, scalar runtime lambda, Higgs mass, pole mass, Yukawa operators,
// CKM/PMNS, flavor hierarchy, or a native HistoryLoopUnit theorem.
package generation2complexhiggsvacuumlineselectorandcp1orbitfirewallaudit

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE762-COMPLEX-HIGGS-VACUUM-LINE-SELECTOR-AND-CP1-ORBIT-FIREWALL-AUDIT"

	StatusGate761RadialProjectorRefinementInherited        = "PASS_GATE761_RADIAL_PROJECTOR_REFINEMENT_INHERITED"
	StatusComplexVacuumLinePromotedToPrimaryTarget         = "PASS_COMPLEX_VACUUM_LINE_PROMOTED_TO_PRIMARY_TARGET"
	StatusCP1OrbitGeometryRecorded                         = "PASS_CP1_ORBIT_GEOMETRY_RECORDED"
	StatusSelectorQuestionFormulated                       = "PASS_SELECTOR_QUESTION_FORMULATED"
	StatusCurrentSourceCandidateAuditCompleted             = "PASS_CURRENT_SOURCE_CANDIDATE_AUDIT_COMPLETED"
	StatusConstructedFromPRadButNotSelectedNativelyAudited = "PASS_CONSTRUCTED_FROM_P_RAD_BUT_NOT_SELECTED_NATIVELY_AUDITED"
	StatusRadialGaugeFixingMarkedSecondary                 = "PASS_RADIAL_GAUGE_FIXING_MARKED_SECONDARY"
	StatusHistoryLoopDependencyRefined                     = "PASS_HISTORYLOOP_DEPENDENCY_REFINED"
	StatusLayerSeparationAudited                           = "PASS_LAYER_SEPARATION_AUDITED"
	StatusPhysicalFirewallsEnforced                        = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusPiVacCRemainsComplexVacuumLineSeal               = "CONDITIONAL_SUPPORT_PI_VAC_C_REMAINS_COMPLEX_VACUUM_LINE_SEAL"
	StatusNoCurrentASHAStructureSelectsCP1Point            = "CONDITIONAL_SUPPORT_NO_CURRENT_ASHA_STRUCTURE_SELECTS_CP1_POINT"
	StatusPRadSelectionRequiresLinePlusGauge               = "CONDITIONAL_SUPPORT_P_RAD_SELECTION_REQUIRES_COMPLEX_LINE_PLUS_RADIAL_GAUGE"
	StatusRadialGaugeSecondaryAfterComplexLine             = "CONDITIONAL_SUPPORT_RADIAL_GAUGE_FIXING_IS_SECONDARY_AFTER_COMPLEX_LINE_SELECTION"
	StatusNoNativeComplexVacuumLineSelector                = "FAILED_ROUTE_NO_NATIVE_COMPLEX_VACUUM_LINE_SELECTOR"
	StatusNoNativeCP1BasePointSelector                     = "FAILED_ROUTE_NO_NATIVE_CP1_BASE_POINT_SELECTOR"
	StatusNSelectsComplexStructureNotCP1Point              = "FAILED_ROUTE_N_SELECTS_COMPLEX_STRUCTURE_NOT_CP1_POINT"
	StatusRhoPlusIsNoBiasStateNotLineSelector              = "FAILED_ROUTE_RHO_PLUS_IS_NO_BIAS_STATE_NOT_LINE_SELECTOR"
	StatusPRadCannotBeUsedAsNativeLineSelector             = "FAILED_ROUTE_P_RAD_CANNOT_BE_USED_AS_NATIVE_LINE_SELECTOR"
	StatusNoNativeRadialGaugeFixingSelector                = "FAILED_ROUTE_NO_NATIVE_RADIAL_GAUGE_FIXING_SELECTOR"
	StatusNoNativeEWSBTheorem                              = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem                   = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                     = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem              = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate762ComplexVacuumLineCP1Boundary              = "FIREWALL_PRESERVED_GATE762_COMPLEX_VACUUM_LINE_CP1_SELECTOR_BOUNDARY"
)

const (
	k7PlusRealDim       = 4
	k7PlusComplexDim    = 2
	complexLineRealRank = 2
	complexLineRank     = 1
	realRadialRank      = 1
	realPhaseRank       = 1
	cp1RealDim          = 2
	s3RealDim           = 3
	s1RealDim           = 1
)

type Gate761Inheritance struct {
	Inherited                       bool
	PriorPRadTyping                 string
	RefinedPRadTyping               string
	ComplexLineFormula              string
	ScalarVacuumSealSplit           string
	PRadGaugeFixedInsideComplexLine bool
	ComplexLineNotNativelySelected  bool
	RadialGaugeNotNativelySelected  bool
	RealRadialWeight                float64
	FullComplexLineWeight           float64
	ActiveHistoryUsesRealRadialHalf bool
	Verdict                         string
}

type PrimaryTarget struct {
	PreviousTarget        string
	NewPrimaryTarget      string
	Reason                string
	ComplexLineBeforePRad bool
	RadialGaugeAfterLine  bool
	LineLivesInK7PlusJ    bool
	LineIsCP1Point        bool
	Verdict               string
}

type CP1OrbitGeometry struct {
	Socket                     string
	ComplexLines               string
	CP1RealDimension           int
	UnitRepresentatives        string
	S3RealDimension            int
	HopfFibration              string
	FiberDimension             int
	BasePointSelectsLine       bool
	FiberGaugeSelectsRadialRep bool
	Recorded                   bool
	Verdict                    string
}

type SelectorQuestion struct {
	Question                  string
	RequiredObject            string
	RequiredRankR             int
	RequiredRankC             int
	RequiresJH                bool
	RequiresCP1BasePoint      bool
	DoesNotRequireRadialGauge bool
	Formulated                bool
	Verdict                   string
}

type SourceCandidate struct {
	Name               string
	SelectsJH          bool
	SelectsCP1Point    bool
	SelectsRadialGauge bool
	Reason             string
}

type SourceCandidateAudit struct {
	Candidates                      []SourceCandidate
	Completed                       bool
	NativeComplexVacuumLineSelector bool
	NativeCP1BasePointSelector      bool
	NativeRadialGaugeSelector       bool
	NSelectsJHOnly                  bool
	RhoPlusNoBias                   bool
	PRadAssumesLineAndGauge         bool
	AllCurrentCandidatesFailCP1     bool
	Verdict                         string
}

type PRadConstructibilityAudit struct {
	CanConstructPiFromSuppliedPRad bool
	ConstructionFormula            string
	ConstructionDependsOnSeal      bool
	ConstructionIsSelectionTheorem bool
	PRadMayNotBeUsedAsNativeCause  bool
	NativeLineSelectorCertified    bool
	Verdict                        string
}

type RadialGaugeHierarchy struct {
	ComplexLineSeal            string
	RadialGaugeFixingSeal      string
	LineSelectionPrecedesGauge bool
	GaugeWithoutLineIllTyped   bool
	PRadRequiresBothChoices    bool
	SecondaryMarked            bool
	Verdict                    string
}

type HistoryLoopDependency struct {
	ActiveLHopfFormula      string
	DependsOnComplexLine    bool
	DependsOnRadialGauge    bool
	RealRadialWeight        float64
	ComplexLineWeight       float64
	FullComplexLineRejected bool
	UnsolvedObjects         []string
	Refined                 bool
	Verdict                 string
}

type LayerSeparation struct {
	NLayer                 string
	CP1Layer               string
	RadialGaugeLayer       string
	HistoryLayer           string
	YukawaLayer            string
	ScalarRuntimeLayer     string
	ObjectsNotSameOperator bool
	Audited                bool
	Verdict                string
}

type Firewalls struct {
	Audited                           bool
	PiVacCNativeVacuumTheorem         bool
	CP1PointNativeEWSBTheorem         bool
	RadialGaugeFixingNativeTheorem    bool
	LHopfNativeHistoryLoopTheorem     bool
	ComplexLineHiggsMassTheorem       bool
	ComplexLineYukawaTheorem          bool
	ScalarRuntimeIndependentTheorem   bool
	HiggsMassOrPoleMassTheorem        bool
	YukawaOperatorOrEigenvalueTheorem bool
	Verdict                           string
}

type Analysis struct {
	Gate761        Gate761Inheritance
	Primary        PrimaryTarget
	Orbit          CP1OrbitGeometry
	Question       SelectorQuestion
	SourceAudit    SourceCandidateAudit
	Constructible  PRadConstructibilityAudit
	GaugeHierarchy RadialGaugeHierarchy
	HistoryLoop    HistoryLoopDependency
	Layers         LayerSeparation
	Firewalls      Firewalls
	Truth          string
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
	a := &Analysis{
		Gate761: Gate761Inheritance{
			Inherited:                       true,
			PriorPRadTyping:                 "supplied real rank-one radial projector",
			RefinedPRadTyping:               "GaugeFixedRadialRepresentativeSeal inside ComplexVacuumLineSeal",
			ComplexLineFormula:              "Pi_vac_C = P_rad + P_phase, P_phase = J_H(n) P_rad J_H(n)^(-1)",
			ScalarVacuumSealSplit:           "ScalarVacuumDirectionSeal = (ComplexVacuumLineSeal, RadialGaugeFixingSeal)",
			PRadGaugeFixedInsideComplexLine: true,
			ComplexLineNotNativelySelected:  true,
			RadialGaugeNotNativelySelected:  true,
			RealRadialWeight:                0.25,
			FullComplexLineWeight:           0.50,
			ActiveHistoryUsesRealRadialHalf: true,
			Verdict:                         StatusGate761RadialProjectorRefinementInherited,
		},
		Primary: PrimaryTarget{
			PreviousTarget:        "P_rad / L_Hopf",
			NewPrimaryTarget:      "Pi_vac_C",
			Reason:                "P_rad is a gauge-fixed representative; the deeper pre-gauge object is the complex Higgs vacuum line in K7+_J(n).",
			ComplexLineBeforePRad: true,
			RadialGaugeAfterLine:  true,
			LineLivesInK7PlusJ:    true,
			LineIsCP1Point:        true,
			Verdict:               StatusComplexVacuumLinePromotedToPrimaryTarget,
		},
		Orbit: CP1OrbitGeometry{
			Socket:                     "K7+_J(n) ~= C^2",
			ComplexLines:               "CP1 = P(C^2) = U(2)/(U(1)xU(1))",
			CP1RealDimension:           cp1RealDim,
			UnitRepresentatives:        "S3",
			S3RealDimension:            s3RealDim,
			HopfFibration:              "S1 -> S3 -> CP1",
			FiberDimension:             s1RealDim,
			BasePointSelectsLine:       true,
			FiberGaugeSelectsRadialRep: true,
			Recorded:                   true,
			Verdict:                    StatusCP1OrbitGeometryRecorded,
		},
		Question: SelectorQuestion{
			Question:                  "Does any current ASHA structure select a point in CP1?",
			RequiredObject:            "complex rank-one projector Pi_vac_C in K7+_J(n)",
			RequiredRankR:             complexLineRealRank,
			RequiredRankC:             complexLineRank,
			RequiresJH:                true,
			RequiresCP1BasePoint:      true,
			DoesNotRequireRadialGauge: true,
			Formulated:                true,
			Verdict:                   StatusSelectorQuestionFormulated,
		},
		SourceAudit: buildSourceAudit(),
		Constructible: PRadConstructibilityAudit{
			CanConstructPiFromSuppliedPRad: true,
			ConstructionFormula:            "Pi_vac_C = P_rad + J_H(n) P_rad J_H(n)^(-1)",
			ConstructionDependsOnSeal:      true,
			ConstructionIsSelectionTheorem: false,
			PRadMayNotBeUsedAsNativeCause:  true,
			NativeLineSelectorCertified:    false,
			Verdict:                        strings.Join([]string{StatusConstructedFromPRadButNotSelectedNativelyAudited, StatusPRadCannotBeUsedAsNativeLineSelector}, "; "),
		},
		GaugeHierarchy: RadialGaugeHierarchy{
			ComplexLineSeal:            "ComplexVacuumLineSeal",
			RadialGaugeFixingSeal:      "RadialGaugeFixingSeal",
			LineSelectionPrecedesGauge: true,
			GaugeWithoutLineIllTyped:   true,
			PRadRequiresBothChoices:    true,
			SecondaryMarked:            true,
			Verdict:                    strings.Join([]string{StatusRadialGaugeFixingMarkedSecondary, StatusRadialGaugeSecondaryAfterComplexLine}, "; "),
		},
		HistoryLoop: HistoryLoopDependency{
			ActiveLHopfFormula:      "L_Hopf = (1/(2*pi)) Tr(rho_plus P_rad) = (1/(2*pi))(1/4)",
			DependsOnComplexLine:    true,
			DependsOnRadialGauge:    true,
			RealRadialWeight:        0.25,
			ComplexLineWeight:       0.50,
			FullComplexLineRejected: true,
			UnsolvedObjects:         []string{"ComplexVacuumLineSeal", "RadialGaugeFixingSeal"},
			Refined:                 true,
			Verdict:                 StatusHistoryLoopDependencyRefined,
		},
		Layers: LayerSeparation{
			NLayer:                 "twistor selector n selects J_H(n), not CP1",
			CP1Layer:               "ComplexVacuumLineSeal selects Pi_vac_C in P(K7+_J(n))",
			RadialGaugeLayer:       "RadialGaugeFixingSeal selects P_rad inside Pi_vac_C",
			HistoryLayer:           "L_Hopf uses the real radial event weight after gauge fixing",
			YukawaLayer:            "N_eff remains finite trace participation, not a CP1 selector",
			ScalarRuntimeLayer:     "lambda_runtime_eff receives scalar collapsed factors only",
			ObjectsNotSameOperator: true,
			Audited:                true,
			Verdict:                StatusLayerSeparationAudited,
		},
		Firewalls: Firewalls{
			Audited:                           true,
			PiVacCNativeVacuumTheorem:         false,
			CP1PointNativeEWSBTheorem:         false,
			RadialGaugeFixingNativeTheorem:    false,
			LHopfNativeHistoryLoopTheorem:     false,
			ComplexLineHiggsMassTheorem:       false,
			ComplexLineYukawaTheorem:          false,
			ScalarRuntimeIndependentTheorem:   false,
			HiggsMassOrPoleMassTheorem:        false,
			YukawaOperatorOrEigenvalueTheorem: false,
			Verdict:                           StatusGate762ComplexVacuumLineCP1Boundary,
		},
		Truth: "Gate 762 finds no current native selector for the CP1 complex Higgs vacuum line; Pi_vac_C remains a ComplexVacuumLineSeal and radial gauge fixing remains a secondary seal.",
	}
	cache = a
	clone := *a
	return &clone, nil
}

func buildSourceAudit() SourceCandidateAudit {
	candidates := []SourceCandidate{
		{Name: "rho_plus", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "no-bias state on K7+; assigns event weights but selects no line"},
		{Name: "n", SelectsJH: true, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "selects complex structure J_H(n), not a complex line in K7+_J(n)"},
		{Name: "q", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "normalizes phase/hypercharge interface, not the CP1 base point"},
		{Name: "P_K7", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "selects the full K7 support, not a line in K7+"},
		{Name: "boundary scalars", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "provide scalar coordinates and wall responses, not a vector in K7+"},
		{Name: "Fano/quaternionic structure", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: false, Reason: "supplies twistor/U(2) socket structure, not a distinguished vacuum line"},
		{Name: "supplied P_rad", SelectsJH: false, SelectsCP1Point: false, SelectsRadialGauge: true, Reason: "can reconstruct a line only after assuming the sealed real representative; cannot serve as native cause of that line"},
	}
	return SourceCandidateAudit{
		Candidates:                      candidates,
		Completed:                       true,
		NativeComplexVacuumLineSelector: false,
		NativeCP1BasePointSelector:      false,
		NativeRadialGaugeSelector:       false,
		NSelectsJHOnly:                  true,
		RhoPlusNoBias:                   true,
		PRadAssumesLineAndGauge:         true,
		AllCurrentCandidatesFailCP1:     true,
		Verdict:                         strings.Join([]string{StatusCurrentSourceCandidateAuditCompleted, StatusNoCurrentASHAStructureSelectsCP1Point, StatusNoNativeComplexVacuumLineSelector, StatusNoNativeCP1BasePointSelector}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate761RadialProjectorRefinementInherited,
		StatusComplexVacuumLinePromotedToPrimaryTarget,
		StatusCP1OrbitGeometryRecorded,
		StatusSelectorQuestionFormulated,
		StatusCurrentSourceCandidateAuditCompleted,
		StatusConstructedFromPRadButNotSelectedNativelyAudited,
		StatusRadialGaugeFixingMarkedSecondary,
		StatusHistoryLoopDependencyRefined,
		StatusLayerSeparationAudited,
		StatusPhysicalFirewallsEnforced,
		StatusPiVacCRemainsComplexVacuumLineSeal,
		StatusNoCurrentASHAStructureSelectsCP1Point,
		StatusPRadSelectionRequiresLinePlusGauge,
		StatusRadialGaugeSecondaryAfterComplexLine,
		StatusNoNativeComplexVacuumLineSelector,
		StatusNoNativeCP1BasePointSelector,
		StatusNSelectsComplexStructureNotCP1Point,
		StatusRhoPlusIsNoBiasStateNotLineSelector,
		StatusPRadCannotBeUsedAsNativeLineSelector,
		StatusNoNativeRadialGaugeFixingSelector,
		StatusNoNativeEWSBTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate762ComplexVacuumLineCP1Boundary,
	}
}

func FormatGate761(v Gate761Inheritance) string {
	return fmt.Sprintf("inherited=%v P_rad=%s refined=%s line=%s split=%s weights(radial=%.2f complex=%.2f) active_radial=%v verdict=%s", v.Inherited, v.PriorPRadTyping, v.RefinedPRadTyping, v.ComplexLineFormula, v.ScalarVacuumSealSplit, v.RealRadialWeight, v.FullComplexLineWeight, v.ActiveHistoryUsesRealRadialHalf, v.Verdict)
}

func FormatPrimary(v PrimaryTarget) string {
	return fmt.Sprintf("previous=%s target=%s reason=%s line_before_prad=%v gauge_after_line=%v cp1=%v verdict=%s", v.PreviousTarget, v.NewPrimaryTarget, v.Reason, v.ComplexLineBeforePRad, v.RadialGaugeAfterLine, v.LineIsCP1Point, v.Verdict)
}

func FormatOrbit(v CP1OrbitGeometry) string {
	return fmt.Sprintf("socket=%s lines=%s dimCP1=%d reps=%s dimS3=%d hopf=%s base_selects_line=%v fiber_selects_radial=%v verdict=%s", v.Socket, v.ComplexLines, v.CP1RealDimension, v.UnitRepresentatives, v.S3RealDimension, v.HopfFibration, v.BasePointSelectsLine, v.FiberGaugeSelectsRadialRep, v.Verdict)
}

func FormatQuestion(v SelectorQuestion) string {
	return fmt.Sprintf("question=%q object=%s rankR=%d rankC=%d requiresJH=%v requiresCP1=%v no_radial_gauge_needed=%v verdict=%s", v.Question, v.RequiredObject, v.RequiredRankR, v.RequiredRankC, v.RequiresJH, v.RequiresCP1BasePoint, v.DoesNotRequireRadialGauge, v.Verdict)
}

func FormatSourceAudit(v SourceCandidateAudit) string {
	parts := make([]string, 0, len(v.Candidates))
	for _, c := range v.Candidates {
		parts = append(parts, fmt.Sprintf("%s(cp1=%v gauge=%v reason=%s)", c.Name, c.SelectsCP1Point, c.SelectsRadialGauge, c.Reason))
	}
	return fmt.Sprintf("completed=%v native_line=%v native_cp1=%v native_gauge=%v n_jh_only=%v rho_nobias=%v prad_assumes=%v all_fail_cp1=%v candidates=[%s] verdict=%s", v.Completed, v.NativeComplexVacuumLineSelector, v.NativeCP1BasePointSelector, v.NativeRadialGaugeSelector, v.NSelectsJHOnly, v.RhoPlusNoBias, v.PRadAssumesLineAndGauge, v.AllCurrentCandidatesFailCP1, strings.Join(parts, "; "), v.Verdict)
}

func FormatConstructible(v PRadConstructibilityAudit) string {
	return fmt.Sprintf("can_construct=%v formula=%s depends_on_seal=%v selection_theorem=%v prad_not_native_cause=%v native_line=%v verdict=%s", v.CanConstructPiFromSuppliedPRad, v.ConstructionFormula, v.ConstructionDependsOnSeal, v.ConstructionIsSelectionTheorem, v.PRadMayNotBeUsedAsNativeCause, v.NativeLineSelectorCertified, v.Verdict)
}

func FormatGaugeHierarchy(v RadialGaugeHierarchy) string {
	return fmt.Sprintf("line=%s gauge=%s line_precedes_gauge=%v gauge_without_line_illtyped=%v prad_requires_both=%v secondary=%v verdict=%s", v.ComplexLineSeal, v.RadialGaugeFixingSeal, v.LineSelectionPrecedesGauge, v.GaugeWithoutLineIllTyped, v.PRadRequiresBothChoices, v.SecondaryMarked, v.Verdict)
}

func FormatHistoryLoop(v HistoryLoopDependency) string {
	return fmt.Sprintf("formula=%s depends_line=%v depends_gauge=%v radial_weight=%.2f line_weight=%.2f reject_full_line=%v unsolved=%s verdict=%s", v.ActiveLHopfFormula, v.DependsOnComplexLine, v.DependsOnRadialGauge, v.RealRadialWeight, v.ComplexLineWeight, v.FullComplexLineRejected, strings.Join(v.UnsolvedObjects, "+"), v.Verdict)
}

func FormatLayers(v LayerSeparation) string {
	return fmt.Sprintf("n=%s cp1=%s gauge=%s history=%s yukawa=%s runtime=%s distinct=%v verdict=%s", v.NLayer, v.CP1Layer, v.RadialGaugeLayer, v.HistoryLayer, v.YukawaLayer, v.ScalarRuntimeLayer, v.ObjectsNotSameOperator, v.Verdict)
}

func FormatFirewalls(v Firewalls) string {
	return fmt.Sprintf("audited=%v pi_native=%v cp1_ewsb=%v gauge_native=%v lhopf_native=%v line_higgs=%v line_yukawa=%v runtime=%v pole=%v yukawa_operator=%v verdict=%s", v.Audited, v.PiVacCNativeVacuumTheorem, v.CP1PointNativeEWSBTheorem, v.RadialGaugeFixingNativeTheorem, v.LHopfNativeHistoryLoopTheorem, v.ComplexLineHiggsMassTheorem, v.ComplexLineYukawaTheorem, v.ScalarRuntimeIndependentTheorem, v.HiggsMassOrPoleMassTheorem, v.YukawaOperatorOrEigenvalueTheorem, v.Verdict)
}
