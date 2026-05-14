// Package twothresholdminimality implements Gate 212: two-threshold solution
// minimality / finite-origin and multiplet-parentage audit.
//
// Gate 211 found 44 ordered two-threshold witnesses for the u*=1 branch.  Gate
// 212 deliberately does not rerun the RG fit.  It audits the degeneracy: do the
// viable pairs have a finite-origin combinatorial signature, a canonical parent
// multiplet, or a derived threshold-spectrum metric selecting a unique physical
// pair?  If not, the gate must log a uniqueness obstruction and require a future
// ThresholdSpectrumSeal rather than promoting a ranked witness into a prediction.
package twothresholdminimality

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
	"github.com/bagherbal/asha-engine/pkg/bridge/twothresholdviability"
)

const (
	StatusFailedCanonicalUniqueness = "FAILED_ROUTE_CANONICAL_THRESHOLD_UNIQUENESS"
	StatusConditionalUnique         = "CONDITIONAL_UNIQUE_THRESHOLD_SPECTRUM"

	contactModeCount = 7
	bSectorGap       = 0.1024649212
	closeSplitTol    = 0.5
)

var contactPartialModes = []float64{0.8975350788094513, 0.7440966379808388, 0.6666666666666659, 0.49999999999999956, 0.44112275728436634, 0.333333333333333, 0.283912192592006}

type Gate211Snapshot struct {
	Gate211Inherited                bool
	ConditionalViabilityInherited   bool
	TopologicalViableOrderedPairs   int
	CentroidViablePairs             int
	LeptoquarkDynamicsSealInherited bool
	EmpiricalCarrierSealInherited   bool
	PhysicalPredictionClaimed       bool
	FiniteCarrierOriginDerived      bool
	MatchingCorrectionsDerived      bool
	TruthStatement                  string
}

func SnapshotFromGate211(a twothresholdviability.Analysis) Gate211Snapshot {
	return Gate211Snapshot{
		Gate211Inherited:                true,
		ConditionalViabilityInherited:   a.Summary.Status == twothresholdviability.StatusConditionalViable && a.Summary.ViableTopological > 0,
		TopologicalViableOrderedPairs:   a.Summary.ViableTopological,
		CentroidViablePairs:             a.Summary.ViableCentroid,
		LeptoquarkDynamicsSealInherited: a.BaryonAnomaly.LeptoquarkDynamicsSealInherited,
		EmpiricalCarrierSealInherited:   true,
		PhysicalPredictionClaimed:       a.Firewall.PhysicalPredictionClaimed,
		FiniteCarrierOriginDerived:      a.ContactMatch.ViableRowsPromotedFromContact,
		MatchingCorrectionsDerived:      a.Firewall.MatchingCorrectionsDerived,
		TruthStatement:                  a.TruthStatement,
	}
}

type RowSemantics struct {
	Name                    string
	SMRepresentation        string
	Statistic               string
	SU3Symbol               string
	SU2Symbol               string
	Hypercharge             string
	SU3Dim                  int
	SU2Dim                  int
	GaugeDimension          int
	WeylEquivalentDimension int
	DeltaB                  representationrowlattice.RationalTriple
}

type PairClass struct {
	Key                          string
	OrderedMultiplicity          int
	Representative               twothresholdviability.PairSolution
	RowA                         RowSemantics
	RowB                         RowSemantics
	TotalGaugeDimension          int
	TotalWeylEquivalentDimension int
	DeltaL                       float64
	MeanThresholdLog             float64
	MStarGeV                     float64
	TotalDeltaBNorm              float64
	GUTRangeDistance             float64
	IndividualDimensionHits7     bool
	TotalDimensionHits7          bool
	WeylDimensionHits7           bool
	BGapSplitMatch               bool
	ContactSplitMatch            bool
	ContactModeCountMatch        bool
	CanonicalFiniteOrigin        bool
	FiniteOriginVerdict          string
	ParentageHint                string
	CompleteParentDerived        bool
	MissingParentPartners        []string
	ParentageVerdict             string
}

