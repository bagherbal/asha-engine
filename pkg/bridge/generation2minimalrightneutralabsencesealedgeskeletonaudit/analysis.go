// Package generation2minimalrightneutralabsencesealedgeskeletonaudit implements
// Gate 843: Minimal RightNeutral Absence Seal and Edge-Skeleton Audit.
//
// Gate 843 follows Gate 842's four-cell right lepto-color rectangle. It audits
// the representation choice that removes the neutral right-lepton singleton
// e_+ tensor P_1 from the active right module at seal level. This explains the
// active rank-seven support as 8-1 inside C_R^2 tensor W, not as a K7 theorem.
// The gate compares the minimal absent-cell branch with the extended
// neutral-inclusive branch, registers the finite-body location of the R2++
// aggregate shadow at seal level, and preserves all native, edge, magnitude,
// particle-assignment, R3, and R4 firewalls.
package generation2minimalrightneutralabsencesealedgeskeletonaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE843-MINIMAL-RIGHT-NEUTRAL-ABSENCE-SEAL-EDGE-SKELETON-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim     = 1
	ColorBlockDim      = 3
	WDim               = LeptonBlockDim + ColorBlockDim
	RightSocketPairDim = 2
	FullRightRank      = RightSocketPairDim * WDim
	EPlusP3Rank        = ColorBlockDim
	EPlusP1Rank        = LeptonBlockDim
	EMinusP3Rank       = ColorBlockDim
	EMinusP1Rank       = LeptonBlockDim
	MinimalRightRank   = EPlusP3Rank + EMinusP3Rank + EMinusP1Rank
	PunctureRank       = EPlusP1Rank
	RestQuartetRank    = EMinusP3Rank + EMinusP1Rank
	FullParticleRank   = 16
	FullFiniteRank     = 32

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate842Inherited             = "PASS_GATE842_FOUR_CELL_LEDGER_INHERITED"
	StatusMinimalAbsenceSealAudited    = "PASS_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL_AUDITED"
	StatusBranchComparisonAudited      = "PASS_MINIMAL_VS_EXTENDED_RIGHT_RECTANGLE_BRANCHES_AUDITED"
	StatusActiveSevenAsEightMinusOne   = "PASS_ACTIVE_RIGHT_SUPPORT_SEVEN_EQUALS_EIGHT_MINUS_ONE_CERTIFIED_AT_SUPPORT_LEVEL"
	StatusPunctureComplementPreserved  = "PASS_NEUTRAL_SINGLETON_PUNCTURE_COMPLEMENT_PRESERVED"
	StatusBMinusLCompensationPreserved = "PASS_B_MINUS_L_COMPENSATION_PATTERN_PRESERVED"
	StatusSealOrientationAudited       = "PASS_DOMINANT_REST_ORIENTATION_SEAL_AUDITED"
	StatusFiniteBodyPlacementSealed    = "PASS_AGGREGATE_FINITE_BODY_LOCATION_REGISTERED_AT_SEAL_LEVEL"
	StatusEdgeSkeletonAudited          = "PASS_D_F_EDGE_SKELETON_ROUTE_AUDITED"
	StatusMagnitudeFirewallPreserved   = "PASS_ABSENCE_SEAL_NOT_TRACE_MAGNITUDE_READOUT"
	StatusAlphaStillSealed             = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_GATE843"
	StatusOfficialLedgersFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusSealedShadow       = "PASS_R2_PLUS_PLUS_SEALED_FINITE_BODY_SHADOW_NOT_R3_OR_R4"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate843              = "FIREWALL_PRESERVED_GATE843_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL"

	SupportGate842Inherited                  = "CONDITIONAL_SUPPORT_GATE842_RIGHT_RECTANGLE_3_1_3_1_INHERITED"
	SupportMinimalRightNeutralAbsenceSeal    = "CONDITIONAL_SUPPORT_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL_ADMITTED_AT_BRIDGE_LAYER"
	SupportHRMinRankSeven                    = "CONDITIONAL_SUPPORT_H_R_MIN_EQUALS_RIGHT_RECTANGLE_MINUS_E_PLUS_P1_HAS_RANK_SEVEN"
	SupportExtendedNeutralInclusiveRankEight = "CONDITIONAL_SUPPORT_EXTENDED_NEUTRAL_INCLUSIVE_RIGHT_RECTANGLE_HAS_RANK_EIGHT"
	SupportR2PrefersMinimalBranch            = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_AGGREGATE_SUPPORT_PREFERS_MINIMAL_ABSENT_CELL_BRANCH"
	SupportPunctureIsCompensatingSingleton   = "CONDITIONAL_SUPPORT_E_PLUS_TENSOR_P1_IS_B_MINUS_L_COMPENSATING_SINGLETON"
	SupportDominantTripletAtSealLevel        = "CONDITIONAL_SUPPORT_I3_LOCATED_ON_E_PLUS_TENSOR_P3_AT_ABSENCE_SEAL_LEVEL"
	SupportRestQuartetAtSealLevel            = "CONDITIONAL_SUPPORT_REST_W_LOCATED_ON_E_MINUS_TENSOR_W_AT_ABSENCE_SEAL_LEVEL"
	SupportAggregateShadowAtSealLevel        = "CONDITIONAL_SUPPORT_H_TOTAL_HAS_FINITE_BODY_TRACE_COMPRESSION_SHADOW_AT_SEAL_LEVEL"
	SupportPunctureAsAbsentNullEdgeCandidate = "CONDITIONAL_SUPPORT_PUNCTURE_AS_ABSENT_NULL_EDGE_CANDIDATE_ONLY"

	FailureAbsenceSealNotNative             = "FAILED_ROUTE_MINIMAL_RIGHT_NEUTRAL_ABSENCE_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoDFEdgeGraph                    = "FAILED_ROUTE_NO_D_F_EDGE_GRAPH_TO_CERTIFY_MINIMAL_ABSENCE"
	FailureNoNullEdgeTheorem                = "FAILED_ROUTE_NO_NULL_EDGE_THEOREM_FOR_E_PLUS_TENSOR_P1"
	FailureNoNativeMinimalAbsenceTheorem    = "FAILED_ROUTE_NO_NATIVE_MINIMAL_RIGHT_NEUTRAL_ABSENCE_THEOREM"
	FailureNoPhysicalParticleAssignment     = "FAILED_ROUTE_NEUTRAL_SINGLETON_NOT_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoRightNeutrinoTheorem           = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureExtendedBranchNeedsProjectionLaw = "FAILED_ROUTE_EXTENDED_NEUTRAL_INCLUSIVE_BRANCH_REQUIRES_EXTRA_PROJECTION_OR_EXCLUSION_LAW"
	FailureNoFullRhoFActionLedger           = "FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoExplicitDFMatrix               = "FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_OR_SYMBOLIC_EDGE_GRAPH_CERTIFIED"
	FailureCompressionSealNotNativeMap      = "FAILED_ROUTE_FINITE_BODY_LOCATION_IS_SEAL_NOT_NATIVE_TRACE_COMPRESSION_MAP"
	FailureNoAggregateCompressionTheorem    = "FAILED_ROUTE_NO_NATIVE_AGGREGATE_TRACE_COMPRESSION_THEOREM"
	FailureNoAlphaDerivation                = "FAILED_ROUTE_MINIMAL_ABSENCE_SEAL_DOES_NOT_DERIVE_ALPHA_B"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout          = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoNEffUpdate                     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                  = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit              = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoThreeGenerationTheorem         = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                            = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlus, R3, R4              bool
	AlphaNative                     bool
}

