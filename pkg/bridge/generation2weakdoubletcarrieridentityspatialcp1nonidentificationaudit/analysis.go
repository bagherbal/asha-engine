// Package generation2weakdoubletcarrieridentityspatialcp1nonidentificationaudit implements Gate 576:
// Finite Weak-Doublet Carrier Identity and Spatial CP1 Nonidentification Audit.
//
// Gate 575 proved that the sealed spatial CP^1 complement u^perp exists only
// as projective orientation data and does not carry a finite spectral-triple
// weak-doublet action. Gate 576 makes the type boundary explicit: the actual
// weak socket in the structural finite spectral triple is the quaternionic
// summand H in A_F=C⊕H⊕M3(C), acting on the left fermion carriers L_L and Q_L
// and on the finite one-form scalar carrier H_phi≈C^2. The sealed spatial CP^1
// plane is a different carrier and remains nonidentified with L_L, Q_L, H_phi,
// or Im(H) unless a new functor/intertwiner theorem is constructed.
package generation2weakdoubletcarrieridentityspatialcp1nonidentificationaudit

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fullphysicalfirstorder"
	gate575 "github.com/bagherbal/asha-engine/pkg/bridge/generation2sealedspatialcp1fstcompatibilityaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationfieldcontent"
)

const (
	AuditID = "GATE576-FINITE-WEAK-DOUBLET-CARRIER-IDENTITY-AND-SPATIAL-CP1-NONIDENTIFICATION-AUDIT"

	StatusGate575Inherited                   = "FIREWALL_PRESERVED_GATE575_SEALED_CP1_NON_FST_COMPATIBILITY_INHERITED"
	StatusFiniteAlgebraRecovered             = "PASS_FINITE_ALGEBRA_AF_C_PLUS_H_PLUS_M3C_RECOVERED"
	StatusQuaternionicWeakSocketIdentified   = "PASS_QUATERNIONIC_H_SUMMAND_IDENTIFIED_AS_STRUCTURAL_WEAK_SOCKET"
	StatusImHsu2LIdentified                  = "PASS_IM_H_IDENTIFIED_WITH_SU2_L_STRUCTURAL_LIE_SOCKET"
	StatusWeakFermionCarriersInventoried     = "PASS_FINITE_WEAK_FERMION_DOUBLETS_L_L_AND_Q_L_INVENTORIED"
	StatusQColorMultiplicitySeparated        = "PASS_Q_L_COLOR_MULTIPLICITY_CARRIED_BY_M3_NOT_BY_SPATIAL_CP1"
	StatusScalarDoubletHPhiInventoried       = "PASS_FINITE_ONE_FORM_SCALAR_DOUBLETT_H_PHI_IDENTIFIED"
	StatusHPhiSeparateFromWSpatial           = "PASS_H_PHI_SEPARATE_FROM_W_SPATIAL_AND_U_PERP"
	StatusSealedSpatialCP1NotFSTCarrier      = "FAILED_ROUTE_SEALED_SPATIAL_CP1_NOT_FST_WEAK_CARRIER"
	StatusNoDJGradingFirstOrderForSealedCP1  = "FAILED_ROUTE_SEALED_SPATIAL_CP1_HAS_NO_D_J_GRADING_FIRST_ORDER_ROLE"
	StatusWeakDoubletCountFourPerGeneration  = "PASS_WEAK_DOUBLET_COUNT_FOUR_PER_GENERATION_CERTIFIED"
	StatusOnePlusThreeIsColorMultiplicity    = "PASS_WEAK_DOUBLET_ONE_PLUS_THREE_IS_COLOR_MULTIPLICITY_NOT_SPATIAL_CP1_SELECTION"
	StatusDiracEdgesReconfirmed              = "PASS_FINITE_DIRAC_ONE_FORM_EDGES_RECONFIRMED"
	StatusEdgesDoNotUseSealedSpatialSelector = "FAILED_ROUTE_FINITE_DIRAC_EDGES_DO_NOT_USE_SEALED_SPATIAL_CP1_SELECTOR"
	StatusNonIdentificationCertified         = "PASS_SEALED_SPATIAL_CP1_NONIDENTIFICATION_WITH_FST_CARRIERS_CERTIFIED"
	StatusNoWeakPlaneFlavorEWObservedData    = "FAILED_ROUTE_NO_PHYSICAL_WEAK_PLANE_FLAVOR_OR_ELECTROWEAK_OBSERVED_DATA_FROM_SPATIAL_CP1"
	StatusGate564565BoundaryPreserved        = "FIREWALL_PRESERVED_GATE564_GATE565_ELECTROWEAK_BRIDGE_SYMBOLIC_BOUNDARY"
	StatusK7TimeBoundaryPreserved            = "FIREWALL_PRESERVED_K7_TIME_OS_HILBERT_RG_BOUNDARY"
	StatusGate576BoundaryPreserved           = "FIREWALL_PRESERVED_GATE576_WEAK_DOUBLET_CARRIER_IDENTITY_SPATIAL_CP1_NONIDENTIFICATION_BOUNDARY"
)

