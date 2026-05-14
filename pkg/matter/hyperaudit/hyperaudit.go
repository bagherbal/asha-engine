// Package hyperaudit audits the two chiral T3_R branches against standard
// hypercharge charge tables.
//
// Gate 21 found that the temporal weak-isospin seed T0=1/2-N0 becomes useful
// only after chiral restriction, but that both even and odd restrictions are
// algebraically possible.  This package performs the next physics-facing audit:
// it computes the matter-side hypercharge spectrum
//
//	Y_m = T3_R + (B-L)/2
//
// for both restrictions and compares the resulting multiplicities against two
// tables:
//
//  1. A right-singlet/conjugate Pati-Salam table containing
//     ν/e and u/d singlet charges together with conjugate partners.
//  2. The usual left-handed Standard-Model Weyl table with a ν^c state.
//
// The result is deliberately disciplined.  The odd branch exactly reproduces
// the right-singlet/conjugate hypercharge multiset, while the even branch
// produces exotic ±5/6 charges.  But the full Standard-Model left-doublet table
// is still not derived, because SU(2)_L doublet structure and charge-conjugation
// orientation remain missing.
package hyperaudit

import (
	"sync"

	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/t3r"
)

type ChargeCount struct {
	Charge float64
	Count  int
}

type BranchAudit struct {
	Name string

	HyperchargeCounts []ChargeCount

	MatchesRightSingletConjugateTable bool
	MatchesLeftHandedSMTable          bool
	ExoticCharges                     []ChargeCount
	ExoticChargePresent               bool

	RightSingletScore int
	LeftSMScore       int
}

type Analysis struct {
	T3R t3r.Analysis

	Even BranchAudit
	Odd  BranchAudit

	PreferredBranchName               string
	ChiralOrientationSelected         bool
	FullStandardModelTableDerived     bool
	RightSingletConjugateTableDerived bool
	SU2LDoubletBridgeMissing          bool
	ConjugationConventionMissing      bool
	RemainingUnknowns                 []string
}

var (
	hyperauditDefaultOnce  sync.Once
	hyperauditDefaultValue Analysis
	hyperauditDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	hyperauditDefaultOnce.Do(func() {
		hyperauditDefaultValue, hyperauditDefaultErr = buildHyperauditDefaultUncached()
	})
	return hyperauditDefaultValue, hyperauditDefaultErr
}

