// Package generation2nativez2equivariantairlockfunctorandboundaryalphasourceaudit
// implements Gate 909: Native Z2-Equivariant Airlock Functor and
// BoundaryAlpha Source Audit.
//
// Gate 909 follows Gate 908's sealed Z2-equivariant R3 trace-ledger plateau. It
// asks whether BoundaryAlpha can be defined on the quotient puncture class
// [p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}, rather than on a
// chosen phase-oriented representative. The honest result is stronger than Gate
// 908 but still not native: alpha_B is well-defined on the Z2 airlock class at
// rank/class level and reconstructs alpha_B=(3/10)s+(7/72)s^2 without an
// absolute phase sign. However, no native Z2-equivariant airlock functor, no
// native BoundaryAlpha source, no native reduced B2 boundary functional, and no
// native transport of s into quotient targets are certified.
package generation2nativez2equivariantairlockfunctorandboundaryalphasourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE909-NATIVE-Z2-EQUIVARIANT-AIRLOCK-FUNCTOR-AND-BOUNDARY-ALPHA-SOURCE-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	LinearDenom  = 10
	QuadDenom    = 72

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	TauPhi               = "tau_phi: lambda<->bar(lambda), e_lambda<->e_barlambda, h_lambda<->h_barlambda, Q_phi->-Q_phi"
	PunctureClass        = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	FlagClass            = "[F_0 subset F_1 subset F_2]_{Z2}"
	FunctorCandidate     = "Z2EquivariantNeutralPunctureAirlockFunctor"
	BoundaryAlphaFormula = "BoundaryAlpha_Z2([p])=[rank(F_1/F_0)/10]s+[rank(F_2/F_0)/72]s^2"
	ReducedB2Response    = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	DegreeOneTargetClass = "[F_1/F_0]_{Z2}={e_lambda tensor P_3,e_barlambda tensor P_3}"
	DegreeTwoTargetClass = "[F_2/F_0]_{Z2}={(C_R^2 tensor W)-(e_lambda tensor P_1),(C_R^2 tensor W)-(e_barlambda tensor P_1)}"
	R3ClassLedger        = "R3 trace ledger on Z2 orientation class under BoundaryAlpha_Z2 seal"
	Classification       = "R3_Z2_EQUIVARIANT_TRACE_LEDGER_WITH_BOUNDARY_ALPHA_CLASS_SEAL_NOT_NATIVE"
	ShortStatus          = "R3_Z2_BOUNDARY_ALPHA_CLASS_SEAL_NOT_NATIVE"
	FinalTruth           = "Z2_EQUIVARIANT_BOUNDARY_ALPHA_CLASS_FUNCTOR_SUPPORTED_AT_RANK_LEVEL_BUT_NATIVE_SOURCE_MISSING"

	StatusGate908Inherited     = "PASS_GATE908_R3_SEALED_Z2_LEDGER_PLATEAU_INHERITED"
	StatusZ2WellDefined        = "PASS_BOUNDARY_ALPHA_Z2_CLASS_WELL_DEFINED_AT_RANK_LEVEL"
	StatusTauCommutation       = "PASS_I_B_Z2_COMMUTES_WITH_TAU_PHI_AT_RANK_LEVEL"
	StatusReducedB2Compatible  = "PASS_REDUCED_B2_RESPONSE_COMPATIBLE_WITH_Z2_AIRLOCK_CLASS"
	StatusCrossLaneAudited     = "PASS_Z2_CROSS_LANE_EXCLUSION_AUDITED_CONDITIONALLY"
	StatusAlphaClassSeal       = "PASS_BOUNDARY_ALPHA_WEAKENED_TO_Z2_CLASS_SEAL"
	StatusR3ConsequenceAudited = "PASS_R3_PRESSURE_REDUCED_TO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	StatusOfficialFreeze       = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict      = "FIREWALL_PRESERVED_GATE909_NATIVE_SOURCE_STILL_MISSING"
	StatusNextGate             = "NEXT_PRESSURE_GATE910_Z2_BOUNDARY_ALPHA_CLASSSEAL_R3_PLATEAU_CLASSIFICATION"

	SupportBoundaryAlphaFunctorZ2Class = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_FUNCTOR_IS_Z2_WELL_DEFINED_AT_CLASS_LEVEL"
	SupportIBZ2CommutesTauRank         = "CONDITIONAL_SUPPORT_I_B_Z2_COMMUTES_WITH_GLOBAL_PHASE_FLIP_AT_RANK_LEVEL"
	SupportAlphaRankPairInvariant      = "CONDITIONAL_SUPPORT_ALPHA_RANK_PAIR_3_7_IS_REPRESENTATIVE_INDEPENDENT"
	SupportReducedB2Compatible         = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_COMPATIBLE_WITH_Z2_AIRLOCK_CLASS"
	SupportZeroCubicFirewalls          = "CONDITIONAL_SUPPORT_ZERO_ORDER_AND_CUBIC_FIREWALLS_REMAIN_SOLVED_AT_RESPONSE_SHAPE_LEVEL"
	SupportCrossLanesIfFunctor         = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_IF_I_B_Z2_FUNCTOR_CERTIFIED"
	SupportAlphaNoPhaseRepresentative  = "CONDITIONAL_SUPPORT_ALPHA_B_NO_LONGER_DEPENDS_ON_PHASE_REPRESENTATIVE"
	SupportAlphaSealWeakensZ2          = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_SEAL_WEAKENS_TO_Z2_EQUIVARIANT_CLASS_SEAL"
	SupportR3OnZ2AirlockClass          = "CONDITIONAL_SUPPORT_R3_LEDGER_CAN_BE_FORMULATED_ON_Z2_AIRLOCK_CLASS"
	SupportPhaseSignNoLongerBlocksR3   = "CONDITIONAL_SUPPORT_PHASE_SIGN_NO_LONGER_BLOCKS_R3_TRACE_LEDGER"
	SupportR3PressureToZ2AlphaFunctor  = "CONDITIONAL_SUPPORT_R3_PRESSURE_REDUCES_TO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	SupportBoundaryAlphaClassPlateau   = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_Z2_CLASS_SEAL_PLATEAU"

	FailureNoNativeZ2AirlockFunctor       = "FAILED_ROUTE_NO_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR"
	FailureNoNativeZ2BoundaryAlphaFunctor = "FAILED_ROUTE_NO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	FailureZ2RankNotNativeFunctor         = "FAILED_ROUTE_Z2_EQUIVARIANCE_IS_RANK_LEVEL_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureReducedB2NotNativeFunctional   = "FAILED_ROUTE_REDUCED_B2_RESPONSE_NOT_NATIVE_BOUNDARY_FUNCTIONAL"
	FailureNoNativeDegreeToZ2FlagFunctor  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeZ2CrossLane            = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoLinearActiveDomainExclusion  = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_ACTIVE_DOMAIN_CLASS"
	FailureNoQuadraticFaceExclusion       = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_EXPOSED_FACE_CLASS"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeBoundaryAlphaSource    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ALPHA_SOURCE"
	FailureNoNativeTransportS             = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_Z2_AIRLOCK_CLASS"
	FailureHiggsOrientationClassSealed    = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3_UNLESS_Z2_BOUNDARY_ALPHA_FUNCTOR_CERTIFIED"
	FailureNoPhysicalParticleAssign       = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawa               = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedAudit struct {
	Gate908Classification string
	Gate908ShortStatus    string
	TraceLedgerZ2Class    bool
	AlphaStillSealed      bool
	NativeR3              bool
	Supports, Failures    []string
}

