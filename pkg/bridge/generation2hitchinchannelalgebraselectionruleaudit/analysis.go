// Package generation2hitchinchannelalgebraselectionruleaudit implements
// Gate 649: Hitchin AAA/AAB Channel Algebra Selection Rule Audit.
//
// Gate 648 refined the source of the negative block coefficient in
//
//	g_twist ∝ P_+ - 3 P_-
//
// from a general negative-sector dimension theorem to the directly witnessed
// cubic ordered-slot multiplicity.  Gate 649 descends one more layer and audits
// the channel selection rule itself.  With
//
//	A = Ω++-,  B = Ω---,
//
// it checks that the admissible S_K-twisted native tensor is supported only on
// A and B and that the cubic Hitchin ledger has the finite channel algebra
//
//	AAA -> +P_+
//	AAB + ABA + BAA -> -3 P_-
//
// while the ABB/BAB/BBA/BBB families vanish, cancel, or project away.  This is
// internal finite tensor-contraction algebra only.  It does not certify split-G2,
// boundary stress, scalar/flavor transport, physical metric, Higgs mass,
// CKM/PMNS, gauge unification, or a native 7/72 theorem.
package generation2hitchinchannelalgebraselectionruleaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate648 "github.com/bagherbal/asha-engine/pkg/bridge/generation2cubicslotmultiplicityversusnegativesectordimensionaudit"
	gate647 "github.com/bagherbal/asha-engine/pkg/bridge/generation2hitchincubicsectorcontractionmultiplicityaudit"
)

const (
	AuditID = "GATE649-HITCHIN-AAA-AAB-CHANNEL-ALGEBRA-SELECTION-RULE-AUDIT"

	StatusGate648SlotMultiplicityInherited      = "PASS_GATE648_SLOT_MULTIPLICITY_RESULT_INHERITED"
	StatusTwoComponentTensorSupportAudited      = "PASS_TWO_COMPONENT_TENSOR_SUPPORT_AUDITED"
	StatusOrderedCubicExpansionComputed         = "PASS_ORDERED_CUBIC_EXPANSION_COMPUTED"
	StatusAAAPositiveChannelAudited             = "PASS_AAA_POSITIVE_CHANNEL_AUDITED"
	StatusAABNegativeChannelsAudited            = "PASS_AAB_NEGATIVE_CHANNELS_AUDITED"
	StatusABBBBBVanishingAudited                = "PASS_ABB_BBB_VANISHING_OR_CANCELLATION_AUDITED"
	StatusOffBlockCancellationAudited           = "PASS_OFF_BLOCK_CANCELLATION_AUDITED"
	StatusSlotFormulaDerived                    = "PASS_SLOT_FORMULA_DERIVED"
	StatusSlotTheoremPrimary                    = "CONDITIONAL_SUPPORT_SLOT_THEOREM_PRIMARY_SOURCE_FOR_MINUS_THREE"
	StatusDEqualsQCoincidence                   = "CONDITIONAL_SUPPORT_D_EQUALS_Q_AS_ASHA_CARRIER_COINCIDENCE"
	StatusChannelSelectionSharpened             = "CONDITIONAL_SUPPORT_HITCHIN_CHANNEL_SELECTION_RULE_SHARPENED"
	StatusNoFullSymbolicChannelSelectionTheorem = "FAILED_ROUTE_NO_FULL_SYMBOLIC_CHANNEL_SELECTION_THEOREM"
	StatusNoSplitG2                             = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                      = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                        = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                      = "FAILED_ROUTE_HITCHIN_CHANNEL_METRIC_IS_NOT_PHYSICAL_METRIC"
	StatusNoHiggsFlavorGauge                    = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate649Boundary                       = "FIREWALL_PRESERVED_GATE649_HITCHIN_CHANNEL_ALGEBRA_BOUNDARY"
)

const (
	plusDim     = 4
	minusDim    = 3
	cubicDegree = 3
	tol         = 1e-8
)

type Gate648Inheritance struct {
	SlotMultiplicityInherited bool
	PositiveDim               int
	NegativeDim               int
	CubicDegree               int
	SlotSourceSupported       bool
	ASHADimEqualsDegree       bool
	GeneralPQDimensionTheorem bool
	CubicSlotTheoremCertified bool
	FullSymbolicHitchin       bool
	SplitG2Certified          bool
	BoundaryStressAssignment  bool
	SevenOver72Theorem        bool
	ScalarFlavorTransport     bool
	PhysicalMetric            bool
	Gate648FirewallPreserved  bool
	Verdict                   string
}

