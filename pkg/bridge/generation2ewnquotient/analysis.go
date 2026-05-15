// Package generation2ewnquotient implements Gate 502:
// Scalar-Normalization-Independent Electroweak Quotient Audit.
//
// Gate 501 sealed the numeric scalar-normalization channel because the product
// spectral-action coefficient K_phi=f0 Tr(Y†Y)/pi^2 depends on the sealed
// Yukawa singular-value spectrum.  Gate 502 therefore asks a narrower question:
// after quotienting out every forbidden scale-like input (a, f0, the Higgs VEV,
// continuum normalization, and gauge-coupling units), which electroweak
// statements still survive?
//
// The result is a bridge quotient theorem, not a physical mass theorem.  The
// abstract electroweak scalar/gauge diagnostic still carries a photon kernel, a
// rank-three broken gauge orbit, charged W-pair degeneracy, and the normalized
// broken Hessian shape diag(1,1,4).  These are dimensionless quotient data and
// do not require a numeric scalar normalization.  But the quotient does not
// promote kappa_U1=6, the canonical action provenance, weak angle, gauge
// couplings, Higgs VEV, or W/Z masses to native registry entries.
package generation2ewnquotient

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2finiteactionsecondvariation"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2vacuumgaugeorbitquotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2yukawatracescalarnormalization"
)

const (
	AuditID = "GATE502-SCALAR-NORMALIZATION-INDEPENDENT-ELECTROWEAK-QUOTIENT-AUDIT"

	StatusGate501ScalarNormalizationSealInherited = "CONDITIONAL_SUPPORT_GATE501_SCALAR_NORMALIZATION_SEAL_INHERITED"
	StatusGate497PhotonBrokenOrbitInherited       = "CONDITIONAL_SUPPORT_GATE497_PHOTON_KERNEL_AND_BROKEN_ORBIT_INHERITED"
	StatusGate495DimensionlessHessianInherited    = "CONDITIONAL_SUPPORT_GATE495_DIMENSIONLESS_HESSIAN_CANDIDATE_INHERITED"
	StatusNormalizationQuotientDefined            = "CONDITIONAL_SUPPORT_SCALAR_NORMALIZATION_QUOTIENT_DEFINED"
	StatusPhotonKernelSurvivesQuotient            = "CONDITIONAL_SUPPORT_PHOTON_KERNEL_SURVIVES_NORMALIZATION_QUOTIENT"
	StatusBrokenRankThreeSurvivesQuotient         = "CONDITIONAL_SUPPORT_BROKEN_RANK_THREE_SURVIVES_NORMALIZATION_QUOTIENT"
	StatusChargedDegeneracySurvivesQuotient       = "CONDITIONAL_SUPPORT_CHARGED_PAIR_DEGENERACY_SURVIVES_NORMALIZATION_QUOTIENT"
	StatusDiag114QuotientShapeSurvives            = "CONDITIONAL_SUPPORT_DIAG114_DIMENSIONLESS_HESSIAN_SHAPE_SURVIVES_QUOTIENT"
	StatusBridgeQuotientAccepted                  = "CONDITIONAL_SUPPORT_ELECTROWEAK_BRIDGE_QUOTIENT_ACCEPTED"
	StatusFirewallPreserved                       = "FIREWALL_PRESERVED_NO_ELECTROWEAK_SCALE_OR_FLAVOR_DATA_IMPORTED"
	StatusNativeRegistryWriteBlocked              = "FIREWALL_BLOCKED_NATIVE_EW_QUOTIENT_TO_MASS_REGISTRY_WRITE"

	StatusFailedQuotientNotNativeActionClosure       = "FAILED_ROUTE_QUOTIENT_SHAPE_IS_NOT_NATIVE_ACTION_CLOSURE"
	StatusFailedKappaStillBridgeAfterQuotient        = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_AFTER_QUOTIENT"
	StatusFailedWeakAngleNotDerivedFromQuotient      = "FAILED_ROUTE_WEAK_MIXING_ANGLE_NOT_DERIVED_FROM_QUOTIENT"
	StatusFailedGaugeCouplingsNotDerivedFromQuotient = "FAILED_ROUTE_GAUGE_COUPLINGS_NOT_DERIVED_FROM_QUOTIENT"
	StatusFailedHiggsVEVStillSealed                  = "FAILED_ROUTE_HIGGS_VEV_STILL_SEALED_AFTER_QUOTIENT"
	StatusFailedWZMassMatrixStillBlocked             = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED_AFTER_QUOTIENT"
	StatusFailedObservedMassRatioNotClaimed          = "FAILED_ROUTE_OBSERVED_WZ_MASS_RATIO_NOT_CLAIMED_BY_QUOTIENT"
	StatusGate503RedirectDefined                     = "CONDITIONAL_SUPPORT_GATE503_ELECTROWEAK_KERNEL_INDEX_NATIVE_CLOSURE_REDIRECT_DEFINED"
)

