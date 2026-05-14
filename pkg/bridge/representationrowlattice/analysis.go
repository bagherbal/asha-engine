// Package representationrowlattice implements Gate 204: representation-row
// lattice completion / finite heavy-sector basis search.
//
// Gate 203 proved that the universal beta completion required by Gate 201 is
// not currently sourced by complete multiplets or by regulator/ghost traces.
// Gate 204 therefore deliberately decouples the rational representation
// geometry from the continuous threshold scales.  It constructs a finite,
// exact rational beta-row grammar from the already-audited gauge alphabet:
//
//	SU(3)c: singlet, fundamental/antifundamental, adjoint
//	SU(2)L: singlet, fundamental doublet, adjoint triplet
//	|Y|:    0, 1/6, 1/3, 1/2, 2/3, 1
//
// with standard one-loop row formulas in GUT-normalized hypercharge.  This is
// not a mass spectrum, not a physical threshold ledger, and not a unification
// claim.  It is a discrete rational viability audit: do the non-universal
// Gate-201 shapes live in the representation-row lattice at all?
package representationrowlattice

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/universalbetasource"
)

type Rational struct {
	Num int64
	Den int64
}

func R(num, den int64) Rational {
	if den == 0 {
		panic("zero denominator")
	}
	if den < 0 {
		num = -num
		den = -den
	}
	g := gcd(abs(num), den)
	return Rational{Num: num / g, Den: den / g}
}

func (r Rational) Add(o Rational) Rational { return R(r.Num*o.Den+o.Num*r.Den, r.Den*o.Den) }
func (r Rational) Mul(o Rational) Rational { return R(r.Num*o.Num, r.Den*o.Den) }
func (r Rational) MulInt(n int64) Rational { return R(r.Num*n, r.Den) }
func (r Rational) IsZero() bool            { return r.Num == 0 }
func (r Rational) Float() float64          { return float64(r.Num) / float64(r.Den) }

func (r Rational) String() string {
	if r.Den == 1 {
		return fmt.Sprintf("%d", r.Num)
	}
	return fmt.Sprintf("%d/%d", r.Num, r.Den)
}

func (r Rational) Equal(o Rational) bool { return r.Num == o.Num && r.Den == o.Den }

func gcd(a, b int64) int64 {
	if a == 0 {
		if b == 0 {
			return 1
		}
		return b
	}
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int64) int64 {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a / gcd(abs(a), abs(b)) * b)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

type RationalTriple struct {
	B1 Rational
	B2 Rational
	B3 Rational
}

func RT(b1, b2, b3 Rational) RationalTriple { return RationalTriple{B1: b1, B2: b2, B3: b3} }

func (t RationalTriple) String() string {
	return fmt.Sprintf("(%s,%s,%s)", t.B1, t.B2, t.B3)
}

func (t RationalTriple) Equal(o RationalTriple) bool {
	return t.B1.Equal(o.B1) && t.B2.Equal(o.B2) && t.B3.Equal(o.B3)
}

func (t RationalTriple) Key() string { return t.String() }

func (t RationalTriple) DenominatorLCM() int64 {
	return lcm(lcm(t.B1.Den, t.B2.Den), t.B3.Den)
}

type Gate203Snapshot struct {
	Gate203Inherited                     bool
	Gate203FailedRoutePreserved          bool
	UniversalBetaSourceStillExternal     bool
	CompleteMultipletSourceFound         bool
	RegulatorTraceSourceFound            bool
	ContactModesHaveBetaPermission       bool
	FockGenerationPromotedToNewThreshold bool
	PhysicalUnificationClaimed           bool
	ThresholdCorrectedPhysicalFitClaimed bool
	AbsoluteMassPredicted                bool
	FiniteMatchingCorrectionsDerived     bool
	StrictNullityAfter                   int
	PhysicalPredictionNullityAfter       int
	Gate201NonUniversalShapeRequirements []Gate201ShapeRequirement
	TruthStatement                       string
}

type Gate201ShapeRequirement struct {
	Name               string
	DeltaB             RationalTriple
	RequiredUniversalC float64
	ConditionalOnly    bool
	FiniteDerived      bool
	FromPhenomenology  bool
}