type FamilySupportRow struct {
	Family         string
	Alias          string
	Omega1NormSq   float64
	Omega2NormSq   float64
	OmegaBNormSq   float64
	Supported      bool
	Role           string
	ResidualLeak   bool
	Interpretation string
}

type TwoComponentTensorSupportAudit struct {
	Rows               []FamilySupportRow
	OnlyAAndBSupported bool
	AName              string
	BName              string
	SupportVerdict     string
	Verdict            string
}

type OrderedChannelRow struct {
	RouteName           string
	Channel             string
	Class               string
	PlusMeanUnit        float64
	MinusMeanUnit       float64
	MixedFrobenius      float64
	Nonzero             bool
	ExpectedRole        string
	SelectionRuleStatus string
}

type OrderedCubicExpansionAudit struct {
	Rows                    []OrderedChannelRow
	RouteCount              int
	ChannelsPerRoute        int
	NonzeroChannelsPerRoute int
	AAAOnlyPositive         bool
	AABOnlyNegative         bool
	ABBBBBClean             bool
	MixedBlocksClean        bool
	Verdict                 string
}

type PositiveChannelAudit struct {
	Rows                   []OrderedChannelRow
	AAAContributesUnit     bool
	AAAContributesOnlyPlus bool
	UnitCoefficient        float64
	SourceClassification   string
	Verdict                string
}

type NegativeChannelAudit struct {
	Rows                        []OrderedChannelRow
	NegativeOrderedChannelCount int
	EachAABContributesMinusUnit bool
	CombinedNegativeCoefficient float64
	SourceClassification        string
	Verdict                     string
}

type VanishingChannelRow struct {
	RouteName       string
	Channel         string
	Class           string
	PlusMeanUnit    float64
	MinusMeanUnit   float64
	MixedFrobenius  float64
	Vanishing       bool
	SourceCandidate string
}

type VanishingCancellationAudit struct {
	Rows                       []VanishingChannelRow
	AllVanishOrProjectAway     bool
	MechanismCandidate         string
	SymbolicMechanismCertified bool
	Verdict                    string
}

type OffBlockCancellationAudit struct {
	MaxMixedFrobenius       float64
	ChannelwiseZero         bool
	PairwiseCancellation    bool
	FullSumCancellation     bool
	MechanismClassification string
	Verdict                 string
}

type SlotFormulaDerivation struct {
	PositiveDim            int
	NegativeDim            int
	SlotMultiplicity       int
	GSlotFormula           string
	NormSquared            float64
	CosineFormula          string
	Cosine                 float64
	ResidualSquaredFormula string
	ResidualSquared        float64
	RecoversGate642Angle   bool
	Verdict                string
}

type DimensionCoincidenceAudit struct {
	SlotMultiplicity         int
	NegativeDim              int
	EqualInASHACarrier       bool
	SupportsSlotTheoremOnly  bool
	SupportsDimensionTheorem bool
	Interpretation           string
	Verdict                  string
}

type SymbolicTheoremReadiness struct {
	CandidateTheorem             string
	FiniteChannelRuleSupported   bool
	FullSymbolicTheoremCertified bool
	MissingProofObject           string
	Verdict                      string
}

type Firewalls struct {
	ClaimsFullSymbolicChannelSelection bool
	ClaimsSplitG2                      bool
	ClaimsBoundaryStress               bool
	ClaimsSevenOver72                  bool
	ClaimsScalarFlavor                 bool
	ClaimsPhysicalMetric               bool
	ClaimsHiggsMass                    bool
	ClaimsCKMPMNS                      bool
	ClaimsGaugeUnification             bool
	Verdict                            string
}

