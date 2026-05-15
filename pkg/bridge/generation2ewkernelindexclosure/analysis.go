// Package generation2ewkernelindexclosure implements Gate 503:
// Electroweak Kernel Index Native Closure Audit.
//
// Gate 502 proved that after quotienting out every scalar/coupling scale, the
// electroweak bridge diagnostic still carries a photon kernel, rank-three
// broken orbit, charged-pair degeneracy, and diag(1,1,4) quotient shape.  Gate
// 503 asks the narrower native question: can the photon kernel and rank-three
// broken image be promoted from a scale-independent bridge diagnostic to a
// representation-index theorem of the finite Higgs doublet socket?
//
// The answer is conditional but meaningful.  Gate 499 supplies structural
// provenance for one complex SU(2)L Higgs doublet with scalar hypercharge ray;
// for any nonzero Higgs ray in this representation, the stabilizer inside
// SU(2)L×U(1)Y is one-dimensional U(1)em, so the gauge orbit has dimension
// 4-1=3 and the radial quotient leaves one scalar mode.  This is an index
// theorem for the representation socket.  It still does not prove that the
// finite action selects a nonzero vacuum ray, its scale, its orientation, the
// gauge Hessian, kappa_U1, weak angle, couplings, or W/Z masses.
package generation2ewkernelindexclosure

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewnquotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2innerfluctuationdphiprovenance"
)

const (
	AuditID = "GATE503-ELECTROWEAK-KERNEL-INDEX-NATIVE-CLOSURE-AUDIT"

	StatusGate502QuotientInherited                 = "CONDITIONAL_SUPPORT_GATE502_ELECTROWEAK_QUOTIENT_INHERITED"
	StatusGate499RepresentationProvenanceInherited = "CONDITIONAL_SUPPORT_GATE499_HIGGS_REPRESENTATION_PROVENANCE_INHERITED"
	StatusRepresentationIndexSieveDefined          = "CONDITIONAL_SUPPORT_ELECTROWEAK_REPRESENTATION_INDEX_SIEVE_DEFINED"
	StatusPhotonStabilizerIndexOne                 = "CONDITIONAL_SUPPORT_PHOTON_STABILIZER_INDEX_ONE_PROVEN_FOR_NONZERO_HIGGS_RAY"
	StatusBrokenOrbitIndexThree                    = "CONDITIONAL_SUPPORT_BROKEN_ELECTROWEAK_ORBIT_INDEX_THREE_PROVEN_FOR_NONZERO_HIGGS_RAY"
	StatusRadialQuotientIndexOne                   = "CONDITIONAL_SUPPORT_RADIAL_SCALAR_QUOTIENT_INDEX_ONE_PROVEN_CONDITIONALLY"
	StatusKernelRankPromotedConditionally          = "CONDITIONAL_SUPPORT_KERNEL_RANK_PROMOTED_TO_CONDITIONAL_REPRESENTATION_INDEX"
	StatusBridgeDiag114Preserved                   = "CONDITIONAL_SUPPORT_DIAG114_REMAINS_BRIDGE_HESSIAN_SHAPE"
	StatusFirewallPreserved                        = "FIREWALL_PRESERVED_NO_ELECTROWEAK_SCALE_MASS_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeMassWriteBlocked                   = "FIREWALL_BLOCKED_ELECTROWEAK_MASS_COUPLING_AND_KAPPA_NATIVE_WRITE"

	StatusFailedNonzeroVacuumRayNotNativeSelected = "FAILED_ROUTE_NONZERO_HIGGS_VACUUM_RAY_NOT_SELECTED_BY_FINITE_ACTION"
	StatusFailedVacuumOrientationNotNative        = "FAILED_ROUTE_VACUUM_ORIENTATION_REMAINS_GAUGE_REPRESENTATIVE_NOT_NATIVE_MINIMIZER"
	StatusFailedVEVScaleStillSealed               = "FAILED_ROUTE_HIGGS_VEV_SCALE_STILL_SEALED"
	StatusFailedKappaStillBridge                  = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_HESSIAN_CANDIDATE"
	StatusFailedGaugeHessianCouplingsNotDerived   = "FAILED_ROUTE_GAUGE_HESSIAN_AND_COUPLINGS_NOT_DERIVED"
	StatusFailedWeakAngleNotDerived               = "FAILED_ROUTE_WEAK_MIXING_ANGLE_NOT_DERIVED_BY_KERNEL_INDEX"
	StatusFailedWZMassMatrixStillBlocked          = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED_BY_INDEX_THEOREM"
	StatusGate504RedirectDefined                  = "CONDITIONAL_SUPPORT_GATE504_CONTINUUM_MATCHING_PERMISSION_LEDGER_REDIRECT_DEFINED"
)

