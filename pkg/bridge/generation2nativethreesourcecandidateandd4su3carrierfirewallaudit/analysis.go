// Package generation2nativethreesourcecandidateandd4su3carrierfirewallaudit implements
// Gate 799: Native Three-Source Candidate Ranking and D4/SU3 Carrier Firewall Audit.
//
// Gate 799 ranks all currently visible sources of the number three by typed
// strength, identifies the missing carrier/readout maps needed to upgrade each
// source from motif or diagnostic to theorem, and preserves the firewall that
// N_eff≈3 is currently only source-typed by color-tripled top dominance.
package generation2nativethreesourcecandidateandd4su3carrierfirewallaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE799-NATIVE-THREE-SOURCE-CANDIDATE-D4-SU3-CARRIER-FIREWALL-AUDIT"

	StatusGate798Inherited      = "PASS_GATE798_HIGH_SCALE_GJ_COLOR_THREE_DIAGNOSTIC_INHERITED"
	StatusCurrentThreeInherited = "PASS_CURRENT_CERTIFIED_THREE_SOURCE_STATUS_INHERITED"
	StatusPackageRequirements   = "PASS_NATIVE_THREE_SOURCE_PACKAGE_REQUIREMENTS_DEFINED"
	StatusColorSU3Audited       = "PASS_COLOR_SU3_MULTIPLICITY_SOURCE_AUDITED"
	StatusGJAudited             = "PASS_GEORGI_JARLSKOG_CLEBSCH_THREE_SOURCE_AUDITED"
	StatusD4Audited             = "PASS_D4_TRIALITY_SOURCE_CANDIDATE_AUDITED"
	StatusA2SU3Audited          = "PASS_SU3_A2_HEXAGONAL_CARRIER_CANDIDATE_AUDITED"
	StatusGenerationAudited     = "PASS_GENERATION_COUNT_SOURCE_CANDIDATE_AUDITED"
	StatusK7HodgeAudited        = "PASS_K7_HODGE_43_SOURCE_CANDIDATE_AUDITED"
	StatusFockProjectiveAudited = "PASS_FOCK_PROJECTIVE_13_SOURCE_CANDIDATE_AUDITED"
	StatusCandidatesRanked      = "PASS_THREE_SOURCE_CANDIDATES_RANKED"
	StatusMethodologicalBranch  = "PASS_METHODOLOGICAL_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusColorStrongestTyped       = "CONDITIONAL_SUPPORT_COLOR_THREE_IS_CURRENT_STRONGEST_TYPED_SOURCE_OF_N_EFF_BASELINE"
	StatusTopColorLimit             = "CONDITIONAL_SUPPORT_TOP_COLOR_DOMINANCE_EXPLAINS_N_EFF_TOP_LIMIT_EQUALS_THREE"
	StatusGJHighScaleCandidate      = "CONDITIONAL_SUPPORT_GJ_THREE_IS_STRONG_HIGH_SCALE_DIAGNOSTIC_CANDIDATE"
	StatusD4DeepNativeCandidate     = "CONDITIONAL_SUPPORT_D4_TRIALITY_IS_DEEP_NATIVE_THREE_SOURCE_CANDIDATE"
	StatusA2SU3Motivation           = "CONDITIONAL_SUPPORT_A2_SU3_GEOMETRY_CAN_MOTIVATE_TYPED_CARRIER_SEARCH"
	StatusGenerationCarrierRequired = "CONDITIONAL_SUPPORT_GENERATION_CARRIER_IS_REQUIRED_FOR_NATIVE_FLAVOR_THEOREM"
	StatusK7NativeResonance         = "CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_IS_NATIVE_THREE_RESONANCE"
	StatusProjectiveResonance       = "CONDITIONAL_SUPPORT_PROJECTIVE_1_PLUS_3_IS_STRUCTURAL_THREE_RESONANCE"
	StatusColorTopRemainsStrongest  = "CONDITIONAL_SUPPORT_COLOR_TOP_DOMINANCE_REMAINS_CURRENT_STRONGEST_TYPED_SOURCE"
	StatusD4HighestNativeBranch     = "CONDITIONAL_SUPPORT_D4_TRIALITY_IS_HIGHEST_VALUE_NATIVE_SEARCH_BRANCH"
	StatusLedgerHighestEmpirical    = "CONDITIONAL_SUPPORT_VALIDATED_YUKAWA_LEDGER_IS_HIGHEST_VALUE_EMPIRICAL_TEST_BRANCH"

	StatusNoReadoutNoYukawa        = "FAILED_ROUTE_THREEFOLD_STRUCTURE_WITHOUT_TRACE_READOUT_IS_NOT_YUKAWA_THEOREM"
	StatusNoSectorOpsNoGeneration  = "FAILED_ROUTE_THREEFOLD_STRUCTURE_WITHOUT_SECTOR_OPERATORS_IS_NOT_GENERATION_THEOREM"
	StatusColorNoEigenvalues       = "FAILED_ROUTE_COLOR_THREE_DOES_NOT_DERIVE_YUKAWA_EIGENVALUES"
	StatusColorNoGeneration        = "FAILED_ROUTE_COLOR_THREE_DOES_NOT_DERIVE_GENERATION_STRUCTURE"
	StatusGJNeedsLedger            = "FAILED_ROUTE_GJ_THREE_REQUIRES_MULTISCALE_YUKAWA_LEDGER"
	StatusGJNoNEffWithoutMap       = "FAILED_ROUTE_GJ_THREE_DOES_NOT_DERIVE_N_EFF_WITHOUT_TRACE_READOUT_MAP"
	StatusGJNotGUT                 = "FAILED_ROUTE_GJ_THREE_NOT_NATIVE_GUT_THEOREM"
	StatusNoD4Package              = "FAILED_ROUTE_NO_CERTIFIED_D4_TRIALITY_CARRIER_PACKAGE"
	StatusNoD4TraceMap             = "FAILED_ROUTE_NO_TYPED_D4_TRIALITY_TO_YUKAWA_TRACE_READOUT_MAP"
	StatusCompactSpin8Firewall     = "FAILED_ROUTE_COMPACT_SPIN8_TRIALITY_NOT_AUTOMATICALLY_NATIVE_IN_CL17_REAL_FORM"
	StatusTrialityNotGeneration    = "FAILED_ROUTE_TRIALITY_FRAMES_NOT_YET_GENERATION_THEOREM"
	StatusHexMotifNotEvidence      = "FAILED_ROUTE_HEXAGONAL_VISUAL_MOTIF_NOT_TYPED_EVIDENCE"
	StatusColorNotFlavorSU3        = "FAILED_ROUTE_COLOR_SU3_NOT_AUTOMATICALLY_FLAVOR_SU3"
	StatusNoA2TraceMap             = "FAILED_ROUTE_NO_A2_SU3_TO_YUKAWA_TRACE_READOUT_MAP"
	StatusNoNativeGeneration       = "FAILED_ROUTE_NO_NATIVE_GENERATION_CARRIER_CERTIFIED"
	StatusGenerationsNoNEff        = "FAILED_ROUTE_THREE_GENERATIONS_ALONE_DO_NOT_DERIVE_N_EFF"
	StatusNoPMNSCKM                = "FAILED_ROUTE_NO_PMNS_CKM_READOUT_WITHOUT_SECTOR_OPERATORS"
	StatusK7MinusNotGeneration     = "FAILED_ROUTE_K7_MINUS_DIMENSION_THREE_NOT_GENERATION_THEOREM"
	StatusNoK7YukawaMap            = "FAILED_ROUTE_NO_K7_POLARITY_TO_YUKAWA_TRACE_READOUT_MAP"
	StatusNoK7PMNSCKMMap           = "FAILED_ROUTE_NO_K7_POLARITY_TO_PMNS_CKM_MAP"
	StatusProjectiveNotYukawa      = "FAILED_ROUTE_PROJECTIVE_1_PLUS_3_NOT_YET_YUKAWA_TRACE_THEOREM"
	StatusProjectiveNotMixing      = "FAILED_ROUTE_PROJECTIVE_1_PLUS_3_NOT_YET_GENERATION_MIXING_THEOREM"
	StatusFirewallPreservedGate799 = "FIREWALL_PRESERVED_GATE799_NATIVE_THREE_SOURCE_CANDIDATE_BOUNDARY"
)

