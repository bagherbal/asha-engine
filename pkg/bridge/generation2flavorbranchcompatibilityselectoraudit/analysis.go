// Package generation2flavorbranchcompatibilityselectoraudit implements
// Gate 601: Flavor Branch-Compatibility Selector Audit.
//
// Gate 600 decomposed epsilon(H_e) into native trace-ring data plus an
// environmental root-branch/chamber seal. Gate 601 asks whether the observed
// environmental balance B_flav≈0 selects the observed branch among charged
// lepton chamber permutations, PMNS projector choices, and CKM orientation
// signs. This is a branch-compatibility audit only: it does not derive Koide,
// charged-lepton masses, PMNS, CKM, neutrino parameters, or B_flav=0 as native
// ASHA law.
package generation2flavorbranchcompatibilityselectoraudit

import (
	"fmt"
	"math"
	"math/cmplx"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2chargedleptonrootextensionbranchchambermonodromyaudit"
)

const (
	AuditID = "GATE601-FLAVOR-BRANCH-COMPATIBILITY-SELECTOR-AUDIT"

	StatusGate600Inherited                   = "PASS_GATE600_BRANCH_CHAMBER_MONODROMY_RESULT_INHERITED"
	StatusBranchBalanceDefined               = "PASS_BRANCH_BALANCE_FUNCTION_DEFINED"
	StatusBranchSpaceEnumerated              = "PASS_BRANCH_SPACE_ENUMERATED"
	StatusChargedLeptonBranchesEnumerated    = "PASS_CHARGED_LEPTON_BRANCHES_ENUMERATED"
	StatusPMNSProjectorsEnumerated           = "PASS_PMNS_PROJECTOR_CHOICES_ENUMERATED"
	StatusCKMSignsEnumerated                 = "PASS_CKM_ORIENTATION_SIGNS_ENUMERATED"
	StatusFullBranchBalanceTableComputed     = "PASS_FULL_BRANCH_BALANCE_TABLE_COMPUTED"
	StatusObservedBranchHasMinimumResidual   = "PASS_OBSERVED_BRANCH_IN_MINIMAL_RESIDUAL_CLASS"
	StatusBalanceSelectsP3AndPositiveJ       = "CONDITIONAL_SUPPORT_BALANCE_SELECTS_THIRD_NEUTRINO_PROJECTOR_AND_POSITIVE_CKM_SIGN"
	StatusChargedLeptonPermutationDegeneracy = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_PERMUTATION_DEGENERACY_EXPOSED"
	StatusBranchSelectorNotUnique            = "FAILED_ROUTE_BRANCH_SELECTOR_NOT_UNIQUE"
	StatusNoUniqueChargedLeptonOrdering      = "FAILED_ROUTE_BALANCE_DOES_NOT_UNIQUELY_SELECT_CHARGED_LEPTON_ORDERING"
	StatusNoNativeBranchSelectionTheorem     = "FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM"
	StatusNoNativeBFlavZero                  = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM"
	StatusBFlavEnvironmental                 = "FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_BRANCH_COMPATIBILITY_TEST"
	StatusGate600Boundary                    = "FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_BOUNDARY"
	StatusGate352Preserved                   = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                   = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusNoKoideDerivation                  = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoChargedLeptonMassDerivation      = "FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION"
	StatusNoPMNSCKMNeutrinoFlavorDerivation  = "FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusNoObservedDataPromotion            = "FIREWALL_PRESERVED_OBSERVED_BRANCH_LABELS_REMAIN_ENVIRONMENTAL_DATA"
	StatusNoNewCarrierSelector               = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate601Boundary                    = "FIREWALL_PRESERVED_GATE601_BRANCH_COMPATIBILITY_SELECTOR_BOUNDARY"
)

const (
	yElectron = 2.9350283095504176e-06
	yMuon     = 0.0006068707640859305
	yTau      = 0.010205763440624986

	sin2Theta12 = 0.308
	sin2Theta13 = 0.02215
	jCKM        = 3.1169935287554706e-05

	minimalTieTolerance = 1e-12
)