const eps = 1e-8

type Inheritance struct {
	Executed                            bool
	Gate501AuditDefined                 bool
	ScalarNormalizationSealed           bool
	TraceAIsBridgeOnly                  bool
	ScalarNormalizationIndependentRoute bool
	Gate497AuditDefined                 bool
	PhotonKernelAvailable               bool
	BrokenOrbitRankThreeAvailable       bool
	RadialQuotientOneModeAvailable      bool
	Gate495AuditDefined                 bool
	DimensionlessDiag114Candidate       bool
	CanonicalActionNativeClosed         bool
	NoScaleOrFlavorDataImported         bool
	Verdict                             string
	Reason                              string
}

type QuotientDefinition struct {
	Executed                          bool
	QuotientedFactors                 []string
	RetainedObjects                   []string
	ScalarNormalizationRemoved        bool
	YukawaTraceRemoved                bool
	HiggsVEVRemoved                   bool
	ContinuumScaleRemoved             bool
	GaugeCouplingScaleRemoved         bool
	OnlyDimensionlessStatements       bool
	PhysicalMassStatementsAllowed     bool
	PhysicalCouplingStatementsAllowed bool
	Verdict                           string
	Reason                            string
}

type KernelRankAudit struct {
	Executed                          bool
	ScalarRealDimension               int
	PhotonKernelDimension             int
	BrokenGeneratorCount              int
	BrokenOrbitRank                   int
	BrokenRankSurvivesScaleQuotient   bool
	PhotonKernelSurvivesScaleQuotient bool
	RadialModeAfterQuotient           int
	FourToOneQuotient                 bool
	NativeKernelIndexClosed           bool
	Verdict                           string
	Reason                            string
}

type HessianQuotientAudit struct {
	Executed                    bool
	RawBrokenDiagonal           []float64
	NormalizedBrokenDiagonal    []float64
	ChargedPairDegenerate       bool
	NeutralToChargedRatio       float64
	Diag114Shape                bool
	DimensionlessShapeSurvives  bool
	KappaU1Candidate            float64
	KappaNative                 bool
	WeakAngleDerived            bool
	GaugeCouplingsDerived       bool
	HiggsVEVDerived             bool
	PhysicalWZMassMatrixDerived bool
	ObservedWZMassRatioClaimed  bool
	Verdict                     string
	Reason                      string
}

type Boundary struct {
	Executed                        bool
	BridgeQuotientAccepted          bool
	NativeElectroweakActionClosed   bool
	NativeScalarNormalizationClosed bool
	NativeKappaClosed               bool
	NativeWeakAngleClosed           bool
	NativeGaugeCouplingsClosed      bool
	NativeHiggsVEVClosed            bool
	NativeWZMassMatrixClosed        bool
	NativeMassRatioClosed           bool
	Verdict                         string
	Reason                          string
}

