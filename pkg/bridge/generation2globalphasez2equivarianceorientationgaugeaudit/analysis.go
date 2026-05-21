// Package generation2globalphasez2equivarianceorientationgaugeaudit implements
// Gate 907: GlobalPhaseZ2 Equivariance and OrientationGauge Audit.
//
// Gate 907 follows Gate 906's reduction of the remaining right-character wound
// to the sign of Q_phi=e_lambda-e_barlambda. It audits whether this sign must be
// natively selected, or whether Q_phi -> -Q_phi is only a global phase-orientation
// gauge. The honest verdict is that all current trace-level data are invariant
// under the lambda/bar(lambda) Z2 reversal at the level of an orientation-class
// candidate, but no native global equivariance theorem, phase-module-to-ASHA-C_R2
// identification, neutral puncture airlock theorem, alpha theorem, Higgs theorem,
// native R3, particle assignment, individual Yukawa values, or official ledger
// update is certified.
package generation2globalphasez2equivarianceorientationgaugeaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE907-GLOBAL-PHASE-Z2-EQUIVARIANCE-ORIENTATION-GAUGE-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	Z2Operation    = "tau_phi: lambda<->bar(lambda), e_lambda<->e_barlambda, Q_phi->-Q_phi"
	CurrentReadout = "1*Pi_lambda3 + alpha_B(1-alpha_B)*Pi_barlambda3 + 3 alpha_B^2*Pi_barlambda1"
	MirrorReadout  = "1*Pi_barlambda3 + alpha_B(1-alpha_B)*Pi_lambda3 + 3 alpha_B^2*Pi_lambda1"
	WeightMultiset = "{(rank 3, weight 1), (rank 3, weight alpha_B(1-alpha_B)), (rank 1, weight 3 alpha_B^2)}"
	Classification = "R3_AIRLOCK_PHASE_SIGN_RECLASSIFIED_AS_ORIENTATION_GAUGE_CANDIDATE"
	ShortStatus    = "R3_PHASE_ORIENTATION_GAUGE_EQUIVARIANCE_CANDIDATE_NOT_NATIVE"
	FinalTruth     = "GLOBAL_PHASE_Z2_EQUIVARIANCE_CAN_REMOVE_ABSOLUTE_SIGN_SELECTION_PRESSURE_IF_CERTIFIED"

	StatusGate906Inherited    = "PASS_GATE906_Q_PHI_SIGN_SELECTION_WOUND_INHERITED"
	StatusZ2Defined           = "PASS_GLOBAL_PHASE_Z2_OPERATION_DEFINED"
	StatusAirlockEquivariance = "PASS_AIRLOCK_PAIR_EQUIVARIANCE_AUDITED"
	StatusEdgeEquivariance    = "PASS_EDGE_TABLE_Z2_MIRROR_AUDITED"
	StatusTraceEquivariance   = "PASS_TRACE_MAGNITUDE_Z2_INVARIANCE_AUDITED"
	StatusLabelFirewall       = "PASS_SOCKET_LABEL_ORIENTATION_DEPENDENCE_AUDITED"
	StatusWoundReclassified   = "PASS_Q_PHI_SIGN_WOUND_WEAKENED_TO_ORIENTATION_GAUGE_CANDIDATE"
	StatusOfficialFreeze      = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict     = "FIREWALL_PRESERVED_GATE907_EQUIVARIANCE_NOT_NATIVE_R3"

	SupportGate906Inherited      = "CONDITIONAL_SUPPORT_GATE906_POSITIVE_GENERATOR_WOUND_INHERITED"
	SupportGlobalZ2Exchanges     = "CONDITIONAL_SUPPORT_GLOBAL_PHASE_Z2_EXCHANGES_AIRLOCK_REPRESENTATIVES"
	SupportAlphaInvariant        = "CONDITIONAL_SUPPORT_ALPHA_RANK_RECONSTRUCTION_IS_Z2_INVARIANT"
	SupportAirlockClass          = "CONDITIONAL_SUPPORT_PUNCTURE_AIRLOCK_EXISTS_AS_ORIENTATION_CLASS"
	SupportEdgeMirror            = "CONDITIONAL_SUPPORT_EDGE_TABLE_HAS_Z2_MIRROR_REPRESENTATIVE"
	SupportEdgeRankKernel        = "CONDITIONAL_SUPPORT_EDGE_RANK_AND_KERNEL_COUNT_ARE_Z2_INVARIANT"
	SupportEdgeGaugeRep          = "CONDITIONAL_SUPPORT_ORIENTED_EDGE_TABLE_IS_A_GAUGE_REPRESENTATIVE_IF_Z2_EQUIVARIANCE_CERTIFIED"
	SupportRowMultiset           = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROW_MULTISET_IS_Z2_INVARIANT"
	SupportNEffInvariant         = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_IS_PHASE_ORIENTATION_INVARIANT"
	SupportCYukawaInvariant      = "CONDITIONAL_SUPPORT_C_YUKAWA_OPERATOR_IS_PHASE_ORIENTATION_INVARIANT"
	SupportLabelsGauge           = "CONDITIONAL_SUPPORT_SOCKET_LABELS_ARE_ORIENTATION_DEPENDENT"
	SupportNoPhysicalLabelsNeed  = "CONDITIONAL_SUPPORT_TRACE_LEDGER_DOES_NOT_REQUIRE_PHYSICAL_LABEL_ASSIGNMENT"
	SupportQPhiGauge             = "CONDITIONAL_SUPPORT_ABSOLUTE_Q_PHI_SIGN_MAY_BE_ORIENTATION_GAUGE_NOT_NATIVE_OBSERVABLE"
	SupportNativePressureZ2      = "CONDITIONAL_SUPPORT_NATIVE_R3_PRESSURE_REDUCES_TO_Z2_EQUIVARIANT_AIRLOCK_THEOREM"
	SupportGeneratorWoundWeakens = "CONDITIONAL_SUPPORT_POSITIVE_PHASE_GENERATOR_WOUND_WEAKENS_TO_ORIENTATION_GAUGE"

	FailureNoNativeGlobalZ2          = "FAILED_ROUTE_NO_NATIVE_GLOBAL_PHASE_Z2_EQUIVARIANCE_THEOREM"
	FailureNoNativePhaseModuleCR2    = "FAILED_ROUTE_NO_NATIVE_PHASE_MODULE_TO_ASHA_C_R2_IDENTIFICATION"
	FailureNoNativeAirlock           = "FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	FailureAlphaStillSealed          = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealed          = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureFullAFDescentBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNotNativeR3               = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoPhysicalAssignment      = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap    = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap    = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawa        = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate      = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate     = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator    = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawa          = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
	FailureEdgeEquivarianceNotNative = "FAILED_ROUTE_EDGE_EQUIVARIANCE_NOT_YET_NATIVE_OPERATOR_THEOREM"
)