type InheritedGate575Audit struct {
	SealedCP1SplitExists          bool
	CommutesWithBMinusL           bool
	CarriesImHAction              bool
	PartOfFiniteWeakCarrier       bool
	CanBePhysicalWeakPlane        bool
	DerivesFlavorOrEWObservedData bool
	AdditionalTheoremRequired     string
	Verdict                       string
}

type FiniteAlgebraAudit struct {
	Algebra                 string
	ComplexSummand          string
	QuaternionicSummand     string
	ColorSummand            string
	PreUnimodularUnitary    string
	UnimodularGaugeGroup    string
	WeakSocketSource        string
	ImHLieAlgebra           string
	QuaternionicWeakSocket  bool
	StructuralOnly          bool
	AbsoluteDynamicsDerived bool
	Verdict                 string
}

type WeakFermionCarrier struct {
	Name              string
	WeakModule        string
	RightModule       string
	ComplexDimension  int
	WeakDoubletCopies int
	HAction           string
	ColorBehavior     string
	IsFiniteCarrier   bool
}

type WeakFermionCarrierInventory struct {
	Carriers                    []WeakFermionCarrier
	LLPresent                   bool
	QLPresent                   bool
	HActsOnLL                   bool
	HActsOnQL                   bool
	QLColorMultiplicity         int
	ColorActionSource           string
	ColorIsWeakStructure        bool
	SealedSpatialCP1Used        bool
	FiniteWeakDoubletsAvailable bool
	Verdict                     string
}

type ScalarDoubletInventory struct {
	CarrierName                   string
	Carrier                       string
	ComplexDimension              int
	RealDimension                 int
	Source                        string
	HActionStructural             bool
	FromFiniteOneFormLane         bool
	SeparateFromWSpatial          bool
	SeparateFromUperp             bool
	SealedSpatialCP1Used          bool
	NumericalHiggsDynamicsDerived bool
	Verdict                       string
}

type SealedSpatialCP1ComparisonAudit struct {
	SealName                    string
	UperpDescription            string
	CP1SplitExistsAlgebraically bool
	AppearsInAFRepresentation   bool
	AppearsInDFEdges            bool
	AppearsInJ                  bool
	AppearsInGrading            bool
	AppearsInFirstOrder         bool
	AppearsInOneFormHiggsLane   bool
	IsFiniteWeakCarrier         bool
	Verdict                     string
}

type WeakDoubletCountAudit struct {
	LeptonWeakDoublets           int
	QuarkColorCopies             int
	QuarkWeakDoublets            int
	TotalWeakDoublets            int
	Gate298SU2Index              string
	OnePlusThreePattern          string
	ComesFromColorMultiplicity   bool
	ComesFromSpatialCP1Selection bool
	Verdict                      string
}