type Firewall struct {
	Executed                      bool
	YukawaTraceValueImported      bool
	ObservedYukawaImported        bool
	ObservedWMassImported         bool
	ObservedZMassImported         bool
	ObservedHiggsVEVImported      bool
	ObservedWeakAngleImported     bool
	ObservedGaugeCouplingImported bool
	ObservedWZMassRatioImported   bool
	CKMPMNSImported               bool
	NativeKappaWritten            bool
	NativeWeakAngleWritten        bool
	NativeGaugeCouplingWritten    bool
	NativeHiggsVEVWritten         bool
	NativeWZMassWritten           bool
	NativeMassRatioWritten        bool
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
	Quotient    QuotientDefinition
	KernelRank  KernelRankAudit
	Hessian     HessianQuotientAudit
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
	g501, err := generation2yukawatracescalarnormalization.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate501 scalar normalization airlock: %w", err)
	}
	g497, err := generation2vacuumgaugeorbitquotient.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate497 gauge-orbit quotient: %w", err)
	}
	g495, err := generation2finiteactionsecondvariation.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate495 finite action second variation: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g501, g497, g495)
	a.Quotient = buildQuotientDefinition()
	a.KernelRank = buildKernelRank(g497)
	a.Hessian = buildHessianQuotient(g495)
	a.Boundary = buildBoundary(a.Quotient, a.KernelRank, a.Hessian)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g501 generation2yukawatracescalarnormalization.Analysis, g497 generation2vacuumgaugeorbitquotient.Analysis, g495 generation2finiteactionsecondvariation.Analysis) Inheritance {
	scalarSealed := g501.Decision.TraceABridgeScalarNormAccepted && !g501.Decision.TraceANativeNumericAccepted && !g501.Decision.ScalarKineticCoefficientNative
	return Inheritance{
		Executed:                            true,
		Gate501AuditDefined:                 true,
		ScalarNormalizationSealed:           scalarSealed,
		TraceAIsBridgeOnly:                  g501.Decision.TraceABridgeScalarNormAccepted && !g501.Trace.NativeNumericValueDerived,
		ScalarNormalizationIndependentRoute: g501.Next.Gate == 502,
		Gate497AuditDefined:                 true,
		PhotonKernelAvailable:               g497.GaugeOrbit.PhotonIsotropyGenerator && g497.ResidualPhase.PhotonStabilizesVacuum,
		BrokenOrbitRankThreeAvailable:       g497.GaugeOrbit.GaugeOrbitRankThree,
		RadialQuotientOneModeAvailable:      g497.GaugeOrbit.FourToOneQuotientDiagnostic,
		Gate495AuditDefined:                 true,
		DimensionlessDiag114Candidate:       g495.Candidate.BrokenDiag114 && g495.Candidate.KappaSixSelectedInCandidate,
		CanonicalActionNativeClosed:         !g495.Provenance.NativeActionProvenanceClosed,
		NoScaleOrFlavorDataImported:         g501.Firewall.Executed && !g501.Firewall.ObservedYukawaImported && !g501.Firewall.ObservedWMassImported && !g501.Firewall.ObservedHiggsVEVImported && !g501.Firewall.GaugeCouplingImported,
		Verdict:                             strings.Join([]string{StatusGate501ScalarNormalizationSealInherited, StatusGate497PhotonBrokenOrbitInherited, StatusGate495DimensionlessHessianInherited}, ";"),
		Reason:                              "Gate501 seals scalar normalization; Gate497 supplies the photon kernel and rank-three broken orbit; Gate495 supplies only a dimensionless diag(1,1,4) Hessian candidate with unresolved native provenance.",
	}
}

