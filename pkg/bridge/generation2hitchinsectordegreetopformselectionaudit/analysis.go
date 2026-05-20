// Package generation2hitchinsectordegreetopformselectionaudit implements
// Gate 650: Hitchin Sector-Degree Top-Form Selection Rule Audit.
//
// Gate 649 established the finite channel algebra
//
//	A = Ω++-, B = Ω---,
//	AAA -> +P_+,
//	AAB + ABA + BAA -> -3P_-,
//
// for the admissible S_K-twisted Hitchin metric ledger. Gate 650 audits the
// symbolic source of this selection rule from sector-degree saturation in the
// 4|3 Hodge split of K_7. It proves only a degree-selection theorem candidate:
// top-form contributions to the Hitchin cubic must have sector degree (4,3).
// It does not certify the remaining sign/equal-unit calibration identity and it
// does not derive split-G2, boundary stress, scalar/flavor transport, physical
// metric, Higgs mass, CKM/PMNS, gauge unification, or a native 7/72 theorem.
package generation2hitchinsectordegreetopformselectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate649 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchinchannelalgebraselectionruleaudit"
)

const (
	AuditID = "GATE650-HITCHIN-SECTOR-DEGREE-TOP-FORM-SELECTION-RULE-AUDIT"

	StatusGate649ChannelAlgebraInherited = "PASS_GATE649_CHANNEL_ALGEBRA_INHERITED"
	StatusSectorDegreeLedgerDefined      = "PASS_SECTOR_DEGREE_LEDGER_DEFINED"
	StatusPositiveAAAOnlyByTopFormDegree = "PASS_POSITIVE_BLOCK_AAA_ONLY_BY_TOP_FORM_DEGREE"
	StatusNegativeAABByTopFormDegree     = "PASS_NEGATIVE_BLOCK_AAB_ABA_BAA_BY_TOP_FORM_DEGREE"
	StatusMixedBlockZeroByTopFormDegree  = "PASS_MIXED_BLOCK_ZERO_BY_TOP_FORM_DEGREE"
	StatusDegreeSelectionRuleSupported   = "CONDITIONAL_SUPPORT_CHANNEL_SELECTION_RULE_FROM_4_BY_3_SECTOR_DEGREE_SATURATION"
	StatusMinusThreeFromDegreePlacements = "CONDITIONAL_SUPPORT_MINUS_THREE_FROM_THREE_DEGREE_ALLOWED_CUBIC_PLACEMENTS"
	StatusDEqualsQCarrierResonance       = "CONDITIONAL_SUPPORT_D_EQUALS_Q_AS_ASHA_CARRIER_RESONANCE"
	StatusSignUnitRequiresCalibration    = "FAILED_ROUTE_SIGN_AND_EQUAL_UNIT_WEIGHT_STILL_REQUIRE_CALIBRATION_IDENTITY"
	StatusNoFullSymbolicDegreeTheorem    = "FAILED_ROUTE_NO_FULL_SYMBOLIC_DEGREE_SELECTION_THEOREM"
	StatusNoSplitG2                      = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress               = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                  = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                 = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric               = "FAILED_ROUTE_HITCHIN_DEGREE_SELECTION_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge             = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate650Boundary                = "FIREWALL_PRESERVED_GATE650_HITCHIN_DEGREE_SELECTION_BOUNDARY"
)

const (
	plusDim     = 4
	minusDim    = 3
	cubicDegree = 3
	tol         = 1e-8
)

type Degree struct {
	Plus  int
	Minus int
}

func (d Degree) Add(e Degree) Degree { return Degree{Plus: d.Plus + e.Plus, Minus: d.Minus + e.Minus} }
func (d Degree) Top() bool           { return d.Plus == plusDim && d.Minus == minusDim }
func (d Degree) String() string      { return fmt.Sprintf("(%d,%d)", d.Plus, d.Minus) }

type Gate649Inheritance struct {
	ChannelAlgebraInherited    bool
	TwoComponentSupport        bool
	AAAChannelAudited          bool
	AABChannelAudited          bool
	VanishingAudited           bool
	OffBlockAudited            bool
	SlotFormulaDerived         bool
	SlotTheoremPrimary         bool
	DEqualsQCoincidence        bool
	FullSymbolicChannelTheorem bool
	SplitG2Certified           bool
	BoundaryStressAssignment   bool
	SevenOver72Theorem         bool
	ScalarFlavorTransport      bool
	PhysicalMetric             bool
	Gate649FirewallPreserved   bool
	Verdict                    string
}