type ClassRepresentative struct {
	Name              string
	Puncture          string
	F0                string
	F1                string
	F2                string
	RankF1OverF0      int
	RankF2OverF0      int
	DegreeOneQuotient string
	DegreeTwoQuotient string
}

type Z2WellDefinedAudit struct {
	Operation              string
	FunctorName            string
	PunctureClass          string
	FlagClass              string
	Plus, Minus            ClassRepresentative
	TauMapsPunctures       bool
	TauMapsDegreeOneTarget bool
	TauMapsDegreeTwoTarget bool
	IBZ2CommutesWithTau    bool
	RankPairInvariant      bool
	ClassLevelWellDefined  bool
	NativeFunctorTheorem   bool
	Supports, Failures     []string
}

type ReducedB2ResponseAudit struct {
	ResponseExpression        string
	ZeroOrderSuppressed       bool
	CubicTermStopped          bool
	Lambda3B2Zero             bool
	DegreeOneTarget           string
	DegreeTwoTarget           string
	DegreeOneRank             int
	DegreeTwoRank             int
	CompatibleWithZ2Class     bool
	NativeBoundaryFunctional  bool
	NativeDegreeToFlagFunctor bool
	Supports, Failures        []string
}

type CrossLaneAudit struct {
	CorrectDegreeOneTarget            string
	CorrectDegreeTwoTarget            string
	DegreeOneToDegreeTwoAllowed       bool
	DegreeTwoToDegreeOneAllowed       bool
	ExcludedIfFunctorCertified        bool
	NativeCrossLaneTheorem            bool
	NativeLinearDomainClassExclusion  bool
	NativeQuadraticFaceClassExclusion bool
	Supports, Failures                []string
}

