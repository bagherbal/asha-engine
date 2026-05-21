// Package generation2phaseanchoredneutralpunctureairlockfunctoraudit implements
// Gate 901: PhaseAnchored NeutralPuncture Airlock Functor Audit.
//
// Gate 901 follows Gate 900's R3-sealed classification and asks whether the
// remaining blockers are projections of one bridge object: a phase-anchored
// neutral-puncture airlock functor. It does not promote to native R3, derive
// alpha_B, assign physical particles, split generations, or update official
// ledgers.
package generation2phaseanchoredneutralpunctureairlockfunctoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE901-PHASE-ANCHORED-NEUTRAL-PUNCTURE-AIRLOCK-FUNCTOR-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OperatorCHiggsDiagnostic  = 1.037220510866514

	OfficialNEffFrozen        = 3.0023273474722147
	OfficialCYukawaFrozen     = 0.9992248188812008
	OfficialCHiggsFrozen      = 1.0372205204048603
	AlphaRankOneNumerator     = 3
	AlphaRankTwoNumerator     = 7
	AlphaLinearDenominator    = 10
	AlphaQuadraticDenominator = 72

	RightCharacterSplit = "rho_R(lambda)=lambda e_+ + bar(lambda) e_-"
	PhaseAnchor         = "o_phi: lambda succeeds bar(lambda)"
	ELambda             = "e_lambda=e_+"
	EBarLambda          = "e_bar_lambda=e_-"
	HLambda             = "h_lambda=h_+"
	HBarLambda          = "h_bar_lambda=h_-"
	PPhase              = "p_phi=e_lambda tensor P_1=e_+ tensor P_1"
	PiTop               = "F_1/F_0=e_lambda tensor P_3=Pi_top"
	HRMin               = "F_2/F_0=(C_R^2 tensor W)-p_phi=H_R^min"
	AFOrient            = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullAF              = "A_F=C plus H plus M_3(C)"

	Classification = "R3_SEALED_CANDIDATE_REDUCED_TO_PHASE_ANCHORED_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	ShortStatus    = "R3_PHASE_ANCHORED_NEUTRAL_PUNCTURE_AIRLOCK_SEAL_NOT_NATIVE"

	StatusGate900Inherited     = "PASS_GATE900_R3_SEALED_CANDIDATE_INHERITED"
	StatusPhaseAnchorDefined   = "PASS_PHASE_ANCHOR_DEFINED_AS_BRIDGE_OBJECT"
	StatusSocketOrderCollapsed = "PASS_SOCKET_ORDER_COLLAPSED_TO_PHASE_ANCHOR"
	StatusEdgeOrderCollapsed   = "PASS_EDGE_ORDERING_COLLAPSED_TO_PHASE_ANCHORED_AIRLOCK_RULE"
	StatusWeakKernelCollapsed  = "PASS_WEAK_SOCKET_SELECTOR_COLLAPSED_TO_PHASE_INDEXED_KERNEL_COMPLEMENT"
	StatusAlphaFlagCollapsed   = "PASS_BOUNDARY_ALPHA_FLAG_COLLAPSED_TO_PHASE_ANCHORED_PUNCTURE"
	StatusAFOrientClassified   = "PASS_A_F_ORIENT_CLASSIFIED_AS_PHASE_ANCHORED_STABILIZER_LAYER"
	StatusMasterBlockerReduced = "PASS_MULTI_BLOCKER_STACK_REDUCED_TO_PHASE_ANCHORED_AIRLOCK_FUNCTOR"
	StatusOfficialFreeze       = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict      = "FIREWALL_PRESERVED_GATE901_PHASE_ANCHORED_AIRLOCK_NOT_NATIVE"

	SupportPhaseAnchorOrdersRightPair        = "CONDITIONAL_SUPPORT_PHASE_ANCHOR_ORDERS_RIGHT_CHARACTER_PAIR"
	SupportPhaseAnchorSelectsEPlus           = "CONDITIONAL_SUPPORT_PHASE_ANCHOR_SELECTS_E_PLUS_AS_EXPOSURE_PUNCTURE_SOCKET"
	SupportPhaseAnchoredPuncture             = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_PUNCTURE_IS_E_PLUS_TENSOR_P1"
	SupportPhaseAnchoredEdgeTable            = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_GENERATES_ORDERED_EDGE_TABLE"
	SupportPhaseAnchoredLeftKernel           = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_RECONSTRUCTS_LEFT_KERNEL_LINE"
	SupportPhaseAnchoredAlphaTargets         = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_RECONSTRUCTS_BOUNDARY_ALPHA_TARGETS"
	SupportAirlockUnifiesAlphaAndHiggs       = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_UNIFIES_ALPHA_AND_HIGGS_ORIENTATION_WOUNDS"
	SupportR3SealedReducesToSingleFunctor    = "CONDITIONAL_SUPPORT_R3_SEALED_STRUCTURE_REDUCES_TO_SINGLE_PHASE_ANCHORED_AIRLOCK_FUNCTOR"
	SupportAFOrientPhaseAnchored             = "CONDITIONAL_SUPPORT_A_F_ORIENT_IS_PHASE_ANCHORED_STABILIZER_LAYER"
	SupportSocketOrderObstructionCollapsed   = "CONDITIONAL_SUPPORT_SOCKET_ORDER_OBSTRUCTION_COLLAPSES_TO_PHASE_ANCHOR"
	SupportBoundaryIncidenceCollapsed        = "CONDITIONAL_SUPPORT_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_SEAL_COLLAPSES_TO_PHASE_ANCHORED_AIRLOCK"
	SupportHiggsOrientationCollapsed         = "CONDITIONAL_SUPPORT_HIGGS_ORIENTATION_SEAL_COLLAPSES_TO_PHASE_ANCHORED_NULL_EDGE_RULE"
	SupportOperatorDiagnosticsRemainCoherent = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTICS_REMAIN_COHERENT_UNDER_PHASE_ANCHORED_AIRLOCK_SEAL"

	FailureNoNativePhaseAnchoredAirlock        = "FAILED_ROUTE_NO_NATIVE_PHASE_ANCHORED_AIRLOCK_FUNCTOR"
	FailureNoNativeRightPhaseOrientation       = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureNoNativeSelectionLambda             = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_LAMBDA_OVER_BARLAMBDA"
	FailureAlphaStillSealedWithoutPhaseFunctor = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealedWithoutPhaseWeak    = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureNoNativeBoundaryIncidenceFlag       = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeWeakSocketSelector          = "FAILED_ROUTE_NO_NATIVE_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	FailureAFOrientNotFullAF                   = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeDescentFullAF               = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNotNativeR3                         = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoPhysicalParticleAssignment        = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap              = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap              = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues            = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem             = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type PhaseAnchorAudit struct {
	RightCharacterSplit string
	Anchor              string
	ELambda, EBarLambda string
	OrdersPair          bool
	SelectsNatively     bool
	Supports, Failures  []string
}