type InheritedAudit struct {
	Gate906Classification string
	QPhiExists            bool
	PositiveSignSelected  bool
	Supports, Failures    []string
}

type Z2OperationAudit struct {
	Operation            string
	SwapsLambdaBarLambda bool
	SwapsProjectors      bool
	FlipsQPhi            bool
	NativeTheorem        bool
	Supports, Failures   []string
}

type AirlockEquivarianceAudit struct {
	PlusPuncture, MinusPuncture   string
	PlusFlagRanks, MinusFlagRanks []int
	AlphaRanksInvariant           bool
	OrientationClassCandidate     bool
	NativeEquivariance            bool
	Supports, Failures            []string
}

type EdgeEquivarianceAudit struct {
	CurrentRankPattern, MirrorRankPattern []int
	CurrentKernelCount, MirrorKernelCount int
	MirrorRepresentativeExists            bool
	RankKernelInvariant                   bool
	NativeOperatorTheorem                 bool
	Supports, Failures                    []string
}

type TraceEquivarianceAudit struct {
	Current, Mirror, Multiset            string
	TraceInvariant, SquareTraceInvariant bool
	NEffInvariant, CYukawaInvariant      bool
	NativeTraceTheorem                   bool
	Supports, Failures                   []string
}

type LabelAudit struct {
	SocketLabelsChange       bool
	PhysicalLabelsCertified  bool
	TraceLedgerNeedsPhysical bool
	Supports, Failures       []string
}

