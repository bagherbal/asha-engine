// Package generation2innerfluctuationdphiprovenance implements Gate 499:
// Native DΦ Inner-Fluctuation Provenance Audit.
//
// Gate 498 proved that the finite scalar/contact response by itself does not
// select the full scalar SU(2)L action: it preserves only the pairwise complex
// U(1) direction.  Gate 499 therefore asks a different provenance question.  Is
// the scalar electroweak representation selected by the completed finite
// spectral-triple / inner-fluctuation field-content ledger rather than by the
// scalar response spectrum alone?
//
// The answer is deliberately split.  The Gate-298 inner-fluctuation theorem
// already recovers exactly one complex weak Higgs doublet from the finite
// one-form module over A_F=C⊕H⊕M3(C), together with the 12 gauge bosons.  This
// closes the structural field-content provenance of the scalar doublet socket
// and explains why a full SU(2)L representation is legitimate even though the
// anisotropic scalar response does not commute with T1,T2.  But it still does
// not derive the finite product-action kinetic projection, scalar normalization,
// vacuum orientation, gauge Hessian/couplings, kappa_U1, or any W/Z mass.
package generation2innerfluctuationdphiprovenance

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarsu2provenance"
	"github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationfieldcontent"
)

const (
	AuditID = "GATE499-INNER-FLUCTUATION-DPHI-PROVENANCE-AUDIT"

	StatusGate498Inherited                  = "CONDITIONAL_SUPPORT_GATE498_SCALAR_SU2_PROVENANCE_OBSTRUCTION_INHERITED"
	StatusInnerFluctuationFieldContent      = "CONDITIONAL_SUPPORT_INNER_FLUCTUATION_FIELD_CONTENT_INHERITED"
	StatusFiniteOneFormHiggsDoublet         = "CONDITIONAL_SUPPORT_FINITE_ONEFORM_HIGGS_DOUBLET_PROVENANCE_CONFIRMED"
	StatusGaugeBosonContentRecovered        = "CONDITIONAL_SUPPORT_GAUGE_BOSON_CONTENT_RECOVERED_FROM_UNITARY_ALGEBRA"
	StatusStructuralDphiSocketFound         = "CONDITIONAL_SUPPORT_STRUCTURAL_DPHI_TRANSFORMATION_SOCKET_FOUND"
	StatusScalarResponseObstructionResolved = "CONDITIONAL_SUPPORT_SCALAR_RESPONSE_SU2_OBSTRUCTION_RECONCILED_AS_POTENTIAL_RESPONSE_NOT_REPRESENTATION_PROVENANCE"
	StatusStructuralRepresentationPromoted  = "CONDITIONAL_SUPPORT_STRUCTURAL_SCALAR_SU2_REPRESENTATION_PROVENANCE_PROMOTED"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_NO_ELECTROWEAK_MASS_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked        = "FIREWALL_BLOCKED_NATIVE_DPHI_ACTION_AND_WZ_REGISTRY_WRITE"

	StatusFailedNativeDphiActionNotDerived      = "FAILED_ROUTE_NATIVE_DPHI_ACTION_AND_KINETIC_PROJECTION_NOT_DERIVED"
	StatusFailedHeatKernelScalarKineticMissing  = "FAILED_ROUTE_HEAT_KERNEL_SCALAR_KINETIC_COEFFICIENT_NOT_DERIVED"
	StatusFailedVacuumOrientationStillBridge    = "FAILED_ROUTE_SCALAR_VACUUM_ORIENTATION_STILL_BRIDGE_LEVEL"
	StatusFailedKappaStillBridge                = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE"
	StatusFailedGaugeHessianCouplingsNotDerived = "FAILED_ROUTE_GAUGE_HESSIAN_AND_COUPLINGS_NOT_DERIVED"
	StatusFailedPhysicalWZMassBlocked           = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusGate500RedirectDefined                = "CONDITIONAL_SUPPORT_GATE500_PRODUCT_SPECTRAL_ACTION_SCALAR_KINETIC_PROJECTION_REDIRECT_DEFINED"
)

