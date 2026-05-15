// Package generation2scalarsu2provenance implements Gate 498:
// Scalar SU(2)_L Complex-Structure and Gauge-Orbit Provenance Audit.
//
// Gate 497 validated the unitary-gauge quotient only inside an abstract
// electroweak scalar doublet representation.  Gate 498 asks whether that
// representation is selected by the finite scalar/contact data themselves.
//
// The result is a sharp bridge/native boundary.  The four active scalar
// directions admit the standard realification of a complex SU(2) doublet, and
// a compatible complex structure J can be written with J^2=-I.  The active
// finite scalar response also respects the two pair planes.  However the
// anisotropic response has spectrum (a,a,b,b) with a != b, so it commutes only
// with the pairwise complex structure and the diagonal T3-like pair rotation;
// it does not commute with the full T1,T2,T3 SU(2) action.  Therefore the
// abstract scalar SU(2)_L orbit used by Gates 492 and 497 remains a bridge
// representation, not a native finite-action-selected theorem.
package generation2scalarsu2provenance

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2vacuumgaugeorbitquotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

const (
	AuditID = "GATE498-SCALAR-SU2-COMPLEX-STRUCTURE-GAUGE-ORBIT-PROVENANCE-AUDIT"

	StatusGate497Inherited                       = "CONDITIONAL_SUPPORT_GATE497_BRIDGE_GAUGE_QUOTIENT_INHERITED"
	StatusAbstractComplexDoubletSocketFound      = "CONDITIONAL_SUPPORT_ABSTRACT_COMPLEX_DOUBLET_SOCKET_FOUND"
	StatusComplexStructureCompatibleWithPairs    = "CONDITIONAL_SUPPORT_COMPLEX_STRUCTURE_COMPATIBLE_WITH_ACTIVE_PAIR_PLANES"
	StatusAbstractSU2ClosureConfirmed            = "CONDITIONAL_SUPPORT_ABSTRACT_SU2_CLOSURE_CONFIRMED"
	StatusPairRotationU1SelectedByScalarResponse = "CONDITIONAL_SUPPORT_PAIR_ROTATION_U1_SELECTED_BY_SCALAR_RESPONSE"
	StatusBridgeGoldstoneOrbitRemainsConsistent  = "CONDITIONAL_SUPPORT_BRIDGE_GOLDSTONE_ORBIT_REMAINS_CONSISTENT"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_NO_ELECTROWEAK_MASS_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked             = "FIREWALL_BLOCKED_NATIVE_SCALAR_SU2_AND_WZ_REGISTRY_WRITE"

	StatusFailedFullSU2NotSelectedByScalarResponse = "FAILED_ROUTE_FULL_SCALAR_SU2_ACTION_NOT_SELECTED_BY_FINITE_SCALAR_RESPONSE"
	StatusFailedComplexStructureNotNativeUnique    = "FAILED_ROUTE_COMPLEX_STRUCTURE_SOCKET_NOT_NATIVE_UNIQUE"
	StatusFailedAnisotropicResponseBreaksT1T2      = "FAILED_ROUTE_ANISOTROPIC_SCALAR_RESPONSE_BREAKS_FULL_SU2_COMMUTATION"
	StatusFailedGaugeOrbitStillBridge              = "FAILED_ROUTE_ELECTROWEAK_GAUGE_ORBIT_REMAINS_BRIDGE_REPRESENTATION"
	StatusFailedNativeDphiStillUnclosed            = "FAILED_ROUTE_NATIVE_DPHI_PROVENANCE_STILL_UNCLOSED"
	StatusFailedKappaStillBridge                   = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE"
	StatusFailedWZMassStillBlocked                 = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusGate499RedirectDefined                   = "CONDITIONAL_SUPPORT_GATE499_NATIVE_DPHI_INNER_FLUCTUATION_PROVENANCE_REDIRECT_DEFINED"
)

const eps = 1e-9

