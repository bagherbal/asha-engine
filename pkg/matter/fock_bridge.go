// Package matter contains bridge-layer analyses that connect the finite spinor
// basis to the Boolean-Octonionic contact/Higgs data.
//
// The package is deliberately conservative: a dimension resonance is not treated
// as a full particle-physics theorem. It records what the finite data supports
// and labels the missing canonical maps explicitly.
package matter

import (
	"sync"

	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/dynamics/higgspotential"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type FockContactBridge struct {
	Fock      spinor.FockSpace
	Potential higgspotential.Analysis

	FockModeCount         int
	FockStateCount        int
	OneParticleStateCount int
	TemporalModeCount     int
	SpatialModeCount      int
	SterileVacuumCount    int
	QuarkSeedCount        int
	LeptonSeedCount       int

	ActiveHiggsDirections       int
	ProtectedContactDirections  int
	PairDegenerateHiggsSpectrum bool

	ModeToActiveScalarMatch      bool
	SpatialToProtectedMatch      bool
	VacuumSterilitySeedAvailable bool

	CanonicalEmbeddingConstructed bool
	YukawaOperatorConstructed     bool
}

var (
	fockContactBridgeDefaultOnce  sync.Once
	fockContactBridgeDefaultValue FockContactBridge
	fockContactBridgeDefaultErr   error
)

func BuildDefaultFockContactBridge() (FockContactBridge, error) {
	fockContactBridgeDefaultOnce.Do(func() {
		fockContactBridgeDefaultValue, fockContactBridgeDefaultErr = buildFockContactBridgeDefaultUncached()
	})
	return fockContactBridgeDefaultValue, fockContactBridgeDefaultErr
}

func buildFockContactBridgeDefaultUncached() (FockContactBridge, error) {
	f, err := spinor.NewCovariantPhaseFockSpace(4)
	if err != nil {
		return FockContactBridge{}, err
	}
	p, err := higgspotential.BuildDefault()
	if err != nil {
		return FockContactBridge{}, err
	}
	return BuildFockContactBridge(f, p, 1e-12)
}

func BuildFockContactBridge(f spinor.FockSpace, p higgspotential.Analysis, eps float64) (FockContactBridge, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if f.ModeCount() == 0 || f.StateCount() == 0 {
		return FockContactBridge{}, fmt.Errorf("empty Fock space")
	}
	oneParticle := f.OneParticleStates()
	sterileVacuumCount := 0
	quarkSeedCount := 0
	leptonSeedCount := 0

	for _, s := range f.States {
		if s.IsSterileVacuumCandidate(eps) {
			sterileVacuumCount++
		}
	}
	for _, s := range oneParticle {
		charge := s.BMinusL()
		if math.Abs(charge-(1.0/3.0)) < eps {
			quarkSeedCount++
		}
		if math.Abs(charge+1.0) < eps {
			leptonSeedCount++
		}
	}

	return FockContactBridge{
		Fock:                          f,
		Potential:                     p,
		FockModeCount:                 f.ModeCount(),
		FockStateCount:                f.StateCount(),
		OneParticleStateCount:         len(oneParticle),
		TemporalModeCount:             f.TemporalModeCount(),
		SpatialModeCount:              f.SpatialModeCount(),
		SterileVacuumCount:            sterileVacuumCount,
		QuarkSeedCount:                quarkSeedCount,
		LeptonSeedCount:               leptonSeedCount,
		ActiveHiggsDirections:         p.ActiveContactDimension,
		ProtectedContactDirections:    p.ProtectedContactDimension,
		PairDegenerateHiggsSpectrum:   p.PairDegenerateSpectrum,
		ModeToActiveScalarMatch:       f.ModeCount() == p.ActiveContactDimension,
		SpatialToProtectedMatch:       f.SpatialModeCount() == p.ProtectedContactDimension,
		VacuumSterilitySeedAvailable:  sterileVacuumCount == 1,
		CanonicalEmbeddingConstructed: false,
		YukawaOperatorConstructed:     false,
	}, nil
}
