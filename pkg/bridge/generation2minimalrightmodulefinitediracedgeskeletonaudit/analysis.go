// Package generation2minimalrightmodulefinitediracedgeskeletonaudit implements
// Gate 844: Minimal RightModule Finite-Dirac Edge-Skeleton Audit.
//
// Gate 844 follows Gate 843's minimal right-neutral absence seal. Gate 843
// located the R2++ aggregate support, at seal level, on the active right
// lepto-color module H_R^min = (C_R^2 tensor W) minus (e_+ tensor P_1), with
// rank seven. Gate 844 asks whether this minimal active right module can serve
// as the support-only domain of a symbolic finite Dirac edge skeleton into the
// left lepto-color doublet H_L = C_L^2 tensor W, while the puncture remains
// absent from the domain. This is an edge-support seal audit only. It does not
// certify an explicit D_F matrix, the first-order condition, Yukawa magnitudes,
// particle assignments, alpha_B, R3, or R4.
package generation2minimalrightmodulefinitediracedgeskeletonaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE844-MINIMAL-RIGHT-MODULE-FINITE-DIRAC-EDGE-SKELETON-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	LeftSocketDim  = 2
	RightPairDim   = 2

	EPlusP3Rank  = ColorBlockDim
	EPlusP1Rank  = LeptonBlockDim
	EMinusP3Rank = ColorBlockDim
	EMinusP1Rank = LeptonBlockDim

	HRMinRank       = EPlusP3Rank + EMinusP3Rank + EMinusP1Rank
	PunctureRank    = EPlusP1Rank
	RightFullRank   = RightPairDim * WDim
	HLRank          = LeftSocketDim * WDim
	HLColorRank     = LeftSocketDim * ColorBlockDim
	HLLeptonRank    = LeftSocketDim * LeptonBlockDim
	FullParticleDim = 16
	FullFiniteDim   = 32

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate843Inherited                  = "PASS_GATE843_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL_INHERITED"
	StatusHRMinDomainAudited                = "PASS_H_R_MIN_ACTIVE_RIGHT_DOMAIN_AUDITED"
	StatusLeftLeptoColorTargetAudited       = "PASS_LEFT_LEPTOCOLOR_DOUBLE_TARGET_AUDITED"
	StatusSymbolicDFSupportConstructed      = "PASS_SYMBOLIC_D_F_EDGE_SUPPORT_CONSTRUCTED_AT_SEAL_LEVEL"
	StatusColorLeptonSupportPreserved       = "PASS_COLOR_LEPTON_SUPPORT_PRESERVATION_AUDITED"
	StatusPunctureAbsenceCompatible         = "PASS_PUNCTURE_ABSENCE_COMPATIBLE_WITH_SYMBOLIC_EDGE_SUPPORT"
	StatusDFEdgeSupportNotMagnitude         = "PASS_D_F_EDGE_SUPPORT_CLASSIFIED_AS_COUPLING_GRAPH_NOT_MAGNITUDE"
	StatusFirstOrderBimoduleFirewallAudited = "PASS_FIRST_ORDER_AND_BIMODULE_FIREWALL_AUDITED"
	StatusAlphaStillSealed                  = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_GATE844"
	StatusOfficialLedgersFrozen             = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusSealedShadow            = "PASS_R2_PLUS_PLUS_SEALED_FINITE_BODY_EDGE_SHADOW_NOT_R3_OR_R4"
	StatusNoObservedDataUsed                = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate844                   = "FIREWALL_PRESERVED_GATE844_MINIMAL_RIGHT_MODULE_D_F_EDGE_SUPPORT"

	SupportHRMinAsRightEdgeDomain       = "CONDITIONAL_SUPPORT_H_R_MIN_IS_ACTIVE_RIGHT_EDGE_DOMAIN_AT_SEAL_LEVEL"
	SupportLeftTargetHL                 = "CONDITIONAL_SUPPORT_H_L_EQUALS_C_L2_TENSOR_W_HAS_RANK_EIGHT"
	SupportSymbolicDFEdgeSkeleton       = "CONDITIONAL_SUPPORT_MINIMAL_RIGHT_MODULE_SUPPORTS_SYMBOLIC_D_F_EDGE_SKELETON"
	SupportPunctureAbsenceCompatible    = "CONDITIONAL_SUPPORT_PUNCTURE_ABSENCE_COMPATIBLE_WITH_EDGE_SUPPORT"
	SupportColorSupportPreserved        = "CONDITIONAL_SUPPORT_COLOR_SUPPORT_EDGES_P3_TO_C_L2_TENSOR_P3"
	SupportLeptonSupportPreserved       = "CONDITIONAL_SUPPORT_LEPTON_SUPPORT_EDGES_P1_TO_C_L2_TENSOR_P1"
	SupportActiveSevenAsDomain          = "CONDITIONAL_SUPPORT_ACTIVE_7_IS_RIGHT_EDGE_DOMAIN_AT_SEAL_LEVEL"
	SupportDFSuppCouplingGraphOnly      = "CONDITIONAL_SUPPORT_D_F_SUPP_IS_COUPLING_GRAPH_ONLY"
	SupportPunctureAsAbsentOnly         = "CONDITIONAL_SUPPORT_E_PLUS_TENSOR_P1_REMAINS_ABSENT_PUNCTURE_CANDIDATE_ONLY"
	SupportFiniteBodyShadowStrengthened = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_FINITE_BODY_SHADOW_STRENGTHENED_BY_EDGE_SUPPORT_SEAL"

	FailureDFSupportSealNotNative         = "FAILED_ROUTE_SYMBOLIC_D_F_EDGE_SKELETON_IS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoExplicitDFMatrix             = "FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_CERTIFIED"
	FailureNoFirstOrderProof              = "FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof       = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailurePunctureAbsenceNotFromNullEdge = "FAILED_ROUTE_PUNCTURE_ABSENCE_NOT_DERIVED_FROM_D_F_NULL_EDGE"
	FailureNoNativeNullEdgeTheorem        = "FAILED_ROUTE_NO_NATIVE_NULL_EDGE_THEOREM_FOR_E_PLUS_TENSOR_P1"
	FailureNoNativeMinimalAbsenceTheorem  = "FAILED_ROUTE_NO_NATIVE_MINIMAL_RIGHT_NEUTRAL_ABSENCE_THEOREM"
	FailureNoPhysicalParticleAssignment   = "FAILED_ROUTE_NEUTRAL_SINGLETON_NOT_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoRightNeutrinoTheorem         = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoFullRhoFActionLedger         = "FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoGammaFJFPackage              = "FAILED_ROUTE_NO_COMPLETE_GAMMA_F_J_F_REFINEMENT_PACKAGE_CERTIFIED"
	FailureDFEdgeSupportNotYukawa         = "FAILED_ROUTE_D_F_EDGE_SUPPORT_NOT_YUKAWA_MAGNITUDE"
	FailureNoNumericalYukawaValues        = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoAlphaDerivation              = "FAILED_ROUTE_MINIMAL_D_F_EDGE_SUPPORT_DOES_NOT_DERIVE_ALPHA_B"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout        = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoAggregateCompressionNative   = "FAILED_ROUTE_NO_NATIVE_AGGREGATE_TRACE_COMPRESSION_THEOREM"
	FailureNoNEffUpdate                   = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit            = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoThreeGenerationTheorem       = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                          = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                          = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlus, R3, R4              bool
	AlphaNative                     bool
}