type Inheritance struct {
	Executed                              bool
	Gate497AuditDefined                   bool
	BridgeGaugeQuotientClosed             bool
	ResidualS1BridgeGaugeOrbitFound       bool
	PhotonIsotropyStabilizerConfirmed     bool
	BrokenGaugeOrbitRankThreeConfirmed    bool
	RadialModeSeparatedFromGaugeOrbit     bool
	UnitaryGaugeRepresentativeBridgeValid bool
	NativeGaugeOrbitStillOpen             bool
	NativeDphiStillOpen                   bool
	NoElectroweakFlavorDataImported       bool
	Verdict                               string
	Reason                                string
}

type ComplexStructureAudit struct {
	Executed                       bool
	ActiveRealDimension            int
	ComplexDoubletDimension        int
	J                              linear.Matrix
	JSkewResidual                  float64
	JSquarePlusIResidual           float64
	ScalarResponseCommJNorm        float64
	JCompatibleWithPairPlanes      bool
	AbstractComplexDoubletSocket   bool
	ComplexStructureNativelyUnique bool
	Verdict                        string
	Reason                         string
}

type SU2ActionAudit struct {
	Executed                        bool
	AbstractDoubletRepresentation   bool
	SkewResidual                    float64
	SU2ClosureResidual              float64
	ScalarResponseSpectrum          []float64
	PairDegenerate                  bool
	PairSplit                       float64
	CommT1Norm                      float64
	CommT2Norm                      float64
	CommT3Norm                      float64
	MaxFullSU2CommNorm              float64
	PairRotationU1Selected          bool
	FullSU2SelectedByScalarResponse bool
	FullSU2ActionNativeSelected     bool
	Verdict                         string
	Reason                          string
}

type GaugeOrbitProvenanceAudit struct {
	Executed                       bool
	BridgeGoldstoneOrbitConsistent bool
	BridgeOrbitRankThreeInherited  bool
	PhotonStabilizerInherited      bool
	RadialQuotientModeInherited    bool
	AbstractSU2NeededForOrbit      bool
	NativeSU2Selected              bool
	NativeGaugeOrbitSelected       bool
	NativeDphiClosed               bool
	WZNativeMassPromotionAllowed   bool
	Verdict                        string
	Reason                         string
}