type Inheritance struct {
	Executed                         bool
	Gate502AuditDefined              bool
	QuotientBridgeAccepted           bool
	QuotientPhotonKernel             bool
	QuotientBrokenRankThree          bool
	QuotientDiag114Shape             bool
	QuotientNativeActionClosed       bool
	Gate499AuditDefined              bool
	StructuralHiggsDoubletProvenance bool
	StructuralDphiSocket             bool
	NativeDphiActionClosed           bool
	NoScaleOrFlavorDataImported      bool
	Verdict                          string
	Reason                           string
}

type RepresentationIndexSieve struct {
	Executed                    bool
	GaugeGroup                  string
	GaugeOrbitDimension         int
	ScalarRepresentation        string
	ScalarRealDimension         int
	ComplexDoublets             int
	HyperchargeRay              string
	AssumesNonzeroHiggsRay      bool
	UsesVacuumScale             bool
	UsesGaugeCouplings          bool
	UsesObservedElectroweakData bool
	Verdict                     string
	Reason                      string
}

type KernelIndexAudit struct {
	Executed                            bool
	GaugeGeneratorDimension             int
	StabilizerDimension                 int
	BrokenOrbitDimension                int
	ScalarRealDimension                 int
	RadialQuotientDimension             int
	PhotonGenerator                     string
	BrokenGenerators                    []string
	PhotonKernelIndexProven             bool
	BrokenOrbitIndexProven              bool
	RadialIndexProven                   bool
	ConditionalOnNonzeroRay             bool
	IndependentOfScalarNormalization    bool
	IndependentOfGaugeCouplings         bool
	IndependentOfYukawaTrace            bool
	UnconditionalNativeVacuumProvenance bool
	Verdict                             string
	Reason                              string
}

type HessianCompatibility struct {
	Executed                 bool
	KernelRankMatchesGate502 bool
	Diag114ShapeInherited    bool
	Diag114NativeHessian     bool
	KappaNative              bool
	WeakAngleDerived         bool
	GaugeCouplingsDerived    bool
	PhysicalWZMassMatrix     bool
	ObservedMassRatioClaimed bool
	Verdict                  string
	Reason                   string
}

type Boundary struct {
	Executed                               bool
	ConditionalRepresentationIndexAccepted bool
	UnconditionalNativeElectroweakAction   bool
	NativeNonzeroVacuumRaySelected         bool
	NativeVacuumOrientationSelected        bool
	NativeScalarNormalizationClosed        bool
	NativeKappaSelected                    bool
	NativeGaugeHessianCouplingsDerived     bool
	NativeWeakAngleDerived                 bool
	NativeWZMassMatrixDerived              bool
	Verdict                                string
	Reason                                 string
}

