package topdownspectraltriple

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func TopDownFockSpectralTripleBoundaryTraceAmplitudeFirewallTheorem() theorem.Theorem {
	const id = "BRIDGE-TOPDOWN-FOCK-SPECTRAL-TRIPLE-BOUNDARY-TRACE-FIREWALL"
	const name = "top-down Fock spectral triple boundary trace reproduction and amplitude firewall"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build top-down Fock spectral triple ansatz", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "16D Fock-spinor carrier matches Gate-25 left/right count", Passed: a.Hilbert.Dimension == 16 && a.Hilbert.FockStateCount == 16 && a.Hilbert.LeftDimension == 8 && a.Hilbert.RightDimension == 8 && a.Hilbert.IdentifiedWithFock16 && !a.Hilbert.UsesObservedInput && a.Hilbert.RequiresBranchChoice && !a.Hilbert.CanonicalBottomUp, Detail: FormatHilbertAudit(a.Hilbert) + " :: " + FormatBasis(a.BasisLeft, a.BasisRight)},
			{Name: "unit-incidence finite Dirac support is the eight Gate-25 Yukawa channels", Passed: a.Triple.HilbertDimension == 16 && a.Triple.DiracSymmetric && a.Triple.DiracOffDiagonal && a.Triple.YukawaChannelCount == 8 && a.Triple.YukawaChannelSupportComplete && !a.Triple.YukawaAmplitudesDerived, Detail: FormatTripleAudit(a.Triple)},
			{Name: "candidate J and gamma obey matrix identities but do not complete a spectral triple", Passed: a.Triple.RealStructureAvailable && a.Triple.RealStructureInvolutive && a.Triple.RealStructureCommutesWithD && a.Triple.RealStructureAnticommutesGamma && a.Triple.GammaAvailable && a.Triple.GammaInvolutive && a.Triple.GammaTraceZero && a.Triple.GammaAnticommutesWithD && !a.Triple.OrderOneAxiomTested && !a.Triple.PromotableSpectralTriple, Detail: FormatTripleAudit(a.Triple)},
			{Name: "Tr(D_F^4 G^2) unit-incidence trace reproduces embedded boundary normalization", Passed: a.Gauge.BoundaryDiagMatched && a.Gauge.WeakAngleSeedMatched && a.Gauge.TraceD4 == 16 && a.Gauge.NormalizedT1 == 1 && a.Gauge.NormalizedT2 == 1 && a.Gauge.NormalizedT3 == 1, Detail: FormatGaugeAudit(a.Gauge) + " :: sectors " + FormatSectorTraces(a.Gauge.SectorTraces)},
			{Name: "weak-angle seed follows from the reproduced sector ratio", Passed: a.Gauge.WeakAngleSeedMatched, Detail: fmt.Sprintf("sin^2_* = K_SU2/(K_SU2+K_Y) = %.12g; K_Y/K_SU2=%.12g", a.Gauge.WeakAngleSeed, a.Gauge.NormalizedY)},
			{Name: "reproduction is not amplitude-invariant", Passed: a.AmplitudeSensitivity.ArbitraryAmplitudesAllowedByPriorGates && !a.AmplitudeSensitivity.UnitAmplitudesDerivedByGate25 && !a.AmplitudeSensitivity.BoundaryRatioStable && !a.AmplitudeSensitivity.WeakAngleStable, Detail: FormatAmplitudeAudit(a.AmplitudeSensitivity) + " :: " + a.AmplitudeSensitivity.ExampleDeformation},
			{Name: "contact/RG/threshold firewall remains closed", Passed: a.Firewall.ContactModeClassificationBypassedForBoundaryTrace && !a.Firewall.ContactModeClassificationSolved && !a.Firewall.ThresholdCorrectionsDerived && !a.Firewall.RGRunningDerived && !a.Firewall.PhysicalCouplingsDerived && !a.Firewall.MassSpectrumDerived && a.Firewall.GaugeKineticRowsDerived == 0 && a.Firewall.BoundaryRowsReproduced == 4 && a.Firewall.ResidualNullityBefore == 3 && a.Firewall.ResidualNullityAfter == 3, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			"Gate 166 is the first successful top-down reproduction of the embedded boundary ratio from the one-generation Fock/Yukawa representation trace.",
			"The reproduction is not a physical Yukawa, mass, threshold, or RG theorem: arbitrary channel amplitudes are still allowed by prior gates and change Tr(D_F^4 G^2).",
			"The honest next target is an amplitude-rigidity theorem: derive unit incidence or another canonical D_F spectrum from the finite action, rather than inserting it as a top-down ansatz.",
		}}
	}}
}
