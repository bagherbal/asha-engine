// Package generation2positivephasegeneratorcharacterweightorientationaudit implements
// Gate 906: PositivePhase Generator and CharacterWeight Orientation Audit.
//
// Gate 906 follows Gate 905's result that C_R^2 has the correct sealed shape
// of a minimal lambda/bar(lambda) phase module, but its orientation is not
// native. It audits whether the remaining lambda-over-bar(lambda) order can be
// reduced to the sign of the phase-weight operator Q_phi=e_lambda-e_barlambda,
// and whether Hopf Reeb orientation or Cl(1,7) chirality sign can source that
// sign. The honest verdict is that the wound is sharpened to a positive phase
// generator / Q_phi sign selection; no native selector, alpha derivation, R3
// promotion, particle assignment, individual Yukawa value, or official ledger
// update is certified.
package generation2positivephasegeneratorcharacterweightorientationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE906-POSITIVE-PHASE-GENERATOR-CHARACTER-WEIGHT-ORIENTATION-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PhaseModule    = "V_phase=C_lambda plus C_barlambda"
	PhaseWeights   = "lambda=e^{+i theta}, bar(lambda)=e^{-i theta}"
	WeightOperator = "Q_phi=e_lambda-e_barlambda"
	RightRhoTarget = "rho_R(lambda)=lambda e_+ + bar(lambda)e_-"
	HopfReebSource = "R=iz"
	CL17Chirality  = "gamma_chi=i omega"

	Classification = "R3_PHASE_WEIGHT_ORIENTATION_OBSTRUCTION"
	ShortStatus    = "R3_AIRLOCK_PHASE_MODULE_INDUCED_POSITIVE_GENERATOR_MISSING"

	StatusGate905Inherited           = "PASS_GATE905_PHASE_MODULE_INDUCTION_PAIR_NOT_ORDER_INHERITED"
	StatusWeightOperatorAudited      = "PASS_PHASE_WEIGHT_OPERATOR_Q_PHI_AUDITED"
	StatusHopfReebAudited            = "PASS_HOPF_REEB_POSITIVE_GENERATOR_SOURCE_AUDITED"
	StatusCL17SignAudited            = "PASS_CL17_CHIRALITY_SIGN_SOURCE_AUDITED"
	StatusJConjugationAudited        = "PASS_J_CONJUGATION_PAIR_NOT_ORIENTATION_AUDITED"
	StatusBoundaryOrientationAudited = "PASS_BOUNDARY_ORIENTATION_FIREWALL_AUDITED"
	StatusWoundSharpened             = "PASS_PHASE_MODULE_ORIENTATION_REDUCED_TO_POSITIVE_GENERATOR_SELECTION"
	StatusOfficialFreeze             = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict            = "FIREWALL_PRESERVED_GATE906_Q_PHI_SIGN_NOT_NATIVE"

	SupportGate905Inherited        = "CONDITIONAL_SUPPORT_GATE905_PHASE_MODULE_INDUCTION_PAIR_NOT_ORDER_INHERITED"
	SupportQPhiExists              = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_PHASE_MODULE_HAS_WEIGHT_OPERATOR_Q_PHI"
	SupportOrderEquivalentQPhi     = "CONDITIONAL_SUPPORT_LAMBDA_BARLAMBDA_ORDER_EQUIVALENT_TO_Q_PHI_SIGN_ORIENTATION"
	SupportSocketOrderAsWeightSign = "CONDITIONAL_SUPPORT_SOCKET_ORDER_CAN_BE_REWRITTEN_AS_PHASE_WEIGHT_SIGN_ORDER"
	SupportPuncturePositiveWeight  = "CONDITIONAL_SUPPORT_PUNCTURE_IS_POSITIVE_PHASE_WEIGHT_LEPTON_SOCKET_IF_Q_PHI_ORIENTED"
	SupportWoundToQPhi             = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_ORDER_REDUCES_TO_ORIENTATION_OF_Q_PHI"
	SupportHopfReebStrongest       = "CONDITIONAL_SUPPORT_HOPF_REEB_DIRECTION_IS_STRONGEST_POSITIVE_PHASE_GENERATOR_SOURCE"
	SupportHopfPositiveWeight      = "CONDITIONAL_SUPPORT_POSITIVE_PHASE_GENERATOR_SELECTS_LAMBDA_WEIGHT_PLUS_ONE"
	SupportHopfOrderIfSealed       = "CONDITIONAL_SUPPORT_LAMBDA_BARLAMBDA_ORDER_CAN_BE_READ_AS_PLUS_MINUS_WEIGHT_ORDER_IF_HOPF_REEB_SEALED"
	SupportCL17SignMatches         = "CONDITIONAL_SUPPORT_CL17_COMPLEX_CHIRALITY_SIGN_MATCHES_PHASE_WEIGHT_SIGN"
	SupportGammaChiCanSource       = "CONDITIONAL_SUPPORT_GAMMA_CHI_ORIENTATION_CAN_SOURCE_Q_PHI_SIGN_IF_TYPED"
	SupportJPair                   = "CONDITIONAL_SUPPORT_J_STRUCTURE_EXPLAINS_PHASE_WEIGHT_CONJUGATION_PAIR"
	SupportOrderedAirlockIfSealed  = "CONDITIONAL_SUPPORT_ORDERED_AIRLOCK_FOLLOWS_IF_POSITIVE_PHASE_GENERATOR_IS_SEALED"
	SupportR3WoundToGenerator      = "CONDITIONAL_SUPPORT_R3_MASTER_WOUND_REDUCES_TO_POSITIVE_PHASE_GENERATOR_SELECTION"
	SupportDiagnosticsRemain       = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTICS_REMAIN_COHERENT_UNDER_POSITIVE_PHASE_GENERATOR_SEAL"

	FailureNoNativePositiveGenerator = "FAILED_ROUTE_NO_NATIVE_POSITIVE_PHASE_GENERATOR_THEOREM"
	FailureNoTypedHopfReebToCR2      = "FAILED_ROUTE_NO_TYPED_HOPF_REEB_TO_C_R2_PHASE_ACTION_MAP"
	FailureHopfReebNotSelector       = "FAILED_ROUTE_HOPF_REEB_ORIENTATION_NOT_YET_NATIVE_RIGHT_CHARACTER_SELECTOR"
	FailureNoTypedGammaChiToQPhi     = "FAILED_ROUTE_NO_TYPED_GAMMA_CHI_SIGN_TO_PHASE_WEIGHT_OPERATOR_MAP"
	FailureCL17DoesNotSelect         = "FAILED_ROUTE_CL17_CHIRALITY_SIGN_DOES_NOT_YET_SELECT_E_LAMBDA_OVER_E_BARLAMBDA"
	FailureQPhiSignNotNative         = "FAILED_ROUTE_Q_PHI_SIGN_IS_NOT_NATIVE_WITHOUT_PHASE_ORIENTATION"
	FailureTwoCharPairNotOrder       = "FAILED_ROUTE_TWO_CHARACTER_MODULE_CERTIFIES_PAIR_NOT_ORDER"
	FailureNoLambdaSelection         = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_LAMBDA_OVER_BARLAMBDA"
	FailureNoPositiveExposure        = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_POSITIVE_PHASE_AS_EXPOSURE_SOCKET"
	FailureJDoesNotOrient            = "FAILED_ROUTE_J_CONJUGATION_CONFIRMS_PAIR_BUT_DOES_NOT_ORIENT_Q_PHI"
	FailureBoundaryNoPhaseSign       = "FAILED_ROUTE_BOUNDARY_ORIENTATION_DOES_NOT_SELECT_PHASE_WEIGHT_SIGN"
	FailurePhaseAnchorSealed         = "FAILED_ROUTE_PHASE_ANCHOR_REMAINS_SEALED"
	FailureAlphaStillSealed          = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealed          = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureNotNativeR3               = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoGenerationCarrierMap    = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap    = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues  = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate      = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate     = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator    = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem   = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedAudit struct {
	Gate905Classification string
	PhaseModuleInduced    bool
	PairCertified         bool
	OrderCertified        bool
	Supports, Failures    []string
}