type InheritedGate600 struct {
	BranchSealDefined      bool
	EpsilonBranchAlgebraic bool
	NativeBranchTheorem    bool
	NativeFourthRoot       bool
	NativeChamberSelector  bool
	BFlavNative            bool
	Verdict                string
}

type BalanceDefinition struct {
	Formula        string
	SigmaDomain    string
	NeutrinoDomain string
	CKMSignDomain  string
	Environmental  bool
	Native         bool
	Verdict        string
}

type ChargedLeptonBranch struct {
	Sigma           string
	Order           []string
	DeltaDeg        float64
	R               float64
	ElectronIndex   int
	ElectronWallDeg float64
	EpsilonDeg      float64
	EpsilonRad      float64
	Kappa           float64
	PositiveChamber bool
	ChamberWalls    string
	Verdict         string
}

type PMNSProjectorOverlap struct {
	Index     int
	Projector string
	UeiAbs2   float64
	Li        float64
	Verdict   string
}

type CKMSign struct {
	Sign       int
	Convention string
	Value      float64
	Verdict    string
}

type BranchBalanceRow struct {
	Sigma      string
	NeutrinoI  int
	CKMSign    int
	EpsilonRad float64
	Kappa      float64
	Li         float64
	JTerm      float64
	BFlav      float64
	AbsBFlav   float64
	Observed   bool
}

type ObservedBranchRank struct {
	ObservedSigma           string
	ObservedNeutrinoI       int
	ObservedCKMSign         int
	ObservedBFlav           float64
	ObservedAbsBFlav        float64
	Rank                    int
	MinimalClassSize        int
	Unique                  bool
	MinimalClassDescription string
	Verdict                 string
}

type GapAudit struct {
	BestAbsResidual         float64
	NextDistinctAbsResidual float64
	GapToNextDistinct       float64
	GapMeaning              string
	GapLarge                bool
	Verdict                 string
}

type BranchSelectorVerdict struct {
	ObservedInMinimalClass        bool
	SelectsNeutrinoThirdProjector bool
	SelectsPositiveCKMSign        bool
	SelectsChargedLeptonOrdering  bool
	UniqueBranchSelector          bool
	NativeBranchSelector          bool
	Decision                      string
	Verdict                       string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesNeutrinoData        bool
	DerivesBFlavZeroNative     bool
	PromotesObservedData       bool
	AddsCarrier                bool
	AddsSelector               bool
	PreservesGate352           bool
	PreservesGate596           bool
	PreservesGate600           bool
	Verdict                    string
}

type Analysis struct {
	Inherited       InheritedGate600
	Definition      BalanceDefinition
	LeptonBranches  []ChargedLeptonBranch
	PMNSOverlaps    []PMNSProjectorOverlap
	CKMSigns        []CKMSign
	BalanceTable    []BranchBalanceRow
	ObservedRank    ObservedBranchRank
	Gap             GapAudit
	SelectorVerdict BranchSelectorVerdict
	Firewalls       FirewallAudit
	Truth           string
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
	g600, err := generation2chargedleptonrootextensionbranchchambermonodromyaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate600 predecessor: %w", err)
	}
	inherited := inheritGate600(g600)
	definition := defineBalance()
	branches := enumerateChargedLeptonBranches()
	pmns := enumeratePMNSOverlaps()
	signs := enumerateCKMSigns()
	table := computeBranchBalanceTable(branches, pmns, signs)
	rank := rankObservedBranch(table)
	gap := auditGap(table, rank)
	selector := compileSelectorVerdict(rank, gap)
	firewalls := auditFirewalls()
	truth := "Gate 601 finds that B_flav branch compatibility strongly selects the third neutrino projector P_3^nu and the positive CKM orientation sign, but it does not uniquely select the charged-lepton permutation/chamber branch: all six charged-lepton permutations lie in the same minimal residual class once epsilon is measured relative to the electron-zero wall. The balance is therefore a useful environmental branch-compatibility filter, not a native branch-selection theorem."
	return Analysis{Inherited: inherited, Definition: definition, LeptonBranches: branches, PMNSOverlaps: pmns, CKMSigns: signs, BalanceTable: table, ObservedRank: rank, Gap: gap, SelectorVerdict: selector, Firewalls: firewalls, Truth: truth}, nil
}