type Inheritance struct {
	Executed                           bool
	Gate498AuditDefined                bool
	ComplexDoubletSocketFound          bool
	AbstractSU2ClosureConfirmed        bool
	ScalarResponseSelectsOnlyU1        bool
	FullScalarSU2NotSelectedByResponse bool
	BridgeGoldstoneOrbitConsistent     bool
	NativeDphiStillOpen                bool
	KappaStillBridge                   bool
	NoElectroweakFlavorDataImported    bool
	Verdict                            string
	Reason                             string
}

type InnerFluctuationAudit struct {
	Executed                     bool
	Gate298FieldContentAvailable bool
	Algebra                      string
	NCGOneFormsFormalized        bool
	FluctuatedDiracFormula       string
	GaugeBosonContentRecovered   bool
	GaugeBosonDimension          int
	HiggsDoubletRecovered        bool
	ComplexDoublets              int
	RealScalarDimension          int
	WeakRepresentation           string
	ColorRepresentation          string
	HyperchargeRay               string
	FiniteOneFormEdges           int
	NumericalYukawaFree          bool
	HiggsPotentialNotDerived     bool
	HeatKernelProjectionMissing  bool
	StructuralFieldContent       bool
	Verdict                      string
	Reason                       string
}

type DphiProvenanceAudit struct {
	Executed                                bool
	StructuralGaugeConnectionAvailable      bool
	StructuralScalarOneFormAvailable        bool
	StructuralLeftRightActionAvailable      bool
	StructuralDphiSocketFound               bool
	ScalarSU2RepresentationProvenanceClosed bool
	ProductGeometryKineticProjectionDerived bool
	NativeDphiActionDerived                 bool
	ScalarKineticNormalizationDerived       bool
	GaugeHessianCouplingsDerived            bool
	PhysicalMassMatrixDerived               bool
	Verdict                                 string
	Reason                                  string
}

type ReconciliationAudit struct {
	Executed                              bool
	ScalarResponseAnisotropic             bool
	ScalarResponseBreaksT1T2              bool
	InnerFluctuationSelectsRepresentation bool
	NoContradiction                       bool
	RepresentationVsResponseSeparated     bool
	GoldstoneBridgeOrbitPreserved         bool
	NativeGaugeEatingStillBlocked         bool
	Verdict                               string
	Reason                                string
}

type Boundary struct {
	Executed                                   bool
	StructuralScalarDoubletProvenancePromoted  bool
	StructuralDphiTransformationSocketPromoted bool
	NativeFullDphiActionClosed                 bool
	NativeScalarKineticProjectionClosed        bool
	NativeVacuumOrientationClosed              bool
	NativeKappaSelected                        bool
	NativeGaugeHessianSelected                 bool
	NativeWZMassMatrixDerived                  bool
	Verdict                                    string
	Reason                                     string
}

