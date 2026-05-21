// Package generation2airlockadmissiblesupportlatticesourceaudit implements
// Gate 931: Airlock AdmissibleSupport Lattice Source Audit.
//
// Gate 931 follows Gate 930's bridge-level AirlockSupportClosureOperator. Gate
// 930 showed that a closure operator exists and has the expected closure laws if
// the candidate support chain {F_0,F_1,F_2} is accepted. Gate 931 audits the
// source type of that support chain itself: puncture rootedness, same-socket
// completion, tensor-factor integrity, no orphan opposite-socket fragments, full
// right-rectangle saturation, minimal sufficiency, and Z2 phase-orientation
// descent. It source-types the chain, but preserves that a native admissibility
// theorem is still missing.
package generation2airlockadmissiblesupportlatticesourceaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE931-AIRLOCK-ADMISSIBLESUPPORT-LATTICE-SOURCE-AUDIT"

	Gate930ShortStatus = "R3_ALPHA_CLOSURE_OPERATOR_EXISTS_NATIVE_SOURCE_MISSING"

	AmbientRightSupportRectangle = "H_R^ambient=C_R^2 tensor W"
	CR2Decomposition             = "C_R^2=e_lambda plus e_barlambda"
	WDecomposition               = "W=P_1 plus P_3"
	AtomicCells                  = "{e_lambda tensor P_1, e_lambda tensor P_3, e_barlambda tensor P_1, e_barlambda tensor P_3}"
	Z2PunctureClass              = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	RepresentativePuncture       = "p=e_phase tensor P_1"
	AdmissibleSupportChain       = "A_airlock={F_0,F_1,F_2} with F_0=e_phase tensor P_1, F_1=e_phase tensor W, F_2=C_R^2 tensor W"
	Z2AdmissibleSupportLattice   = "A_airlock^Z2={[F_0]_{Z2},[F_1]_{Z2},[F_2]_{Z2}}"
	ClosureOperatorConsequence   = "Cl_airlock^Z2(0)=[F_0]_{Z2}, Cl_airlock^Z2(1)=[F_1]_{Z2}, Cl_airlock^Z2(2)=[F_2]_{Z2}"
	ThetaFromSupportLattice      = "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"
	MuBFromSupportLattice        = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	NextGate                     = "NEXT_PRESSURE_GATE932_TENSORSTRUCTURED_AIRLOCKSUPPORT_ADMISSIBILITY_AUDIT"

	RankF0       = 1
	RankF1       = 4
	RankF2       = 8
	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_AIRLOCK_ADMISSIBLE_SUPPORT_LATTICE_SOURCE_TYPED_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_SUPPORT_LATTICE_SOURCED_NATIVE_ADMISSIBILITY_MISSING"
	FinalTruth     = "AIRLOCK_ADMISSIBLE_SUPPORT_CHAIN_SOURCE_TYPED_BY_PUNCTURE_ROOT_SAME_SOCKET_COMPLETION_AND_FULL_RECTANGLE_SATURATION_BUT_NATIVE_ADMISSIBILITY_THEOREM_MISSING"

	StatusInheritedGate930       = "PASS_GATE930_CLOSURE_OPERATOR_GAP_INHERITED"
	StatusPunctureRooted         = "PASS_ADMISSIBLE_SUPPORT_CHAIN_ROOTED_AT_NEUTRAL_PUNCTURE"
	StatusSameSocketCompletion   = "PASS_SAME_SOCKET_COMPLETION_FORCES_F1"
	StatusTensorFactorIntegrity  = "PASS_ADMISSIBLE_SUPPORTS_ARE_TENSOR_STRUCTURED_COMPLETIONS"
	StatusNoOrphanFragments      = "PASS_ORPHAN_OPPOSITE_SOCKET_FRAGMENTS_EXCLUDED_AT_BRIDGE_LEVEL"
	StatusFullRectangleSaturated = "PASS_FULL_PAIR_ACTIVATION_FORCES_F2"
	StatusMinimalChainSufficient = "PASS_F0_F1_F2_MINIMAL_SUFFICIENT_CHAIN"
	StatusZ2ClassLattice         = "PASS_ADMISSIBLE_SUPPORT_LATTICE_DESCENDS_TO_Z2_CLASS"
	StatusMeasureReconstructed   = "PASS_THETA_AND_MU_B_RECONSTRUCTED_FROM_SOURCED_SUPPORT_LATTICE"
	StatusNativeAdmissibilityGap = "FIREWALL_PRESERVED_NATIVE_ADMISSIBLE_SUPPORT_LATTICE_THEOREM_MISSING"
	StatusAlphaR3StillBlocked    = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportChainRootedAtPuncture             = "CONDITIONAL_SUPPORT_ADMISSIBLE_SUPPORT_CHAIN_IS_ROOTED_AT_NEUTRAL_PUNCTURE"
	SupportF0EqualsPuncture                  = "CONDITIONAL_SUPPORT_F0_EQUALS_PUNCTURE_BASEPOINT"
	SupportActiveTargetsRelativeToF0         = "CONDITIONAL_SUPPORT_ALL_ACTIVE_TARGETS_ARE_RELATIVE_TO_F0"
	SupportSameSocketForcesF1                = "CONDITIONAL_SUPPORT_SAME_SOCKET_COMPLETION_FORCES_F1_EQUALS_E_PHASE_TENSOR_W"
	SupportFirstNonbaseIsExposedFace         = "CONDITIONAL_SUPPORT_FIRST_NONBASE_SUPPORT_IS_EXPOSED_PHASE_FACE_COMPLETION"
	SupportF1OverF0EqualsP3                  = "CONDITIONAL_SUPPORT_F1_OVER_F0_EQUALS_E_PHASE_TENSOR_P3"
	SupportRankThreeFromSameSocket           = "CONDITIONAL_SUPPORT_RANK_THREE_EXPOSED_TARGET_FOLLOWS_FROM_SAME_SOCKET_COMPLETION"
	SupportTensorStructuredNotArbitrary      = "CONDITIONAL_SUPPORT_ADMISSIBLE_SUPPORTS_ARE_TENSOR_STRUCTURED_COMPLETIONS_NOT_ARBITRARY_SUBSPACES"
	SupportPartialOppositeFragmentsExcluded  = "CONDITIONAL_SUPPORT_PARTIAL_OPPOSITE_SOCKET_FRAGMENTS_ARE_NOT_AIRLOCK_ADMISSIBLE"
	SupportSocketWIntegrity                  = "CONDITIONAL_SUPPORT_AIRLOCK_CHAIN_PRESERVES_SOCKET_AND_W_BLOCK_INTEGRITY"
	SupportOrphanOppositeExcluded            = "CONDITIONAL_SUPPORT_ORPHAN_OPPOSITE_SOCKET_FRAGMENTS_ARE_EXCLUDED"
	SupportNoPartialOppositeSocketLevel      = "CONDITIONAL_SUPPORT_ADMISSIBLE_CHAIN_HAS_NO_PARTIAL_OPPOSITE_SOCKET_LEVEL"
	SupportOppositeOnlyAtFullSaturation      = "CONDITIONAL_SUPPORT_OPPOSITE_SOCKET_APPEARS_ONLY_AT_FULL_SATURATION_LEVEL"
	SupportFullPairForcesF2                  = "CONDITIONAL_SUPPORT_FULL_PAIR_ACTIVATION_FORCES_F2_EQUALS_C_R2_TENSOR_W"
	SupportF2FullRightRectangle              = "CONDITIONAL_SUPPORT_F2_IS_FULL_RIGHT_RECTANGLE_SATURATION"
	SupportF2OverF0RankSeven                 = "CONDITIONAL_SUPPORT_F2_OVER_F0_HAS_RANK_SEVEN"
	SupportRankSevenFromSaturation           = "CONDITIONAL_SUPPORT_RANK_SEVEN_FULL_ENCLOSURE_TARGET_FOLLOWS_FROM_SATURATION"
	SupportMinimalSufficientChain            = "CONDITIONAL_SUPPORT_F0_F1_F2_FORM_MINIMAL_SUFFICIENT_AIRLOCK_SUPPORT_CHAIN"
	SupportNoExtraIntermediateLevel          = "CONDITIONAL_SUPPORT_NO_EXTRA_INTERMEDIATE_SUPPORT_LEVEL_IS_REQUIRED"
	SupportAdmissibleLatticeThreeLevel       = "CONDITIONAL_SUPPORT_ADMISSIBLE_LATTICE_COLLAPSES_TO_THREE_LEVEL_CHAIN"
	SupportAdmissibleLatticeDescendsToZ2     = "CONDITIONAL_SUPPORT_ADMISSIBLE_SUPPORT_LATTICE_DESCENDS_TO_Z2_CLASS"
	SupportPhaseFlipExchangesRepresentatives = "CONDITIONAL_SUPPORT_PHASE_FLIP_EXCHANGES_F0_AND_F1_REPRESENTATIVES_WHILE_FIXING_F2"
	SupportRanksZ2Independent                = "CONDITIONAL_SUPPORT_SUPPORT_RANKS_ARE_Z2_REPRESENTATIVE_INDEPENDENT"
	SupportThetaMuBReconstructed             = "CONDITIONAL_SUPPORT_THETA_B_Z2_AND_MU_B_RECONSTRUCTED_FROM_SOURCED_SUPPORT_LATTICE"
	SupportClosureDomainFromLattice          = "CONDITIONAL_SUPPORT_CLOSURE_OPERATOR_DOMAIN_SOURCED_BY_ADMISSIBLE_SUPPORT_LATTICE"

	FailurePunctureRootednessNotNative      = "FAILED_ROUTE_PUNCTURE_ROOTEDNESS_NOT_NATIVE_SUPPORT_LATTICE_THEOREM"
	FailureSameSocketCompletionNotNative    = "FAILED_ROUTE_SAME_SOCKET_COMPLETION_NOT_NATIVE_ADMISSIBILITY_THEOREM"
	FailureTensorFactorIntegrityNotNative   = "FAILED_ROUTE_TENSOR_FACTOR_INTEGRITY_NOT_NATIVE_ADMISSIBILITY_THEOREM"
	FailureNoOrphanRuleNotNative            = "FAILED_ROUTE_NO_ORPHAN_FRAGMENT_RULE_NOT_NATIVE_THEOREM"
	FailureFullRectangleSaturationNotNative = "FAILED_ROUTE_FULL_RIGHT_RECTANGLE_SATURATION_NOT_NATIVE_ADMISSIBILITY_THEOREM"
	FailureMinimalSufficientChainNotNative  = "FAILED_ROUTE_MINIMAL_SUFFICIENT_SUPPORT_CHAIN_NOT_NATIVE_UNIQUENESS_THEOREM"
	FailureZ2SupportLatticeNotNative        = "FAILED_ROUTE_Z2_SUPPORT_LATTICE_NOT_NATIVE_GLOBAL_PHASE_THEOREM"
	FailureNoNativeAdmissibleSupportLattice = "FAILED_ROUTE_NO_NATIVE_ADMISSIBLE_AIRLOCK_SUPPORT_LATTICE_THEOREM"
	FailureAlphaBridgeCandidateNotNative    = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                      = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked        = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type PunctureRootAudit struct {
	RootedAtPuncture bool
	F0Rank           int
	RelativeTargets  bool
	NativeTheorem    bool
	Supports         []string
	Failures         []string
}