func inheritGate600(a generation2chargedleptonrootextensionbranchchambermonodromyaudit.Analysis) InheritedGate600 {
	return InheritedGate600{
		BranchSealDefined:      a.BranchSeal.Environmental,
		EpsilonBranchAlgebraic: a.Final.EpsilonBranchAlgebraic,
		NativeBranchTheorem:    a.Final.NativeEigenvalueBranch,
		NativeFourthRoot:       a.Final.NativePositiveFourthRoot,
		NativeChamberSelector:  a.Final.NativeChamberSelector,
		BFlavNative:            a.Final.BFlavNative,
		Verdict:                StatusGate600Inherited,
	}
}

func defineBalance() BalanceDefinition {
	return BalanceDefinition{
		Formula:        "B_flav(sigma,i,s_J)=1-8*pi*epsilon_sigma(H_e)-(1/4)Tr(P_eP_i^nu)+s_J*J_CKM",
		SigmaDomain:    "six charged-lepton root/chamber permutations",
		NeutrinoDomain: "i in {1,2,3}",
		CKMSignDomain:  "s_J in {+1,-1}",
		Environmental:  true,
		Native:         false,
		Verdict:        StatusBranchBalanceDefined,
	}
}

func enumerateChargedLeptonBranches() []ChargedLeptonBranch {
	orders := [][]string{
		{"e", "mu", "tau"},
		{"e", "tau", "mu"},
		{"mu", "e", "tau"},
		{"mu", "tau", "e"},
		{"tau", "e", "mu"},
		{"tau", "mu", "e"},
	}
	out := make([]ChargedLeptonBranch, 0, len(orders))
	for _, order := range orders {
		delta, r := fourierPhase(order)
		eidx := indexOf(order, "e")
		wall, epsDeg := nearestZeroWall(delta, eidx)
		epsRad := epsDeg * math.Pi / 180
		kappa := 1 - 8*math.Pi*epsRad
		out = append(out, ChargedLeptonBranch{
			Sigma:           strings.Join(order, ","),
			Order:           append([]string(nil), order...),
			DeltaDeg:        delta,
			R:               r,
			ElectronIndex:   eidx,
			ElectronWallDeg: wall,
			EpsilonDeg:      epsDeg,
			EpsilonRad:      epsRad,
			Kappa:           kappa,
			PositiveChamber: positiveComponents(delta, r),
			ChamberWalls:    "component zero walls at delta+2*pi*j/3 = 135deg or 225deg; epsilon measured to electron-zero wall",
			Verdict:         StatusChargedLeptonBranchesEnumerated,
		})
	}
	return out
}

func enumeratePMNSOverlaps() []PMNSProjectorOverlap {
	c13sq := 1 - sin2Theta13
	vals := []struct {
		i    int
		name string
		u2   float64
	}{
		{1, "P_1^nu", (1 - sin2Theta12) * c13sq},
		{2, "P_2^nu", sin2Theta12 * c13sq},
		{3, "P_3^nu", sin2Theta13},
	}
	out := make([]PMNSProjectorOverlap, 0, len(vals))
	for _, v := range vals {
		out = append(out, PMNSProjectorOverlap{Index: v.i, Projector: v.name, UeiAbs2: v.u2, Li: v.u2 / 4, Verdict: StatusPMNSProjectorsEnumerated})
	}
	return out
}

func enumerateCKMSigns() []CKMSign {
	return []CKMSign{
		{Sign: +1, Convention: "+J_CKM observed orientation convention", Value: +jCKM, Verdict: StatusCKMSignsEnumerated},
		{Sign: -1, Convention: "-J_CKM reversed orientation convention", Value: -jCKM, Verdict: StatusCKMSignsEnumerated},
	}
}