type Firewall struct {
	Executed                  bool
	ObservedWMassImported     bool
	ObservedZMassImported     bool
	ObservedHiggsMassImported bool
	FermiConstantImported     bool
	WeakAngleImported         bool
	FineStructureImported     bool
	GaugeCouplingImported     bool
	HiggsVEVImported          bool
	YukawaImported            bool
	CKMPMNSImported           bool
	NativeDphiActionWritten   bool
	NativeKineticWritten      bool
	NativeKappaWritten        bool
	NativeWZMassWritten       bool
	Verdict                   string
	Reason                    string
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
	InnerFluctuation InnerFluctuationAudit
	Dphi             DphiProvenanceAudit
	Reconciliation   ReconciliationAudit
	Boundary         Boundary
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
	g498, err := generation2scalarsu2provenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate498 scalar SU2 provenance audit: %w", err)
	}
	g298, err := innerfluctuationfieldcontent.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate298 inner-fluctuation field-content audit: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g498)
	a.InnerFluctuation = buildInnerFluctuation(g298)
	a.Dphi = buildDphi(a.InnerFluctuation)
	a.Reconciliation = buildReconciliation(g498, a.InnerFluctuation, a.Dphi)
	a.Boundary = buildBoundary(a.Dphi)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2scalarsu2provenance.Analysis) Inheritance {
	return Inheritance{
		Executed:                           true,
		Gate498AuditDefined:                true,
		ComplexDoubletSocketFound:          g.ComplexStructure.AbstractComplexDoubletSocket,
		AbstractSU2ClosureConfirmed:        g.SU2Action.AbstractDoubletRepresentation && g.SU2Action.SU2ClosureResidual < 1e-9,
		ScalarResponseSelectsOnlyU1:        g.SU2Action.PairRotationU1Selected,
		FullScalarSU2NotSelectedByResponse: !g.SU2Action.FullSU2SelectedByScalarResponse && !g.SU2Action.FullSU2ActionNativeSelected,
		BridgeGoldstoneOrbitConsistent:     g.GaugeOrbit.BridgeGoldstoneOrbitConsistent,
		NativeDphiStillOpen:                !g.GaugeOrbit.NativeDphiClosed && !g.Boundary.NativeDphiClosed,
		KappaStillBridge:                   !g.Boundary.NativeKappaSelected,
		NoElectroweakFlavorDataImported:    g.Firewall.Executed && !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.WeakAngleImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                            StatusGate498Inherited,
		Reason:                             "Gate498 leaves a clean provenance gap: scalar response alone selects pair U(1), not the full scalar SU(2)L or native DΦ.",
	}
}

func buildInnerFluctuation(g innerfluctuationfieldcontent.Analysis) InnerFluctuationAudit {
	return InnerFluctuationAudit{
		Executed:                     true,
		Gate298FieldContentAvailable: true,
		Algebra:                      g.Input.Algebra,
		NCGOneFormsFormalized:        g.NCG.Formalized,
		FluctuatedDiracFormula:       g.NCG.FluctuatedDirac,
		GaugeBosonContentRecovered:   g.Gauge.GaugeContentRecovered,
		GaugeBosonDimension:          g.Gauge.TotalDimension,
		HiggsDoubletRecovered:        g.Higgs.SingleDoubletRecovered,
		ComplexDoublets:              g.Higgs.ComplexDoublets,
		RealScalarDimension:          g.Higgs.RealScalarDimension,
		WeakRepresentation:           g.Higgs.WeakRepresentation,
		ColorRepresentation:          g.Higgs.ColorRepresentation,
		HyperchargeRay:               g.Higgs.HyperchargeAbs,
		FiniteOneFormEdges:           len(g.Higgs.Edges),
		NumericalYukawaFree:          g.Higgs.NumericalYukawaFree,
		HiggsPotentialNotDerived:     g.Firewalls.DoesNotClaimHiggsPotential,
		HeatKernelProjectionMissing:  g.Firewalls.DoesNotClaimHeatKernel,
		StructuralFieldContent:       g.Summary.FullSMFieldContentStructural,
		Verdict:                      strings.Join([]string{StatusInnerFluctuationFieldContent, StatusGaugeBosonContentRecovered, StatusFiniteOneFormHiggsDoublet}, ";"),
		Reason:                       "The completed finite spectral-triple inner-fluctuation ledger recovers the Standard Model gauge field content and exactly one complex weak Higgs doublet from finite one-forms, while leaving numerical dynamics firewalled.",
	}
}