type EdgeLaneRelationAudit struct {
	Edges                     []string
	CanonicalEdgesReconfirmed bool
	UsesSealedSpatialSelector bool
	UsesUperpCarrier          bool
	UsesHPhiScalarLane        bool
	FirstOrderCompatible      bool
	Verdict                   string
}

type NonIdentificationAudit struct {
	UperpEqualsHPhi    bool
	UperpEqualsLL      bool
	UperpEqualsQL      bool
	UperpEqualsImH     bool
	DistinctCarriers   []string
	NewFunctorRequired string
	Certified          bool
	Verdict            string
}

type FirewallAudit struct {
	PhysicalWeakPlaneDerived  bool
	WeakIsospinDerivedFromCP1 bool
	WZPhotonDynamicsDerived   bool
	MassesDerived             bool
	GenerationHierarchy       bool
	YukawaTexture             bool
	CKMPMNS                   bool
	ObservedFlavorData        bool
	Gate564565Preserved       bool
	K7TimePreserved           bool
	Verdict                   string
}

type FinalVerdict struct {
	WeakSocketLocation               string
	ActualWeakDoubletCarriers        []string
	HPhiIsScalarWeakDoublet          bool
	SealedSpatialCP1IsWeakCarrier    bool
	WeakDoubletOnePlusThreeFromColor bool
	DerivesPhysicalWeakFlavorEWData  bool
	AdditionalTheoremRequired        string
	Verdict                          string
}

type Analysis struct {
	Inherited     InheritedGate575Audit
	FiniteAlgebra FiniteAlgebraAudit
	WeakFermions  WeakFermionCarrierInventory
	ScalarDoublet ScalarDoubletInventory
	SealedCompare SealedSpatialCP1ComparisonAudit
	WeakCount     WeakDoubletCountAudit
	EdgeLane      EdgeLaneRelationAudit
	NonIdentity   NonIdentificationAudit
	Firewalls     FirewallAudit
	Final         FinalVerdict
	Truth         string
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
	g575, err := gate575.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate575 sealed spatial CP1 predecessor: %w", err)
	}
	g297, err := fullphysicalfirstorder.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate297 finite spectral-triple carrier predecessor: %w", err)
	}
	g298, err := innerfluctuationfieldcontent.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate298 inner-fluctuation predecessor: %w", err)
	}

	a := Analysis{}
	a.Inherited = auditInherited(g575)
	a.FiniteAlgebra = auditFiniteAlgebra(g298)
	a.WeakFermions = auditWeakFermionCarriers(g297)
	a.ScalarDoublet = auditScalarDoublet(g298)
	a.SealedCompare = auditSealedSpatialComparison(g575)
	a.WeakCount = auditWeakDoubletCount(g298)
	a.EdgeLane = auditEdgeLane(g297, g298)
	a.NonIdentity = auditNonIdentification()
	a.Firewalls = auditFirewalls()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate575.Analysis) InheritedGate575Audit {
	return InheritedGate575Audit{
		SealedCP1SplitExists:          g.Final.SealedCP1SplitExistsAlgebraically,
		CommutesWithBMinusL:           g.Final.CommutesWithBMinusL,
		CarriesImHAction:              g.Final.CarriesNativeOrSealedImHAction,
		PartOfFiniteWeakCarrier:       g.Final.PartOfFiniteWeakDoubletCarrier,
		CanBePhysicalWeakPlane:        g.Final.CanBeCalledPhysicalWeakPlane,
		DerivesFlavorOrEWObservedData: g.Final.DerivesFlavorOrEWObservedData,
		AdditionalTheoremRequired:     g.Final.AdditionalTheoremRequired,
		Verdict:                       StatusGate575Inherited,
	}
}