const (
	NEffSnapshot    = 3.0023273474722147
	CYukawaSnapshot = 0.9992248188812008
)

type Gate798Inheritance struct {
	Inherited              bool
	NEff                   float64
	CYukawa                float64
	CurrentCertifiedSource string
	NotGenerationTheorem   bool
	NotD4Theorem           bool
	NotNativeYukawaTheorem bool
	GJRequiresLedger       bool
	MotifsNotEvidence      bool
	Verdict                string
}

type NativeThreeSourceRequirement struct {
	Defined     bool
	PackageName string
	Fields      []string
	NEffReadout []string
	GJReadout   []string
	GenReadout  []string
	RejectNoMap bool
	RejectNoOps bool
	Verdict     string
}

type Candidate struct {
	Rank        int
	Name        string
	Audited     bool
	TypedSource string
	Strengths   []string
	Limitations []string
	RequiredMap []string
	Verdict     string
	Support     string
	Failures    []string
}

type CandidateRanking struct {
	Recorded bool
	Ranks    []string
	Verdict  string
}

type MethodologicalBranch struct {
	Recorded          bool
	EmpiricalPath     []string
	NativePath        []string
	ForbiddenPath     string
	Recommended       string
	RecommendationWhy string
	Verdict           string
}

type Firewalls struct {
	Enforced                 bool
	ThreeIsProof             bool
	NEffD4Theorem            bool
	NEffGenerationTheorem    bool
	ColorFullFlavorTheorem   bool
	GJGUTTheorem             bool
	HexagramEvidence         bool
	K7GenerationTheorem      bool
	ProjectiveYukawaTheorem  bool
	D4WithoutRealFormTheorem bool
	CHiggsLevelC             bool
	TreeProxyPoleMass        bool
	Verdict                  string
}