type SameSocketAudit struct {
	F1Rank              int
	F1OverF0Rank        int
	CompletesSameSocket bool
	F1EqualsPhaseW      bool
	NativeTheorem       bool
	Supports            []string
	Failures            []string
}

type TensorIntegrityAudit struct {
	StructuredCompletions bool
	ArbitrarySubspaces    bool
	PreservesSocketW      bool
	NativeTheorem         bool
	Supports              []string
	Failures              []string
}

type NoOrphanAudit struct {
	ExcludesOppositeLeptonSingleton bool
	ExcludesOppositeColorFragment   bool
	OppositeOnlyAtFullSaturation    bool
	NativeTheorem                   bool
	Supports                        []string
	Failures                        []string
}

type SaturationAudit struct {
	F2Rank           int
	F2OverF0Rank     int
	FullPairForcesF2 bool
	FullRectangle    bool
	RejectsStopAtF1  bool
	NativeTheorem    bool
	Supports         []string
	Failures         []string
}

type MinimalChainAudit struct {
	Chain                 string
	MinimalSufficient     bool
	NoExtraIntermediate   bool
	ThreeLevelCollapse    bool
	NativeUniquenessProof bool
	Supports              []string
	Failures              []string
}

type Z2LatticeAudit struct {
	DescendsToZ2Class       bool
	PhaseFlipExchangesF0F1  bool
	PhaseFlipFixesF2        bool
	RanksRepresentativeFree bool
	NativePhaseTheorem      bool
	Supports                []string
	Failures                []string
}

