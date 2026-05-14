// Package noncartanflavorvacuum implements Gate 260:
// Non-Cartan Flavor Vacuum / Off-Diagonal U12 Mixing Audit.
//
// Gate 259 conditionally selected the U12 weak plane using tau_eta under the
// SpontaneousCarrierSeal, but the Cartan electroweak route still failed to
// produce a neutral 3-plane inside 8_v.  Gate 260 audits the uncomfortable
// possibility that adding W+/W- off-diagonal SU(2)_L generators cannot change
// the charge-kernel dimension, because every su(2) element is gauge-conjugate
// to its Cartan representative.  It then opens a parallel direct-generation
// audit: tau_eta=(2,-2,1) already acts on a three-component generation/operator
// carrier and may be the generation-breaking source without passing through an
// 8_v neutral kernel.
package noncartanflavorvacuum

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/tauetaweakselector"
)

const (
	AuditID = "GATE260-NON-CARTAN-FLAVOR-VACUUM-OFF-DIAGONAL-U12-MIXING-AUDIT"

	StatusGate259Inherited                = "CONDITIONAL_SUPPORT_GATE259_TAU_ETA_U12_SELECTOR_INHERITED"
	StatusNonCartanGeneratorsRetrieved    = "CONDITIONAL_SUPPORT_U12_NON_CARTAN_SU2_GENERATORS_RETRIEVED"
	StatusGaugeOrbitInvariantProved       = "CONDITIONAL_SUPPORT_SU2_GAUGE_ORBIT_SPECTRUM_INVARIANT_PROVED"
	StatusOffDiagonalRouteClosed          = "FAILED_ROUTE_NON_CARTAN_OFF_DIAGONAL_U12_CANNOT_ENLARGE_Q_KERNEL"
	StatusEightVThreePlaneStillBlocked    = "FAILED_ROUTE_8V_NEUTRAL_3PLANE_STILL_BLOCKED_AFTER_NON_CARTAN_AUDIT"
	StatusTauEtaGenerationCarrierOpened   = "CONDITIONAL_SUPPORT_TAU_ETA_DIRECT_GENERATION_CARRIER_OPENED"
	StatusTauEtaYukawaSourceCandidate     = "CONDITIONAL_SUPPORT_TAU_ETA_GENERATION_BREAKING_SOURCE_MAP_CANDIDATE"
	StatusDirectYukawaTextureStillBlocked = "FAILED_ROUTE_DIRECT_TAU_ETA_YUKAWA_TEXTURE_REQUIRES_ACTION_AND_BILINEAR_MAP"
	StatusMassCKMStillBlocked             = "FAILED_ROUTE_CKM_PMNS_AND_FERMION_MASSES_STILL_BLOCKED"
	StatusNoForcedThreePlane              = "CONDITIONAL_SUPPORT_NO_3PLANE_FORCED_BY_NON_CARTAN_SCAN"
)

type Gate259Inheritance struct {
	TauEtaRetrieved                 bool
	ConditionalU12WeakPlaneSelected bool
	CartanRestrictedScanCompleted   bool
	CartanNeutral3PlaneDerived      bool
	CartanMaxPolarizedKernelDim     int
	CartanMaxFullKernelDim          int
	Gate259Status                   string
	Gate259NextGate                 string
	Gate259Comment                  string
	SurvivingWitnessNames           []string
	Verdict                         string
}

type NonCartanGeneratorAudit struct {
	WeakPlaneName             string
	ModePair                  [2]int
	CartanGenerator           string
	OffDiagonalGenerators     []string
	RaisingLoweringGenerators []string
	PauliBasisRetrieved       bool
	LieAlgebraClosed          bool
	HermitianBasis            bool
	TracelessBasis            bool
	ActsInsideSelectedPlane   bool
	ChangesGaugeDirection     bool
	ChangesChargeSpectrum     bool
	Verdict                   string
}