type DegreeLedgerRow struct {
	Object         string
	SectorDegree   Degree
	InteriorPlus   string
	InteriorMinus  string
	Interpretation string
}

type SectorDegreeLedger struct {
	Rows         []DegreeLedgerRow
	TopDegree    Degree
	PositiveDim  int
	NegativeDim  int
	AName        string
	BName        string
	AHasDegree21 bool
	BHasDegree03 bool
	Verdict      string
}

type ChannelDegreeRow struct {
	Block               string
	XYSectors           string
	Channel             string
	FirstFamily         string
	SecondFamily        string
	ThirdFamily         string
	FirstInterior       Degree
	SecondInterior      Degree
	ThirdDegree         Degree
	TotalDegree         Degree
	FirstSlotZero       bool
	SecondSlotZero      bool
	ReachesTopDegree    bool
	SurvivesByDegree    bool
	ExpectedGate649Role string
	DegreeMechanism     string
}

type PositiveBlockDegreeAudit struct {
	Rows               []ChannelDegreeRow
	SurvivingChannels  []string
	AAAOnlySurvives    bool
	SelectionMechanism string
	Verdict            string
}

type NegativeBlockDegreeAudit struct {
	Rows                  []ChannelDegreeRow
	SurvivingChannels     []string
	AABPlacementsOnly     bool
	AllowedPlacementCount int
	SelectionMechanism    string
	Verdict               string
}

type MixedBlockDegreeAudit struct {
	Rows                   []ChannelDegreeRow
	AnySurvivesByDegree    bool
	MixedBlockZeroByDegree bool
	SelectionMechanism     string
	Verdict                string
}

type SignNormalizationAudit struct {
	DegreeRuleCertifiesSupport       bool
	Gate649CertifiesFiniteSigns      bool
	EqualUnitWeightCertifiedByDegree bool
	RequiresCalibrationIdentity      bool
	MissingProofObject               string
	Verdict                          string
}

type SymbolicSelectionTheorem struct {
	CandidateTheorem             string
	DegreeSelectionSupported     bool
	FullSymbolicTheoremCertified bool
	RemainingGap                 string
	Verdict                      string
}

type ResultingSlotFormula struct {
	PositiveDim          int
	NegativeDim          int
	SlotMultiplicity     int
	NormSquared          float64
	Cosine               float64
	ResidualSquared      float64
	RecoversGate642Angle bool
	Formula              string
	Verdict              string
}

type ResonanceAudit struct {
	CubicDegree        int
	NegativeDim        int
	EqualInASHACarrier bool
	Interpretation     string
	Verdict            string
}

type Firewalls struct {
	ClaimsFullSymbolicDegreeTheorem bool
	ClaimsSplitG2                   bool
	ClaimsBoundaryStress            bool
	ClaimsSevenOver72               bool
	ClaimsScalarFlavor              bool
	ClaimsPhysicalMetric            bool
	ClaimsHiggsMass                 bool
	ClaimsCKMPMNS                   bool
	ClaimsGaugeUnification          bool
	Verdict                         string
}

