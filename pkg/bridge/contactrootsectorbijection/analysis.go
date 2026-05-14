// Package contactrootsectorbijection implements Gate 278:
// Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit.
//
// Gate 277 selected the physical sector-level pairing {u,d}|{e,nu} using
// tau_eta and B_gap tags, but it did not select a contact resolvent root because
// the four quartic contact roots remained an unlabeled Galois orbit. Gate 278
// audits whether the internal arithmetic of those roots, together with the
// Morita multiplicity kappa_C:kappa_Q=1:3 and the B_gap/Majorana neutrino tag,
// supplies a native root-to-sector bijection.
//
// The result is intentionally firewalled. The roots are retrieved and their
// magnitude/order/projector diagnostics are audited. No root is natively null or
// Majorana-suppressed, the 1+3 multiplicity counts Hilbert trace sectors rather
// than individual quartic roots, and irreducibility prevents rational root or
// pair projectors over the base field. Thus no native bijection q_i <->
// {u,d,e,nu}, no resolvent-root selection, and no r_+/r_- branch locking is
// derived.
package contactrootsectorbijection

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE278-QUARTIC-ROOT-YUKAWA-SECTOR-BIJECTION-CONTACT-PROJECTOR-SEMANTICS-AUDIT"

	StatusRootsRetrieved          = "CONDITIONAL_SUPPORT_CONTACT_QUARTIC_ROOTS_RETRIEVED"
	StatusMagnitudeSieveCompleted = "CONDITIONAL_SUPPORT_CONTACT_ROOT_MAGNITUDE_SIEVE_COMPLETED"
	StatusConstraintsApplied      = "CONDITIONAL_SUPPORT_MORITA_AND_B_GAP_CONSTRAINTS_APPLIED_AS_AUDIT_TESTS"
	StatusResolventInherited      = "CONDITIONAL_SUPPORT_RESOLVENT_PAIRINGS_INHERITED"
	StatusFirewallsPreserved      = "CONDITIONAL_SUPPORT_ROOT_BIJECTION_FIREWALLS_PRESERVED"

	StatusFailedNoNullRoot             = "FAILED_ROUTE_NO_CONTACT_ROOT_IS_NATIVE_NULL_OR_MAJORANA_SUPPRESSED"
	StatusFailedMultiplicityNoLabel    = "FAILED_ROUTE_1_PLUS_3_MULTIPLICITY_DOES_NOT_LABEL_INDIVIDUAL_ROOTS"
	StatusFailedProjectorsMissing      = "FAILED_ROUTE_CONTACT_ROOT_PROJECTOR_SEMANTICS_NOT_DERIVED"
	StatusFailedRootSectorBijection    = "FAILED_ROUTE_ROOT_TO_YUKAWA_SECTOR_BIJECTION_MISSING"
	StatusFailedResolventRootSelection = "FAILED_ROUTE_CONTACT_RESOLVENT_ROOT_NOT_SELECTED"
	StatusFailedRBranchLock            = "FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED"
	StatusFailedHiggsRatio             = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

type ContactRoot struct {
	Label                    string
	Approx                   float64
	Interval                 string
	OrderRank                int
	DistanceToZero           float64
	NearZeroSuppressed       bool
	ClosestSimpleRational    string
	DistanceToSimpleRational float64
	CanBeSelectedByMagnitude bool
	Verdict                  string
}

type RootMagnitudeAudit struct {
	Polynomial                 string
	Roots                      []ContactRoot
	AllRootsO1                 bool
	AnyNativeNullRoot          bool
	OrderingAvailable          bool
	OrderingInvariant          bool
	MagnitudeBijectionDerived  bool
	BGapScaleComparableToRoots bool
	Verdict                    string
}

type Constraint struct {
	Name                 string
	SourceGate           string
	Data                 string
	ActsOn               string
	CanLabelRoot         bool
	CanSelectPairing     bool
	RequiresExtraFunctor bool
	Verdict              string
}

type ConstraintAudit struct {
	Constraints                 []Constraint
	MoritaMultiplicityAvailable bool
	BGapTagAvailable            bool
	TauEtaPairingAvailable      bool
	ConstraintsReachSectors     bool
	ConstraintsReachRoots       bool
	UsesObservedMasses          bool
	Verdict                     string
}

type ProjectorAudit struct {
	QuarticIrreducibleOverQ         bool
	ResolventIrreducibleOverQ       bool
	IndividualRootProjectorsOverQ   bool
	TwoPlusTwoPairProjectorsOverQ   bool
	RequiresSplittingField          bool
	RequiresResolventRootAdjunction bool
	RationalContactProjectorDerived bool
	Verdict                         string
}

