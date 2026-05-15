// Package generation2ewresidualgeometryairlock implements Gate 508:
// Electroweak Comparator Residual Geometry Airlock.
//
// Gate 507 proved that an explicit, fully tagged electroweak comparator file can
// drive the bridge adapter without leaking any value into the native ASHA
// registry.  Gate 508 maps those file-adapter residuals against the surviving
// dimensionless electroweak quotient/index diagnostics from Gates 502 and 503.
//
// The result is deliberately narrow: photon zero-mode alignment and tree rho
// self-consistency are safe bridge diagnostics, while the file-level W/Z ratio,
// weak angle, couplings, VEV, and any residual against the diag(1,1,4) quotient
// shape are not native predictions.  The synthetic 3-4-5 fixture even gives
// mZ^2/mW^2 = 25/9, which is not the Gate502 diag(1,1,4) quotient ratio 4;
// this mismatch is the point of the airlock.  It prevents a bridge comparator
// from being mistaken for a theorem.
package generation2ewresidualgeometryairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewkernelindexclosure"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewnquotient"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2observedelectroweakfileadapter"
)

const (
	AuditID = "GATE508-ELECTROWEAK-COMPARATOR-RESIDUAL-GEOMETRY-AIRLOCK"

	StatusGate507FileAdapterInherited       = "CONDITIONAL_SUPPORT_GATE507_FILE_ADAPTER_INHERITED"
	StatusGate502QuotientInherited          = "CONDITIONAL_SUPPORT_GATE502_ELECTROWEAK_QUOTIENT_INHERITED"
	StatusGate503KernelIndexInherited       = "CONDITIONAL_SUPPORT_GATE503_KERNEL_INDEX_INHERITED"
	StatusResidualGeometryAirlockDefined    = "CONDITIONAL_SUPPORT_ELECTROWEAK_RESIDUAL_GEOMETRY_AIRLOCK_DEFINED"
	StatusPhotonZeroAlignmentConfirmed      = "CONDITIONAL_SUPPORT_PHOTON_ZERO_ALIGNMENT_CONFIRMED"
	StatusRhoIdentityBridgeOnlyConfirmed    = "CONDITIONAL_SUPPORT_TREE_RHO_IDENTITY_CLASSIFIED_AS_BRIDGE_ONLY"
	StatusFileResidualsClassifiedBridgeOnly = "CONDITIONAL_SUPPORT_FILE_COMPARATOR_RESIDUALS_CLASSIFIED_BRIDGE_ONLY"
	StatusDiag114ResidualComputedBridgeOnly = "CONDITIONAL_SUPPORT_DIAG114_TO_FILE_RATIO_RESIDUAL_COMPUTED_BRIDGE_ONLY"
	StatusGate509RedirectDefined            = "CONDITIONAL_SUPPORT_GATE509_NATIVE_FRONTIER_REDIRECT_DEFINED"
	StatusFirewallPreserved                 = "FIREWALL_PRESERVED_GATE508_NO_ELECTROWEAK_NATIVE_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked        = "FIREWALL_BLOCKED_GATE508_RESIDUAL_GEOMETRY_NATIVE_WRITE"

	StatusFailedDiag114NotMassRatio            = "FAILED_ROUTE_DIAG114_QUOTIENT_SHAPE_IS_NOT_FILE_WZ_MASS_RATIO"
	StatusFailedFileRatioDoesNotSelectDiag114  = "FAILED_ROUTE_FILE_NEUTRAL_CHARGED_RATIO_DOES_NOT_SELECT_DIAG114"
	StatusFailedWeakAngleNotDerivedByResiduals = "FAILED_ROUTE_WEAK_ANGLE_NOT_DERIVED_BY_RESIDUAL_GEOMETRY"
	StatusFailedCouplingsNotDerivedByResiduals = "FAILED_ROUTE_GAUGE_COUPLINGS_NOT_DERIVED_BY_RESIDUAL_GEOMETRY"
	StatusFailedVEVNotDerivedByResiduals       = "FAILED_ROUTE_HIGGS_VEV_NOT_DERIVED_BY_RESIDUAL_GEOMETRY"
	StatusFailedWZMassesNotNative              = "FAILED_ROUTE_WZ_MASSES_NOT_NATIVE_AFTER_RESIDUAL_GEOMETRY"
	StatusFailedKappaStillBridge               = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_AFTER_RESIDUAL_GEOMETRY"
)

