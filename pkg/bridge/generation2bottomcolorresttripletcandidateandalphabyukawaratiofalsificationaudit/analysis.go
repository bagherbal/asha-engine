// Package generation2bottomcolorresttripletcandidateandalphabyukawaratiofalsificationaudit implements
// Gate 820: BottomColor RestTriplet Candidate and AlphaB Yukawa-Ratio Falsification Audit.
//
// Gate 820 tests whether Gate 819's 1+3 rest simplex can be read as one tiny
// rest dust atom plus one colored rest triplet, without assuming whether that
// triplet is bottom-like, charm-like, or only abstract.
package generation2bottomcolorresttripletcandidateandalphabyukawaratiofalsificationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE820-BOTTOM-COLOR-REST-TRIPLET-CANDIDATE-ALPHAB-YUKAWA-RATIO-FALSIFICATION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate819Inherited   = "PASS_GATE819_ONE_PLUS_THREE_SIMPLEX_INHERITED"
	StatusRestTripletDefined = "PASS_REST_TRIPLET_INTERPRETATION_DEFINED"
	StatusBOverTComputed     = "PASS_B_OVER_T_ALPHA_RATIO_COMPUTED"
	StatusDustScaleComputed  = "PASS_DUST_SCALE_COMPUTED"
	StatusBottomAudited      = "PASS_BOTTOM_COLOR_TRIPLET_CANDIDATE_AUDITED"
	StatusCharmAudited       = "PASS_CHARM_COLOR_TRIPLET_CANDIDATE_AUDITED"
	StatusAbstractAudited    = "PASS_ABSTRACT_COLORED_TRIPLET_CANDIDATE_AUDITED"
	StatusExternalProtocol   = "PASS_EXTERNAL_LEDGER_FALSIFICATION_TEST_DEFINED"
	StatusNativeSourceAudit  = "PASS_NATIVE_SOURCE_AUDIT_DEFINED"
	StatusClassification     = "PASS_STATUS_CLASSIFICATION_DEFINED"
	StatusImpactRecorded     = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusPhysicalFirewalls  = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	SupportSimplexAsTriplet      = "CONDITIONAL_SUPPORT_ONE_PLUS_THREE_SIMPLEX_CAN_BE_READ_AS_TINY_DUST_PLUS_COLORED_REST_TRIPLET"
	SupportAlphaPredictsBOverT   = "CONDITIONAL_SUPPORT_ALPHA_B_PREDICTS_COLORED_REST_TO_TOP_RATIO_B_OVER_T"
	SupportSqrtRatioSharpTest    = "CONDITIONAL_SUPPORT_SQRT_ALPHA_B_RATIO_IS_A_SHARP_EXTERNAL_LEDGER_TEST"
	SupportBottomSeriousIfLedger = "CONDITIONAL_SUPPORT_BOTTOM_LIKE_TRIPLET_IS_A_SERIOUS_CANDIDATE_IF_LEDGER_SUPPORTS_IT"
	SupportGJAfterBottom         = "CONDITIONAL_SUPPORT_GJ_BTAU_LANE_BECOMES_RELEVANT_ONLY_AFTER_BOTTOM_LIKE_IDENTIFICATION"
	SupportExternalFalsifies     = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_FALSIFY_THE_COLORED_TRIPLET_READING"
	SupportStrengthenedPartialR2 = "CONDITIONAL_SUPPORT_EXPECTED_STATUS_IS_STRENGTHENED_PARTIAL_R2_WITH_SHARPER_TRIPLET_PREDICTION"

	FailureTripletNotNative     = "FAILED_ROUTE_SIMPLEX_TRIPLET_READING_NOT_NATIVE_WITHOUT_TRACE_ATOM_MAP"
	FailureBottomNotAllowed     = "FAILED_ROUTE_BOTTOM_COLOR_IDENTIFICATION_NOT_ALLOWED_WITHOUT_LEDGER_OR_NATIVE_OPERATOR"
	FailureCharmNotAllowed      = "FAILED_ROUTE_CHARM_COLOR_IDENTIFICATION_NOT_ALLOWED_WITHOUT_LEDGER_OR_NATIVE_OPERATOR"
	FailureProjectiveNotTheorem = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_YUKAWA_TRACE_THEOREM"
	FailureK7NotTriplet         = "FAILED_ROUTE_K7_4_3_NOT_COLORED_REST_TRIPLET_THEOREM"
	FailureAlphaNotRatioTheorem = "FAILED_ROUTE_BOUNDARY_ALPHA_B_NOT_YUKAWA_RATIO_THEOREM"
	FailureAbstractNoSector     = "FAILED_ROUTE_ABSTRACT_DUST_PLUS_TRIPLET_DOES_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	FailureExternalNotNative    = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	FailureNoUpdateCYukawa      = "FAILED_ROUTE_GATE820_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFIED_TRACE_MAGNITUDE_MAP_OR_VALIDATED_EXTERNAL_LEDGER"
	FailureCHiggsLevelB         = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureTreeProxyNotPole     = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate820       = "FIREWALL_PRESERVED_GATE820_BOTTOM_COLOR_REST_TRIPLET_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2          float64
	AlphaB, NEffBFN, QSimplex       float64
	BOverT, SqrtBOverT              float64
	DustOverT, SqrtDustOverT        float64
	CYukawaSimplex, CHiggsSimplex   float64
	OfficialCYukawa, OfficialCHiggs float64
	Verdicts, Supports, Failures    []string
}