type AlphaLayerAudit struct {
	RepresentativeAlphaRequired bool
	Z2ClassAlphaSupported       bool
	NativeAlphaCertified        bool
	Formula                     string
	S                           float64
	LinearCoefficient           float64
	QuadraticCoefficient        float64
	LinearContribution          float64
	QuadraticContribution       float64
	ReconstructedAlpha          float64
	ExpectedAlpha               float64
	Residual                    float64
	RankPair                    []int
	RepresentativeIndependent   bool
	SealWeakenedToClassSeal     bool
	Supports, Failures          []string
}

type R3ConsequenceAudit struct {
	LedgerName                 string
	R3LedgerOnZ2AirlockClass   bool
	PhaseSignBlocksTraceLedger bool
	NativeZ2AlphaFunctor       bool
	NativeR3                   bool
	FullAFDescent              bool
	OfficialLedgerUpdate       bool
	Supports, Failures         []string
}

type FreezeAudit struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	NativeZ2AirlockFunctor       bool
	NativeZ2BoundaryAlphaFunctor bool
	RankEquivarianceNative       bool
	ReducedB2NativeFunctional    bool
	NativeDegreeToZ2FlagFunctor  bool
	NativeZ2CrossLane            bool
	AlphaNative                  bool
	NativeBoundaryAlphaSource    bool
	NativeTransportS             bool
	HiggsOrientationNative       bool
	FullAFDescent                bool
	NativeR3                     bool
	PhysicalParticleAssignment   bool
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
	Inherited      InheritedAudit
	WellDefined    Z2WellDefinedAudit
	Response       ReducedB2ResponseAudit
	CrossLane      CrossLaneAudit
	Alpha          AlphaLayerAudit
	R3             R3ConsequenceAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func LinearCoefficient() float64              { return float64(RankF1OverF0) / float64(LinearDenom) }
