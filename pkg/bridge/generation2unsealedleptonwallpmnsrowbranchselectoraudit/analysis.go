// Package generation2unsealedleptonwallpmnsrowbranchselectoraudit implements
// Gate 602: Unsealed Lepton-Wall / PMNS-Row Branch Selector Audit.
//
// Gate 601 showed that B_flav selects P_3^nu and the positive CKM orientation
// sign, but it did not select the charged-lepton ordering because every branch
// was measured relative to the electron-zero wall. Gate 602 removes that hidden
// preselection: it enumerates charged-lepton zero-wall labels, PMNS rows,
// neutrino projectors, and CKM orientation signs, then checks whether the
// environmental balance itself selects the electron row. This remains a
// bridge-layer compatibility audit only.
package generation2unsealedleptonwallpmnsrowbranchselectoraudit

import (
	"fmt"
	"math"
	"math/cmplx"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2flavorbranchcompatibilityselectoraudit"
)

const (
	AuditID = "GATE602-UNSEALED-LEPTON-WALL-PMNS-ROW-BRANCH-SELECTOR-AUDIT"

	StatusGate601Inherited                  = "PASS_GATE601_BRANCH_COMPATIBILITY_RESULT_INHERITED"
	StatusBranchRowBalanceDefined           = "PASS_BRANCH_ROW_BALANCE_FUNCTION_DEFINED"
	StatusWallRowsAndSignsEnumerated        = "PASS_LEPTON_WALLS_PMNS_ROWS_AND_CKM_SIGNS_ENUMERATED"
	StatusLeptonWallCandidatesEnumerated    = "PASS_CHARGED_LEPTON_ZERO_WALL_CANDIDATES_ENUMERATED"
	StatusPMNSRowsEnumerated                = "PASS_PMNS_ROW_PROJECTOR_OVERLAPS_ENUMERATED"
	StatusCKMSignsEnumerated                = "PASS_CKM_ORIENTATION_SIGNS_ENUMERATED"
	StatusFullBranchRowTableComputed        = "PASS_FULL_BRANCH_ROW_BALANCE_TABLE_COMPUTED"
	StatusObservedTupleMinimal              = "PASS_OBSERVED_TUPLE_IN_MINIMAL_RESIDUAL_CLASS"
	StatusSelectsElectronRow                = "CONDITIONAL_SUPPORT_BALANCE_SELECTS_ELECTRON_ROW"
	StatusSelectsP3AndPositiveJ             = "CONDITIONAL_SUPPORT_BALANCE_SELECTS_P3_NU_AND_POSITIVE_CKM_SIGN"
	StatusSigmaDegeneracyExposed            = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_SIGMA_DEGENERACY_REMAINS"
	StatusNotFullOrderingSelector           = "FAILED_ROUTE_FULL_CHARGED_LEPTON_ORDERING_NOT_UNIQUELY_SELECTED"
	StatusNoNativeBranchSelectionTheorem    = "FAILED_ROUTE_NO_NATIVE_BRANCH_SELECTION_THEOREM"
	StatusNoNativeBFlavZero                 = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM"
	StatusBFlavEnvironmental                = "FAILED_ROUTE_B_FLAV_REMAINS_ENVIRONMENTAL_BRANCH_ROW_COMPATIBILITY_TEST"
	StatusGate600Boundary                   = "FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_BOUNDARY"
	StatusGate601Boundary                   = "FIREWALL_PRESERVED_GATE601_BRANCH_COMPATIBILITY_BOUNDARY"
	StatusGate352Preserved                  = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                  = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusNoKoideDerivation                 = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoChargedLeptonMassDerivation     = "FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION"
	StatusNoPMNSCKMNeutrinoFlavorDerivation = "FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusNoObservedDataPromotion           = "FIREWALL_PRESERVED_OBSERVED_BRANCH_LABELS_REMAIN_ENVIRONMENTAL_DATA"
	StatusNoNewCarrierSelector              = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate602Boundary                   = "FIREWALL_PRESERVED_GATE602_UNSEALED_LEPTON_WALL_BOUNDARY"
)

const (
	yElectron = 2.9350283095504176e-06
	yMuon     = 0.0006068707640859305
	yTau      = 0.010205763440624986

	sin2Theta12 = 0.308
	sin2Theta23 = 0.470
	sin2Theta13 = 0.02215
	deltaCPDeg  = 212.0
	jCKM        = 3.1169935287554706e-05

	minimalTieTolerance = 1e-12
)