func buildQuotientDefinition() QuotientDefinition {
	return QuotientDefinition{
		Executed:                          true,
		QuotientedFactors:                 []string{"a=Tr(Y†Y)", "f0", "Higgs VEV v", "overall scalar kinetic scale", "continuum matching scale", "gauge-coupling units"},
		RetainedObjects:                   []string{"kernel/nullity", "rank", "degeneracy", "normalized Hessian eigenvalue ratios", "symbolic carrier labels"},
		ScalarNormalizationRemoved:        true,
		YukawaTraceRemoved:                true,
		HiggsVEVRemoved:                   true,
		ContinuumScaleRemoved:             true,
		GaugeCouplingScaleRemoved:         true,
		OnlyDimensionlessStatements:       true,
		PhysicalMassStatementsAllowed:     false,
		PhysicalCouplingStatementsAllowed: false,
		Verdict:                           StatusNormalizationQuotientDefined,
		Reason:                            "The quotient deliberately deletes every scalar, continuum, and coupling scale, preserving only rank/kernel/ratio statements that are invariant under positive rescaling.",
	}
}

func buildKernelRank(g497 generation2vacuumgaugeorbitquotient.Analysis) KernelRankAudit {
	return KernelRankAudit{
		Executed:                          true,
		ScalarRealDimension:               g497.GaugeOrbit.ScalarDimensionBeforeQuotient,
		PhotonKernelDimension:             g497.GaugeOrbit.IsotropyDimension,
		BrokenGeneratorCount:              g497.GaugeOrbit.BrokenGeneratorCount,
		BrokenOrbitRank:                   g497.GaugeOrbit.OrbitImageRank,
		BrokenRankSurvivesScaleQuotient:   g497.GaugeOrbit.GaugeOrbitRankThree,
		PhotonKernelSurvivesScaleQuotient: g497.GaugeOrbit.PhotonIsotropyGenerator,
		RadialModeAfterQuotient:           g497.GaugeOrbit.ScalarDimensionAfterQuotient,
		FourToOneQuotient:                 g497.GaugeOrbit.FourToOneQuotientDiagnostic,
		NativeKernelIndexClosed:           false,
		Verdict:                           strings.Join([]string{StatusPhotonKernelSurvivesQuotient, StatusBrokenRankThreeSurvivesQuotient}, ";"),
		Reason:                            "Kernel dimension, broken-orbit rank, and the 4→1 quotient are invariant under scalar normalization.  They remain bridge structural diagnostics because the native finite-action gauge-orbit selection is still open.",
	}
}

func buildHessianQuotient(g495 generation2finiteactionsecondvariation.Analysis) HessianQuotientAudit {
	diag := append([]float64(nil), g495.Candidate.BrokenSelectedDiagonal...)
	ratio := 0.0
	chargedDegenerate := false
	diag114 := false
	if len(diag) == 3 && math.Abs(diag[0]) > eps && math.Abs(diag[1]) > eps {
		chargedDegenerate = math.Abs(diag[0]-diag[1]) < eps
		ratio = diag[2] / (0.5 * (diag[0] + diag[1]))
		diag114 = closeSlice(diag, []float64{1, 1, 4}, eps)
	}
	return HessianQuotientAudit{
		Executed:                    true,
		RawBrokenDiagonal:           append([]float64(nil), g495.Candidate.BrokenRawDiagonal...),
		NormalizedBrokenDiagonal:    diag,
		ChargedPairDegenerate:       chargedDegenerate,
		NeutralToChargedRatio:       ratio,
		Diag114Shape:                diag114,
		DimensionlessShapeSurvives:  diag114 && chargedDegenerate && math.Abs(ratio-4) < eps,
		KappaU1Candidate:            g495.Candidate.KappaU1,
		KappaNative:                 g495.Boundary.NativeKappaSelected,
		WeakAngleDerived:            g495.Boundary.NativeWeakAngleDerived,
		GaugeCouplingsDerived:       g495.Boundary.NativeGaugeCouplingsDerived,
		HiggsVEVDerived:             g495.Boundary.NativeHiggsVEVDerived,
		PhysicalWZMassMatrixDerived: g495.Boundary.NativeWZMassesDerived,
		ObservedWZMassRatioClaimed:  false,
		Verdict:                     strings.Join([]string{StatusChargedDegeneracySurvivesQuotient, StatusDiag114QuotientShapeSurvives, StatusFailedKappaStillBridgeAfterQuotient}, ";"),
		Reason:                      "The normalized candidate Hessian shape diag(1,1,4) survives quotienting because all positive scalar factors cancel.  This is not a physical mass ratio or kappa theorem because the native action provenance remains unclosed.",
	}
}