type Cell struct {
	Name, Expression, SafeLabel string
	Rank                        int
	BMinusLTrace                float64
	Active, Puncture            bool
}

type RightRectangle struct {
	Cells                                       []Cell
	FullRank, ActiveRank, PunctureRank          int
	MinimalRank, ExtendedRank                   int
	RankPattern, MinimalPattern                 string
	InheritedFromGate842, Orthogonal            bool
	Complete, ActiveIsFullMinusPuncture         bool
	BMinusLActive, BMinusLPuncture, BMinusLFull float64
	Supports, Failures                          []string
}

type BranchComparison struct {
	MinimalExpression, ExtendedExpression         string
	MinimalRank, ExtendedRank, PunctureRank       int
	MinimalBranchAdmittedAsSeal                   bool
	MinimalBranchNative                           bool
	ExtendedBranchAvailable                       bool
	ExtendedBranchMatchesR2Support                bool
	ExtendedBranchNeedsExtraProjectionOrExclusion bool
	R2PlusPlusPrefersMinimalBranch                bool
	Supports, Failures                            []string
}

type OrientationSeal struct {
	DominantExpression, RestExpression, PunctureExpression string
	DominantRank, RestRank, TotalRank                      int
	DominantLocationSealed, RestLocationSealed             bool
	NativeOrientationTheorem                               bool
	PhysicalParticleAssignment                             bool
	Supports, Failures                                     []string
}