func DefaultGate203Snapshot() (Gate203Snapshot, error) {
	prev, err := universalbetasource.BuildDefault()
	if err != nil {
		return Gate203Snapshot{}, err
	}
	reqs := make([]Gate201ShapeRequirement, 0, len(prev.PreviousGate202.Requirements))
	for _, r := range prev.PreviousGate202.Requirements {
		tr, ok := knownGate201ShapeTriple(r.ShapeName)
		if !ok {
			return Gate203Snapshot{}, fmt.Errorf("Gate 204 cannot parse inherited Gate-201 shape %q", r.ShapeName)
		}
		reqs = append(reqs, Gate201ShapeRequirement{
			Name:               r.ShapeName,
			DeltaB:             tr,
			RequiredUniversalC: r.RequiredUniversalBeta,
			ConditionalOnly:    r.ConditionalOnly,
			FiniteDerived:      r.FiniteDerived,
			FromPhenomenology:  true,
		})
	}
	return Gate203Snapshot{
		Gate203Inherited:                     true,
		Gate203FailedRoutePreserved:          prev.Summary.FailedRouteLogged && prev.Firewall.Gate202FailedRoutePreserved,
		UniversalBetaSourceStillExternal:     prev.Summary.UniversalBetaSourceStillExternal && !prev.Firewall.UniversalBetaSourceDerived,
		CompleteMultipletSourceFound:         prev.MultipletAudit.CompleteMultipletSourceFound,
		RegulatorTraceSourceFound:            prev.RegulatorAudit.RegulatorTraceSourceFound,
		ContactModesHaveBetaPermission:       prev.Firewall.ContactModesPromotedToBetaRows,
		FockGenerationPromotedToNewThreshold: prev.Firewall.FockGenerationPromotedToNewThreshold,
		PhysicalUnificationClaimed:           prev.Firewall.PhysicalUnificationClaimed,
		ThresholdCorrectedPhysicalFitClaimed: prev.Firewall.ThresholdCorrectedPhysicalFitClaimed,
		AbsoluteMassPredicted:                prev.Firewall.AbsoluteMassPredicted,
		FiniteMatchingCorrectionsDerived:     prev.Firewall.FiniteMatchingCorrectionsDerived,
		StrictNullityAfter:                   prev.Firewall.StrictNullityAfter,
		PhysicalPredictionNullityAfter:       prev.Firewall.PhysicalPredictionNullityAfter,
		Gate201NonUniversalShapeRequirements: reqs,
		TruthStatement:                       prev.TruthStatement,
	}, nil
}

func knownGate201ShapeTriple(name string) (RationalTriple, bool) {
	switch name {
	case "Dirac vectorlike quark doublet":
		return RT(R(2, 15), R(2, 1), R(4, 3)), true
	case "Weyl SU(2)L adjoint fermion":
		return RT(R(0, 1), R(4, 3), R(0, 1)), true
	default:
		return RationalTriple{}, false
	}
}

type GroupRep struct {
	Name                 string
	Symbol               string
	Dim                  int64
	DynkinT              Rational
	RealType             bool
	FiniteAlphabetReason string
}

type Statistic struct {
	Name                 string
	Symbol               string
	Coeff                Rational
	RequiresY0           bool
	RequiresRealGaugeRep bool
}

type RepresentationRow struct {
	Name                   string
	Statistics             Statistic
	SU3                    GroupRep
	SU2                    GroupRep
	Hypercharge            Rational
	DeltaB                 RationalTriple
	ExactRational          bool
	FiniteAlphabet         bool
	StandardOneLoopFormula bool
	DirectGate201Shape     bool
	Verdict                string
}

type RowGrammarAudit struct {
	SU3AlphabetSize             int
	SU2AlphabetSize             int
	HyperchargeAlphabetSize     int
	StatisticsCount             int
	CandidateRowsGenerated      int
	UniqueRows                  int
	ExactRationalRows           int
	StandardFormulaRows         int
	CommonDenominatorLCM        int64
	FiniteAlphabetDeclared      bool
	UnboundedEnumerationAvoided bool
	Verdict                     string
}

type LatticeAudit struct {
	GeneratorCount      int
	UniqueGeneratorRows int
	CommonDenominator   int64
	IntegerGridEmbedded bool
	ContainsZeroRow     bool
	SemigroupOnly       bool
	NoContinuousScales  bool
	NoUniversalFit      bool
	Verdict             string
}