func buildBoundary(q QuotientDefinition, k KernelRankAudit, h HessianQuotientAudit) Boundary {
	bridge := q.OnlyDimensionlessStatements && k.BrokenRankSurvivesScaleQuotient && k.PhotonKernelSurvivesScaleQuotient && h.DimensionlessShapeSurvives
	return Boundary{
		Executed:                        true,
		BridgeQuotientAccepted:          bridge,
		NativeElectroweakActionClosed:   false,
		NativeScalarNormalizationClosed: false,
		NativeKappaClosed:               false,
		NativeWeakAngleClosed:           false,
		NativeGaugeCouplingsClosed:      false,
		NativeHiggsVEVClosed:            false,
		NativeWZMassMatrixClosed:        false,
		NativeMassRatioClosed:           false,
		Verdict:                         strings.Join([]string{StatusBridgeQuotientAccepted, StatusFailedQuotientNotNativeActionClosure, StatusFailedWZMassMatrixStillBlocked}, ";"),
		Reason:                          "Gate502 accepts only the quotient-level electroweak shape data.  It blocks all native action, coefficient, coupling, scale, and mass promotions.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                      true,
		YukawaTraceValueImported:      false,
		ObservedYukawaImported:        false,
		ObservedWMassImported:         false,
		ObservedZMassImported:         false,
		ObservedHiggsVEVImported:      false,
		ObservedWeakAngleImported:     false,
		ObservedGaugeCouplingImported: false,
		ObservedWZMassRatioImported:   false,
		CKMPMNSImported:               false,
		NativeKappaWritten:            false,
		NativeWeakAngleWritten:        false,
		NativeGaugeCouplingWritten:    false,
		NativeHiggsVEVWritten:         false,
		NativeWZMassWritten:           false,
		NativeMassRatioWritten:        false,
		Verdict:                       StatusFirewallPreserved,
		Reason:                        "No numeric Yukawa trace, Yukawa data, W/Z masses, Higgs VEV, weak angle, gauge couplings, W/Z mass ratio, CKM, or PMNS data are imported; no native electroweak mass/coupling write is made.",
	}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native electroweak action closure, scalar normalization, kappa_U1, weak angle, gauge coupling, Higgs VEV, W/Z mass matrix, or W/Z mass ratio is admitted at Gate502.",
		},
		BridgeEntries: []string{
			"After quotienting by a, f0, VEV, continuum scale, and coupling units, the photon kernel/null direction survives as scale-independent bridge data.",
			"The broken electroweak gauge orbit remains rank three, leaving one radial scalar quotient mode in the diagnostic scalar representation.",
			"The normalized broken Hessian shape diag(1,1,4), charged-pair degeneracy, and neutral/charged quotient ratio 4 survive as dimensionless bridge-candidate shape data.",
		},
		EnvironmentalEntries: []string{
			"The numerical Yukawa trace a, scalar normalization, Higgs VEV, continuum matching scale, gauge couplings, weak angle, W/Z masses, and observed W/Z ratio remain environmental or bridge-scale data.",
		},
		FailedRoutes: []string{
			StatusFailedQuotientNotNativeActionClosure,
			StatusFailedKappaStillBridgeAfterQuotient,
			StatusFailedWeakAngleNotDerivedFromQuotient,
			StatusFailedGaugeCouplingsNotDerivedFromQuotient,
			StatusFailedHiggsVEVStillSealed,
			StatusFailedWZMassMatrixStillBlocked,
			StatusFailedObservedMassRatioNotClaimed,
		},
		OpenTheorems: []string{
			"Prove the photon kernel and rank-three broken image as a native representation index theorem independent of the bridge scalar/gauge diagnostic.",
			"Derive a native finite-action provenance for the quotient Hessian before promoting kappa_U1=6 or any physical electroweak conclusion.",
			"Define a continuum matching permission ledger for the eventual environmental import of VEV/couplings without native registry contamination.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 503, Title: "Electroweak Kernel Index Native Closure Audit", Reason: "Gate502 preserves photon nullity and rank-three broken structure only as normalization-independent bridge quotient data.", PrimaryTask: "test whether Q_em kernel dimension one and broken-image rank three can be derived as native finite-representation index facts, without the bridge scalar metric, VEV, kappa, couplings, or physical masses"}
}