type WeightOperatorAudit struct {
	LambdaWeight, BarLambdaWeight int
	Operator                      string
	Exists                        bool
	OrderEquivalentToSign         bool
	PositiveSignNative            bool
	Supports, Failures            []string
}

type HopfReebAudit struct {
	Reeb                               string
	SuppliesPositiveGeneratorCandidate bool
	SelectsLambdaWeightIfSealed        bool
	TypedActionOnCR2                   bool
	NativeSelector                     bool
	Supports, Failures                 []string
}

type CL17SignAudit struct {
	GammaChi               string
	SignMatchesPhaseWeight bool
	CanSourceQPhiIfTyped   bool
	TypedMapToQPhi         bool
	SelectsSocketOrder     bool
	Supports, Failures     []string
}

type JConjugationAudit struct {
	ExchangesWeights   bool
	ExplainsPair       bool
	OrientsSign        bool
	Supports, Failures []string
}

type BoundaryOrientationAudit struct {
	OrientsExteriorDegree  bool
	SelectsPhaseWeightSign bool
	Supports, Failures     []string
}

type WoundAudit struct {
	PreviousWound      string
	ReducedWound       string
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
	NativePositiveGenerator bool
	TypedHopfReebToCR2      bool
	TypedGammaChiToQPhi     bool
	QPhiSignNative          bool
	PairOrdersItself        bool
	LambdaSelectionNative   bool
	PhaseAnchorNative       bool
	NativeR3                bool
	PhysicalOrYukawaClaim   bool
	OfficialLedgerUpdated   bool
}

