// Package su2l audits whether the right-singlet/conjugate hypercharge table
// selected by Gate 22, together with the finite scalar doublet charge from Gate
// 20, is sufficient to reconstruct the left SU(2)_L doublet hypercharges as a
// Yukawa-selection consequence.
//
// The rule tested here is the finite charge-balance equation
//
//	Y_L = Y_R - Y_Φ
//
// where Y_R is the matter-side right-singlet hypercharge produced by the odd
// T3_R branch, and Y_Φ is one of the two scalar doublet charges ±1/2 derived
// from the 2+2 Higgs/contact spectrum.  This is not yet a nonabelian SU(2)_L
// generator theorem; it is a charge-level doublet audit.
package su2l

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/matter/hyperaudit"
)

type Pairing struct {
	Name string

	RightCharge  float64
	ScalarCharge float64
	LeftCharge   float64
	Multiplicity int
}

type Orientation struct {
	Name string

	Up       Pairing
	Down     Pairing
	Neutrino Pairing
	Electron Pairing

	QuarkLeftHypercharge  float64
	LeptonLeftHypercharge float64
	QuarkDoubletDim       int
	LeptonDoubletDim      int

	MatchesStandardOrientation  bool
	MatchesConjugateOrientation bool
	ChargeBalanceExact          bool
}

type Analysis struct {
	Audit hyperaudit.Analysis

	ScalarWeight  float64
	ScalarCharges []float64

	OddRightTableAvailable bool
	Standard               Orientation
	Conjugate              Orientation

	ChargeLevelSU2LDoubletsDerived  bool
	NonabelianSU2LGeneratorsDerived bool
	NeutralSeedAmbiguity            bool
	ConjugationConventionMissing    bool
	YukawaIntertwinerDerived        bool
	RemainingUnknowns               []string
}

