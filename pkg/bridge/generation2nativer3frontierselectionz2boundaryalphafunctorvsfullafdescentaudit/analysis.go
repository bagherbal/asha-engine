// Package generation2nativer3frontierselectionz2boundaryalphafunctorvsfullafdescentaudit
// implements Gate 911: Native R3 Frontier Selection — Z2 BoundaryAlpha
// Functor vs Full A_F Descent Audit.
//
// Gate 911 is intentionally classificatory and strategic. It follows Gate 910's
// frozen plateau, where the R3 trace ledger exists only under a Z2 BoundaryAlpha
// class seal. The gate does not derive alpha_B, update official ledgers, assign
// physical sectors, or enter generation/flavor/Yukawa-spectrum territory. It
// selects the next native-R3 rail: attack the native Z2 BoundaryAlpha functor
// before the deeper full A_F descent / spontaneous-orientation problem.
package generation2nativer3frontierselectionz2boundaryalphafunctorvsfullafdescentaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE911-NATIVE-R3-FRONTIER-SELECTION-Z2-BOUNDARY-ALPHA-FUNCTOR-VS-FULL-AF-DESCENT-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OperatorCHiggsDiagnostic  = 1.037220510866514

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	LinearDenom  = 10
	QuadDenom    = 72

	Gate910Classification = "R3_SEALED_Z2_EQUIVARIANT_TRACE_LEDGER_WITH_BOUNDARY_ALPHA_CLASS_SEAL_NATIVE_PROMOTION_BLOCKED"
	Gate910ShortStatus    = "R3_Z2_BOUNDARY_ALPHA_CLASS_SEAL_PLATEAU_NOT_NATIVE"
	Gate910Verdict        = "R3_SEALED_PLATEAU_CONFIRMED_NATIVE_R3_BLOCKED_BY_Z2_BOUNDARY_ALPHA_FUNCTOR"

	BoundaryAlphaFormula = "alpha_B=(3/10)s+(7/72)s^2"
	Z2PunctureClass      = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	TraceRowMultiset     = "{(3,1),(3,alpha_B(1-alpha_B)),(1,3 alpha_B^2)}"
	NEffFormula          = "N_eff^operator=3(1+alpha_B)^2/(1+alpha_B^2-2alpha_B^3+4alpha_B^4)"
	CYukawaFormula       = "C_Yukawa^operator=3/N_eff^operator"
	OrientedAlgebra      = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullAlgebra          = "A_F=C plus H plus M_3(C)"
	ReducedB2Response    = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"

	FrontierAName = "Frontier A — Native Z2 BoundaryAlpha Functor"
	FrontierBName = "Frontier B — Full A_F Descent / Spontaneous-Orientation Status"
	FrontierCName = "Frontier C — Generation / Flavor / Individual Yukawa Branch"

	MissingZ2BoundaryAlphaFunctor = "Z2EquivariantNeutralPunctureAirlockFunctor"
	NextGate                      = "NEXT_PRESSURE_GATE912_Z2_BOUNDARY_ALPHA_FUNCTOR_SOURCE_DECOMPOSITION_AUDIT"

	Classification      = "R3_FRONTIER_SELECTED_Z2_BOUNDARY_ALPHA_FUNCTOR_BEFORE_FULL_AF_DESCENT"
	ShortStatus         = "R3_NEXT_RAIL_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	FinalTruth          = "FRONTIER_A_SELECTED_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR_FIRST"
	StrategicConclusion = "Native R3 pressure should now attack the Z2 BoundaryAlpha functor first; full A_F descent remains required but should follow after the alpha source is sharpened."

	StatusGate910Inherited       = "PASS_GATE910_R3_Z2_CLASS_SEAL_PLATEAU_INHERITED"
	StatusNoLoopBack             = "PASS_BRANCH_DOES_NOT_LOOP_BACK_TO_PHASE_SIGN_REPRESENTATIVE_ALPHA_OR_SOCKET_ORDER"
	StatusFrontierASelected      = "PASS_FRONTIER_A_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR_SELECTED_FIRST"
	StatusFrontierBDeferred      = "PASS_FRONTIER_B_FULL_A_F_DESCENT_DEFERRED_BUT_RECORDED_AS_REQUIRED"
	StatusFrontierCDeferred      = "PASS_FRONTIER_C_GENERATION_FLAVOR_YUKAWA_CLASSIFIED_R4_OR_LATER"
	StatusTraceDiagnosticsSealed = "PASS_TRACE_DIAGNOSTICS_REMAIN_SEALED_AND_NON_OFFICIAL"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE911_NOT_NATIVE_R3"

	SupportNativeR3TwoBlockers           = "CONDITIONAL_SUPPORT_NATIVE_R3_FRONTIER_HAS_TWO_CORE_BLOCKERS"
	SupportFrontierAZ2AlphaPrimary       = "CONDITIONAL_SUPPORT_FRONTIER_A_Z2_BOUNDARY_ALPHA_FUNCTOR_IS_PRIMARY"
	SupportFrontierADirectTraceLedger    = "CONDITIONAL_SUPPORT_FRONTIER_A_DIRECTLY_TARGETS_NATIVE_R3_TRACE_LEDGER"
	SupportFrontierAControlsAlphaWeights = "CONDITIONAL_SUPPORT_FRONTIER_A_CONTROLS_ALPHA_B_AND_SOCKET_TRACE_WEIGHTS"
	SupportFrontierAPrimaryBlocker       = "CONDITIONAL_SUPPORT_FRONTIER_A_IS_THE_PRIMARY_NATIVE_R3_BLOCKER"
	SupportFrontierBRequired             = "CONDITIONAL_SUPPORT_FRONTIER_B_IS_REQUIRED_FOR_FULL_NATIVE_FINITE_SECTOR_THEOREM"
	SupportFrontierBSecond               = "CONDITIONAL_SUPPORT_FRONTIER_B_FULL_A_F_DESCENT_IS_SECONDARY_BUT_REQUIRED"
	SupportFrontierBAfterAlpha           = "CONDITIONAL_SUPPORT_FRONTIER_B_SHOULD_FOLLOW_AFTER_ALPHA_SOURCE_IS_SHARPENED"
	SupportFrontierCR4Later              = "CONDITIONAL_SUPPORT_FRONTIER_C_GENERATION_FLAVOR_YUKAWA_IS_R4_OR_LATER"
	SupportNextBranchZ2Alpha             = "CONDITIONAL_SUPPORT_NEXT_BRANCH_SHOULD_ATTACK_Z2_BOUNDARY_ALPHA_FUNCTOR"
	SupportNoLoopBack                    = "CONDITIONAL_SUPPORT_BRANCH_SHOULD_NOT_LOOP_BACK_TO_PHASE_SIGN_OR_REPRESENTATIVE_ALPHA"

	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureFrontierANotCertified          = "FAILED_ROUTE_FRONTIER_A_NOT_CERTIFIED_YET"
	FailureNoNativeZ2BoundaryAlphaFunctor = "FAILED_ROUTE_NO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	FailureNoNativeZ2AirlockFunctor       = "FAILED_ROUTE_NO_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR"
	FailureNoNativeDegreeToZ2FlagFunctor  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeZ2CrossLane            = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureReducedB2NotNativeFunctional   = "FAILED_ROUTE_REDUCED_B2_RESPONSE_NOT_NATIVE_BOUNDARY_FUNCTIONAL"
	FailureAlphaSealedUntilFunctor        = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_UNTIL_Z2_BOUNDARY_ALPHA_FUNCTOR_EXISTS"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureAFOrientNotFullAF              = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureHiggsOrientationSealed         = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawa               = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedPlateau struct {
	Gate910Classification string
	Gate910ShortStatus    string
	Gate910Verdict        string
	R3TraceLedgerSealed   bool
	NativeR3              bool
	LoopBackPhase         bool
	LoopBackRepAlpha      bool
	LoopBackSocketOrder   bool
	Supports, Failures    []string
}

