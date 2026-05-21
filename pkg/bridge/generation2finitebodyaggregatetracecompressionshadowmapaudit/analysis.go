// Package generation2finitebodyaggregatetracecompressionshadowmapaudit implements
// Gate 845: Finite-Body Aggregate Trace-Compression Shadow Map Audit.
//
// Gate 845 follows Gate 844's symbolic finite-Dirac edge-domain seal. Gate 844
// made H_R^min the support-only right edge domain at seal level. Gate 845 asks
// whether the already-derived R2++ aggregate trace-magnitude operator can be
// placed on this finite body as a trace-compression shadow:
//
//	H_R^min = (e_+ tensor P_3) plus (e_- tensor W)
//	H_agg/T = I_{e_+ tensor P_3} plus
//	          [alpha_B P_3 - 3 alpha_B^2(B-L)]_{e_- tensor W}.
//
// This gate reconstructs the Gate 829 trace and square-trace diagnostics from
// the finite-body location. It does not derive alpha_B, certify a native
// compression theorem, produce Yukawa magnitudes, identify particles, promote to
// R3/R4, or update official ledgers.
package generation2finitebodyaggregatetracecompressionshadowmapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE845-FINITE-BODY-AGGREGATE-TRACE-COMPRESSION-SHADOW-MAP-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	RightPairDim   = 2
	LeftSocketDim  = 2

	TopRank       = ColorBlockDim
	RestRank      = WDim
	PunctureRank  = LeptonBlockDim
	HRMinRank     = TopRank + RestRank
	RightFullRank = RightPairDim * WDim
	HLRank        = LeftSocketDim * WDim

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate844Inherited                 = "PASS_GATE844_MINIMAL_RIGHT_EDGE_DOMAIN_INHERITED"
	StatusHRMinTopRestDecompositionAudited = "PASS_H_R_MIN_DECOMPOSES_AS_TOP_PLUS_REST"
	StatusAggregatePlacementAudited        = "PASS_AGGREGATE_OPERATOR_PLACED_ON_FINITE_BODY_SUPPORT_AT_SEAL_LEVEL"
	StatusTraceReconstructionAudited       = "PASS_TRACE_AND_SQUARE_TRACE_RECONSTRUCT_GATE829_OPERATOR"
	StatusOperatorNEffReproduced           = "PASS_OPERATOR_N_EFF_REPRODUCES_GATE829_DIAGNOSTIC"
	StatusPunctureExcludedAudited          = "PASS_PUNCTURE_EXCLUDED_FROM_COMPRESSION_SUPPORT"
	StatusEdgeCompatibilityAudited         = "PASS_EDGE_SUPPORT_COMPATIBILITY_WITH_GATE844_AUDITED"
	StatusAlphaStillSealed                 = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_GATE845"
	StatusOfficialLedgersFrozen            = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusPlusShadow             = "PASS_R2_PLUS_PLUS_PLUS_FINITE_BODY_LOCATED_SHADOW_NOT_R3_OR_R4"
	StatusNoObservedDataUsed               = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate845                  = "FIREWALL_PRESERVED_GATE845_FINITE_BODY_AGGREGATE_SHADOW"

	SupportHRMinAsTopPlusRest                = "CONDITIONAL_SUPPORT_H_R_MIN_EQUALS_E_PLUS_P3_PLUS_E_MINUS_W"
	SupportPiTopRankThree                    = "CONDITIONAL_SUPPORT_PI_TOP_EQUALS_E_PLUS_TENSOR_P3_RANK_THREE"
	SupportPiRestRankFour                    = "CONDITIONAL_SUPPORT_PI_REST_EQUALS_E_MINUS_TENSOR_W_RANK_FOUR"
	SupportFiniteBodyLocation                = "CONDITIONAL_SUPPORT_AGGREGATE_OPERATOR_HAS_FINITE_BODY_LOCATION_AT_SEAL_LEVEL"
	SupportRestBLTransferOnW                 = "CONDITIONAL_SUPPORT_REST_OPERATOR_ACTS_NATURALLY_ON_W_FACTOR"
	SupportTraceReconstructionMatchesGate829 = "CONDITIONAL_SUPPORT_TRACE_RECONSTRUCTION_MATCHES_GATE829"
	SupportOperatorNEffDiagnostic            = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_DIAGNOSTIC_RECONSTRUCTED_FROM_FINITE_BODY_SHADOW"
	SupportPunctureExcluded                  = "CONDITIONAL_SUPPORT_E_PLUS_TENSOR_P1_EXCLUDED_FROM_AGGREGATE_SHADOW_SUPPORT"
	SupportEdgeCompatibleDomain              = "CONDITIONAL_SUPPORT_AGGREGATE_SHADOW_SUPPORT_COMPATIBLE_WITH_SYMBOLIC_EDGE_DOMAIN"
	SupportR2PlusPlusPlus                    = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_FINITE_BODY_LOCATED_AGGREGATE_SHADOW"

	FailureCompressionSealNotNative      = "FAILED_ROUTE_FINITE_BODY_AGGREGATE_COMPRESSION_IS_SEAL_NOT_NATIVE_THEOREM"
	FailureNoNativeCompressionMap        = "FAILED_ROUTE_NO_NATIVE_AGGREGATE_TRACE_COMPRESSION_MAP_CERTIFIED"
	FailureNoNativeShadowFunctional      = "FAILED_ROUTE_NO_NATIVE_TRACE_COMPRESSION_FUNCTIONAL_CERTIFIED"
	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureCompressionDoesNotDeriveAlpha = "FAILED_ROUTE_AGGREGATE_COMPRESSION_DOES_NOT_DERIVE_ALPHA_B"
	FailureDFSupportNotMatrix            = "FAILED_ROUTE_D_F_SUPPORT_GRAPH_IS_NOT_D_F_MATRIX"
	FailureNoExplicitDFMatrix            = "FAILED_ROUTE_NO_EXPLICIT_D_F_MATRIX_CERTIFIED"
	FailureNoFirstOrderProof             = "FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof      = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureDFEdgeSupportNotYukawa        = "FAILED_ROUTE_D_F_EDGE_SUPPORT_NOT_YUKAWA_MAGNITUDE"
	FailureNoTraceMagnitudeReadout       = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoNumericalYukawaValues       = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoNEffUpdate                  = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoPhysicalParticleAssignment  = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT_FROM_AGGREGATE_SHADOW"
	FailureNoRightNeutrinoTheorem        = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoThreeGenerationTheorem      = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                         = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_NOT_R3"
	FailureNotR4                         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlusPlus, R3, R4          bool
	AlphaNative                     bool
}