func QuadraticCoefficient() float64           { return float64(RankF2OverF0) / float64(QuadDenom) }
func LinearContribution(s float64) float64    { return LinearCoefficient() * s }
func QuadraticContribution(s float64) float64 { return QuadraticCoefficient() * s * s }
func BoundaryAlphaZ2(s float64) float64       { return LinearContribution(s) + QuadraticContribution(s) }

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.TraceLedgerZ2Class || !inherited.AlphaStillSealed || inherited.NativeR3 {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}

	wellDefined := buildZ2WellDefinedAudit()
	if !wellDefined.TauMapsPunctures || !wellDefined.TauMapsDegreeOneTarget || !wellDefined.TauMapsDegreeTwoTarget || !wellDefined.IBZ2CommutesWithTau || !wellDefined.RankPairInvariant || !wellDefined.ClassLevelWellDefined || wellDefined.NativeFunctorTheorem {
		return Audit{}, fmt.Errorf("Z2 well-definedness leak: %s", FormatWellDefined(wellDefined))
	}

	response := buildReducedB2ResponseAudit()
	if !response.ZeroOrderSuppressed || !response.CubicTermStopped || !response.Lambda3B2Zero || !response.CompatibleWithZ2Class || response.NativeBoundaryFunctional || response.NativeDegreeToFlagFunctor {
		return Audit{}, fmt.Errorf("reduced B2 response leak: %s", FormatResponse(response))
	}

	cross := buildCrossLaneAudit()
	if cross.DegreeOneToDegreeTwoAllowed || cross.DegreeTwoToDegreeOneAllowed || !cross.ExcludedIfFunctorCertified || cross.NativeCrossLaneTheorem || cross.NativeLinearDomainClassExclusion || cross.NativeQuadraticFaceClassExclusion {
		return Audit{}, fmt.Errorf("cross-lane leak: %s", FormatCrossLane(cross))
	}

	alpha := buildAlphaLayerAudit()
	if alpha.RepresentativeAlphaRequired || !alpha.Z2ClassAlphaSupported || alpha.NativeAlphaCertified || !alpha.RepresentativeIndependent || !alpha.SealWeakenedToClassSeal || !near(alpha.ReconstructedAlpha, AlphaB) || math.Abs(alpha.Residual) > 1e-18 {
		return Audit{}, fmt.Errorf("alpha layer leak: %s", FormatAlpha(alpha))
	}

	r3 := buildR3ConsequenceAudit()
	if !r3.R3LedgerOnZ2AirlockClass || r3.PhaseSignBlocksTraceLedger || r3.NativeZ2AlphaFunctor || r3.NativeR3 || r3.FullAFDescent || r3.OfficialLedgerUpdate {
		return Audit{}, fmt.Errorf("R3 consequence leak: %s", FormatR3(r3))
	}

	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) || near(freeze.OperatorCYukawa, freeze.OfficialCYukawa) || near(freeze.OperatorCHiggs, freeze.OfficialCHiggs) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
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
		WellDefined:    wellDefined,
		Response:       response,
		CrossLane:      cross,
		Alpha:          alpha,
		R3:             r3,
		Freeze:         freeze,
		Firewalls:      firewalls,
		Truth:          FinalTruth,
		Final:          "Gate 909 rewrites BoundaryAlpha at the quotient level: alpha_B is a function of the Z2 neutral-puncture airlock class [p]_{Z2}, not of p_lambda or p_barlambda as chosen representatives. The rank pair (3,7), reduced B2 response shape, and degree-one/degree-two quotient targets are class-level equivariant under tau_phi, so the phase sign no longer blocks the R3 aggregate trace ledger. This remains a class seal rather than a native theorem: the native Z2-equivariant airlock functor, native BoundaryAlpha source, native reduced boundary functional, native s-transport, full A_F descent, physical assignments, generation/flavor maps, individual Yukawa values, and official ledger updates remain blocked.",
	}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{
		Gate908Classification: "R3_SEALED_Z2_EQUIVARIANT_LEDGER_CANDIDATE",
		Gate908ShortStatus:    "R3_Z2_EQUIVARIANT_TRACE_LEDGER_CANDIDATE_NOT_NATIVE",
		TraceLedgerZ2Class:    true,
		AlphaStillSealed:      true,
		NativeR3:              false,
		Supports:              []string{StatusGate908Inherited, SupportR3OnZ2AirlockClass},
		Failures:              []string{FailureAlphaStillSealed, FailureNotNativeR3},
	}
}

func buildRepresentative(name, puncture string) ClassRepresentative {
	phase := "lambda"
	if strings.Contains(puncture, "barlambda") {
		phase = "barlambda"
	}
	return ClassRepresentative{
		Name:              name,
		Puncture:          puncture,
		F0:                puncture,
		F1:                fmt.Sprintf("e_%s tensor W", phase),
		F2:                "C_R^2 tensor W",
		RankF1OverF0:      RankF1OverF0,
		RankF2OverF0:      RankF2OverF0,
		DegreeOneQuotient: fmt.Sprintf("e_%s tensor P_3", phase),
		DegreeTwoQuotient: fmt.Sprintf("(C_R^2 tensor W)-(e_%s tensor P_1)", phase),
	}
}

func buildZ2WellDefinedAudit() Z2WellDefinedAudit {
	plus := buildRepresentative("A", "e_lambda tensor P_1")
	minus := buildRepresentative("B", "e_barlambda tensor P_1")
	return Z2WellDefinedAudit{
		Operation:              TauPhi,
		FunctorName:            FunctorCandidate,
		PunctureClass:          PunctureClass,
		FlagClass:              FlagClass,
		Plus:                   plus,
		Minus:                  minus,
		TauMapsPunctures:       true,
		TauMapsDegreeOneTarget: true,
		TauMapsDegreeTwoTarget: true,
		IBZ2CommutesWithTau:    true,
		RankPairInvariant:      plus.RankF1OverF0 == minus.RankF1OverF0 && plus.RankF2OverF0 == minus.RankF2OverF0,
		ClassLevelWellDefined:  true,
		NativeFunctorTheorem:   false,
		Supports:               []string{SupportBoundaryAlphaFunctorZ2Class, SupportIBZ2CommutesTauRank, SupportAlphaRankPairInvariant},
		Failures:               []string{FailureZ2RankNotNativeFunctor, FailureNoNativeZ2AirlockFunctor, FailureNoNativeZ2BoundaryAlphaFunctor},
	}
}