type EdgeSkeleton struct {
	PunctureExpression                              string
	DFEdgeGraphAvailable, ExplicitDFMatrixAvailable bool
	NullEdgeCertified, MinimalAbsenceEdgeCertified  bool
	AbsentNullEdgeCandidateOnly                     bool
	PhysicalRightNeutrinoTheorem                    bool
	Supports, Failures                              []string
}

type AggregatePlacement struct {
	Expression                                              string
	FiniteBodyLocationAtSealLevel, NativeCompressionTheorem bool
	TraceCompressionShadowAtSealLevel                       bool
	AlphaDerived, TraceMagnitudeReadout, R3, R4             bool
	Supports, Failures                                      []string
}

type Firewalls struct {
	Enforced                                                     bool
	AbsenceSealNotNative, NoDFEdgeGraph, NoNullEdgeTheorem       bool
	NoNativeMinimalAbsenceTheorem, NoPhysicalParticleAssignment  bool
	NoRightNeutrinoTheorem, ExtendedBranchNeedsProjectionLaw     bool
	NoFullRhoFActionLedger, NoExplicitDFMatrix                   bool
	CompressionSealNotNativeMap, NoAggregateCompressionTheorem   bool
	NoAlphaDerivation, AlphaStillSealed, NoTraceMagnitudeReadout bool
	NoNEffUpdate, NoCYukawaUpdate, NoObservedYukawaFit           bool
	NoThreeGenerationTheorem, NotR3, NotR4                       bool
	Verdict                                                      string
}

type Impact struct {
	Classification                                          string
	MinimalAbsenceSealAdmitted, FiniteBodyLocationSealed    bool
	NullEdgeStillUncertified, NativeCompressionStillMissing bool
	AlphaStillSealed, MagnitudesStillMissing                bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs        bool
	CanPromoteToR3, CanPromoteToR4                          bool
}

type Audit struct {
	Ledger       Ledger
	Rectangle    RightRectangle
	Branches     BranchComparison
	Orientation  OrientationSeal
	Edge         EdgeSkeleton
	Placement    AggregatePlacement
	Firewalls    Firewalls
	Impact       Impact
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	rect := buildRectangle()
	branches := buildBranches()
	orientation := buildOrientation()
	edge := buildEdge()
	placement := buildPlacement()
	firewalls := buildFirewalls()
	impact := Impact{
		Classification:                "R2++ sealed finite-body trace-compression shadow: minimal right neutral absence admitted as bridge seal; not R3 or R4",
		MinimalAbsenceSealAdmitted:    true,
		FiniteBodyLocationSealed:      true,
		NullEdgeStillUncertified:      true,
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
		Ledger:      Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaNative: false},
		Rectangle:   rect,
		Branches:    branches,
		Orientation: orientation,
		Edge:        edge,
		Placement:   placement,
		Firewalls:   firewalls,
		Impact:      impact,
		Truth:       "Gate 843 admits a minimal right-neutral absence seal: the active right support is 8-1 = 7 inside C_R^2 tensor W, with e_+ tensor P_1 as the B-L compensating puncture. This is a finite-body location seal, not a native edge theorem or trace-magnitude readout.",
		Final:       "The R2++ aggregate operator may now be located, at seal level, on I_{e_+ tensor P_3} plus the e_- tensor W rest quartet. Alpha_B, D_F null-edge proof, trace magnitudes, physical particle assignment, R3, and R4 remain blocked.",
	}
	return a, validate(a)
}