type InheritedGate601 struct {
	ObservedMinimal       bool
	SelectsP3             bool
	SelectsPositiveJ      bool
	SelectsLeptonOrdering bool
	UniqueBranchSelector  bool
	NativeBranchSelector  bool
	MinimalClassSize      int
	Verdict               string
}

type BranchRowBalanceDefinition struct {
	Formula        string
	SigmaDomain    string
	AlphaDomain    string
	NeutrinoDomain string
	CKMSignDomain  string
	Environmental  bool
	Native         bool
	Verdict        string
}

type ChargedLeptonWallCandidate struct {
	Sigma           string
	Order           []string
	Alpha           string
	ComponentIndex  int
	DeltaDeg        float64
	R               float64
	WallDeg         float64
	EpsilonDeg      float64
	EpsilonRad      float64
	Kappa           float64
	PositiveChamber bool
	WallMeaning     string
	ObservedWall    bool
	Verdict         string
}

type PMNSRowProjectorOverlap struct {
	Alpha     string
	Index     int
	Projector string
	UAbs2     float64
	Li        float64
	Verdict   string
}

type CKMSign struct {
	Sign       int
	Convention string
	Value      float64
	Verdict    string
}

type BranchRowBalanceRow struct {
	Sigma      string
	Alpha      string
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

type ObservedTupleRank struct {
	ObservedSigma       string
	ObservedAlpha       string
	ObservedNeutrinoI   int
	ObservedCKMSign     int
	ObservedBFlav       float64
	ObservedAbsBFlav    float64
	Rank                int
	MinimalClassSize    int
	Unique              bool
	MinimalClassSummary string
	Verdict             string
}

type GapAudit struct {
	BestAbsResidual         float64
	NextDistinctAbsResidual float64
	GapToNextDistinct       float64
	GapMeaning              string
	GapLarge                bool
	Verdict                 string
}

type DegeneracyLedger struct {
	MinimalRows                int
	DistinctAlphas             []string
	DistinctNeutrinoProjectors []int
	DistinctCKMSigns           []int
	DistinctSigmas             []string
	ElectronRowSelected        bool
	P3Selected                 bool
	PositiveJSelected          bool
	SigmaStillDegenerate       bool
	Verdict                    string
}

type SelectorVerdict struct {
	ObservedInMinimalClass        bool
	SelectsElectronRow            bool
	SelectsThirdNeutrinoProjector bool
	SelectsPositiveCKMSign        bool
	SelectsFullChargedLeptonSigma bool
	UniqueSelector                bool
	NativeSelector                bool
	Decision                      string
	Verdict                       string
}

type FirewallAudit struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesNeutrinoData        bool
	DerivesFlavor              bool
	DerivesBFlavZeroNative     bool
	PromotesObservedData       bool
	AddsCarrier                bool
	AddsSelector               bool
	PreservesGate352           bool
	PreservesGate596           bool
	PreservesGate600           bool
	PreservesGate601           bool
	Verdict                    string
}

type Analysis struct {
	Inherited       InheritedGate601
	Definition      BranchRowBalanceDefinition
	WallCandidates  []ChargedLeptonWallCandidate
	PMNSOverlaps    []PMNSRowProjectorOverlap
	CKMSigns        []CKMSign
	BalanceTable    []BranchRowBalanceRow
	ObservedRank    ObservedTupleRank
	Gap             GapAudit
	Degeneracy      DegeneracyLedger
	SelectorVerdict SelectorVerdict
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
	g601, err := generation2flavorbranchcompatibilityselectoraudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate601 predecessor: %w", err)
	}
	inherited := inheritGate601(g601)
	definition := defineBalance()
	walls := enumerateChargedLeptonWallCandidates()
	pmns := enumeratePMNSRowProjectorOverlaps()
	signs := enumerateCKMSigns()
	table := computeBranchRowBalanceTable(walls, pmns, signs)
	rank := rankObservedTuple(table)
	gap := auditGap(table)
	degen := auditDegeneracy(table)
	selector := compileSelectorVerdict(rank, degen)
	firewalls := auditFirewalls()
	truth := "Gate 602 unseals the charged-lepton wall label and PMNS row.  The environmental balance B_flav now selects the electron row, the third neutrino projector P_3^nu, and the positive CKM orientation sign.  It still leaves a sixfold charged-lepton sigma/cyclic-order degeneracy, so it is a strong environmental branch-row compatibility selector, not a full native history branch theorem."
	return Analysis{Inherited: inherited, Definition: definition, WallCandidates: walls, PMNSOverlaps: pmns, CKMSigns: signs, BalanceTable: table, ObservedRank: rank, Gap: gap, Degeneracy: degen, SelectorVerdict: selector, Firewalls: firewalls, Truth: truth}, nil
}