type PairingCandidate struct {
	Label                   string
	RootPairing             string
	PairSums                [2]float64
	PairProducts            [2]float64
	PairMeanGap             float64
	CompatibleWithUD_ENU    bool
	SelectedByMultiplicity  bool
	SelectedByBGap          bool
	SelectedAsResolventRoot bool
	Verdict                 string
}

type PairingAudit struct {
	Candidates                []PairingCandidate
	TotalCandidates           int
	CompatibleWithSectorSplit int
	SelectedPairings          int
	UniqueRootPairing         bool
	SelectedRootPairing       string
	Verdict                   string
}

type BijectionAudit struct {
	TotalRootSectorBijections    int
	BijectionsAfterSectorPairing int
	BijectionsAfterBGapNuTag     int
	BijectionsAfterUDTauTag      int
	UniqueBijection              bool
	DerivedAssignment            map[string]string
	Verdict                      string
}

type BranchProjectionAudit struct {
	RPlus                   float64
	RMinus                  float64
	AbsYOverXPlus           float64
	AbsYOverXMinus          float64
	ResolventRootSelected   bool
	RootPairingToRBranchMap bool
	UniqueAmplitudeBranch   bool
	SelectedBranch          string
	Verdict                 string
}

type FirewallAudit struct {
	NoObservedMassesUsed      bool
	NoCKMPMNSUsed             bool
	NoEmpiricalYukawaInserted bool
	NoRootOrderingPromotion   bool
	NoArbitraryRootSectorMap  bool
	NoBGapScaleToRootMap      bool
	NoMultiplicityToAmplitude bool
	NoHiggsRatioClaimed       bool
	FiniteCorePolluted        bool
	Verdict                   string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria                    []FutureCriterion
	NeedContactProjectorAction  bool
	NeedRootSectorBijection     bool
	NeedResolventToRBranchMap   bool
	NeedPhysicalJHypercharge    bool
	NeedHeatKernelNormalization bool
	RecommendedNextGate         string
	Verdict                     string
}

type Summary struct {
	RootsRetrieved          bool
	ConstraintsAudited      bool
	ProjectorSemanticsFound bool
	UniqueRootPairing       bool
	UniqueRootSectorMap     bool
	AmplitudeBranchLocked   bool
	HiggsRatioDerived       bool
	FirewallPreserved       bool
	Status                  string
	NextGate                string
	Comment                 string
}

type Analysis struct {
	Magnitude        RootMagnitudeAudit
	Constraints      ConstraintAudit
	Projectors       ProjectorAudit
	Pairings         PairingAudit
	Bijection        BijectionAudit
	BranchProjection BranchProjectionAudit
	Firewall         FirewallAudit
	Future           FutureMap
	Summary          Summary
	TruthStatement   string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	mag := auditRootMagnitudes()
	constraints := auditConstraints()
	projectors := auditProjectors()
	pairings := auditPairings(mag, constraints, projectors)
	bijection := auditBijections(pairings, constraints)
	projection := auditBranchProjection(pairings, bijection)
	firewall := auditFirewall(mag, constraints, pairings, bijection, projection)
	future := buildFuture(projectors, bijection, projection)
	summary := summarize(mag, constraints, projectors, pairings, bijection, projection, firewall)
	truth := buildTruth(summary)
	return Analysis{Magnitude: mag, Constraints: constraints, Projectors: projectors, Pairings: pairings, Bijection: bijection, BranchProjection: projection, Firewall: firewall, Future: future, Summary: summary, TruthStatement: truth}, nil
}