type FiniteOriginAudit struct {
	OrderedPairsAudited             int
	UnorderedPairClasses            int
	ContactPartialOverlapModes      int
	BGap                            float64
	PairDimensionHitsSeven          int
	PairWeylDimensionHitsSeven      int
	IndividualRowDimensionHitsSeven int
	BGapSplitMatches                int
	ContactSplitMatches             int
	CanonicalFiniteOriginMatches    int
	CarrierActivationSealIntact     bool
	Verdict                         string
}

type MultipletParentageAudit struct {
	PairClassesAudited            int
	CloseSplitPairs               int
	MinDeltaL                     float64
	MaxDeltaL                     float64
	MeanDeltaL                    float64
	BestCloseSplitKey             string
	BestCloseSplitRows            string
	ExternalParentHints           int
	CompleteParentageDerived      int
	ThresholdSplittingRuleDerived bool
	UnifiedParentGaugeImported    bool
	Verdict                       string
}

type DegeneracyAudit struct {
	OrderedViablePairs            int
	PhysicalUnorderedClasses      int
	Gate211RankedBestExists       bool
	Gate211RankingIsFiniteMetric  bool
	UniqueAfterFiniteOrigin       bool
	UniqueAfterParentage          bool
	CanonicalUniquePairFound      bool
	ThresholdSpectrumSealRequired bool
	ProposedSealName              string
	BindingObstruction            string
	Verdict                       string
}

type FirewallAudit struct {
	Gate211Inherited                      bool
	LeptoquarkDynamicsSealInherited       bool
	EmpiricalCarrierSealInherited         bool
	EmpiricalLedgerQuarantined            bool
	ContactModesPromotedToCarriers        bool
	BGapPromotedToMass                    bool
	SU5OrPatiSalamGaugeImported           bool
	MatchingCorrectionsDerived            bool
	PhysicalPredictionClaimed             bool
	ProtonLifetimeComputed                bool
	UniqueThresholdSpectrumClaimed        bool
	ThresholdSpectrumSealDeclaredAsFuture bool
	RecommendedNextGate                   string
	OpenRequirements                      []string
	Verdict                               string
}

type Summary struct {
	TestsAudited                  int
	Gate211ViabilityInherited     bool
	OrderedViablePairs            int
	UnorderedPairClasses          int
	FiniteOriginMatches           int
	CompleteParentageMatches      int
	CanonicalUniquePairFound      bool
	ThresholdSpectrumSealRequired bool
	Status                        string
	Comment                       string
}

type Analysis struct {
	Gate211         Gate211Snapshot
	Gate211Analysis twothresholdviability.Analysis
	PairClasses     []PairClass
	FiniteOrigin    FiniteOriginAudit
	Parentage       MultipletParentageAudit
	Degeneracy      DegeneracyAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g211, err := twothresholdviability.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(g211)
	})
	return defaultA, defaultErr
}

