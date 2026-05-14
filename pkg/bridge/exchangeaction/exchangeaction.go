// Package exchangeaction performs Gate 63: finite exchange-action / propagator
// normalization search.
//
// Gates 60--62 built a chain:
//
//	finite current-sector kinetic traces
//	-> signed Clifford/Lorentz Fierz coefficient c_LR = -2
//	-> conditional scalar exchange kernel under the two possible exchange
//	   orientations.
//
// Gate 63 asks whether the finite engine itself selects the exchange-action
// orientation and propagator weights.  The answer, at this stage, is no: the
// current-sector kinetic traces provide normalization diagnostics, but they do
// not by themselves define propagator denominators, relative couplings, or the
// sign of the integrated-out exchange action.  This package computes the
// natural diagnostic families and keeps the physical kernel marked open.
package exchangeaction

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/exchangekernel"
)

type SectorPropagatorDiagnostic struct {
	Sector         string
	GeneratorCount int
	KineticTrace   float64
	TraceWeight    float64

	AttractiveUnitContribution  float64
	UnitPropagatorContribution  float64
	InverseKineticContribution  float64
	KineticWeightedContribution float64

	PropagatorWeightDerived bool
}

type Analysis struct {
	Exchange exchangekernel.Analysis

	SectorDiagnostics []SectorPropagatorDiagnostic

	UnitAttractiveKernel          float64
	InverseKineticDiagnostic      float64
	KineticWeightedDiagnostic     float64
	DominantUnitSector            string
	DominantKineticWeightedSector string

	CurrentSectorKineticDataAvailable bool
	UnitPropagatorBranchAvailable     bool
	InverseKineticBranchAvailable     bool
	KineticWeightedBranchAvailable    bool

	ExchangeActionSignDerived       bool
	PropagatorWeightsDerived        bool
	RelativeCurrentCouplingsDerived bool
	FiniteExchangeKernelDerived     bool
	AttractiveScalarChannelDerived  bool
	UpDownSplittingDerived          bool
	RegulatorCriticalityDerived     bool
	CondensationClaimAllowed        bool
	HiddenObservedInputUsed         bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ex, err := exchangekernel.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ex, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(ex exchangekernel.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if !ex.ConditionalAttractiveBranchAvailable {
		return Analysis{}, fmt.Errorf("Gate 63 requires Gate 62 conditional exchange-kernel diagnostics")
	}
	if len(ex.SectorDiagnostics) == 0 {
		return Analysis{}, fmt.Errorf("no sector exchange diagnostics available")
	}

	sectors := make([]SectorPropagatorDiagnostic, 0, len(ex.SectorDiagnostics))
	unit := 0.0
	inverseKinetic := 0.0
	kineticWeighted := 0.0
	dominantUnit := ""
	dominantKinetic := ""
	maxUnit := math.Inf(-1)
	maxKinetic := math.Inf(-1)
	kineticData := true

	for _, s := range ex.SectorDiagnostics {
		if s.TraceWeight < -eps || s.AttractiveUnitContribution < -eps {
			return Analysis{}, fmt.Errorf("sector %s has invalid weights", s.Sector)
		}
		if s.TraceWeight <= eps {
			kineticData = false
		}
		// The unit branch is the conditional attractive branch from Gate 62.
		unitContribution := s.AttractiveUnitContribution
		// The inverse-kinetic branch is a diagnostic only: it asks what would
		// happen if the current-sector propagator denominator scaled with the
		// sector kinetic trace. This is not a derivation.
		invContribution := 0.0
		if s.TraceWeight > eps {
			// Since TraceWeight = K_sector / K_total, dividing by the trace weight
			// removes the kinetic suppression and exposes how sensitive the kernel
			// would be to a naive inverse-kinetic denominator.
			invContribution = unitContribution / s.TraceWeight
		}
		// The kinetic-weighted branch is also diagnostic: it asks what happens
		// if kinetic trace weights are also used as propagator/coupling weights.
		weightedContribution := unitContribution * s.TraceWeight

		unit += unitContribution
		inverseKinetic += invContribution
		kineticWeighted += weightedContribution
		if unitContribution > maxUnit {
			maxUnit = unitContribution
			dominantUnit = s.Sector
		}
		if weightedContribution > maxKinetic {
			maxKinetic = weightedContribution
			dominantKinetic = s.Sector
		}
		sectors = append(sectors, SectorPropagatorDiagnostic{
			Sector:                      s.Sector,
			GeneratorCount:              s.GeneratorCount,
			KineticTrace:                s.TraceWeight * 66.66666666666667, // diagnostic reconstruction from previous finite trace total
			TraceWeight:                 s.TraceWeight,
			AttractiveUnitContribution:  s.AttractiveUnitContribution,
			UnitPropagatorContribution:  unitContribution,
			InverseKineticContribution:  invContribution,
			KineticWeightedContribution: weightedContribution,
			PropagatorWeightDerived:     false,
		})
	}
	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Sector < sectors[j].Sector })

	truth := "Gate 63 finds no finite exchange-action theorem yet. The signed Fierz data plus the attractive branch from Gate 62 gives a unit-kernel diagnostic G_hat=2, and finite current-sector kinetic weights allow inverse-kinetic and kinetic-weighted diagnostic families. However, none of these families is selected by a finite action, propagator denominator, or relative coupling theorem. Therefore the native four-fermion/NJL kernel remains open and condensation cannot be claimed."

	return Analysis{
		Exchange:                          ex,
		SectorDiagnostics:                 sectors,
		UnitAttractiveKernel:              unit,
		InverseKineticDiagnostic:          inverseKinetic,
		KineticWeightedDiagnostic:         kineticWeighted,
		DominantUnitSector:                dominantUnit,
		DominantKineticWeightedSector:     dominantKinetic,
		CurrentSectorKineticDataAvailable: kineticData,
		UnitPropagatorBranchAvailable:     unit > eps,
		InverseKineticBranchAvailable:     inverseKinetic > eps,
		KineticWeightedBranchAvailable:    kineticWeighted > eps,
		ExchangeActionSignDerived:         false,
		PropagatorWeightsDerived:          false,
		RelativeCurrentCouplingsDerived:   false,
		FiniteExchangeKernelDerived:       false,
		AttractiveScalarChannelDerived:    false,
		UpDownSplittingDerived:            false,
		RegulatorCriticalityDerived:       false,
		CondensationClaimAllowed:          false,
		HiddenObservedInputUsed:           false,
		TruthStatement:                    truth,
		RecommendedNextGate:               "Gate 64 — Finite Propagator from B-sector / Contact Spectrum Search",
		RemainingUnknowns: []string{
			"U-20D2A-EXCHANGE-ACTION-SIGN: derive the sign of finite current exchange from an action, not from an NJL convention",
			"U-20D2B-PROPAGATOR-DENOMINATORS: derive sector propagator weights from finite spectra or kinetic operators",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector couplings instead of unit or trace-weight diagnostics",
			"U-20D4-UP-DOWN-SPLITTING: derive an operator that selects top-like up over bottom-like down",
			"U-20D6-CRITICAL-REGULATOR: derive C_reg for the finite NJL gap equation",
		},
	}, nil
}

func FormatSectorDiagnostics(xs []SectorPropagatorDiagnostic) string {
	ys := append([]SectorPropagatorDiagnostic(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector < ys[j].Sector })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d, w=%.6f, unit=%.10f, invK=%.10f, Kweighted=%.10f)", x.Sector, x.GeneratorCount, x.TraceWeight, x.UnitPropagatorContribution, x.InverseKineticContribution, x.KineticWeightedContribution))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