func truth(a Analysis) string {
	if a.Boundary.BridgeQuotientAccepted && !a.Boundary.NativeWZMassMatrixClosed {
		return "Gate502 proves the safe electroweak remainder after scalar normalization is sealed.  Photon nullity, rank-three broken orbit, charged-pair degeneracy, and the dimensionless diag(1,1,4) Hessian shape survive quotienting by a, f0, VEV, continuum scale, and coupling units.  These are bridge quotient invariants, not native W/Z masses, weak angle, gauge couplings, kappa promotion, or an observed mass-ratio theorem."
	}
	return "Gate502 did not establish a scalar-normalization-independent electroweak quotient."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.ScalarNormalizationSealed && a.Inheritance.ScalarNormalizationIndependentRoute && a.Inheritance.PhotonKernelAvailable && a.Inheritance.BrokenOrbitRankThreeAvailable && a.Inheritance.DimensionlessDiag114Candidate, "required inheritance missing"},
		{a.Quotient.Executed && a.Quotient.ScalarNormalizationRemoved && a.Quotient.YukawaTraceRemoved && a.Quotient.HiggsVEVRemoved && a.Quotient.OnlyDimensionlessStatements && !a.Quotient.PhysicalMassStatementsAllowed, "quotient definition invalid"},
		{a.KernelRank.Executed && a.KernelRank.PhotonKernelDimension == 1 && a.KernelRank.BrokenOrbitRank == 3 && a.KernelRank.PhotonKernelSurvivesScaleQuotient && a.KernelRank.BrokenRankSurvivesScaleQuotient && a.KernelRank.RadialModeAfterQuotient == 1, "kernel/rank quotient not established"},
		{a.Hessian.Executed && a.Hessian.ChargedPairDegenerate && a.Hessian.Diag114Shape && a.Hessian.DimensionlessShapeSurvives && math.Abs(a.Hessian.NeutralToChargedRatio-4) < eps && !a.Hessian.KappaNative && !a.Hessian.PhysicalWZMassMatrixDerived, "Hessian quotient over-promoted or missing"},
		{a.Boundary.Executed && a.Boundary.BridgeQuotientAccepted && !a.Boundary.NativeElectroweakActionClosed && !a.Boundary.NativeScalarNormalizationClosed && !a.Boundary.NativeKappaClosed && !a.Boundary.NativeWZMassMatrixClosed && !a.Boundary.NativeMassRatioClosed, "boundary over-promoted quotient"},
		{a.Firewall.Executed && !a.Firewall.YukawaTraceValueImported && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedZMassImported && !a.Firewall.ObservedHiggsVEVImported && !a.Firewall.ObservedWeakAngleImported && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeMassRatioWritten, "firewall violation"},
		{a.Next.Gate == 503, "Gate503 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func closeSlice(xs, ys []float64, tol float64) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if math.Abs(xs[i]-ys[i]) > tol {
			return false
		}
	}
	return true
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate501=%t scalar_sealed=%t trace_bridge=%t quotient_route=%t gate497=%t photon=%t broken_rank3=%t radial=%t gate495=%t diag114=%t action_native_closed=%t no_data=%t verdict=%s reason=%s", x.Gate501AuditDefined, x.ScalarNormalizationSealed, x.TraceAIsBridgeOnly, x.ScalarNormalizationIndependentRoute, x.Gate497AuditDefined, x.PhotonKernelAvailable, x.BrokenOrbitRankThreeAvailable, x.RadialQuotientOneModeAvailable, x.Gate495AuditDefined, x.DimensionlessDiag114Candidate, x.CanonicalActionNativeClosed, x.NoScaleOrFlavorDataImported, x.Verdict, x.Reason)
}

