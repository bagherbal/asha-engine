// Package tensor constructs the first honest tensor-factor bridge between the
// matter/Fock charge sector and the finite Higgs/contact scalar sector.
//
// Earlier gates showed two facts that must both be respected:
//   - the Fock matter sector has a natural B-L charge polarization 1+3;
//   - the finite Higgs/contact sector has a pair-degenerate 2+2 scalar spectrum.
//
// The correct bridge is not to identify those eigenvalue lists. The correct
// bridge is a tensor product:
//
//	H_total = H_Fock ⊗ H_Φ,
//
// where Q_{B-L} acts on H_Fock and the scalar/contact response acts on H_Φ.
// Future Yukawa structure must be an intertwiner between these factors. This
// package establishes that clean separation as a typed theorem gate.
package tensor

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/charge"
)

type ChargeFiber struct {
	Charge   float64
	States   int
	FiberDim int
	TotalDim int
}

type Analysis struct {
	Charge charge.Analysis

	MatterDimension int
	ScalarDimension int
	TensorDimension int

	MatterChargeOperator linear.Matrix // Q_F ⊗ I_Φ lives on the tensor space after lifting.
	ScalarResponse       linear.Matrix // I_F ⊗ S_Φ lives on the tensor space.
	MatterResponse       linear.Matrix // H_F ⊗ I_Φ lives on the tensor space.

	ScalarOneParticleWeights  []float64
	ScalarTrace               float64
	TensorScalarTrace         float64
	ExpectedTensorScalarTrace float64
	TensorScalarTraceResidual float64

	MatterChargeTrace         float64
	TensorChargeTrace         float64
	ExpectedTensorChargeTrace float64
	TensorChargeTraceResidual float64

	ChargeScalarCommutatorNorm float64
	ChargeMatterCommutatorNorm float64
	MatterScalarCommutatorNorm float64

	VacuumScalarFiberDimension int
	VacuumFiberCharge          float64
	ChargeFibers               []ChargeFiber

	DirectEigenvalueIdentificationRejected bool
	YukawaIntertwinerConstructed           bool
	SelectionRuleFormulated                bool
}

var (
	tensorDefaultOnce  sync.Once
	tensorDefaultValue Analysis
	tensorDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	tensorDefaultOnce.Do(func() {
		tensorDefaultValue, tensorDefaultErr = buildTensorDefaultUncached()
	})
	return tensorDefaultValue, tensorDefaultErr
}

func buildTensorDefaultUncached() (Analysis, error) {
	a, err := charge.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-10)
}

func Build(c charge.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	matterDim := c.Action.Operator.Rows()
	if matterDim == 0 || c.Action.Operator.Cols() != matterDim {
		return Analysis{}, fmt.Errorf("matter response operator must be square, got %dx%d", c.Action.Operator.Rows(), c.Action.Operator.Cols())
	}
	scalarWeights := append([]float64(nil), c.Action.OneParticleWeights...)
	if len(scalarWeights) == 0 {
		return Analysis{}, fmt.Errorf("empty scalar/contact one-particle spectrum")
	}
	scalarDim := len(scalarWeights)
	tensorDim := matterDim * scalarDim

	idMatter := linear.Identity(matterDim)
	idScalar := linear.Identity(scalarDim)
	scalarOne := linear.Diagonal(scalarWeights)

	chargeTensor := linear.Kronecker(c.FockChargeOperator, idScalar)
	scalarTensor := linear.Kronecker(idMatter, scalarOne)
	matterTensor := linear.Kronecker(c.Action.Operator, idScalar)

	commQS, err := linear.Commutator(chargeTensor, scalarTensor)
	if err != nil {
		return Analysis{}, err
	}
	commQM, err := linear.Commutator(chargeTensor, matterTensor)
	if err != nil {
		return Analysis{}, err
	}
	commMS, err := linear.Commutator(matterTensor, scalarTensor)
	if err != nil {
		return Analysis{}, err
	}

	scalarTrace := sum(scalarWeights)
	tensorScalarTrace, err := scalarTensor.Trace()
	if err != nil {
		return Analysis{}, err
	}
	matterChargeTrace, err := c.FockChargeOperator.Trace()
	if err != nil {
		return Analysis{}, err
	}
	tensorChargeTrace, err := chargeTensor.Trace()
	if err != nil {
		return Analysis{}, err
	}
	expectedScalarTrace := float64(matterDim) * scalarTrace
	expectedChargeTrace := float64(scalarDim) * matterChargeTrace

	vacuumFiberDim := 0
	vacuumCharge := math.NaN()
	fibers := make([]ChargeFiber, 0)
	for _, state := range c.Action.Bridge.Fock.States {
		q := state.BMinusL()
		if math.Abs(q) < eps {
			q = 0
		}
		fibers = addToChargeFiber(fibers, q, scalarDim, eps)
		if state.IsVacuum() {
			vacuumFiberDim = scalarDim
			vacuumCharge = q
		}
	}
	sort.Slice(fibers, func(i, j int) bool { return fibers[i].Charge < fibers[j].Charge })

	return Analysis{
		Charge:                                 c,
		MatterDimension:                        matterDim,
		ScalarDimension:                        scalarDim,
		TensorDimension:                        tensorDim,
		MatterChargeOperator:                   chargeTensor,
		ScalarResponse:                         scalarTensor,
		MatterResponse:                         matterTensor,
		ScalarOneParticleWeights:               scalarWeights,
		ScalarTrace:                            scalarTrace,
		TensorScalarTrace:                      tensorScalarTrace,
		ExpectedTensorScalarTrace:              expectedScalarTrace,
		TensorScalarTraceResidual:              math.Abs(tensorScalarTrace - expectedScalarTrace),
		MatterChargeTrace:                      matterChargeTrace,
		TensorChargeTrace:                      tensorChargeTrace,
		ExpectedTensorChargeTrace:              expectedChargeTrace,
		TensorChargeTraceResidual:              math.Abs(tensorChargeTrace - expectedChargeTrace),
		ChargeScalarCommutatorNorm:             commQS.FrobeniusNorm(),
		ChargeMatterCommutatorNorm:             commQM.FrobeniusNorm(),
		MatterScalarCommutatorNorm:             commMS.FrobeniusNorm(),
		VacuumScalarFiberDimension:             vacuumFiberDim,
		VacuumFiberCharge:                      vacuumCharge,
		ChargeFibers:                           fibers,
		DirectEigenvalueIdentificationRejected: !c.DirectScalarToColorIsotropyPossible,
		YukawaIntertwinerConstructed:           false,
		SelectionRuleFormulated:                false,
	}, nil
}

func addToChargeFiber(fibers []ChargeFiber, charge float64, scalarDim int, eps float64) []ChargeFiber {
	for i := range fibers {
		if math.Abs(fibers[i].Charge-charge) <= eps {
			fibers[i].States++
			fibers[i].TotalDim = fibers[i].States * scalarDim
			return fibers
		}
	}
	return append(fibers, ChargeFiber{Charge: charge, States: 1, FiberDim: scalarDim, TotalDim: scalarDim})
}

func sum(values []float64) float64 {
	out := 0.0
	for _, v := range values {
		out += v
	}
	return out
}

func FormatFibers(fibers []ChargeFiber) string {
	out := "["
	for i, f := range fibers {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("q=%.6g: states=%d, scalar-fiber=%d, total=%d", f.Charge, f.States, f.FiberDim, f.TotalDim)
	}
	return out + "]"
}