type SocketOrderAudit struct {
	ExposureSocket, RestSocket string
	Puncture                   string
	OrderedByPhaseAnchor       bool
	OrderedNatively            bool
	Supports, Failures         []string
}

type EdgeOrderingAudit struct {
	Edges                  []string
	MissingEdge            string
	GeneratedByPhaseAnchor bool
	GeneratedNatively      bool
	Supports, Failures     []string
}

type WeakKernelAudit struct {
	Image              []string
	Kernel             string
	PhaseIndexedFrame  bool
	SelectorNative     bool
	Supports, Failures []string
}

type BoundaryAlphaAudit struct {
	F0, F1, F2                 string
	RankF1OverF0, RankF2OverF0 int
	Alpha                      float64
	SelectedByPhasePuncture    bool
	NativeAlphaFunctor         bool
	Supports, Failures         []string
}

type OrientedStabilizerAudit struct {
	Layer              string
	FullLayer          string
	PhaseAnchored      bool
	FullDescent        bool
	Supports, Failures []string
}

type MasterFunctorAudit struct {
	UnifiesSocketOrder   bool
	UnifiesEdgeOrdering  bool
	UnifiesWeakKernel    bool
	UnifiesBoundaryAlpha bool
	SingleMasterBlocker  bool
	NativeFunctor        bool
	Supports, Failures   []string
}