type WoundAudit struct {
	PreviousWound      string
	ReducedWound       string
	AbsoluteSignNeeded bool
	NativeSolved       bool
	Supports, Failures []string
}

type FreezeAudit struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
}

type Firewalls struct {
	NativeGlobalZ2       bool
	NativePhaseModuleCR2 bool
	NativeAirlock        bool
	AlphaNative          bool
	HiggsNative          bool
	FullAFDescent        bool
	NativeR3             bool
	PhysicalAssignment   bool
	GenerationCarrier    bool
	FlavorOrientation    bool
	IndividualYukawa     bool
	OfficialLedgerUpdate bool
	NativeYukawaOperator bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedAudit
	Z2             Z2OperationAudit
	Airlock        AirlockEquivarianceAudit
	Edge           EdgeEquivarianceAudit
	Trace          TraceEquivarianceAudit
	Labels         LabelAudit
	Wound          WoundAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.QPhiExists || inherited.PositiveSignSelected {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}
	z2 := buildZ2OperationAudit()
	if !z2.SwapsLambdaBarLambda || !z2.SwapsProjectors || !z2.FlipsQPhi || z2.NativeTheorem {
		return Audit{}, fmt.Errorf("Z2 leak: %s", FormatZ2(z2))
	}
	airlock := buildAirlockEquivarianceAudit()
	if !airlock.AlphaRanksInvariant || !airlock.OrientationClassCandidate || airlock.NativeEquivariance || !sameInts(airlock.PlusFlagRanks, airlock.MinusFlagRanks) {
		return Audit{}, fmt.Errorf("airlock leak: %s", FormatAirlock(airlock))
	}
	edge := buildEdgeEquivarianceAudit()
	if !edge.MirrorRepresentativeExists || !edge.RankKernelInvariant || edge.NativeOperatorTheorem || !sameInts(edge.CurrentRankPattern, edge.MirrorRankPattern) || edge.CurrentKernelCount != edge.MirrorKernelCount {
		return Audit{}, fmt.Errorf("edge leak: %s", FormatEdge(edge))
	}
	trace := buildTraceEquivarianceAudit()
	if !trace.TraceInvariant || !trace.SquareTraceInvariant || !trace.NEffInvariant || !trace.CYukawaInvariant || trace.NativeTraceTheorem {
		return Audit{}, fmt.Errorf("trace leak: %s", FormatTrace(trace))
	}
	labels := buildLabelAudit()
	if !labels.SocketLabelsChange || labels.PhysicalLabelsCertified || labels.TraceLedgerNeedsPhysical {
		return Audit{}, fmt.Errorf("label leak: %s", FormatLabels(labels))
	}
	wound := buildWoundAudit()
	if wound.AbsoluteSignNeeded || wound.NativeSolved || !strings.Contains(wound.ReducedWound, "Z2-equivariant") {
		return Audit{}, fmt.Errorf("wound leak: %s", FormatWound(wound))
	}
	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{ID: AuditID, Classification: Classification, ShortStatus: ShortStatus, Inherited: inherited, Z2: z2, Airlock: airlock, Edge: edge, Trace: trace, Labels: labels, Wound: wound, Freeze: freeze, Firewalls: firewalls, Truth: FinalTruth, Final: "Gate 907 reclassifies the absolute Q_phi sign as a possible global orientation gauge: the two airlock representatives are exchanged by lambda/bar(lambda) reversal, while alpha ranks, edge rank/kernel count, trace row multiset, operator N_eff, and operator C_Yukawa remain invariant. Native R3 is still blocked until a global phase-Z2 equivariant airlock theorem and the remaining phase-module, airlock, alpha, Higgs, and full-descent firewalls are certified."}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{Gate906Classification: "R3_PHASE_WEIGHT_ORIENTATION_OBSTRUCTION", QPhiExists: true, PositiveSignSelected: false, Supports: []string{SupportGate906Inherited}, Failures: []string{FailureNoNativeGlobalZ2}}
}
func buildZ2OperationAudit() Z2OperationAudit {
	return Z2OperationAudit{Operation: Z2Operation, SwapsLambdaBarLambda: true, SwapsProjectors: true, FlipsQPhi: true, NativeTheorem: false, Supports: []string{SupportGlobalZ2Exchanges}, Failures: []string{FailureNoNativeGlobalZ2}}
}
func buildAirlockEquivarianceAudit() AirlockEquivarianceAudit {
	return AirlockEquivarianceAudit{PlusPuncture: "p_lambda=e_lambda tensor P_1", MinusPuncture: "p_barlambda=e_barlambda tensor P_1", PlusFlagRanks: []int{3, 7}, MinusFlagRanks: []int{3, 7}, AlphaRanksInvariant: true, OrientationClassCandidate: true, NativeEquivariance: false, Supports: []string{SupportGlobalZ2Exchanges, SupportAlphaInvariant, SupportAirlockClass}, Failures: []string{FailureNoNativeGlobalZ2, FailureNoNativeAirlock}}
}
func buildEdgeEquivarianceAudit() EdgeEquivarianceAudit {
	return EdgeEquivarianceAudit{CurrentRankPattern: []int{3, 3, 1}, MirrorRankPattern: []int{3, 3, 1}, CurrentKernelCount: 1, MirrorKernelCount: 1, MirrorRepresentativeExists: true, RankKernelInvariant: true, NativeOperatorTheorem: false, Supports: []string{SupportEdgeMirror, SupportEdgeRankKernel, SupportEdgeGaugeRep}, Failures: []string{FailureEdgeEquivarianceNotNative, FailureNoNativeGlobalZ2}}
}
func buildTraceEquivarianceAudit() TraceEquivarianceAudit {
	return TraceEquivarianceAudit{Current: CurrentReadout, Mirror: MirrorReadout, Multiset: WeightMultiset, TraceInvariant: true, SquareTraceInvariant: true, NEffInvariant: true, CYukawaInvariant: true, NativeTraceTheorem: false, Supports: []string{SupportRowMultiset, SupportNEffInvariant, SupportCYukawaInvariant}, Failures: []string{FailureNoNativeGlobalZ2}}
}
func buildLabelAudit() LabelAudit {
	return LabelAudit{SocketLabelsChange: true, PhysicalLabelsCertified: false, TraceLedgerNeedsPhysical: false, Supports: []string{SupportLabelsGauge, SupportNoPhysicalLabelsNeed}, Failures: []string{FailureNoPhysicalAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap}}
}
func buildWoundAudit() WoundAudit {
	return WoundAudit{PreviousWound: "select +Q_phi rather than -Q_phi", ReducedWound: "certify Z2-equivariant airlock orientation class", AbsoluteSignNeeded: false, NativeSolved: false, Supports: []string{SupportQPhiGauge, SupportNativePressureZ2, SupportGeneratorWoundWeakens}, Failures: []string{FailureNoNativeGlobalZ2, FailureNoNativePhaseModuleCR2, FailureNoNativeAirlock}}
}
func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{Alpha: AlphaB, OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false}
}
func buildFirewalls() Firewalls {
	return Firewalls{NativeGlobalZ2: false, NativePhaseModuleCR2: false, NativeAirlock: false, AlphaNative: false, HiggsNative: false, FullAFDescent: false, NativeR3: false, PhysicalAssignment: false, GenerationCarrier: false, FlavorOrientation: false, IndividualYukawa: false, OfficialLedgerUpdate: false, NativeYukawaOperator: false}
}

