// Package spin8trialityfunctor implements Gate 247:
// Spin(8) Triality Automorphism / Scalar-to-Spinor Functor Audit.
//
// Gate 246 proved that tau_eta=(2,-2,1) has exactly the right 1+1+1
// generation-breaking and non-commuting texture capacity, but it also proved
// that the scalar-bundle -> triality-carrier functor is missing. Gate 247 tests
// whether Spin(8) triality itself supplies that functor. The audit is strict:
// Spin(8) triality is an automorphism among representation categories
// (8_v,8_s,8_c), not a magic map from an arbitrary scalar trace triple into a
// spinor endomorphism. Without an 8_v/exterior representative of tau_eta and
// explicit triality automorphism data on S_C, the pullback remains blocked.
package spin8trialityfunctor

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalartrialitytexture"
)

const (
	AuditID = "GATE247-SPIN8-TRIALITY-AUTOMORPHISM-SCALAR-TO-SPINOR-FUNCTOR-AUDIT"

	StatusSpin8TrialityPreflight          = "CONDITIONAL_SUPPORT_SPIN8_TRIALITY_AUTOMORPHISM_PREFLIGHT"
	StatusTrialityDimensionMatch          = "CONDITIONAL_SUPPORT_TRIALITY_SCALAR_SPINOR_DIMENSION_MATCH"
	StatusTauEtaTextureCapacityInherited  = "CONDITIONAL_SUPPORT_TAU_ETA_TEXTURE_CAPACITY_INHERITED"
	StatusScalarTraceNotVectorRep         = "FAILED_ROUTE_SCALAR_TRACE_NOT_VECTOR_REPRESENTATIVE"
	StatusTrialityFunctorPullbackBlocked  = "FAILED_ROUTE_TRIALITY_FUNCTOR_PULLBACK_DERIVATION"
	StatusYukawaTextureBlocked            = "FAILED_ROUTE_TRIALITY_FUNCTOR_YUKAWA_DERIVATION"
	StatusCKMPMNSBlocked                  = "FAILED_ROUTE_CKM_PMNS_DERIVATION"
	StatusYukawaAmplitudeSealStillBinding = "YUKAWA_AMPLITUDE_SEAL_REMAINS_BINDING"
)

type InheritedGate246Audit struct {
	ScalarOriginKnown              bool
	ScalarToTrialityFunctorDerived bool
	TauGenerationCapacity          bool
	GenerationTextureDerived       bool
	RawNonCommutingCapacity        bool
	QualifiedTexturePairDerived    bool
	CKMPMNSDerived                 bool
	FermionMassesDerived           bool
	TruthStatement                 string
}

type Spin8TrialityAudit struct {
	AbstractSpin8TrialityAvailable bool
	TrialityRepresentations        []string
	AutomorphismGroup              string
	VectorToSpinorFunctorKnown     bool
	ExplicitMatricesOnSC           bool
	ScalarBundleIsVectorRep        bool
	ScalarTraceTripleIsVector      bool
	Verdict                        string
}

type ScalarToSpinorFunctorAudit struct {
	SourceObject                        string
	SourceCarrier                       string
	RequiredSourceForTriality           string
	TargetCarrier                       string
	DimensionOfTraceTriple              int
	GenerationCarrierDimension          int
	DimensionMatch                      bool
	ExteriorOrVectorRepresentativeKnown bool
	CharacteristicRepresentativeKnown   bool
	ExplicitTrialityAutomorphismKnown   bool
	PullbackFunctorDerived              bool
	ManualPullback                      string
	ManualPullbackRejected              bool
	Verdict                             string
}

type Matrix3 [3][3]float64

type TextureRealizationAudit struct {
	CandidateName               string
	CandidateEigenvalues        []int
	DistinctEigenvalues         int
	BreaksGenerationDegeneracy  bool
	TrialityCycle               Matrix3
	TrialityReflection          Matrix3
	CommutatorWithCycle         Matrix3
	CommutatorWithReflection    Matrix3
	CycleCommutatorNorm         float64
	ReflectionCommutatorNorm    float64
	RawNonCommutingCapacity     bool
	LawfulPullbackDerived       bool
	DiagonalOperatorConstructed bool
	YukawaTextureDerived        bool
	CKMDerived                  bool
	PMNSDerived                 bool
	Verdict                     string
}