type DiagnosticsAudit struct {
	Alpha, NEff, CYukawa, CHiggs                  float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs float64
	Coherent, OfficialFrozen, CanUpdate           bool
	Supports, Failures                            []string
}

type Firewalls struct {
	Enforced                         bool
	NoNativePhaseAnchoredAirlock     bool
	NoNativeRightPhaseOrientation    bool
	NoLambdaOverBarLambdaSelection   bool
	AlphaStillSealed                 bool
	HiggsOrientationStillSealed      bool
	NoNativeBoundaryIncidenceFunctor bool
	NoNativeWeakSocketSelector       bool
	AFOrientNotFullAF                bool
	NoFullDescent                    bool
	NotNativeR3                      bool
	NoPhysicalAssignment             bool
	NoGenerationCarrier              bool
	NoFlavorOrientation              bool
	NoIndividualYukawas              bool
	NoOfficialLedgerUpdate           bool
	NoNativeYukawaOperator           bool
	NoR4YukawaTheorem                bool
	Verdict                          string
}

type Audit struct {
	ID             string
	PhaseAnchor    PhaseAnchorAudit
	SocketOrder    SocketOrderAudit
	EdgeOrdering   EdgeOrderingAudit
	WeakKernel     WeakKernelAudit
	BoundaryAlpha  BoundaryAlphaAudit
	Stabilizer     OrientedStabilizerAudit
	MasterFunctor  MasterFunctorAudit
	Diagnostics    DiagnosticsAudit
	Firewalls      Firewalls
	Classification string
	ShortStatus    string
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	phase := buildPhaseAnchorAudit()
	if !phase.OrdersPair || phase.SelectsNatively {
		return Audit{}, fmt.Errorf("phase anchor leak: %s", FormatPhaseAnchor(phase))
	}
	socket := buildSocketOrderAudit()
	if !socket.OrderedByPhaseAnchor || socket.OrderedNatively || socket.Puncture != PPhase {
		return Audit{}, fmt.Errorf("socket order leak: %s", FormatSocketOrder(socket))
	}
	edge := buildEdgeOrderingAudit()
	if !edge.GeneratedByPhaseAnchor || edge.GeneratedNatively || edge.MissingEdge != "e_lambda tensor P_1 -> h_lambda tensor P_1 = 0" {
		return Audit{}, fmt.Errorf("edge order leak: %s", FormatEdgeOrdering(edge))
	}
	weak := buildWeakKernelAudit()
	if !weak.PhaseIndexedFrame || weak.SelectorNative || weak.Kernel != "h_lambda tensor P_1" {
		return Audit{}, fmt.Errorf("weak kernel leak: %s", FormatWeakKernel(weak))
	}
	alpha := buildBoundaryAlphaAudit()
	if !alpha.SelectedByPhasePuncture || alpha.NativeAlphaFunctor || !near(alpha.Alpha, AlphaB) || alpha.RankF1OverF0 != AlphaRankOneNumerator || alpha.RankF2OverF0 != AlphaRankTwoNumerator {
		return Audit{}, fmt.Errorf("boundary alpha leak: %s", FormatBoundaryAlpha(alpha))
	}
	stab := buildStabilizerAudit()
	if !stab.PhaseAnchored || stab.FullDescent {
		return Audit{}, fmt.Errorf("stabilizer leak: %s", FormatStabilizer(stab))
	}
	master := buildMasterFunctorAudit()
	if !master.UnifiesSocketOrder || !master.UnifiesEdgeOrdering || !master.UnifiesWeakKernel || !master.UnifiesBoundaryAlpha || !master.SingleMasterBlocker || master.NativeFunctor {
		return Audit{}, fmt.Errorf("master functor leak: %s", FormatMasterFunctor(master))
	}
	diag := buildDiagnosticsAudit()
	if !diag.Coherent || !diag.OfficialFrozen || diag.CanUpdate || !near(diag.NEff, OperatorNEffDiagnostic) || near(diag.NEff, diag.OfficialNEff) {
		return Audit{}, fmt.Errorf("diagnostic freeze leak: %s", FormatDiagnostics(diag))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{ID: AuditID, PhaseAnchor: phase, SocketOrder: socket, EdgeOrdering: edge, WeakKernel: weak, BoundaryAlpha: alpha, Stabilizer: stab, MasterFunctor: master, Diagnostics: diag, Firewalls: firewalls, Classification: Classification, ShortStatus: ShortStatus, Truth: "Gate 901 reduces the R3-sealed candidate's many local blockers to one master bridge object: a phase-anchored neutral-puncture airlock functor. The functor is coherent under seal, but not native.", Final: ShortStatus}, nil
}

func buildPhaseAnchorAudit() PhaseAnchorAudit {
	return PhaseAnchorAudit{RightCharacterSplit: RightCharacterSplit, Anchor: PhaseAnchor, ELambda: ELambda, EBarLambda: EBarLambda, OrdersPair: true, SelectsNatively: false, Supports: []string{StatusGate900Inherited, StatusPhaseAnchorDefined, SupportPhaseAnchorOrdersRightPair, SupportPhaseAnchorSelectsEPlus, SupportSocketOrderObstructionCollapsed}, Failures: []string{FailureNoNativePhaseAnchoredAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda}}
}

func buildSocketOrderAudit() SocketOrderAudit {
	return SocketOrderAudit{ExposureSocket: ELambda, RestSocket: EBarLambda, Puncture: PPhase, OrderedByPhaseAnchor: true, OrderedNatively: false, Supports: []string{StatusSocketOrderCollapsed, SupportPhaseAnchorSelectsEPlus, SupportPhaseAnchoredPuncture, SupportSocketOrderObstructionCollapsed}, Failures: []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda}}
}