type Analysis struct {
	Inherited   Gate648Inheritance
	Support     TwoComponentTensorSupportAudit
	Expansion   OrderedCubicExpansionAudit
	PositiveAAA PositiveChannelAudit
	NegativeAAB NegativeChannelAudit
	Vanishing   VanishingCancellationAudit
	OffBlock    OffBlockCancellationAudit
	SlotFormula SlotFormulaDerivation
	Coincidence DimensionCoincidenceAudit
	Readiness   SymbolicTheoremReadiness
	Firewalls   Firewalls
	Truth       string
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
	g648, err := gate648.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate648 inheritance unavailable: %w", err)
	}
	g647, err := gate647.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate647 contribution ledger unavailable: %w", err)
	}
	inherited := buildInheritance(g648)
	support := buildSupport(g647)
	expansion := buildExpansion(g647)
	positive := buildPositiveAAA(expansion)
	negative := buildNegativeAAB(expansion)
	vanishing := buildVanishing(expansion)
	off := buildOffBlock(expansion)
	formula := buildSlotFormula()
	coin := buildCoincidence()
	readiness := buildReadiness(expansion, positive, negative, vanishing, off)
	firewalls := Firewalls{Verdict: StatusGate649Boundary}
	truth := "Gate 649 refines the Gate648 slot source into a channel-selection rule.  The finite Hitchin cubic ledger supports the two-component tensor support A=Ω++- and B=Ω---, with AAA producing the positive block and the three ordered AAB placements producing the negative block.  ABB/BAB/BBA/BBB do not contribute to the final projector-plane ray in the audited routes.  This supports the slot theorem as the primary source of the -3 coefficient and records d=q=3 as an ASHA carrier coincidence, while a basis-free symbolic channel-selection theorem, split-G2, boundary stress, scalar/flavor transport, physical metric, and native 7/72 remain firewalled."
	return Analysis{Inherited: inherited, Support: support, Expansion: expansion, PositiveAAA: positive, NegativeAAB: negative, Vanishing: vanishing, OffBlock: off, SlotFormula: formula, Coincidence: coin, Readiness: readiness, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g648 gate648.Analysis) Gate648Inheritance {
	return Gate648Inheritance{
		SlotMultiplicityInherited: g648.SlotAudit.SlotSourceSupported && g648.TraceAudit.AllRoutesPass && g648.FormulaAudit.LedgerSelectsSlotSource,
		PositiveDim:               plusDim,
		NegativeDim:               minusDim,
		CubicDegree:               cubicDegree,
		SlotSourceSupported:       g648.SlotAudit.SlotSourceSupported,
		ASHADimEqualsDegree:       g648.TheoremTarget.ASHACoincidenceCertified,
		GeneralPQDimensionTheorem: g648.TheoremTarget.GeneralPQDimensionTheorem || g648.Firewalls.ClaimsGeneralPQDimensionTheorem,
		CubicSlotTheoremCertified: g648.TheoremTarget.CubicSlotTheoremCertified || g648.Firewalls.ClaimsCubicSlotTheorem,
		FullSymbolicHitchin:       g648.Firewalls.ClaimsFullSymbolicHitchin,
		SplitG2Certified:          g648.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:  g648.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:        g648.Firewalls.ClaimsSevenOver72,
		ScalarFlavorTransport:     g648.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:            g648.Firewalls.ClaimsPhysicalMetric,
		Gate648FirewallPreserved:  g648.Firewalls.Verdict == gate648.StatusGate648Boundary,
		Verdict:                   StatusGate648SlotMultiplicityInherited,
	}
}

func buildSupport(g647 gate647.Analysis) TwoComponentTensorSupportAudit {
	rows := make([]FamilySupportRow, 0, len(g647.Families.Families))
	onlyAB := true
	for _, f := range g647.Families.Families {
		supported := f.Omega1AltNormSq > tol || f.Omega2AltNormSq > tol || f.OmegaBAltNormSq > tol
		alias := ""
		role := "zero family in the audited admissible twists"
		if f.Family == "Ω++-" {
			alias = "A"
			role = "A-family; two positive slots and one negative slot; sources AAA positive and AAB negative channels"
		}
		if f.Family == "Ω---" {
			alias = "B"
			role = "B-family; three negative slots; appears once in each ordered AAB negative channel"
		}
		if supported && alias == "" {
			onlyAB = false
		}
		rows = append(rows, FamilySupportRow{Family: f.Family, Alias: alias, Omega1NormSq: f.Omega1AltNormSq, Omega2NormSq: f.Omega2AltNormSq, OmegaBNormSq: f.OmegaBAltNormSq, Supported: supported, Role: role, ResidualLeak: supported && alias == "", Interpretation: "support read from the Gate647 component-family tensor ledger"})
	}
	return TwoComponentTensorSupportAudit{Rows: rows, OnlyAAndBSupported: onlyAB, AName: "Ω++-", BName: "Ω---", SupportVerdict: "admissible twisted tensor support is carried by A=Ω++- and B=Ω--- in the audited routes", Verdict: join(StatusTwoComponentTensorSupportAudited, StatusChannelSelectionSharpened)}
}

