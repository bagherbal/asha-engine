// Package yukawa formulates the first honest Yukawa/intertwiner selection rule
// on the tensor-factor bridge H_Fock ⊗ H_Φ.
//
// Gate 17 established that charge and scalar response live on different tensor
// factors:
//
//	Q_total = Q_{B-L} ⊗ I_Φ,
//	S_total = I_Fock ⊗ S_Φ.
//
// A neutral finite Yukawa/intertwiner Y cannot be guessed from scalar eigenvalues.
// Its first necessary condition is charge compatibility:
//
//	[Q_total, Y] = 0.
//
// This package computes the size and structure of the linear space of such
// charge-preserving intertwiners. It intentionally does not claim a Standard
// Model Yukawa matrix, fermion masses, chirality, or hypercharge. Instead it
// exposes the next missing object: a scalar charge/hypercharge/chirality bridge
// if one wants electroweak charge-changing couplings.
package yukawa

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/tensor"
)

type ChargeSector struct {
	Charge   float64
	Dim      int
	States   int
	FiberDim int
}

type Analysis struct {
	Tensor tensor.Analysis

	MatterDimension     int
	ScalarDimension     int
	TensorDimension     int
	FullLinearDimension int

	ChargeSectors []ChargeSector

	ChargePreservingDimension int
	ChargeChangingDimension   int
	NeutralSelectionFraction  float64

	OneParticleChargePreservingDimension int
	OneParticleTotalDimension            int
	OneParticleLinearDimension           int

	ParityPreservingDimension int
	ParityFlippingDimension   int
	ParityBalanceResidual     int

	ScalarChargeDimension       int
	ScalarChargeOperatorPresent bool
	NeutralScalarOnly           bool

	ChargeRuleResidual       float64
	SampleNeutralIntertwiner linear.Matrix

	YukawaSelectionRuleFormulated bool
	PhysicalYukawaTextureDerived  bool
	MassMatrixDerived             bool
	RemainingUnknowns             []string
}

var (
	yukawaDefaultOnce  sync.Once
	yukawaDefaultValue Analysis
	yukawaDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	yukawaDefaultOnce.Do(func() {
		yukawaDefaultValue, yukawaDefaultErr = buildYukawaDefaultUncached()
	})
	return yukawaDefaultValue, yukawaDefaultErr
}

func buildYukawaDefaultUncached() (Analysis, error) {
	t, err := tensor.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(t, 1e-10)
}

func Build(t tensor.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if t.TensorDimension <= 0 || t.MatterDimension <= 0 || t.ScalarDimension <= 0 {
		return Analysis{}, fmt.Errorf("invalid tensor dimensions: matter=%d scalar=%d total=%d", t.MatterDimension, t.ScalarDimension, t.TensorDimension)
	}

	sectors := sectorsFromFibers(t.ChargeFibers)
	chargePreserving := 0
	for _, s := range sectors {
		chargePreserving += s.Dim * s.Dim
	}
	full := t.TensorDimension * t.TensorDimension
	chargeChanging := full - chargePreserving
	fraction := 0.0
	if full > 0 {
		fraction = float64(chargePreserving) / float64(full)
	}

	oneParticleDims := oneParticleSectorDimensions(t)
	oneParticleTotal := 0
	oneParticlePreserving := 0
	for _, d := range oneParticleDims {
		oneParticleTotal += d
		oneParticlePreserving += d * d
	}
	oneParticleLinear := oneParticleTotal * oneParticleTotal

	parityPreserving, parityFlipping, parityResidual := paritySelectionDimensions(t)

	sample := buildSampleNeutralIntertwiner(t, eps)
	comm, err := linear.Commutator(t.MatterChargeOperator, sample)
	if err != nil {
		return Analysis{}, err
	}

	return Analysis{
		Tensor:                               t,
		MatterDimension:                      t.MatterDimension,
		ScalarDimension:                      t.ScalarDimension,
		TensorDimension:                      t.TensorDimension,
		FullLinearDimension:                  full,
		ChargeSectors:                        sectors,
		ChargePreservingDimension:            chargePreserving,
		ChargeChangingDimension:              chargeChanging,
		NeutralSelectionFraction:             fraction,
		OneParticleChargePreservingDimension: oneParticlePreserving,
		OneParticleTotalDimension:            oneParticleTotal,
		OneParticleLinearDimension:           oneParticleLinear,
		ParityPreservingDimension:            parityPreserving,
		ParityFlippingDimension:              parityFlipping,
		ParityBalanceResidual:                parityResidual,
		ScalarChargeDimension:                t.ScalarDimension,
		ScalarChargeOperatorPresent:          false,
		NeutralScalarOnly:                    true,
		ChargeRuleResidual:                   comm.FrobeniusNorm(),
		SampleNeutralIntertwiner:             sample,
		YukawaSelectionRuleFormulated:        true,
		PhysicalYukawaTextureDerived:         false,
		MassMatrixDerived:                    false,
		RemainingUnknowns: []string{
			"U-06-HYPERCHARGE: define scalar/fiber charge operator beyond neutral B-L bookkeeping",
			"U-07-YUKAWA: construct a non-arbitrary chirality/flavor intertwiner, not a fitted mass table",
			"U-11-CHIRALITY: identify the correct finite left/right projector on the Fock-contact tensor product",
		},
	}, nil
}

func sectorsFromFibers(fibers []tensor.ChargeFiber) []ChargeSector {
	out := make([]ChargeSector, 0, len(fibers))
	for _, f := range fibers {
		out = append(out, ChargeSector{Charge: f.Charge, Dim: f.TotalDim, States: f.States, FiberDim: f.FiberDim})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Charge < out[j].Charge })
	return out
}

func oneParticleSectorDimensions(t tensor.Analysis) map[string]int {
	// The current Fock convention has one temporal/lepton seed with B-L=-1 and
	// three spatial/color seeds with B-L=1/3. Tensoring with H_Φ gives dimensions
	// 1*dim(H_Φ) and 3*dim(H_Φ).
	return map[string]int{
		"lepton-temporal": t.ScalarDimension,
		"spatial-color":   3 * t.ScalarDimension,
	}
}

func paritySelectionDimensions(t tensor.Analysis) (preserving int, flipping int, residual int) {
	evenDim := 0
	oddDim := 0
	for _, state := range t.Charge.Action.Bridge.Fock.States {
		if state.ExcitationNumber()%2 == 0 {
			evenDim += t.ScalarDimension
		} else {
			oddDim += t.ScalarDimension
		}
	}
	preserving = evenDim*evenDim + oddDim*oddDim
	flipping = 2 * evenDim * oddDim
	residual = int(math.Abs(float64(preserving - flipping)))
	return preserving, flipping, residual
}

func buildSampleNeutralIntertwiner(t tensor.Analysis, eps float64) linear.Matrix {
	// A canonical sample neutral intertwiner is the scalar response itself:
	// S_total=I_Fock⊗S_Φ. It is not the physical Yukawa matrix; it is a test
	// witness for the neutral selection rule [Q_total,Y]=0.
	_ = eps
	return t.ScalarResponse.Clone()
}

func FormatChargeSectors(sectors []ChargeSector) string {
	out := "["
	for i, s := range sectors {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("q=%.6g: dim=%d", s.Charge, s.Dim)
	}
	return out + "]"
}

func FormatUnknowns(unknowns []string) string {
	out := "["
	for i, u := range unknowns {
		if i > 0 {
			out += "; "
		}
		out += u
	}
	return out + "]"
}