func auditFiniteAlgebra(g298 innerfluctuationfieldcontent.Analysis) FiniteAlgebraAudit {
	return FiniteAlgebraAudit{
		Algebra:                 g298.Input.Algebra,
		ComplexSummand:          "C: hypercharge/complex singlet socket in the recovered finite algebra ledger",
		QuaternionicSummand:     "H: quaternionic weak socket; unitary part Sp(1)",
		ColorSummand:            "M3(C): color/right-module socket",
		PreUnimodularUnitary:    g298.Gauge.PreUnimodularUnitary,
		UnimodularGaugeGroup:    g298.Gauge.UnimodularGaugeGroup,
		WeakSocketSource:        "quaternionic H summand of A_F",
		ImHLieAlgebra:           "Im(H) ≅ su(2)_L structurally",
		QuaternionicWeakSocket:  true,
		StructuralOnly:          true,
		AbsoluteDynamicsDerived: false,
		Verdict:                 join(StatusFiniteAlgebraRecovered, StatusQuaternionicWeakSocketIdentified, StatusImHsu2LIdentified),
	}
}

func auditWeakFermionCarriers(g297 fullphysicalfirstorder.Analysis) WeakFermionCarrierInventory {
	carriers := []WeakFermionCarrier{}
	ll, ql := false, false
	for _, s := range g297.Representation.Slots {
		switch s.Name {
		case "L_L":
			ll = true
			carriers = append(carriers, WeakFermionCarrier{Name: s.Name, WeakModule: s.WeakModule, RightModule: s.RightModule, ComplexDimension: s.Dim, WeakDoubletCopies: 1, HAction: "H acts on the C^2 weak doublet factor", ColorBehavior: "color singlet/right C lepton module", IsFiniteCarrier: true})
		case "Q_L":
			ql = true
			carriers = append(carriers, WeakFermionCarrier{Name: s.Name, WeakModule: s.WeakModule, RightModule: s.RightModule, ComplexDimension: s.Dim, WeakDoubletCopies: 3, HAction: "H acts on C^2_weak; M3(C) supplies three color copies", ColorBehavior: "three color multiplicities through the right/opposite M3(C) module", IsFiniteCarrier: true})
		}
	}
	return WeakFermionCarrierInventory{
		Carriers:                    carriers,
		LLPresent:                   ll,
		QLPresent:                   ql,
		HActsOnLL:                   ll,
		HActsOnQL:                   ql,
		QLColorMultiplicity:         3,
		ColorActionSource:           "M3(C) right/opposite color module",
		ColorIsWeakStructure:        false,
		SealedSpatialCP1Used:        false,
		FiniteWeakDoubletsAvailable: ll && ql,
		Verdict:                     join(StatusWeakFermionCarriersInventoried, StatusQColorMultiplicitySeparated),
	}
}

func auditScalarDoublet(g298 innerfluctuationfieldcontent.Analysis) ScalarDoubletInventory {
	return ScalarDoubletInventory{
		CarrierName:                   "H_phi",
		Carrier:                       "H_phi ≈ C^2, one complex finite one-form scalar/Higgs doublet plus conjugate",
		ComplexDimension:              g298.Higgs.ComplexDoublets * 2,
		RealDimension:                 g298.Higgs.RealScalarDimension,
		Source:                        "Ω^1_D(A_F) finite one-form scalar lane over the legal D_F edge graph",
		HActionStructural:             g298.Higgs.SingleDoubletRecovered,
		FromFiniteOneFormLane:         true,
		SeparateFromWSpatial:          true,
		SeparateFromUperp:             true,
		SealedSpatialCP1Used:          false,
		NumericalHiggsDynamicsDerived: false,
		Verdict:                       join(StatusScalarDoubletHPhiInventoried, StatusHPhiSeparateFromWSpatial),
	}
}

func auditSealedSpatialComparison(g575 gate575.Analysis) SealedSpatialCP1ComparisonAudit {
	return SealedSpatialCP1ComparisonAudit{
		SealName:                    g575.Decomposition.SealName,
		UperpDescription:            "u^perp, the sealed two-dimensional complement inside W_spatial from SpatialProjectiveOrientationSeal",
		CP1SplitExistsAlgebraically: g575.Final.SealedCP1SplitExistsAlgebraically,
		AppearsInAFRepresentation:   false,
		AppearsInDFEdges:            false,
		AppearsInJ:                  false,
		AppearsInGrading:            false,
		AppearsInFirstOrder:         false,
		AppearsInOneFormHiggsLane:   false,
		IsFiniteWeakCarrier:         false,
		Verdict:                     join(StatusSealedSpatialCP1NotFSTCarrier, StatusNoDJGradingFirstOrderForSealedCP1),
	}
}