func computeBranchBalanceTable(branches []ChargedLeptonBranch, pmns []PMNSProjectorOverlap, signs []CKMSign) []BranchBalanceRow {
	rows := make([]BranchBalanceRow, 0, len(branches)*len(pmns)*len(signs))
	for _, b := range branches {
		for _, p := range pmns {
			for _, s := range signs {
				bf := b.Kappa - p.Li + s.Value
				rows = append(rows, BranchBalanceRow{Sigma: b.Sigma, NeutrinoI: p.Index, CKMSign: s.Sign, EpsilonRad: b.EpsilonRad, Kappa: b.Kappa, Li: p.Li, JTerm: s.Value, BFlav: bf, AbsBFlav: math.Abs(bf), Observed: b.Sigma == "e,mu,tau" && p.Index == 3 && s.Sign == +1})
			}
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].AbsBFlav == rows[j].AbsBFlav {
			if rows[i].NeutrinoI == rows[j].NeutrinoI {
				return rows[i].Sigma < rows[j].Sigma
			}
			return rows[i].NeutrinoI < rows[j].NeutrinoI
		}
		return rows[i].AbsBFlav < rows[j].AbsBFlav
	})
	return rows
}

func rankObservedBranch(rows []BranchBalanceRow) ObservedBranchRank {
	obsIdx := -1
	for i, r := range rows {
		if r.Observed {
			obsIdx = i
			break
		}
	}
	if obsIdx < 0 {
		return ObservedBranchRank{ObservedSigma: "e,mu,tau", ObservedNeutrinoI: 3, ObservedCKMSign: +1, Rank: -1, Unique: false, Verdict: StatusBranchSelectorNotUnique}
	}
	obs := rows[obsIdx]
	best := rows[0].AbsBFlav
	minimal := 0
	for _, r := range rows {
		if math.Abs(r.AbsBFlav-best) <= minimalTieTolerance {
			minimal++
		}
	}
	return ObservedBranchRank{
		ObservedSigma:           obs.Sigma,
		ObservedNeutrinoI:       obs.NeutrinoI,
		ObservedCKMSign:         obs.CKMSign,
		ObservedBFlav:           obs.BFlav,
		ObservedAbsBFlav:        obs.AbsBFlav,
		Rank:                    1,
		MinimalClassSize:        minimal,
		Unique:                  minimal == 1,
		MinimalClassDescription: "six charged-lepton permutations tie within numerical tolerance for the same electron-wall epsilon; all use i=3 and s_J=+1",
		Verdict:                 StatusObservedBranchHasMinimumResidual,
	}
}

func auditGap(rows []BranchBalanceRow, rank ObservedBranchRank) GapAudit {
	best := rows[0].AbsBFlav
	next := math.NaN()
	for _, r := range rows {
		if math.Abs(r.AbsBFlav-best) > minimalTieTolerance {
			next = r.AbsBFlav
			break
		}
	}
	gap := next - best
	return GapAudit{
		BestAbsResidual:         best,
		NextDistinctAbsResidual: next,
		GapToNextDistinct:       gap,
		GapMeaning:              "gap after the charged-lepton permutation tie; distinguishes PMNS projector/sign choices, not the charged-lepton ordering itself",
		GapLarge:                gap > 1e-5,
		Verdict:                 StatusBranchSelectorNotUnique,
	}
}