type TraceDiagnostics struct {
	BoundaryAlphaFormula string
	Z2PunctureClass      string
	TraceRows            string
	NEffFormula          string
	CYukawaFormula       string
	S                    float64
	Alpha                float64
	OperatorNEff         float64
	OperatorCYukawa      float64
	OperatorCHiggs       float64
	DiagnosticsOnly      bool
	OfficialUpdates      bool
	Supports, Failures   []string
}

type FrontierAssessment struct {
	Name                         string
	SelectedFirst                bool
	RequiredForNativeR3          bool
	DirectlyControlsAlpha        bool
	DirectlyControlsTraceWeights bool
	DirectlyControlsNEff         bool
	DirectlyControlsCYukawa      bool
	DirectlyControlsCHiggs       bool
	FullAFDescentProblem         bool
	GenerationFlavorYukawaBranch bool
	R4OrLater                    bool
	Deferred                     bool
	MissingObjects               []string
	Supports, Failures           []string
}

type FrontierSelection struct {
	Classification       string
	ShortStatus          string
	Verdict              string
	NextGate             string
	SelectedFrontier     string
	DeferredFrontier     string
	R4OrLaterFrontier    string
	NativeR3             bool
	AttackAFirst         bool
	AttackBBeforeA       bool
	EnterGenerationNow   bool
	LoopBackToPhase      bool
	LoopBackToRepAlpha   bool
	LoopBackToSocket     bool
	UpdateOfficialLedger bool
	Supports, Failures   []string
}