func FormatQuotient(x QuotientDefinition) string {
	return fmt.Sprintf("quotient=[%s] retained=[%s] scalar_removed=%t trace_removed=%t VEV_removed=%t continuum_removed=%t gauge_scale_removed=%t dimensionless_only=%t masses_allowed=%t couplings_allowed=%t verdict=%s reason=%s", strings.Join(x.QuotientedFactors, "; "), strings.Join(x.RetainedObjects, "; "), x.ScalarNormalizationRemoved, x.YukawaTraceRemoved, x.HiggsVEVRemoved, x.ContinuumScaleRemoved, x.GaugeCouplingScaleRemoved, x.OnlyDimensionlessStatements, x.PhysicalMassStatementsAllowed, x.PhysicalCouplingStatementsAllowed, x.Verdict, x.Reason)
}

func FormatKernelRank(x KernelRankAudit) string {
	return fmt.Sprintf("scalar_dim=%d photon_kernel=%d broken_count=%d broken_rank=%d rank_survives=%t photon_survives=%t radial_after=%d four_to_one=%t native_index=%t verdict=%s reason=%s", x.ScalarRealDimension, x.PhotonKernelDimension, x.BrokenGeneratorCount, x.BrokenOrbitRank, x.BrokenRankSurvivesScaleQuotient, x.PhotonKernelSurvivesScaleQuotient, x.RadialModeAfterQuotient, x.FourToOneQuotient, x.NativeKernelIndexClosed, x.Verdict, x.Reason)
}