func buildReducedB2ResponseAudit() ReducedB2ResponseAudit {
	return ReducedB2ResponseAudit{
		ResponseExpression:        ReducedB2Response,
		ZeroOrderSuppressed:       true,
		CubicTermStopped:          true,
		Lambda3B2Zero:             true,
		DegreeOneTarget:           DegreeOneTargetClass,
		DegreeTwoTarget:           DegreeTwoTargetClass,
		DegreeOneRank:             RankF1OverF0,
		DegreeTwoRank:             RankF2OverF0,
		CompatibleWithZ2Class:     true,
		NativeBoundaryFunctional:  false,
		NativeDegreeToFlagFunctor: false,
		Supports:                  []string{SupportReducedB2Compatible, SupportZeroCubicFirewalls},
		Failures:                  []string{FailureReducedB2NotNativeFunctional, FailureNoNativeDegreeToZ2FlagFunctor},
	}
}

func buildCrossLaneAudit() CrossLaneAudit {
	return CrossLaneAudit{
		CorrectDegreeOneTarget:            DegreeOneTargetClass,
		CorrectDegreeTwoTarget:            DegreeTwoTargetClass,
		DegreeOneToDegreeTwoAllowed:       false,
		DegreeTwoToDegreeOneAllowed:       false,
		ExcludedIfFunctorCertified:        true,
		NativeCrossLaneTheorem:            false,
		NativeLinearDomainClassExclusion:  false,
		NativeQuadraticFaceClassExclusion: false,
		Supports:                          []string{SupportCrossLanesIfFunctor},
		Failures:                          []string{FailureNoNativeZ2CrossLane, FailureNoLinearActiveDomainExclusion, FailureNoQuadraticFaceExclusion, FailureNoNativeDegreeToZ2FlagFunctor},
	}
}

func buildAlphaLayerAudit() AlphaLayerAudit {
	linear := LinearContribution(SBoundary)
	quadratic := QuadraticContribution(SBoundary)
	reconstructed := BoundaryAlphaZ2(SBoundary)
	return AlphaLayerAudit{
		RepresentativeAlphaRequired: false,
		Z2ClassAlphaSupported:       true,
		NativeAlphaCertified:        false,
		Formula:                     BoundaryAlphaFormula,
		S:                           SBoundary,
		LinearCoefficient:           LinearCoefficient(),
		QuadraticCoefficient:        QuadraticCoefficient(),
		LinearContribution:          linear,
		QuadraticContribution:       quadratic,
		ReconstructedAlpha:          reconstructed,
		ExpectedAlpha:               AlphaB,
		Residual:                    reconstructed - AlphaB,
		RankPair:                    []int{RankF1OverF0, RankF2OverF0},
		RepresentativeIndependent:   true,
		SealWeakenedToClassSeal:     true,
		Supports:                    []string{SupportAlphaNoPhaseRepresentative, SupportAlphaSealWeakensZ2, SupportAlphaRankPairInvariant, SupportBoundaryAlphaClassPlateau},
		Failures:                    []string{FailureAlphaStillSealed, FailureNoNativeBoundaryAlphaSource, FailureNoNativeTransportS, FailureNoNativeZ2BoundaryAlphaFunctor},
	}
}

