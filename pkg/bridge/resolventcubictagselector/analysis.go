// Package resolventcubictagselector implements Gate 277:
// Resolvent Cubic Selector / B-Gap and Tau-Eta Symmetry Breaking Audit.
//
// Gate 276 left a two-branch scalar-Morita amplitude ambiguity. Earlier scalar
// gates left a three-branch quartic-contact resolvent ambiguity: a root of the
// resolvent cubic selects one of the three unordered 2+2 partitions of the four
// quartic contact roots. Gate 277 audits whether native topological tags
// (tau_eta and B_gap/Majorana scale diagnostics) break the sector-level S4
// degeneracy enough to select the physical {u,d}|{e,nu} pairing and whether that
// selection can be pushed all the way down to a contact resolvent root and a
// Gate-275 r_+/r_- branch.
//
// The result is deliberately firewalled. The tags do uniquely select the
// Standard-Model sector pairing among the three abstract 2+2 partitions of
// {u,d,e,nu}. However, the project still lacks a native bijection between the
// four quartic contact roots and the four Yukawa sectors, and it lacks a native
// map from the selected resolvent root to the two scalar-Morita branches. Thus
// the sector pairing is conditionally selected, while the contact root and Higgs
// amplitude branch remain blocked.
package resolventcubictagselector

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE277-RESOLVENT-CUBIC-SELECTOR-BGAP-TAUETA-SYMMETRY-BREAKING-AUDIT"

	StatusResolventRetrieved     = "CONDITIONAL_SUPPORT_RESOLVENT_CUBIC_RETRIEVED"
	StatusTopologicalTagsApplied = "CONDITIONAL_SUPPORT_TAU_ETA_AND_B_GAP_TAGS_APPLIED"
	StatusSectorPairingSelected  = "CONDITIONAL_SUPPORT_SECTOR_LEVEL_UD_ENU_PAIRING_SELECTED"
	StatusGaloisSieveCompleted   = "CONDITIONAL_SUPPORT_GALOIS_ORBIT_SIEVE_COMPLETED"
	StatusGate275BranchInherited = "CONDITIONAL_SUPPORT_GATE275_AMPLITUDE_BRANCHES_INHERITED"
	StatusFirewallPreserved      = "CONDITIONAL_SUPPORT_RESOLVENT_SELECTOR_FIREWALLS_PRESERVED"

	StatusFailedRootSectorMap   = "FAILED_ROUTE_QUARTIC_ROOT_TO_YUKAWA_SECTOR_BIJECTION_MISSING"
	StatusFailedContactRoot     = "FAILED_ROUTE_CONTACT_RESOLVENT_ROOT_NOT_SELECTED"
	StatusFailedRBranchMap      = "FAILED_ROUTE_RESOLVENT_TO_RPLUS_RMINUS_BRANCH_MAP_MISSING"
	StatusFailedAmplitudeBranch = "FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED"
	StatusFailedHiggsRatio      = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

type QuarticRoot struct {
	Label        string
	Interval     string
	Approx       float64
	Polynomial   string
	Sector       string
	SectorMapped bool
	Verdict      string
}

type ResolventBranch struct {
	Label                 string
	PairingOnRoots        string
	PairingOnSectors      string
	SelectedByTags        bool
	EliminatedByTauEta    bool
	EliminatedByBGap      bool
	RequiresRootSectorMap bool
	ContactRootSelected   bool
	Verdict               string
}

type ResolventAudit struct {
	QuarticPolynomial               string
	ResolventIntegerPolynomial      string
	ResolventMonicCoefficients      []string
	QuarticRoots                    []QuarticRoot
	Branches                        []ResolventBranch
	EncodesTwoPlusTwo               bool
	IrreducibleOverQ                bool
	CanonicalRootPreviouslySelected bool
	Retrieved                       bool
	Verdict                         string
}

type TopologicalTag struct {
	Name                            string
	SourceGate                      string
	ActsOn                          string
	NativeData                      string
	SelectionRule                   string
	IsSealed                        bool
	DerivedAsOperatorOnQuarticRoots bool
	Verdict                         string
}