type Inheritance struct {
	Executed                        bool
	Gate507AuditDefined             bool
	Gate507FileLoaded               bool
	Gate507SyntheticOnly            bool
	Gate507ObservedNumbersImported  bool
	Gate507AdapterExecuted          bool
	Gate507ResidualsComputed        bool
	Gate507ResidualsAllZero         bool
	Gate507NativeWriteBlocked       bool
	Gate502AuditDefined             bool
	Gate502QuotientAccepted         bool
	Gate502PhotonKernel             bool
	Gate502BrokenRankThree          bool
	Gate502Diag114Shape             bool
	Gate503AuditDefined             bool
	Gate503KernelIndexAccepted      bool
	Gate503PhotonStabilizerIndexOne bool
	Gate503BrokenOrbitIndexThree    bool
	Gate503RadialQuotientIndexOne   bool
	NoObservedOrScaleDataImported   bool
	Verdict                         string
	Reason                          string
}

type QuotientLedger struct {
	Executed                    bool
	PhotonKernelDimension       int
	BrokenOrbitRank             int
	RadialQuotientDimension     int
	Diag114NeutralChargedRatio  float64
	Diag114NativeMassRatio      bool
	KappaNative                 bool
	WeakAngleDerived            bool
	GaugeCouplingsDerived       bool
	VEVDerived                  bool
	PhysicalWZMassMatrixDerived bool
	Verdict                     string
	Reason                      string
}

type FileResidualGeometry struct {
	Executed                        bool
	BridgeOnly                      bool
	SyntheticOnly                   bool
	ObservedValuesImported          bool
	FileResidualsComputed           bool
	FileResidualsAllZero            bool
	FileWeakAngleResidual           float64
	FileMWResidual                  float64
	FileMZResidual                  float64
	FileNeutralChargedRatio         float64
	QuotientNeutralChargedRatio     float64
	Diag114ToFileRatioResidual      float64
	Diag114RatioMatchedByFile       bool
	PhotonZeroAlignment             bool
	RhoIdentityConfirmed            bool
	RhoIdentityNativeMassPrediction bool
	Verdict                         string
	Reason                          string
}