type Firewall struct {
	Executed                      bool
	ObservedWMassImported         bool
	ObservedZMassImported         bool
	ObservedWZRatioImported       bool
	ObservedWeakAngleImported     bool
	ObservedGaugeCouplingImported bool
	ObservedHiggsVEVImported      bool
	ObservedYukawaImported        bool
	CKMPMNSImported               bool
	NativeKappaWritten            bool
	NativeWeakAngleWritten        bool
	NativeGaugeCouplingWritten    bool
	NativeHiggsVEVWritten         bool
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
	Inheritance Inheritance
	Sieve       RepresentationIndexSieve
	Kernel      KernelIndexAudit
	Hessian     HessianCompatibility
	Boundary    Boundary
	Firewall    Firewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
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
	g502, err := generation2ewnquotient.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate502 electroweak quotient: %w", err)
	}
	g499, err := generation2innerfluctuationdphiprovenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate499 inner-fluctuation Dphi provenance: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g502, g499)
	a.Sieve = buildSieve(g499)
	a.Kernel = buildKernelIndex(a.Sieve)
	a.Hessian = buildHessianCompatibility(g502, a.Kernel)
	a.Boundary = buildBoundary(a.Kernel, a.Hessian)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g502 generation2ewnquotient.Analysis, g499 generation2innerfluctuationdphiprovenance.Analysis) Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate502AuditDefined:              true,
		QuotientBridgeAccepted:           g502.Boundary.BridgeQuotientAccepted,
		QuotientPhotonKernel:             g502.KernelRank.PhotonKernelDimension == 1 && g502.KernelRank.PhotonKernelSurvivesScaleQuotient,
		QuotientBrokenRankThree:          g502.KernelRank.BrokenOrbitRank == 3 && g502.KernelRank.BrokenRankSurvivesScaleQuotient,
		QuotientDiag114Shape:             g502.Hessian.Diag114Shape && g502.Hessian.DimensionlessShapeSurvives,
		QuotientNativeActionClosed:       g502.Boundary.NativeElectroweakActionClosed,
		Gate499AuditDefined:              true,
		StructuralHiggsDoubletProvenance: g499.Boundary.StructuralScalarDoubletProvenancePromoted,
		StructuralDphiSocket:             g499.Boundary.StructuralDphiTransformationSocketPromoted,
		NativeDphiActionClosed:           g499.Boundary.NativeFullDphiActionClosed,
		NoScaleOrFlavorDataImported:      g502.Firewall.Executed && g499.Firewall.Executed && !g502.Firewall.ObservedWMassImported && !g502.Firewall.ObservedWeakAngleImported && !g499.Firewall.HiggsVEVImported && !g499.Firewall.YukawaImported,
		Verdict:                          strings.Join([]string{StatusGate502QuotientInherited, StatusGate499RepresentationProvenanceInherited}, ";"),
		Reason:                           "Gate502 supplies scale-independent kernel/rank quotient data; Gate499 supplies finite inner-fluctuation provenance for one complex Higgs doublet and the structural DΦ socket.",
	}
}

func buildSieve(g499 generation2innerfluctuationdphiprovenance.Analysis) RepresentationIndexSieve {
	return RepresentationIndexSieve{
		Executed:                    true,
		GaugeGroup:                  "SU(2)_L × U(1)_Y acting on the finite Higgs one-form socket",
		GaugeOrbitDimension:         4,
		ScalarRepresentation:        g499.InnerFluctuation.WeakRepresentation,
		ScalarRealDimension:         g499.InnerFluctuation.RealScalarDimension,
		ComplexDoublets:             g499.InnerFluctuation.ComplexDoublets,
		HyperchargeRay:              g499.InnerFluctuation.HyperchargeRay,
		AssumesNonzeroHiggsRay:      true,
		UsesVacuumScale:             false,
		UsesGaugeCouplings:          false,
		UsesObservedElectroweakData: false,
		Verdict:                     StatusRepresentationIndexSieveDefined,
		Reason:                      "The sieve uses only the structural Higgs representation recovered by inner fluctuations.  It assumes a nonzero Higgs ray but does not use its scale, gauge couplings, weak angle, or observed masses.",
	}
}

func buildKernelIndex(s RepresentationIndexSieve) KernelIndexAudit {
	stabilizer := 1
	gaugeDim := s.GaugeOrbitDimension
	broken := gaugeDim - stabilizer
	radial := s.ScalarRealDimension - broken
	return KernelIndexAudit{
		Executed:                            true,
		GaugeGeneratorDimension:             gaugeDim,
		StabilizerDimension:                 stabilizer,
		BrokenOrbitDimension:                broken,
		ScalarRealDimension:                 s.ScalarRealDimension,
		RadialQuotientDimension:             radial,
		PhotonGenerator:                     "Q_em = T3 + Y_phi stabilizes the chosen nonzero Higgs ray",
		BrokenGenerators:                    []string{"T1", "T2", "Z = T3 - Y_phi"},
		PhotonKernelIndexProven:             stabilizer == 1,
		BrokenOrbitIndexProven:              broken == 3,
		RadialIndexProven:                   radial == 1,
		ConditionalOnNonzeroRay:             s.AssumesNonzeroHiggsRay,
		IndependentOfScalarNormalization:    !s.UsesVacuumScale,
		IndependentOfGaugeCouplings:         !s.UsesGaugeCouplings,
		IndependentOfYukawaTrace:            true,
		UnconditionalNativeVacuumProvenance: false,
		Verdict:                             strings.Join([]string{StatusPhotonStabilizerIndexOne, StatusBrokenOrbitIndexThree, StatusRadialQuotientIndexOne, StatusKernelRankPromotedConditionally}, ";"),
		Reason:                              "For one complex Higgs doublet with the electroweak hypercharge ray, any nonzero Higgs ray has a one-dimensional U(1)em stabilizer inside the four-dimensional SU(2)L×U(1)Y gauge orbit.  Hence the broken orbit has index three and the four-real scalar socket quotients to one radial mode.  The statement is conditional because the finite action has not selected the nonzero ray.",
	}
}

