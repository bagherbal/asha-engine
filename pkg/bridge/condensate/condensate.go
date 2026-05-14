// Package condensate audits the composite-Higgs / fermion-condensate direction.
//
// Gates 37-50 established a strong finite scalar/contact sector, but they also
// exposed a limit: scalar geometry alone does not derive the physical Higgs
// scale, full scalar SU(2)_L action, or fermion masses.  The older Higgs prelude
// points to the better next calculation: treat the Higgs candidate as a scalar
// bilinear / condensate built from the already-derived Fock matter and the
// gauge-compatible Yukawa channels.
//
// This package does not solve the one-loop effective potential, NJL gap
// equation, or CPT boundary problem.  It creates a theorem-gated ledger of what
// is already available for that computation and what remains genuinely open.
package condensate

import (
	"fmt"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawaintertwiner"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type CriticalComputation struct {
	ID          string
	Name        string
	Status      string
	Requirement string
}

type Analysis struct {
	Fock            spinor.FockSpace
	Yukawa          yukawaintertwiner.Analysis
	ScalarPotential scalarpotential.Analysis

	FockModes     int
	FockStates    int
	TemporalModes int
	SpatialModes  int
	VacuumNeutral bool

	ScalarActiveRealDimension int
	ScalarComplexDimension    int
	ScalarAngularDirections   int
	ScalarRadialDirections    int

	GaugeCompatibleYukawaChannels int
	UpTypeColorChannels           int
	DownTypeColorChannels         int
	LeptonicChannels              int
	ColorAmplificationAvailable   bool

	BilinearScalarCandidateAvailable bool
	CompositeHiggsDirectionPreferred bool

	NativeOneLoopPotentialComputed bool
	NJLGapEquationSolved           bool
	CPTBoundarySelectionComputed   bool
	CondensationScaleDerived       bool
	PhysicalHiggsIdentityDerived   bool

	CriticalComputations []CriticalComputation
	TruthStatement       string
	RecommendedNextGate  string
	RemainingUnknowns    []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		fock, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = err
			return
		}
		y, err := yukawaintertwiner.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		sp, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(fock, y, sp)
	})
	return defaultValue, defaultErr
}

func Build(fock spinor.FockSpace, y yukawaintertwiner.Analysis, sp scalarpotential.Analysis) (Analysis, error) {
	if fock.ModeCount() != 4 {
		return Analysis{}, fmt.Errorf("condensate audit expects 4 Fock modes, got %d", fock.ModeCount())
	}
	if fock.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("condensate audit expects 16 Fock states, got %d", fock.StateCount())
	}
	vac, err := fock.Vacuum()
	if err != nil {
		return Analysis{}, err
	}
	if sp.ActiveRealDimension != 4 {
		return Analysis{}, fmt.Errorf("condensate audit expects 4 active scalar directions, got %d", sp.ActiveRealDimension)
	}
	if y.MinimalChannelCount == 0 {
		return Analysis{}, fmt.Errorf("condensate audit requires gauge-compatible Yukawa channels")
	}

	leptonic := y.NeutrinoChannels + y.ElectronChannels
	colorAmp := y.UpChannels == 3 && y.DownChannels == 3 && fock.SpatialModeCount() == 3
	bilinearAvailable := y.MinimalChannelCount == 8 && sp.ComplexDoubletDimension == 2

	computations := []CriticalComputation{
		{
			ID:          "C1-NATIVE-ONE-LOOP-POTENTIAL",
			Name:        "native one-loop effective potential in Cℓ(1,7) Fock space",
			Status:      "OPEN",
			Requirement: "compute fermion/boson loop signs and scalar mass-parameter flow using the finite Fock/Yukawa operators rather than importing the Standard Model RGE",
		},
		{
			ID:          "C2-NJL-CONDENSATE-GAP",
			Name:        "NJL/top-condensate gap equation from the x∧p gauge sector",
			Status:      "OPEN",
			Requirement: "derive a four-fermion kernel, solve the nonlinear gap equation, and obtain a non-fitted condensate radius/scale",
		},
		{
			ID:          "C3-CPT-BOUNDARY-VACUUM-SELECTION",
			Name:        "CPT boundary condition and vacuum selection",
			Status:      "OPEN",
			Requirement: "test whether a CPT-symmetric boundary selects the vacuum orientation and potential parameters",
		},
	}

	return Analysis{
		Fock:                             fock,
		Yukawa:                           y,
		ScalarPotential:                  sp,
		FockModes:                        fock.ModeCount(),
		FockStates:                       fock.StateCount(),
		TemporalModes:                    fock.TemporalModeCount(),
		SpatialModes:                     fock.SpatialModeCount(),
		VacuumNeutral:                    vac.BMinusL() == 0,
		ScalarActiveRealDimension:        sp.ActiveRealDimension,
		ScalarComplexDimension:           sp.ComplexDoubletDimension,
		ScalarAngularDirections:          sp.ActiveRealDimension - 1,
		ScalarRadialDirections:           1,
		GaugeCompatibleYukawaChannels:    y.MinimalChannelCount,
		UpTypeColorChannels:              y.UpChannels,
		DownTypeColorChannels:            y.DownChannels,
		LeptonicChannels:                 leptonic,
		ColorAmplificationAvailable:      colorAmp,
		BilinearScalarCandidateAvailable: bilinearAvailable,
		CompositeHiggsDirectionPreferred: bilinearAvailable && colorAmp,
		NativeOneLoopPotentialComputed:   false,
		NJLGapEquationSolved:             false,
		CPTBoundarySelectionComputed:     false,
		CondensationScaleDerived:         false,
		PhysicalHiggsIdentityDerived:     false,
		CriticalComputations:             computations,
		TruthStatement:                   truth(bilinearAvailable, colorAmp),
		RecommendedNextGate:              "Gate 52 — Native One-Loop Effective Potential Ledger",
		RemainingUnknowns: []string{
			"U-20A-NATIVE-EFFECTIVE-POTENTIAL: compute the scalar mass-parameter correction using finite Fock/Yukawa operators",
			"U-20B-CONDENSATE-KERNEL: derive the four-fermion/NJL kernel from the x∧p gauge sector instead of inserting it",
			"U-20C-GAP-SCALE: solve the gap equation and derive a non-fitted scalar scale μ or v",
			"U-20D-CPT-VACUUM-ORIENTATION: test whether the boundary condition selects the condensate orientation/sign",
		},
	}, nil
}

func truth(bilinearAvailable, colorAmp bool) string {
	switch {
	case bilinearAvailable && colorAmp:
		return "The finite engine now has the structural ingredients for the composite-Higgs direction: 16-state Fock matter, gauge-compatible left/right scalar channels, a four-real-dimensional scalar/contact doublet, and three-color amplification. The physical Higgs identity and scale still require the native one-loop potential and NJL gap computations."
	case bilinearAvailable:
		return "The finite engine has a scalar bilinear channel, but the color amplification needed for a top-condensate direction is not established."
	default:
		return "The current finite data does not yet support a composite-Higgs/condensate audit."
	}
}

func FormatComputations(values []CriticalComputation) string {
	if len(values) == 0 {
		return "none"
	}
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v.ID + "=" + v.Status
	}
	return out
}

func FormatUnknowns(values []string) string {
	if len(values) == 0 {
		return "none"
	}
	out := ""
	for i, v := range values {
		if i > 0 {
			out += "; "
		}
		out += v
	}
	return out
}