type TagAudit struct {
	Tags                  []TopologicalTag
	TauEtaBindsUD         bool
	BGapTagsNeutrino      bool
	TagsReachSectorLabels bool
	TagsReachQuarticRoots bool
	UsesObservedMasses    bool
	Verdict               string
}

type PairingSieve struct {
	CandidatePairings       []ResolventBranch
	TotalCandidates         int
	SurvivingSectorPairings int
	SelectedSectorPairing   string
	EliminatedCandidates    int
	UniqueSectorPairing     bool
	UniqueContactRoot       bool
	Verdict                 string
}

type Gate275Branch struct {
	Name              string
	ExactR            string
	R                 float64
	AbsYOverX         float64
	ContactPairingMap string
	Selected          bool
	Verdict           string
}

type BranchProjectionAudit struct {
	BranchesInherited         []Gate275Branch
	SelectedSectorPairing     string
	ResolventRootSelected     bool
	ResolventRootToRBranchMap bool
	UniqueRBranchSelected     bool
	SelectedRBranch           string
	Verdict                   string
}

type FirewallAudit struct {
	NoObservedMassesUsed         bool
	NoCKMPMNSUsed                bool
	NoEmpiricalYukawaInserted    bool
	NoArbitraryRootSectorMap     bool
	SectorPairingNotOverpromoted bool
	ContactRootNotOverpromoted   bool
	RBranchNotOverpromoted       bool
	HiggsRatioNotClaimed         bool
	FiniteCorePolluted           bool
	Verdict                      string
}

type FutureCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type FutureMap struct {
	Criteria              []FutureCriterion
	NeedRootSectorMap     bool
	NeedContactProjectors bool
	NeedResolventToRMap   bool
	NeedBranchSelector    bool
	NeedHeatKernelMap     bool
	RecommendedNextGate   string
	Verdict               string
}

type Summary struct {
	ResolventRetrieved    bool
	TagsApplied           bool
	UniqueSectorPairing   bool
	UniqueContactRoot     bool
	UniqueAmplitudeBranch bool
	HiggsRatioDerived     bool
	FirewallPreserved     bool
	Status                string
	NextGate              string
	Comment               string
}

type Analysis struct {
	Resolvent        ResolventAudit
	Tags             TagAudit
	Sieve            PairingSieve
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
	defaultOnce.Do(func() {
		defaultA, defaultErr = Build()
	})
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	resolvent := retrieveResolvent()
	tags := auditTags()
	sieve := runPairingSieve(resolvent, tags)
	projection := auditGate275Projection(sieve)
	firewall := auditFirewall(sieve, projection)
	future := buildFuture(sieve, projection)
	summary := summarize(resolvent, tags, sieve, projection, firewall)
	truth := buildTruth(sieve, projection)
	return Analysis{Resolvent: resolvent, Tags: tags, Sieve: sieve, BranchProjection: projection, Firewall: firewall, Future: future, Summary: summary, TruthStatement: truth}, nil
}