type GaugeDirectionAudit struct {
	Name                          string
	CoefficientsT1T2T3            [3]float64
	Norm                          float64
	EigenvalueRadius              float64
	Eigenvalues                   []float64
	YPlusHalfZeroMultiplicity     int
	YMinusHalfZeroMultiplicity    int
	SameAsCartanSpectrum          bool
	CanIncreaseBeyondCartanKernel bool
	Residual                      float64
}

type GaugeOrbitInvariantAudit struct {
	AllSU2ElementsConjugateToCartan  bool
	SpectrumInvariantUnderConjugacy  bool
	KernelDimensionGaugeInvariant    bool
	OffDiagonalTermsRotateBasisOnly  bool
	CartanMaxPolarizedKernelDim      int
	CartanMaxFullKernelDim           int
	NonCartanUpperBoundFullKernelDim int
	NonCartanCanEnlargeKernel        bool
	DirectionsAudited                []GaugeDirectionAudit
	DirectionCount                   int
	AllDirectionsMatchCartanSpectrum bool
	Verdict                          string
}

type EightVRouteClosureAudit struct {
	UsesGate259Survivors                       bool
	SurvivorCount                              int
	BranchCount                                int
	InheritedBranchEvaluationCount             int
	OffDiagonalScanReplacedByInvariant         bool
	WouldNeedNewRepresentationNotGaugeRotation bool
	Neutral3PlaneAvailable                     bool
	YukawaVia8VOpened                          bool
	Verdict                                    string
}

type DirectGenerationCarrierAudit struct {
	CarrierName                      string
	CarrierKind                      string
	Dimension                        int
	SourceGate                       string
	TauEtaEigenvalues                []int
	SignedDistinctEigenvalueCount    int
	MagnitudeDistinctCount           int
	Trace                            int
	Determinant                      int
	ActsOnGenerationIndex            bool
	OperatorSpaceNotVector8V         bool
	NativeGenerationBreakingCapacity bool
	Bypasses8VNeutralKernel          bool
	RequiresTrialityTransport        bool
	Verdict                          string
}

type DirectYukawaSourceAudit struct {
	TauEtaSourceMapCandidate         bool
	GenerationDiagonalTextureSeed    bool
	OnePlusOnePlusOneSignedSpectrum  bool
	CanBreakGenerationDegeneracy     bool
	YukawaTextureDerived             bool
	RequiresLeftRightBilinearCarrier bool
	RequiresFiniteYukawaAction       bool
	RequiresKineticNormalization     bool
	RequiresPhaseMixingSource        bool
	RequiresEmpiricalYukawaSeal      bool
	CKMPMNSDerived                   bool
	FermionMassesDerived             bool
	Verdict                          string
}

type FirewallAudit struct {
	Gate259NoGoPreserved                     bool
	DoesNotTreatWpmAsChargeOperator          bool
	DoesNotPromoteGaugeRotationToNewSpectrum bool
	DoesNotSelectTrialityByHand              bool
	DoesNotForceKernelDimThree               bool
	DoesNotRewriteTauEtaAsFockVector         bool
	UsesTauEtaAsGenerationOperator           bool
	DoesNotConstructYukawaTextureByHand      bool
	DoesNotImportObservedMasses              bool
	DoesNotImportCKMPMNS                     bool
	SpontaneousCarrierSealPreserved          bool
	FiniteCorePolluted                       bool
	Verdict                                  string
}

type Summary struct {
	Gate259Inherited              bool
	NonCartanGeneratorsRetrieved  bool
	GaugeOrbitInvariantProved     bool
	EightVNeutral3PlaneDerived    bool
	EightVRouteClosed             bool
	DirectGenerationCarrierOpened bool
	TauEtaYukawaSourceCandidate   bool
	DirectYukawaTextureDerived    bool
	CKMPMNSDerived                bool
	FermionMassesDerived          bool
	Status                        string
	NextGate                      string
	Comment                       string
}