func Statuses() []string {
	return []string{StatusGate906Inherited, StatusZ2Defined, StatusAirlockEquivariance, StatusEdgeEquivariance, StatusTraceEquivariance, StatusLabelFirewall, StatusWoundReclassified, StatusOfficialFreeze, StatusFirewallVerdict, SupportGlobalZ2Exchanges, SupportAlphaInvariant, SupportEdgeMirror, SupportRowMultiset, SupportNEffInvariant, SupportCYukawaInvariant, SupportQPhiGauge, SupportNativePressureZ2, FailureNoNativeGlobalZ2, FailureNoNativePhaseModuleCR2, FailureNoNativeAirlock, FailureAlphaStillSealed, FailureHiggsStillSealed, FailureNotNativeR3}
}

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("inherited=%s Q_phi_exists=%t positive_sign_selected=%t supports=%s failures=%s", a.Gate906Classification, a.QPhiExists, a.PositiveSignSelected, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatZ2(a Z2OperationAudit) string {
	return fmt.Sprintf("operation=%s swaps_lambda=%t swaps_projectors=%t flips_Qphi=%t native=%t supports=%s failures=%s", a.Operation, a.SwapsLambdaBarLambda, a.SwapsProjectors, a.FlipsQPhi, a.NativeTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatAirlock(a AirlockEquivarianceAudit) string {
	return fmt.Sprintf("plus=%s minus=%s plus_ranks=%v minus_ranks=%v alpha_invariant=%t orientation_class=%t native=%t supports=%s failures=%s", a.PlusPuncture, a.MinusPuncture, a.PlusFlagRanks, a.MinusFlagRanks, a.AlphaRanksInvariant, a.OrientationClassCandidate, a.NativeEquivariance, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatEdge(a EdgeEquivarianceAudit) string {
	return fmt.Sprintf("current_ranks=%v mirror_ranks=%v current_kernel=%d mirror_kernel=%d mirror=%t rank_kernel_invariant=%t native=%t supports=%s failures=%s", a.CurrentRankPattern, a.MirrorRankPattern, a.CurrentKernelCount, a.MirrorKernelCount, a.MirrorRepresentativeExists, a.RankKernelInvariant, a.NativeOperatorTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatTrace(a TraceEquivarianceAudit) string {
	return fmt.Sprintf("current=%s mirror=%s multiset=%s trace_inv=%t square_inv=%t N_eff_inv=%t C_Y_inv=%t native=%t supports=%s failures=%s", a.Current, a.Mirror, a.Multiset, a.TraceInvariant, a.SquareTraceInvariant, a.NEffInvariant, a.CYukawaInvariant, a.NativeTraceTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatLabels(a LabelAudit) string {
	return fmt.Sprintf("socket_labels_change=%t physical_labels=%t trace_needs_physical=%t supports=%s failures=%s", a.SocketLabelsChange, a.PhysicalLabelsCertified, a.TraceLedgerNeedsPhysical, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatWound(a WoundAudit) string {
	return fmt.Sprintf("previous=%q reduced=%q absolute_sign_needed=%t native_solved=%t supports=%s failures=%s", a.PreviousWound, a.ReducedWound, a.AbsoluteSignNeeded, a.NativeSolved, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFreeze(a FreezeAudit) string {
	return fmt.Sprintf("alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g operator_CY=%.16g official_CY=%.16g operator_CH=%.16g official_CH=%.16g frozen=%t diagnostic_only=%t can_update=%t", a.Alpha, a.OperatorNEff, a.OfficialNEff, a.OperatorCYukawa, a.OfficialCYukawa, a.OperatorCHiggs, a.OfficialCHiggs, a.Frozen, a.DiagnosticOnly, a.CanUpdate)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_Z2=%t native_phase_module_CR2=%t native_airlock=%t alpha_native=%t higgs_native=%t full_AF_descent=%t native_R3=%t physical_assignment=%t generation=%t flavor=%t individual_yukawa=%t official_update=%t native_yukawa_operator=%t", f.NativeGlobalZ2, f.NativePhaseModuleCR2, f.NativeAirlock, f.AlphaNative, f.HiggsNative, f.FullAFDescent, f.NativeR3, f.PhysicalAssignment, f.GenerationCarrier, f.FlavorOrientation, f.IndividualYukawa, f.OfficialLedgerUpdate, f.NativeYukawaOperator)
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoNativeGlobalZ2, FailureNoNativePhaseModuleCR2, FailureNoNativeAirlock, FailureAlphaStillSealed, FailureHiggsStillSealed, FailureFullAFDescentBlocked, FailureNotNativeR3, FailureNoPhysicalAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawa, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawa}
}
func firewallsOK(f Firewalls) bool {
	return !f.NativeGlobalZ2 && !f.NativePhaseModuleCR2 && !f.NativeAirlock && !f.AlphaNative && !f.HiggsNative && !f.FullAFDescent && !f.NativeR3 && !f.PhysicalAssignment && !f.GenerationCarrier && !f.FlavorOrientation && !f.IndividualYukawa && !f.OfficialLedgerUpdate && !f.NativeYukawaOperator
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
func sameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