type Analysis struct {
	Gate798     Gate798Inheritance
	Requirement NativeThreeSourceRequirement
	Candidates  []Candidate
	Ranking     CandidateRanking
	Branch      MethodologicalBranch
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func BuildDefault() (Analysis, error) {
	if NEffSnapshot <= 3 || CYukawaSnapshot <= 0 || math.IsNaN(NEffSnapshot) {
		return Analysis{}, fmt.Errorf("invalid Gate 799 inherited N_eff snapshot")
	}
	candidates := []Candidate{
		{
			Rank: 1, Name: "Color SU(3) top dominance", Audited: true,
			TypedSource: "SU(3)_c color multiplicity directly in finite spectral-action a,b trace formulas",
			Strengths:   []string{"a_u=3Tr(Y_u†Y_u)", "b_u=3Tr((Y_u†Y_u)^2)", "single dominant top channel gives a_top=3T, b_top=3T^2, N_eff_top=3"},
			Limitations: []string{"does not derive y_t", "does not explain why top dominates", "does not derive non-top rest pressure", "does not derive generation structure", "does not derive PMNS/CKM", "does not prove Georgi-Jarlskog"},
			RequiredMap: []string{"none for top-color trace limit", "native Yukawa/eigenvalue theorem still required for actual y_t and rest atoms"},
			Verdict:     StatusColorSU3Audited, Support: StatusColorStrongestTyped,
			Failures: []string{StatusColorNoEigenvalues, StatusColorNoGeneration},
		},
		{
			Rank: 2, Name: "External validated Yukawa atom ledger", Audited: true,
			TypedSource: "validated sector/atom data interface for assigning N_eff-3 to actual atoms",
			Strengths:   []string{"empirically testable", "sector-auditable", "can validate or update C_Yukawa"},
			Limitations: []string{"external data not native", "does not prove PMNS/CKM", "does not prove triality"},
			RequiredMap: []string{"ExternalYukawaLedgerConventionSeal", "DecomposedYukawaTraceLedgerSeal", "validation against a,b,N_eff"},
			Verdict:     "PASS_EXTERNAL_VALIDATED_YUKAWA_LEDGER_SOURCE_AUDITED", Support: StatusLedgerHighestEmpirical,
			Failures: []string{"FAILED_ROUTE_VALIDATED_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"},
		},
		{
			Rank: 3, Name: "D4 / Spin(8) triality", Audited: true,
			TypedSource: "future real-form-compatible D4 triality carrier candidate",
			Strengths:   []string{"native threefold outer automorphism candidate", "three 8-dimensional frames candidate", "invariant trilinear coupling candidate"},
			Limitations: []string{"compact Spin(8) triality not automatically native in Cl(1,7)", "triality frames not yet generations", "trilinear form not yet Yukawa trace ledger", "no breaking operator for N_eff-3"},
			RequiredMap: []string{"D4TrialityCarrierPackage", "real-form airlock", "trace-readout map into a,b,N_eff", "sector operator map", "breaking/deformation object"},
			Verdict:     StatusD4Audited, Support: StatusD4DeepNativeCandidate,
			Failures: []string{StatusNoD4Package, StatusNoD4TraceMap, StatusCompactSpin8Firewall, StatusTrialityNotGeneration},
		},
		{
			Rank: 4, Name: "Generation carrier", Audited: true,
			TypedSource: "three observed generations as an empirical/native-carrier target",
			Strengths:   []string{"directly aligned with family structure", "required for PMNS/CKM and flavor orientation"},
			Limitations: []string{"three generations alone do not derive N_eff", "does not derive top dominance", "does not derive D4/triality"},
			RequiredMap: []string{"GenerationCarrierPackage", "sector operators", "diagonalization frames", "PMNS/CKM readout maps", "trace atom map into a,b"},
			Verdict:     StatusGenerationAudited, Support: StatusGenerationCarrierRequired,
			Failures: []string{StatusNoNativeGeneration, StatusGenerationsNoNEff, StatusNoPMNSCKM},
		},
		{
			Rank: 5, Name: "Georgi-Jarlskog Clebsch-three", Audited: true,
			TypedSource: "high-scale down/lepton Clebsch diagnostic",
			Strengths:   []string{"representation/Clebsch factor diagnostic", "connects down quarks and charged leptons", "strong high-scale flavor embedding candidate"},
			Limitations: []string{"requires multi-scale Yukawa ledger", "not available from aggregate a,b", "not a low-scale N_eff theorem", "not native GUT theorem"},
			RequiredMap: []string{"multi-scale Yukawa ledger", "RG/threshold package", "representation convention", "trace-readout map if used for N_eff"},
			Verdict:     StatusGJAudited, Support: StatusGJHighScaleCandidate,
			Failures: []string{StatusGJNeedsLedger, StatusGJNoNEffWithoutMap, StatusGJNotGUT},
		},
		{
			Rank: 6, Name: "SU(3) / A2 hexagonal carrier", Audited: true,
			TypedSource: "A2 root/weight geometry search motif",
			Strengths:   []string{"mathematically lawful carrier search motif", "can motivate color/flavor/boundary representation search"},
			Limitations: []string{"visual motif is not typed evidence", "color SU(3) not automatically flavor SU(3)", "no Yukawa trace-readout map"},
			RequiredMap: []string{"A2SU3CarrierPackage", "root/weight basis", "representation labels", "map into color/flavor/boundary axes", "trace-readout map into a,b,N_eff or GJ"},
			Verdict:     StatusA2SU3Audited, Support: StatusA2SU3Motivation,
			Failures: []string{StatusHexMotifNotEvidence, StatusColorNotFlavorSU3, StatusNoA2TraceMap},
		},
		{
			Rank: 7, Name: "K7 Hodge 4|3 and Fock/projective 1+3", Audited: true,
			TypedSource: "native ASHA structural resonances with a three-sector component",
			Strengths:   []string{"K7 Hodge polarity is native", "projective 1+3 appears in selector/Higgs socket lanes", "strong structural resonance"},
			Limitations: []string{"K7^- dimension three is not a generation carrier", "no K7 to Yukawa trace readout", "projective 1+3 is not a Yukawa operator", "no PMNS/CKM readout"},
			RequiredMap: []string{"K7 polarity to Yukawa trace readout", "K7 polarity to PMNS/CKM map", "projective selector to generation/mixing map"},
			Verdict:     StatusK7HodgeAudited + "; " + StatusFockProjectiveAudited, Support: StatusK7NativeResonance + "; " + StatusProjectiveResonance,
			Failures: []string{StatusK7MinusNotGeneration, StatusNoK7YukawaMap, StatusNoK7PMNSCKMMap, StatusProjectiveNotYukawa, StatusProjectiveNotMixing},
		},
	}
	return Analysis{
		Gate798: Gate798Inheritance{
			Inherited: true, NEff: NEffSnapshot, CYukawa: CYukawaSnapshot,
			CurrentCertifiedSource: "color-tripled top dominance",
			NotGenerationTheorem:   true, NotD4Theorem: true, NotNativeYukawaTheorem: true,
			GJRequiresLedger: true, MotifsNotEvidence: true, Verdict: StatusGate798Inherited,
		},
		Requirement: NativeThreeSourceRequirement{
			Defined: true, PackageName: "NativeThreeSourcePackage",
			Fields:      []string{"typed carrier", "symmetry or representation action", "distinguished threefold structure", "trace/readout map", "breaking or deformation object", "scale/convention airlock", "noncircularity proof"},
			NEffReadout: []string{"a=sum_i x_i", "b=sum_i x_i^2", "N_eff=a^2/b"},
			GJReadout:   []string{"R_GJ_3=y_b/y_tau", "R_GJ_2=y_mu/(3y_s)", "R_GJ_1=3y_e/y_d"},
			GenReadout:  []string{"generation carrier", "sector operators", "PMNS/CKM readout maps"},
			RejectNoMap: true, RejectNoOps: true, Verdict: StatusPackageRequirements,
		},
		Candidates: candidates,
		Ranking: CandidateRanking{Recorded: true, Ranks: []string{
			"1 Color SU(3) top dominance — current strongest typed source for N_eff baseline",
			"2 External validated Yukawa atom ledger — highest empirical test branch",
			"3 D4 / Spin(8) triality — highest-value native search branch, not certified",
			"4 Generation carrier — required for flavor/mixing theorem",
			"5 Georgi-Jarlskog Clebsch-three — strong high-scale diagnostic",
			"6 SU(3)/A2 hexagonal carrier — lawful motif for typed carrier search",
			"7 K7 Hodge 4|3 and Fock/projective 1+3 — native resonances without readout",
		}, Verdict: StatusCandidatesRanked},
		Branch: MethodologicalBranch{Recorded: true,
			EmpiricalPath:     []string{"acquire external Yukawa atom ledger", "validate N_eff", "audit sector contributions", "run GJ/FN/Koide diagnostics"},
			NativePath:        []string{"construct D4TrialityCarrierPackage or GenerationCarrierPackage", "prove trace-readout map", "test whether it sources N_eff or flavor orientation"},
			ForbiddenPath:     "use symbolic threefold motifs to bypass trace data or carrier construction",
			Recommended:       "Gate 800 — D4 Triality Carrier Package Requirement and Cl(1,7) Real-Form Audit",
			RecommendationWhy: "test whether D4 triality can be typed inside the actual Cl(1,7) board before using it for Yukawa traces",
			Verdict:           StatusMethodologicalBranch,
		},
		Firewalls: Firewalls{Enforced: true, Verdict: StatusFirewallPreservedGate799},
		Truth:     "Gate 799 ranks three-source candidates without promoting any motif into a theorem; color-tripled top dominance remains the strongest certified source of the N_eff≈3 baseline.",
		Final:     "Gate 799 does not prove any native three-source theorem. It ranks the current candidates honestly: the strongest certified source of N_eff≈3 remains color-tripled top dominance; the highest-value empirical path is a validated Yukawa atom ledger; the highest-value native path is a D4 triality carrier audit inside the actual Cl(1,7) real-form board. The next native gate should test whether D4 triality can even be typed lawfully before any Yukawa trace readout is attempted.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate798Inherited, StatusCurrentThreeInherited, StatusPackageRequirements, StatusColorSU3Audited, StatusGJAudited, StatusD4Audited, StatusA2SU3Audited, StatusGenerationAudited, StatusK7HodgeAudited, StatusFockProjectiveAudited, StatusCandidatesRanked, StatusMethodologicalBranch, StatusPhysicalFirewalls,
		StatusColorStrongestTyped, StatusTopColorLimit, StatusGJHighScaleCandidate, StatusD4DeepNativeCandidate, StatusA2SU3Motivation, StatusGenerationCarrierRequired, StatusK7NativeResonance, StatusProjectiveResonance, StatusColorTopRemainsStrongest, StatusD4HighestNativeBranch, StatusLedgerHighestEmpirical,
		StatusNoReadoutNoYukawa, StatusNoSectorOpsNoGeneration, StatusColorNoEigenvalues, StatusColorNoGeneration, StatusGJNeedsLedger, StatusGJNoNEffWithoutMap, StatusGJNotGUT, StatusNoD4Package, StatusNoD4TraceMap, StatusCompactSpin8Firewall, StatusTrialityNotGeneration, StatusHexMotifNotEvidence, StatusColorNotFlavorSU3, StatusNoA2TraceMap, StatusNoNativeGeneration, StatusGenerationsNoNEff, StatusNoPMNSCKM, StatusK7MinusNotGeneration, StatusNoK7YukawaMap, StatusNoK7PMNSCKMMap, StatusProjectiveNotYukawa, StatusProjectiveNotMixing, StatusFirewallPreservedGate799,
	}
}

func CandidateByName(candidates []Candidate, namePart string) (Candidate, bool) {
	for _, c := range candidates {
		if strings.Contains(c.Name, namePart) {
			return c, true
		}
	}
	return Candidate{}, false
}

func FormatRequirement(r NativeThreeSourceRequirement) string {
	return r.PackageName + ": fields=" + strings.Join(r.Fields, ", ") + "; N_eff=" + strings.Join(r.NEffReadout, ", ")
}
func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("rank %d %s: %s; required=%s", c.Rank, c.Name, c.TypedSource, strings.Join(c.RequiredMap, ", "))
}
func FormatRanking(r CandidateRanking) string { return strings.Join(r.Ranks, " | ") }
func FormatBranch(b MethodologicalBranch) string {
	return b.Recommended + "; why=" + b.RecommendationWhy + "; forbidden=" + b.ForbiddenPath
}
func closeAbs(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}
func containsAll(hay []string, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range hay {
			if strings.Contains(h, n) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