func BuildDefault() (Analysis, error) {
	a, err := hyperaudit.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(a hyperaudit.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !a.ChiralOrientationSelected || !a.Odd.MatchesRightSingletConjugateTable {
		return Analysis{}, fmt.Errorf("Gate 23 requires the Gate 22 odd right-singlet/conjugate table")
	}
	weight := a.T3R.Hypercharge.FundamentalWeight
	if math.Abs(weight) <= eps {
		return Analysis{}, fmt.Errorf("scalar doublet weight is zero")
	}
	counts := chargeMap(a.Odd.HyperchargeCounts, eps)

	std, err := buildOrientation("standard particle orientation", counts, weight, +1, eps)
	if err != nil {
		return Analysis{}, err
	}
	conj, err := buildOrientation("conjugate mirror orientation", counts, weight, -1, eps)
	if err != nil {
		return Analysis{}, err
	}

	derived := std.MatchesStandardOrientation && conj.MatchesConjugateOrientation
	return Analysis{
		Audit:                           a,
		ScalarWeight:                    weight,
		ScalarCharges:                   []float64{weight, -weight},
		OddRightTableAvailable:          true,
		Standard:                        std,
		Conjugate:                       conj,
		ChargeLevelSU2LDoubletsDerived:  derived,
		NonabelianSU2LGeneratorsDerived: false,
		NeutralSeedAmbiguity:            counts[canon(0, eps)] >= 2,
		ConjugationConventionMissing:    true,
		YukawaIntertwinerDerived:        false,
		RemainingUnknowns: []string{
			"U-13A-SU2L-GENERATORS: derive actual nonabelian SU(2)_L raising/lowering generators, not only charge-level doublet hypercharges",
			"U-14-CONJUGATION-CONVENTION: select particle vs conjugate orientation from a finite reality/CPT structure",
			"U-07-YUKAWA: construct explicit intertwiners after left-doublet charge channels are fixed",
		},
	}, nil
}

// orientationSign=+1 returns the usual particle orientation:
//
//	u_R(2/3)-(+1/2)=1/6, d_R(-1/3)-(-1/2)=1/6,
//	ν_R(0)-(+1/2)=-1/2, e_R(-1)-(-1/2)=-1/2.
//
// orientationSign=-1 returns the conjugate mirror orientation.
func buildOrientation(name string, counts map[int]int, w float64, orientationSign int, eps float64) (Orientation, error) {
	s := float64(orientationSign)
	up := Pairing{Name: "up-type", RightCharge: s * 2.0 / 3.0, ScalarCharge: s * w, Multiplicity: 3}
	down := Pairing{Name: "down-type", RightCharge: -s * 1.0 / 3.0, ScalarCharge: -s * w, Multiplicity: 3}
	neutrino := Pairing{Name: "neutrino-type", RightCharge: 0, ScalarCharge: s * w, Multiplicity: 1}
	electron := Pairing{Name: "electron-type", RightCharge: -s, ScalarCharge: -s * w, Multiplicity: 1}

	pairings := []*Pairing{&up, &down, &neutrino, &electron}
	for _, p := range pairings {
		have := counts[canon(p.RightCharge, eps)]
		if have < p.Multiplicity {
			return Orientation{}, fmt.Errorf("%s requires charge %.12g multiplicity %d; only have %d", name, p.RightCharge, p.Multiplicity, have)
		}
		p.LeftCharge = canonFloat(p.RightCharge-p.ScalarCharge, eps)
	}

	quarkExact := close(up.LeftCharge, down.LeftCharge, eps)
	leptonExact := close(neutrino.LeftCharge, electron.LeftCharge, eps)
	o := Orientation{
		Name:                  name,
		Up:                    up,
		Down:                  down,
		Neutrino:              neutrino,
		Electron:              electron,
		QuarkLeftHypercharge:  up.LeftCharge,
		LeptonLeftHypercharge: neutrino.LeftCharge,
		QuarkDoubletDim:       up.Multiplicity + down.Multiplicity,
		LeptonDoubletDim:      neutrino.Multiplicity + electron.Multiplicity,
		ChargeBalanceExact:    quarkExact && leptonExact,
	}
	o.MatchesStandardOrientation = close(o.QuarkLeftHypercharge, 1.0/6.0, eps) && close(o.LeptonLeftHypercharge, -1.0/2.0, eps) && o.QuarkDoubletDim == 6 && o.LeptonDoubletDim == 2
	o.MatchesConjugateOrientation = close(o.QuarkLeftHypercharge, -1.0/6.0, eps) && close(o.LeptonLeftHypercharge, 1.0/2.0, eps) && o.QuarkDoubletDim == 6 && o.LeptonDoubletDim == 2
	return o, nil
}

func chargeMap(counts []hyperaudit.ChargeCount, eps float64) map[int]int {
	out := map[int]int{}
	for _, c := range counts {
		out[canon(c.Charge, eps)] = c.Count
	}
	return out
}

func canon(x float64, eps float64) int {
	return int(math.Round(canonFloat(x, eps) * 6.0))
}

func canonFloat(x float64, eps float64) float64 {
	if math.Abs(x) < eps {
		return 0
	}
	return math.Round(x*1e12) / 1e12
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatPairing(p Pairing) string {
	return fmt.Sprintf("%s: Y_R=%.6g, Y_Φ=%.6g ⇒ Y_L=%.6g ×%d", p.Name, p.RightCharge, p.ScalarCharge, p.LeftCharge, p.Multiplicity)
}

func FormatOrientation(o Orientation) string {
	parts := []string{
		fmt.Sprintf("%s", o.Name),
		FormatPairing(o.Up),
		FormatPairing(o.Down),
		FormatPairing(o.Neutrino),
		FormatPairing(o.Electron),
		fmt.Sprintf("Q_L: Y=%.6g ×%d", o.QuarkLeftHypercharge, o.QuarkDoubletDim),
		fmt.Sprintf("L_L: Y=%.6g ×%d", o.LeptonLeftHypercharge, o.LeptonDoubletDim),
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(unknowns []string) string {
	xs := append([]string(nil), unknowns...)
	sort.Strings(xs)
	return "[" + strings.Join(xs, "; ") + "]"
}