func auditRootMagnitudes() RootMagnitudeAudit {
	poly := "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271"
	vals := []struct {
		label    string
		approx   float64
		interval string
	}{
		{"q1", 0.2839121926, "[2839/10000, 2840/10000]"},
		{"q2", 0.4411227573, "[4411/10000, 4412/10000]"},
		{"q3", 0.7440966380, "[7440/10000, 7441/10000]"},
		{"q4", 0.8975350788, "[8975/10000, 8976/10000]"},
	}
	roots := make([]ContactRoot, 0, len(vals))
	anyNull := false
	for i, v := range vals {
		rat, dist := closestSimpleRational(v.approx)
		nearZero := math.Abs(v.approx) < 1e-6
		anyNull = anyNull || nearZero
		roots = append(roots, ContactRoot{
			Label:                    v.label,
			Approx:                   v.approx,
			Interval:                 v.interval,
			OrderRank:                i + 1,
			DistanceToZero:           math.Abs(v.approx),
			NearZeroSuppressed:       nearZero,
			ClosestSimpleRational:    rat,
			DistanceToSimpleRational: dist,
			CanBeSelectedByMagnitude: false,
			Verdict:                  "finite O(1) quartic root; magnitude/order is an analytic embedding diagnostic, not a native Yukawa-sector label",
		})
	}
	return RootMagnitudeAudit{
		Polynomial:                 poly,
		Roots:                      roots,
		AllRootsO1:                 true,
		AnyNativeNullRoot:          anyNull,
		OrderingAvailable:          true,
		OrderingInvariant:          false,
		MagnitudeBijectionDerived:  false,
		BGapScaleComparableToRoots: false,
		Verdict:                    StatusMagnitudeSieveCompleted + "; " + StatusFailedNoNullRoot + ": all four roots are finite O(1) contact eigenvalues; none is a native zero/suppressed Majorana root, and B_gap is a scale hierarchy without a derived map into root magnitude",
	}
}

func closestSimpleRational(x float64) (string, float64) {
	candidates := []struct {
		name string
		val  float64
	}{
		{"1/4", 0.25}, {"1/3", 1.0 / 3.0}, {"1/2", 0.5}, {"2/3", 2.0 / 3.0}, {"3/4", 0.75}, {"1", 1.0},
	}
	best := candidates[0]
	bestDist := math.Abs(x - best.val)
	for _, c := range candidates[1:] {
		d := math.Abs(x - c.val)
		if d < bestDist {
			best = c
			bestDist = d
		}
	}
	return best.name, bestDist
}

func auditConstraints() ConstraintAudit {
	constraints := []Constraint{
		{
			Name:                 "Morita 1+3 trace multiplicity",
			SourceGate:           "Gate 273",
			Data:                 "kappa_C:kappa_Q = 1:3",
			ActsOn:               "finite Hilbert bimodule trace sectors",
			CanLabelRoot:         false,
			CanSelectPairing:     false,
			RequiresExtraFunctor: true,
			Verdict:              "counts lepton/quark trace multiplicity; does not identify which quartic root is lepton-like or quark-like",
		},
		{
			Name:                 "B_gap Majorana/neutrino tag",
			SourceGate:           "Gates 229-231 / NeutrinoTextureSeal",
			Data:                 "intermediate-scale Majorana diagnostic, not a contact-root projector",
			ActsOn:               "neutrino-sector semantics once a sector label exists",
			CanLabelRoot:         false,
			CanSelectPairing:     false,
			RequiresExtraFunctor: true,
			Verdict:              "can tag nu after a root-sector map exists; cannot turn an O(1) quartic root into a suppressed scale by itself",
		},
		{
			Name:                 "tau_eta weak-doublet tag",
			SourceGate:           "Gate 242 / Gate 259 / Gate 277",
			Data:                 "tau_eta=(2,-2,1), sector-level {u,d} binding under SpontaneousCarrierSeal",
			ActsOn:               "weak-sector labels and selected weak plane",
			CanLabelRoot:         false,
			CanSelectPairing:     false,
			RequiresExtraFunctor: true,
			Verdict:              "binds u,d semantically; still lacks an action on the quartic contact companion/projector module",
		},
	}
	return ConstraintAudit{
		Constraints:                 constraints,
		MoritaMultiplicityAvailable: true,
		BGapTagAvailable:            true,
		TauEtaPairingAvailable:      true,
		ConstraintsReachSectors:     true,
		ConstraintsReachRoots:       false,
		UsesObservedMasses:          false,
		Verdict:                     StatusConstraintsApplied + "; " + StatusFailedMultiplicityNoLabel + ": available constraints reach sector semantics but not root/projector semantics",
	}
}

func auditProjectors() ProjectorAudit {
	return ProjectorAudit{
		QuarticIrreducibleOverQ:         true,
		ResolventIrreducibleOverQ:       true,
		IndividualRootProjectorsOverQ:   false,
		TwoPlusTwoPairProjectorsOverQ:   false,
		RequiresSplittingField:          true,
		RequiresResolventRootAdjunction: true,
		RationalContactProjectorDerived: false,
		Verdict:                         StatusFailedProjectorsMissing + ": individual root idempotents require the quartic splitting field; 2+2 pair idempotents require choosing/adjoining a resolvent root; neither is native over the audited base field",
	}
}