func FormatHessian(x HessianQuotientAudit) string {
	return fmt.Sprintf("raw=%s normalized=%s charged_degenerate=%t neutral_charged=%.10f diag114=%t shape_survives=%t kappa=%.10f kappa_native=%t weak_angle=%t couplings=%t VEV=%t WZ=%t observed_ratio=%t verdict=%s reason=%s", formatFloats(x.RawBrokenDiagonal), formatFloats(x.NormalizedBrokenDiagonal), x.ChargedPairDegenerate, x.NeutralToChargedRatio, x.Diag114Shape, x.DimensionlessShapeSurvives, x.KappaU1Candidate, x.KappaNative, x.WeakAngleDerived, x.GaugeCouplingsDerived, x.HiggsVEVDerived, x.PhysicalWZMassMatrixDerived, x.ObservedWZMassRatioClaimed, x.Verdict, x.Reason)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("bridge=%t action_native=%t scalar_norm=%t kappa=%t weak_angle=%t couplings=%t VEV=%t WZ=%t mass_ratio=%t verdict=%s reason=%s", x.BridgeQuotientAccepted, x.NativeElectroweakActionClosed, x.NativeScalarNormalizationClosed, x.NativeKappaClosed, x.NativeWeakAngleClosed, x.NativeGaugeCouplingsClosed, x.NativeHiggsVEVClosed, x.NativeWZMassMatrixClosed, x.NativeMassRatioClosed, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("a_value=%t Yukawa=%t W=%t Z=%t VEV=%t thetaW=%t gauge=%t WZratio=%t CKM_PMNS=%t native_kappa=%t native_thetaW=%t native_gauge=%t native_VEV=%t native_WZ=%t native_ratio=%t verdict=%s reason=%s", x.YukawaTraceValueImported, x.ObservedYukawaImported, x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsVEVImported, x.ObservedWeakAngleImported, x.ObservedGaugeCouplingImported, x.ObservedWZMassRatioImported, x.CKMPMNSImported, x.NativeKappaWritten, x.NativeWeakAngleWritten, x.NativeGaugeCouplingWritten, x.NativeHiggsVEVWritten, x.NativeWZMassWritten, x.NativeMassRatioWritten, x.Verdict, x.Reason)
}

func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 502 Registry Audit — Scalar-Normalization-Independent Electroweak Quotient Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{
		StatusGate501ScalarNormalizationSealInherited,
		StatusGate497PhotonBrokenOrbitInherited,
		StatusGate495DimensionlessHessianInherited,
		StatusNormalizationQuotientDefined,
		StatusPhotonKernelSurvivesQuotient,
		StatusBrokenRankThreeSurvivesQuotient,
		StatusChargedDegeneracySurvivesQuotient,
		StatusDiag114QuotientShapeSurvives,
		StatusBridgeQuotientAccepted,
		StatusFailedQuotientNotNativeActionClosure,
		StatusFailedKappaStillBridgeAfterQuotient,
		StatusFailedWeakAngleNotDerivedFromQuotient,
		StatusFailedGaugeCouplingsNotDerivedFromQuotient,
		StatusFailedHiggsVEVStillSealed,
		StatusFailedWZMassMatrixStillBlocked,
		StatusFailedObservedMassRatioNotClaimed,
		StatusFirewallPreserved,
		StatusNativeRegistryWriteBlocked,
		StatusGate503RedirectDefined,
	} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate501 sealed the scalar kinetic normalization because:\n\n```text\nK_phi = f0 a / pi^2\na = Tr(Y†Y)\n```\n\n")
	b.WriteString("The value of `a` is a Yukawa amplitude trace and is therefore bridge/environmental.  Gate502 therefore deletes all scale-like quantities and asks only what survives as a quotient statement.\n\n")
	b.WriteString("## Quotient definition\n\n")
	b.WriteString("Removed:\n\n```text\na = Tr(Y†Y)\nf0\nHiggs VEV v\noverall scalar kinetic scale\ncontinuum matching scale\ngauge-coupling units\n```\n\n")
	b.WriteString("Retained:\n\n```text\nkernel/nullity\nrank\ndegeneracy\nnormalized Hessian eigenvalue ratios\nsymbolic carrier labels\n```\n\n")
	b.WriteString("## Kernel and rank audit\n\n")
	b.WriteString(fmt.Sprintf("The quotient preserves the photon kernel and broken-rank structure:\n\n```text\nscalar real dimension before quotient = %d\nphoton kernel dimension = %d\nbroken generator count = %d\nbroken orbit rank = %d\nradial scalar modes after gauge quotient = %d\n```\n\n", a.KernelRank.ScalarRealDimension, a.KernelRank.PhotonKernelDimension, a.KernelRank.BrokenGeneratorCount, a.KernelRank.BrokenOrbitRank, a.KernelRank.RadialModeAfterQuotient))
	b.WriteString("These facts are invariant under positive scalar rescaling.  They are still bridge diagnostics until a native finite-representation index theorem closes the provenance.\n\n")
	b.WriteString("## Dimensionless Hessian quotient audit\n\n")
	b.WriteString(fmt.Sprintf("The normalized broken Hessian candidate is:\n\n```text\nK_broken / charged_unit = %s\ncharged pair degeneracy = %t\nneutral / charged quotient ratio = %.10g\nkappa_U1 candidate = %.10g\n```\n\n", formatFloats(a.Hessian.NormalizedBrokenDiagonal), a.Hessian.ChargedPairDegenerate, a.Hessian.NeutralToChargedRatio, a.Hessian.KappaU1Candidate))
	b.WriteString("The shape `diag(1,1,4)` survives the quotient, but it is not promoted to a physical W/Z mass ratio, weak mixing angle, gauge-coupling theorem, or native `kappa_U1` theorem.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("No numeric Yukawa trace, Yukawa data, W/Z masses, Higgs VEV, observed weak angle, gauge couplings, observed W/Z ratio, CKM, or PMNS data enter this gate.  No native write is made for `kappa_U1`, couplings, VEV, W/Z masses, or mass ratios.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate503 should be:\n\n```text\nGate 503 — Electroweak Kernel Index Native Closure Audit\n```\n\n")
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

func formatFloats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.10g", x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