func buildDphi(in InnerFluctuationAudit) DphiProvenanceAudit {
	structural := in.NCGOneFormsFormalized && in.GaugeBosonContentRecovered && in.HiggsDoubletRecovered && in.ComplexDoublets == 1 && in.RealScalarDimension == 4
	return DphiProvenanceAudit{
		Executed:                                true,
		StructuralGaugeConnectionAvailable:      in.GaugeBosonContentRecovered && in.GaugeBosonDimension == 12,
		StructuralScalarOneFormAvailable:        in.HiggsDoubletRecovered && in.FiniteOneFormEdges == 4,
		StructuralLeftRightActionAvailable:      in.WeakRepresentation != "" && in.HyperchargeRay != "",
		StructuralDphiSocketFound:               structural,
		ScalarSU2RepresentationProvenanceClosed: structural,
		ProductGeometryKineticProjectionDerived: false,
		NativeDphiActionDerived:                 false,
		ScalarKineticNormalizationDerived:       false,
		GaugeHessianCouplingsDerived:            false,
		PhysicalMassMatrixDerived:               false,
		Verdict:                                 strings.Join([]string{StatusStructuralDphiSocketFound, StatusStructuralRepresentationPromoted, StatusFailedNativeDphiActionNotDerived}, ";"),
		Reason:                                  "Inner fluctuations supply the representation-level ingredients for DΦ: gauge connection, finite scalar one-form, and left/right gauge action. They do not yet supply the product spectral-action kinetic projection that would turn this socket into a normalized native electroweak action.",
	}
}

func buildReconciliation(g generation2scalarsu2provenance.Analysis, in InnerFluctuationAudit, d DphiProvenanceAudit) ReconciliationAudit {
	anisotropic := g.SU2Action.PairDegenerate && g.SU2Action.PairSplit > 0
	breaks := g.SU2Action.CommT1Norm > 1e-8 && g.SU2Action.CommT2Norm > 1e-8 && g.SU2Action.CommT3Norm < 1e-8
	return ReconciliationAudit{
		Executed:                              true,
		ScalarResponseAnisotropic:             anisotropic,
		ScalarResponseBreaksT1T2:              breaks,
		InnerFluctuationSelectsRepresentation: d.ScalarSU2RepresentationProvenanceClosed,
		NoContradiction:                       anisotropic && breaks && d.ScalarSU2RepresentationProvenanceClosed,
		RepresentationVsResponseSeparated:     true,
		GoldstoneBridgeOrbitPreserved:         g.GaugeOrbit.BridgeGoldstoneOrbitConsistent,
		NativeGaugeEatingStillBlocked:         !d.NativeDphiActionDerived && !d.PhysicalMassMatrixDerived,
		Verdict:                               StatusScalarResponseObstructionResolved,
		Reason:                                "The scalar response commutator test and the inner-fluctuation field-content theorem answer different questions: response anisotropy does not select the full symmetry, while finite one-forms identify the Higgs representation socket.",
	}
}