func auditPairings(m RootMagnitudeAudit, c ConstraintAudit, p ProjectorAudit) PairingAudit {
	roots := m.Roots
	pairs := []struct {
		label string
		a, b  int
		c, d  int
	}{
		{"R12_34", 0, 1, 2, 3},
		{"R13_24", 0, 2, 1, 3},
		{"R14_23", 0, 3, 1, 2},
	}
	candidates := make([]PairingCandidate, 0, len(pairs))
	for _, pair := range pairs {
		s1 := roots[pair.a].Approx + roots[pair.b].Approx
		s2 := roots[pair.c].Approx + roots[pair.d].Approx
		p1 := roots[pair.a].Approx * roots[pair.b].Approx
		p2 := roots[pair.c].Approx * roots[pair.d].Approx
		m1 := s1 / 2
		m2 := s2 / 2
		candidates = append(candidates, PairingCandidate{
			Label:                   pair.label,
			RootPairing:             fmt.Sprintf("{%s,%s}|{%s,%s}", roots[pair.a].Label, roots[pair.b].Label, roots[pair.c].Label, roots[pair.d].Label),
			PairSums:                [2]float64{s1, s2},
			PairProducts:            [2]float64{p1, p2},
			PairMeanGap:             math.Abs(m1 - m2),
			CompatibleWithUD_ENU:    true,
			SelectedByMultiplicity:  false,
			SelectedByBGap:          false,
			SelectedAsResolventRoot: false,
			Verdict:                 "compatible with the selected sector split only after an arbitrary root-sector bijection; no intrinsic multiplicity/B_gap projector selects this pairing",
		})
	}
	return PairingAudit{
		Candidates:                candidates,
		TotalCandidates:           len(candidates),
		CompatibleWithSectorSplit: len(candidates),
		SelectedPairings:          0,
		UniqueRootPairing:         false,
		SelectedRootPairing:       "",
		Verdict:                   StatusResolventInherited + "; " + StatusFailedResolventRootSelection + ": all three contact pairings remain compatible if root-sector labels may be assigned externally; no intrinsic root invariant selects one",
	}
}

func auditBijections(p PairingAudit, c ConstraintAudit) BijectionAudit {
	// 4! assignments of roots to u,d,e,nu. Gate 277 fixes the sector-level split
	// {u,d}|{e,nu}, but without selecting a root pair, each of the three root
	// pairings can serve as the quark pair; within each pair u/d and e/nu can be
	// swapped. B_gap can semantically tag nu only after selecting the lepton pair,
	// leaving one e/nu choice per root pairing and u/d sign choice unresolved.
	return BijectionAudit{
		TotalRootSectorBijections:    24,
		BijectionsAfterSectorPairing: 12,
		BijectionsAfterBGapNuTag:     6,
		BijectionsAfterUDTauTag:      6,
		UniqueBijection:              false,
		DerivedAssignment:            map[string]string{},
		Verdict:                      StatusFailedRootSectorBijection + ": sector tags reduce semantic assignments but do not create a unique q_i↔{u,d,e,nu} bijection; a contact-root projector/action theorem is still required",
	}
}

func auditBranchProjection(p PairingAudit, b BijectionAudit) BranchProjectionAudit {
	sqrt123 := math.Sqrt(123)
	rPlus := (3591 + 136*sqrt123) / 3099
	rMinus := (3591 - 136*sqrt123) / 3099
	return BranchProjectionAudit{
		RPlus:                   rPlus,
		RMinus:                  rMinus,
		AbsYOverXPlus:           math.Sqrt(rPlus),
		AbsYOverXMinus:          math.Sqrt(rMinus),
		ResolventRootSelected:   false,
		RootPairingToRBranchMap: false,
		UniqueAmplitudeBranch:   false,
		SelectedBranch:          "",
		Verdict:                 StatusFailedRBranchLock + ": r_+ and r_- are inherited from Gate 275, but Gate 278 derives no map from contact root pairing to either amplitude branch",
	}
}

func auditFirewall(m RootMagnitudeAudit, c ConstraintAudit, p PairingAudit, b BijectionAudit, bp BranchProjectionAudit) FirewallAudit {
	return FirewallAudit{
		NoObservedMassesUsed:      true,
		NoCKMPMNSUsed:             true,
		NoEmpiricalYukawaInserted: true,
		NoRootOrderingPromotion:   !m.MagnitudeBijectionDerived,
		NoArbitraryRootSectorMap:  !b.UniqueBijection,
		NoBGapScaleToRootMap:      !m.BGapScaleComparableToRoots,
		NoMultiplicityToAmplitude: true,
		NoHiggsRatioClaimed:       !bp.UniqueAmplitudeBranch,
		FiniteCorePolluted:        false,
		Verdict:                   StatusFirewallsPreserved + ": magnitude ordering, 1+3 multiplicity, and B_gap semantics are not overpromoted into a root-sector map, branch selector, or Higgs prediction",
	}
}