func buildRectangle() RightRectangle {
	cells := []Cell{
		{Name: "character-plus color triplet", Expression: "e_+ tensor P_3", SafeLabel: "dominant character-color triplet candidate", Rank: EPlusP3Rank, BMinusLTrace: 1, Active: true},
		{Name: "character-plus lepton singleton", Expression: "e_+ tensor P_1", SafeLabel: "neutral right-lepton puncture / B-L compensating singleton", Rank: EPlusP1Rank, BMinusLTrace: -1, Puncture: true},
		{Name: "character-minus color triplet", Expression: "e_- tensor P_3", SafeLabel: "rest character-color triplet", Rank: EMinusP3Rank, BMinusLTrace: 1, Active: true},
		{Name: "character-minus lepton singleton", Expression: "e_- tensor P_1", SafeLabel: "rest character-lepton singleton", Rank: EMinusP1Rank, BMinusLTrace: -1, Active: true},
	}
	active := 0
	puncture := 0
	bActive := 0.0
	bPuncture := 0.0
	for _, c := range cells {
		if c.Active {
			active += c.Rank
			bActive += c.BMinusLTrace
		}
		if c.Puncture {
			puncture += c.Rank
			bPuncture += c.BMinusLTrace
		}
	}
	return RightRectangle{
		Cells:                     cells,
		FullRank:                  FullRightRank,
		ActiveRank:                active,
		PunctureRank:              puncture,
		MinimalRank:               MinimalRightRank,
		ExtendedRank:              FullRightRank,
		RankPattern:               "8=3+1+3+1",
		MinimalPattern:            "7=8-1=3+3+1",
		InheritedFromGate842:      true,
		Orthogonal:                true,
		Complete:                  true,
		ActiveIsFullMinusPuncture: active == FullRightRank-puncture,
		BMinusLActive:             bActive,
		BMinusLPuncture:           bPuncture,
		BMinusLFull:               bActive + bPuncture,
		Supports:                  []string{SupportGate842Inherited, SupportHRMinRankSeven, SupportPunctureIsCompensatingSingleton},
		Failures:                  []string{FailureAbsenceSealNotNative, FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem},
	}
}

func buildBranches() BranchComparison {
	return BranchComparison{
		MinimalExpression:              "H_R^min = (C_R^2 tensor W) minus (e_+ tensor P_1)",
		ExtendedExpression:             "H_R^ext = C_R^2 tensor W",
		MinimalRank:                    MinimalRightRank,
		ExtendedRank:                   FullRightRank,
		PunctureRank:                   PunctureRank,
		MinimalBranchAdmittedAsSeal:    true,
		MinimalBranchNative:            false,
		ExtendedBranchAvailable:        true,
		ExtendedBranchMatchesR2Support: false,
		ExtendedBranchNeedsExtraProjectionOrExclusion: true,
		R2PlusPlusPrefersMinimalBranch:                true,
		Supports:                                      []string{SupportMinimalRightNeutralAbsenceSeal, SupportExtendedNeutralInclusiveRankEight, SupportR2PrefersMinimalBranch},
		Failures:                                      []string{FailureAbsenceSealNotNative, FailureExtendedBranchNeedsProjectionLaw, FailureNoNativeMinimalAbsenceTheorem},
	}
}

func buildOrientation() OrientationSeal {
	return OrientationSeal{
		DominantExpression:         "I_{e_+ tensor P_3}",
		RestExpression:             "[alpha_B P_3 - 3 alpha_B^2(B-L)] on e_- tensor W",
		PunctureExpression:         "e_+ tensor P_1",
		DominantRank:               EPlusP3Rank,
		RestRank:                   RestQuartetRank,
		TotalRank:                  MinimalRightRank,
		DominantLocationSealed:     true,
		RestLocationSealed:         true,
		NativeOrientationTheorem:   false,
		PhysicalParticleAssignment: false,
		Supports:                   []string{SupportDominantTripletAtSealLevel, SupportRestQuartetAtSealLevel, SupportAggregateShadowAtSealLevel},
		Failures:                   []string{FailureAbsenceSealNotNative, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem},
	}
}

func buildEdge() EdgeSkeleton {
	return EdgeSkeleton{
		PunctureExpression:           "e_+ tensor P_1",
		DFEdgeGraphAvailable:         false,
		ExplicitDFMatrixAvailable:    false,
		NullEdgeCertified:            false,
		MinimalAbsenceEdgeCertified:  false,
		AbsentNullEdgeCandidateOnly:  true,
		PhysicalRightNeutrinoTheorem: false,
		Supports:                     []string{SupportPunctureAsAbsentNullEdgeCandidate},
		Failures:                     []string{FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoNativeMinimalAbsenceTheorem, FailureNoExplicitDFMatrix, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem},
	}
}