type PullbackObstructionAudit struct {
	MissingPieces           []string
	BindingTypeMismatch     string
	WhyTrialityInsufficient string
	PullbackDerived         bool
	ObstructionLevel        string
	Verdict                 string
}

type FirewallAudit struct {
	ImportedConnesAlgebra      bool
	InventedSpin8Matrices      bool
	ForcedScalarToSpinorMap    bool
	InsertedDTauAsTexture      bool
	ImportedYukawaMasses       bool
	ImportedCKM                bool
	ImportedPMNS               bool
	ClaimedFermionMasses       bool
	ClaimedFiniteFlavorTheorem bool
	PollutedFiniteCore         bool
	Verdict                    string
}

type Summary struct {
	Spin8TrialityAvailable      bool
	DimensionMatch              bool
	TauTextureCapacityInherited bool
	ScalarTraceIsVectorRep      bool
	TrialityFunctorDerived      bool
	DiagonalTextureConstructed  bool
	QualifiedTextureDerived     bool
	CKMPMNSDerived              bool
	FermionMassesDerived        bool
	Status                      string
	NextGate                    string
	Comment                     string
}

type Analysis struct {
	PreviousGate246 InheritedGate246Audit
	Spin8Triality   Spin8TrialityAudit
	ScalarSpinor    ScalarToSpinorFunctorAudit
	Texture         TextureRealizationAudit
	Obstruction     PullbackObstructionAudit
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
		prevRaw, err := scalartrialitytexture.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		prev := inheritGate246(prevRaw)
		tr := auditSpin8Triality()
		ss := auditScalarToSpinor(prev, tr)
		tx := auditTextureRealization(prev, ss)
		ob := auditObstruction(ss, tr)
		fw := auditFirewall()
		sum := summarize(prev, tr, ss, tx)
		truth := buildTruth(tr, ss, tx, ob)
		defaultA = Analysis{PreviousGate246: prev, Spin8Triality: tr, ScalarSpinor: ss, Texture: tx, Obstruction: ob, Firewall: fw, Summary: sum, TruthStatement: truth}
	})
	return defaultA, defaultErr
}

func inheritGate246(a scalartrialitytexture.Analysis) InheritedGate246Audit {
	return InheritedGate246Audit{
		ScalarOriginKnown:              a.Summary.ScalarOriginKnown,
		ScalarToTrialityFunctorDerived: a.Summary.ScalarToTrialityFunctorDerived,
		TauGenerationCapacity:          a.Summary.TauGenerationCapacity,
		GenerationTextureDerived:       a.Summary.GenerationTextureDerived,
		RawNonCommutingCapacity:        a.Summary.RawNonCommutingCapacity,
		QualifiedTexturePairDerived:    a.Summary.QualifiedTexturePairDerived,
		CKMPMNSDerived:                 a.Summary.CKMPMNSDerived,
		FermionMassesDerived:           a.Summary.FermionMassesDerived,
		TruthStatement:                 a.TruthStatement,
	}
}

func auditSpin8Triality() Spin8TrialityAudit {
	return Spin8TrialityAudit{
		AbstractSpin8TrialityAvailable: true,
		TrialityRepresentations:        []string{"8_v vector", "8_s left spinor", "8_c right spinor"},
		AutomorphismGroup:              "Out(Spin(8)) ≅ S3 permuting 8_v, 8_s, 8_c",
		VectorToSpinorFunctorKnown:     true,
		ExplicitMatricesOnSC:           false,
		ScalarBundleIsVectorRep:        false,
		ScalarTraceTripleIsVector:      false,
		Verdict:                        "abstract Spin(8) triality is the correct kind of representation-level bridge, but the project has not derived explicit triality automorphism matrices on S_C nor shown that the H_Phi scalar trace triple is an 8_v vector representative",
	}
}