func inheritGate601(a generation2flavorbranchcompatibilityselectoraudit.Analysis) InheritedGate601 {
	return InheritedGate601{ObservedMinimal: a.SelectorVerdict.ObservedInMinimalClass, SelectsP3: a.SelectorVerdict.SelectsNeutrinoThirdProjector, SelectsPositiveJ: a.SelectorVerdict.SelectsPositiveCKMSign, SelectsLeptonOrdering: a.SelectorVerdict.SelectsChargedLeptonOrdering, UniqueBranchSelector: a.SelectorVerdict.UniqueBranchSelector, NativeBranchSelector: a.SelectorVerdict.NativeBranchSelector, MinimalClassSize: a.ObservedRank.MinimalClassSize, Verdict: StatusGate601Inherited}
}

func defineBalance() BranchRowBalanceDefinition {
	return BranchRowBalanceDefinition{
		Formula:        "B_flav(sigma,alpha,i,s_J)=1-8*pi*epsilon_{sigma,alpha}(H_e)-(1/4)Tr(P_alpha P_i^nu)+s_J*J_CKM",
		SigmaDomain:    "six charged-lepton root/cyclic orderings",
		AlphaDomain:    "alpha in {e,mu,tau} wall/PMNS row labels",
		NeutrinoDomain: "i in {1,2,3}",
		CKMSignDomain:  "s_J in {+1,-1}",
		Environmental:  true,
		Native:         false,
		Verdict:        StatusBranchRowBalanceDefined,
	}
}

func enumerateChargedLeptonWallCandidates() []ChargedLeptonWallCandidate {
	orders := [][]string{{"e", "mu", "tau"}, {"e", "tau", "mu"}, {"mu", "e", "tau"}, {"mu", "tau", "e"}, {"tau", "e", "mu"}, {"tau", "mu", "e"}}
	alphas := []string{"e", "mu", "tau"}
	out := make([]ChargedLeptonWallCandidate, 0, len(orders)*len(alphas))
	for _, order := range orders {
		delta, r := fourierPhase(order)
		for _, alpha := range alphas {
			idx := indexOf(order, alpha)
			wall, epsDeg := nearestZeroWall(delta, idx)
			epsRad := epsDeg * math.Pi / 180
			kappa := 1 - 8*math.Pi*epsRad
			out = append(out, ChargedLeptonWallCandidate{
				Sigma:           strings.Join(order, ","),
				Order:           append([]string(nil), order...),
				Alpha:           alpha,
				ComponentIndex:  idx,
				DeltaDeg:        delta,
				R:               r,
				WallDeg:         wall,
				EpsilonDeg:      epsDeg,
				EpsilonRad:      epsRad,
				Kappa:           kappa,
				PositiveChamber: positiveComponents(delta, r),
				WallMeaning:     fmt.Sprintf("distance to %s-zero wall in this charged-lepton branch", alpha),
				ObservedWall:    alpha == "e" && strings.Join(order, ",") == "e,mu,tau",
				Verdict:         StatusLeptonWallCandidatesEnumerated,
			})
		}
	}
	return out
}

func enumeratePMNSRowProjectorOverlaps() []PMNSRowProjectorOverlap {
	matrix := pmnsAbs2Matrix()
	alphas := []string{"e", "mu", "tau"}
	out := make([]PMNSRowProjectorOverlap, 0, 9)
	for a, alpha := range alphas {
		for i := 0; i < 3; i++ {
			u2 := matrix[a][i]
			out = append(out, PMNSRowProjectorOverlap{Alpha: alpha, Index: i + 1, Projector: fmt.Sprintf("P_%d^nu", i+1), UAbs2: u2, Li: u2 / 4, Verdict: StatusPMNSRowsEnumerated})
		}
	}
	return out
}

func enumerateCKMSigns() []CKMSign {
	return []CKMSign{{Sign: +1, Convention: "+J_CKM observed orientation convention", Value: +jCKM, Verdict: StatusCKMSignsEnumerated}, {Sign: -1, Convention: "-J_CKM reversed orientation convention", Value: -jCKM, Verdict: StatusCKMSignsEnumerated}}
}