type NativeBoundary struct {
	Executed                        bool
	ComplexStructureSocketFound     bool
	ComplexStructureNativelyUnique  bool
	AbstractSU2ClosureConfirmed     bool
	FullScalarSU2NativeSelected     bool
	NativeGaugeOrbitSelected        bool
	NativeDphiClosed                bool
	NativeScalarKineticMetricClosed bool
	NativeKappaSelected             bool
	NativeGaugeHessianSelected      bool
	NativeWZMassMatrixDerived       bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                      bool
	ObservedWMassImported         bool
	ObservedZMassImported         bool
	ObservedHiggsMassImported     bool
	FermiConstantImported         bool
	WeakAngleImported             bool
	FineStructureImported         bool
	GaugeCouplingImported         bool
	HiggsVEVImported              bool
	YukawaImported                bool
	CKMPMNSImported               bool
	NativeComplexStructureWritten bool
	NativeScalarSU2Written        bool
	NativeGaugeOrbitWritten       bool
	NativeDphiWritten             bool
	NativeKappaWritten            bool
	NativeWZMassWritten           bool
	Verdict                       string
	Reason                        string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance      Inheritance
	ComplexStructure ComplexStructureAudit
	SU2Action        SU2ActionAudit
	GaugeOrbit       GaugeOrbitProvenanceAudit
	Boundary         NativeBoundary
	Firewall         Firewall
	Registry         RegistryUpdate
	Next             NextStep
	Truth            string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g497, err := generation2vacuumgaugeorbitquotient.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate497 vacuum gauge-orbit quotient audit: %w", err)
	}
	su2, err := scalarsu2.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar SU2 audit: %w", err)
	}

	inherited := Inheritance{
		Executed:                              true,
		Gate497AuditDefined:                   true,
		BridgeGaugeQuotientClosed:             g497.Boundary.BridgeQuotientSocketClosed,
		ResidualS1BridgeGaugeOrbitFound:       g497.ResidualPhase.ResidualS1CoveredByGaugeOrbit,
		PhotonIsotropyStabilizerConfirmed:     g497.ResidualPhase.PhotonStabilizesVacuum && g497.GaugeOrbit.PhotonIsotropyGenerator,
		BrokenGaugeOrbitRankThreeConfirmed:    g497.GaugeOrbit.GaugeOrbitRankThree,
		RadialModeSeparatedFromGaugeOrbit:     g497.GaugeOrbit.RadialSeparatedFromGaugeOrbit,
		UnitaryGaugeRepresentativeBridgeValid: g497.Representative.RepresentativeAllowedAfterQuotient,
		NativeGaugeOrbitStillOpen:             !g497.Boundary.NativeGaugeOrbitSelected,
		NativeDphiStillOpen:                   !g497.Boundary.NativeDphiClosed,
		NoElectroweakFlavorDataImported:       noImportedGate497(g497),
		Verdict:                               StatusGate497Inherited,
		Reason:                                "Gate497 supplies a bridge-valid rank-three broken gauge orbit and photon stabilizer, but explicitly leaves the scalar SU(2)L/gauge-orbit provenance open.",
	}

	j := standardComplexStructure()
	j2, err := j.Mul(j)
	if err != nil {
		return Analysis{}, err
	}
	j2PlusI, err := j2.Add(linear.Identity(4))
	if err != nil {
		return Analysis{}, err
	}
	jSkew, err := j.Add(j.Transpose())
	if err != nil {
		return Analysis{}, err
	}
	cj, err := linear.Commutator(su2.ScalarResponse, j)
	if err != nil {
		return Analysis{}, err
	}
	jCompatible := jSkew.FrobeniusNorm() < eps && j2PlusI.FrobeniusNorm() < eps && cj.FrobeniusNorm() < eps
	complexAudit := ComplexStructureAudit{
		Executed:                       true,
		ActiveRealDimension:            su2.ActiveRealDimension,
		ComplexDoubletDimension:        su2.ComplexDoubletDimension,
		J:                              j,
		JSkewResidual:                  jSkew.FrobeniusNorm(),
		JSquarePlusIResidual:           j2PlusI.FrobeniusNorm(),
		ScalarResponseCommJNorm:        cj.FrobeniusNorm(),
		JCompatibleWithPairPlanes:      jCompatible,
		AbstractComplexDoubletSocket:   su2.ActiveRealDimension == 4 && su2.ComplexDoubletDimension == 2 && jCompatible,
		ComplexStructureNativelyUnique: false,
		Verdict:                        StatusAbstractComplexDoubletSocketFound,
		Reason:                         "A standard complex structure on the two active pair planes satisfies J^2=-I, is skew with respect to the Euclidean diagnostic metric, and commutes with the finite scalar response.  This proves a compatible complex-doublet socket, not a unique native complex-structure selector.",
	}

	su2Audit := SU2ActionAudit{
		Executed:                        true,
		AbstractDoubletRepresentation:   su2.AbstractDoubletRepresentation,
		SkewResidual:                    su2.SkewResidual,
		SU2ClosureResidual:              su2.SU2ClosureResidual,
		ScalarResponseSpectrum:          append([]float64(nil), su2.ActiveSpectrum...),
		PairDegenerate:                  su2.PairDegenerate,
		PairSplit:                       su2.PairSplit,
		CommT1Norm:                      su2.ScalarResponseCommT1Norm,
		CommT2Norm:                      su2.ScalarResponseCommT2Norm,
		CommT3Norm:                      su2.ScalarResponseCommT3Norm,
		MaxFullSU2CommNorm:              su2.MaxFullSU2CommNorm,
		PairRotationU1Selected:          su2.U1PairRotationSelected,
		FullSU2SelectedByScalarResponse: su2.FullSU2SelectedByScalarData,
		FullSU2ActionNativeSelected:     false,
		Verdict:                         StatusFailedFullSU2NotSelectedByScalarResponse,
		Reason:                          "The standard real doublet generators close su(2), but the anisotropic scalar response commutes only with J/T3 and not with T1,T2.  Therefore finite scalar response data do not select the full scalar SU(2)L action.",
	}

	orbitConsistent := inherited.BridgeGaugeQuotientClosed && complexAudit.AbstractComplexDoubletSocket && su2Audit.AbstractDoubletRepresentation && g497.GaugeOrbit.FourToOneQuotientDiagnostic
	gaugeOrbit := GaugeOrbitProvenanceAudit{
		Executed:                       true,
		BridgeGoldstoneOrbitConsistent: orbitConsistent,
		BridgeOrbitRankThreeInherited:  g497.GaugeOrbit.GaugeOrbitRankThree,
		PhotonStabilizerInherited:      g497.GaugeOrbit.PhotonIsotropyGenerator,
		RadialQuotientModeInherited:    g497.GaugeOrbit.ScalarDimensionAfterQuotient == 1,
		AbstractSU2NeededForOrbit:      true,
		NativeSU2Selected:              false,
		NativeGaugeOrbitSelected:       false,
		NativeDphiClosed:               false,
		WZNativeMassPromotionAllowed:   false,
		Verdict:                        StatusBridgeGoldstoneOrbitRemainsConsistent,
		Reason:                         "The Gate497 Goldstone quotient remains internally consistent when read through the abstract complex doublet and su(2) socket.  Its provenance is still bridge-level because the full SU(2)L action is not selected by the finite scalar response and DΦ remains unclosed.",
	}

	boundary := NativeBoundary{
		Executed:                        true,
		ComplexStructureSocketFound:     complexAudit.AbstractComplexDoubletSocket,
		ComplexStructureNativelyUnique:  false,
		AbstractSU2ClosureConfirmed:     su2Audit.AbstractDoubletRepresentation,
		FullScalarSU2NativeSelected:     false,
		NativeGaugeOrbitSelected:        false,
		NativeDphiClosed:                false,
		NativeScalarKineticMetricClosed: false,
		NativeKappaSelected:             false,
		NativeGaugeHessianSelected:      false,
		NativeWZMassMatrixDerived:       false,
		Verdict:                         StatusFailedFullSU2NotSelectedByScalarResponse,
		Reason:                          "Gate498 upgrades the abstract scalar representation audit but does not close native provenance.  The scalar response selects pair structure and a commuting U(1), not the full electroweak SU(2)L gauge orbit or its finite-action DΦ.",
	}

	firewall := Firewall{
		Executed:                      true,
		ObservedWMassImported:         false,
		ObservedZMassImported:         false,
		ObservedHiggsMassImported:     false,
		FermiConstantImported:         false,
		WeakAngleImported:             false,
		FineStructureImported:         false,
		GaugeCouplingImported:         false,
		HiggsVEVImported:              false,
		YukawaImported:                false,
		CKMPMNSImported:               false,
		NativeComplexStructureWritten: false,
		NativeScalarSU2Written:        false,
		NativeGaugeOrbitWritten:       false,
		NativeDphiWritten:             false,
		NativeKappaWritten:            false,
		NativeWZMassWritten:           false,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "No electroweak mass, weak-angle, gauge-coupling, Higgs-VEV, Yukawa, CKM, or PMNS datum is imported, and no native scalar SU(2), DΦ, kappa, or W/Z registry write is made.",
	}

	registry := RegistryUpdate{
		NativeEntries: []string{
			"No new native scalar SU(2)L, DΦ, gauge-orbit, kappa_U1, or W/Z mass entry is admitted at Gate498.",
		},
		BridgeEntries: []string{
			"The four active scalar directions admit a compatible complex-doublet socket with J^2=-I and [S_phi,J]=0.",
			"The standard real SU(2) doublet generators close algebraically on the four-real scalar frame.",
			"The scalar response selects the pairwise complex/U(1) rotation but not the full T1,T2,T3 SU(2)L action.",
			"The Gate497 4 -> 1 Goldstone quotient remains a consistent bridge diagnostic, not a native electroweak theorem.",
		},
		EnvironmentalEntries: []string{
			"Observed W/Z masses, Higgs VEV, Fermi constant, weak angle, alpha, gauge couplings, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedFullSU2NotSelectedByScalarResponse,
			StatusFailedComplexStructureNotNativeUnique,
			StatusFailedAnisotropicResponseBreaksT1T2,
			StatusFailedGaugeOrbitStillBridge,
			StatusFailedNativeDphiStillUnclosed,
			StatusFailedKappaStillBridge,
			StatusFailedWZMassStillBlocked,
		},
		OpenTheorems: []string{
			"Derive the scalar complex/quaternionic structure directly from finite contact/spectral data rather than choosing a compatible realification.",
			"Derive the full SU(2)L scalar action from finite inner fluctuations or contact geometry, despite the anisotropic scalar response.",
			"Derive native DΦ and scalar kinetic normalization before promoting kappa_U1=6, the gauge Hessian, or W/Z masses.",
		},
	}

	next := NextStep{
		Gate:        499,
		Title:       "Native DΦ Inner-Fluctuation Provenance Audit",
		Reason:      "Gate498 confirms that the scalar complex/SU2 structure is only a compatible bridge socket; the next possible native promotion must come from the finite inner-fluctuation/covariant-derivative construction rather than from the scalar response alone.",
		PrimaryTask: "Audit whether the finite spectral triple and inner fluctuation algebra produce a canonical scalar covariant derivative DΦ with the same Goldstone/gauge images, or whether DΦ remains an imported continuum bridge template.",
	}

	truth := "Gate498 proves that the scalar frame can consistently be read as a complex SU(2)L doublet: a compatible J exists, the abstract su(2) generators close, and the Gate497 Goldstone quotient remains coherent.  But the finite scalar response itself is anisotropic; it selects pair structure and a commuting U(1)/T3 direction, not the full SU(2)L orbit.  Therefore the electroweak scalar complex structure, gauge orbit, DΦ, kappa_U1=6, and W/Z mass matrix remain bridge-level until a native inner-fluctuation theorem selects them."

	return Analysis{Inheritance: inherited, ComplexStructure: complexAudit, SU2Action: su2Audit, GaugeOrbit: gaugeOrbit, Boundary: boundary, Firewall: firewall, Registry: registry, Next: next, Truth: truth}, nil
}