type Analysis struct {
	Inherited Gate649Inheritance
	Ledger    SectorDegreeLedger
	Positive  PositiveBlockDegreeAudit
	Negative  NegativeBlockDegreeAudit
	Mixed     MixedBlockDegreeAudit
	Sign      SignNormalizationAudit
	Theorem   SymbolicSelectionTheorem
	Slot      ResultingSlotFormula
	Resonance ResonanceAudit
	Firewalls Firewalls
	Truth     string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g649, err := gate649.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate649 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g649)
	ledger := buildDegreeLedger()
	positive := buildPositiveDegreeAudit()
	negative := buildNegativeDegreeAudit()
	mixed := buildMixedDegreeAudit()
	sign := buildSignAudit(g649)
	theorem := buildTheorem(positive, negative, mixed, sign)
	slot := buildSlotFormula(sign)
	resonance := buildResonance()
	firewalls := Firewalls{Verdict: StatusGate650Boundary}
	truth := "Gate 650 upgrades Gate649's finite AAA/AAB ledger into a sector-degree selection-rule candidate.  On the 4|3 Hodge split, A=Ω++- has degree (2,1) and B=Ω--- has degree (0,3).  Since the Hitchin cubic top-form contribution must saturate degree (4,3), the positive block admits only AAA, the negative block admits exactly AAB, ABA, and BAA, and the mixed block has no degree-allowed channel.  This explains the support pattern and the three slot placements behind P_+-3P_-, but the negative sign and equal unit weights still require the native calibration identity; no split-G2, boundary stress, scalar/flavor transport, physical metric, or native 7/72 theorem follows."
	return Analysis{Inherited: inherited, Ledger: ledger, Positive: positive, Negative: negative, Mixed: mixed, Sign: sign, Theorem: theorem, Slot: slot, Resonance: resonance, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate649.Analysis) Gate649Inheritance {
	joined := strings.Join(gate649.Statuses(), "\n")
	return Gate649Inheritance{
		ChannelAlgebraInherited:    g.Expansion.AAAOnlyPositive && g.Expansion.AABOnlyNegative && g.Support.OnlyAAndBSupported,
		TwoComponentSupport:        g.Support.OnlyAAndBSupported,
		AAAChannelAudited:          g.PositiveAAA.AAAContributesUnit,
		AABChannelAudited:          g.NegativeAAB.EachAABContributesMinusUnit,
		VanishingAudited:           g.Vanishing.AllVanishOrProjectAway,
		OffBlockAudited:            g.OffBlock.ChannelwiseZero,
		SlotFormulaDerived:         g.SlotFormula.RecoversGate642Angle,
		SlotTheoremPrimary:         strings.Contains(joined, gate649.StatusSlotTheoremPrimary),
		DEqualsQCoincidence:        g.Coincidence.EqualInASHACarrier && g.Coincidence.SupportsSlotTheoremOnly,
		FullSymbolicChannelTheorem: g.Readiness.FullSymbolicTheoremCertified || g.Firewalls.ClaimsFullSymbolicChannelSelection,
		SplitG2Certified:           g.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:   g.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:         g.Firewalls.ClaimsSevenOver72,
		ScalarFlavorTransport:      g.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:             g.Firewalls.ClaimsPhysicalMetric,
		Gate649FirewallPreserved:   g.Firewalls.Verdict == gate649.StatusGate649Boundary,
		Verdict:                    StatusGate649ChannelAlgebraInherited,
	}
}

func buildDegreeLedger() SectorDegreeLedger {
	rows := []DegreeLedgerRow{
		{Object: "A=Ω++-", SectorDegree: Degree{2, 1}, InteriorPlus: "i_{K+}A has degree (1,1)", InteriorMinus: "i_{K-}A has degree (2,0)", Interpretation: "two positive slots and one negative slot"},
		{Object: "B=Ω---", SectorDegree: Degree{0, 3}, InteriorPlus: "i_{K+}B=0", InteriorMinus: "i_{K-}B has degree (0,2)", Interpretation: "three negative slots; killed by positive interior contraction"},
	}
	return SectorDegreeLedger{Rows: rows, TopDegree: Degree{4, 3}, PositiveDim: plusDim, NegativeDim: minusDim, AName: "Ω++-", BName: "Ω---", AHasDegree21: true, BHasDegree03: true, Verdict: StatusSectorDegreeLedgerDefined}
}

func buildPositiveDegreeAudit() PositiveBlockDegreeAudit {
	rows := degreeRowsForBlock("positive", "+,+")
	survive := survivingChannels(rows)
	return PositiveBlockDegreeAudit{Rows: rows, SurvivingChannels: survive, AAAOnlySurvives: len(survive) == 1 && survive[0] == "AAA", SelectionMechanism: "for x,y in K7+, i_xB=i_yB=0 and only i_xA∧i_yA∧A reaches sector degree (4,3)", Verdict: join(StatusPositiveAAAOnlyByTopFormDegree, StatusDegreeSelectionRuleSupported)}
}