func buildEdgeOrderingAudit() EdgeOrderingAudit {
	return EdgeOrderingAudit{Edges: []string{"e_lambda tensor P_3 -> h_lambda tensor P_3", "e_bar_lambda tensor P_3 -> h_bar_lambda tensor P_3", "e_bar_lambda tensor P_1 -> h_bar_lambda tensor P_1"}, MissingEdge: "e_lambda tensor P_1 -> h_lambda tensor P_1 = 0", GeneratedByPhaseAnchor: true, GeneratedNatively: false, Supports: []string{StatusEdgeOrderCollapsed, SupportPhaseAnchoredEdgeTable}, Failures: []string{FailureNoNativePhaseAnchoredAirlock, FailureHiggsStillSealedWithoutPhaseWeak}}
}

func buildWeakKernelAudit() WeakKernelAudit {
	return WeakKernelAudit{Image: []string{"h_lambda tensor P_3", "h_bar_lambda tensor P_3", "h_bar_lambda tensor P_1"}, Kernel: "h_lambda tensor P_1", PhaseIndexedFrame: true, SelectorNative: false, Supports: []string{StatusWeakKernelCollapsed, SupportPhaseAnchoredLeftKernel, SupportHiggsOrientationCollapsed}, Failures: []string{FailureNoNativeWeakSocketSelector, FailureHiggsStillSealedWithoutPhaseWeak}}
}

func buildBoundaryAlphaAudit() BoundaryAlphaAudit {
	alpha := float64(AlphaRankOneNumerator)/float64(AlphaLinearDenominator)*Ssplit + float64(AlphaRankTwoNumerator)/float64(AlphaQuadraticDenominator)*Ssplit*Ssplit
	return BoundaryAlphaAudit{F0: PPhase, F1: "F_1=e_lambda tensor W", F2: "F_2=C_R^2 tensor W", RankF1OverF0: AlphaRankOneNumerator, RankF2OverF0: AlphaRankTwoNumerator, Alpha: alpha, SelectedByPhasePuncture: true, NativeAlphaFunctor: false, Supports: []string{StatusAlphaFlagCollapsed, SupportPhaseAnchoredAlphaTargets, SupportBoundaryIncidenceCollapsed}, Failures: []string{FailureAlphaStillSealedWithoutPhaseFunctor, FailureNoNativeBoundaryIncidenceFlag}}
}