type RestTripletInterpretation struct {
	Assumption                   string
	Equations                    []string
	DiagnosticOnly               bool
	Verdicts, Supports, Failures []string
}

type SectorCandidate struct {
	Name                         string
	CandidateRole                string
	Prediction                   string
	RequiresLedger               bool
	AllowedToIdentifyNow         bool
	Verdicts, Supports, Failures []string
}

type ExternalLedgerProtocol struct {
	RequiredObjects              []string
	Tests                        []string
	Forbidden                    []string
	CanFalsify                   bool
	CanUpgradeExternalR3         bool
	Verdicts, Supports, Failures []string
}

type NativeSourceAudit struct {
	Lane                         string
	Supplies                     []string
	DoesNotSupply                []string
	Verdicts, Supports, Failures []string
}

type Status struct {
	Outcome                      string
	Level                        string
	NativeSourceFound            bool
	ExternalLedgerSupplied       bool
	CanUpdateCYukawa             bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type BranchDecision struct {
	NextGate                     string
	Reason                       string
	Verdicts, Supports, Failures []string
}

type Firewalls struct {
	Enforced                                                                                                      bool
	TripletNotNative, BottomNotAllowed, CharmNotAllowed, ProjectiveNotTheorem, K7NotTriplet, AlphaNotRatioTheorem bool
	AbstractNotSector, ExternalNotNative, NoCYukawaUpdate, CHiggsLevelB, TreeProxyNotPole                         bool
	Verdict                                                                                                       string
}

type Analysis struct {
	Ledger         Ledger
	Interpretation RestTripletInterpretation
	Candidates     []SectorCandidate
	Protocol       ExternalLedgerProtocol
	NativeSources  []NativeSourceAudit
	Status         Status
	Impact         Impact
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func NEffBFN(s float64) float64            { return 3.0 + 6.0*AlphaB(s) }
func QSimplex(t float64) float64           { return t*t + math.Pow(1.0-t, 2)/3.0 }
func BOverT(alpha float64) float64         { return alpha * (1.0 - alpha) }
func DustOverT(alpha float64) float64      { return 3.0 * alpha * alpha }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	nEffBFN := NEffBFN(SBoundary)
	qSimplex := QSimplex(alpha)
	bOverT := BOverT(alpha)
	dustOverT := DustOverT(alpha)
	if bOverT <= 0 || dustOverT <= 0 {
		return Analysis{}, fmt.Errorf("invalid positive triplet/dust scales: B/T=%g D/T=%g", bOverT, dustOverT)
	}
	ledger := Ledger{
		NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2,
		AlphaB: alpha, NEffBFN: nEffBFN, QSimplex: qSimplex,
		BOverT: bOverT, SqrtBOverT: math.Sqrt(bOverT),
		DustOverT: dustOverT, SqrtDustOverT: math.Sqrt(dustOverT),
		CYukawaSimplex: CYukawaFromNEff(nEffBFN), CHiggsSimplex: CHiggsFromNEff(nEffBFN),
		OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		Verdicts: []string{StatusGate819Inherited, StatusBOverTComputed, StatusDustScaleComputed},
		Supports: []string{SupportAlphaPredictsBOverT, SupportSqrtRatioSharpTest},
		Failures: []string{FailureAlphaNotRatioTheorem},
	}

	interpretation := RestTripletInterpretation{
		Assumption:     "audit-only reading: one tiny dust component plus one colored rest triplet",
		Equations:      []string{"a_rest = 3T alpha_B", "B/T = alpha_B(1-alpha_B)", "D/T = 3 alpha_B^2", "sqrt(B/T) is the candidate Yukawa ratio only after a trace ledger identifies the triplet"},
		DiagnosticOnly: true,
		Verdicts:       []string{StatusRestTripletDefined},
		Supports:       []string{SupportSimplexAsTriplet, SupportAlphaPredictsBOverT},
		Failures:       []string{FailureTripletNotNative, FailureAbstractNoSector},
	}

	candidates := []SectorCandidate{
		{Name: "Candidate A — bottom-color triplet", CandidateRole: "triplet rest block = bottom-color triplet", Prediction: "y_b^2/y_t^2 ≈ alpha_B(1-alpha_B); y_b/y_t ≈ sqrt(alpha_B(1-alpha_B))", RequiresLedger: true, AllowedToIdentifyNow: false, Verdicts: []string{StatusBottomAudited}, Supports: []string{SupportBottomSeriousIfLedger, SupportGJAfterBottom}, Failures: []string{FailureBottomNotAllowed}},
		{Name: "Candidate B — charm-color triplet", CandidateRole: "triplet rest block = charm-color triplet", Prediction: "y_c^2/y_t^2 ≈ alpha_B(1-alpha_B); y_c/y_t ≈ sqrt(alpha_B(1-alpha_B))", RequiresLedger: true, AllowedToIdentifyNow: false, Verdicts: []string{StatusCharmAudited}, Failures: []string{FailureCharmNotAllowed}},
		{Name: "Candidate C — abstract colored rest chamber", CandidateRole: "ledger should contain one color-tripled rest atom with B/T ≈ alpha_B(1-alpha_B)", Prediction: "abstract colored rest triplet + dust; no SM sector assignment", RequiresLedger: true, AllowedToIdentifyNow: false, Verdicts: []string{StatusAbstractAudited}, Supports: []string{SupportSimplexAsTriplet}, Failures: []string{FailureAbstractNoSector}},
		{Name: "Candidate D — no sector triplet matches", CandidateRole: "simplex remains abstract or downgrades under ledger failure", Prediction: "if no colored triplet/dust pattern survives, the branch downgrades toward R1 scalar closure", RequiresLedger: true, AllowedToIdentifyNow: true, Verdicts: []string{StatusClassification}, Failures: []string{FailureTripletNotNative}},
	}

	protocol := ExternalLedgerProtocol{
		RequiredObjects: []string{"typed top-like trace atom T=h_t", "colored non-top candidates f in {b,c,s,u,d}", "B_f=h_f and ratio_f=B_f/T", "a_rest,b_rest,alpha_ext,beta_ext,q_rest_ext", "scale/scheme/color/neutrino conventions"},
		Tests:           []string{"T1: one colored rest triplet has B_f/T near alpha_B(1-alpha_B)", "T2: remaining dust D_ext/T near 3 alpha_B^2", "T3: q_rest_ext matches q_simplex(alpha_B)", "T4: selected triplet is largest non-top colored block without forced selection", "T5: result is stable under declared scale/scheme convention"},
		Forbidden:       []string{"choose triplet after rescaling atoms", "merge unrelated atoms to fake color triplet", "discard atoms to force D_ext", "use Higgs mass or C_Higgs to tune ledger", "retune alpha_B, 9/5, or 6 after seeing data"},
		CanFalsify:      true, CanUpgradeExternalR3: true,
		Verdicts: []string{StatusExternalProtocol},
		Supports: []string{SupportExternalFalsifies, SupportSqrtRatioSharpTest},
		Failures: []string{FailureExternalNotNative},
	}

	nativeSources := []NativeSourceAudit{
		{Lane: "finite spectral triple", Supplies: []string{"color multiplicity", "sector edge templates"}, DoesNotSupply: []string{"B/T", "bottom/charm trace atom", "Yukawa hierarchy"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureTripletNotNative}},
		{Lane: "projective/Fock 1+3", Supplies: []string{"structural one-plus-three resonance"}, DoesNotSupply: []string{"colored Yukawa trace atom", "bottom/charm identification"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureProjectiveNotTheorem}},
		{Lane: "K7 4|3", Supplies: []string{"carrier resonance"}, DoesNotSupply: []string{"colored rest triplet theorem", "bottom/charm trace magnitude"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureK7NotTriplet}},
		{Lane: "boundary alpha_B", Supplies: []string{"small dust/rest-size parameter"}, DoesNotSupply: []string{"triplet sector identification", "Yukawa ratio theorem"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureAlphaNotRatioTheorem}},
		{Lane: "Georgi-Jarlskog", Supplies: []string{"future high-scale down/lepton diagnostic if bottom-like triplet is ledger-supported"}, DoesNotSupply: []string{"low-scale proof", "triplet atom"}, Verdicts: []string{StatusNativeSourceAudit}, Supports: []string{SupportGJAfterBottom}, Failures: []string{FailureTripletNotNative}},
		{Lane: "D4/triality", Supplies: []string{"airlocked structural search geometry"}, DoesNotSupply: []string{"trace atom", "bottom/charm triplet"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureTripletNotNative}},
	}

	status := Status{Outcome: "Outcome C — only abstract simplex survives without external ledger", Level: "strengthened partial R2 with sharper falsifiable colored-triplet prediction; not external R3; not native R4", NativeSourceFound: false, ExternalLedgerSupplied: false, CanUpdateCYukawa: false, Verdicts: []string{StatusClassification}, Supports: []string{SupportStrengthenedPartialR2}, Failures: []string{FailureTripletNotNative, FailureExternalNotNative}}

	impact := Impact{CandidateNEff: nEffBFN, CandidateCYukawa: CYukawaFromNEff(nEffBFN), CandidateCHiggs: CHiggsFromNEff(nEffBFN), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Failures: []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole}}

	branch := BranchDecision{NextGate: "Gate 821 — Colored RestTriplet External Ledger Checklist and Bottom/Charm Discriminator Audit", Reason: "Freeze the B/T and dust predictions, then specify the ledger intake needed to distinguish bottom-like, charm-like, abstract-colored, or failed triplet readings.", Verdicts: []string{StatusClassification}, Supports: []string{SupportExternalFalsifies}, Failures: []string{FailureTripletNotNative}}

	firewalls := Firewalls{Enforced: true, TripletNotNative: true, BottomNotAllowed: true, CharmNotAllowed: true, ProjectiveNotTheorem: true, K7NotTriplet: true, AlphaNotRatioTheorem: true, AbstractNotSector: true, ExternalNotNative: true, NoCYukawaUpdate: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate820}

	truth := "Gate 820 turns the 1+3 simplex into a sharper falsifiable colored-rest-triplet prediction, but no native trace atom map or external ledger identifies the triplet."
	final := "The simplex predicts B/T = alpha_B(1-alpha_B) and sqrt(B/T) ≈ 0.01969125 for any colored rest-triplet interpretation; bottom-like and charm-like readings remain forbidden without a validated trace ledger or native operator."

	return Analysis{Ledger: ledger, Interpretation: interpretation, Candidates: candidates, Protocol: protocol, NativeSources: nativeSources, Status: status, Impact: impact, Branch: branch, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func Statuses() []string {
	return []string{StatusGate819Inherited, StatusRestTripletDefined, StatusBOverTComputed, StatusDustScaleComputed, StatusBottomAudited, StatusCharmAudited, StatusAbstractAudited, StatusExternalProtocol, StatusNativeSourceAudit, StatusClassification, StatusImpactRecorded, StatusPhysicalFirewalls, SupportSimplexAsTriplet, SupportAlphaPredictsBOverT, SupportSqrtRatioSharpTest, SupportBottomSeriousIfLedger, SupportGJAfterBottom, SupportExternalFalsifies, SupportStrengthenedPartialR2, FailureTripletNotNative, FailureBottomNotAllowed, FailureCharmNotAllowed, FailureProjectiveNotTheorem, FailureK7NotTriplet, FailureAlphaNotRatioTheorem, FailureAbstractNoSector, FailureExternalNotNative, FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole, StatusFirewallGate820}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g N_eff_BFN=%.16g q_simplex=%.16g B/T=%.16g sqrt(B/T)=%.16g D/T=%.16g sqrt(D/T)=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.NEffBFN, a.QSimplex, a.BOverT, a.SqrtBOverT, a.DustOverT, a.SqrtDustOverT)
}

func FormatInterpretation(a RestTripletInterpretation) string {
	return fmt.Sprintf("%s diagnosticOnly=%t equations=[%s]", a.Assumption, a.DiagnosticOnly, strings.Join(a.Equations, "; "))
}

func FormatCandidates(rows []SectorCandidate) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s role=%s prediction=%s requiresLedger=%t identifyNow=%t", r.Name, r.CandidateRole, r.Prediction, r.RequiresLedger, r.AllowedToIdentifyNow))
	}
	return strings.Join(out, " | ")
}

func FormatProtocol(a ExternalLedgerProtocol) string {
	return fmt.Sprintf("required=[%s] tests=[%s] forbidden=[%s] canFalsify=%t externalR3=%t", strings.Join(a.RequiredObjects, "; "), strings.Join(a.Tests, "; "), strings.Join(a.Forbidden, "; "), a.CanFalsify, a.CanUpgradeExternalR3)
}

func FormatNativeSources(rows []NativeSourceAudit) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s supplies=[%s] doesNotSupply=[%s]", r.Lane, strings.Join(r.Supplies, ","), strings.Join(r.DoesNotSupply, ",")))
	}
	return strings.Join(out, " | ")
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.CandidateNEff, a.CandidateCYukawa, a.CandidateCHiggs, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range hay {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}