func auditScalarToSpinor(prev InheritedGate246Audit, tr Spin8TrialityAudit) ScalarToSpinorFunctorAudit {
	return ScalarToSpinorFunctorAudit{
		SourceObject:                        "tau_eta = (2,-2,1)",
		SourceCarrier:                       "neutral scalar/Higgs trace ledger H_Phi",
		RequiredSourceForTriality:           "a vector/exterior representative in 8_v or a finite characteristic class representative with explicit Spin(8) triality action",
		TargetCarrier:                       "3-dimensional triality generation carrier inside the spinor/Yukawa texture ledger",
		DimensionOfTraceTriple:              3,
		GenerationCarrierDimension:          3,
		DimensionMatch:                      true,
		ExteriorOrVectorRepresentativeKnown: false,
		CharacteristicRepresentativeKnown:   false,
		ExplicitTrialityAutomorphismKnown:   tr.ExplicitMatricesOnSC,
		PullbackFunctorDerived:              false,
		ManualPullback:                      "tau_eta ?-> diag(2,-2,1) on generations via Spin(8) triality",
		ManualPullbackRejected:              true,
		Verdict:                             "the trace triple and generation carrier both have three slots, but triality acts on vector/spinor representations, not on an unrepresented scalar trace sequence; the required source representative is missing",
	}
}

func auditTextureRealization(prev InheritedGate246Audit, ss ScalarToSpinorFunctorAudit) TextureRealizationAudit {
	eigs := []int{2, -2, 1}
	D := diag(float64(eigs[0]), float64(eigs[1]), float64(eigs[2]))
	C := Matrix3{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}
	R := Matrix3{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}}
	commC := commutator(D, C)
	commR := commutator(D, R)
	raw := frob(commC) > 0 || frob(commR) > 0
	return TextureRealizationAudit{
		CandidateName:               "D_tau = diag(2,-2,1) on spinor generation carrier (conditional only)",
		CandidateEigenvalues:        eigs,
		DistinctEigenvalues:         distinctInts(eigs),
		BreaksGenerationDegeneracy:  distinctInts(eigs) == 3,
		TrialityCycle:               C,
		TrialityReflection:          R,
		CommutatorWithCycle:         commC,
		CommutatorWithReflection:    commR,
		CycleCommutatorNorm:         frob(commC),
		ReflectionCommutatorNorm:    frob(commR),
		RawNonCommutingCapacity:     raw && prev.RawNonCommutingCapacity,
		LawfulPullbackDerived:       ss.PullbackFunctorDerived,
		DiagonalOperatorConstructed: false,
		YukawaTextureDerived:        false,
		CKMDerived:                  false,
		PMNSDerived:                 false,
		Verdict:                     "D_tau would break S3 and fail to commute with triality permutations, but it remains an unconstructed diagnostic until Spin(8) triality receives a lawful scalar-trace representative to act on",
	}
}

func auditObstruction(ss ScalarToSpinorFunctorAudit, tr Spin8TrialityAudit) PullbackObstructionAudit {
	return PullbackObstructionAudit{
		MissingPieces: []string{
			"tau_eta as an element of 8_v or Λ*W",
			"explicit Spin(8) triality automorphism matrices on S_C",
			"H_Phi scalar trace representation as vector/scalar bundle object compatible with 8_v",
			"basis-independent map from neutral scalar trace slots to generation carrier",
			"order-one/spectral-triple permission to use the resulting operator in Yukawa matrices",
			"normalization into dimensionless Yukawa amplitudes",
		},
		BindingTypeMismatch:     "Spin(8) triality rotates representations 8_v,8_s,8_c; tau_eta is currently a scalar trace ledger, not a vector representative",
		WhyTrialityInsufficient: "abstract triality is a representation equivalence; it does not assign coordinates to an arbitrary three-entry trace sequence or turn it into a spinor endomorphism without a chosen representative",
		PullbackDerived:         false,
		ObstructionLevel:        "representation-domain mismatch",
		Verdict:                 "triality is the right mathematical arena, but it does not by itself solve the scalar-to-spinor pullback missing since Gate 246",
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ImportedConnesAlgebra:      false,
		InventedSpin8Matrices:      false,
		ForcedScalarToSpinorMap:    false,
		InsertedDTauAsTexture:      false,
		ImportedYukawaMasses:       false,
		ImportedCKM:                false,
		ImportedPMNS:               false,
		ClaimedFermionMasses:       false,
		ClaimedFiniteFlavorTheorem: false,
		PollutedFiniteCore:         false,
		Verdict:                    "Spin(8) triality is audited as native representation theory only; no scalar trace is forced into a spinor texture and no observed flavor data are inserted",
	}
}