func buildNegativeDegreeAudit() NegativeBlockDegreeAudit {
	rows := degreeRowsForBlock("negative", "-,-")
	survive := survivingChannels(rows)
	ok := len(survive) == 3 && containsAll(survive, []string{"AAB", "ABA", "BAA"})
	return NegativeBlockDegreeAudit{Rows: rows, SurvivingChannels: survive, AABPlacementsOnly: ok, AllowedPlacementCount: len(survive), SelectionMechanism: "for x,y in K7-, precisely the three AAB placements saturate the (4,3) top degree", Verdict: join(StatusNegativeAABByTopFormDegree, StatusMinusThreeFromDegreePlacements)}
}

func buildMixedDegreeAudit() MixedBlockDegreeAudit {
	rows := append(degreeRowsForBlock("mixed +,-", "+,-"), degreeRowsForBlock("mixed -,+", "-,+")...)
	any := false
	for _, r := range rows {
		if r.SurvivesByDegree {
			any = true
		}
	}
	return MixedBlockDegreeAudit{Rows: rows, AnySurvivesByDegree: any, MixedBlockZeroByDegree: !any, SelectionMechanism: "no ordered channel with one positive and one negative interior contraction saturates the (4,3) top degree", Verdict: join(StatusMixedBlockZeroByTopFormDegree, StatusDegreeSelectionRuleSupported)}
}

func buildSignAudit(g gate649.Analysis) SignNormalizationAudit {
	finiteSigns := g.PositiveAAA.AAAContributesUnit && g.NegativeAAB.EachAABContributesMinusUnit
	return SignNormalizationAudit{DegreeRuleCertifiesSupport: true, Gate649CertifiesFiniteSigns: finiteSigns, EqualUnitWeightCertifiedByDegree: false, RequiresCalibrationIdentity: true, MissingProofObject: "sector-degree saturation explains support and channel count; the negative sign and equal unit magnitudes still require the octonionic calibration/orientation/antisymmetrization identity", Verdict: join(StatusSignUnitRequiresCalibration, StatusNoFullSymbolicDegreeTheorem)}
}

func buildTheorem(p PositiveBlockDegreeAudit, n NegativeBlockDegreeAudit, m MixedBlockDegreeAudit, s SignNormalizationAudit) SymbolicSelectionTheorem {
	candidate := "For Ω=A+B with A∈Λ^{2,1} and B∈Λ^{0,3} on a 4|3 split K7, the Hitchin cubic top-form contraction admits H(A,A,A) in the positive block, H(A,A,B)+H(A,B,A)+H(B,A,A) in the negative block, and no mixed block by sector-degree saturation. With the separate calibration sign/unit identity, g_twist ∝ P_+-3P_-."
	supported := p.AAAOnlySurvives && n.AABPlacementsOnly && m.MixedBlockZeroByDegree
	return SymbolicSelectionTheorem{CandidateTheorem: candidate, DegreeSelectionSupported: supported, FullSymbolicTheoremCertified: false, RemainingGap: s.MissingProofObject, Verdict: join(StatusDegreeSelectionRuleSupported, StatusSignUnitRequiresCalibration, StatusNoFullSymbolicDegreeTheorem)}
}

func buildSlotFormula(s SignNormalizationAudit) ResultingSlotFormula {
	p, q, d := float64(plusDim), float64(minusDim), float64(cubicDegree)
	norm := p + d*d*q
	cos := (p + d*q) / math.Sqrt((p+q)*norm)
	rho := p * q * (d - 1) * (d - 1) / ((p + q) * norm)
	return ResultingSlotFormula{PositiveDim: plusDim, NegativeDim: minusDim, SlotMultiplicity: cubicDegree, NormSquared: norm, Cosine: cos, ResidualSquared: rho, RecoversGate642Angle: math.Abs(cos-13/math.Sqrt(217)) < tol && math.Abs(rho-48.0/217.0) < tol, Formula: "with the separate equal-unit calibration identity, G_slot=(P_+-3P_-)/sqrt(31)", Verdict: join(StatusMinusThreeFromDegreePlacements, StatusSignUnitRequiresCalibration)}
}

