package spectralgraphtracenormalization

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func SpectralGraphTraceNodeToEdgeKineticNormalizationSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-SPECTRAL-GRAPH-TRACE-NODE-TO-EDGE-KINETIC-NORMALIZATION-SIEVE"
	const name = "Spectral Graph Trace / Node-to-Edge Kinetic Normalization Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build spectral graph trace normalization audit", Passed: false, Detail: err.Error()}}}
		}
		c := a.Calculation
		var contact, edge, unit HiggsLane
		for _, l := range c.Lanes {
			switch l.Name {
			case "contact-node denominator":
				contact = l
			case "edge-trace denominator":
				edge = l
			case "unit sharp-cutoff denominator":
				unit = l
			}
		}
		checks := []theorem.Check{
			{Name: "node and edge trace domains are separated", Passed: len(c.Domains) == 2 && math.Abs(c.Bridge.NodeCount-7) < 1e-12 && math.Abs(c.Bridge.EdgeCount-10) < 1e-12, Detail: "Contact nodes and J-doubled D_F edge slots are distinct finite trace domains."},
			{Name: "Higgs kinetic term is structurally edge-supported", Passed: c.Kinetic.UsesDFCommutator && c.Kinetic.MandatesEdgeSupport && !c.Kinetic.MandatesEdgeDenominator, Detail: c.Kinetic.Verdict},
			{Name: "10/7 node-to-edge bridge is computed", Passed: c.Bridge.CombinatorialNative && math.Abs(c.Bridge.NodeToEdgeRatio-10.0/7.0) < 1e-12 && !c.Bridge.CCMNormalizationNative, Detail: c.Bridge.Verdict},
			{Name: "contact-node denominator overpredicts Higgs", Passed: contact.MassPfaffianGeV > 145 && contact.Native && !contact.Sealed, Detail: contact.Verdict},
			{Name: "edge-trace denominator near-closes Higgs", Passed: math.Abs(edge.Denominator-10) < 1e-12 && math.Abs(edge.MassPfaffianGeV-HiggsTargetGeV) < 0.3 && !edge.Sealed, Detail: edge.Verdict},
			{Name: "unit sharp cutoff without finite normalization overpredicts", Passed: unit.MassPfaffianGeV > 390 && unit.Native && !unit.Sealed, Detail: unit.Verdict},
			{Name: "raw a/e trace recomputation remains missing", Passed: strings.Contains(strings.Join(c.Closure.OpenObstructions, "\n"), "raw a/e") && strings.Contains(StatusLine(c), StatusTensionTenOverSevenRequiresRawTraceAudit), Detail: c.Closure.Conclusion},
			{Name: "Higgs mass geometric sealing is not claimed", Passed: !c.EdgeTraceDerived && !c.TenOverSevenDerived && !c.HiggsMassSealed && strings.Contains(StatusLine(c), StatusFailedHiggsMassNotGeometricallySealed), Detail: c.Truth},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{c.Truth}}
	}}
}