func summarize(prev InheritedGate246Audit, tr Spin8TrialityAudit, ss ScalarToSpinorFunctorAudit, tx TextureRealizationAudit) Summary {
	status := strings.Join([]string{
		StatusSpin8TrialityPreflight,
		StatusTrialityDimensionMatch,
		StatusTauEtaTextureCapacityInherited,
		StatusScalarTraceNotVectorRep,
		StatusTrialityFunctorPullbackBlocked,
		StatusYukawaTextureBlocked,
		StatusCKMPMNSBlocked,
		StatusYukawaAmplitudeSealStillBinding,
	}, "\n")
	return Summary{
		Spin8TrialityAvailable:      tr.AbstractSpin8TrialityAvailable && tr.VectorToSpinorFunctorKnown,
		DimensionMatch:              ss.DimensionMatch,
		TauTextureCapacityInherited: prev.TauGenerationCapacity && tx.RawNonCommutingCapacity,
		ScalarTraceIsVectorRep:      tr.ScalarTraceTripleIsVector || ss.ExteriorOrVectorRepresentativeKnown,
		TrialityFunctorDerived:      ss.PullbackFunctorDerived,
		DiagonalTextureConstructed:  tx.DiagonalOperatorConstructed,
		QualifiedTextureDerived:     tx.YukawaTextureDerived,
		CKMPMNSDerived:              tx.CKMDerived || tx.PMNSDerived,
		FermionMassesDerived:        false,
		Status:                      status,
		NextGate:                    "derive an 8_v/exterior representative of tau_eta or a scalar-bundle representation compatible with Spin(8) triality before using triality as a pullback functor",
		Comment:                     "Gate 247 confirms that Spin(8) triality is the right representation-theoretic bridge but rejects its use as a scalar-trace-to-spinor functor without a representative in the automorphism domain.",
	}
}

func buildTruth(tr Spin8TrialityAudit, ss ScalarToSpinorFunctorAudit, tx TextureRealizationAudit, ob PullbackObstructionAudit) string {
	return fmt.Sprintf("Gate 247 audits Spin(8) triality as the missing scalar-to-spinor functor. Abstract triality among %s is available, and the tau_eta texture capacity inherited from Gate 246 remains strong: D_tau would have %d distinct eigenvalues and nonzero triality commutators (cycle norm %.6g). But tau_eta is not an 8_v/vector representative; it is a scalar trace ledger. Therefore Spin(8) triality cannot lawfully pull it into the spinor generation carrier. Binding obstruction: %s.", strings.Join(tr.TrialityRepresentations, ", "), tx.DistinctEigenvalues, tx.CycleCommutatorNorm, ob.BindingTypeMismatch)
}

func distinctInts(xs []int) int {
	m := map[int]bool{}
	for _, x := range xs {
		m[x] = true
	}
	return len(m)
}

func diag(a, b, c float64) Matrix3 { return Matrix3{{a, 0, 0}, {0, b, 0}, {0, 0, c}} }

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

func commutator(a, b Matrix3) Matrix3 { return sub(mul(a, b), mul(b, a)) }

func frob(a Matrix3) float64 {
	var s float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s += a[i][j] * a[i][j]
		}
	}
	return math.Sqrt(s)
}