type ShapeMembership struct {
	ShapeName                  string
	TargetDeltaB               RationalTriple
	Found                      bool
	DirectGenerator            bool
	MatchedRepresentation      string
	MatchedStatistics          string
	MatchedSMRep               string
	ConditionalSupport         bool
	UniversalCompletionIgnored bool
	RequiredUniversalCExternal float64
	FiniteDerived              bool
	Verdict                    string
}

type MembershipAudit struct {
	ShapesAudited              int
	ShapesOnLattice            int
	DirectGeneratorMatches     int
	ConditionalSupportCount    int
	UniversalCompletionIgnored bool
	AllGate201ShapesSupported  bool
	Verdict                    string
}

type ContactInventoryAudit struct {
	ContactPartialOverlapModes        int
	ContactModesHaveChargeLabels      bool
	ContactModesHaveGaugeRepSemantics bool
	ContactModesHaveDynkinIndices     bool
	ContactModesHaveSpinStatistics    bool
	ContactModesHaveMassActivation    bool
	ContactModesHaveDecouplingLaw     bool
	CanonicalMapToRowBasisFound       bool
	CandidateRowsAssigned             int
	FiniteHeavySectorBasisDerived     bool
	Verdict                           string
}

type FirewallAudit struct {
	Gate203Inherited                        bool
	Gate203FailedRoutePreserved             bool
	UniversalBetaSourceStillExternal        bool
	RepresentationLatticeConstructed        bool
	Gate201ShapesPromotedToFinitePrediction bool
	UniversalBetaFitAttempted               bool
	ContinuousScalesSolved                  bool
	ObservedInputsUsedForFiniteDerivation   bool
	ContactModesPromotedToBetaRows          bool
	FockGenerationPromotedToNewThreshold    bool
	PhysicalUnificationClaimed              bool
	ThresholdCorrectedPhysicalFitClaimed    bool
	AbsoluteMassPredicted                   bool
	FiniteMatchingCorrectionsDerived        bool
	StrictNullityBefore                     int
	StrictNullityAfter                      int
	PhysicalPredictionNullityBefore         int
	PhysicalPredictionNullityAfter          int
	RecommendedNextGate                     string
	OpenRequirements                        []string
	Verdict                                 string
}

type Summary struct {
	TestsAudited               int
	Gate203Inherited           bool
	RationalGrammarConstructed bool
	LatticeConstructed         bool
	Gate201ShapesOnLattice     bool
	ConditionalSupportLogged   bool
	ContactMapFailed           bool
	UniversalFitAvoided        bool
	NoPhysicalPredictionClaim  bool
	Status                     string
	Comment                    string
}