type MeasureConsequenceAudit struct {
	ThetaRecovered       bool
	ClosureDomainSourced bool
	MuBReconstructed     bool
	ThetaOneRank         int
	ThetaTwoRank         int
	H10Rank              int
	H72Rank              int
	NativeAlpha          bool
	Supports             []string
	Failures             []string
}

type Gate931Audit struct {
	ID              string
	Inherited       string
	Classification  string
	ShortStatus     string
	Truth           string
	PunctureRoot    PunctureRootAudit
	SameSocket      SameSocketAudit
	TensorIntegrity TensorIntegrityAudit
	NoOrphan        NoOrphanAudit
	Saturation      SaturationAudit
	MinimalChain    MinimalChainAudit
	Z2Lattice       Z2LatticeAudit
	Measure         MeasureConsequenceAudit
	Firewalls       []string
	Final           string
}

func BuildDefault() (Gate931Audit, error) {
	a := Gate931Audit{
		ID:              AuditID,
		Inherited:       Gate930ShortStatus,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		Truth:           FinalTruth,
		PunctureRoot:    PunctureRootAudit{RootedAtPuncture: true, F0Rank: RankF0, RelativeTargets: true, NativeTheorem: false, Supports: []string{SupportChainRootedAtPuncture, SupportF0EqualsPuncture, SupportActiveTargetsRelativeToF0}, Failures: []string{FailurePunctureRootednessNotNative}},
		SameSocket:      SameSocketAudit{F1Rank: RankF1, F1OverF0Rank: RankF1OverF0, CompletesSameSocket: true, F1EqualsPhaseW: true, NativeTheorem: false, Supports: []string{SupportSameSocketForcesF1, SupportFirstNonbaseIsExposedFace, SupportF1OverF0EqualsP3, SupportRankThreeFromSameSocket}, Failures: []string{FailureSameSocketCompletionNotNative}},
		TensorIntegrity: TensorIntegrityAudit{StructuredCompletions: true, ArbitrarySubspaces: false, PreservesSocketW: true, NativeTheorem: false, Supports: []string{SupportTensorStructuredNotArbitrary, SupportPartialOppositeFragmentsExcluded, SupportSocketWIntegrity}, Failures: []string{FailureTensorFactorIntegrityNotNative}},
		NoOrphan:        NoOrphanAudit{ExcludesOppositeLeptonSingleton: true, ExcludesOppositeColorFragment: true, OppositeOnlyAtFullSaturation: true, NativeTheorem: false, Supports: []string{SupportOrphanOppositeExcluded, SupportNoPartialOppositeSocketLevel, SupportOppositeOnlyAtFullSaturation}, Failures: []string{FailureNoOrphanRuleNotNative}},
		Saturation:      SaturationAudit{F2Rank: RankF2, F2OverF0Rank: RankF2OverF0, FullPairForcesF2: true, FullRectangle: true, RejectsStopAtF1: true, NativeTheorem: false, Supports: []string{SupportFullPairForcesF2, SupportF2FullRightRectangle, SupportF2OverF0RankSeven, SupportRankSevenFromSaturation}, Failures: []string{FailureFullRectangleSaturationNotNative}},
		MinimalChain:    MinimalChainAudit{Chain: AdmissibleSupportChain, MinimalSufficient: true, NoExtraIntermediate: true, ThreeLevelCollapse: true, NativeUniquenessProof: false, Supports: []string{SupportMinimalSufficientChain, SupportNoExtraIntermediateLevel, SupportAdmissibleLatticeThreeLevel}, Failures: []string{FailureMinimalSufficientChainNotNative}},
		Z2Lattice:       Z2LatticeAudit{DescendsToZ2Class: true, PhaseFlipExchangesF0F1: true, PhaseFlipFixesF2: true, RanksRepresentativeFree: true, NativePhaseTheorem: false, Supports: []string{SupportAdmissibleLatticeDescendsToZ2, SupportPhaseFlipExchangesRepresentatives, SupportRanksZ2Independent}, Failures: []string{FailureZ2SupportLatticeNotNative}},
		Measure:         MeasureConsequenceAudit{ThetaRecovered: true, ClosureDomainSourced: true, MuBReconstructed: true, ThetaOneRank: RankF1OverF0, ThetaTwoRank: RankF2OverF0, H10Rank: RankH10, H72Rank: RankH72, NativeAlpha: false, Supports: []string{SupportThetaMuBReconstructed, SupportClosureDomainFromLattice}, Failures: []string{FailureNoNativeAdmissibleSupportLattice, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}},
		Firewalls:       []string{FailurePunctureRootednessNotNative, FailureSameSocketCompletionNotNative, FailureTensorFactorIntegrityNotNative, FailureNoOrphanRuleNotNative, FailureFullRectangleSaturationNotNative, FailureMinimalSufficientChainNotNative, FailureZ2SupportLatticeNotNative, FailureNoNativeAdmissibleSupportLattice, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator},
		Final:           "Gate 931 source-types {F_0,F_1,F_2} by puncture root, same-socket completion, tensor-factor integrity, no-orphan fragments, full right-rectangle saturation, minimal sufficiency, and Z2 class descent; the native admissibility theorem remains missing.",
	}
	if err := a.Validate(); err != nil {
		return Gate931Audit{}, err
	}
	return a, nil
}