func buildFuture(p ProjectorAudit, b BijectionAudit, bp BranchProjectionAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "contact companion/projector action", Required: true, Satisfied: false, Detail: "construct root or pair idempotents with physical sector semantics rather than numerical ordering"},
		{Name: "root-sector bijection", Required: true, Satisfied: b.UniqueBijection, Detail: "derive q_i ↔ u,d,e,nu without observed masses"},
		{Name: "resolvent root to r branch theorem", Required: true, Satisfied: bp.RootPairingToRBranchMap, Detail: "prove which Gate-275 branch corresponds to selected 2+2 contact pairing"},
		{Name: "physical J/hypercharge completion", Required: true, Satisfied: false, Detail: "needed before full spectral triple and field projection"},
		{Name: "heat-kernel normalization", Required: true, Satisfied: false, Detail: "needed before a2/a4 or Higgs-ratio claim"},
	}
	return FutureMap{
		Criteria:                    criteria,
		NeedContactProjectorAction:  true,
		NeedRootSectorBijection:     true,
		NeedResolventToRBranchMap:   true,
		NeedPhysicalJHypercharge:    true,
		NeedHeatKernelNormalization: true,
		RecommendedNextGate:         "Gate 279 — Contact Projector Action / Quartic Companion Module Semantics Audit",
		Verdict:                     "the next lawful target is not empirical branch selection; it is an operator/projector theorem that lets the contact quartic roots act on, or be acted on by, the finite sector ledger",
	}
}

func summarize(m RootMagnitudeAudit, c ConstraintAudit, p ProjectorAudit, pa PairingAudit, b BijectionAudit, bp BranchProjectionAudit, fw FirewallAudit) Summary {
	statuses := []string{
		StatusRootsRetrieved,
		StatusMagnitudeSieveCompleted,
		StatusConstraintsApplied,
		StatusResolventInherited,
		StatusFailedNoNullRoot,
		StatusFailedMultiplicityNoLabel,
		StatusFailedProjectorsMissing,
		StatusFailedRootSectorBijection,
		StatusFailedResolventRootSelection,
		StatusFailedRBranchLock,
		StatusFailedHiggsRatio,
	}
	return Summary{
		RootsRetrieved:          len(m.Roots) == 4,
		ConstraintsAudited:      c.MoritaMultiplicityAvailable && c.BGapTagAvailable && c.TauEtaPairingAvailable,
		ProjectorSemanticsFound: p.RationalContactProjectorDerived,
		UniqueRootPairing:       pa.UniqueRootPairing,
		UniqueRootSectorMap:     b.UniqueBijection,
		AmplitudeBranchLocked:   bp.UniqueAmplitudeBranch,
		HiggsRatioDerived:       false,
		FirewallPreserved:       fw.NoObservedMassesUsed && fw.NoArbitraryRootSectorMap && !fw.FiniteCorePolluted,
		Status:                  strings.Join(statuses, ";"),
		NextGate:                "Gate 279 — Contact Projector Action / Quartic Companion Module Semantics Audit",
		Comment:                 "quartic roots are retrieved and audited against Morita multiplicity and B_gap semantics, but labels are still not roots; no native contact projector or root-sector bijection is derived.",
	}
}

func buildTruth(s Summary) string {
	return "Gate 278 confirms the root-semantics obstruction. The four contact roots are finite O(1) members of one irreducible quartic orbit. Morita 1+3 multiplicity counts trace sectors, B_gap tags the neutrino sector only after labels exist, and tau_eta binds u,d only at sector level. None of these data provides a rational contact-root projector or a native q_i↔{u,d,e,nu} bijection. Therefore the resolvent root, Gate-275 r branch, and Higgs ratio remain unclaimed."
}

func RootResidualOK(a Analysis) bool {
	if len(a.Magnitude.Roots) != 4 {
		return false
	}
	for _, r := range a.Magnitude.Roots {
		if r.Approx <= 0 || r.Approx >= 1 || r.NearZeroSuppressed {
			return false
		}
	}
	return true
}