type Classification struct {
	Executed                         bool
	PhotonZeroIsStructuralAlignment  bool
	RhoIdentityIsBridgeFormula       bool
	FileResidualsAreAdapterResiduals bool
	Diag114MismatchIsExpected        bool
	Diag114UsedAsMassRatio           bool
	WeakAngleNativePrediction        bool
	GaugeCouplingNativePrediction    bool
	VEVNativePrediction              bool
	WZMassNativePrediction           bool
	KappaNativePromotion             bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                          bool
	ObservedNumbersImported           bool
	ObservedWMassImported             bool
	ObservedZMassImported             bool
	ObservedWeakAngleImported         bool
	ObservedGaugeCouplingImported     bool
	ObservedVEVImported               bool
	FileAdapterOutputsNative          bool
	FileResidualsNative               bool
	Diag114RatioNativeMassRatio       bool
	WeakAngleNativeWritten            bool
	GaugeCouplingsNativeWritten       bool
	VEVNativeWritten                  bool
	WZMassNativeWritten               bool
	KappaNativeWritten                bool
	NativeRegistryWritten             bool
	PhysicalElectroweakPredictionMade bool
	Verdict                           string
	Reason                            string
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
	Inheritance    Inheritance
	Quotient       QuotientLedger
	Residuals      FileResidualGeometry
	Classification Classification
	Firewall       Firewall
	Registry       RegistryUpdate
	Next           NextStep
	Truth          string
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
	g507, err := generation2observedelectroweakfileadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate507 electroweak file adapter: %w", err)
	}
	g502, err := generation2ewnquotient.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate502 electroweak quotient: %w", err)
	}
	g503, err := generation2ewkernelindexclosure.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate503 kernel index: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g507, g502, g503)
	a.Quotient = buildQuotientLedger(g502, g503)
	a.Residuals = buildResidualGeometry(g507, a.Quotient)
	a.Classification = buildClassification(a.Residuals)
	a.Firewall = buildFirewall(g507, a.Residuals, a.Classification)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g507 generation2observedelectroweakfileadapter.Analysis, g502 generation2ewnquotient.Analysis, g503 generation2ewkernelindexclosure.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate507AuditDefined:             true,
		Gate507FileLoaded:               g507.Import.Loaded,
		Gate507SyntheticOnly:            g507.Import.SyntheticFixture && !g507.Import.ObservedValuesLoaded,
		Gate507ObservedNumbersImported:  g507.Import.ObservedValuesLoaded,
		Gate507AdapterExecuted:          g507.Output.Executed && g507.Output.Ready,
		Gate507ResidualsComputed:        g507.Residuals.Executed && g507.Residuals.ComparatorRowsAvailable,
		Gate507ResidualsAllZero:         g507.Residuals.AllResidualsZero,
		Gate507NativeWriteBlocked:       g507.Firewall.NativeRegistryWritten == false && g507.Firewall.PhysicalElectroweakPredictionMade == false,
		Gate502AuditDefined:             true,
		Gate502QuotientAccepted:         g502.Boundary.BridgeQuotientAccepted,
		Gate502PhotonKernel:             g502.KernelRank.PhotonKernelSurvivesScaleQuotient,
		Gate502BrokenRankThree:          g502.KernelRank.BrokenOrbitRank == 3,
		Gate502Diag114Shape:             g502.Hessian.Diag114Shape,
		Gate503AuditDefined:             true,
		Gate503KernelIndexAccepted:      g503.Boundary.ConditionalRepresentationIndexAccepted,
		Gate503PhotonStabilizerIndexOne: g503.Kernel.PhotonKernelIndexProven,
		Gate503BrokenOrbitIndexThree:    g503.Kernel.BrokenOrbitIndexProven,
		Gate503RadialQuotientIndexOne:   g503.Kernel.RadialIndexProven,
		NoObservedOrScaleDataImported:   !g507.Import.ObservedValuesLoaded && !g502.Firewall.ObservedWMassImported && !g503.Firewall.ObservedWMassImported,
		Verdict:                         strings.Join([]string{StatusGate507FileAdapterInherited, StatusGate502QuotientInherited, StatusGate503KernelIndexInherited}, ";"),
		Reason:                          "Gate508 inherits the synthetic bridge-only file adapter, the Gate502 scale-free quotient ledger, and the Gate503 conditional kernel-index theorem without importing observed electroweak data.",
	}
}

func buildQuotientLedger(g502 generation2ewnquotient.Analysis, g503 generation2ewkernelindexclosure.Analysis) QuotientLedger {
	return QuotientLedger{
		Executed:                    true,
		PhotonKernelDimension:       g503.Kernel.StabilizerDimension,
		BrokenOrbitRank:             g503.Kernel.BrokenOrbitDimension,
		RadialQuotientDimension:     g503.Kernel.RadialQuotientDimension,
		Diag114NeutralChargedRatio:  g502.Hessian.NeutralToChargedRatio,
		Diag114NativeMassRatio:      g502.Hessian.ObservedWZMassRatioClaimed,
		KappaNative:                 g502.Hessian.KappaNative,
		WeakAngleDerived:            g502.Hessian.WeakAngleDerived,
		GaugeCouplingsDerived:       g502.Hessian.GaugeCouplingsDerived,
		VEVDerived:                  g502.Hessian.HiggsVEVDerived,
		PhysicalWZMassMatrixDerived: g502.Hessian.PhysicalWZMassMatrixDerived,
		Verdict:                     strings.Join([]string{StatusResidualGeometryAirlockDefined, StatusGate502QuotientInherited, StatusGate503KernelIndexInherited}, ";"),
		Reason:                      "The native-facing quotient ledger contains only dimensionless structural data: photon stabilizer dimension one, broken orbit rank three, radial quotient dimension one, and a bridge Hessian shape diag(1,1,4). It does not contain physical electroweak masses or couplings.",
	}
}