func (a Gate931Audit) Validate() error {
	if a.ID != AuditID || a.Inherited != Gate930ShortStatus {
		return fmt.Errorf("bad Gate931 audit identity")
	}
	if !a.PunctureRoot.RootedAtPuncture || a.PunctureRoot.F0Rank != RankF0 || a.PunctureRoot.NativeTheorem {
		return fmt.Errorf("bad puncture root audit")
	}
	if !a.SameSocket.CompletesSameSocket || !a.SameSocket.F1EqualsPhaseW || a.SameSocket.F1OverF0Rank != RankF1OverF0 || a.SameSocket.NativeTheorem {
		return fmt.Errorf("bad same-socket audit")
	}
	if !a.TensorIntegrity.StructuredCompletions || a.TensorIntegrity.ArbitrarySubspaces || !a.TensorIntegrity.PreservesSocketW || a.TensorIntegrity.NativeTheorem {
		return fmt.Errorf("bad tensor integrity audit")
	}
	if !a.NoOrphan.ExcludesOppositeLeptonSingleton || !a.NoOrphan.ExcludesOppositeColorFragment || !a.NoOrphan.OppositeOnlyAtFullSaturation || a.NoOrphan.NativeTheorem {
		return fmt.Errorf("bad no-orphan audit")
	}
	if !a.Saturation.FullPairForcesF2 || !a.Saturation.FullRectangle || !a.Saturation.RejectsStopAtF1 || a.Saturation.F2OverF0Rank != RankF2OverF0 || a.Saturation.NativeTheorem {
		return fmt.Errorf("bad saturation audit")
	}
	if !a.MinimalChain.MinimalSufficient || !a.MinimalChain.NoExtraIntermediate || !a.MinimalChain.ThreeLevelCollapse || a.MinimalChain.NativeUniquenessProof {
		return fmt.Errorf("bad minimal chain audit")
	}
	if !a.Z2Lattice.DescendsToZ2Class || !a.Z2Lattice.PhaseFlipExchangesF0F1 || !a.Z2Lattice.PhaseFlipFixesF2 || !a.Z2Lattice.RanksRepresentativeFree || a.Z2Lattice.NativePhaseTheorem {
		return fmt.Errorf("bad z2 lattice audit")
	}
	if !a.Measure.ThetaRecovered || !a.Measure.ClosureDomainSourced || !a.Measure.MuBReconstructed || a.Measure.NativeAlpha || a.Measure.ThetaOneRank != RankF1OverF0 || a.Measure.ThetaTwoRank != RankF2OverF0 || a.Measure.H10Rank != RankH10 || a.Measure.H72Rank != RankH72 {
		return fmt.Errorf("bad measure consequence audit")
	}
	if !firewallsOK(a.Firewalls) {
		return fmt.Errorf("missing preserved firewalls")
	}
	return nil
}