type SupportBlock struct {
	Name, Expression, Role string
	Rank                   int
	Included               bool
	BMinusLTrace           float64
}

type DomainDecomposition struct {
	Expression, RankPattern                       string
	Top, Rest, Puncture                           SupportBlock
	HRMinRank, RightFullRank, PunctureRank        int
	Orthogonal, CompleteOnHRMin, PunctureExcluded bool
	InheritedGate844, EdgeDomainCompatible        bool
	Supports, Failures                            []string
}

type AggregateOperator struct {
	Expression                                                       string
	AlphaB, TopEigenvalue, RestLeptonEigenvalue, RestColorEigenvalue float64
	TopTrace, RestTrace, TotalTrace                                  float64
	TopSquareTrace, RestSquareTrace, TotalSquareTrace                float64
	OperatorNEff, OfficialNEff, NEffGap                              float64
	FiniteBodyLocationAtSealLevel, NativeCompressionTheorem          bool
	NativeTraceCompressionFunctional, AlphaDerived                   bool
	TraceMagnitudeReadout, R3, R4                                    bool
	Supports, Failures                                               []string
}

type EdgeCompatibility struct {
	Expression                                                                string
	TopTarget, RestColorTarget, RestLeptonTarget                              string
	CompatibleWithGate844, SupportOnly, ExplicitDFMatrix, FirstOrderCertified bool
	BimoduleCommutantCertified, Magnitudes                                    bool
	Supports, Failures                                                        []string
}

type Impact struct {
	Classification                                        string
	FiniteBodyLocation, EdgeCompatible, NativeCompression bool
	AlphaStillSealed, MagnitudesStillMissing              bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs      bool
	CanPromoteToR3, CanPromoteToR4                        bool
}