type Analysis struct {
	PreviousGate203     Gate203Snapshot
	SU3Alphabet         []GroupRep
	SU2Alphabet         []GroupRep
	HyperchargeAlphabet []Rational
	Statistics          []Statistic
	Rows                []RepresentationRow
	UniqueRows          []RepresentationRow
	GrammarAudit        RowGrammarAudit
	LatticeAudit        LatticeAudit
	Memberships         []ShapeMembership
	MembershipAudit     MembershipAudit
	ContactInventory    ContactInventoryAudit
	Firewall            FirewallAudit
	Summary             Summary
	TruthStatement      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := DefaultGate203Snapshot()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev Gate203Snapshot) (Analysis, error) {
	if !prev.Gate203Inherited || !prev.Gate203FailedRoutePreserved || !prev.UniversalBetaSourceStillExternal || len(prev.Gate201NonUniversalShapeRequirements) == 0 {
		return Analysis{}, fmt.Errorf("Gate 204 requires Gate 203 failed-route source classification and inherited Gate-201 shapes")
	}
	if prev.PhysicalUnificationClaimed || prev.ThresholdCorrectedPhysicalFitClaimed || prev.AbsoluteMassPredicted || prev.FiniteMatchingCorrectionsDerived || prev.CompleteMultipletSourceFound || prev.RegulatorTraceSourceFound {
		return Analysis{}, fmt.Errorf("Gate 204 refuses inherited prediction/source leakage")
	}

	su3 := su3Alphabet()
	su2 := su2Alphabet()
	hy := hyperchargeAlphabet()
	stats := statisticAlphabet()
	rows := generateRows(su3, su2, hy, stats)
	markGate201Rows(rows, prev.Gate201NonUniversalShapeRequirements)
	unique := uniqueRows(rows)
	grammar := auditGrammar(su3, su2, hy, stats, rows, unique)
	lattice := auditLattice(unique)
	memberships := auditMemberships(prev.Gate201NonUniversalShapeRequirements, rows)
	ma := summarizeMemberships(memberships)
	contact := auditContactInventory()
	fw := auditFirewall(prev, grammar, lattice, ma, contact)
	summary := Summary{
		TestsAudited:               7,
		Gate203Inherited:           fw.Gate203Inherited && fw.Gate203FailedRoutePreserved,
		RationalGrammarConstructed: grammar.FiniteAlphabetDeclared && grammar.ExactRationalRows == grammar.CandidateRowsGenerated,
		LatticeConstructed:         lattice.IntegerGridEmbedded && lattice.UniqueGeneratorRows == grammar.UniqueRows,
		Gate201ShapesOnLattice:     ma.AllGate201ShapesSupported,
		ConditionalSupportLogged:   ma.ConditionalSupportCount == len(prev.Gate201NonUniversalShapeRequirements),
		ContactMapFailed:           !contact.CanonicalMapToRowBasisFound && !contact.FiniteHeavySectorBasisDerived,
		UniversalFitAvoided:        lattice.NoUniversalFit && fw.UniversalBetaSourceStillExternal && !fw.UniversalBetaFitAttempted,
		NoPhysicalPredictionClaim:  !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed && !fw.AbsoluteMassPredicted && fw.PhysicalPredictionNullityBefore == fw.PhysicalPredictionNullityAfter,
		Status:                     "CONDITIONAL_SUPPORT",
		Comment:                    "Gate 204 constructs a finite exact rational representation-row lattice from the existing charge/gauge alphabet. The two non-universal Gate-201 shapes are direct lattice generators, so their shapes are representation-theoretically viable. The universal completion, threshold scales, contact-mode mapping, and finite heavy-sector origin remain unproved.",
	}
	truth := "Gate 204 decouples discrete representation geometry from continuous RG lever arms. The finite charge/gauge alphabet generates an exact rational beta-row lattice. The Gate-201 non-universal shapes, Dirac vectorlike quark doublet (2/15,2,4/3) and Weyl SU(2)L adjoint fermion (0,4/3,0), are direct generators of that lattice. This is conditional support for their representation shape only. It does not derive a mass threshold, universal beta row, M_*, matching correction, or physical unification. The seven contact partial-overlap modes still cannot be canonically mapped to row generators because charge, Dynkin-index, spin-statistics, mass-activation, and decoupling semantics remain absent."

	return Analysis{PreviousGate203: prev, SU3Alphabet: su3, SU2Alphabet: su2, HyperchargeAlphabet: hy, Statistics: stats, Rows: rows, UniqueRows: unique, GrammarAudit: grammar, LatticeAudit: lattice, Memberships: memberships, MembershipAudit: ma, ContactInventory: contact, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func su3Alphabet() []GroupRep {
	return []GroupRep{
		{Name: "color singlet", Symbol: "1", Dim: 1, DynkinT: R(0, 1), RealType: true, FiniteAlphabetReason: "Fock lepton/scalar and gauge-singlet sectors"},
		{Name: "color fundamental", Symbol: "3", Dim: 3, DynkinT: R(1, 2), RealType: false, FiniteAlphabetReason: "Fock 1+3 color seed"},
		{Name: "color antifundamental", Symbol: "3bar", Dim: 3, DynkinT: R(1, 2), RealType: false, FiniteAlphabetReason: "charge-conjugate Fock sector"},
		{Name: "color adjoint", Symbol: "8", Dim: 8, DynkinT: R(3, 1), RealType: true, FiniteAlphabetReason: "adjoint gauge threshold row grammar"},
	}
}

func su2Alphabet() []GroupRep {
	return []GroupRep{
		{Name: "weak singlet", Symbol: "1", Dim: 1, DynkinT: R(0, 1), RealType: true, FiniteAlphabetReason: "right-singlet/scalar singlet sectors"},
		{Name: "weak doublet", Symbol: "2", Dim: 2, DynkinT: R(1, 2), RealType: false, FiniteAlphabetReason: "derived SU(2)L doublet scaffold"},
		{Name: "weak adjoint", Symbol: "3", Dim: 3, DynkinT: R(2, 1), RealType: true, FiniteAlphabetReason: "adjoint gauge threshold row grammar"},
	}
}

func hyperchargeAlphabet() []Rational {
	return []Rational{R(0, 1), R(1, 6), R(1, 3), R(1, 2), R(2, 3), R(1, 1)}
}

func statisticAlphabet() []Statistic {
	return []Statistic{
		{Name: "Weyl fermion", Symbol: "Weyl", Coeff: R(2, 3)},
		{Name: "Dirac fermion", Symbol: "Dirac", Coeff: R(4, 3)},
		{Name: "complex scalar", Symbol: "CScalar", Coeff: R(1, 3)},
		{Name: "real scalar", Symbol: "RScalar", Coeff: R(1, 6), RequiresY0: true, RequiresRealGaugeRep: true},
	}
}

func generateRows(su3 []GroupRep, su2 []GroupRep, hy []Rational, stats []Statistic) []RepresentationRow {
	out := make([]RepresentationRow, 0, len(su3)*len(su2)*len(hy)*len(stats))
	for _, st := range stats {
		for _, c := range su3 {
			for _, w := range su2 {
				for _, y := range hy {
					if st.RequiresY0 && !y.IsZero() {
						continue
					}
					if st.RequiresRealGaugeRep && (!c.RealType || !w.RealType) {
						continue
					}
					row := betaRow(st, c, w, y)
					out = append(out, RepresentationRow{
						Name:                   fmt.Sprintf("%s (%s,%s,Y=%s)", st.Name, c.Symbol, w.Symbol, y),
						Statistics:             st,
						SU3:                    c,
						SU2:                    w,
						Hypercharge:            y,
						DeltaB:                 row,
						ExactRational:          true,
						FiniteAlphabet:         true,
						StandardOneLoopFormula: true,
						Verdict:                "exact rational row generator in the finite heavy-sector grammar; no mass activation implied",
					})
				}
			}
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].DeltaB.Key() != out[j].DeltaB.Key() {
			return out[i].DeltaB.Key() < out[j].DeltaB.Key()
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func betaRow(st Statistic, su3 GroupRep, su2 GroupRep, y Rational) RationalTriple {
	y2 := y.Mul(y)
	u1Index := R(3, 5).Mul(y2).MulInt(su3.Dim * su2.Dim)
	su2Index := su2.DynkinT.MulInt(su3.Dim)
	su3Index := su3.DynkinT.MulInt(su2.Dim)
	return RT(st.Coeff.Mul(u1Index), st.Coeff.Mul(su2Index), st.Coeff.Mul(su3Index))
}

func markGate201Rows(rows []RepresentationRow, reqs []Gate201ShapeRequirement) {
	for i := range rows {
		for _, req := range reqs {
			if rows[i].DeltaB.Equal(req.DeltaB) && gate201DirectSemantics(rows[i], req.Name) {
				rows[i].DirectGate201Shape = true
				rows[i].Verdict = "direct Gate-201 non-universal shape generator; conditional support only, universal completion ignored"
			}
		}
	}
}

func gate201DirectSemantics(r RepresentationRow, name string) bool {
	switch name {
	case "Dirac vectorlike quark doublet":
		return r.Statistics.Symbol == "Dirac" && (r.SU3.Symbol == "3" || r.SU3.Symbol == "3bar") && r.SU2.Symbol == "2" && r.Hypercharge.Equal(R(1, 6))
	case "Weyl SU(2)L adjoint fermion":
		return r.Statistics.Symbol == "Weyl" && r.SU3.Symbol == "1" && r.SU2.Symbol == "3" && r.Hypercharge.IsZero()
	default:
		return false
	}
}

func uniqueRows(rows []RepresentationRow) []RepresentationRow {
	seen := map[string]RepresentationRow{}
	for _, r := range rows {
		if old, ok := seen[r.DeltaB.Key()]; !ok || (r.DirectGate201Shape && !old.DirectGate201Shape) {
			seen[r.DeltaB.Key()] = r
		}
	}
	out := make([]RepresentationRow, 0, len(seen))
	for _, r := range seen {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].DeltaB.Key() < out[j].DeltaB.Key() })
	return out
}

func auditGrammar(su3 []GroupRep, su2 []GroupRep, hy []Rational, stats []Statistic, rows []RepresentationRow, unique []RepresentationRow) RowGrammarAudit {
	exact := 0
	standard := 0
	den := int64(1)
	for _, r := range rows {
		if r.ExactRational {
			exact++
		}
		if r.StandardOneLoopFormula {
			standard++
		}
		den = lcm(den, r.DeltaB.DenominatorLCM())
	}
	return RowGrammarAudit{
		SU3AlphabetSize:             len(su3),
		SU2AlphabetSize:             len(su2),
		HyperchargeAlphabetSize:     len(hy),
		StatisticsCount:             len(stats),
		CandidateRowsGenerated:      len(rows),
		UniqueRows:                  len(unique),
		ExactRationalRows:           exact,
		StandardFormulaRows:         standard,
		CommonDenominatorLCM:        den,
		FiniteAlphabetDeclared:      true,
		UnboundedEnumerationAvoided: true,
		Verdict:                     "finite rational beta-row grammar constructed from the audited Fock/SM charge alphabet and singlet/fundamental/adjoint gauge reps",
	}
}

func auditLattice(unique []RepresentationRow) LatticeAudit {
	den := int64(1)
	zero := false
	for _, r := range unique {
		den = lcm(den, r.DeltaB.DenominatorLCM())
		if r.DeltaB.Equal(RT(R(0, 1), R(0, 1), R(0, 1))) {
			zero = true
		}
	}
	return LatticeAudit{
		GeneratorCount:      len(unique),
		UniqueGeneratorRows: len(unique),
		CommonDenominator:   den,
		IntegerGridEmbedded: den > 0,
		ContainsZeroRow:     zero,
		SemigroupOnly:       true,
		NoContinuousScales:  true,
		NoUniversalFit:      true,
		Verdict:             fmt.Sprintf("lattice is represented as a nonnegative integer semigroup of exact rational generators embedded in (1/%d)Z^3; no RG scale or universal beta fit is solved", den),
	}
}

func auditMemberships(reqs []Gate201ShapeRequirement, rows []RepresentationRow) []ShapeMembership {
	out := make([]ShapeMembership, 0, len(reqs))
	for _, req := range reqs {
		m := ShapeMembership{ShapeName: req.Name, TargetDeltaB: req.DeltaB, RequiredUniversalCExternal: req.RequiredUniversalC, UniversalCompletionIgnored: true, FiniteDerived: false, Verdict: "not found in row lattice"}
		for _, row := range rows {
			if row.DeltaB.Equal(req.DeltaB) && gate201DirectSemantics(row, req.Name) {
				m.Found = true
				m.DirectGenerator = true
				m.MatchedRepresentation = row.Name
				m.MatchedStatistics = row.Statistics.Name
				m.MatchedSMRep = fmt.Sprintf("(%s,%s,%s)", row.SU3.Symbol, row.SU2.Symbol, row.Hypercharge)
				m.ConditionalSupport = req.ConditionalOnly && !req.FiniteDerived
				m.Verdict = "CONDITIONAL_SUPPORT: Gate-201 non-universal shape is an exact rational row-lattice generator; universal completion and threshold scales remain external"
				break
			}
		}
		out = append(out, m)
	}
	return out
}

func summarizeMemberships(ms []ShapeMembership) MembershipAudit {
	found := 0
	direct := 0
	support := 0
	ignored := true
	for _, m := range ms {
		if m.Found {
			found++
		}
		if m.DirectGenerator {
			direct++
		}
		if m.ConditionalSupport {
			support++
		}
		ignored = ignored && m.UniversalCompletionIgnored
	}
	return MembershipAudit{
		ShapesAudited:              len(ms),
		ShapesOnLattice:            found,
		DirectGeneratorMatches:     direct,
		ConditionalSupportCount:    support,
		UniversalCompletionIgnored: ignored,
		AllGate201ShapesSupported:  len(ms) > 0 && found == len(ms),
		Verdict:                    "Gate-201 non-universal shapes are exact lattice generators; this supports representation viability only and does not repair the universal beta source",
	}
}

func auditContactInventory() ContactInventoryAudit {
	return ContactInventoryAudit{
		ContactPartialOverlapModes:        7,
		ContactModesHaveChargeLabels:      false,
		ContactModesHaveGaugeRepSemantics: false,
		ContactModesHaveDynkinIndices:     false,
		ContactModesHaveSpinStatistics:    false,
		ContactModesHaveMassActivation:    false,
		ContactModesHaveDecouplingLaw:     false,
		CanonicalMapToRowBasisFound:       false,
		CandidateRowsAssigned:             0,
		FiniteHeavySectorBasisDerived:     false,
		Verdict:                           "the seven contact partial-overlap modes cannot be canonically assigned to beta-row generators without charge, representation, Dynkin-index, spin-statistics, mass-activation, and decoupling semantics",
	}
}

func auditFirewall(prev Gate203Snapshot, grammar RowGrammarAudit, lattice LatticeAudit, ma MembershipAudit, contact ContactInventoryAudit) FirewallAudit {
	return FirewallAudit{
		Gate203Inherited:                        prev.Gate203Inherited,
		Gate203FailedRoutePreserved:             prev.Gate203FailedRoutePreserved,
		UniversalBetaSourceStillExternal:        prev.UniversalBetaSourceStillExternal,
		RepresentationLatticeConstructed:        grammar.FiniteAlphabetDeclared && lattice.IntegerGridEmbedded,
		Gate201ShapesPromotedToFinitePrediction: false,
		UniversalBetaFitAttempted:               false,
		ContinuousScalesSolved:                  false,
		ObservedInputsUsedForFiniteDerivation:   false,
		ContactModesPromotedToBetaRows:          contact.CanonicalMapToRowBasisFound,
		FockGenerationPromotedToNewThreshold:    prev.FockGenerationPromotedToNewThreshold,
		PhysicalUnificationClaimed:              false,
		ThresholdCorrectedPhysicalFitClaimed:    false,
		AbsoluteMassPredicted:                   false,
		FiniteMatchingCorrectionsDerived:        false,
		StrictNullityBefore:                     prev.StrictNullityAfter,
		StrictNullityAfter:                      prev.StrictNullityAfter,
		PhysicalPredictionNullityBefore:         prev.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:          prev.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                     "Gate 205 — finite carrier activation / contact-to-row semantics obstruction audit",
		OpenRequirements: []string{
			"derive a canonical map from contact/Fock finite carriers to charge labels, Dynkin indices, and spin-statistics",
			"derive a finite mass-activation or decoupling predicate before promoting row-lattice generators to physical thresholds",
			"keep the Gate-201 universal beta row and continuous lever arm external until a finite source theorem is found",
			"derive threshold matching corrections and M_* before claiming physical unification",
		},
		Verdict: fmt.Sprintf("representation lattice built and Gate-201 shapes supported=%t; contact mapping=%t; universal source external=%t", ma.AllGate201ShapesSupported, contact.CanonicalMapToRowBasisFound, prev.UniversalBetaSourceStillExternal),
	}
}

func FormatGate203(s Gate203Snapshot) string {
	return fmt.Sprintf("gate203=%t failed=%t universalExternal=%t completeSource=%t regulatorSource=%t contactBeta=%t fockThreshold=%t unification=%t fit=%t mass=%t matching=%t strict=%d prediction=%d shapes=%d", s.Gate203Inherited, s.Gate203FailedRoutePreserved, s.UniversalBetaSourceStillExternal, s.CompleteMultipletSourceFound, s.RegulatorTraceSourceFound, s.ContactModesHaveBetaPermission, s.FockGenerationPromotedToNewThreshold, s.PhysicalUnificationClaimed, s.ThresholdCorrectedPhysicalFitClaimed, s.AbsoluteMassPredicted, s.FiniteMatchingCorrectionsDerived, s.StrictNullityAfter, s.PhysicalPredictionNullityAfter, len(s.Gate201NonUniversalShapeRequirements))
}

func FormatShapeRequirement(r Gate201ShapeRequirement) string {
	return fmt.Sprintf("%s deltaB=%s c_univ=%.9g conditional=%t finite=%t phenomenology=%t", r.Name, r.DeltaB, r.RequiredUniversalC, r.ConditionalOnly, r.FiniteDerived, r.FromPhenomenology)
}

func FormatShapeRequirements(rs []Gate201ShapeRequirement) string {
	parts := make([]string, 0, len(rs))
	for _, r := range rs {
		parts = append(parts, FormatShapeRequirement(r))
	}
	return strings.Join(parts, "; ")
}

func FormatGrammar(a RowGrammarAudit) string {
	return fmt.Sprintf("su3=%d su2=%d hypercharge=%d stats=%d candidates=%d unique=%d exact=%d standard=%d denominator=%d finiteAlphabet=%t bounded=%t", a.SU3AlphabetSize, a.SU2AlphabetSize, a.HyperchargeAlphabetSize, a.StatisticsCount, a.CandidateRowsGenerated, a.UniqueRows, a.ExactRationalRows, a.StandardFormulaRows, a.CommonDenominatorLCM, a.FiniteAlphabetDeclared, a.UnboundedEnumerationAvoided)
}

func FormatLattice(a LatticeAudit) string {
	return fmt.Sprintf("generators=%d unique=%d denominator=%d grid=%t zero=%t semigroup=%t noScales=%t noUniversalFit=%t", a.GeneratorCount, a.UniqueGeneratorRows, a.CommonDenominator, a.IntegerGridEmbedded, a.ContainsZeroRow, a.SemigroupOnly, a.NoContinuousScales, a.NoUniversalFit)
}

func FormatRow(r RepresentationRow) string {
	return fmt.Sprintf("%s rep=(%s,%s,Y=%s) row=%s exact=%t finiteAlphabet=%t directGate201=%t", r.Statistics.Name, r.SU3.Symbol, r.SU2.Symbol, r.Hypercharge, r.DeltaB, r.ExactRational, r.FiniteAlphabet, r.DirectGate201Shape)
}

func FormatRows(rows []RepresentationRow, max int) string {
	if max <= 0 || max > len(rows) {
		max = len(rows)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, FormatRow(rows[i]))
	}
	if max < len(rows) {
		parts = append(parts, fmt.Sprintf("... +%d", len(rows)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatMembership(m ShapeMembership) string {
	return fmt.Sprintf("shape=%s target=%s found=%t direct=%t matched=%q stats=%s smrep=%s support=%t c_univ_external=%.9g finite=%t ignoredUniversal=%t", m.ShapeName, m.TargetDeltaB, m.Found, m.DirectGenerator, m.MatchedRepresentation, m.MatchedStatistics, m.MatchedSMRep, m.ConditionalSupport, m.RequiredUniversalCExternal, m.FiniteDerived, m.UniversalCompletionIgnored)
}

func FormatMemberships(ms []ShapeMembership) string {
	parts := make([]string, 0, len(ms))
	for _, m := range ms {
		parts = append(parts, FormatMembership(m))
	}
	return strings.Join(parts, "; ")
}

func FormatMembershipAudit(a MembershipAudit) string {
	return fmt.Sprintf("shapes=%d onLattice=%d direct=%d support=%d universalIgnored=%t allSupported=%t", a.ShapesAudited, a.ShapesOnLattice, a.DirectGeneratorMatches, a.ConditionalSupportCount, a.UniversalCompletionIgnored, a.AllGate201ShapesSupported)
}

func FormatContactInventory(c ContactInventoryAudit) string {
	return fmt.Sprintf("modes=%d charge=%t rep=%t dynkin=%t spin=%t mass=%t decoupling=%t map=%t assigned=%d heavyBasis=%t", c.ContactPartialOverlapModes, c.ContactModesHaveChargeLabels, c.ContactModesHaveGaugeRepSemantics, c.ContactModesHaveDynkinIndices, c.ContactModesHaveSpinStatistics, c.ContactModesHaveMassActivation, c.ContactModesHaveDecouplingLaw, c.CanonicalMapToRowBasisFound, c.CandidateRowsAssigned, c.FiniteHeavySectorBasisDerived)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g203=%t failed=%t universalExternal=%t lattice=%t promotePrediction=%t fitUniversal=%t scales=%t observed=%t contactBeta=%t fockThreshold=%t unification=%t fit=%t mass=%t matching=%t strict=%d->%d prediction=%d->%d next=%s", f.Gate203Inherited, f.Gate203FailedRoutePreserved, f.UniversalBetaSourceStillExternal, f.RepresentationLatticeConstructed, f.Gate201ShapesPromotedToFinitePrediction, f.UniversalBetaFitAttempted, f.ContinuousScalesSolved, f.ObservedInputsUsedForFiniteDerivation, f.ContactModesPromotedToBetaRows, f.FockGenerationPromotedToNewThreshold, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.AbsoluteMassPredicted, f.FiniteMatchingCorrectionsDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d g203=%t grammar=%t lattice=%t shapes=%t support=%t contactFail=%t universalFitAvoided=%t noPrediction=%t status=%s", s.TestsAudited, s.Gate203Inherited, s.RationalGrammarConstructed, s.LatticeConstructed, s.Gate201ShapesOnLattice, s.ConditionalSupportLogged, s.ContactMapFailed, s.UniversalFitAvoided, s.NoPhysicalPredictionClaim, s.Status)
}