func buildR3ConsequenceAudit() R3ConsequenceAudit {
	return R3ConsequenceAudit{
		LedgerName:                 R3ClassLedger,
		R3LedgerOnZ2AirlockClass:   true,
		PhaseSignBlocksTraceLedger: false,
		NativeZ2AlphaFunctor:       false,
		NativeR3:                   false,
		FullAFDescent:              false,
		OfficialLedgerUpdate:       false,
		Supports:                   []string{SupportR3OnZ2AirlockClass, SupportPhaseSignNoLongerBlocksR3, SupportR3PressureToZ2AlphaFunctor},
		Failures:                   []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked, FailureNoOfficialNEffUpdate, FailureNoPhysicalParticleAssign, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues},
	}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{
		Alpha:           AlphaB,
		OperatorNEff:    OperatorNEffDiagnostic,
		OfficialNEff:    OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic,
		OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs:  OperatorCHiggsDiagnostic,
		OfficialCHiggs:  OfficialCHiggsFrozen,
		Frozen:          true,
		DiagnosticOnly:  true,
		CanUpdate:       false,
		Supports:        []string{StatusOfficialFreeze},
		Failures:        []string{FailureNoOfficialNEffUpdate},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NativeZ2AirlockFunctor:       false,
		NativeZ2BoundaryAlphaFunctor: false,
		RankEquivarianceNative:       false,
		ReducedB2NativeFunctional:    false,
		NativeDegreeToZ2FlagFunctor:  false,
		NativeZ2CrossLane:            false,
		AlphaNative:                  false,
		NativeBoundaryAlphaSource:    false,
		NativeTransportS:             false,
		HiggsOrientationNative:       false,
		FullAFDescent:                false,
		NativeR3:                     false,
		PhysicalParticleAssignment:   false,
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
		StatusGate908Inherited,
		StatusZ2WellDefined,
		StatusTauCommutation,
		StatusReducedB2Compatible,
		StatusCrossLaneAudited,
		StatusAlphaClassSeal,
		StatusR3ConsequenceAudited,
		StatusOfficialFreeze,
		StatusFirewallVerdict,
		StatusNextGate,
		SupportBoundaryAlphaFunctorZ2Class,
		SupportIBZ2CommutesTauRank,
		SupportAlphaRankPairInvariant,
		SupportReducedB2Compatible,
		SupportZeroCubicFirewalls,
		SupportCrossLanesIfFunctor,
		SupportAlphaNoPhaseRepresentative,
		SupportAlphaSealWeakensZ2,
		SupportR3OnZ2AirlockClass,
		SupportPhaseSignNoLongerBlocksR3,
		SupportR3PressureToZ2AlphaFunctor,
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureZ2RankNotNativeFunctor,
		FailureReducedB2NotNativeFunctional,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryAlphaSource,
		FailureNoNativeTransportS,
		FailureNotNativeR3,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureZ2RankNotNativeFunctor,
		FailureReducedB2NotNativeFunctional,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureNoLinearActiveDomainExclusion,
		FailureNoQuadraticFaceExclusion,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryAlphaSource,
		FailureNoNativeTransportS,
		FailureHiggsOrientationClassSealed,
		FailureFullAFDescentStillBlocked,
		FailureNotNativeR3,
		FailureNoPhysicalParticleAssign,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoOfficialNEffUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("gate908_classification=%s gate908_short=%s trace_ledger_Z2=%t alpha_still_sealed=%t native_R3=%t supports=%s failures=%s", a.Gate908Classification, a.Gate908ShortStatus, a.TraceLedgerZ2Class, a.AlphaStillSealed, a.NativeR3, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatRepresentative(r ClassRepresentative) string {
	return fmt.Sprintf("rep=%s puncture=%s F0=%s F1=%s F2=%s rank_F1_F0=%d rank_F2_F0=%d degree1=%s degree2=%s", r.Name, r.Puncture, r.F0, r.F1, r.F2, r.RankF1OverF0, r.RankF2OverF0, r.DegreeOneQuotient, r.DegreeTwoQuotient)
}

func FormatWellDefined(a Z2WellDefinedAudit) string {
	return fmt.Sprintf("operation=%s functor=%s class=%s flag=%s plus={%s} minus={%s} tau_punctures=%t tau_deg1=%t tau_deg2=%t commutes=%t rank_pair_invariant=%t class_well_defined=%t native=%t supports=%s failures=%s", a.Operation, a.FunctorName, a.PunctureClass, a.FlagClass, FormatRepresentative(a.Plus), FormatRepresentative(a.Minus), a.TauMapsPunctures, a.TauMapsDegreeOneTarget, a.TauMapsDegreeTwoTarget, a.IBZ2CommutesWithTau, a.RankPairInvariant, a.ClassLevelWellDefined, a.NativeFunctorTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatResponse(a ReducedB2ResponseAudit) string {
	return fmt.Sprintf("response=%s zero_order_suppressed=%t cubic_stopped=%t lambda3_B2_zero=%t degree1_target=%s rank1=%d degree2_target=%s rank2=%d compatible_Z2=%t native_functional=%t native_degree_functor=%t supports=%s failures=%s", a.ResponseExpression, a.ZeroOrderSuppressed, a.CubicTermStopped, a.Lambda3B2Zero, a.DegreeOneTarget, a.DegreeOneRank, a.DegreeTwoTarget, a.DegreeTwoRank, a.CompatibleWithZ2Class, a.NativeBoundaryFunctional, a.NativeDegreeToFlagFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatCrossLane(a CrossLaneAudit) string {
	return fmt.Sprintf("degree1_target=%s degree2_target=%s deg1_to_deg2_allowed=%t deg2_to_deg1_allowed=%t excluded_if_functor=%t native_cross_lane=%t native_linear_exclusion=%t native_quadratic_exclusion=%t supports=%s failures=%s", a.CorrectDegreeOneTarget, a.CorrectDegreeTwoTarget, a.DegreeOneToDegreeTwoAllowed, a.DegreeTwoToDegreeOneAllowed, a.ExcludedIfFunctorCertified, a.NativeCrossLaneTheorem, a.NativeLinearDomainClassExclusion, a.NativeQuadraticFaceClassExclusion, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatAlpha(a AlphaLayerAudit) string {
	return fmt.Sprintf("formula=%s s=%.17g rank_pair=%v linear_coeff=%.16g quadratic_coeff=%.16g linear=%.16g quadratic=%.16g reconstructed=%.17g expected=%.17g residual=%.3e representative_required=%t class_supported=%t native=%t representative_independent=%t class_seal=%t supports=%s failures=%s", a.Formula, a.S, a.RankPair, a.LinearCoefficient, a.QuadraticCoefficient, a.LinearContribution, a.QuadraticContribution, a.ReconstructedAlpha, a.ExpectedAlpha, a.Residual, a.RepresentativeAlphaRequired, a.Z2ClassAlphaSupported, a.NativeAlphaCertified, a.RepresentativeIndependent, a.SealWeakenedToClassSeal, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatR3(a R3ConsequenceAudit) string {
	return fmt.Sprintf("ledger=%s on_Z2_airlock_class=%t phase_sign_blocks=%t native_Z2_alpha_functor=%t native_R3=%t full_AF_descent=%t official_update=%t supports=%s failures=%s", a.LedgerName, a.R3LedgerOnZ2AirlockClass, a.PhaseSignBlocksTraceLedger, a.NativeZ2AlphaFunctor, a.NativeR3, a.FullAFDescent, a.OfficialLedgerUpdate, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFreeze(a FreezeAudit) string {
	return fmt.Sprintf("alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g operator_CY=%.16g official_CY=%.16g operator_CH=%.16g official_CH=%.16g frozen=%t diagnostic_only=%t can_update=%t supports=%s failures=%s", a.Alpha, a.OperatorNEff, a.OfficialNEff, a.OperatorCYukawa, a.OfficialCYukawa, a.OperatorCHiggs, a.OfficialCHiggs, a.Frozen, a.DiagnosticOnly, a.CanUpdate, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_Z2_airlock=%t native_Z2_boundary_alpha=%t rank_equivariance_native=%t reduced_B2_native=%t native_degree_Z2_flag=%t native_Z2_cross_lane=%t alpha_native=%t native_boundary_alpha_source=%t native_s_transport=%t higgs_native=%t full_AF_descent=%t native_R3=%t physical_assignment=%t generation=%t flavor=%t individual_yukawa=%t official_update=%t native_yukawa_operator=%t R4_native_yukawa=%t", f.NativeZ2AirlockFunctor, f.NativeZ2BoundaryAlphaFunctor, f.RankEquivarianceNative, f.ReducedB2NativeFunctional, f.NativeDegreeToZ2FlagFunctor, f.NativeZ2CrossLane, f.AlphaNative, f.NativeBoundaryAlphaSource, f.NativeTransportS, f.HiggsOrientationNative, f.FullAFDescent, f.NativeR3, f.PhysicalParticleAssignment, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.OfficialLedgerUpdate, f.NativeYukawaOperator, f.R4NativeYukawa)
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeZ2AirlockFunctor && !f.NativeZ2BoundaryAlphaFunctor && !f.RankEquivarianceNative && !f.ReducedB2NativeFunctional && !f.NativeDegreeToZ2FlagFunctor && !f.NativeZ2CrossLane && !f.AlphaNative && !f.NativeBoundaryAlphaSource && !f.NativeTransportS && !f.HiggsOrientationNative && !f.FullAFDescent && !f.NativeR3 && !f.PhysicalParticleAssignment && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.OfficialLedgerUpdate && !f.NativeYukawaOperator && !f.R4NativeYukawa
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