func Statuses() []string {
	return []string{StatusInheritedGate930, StatusPunctureRooted, StatusSameSocketCompletion, StatusTensorFactorIntegrity, StatusNoOrphanFragments, StatusFullRectangleSaturated, StatusMinimalChainSufficient, StatusZ2ClassLattice, StatusMeasureReconstructed, StatusNativeAdmissibilityGap, StatusAlphaR3StillBlocked}
}

func Supports() []string {
	return []string{SupportChainRootedAtPuncture, SupportF0EqualsPuncture, SupportActiveTargetsRelativeToF0, SupportSameSocketForcesF1, SupportFirstNonbaseIsExposedFace, SupportF1OverF0EqualsP3, SupportRankThreeFromSameSocket, SupportTensorStructuredNotArbitrary, SupportPartialOppositeFragmentsExcluded, SupportSocketWIntegrity, SupportOrphanOppositeExcluded, SupportNoPartialOppositeSocketLevel, SupportOppositeOnlyAtFullSaturation, SupportFullPairForcesF2, SupportF2FullRightRectangle, SupportF2OverF0RankSeven, SupportRankSevenFromSaturation, SupportMinimalSufficientChain, SupportNoExtraIntermediateLevel, SupportAdmissibleLatticeThreeLevel, SupportAdmissibleLatticeDescendsToZ2, SupportPhaseFlipExchangesRepresentatives, SupportRanksZ2Independent, SupportThetaMuBReconstructed, SupportClosureDomainFromLattice}
}