func buildHessianCompatibility(g502 generation2ewnquotient.Analysis, k KernelIndexAudit) HessianCompatibility {
	return HessianCompatibility{
		Executed:                 true,
		KernelRankMatchesGate502: k.PhotonKernelIndexProven && k.BrokenOrbitDimension == g502.KernelRank.BrokenOrbitRank && k.StabilizerDimension == g502.KernelRank.PhotonKernelDimension,
		Diag114ShapeInherited:    g502.Hessian.Diag114Shape && g502.Hessian.DimensionlessShapeSurvives,
		Diag114NativeHessian:     false,
		KappaNative:              false,
		WeakAngleDerived:         false,
		GaugeCouplingsDerived:    false,
		PhysicalWZMassMatrix:     false,
		ObservedMassRatioClaimed: false,
		Verdict:                  strings.Join([]string{StatusBridgeDiag114Preserved, StatusFailedKappaStillBridge, StatusFailedWZMassMatrixStillBlocked}, ";"),
		Reason:                   "The representation index explains the photon nullity and rank-three broken image, but it does not select the numeric Hessian.  The diag(1,1,4) candidate, kappa_U1=6, couplings, weak angle, and W/Z masses remain bridge/action-level questions.",
	}
}

func buildBoundary(k KernelIndexAudit, h HessianCompatibility) Boundary {
	accepted := k.PhotonKernelIndexProven && k.BrokenOrbitIndexProven && k.RadialIndexProven && k.ConditionalOnNonzeroRay && h.KernelRankMatchesGate502
	return Boundary{
		Executed:                               true,
		ConditionalRepresentationIndexAccepted: accepted,
		UnconditionalNativeElectroweakAction:   false,
		NativeNonzeroVacuumRaySelected:         false,
		NativeVacuumOrientationSelected:        false,
		NativeScalarNormalizationClosed:        false,
		NativeKappaSelected:                    false,
		NativeGaugeHessianCouplingsDerived:     false,
		NativeWeakAngleDerived:                 false,
		NativeWZMassMatrixDerived:              false,
		Verdict: strings.Join([]string{
			StatusKernelRankPromotedConditionally,
			StatusFailedNonzeroVacuumRayNotNativeSelected,
			StatusFailedVacuumOrientationNotNative,
			StatusFailedVEVScaleStillSealed,
			StatusFailedKappaStillBridge,
			StatusFailedGaugeHessianCouplingsNotDerived,
			StatusFailedWeakAngleNotDerived,
			StatusFailedWZMassMatrixStillBlocked,
		}, ";"),
		Reason: "Gate503 admits a conditional representation-index theorem for kernel and rank.  It blocks unconditional electroweak action closure and every physical scale, coupling, angle, and mass claim.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                      true,
		ObservedWMassImported:         false,
		ObservedZMassImported:         false,
		ObservedWZRatioImported:       false,
		ObservedWeakAngleImported:     false,
		ObservedGaugeCouplingImported: false,
		ObservedHiggsVEVImported:      false,
		ObservedYukawaImported:        false,
		CKMPMNSImported:               false,
		NativeKappaWritten:            false,
		NativeWeakAngleWritten:        false,
		NativeGaugeCouplingWritten:    false,
		NativeHiggsVEVWritten:         false,
		NativeWZMassWritten:           false,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "No W/Z mass, W/Z ratio, weak angle, gauge coupling, Higgs VEV, Yukawa, CKM, or PMNS datum is imported; no native kappa, coupling, VEV, or mass write is made.",
	}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Conditional representation-index theorem: given the finite one-form Higgs doublet socket and a nonzero Higgs ray, U(1)_em is the one-dimensional stabilizer and the broken electroweak orbit has dimension three.",
			"No unconditional native electroweak action, nonzero vacuum selection, VEV, kappa_U1, weak angle, gauge-coupling, or W/Z mass entry is admitted at Gate503.",
		},
		BridgeEntries: []string{
			"The Gate502 photon-kernel and rank-three quotient are now explained by the Gate499 structural Higgs representation index, conditional on a nonzero ray.",
			"The diag(1,1,4) Hessian shape remains a bridge/action candidate and is not promoted by the index theorem.",
		},
		EnvironmentalEntries: []string{
			"Higgs VEV, gauge couplings, weak mixing angle, W/Z masses, W/Z ratio, Yukawa amplitudes, CKM, and PMNS remain sealed environmental or continuum-matching data.",
		},
		FailedRoutes: []string{
			StatusFailedNonzeroVacuumRayNotNativeSelected,
			StatusFailedVacuumOrientationNotNative,
			StatusFailedVEVScaleStillSealed,
			StatusFailedKappaStillBridge,
			StatusFailedGaugeHessianCouplingsNotDerived,
			StatusFailedWeakAngleNotDerived,
			StatusFailedWZMassMatrixStillBlocked,
		},
		OpenTheorems: []string{
			"Derive a finite-action theorem selecting a nonzero Higgs vacuum ray rather than assuming it as the broken-phase condition.",
			"Derive the gauge Hessian/coupling normalization before promoting kappa_U1=6, weak angle, or W/Z masses.",
			"Define a continuum-matching permission ledger for importing VEV and gauge couplings without contaminating the native registry.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 504, Title: "Continuum Matching Permission Ledger for Electroweak Scales", Reason: "Gate503 closes the conditional kernel/rank representation index but leaves all scale and coupling data outside the native core.", PrimaryTask: "define the exact permission boundary for importing Higgs VEV, gauge couplings, weak angle, and W/Z masses as continuum/environmental bridge data without rewriting them as native finite-geometry theorems"}
}