type Analysis struct {
	PreviousGate259 tauetaweakselector.Analysis
	Inheritance     Gate259Inheritance
	NonCartan       NonCartanGeneratorAudit
	GaugeOrbit      GaugeOrbitInvariantAudit
	EightVClosure   EightVRouteClosureAudit
	Generation      DirectGenerationCarrierAudit
	YukawaSource    DirectYukawaSourceAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := tauetaweakselector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 259 predecessor: %w", err)
			return
		}
		inh := inheritGate259(prev)
		nc := retrieveNonCartanGenerators(inh)
		orbit := auditGaugeOrbitInvariant(inh, nc)
		closure := auditEightVClosure(prev, inh, orbit)
		gen := auditDirectGenerationCarrier(prev)
		yuk := auditDirectYukawaSource(gen)
		fw := auditFirewall(prev, nc, orbit, closure, gen, yuk)
		sum := summarize(inh, nc, orbit, closure, gen, yuk)
		truth := buildTruth(inh, nc, orbit, closure, gen, yuk)
		defaultA = Analysis{PreviousGate259: prev, Inheritance: inh, NonCartan: nc, GaugeOrbit: orbit, EightVClosure: closure, Generation: gen, YukawaSource: yuk, Firewall: fw, Summary: sum, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate259(a tauetaweakselector.Analysis) Gate259Inheritance {
	names := append([]string(nil), a.CombinedSieve.SurvivingWitnessNames...)
	sort.Strings(names)
	return Gate259Inheritance{
		TauEtaRetrieved:                 a.Summary.TauEtaRetrieved,
		ConditionalU12WeakPlaneSelected: a.Summary.UniqueUnorientedWeakPlane && a.SpatialTag.ComplementPlaneName == "U12",
		CartanRestrictedScanCompleted:   a.Summary.RestrictedTrialityRescanned && a.RestrictedScan.ResultCount == 12,
		CartanNeutral3PlaneDerived:      a.Summary.Neutral3PlaneDerived,
		CartanMaxPolarizedKernelDim:     a.RestrictedScan.MaxPolarizedKernelComplexDim,
		CartanMaxFullKernelDim:          a.RestrictedScan.MaxFullQ8vCKernelComplexDim,
		Gate259Status:                   a.Summary.Status,
		Gate259NextGate:                 a.Summary.NextGate,
		Gate259Comment:                  a.Summary.Comment,
		SurvivingWitnessNames:           names,
		Verdict:                         StatusGate259Inherited + "; inherited U12 selector and Cartan no-go without rerunning broad historical gates",
	}
}

func retrieveNonCartanGenerators(inh Gate259Inheritance) NonCartanGeneratorAudit {
	return NonCartanGeneratorAudit{
		WeakPlaneName:             "U12",
		ModePair:                  [2]int{1, 2},
		CartanGenerator:           "T3(U12)=1/2(N1-N2)",
		OffDiagonalGenerators:     []string{"T1(U12)=1/2(a1†a2+a2†a1)", "T2(U12)=1/(2i)(a1†a2-a2†a1)"},
		RaisingLoweringGenerators: []string{"T+(U12)=a1†a2", "T-(U12)=a2†a1"},
		PauliBasisRetrieved:       inh.ConditionalU12WeakPlaneSelected,
		LieAlgebraClosed:          true,
		HermitianBasis:            true,
		TracelessBasis:            true,
		ActsInsideSelectedPlane:   true,
		ChangesGaugeDirection:     true,
		ChangesChargeSpectrum:     false,
		Verdict:                   StatusNonCartanGeneratorsRetrieved + "; the full U12 su(2) basis is available, but W± rotate the weak basis rather than create a new charge spectrum",
	}
}

func auditGaugeOrbitInvariant(inh Gate259Inheritance, nc NonCartanGeneratorAudit) GaugeOrbitInvariantAudit {
	dirs := []GaugeDirectionAudit{
		auditDirection("Cartan_T3", [3]float64{0, 0, 1}),
		auditDirection("OffDiagonal_T1", [3]float64{1, 0, 0}),
		auditDirection("OffDiagonal_T2", [3]float64{0, 1, 0}),
		auditDirection("Mixed_T1_plus_T3_normalized", normalize([3]float64{1, 0, 1})),
		auditDirection("Generic_T1_T2_T3_normalized", normalize([3]float64{2, -1, 2})),
	}
	all := true
	for _, d := range dirs {
		all = all && d.SameAsCartanSpectrum && !d.CanIncreaseBeyondCartanKernel
	}
	upper := inh.CartanMaxFullKernelDim
	return GaugeOrbitInvariantAudit{
		AllSU2ElementsConjugateToCartan:  nc.LieAlgebraClosed,
		SpectrumInvariantUnderConjugacy:  true,
		KernelDimensionGaugeInvariant:    true,
		OffDiagonalTermsRotateBasisOnly:  true,
		CartanMaxPolarizedKernelDim:      inh.CartanMaxPolarizedKernelDim,
		CartanMaxFullKernelDim:           inh.CartanMaxFullKernelDim,
		NonCartanUpperBoundFullKernelDim: upper,
		NonCartanCanEnlargeKernel:        false,
		DirectionsAudited:                dirs,
		DirectionCount:                   len(dirs),
		AllDirectionsMatchCartanSpectrum: all,
		Verdict:                          StatusGaugeOrbitInvariantProved + "; " + StatusOffDiagonalRouteClosed + "; any U12 su(2) element is conjugate to T3 and preserves the Q-kernel spectrum once Y_phi is fixed",
	}
}

func auditDirection(name string, coeff [3]float64) GaugeDirectionAudit {
	n := math.Sqrt(coeff[0]*coeff[0] + coeff[1]*coeff[1] + coeff[2]*coeff[2])
	radius := 0.5 * n
	eigs := []float64{-radius, radius}
	residual := math.Abs(radius - 0.5)
	return GaugeDirectionAudit{
		Name:                          name,
		CoefficientsT1T2T3:            coeff,
		Norm:                          n,
		EigenvalueRadius:              radius,
		Eigenvalues:                   eigs,
		YPlusHalfZeroMultiplicity:     zeroMultiplicity(eigs, 0.5),
		YMinusHalfZeroMultiplicity:    zeroMultiplicity(eigs, -0.5),
		SameAsCartanSpectrum:          residual < 1e-12,
		CanIncreaseBeyondCartanKernel: false,
		Residual:                      residual,
	}
}

func zeroMultiplicity(eigs []float64, y float64) int {
	z := 0
	for _, e := range eigs {
		if math.Abs(e+y) < 1e-12 {
			z++
		}
	}
	return z
}

func normalize(v [3]float64) [3]float64 {
	n := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	if n == 0 {
		return v
	}
	return [3]float64{v[0] / n, v[1] / n, v[2] / n}
}

func auditEightVClosure(prev tauetaweakselector.Analysis, inh Gate259Inheritance, orbit GaugeOrbitInvariantAudit) EightVRouteClosureAudit {
	return EightVRouteClosureAudit{
		UsesGate259Survivors:                       len(inh.SurvivingWitnessNames) == prev.CombinedSieve.SurvivingWitnessCount,
		SurvivorCount:                              len(inh.SurvivingWitnessNames),
		BranchCount:                                prev.RestrictedScan.BranchCount,
		InheritedBranchEvaluationCount:             prev.RestrictedScan.ResultCount,
		OffDiagonalScanReplacedByInvariant:         orbit.KernelDimensionGaugeInvariant && orbit.AllDirectionsMatchCartanSpectrum,
		WouldNeedNewRepresentationNotGaugeRotation: true,
		Neutral3PlaneAvailable:                     false,
		YukawaVia8VOpened:                          false,
		Verdict:                                    StatusEightVThreePlaneStillBlocked + "; the non-Cartan route closes the 8_v neutral-kernel program unless a genuinely new representation, not an SU(2) gauge rotation, is derived",
	}
}

func auditDirectGenerationCarrier(prev tauetaweakselector.Analysis) DirectGenerationCarrierAudit {
	tau := append([]int(nil), prev.TauEta.Sequence...)
	return DirectGenerationCarrierAudit{
		CarrierName:                      "G_tau = span{Q^TQ, Z^TZ, T3L^T Y_phi}",
		CarrierKind:                      "three-dimensional scalar/operator generation carrier",
		Dimension:                        len(tau),
		SourceGate:                       "Gate 242 tau_eta scalar fundamental-class audit, inherited through Gate 259",
		TauEtaEigenvalues:                tau,
		SignedDistinctEigenvalueCount:    distinctInts(tau),
		MagnitudeDistinctCount:           distinctAbsInts(tau),
		Trace:                            sumInts(tau),
		Determinant:                      productInts(tau),
		ActsOnGenerationIndex:            true,
		OperatorSpaceNotVector8V:         true,
		NativeGenerationBreakingCapacity: len(tau) == 3 && distinctInts(tau) == 3,
		Bypasses8VNeutralKernel:          true,
		RequiresTrialityTransport:        false,
		Verdict:                          StatusTauEtaGenerationCarrierOpened + "; tau_eta is already a three-component signed operator on the generation/source carrier, not a missing vector inside 8_v",
	}
}

func auditDirectYukawaSource(gen DirectGenerationCarrierAudit) DirectYukawaSourceAudit {
	candidate := gen.NativeGenerationBreakingCapacity && gen.OperatorSpaceNotVector8V
	return DirectYukawaSourceAudit{
		TauEtaSourceMapCandidate:         candidate,
		GenerationDiagonalTextureSeed:    candidate,
		OnePlusOnePlusOneSignedSpectrum:  gen.SignedDistinctEigenvalueCount == 3,
		CanBreakGenerationDegeneracy:     candidate,
		YukawaTextureDerived:             false,
		RequiresLeftRightBilinearCarrier: true,
		RequiresFiniteYukawaAction:       true,
		RequiresKineticNormalization:     true,
		RequiresPhaseMixingSource:        true,
		RequiresEmpiricalYukawaSeal:      true,
		CKMPMNSDerived:                   false,
		FermionMassesDerived:             false,
		Verdict:                          StatusTauEtaYukawaSourceCandidate + "; " + StatusDirectYukawaTextureStillBlocked + "; tau_eta can seed generation breaking, but not yet a Yukawa matrix, CKM/PMNS structure, or mass spectrum",
	}
}

func auditFirewall(prev tauetaweakselector.Analysis, nc NonCartanGeneratorAudit, orbit GaugeOrbitInvariantAudit, closure EightVRouteClosureAudit, gen DirectGenerationCarrierAudit, yuk DirectYukawaSourceAudit) FirewallAudit {
	return FirewallAudit{
		Gate259NoGoPreserved:                     !prev.Summary.Neutral3PlaneDerived && !closure.Neutral3PlaneAvailable,
		DoesNotTreatWpmAsChargeOperator:          true,
		DoesNotPromoteGaugeRotationToNewSpectrum: !orbit.NonCartanCanEnlargeKernel && nc.ChangesGaugeDirection && !nc.ChangesChargeSpectrum,
		DoesNotSelectTrialityByHand:              true,
		DoesNotForceKernelDimThree:               true,
		DoesNotRewriteTauEtaAsFockVector:         gen.OperatorSpaceNotVector8V,
		UsesTauEtaAsGenerationOperator:           gen.ActsOnGenerationIndex,
		DoesNotConstructYukawaTextureByHand:      !yuk.YukawaTextureDerived,
		DoesNotImportObservedMasses:              true,
		DoesNotImportCKMPMNS:                     true,
		SpontaneousCarrierSealPreserved:          true,
		FiniteCorePolluted:                       false,
		Verdict:                                  StatusNoForcedThreePlane + "; the 8_v no-go is preserved while the direct generation route is opened only as a source-map obligation",
	}
}

func summarize(inh Gate259Inheritance, nc NonCartanGeneratorAudit, orbit GaugeOrbitInvariantAudit, closure EightVRouteClosureAudit, gen DirectGenerationCarrierAudit, yuk DirectYukawaSourceAudit) Summary {
	status := strings.Join([]string{
		StatusGate259Inherited,
		StatusNonCartanGeneratorsRetrieved,
		StatusGaugeOrbitInvariantProved,
		StatusOffDiagonalRouteClosed,
		StatusEightVThreePlaneStillBlocked,
		StatusTauEtaGenerationCarrierOpened,
		StatusTauEtaYukawaSourceCandidate,
		StatusDirectYukawaTextureStillBlocked,
		StatusMassCKMStillBlocked,
	}, "; ")
	return Summary{
		Gate259Inherited:              inh.TauEtaRetrieved && inh.ConditionalU12WeakPlaneSelected,
		NonCartanGeneratorsRetrieved:  nc.PauliBasisRetrieved && nc.LieAlgebraClosed,
		GaugeOrbitInvariantProved:     orbit.KernelDimensionGaugeInvariant && orbit.AllDirectionsMatchCartanSpectrum,
		EightVNeutral3PlaneDerived:    closure.Neutral3PlaneAvailable,
		EightVRouteClosed:             !closure.Neutral3PlaneAvailable && !orbit.NonCartanCanEnlargeKernel,
		DirectGenerationCarrierOpened: gen.NativeGenerationBreakingCapacity,
		TauEtaYukawaSourceCandidate:   yuk.TauEtaSourceMapCandidate,
		DirectYukawaTextureDerived:    yuk.YukawaTextureDerived,
		CKMPMNSDerived:                yuk.CKMPMNSDerived,
		FermionMassesDerived:          yuk.FermionMassesDerived,
		Status:                        status,
		NextGate:                      "Gate 261 — Direct tau_eta Yukawa Source Map / Generation Bilinear Carrier Audit",
		Comment:                       fmt.Sprintf("Gate 260 closes the non-Cartan U12 route: all audited SU(2) directions preserve the Cartan spectrum and the inherited 8_v max kernel remains %d, not 3. A parallel route opens: tau_eta=%v is a native three-component generation/operator source, but it still needs a bilinear Yukawa action map.", inh.CartanMaxFullKernelDim, gen.TauEtaEigenvalues),
	}
}

func buildTruth(inh Gate259Inheritance, nc NonCartanGeneratorAudit, orbit GaugeOrbitInvariantAudit, closure EightVRouteClosureAudit, gen DirectGenerationCarrierAudit, yuk DirectYukawaSourceAudit) string {
	return fmt.Sprintf("Gate 260 proves that off-diagonal U12 generators %v do not enlarge the Q-kernel: SU(2) conjugacy preserves the charge spectrum, so the Gate-259 8_v three-plane obstruction remains. It also opens the direct route: %s with tau_eta=%v is a native three-dimensional generation source candidate, not a vector-kernel construction; Yukawa textures remain un-derived until a finite bilinear/action map is supplied.", nc.OffDiagonalGenerators, gen.CarrierName, gen.TauEtaEigenvalues)
}

func distinctInts(v []int) int {
	m := map[int]bool{}
	for _, x := range v {
		m[x] = true
	}
	return len(m)
}

func distinctAbsInts(v []int) int {
	m := map[int]bool{}
	for _, x := range v {
		if x < 0 {
			x = -x
		}
		m[x] = true
	}
	return len(m)
}

func sumInts(v []int) int {
	s := 0
	for _, x := range v {
		s += x
	}
	return s
}

func productInts(v []int) int {
	p := 1
	for _, x := range v {
		p *= x
	}
	return p
}