func Build(g211 twothresholdviability.Analysis) (Analysis, error) {
	snap := SnapshotFromGate211(g211)
	if !snap.Gate211Inherited || !snap.ConditionalViabilityInherited || snap.TopologicalViableOrderedPairs == 0 {
		return Analysis{}, fmt.Errorf("Gate 212 requires Gate 211 topological viable witnesses")
	}
	target := firstTarget(g211, "u_topological")
	if target.ViablePairs == 0 {
		return Analysis{}, fmt.Errorf("Gate 212 found no topological Gate 211 viable pairs")
	}
	classes := buildPairClasses(target.AllViableSolutions)
	finite := auditFiniteOrigin(classes, len(target.AllViableSolutions))
	parent := auditParentage(classes)
	degen := auditDegeneracy(snap, classes, finite, parent)
	fw := auditFirewall(snap, finite, parent, degen)
	status := StatusFailedCanonicalUniqueness
	if degen.CanonicalUniquePairFound {
		status = StatusConditionalUnique
	}
	summary := Summary{
		TestsAudited:                  7,
		Gate211ViabilityInherited:     snap.ConditionalViabilityInherited,
		OrderedViablePairs:            len(target.AllViableSolutions),
		UnorderedPairClasses:          len(classes),
		FiniteOriginMatches:           finite.CanonicalFiniteOriginMatches,
		CompleteParentageMatches:      parent.CompleteParentageDerived,
		CanonicalUniquePairFound:      degen.CanonicalUniquePairFound,
		ThresholdSpectrumSealRequired: degen.ThresholdSpectrumSealRequired,
		Status:                        status,
		Comment:                       "Gate 212 audits whether the 44 ordered Gate-211 witnesses reduce to a unique finite-origin or parent-multiplet threshold spectrum. The audit preserves the carrier and leptoquark seals and refuses to promote ranking into derivation.",
	}
	truth := buildTruth(summary, finite, parent, degen)
	return Analysis{Gate211: snap, Gate211Analysis: g211, PairClasses: classes, FiniteOrigin: finite, Parentage: parent, Degeneracy: degen, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func firstTarget(a twothresholdviability.Analysis, name string) twothresholdviability.TargetAudit {
	for _, t := range a.TargetAudits {
		if t.Target.Name == name {
			return t
		}
	}
	return twothresholdviability.TargetAudit{}
}

func buildPairClasses(solutions []twothresholdviability.PairSolution) []PairClass {
	m := map[string]*PairClass{}
	for _, s := range solutions {
		r1 := parseRow(s.Row1Name, s.Row1Rep, s.Row1DeltaB)
		r2 := parseRow(s.Row2Name, s.Row2Rep, s.Row2DeltaB)
		keyA := r1.DeltaB.Key() + "|" + r1.Name
		keyB := r2.DeltaB.Key() + "|" + r2.Name
		key := keyA + "||" + keyB
		if keyB < keyA {
			key = keyB + "||" + keyA
		}
		pc, ok := m[key]
		if !ok {
			a, b := r1, r2
			if keyB < keyA {
				a, b = r2, r1
			}
			pc = &PairClass{Key: key, Representative: s, RowA: a, RowB: b}
			pc.TotalGaugeDimension = a.GaugeDimension + b.GaugeDimension
			pc.TotalWeylEquivalentDimension = a.WeylEquivalentDimension + b.WeylEquivalentDimension
			pc.DeltaL = math.Abs(s.LB2 - s.LB1)
			pc.MeanThresholdLog = 0.5 * (s.LB1 + s.LB2)
			pc.MStarGeV = s.MStarGeV
			pc.TotalDeltaBNorm = s.TotalDeltaBNorm
			pc.GUTRangeDistance = s.GUTRangeDistance
			pc.IndividualDimensionHits7 = a.GaugeDimension == contactModeCount || b.GaugeDimension == contactModeCount
			pc.TotalDimensionHits7 = pc.TotalGaugeDimension == contactModeCount
			pc.WeylDimensionHits7 = pc.TotalWeylEquivalentDimension == contactModeCount
			pc.BGapSplitMatch = nearly(pc.DeltaL, bSectorGap, 1e-9)
			pc.ContactSplitMatch = matchesAny(pc.DeltaL, contactPartialModes, 1e-9)
			pc.ContactModeCountMatch = pc.TotalGaugeDimension == contactModeCount || pc.TotalWeylEquivalentDimension == contactModeCount
			pc.CanonicalFiniteOrigin = pc.ContactModeCountMatch && (pc.BGapSplitMatch || pc.ContactSplitMatch)
			pc.FiniteOriginVerdict = finiteOriginVerdict(*pc)
			pc.ParentageHint, pc.MissingParentPartners = parentageHint(a, b)
			pc.CompleteParentDerived = false
			pc.ParentageVerdict = parentageVerdict(*pc)
			m[key] = pc
		}
		pc.OrderedMultiplicity++
		// Keep the Gate-211 best-ranked representative when present.
		if solutionRankLess(s, pc.Representative) {
			pc.Representative = s
			pc.DeltaL = math.Abs(s.LB2 - s.LB1)
			pc.MeanThresholdLog = 0.5 * (s.LB1 + s.LB2)
			pc.MStarGeV = s.MStarGeV
			pc.TotalDeltaBNorm = s.TotalDeltaBNorm
			pc.GUTRangeDistance = s.GUTRangeDistance
		}
	}
	out := make([]PairClass, 0, len(m))
	for _, pc := range m {
		out = append(out, *pc)
	}
	sort.Slice(out, func(i, j int) bool {
		if math.Abs(out[i].TotalDeltaBNorm-out[j].TotalDeltaBNorm) > 1e-12 {
			return out[i].TotalDeltaBNorm < out[j].TotalDeltaBNorm
		}
		if math.Abs(out[i].GUTRangeDistance-out[j].GUTRangeDistance) > 1e-12 {
			return out[i].GUTRangeDistance < out[j].GUTRangeDistance
		}
		return out[i].Key < out[j].Key
	})
	return out
}

var repRE = regexp.MustCompile(`\(([0-9a-zA-Z]+),([0-9a-zA-Z]+),Y=([^\)]+)\)`)

func parseRow(name, rep string, db representationrowlattice.RationalTriple) RowSemantics {
	stat := "unknown"
	switch {
	case strings.HasPrefix(name, "Dirac fermion"):
		stat = "Dirac fermion"
	case strings.HasPrefix(name, "Weyl fermion"):
		stat = "Weyl fermion"
	case strings.HasPrefix(name, "complex scalar"):
		stat = "complex scalar"
	case strings.HasPrefix(name, "real scalar"):
		stat = "real scalar"
	}
	su3, su2, y := "?", "?", "?"
	if m := repRE.FindStringSubmatch(rep); len(m) == 4 {
		su3, su2, y = m[1], m[2], m[3]
	}
	d3 := repDim(su3)
	d2 := repDim(su2)
	gd := d3 * d2
	factor := 1
	switch stat {
	case "Dirac fermion":
		factor = 2
	case "complex scalar":
		factor = 1
	case "real scalar":
		factor = 1
	}
	return RowSemantics{Name: name, SMRepresentation: rep, Statistic: stat, SU3Symbol: su3, SU2Symbol: su2, Hypercharge: y, SU3Dim: d3, SU2Dim: d2, GaugeDimension: gd, WeylEquivalentDimension: gd * factor, DeltaB: db}
}

func repDim(sym string) int {
	switch sym {
	case "1":
		return 1
	case "2":
		return 2
	case "3", "3bar":
		return 3
	case "8":
		return 8
	default:
		n, err := strconv.Atoi(sym)
		if err == nil && n > 0 {
			return n
		}
		return 0
	}
}

func solutionRankLess(a, b twothresholdviability.PairSolution) bool {
	if math.Abs(a.TotalDeltaBNorm-b.TotalDeltaBNorm) > 1e-12 {
		return a.TotalDeltaBNorm < b.TotalDeltaBNorm
	}
	if math.Abs(a.GUTRangeDistance-b.GUTRangeDistance) > 1e-12 {
		return a.GUTRangeDistance < b.GUTRangeDistance
	}
	return a.Row1Name+a.Row2Name < b.Row1Name+b.Row2Name
}

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func matchesAny(x float64, ys []float64, tol float64) bool {
	for _, y := range ys {
		if nearly(x, y, tol) {
			return true
		}
	}
	return false
}

func finiteOriginVerdict(pc PairClass) string {
	if pc.CanonicalFiniteOrigin {
		return "CANONICAL_FINITE_ORIGIN_CANDIDATE: exact dimension/spectral match found; still requires carrier semantics"
	}
	return fmt.Sprintf("no finite-origin match: gaugeDim=%d weylEqDim=%d contactModes=%d ΔL=%.9g Bgap=%.9g; charge/spin/mass semantics remain sealed", pc.TotalGaugeDimension, pc.TotalWeylEquivalentDimension, contactModeCount, pc.DeltaL, bSectorGap)
}

func parentageHint(a, b RowSemantics) (string, []string) {
	// These are external branching-pattern preflight hints only.  The engine has
	// no derived SU(5), SO(10), or Pati-Salam parent gauge connection here.
	reps := map[string]bool{canonicalRep(a): true, canonicalRep(b): true}
	if reps["(1,3,Y=1)"] {
		return "external SU(5)-15 fragment hint", []string{"(3,2,Y=1/6)", "(6,1,Y=2/3) or conjugate partner"}
	}
	if reps["(1,3,Y=0)"] || reps["(8,1,Y=0)"] {
		return "external SU(5)-24 fragment hint", []string{"(8,1,Y=0)", "(1,3,Y=0)", "(3,2,Y=5/6)", "(3bar,2,Y=5/6)"}
	}
	if reps["(8,2,Y=1/2)"] {
		return "external SU(5)-45/50/70 style color-octet-doublet fragment hint", []string{"additional electroweak doublet/singlet and colored partners depend on parent choice"}
	}
	if strings.HasPrefix(a.SU3Symbol, "8") || strings.HasPrefix(b.SU3Symbol, "8") {
		return "external colored-adjoint fragment hint", []string{"complete parent branching not derived by ASHA finite core"}
	}
	return "none", nil
}

func canonicalRep(r RowSemantics) string {
	return fmt.Sprintf("(%s,%s,Y=%s)", r.SU3Symbol, r.SU2Symbol, r.Hypercharge)
}

func parentageVerdict(pc PairClass) string {
	if pc.CompleteParentDerived {
		return "complete parent multiplet derived"
	}
	if pc.ParentageHint != "none" {
		return fmt.Sprintf("%s only; missing partners=%v; no ASHA-derived parent gauge connection or branching theorem", pc.ParentageHint, pc.MissingParentPartners)
	}
	return "no parent multiplet hint and no derived parentage theorem"
}

func auditFiniteOrigin(classes []PairClass, ordered int) FiniteOriginAudit {
	dim7, weyl7, indiv7, bgap, contact, canon := 0, 0, 0, 0, 0, 0
	for _, pc := range classes {
		if pc.TotalDimensionHits7 {
			dim7++
		}
		if pc.WeylDimensionHits7 {
			weyl7++
		}
		if pc.IndividualDimensionHits7 {
			indiv7++
		}
		if pc.BGapSplitMatch {
			bgap++
		}
		if pc.ContactSplitMatch {
			contact++
		}
		if pc.CanonicalFiniteOrigin {
			canon++
		}
	}
	return FiniteOriginAudit{
		OrderedPairsAudited:             ordered,
		UnorderedPairClasses:            len(classes),
		ContactPartialOverlapModes:      contactModeCount,
		BGap:                            bSectorGap,
		PairDimensionHitsSeven:          dim7,
		PairWeylDimensionHitsSeven:      weyl7,
		IndividualRowDimensionHitsSeven: indiv7,
		BGapSplitMatches:                bgap,
		ContactSplitMatches:             contact,
		CanonicalFiniteOriginMatches:    canon,
		CarrierActivationSealIntact:     true,
		Verdict:                         fmt.Sprintf("no canonical finite-origin selector: %d unordered classes, %d exact dimension-7 pair matches, %d B-gap split matches, %d contact-split matches; contact modes remain semantic anchors only", len(classes), dim7, bgap, contact),
	}
}

func auditParentage(classes []PairClass) MultipletParentageAudit {
	minDL := math.Inf(1)
	maxDL := 0.0
	sum := 0.0
	close := 0
	hints := 0
	bestKey := ""
	bestRows := ""
	for _, pc := range classes {
		if pc.DeltaL < minDL {
			minDL = pc.DeltaL
			bestKey = pc.Key
			bestRows = pc.RowA.SMRepresentation + " + " + pc.RowB.SMRepresentation
		}
		if pc.DeltaL > maxDL {
			maxDL = pc.DeltaL
		}
		sum += pc.DeltaL
		if pc.DeltaL <= closeSplitTol {
			close++
		}
		if pc.ParentageHint != "none" {
			hints++
		}
	}
	mean := 0.0
	if len(classes) > 0 {
		mean = sum / float64(len(classes))
	}
	return MultipletParentageAudit{
		PairClassesAudited:            len(classes),
		CloseSplitPairs:               close,
		MinDeltaL:                     minDL,
		MaxDeltaL:                     maxDL,
		MeanDeltaL:                    mean,
		BestCloseSplitKey:             bestKey,
		BestCloseSplitRows:            bestRows,
		ExternalParentHints:           hints,
		CompleteParentageDerived:      0,
		ThresholdSplittingRuleDerived: false,
		UnifiedParentGaugeImported:    false,
		Verdict:                       fmt.Sprintf("threshold splitting gives %d close classes (ΔL≤%.3g), min ΔL=%.9g for %s, but no finite parent-gauge/branching theorem derives a complete parent multiplet", close, closeSplitTol, minDL, bestRows),
	}
}

func auditDegeneracy(s Gate211Snapshot, classes []PairClass, f FiniteOriginAudit, p MultipletParentageAudit) DegeneracyAudit {
	uniqueFinite := f.CanonicalFiniteOriginMatches == 1
	uniqueParent := p.CompleteParentageDerived == 1
	unique := uniqueFinite && uniqueParent
	return DegeneracyAudit{
		OrderedViablePairs:            s.TopologicalViableOrderedPairs,
		PhysicalUnorderedClasses:      len(classes),
		Gate211RankedBestExists:       len(classes) > 0,
		Gate211RankingIsFiniteMetric:  false,
		UniqueAfterFiniteOrigin:       uniqueFinite,
		UniqueAfterParentage:          uniqueParent,
		CanonicalUniquePairFound:      unique,
		ThresholdSpectrumSealRequired: !unique,
		ProposedSealName:              "ThresholdSpectrumSeal",
		BindingObstruction:            "finite-origin dimensions, B-sector/contact spectra, and external parentage hints do not define a canonical threshold-spectrum selector",
		Verdict:                       "FAILED_ROUTE: Gate-211 ranking is useful phenomenological ordering, but no finite algebraic metric uniquely selects one unordered pair class; exact heavy spectrum remains a sealed phenomenological boundary condition",
	}
}

func auditFirewall(s Gate211Snapshot, f FiniteOriginAudit, p MultipletParentageAudit, d DegeneracyAudit) FirewallAudit {
	return FirewallAudit{
		Gate211Inherited:                      s.Gate211Inherited,
		LeptoquarkDynamicsSealInherited:       s.LeptoquarkDynamicsSealInherited,
		EmpiricalCarrierSealInherited:         s.EmpiricalCarrierSealInherited,
		EmpiricalLedgerQuarantined:            true,
		ContactModesPromotedToCarriers:        f.CanonicalFiniteOriginMatches > 0,
		BGapPromotedToMass:                    false,
		SU5OrPatiSalamGaugeImported:           p.UnifiedParentGaugeImported,
		MatchingCorrectionsDerived:            s.MatchingCorrectionsDerived,
		PhysicalPredictionClaimed:             s.PhysicalPredictionClaimed,
		ProtonLifetimeComputed:                false,
		UniqueThresholdSpectrumClaimed:        d.CanonicalUniquePairFound,
		ThresholdSpectrumSealDeclaredAsFuture: d.ThresholdSpectrumSealRequired,
		RecommendedNextGate:                   "Gate 213 — ThresholdSpectrumSeal / matching-correction and two-loop stability preflight audit",
		OpenRequirements: []string{
			"derive or seal a threshold-spectrum selector before choosing one of the Gate-211 pair classes",
			"derive finite matching corrections for two separated thresholds",
			"audit two-loop stability and scheme dependence of the viable witnesses",
			"keep EmpiricalCarrierSeal and LeptoquarkDynamicsSeal active until native carrier dynamics are derived",
		},
		Verdict: "firewall preserved: finite contact/B-sector data are not promoted, external parent groups are not imported as dynamics, and no unique physical spectrum is claimed",
	}
}

func buildTruth(s Summary, f FiniteOriginAudit, p MultipletParentageAudit, d DegeneracyAudit) string {
	if s.CanonicalUniquePairFound {
		return "Gate 212 finds a canonical finite-origin and parentage selector for a unique Gate-211 two-threshold pair. This remains conditional on the active carrier seal."
	}
	return fmt.Sprintf("Gate 212 preserves the Gate-211 viable bridge but proves a canonical uniqueness obstruction. The 44 ordered viable witnesses reduce to %d unordered pair classes. None is selected by exact contact-mode dimension, B-sector/contact spectral matching, or an ASHA-derived parent-multiplet theorem. Close threshold splittings exist (%d classes with ΔL≤%.3g), but splitting is not a finite selector. Therefore a future %s is required before any specific two-threshold spectrum can be promoted beyond conditional phenomenology.", s.UnorderedPairClasses, p.CloseSplitPairs, closeSplitTol, d.ProposedSealName)
}

func FormatRow(r RowSemantics) string {
	return fmt.Sprintf("%s %s dim=%d weylEq=%d Δb=%s", r.Name, r.SMRepresentation, r.GaugeDimension, r.WeylEquivalentDimension, r.DeltaB)
}

func FormatPairClass(p PairClass) string {
	return fmt.Sprintf("rows=[%s; %s] ordered=%d gaugeDim=%d weylEq=%d ΔL=%.9g M*=% .9g finiteOrigin=%t parent=%q", p.RowA.SMRepresentation, p.RowB.SMRepresentation, p.OrderedMultiplicity, p.TotalGaugeDimension, p.TotalWeylEquivalentDimension, p.DeltaL, p.MStarGeV, p.CanonicalFiniteOrigin, p.ParentageHint)
}

func FormatFiniteOrigin(a FiniteOriginAudit) string {
	return fmt.Sprintf("ordered=%d unordered=%d contactModes=%d Bgap=%.10f dim7=%d weyl7=%d indiv7=%d bgapSplit=%d contactSplit=%d canonical=%d sealIntact=%t", a.OrderedPairsAudited, a.UnorderedPairClasses, a.ContactPartialOverlapModes, a.BGap, a.PairDimensionHitsSeven, a.PairWeylDimensionHitsSeven, a.IndividualRowDimensionHitsSeven, a.BGapSplitMatches, a.ContactSplitMatches, a.CanonicalFiniteOriginMatches, a.CarrierActivationSealIntact)
}

func FormatParentage(a MultipletParentageAudit) string {
	return fmt.Sprintf("classes=%d close=%d minΔL=%.9g maxΔL=%.9g meanΔL=%.9g best=%q hints=%d complete=%d splittingRule=%t importedParent=%t", a.PairClassesAudited, a.CloseSplitPairs, a.MinDeltaL, a.MaxDeltaL, a.MeanDeltaL, a.BestCloseSplitRows, a.ExternalParentHints, a.CompleteParentageDerived, a.ThresholdSplittingRuleDerived, a.UnifiedParentGaugeImported)
}

func FormatDegeneracy(a DegeneracyAudit) string {
	return fmt.Sprintf("ordered=%d unordered=%d rankedBest=%t finiteMetric=%t uniqueFinite=%t uniqueParent=%t unique=%t sealRequired=%t seal=%s obstruction=%q", a.OrderedViablePairs, a.PhysicalUnorderedClasses, a.Gate211RankedBestExists, a.Gate211RankingIsFiniteMetric, a.UniqueAfterFiniteOrigin, a.UniqueAfterParentage, a.CanonicalUniquePairFound, a.ThresholdSpectrumSealRequired, a.ProposedSealName, a.BindingObstruction)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate211=%t lqSeal=%t carrierSeal=%t ledger=%t contactPromoted=%t bgapMass=%t parentImported=%t matching=%t prediction=%t lifetime=%t uniqueClaim=%t futureSeal=%t next=%s", a.Gate211Inherited, a.LeptoquarkDynamicsSealInherited, a.EmpiricalCarrierSealInherited, a.EmpiricalLedgerQuarantined, a.ContactModesPromotedToCarriers, a.BGapPromotedToMass, a.SU5OrPatiSalamGaugeImported, a.MatchingCorrectionsDerived, a.PhysicalPredictionClaimed, a.ProtonLifetimeComputed, a.UniqueThresholdSpectrumClaimed, a.ThresholdSpectrumSealDeclaredAsFuture, a.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate211=%t ordered=%d unordered=%d finiteOrigin=%d parentage=%d unique=%t sealRequired=%t status=%s comment=%q", s.TestsAudited, s.Gate211ViabilityInherited, s.OrderedViablePairs, s.UnorderedPairClasses, s.FiniteOriginMatches, s.CompleteParentageMatches, s.CanonicalUniquePairFound, s.ThresholdSpectrumSealRequired, s.Status, s.Comment)
}