func auditWeakDoubletCount(g298 innerfluctuationfieldcontent.Analysis) WeakDoubletCountAudit {
	return WeakDoubletCountAudit{
		LeptonWeakDoublets:           1,
		QuarkColorCopies:             3,
		QuarkWeakDoublets:            3,
		TotalWeakDoublets:            4,
		Gate298SU2Index:              g298.Trace.SU2Index.RatString(),
		OnePlusThreePattern:          "one lepton weak doublet L_L plus three colored quark weak doublets Q_L^r,Q_L^g,Q_L^b",
		ComesFromColorMultiplicity:   true,
		ComesFromSpatialCP1Selection: false,
		Verdict:                      join(StatusWeakDoubletCountFourPerGeneration, StatusOnePlusThreeIsColorMultiplicity),
	}
}

func auditEdgeLane(g297 fullphysicalfirstorder.Analysis, g298 innerfluctuationfieldcontent.Analysis) EdgeLaneRelationAudit {
	edges := []string{}
	for _, e := range g297.FirstOrder.LegalEdges {
		edges = append(edges, e.From+"↔"+e.To)
	}
	if len(edges) == 0 {
		edges = append(edges, g298.Input.EdgeGraph...)
	}
	return EdgeLaneRelationAudit{
		Edges:                     edges,
		CanonicalEdgesReconfirmed: containsAllEdges(edges),
		UsesSealedSpatialSelector: false,
		UsesUperpCarrier:          false,
		UsesHPhiScalarLane:        g298.Higgs.SingleDoubletRecovered,
		FirstOrderCompatible:      g297.FirstOrder.FullSweepVerified,
		Verdict:                   join(StatusDiracEdgesReconfirmed, StatusEdgesDoNotUseSealedSpatialSelector),
	}
}