type DomainCell struct {
	Name, Expression, SafeLabel string
	Rank                        int
	BMinusLTrace                float64
	ActiveDomain, Puncture      bool
}

type RightDomain struct {
	Cells                                       []DomainCell
	Expression, PunctureExpression              string
	FullRightRank, ActiveRank, PunctureRank     int
	RankPattern, ActivePattern                  string
	InheritedFromGate843, MinimalAbsenceSeal    bool
	ActiveIsFullMinusPuncture                   bool
	BMinusLActive, BMinusLPuncture, BMinusLFull float64
	Supports, Failures                          []string
}

type LeftTarget struct {
	Expression                    string
	Rank, ColorRank, LeptonRank   int
	Complete, LeptoColorPreserved bool
	Supports, Failures            []string
}

type Edge struct {
	Name, DomainExpression, TargetExpression, SupportType string
	DomainRank, TargetRank                                int
	ColorLeptonPreserving, SymbolicOnly                   bool
}

type EdgeSkeleton struct {
	Expression                                                      string
	Edges                                                           []Edge
	DomainRank, TargetRank, PunctureRank                            int
	ConstructedAtSealLevel, NativeDFMatrixCertified                 bool
	ExplicitDFMatrixCertified, FirstOrderConditionCertified         bool
	BimoduleCommutantCertified, EdgeSupportOnly                     bool
	YukawaMagnitudes, NumericalValues                               bool
	PunctureInDomain, PunctureNullEdgeCertified                     bool
	PunctureAbsenceCompatible, PunctureAbsenceDerivedFromDFNullEdge bool
	Supports, Failures                                              []string
}