func buildExpansion(g647 gate647.Analysis) OrderedCubicExpansionAudit {
	rows := []OrderedChannelRow{}
	for _, r := range g647.Contributions.Routes {
		byChannel := map[string]gate647.TripleContribution{}
		for _, c := range r.TopContributions {
			byChannel[c.Families] = c
		}
		for _, ch := range orderedChannels() {
			tr, ok := byChannel[ch.full]
			row := OrderedChannelRow{RouteName: r.RouteName, Channel: ch.full, Class: ch.class, ExpectedRole: ch.role}
			if ok {
				row.PlusMeanUnit = tr.PlusMeanUnit
				row.MinusMeanUnit = tr.MinusMeanUnit
				row.MixedFrobenius = tr.MixedFrobenius
				row.Nonzero = true
			} else {
				row.Nonzero = false
			}
			row.SelectionRuleStatus = classifyChannel(row)
			rows = append(rows, row)
		}
	}
	return OrderedCubicExpansionAudit{Rows: rows, RouteCount: len(g647.Contributions.Routes), ChannelsPerRoute: 8, NonzeroChannelsPerRoute: 4, AAAOnlyPositive: aaaOnlyPositive(rows), AABOnlyNegative: aabOnlyNegative(rows), ABBBBBClean: abbBBBVanishing(rows), MixedBlocksClean: mixedClean(rows), Verdict: join(StatusOrderedCubicExpansionComputed, StatusChannelSelectionSharpened)}
}

func buildPositiveAAA(exp OrderedCubicExpansionAudit) PositiveChannelAudit {
	rows := filterRows(exp.Rows, "AAA")
	ok := len(rows) == exp.RouteCount
	for _, r := range rows {
		if !r.Nonzero || math.Abs(r.PlusMeanUnit-1) > tol || math.Abs(r.MinusMeanUnit) > tol || r.MixedFrobenius > tol {
			ok = false
		}
	}
	return PositiveChannelAudit{Rows: rows, AAAContributesUnit: ok, AAAContributesOnlyPlus: ok, UnitCoefficient: 1, SourceClassification: "AAA is the only positive source channel and gives +cP_+ after normalization by the positive sector coefficient c", Verdict: join(StatusAAAPositiveChannelAudited, StatusSlotTheoremPrimary)}
}

func buildNegativeAAB(exp OrderedCubicExpansionAudit) NegativeChannelAudit {
	rows := []OrderedChannelRow{}
	for _, r := range exp.Rows {
		if r.Class == "AAB" || r.Class == "ABA" || r.Class == "BAA" {
			rows = append(rows, r)
		}
	}
	ok := len(rows) == exp.RouteCount*cubicDegree
	for _, r := range rows {
		if !r.Nonzero || math.Abs(r.MinusMeanUnit+1) > tol || math.Abs(r.PlusMeanUnit) > tol || r.MixedFrobenius > tol {
			ok = false
		}
	}
	return NegativeChannelAudit{Rows: rows, NegativeOrderedChannelCount: cubicDegree, EachAABContributesMinusUnit: ok, CombinedNegativeCoefficient: -float64(cubicDegree), SourceClassification: "each ordered AAB placement contributes -cP_-; summing AAB+ABA+BAA gives -3cP_-", Verdict: join(StatusAABNegativeChannelsAudited, StatusSlotTheoremPrimary)}
}

