// Package charge constructs the first explicit charge-polarizing bridge for the
// finite matter sector.
//
// Gate 15 showed that the active Higgs/contact spectrum alone is pair-
// degenerate (2+2), so it cannot canonically produce the physical 1+3 split of
// one temporal/lepton seed and three spatial/color seeds. The missing object is
// a charge-polarizing operator. In the covariant Witt/Fock sector the standard
// candidate is the B-L number operator
//
//	Q_{B-L} = (1/3) Σ_{i=1}^3 N_i - N_0,
//
// where mode 0 is the temporal/lepton mode and modes 1,2,3 are spatial/color
// seeds. This package verifies what that operator solves and what it does not:
// it supplies the required 1+3 charge polarization, but it does not license a
// direct identification of the pair-degenerate Higgs/contact eigenvalues with
// the three color modes.
package charge

import (
	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	matteraction "github.com/bagherbal/asha-engine/pkg/matter/action"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type Cluster struct {
	Multiplicity int
	Mean         float64
	Spread       float64
}

type Analysis struct {
	Action matteraction.Analysis

	OneParticleChargeOperator linear.Matrix // 4x4 diag(-1,1/3,1/3,1/3).
	FockChargeOperator        linear.Matrix // 16x16 diagonal B-L operator.

	OneParticleChargeSpectrum   []float64
	OneParticleChargeClusters   []Cluster
	ChargePolarizesOnePlusThree bool

	VacuumCharge                  float64
	TraceOneParticleCharge        float64
	TraceOneParticleChargeSquared float64
	FockTrace                     float64
	FockEvenTrace                 float64
	FockOddTrace                  float64

	CommutatorWithFockResponseNorm float64

	DirectScalarToColorIsotropyPossible bool
	BestSpatialScalarAnisotropy         float64
	BestTemporalScalarWeight            float64
	BestSpatialScalarWeights            []float64

	Resolution       string
	RemainingUnknown string
}

func BuildDefault() (Analysis, error) {
	a, err := matteraction.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(a matteraction.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if a.Bridge.Fock.ModeCount() != 4 {
		return Analysis{}, fmt.Errorf("B-L bridge expects four covariant Fock modes, got %d", a.Bridge.Fock.ModeCount())
	}

	oneCharges := []float64{-1.0, 1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}
	q1 := linear.Diagonal(oneCharges)
	clusters := clustersFromValues(oneCharges, eps)
	polarizes := hasClusterMultiset(clusters, []int{1, 3})

	qF := linear.NewMatrix(a.Bridge.Fock.StateCount(), a.Bridge.Fock.StateCount())
	vacuumCharge := math.NaN()
	fockTrace := 0.0
	evenTrace := 0.0
	oddTrace := 0.0
	for i, state := range a.Bridge.Fock.States {
		charge := bMinusL(state)
		qF.Set(i, i, charge)
		fockTrace += charge
		if state.ExcitationNumber()%2 == 0 {
			evenTrace += charge
		} else {
			oddTrace += charge
		}
		if state.IsVacuum() {
			vacuumCharge = charge
		}
	}

	comm, err := commutator(qF, a.Operator)
	if err != nil {
		return Analysis{}, err
	}

	scalarWeights := append([]float64(nil), a.OneParticleWeights...)
	bestAniso, bestTemporal, bestSpatial := bestSpatialAnisotropy(scalarWeights)
	scalarColorIsotropic := bestAniso <= eps

	traceQ := 0.0
	traceQ2 := 0.0
	for _, q := range oneCharges {
		traceQ += q
		traceQ2 += q * q
	}

	resolution := "B-L supplies the missing 1+3 charge polarization on the Fock one-particle modes."
	remaining := "The pair-degenerate Higgs/contact spectrum must not be identified directly with three color modes; the next bridge needs a tensor-factor or representation action separating charge polarization from scalar mixing."

	return Analysis{
		Action:                              a,
		OneParticleChargeOperator:           q1,
		FockChargeOperator:                  qF,
		OneParticleChargeSpectrum:           oneCharges,
		OneParticleChargeClusters:           clusters,
		ChargePolarizesOnePlusThree:         polarizes,
		VacuumCharge:                        vacuumCharge,
		TraceOneParticleCharge:              traceQ,
		TraceOneParticleChargeSquared:       traceQ2,
		FockTrace:                           fockTrace,
		FockEvenTrace:                       evenTrace,
		FockOddTrace:                        oddTrace,
		CommutatorWithFockResponseNorm:      comm.FrobeniusNorm(),
		DirectScalarToColorIsotropyPossible: scalarColorIsotropic,
		BestSpatialScalarAnisotropy:         bestAniso,
		BestTemporalScalarWeight:            bestTemporal,
		BestSpatialScalarWeights:            bestSpatial,
		Resolution:                          resolution,
		RemainingUnknown:                    remaining,
	}, nil
}

func bMinusL(state spinor.FockState) float64 {
	if len(state.Occupation) == 0 {
		return 0
	}
	charge := 0.0
	if state.Occupation[0] {
		charge -= 1.0
	}
	for i := 1; i < len(state.Occupation); i++ {
		if state.Occupation[i] {
			charge += 1.0 / 3.0
		}
	}
	return charge
}

func commutator(a, b linear.Matrix) (linear.Matrix, error) {
	ab, err := a.Mul(b)
	if err != nil {
		return linear.Matrix{}, err
	}
	ba, err := b.Mul(a)
	if err != nil {
		return linear.Matrix{}, err
	}
	return ab.Sub(ba)
}

func clustersFromValues(values []float64, eps float64) []Cluster {
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)
	out := make([]Cluster, 0)
	for i := 0; i < len(sorted); {
		start := sorted[i]
		count := 1
		sum := start
		minV, maxV := start, start
		j := i + 1
		for ; j < len(sorted); j++ {
			if math.Abs(sorted[j]-start) > eps {
				break
			}
			count++
			sum += sorted[j]
			if sorted[j] < minV {
				minV = sorted[j]
			}
			if sorted[j] > maxV {
				maxV = sorted[j]
			}
		}
		out = append(out, Cluster{Multiplicity: count, Mean: sum / float64(count), Spread: maxV - minV})
		i = j
	}
	return out
}

func hasClusterMultiset(clusters []Cluster, wanted []int) bool {
	if len(clusters) != len(wanted) {
		return false
	}
	seen := map[int]int{}
	for _, c := range clusters {
		seen[c.Multiplicity]++
	}
	for _, w := range wanted {
		seen[w]--
	}
	for _, left := range seen {
		if left != 0 {
			return false
		}
	}
	return true
}

func bestSpatialAnisotropy(weights []float64) (best float64, temporal float64, spatial []float64) {
	if len(weights) != 4 {
		return math.Inf(1), math.NaN(), nil
	}
	best = math.Inf(1)
	for t := range weights {
		s := make([]float64, 0, 3)
		for i, w := range weights {
			if i != t {
				s = append(s, w)
			}
		}
		minV, maxV := s[0], s[0]
		for _, w := range s[1:] {
			if w < minV {
				minV = w
			}
			if w > maxV {
				maxV = w
			}
		}
		aniso := maxV - minV
		if aniso < best {
			best = aniso
			temporal = weights[t]
			spatial = s
		}
	}
	return best, temporal, spatial
}

func FormatClusters(clusters []Cluster) string {
	out := "["
	for i, c := range clusters {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("m=%d@%.10g", c.Multiplicity, c.Mean)
	}
	return out + "]"
}

func FormatFloatSlice(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10g", v)
	}
	return out + "]"
}