type Firewalls struct {
	NativeR3                     bool
	AlphaNative                  bool
	NativeZ2BoundaryAlphaFunctor bool
	NativeZ2AirlockFunctor       bool
	NativeDegreeToZ2FlagFunctor  bool
	NativeZ2CrossLane            bool
	ReducedB2NativeFunctional    bool
	FullAFDescent                bool
	AFOrientEqualsFullAF         bool
	HiggsOrientationNative       bool
	GenerationCarrierMap         bool
	FlavorOrientationMap         bool
	IndividualYukawaValues       bool
	OfficialLedgerUpdate         bool
	NativeYukawaOperator         bool
	R4NativeYukawa               bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedPlateau
	Trace          TraceDiagnostics
	FrontierA      FrontierAssessment
	FrontierB      FrontierAssessment
	FrontierC      FrontierAssessment
	Selection      FrontierSelection
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func LinearContribution(s float64) float64 { return float64(RankF1OverF0) / float64(LinearDenom) * s }
func QuadraticContribution(s float64) float64 {
	return float64(RankF2OverF0) / float64(QuadDenom) * s * s
}
func BoundaryAlphaZ2(s float64) float64 { return LinearContribution(s) + QuadraticContribution(s) }
func OperatorNEff(alpha float64) float64 {
	return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))
}
func OperatorCYukawa(neff float64) float64 { return 3.0 / neff }

func BuildDefault() (Audit, error) {
	inherited := buildInheritedPlateau()
	if !inherited.R3TraceLedgerSealed || inherited.NativeR3 || inherited.LoopBackPhase || inherited.LoopBackRepAlpha || inherited.LoopBackSocketOrder {
		return Audit{}, fmt.Errorf("inherited plateau leak: %s", FormatInherited(inherited))
	}

	trace := buildTraceDiagnostics()
	if !trace.DiagnosticsOnly || trace.OfficialUpdates || !near(trace.Alpha, AlphaB) || !near(trace.OperatorNEff, OperatorNEffDiagnostic) || !near(trace.OperatorCYukawa, OperatorCYukawaDiagnostic) || !near(trace.OperatorCHiggs, OperatorCHiggsDiagnostic) {
		return Audit{}, fmt.Errorf("trace diagnostic leak: %s", FormatTrace(trace))
	}

	frontierA := buildFrontierA()
	if !frontierA.SelectedFirst || !frontierA.RequiredForNativeR3 || !frontierA.DirectlyControlsAlpha || !frontierA.DirectlyControlsTraceWeights || !frontierA.DirectlyControlsNEff || frontierA.Deferred || frontierA.R4OrLater {
		return Audit{}, fmt.Errorf("frontier A selection leak: %s", FormatFrontier(frontierA))
	}

	frontierB := buildFrontierB()
	if frontierB.SelectedFirst || !frontierB.RequiredForNativeR3 || !frontierB.FullAFDescentProblem || !frontierB.Deferred || frontierB.DirectlyControlsAlpha || frontierB.R4OrLater {
		return Audit{}, fmt.Errorf("frontier B classification leak: %s", FormatFrontier(frontierB))
	}

	frontierC := buildFrontierC()
	if frontierC.SelectedFirst || frontierC.RequiredForNativeR3 || !frontierC.GenerationFlavorYukawaBranch || !frontierC.R4OrLater || !frontierC.Deferred {
		return Audit{}, fmt.Errorf("frontier C classification leak: %s", FormatFrontier(frontierC))
	}

	selection := buildSelection()
	if !selection.AttackAFirst || selection.AttackBBeforeA || selection.EnterGenerationNow || selection.NativeR3 || selection.LoopBackToPhase || selection.LoopBackToRepAlpha || selection.LoopBackToSocket || selection.UpdateOfficialLedger {
		return Audit{}, fmt.Errorf("frontier selection leak: %s", FormatSelection(selection))
	}

	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited:      inherited,
		Trace:          trace,
		FrontierA:      frontierA,
		FrontierB:      frontierB,
		FrontierC:      frontierC,
		Selection:      selection,
		Firewalls:      firewalls,
		Truth:          FinalTruth,
		Final:          "Gate 911 selects the next native-R3 rail without reopening phase sign, representative alpha, or socket-order wounds. Frontier A, the native Z2 BoundaryAlpha functor/source, is selected first because it directly controls alpha_B, the socket trace weights, Y^dagger Y readout, and the sealed N_eff/C_Yukawa/C_Higgs operator diagnostics. Frontier B, full A_F descent or lawful spontaneous-orientation status, remains required but is deferred until the alpha source is sharpened. Frontier C, generation/flavor/individual Yukawa spectrum, remains R4-or-later. Native R3, official ledger updates, and physical assignments remain blocked.",
	}, nil
}