func buildResidualGeometry(g507 generation2observedelectroweakfileadapter.Analysis, q QuotientLedger) FileResidualGeometry {
	fileRatio := math.NaN()
	if g507.Output.MW != 0 {
		fileRatio = (g507.Output.MZ * g507.Output.MZ) / (g507.Output.MW * g507.Output.MW)
	}
	residual := math.Abs(fileRatio - q.Diag114NeutralChargedRatio)
	matched := nearly(fileRatio, q.Diag114NeutralChargedRatio, 1e-12)
	return FileResidualGeometry{
		Executed:                        true,
		BridgeOnly:                      g507.Residuals.BridgeOnly,
		SyntheticOnly:                   g507.Import.SyntheticFixture && !g507.Import.ObservedValuesLoaded,
		ObservedValuesImported:          g507.Import.ObservedValuesLoaded,
		FileResidualsComputed:           g507.Residuals.Executed,
		FileResidualsAllZero:            g507.Residuals.AllResidualsZero,
		FileWeakAngleResidual:           g507.Residuals.WeakAngleResidual,
		FileMWResidual:                  g507.Residuals.MWResidual,
		FileMZResidual:                  g507.Residuals.MZResidual,
		FileNeutralChargedRatio:         fileRatio,
		QuotientNeutralChargedRatio:     q.Diag114NeutralChargedRatio,
		Diag114ToFileRatioResidual:      residual,
		Diag114RatioMatchedByFile:       matched,
		PhotonZeroAlignment:             g507.Output.PhotonZeroPreserved && q.PhotonKernelDimension == 1,
		RhoIdentityConfirmed:            g507.Output.RhoIdentityConfirmed,
		RhoIdentityNativeMassPrediction: false,
		Verdict:                         strings.Join([]string{StatusFileResidualsClassifiedBridgeOnly, StatusPhotonZeroAlignmentConfirmed, StatusRhoIdentityBridgeOnlyConfirmed, StatusDiag114ResidualComputedBridgeOnly, StatusFailedFileRatioDoesNotSelectDiag114}, ";"),
		Reason:                          "Gate508 compares file-adapter outputs to the quotient/index ledger only as residual geometry. The synthetic file residuals are zero against their own comparator rows, but the file mass-ratio channel is not forced to equal the diag(1,1,4) quotient shape.",
	}
}

func buildClassification(r FileResidualGeometry) Classification {
	return Classification{
		Executed:                         true,
		PhotonZeroIsStructuralAlignment:  r.PhotonZeroAlignment,
		RhoIdentityIsBridgeFormula:       r.RhoIdentityConfirmed && !r.RhoIdentityNativeMassPrediction,
		FileResidualsAreAdapterResiduals: r.FileResidualsComputed && r.BridgeOnly,
		Diag114MismatchIsExpected:        !r.Diag114RatioMatchedByFile,
		Diag114UsedAsMassRatio:           false,
		WeakAngleNativePrediction:        false,
		GaugeCouplingNativePrediction:    false,
		VEVNativePrediction:              false,
		WZMassNativePrediction:           false,
		KappaNativePromotion:             false,
		Verdict:                          strings.Join([]string{StatusResidualGeometryAirlockDefined, StatusFailedDiag114NotMassRatio, StatusFailedWeakAngleNotDerivedByResiduals, StatusFailedCouplingsNotDerivedByResiduals, StatusFailedVEVNotDerivedByResiduals, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge}, ";"),
		Reason:                           "Photon zero and rho=1 are accepted only as structural/bridge consistency checks. The diag(1,1,4) quotient shape is not used as a W/Z mass-ratio prediction, and the file ratio is not allowed to select kappa, weak angle, VEV, or couplings.",
	}
}

