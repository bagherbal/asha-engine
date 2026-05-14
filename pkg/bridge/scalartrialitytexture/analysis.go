// Package scalartrialitytexture implements Gate 246:
// Scalar Bundle to Triality Pullback / Yukawa Generation Texture Audit.
//
// Gate 245 proved that tau_eta=(2,-2,1) is not a spatial-axis selector: its
// source records are neutral electroweak scalar-bundle observables Q^TQ,
// Z^TZ, and T3L^T Y_phi. Gate 246 redirects the audit to the mathematically
// appropriate target: generation/flavor texture. The audit is strict. The
// signed triple has exactly the right 1+1+1 eigenvalue capacity, and if it were
// lawfully pulled back to the 3D triality generation carrier it would provide a
// diagonal generation-breaking source which does not commute with triality
// permutations. But the scalar-bundle -> triality-carrier functor is not yet
// derived, so no Yukawa texture, masses, CKM, or PMNS theorem is claimed.
package scalartrialitytexture

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/liecarrierprojection"
)

const (
	AuditID = "GATE246-SCALAR-BUNDLE-TO-TRIALITY-PULLBACK-YUKAWA-GENERATION-TEXTURE-AUDIT"

	StatusScalarOriginInherited          = "CONDITIONAL_SUPPORT_SCALAR_BUNDLE_ORIGIN_INHERITED"
	StatusTauEtaGenerationCapacity       = "CONDITIONAL_SUPPORT_TAU_ETA_GENERATION_BREAKING_CAPACITY"
	StatusNonCommutingCapacity           = "CONDITIONAL_SUPPORT_TAU_ETA_TRIALITY_NONCOMMUTING_CAPACITY"
	StatusScalarTrialityPullbackBlocked  = "FAILED_ROUTE_SCALAR_TO_TRIALITY_PULLBACK"
	StatusYukawaTextureBlocked           = "FAILED_ROUTE_TAU_ETA_YUKAWA_TEXTURE_DERIVATION"
	StatusCKMPMNSBlocked                 = "FAILED_ROUTE_CKM_PMNS_DERIVATION"
	StatusFermionMassesBlocked           = "FAILED_ROUTE_FERMION_MASS_DERIVATION"
	StatusYukawaAmplitudeSealStillNeeded = "YUKAWA_AMPLITUDE_SEAL_REMAINS_BINDING"
)

type InheritedGate245Audit struct {
	TauEtaOriginTraced            bool
	TauSlotsAreSU2Basis           bool
	CarrierProjectionDerived      bool
	ExteriorRepresentativeDerived bool
	WeakPlaneDerived              bool
	GenerationBreakingCapacity    bool
	GenerationTextureDerived      bool
	SourceIsNeutralScalarEWPlane  bool
	TruthStatement                string
}

type ScalarFlavorAlignmentAudit struct {
	TauSequence                    []int
	SourceBundle                   string
	SourceOperators                []string
	NativeHiggsSectorObservable    bool
	ScalarOriginKnown              bool
	ScalarToTrialityFunctorDerived bool
	TrialityCarrierDimension       int
	TrialityCarrierDerivedEarlier  bool
	MapWouldBeTypeCorrect          bool
	MapActuallyDerived             bool
	Verdict                        string
}

type GenerationTextureAudit struct {
	CandidateOperatorName     string
	Eigenvalues               []int
	DistinctEigenvalues       int
	BreaksS3Degeneracy        bool
	Trace                     int
	Determinant               int
	FrobeniusNormSquared      int
	HermitianDiagonalCapacity bool
	ScalarToGenerationMap     bool
	GenerationOperatorDerived bool
	YukawaTextureDerived      bool
	TextureIfMapExisted       string
	Verdict                   string
}

type Matrix3 [3][3]float64

type NonCommutingTextureAudit struct {
	Gate173NeedsNonCommutingPair       bool
	Gate173FoundQualifiedPair          bool
	TrialityCycleName                  string
	TrialityReflectionName             string
	TauDiagonalName                    string
	CommutatorWithCycle                Matrix3
	CommutatorWithReflection           Matrix3
	CycleCommutatorNorm                float64
	ReflectionCommutatorNorm           float64
	RawNonCommutingWithTriality        bool
	PairWouldBeQualifiedIfPullbackHeld bool
	PairActuallyQualified              bool
	ReasonNotQualified                 string
	CKMPrerequisiteCapacity            bool
	PMNSPrerequisiteCapacity           bool
	CKMDerived                         bool
	PMNSDerived                        bool
	Verdict                            string
}