type AggregateShadow struct {
	Expression                                          string
	FiniteBodyLocationAtSealLevel, EdgeSupportSealLevel bool
	NativeCompressionTheorem, AlphaDerived              bool
	TraceMagnitudeReadout, R3, R4                       bool
	Supports, Failures                                  []string
}

type Firewalls struct {
	Enforced                                                      bool
	DFSupportSealNotNative, NoExplicitDFMatrix, NoFirstOrderProof bool
	NoBimoduleCommutantProof, PunctureAbsenceNotFromNullEdge      bool
	NoNativeNullEdgeTheorem, NoNativeMinimalAbsenceTheorem        bool
	NoPhysicalParticleAssignment, NoRightNeutrinoTheorem          bool
	NoFullRhoFActionLedger, NoGammaFJFPackage                     bool
	DFEdgeSupportNotYukawa, NoNumericalYukawaValues               bool
	NoAlphaDerivation, AlphaStillSealed, NoTraceMagnitudeReadout  bool
	NoAggregateCompressionNative, NoNEffUpdate, NoCYukawaUpdate   bool
	NoObservedYukawaFit, NoThreeGenerationTheorem, NotR3, NotR4   bool
	Verdict                                                       string
}

type Impact struct {
	Classification                                     string
	SymbolicEdgeSupportSeal, PunctureAbsenceCompatible bool
	PunctureNullEdgeNative, NativeDFMatrixCertified    bool
	FirstOrderCertified, NativeCompressionStillMissing bool
	AlphaStillSealed, MagnitudesStillMissing           bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs   bool
	CanPromoteToR3, CanPromoteToR4                     bool
}

type Audit struct {
	Ledger    Ledger
	Domain    RightDomain
	Target    LeftTarget
	Edges     EdgeSkeleton
	Shadow    AggregateShadow
	Firewalls Firewalls
	Impact    Impact
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	domain := buildDomain()
	target := buildTarget()
	edges := buildEdges()
	shadow := buildShadow()
	firewalls := buildFirewalls()
	impact := Impact{
		Classification:                "R2++ sealed finite-body edge-support shadow: H_R^min supports a symbolic D_F coupling graph into H_L, but no native D_F matrix, first-order proof, magnitude readout, R3, or R4 theorem is certified",
		SymbolicEdgeSupportSeal:       true,
		PunctureAbsenceCompatible:     true,
		PunctureNullEdgeNative:        false,
		NativeDFMatrixCertified:       false,
		FirstOrderCertified:           false,
		NativeCompressionStillMissing: true,
		AlphaStillSealed:              true,
		MagnitudesStillMissing:        true,
		CanUpdateNEff:                 false,
		CanUpdateCYukawa:              false,
		CanUpdateCHiggs:               false,
		CanPromoteToR3:                false,
		CanPromoteToR4:                false,
	}
	a := Audit{
		Ledger:    Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaNative: false},
		Domain:    domain,
		Target:    target,
		Edges:     edges,
		Shadow:    shadow,
		Firewalls: firewalls,
		Impact:    impact,
		Truth:     "Gate 844 builds a support-only symbolic D_F edge skeleton from H_R^min to H_L. The minimal right module remains a seal-level domain; the puncture is absent from the domain, but its absence is not derived from a native D_F null-edge theorem.",
		Final:     "The R2++ aggregate shadow is strengthened: the rank-seven support can act as a sealed right-edge domain into the left lepto-color doublet. This is a coupling graph only, not a Yukawa magnitude theorem, not an alpha source, not R3, and not R4.",
	}
	return a, validate(a)
}