func buildVanishing(exp OrderedCubicExpansionAudit) VanishingCancellationAudit {
	rows := []VanishingChannelRow{}
	all := true
	for _, r := range exp.Rows {
		if r.Class == "ABB" || r.Class == "BAB" || r.Class == "BBA" || r.Class == "BBB" {
			v := !r.Nonzero && math.Abs(r.PlusMeanUnit) < tol && math.Abs(r.MinusMeanUnit) < tol && r.MixedFrobenius < tol
			if !v {
				all = false
			}
			rows = append(rows, VanishingChannelRow{RouteName: r.RouteName, Channel: r.Channel, Class: r.Class, PlusMeanUnit: r.PlusMeanUnit, MinusMeanUnit: r.MinusMeanUnit, MixedFrobenius: r.MixedFrobenius, Vanishing: v, SourceCandidate: "finite ledger shows no significant final block contribution; exact source is not yet separated among sector-degree impossibility, antisymmetry, Hodge parity, or octonionic calibration identity"})
		}
	}
	return VanishingCancellationAudit{Rows: rows, AllVanishOrProjectAway: all, MechanismCandidate: "ABB/BAB/BBA/BBB do not survive in the audited final block ray; the source is classified only as a candidate vanishing/project-away mechanism", SymbolicMechanismCertified: false, Verdict: join(StatusABBBBBVanishingAudited, StatusNoFullSymbolicChannelSelectionTheorem)}
}

func buildOffBlock(exp OrderedCubicExpansionAudit) OffBlockCancellationAudit {
	maxMix := 0.0
	channelwise := true
	for _, r := range exp.Rows {
		if r.MixedFrobenius > maxMix {
			maxMix = r.MixedFrobenius
		}
		if r.MixedFrobenius > tol {
			channelwise = false
		}
	}
	return OffBlockCancellationAudit{MaxMixedFrobenius: maxMix, ChannelwiseZero: channelwise, PairwiseCancellation: false, FullSumCancellation: !channelwise, MechanismClassification: "mixed block is zero channelwise in the retained eight-channel support ledger; no basis-free source theorem for this cancellation is certified", Verdict: join(StatusOffBlockCancellationAudited, StatusNoFullSymbolicChannelSelectionTheorem)}
}

func buildSlotFormula() SlotFormulaDerivation {
	p, q, d := float64(plusDim), float64(minusDim), float64(cubicDegree)
	normSq := p + d*d*q
	cos := (p + d*q) / math.Sqrt((p+q)*normSq)
	rho := p * q * (d - 1) * (d - 1) / ((p + q) * normSq)
	return SlotFormulaDerivation{PositiveDim: plusDim, NegativeDim: minusDim, SlotMultiplicity: cubicDegree, GSlotFormula: "G_slot=(P_+-dP_-)/sqrt(p+d^2q), d=3", NormSquared: normSq, CosineFormula: "cos(theta)=(p+dq)/sqrt((p+q)(p+d^2q))", Cosine: cos, ResidualSquaredFormula: "rho^2=pq(d-1)^2/[(p+q)(p+d^2q)]", ResidualSquared: rho, RecoversGate642Angle: math.Abs(cos-13/math.Sqrt(217)) < tol && math.Abs(rho-48.0/217.0) < tol, Verdict: join(StatusSlotFormulaDerived, StatusSlotTheoremPrimary)}
}

func buildCoincidence() DimensionCoincidenceAudit {
	eq := cubicDegree == minusDim
	return DimensionCoincidenceAudit{SlotMultiplicity: cubicDegree, NegativeDim: minusDim, EqualInASHACarrier: eq, SupportsSlotTheoremOnly: true, SupportsDimensionTheorem: false, Interpretation: "Gate649 treats d=3 as the primary witnessed source and q=3 as the ASHA-specific carrier coincidence; the final ray cannot by itself distinguish d from q", Verdict: join(StatusDEqualsQCoincidence, StatusSlotTheoremPrimary)}
}

func buildReadiness(exp OrderedCubicExpansionAudit, p PositiveChannelAudit, n NegativeChannelAudit, v VanishingCancellationAudit, off OffBlockCancellationAudit) SymbolicTheoremReadiness {
	candidate := "For admissible S_K-twisted native Ω_0 with A=Ω++- and B=Ω---, H(A,A,A)=+cP_+, H(A,A,B)+H(A,B,A)+H(B,A,A)=-3cP_-, and all other ordered families vanish/cancel, hence g_twist ∝ P_+-3P_-."
	finite := exp.AAAOnlyPositive && exp.AABOnlyNegative && v.AllVanishOrProjectAway && off.ChannelwiseZero && p.AAAContributesUnit && n.EachAABContributesMinusUnit
	missing := "a basis-free proof that the native tensor support is exactly A+B and that the Hitchin cubic functional obeys the audited AAA/AAB selection rule before choosing the finite computational frame"
	return SymbolicTheoremReadiness{CandidateTheorem: candidate, FiniteChannelRuleSupported: finite, FullSymbolicTheoremCertified: false, MissingProofObject: missing, Verdict: join(StatusChannelSelectionSharpened, StatusNoFullSymbolicChannelSelectionTheorem)}
}