type Firewalls struct {
	Enforced                                                                            bool
	CompressionSealNotNative, NoNativeCompressionMap, NoNativeShadowFunctional          bool
	AlphaStillSealed, CompressionDoesNotDeriveAlpha                                     bool
	DFSupportNotMatrix, NoExplicitDFMatrix, NoFirstOrderProof, NoBimoduleCommutantProof bool
	DFEdgeSupportNotYukawa, NoTraceMagnitudeReadout, NoNumericalYukawaValues            bool
	NoNEffUpdate, NoCYukawaUpdate, NoPhysicalParticleAssignment, NoRightNeutrinoTheorem bool
	NoThreeGenerationTheorem, NotR3, NotR4                                              bool
	Verdict                                                                             string
}

type Audit struct {
	Ledger       Ledger
	Domain       DomainDecomposition
	Operator     AggregateOperator
	Edges        EdgeCompatibility
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:    buildLedger(),
		Domain:    buildDomain(),
		Operator:  buildOperator(),
		Edges:     buildEdgeCompatibility(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 845 places the Gate 829 aggregate operator on the Gate 844 minimal right edge-domain at seal level: H_R^min=(e_+ tensor P_3) plus (e_- tensor W). It reconstructs the same trace and operator_N_eff diagnostics, but does not derive alpha_B or produce a sector trace-magnitude ledger.",
		Final:     "R2+++ sealed finite-body aggregate trace-compression shadow: finite-body located and edge-compatible, but alpha_B remains sealed, trace magnitudes remain absent, and R3/R4 promotion is blocked.",
	}
	return a, validate(a)
}

func buildLedger() Ledger {
	return Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: operatorNEff(AlphaB), OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlusPlus: true, R3: false, R4: false, AlphaNative: false}
}

func buildDomain() DomainDecomposition {
	top := SupportBlock{Name: "Pi_top", Expression: "e_+ tensor P_3", Role: "dominant color triplet aggregate atom", Rank: TopRank, Included: true, BMinusLTrace: float64(ColorBlockDim) * BMinusLColorWeight}
	rest := SupportBlock{Name: "Pi_rest", Expression: "e_- tensor W", Role: "rest lepto-color quartet", Rank: RestRank, Included: true, BMinusLTrace: BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight}
	puncture := SupportBlock{Name: "Pi_puncture", Expression: "e_+ tensor P_1", Role: "neutral right-lepton puncture excluded from aggregate support", Rank: PunctureRank, Included: false, BMinusLTrace: BMinusLLeptonWeight}
	return DomainDecomposition{
		Expression:  "H_R^min = (e_+ tensor P_3) plus (e_- tensor W)",
		RankPattern: "7=3+4=(8-1)",
		Top:         top, Rest: rest, Puncture: puncture,
		HRMinRank: HRMinRank, RightFullRank: RightFullRank, PunctureRank: PunctureRank,
		Orthogonal: true, CompleteOnHRMin: true, PunctureExcluded: true,
		InheritedGate844: true, EdgeDomainCompatible: true,
		Supports: []string{SupportHRMinAsTopPlusRest, SupportPiTopRankThree, SupportPiRestRankFour, SupportPunctureExcluded, SupportEdgeCompatibleDomain},
		Failures: []string{FailureCompressionSealNotNative, FailureNoNativeCompressionMap},
	}
}

func buildOperator() AggregateOperator {
	a := AlphaB
	restLepton := 3 * a * a
	restColor := a * (1 - a)
	topTrace := float64(TopRank)
	restTrace := restLepton + float64(ColorBlockDim)*restColor
	totalTrace := topTrace + restTrace
	topSquare := float64(TopRank)
	restSquare := restLepton*restLepton + float64(ColorBlockDim)*restColor*restColor
	totalSquare := topSquare + restSquare
	neff := totalTrace * totalTrace / totalSquare
	return AggregateOperator{
		Expression: "H_agg/T = I_{e_+ tensor P_3} plus [alpha_B P_3 - 3 alpha_B^2(B-L)]_{e_- tensor W}",
		AlphaB:     a, TopEigenvalue: 1, RestLeptonEigenvalue: restLepton, RestColorEigenvalue: restColor,
		TopTrace: topTrace, RestTrace: restTrace, TotalTrace: totalTrace,
		TopSquareTrace: topSquare, RestSquareTrace: restSquare, TotalSquareTrace: totalSquare,
		OperatorNEff: neff, OfficialNEff: OfficialNEff, NEffGap: neff - OfficialNEff,
		FiniteBodyLocationAtSealLevel: true, NativeCompressionTheorem: false,
		NativeTraceCompressionFunctional: false, AlphaDerived: false, TraceMagnitudeReadout: false, R3: false, R4: false,
		Supports: []string{SupportFiniteBodyLocation, SupportRestBLTransferOnW, SupportTraceReconstructionMatchesGate829, SupportOperatorNEffDiagnostic, SupportR2PlusPlusPlus},
		Failures: []string{FailureCompressionSealNotNative, FailureNoNativeCompressionMap, FailureNoNativeShadowFunctional, FailureAlphaStillSealed, FailureCompressionDoesNotDeriveAlpha, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4},
	}
}