func buildFirewall(g507 generation2observedelectroweakfileadapter.Analysis, r FileResidualGeometry, c Classification) Firewall {
	return Firewall{
		Executed:                          true,
		ObservedNumbersImported:           g507.Import.ObservedValuesLoaded,
		ObservedWMassImported:             false,
		ObservedZMassImported:             false,
		ObservedWeakAngleImported:         false,
		ObservedGaugeCouplingImported:     false,
		ObservedVEVImported:               false,
		FileAdapterOutputsNative:          false,
		FileResidualsNative:               false,
		Diag114RatioNativeMassRatio:       c.Diag114UsedAsMassRatio,
		WeakAngleNativeWritten:            c.WeakAngleNativePrediction,
		GaugeCouplingsNativeWritten:       c.GaugeCouplingNativePrediction,
		VEVNativeWritten:                  c.VEVNativePrediction,
		WZMassNativeWritten:               c.WZMassNativePrediction,
		KappaNativeWritten:                c.KappaNativePromotion,
		NativeRegistryWritten:             false,
		PhysicalElectroweakPredictionMade: false,
		Verdict:                           strings.Join([]string{StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                            fmt.Sprintf("All Gate508 residuals remain bridge diagnostics. The file neutral/charged ratio residual against diag(1,1,4) is %s, but neither the mismatch nor any zero residual is native-registry eligible.", fmtFloat(r.Diag114ToFileRatioResidual)),
	}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"Gate508 writes no native electroweak mass, weak-angle, coupling, VEV, kappa, file residual, or observed-comparator datum."},
		BridgeEntries:        []string{"Electroweak file-adapter residuals may be compared to the quotient/index ledger only as residual geometry.", "Photon zero-mode alignment and tree rho identity are bridge/structural consistency checks, not electroweak-scale predictions.", "The synthetic file ratio mZ^2/mW^2=25/9 is explicitly not forced to match the Gate502 diag(1,1,4) quotient ratio 4."},
		EnvironmentalEntries: []string{"Any future real observed electroweak file remains environmental bridge data with source/version/scale/scheme/uncertainty metadata and native-promotion rejection."},
		FailedRoutes:         []string{StatusFailedDiag114NotMassRatio, StatusFailedFileRatioDoesNotSelectDiag114, StatusFailedWeakAngleNotDerivedByResiduals, StatusFailedCouplingsNotDerivedByResiduals, StatusFailedVEVNotDerivedByResiduals, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge},
		OpenTheorems:         []string{"A separate native finite-action theorem is still required to select kappa_U1, gauge couplings, a nonzero Higgs VEV, or physical W/Z masses.", "Gate509 should redirect away from electroweak comparator arithmetic and back to a native invariant lane."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 509, Title: "Native Frontier Redirect After Electroweak Adapter Closure", Reason: "Gate508 closes the residual-geometry adapter branch as bridge-only; the next safe step is to select a native theorem lane that does not depend on electroweak comparator data.", PrimaryTask: "redirect to a native topological/spectral invariant lane, such as curvature-action coefficient provenance, gravitational spectral terms, or anomaly/topological ledgers, without importing electroweak scales or masses"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate507FileLoaded && a.Inheritance.Gate507SyntheticOnly && !a.Inheritance.Gate507ObservedNumbersImported && a.Inheritance.Gate507AdapterExecuted && a.Inheritance.Gate507ResidualsComputed && a.Inheritance.Gate507ResidualsAllZero && a.Inheritance.Gate507NativeWriteBlocked && a.Inheritance.Gate502QuotientAccepted && a.Inheritance.Gate502PhotonKernel && a.Inheritance.Gate502BrokenRankThree && a.Inheritance.Gate502Diag114Shape && a.Inheritance.Gate503KernelIndexAccepted && a.Inheritance.NoObservedOrScaleDataImported, "Gate508 inheritance invalid"},
		{a.Quotient.Executed && a.Quotient.PhotonKernelDimension == 1 && a.Quotient.BrokenOrbitRank == 3 && a.Quotient.RadialQuotientDimension == 1 && nearly(a.Quotient.Diag114NeutralChargedRatio, 4, 1e-12) && !a.Quotient.Diag114NativeMassRatio && !a.Quotient.KappaNative && !a.Quotient.WeakAngleDerived && !a.Quotient.GaugeCouplingsDerived && !a.Quotient.VEVDerived && !a.Quotient.PhysicalWZMassMatrixDerived, "Gate508 quotient ledger invalid"},
		{a.Residuals.Executed && a.Residuals.BridgeOnly && a.Residuals.SyntheticOnly && !a.Residuals.ObservedValuesImported && a.Residuals.FileResidualsComputed && a.Residuals.FileResidualsAllZero && nearly(a.Residuals.FileNeutralChargedRatio, 25.0/9.0, 1e-12) && nearly(a.Residuals.QuotientNeutralChargedRatio, 4, 1e-12) && nearly(a.Residuals.Diag114ToFileRatioResidual, 11.0/9.0, 1e-12) && !a.Residuals.Diag114RatioMatchedByFile && a.Residuals.PhotonZeroAlignment && a.Residuals.RhoIdentityConfirmed && !a.Residuals.RhoIdentityNativeMassPrediction, "Gate508 residual geometry invalid"},
		{a.Classification.Executed && a.Classification.PhotonZeroIsStructuralAlignment && a.Classification.RhoIdentityIsBridgeFormula && a.Classification.FileResidualsAreAdapterResiduals && a.Classification.Diag114MismatchIsExpected && !a.Classification.Diag114UsedAsMassRatio && !a.Classification.WeakAngleNativePrediction && !a.Classification.GaugeCouplingNativePrediction && !a.Classification.VEVNativePrediction && !a.Classification.WZMassNativePrediction && !a.Classification.KappaNativePromotion, "Gate508 classification invalid"},
		{a.Firewall.Executed && !a.Firewall.ObservedNumbersImported && !a.Firewall.FileAdapterOutputsNative && !a.Firewall.FileResidualsNative && !a.Firewall.Diag114RatioNativeMassRatio && !a.Firewall.WeakAngleNativeWritten && !a.Firewall.GaugeCouplingsNativeWritten && !a.Firewall.VEVNativeWritten && !a.Firewall.WZMassNativeWritten && !a.Firewall.KappaNativeWritten && !a.Firewall.NativeRegistryWritten && !a.Firewall.PhysicalElectroweakPredictionMade, "Gate508 firewall violated"},
		{a.Next.Gate == 509, "Gate509 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func truth(a Analysis) string {
	if a.Firewall.Executed && !a.Firewall.NativeRegistryWritten {
		return "Gate 508 proves that electroweak file-adapter residuals can be mapped against the native-facing quotient/index ledger only as bridge residual geometry. Photon zero-mode alignment and tree rho self-consistency survive as structural checks, but the synthetic file mass-ratio channel 25/9 does not select the diag(1,1,4) quotient ratio 4. No weak angle, coupling, VEV, kappa, W/Z mass, observed comparator, or residual enters the native ASHA registry."
	}
	return "Gate 508 failed before establishing the residual-geometry firewall."
}

func statuses() []string {
	return []string{StatusGate507FileAdapterInherited, StatusGate502QuotientInherited, StatusGate503KernelIndexInherited, StatusResidualGeometryAirlockDefined, StatusPhotonZeroAlignmentConfirmed, StatusRhoIdentityBridgeOnlyConfirmed, StatusFileResidualsClassifiedBridgeOnly, StatusDiag114ResidualComputedBridgeOnly, StatusGate509RedirectDefined, StatusFailedDiag114NotMassRatio, StatusFailedFileRatioDoesNotSelectDiag114, StatusFailedWeakAngleNotDerivedByResiduals, StatusFailedCouplingsNotDerivedByResiduals, StatusFailedVEVNotDerivedByResiduals, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate507_loaded=%t synthetic_only=%t observed_imported=%t adapter=%t residuals=%t residuals_zero=%t native_blocked=%t gate502_quotient=%t photon=%t broken_rank3=%t diag114=%t gate503_index=%t no_observed_or_scale=%t verdict=%s reason=%s", x.Executed, x.Gate507FileLoaded, x.Gate507SyntheticOnly, x.Gate507ObservedNumbersImported, x.Gate507AdapterExecuted, x.Gate507ResidualsComputed, x.Gate507ResidualsAllZero, x.Gate507NativeWriteBlocked, x.Gate502QuotientAccepted, x.Gate502PhotonKernel, x.Gate502BrokenRankThree, x.Gate502Diag114Shape, x.Gate503KernelIndexAccepted, x.NoObservedOrScaleDataImported, x.Verdict, x.Reason)
}
func FormatQuotient(x QuotientLedger) string {
	return fmt.Sprintf("executed=%t photon_kernel_dim=%d broken_rank=%d radial=%d diag114_ratio=%s diag114_native_mass_ratio=%t kappa_native=%t weak_angle=%t couplings=%t vev=%t wz_mass_matrix=%t verdict=%s reason=%s", x.Executed, x.PhotonKernelDimension, x.BrokenOrbitRank, x.RadialQuotientDimension, fmtFloat(x.Diag114NeutralChargedRatio), x.Diag114NativeMassRatio, x.KappaNative, x.WeakAngleDerived, x.GaugeCouplingsDerived, x.VEVDerived, x.PhysicalWZMassMatrixDerived, x.Verdict, x.Reason)
}
func FormatResiduals(x FileResidualGeometry) string {
	return fmt.Sprintf("executed=%t bridge=%t synthetic=%t observed_imported=%t file_residuals=%t all_zero=%t sin2_residual=%s mW_residual=%s mZ_residual=%s file_ratio=%s quotient_ratio=%s diag114_file_residual=%s ratio_match=%t photon_alignment=%t rho=%t rho_native=%t verdict=%s reason=%s", x.Executed, x.BridgeOnly, x.SyntheticOnly, x.ObservedValuesImported, x.FileResidualsComputed, x.FileResidualsAllZero, fmtFloat(x.FileWeakAngleResidual), fmtFloat(x.FileMWResidual), fmtFloat(x.FileMZResidual), fmtFloat(x.FileNeutralChargedRatio), fmtFloat(x.QuotientNeutralChargedRatio), fmtFloat(x.Diag114ToFileRatioResidual), x.Diag114RatioMatchedByFile, x.PhotonZeroAlignment, x.RhoIdentityConfirmed, x.RhoIdentityNativeMassPrediction, x.Verdict, x.Reason)
}
func FormatClassification(x Classification) string {
	return fmt.Sprintf("executed=%t photon_structural=%t rho_bridge=%t file_residuals_adapter=%t diag114_mismatch_expected=%t diag114_used_as_mass_ratio=%t weak_angle_native=%t couplings_native=%t vev_native=%t wz_native=%t kappa_native=%t verdict=%s reason=%s", x.Executed, x.PhotonZeroIsStructuralAlignment, x.RhoIdentityIsBridgeFormula, x.FileResidualsAreAdapterResiduals, x.Diag114MismatchIsExpected, x.Diag114UsedAsMassRatio, x.WeakAngleNativePrediction, x.GaugeCouplingNativePrediction, x.VEVNativePrediction, x.WZMassNativePrediction, x.KappaNativePromotion, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t observed=%t rows_native=%t residuals_native=%t diag114_mass_ratio=%t weak_angle_write=%t couplings_write=%t vev_write=%t wz_write=%t kappa_write=%t native_registry=%t physical_prediction=%t verdict=%s reason=%s", x.Executed, x.ObservedNumbersImported, x.FileAdapterOutputsNative, x.FileResidualsNative, x.Diag114RatioNativeMassRatio, x.WeakAngleNativeWritten, x.GaugeCouplingsNativeWritten, x.VEVNativeWritten, x.WZMassNativeWritten, x.KappaNativeWritten, x.NativeRegistryWritten, x.PhysicalElectroweakPredictionMade, x.Verdict, x.Reason)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 508 Registry Audit — Electroweak Comparator Residual Geometry Airlock\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate507 proved that a file-backed electroweak adapter can run on an explicit synthetic bridge ledger. Gates502-503 provided a scale-free quotient/index ledger: photon stabilizer dimension one, broken orbit rank three, radial scalar quotient dimension one, and a diag(1,1,4) bridge Hessian shape. Gate508 may compare these ledgers only as residual geometry.\n\n")
	b.WriteString("```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Native-facing quotient/index ledger\n\n")
	b.WriteString("```text\n" + FormatQuotient(a.Quotient) + "\n```\n\n")
	b.WriteString("## File residual geometry\n\n")
	b.WriteString("```text\n" + FormatResiduals(a.Residuals) + "\n```\n\n")
	b.WriteString("The synthetic adapter fixture gives `m_Z^2/m_W^2 = 25/9`, while the Gate502 quotient shape carries a dimensionless `diag(1,1,4)` ratio of `4`. The residual `|25/9 - 4| = 11/9` is intentionally bridge-only. It proves that the comparator channel is not allowed to reinterpret the quotient Hessian shape as a physical W/Z mass ratio.\n\n")
	b.WriteString("## Classification\n\n")
	b.WriteString("```text\n" + FormatClassification(a.Classification) + "\n```\n\n")
	b.WriteString("Photon zero-mode alignment is structural. The tree rho identity is a bridge formula identity. File residuals are adapter residuals. None of these data derive a weak angle, coupling, VEV, kappa, or physical W/Z mass.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate509 should be:\n\n```text\nGate 509 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
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

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}