func buildHyperauditDefaultUncached() (Analysis, error) {
	a, err := t3r.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(a t3r.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	even, err := auditBranch(a.ChiralEven, eps)
	if err != nil {
		return Analysis{}, err
	}
	odd, err := auditBranch(a.ChiralOdd, eps)
	if err != nil {
		return Analysis{}, err
	}

	preferred := ""
	selected := false
	if odd.MatchesRightSingletConjugateTable && !even.MatchesRightSingletConjugateTable {
		preferred = odd.Name
		selected = true
	}
	if even.MatchesRightSingletConjugateTable && !odd.MatchesRightSingletConjugateTable {
		preferred = even.Name
		selected = true
	}

	return Analysis{
		T3R:                               a,
		Even:                              even,
		Odd:                               odd,
		PreferredBranchName:               preferred,
		ChiralOrientationSelected:         selected,
		RightSingletConjugateTableDerived: selected,
		FullStandardModelTableDerived:     even.MatchesLeftHandedSMTable || odd.MatchesLeftHandedSMTable,
		SU2LDoubletBridgeMissing:          !(even.MatchesLeftHandedSMTable || odd.MatchesLeftHandedSMTable),
		ConjugationConventionMissing:      true,
		RemainingUnknowns: []string{
			"U-13-SU2L-DOUBLETS: derive the finite SU(2)_L doublet sector producing Q_L and L_L hypercharges",
			"U-14-CONJUGATION-CONVENTION: decide particle vs left-handed-conjugate orientation for the 16 Fock states",
			"U-07-YUKAWA: build Yukawa intertwiners only after SU(2)_L and conjugation are fixed",
		},
	}, nil
}

func auditBranch(c t3r.Candidate, eps float64) (BranchAudit, error) {
	counts, err := countsFromDiagonal(c.MatterHypercharge, eps)
	if err != nil {
		return BranchAudit{}, err
	}
	expectedRight := rightSingletConjugateExpected()
	expectedLeft := leftHandedSMExpected()
	rightScore := matchingScore(counts, expectedRight, eps)
	leftScore := matchingScore(counts, expectedLeft, eps)
	exotic := exoticCharges(counts, allowedUnion(expectedRight, expectedLeft), eps)
	return BranchAudit{
		Name:                              c.Name,
		HyperchargeCounts:                 counts,
		MatchesRightSingletConjugateTable: sameCounts(counts, expectedRight, eps),
		MatchesLeftHandedSMTable:          sameCounts(counts, expectedLeft, eps),
		ExoticCharges:                     exotic,
		ExoticChargePresent:               len(exotic) > 0,
		RightSingletScore:                 rightScore,
		LeftSMScore:                       leftScore,
	}, nil
}

func countsFromDiagonal(m linear.Matrix, eps float64) ([]ChargeCount, error) {
	if m.Rows() != m.Cols() {
		return nil, fmt.Errorf("expected square hypercharge matrix, got %dx%d", m.Rows(), m.Cols())
	}
	counts := make([]ChargeCount, 0)
	for i := 0; i < m.Rows(); i++ {
		q := canonical(m.At(i, i), eps)
		idx := -1
		for j := range counts {
			if close(counts[j].Charge, q, eps) {
				idx = j
				break
			}
		}
		if idx < 0 {
			counts = append(counts, ChargeCount{Charge: q})
		} else {
			counts[idx].Count++
			continue
		}
		counts[len(counts)-1].Count = 1
	}
	sortCounts(counts)
	return counts, nil
}

func rightSingletConjugateExpected() []ChargeCount {
	// Vectorlike/right-singlet-plus-conjugate multiset:
	// ν_R/ν^c-like neutral pair, e_R/e^c, u/d color triplets and conjugates.
	return []ChargeCount{
		{Charge: -1, Count: 1},
		{Charge: -2.0 / 3.0, Count: 3},
		{Charge: -1.0 / 3.0, Count: 3},
		{Charge: 0, Count: 2},
		{Charge: 1.0 / 3.0, Count: 3},
		{Charge: 2.0 / 3.0, Count: 3},
		{Charge: 1, Count: 1},
	}
}

func leftHandedSMExpected() []ChargeCount {
	// Standard left-handed Weyl one-generation table including ν^c:
	// Q_L: 6 states at 1/6, L_L: 2 states at -1/2,
	// u^c: 3 at -2/3, d^c: 3 at 1/3, e^c: 1 at 1, ν^c: 1 at 0.
	return []ChargeCount{
		{Charge: -2.0 / 3.0, Count: 3},
		{Charge: -1.0 / 2.0, Count: 2},
		{Charge: 0, Count: 1},
		{Charge: 1.0 / 6.0, Count: 6},
		{Charge: 1.0 / 3.0, Count: 3},
		{Charge: 1, Count: 1},
	}
}

func sameCounts(a, b []ChargeCount, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	aa := append([]ChargeCount(nil), a...)
	bb := append([]ChargeCount(nil), b...)
	sortCounts(aa)
	sortCounts(bb)
	for i := range aa {
		if !close(aa[i].Charge, bb[i].Charge, eps) || aa[i].Count != bb[i].Count {
			return false
		}
	}
	return true
}

func matchingScore(actual, expected []ChargeCount, eps float64) int {
	score := 0
	for _, e := range expected {
		for _, a := range actual {
			if close(a.Charge, e.Charge, eps) {
				if a.Count < e.Count {
					score += a.Count
				} else {
					score += e.Count
				}
				break
			}
		}
	}
	return score
}

func exoticCharges(actual, allowed []ChargeCount, eps float64) []ChargeCount {
	out := make([]ChargeCount, 0)
	for _, a := range actual {
		found := false
		for _, x := range allowed {
			if close(a.Charge, x.Charge, eps) {
				found = true
				break
			}
		}
		if !found {
			out = append(out, a)
		}
	}
	return out
}

func allowedUnion(groups ...[]ChargeCount) []ChargeCount {
	out := make([]ChargeCount, 0)
	for _, g := range groups {
		for _, c := range g {
			found := false
			for _, existing := range out {
				if close(existing.Charge, c.Charge, 1e-10) {
					found = true
					break
				}
			}
			if !found {
				out = append(out, ChargeCount{Charge: c.Charge})
			}
		}
	}
	sortCounts(out)
	return out
}

func sortCounts(c []ChargeCount) {
	sort.Slice(c, func(i, j int) bool { return c[i].Charge < c[j].Charge })
}

func canonical(v, eps float64) float64 {
	if math.Abs(v) < eps {
		return 0
	}
	rationals := []float64{-1, -5.0 / 6.0, -2.0 / 3.0, -1.0 / 2.0, -1.0 / 3.0, -1.0 / 6.0, 0, 1.0 / 6.0, 1.0 / 3.0, 1.0 / 2.0, 2.0 / 3.0, 5.0 / 6.0, 1}
	for _, r := range rationals {
		if math.Abs(v-r) <= eps {
			return r
		}
	}
	return v
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatCounts(counts []ChargeCount) string {
	parts := make([]string, 0, len(counts))
	for _, c := range counts {
		parts = append(parts, fmt.Sprintf("Y=%s×%d", formatCharge(c.Charge), c.Count))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func formatCharge(q float64) string {
	candidates := []struct {
		Value float64
		Text  string
	}{
		{-1, "-1"}, {-5.0 / 6.0, "-5/6"}, {-2.0 / 3.0, "-2/3"}, {-1.0 / 2.0, "-1/2"}, {-1.0 / 3.0, "-1/3"}, {-1.0 / 6.0, "-1/6"}, {0, "0"}, {1.0 / 6.0, "1/6"}, {1.0 / 3.0, "1/3"}, {1.0 / 2.0, "1/2"}, {2.0 / 3.0, "2/3"}, {5.0 / 6.0, "5/6"}, {1, "1"},
	}
	for _, c := range candidates {
		if math.Abs(q-c.Value) < 1e-10 {
			return c.Text
		}
	}
	return fmt.Sprintf("%.6g", q)
}

func FormatUnknowns(unknowns []string) string {
	return "[" + strings.Join(unknowns, "; ") + "]"
}