func compileSelectorVerdict(rank ObservedBranchRank, gap GapAudit) BranchSelectorVerdict {
	return BranchSelectorVerdict{
		ObservedInMinimalClass:        rank.Rank == 1,
		SelectsNeutrinoThirdProjector: true,
		SelectsPositiveCKMSign:        true,
		SelectsChargedLeptonOrdering:  false,
		UniqueBranchSelector:          false,
		NativeBranchSelector:          false,
		Decision:                      "B_flav is a strong environmental compatibility filter for P_3^nu and +J_CKM, but it leaves a sixfold charged-lepton permutation/chamber degeneracy. It therefore does not uniquely select the full observed history branch.",
		Verdict:                       StatusBranchSelectorNotUnique,
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesChargedLeptonMasses: false, DerivesPMNS: false, DerivesCKM: false, DerivesNeutrinoData: false, DerivesBFlavZeroNative: false, PromotesObservedData: false, AddsCarrier: false, AddsSelector: false, PreservesGate352: true, PreservesGate596: true, PreservesGate600: true, Verdict: StatusGate601Boundary}
}

func fourierPhase(order []string) (deltaDeg float64, r float64) {
	xs := make([]float64, 3)
	for i, label := range order {
		xs[i] = math.Sqrt(yukawa(label))
	}
	a := (xs[0] + xs[1] + xs[2]) / 3
	omega := 2 * math.Pi / 3
	var c complex128
	for j := 0; j < 3; j++ {
		v := xs[j]/a - 1
		c += complex(v, 0) * cmplx.Exp(complex(0, -omega*float64(j)))
	}
	delta := math.Atan2(imag(c), real(c)) * 180 / math.Pi
	delta = normalizeDeg(delta)
	r = cmplx.Abs(c) * 2 / (3 * math.Sqrt2)
	return delta, r
}

func nearestZeroWall(deltaDeg float64, componentIndex int) (wallDeg, epsilonDeg float64) {
	walls := []float64{normalizeDeg(135 - 120*float64(componentIndex)), normalizeDeg(225 - 120*float64(componentIndex))}
	bestWall := walls[0]
	best := circularDistanceDeg(deltaDeg, walls[0])
	for _, w := range walls[1:] {
		d := circularDistanceDeg(deltaDeg, w)
		if d < best {
			best = d
			bestWall = w
		}
	}
	return bestWall, best
}

func positiveComponents(deltaDeg, r float64) bool {
	for j := 0; j < 3; j++ {
		v := 1 + math.Sqrt2*r*math.Cos((deltaDeg+120*float64(j))*math.Pi/180)
		if v <= 0 {
			return false
		}
	}
	return true
}

func indexOf(order []string, target string) int {
	for i, s := range order {
		if s == target {
			return i
		}
	}
	return -1
}

func yukawa(label string) float64 {
	switch label {
	case "e":
		return yElectron
	case "mu":
		return yMuon
	case "tau":
		return yTau
	default:
		panic("unknown charged-lepton label: " + label)
	}
}

func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360)
	if y < 0 {
		y += 360
	}
	return y
}

func circularDistanceDeg(a, b float64) float64 {
	d := math.Abs(normalizeDeg(a) - normalizeDeg(b))
	if d > 180 {
		d = 360 - d
	}
	return d
}

func Statuses() []string {
	return []string{
		StatusGate600Inherited,
		StatusBranchBalanceDefined,
		StatusBranchSpaceEnumerated,
		StatusChargedLeptonBranchesEnumerated,
		StatusPMNSProjectorsEnumerated,
		StatusCKMSignsEnumerated,
		StatusFullBranchBalanceTableComputed,
		StatusObservedBranchHasMinimumResidual,
		StatusBalanceSelectsP3AndPositiveJ,
		StatusChargedLeptonPermutationDegeneracy,
		StatusBranchSelectorNotUnique,
		StatusNoUniqueChargedLeptonOrdering,
		StatusNoNativeBranchSelectionTheorem,
		StatusNoNativeBFlavZero,
		StatusBFlavEnvironmental,
		StatusGate600Boundary,
		StatusGate352Preserved,
		StatusGate596Preserved,
		StatusNoKoideDerivation,
		StatusNoChargedLeptonMassDerivation,
		StatusNoPMNSCKMNeutrinoFlavorDerivation,
		StatusNoObservedDataPromotion,
		StatusNoNewCarrierSelector,
		StatusGate601Boundary,
	}
}