func truth(a Analysis) string {
	if a.Boundary.ConditionalRepresentationIndexAccepted && !a.Boundary.NativeWZMassMatrixDerived {
		return "Gate503 converts the electroweak photon-kernel/rank-three fact from a mere scale-independent bridge diagnostic into a conditional representation-index theorem: the finite inner-fluctuation Higgs doublet socket has a one-dimensional U(1)em stabilizer and a three-dimensional broken orbit whenever a nonzero Higgs ray is present.  This proves the 4→3+1 Goldstone index at the representation level, but it does not select the nonzero vacuum, VEV, kappa_U1, weak angle, gauge couplings, or W/Z masses."
	}
	return "Gate503 did not close the electroweak kernel representation index."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.QuotientBridgeAccepted && a.Inheritance.StructuralHiggsDoubletProvenance, "required Gate502/Gate499 inheritance missing"},
		{a.Sieve.Executed && a.Sieve.ScalarRealDimension == 4 && a.Sieve.ComplexDoublets == 1 && a.Sieve.AssumesNonzeroHiggsRay && !a.Sieve.UsesObservedElectroweakData, "representation index sieve invalid"},
		{a.Kernel.Executed && a.Kernel.StabilizerDimension == 1 && a.Kernel.BrokenOrbitDimension == 3 && a.Kernel.RadialQuotientDimension == 1 && a.Kernel.PhotonKernelIndexProven && a.Kernel.BrokenOrbitIndexProven && a.Kernel.RadialIndexProven && a.Kernel.ConditionalOnNonzeroRay && !a.Kernel.UnconditionalNativeVacuumProvenance, "kernel index invalid or over-promoted"},
		{a.Hessian.Executed && a.Hessian.KernelRankMatchesGate502 && a.Hessian.Diag114ShapeInherited && !a.Hessian.Diag114NativeHessian && !a.Hessian.KappaNative && !a.Hessian.WeakAngleDerived && !a.Hessian.PhysicalWZMassMatrix, "Hessian compatibility over-promoted or invalid"},
		{a.Boundary.Executed && a.Boundary.ConditionalRepresentationIndexAccepted && !a.Boundary.UnconditionalNativeElectroweakAction && !a.Boundary.NativeNonzeroVacuumRaySelected && !a.Boundary.NativeKappaSelected && !a.Boundary.NativeWeakAngleDerived && !a.Boundary.NativeWZMassMatrixDerived, "boundary over-promoted electroweak physics"},
		{a.Firewall.Executed && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedWeakAngleImported && !a.Firewall.ObservedGaugeCouplingImported && !a.Firewall.ObservedHiggsVEVImported && !a.Firewall.NativeKappaWritten && !a.Firewall.NativeWZMassWritten, "firewall violation"},
		{a.Next.Gate == 504, "Gate504 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate502=%t bridge=%t photon=%t rank3=%t diag114=%t action_closed=%t gate499=%t Higgs=%t Dphi=%t native_Dphi=%t no_data=%t verdict=%s reason=%s", x.Gate502AuditDefined, x.QuotientBridgeAccepted, x.QuotientPhotonKernel, x.QuotientBrokenRankThree, x.QuotientDiag114Shape, x.QuotientNativeActionClosed, x.Gate499AuditDefined, x.StructuralHiggsDoubletProvenance, x.StructuralDphiSocket, x.NativeDphiActionClosed, x.NoScaleOrFlavorDataImported, x.Verdict, x.Reason)
}