func buildBoundary(d DphiProvenanceAudit) Boundary {
	return Boundary{
		Executed: true,
		StructuralScalarDoubletProvenancePromoted:  d.ScalarSU2RepresentationProvenanceClosed,
		StructuralDphiTransformationSocketPromoted: d.StructuralDphiSocketFound,
		NativeFullDphiActionClosed:                 d.NativeDphiActionDerived,
		NativeScalarKineticProjectionClosed:        d.ProductGeometryKineticProjectionDerived && d.ScalarKineticNormalizationDerived,
		NativeVacuumOrientationClosed:              false,
		NativeKappaSelected:                        false,
		NativeGaugeHessianSelected:                 d.GaugeHessianCouplingsDerived,
		NativeWZMassMatrixDerived:                  d.PhysicalMassMatrixDerived,
		Verdict:                                    strings.Join([]string{StatusStructuralRepresentationPromoted, StatusFailedNativeDphiActionNotDerived, StatusFailedPhysicalWZMassBlocked}, ";"),
		Reason:                                     "Gate499 promotes the scalar doublet/DΦ transformation socket as structural inner-fluctuation field content, but blocks the native action, kinetic, gauge-Hessian, kappa, vacuum, and W/Z mass writes.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                  true,
		ObservedWMassImported:     false,
		ObservedZMassImported:     false,
		ObservedHiggsMassImported: false,
		FermiConstantImported:     false,
		WeakAngleImported:         false,
		FineStructureImported:     false,
		GaugeCouplingImported:     false,
		HiggsVEVImported:          false,
		YukawaImported:            false,
		CKMPMNSImported:           false,
		NativeDphiActionWritten:   false,
		NativeKineticWritten:      false,
		NativeKappaWritten:        false,
		NativeWZMassWritten:       false,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "No W/Z/Higgs mass, Fermi constant, weak angle, alpha, gauge coupling, VEV, Yukawa, CKM, or PMNS datum is imported; no native action or W/Z mass registry write is made.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No physical electroweak mass, coupling, kappa, or kinetic-normalization native entry is admitted at Gate499.",
		},
		BridgeEntries: []string{
			"The finite inner-fluctuation ledger structurally recovers one complex Higgs doublet from Ω¹_D(A_F).",
			"The structural DΦ transformation socket is accepted: gauge connection plus scalar one-form plus left/right action.",
			"Gate498 scalar-response obstruction is reconciled: response anisotropy does not disprove the Higgs representation provenance from A_F.",
		},
		EnvironmentalEntries: []string{
			"Observed W/Z/Higgs masses, Higgs VEV, weak angle, alpha, gauge couplings, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedNativeDphiActionNotDerived,
			StatusFailedHeatKernelScalarKineticMissing,
			StatusFailedVacuumOrientationStillBridge,
			StatusFailedKappaStillBridge,
			StatusFailedGaugeHessianCouplingsNotDerived,
			StatusFailedPhysicalWZMassBlocked,
		},
		OpenTheorems: []string{
			"Derive the product spectral-action scalar kinetic projection from the finite/continuum product geometry.",
			"Derive scalar kinetic normalization, gauge Hessian/couplings, and vacuum orientation before promoting W/Z masses.",
			"Prove whether kappa_U1=6 follows from the normalized product-action Hessian rather than whitening diagnostics.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 500, Title: "Product Spectral-Action Scalar Kinetic Projection Audit", Reason: "Gate499 closes representation-level provenance but leaves the action-level kinetic projection unproved.", PrimaryTask: "derive or block the scalar kinetic term and normalized DΦ†DΦ coefficient from the product spectral action without importing W/Z masses, weak angle, VEV, or gauge couplings"}
}

func truth(a Analysis) string {
	if a.Boundary.StructuralScalarDoubletProvenancePromoted && !a.Boundary.NativeFullDphiActionClosed {
		return "Gate499 separates representation provenance from action dynamics.  The finite inner-fluctuation ledger does select the structural Higgs doublet/DΦ transformation socket: one complex weak doublet arises from finite one-forms over A_F=C⊕H⊕M3(C), so the scalar SU(2)L representation is not a random bridge realification.  But the normalized product-action kinetic projection, vacuum orientation, gauge Hessian/couplings, kappa_U1, and W/Z mass matrix remain unproved and firewalled."
	}
	return "Gate499 did not close the structural inner-fluctuation scalar representation provenance."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate498AuditDefined, "Gate498 inheritance missing"},
		{a.Inheritance.FullScalarSU2NotSelectedByResponse && a.Inheritance.NativeDphiStillOpen, "Gate498 obstruction not inherited"},
		{a.InnerFluctuation.Executed && a.InnerFluctuation.GaugeBosonContentRecovered && a.InnerFluctuation.GaugeBosonDimension == 12, "inner-fluctuation gauge content not recovered"},
		{a.InnerFluctuation.HiggsDoubletRecovered && a.InnerFluctuation.ComplexDoublets == 1 && a.InnerFluctuation.RealScalarDimension == 4, "single complex Higgs doublet not recovered"},
		{a.Dphi.StructuralDphiSocketFound && a.Dphi.ScalarSU2RepresentationProvenanceClosed, "Dphi structural socket not promoted"},
		{!a.Dphi.NativeDphiActionDerived && !a.Dphi.ProductGeometryKineticProjectionDerived && !a.Dphi.PhysicalMassMatrixDerived, "Gate499 over-promoted Dphi action or masses"},
		{a.Reconciliation.NoContradiction && a.Reconciliation.NativeGaugeEatingStillBlocked, "response/representation reconciliation failed"},
		{a.Boundary.StructuralScalarDoubletProvenancePromoted && !a.Boundary.NativeFullDphiActionClosed && !a.Boundary.NativeWZMassMatrixDerived, "boundary over-promoted native Dphi or W/Z"},
		{a.Firewall.Executed && !a.Firewall.ObservedWMassImported && !a.Firewall.WeakAngleImported && !a.Firewall.HiggsVEVImported && !a.Firewall.NativeWZMassWritten, "firewall violation"},
		{a.Next.Gate == 500, "Gate500 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate498=%t J_socket=%t abstract_SU2=%t response_U1_only=%t full_SU2_response_fail=%t bridge_goldstone=%t native_Dphi_open=%t kappa_bridge=%t no_data=%t verdict=%s reason=%s", x.Gate498AuditDefined, x.ComplexDoubletSocketFound, x.AbstractSU2ClosureConfirmed, x.ScalarResponseSelectsOnlyU1, x.FullScalarSU2NotSelectedByResponse, x.BridgeGoldstoneOrbitConsistent, x.NativeDphiStillOpen, x.KappaStillBridge, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatInnerFluctuation(x InnerFluctuationAudit) string {
	return fmt.Sprintf("gate298=%t A=%s oneforms=%t DA=%q gauge=%t dim=%d Higgs=%t complex=%d real=%d weak=%s color=%s Y=%s edges=%d YukawaFree=%t potentialMissing=%t heatMissing=%t structural=%t verdict=%s reason=%s", x.Gate298FieldContentAvailable, x.Algebra, x.NCGOneFormsFormalized, x.FluctuatedDiracFormula, x.GaugeBosonContentRecovered, x.GaugeBosonDimension, x.HiggsDoubletRecovered, x.ComplexDoublets, x.RealScalarDimension, x.WeakRepresentation, x.ColorRepresentation, x.HyperchargeRay, x.FiniteOneFormEdges, x.NumericalYukawaFree, x.HiggsPotentialNotDerived, x.HeatKernelProjectionMissing, x.StructuralFieldContent, x.Verdict, x.Reason)
}

func FormatDphi(x DphiProvenanceAudit) string {
	return fmt.Sprintf("gauge_conn=%t scalar_oneform=%t left_right=%t socket=%t representation_closed=%t product_kinetic=%t native_action=%t scalar_norm=%t hessian_couplings=%t masses=%t verdict=%s reason=%s", x.StructuralGaugeConnectionAvailable, x.StructuralScalarOneFormAvailable, x.StructuralLeftRightActionAvailable, x.StructuralDphiSocketFound, x.ScalarSU2RepresentationProvenanceClosed, x.ProductGeometryKineticProjectionDerived, x.NativeDphiActionDerived, x.ScalarKineticNormalizationDerived, x.GaugeHessianCouplingsDerived, x.PhysicalMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatReconciliation(x ReconciliationAudit) string {
	return fmt.Sprintf("anisotropic=%t breaks_T1T2=%t inner_selects_rep=%t no_contradiction=%t separate_rep_response=%t goldstone_bridge=%t native_eating_blocked=%t verdict=%s reason=%s", x.ScalarResponseAnisotropic, x.ScalarResponseBreaksT1T2, x.InnerFluctuationSelectsRepresentation, x.NoContradiction, x.RepresentationVsResponseSeparated, x.GoldstoneBridgeOrbitPreserved, x.NativeGaugeEatingStillBlocked, x.Verdict, x.Reason)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("structural_doublet=%t structural_Dphi=%t native_Dphi_action=%t kinetic_projection=%t vacuum=%t kappa=%t hessian=%t WZ=%t verdict=%s reason=%s", x.StructuralScalarDoubletProvenancePromoted, x.StructuralDphiTransformationSocketPromoted, x.NativeFullDphiActionClosed, x.NativeScalarKineticProjectionClosed, x.NativeVacuumOrientationClosed, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("W=%t Z=%t H=%t Fermi=%t theta=%t alpha=%t gauge=%t v=%t Yukawa=%t CKM_PMNS=%t native_Dphi=%t native_kinetic=%t native_kappa=%t native_WZ=%t verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.HiggsVEVImported, x.YukawaImported, x.CKMPMNSImported, x.NativeDphiActionWritten, x.NativeKineticWritten, x.NativeKappaWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}