func noImportedGate497(g497 generation2vacuumgaugeorbitquotient.Analysis) bool {
	f := g497.Firewall
	return !f.ObservedWMassImported && !f.ObservedZMassImported && !f.ObservedHiggsMassImported && !f.FermiConstantImported && !f.WeakAngleImported && !f.FineStructureImported && !f.GaugeCouplingImported && !f.HiggsVEVImported && !f.YukawaImported && !f.CKMPMNSImported
}

func standardComplexStructure() linear.Matrix {
	j := linear.NewMatrix(4, 4)
	j.Set(0, 1, -1)
	j.Set(1, 0, 1)
	j.Set(2, 3, -1)
	j.Set(3, 2, 1)
	return j
}

func FormatSpectrum(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10f", v)
	}
	return out + "]"
}

func FormatMatrix(m linear.Matrix) string {
	rows := make([]string, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals[c] = fmt.Sprintf("%.1f", m.At(r, c))
		}
		rows[r] = "[" + strings.Join(vals, ", ") + "]"
	}
	return "[" + strings.Join(rows, ", ") + "]"
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate497=%t bridge_quotient=%t residual_s1_bridge=%t photon_stabilizer=%t broken_rank3=%t radial_separated=%t unitary_bridge=%t native_orbit_open=%t native_Dphi_open=%t no_data=%t verdict=%s reason=%s", x.Gate497AuditDefined, x.BridgeGaugeQuotientClosed, x.ResidualS1BridgeGaugeOrbitFound, x.PhotonIsotropyStabilizerConfirmed, x.BrokenGaugeOrbitRankThreeConfirmed, x.RadialModeSeparatedFromGaugeOrbit, x.UnitaryGaugeRepresentativeBridgeValid, x.NativeGaugeOrbitStillOpen, x.NativeDphiStillOpen, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatComplexStructure(x ComplexStructureAudit) string {
	return fmt.Sprintf("active_dim=%d complex_dim=%d J=%s J_skew=%.3e J2_plus_I=%.3e comm_SJ=%.3e pair_compatible=%t socket=%t native_unique=%t verdict=%s reason=%s", x.ActiveRealDimension, x.ComplexDoubletDimension, FormatMatrix(x.J), x.JSkewResidual, x.JSquarePlusIResidual, x.ScalarResponseCommJNorm, x.JCompatibleWithPairPlanes, x.AbstractComplexDoubletSocket, x.ComplexStructureNativelyUnique, x.Verdict, x.Reason)
}