func retrieveResolvent() ResolventAudit {
	poly := "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271"
	roots := []QuarticRoot{
		{Label: "q1", Interval: "[2839/10000, 2840/10000]", Approx: 0.2839121926, Polynomial: poly, Sector: "", SectorMapped: false, Verdict: "isolated quartic contact root; no native Yukawa-sector label assigned"},
		{Label: "q2", Interval: "[4411/10000, 4412/10000]", Approx: 0.4411227573, Polynomial: poly, Sector: "", SectorMapped: false, Verdict: "isolated quartic contact root; no native Yukawa-sector label assigned"},
		{Label: "q3", Interval: "[7440/10000, 7441/10000]", Approx: 0.7440966380, Polynomial: poly, Sector: "", SectorMapped: false, Verdict: "isolated quartic contact root; no native Yukawa-sector label assigned"},
		{Label: "q4", Interval: "[8975/10000, 8976/10000]", Approx: 0.8975350788, Polynomial: poly, Sector: "", SectorMapped: false, Verdict: "isolated quartic contact root; no native Yukawa-sector label assigned"},
	}
	branches := []ResolventBranch{
		{Label: "R12_34", PairingOnRoots: "{q1,q2}|{q3,q4}", PairingOnSectors: "{u,d}|{e,nu}", RequiresRootSectorMap: true, Verdict: "formal resolvent branch; sector pairing shown only as a semantic target, not as a root theorem"},
		{Label: "R13_24", PairingOnRoots: "{q1,q3}|{q2,q4}", PairingOnSectors: "{u,e}|{d,nu}", RequiresRootSectorMap: true, Verdict: "formal resolvent branch; sector pairing shown only as a semantic target, not as a root theorem"},
		{Label: "R14_23", PairingOnRoots: "{q1,q4}|{q2,q3}", PairingOnSectors: "{u,nu}|{d,e}", RequiresRootSectorMap: true, Verdict: "formal resolvent branch; sector pairing shown only as a semantic target, not as a root theorem"},
	}
	return ResolventAudit{
		QuarticPolynomial:               poly,
		ResolventIntegerPolynomial:      "5832000z^3 - 11566800z^2 + 7569900z - 1637467",
		ResolventMonicCoefficients:      []string{"1", "-119/60", "8411/6480", "-1637467/5832000"},
		QuarticRoots:                    roots,
		Branches:                        branches,
		EncodesTwoPlusTwo:               true,
		IrreducibleOverQ:                true,
		CanonicalRootPreviouslySelected: false,
		Retrieved:                       true,
		Verdict:                         StatusResolventRetrieved + ": exact quartic roots and three 2+2 resolvent branches are retrieved as audited finite data; no root-sector assignment is retrieved",
	}
}

func auditTags() TagAudit {
	tags := []TopologicalTag{
		{
			Name:                            "tau_eta weak doublet tag",
			SourceGate:                      "Gate 242 / Gate 259",
			ActsOn:                          "weak-plane / generation-breaking selector",
			NativeData:                      "tau_eta=(2,-2,1), |tau_eta|=(2,2,1)",
			SelectionRule:                   "bind u and d as the weak-isospin pair when sector labels are already available",
			IsSealed:                        true,
			DerivedAsOperatorOnQuarticRoots: false,
			Verdict:                         "conditioned by SpontaneousCarrierSeal; reaches weak sector labels but not quartic contact roots",
		},
		{
			Name:                            "B_gap Majorana/neutrino tag",
			SourceGate:                      "Gates 229-231 / NeutrinoTextureSeal",
			ActsOn:                          "intermediate-scale neutrino/Majorana diagnostic",
			NativeData:                      "B_gap hierarchy prepares a sealed intermediate/Majorana scale diagnostic, not a finite mass matrix",
			SelectionRule:                   "isolate nu from the quartic Yukawa-sector orbit when sector labels are already available",
			IsSealed:                        true,
			DerivedAsOperatorOnQuarticRoots: false,
			Verdict:                         "reaches the neutrino seal semantically, but no operator maps B_gap to a quartic contact projector",
		},
	}
	return TagAudit{
		Tags:                  tags,
		TauEtaBindsUD:         true,
		BGapTagsNeutrino:      true,
		TagsReachSectorLabels: true,
		TagsReachQuarticRoots: false,
		UsesObservedMasses:    false,
		Verdict:               StatusTopologicalTagsApplied + ": tags select sector-level semantics under existing seals, but do not act as contact-root projectors",
	}
}