func buildDomain() RightDomain {
	cells := []DomainCell{
		{Name: "character-plus color triplet", Expression: "e_+ tensor P_3", SafeLabel: "dominant character-color triplet edge-domain cell", Rank: EPlusP3Rank, BMinusLTrace: 1, ActiveDomain: true},
		{Name: "character-plus lepton singleton", Expression: "e_+ tensor P_1", SafeLabel: "neutral right-lepton puncture / absent domain cell", Rank: EPlusP1Rank, BMinusLTrace: -1, Puncture: true},
		{Name: "character-minus color triplet", Expression: "e_- tensor P_3", SafeLabel: "rest character-color triplet edge-domain cell", Rank: EMinusP3Rank, BMinusLTrace: 1, ActiveDomain: true},
		{Name: "character-minus lepton singleton", Expression: "e_- tensor P_1", SafeLabel: "rest character-lepton singleton edge-domain cell", Rank: EMinusP1Rank, BMinusLTrace: -1, ActiveDomain: true},
	}
	active := 0
	puncture := 0
	bActive := 0.0
	bPuncture := 0.0
	for _, c := range cells {
		if c.ActiveDomain {
			active += c.Rank
			bActive += c.BMinusLTrace
		}
		if c.Puncture {
			puncture += c.Rank
			bPuncture += c.BMinusLTrace
		}
	}
	return RightDomain{
		Cells:                     cells,
		Expression:                "H_R^min = (e_+ tensor P_3) plus (e_- tensor P_3) plus (e_- tensor P_1)",
		PunctureExpression:        "e_+ tensor P_1",
		FullRightRank:             RightFullRank,
		ActiveRank:                active,
		PunctureRank:              puncture,
		RankPattern:               "8=3+1+3+1",
		ActivePattern:             "7=3+3+1=8-1",
		InheritedFromGate843:      true,
		MinimalAbsenceSeal:        true,
		ActiveIsFullMinusPuncture: active == RightFullRank-puncture,
		BMinusLActive:             bActive,
		BMinusLPuncture:           bPuncture,
		BMinusLFull:               bActive + bPuncture,
		Supports:                  []string{SupportHRMinAsRightEdgeDomain, SupportActiveSevenAsDomain, SupportPunctureAbsenceCompatible},
		Failures:                  []string{FailureDFSupportSealNotNative, FailurePunctureAbsenceNotFromNullEdge, FailureNoNativeMinimalAbsenceTheorem},
	}
}

func buildTarget() LeftTarget {
	return LeftTarget{
		Expression:          "H_L = C_L^2 tensor W = (C_L^2 tensor P_3) plus (C_L^2 tensor P_1)",
		Rank:                HLRank,
		ColorRank:           HLColorRank,
		LeptonRank:          HLLeptonRank,
		Complete:            true,
		LeptoColorPreserved: true,
		Supports:            []string{SupportLeftTargetHL, SupportColorSupportPreserved, SupportLeptonSupportPreserved},
		Failures:            []string{FailureNoFullRhoFActionLedger, FailureNoGammaFJFPackage},
	}
}