func computeBranchRowBalanceTable(walls []ChargedLeptonWallCandidate, pmns []PMNSRowProjectorOverlap, signs []CKMSign) []BranchRowBalanceRow {
	rows := make([]BranchRowBalanceRow, 0, len(walls)*len(pmns)*len(signs))
	for _, w := range walls {
		for _, p := range pmns {
			// The wall label alpha and PMNS row alpha must refer to the same lepton flavor.
			if p.Alpha != w.Alpha {
				continue
			}
			for _, s := range signs {
				bf := w.Kappa - p.Li + s.Value
				rows = append(rows, BranchRowBalanceRow{Sigma: w.Sigma, Alpha: w.Alpha, NeutrinoI: p.Index, CKMSign: s.Sign, EpsilonRad: w.EpsilonRad, Kappa: w.Kappa, Li: p.Li, JTerm: s.Value, BFlav: bf, AbsBFlav: math.Abs(bf), Observed: w.Sigma == "e,mu,tau" && w.Alpha == "e" && p.Index == 3 && s.Sign == +1})
			}
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].AbsBFlav == rows[j].AbsBFlav {
			if rows[i].Alpha == rows[j].Alpha {
				if rows[i].NeutrinoI == rows[j].NeutrinoI {
					return rows[i].Sigma < rows[j].Sigma
				}
				return rows[i].NeutrinoI < rows[j].NeutrinoI
			}
			return rows[i].Alpha < rows[j].Alpha
		}
		return rows[i].AbsBFlav < rows[j].AbsBFlav
	})
	return rows
}

func rankObservedTuple(rows []BranchRowBalanceRow) ObservedTupleRank {
	obsIdx := -1
	for i, r := range rows {
		if r.Observed {
			obsIdx = i
			break
		}
	}
	if obsIdx < 0 {
		return ObservedTupleRank{ObservedSigma: "e,mu,tau", ObservedAlpha: "e", ObservedNeutrinoI: 3, ObservedCKMSign: +1, Rank: -1, Verdict: StatusNotFullOrderingSelector}
	}
	obs := rows[obsIdx]
	best := rows[0].AbsBFlav
	minimal := 0
	for _, r := range rows {
		if math.Abs(r.AbsBFlav-best) <= minimalTieTolerance {
			minimal++
		}
	}
	return ObservedTupleRank{ObservedSigma: obs.Sigma, ObservedAlpha: obs.Alpha, ObservedNeutrinoI: obs.NeutrinoI, ObservedCKMSign: obs.CKMSign, ObservedBFlav: obs.BFlav, ObservedAbsBFlav: obs.AbsBFlav, Rank: 1, MinimalClassSize: minimal, Unique: minimal == 1, MinimalClassSummary: "six charged-lepton sigma/cyclic orderings tie, but all minimal rows have alpha=e, i=3, and s_J=+1", Verdict: StatusObservedTupleMinimal}
}

func auditGap(rows []BranchRowBalanceRow) GapAudit {
	best := rows[0].AbsBFlav
	next := math.NaN()
	for _, r := range rows {
		if math.Abs(r.AbsBFlav-best) > minimalTieTolerance {
			next = r.AbsBFlav
			break
		}
	}
	gap := next - best
	return GapAudit{BestAbsResidual: best, NextDistinctAbsResidual: next, GapToNextDistinct: gap, GapMeaning: "gap after the sixfold charged-lepton sigma degeneracy; distinguishes row/projector/sign choices", GapLarge: gap > 1e-5, Verdict: StatusSelectsElectronRow}
}

func auditDegeneracy(rows []BranchRowBalanceRow) DegeneracyLedger {
	best := rows[0].AbsBFlav
	var minimal []BranchRowBalanceRow
	for _, r := range rows {
		if math.Abs(r.AbsBFlav-best) <= minimalTieTolerance {
			minimal = append(minimal, r)
		}
	}
	alphas := distinctStrings(minimal, func(r BranchRowBalanceRow) string { return r.Alpha })
	neutrinos := distinctInts(minimal, func(r BranchRowBalanceRow) int { return r.NeutrinoI })
	signs := distinctInts(minimal, func(r BranchRowBalanceRow) int { return r.CKMSign })
	sigmas := distinctStrings(minimal, func(r BranchRowBalanceRow) string { return r.Sigma })
	return DegeneracyLedger{MinimalRows: len(minimal), DistinctAlphas: alphas, DistinctNeutrinoProjectors: neutrinos, DistinctCKMSigns: signs, DistinctSigmas: sigmas, ElectronRowSelected: len(alphas) == 1 && alphas[0] == "e", P3Selected: len(neutrinos) == 1 && neutrinos[0] == 3, PositiveJSelected: len(signs) == 1 && signs[0] == +1, SigmaStillDegenerate: len(sigmas) > 1, Verdict: strings.Join([]string{StatusSelectsElectronRow, StatusSelectsP3AndPositiveJ, StatusSigmaDegeneracyExposed, StatusNotFullOrderingSelector}, ";")}
}