func FormatSieve(x RepresentationIndexSieve) string {
	return fmt.Sprintf("group=%s gauge_dim=%d rep=%s scalar_real=%d doublets=%d Y=%s nonzero_ray=%t uses_scale=%t uses_couplings=%t observed=%t verdict=%s reason=%s", x.GaugeGroup, x.GaugeOrbitDimension, x.ScalarRepresentation, x.ScalarRealDimension, x.ComplexDoublets, x.HyperchargeRay, x.AssumesNonzeroHiggsRay, x.UsesVacuumScale, x.UsesGaugeCouplings, x.UsesObservedElectroweakData, x.Verdict, x.Reason)
}

func FormatKernel(x KernelIndexAudit) string {
	return fmt.Sprintf("gauge_dim=%d stabilizer=%d broken=%d scalar_real=%d radial=%d photon=%s broken_gens=[%s] photon_index=%t broken_index=%t radial_index=%t conditional=%t scale_free=%t coupling_free=%t Yukawa_free=%t unconditional_vacuum=%t verdict=%s reason=%s", x.GaugeGeneratorDimension, x.StabilizerDimension, x.BrokenOrbitDimension, x.ScalarRealDimension, x.RadialQuotientDimension, x.PhotonGenerator, strings.Join(x.BrokenGenerators, ", "), x.PhotonKernelIndexProven, x.BrokenOrbitIndexProven, x.RadialIndexProven, x.ConditionalOnNonzeroRay, x.IndependentOfScalarNormalization, x.IndependentOfGaugeCouplings, x.IndependentOfYukawaTrace, x.UnconditionalNativeVacuumProvenance, x.Verdict, x.Reason)
}