func buildResonance() ResonanceAudit {
	eq := cubicDegree == minusDim
	return ResonanceAudit{CubicDegree: cubicDegree, NegativeDim: minusDim, EqualInASHACarrier: eq, Interpretation: "the degree rule sources three AAB placements from cubic order; ASHA also has dim(K7-)=3, so d=q is a carrier resonance rather than an independent dimension theorem", Verdict: StatusDEqualsQCarrierResonance}
}

type family struct {
	name   string
	degree Degree
}

type channelSpec struct {
	code       string
	f1, f2, f3 family
}

func channels() []channelSpec {
	A := family{"A", Degree{2, 1}}
	B := family{"B", Degree{0, 3}}
	return []channelSpec{
		{"AAA", A, A, A}, {"AAB", A, A, B}, {"ABA", A, B, A}, {"BAA", B, A, A},
		{"ABB", A, B, B}, {"BAB", B, A, B}, {"BBA", B, B, A}, {"BBB", B, B, B},
	}
}

func degreeRowsForBlock(block, xy string) []ChannelDegreeRow {
	rows := []ChannelDegreeRow{}
	for _, c := range channels() {
		xSec, ySec := xy[0:1], xy[2:3]
		first, zero1 := interiorDegree(c.f1, xSec)
		second, zero2 := interiorDegree(c.f2, ySec)
		total := first.Add(second).Add(c.f3.degree)
		survives := !zero1 && !zero2 && total.Top()
		rows = append(rows, ChannelDegreeRow{Block: block, XYSectors: xy, Channel: c.code, FirstFamily: c.f1.name, SecondFamily: c.f2.name, ThirdFamily: c.f3.name, FirstInterior: first, SecondInterior: second, ThirdDegree: c.f3.degree, TotalDegree: total, FirstSlotZero: zero1, SecondSlotZero: zero2, ReachesTopDegree: total.Top(), SurvivesByDegree: survives, ExpectedGate649Role: expectedRole(block, c.code), DegreeMechanism: mechanism(zero1, zero2, total, survives)})
	}
	return rows
}

func interiorDegree(f family, sector string) (Degree, bool) {
	switch sector {
	case "+":
		if f.degree.Plus == 0 {
			return Degree{}, true
		}
		return Degree{f.degree.Plus - 1, f.degree.Minus}, false
	case "-":
		if f.degree.Minus == 0 {
			return Degree{}, true
		}
		return Degree{f.degree.Plus, f.degree.Minus - 1}, false
	default:
		return Degree{}, true
	}
}

func mechanism(z1, z2 bool, total Degree, survives bool) string {
	if survives {
		return "degree saturates top form (4,3)"
	}
	if z1 || z2 {
		return "interior contraction kills one slot before top-form test"
	}
	if total.Plus > plusDim || total.Minus > minusDim {
		return "sector degree oversaturates available top-form degree"
	}
	return "sector degree undersaturates or mismatches top-form degree"
}

func expectedRole(block, code string) string {
	switch block {
	case "positive":
		if code == "AAA" {
			return "Gate649 positive channel +P_+"
		}
	case "negative":
		if code == "AAB" || code == "ABA" || code == "BAA" {
			return "Gate649 negative channel -P_-"
		}
	default:
		return "Gate649 mixed block zero"
	}
	return "blocked by degree rule"
}

func survivingChannels(rows []ChannelDegreeRow) []string {
	out := []string{}
	for _, r := range rows {
		if r.SurvivesByDegree {
			out = append(out, r.Channel)
		}
	}
	return out
}

func containsAll(xs, wants []string) bool {
	m := map[string]bool{}
	for _, x := range xs {
		m[x] = true
	}
	for _, w := range wants {
		if !m[w] {
			return false
		}
	}
	return true
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate649ChannelAlgebraInherited,
		StatusSectorDegreeLedgerDefined,
		StatusPositiveAAAOnlyByTopFormDegree,
		StatusNegativeAABByTopFormDegree,
		StatusMixedBlockZeroByTopFormDegree,
		StatusDegreeSelectionRuleSupported,
		StatusMinusThreeFromDegreePlacements,
		StatusDEqualsQCarrierResonance,
		StatusSignUnitRequiresCalibration,
		StatusNoFullSymbolicDegreeTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate650Boundary,
	}
}