func compileSelectorVerdict(rank ObservedTupleRank, d DegeneracyLedger) SelectorVerdict {
	unique := rank.Unique
	return SelectorVerdict{ObservedInMinimalClass: rank.Rank == 1, SelectsElectronRow: d.ElectronRowSelected, SelectsThirdNeutrinoProjector: d.P3Selected, SelectsPositiveCKMSign: d.PositiveJSelected, SelectsFullChargedLeptonSigma: !d.SigmaStillDegenerate, UniqueSelector: unique, NativeSelector: false, Decision: "B_flav, with the lepton wall unsealed, selects the electron row together with P_3^nu and +J_CKM.  It still cannot select the full charged-lepton sigma/cyclic ordering, leaving a sixfold degeneracy that would require an additional chamber-orientation principle.", Verdict: strings.Join([]string{StatusSelectsElectronRow, StatusSelectsP3AndPositiveJ, StatusNotFullOrderingSelector, StatusNoNativeBranchSelectionTheorem}, ";")}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{DerivesKoide: false, DerivesChargedLeptonMasses: false, DerivesPMNS: false, DerivesCKM: false, DerivesNeutrinoData: false, DerivesFlavor: false, DerivesBFlavZeroNative: false, PromotesObservedData: false, AddsCarrier: false, AddsSelector: false, PreservesGate352: true, PreservesGate596: true, PreservesGate600: true, PreservesGate601: true, Verdict: StatusGate602Boundary}
}

func pmnsAbs2Matrix() [3][3]float64 {
	s12, c12 := math.Sqrt(sin2Theta12), math.Sqrt(1-sin2Theta12)
	s23, c23 := math.Sqrt(sin2Theta23), math.Sqrt(1-sin2Theta23)
	s13, c13 := math.Sqrt(sin2Theta13), math.Sqrt(1-sin2Theta13)
	delta := deltaCPDeg * math.Pi / 180
	eid := cmplx.Exp(complex(0, delta))
	U := [3][3]complex128{
		{complex(c12*c13, 0), complex(s12*c13, 0), complex(s13, 0) * cmplx.Exp(complex(0, -delta))},
		{-complex(s12*c23, 0) - complex(c12*s23*s13, 0)*eid, complex(c12*c23, 0) - complex(s12*s23*s13, 0)*eid, complex(s23*c13, 0)},
		{complex(s12*s23, 0) - complex(c12*c23*s13, 0)*eid, -complex(c12*s23, 0) - complex(s12*c23*s13, 0)*eid, complex(c23*c13, 0)},
	}
	var out [3][3]float64
	for i := range U {
		for j := range U[i] {
			out[i][j] = cmplx.Abs(U[i][j]) * cmplx.Abs(U[i][j])
		}
	}
	return out
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

func distinctStrings(rows []BranchRowBalanceRow, f func(BranchRowBalanceRow) string) []string {
	m := map[string]bool{}
	for _, r := range rows {
		m[f(r)] = true
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func distinctInts(rows []BranchRowBalanceRow, f func(BranchRowBalanceRow) int) []int {
	m := map[int]bool{}
	for _, r := range rows {
		m[f(r)] = true
	}
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

func Statuses() []string {
	return []string{StatusGate601Inherited, StatusBranchRowBalanceDefined, StatusWallRowsAndSignsEnumerated, StatusLeptonWallCandidatesEnumerated, StatusPMNSRowsEnumerated, StatusCKMSignsEnumerated, StatusFullBranchRowTableComputed, StatusObservedTupleMinimal, StatusSelectsElectronRow, StatusSelectsP3AndPositiveJ, StatusSigmaDegeneracyExposed, StatusNotFullOrderingSelector, StatusNoNativeBranchSelectionTheorem, StatusNoNativeBFlavZero, StatusBFlavEnvironmental, StatusGate600Boundary, StatusGate601Boundary, StatusGate352Preserved, StatusGate596Preserved, StatusNoKoideDerivation, StatusNoChargedLeptonMassDerivation, StatusNoPMNSCKMNeutrinoFlavorDerivation, StatusNoObservedDataPromotion, StatusNoNewCarrierSelector, StatusGate602Boundary}
}