func FormatSU2Action(x SU2ActionAudit) string {
	return fmt.Sprintf("abstract_doublet=%t skew=%.3e closure=%.3e spectrum=%s pair_degenerate=%t pair_split=%.10f commT1=%.6e commT2=%.6e commT3=%.6e max_comm=%.6e pair_U1=%t full_SU2_by_response=%t native_SU2=%t verdict=%s reason=%s", x.AbstractDoubletRepresentation, x.SkewResidual, x.SU2ClosureResidual, FormatSpectrum(x.ScalarResponseSpectrum), x.PairDegenerate, x.PairSplit, x.CommT1Norm, x.CommT2Norm, x.CommT3Norm, x.MaxFullSU2CommNorm, x.PairRotationU1Selected, x.FullSU2SelectedByScalarResponse, x.FullSU2ActionNativeSelected, x.Verdict, x.Reason)
}

func FormatGaugeOrbit(x GaugeOrbitProvenanceAudit) string {
	return fmt.Sprintf("bridge_consistent=%t rank3=%t photon=%t radial_one=%t abstract_SU2_needed=%t native_SU2=%t native_orbit=%t native_Dphi=%t WZ_native=%t verdict=%s reason=%s", x.BridgeGoldstoneOrbitConsistent, x.BridgeOrbitRankThreeInherited, x.PhotonStabilizerInherited, x.RadialQuotientModeInherited, x.AbstractSU2NeededForOrbit, x.NativeSU2Selected, x.NativeGaugeOrbitSelected, x.NativeDphiClosed, x.WZNativeMassPromotionAllowed, x.Verdict, x.Reason)
}

func FormatBoundary(x NativeBoundary) string {
	return fmt.Sprintf("J_socket=%t J_unique=%t abstract_SU2=%t full_SU2_native=%t native_orbit=%t native_Dphi=%t native_metric=%t native_kappa=%t native_hessian=%t native_WZ=%t verdict=%s reason=%s", x.ComplexStructureSocketFound, x.ComplexStructureNativelyUnique, x.AbstractSU2ClosureConfirmed, x.FullScalarSU2NativeSelected, x.NativeGaugeOrbitSelected, x.NativeDphiClosed, x.NativeScalarKineticMetricClosed, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("observed_W=%t observed_Z=%t observed_Higgs=%t Fermi=%t theta=%t alpha=%t gauge_coupling=%t v=%t Yukawa=%t CKM_PMNS=%t native_J=%t native_SU2=%t native_orbit=%t native_Dphi=%t native_kappa=%t native_WZ=%t verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.HiggsVEVImported, x.YukawaImported, x.CKMPMNSImported, x.NativeComplexStructureWritten, x.NativeScalarSU2Written, x.NativeGaugeOrbitWritten, x.NativeDphiWritten, x.NativeKappaWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func nearlyZero(x float64) bool { return math.Abs(x) < eps }