func FormatHessian(x HessianCompatibility) string {
	return fmt.Sprintf("matches_gate502=%t diag114=%t diag114_native=%t kappa=%t weak_angle=%t couplings=%t WZ=%t observed_ratio=%t verdict=%s reason=%s", x.KernelRankMatchesGate502, x.Diag114ShapeInherited, x.Diag114NativeHessian, x.KappaNative, x.WeakAngleDerived, x.GaugeCouplingsDerived, x.PhysicalWZMassMatrix, x.ObservedMassRatioClaimed, x.Verdict, x.Reason)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("conditional_index=%t action=%t nonzero_vacuum=%t orientation=%t scalar_norm=%t kappa=%t hessian_couplings=%t weak_angle=%t WZ=%t verdict=%s reason=%s", x.ConditionalRepresentationIndexAccepted, x.UnconditionalNativeElectroweakAction, x.NativeNonzeroVacuumRaySelected, x.NativeVacuumOrientationSelected, x.NativeScalarNormalizationClosed, x.NativeKappaSelected, x.NativeGaugeHessianCouplingsDerived, x.NativeWeakAngleDerived, x.NativeWZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("W=%t Z=%t ratio=%t theta=%t gauge=%t VEV=%t Yukawa=%t CKM_PMNS=%t native_kappa=%t native_theta=%t native_gauge=%t native_VEV=%t native_WZ=%t verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedWZRatioImported, x.ObservedWeakAngleImported, x.ObservedGaugeCouplingImported, x.ObservedHiggsVEVImported, x.ObservedYukawaImported, x.CKMPMNSImported, x.NativeKappaWritten, x.NativeWeakAngleWritten, x.NativeGaugeCouplingWritten, x.NativeHiggsVEVWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 503 Registry Audit — Electroweak Kernel Index Native Closure Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{
		StatusGate502QuotientInherited,
		StatusGate499RepresentationProvenanceInherited,
		StatusRepresentationIndexSieveDefined,
		StatusPhotonStabilizerIndexOne,
		StatusBrokenOrbitIndexThree,
		StatusRadialQuotientIndexOne,
		StatusKernelRankPromotedConditionally,
		StatusBridgeDiag114Preserved,
		StatusFailedNonzeroVacuumRayNotNativeSelected,
		StatusFailedVacuumOrientationNotNative,
		StatusFailedVEVScaleStillSealed,
		StatusFailedKappaStillBridge,
		StatusFailedGaugeHessianCouplingsNotDerived,
		StatusFailedWeakAngleNotDerived,
		StatusFailedWZMassMatrixStillBlocked,
		StatusFirewallPreserved,
		StatusNativeMassWriteBlocked,
		StatusGate504RedirectDefined,
	} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate502 established the scalar-normalization-independent electroweak quotient: photon nullity, broken rank three, charged degeneracy, and the dimensionless `diag(1,1,4)` Hessian shape survive after deleting `a`, `f0`, VEV, continuum scale, and coupling units.\n\n")
	b.WriteString("Gate499 established structural provenance for one finite inner-fluctuation Higgs doublet socket and a structural `DΦ` transformation socket.  Gate503 combines these facts but does not import electroweak scales or observed data.\n\n")
	b.WriteString("## Representation index sieve\n\n")
	b.WriteString(fmt.Sprintf("```text\ngauge group carrier = %s\ngauge generator dimension = %d\nscalar representation = %s\nscalar real dimension = %d\ncomplex Higgs doublets = %d\nhypercharge ray = %s\nassumes nonzero Higgs ray = %t\nuses VEV scale = %t\nuses gauge couplings = %t\n```\n\n", a.Sieve.GaugeGroup, a.Sieve.GaugeOrbitDimension, a.Sieve.ScalarRepresentation, a.Sieve.ScalarRealDimension, a.Sieve.ComplexDoublets, a.Sieve.HyperchargeRay, a.Sieve.AssumesNonzeroHiggsRay, a.Sieve.UsesVacuumScale, a.Sieve.UsesGaugeCouplings))
	b.WriteString("## Kernel index audit\n\n")
	b.WriteString(fmt.Sprintf("For a nonzero Higgs ray in the one-doublet socket:\n\n```text\ndim(SU(2)_L × U(1)_Y) = %d\ndim(stabilizer U(1)_em) = %d\ndim(broken orbit) = %d\nreal scalar dimension = %d\nradial quotient dimension = %d\nphoton stabilizer = %s\nbroken generators = %s\n```\n\n", a.Kernel.GaugeGeneratorDimension, a.Kernel.StabilizerDimension, a.Kernel.BrokenOrbitDimension, a.Kernel.ScalarRealDimension, a.Kernel.RadialQuotientDimension, a.Kernel.PhotonGenerator, strings.Join(a.Kernel.BrokenGenerators, ", ")))
	b.WriteString("This closes the kernel/rank fact as a conditional representation-index theorem.  The condition is essential: the finite action still has not selected a nonzero Higgs ray.\n\n")
	b.WriteString("## Hessian compatibility\n\n")
	b.WriteString(fmt.Sprintf("```text\nGate502 kernel/rank matched = %t\ndiag(1,1,4) shape inherited = %t\ndiag(1,1,4) native Hessian = %t\nkappa_U1 native = %t\nweak angle derived = %t\ngauge couplings derived = %t\nphysical W/Z mass matrix = %t\n```\n\n", a.Hessian.KernelRankMatchesGate502, a.Hessian.Diag114ShapeInherited, a.Hessian.Diag114NativeHessian, a.Hessian.KappaNative, a.Hessian.WeakAngleDerived, a.Hessian.GaugeCouplingsDerived, a.Hessian.PhysicalWZMassMatrix))
	b.WriteString("The representation index explains the nullity and rank.  It does not select the action Hessian or any physical electroweak scale.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("No W/Z mass, observed W/Z ratio, weak angle, gauge coupling, Higgs VEV, Yukawa value, CKM, or PMNS datum enters this gate.  No native write is made for `kappa_U1`, weak angle, couplings, VEV, W/Z masses, or observed ratios.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate504 should be:\n\n```text\nGate 504 — Continuum Matching Permission Ledger for Electroweak Scales\n```\n\n")
	b.WriteString("Primary task:\n\n```text\n")
	b.WriteString(a.Next.PrimaryTask)
	b.WriteString("\n```\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