func buildEdges() EdgeSkeleton {
	edges := []Edge{
		{Name: "character-plus color edge", DomainExpression: "e_+ tensor P_3", TargetExpression: "C_L^2 tensor P_3", SupportType: "color-preserving symbolic edge", DomainRank: EPlusP3Rank, TargetRank: HLColorRank, ColorLeptonPreserving: true, SymbolicOnly: true},
		{Name: "character-minus color edge", DomainExpression: "e_- tensor P_3", TargetExpression: "C_L^2 tensor P_3", SupportType: "color-preserving symbolic edge", DomainRank: EMinusP3Rank, TargetRank: HLColorRank, ColorLeptonPreserving: true, SymbolicOnly: true},
		{Name: "character-minus lepton edge", DomainExpression: "e_- tensor P_1", TargetExpression: "C_L^2 tensor P_1", SupportType: "lepton-preserving symbolic edge", DomainRank: EMinusP1Rank, TargetRank: HLLeptonRank, ColorLeptonPreserving: true, SymbolicOnly: true},
	}
	return EdgeSkeleton{
		Expression:                           "D_F^supp: H_R^min -> H_L",
		Edges:                                edges,
		DomainRank:                           HRMinRank,
		TargetRank:                           HLRank,
		PunctureRank:                         PunctureRank,
		ConstructedAtSealLevel:               true,
		NativeDFMatrixCertified:              false,
		ExplicitDFMatrixCertified:            false,
		FirstOrderConditionCertified:         false,
		BimoduleCommutantCertified:           false,
		EdgeSupportOnly:                      true,
		YukawaMagnitudes:                     false,
		NumericalValues:                      false,
		PunctureInDomain:                     false,
		PunctureNullEdgeCertified:            false,
		PunctureAbsenceCompatible:            true,
		PunctureAbsenceDerivedFromDFNullEdge: false,
		Supports:                             []string{SupportSymbolicDFEdgeSkeleton, SupportPunctureAbsenceCompatible, SupportPunctureAsAbsentOnly, SupportColorSupportPreserved, SupportLeptonSupportPreserved, SupportDFSuppCouplingGraphOnly},
		Failures:                             []string{FailureDFSupportSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailurePunctureAbsenceNotFromNullEdge, FailureNoNativeNullEdgeTheorem, FailureDFEdgeSupportNotYukawa, FailureNoNumericalYukawaValues},
	}
}