func buildEdgeCompatibility() EdgeCompatibility {
	return EdgeCompatibility{
		Expression:            "Gate844 D_F^supp-compatible domain restriction on H_R^min",
		TopTarget:             "e_+ tensor P_3 -> C_L^2 tensor P_3",
		RestColorTarget:       "e_- tensor P_3 -> C_L^2 tensor P_3",
		RestLeptonTarget:      "e_- tensor P_1 -> C_L^2 tensor P_1",
		CompatibleWithGate844: true, SupportOnly: true, ExplicitDFMatrix: false,
		FirstOrderCertified: false, BimoduleCommutantCertified: false, Magnitudes: false,
		Supports: []string{SupportEdgeCompatibleDomain},
		Failures: []string{FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa, FailureNoNumericalYukawaValues},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++ finite-body located aggregate trace-compression shadow; not R3/R4", FiniteBodyLocation: true, EdgeCompatible: true, NativeCompression: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, CompressionSealNotNative: true, NoNativeCompressionMap: true, NoNativeShadowFunctional: true, AlphaStillSealed: true, CompressionDoesNotDeriveAlpha: true, DFSupportNotMatrix: true, NoExplicitDFMatrix: true, NoFirstOrderProof: true, NoBimoduleCommutantProof: true, DFEdgeSupportNotYukawa: true, NoTraceMagnitudeReadout: true, NoNumericalYukawaValues: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoPhysicalParticleAssignment: true, NoRightNeutrinoTheorem: true, NoThreeGenerationTheorem: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate845}
}

func validate(a Audit) error {
	if !a.Domain.InheritedGate844 || !a.Domain.Orthogonal || !a.Domain.CompleteOnHRMin || !a.Domain.PunctureExcluded || !a.Domain.EdgeDomainCompatible || a.Domain.HRMinRank != HRMinRank || a.Domain.Top.Rank != TopRank || a.Domain.Rest.Rank != RestRank || a.Domain.Puncture.Rank != PunctureRank {
		return fmt.Errorf("domain invalid: %s", FormatDomain(a.Domain))
	}
	if math.Abs(a.Domain.Top.BMinusLTrace-1) > 1e-12 || math.Abs(a.Domain.Rest.BMinusLTrace) > 1e-12 || math.Abs(a.Domain.Puncture.BMinusLTrace+1) > 1e-12 {
		return fmt.Errorf("domain B-L traces invalid: %s", FormatDomain(a.Domain))
	}
	if !a.Operator.FiniteBodyLocationAtSealLevel || a.Operator.NativeCompressionTheorem || a.Operator.NativeTraceCompressionFunctional || a.Operator.AlphaDerived || a.Operator.TraceMagnitudeReadout || a.Operator.R3 || a.Operator.R4 {
		return fmt.Errorf("operator over/under-certified: %s", FormatOperator(a.Operator))
	}
	wantTrace := 3 + 3*AlphaB
	wantSquare := 3 + 3*AlphaB*AlphaB - 6*math.Pow(AlphaB, 3) + 12*math.Pow(AlphaB, 4)
	if math.Abs(a.Operator.TotalTrace-wantTrace) > 1e-14 || math.Abs(a.Operator.TotalSquareTrace-wantSquare) > 1e-14 || math.Abs(a.Operator.OperatorNEff-operatorNEff(AlphaB)) > 1e-14 {
		return fmt.Errorf("trace reconstruction invalid: %s", FormatOperator(a.Operator))
	}
	if !a.Edges.CompatibleWithGate844 || !a.Edges.SupportOnly || a.Edges.ExplicitDFMatrix || a.Edges.FirstOrderCertified || a.Edges.BimoduleCommutantCertified || a.Edges.Magnitudes {
		return fmt.Errorf("edge compatibility invalid: %s", FormatEdges(a.Edges))
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.R2PlusPlusPlus || a.Ledger.R3 || a.Ledger.R4 || a.Ledger.AlphaNative {
		return fmt.Errorf("ledger invalid: %s", FormatLedger(a.Ledger))
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate845 || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoNEffUpdate || !a.Firewalls.NoCYukawaUpdate {
		return fmt.Errorf("firewall invalid: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func operatorNEff(alpha float64) float64 {
	return 3 * (1 + alpha) * (1 + alpha) / (1 + alpha*alpha - 2*math.Pow(alpha, 3) + 4*math.Pow(alpha, 4))
}

func Statuses() []string {
	return []string{
		StatusGate844Inherited, StatusHRMinTopRestDecompositionAudited, StatusAggregatePlacementAudited, StatusTraceReconstructionAudited, StatusOperatorNEffReproduced, StatusPunctureExcludedAudited, StatusEdgeCompatibilityAudited, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusPlusShadow, StatusNoObservedDataUsed, StatusFirewallGate845,
		SupportHRMinAsTopPlusRest, SupportPiTopRankThree, SupportPiRestRankFour, SupportFiniteBodyLocation, SupportRestBLTransferOnW, SupportTraceReconstructionMatchesGate829, SupportOperatorNEffDiagnostic, SupportPunctureExcluded, SupportEdgeCompatibleDomain, SupportR2PlusPlusPlus,
		FailureCompressionSealNotNative, FailureNoNativeCompressionMap, FailureNoNativeShadowFunctional, FailureAlphaStillSealed, FailureCompressionDoesNotDeriveAlpha, FailureDFSupportNotMatrix, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureDFEdgeSupportNotYukawa, FailureNoTraceMagnitudeReadout, FailureNoNumericalYukawaValues, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoPhysicalParticleAssignment, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t R2+++ =%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.R2PlusPlusPlus, l.R3, l.R4, l.AlphaNative)
}

func FormatDomain(d DomainDecomposition) string {
	return fmt.Sprintf("%s; %s; top=%s rank=%d B-L=%.0f; rest=%s rank=%d B-L=%.0f; puncture=%s rank=%d B-L=%.0f included=%t; orthogonal=%t complete=%t", d.RankPattern, d.Expression, d.Top.Expression, d.Top.Rank, d.Top.BMinusLTrace, d.Rest.Expression, d.Rest.Rank, d.Rest.BMinusLTrace, d.Puncture.Expression, d.Puncture.Rank, d.Puncture.BMinusLTrace, d.Puncture.Included, d.Orthogonal, d.CompleteOnHRMin)
}

func FormatOperator(o AggregateOperator) string {
	return fmt.Sprintf("%s; alpha=%.16g eig(top,lepton,color)=(%.16g,%.16g,%.16g); trace(top,rest,total)=(%.16g,%.16g,%.16g); square(top,rest,total)=(%.16g,%.16g,%.16g); operator_N_eff=%.16g official_N_eff=%.16g gap=%.16g finite_body_seal=%t native_compression=%t alpha_derived=%t trace_magnitude=%t R3=%t R4=%t", o.Expression, o.AlphaB, o.TopEigenvalue, o.RestLeptonEigenvalue, o.RestColorEigenvalue, o.TopTrace, o.RestTrace, o.TotalTrace, o.TopSquareTrace, o.RestSquareTrace, o.TotalSquareTrace, o.OperatorNEff, o.OfficialNEff, o.NEffGap, o.FiniteBodyLocationAtSealLevel, o.NativeCompressionTheorem, o.AlphaDerived, o.TraceMagnitudeReadout, o.R3, o.R4)
}

func FormatEdges(e EdgeCompatibility) string {
	return fmt.Sprintf("%s; compatible=%t support_only=%t explicit_DF=%t first_order=%t bimodule=%t magnitudes=%t edges=[%s; %s; %s]", e.Expression, e.CompatibleWithGate844, e.SupportOnly, e.ExplicitDFMatrix, e.FirstOrderCertified, e.BimoduleCommutantCertified, e.Magnitudes, e.TopTarget, e.RestColorTarget, e.RestLeptonTarget)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("%s; finite_body=%t edge_compatible=%t native_compression=%t alpha_sealed=%t magnitudes_missing=%t updates=(%t,%t,%t) promotions=(R3:%t,R4:%t)", i.Classification, i.FiniteBodyLocation, i.EdgeCompatible, i.NativeCompression, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string { return f.Verdict }

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

func JoinStatuses() string { return strings.Join(Statuses(), "\n") }