type Audit struct {
	ID                  string
	Classification      string
	ShortStatus         string
	Inherited           InheritedAudit
	WeightOperator      WeightOperatorAudit
	HopfReeb            HopfReebAudit
	CL17Sign            CL17SignAudit
	JConjugation        JConjugationAudit
	BoundaryOrientation BoundaryOrientationAudit
	Wound               WoundAudit
	Freeze              FreezeAudit
	Firewalls           Firewalls
	Truth               string
	Final               string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.PhaseModuleInduced || !inherited.PairCertified || inherited.OrderCertified {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}
	weight := buildWeightOperatorAudit()
	if !weight.Exists || !weight.OrderEquivalentToSign || weight.PositiveSignNative {
		return Audit{}, fmt.Errorf("weight leak: %s", FormatWeightOperator(weight))
	}
	hopf := buildHopfReebAudit()
	if !hopf.SuppliesPositiveGeneratorCandidate || !hopf.SelectsLambdaWeightIfSealed || hopf.TypedActionOnCR2 || hopf.NativeSelector {
		return Audit{}, fmt.Errorf("hopf leak: %s", FormatHopfReeb(hopf))
	}
	cl17 := buildCL17SignAudit()
	if !cl17.SignMatchesPhaseWeight || !cl17.CanSourceQPhiIfTyped || cl17.TypedMapToQPhi || cl17.SelectsSocketOrder {
		return Audit{}, fmt.Errorf("cl17 leak: %s", FormatCL17Sign(cl17))
	}
	j := buildJConjugationAudit()
	if !j.ExchangesWeights || !j.ExplainsPair || j.OrientsSign {
		return Audit{}, fmt.Errorf("J leak: %s", FormatJConjugation(j))
	}
	boundary := buildBoundaryOrientationAudit()
	if !boundary.OrientsExteriorDegree || boundary.SelectsPhaseWeightSign {
		return Audit{}, fmt.Errorf("boundary leak: %s", FormatBoundaryOrientation(boundary))
	}
	wound := buildWoundAudit()
	if wound.NativeSolved || !strings.Contains(wound.ReducedWound, "positive phase generator") {
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
	return Audit{ID: AuditID, Classification: Classification, ShortStatus: ShortStatus, Inherited: inherited, WeightOperator: weight, HopfReeb: hopf, CL17Sign: cl17, JConjugation: j, BoundaryOrientation: boundary, Wound: wound, Freeze: freeze, Firewalls: firewalls, Truth: "PHASE_MODULE_ORIENTATION_REDUCED_TO_POSITIVE_GENERATOR_OR_Q_PHI_SIGN_SELECTION", Final: "Gate 906 compresses the remaining socket-order wound to the sign of Q_phi=e_lambda-e_barlambda. Hopf Reeb direction and Cl(1,7) chirality sign are the strongest sign-source candidates, but no native positive phase generator theorem or typed action on C_R^2 is certified."}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{Gate905Classification: "R3_PHASE_MODULE_INDUCTION_SUPPORT_ORDER_OBSTRUCTION", PhaseModuleInduced: true, PairCertified: true, OrderCertified: false, Supports: []string{SupportGate905Inherited}, Failures: []string{FailureTwoCharPairNotOrder, FailureNoLambdaSelection}}
}

func buildWeightOperatorAudit() WeightOperatorAudit {
	return WeightOperatorAudit{LambdaWeight: +1, BarLambdaWeight: -1, Operator: WeightOperator, Exists: true, OrderEquivalentToSign: true, PositiveSignNative: false, Supports: []string{SupportQPhiExists, SupportOrderEquivalentQPhi, SupportSocketOrderAsWeightSign, SupportPuncturePositiveWeight, SupportWoundToQPhi}, Failures: []string{FailureQPhiSignNotNative, FailureNoLambdaSelection, FailureNoPositiveExposure}}
}

func buildHopfReebAudit() HopfReebAudit {
	return HopfReebAudit{Reeb: HopfReebSource, SuppliesPositiveGeneratorCandidate: true, SelectsLambdaWeightIfSealed: true, TypedActionOnCR2: false, NativeSelector: false, Supports: []string{SupportHopfReebStrongest, SupportHopfPositiveWeight, SupportHopfOrderIfSealed}, Failures: []string{FailureNoTypedHopfReebToCR2, FailureHopfReebNotSelector, FailureNoNativePositiveGenerator}}
}

func buildCL17SignAudit() CL17SignAudit {
	return CL17SignAudit{GammaChi: CL17Chirality, SignMatchesPhaseWeight: true, CanSourceQPhiIfTyped: true, TypedMapToQPhi: false, SelectsSocketOrder: false, Supports: []string{SupportCL17SignMatches, SupportGammaChiCanSource}, Failures: []string{FailureNoTypedGammaChiToQPhi, FailureCL17DoesNotSelect, FailureNoNativePositiveGenerator}}
}

func buildJConjugationAudit() JConjugationAudit {
	return JConjugationAudit{ExchangesWeights: true, ExplainsPair: true, OrientsSign: false, Supports: []string{SupportJPair}, Failures: []string{FailureJDoesNotOrient}}
}

func buildBoundaryOrientationAudit() BoundaryOrientationAudit {
	return BoundaryOrientationAudit{OrientsExteriorDegree: true, SelectsPhaseWeightSign: false, Supports: []string{}, Failures: []string{FailureBoundaryNoPhaseSign}}
}

func buildWoundAudit() WoundAudit {
	return WoundAudit{PreviousWound: "orientation of induced two-character phase module", ReducedWound: "positive phase generator / Q_phi sign selection", NativeSolved: false, Supports: []string{SupportOrderedAirlockIfSealed, SupportR3WoundToGenerator}, Failures: []string{FailureNoNativePositiveGenerator, FailureNoLambdaSelection, FailurePhaseAnchorSealed}}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{Alpha: AlphaB, OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{NativePositiveGenerator: false, TypedHopfReebToCR2: false, TypedGammaChiToQPhi: false, QPhiSignNative: false, PairOrdersItself: false, LambdaSelectionNative: false, PhaseAnchorNative: false, NativeR3: false, PhysicalOrYukawaClaim: false, OfficialLedgerUpdated: false}
}

func Statuses() []string {
	return []string{StatusGate905Inherited, StatusWeightOperatorAudited, StatusHopfReebAudited, StatusCL17SignAudited, StatusJConjugationAudited, StatusBoundaryOrientationAudited, StatusWoundSharpened, StatusOfficialFreeze, StatusFirewallVerdict, SupportQPhiExists, SupportOrderEquivalentQPhi, SupportHopfReebStrongest, SupportCL17SignMatches, SupportR3WoundToGenerator, FailureNoNativePositiveGenerator, FailureNoTypedHopfReebToCR2, FailureNoTypedGammaChiToQPhi, FailureQPhiSignNotNative, FailureNoLambdaSelection, FailureNotNativeR3}
}

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("inherited=%s module_induced=%t pair=%t order=%t supports=%s failures=%s", a.Gate905Classification, a.PhaseModuleInduced, a.PairCertified, a.OrderCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatWeightOperator(a WeightOperatorAudit) string {
	return fmt.Sprintf("operator=%s weights=(%+d,%+d) exists=%t order_equiv=%t native_sign=%t supports=%s failures=%s", a.Operator, a.LambdaWeight, a.BarLambdaWeight, a.Exists, a.OrderEquivalentToSign, a.PositiveSignNative, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatHopfReeb(a HopfReebAudit) string {
	return fmt.Sprintf("Reeb=%s candidate=%t sealed_selects_lambda=%t typed_CR2=%t native=%t supports=%s failures=%s", a.Reeb, a.SuppliesPositiveGeneratorCandidate, a.SelectsLambdaWeightIfSealed, a.TypedActionOnCR2, a.NativeSelector, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatCL17Sign(a CL17SignAudit) string {
	return fmt.Sprintf("gamma=%s sign_matches=%t can_source_if_typed=%t typed_Qphi=%t selects=%t supports=%s failures=%s", a.GammaChi, a.SignMatchesPhaseWeight, a.CanSourceQPhiIfTyped, a.TypedMapToQPhi, a.SelectsSocketOrder, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatJConjugation(a JConjugationAudit) string {
	return fmt.Sprintf("J_exchanges_weights=%t explains_pair=%t orients_sign=%t supports=%s failures=%s", a.ExchangesWeights, a.ExplainsPair, a.OrientsSign, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatBoundaryOrientation(a BoundaryOrientationAudit) string {
	return fmt.Sprintf("boundary_orients_degree=%t selects_phase_weight_sign=%t supports=%s failures=%s", a.OrientsExteriorDegree, a.SelectsPhaseWeightSign, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatWound(a WoundAudit) string {
	return fmt.Sprintf("previous=%q reduced=%q native_solved=%t supports=%s failures=%s", a.PreviousWound, a.ReducedWound, a.NativeSolved, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFreeze(a FreezeAudit) string {
	return fmt.Sprintf("alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g operator_CY=%.16g official_CY=%.16g operator_CH=%.16g official_CH=%.16g frozen=%t diagnostic_only=%t can_update=%t", a.Alpha, a.OperatorNEff, a.OfficialNEff, a.OperatorCYukawa, a.OfficialCYukawa, a.OperatorCHiggs, a.OfficialCHiggs, a.Frozen, a.DiagnosticOnly, a.CanUpdate)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_positive_generator=%t typed_hopf_CR2=%t typed_gamma_Qphi=%t Qphi_native=%t pair_orders_itself=%t lambda_native=%t phase_anchor_native=%t native_R3=%t physical_or_yukawa_claim=%t official_update=%t", f.NativePositiveGenerator, f.TypedHopfReebToCR2, f.TypedGammaChiToQPhi, f.QPhiSignNative, f.PairOrdersItself, f.LambdaSelectionNative, f.PhaseAnchorNative, f.NativeR3, f.PhysicalOrYukawaClaim, f.OfficialLedgerUpdated)
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoNativePositiveGenerator, FailureNoTypedHopfReebToCR2, FailureNoTypedGammaChiToQPhi, FailureQPhiSignNotNative, FailureTwoCharPairNotOrder, FailureNoLambdaSelection, FailureNoPositiveExposure, FailurePhaseAnchorSealed, FailureAlphaStillSealed, FailureHiggsStillSealed, FailureNotNativeR3, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return !f.NativePositiveGenerator && !f.TypedHopfReebToCR2 && !f.TypedGammaChiToQPhi && !f.QPhiSignNative && !f.PairOrdersItself && !f.LambdaSelectionNative && !f.PhaseAnchorNative && !f.NativeR3 && !f.PhysicalOrYukawaClaim && !f.OfficialLedgerUpdated
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