func Failures() []string {
	return []string{FailurePunctureRootednessNotNative, FailureSameSocketCompletionNotNative, FailureTensorFactorIntegrityNotNative, FailureNoOrphanRuleNotNative, FailureFullRectangleSaturationNotNative, FailureMinimalSufficientChainNotNative, FailureZ2SupportLatticeNotNative, FailureNoNativeAdmissibleSupportLattice, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatPunctureRoot(a PunctureRootAudit) string {
	return fmt.Sprintf("rooted=%t F0_rank=%d relative_targets=%t native=%t", a.RootedAtPuncture, a.F0Rank, a.RelativeTargets, a.NativeTheorem)
}
func FormatSameSocket(a SameSocketAudit) string {
	return fmt.Sprintf("same_socket=%t F1=e_phase tensor W:%t F1_rank=%d F1/F0_rank=%d native=%t", a.CompletesSameSocket, a.F1EqualsPhaseW, a.F1Rank, a.F1OverF0Rank, a.NativeTheorem)
}
func FormatTensorIntegrity(a TensorIntegrityAudit) string {
	return fmt.Sprintf("structured=%t arbitrary_subspaces=%t socket_W_integrity=%t native=%t", a.StructuredCompletions, a.ArbitrarySubspaces, a.PreservesSocketW, a.NativeTheorem)
}
func FormatNoOrphan(a NoOrphanAudit) string {
	return fmt.Sprintf("exclude_opposite_lepton=%t exclude_opposite_color=%t opposite_only_full=%t native=%t", a.ExcludesOppositeLeptonSingleton, a.ExcludesOppositeColorFragment, a.OppositeOnlyAtFullSaturation, a.NativeTheorem)
}
func FormatSaturation(a SaturationAudit) string {
	return fmt.Sprintf("full_pair_forces_F2=%t full_rectangle=%t rejects_F1=%t F2_rank=%d F2/F0_rank=%d native=%t", a.FullPairForcesF2, a.FullRectangle, a.RejectsStopAtF1, a.F2Rank, a.F2OverF0Rank, a.NativeTheorem)
}
func FormatMinimalChain(a MinimalChainAudit) string {
	return fmt.Sprintf("chain=%s minimal=%t no_extra=%t three_level=%t native_unique=%t", a.Chain, a.MinimalSufficient, a.NoExtraIntermediate, a.ThreeLevelCollapse, a.NativeUniquenessProof)
}
func FormatZ2Lattice(a Z2LatticeAudit) string {
	return fmt.Sprintf("descends=%t flip_F0F1=%t fixes_F2=%t ranks_free=%t native_phase=%t", a.DescendsToZ2Class, a.PhaseFlipExchangesF0F1, a.PhaseFlipFixesF2, a.RanksRepresentativeFree, a.NativePhaseTheorem)
}
func FormatMeasure(a MeasureConsequenceAudit) string {
	return fmt.Sprintf("theta=%t closure_domain=%t muB=%t ranks=(%d,%d) chambers=(%d,%d) native_alpha=%t", a.ThetaRecovered, a.ClosureDomainSourced, a.MuBReconstructed, a.ThetaOneRank, a.ThetaTwoRank, a.H10Rank, a.H72Rank, a.NativeAlpha)
}
func FormatFirewalls(f []string) string { return "firewalls=" + strings.Join(f, ",") }

func containsAll(have []string, want []string) bool {
	m := map[string]bool{}
	for _, v := range have {
		m[v] = true
	}
	for _, v := range want {
		if !m[v] {
			return false
		}
	}
	return true
}

func firewallsOK(f []string) bool {
	return containsAll(f, []string{FailureNoNativeAdmissibleSupportLattice, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3})
}