func runPairingSieve(res ResolventAudit, tags TagAudit) PairingSieve {
	candidates := make([]ResolventBranch, 0, len(res.Branches))
	survivors := 0
	selected := ""
	for _, b := range res.Branches {
		nb := b
		switch b.PairingOnSectors {
		case "{u,d}|{e,nu}":
			nb.SelectedByTags = tags.TauEtaBindsUD && tags.BGapTagsNeutrino
			nb.EliminatedByTauEta = false
			nb.EliminatedByBGap = false
			nb.ContactRootSelected = false
			nb.Verdict = "sector-level pass: tau_eta binds {u,d} and B_gap places nu in the complementary lepton pair; still not a contact-root selection because root-sector bijection is missing"
			if nb.SelectedByTags {
				survivors++
				selected = b.PairingOnSectors
			}
		case "{u,e}|{d,nu}":
			nb.SelectedByTags = false
			nb.EliminatedByTauEta = true
			nb.EliminatedByBGap = true
			nb.Verdict = "sector-level reject: u is separated from d and nu is paired with d rather than with the lepton partner"
		case "{u,nu}|{d,e}":
			nb.SelectedByTags = false
			nb.EliminatedByTauEta = true
			nb.EliminatedByBGap = true
			nb.Verdict = "sector-level reject: u is separated from d and nu is paired with u rather than with the lepton partner"
		default:
			nb.SelectedByTags = false
			nb.Verdict = "unknown sector pairing; rejected by selector audit"
		}
		candidates = append(candidates, nb)
	}
	uniqueSector := survivors == 1
	return PairingSieve{
		CandidatePairings:       candidates,
		TotalCandidates:         len(candidates),
		SurvivingSectorPairings: survivors,
		SelectedSectorPairing:   selected,
		EliminatedCandidates:    len(candidates) - survivors,
		UniqueSectorPairing:     uniqueSector,
		UniqueContactRoot:       false,
		Verdict:                 StatusGaloisSieveCompleted + ": sector-pairing degeneracy is reduced from 3 to 1, but the contact resolvent root remains unselected without a quartic-root/Yukawa-sector bijection",
	}
}

func auditGate275Projection(s PairingSieve) BranchProjectionAudit {
	branches := gate275Branches()
	for i := range branches {
		branches[i].ContactPairingMap = "missing: no derived map from selected sector pairing to r_+/r_-"
		branches[i].Selected = false
		branches[i].Verdict = "carried as Gate-275 scalar-Morita branch; not selected by Gate-277 tags"
	}
	return BranchProjectionAudit{
		BranchesInherited:         branches,
		SelectedSectorPairing:     s.SelectedSectorPairing,
		ResolventRootSelected:     false,
		ResolventRootToRBranchMap: false,
		UniqueRBranchSelected:     false,
		SelectedRBranch:           "",
		Verdict:                   StatusGate275BranchInherited + "; " + StatusFailedRBranchMap + ": the selected sector pairing has no native projection onto r_+ or r_-",
	}
}

func gate275Branches() []Gate275Branch {
	sqrt123 := math.Sqrt(123)
	rPlus := (3591 + 136*sqrt123) / 3099
	rMinus := (3591 - 136*sqrt123) / 3099
	return []Gate275Branch{
		{Name: "r_plus", ExactR: "(3591 + 136√123)/3099", R: rPlus, AbsYOverX: math.Sqrt(rPlus)},
		{Name: "r_minus", ExactR: "(3591 - 136√123)/3099", R: rMinus, AbsYOverX: math.Sqrt(rMinus)},
	}
}

func auditFirewall(s PairingSieve, p BranchProjectionAudit) FirewallAudit {
	return FirewallAudit{
		NoObservedMassesUsed:         true,
		NoCKMPMNSUsed:                true,
		NoEmpiricalYukawaInserted:    true,
		NoArbitraryRootSectorMap:     true,
		SectorPairingNotOverpromoted: s.UniqueSectorPairing && !s.UniqueContactRoot,
		ContactRootNotOverpromoted:   !s.UniqueContactRoot && !p.ResolventRootSelected,
		RBranchNotOverpromoted:       !p.UniqueRBranchSelected,
		HiggsRatioNotClaimed:         true,
		FiniteCorePolluted:           false,
		Verdict:                      StatusFirewallPreserved + ": semantic sector selection is not promoted to quartic root selection, amplitude branch selection, or Higgs prediction",
	}
}