func buildInheritedPlateau() InheritedPlateau {
	return InheritedPlateau{
		Gate910Classification: Gate910Classification,
		Gate910ShortStatus:    Gate910ShortStatus,
		Gate910Verdict:        Gate910Verdict,
		R3TraceLedgerSealed:   true,
		NativeR3:              false,
		LoopBackPhase:         false,
		LoopBackRepAlpha:      false,
		LoopBackSocketOrder:   false,
		Supports:              []string{StatusGate910Inherited, StatusNoLoopBack, SupportNoLoopBack},
		Failures:              []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeZ2BoundaryAlphaFunctor},
	}
}

func buildTraceDiagnostics() TraceDiagnostics {
	alpha := BoundaryAlphaZ2(SBoundary)
	neff := OperatorNEff(alpha)
	return TraceDiagnostics{
		BoundaryAlphaFormula: BoundaryAlphaFormula,
		Z2PunctureClass:      Z2PunctureClass,
		TraceRows:            TraceRowMultiset,
		NEffFormula:          NEffFormula,
		CYukawaFormula:       CYukawaFormula,
		S:                    SBoundary,
		Alpha:                alpha,
		OperatorNEff:         neff,
		OperatorCYukawa:      OperatorCYukawa(neff),
		OperatorCHiggs:       OperatorCHiggsDiagnostic,
		DiagnosticsOnly:      true,
		OfficialUpdates:      false,
		Supports:             []string{StatusTraceDiagnosticsSealed, SupportFrontierADirectTraceLedger, SupportFrontierAControlsAlphaWeights},
		Failures:             []string{FailureAlphaStillSealed, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
}

func buildFrontierA() FrontierAssessment {
	return FrontierAssessment{
		Name:                         FrontierAName,
		SelectedFirst:                true,
		RequiredForNativeR3:          true,
		DirectlyControlsAlpha:        true,
		DirectlyControlsTraceWeights: true,
		DirectlyControlsNEff:         true,
		DirectlyControlsCYukawa:      true,
		DirectlyControlsCHiggs:       true,
		FullAFDescentProblem:         false,
		GenerationFlavorYukawaBranch: false,
		R4OrLater:                    false,
		Deferred:                     false,
		MissingObjects: []string{
			MissingZ2BoundaryAlphaFunctor,
			"BoundaryAlpha_Z2 native source",
			"degree-to-Z2-flag-class functor",
			"native Z2 cross-lane exclusion theorem",
			"native S_split transport law into Z2 airlock class",
		},
		Supports: []string{
			StatusFrontierASelected,
			SupportNativeR3TwoBlockers,
			SupportFrontierAZ2AlphaPrimary,
			SupportFrontierADirectTraceLedger,
			SupportFrontierAControlsAlphaWeights,
			SupportFrontierAPrimaryBlocker,
			SupportNextBranchZ2Alpha,
		},
		Failures: []string{
			FailureFrontierANotCertified,
			FailureNoNativeZ2BoundaryAlphaFunctor,
			FailureNoNativeZ2AirlockFunctor,
			FailureNoNativeDegreeToZ2FlagFunctor,
			FailureNoNativeZ2CrossLane,
			FailureReducedB2NotNativeFunctional,
			FailureAlphaSealedUntilFunctor,
		},
	}
}

func buildFrontierB() FrontierAssessment {
	return FrontierAssessment{
		Name:                         FrontierBName,
		SelectedFirst:                false,
		RequiredForNativeR3:          true,
		DirectlyControlsAlpha:        false,
		DirectlyControlsTraceWeights: false,
		DirectlyControlsNEff:         false,
		DirectlyControlsCYukawa:      false,
		DirectlyControlsCHiggs:       false,
		FullAFDescentProblem:         true,
		GenerationFlavorYukawaBranch: false,
		R4OrLater:                    false,
		Deferred:                     true,
		MissingObjects: []string{
			"native descent from full A_F to A_F^orient",
			"lawful spontaneous-orientation interpretation",
			"proof that post-orientation projectors have native finite-sector status",
		},
		Supports: []string{
			StatusFrontierBDeferred,
			SupportNativeR3TwoBlockers,
			SupportFrontierBRequired,
			SupportFrontierBSecond,
			SupportFrontierBAfterAlpha,
		},
		Failures: []string{
			FailureFullAFDescentStillBlocked,
			FailureAFOrientNotFullAF,
			FailureHiggsOrientationSealed,
		},
	}
}

func buildFrontierC() FrontierAssessment {
	return FrontierAssessment{
		Name:                         FrontierCName,
		SelectedFirst:                false,
		RequiredForNativeR3:          false,
		DirectlyControlsAlpha:        false,
		DirectlyControlsTraceWeights: false,
		DirectlyControlsNEff:         false,
		DirectlyControlsCYukawa:      false,
		DirectlyControlsCHiggs:       false,
		FullAFDescentProblem:         false,
		GenerationFlavorYukawaBranch: true,
		R4OrLater:                    true,
		Deferred:                     true,
		MissingObjects: []string{
			"GenerationCarrierMap",
			"FlavorOrientationMap",
			"IndividualYukawaSpectrumMap",
			"physical particle assignment",
		},
		Supports: []string{StatusFrontierCDeferred, SupportFrontierCR4Later},
		Failures: []string{FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawa},
	}
}

func buildSelection() FrontierSelection {
	return FrontierSelection{
		Classification:       Classification,
		ShortStatus:          ShortStatus,
		Verdict:              FinalTruth,
		NextGate:             NextGate,
		SelectedFrontier:     FrontierAName,
		DeferredFrontier:     FrontierBName,
		R4OrLaterFrontier:    FrontierCName,
		NativeR3:             false,
		AttackAFirst:         true,
		AttackBBeforeA:       false,
		EnterGenerationNow:   false,
		LoopBackToPhase:      false,
		LoopBackToRepAlpha:   false,
		LoopBackToSocket:     false,
		UpdateOfficialLedger: false,
		Supports: []string{
			SupportNativeR3TwoBlockers,
			SupportFrontierAZ2AlphaPrimary,
			SupportFrontierBSecond,
			SupportFrontierCR4Later,
			SupportNextBranchZ2Alpha,
			SupportNoLoopBack,
		},
		Failures: []string{
			FailureNotNativeR3,
			FailureAlphaStillSealed,
			FailureNoNativeZ2BoundaryAlphaFunctor,
			FailureFullAFDescentStillBlocked,
			FailureNoGenerationCarrierMap,
		},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NativeR3:                     false,
		AlphaNative:                  false,
		NativeZ2BoundaryAlphaFunctor: false,
		NativeZ2AirlockFunctor:       false,
		NativeDegreeToZ2FlagFunctor:  false,
		NativeZ2CrossLane:            false,
		ReducedB2NativeFunctional:    false,
		FullAFDescent:                false,
		AFOrientEqualsFullAF:         false,
		HiggsOrientationNative:       false,
		GenerationCarrierMap:         false,
		FlavorOrientationMap:         false,
		IndividualYukawaValues:       false,
		OfficialLedgerUpdate:         false,
		NativeYukawaOperator:         false,
		R4NativeYukawa:               false,
	}
}

func Statuses() []string {
	return []string{
		StatusGate910Inherited,
		StatusNoLoopBack,
		StatusFrontierASelected,
		StatusFrontierBDeferred,
		StatusFrontierCDeferred,
		StatusTraceDiagnosticsSealed,
		StatusFirewallVerdict,
		NextGate,
		SupportNativeR3TwoBlockers,
		SupportFrontierAZ2AlphaPrimary,
		SupportFrontierADirectTraceLedger,
		SupportFrontierAControlsAlphaWeights,
		SupportFrontierAPrimaryBlocker,
		SupportFrontierBRequired,
		SupportFrontierBSecond,
		SupportFrontierBAfterAlpha,
		SupportFrontierCR4Later,
		SupportNextBranchZ2Alpha,
		SupportNoLoopBack,
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureFrontierANotCertified,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureReducedB2NotNativeFunctional,
		FailureAlphaSealedUntilFunctor,
		FailureFullAFDescentStillBlocked,
		FailureAFOrientNotFullAF,
		FailureHiggsOrientationSealed,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoOfficialNEffUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNotNativeR3,
		FailureAlphaStillSealed,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureReducedB2NotNativeFunctional,
		FailureFullAFDescentStillBlocked,
		FailureAFOrientNotFullAF,
		FailureHiggsOrientationSealed,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func FormatInherited(a InheritedPlateau) string {
	return fmt.Sprintf("gate910_classification=%s gate910_short=%s gate910_verdict=%s sealed_R3=%t native_R3=%t loop_phase=%t loop_rep_alpha=%t loop_socket=%t supports=%s failures=%s", a.Gate910Classification, a.Gate910ShortStatus, a.Gate910Verdict, a.R3TraceLedgerSealed, a.NativeR3, a.LoopBackPhase, a.LoopBackRepAlpha, a.LoopBackSocketOrder, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatTrace(a TraceDiagnostics) string {
	return fmt.Sprintf("class=%s formula=%s rows=%s N_eff_formula=%s CY_formula=%s s=%.17g alpha=%.17g operator_N_eff=%.16g operator_CY=%.16g operator_CH=%.16g diagnostics_only=%t official_updates=%t supports=%s failures=%s", a.Z2PunctureClass, a.BoundaryAlphaFormula, a.TraceRows, a.NEffFormula, a.CYukawaFormula, a.S, a.Alpha, a.OperatorNEff, a.OperatorCYukawa, a.OperatorCHiggs, a.DiagnosticsOnly, a.OfficialUpdates, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFrontier(a FrontierAssessment) string {
	return fmt.Sprintf("name=%s selected_first=%t required_native_R3=%t controls_alpha=%t controls_trace_weights=%t controls_N_eff=%t controls_CY=%t controls_CH=%t full_AF_problem=%t generation_flavor_yukawa=%t R4_or_later=%t deferred=%t missing=%s supports=%s failures=%s", a.Name, a.SelectedFirst, a.RequiredForNativeR3, a.DirectlyControlsAlpha, a.DirectlyControlsTraceWeights, a.DirectlyControlsNEff, a.DirectlyControlsCYukawa, a.DirectlyControlsCHiggs, a.FullAFDescentProblem, a.GenerationFlavorYukawaBranch, a.R4OrLater, a.Deferred, strings.Join(a.MissingObjects, ";"), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatSelection(a FrontierSelection) string {
	return fmt.Sprintf("classification=%s short=%s verdict=%s next=%s selected=%s deferred=%s R4=%s native_R3=%t attack_A_first=%t attack_B_before_A=%t enter_generation_now=%t loop_phase=%t loop_rep_alpha=%t loop_socket=%t update_official=%t supports=%s failures=%s", a.Classification, a.ShortStatus, a.Verdict, a.NextGate, a.SelectedFrontier, a.DeferredFrontier, a.R4OrLaterFrontier, a.NativeR3, a.AttackAFirst, a.AttackBBeforeA, a.EnterGenerationNow, a.LoopBackToPhase, a.LoopBackToRepAlpha, a.LoopBackToSocket, a.UpdateOfficialLedger, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_R3=%t alpha_native=%t native_Z2_boundary_alpha=%t native_Z2_airlock=%t native_degree_Z2_flag=%t native_Z2_cross_lane=%t reduced_B2_native=%t full_AF_descent=%t AF_orient_equals_full_AF=%t higgs_native=%t generation=%t flavor=%t individual_yukawa=%t official_update=%t native_yukawa_operator=%t R4_native_yukawa=%t", f.NativeR3, f.AlphaNative, f.NativeZ2BoundaryAlphaFunctor, f.NativeZ2AirlockFunctor, f.NativeDegreeToZ2FlagFunctor, f.NativeZ2CrossLane, f.ReducedB2NativeFunctional, f.FullAFDescent, f.AFOrientEqualsFullAF, f.HiggsOrientationNative, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.OfficialLedgerUpdate, f.NativeYukawaOperator, f.R4NativeYukawa)
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeR3 && !f.AlphaNative && !f.NativeZ2BoundaryAlphaFunctor && !f.NativeZ2AirlockFunctor && !f.NativeDegreeToZ2FlagFunctor && !f.NativeZ2CrossLane && !f.ReducedB2NativeFunctional && !f.FullAFDescent && !f.AFOrientEqualsFullAF && !f.HiggsOrientationNative && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.OfficialLedgerUpdate && !f.NativeYukawaOperator && !f.R4NativeYukawa
}

func containsAll(haystack []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