func buildShadow() AggregateShadow {
	return AggregateShadow{
		Expression:                    "H_total/T = I_{e_+ tensor P_3} plus [alpha_B P_3 - 3 alpha_B^2(B-L)]_{e_- tensor W}, with H_R^min as D_F^supp domain",
		FiniteBodyLocationAtSealLevel: true,
		EdgeSupportSealLevel:          true,
		NativeCompressionTheorem:      false,
		AlphaDerived:                  false,
		TraceMagnitudeReadout:         false,
		R3:                            false,
		R4:                            false,
		Supports:                      []string{SupportFiniteBodyShadowStrengthened, SupportHRMinAsRightEdgeDomain, SupportSymbolicDFEdgeSkeleton},
		Failures:                      []string{FailureNoAggregateCompressionNative, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		Enforced:                       true,
		DFSupportSealNotNative:         true,
		NoExplicitDFMatrix:             true,
		NoFirstOrderProof:              true,
		NoBimoduleCommutantProof:       true,
		PunctureAbsenceNotFromNullEdge: true,
		NoNativeNullEdgeTheorem:        true,
		NoNativeMinimalAbsenceTheorem:  true,
		NoPhysicalParticleAssignment:   true,
		NoRightNeutrinoTheorem:         true,
		NoFullRhoFActionLedger:         true,
		NoGammaFJFPackage:              true,
		DFEdgeSupportNotYukawa:         true,
		NoNumericalYukawaValues:        true,
		NoAlphaDerivation:              true,
		AlphaStillSealed:               true,
		NoTraceMagnitudeReadout:        true,
		NoAggregateCompressionNative:   true,
		NoNEffUpdate:                   true,
		NoCYukawaUpdate:                true,
		NoObservedYukawaFit:            true,
		NoThreeGenerationTheorem:       true,
		NotR3:                          true,
		NotR4:                          true,
		Verdict:                        StatusFirewallGate844,
	}
}

func validate(a Audit) error {
	if !a.Domain.InheritedFromGate843 || !a.Domain.MinimalAbsenceSeal || a.Domain.ActiveRank != HRMinRank || a.Domain.PunctureRank != PunctureRank || !a.Domain.ActiveIsFullMinusPuncture {
		return fmt.Errorf("right domain invalid: %s", FormatDomain(a.Domain))
	}
	if math.Abs(a.Domain.BMinusLActive-1) > 1e-12 || math.Abs(a.Domain.BMinusLPuncture+1) > 1e-12 || math.Abs(a.Domain.BMinusLFull) > 1e-12 {
		return fmt.Errorf("B-L compensation mismatch: %s", FormatDomain(a.Domain))
	}
	if !a.Target.Complete || !a.Target.LeptoColorPreserved || a.Target.Rank != HLRank || a.Target.ColorRank != HLColorRank || a.Target.LeptonRank != HLLeptonRank {
		return fmt.Errorf("left target invalid: %s", FormatTarget(a.Target))
	}
	if !a.Edges.ConstructedAtSealLevel || !a.Edges.EdgeSupportOnly || a.Edges.NativeDFMatrixCertified || a.Edges.ExplicitDFMatrixCertified || a.Edges.FirstOrderConditionCertified || a.Edges.BimoduleCommutantCertified || a.Edges.YukawaMagnitudes || a.Edges.NumericalValues {
		return fmt.Errorf("edge skeleton over/under-certified: %s", FormatEdges(a.Edges))
	}
	if a.Edges.PunctureInDomain || a.Edges.PunctureNullEdgeCertified || !a.Edges.PunctureAbsenceCompatible || a.Edges.PunctureAbsenceDerivedFromDFNullEdge {
		return fmt.Errorf("puncture edge status invalid: %s", FormatEdges(a.Edges))
	}
	if len(a.Edges.Edges) != 3 {
		return fmt.Errorf("wrong number of symbolic edges: %s", FormatEdges(a.Edges))
	}
	for _, e := range a.Edges.Edges {
		if !e.ColorLeptonPreserving || !e.SymbolicOnly || e.DomainRank <= 0 || e.TargetRank <= 0 {
			return fmt.Errorf("edge invalid: %+v", e)
		}
	}
	if !a.Shadow.FiniteBodyLocationAtSealLevel || !a.Shadow.EdgeSupportSealLevel || a.Shadow.NativeCompressionTheorem || a.Shadow.AlphaDerived || a.Shadow.TraceMagnitudeReadout || a.Shadow.R3 || a.Shadow.R4 {
		return fmt.Errorf("shadow invalid: %s", FormatShadow(a.Shadow))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate844 || !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		return fmt.Errorf("firewall or ledger invalid")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate843Inherited, StatusHRMinDomainAudited, StatusLeftLeptoColorTargetAudited, StatusSymbolicDFSupportConstructed, StatusColorLeptonSupportPreserved, StatusPunctureAbsenceCompatible, StatusDFEdgeSupportNotMagnitude, StatusFirstOrderBimoduleFirewallAudited, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusSealedShadow, StatusNoObservedDataUsed, StatusFirewallGate844,
		SupportHRMinAsRightEdgeDomain, SupportLeftTargetHL, SupportSymbolicDFEdgeSkeleton, SupportPunctureAbsenceCompatible, SupportColorSupportPreserved, SupportLeptonSupportPreserved, SupportActiveSevenAsDomain, SupportDFSuppCouplingGraphOnly, SupportPunctureAsAbsentOnly, SupportFiniteBodyShadowStrengthened,
		FailureDFSupportSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailurePunctureAbsenceNotFromNullEdge, FailureNoNativeNullEdgeTheorem, FailureNoNativeMinimalAbsenceTheorem, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem, FailureNoFullRhoFActionLedger, FailureNoGammaFJFPackage, FailureDFEdgeSupportNotYukawa, FailureNoNumericalYukawaValues, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoAggregateCompressionNative, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t R2++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4, l.AlphaNative)
}

func FormatDomain(d RightDomain) string {
	parts := make([]string, 0, len(d.Cells))
	for _, c := range d.Cells {
		flag := ""
		if c.ActiveDomain {
			flag += " domain"
		}
		if c.Puncture {
			flag += " puncture"
		}
		parts = append(parts, fmt.Sprintf("%s rank=%d B-L=%.0f%s", c.Expression, c.Rank, c.BMinusLTrace, flag))
	}
	return fmt.Sprintf("%s; %s; full=%d domain=%d puncture=%d B-L(domain,puncture,full)=(%.0f,%.0f,%.0f); cells=[%s]", d.RankPattern, d.ActivePattern, d.FullRightRank, d.ActiveRank, d.PunctureRank, d.BMinusLActive, d.BMinusLPuncture, d.BMinusLFull, strings.Join(parts, "; "))
}

func FormatTarget(t LeftTarget) string {
	return fmt.Sprintf("%s; rank=%d color_target=%d lepton_target=%d complete=%t lepto_color_preserved=%t", t.Expression, t.Rank, t.ColorRank, t.LeptonRank, t.Complete, t.LeptoColorPreserved)
}

func FormatEdges(e EdgeSkeleton) string {
	parts := make([]string, 0, len(e.Edges))
	for _, edge := range e.Edges {
		parts = append(parts, fmt.Sprintf("%s: %s -> %s domain_rank=%d target_rank=%d symbolic=%t", edge.Name, edge.DomainExpression, edge.TargetExpression, edge.DomainRank, edge.TargetRank, edge.SymbolicOnly))
	}
	return fmt.Sprintf("%s domain=%d target=%d support_only=%t seal=%t native_DF=%t explicit_DF=%t first_order=%t bimodule=%t puncture_in_domain=%t puncture_null_edge=%t puncture_compatible=%t magnitudes=%t values=%t edges=[%s]", e.Expression, e.DomainRank, e.TargetRank, e.EdgeSupportOnly, e.ConstructedAtSealLevel, e.NativeDFMatrixCertified, e.ExplicitDFMatrixCertified, e.FirstOrderConditionCertified, e.BimoduleCommutantCertified, e.PunctureInDomain, e.PunctureNullEdgeCertified, e.PunctureAbsenceCompatible, e.YukawaMagnitudes, e.NumericalValues, strings.Join(parts, "; "))
}

func FormatShadow(s AggregateShadow) string {
	return fmt.Sprintf("%s; finite_body_seal=%t edge_support_seal=%t native_compression=%t alpha_derived=%t trace_magnitude=%t R3=%t R4=%t", s.Expression, s.FiniteBodyLocationAtSealLevel, s.EdgeSupportSealLevel, s.NativeCompressionTheorem, s.AlphaDerived, s.TraceMagnitudeReadout, s.R3, s.R4)
}

func FormatFirewalls(f Firewalls) string { return f.Verdict }

func FormatImpact(i Impact) string {
	return fmt.Sprintf("%s; edge_seal=%t puncture_compatible=%t puncture_null_native=%t native_DF=%t first_order=%t native_compression_missing=%t alpha_sealed=%t magnitudes_missing=%t updates=(%t,%t,%t) promotions=(R3:%t,R4:%t)", i.Classification, i.SymbolicEdgeSupportSeal, i.PunctureAbsenceCompatible, i.PunctureNullEdgeNative, i.NativeDFMatrixCertified, i.FirstOrderCertified, i.NativeCompressionStillMissing, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func containsAll(haystack, needles []string) bool {
	set := make(map[string]bool, len(haystack))
	for _, h := range haystack {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}
