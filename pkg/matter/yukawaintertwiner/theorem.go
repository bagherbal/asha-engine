package yukawaintertwiner

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeCompatibleYukawaTheorem() theorem.Theorem {
	const id = "MATTER-GAUGE-COMPATIBLE-YUKAWA-AUDIT"
	const name = "gauge-compatible finite Yukawa/intertwiner channel audit"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Yukawa/intertwiner audit", Passed: false, Detail: err.Error()}},
				}
			}
			eps := 1e-10
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "domain and codomain", Passed: a.LeftDimension == 8 && a.RightDimension == 8, Detail: fmt.Sprintf("left doublet states=%d, right singlet states=%d", a.LeftDimension, a.RightDimension)},
					{Name: "finite scalar branches", Passed: len(a.ScalarBranches) == 2, Detail: FormatScalarBranches(a.ScalarBranches)},
					{Name: "hypercharge balance rule", Passed: a.HyperchargeResidualMax < eps, Detail: fmt.Sprintf("max |Y_R−(Y_L+Y_Φ)| = %.3e", a.HyperchargeResidualMax)},
					{Name: "allowed one-generation channels", Passed: a.UpChannels == 3 && a.DownChannels == 3 && a.NeutrinoChannels == 1 && a.ElectronChannels == 1, Detail: fmt.Sprintf("up=%d, down=%d, neutrino=%d, electron=%d", a.UpChannels, a.DownChannels, a.NeutrinoChannels, a.ElectronChannels)},
					{Name: "color and lepton selection", Passed: a.ColorPreserving && a.LeptonPreserving, Detail: "quark channels preserve color index and lepton channels stay colorless"},
					{Name: "minimal channel list", Passed: a.MinimalChannelCount == 8, Detail: FormatChannels(a.Channels)},
					{Name: "scalar-fiber multiplicity", Passed: a.FiberEntryCount == 16, Detail: FormatSummaries(a.Summaries)},
					{Name: "charge-compatible Yukawa channels derived", Passed: a.ChargeCompatibleYukawaChannelsDerived, Detail: "selection channels are derived from Y_R=Y_L+Y_Φ; no coupling strengths are fitted"},
					{Name: "coupling constants not derived", Passed: !a.GaugeInvariantCouplingConstantsDerived && !a.MassMatrixDerived, Detail: "this gate derives allowed intertwiners only, not masses or numerical Yukawa couplings"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"Gate 25 turns the charge table plus finite scalar doublet branch into explicit allowed one-generation Yukawa channels.",
					"This is a selection-rule theorem, not a mass theorem: coupling strengths, flavor mixing, and generation structure remain bridge problems.",
				},
			}
		},
	}
}