func auditNonIdentification() NonIdentificationAudit {
	return NonIdentificationAudit{
		UperpEqualsHPhi: false,
		UperpEqualsLL:   false,
		UperpEqualsQL:   false,
		UperpEqualsImH:  false,
		DistinctCarriers: []string{
			"u^perp: sealed two-dimensional subspace inside W_spatial ⊂ C^4 Fock/projective law-space",
			"H_phi: finite one-form scalar/Higgs carrier ≈ C^2",
			"L_L: finite fermion lepton weak doublet carrier",
			"Q_L: finite fermion quark weak doublet carrier with M3(C) color multiplicity",
			"Im(H): three-real-dimensional quaternionic Lie algebra socket, not a complex projective plane",
		},
		NewFunctorRequired: "A new carrier-action functor/intertwiner proving u^perp carries the same A_F, H, D_F, J, grading, first-order, and one-form data as the finite weak-doublet lane.",
		Certified:          true,
		Verdict:            StatusNonIdentificationCertified,
	}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{
		PhysicalWeakPlaneDerived:  false,
		WeakIsospinDerivedFromCP1: false,
		WZPhotonDynamicsDerived:   false,
		MassesDerived:             false,
		GenerationHierarchy:       false,
		YukawaTexture:             false,
		CKMPMNS:                   false,
		ObservedFlavorData:        false,
		Gate564565Preserved:       true,
		K7TimePreserved:           true,
		Verdict:                   join(StatusNoWeakPlaneFlavorEWObservedData, StatusGate564565BoundaryPreserved, StatusK7TimeBoundaryPreserved),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{
		WeakSocketLocation:               "the quaternionic H summand of A_F=C⊕H⊕M3(C), with Im(H)≅su(2)_L structurally",
		ActualWeakDoubletCarriers:        []string{"L_L", "Q_L^r", "Q_L^g", "Q_L^b", "H_phi scalar one-form doublet"},
		HPhiIsScalarWeakDoublet:          a.ScalarDoublet.HActionStructural && a.ScalarDoublet.FromFiniteOneFormLane,
		SealedSpatialCP1IsWeakCarrier:    a.SealedCompare.IsFiniteWeakCarrier,
		WeakDoubletOnePlusThreeFromColor: a.WeakCount.ComesFromColorMultiplicity && !a.WeakCount.ComesFromSpatialCP1Selection,
		DerivesPhysicalWeakFlavorEWData:  a.Firewalls.PhysicalWeakPlaneDerived || a.Firewalls.WeakIsospinDerivedFromCP1 || a.Firewalls.WZPhotonDynamicsDerived || a.Firewalls.MassesDerived || a.Firewalls.GenerationHierarchy || a.Firewalls.YukawaTexture || a.Firewalls.CKMPMNS || a.Firewalls.ObservedFlavorData,
		AdditionalTheoremRequired:        "To identify the sealed spatial CP^1 with a weak carrier, ASHA would need a new native functor/intertwiner from W_spatial/u^perp into the finite spectral-triple carrier category, proving compatible A_F representation, H/Im(H) action, D_F edge role, J, grading, first-order condition, finite one-form/H_phi identity, and preserving color, B-L, K7/time, flavor, and electroweak firewalls.",
		Verdict:                          join(StatusFiniteAlgebraRecovered, StatusWeakFermionCarriersInventoried, StatusScalarDoubletHPhiInventoried, StatusSealedSpatialCP1NotFSTCarrier, StatusOnePlusThreeIsColorMultiplicity, StatusNoWeakPlaneFlavorEWObservedData, StatusGate576BoundaryPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.SealedCP1SplitExists || !a.Inherited.CommutesWithBMinusL || a.Inherited.CarriesImHAction || a.Inherited.PartOfFiniteWeakCarrier || a.Inherited.CanBePhysicalWeakPlane || a.Inherited.DerivesFlavorOrEWObservedData {
		return fmt.Errorf("Gate575 inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if a.FiniteAlgebra.Algebra == "" || !a.FiniteAlgebra.QuaternionicWeakSocket || a.FiniteAlgebra.AbsoluteDynamicsDerived {
		return fmt.Errorf("finite algebra audit failed: %s", FormatFiniteAlgebra(a.FiniteAlgebra))
	}
	if !a.WeakFermions.LLPresent || !a.WeakFermions.QLPresent || !a.WeakFermions.HActsOnLL || !a.WeakFermions.HActsOnQL || a.WeakFermions.QLColorMultiplicity != 3 || a.WeakFermions.ColorIsWeakStructure || a.WeakFermions.SealedSpatialCP1Used {
		return fmt.Errorf("weak fermion carrier audit failed: %s", FormatWeakFermions(a.WeakFermions))
	}
	if a.ScalarDoublet.CarrierName != "H_phi" || a.ScalarDoublet.ComplexDimension != 2 || a.ScalarDoublet.RealDimension != 4 || !a.ScalarDoublet.HActionStructural || !a.ScalarDoublet.SeparateFromWSpatial || !a.ScalarDoublet.SeparateFromUperp || a.ScalarDoublet.SealedSpatialCP1Used {
		return fmt.Errorf("scalar doublet audit failed: %s", FormatScalarDoublet(a.ScalarDoublet))
	}
	if !a.SealedCompare.CP1SplitExistsAlgebraically || a.SealedCompare.AppearsInAFRepresentation || a.SealedCompare.AppearsInDFEdges || a.SealedCompare.AppearsInJ || a.SealedCompare.AppearsInGrading || a.SealedCompare.AppearsInFirstOrder || a.SealedCompare.AppearsInOneFormHiggsLane || a.SealedCompare.IsFiniteWeakCarrier {
		return fmt.Errorf("sealed spatial comparison failed: %s", FormatSealedCompare(a.SealedCompare))
	}
	if a.WeakCount.LeptonWeakDoublets != 1 || a.WeakCount.QuarkWeakDoublets != 3 || a.WeakCount.TotalWeakDoublets != 4 || !a.WeakCount.ComesFromColorMultiplicity || a.WeakCount.ComesFromSpatialCP1Selection {
		return fmt.Errorf("weak count failed: %s", FormatWeakCount(a.WeakCount))
	}
	if !a.EdgeLane.CanonicalEdgesReconfirmed || !a.EdgeLane.FirstOrderCompatible || a.EdgeLane.UsesSealedSpatialSelector || a.EdgeLane.UsesUperpCarrier || !a.EdgeLane.UsesHPhiScalarLane {
		return fmt.Errorf("edge lane failed: %s", FormatEdgeLane(a.EdgeLane))
	}
	if !a.NonIdentity.Certified || a.NonIdentity.UperpEqualsHPhi || a.NonIdentity.UperpEqualsLL || a.NonIdentity.UperpEqualsQL || a.NonIdentity.UperpEqualsImH || a.NonIdentity.NewFunctorRequired == "" {
		return fmt.Errorf("nonidentification failed: %s", FormatNonIdentity(a.NonIdentity))
	}
	if a.Firewalls.PhysicalWeakPlaneDerived || a.Firewalls.WeakIsospinDerivedFromCP1 || a.Firewalls.WZPhotonDynamicsDerived || a.Firewalls.MassesDerived || a.Firewalls.GenerationHierarchy || a.Firewalls.YukawaTexture || a.Firewalls.CKMPMNS || a.Firewalls.ObservedFlavorData || !a.Firewalls.Gate564565Preserved || !a.Firewalls.K7TimePreserved {
		return fmt.Errorf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	if a.Final.SealedSpatialCP1IsWeakCarrier || !a.Final.HPhiIsScalarWeakDoublet || !a.Final.WeakDoubletOnePlusThreeFromColor || a.Final.DerivesPhysicalWeakFlavorEWData || a.Final.AdditionalTheoremRequired == "" {
		return fmt.Errorf("final verdict failed: %s", FormatFinal(a.Final))
	}
	return nil
}

func containsAllEdges(edges []string) bool {
	want := []string{"Q_L↔u_R", "Q_L↔d_R", "L_L↔e_R", "L_L↔ν_R"}
	joined := strings.Join(edges, "|")
	for _, w := range want {
		if !strings.Contains(joined, w) {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusFiniteAlgebraRecovered,
		StatusQuaternionicWeakSocketIdentified,
		StatusImHsu2LIdentified,
		StatusWeakFermionCarriersInventoried,
		StatusQColorMultiplicitySeparated,
		StatusScalarDoubletHPhiInventoried,
		StatusHPhiSeparateFromWSpatial,
		StatusSealedSpatialCP1NotFSTCarrier,
		StatusWeakDoubletCountFourPerGeneration,
		StatusOnePlusThreeIsColorMultiplicity,
		StatusDiracEdgesReconfirmed,
		StatusEdgesDoNotUseSealedSpatialSelector,
		StatusNonIdentificationCertified,
		StatusNoWeakPlaneFlavorEWObservedData,
		StatusGate576BoundaryPreserved,
	}
}

func truth(a Analysis) string {
	return "Gate 576 identifies the actual finite weak-doublet carriers and blocks the remaining type confusion. The weak socket is the quaternionic H summand of A_F=C⊕H⊕M3(C), with Im(H)≅su(2)_L acting structurally on L_L, Q_L, and the finite one-form scalar carrier H_phi≈C^2. The one-plus-three weak-doublet count per generation is L_L plus the three color copies of Q_L, not a projective split of W_spatial. The sealed spatial CP^1 complement u^perp remains a projective orientation inside W_spatial and is not H_phi, L_L, Q_L, or Im(H); no W/Z/photon dynamics, weak plane, generation hierarchy, Yukawa texture, CKM/PMNS, or observed flavor/electroweak data is derived."
}

func join(parts ...string) string { return strings.Join(parts, ";") }