func buildFuture(s PairingSieve, p BranchProjectionAudit) FutureMap {
	criteria := []FutureCriterion{
		{Name: "quartic root to Yukawa sector bijection", Required: true, Satisfied: false, Detail: "must assign q1..q4 to u,d,e,nu without observed masses or arbitrary ordering"},
		{Name: "branchwise contact projectors", Required: true, Satisfied: false, Detail: "must construct projectors/idempotents on the quartic companion module for the selected sector pairing"},
		{Name: "resolvent root to scalar-Morita r branch map", Required: true, Satisfied: false, Detail: "must prove whether the selected 2+2 pairing corresponds to r_plus or r_minus"},
		{Name: "physical J/hypercharge completion", Required: true, Satisfied: false, Detail: "needed before physical spectral triple and heat-kernel projection"},
		{Name: "heat-kernel and field-normalization projection", Required: true, Satisfied: false, Detail: "needed before claiming a2/a4 or Higgs ratio"},
		{Name: "sector-level tau_eta/B_gap pairing selector", Required: true, Satisfied: s.UniqueSectorPairing && p.SelectedSectorPairing == "{u,d}|{e,nu}", Detail: "Gate 277 selects the sector pairing only"},
	}
	return FutureMap{
		Criteria:              criteria,
		NeedRootSectorMap:     true,
		NeedContactProjectors: true,
		NeedResolventToRMap:   true,
		NeedBranchSelector:    true,
		NeedHeatKernelMap:     true,
		RecommendedNextGate:   "Gate 278 — Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit",
		Verdict:               "the next theorem must turn sector labels into contact-root/projector semantics before the r branch can be selected",
	}
}

func summarize(res ResolventAudit, tags TagAudit, s PairingSieve, p BranchProjectionAudit, fw FirewallAudit) Summary {
	statuses := []string{StatusResolventRetrieved, StatusTopologicalTagsApplied, StatusSectorPairingSelected, StatusGaloisSieveCompleted, StatusGate275BranchInherited, StatusFailedRootSectorMap, StatusFailedContactRoot, StatusFailedRBranchMap, StatusFailedAmplitudeBranch, StatusFailedHiggsRatio}
	return Summary{
		ResolventRetrieved:    res.Retrieved && res.EncodesTwoPlusTwo,
		TagsApplied:           tags.TauEtaBindsUD && tags.BGapTagsNeutrino,
		UniqueSectorPairing:   s.UniqueSectorPairing,
		UniqueContactRoot:     s.UniqueContactRoot,
		UniqueAmplitudeBranch: p.UniqueRBranchSelected,
		HiggsRatioDerived:     false,
		FirewallPreserved:     fw.NoObservedMassesUsed && fw.NoArbitraryRootSectorMap && !fw.FiniteCorePolluted,
		Status:                strings.Join(statuses, ";"),
		NextGate:              "Gate 278 — Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit",
		Comment:               "tau_eta+B_gap break the sector-level S4 ambiguity to {u,d}|{e,nu}, but no native map sends that semantic pairing to a contact resolvent root or Gate-275 r branch.",
	}
}

func buildTruth(s PairingSieve, p BranchProjectionAudit) string {
	return fmt.Sprintf("Gate 277 shows that tau_eta and B_gap can uniquely select the sector-level 2+2 pairing %s among the three Yukawa-sector pairings. This is a genuine topological-semantic sieve. It does not yet select a quartic contact resolvent root or the Gate-275 r branch, because the project has no native bijection q_i↔{u,d,e,nu} and no theorem mapping the selected resolvent pairing to r_+ or r_-. The Higgs ratio remains unclaimed.", s.SelectedSectorPairing)
}

func BranchResidualOK(branches []Gate275Branch) bool {
	if len(branches) != 2 {
		return false
	}
	for _, b := range branches {
		if b.R <= 0 || b.AbsYOverX <= 0 {
			return false
		}
	}
	return true
}