func buildPlacement() AggregatePlacement {
	return AggregatePlacement{
		Expression:                        "H_total/T = I_{e_+ tensor P_3} plus [alpha_B P_3 - 3 alpha_B^2(B-L)]_{e_- tensor W}",
		FiniteBodyLocationAtSealLevel:     true,
		NativeCompressionTheorem:          false,
		TraceCompressionShadowAtSealLevel: true,
		AlphaDerived:                      false,
		TraceMagnitudeReadout:             false,
		R3:                                false,
		R4:                                false,
		Supports:                          []string{SupportAggregateShadowAtSealLevel, SupportDominantTripletAtSealLevel, SupportRestQuartetAtSealLevel},
		Failures:                          []string{FailureCompressionSealNotNativeMap, FailureNoAggregateCompressionTheorem, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		Enforced:                         true,
		AbsenceSealNotNative:             true,
		NoDFEdgeGraph:                    true,
		NoNullEdgeTheorem:                true,
		NoNativeMinimalAbsenceTheorem:    true,
		NoPhysicalParticleAssignment:     true,
		NoRightNeutrinoTheorem:           true,
		ExtendedBranchNeedsProjectionLaw: true,
		NoFullRhoFActionLedger:           true,
		NoExplicitDFMatrix:               true,
		CompressionSealNotNativeMap:      true,
		NoAggregateCompressionTheorem:    true,
		NoAlphaDerivation:                true,
		AlphaStillSealed:                 true,
		NoTraceMagnitudeReadout:          true,
		NoNEffUpdate:                     true,
		NoCYukawaUpdate:                  true,
		NoObservedYukawaFit:              true,
		NoThreeGenerationTheorem:         true,
		NotR3:                            true,
		NotR4:                            true,
		Verdict:                          StatusFirewallGate843,
	}
}

func validate(a Audit) error {
	if a.Rectangle.FullRank != FullRightRank || a.Rectangle.ActiveRank != MinimalRightRank || a.Rectangle.PunctureRank != PunctureRank {
		return fmt.Errorf("right rectangle rank mismatch: %s", FormatRectangle(a.Rectangle))
	}
	if !a.Rectangle.ActiveIsFullMinusPuncture || a.Rectangle.ActiveRank+a.Rectangle.PunctureRank != a.Rectangle.FullRank {
		return fmt.Errorf("active support is not full minus puncture: %s", FormatRectangle(a.Rectangle))
	}
	if math.Abs(a.Rectangle.BMinusLActive-1) > 1e-12 || math.Abs(a.Rectangle.BMinusLPuncture+1) > 1e-12 || math.Abs(a.Rectangle.BMinusLFull) > 1e-12 {
		return fmt.Errorf("B-L compensation mismatch: %s", FormatRectangle(a.Rectangle))
	}
	if !a.Branches.MinimalBranchAdmittedAsSeal || a.Branches.MinimalBranchNative || !a.Branches.ExtendedBranchNeedsExtraProjectionOrExclusion {
		return fmt.Errorf("branch comparison invalid: %s", FormatBranches(a.Branches))
	}
	if !a.Orientation.DominantLocationSealed || !a.Orientation.RestLocationSealed || a.Orientation.NativeOrientationTheorem || a.Orientation.PhysicalParticleAssignment {
		return fmt.Errorf("orientation invalid: %s", FormatOrientation(a.Orientation))
	}
	if a.Edge.DFEdgeGraphAvailable || a.Edge.NullEdgeCertified || a.Edge.MinimalAbsenceEdgeCertified || !a.Edge.AbsentNullEdgeCandidateOnly {
		return fmt.Errorf("edge route over-certified: %s", FormatEdge(a.Edge))
	}
	if !a.Placement.FiniteBodyLocationAtSealLevel || !a.Placement.TraceCompressionShadowAtSealLevel || a.Placement.NativeCompressionTheorem || a.Placement.AlphaDerived || a.Placement.TraceMagnitudeReadout || a.Placement.R3 || a.Placement.R4 {
		return fmt.Errorf("placement invalid: %s", FormatPlacement(a.Placement))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate843 || !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		return fmt.Errorf("firewall or ledger invalid")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate842Inherited, StatusMinimalAbsenceSealAudited, StatusBranchComparisonAudited, StatusActiveSevenAsEightMinusOne, StatusPunctureComplementPreserved, StatusBMinusLCompensationPreserved, StatusSealOrientationAudited, StatusFiniteBodyPlacementSealed, StatusEdgeSkeletonAudited, StatusMagnitudeFirewallPreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusSealedShadow, StatusNoObservedDataUsed, StatusFirewallGate843,
		SupportGate842Inherited, SupportMinimalRightNeutralAbsenceSeal, SupportHRMinRankSeven, SupportExtendedNeutralInclusiveRankEight, SupportR2PrefersMinimalBranch, SupportPunctureIsCompensatingSingleton, SupportDominantTripletAtSealLevel, SupportRestQuartetAtSealLevel, SupportAggregateShadowAtSealLevel, SupportPunctureAsAbsentNullEdgeCandidate,
		FailureAbsenceSealNotNative, FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoNativeMinimalAbsenceTheorem, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem, FailureExtendedBranchNeedsProjectionLaw, FailureNoFullRhoFActionLedger, FailureNoExplicitDFMatrix, FailureCompressionSealNotNativeMap, FailureNoAggregateCompressionTheorem, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t R2++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4, l.AlphaNative)
}

func FormatRectangle(r RightRectangle) string {
	parts := make([]string, 0, len(r.Cells))
	for _, c := range r.Cells {
		flag := ""
		if c.Active {
			flag += " active"
		}
		if c.Puncture {
			flag += " puncture"
		}
		parts = append(parts, fmt.Sprintf("%s rank=%d B-L=%.0f%s", c.Expression, c.Rank, c.BMinusLTrace, flag))
	}
	return fmt.Sprintf("%s; %s; full=%d active=%d puncture=%d B-L(active,puncture,full)=(%.0f,%.0f,%.0f); cells=[%s]", r.RankPattern, r.MinimalPattern, r.FullRank, r.ActiveRank, r.PunctureRank, r.BMinusLActive, r.BMinusLPuncture, r.BMinusLFull, strings.Join(parts, "; "))
}

func FormatBranches(b BranchComparison) string {
	return fmt.Sprintf("minimal=%q rank=%d seal=%t native=%t; extended=%q rank=%d matches_R2=%t needs_projection=%t prefers_minimal=%t", b.MinimalExpression, b.MinimalRank, b.MinimalBranchAdmittedAsSeal, b.MinimalBranchNative, b.ExtendedExpression, b.ExtendedRank, b.ExtendedBranchMatchesR2Support, b.ExtendedBranchNeedsExtraProjectionOrExclusion, b.R2PlusPlusPrefersMinimalBranch)
}

func FormatOrientation(o OrientationSeal) string {
	return fmt.Sprintf("dominant=%s rank=%d sealed=%t; rest=%s rank=%d sealed=%t; puncture=%s; native_orientation=%t physical_assignment=%t", o.DominantExpression, o.DominantRank, o.DominantLocationSealed, o.RestExpression, o.RestRank, o.RestLocationSealed, o.PunctureExpression, o.NativeOrientationTheorem, o.PhysicalParticleAssignment)
}

func FormatEdge(e EdgeSkeleton) string {
	return fmt.Sprintf("puncture=%s D_F_graph=%t explicit_D_F=%t null_edge=%t minimal_absence_edge=%t candidate_only=%t physical_RN=%t", e.PunctureExpression, e.DFEdgeGraphAvailable, e.ExplicitDFMatrixAvailable, e.NullEdgeCertified, e.MinimalAbsenceEdgeCertified, e.AbsentNullEdgeCandidateOnly, e.PhysicalRightNeutrinoTheorem)
}

func FormatPlacement(p AggregatePlacement) string {
	return fmt.Sprintf("%s; finite_body_seal=%t native_compression=%t trace_shadow_seal=%t alpha_derived=%t trace_magnitude=%t R3=%t R4=%t", p.Expression, p.FiniteBodyLocationAtSealLevel, p.NativeCompressionTheorem, p.TraceCompressionShadowAtSealLevel, p.AlphaDerived, p.TraceMagnitudeReadout, p.R3, p.R4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("%s; absence_seal=%t finite_location=%t null_edge_uncertified=%t native_compression_missing=%t alpha_sealed=%t magnitudes_missing=%t updates=(%t,%t,%t) promotions=(R3:%t,R4:%t)", i.Classification, i.MinimalAbsenceSealAdmitted, i.FiniteBodyLocationSealed, i.NullEdgeStillUncertified, i.NativeCompressionStillMissing, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
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