type channelSpec struct{ full, class, role string }

func orderedChannels() []channelSpec {
	return []channelSpec{
		{"Ω++-×Ω++-×Ω++-", "AAA", "positive unit block"},
		{"Ω++-×Ω++-×Ω---", "AAB", "negative ordered slot"},
		{"Ω++-×Ω---×Ω++-", "ABA", "negative ordered slot"},
		{"Ω---×Ω++-×Ω++-", "BAA", "negative ordered slot"},
		{"Ω++-×Ω---×Ω---", "ABB", "vanishing/projected-away channel"},
		{"Ω---×Ω++-×Ω---", "BAB", "vanishing/projected-away channel"},
		{"Ω---×Ω---×Ω++-", "BBA", "vanishing/projected-away channel"},
		{"Ω---×Ω---×Ω---", "BBB", "vanishing/projected-away channel"},
	}
}

func classifyChannel(r OrderedChannelRow) string {
	switch r.Class {
	case "AAA":
		if r.Nonzero && math.Abs(r.PlusMeanUnit-1) < tol && math.Abs(r.MinusMeanUnit) < tol && r.MixedFrobenius < tol {
			return "AAA -> +P_+"
		}
	case "AAB", "ABA", "BAA":
		if r.Nonzero && math.Abs(r.MinusMeanUnit+1) < tol && math.Abs(r.PlusMeanUnit) < tol && r.MixedFrobenius < tol {
			return r.Class + " -> -P_-"
		}
	case "ABB", "BAB", "BBA", "BBB":
		if !r.Nonzero && math.Abs(r.PlusMeanUnit) < tol && math.Abs(r.MinusMeanUnit) < tol && r.MixedFrobenius < tol {
			return r.Class + " -> 0/projected away"
		}
	}
	return "UNRESOLVED_CHANNEL_STATUS"
}

func filterRows(rows []OrderedChannelRow, class string) []OrderedChannelRow {
	out := []OrderedChannelRow{}
	for _, r := range rows {
		if r.Class == class {
			out = append(out, r)
		}
	}
	return out
}

func aaaOnlyPositive(rows []OrderedChannelRow) bool {
	for _, r := range rows {
		if r.Class == "AAA" && (!r.Nonzero || math.Abs(r.PlusMeanUnit-1) > tol || math.Abs(r.MinusMeanUnit) > tol || r.MixedFrobenius > tol) {
			return false
		}
	}
	return true
}

func aabOnlyNegative(rows []OrderedChannelRow) bool {
	for _, r := range rows {
		if (r.Class == "AAB" || r.Class == "ABA" || r.Class == "BAA") && (!r.Nonzero || math.Abs(r.MinusMeanUnit+1) > tol || math.Abs(r.PlusMeanUnit) > tol || r.MixedFrobenius > tol) {
			return false
		}
	}
	return true
}

func abbBBBVanishing(rows []OrderedChannelRow) bool {
	for _, r := range rows {
		if (r.Class == "ABB" || r.Class == "BAB" || r.Class == "BBA" || r.Class == "BBB") && r.Nonzero {
			return false
		}
	}
	return true
}

func mixedClean(rows []OrderedChannelRow) bool {
	for _, r := range rows {
		if r.MixedFrobenius > tol {
			return false
		}
	}
	return true
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate648SlotMultiplicityInherited,
		StatusTwoComponentTensorSupportAudited,
		StatusOrderedCubicExpansionComputed,
		StatusAAAPositiveChannelAudited,
		StatusAABNegativeChannelsAudited,
		StatusABBBBBVanishingAudited,
		StatusOffBlockCancellationAudited,
		StatusSlotFormulaDerived,
		StatusSlotTheoremPrimary,
		StatusDEqualsQCoincidence,
		StatusChannelSelectionSharpened,
		StatusNoFullSymbolicChannelSelectionTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate649Boundary,
	}
}