type PullbackObstructionAudit struct {
	ScalarBundleCarrier     string
	TrialityCarrier         string
	KnownSharedStructure    []string
	MissingFunctor          string
	MissingCompatibility    []string
	ManualDiagonalInsertion string
	ManualInsertionRejected bool
	PullbackDerived         bool
	Reason                  string
	Verdict                 string
}

type FirewallAudit struct {
	ForcedScalarToGenerationMap bool
	ForcedTauDiagonalTexture    bool
	ImportedYukawaMasses        bool
	ImportedCKM                 bool
	ImportedPMNS                bool
	InsertedObservedMasses      bool
	ClaimedFermionMasses        bool
	ClaimedFiniteFlavorTheorem  bool
	ClaimedWeakPlane            bool
	FiniteCorePolluted          bool
	Verdict                     string
}

type Summary struct {
	ScalarOriginKnown              bool
	ScalarToTrialityFunctorDerived bool
	TauGenerationCapacity          bool
	GenerationTextureDerived       bool
	RawNonCommutingCapacity        bool
	QualifiedTexturePairDerived    bool
	CKMPMNSDerived                 bool
	FermionMassesDerived           bool
	WeakPlaneDerived               bool
	Status                         string
	NextGate                       string
	Comment                        string
}

type Analysis struct {
	PreviousGate245     InheritedGate245Audit
	ScalarFlavor        ScalarFlavorAlignmentAudit
	GenerationTexture   GenerationTextureAudit
	NonCommutingTexture NonCommutingTextureAudit
	PullbackObstruction PullbackObstructionAudit
	Firewall            FirewallAudit
	Summary             Summary
	TruthStatement      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prevRaw, err := liecarrierprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate245(prevRaw)
		sf := auditScalarFlavor(prev)
		gt := auditGenerationTexture(sf)
		nc := auditNonCommutingTexture(gt)
		po := auditPullbackObstruction(sf)
		fw := auditFirewall()
		sum := summarize(sf, gt, nc, po)
		truth := buildTruth(sf, gt, nc, po)
		defaultA = Analysis{PreviousGate245: prev, ScalarFlavor: sf, GenerationTexture: gt, NonCommutingTexture: nc, PullbackObstruction: po, Firewall: fw, Summary: sum, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate245(a liecarrierprojection.Analysis) InheritedGate245Audit {
	return InheritedGate245Audit{
		TauEtaOriginTraced:            a.Summary.OperatorDecompositionTraced,
		TauSlotsAreSU2Basis:           a.Summary.TauSlotsAreSU2Basis,
		CarrierProjectionDerived:      a.Summary.CarrierProjectionDerived,
		ExteriorRepresentativeDerived: a.Summary.ExteriorRepresentativeDerived,
		WeakPlaneDerived:              a.Summary.WeakPlaneDerived,
		GenerationBreakingCapacity:    a.Summary.GenerationBreakingCapacity,
		GenerationTextureDerived:      a.Summary.GenerationTextureDerived,
		SourceIsNeutralScalarEWPlane:  a.OperatorDecomposition.QZMixT3AndYPhi && a.OperatorDecomposition.SlotsAreQuadraticScalarRecords,
		TruthStatement:                a.TruthStatement,
	}
}

func auditScalarFlavor(prev InheritedGate245Audit) ScalarFlavorAlignmentAudit {
	return ScalarFlavorAlignmentAudit{
		TauSequence:                    []int{2, -2, 1},
		SourceBundle:                   "neutral electroweak scalar bundle H_Phi",
		SourceOperators:                []string{"tau_eta(Q^T Q)", "tau_eta(Z^T Z)", "tau_eta(T3L^T Y_phi)"},
		NativeHiggsSectorObservable:    prev.SourceIsNeutralScalarEWPlane,
		ScalarOriginKnown:              prev.TauEtaOriginTraced && prev.SourceIsNeutralScalarEWPlane,
		ScalarToTrialityFunctorDerived: false,
		TrialityCarrierDimension:       3,
		TrialityCarrierDerivedEarlier:  true,
		MapWouldBeTypeCorrect:          true,
		MapActuallyDerived:             false,
		Verdict:                        "tau_eta lives in the scalar/Higgs-sector trace ledger and therefore points to flavor texture more naturally than spatial orientation, but the scalar-bundle -> triality-carrier functor is not derived",
	}
}

func auditGenerationTexture(sf ScalarFlavorAlignmentAudit) GenerationTextureAudit {
	eigs := []int{2, -2, 1}
	return GenerationTextureAudit{
		CandidateOperatorName:     "D_tau = diag(2,-2,1) on triality generation carrier (conditional only)",
		Eigenvalues:               eigs,
		DistinctEigenvalues:       distinctInts(eigs),
		BreaksS3Degeneracy:        distinctInts(eigs) == 3,
		Trace:                     eigs[0] + eigs[1] + eigs[2],
		Determinant:               eigs[0] * eigs[1] * eigs[2],
		FrobeniusNormSquared:      eigs[0]*eigs[0] + eigs[1]*eigs[1] + eigs[2]*eigs[2],
		HermitianDiagonalCapacity: true,
		ScalarToGenerationMap:     sf.ScalarToTrialityFunctorDerived,
		GenerationOperatorDerived: false,
		YukawaTextureDerived:      false,
		TextureIfMapExisted:       "D_tau would be a self-adjoint 3-distinct-eigenvalue generation spurion, splitting triality as 1+1+1",
		Verdict:                   "the signed tau_eta sequence has exact diagonal generation-breaking capacity, but remains a conditional texture until the scalar-to-triality pullback is derived",
	}
}

func auditNonCommutingTexture(gt GenerationTextureAudit) NonCommutingTextureAudit {
	D := diag(float64(gt.Eigenvalues[0]), float64(gt.Eigenvalues[1]), float64(gt.Eigenvalues[2]))
	C := Matrix3{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	R := Matrix3{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	commC := commutator(D, C)
	commR := commutator(D, R)
	nC := frob(commC)
	nR := frob(commR)
	raw := nC > 0 || nR > 0
	return NonCommutingTextureAudit{
		Gate173NeedsNonCommutingPair:       true,
		Gate173FoundQualifiedPair:          false,
		TrialityCycleName:                  "C3_cycle",
		TrialityReflectionName:             "S3_reflection",
		TauDiagonalName:                    "D_tau",
		CommutatorWithCycle:                commC,
		CommutatorWithReflection:           commR,
		CycleCommutatorNorm:                nC,
		ReflectionCommutatorNorm:           nR,
		RawNonCommutingWithTriality:        raw,
		PairWouldBeQualifiedIfPullbackHeld: raw && gt.BreaksS3Degeneracy,
		PairActuallyQualified:              false,
		ReasonNotQualified:                 "D_tau is not derived as a scalar-to-triality endomorphism; raw commutators are capacity diagnostics, not Yukawa texture theorems",
		CKMPrerequisiteCapacity:            raw,
		PMNSPrerequisiteCapacity:           raw,
		CKMDerived:                         false,
		PMNSDerived:                        false,
		Verdict:                            "if D_tau were lawfully pulled back, it would not commute with triality permutations and would supply the missing kind of texture capacity; the qualified pair is still blocked by the missing pullback functor",
	}
}

func auditPullbackObstruction(sf ScalarFlavorAlignmentAudit) PullbackObstructionAudit {
	return PullbackObstructionAudit{
		ScalarBundleCarrier:     "H_Phi neutral scalar/electroweak trace bundle",
		TrialityCarrier:         "3-dimensional generation carrier from exact triality bookkeeping",
		KnownSharedStructure:    []string{"both are finite ledgers", "both have three-entry diagnostic data", "scalar sector supplies Higgs/Yukawa context", "triality sector supplies generation labels"},
		MissingFunctor:          "H_Phi scalar trace functional -> generation-carrier endomorphism",
		MissingCompatibility:    []string{"basis identification", "charge/Yukawa block compatibility", "left/right chirality placement", "fermion species replication", "order-one/spectral-triple permission", "normalization into Yukawa amplitudes"},
		ManualDiagonalInsertion: "D_tau ?= diag(2,-2,1) on generations",
		ManualInsertionRejected: true,
		PullbackDerived:         false,
		Reason:                  "three scalar neutral-current trace slots are not yet a generation basis; assigning them to generations would be a phenomenological ansatz, not finite algebra",
		Verdict:                 "scalar-to-triality pullback remains the missing theorem; tau_eta is the right type of scalar-sector object for flavor only after that functor is derived or sealed",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ForcedScalarToGenerationMap: false,
		ForcedTauDiagonalTexture:    false,
		ImportedYukawaMasses:        false,
		ImportedCKM:                 false,
		ImportedPMNS:                false,
		InsertedObservedMasses:      false,
		ClaimedFermionMasses:        false,
		ClaimedFiniteFlavorTheorem:  false,
		ClaimedWeakPlane:            false,
		FiniteCorePolluted:          false,
		Verdict:                     "tau_eta is not manually inserted into any Yukawa matrix; observed masses and CKM/PMNS remain sealed empirical data",
	}
}

func summarize(sf ScalarFlavorAlignmentAudit, gt GenerationTextureAudit, nc NonCommutingTextureAudit, po PullbackObstructionAudit) Summary {
	status := strings.Join([]string{
		StatusScalarOriginInherited,
		StatusTauEtaGenerationCapacity,
		StatusNonCommutingCapacity,
		StatusScalarTrialityPullbackBlocked,
		StatusYukawaTextureBlocked,
		StatusCKMPMNSBlocked,
		StatusFermionMassesBlocked,
		StatusYukawaAmplitudeSealStillNeeded,
	}, "\n")
	return Summary{
		ScalarOriginKnown:              sf.ScalarOriginKnown,
		ScalarToTrialityFunctorDerived: sf.ScalarToTrialityFunctorDerived,
		TauGenerationCapacity:          gt.BreaksS3Degeneracy && gt.DistinctEigenvalues == 3,
		GenerationTextureDerived:       gt.GenerationOperatorDerived && gt.YukawaTextureDerived,
		RawNonCommutingCapacity:        nc.RawNonCommutingWithTriality,
		QualifiedTexturePairDerived:    nc.PairActuallyQualified,
		CKMPMNSDerived:                 nc.CKMDerived || nc.PMNSDerived,
		FermionMassesDerived:           false,
		WeakPlaneDerived:               false,
		Status:                         status,
		NextGate:                       "derive or seal the scalar-to-triality representation functor before using tau_eta as a Yukawa texture source",
		Comment:                        "Gate 246 redirects tau_eta to flavor where it belongs, proves strong generation-breaking/noncommuting capacity, and blocks the theorem because the carrier pullback is still missing.",
	}
}

func buildTruth(sf ScalarFlavorAlignmentAudit, gt GenerationTextureAudit, nc NonCommutingTextureAudit, po PullbackObstructionAudit) string {
	return fmt.Sprintf("Gate 246 proves that tau_eta's neutral scalar-bundle origin makes it a plausible flavor/generation texture source, not a spatial weak-plane selector. The conditional diagonal D_tau=diag(2,-2,1) has three distinct eigenvalues and would not commute with triality permutations (cycle commutator norm %.6g), supplying exactly the kind of generation-breaking capacity Gate 173 lacked. But the scalar-to-triality functor is not derived, so D_tau remains capacity only and no Yukawa matrices, fermion masses, CKM, or PMNS data are claimed. Missing functor: %s.", nc.CycleCommutatorNorm, po.MissingFunctor)
}

func distinctInts(xs []int) int {
	m := map[int]bool{}
	for _, x := range xs {
		m[x] = true
	}
	return len(m)
}

func diag(a, b, c float64) Matrix3 {
	return Matrix3{{a, 0, 0}, {0, b, 0}, {0, 0, c}}
}

func commutator(a, b Matrix3) Matrix3 {
	return sub(mul(a, b), mul(b, a))
}

func mul(a, b Matrix3) Matrix3 {
	var r Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				r[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return r
}

func sub(a, b Matrix3) Matrix3 {
	var r Matrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r[i][j] = a[i][j] - b[i][j]
		}
	}
	return r
}

func frob(a Matrix3) float64 {
	var s float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