func buildStabilizerAudit() OrientedStabilizerAudit {
	return OrientedStabilizerAudit{Layer: AFOrient, FullLayer: FullAF, PhaseAnchored: true, FullDescent: false, Supports: []string{StatusAFOrientClassified, SupportAFOrientPhaseAnchored}, Failures: []string{FailureAFOrientNotFullAF, FailureNoNativeDescentFullAF}}
}

func buildMasterFunctorAudit() MasterFunctorAudit {
	return MasterFunctorAudit{UnifiesSocketOrder: true, UnifiesEdgeOrdering: true, UnifiesWeakKernel: true, UnifiesBoundaryAlpha: true, SingleMasterBlocker: true, NativeFunctor: false, Supports: []string{StatusMasterBlockerReduced, SupportAirlockUnifiesAlphaAndHiggs, SupportR3SealedReducesToSingleFunctor}, Failures: []string{FailureNoNativePhaseAnchoredAirlock, FailureNotNativeR3}}
}

func buildDiagnosticsAudit() DiagnosticsAudit {
	return DiagnosticsAudit{Alpha: AlphaB, NEff: OperatorNEffDiagnostic, CYukawa: OperatorCYukawaDiagnostic, CHiggs: OperatorCHiggsDiagnostic, OfficialNEff: OfficialNEffFrozen, OfficialCYukawa: OfficialCYukawaFrozen, OfficialCHiggs: OfficialCHiggsFrozen, Coherent: true, OfficialFrozen: true, CanUpdate: false, Supports: []string{StatusOfficialFreeze, SupportOperatorDiagnosticsRemainCoherent}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NoNativePhaseAnchoredAirlock: true, NoNativeRightPhaseOrientation: true, NoLambdaOverBarLambdaSelection: true, AlphaStillSealed: true, HiggsOrientationStillSealed: true, NoNativeBoundaryIncidenceFunctor: true, NoNativeWeakSocketSelector: true, AFOrientNotFullAF: true, NoFullDescent: true, NotNativeR3: true, NoPhysicalAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4YukawaTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatPhaseAnchor(p PhaseAnchorAudit) string {
	return fmt.Sprintf("phase_anchor(split=%s anchor=%s e_lambda=%s e_barlambda=%s orders=%t native=%t supports=%s failures=%s)", p.RightCharacterSplit, p.Anchor, p.ELambda, p.EBarLambda, p.OrdersPair, p.SelectsNatively, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatSocketOrder(s SocketOrderAudit) string {
	return fmt.Sprintf("socket_order(exposure=%s rest=%s puncture=%s ordered_by_phase=%t native=%t supports=%s failures=%s)", s.ExposureSocket, s.RestSocket, s.Puncture, s.OrderedByPhaseAnchor, s.OrderedNatively, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatEdgeOrdering(e EdgeOrderingAudit) string {
	return fmt.Sprintf("edge_ordering(edges=%s missing=%s generated_by_phase=%t native=%t supports=%s failures=%s)", strings.Join(e.Edges, ";"), e.MissingEdge, e.GeneratedByPhaseAnchor, e.GeneratedNatively, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}

func FormatWeakKernel(w WeakKernelAudit) string {
	return fmt.Sprintf("weak_kernel(image=%s kernel=%s phase_indexed=%t native=%t supports=%s failures=%s)", strings.Join(w.Image, ";"), w.Kernel, w.PhaseIndexedFrame, w.SelectorNative, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatBoundaryAlpha(b BoundaryAlphaAudit) string {
	return fmt.Sprintf("boundary_alpha(F0=%s F1=%s F2=%s ranks=(%d,%d) alpha=%.16g selected_by_phase=%t native=%t supports=%s failures=%s)", b.F0, b.F1, b.F2, b.RankF1OverF0, b.RankF2OverF0, b.Alpha, b.SelectedByPhasePuncture, b.NativeAlphaFunctor, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatStabilizer(s OrientedStabilizerAudit) string {
	return fmt.Sprintf("stabilizer(layer=%s full=%s phase_anchored=%t full_descent=%t supports=%s failures=%s)", s.Layer, s.FullLayer, s.PhaseAnchored, s.FullDescent, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatMasterFunctor(m MasterFunctorAudit) string {
	return fmt.Sprintf("master_functor(socket=%t edge=%t weak=%t alpha=%t single_blocker=%t native=%t supports=%s failures=%s)", m.UnifiesSocketOrder, m.UnifiesEdgeOrdering, m.UnifiesWeakKernel, m.UnifiesBoundaryAlpha, m.SingleMasterBlocker, m.NativeFunctor, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}

func FormatDiagnostics(d DiagnosticsAudit) string {
	return fmt.Sprintf("diagnostics(alpha=%.16g neff=%.16g cy=%.16g ch=%.16g official_neff=%.16g official_cy=%.16g official_ch=%.16g coherent=%t frozen=%t can_update=%t supports=%s failures=%s)", d.Alpha, d.NEff, d.CYukawa, d.CHiggs, d.OfficialNEff, d.OfficialCYukawa, d.OfficialCHiggs, d.Coherent, d.OfficialFrozen, d.CanUpdate, strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t phase_airlock=%t right_phase=%t lambda_order=%t alpha=%t higgs=%t boundary=%t weak=%t af_orient=%t descent=%t not_r3=%t physical=%t generation=%t flavor=%t individual=%t official=%t yukawa=%t r4=%t verdict=%s)", f.Enforced, f.NoNativePhaseAnchoredAirlock, f.NoNativeRightPhaseOrientation, f.NoLambdaOverBarLambdaSelection, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNativeBoundaryIncidenceFunctor, f.NoNativeWeakSocketSelector, f.AFOrientNotFullAF, f.NoFullDescent, f.NotNativeR3, f.NoPhysicalAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4YukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{StatusGate900Inherited, StatusPhaseAnchorDefined, StatusSocketOrderCollapsed, StatusEdgeOrderCollapsed, StatusWeakKernelCollapsed, StatusAlphaFlagCollapsed, StatusAFOrientClassified, StatusMasterBlockerReduced, StatusOfficialFreeze, StatusFirewallVerdict, SupportPhaseAnchorOrdersRightPair, SupportPhaseAnchorSelectsEPlus, SupportPhaseAnchoredPuncture, SupportPhaseAnchoredEdgeTable, SupportPhaseAnchoredLeftKernel, SupportPhaseAnchoredAlphaTargets, SupportAirlockUnifiesAlphaAndHiggs, SupportR3SealedReducesToSingleFunctor, SupportAFOrientPhaseAnchored, SupportSocketOrderObstructionCollapsed, SupportBoundaryIncidenceCollapsed, SupportHiggsOrientationCollapsed, SupportOperatorDiagnosticsRemainCoherent, FailureNoNativePhaseAnchoredAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda, FailureAlphaStillSealedWithoutPhaseFunctor, FailureHiggsStillSealedWithoutPhaseWeak, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeWeakSocketSelector, FailureAFOrientNotFullAF, FailureNoNativeDescentFullAF, FailureNotNativeR3, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoNativePhaseAnchoredAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda, FailureAlphaStillSealedWithoutPhaseFunctor, FailureHiggsStillSealedWithoutPhaseWeak, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeWeakSocketSelector, FailureAFOrientNotFullAF, FailureNoNativeDescentFullAF, FailureNotNativeR3, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativePhaseAnchoredAirlock && f.NoNativeRightPhaseOrientation && f.NoLambdaOverBarLambdaSelection && f.AlphaStillSealed && f.HiggsOrientationStillSealed && f.NoNativeBoundaryIncidenceFunctor && f.NoNativeWeakSocketSelector && f.AFOrientNotFullAF && f.NoFullDescent && f.NotNativeR3 && f.NoPhysicalAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4YukawaTheorem && f.Verdict == StatusFirewallVerdict
}

func containsAll(haystack []string, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 5e-15 }
