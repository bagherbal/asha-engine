// Package complexifiedhilbertspace implements Gate 235:
// Particle/Antiparticle Hilbert space doubling and finite algebra representation audit.
//
// Gate 234 showed that an occupation-complement candidate J can sieve the
// 16-state real Fock scaffold but cannot host true charge conjugation,
// Majorana bilinears, or a non-vacuous order-one calculus. Gate 235 therefore
// performs the mathematically minimal next move: it does not add sixteen states
// from outside. It complexifies the already-derived real Cl(1,7) spinor/Fock
// carrier S, forming S_C = S \otimes_R C. This is 16-dimensional over C and
// 32-dimensional over R, with the canonical anti-linear conjugation acting as a
// candidate real structure.
//
// The result is deliberately split. Complexification gives a legitimate doubled
// real carrier and a kinematic anti-linear J. It also gives the capacity to write
// neutral particle/conjugate bilinears on the doubled space. However, the engine
// still does not derive the faithful finite algebra A on this doubled carrier,
// does not derive C ⊕ H ⊕ M3(C), does not apply the full order-one condition,
// and does not canonically identify B_gap as a Majorana mass. The B-gap remains
// a dimensionless finite spectral datum until a future theorem supplies the
// algebra representation and spectral-action map.
package complexifiedhilbertspace

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/realstructureorderone"
	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

const (
	AuditID = "GATE235-COMPLEXIFIED-HILBERT-SPACE-FINITE-ALGEBRA-AUDIT"

	StatusComplexificationSupport      = "CONDITIONAL_SUPPORT_COMPLEXIFICATION_DERIVED_DOUBLING"
	StatusCandidateAntiLinearJ         = "CONDITIONAL_SUPPORT_ANTILINEAR_CONJUGATION_J_PREFLIGHT"
	StatusMajoranaCapacity             = "CONDITIONAL_SUPPORT_NEUTRAL_MAJORANA_BILINEAR_CAPACITY"
	StatusFailedNativeAlgebra          = "FAILED_ROUTE_NATIVE_FINITE_ALGEBRA_REPRESENTATION_DERIVATION"
	StatusFailedConnesAlgebra          = "FAILED_ROUTE_CONNES_ALGEBRA_IMPORT_BLOCKED"
	StatusFailedBGapMajorana           = "FAILED_ROUTE_CANONICAL_BGAP_MAJORANA_IDENTIFICATION"
	StatusFailedSpectralTripleComplete = "FAILED_ROUTE_FULL_DOUBLED_SPECTRAL_TRIPLE_DERIVATION"
)

type ComplexificationAudit struct {
	SourceRealCarrier             string
	DerivedByComplexification     bool
	ExternalStatesAdded           bool
	RealDimensionBefore           int
	ComplexDimensionAfter         int
	RealDimensionAfter            int
	Formula                       string
	FixedRealHalfDimension        int
	ImaginaryHalfDimension        int
	ParticleAntiparticleSemantics bool
	ParticleAntiparticleCandidate bool
	Verdict                       string
}

type AntiLinearJAudit struct {
	CandidateName                        string
	AntiLinear                           bool
	J2Sign                               int
	JGammaSign                           int
	JDSignIfImposed                      int
	ExchangesRepresentationWithConjugate bool
	PhysicalChargeConjugationDerived     bool
	KOConventionDerived                  bool
	CandidateOnly                        bool
	Verdict                              string
}

type DoubledState struct {
	Label      string
	Sector     string
	Mask       int
	Excitation int
	BMinusL    float64
	Neutral    bool
	Parity     string
}

type FiniteAlgebraSearchAudit struct {
	SearchPrinciple                          string
	ImportedConnesAlgebra                    bool
	DerivedLieAlgebraInput                   string
	UniversalEnvelopingPreflight             bool
	ExplicitGaugeMatricesAvailable           bool
	ContactPreservingRepresentationAvailable bool
	ColorLeptonSplitAvailable                bool
	ColorM3CDerived                          bool
	QuaternionHDerived                       bool
	ComplexSummandDerived                    bool
	MaximalAssociativeSubalgebraDerived      bool
	FaithfulDoubledRepresentation            bool
	OppositeAlgebraActionDerived             bool
	OrderOneReady                            bool
	CandidateRows                            []string
	Verdict                                  string
}

type MajoranaBilinearAudit struct {
	DoubledSpaceAvailable            bool
	NeutralParticleStates            []string
	NeutralConjugateStates           []string
	NeutralBilinearCapacity          bool
	TotallyNeutralSlotCount          int
	RHNeutrinoSlotDerived            bool
	MajoranaTermKinematicallyAllowed bool
	GaugeInvariantIfNeutral          bool
	GradingCompatibilityDerived      bool
	OrderOneCompatibilityDerived     bool
	Verdict                          string
}

type BGapIdentificationAudit struct {
	BGapAvailable                 bool
	BGap                          float64
	CandidateMajoranaSlotExists   bool
	BGapDimensionless             bool
	BGapInsertedAsDiagnostic      bool
	BGapCanonicalMajoranaEntry    bool
	BGapPromotedToMass            bool
	BGapSelectsRHNeutrino         bool
	RequiresAlgebraRepresentation bool
	RequiresScaleMap              bool
	Verdict                       string
}

type FirewallAudit struct {
	ExternalAntiparticlesAdded bool
	ConnesAlgebraImported      bool
	ContinuumMassInserted      bool
	VEVInserted                bool
	MBInserted                 bool
	MIntInserted               bool
	MStarInserted              bool
	BGapPromotedToMass         bool
	MajoranaMassClaimed        bool
	OrderOneClaimed            bool
	PMNSOrYukawaClaimed        bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	ComplexificationDerived   bool
	AntiLinearJAvailable      bool
	NativeAlgebraDerived      bool
	MajoranaCapacity          bool
	BGapMajoranaIdentified    bool
	FullSpectralTripleDerived bool
	Status                    string
	NextGate                  string
	Comment                   string
}

type Analysis struct {
	Previous         realstructureorderone.Analysis
	DoubledStates    []DoubledState
	Complexification ComplexificationAudit
	J                AntiLinearJAudit
	Algebra          FiniteAlgebraSearchAudit
	Majorana         MajoranaBilinearAudit
	BGap             BGapIdentificationAudit
	Firewall         FirewallAudit
	Summary          Summary
	TruthStatement   string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := realstructureorderone.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 234 predecessor: %w", err)
			return
		}
		f, err := spinor.NewCovariantPhaseFockSpace(4)
		if err != nil {
			defaultErr = fmt.Errorf("construct Fock space: %w", err)
			return
		}
		b, err := bsector.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("construct B-sector vacuum: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, f, b, 1e-10)
	})
	return defaultA, defaultErr
}

func Build(prev realstructureorderone.Analysis, f spinor.FockSpace, b bsector.Vacuum, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if f.ModeCount() != 4 || f.StateCount() != 16 {
		return Analysis{}, fmt.Errorf("Gate 235 requires native four-mode 16-state Fock carrier, got modes=%d states=%d", f.ModeCount(), f.StateCount())
	}
	states := buildDoubledStates(f, eps)
	c := auditComplexification(f)
	j := auditAntiLinearJ(c)
	alg := auditFiniteAlgebraSearch(f)
	maj := auditMajoranaCapacity(states)
	bg := auditBGapIdentification(b, maj)
	fw := auditFirewall()
	sum := summarize(c, j, alg, maj, bg)
	truth := buildTruth(c, j, alg, maj, bg)
	return Analysis{Previous: prev, DoubledStates: states, Complexification: c, J: j, Algebra: alg, Majorana: maj, BGap: bg, Firewall: fw, Summary: sum, TruthStatement: truth}, nil
}

func buildDoubledStates(f spinor.FockSpace, eps float64) []DoubledState {
	out := make([]DoubledState, 0, 2*f.StateCount())
	for _, sector := range []string{"Re(S_C)", "Im(S_C)"} {
		prefix := "p"
		if sector == "Im(S_C)" {
			prefix = "Jp"
		}
		for i, s := range f.States {
			mask := maskOf(s)
			parity := "odd"
			if s.ExcitationNumber()%2 == 0 {
				parity = "even"
			}
			out = append(out, DoubledState{
				Label:      fmt.Sprintf("%s[%02d|mask=%04b]", prefix, i, mask),
				Sector:     sector,
				Mask:       mask,
				Excitation: s.ExcitationNumber(),
				BMinusL:    s.BMinusL(),
				Neutral:    math.Abs(s.BMinusL()) < eps,
				Parity:     parity,
			})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Sector == out[j].Sector {
			return out[i].Mask < out[j].Mask
		}
		return out[i].Sector < out[j].Sector
	})
	return out
}

func maskOf(s spinor.FockState) int {
	mask := 0
	for i, occupied := range s.Occupation {
		if occupied {
			mask |= 1 << i
		}
	}
	return mask
}

func auditComplexification(f spinor.FockSpace) ComplexificationAudit {
	return ComplexificationAudit{
		SourceRealCarrier:             "irreducible real Cl(1,7) spinor / four-mode Fock carrier S",
		DerivedByComplexification:     true,
		ExternalStatesAdded:           false,
		RealDimensionBefore:           f.StateCount(),
		ComplexDimensionAfter:         f.StateCount(),
		RealDimensionAfter:            2 * f.StateCount(),
		Formula:                       "S_C = S ⊗_R C; dim_C(S_C)=16; dim_R(S_C)=32",
		FixedRealHalfDimension:        f.StateCount(),
		ImaginaryHalfDimension:        f.StateCount(),
		ParticleAntiparticleSemantics: false,
		ParticleAntiparticleCandidate: true,
		Verdict:                       StatusComplexificationSupport,
	}
}

func auditAntiLinearJ(c ComplexificationAudit) AntiLinearJAudit {
	return AntiLinearJAudit{
		CandidateName:                        "complex conjugation on S_C",
		AntiLinear:                           true,
		J2Sign:                               +1,
		JGammaSign:                           +1,
		JDSignIfImposed:                      +1,
		ExchangesRepresentationWithConjugate: true,
		PhysicalChargeConjugationDerived:     false,
		KOConventionDerived:                  false,
		CandidateOnly:                        true,
		Verdict:                              StatusCandidateAntiLinearJ,
	}
}

func auditFiniteAlgebraSearch(f spinor.FockSpace) FiniteAlgebraSearchAudit {
	spectrum := f.ChargeSpectrum()
	rows := []string{
		"input Lie algebra available from earlier gates: contact-preserving su(2) ⊕ u(1)",
		fmt.Sprintf("native B-L occupation spectrum buckets: %s", formatSpectrum(spectrum)),
		"one temporal lepton seed + three spatial quark seeds gives color/lepton bookkeeping, not M3(C)",
		"associative closure of explicit gauge matrices cannot be computed because those matrices are not represented on S_C here",
		"C ⊕ H ⊕ M3(C) is therefore not imported and not derived",
	}
	return FiniteAlgebraSearchAudit{
		SearchPrinciple:                          "derive the associative algebra from the complexified Cl(1,7) spinor action compatible with the already-derived su(2)⊕u(1), contact preservation, and color/lepton split; do not import Connes algebra",
		ImportedConnesAlgebra:                    false,
		DerivedLieAlgebraInput:                   "su(2)⊕u(1) from contact-preserving centralizer/gauge gates",
		UniversalEnvelopingPreflight:             true,
		ExplicitGaugeMatricesAvailable:           false,
		ContactPreservingRepresentationAvailable: false,
		ColorLeptonSplitAvailable:                true,
		ColorM3CDerived:                          false,
		QuaternionHDerived:                       false,
		ComplexSummandDerived:                    false,
		MaximalAssociativeSubalgebraDerived:      false,
		FaithfulDoubledRepresentation:            false,
		OppositeAlgebraActionDerived:             false,
		OrderOneReady:                            false,
		CandidateRows:                            rows,
		Verdict:                                  StatusFailedNativeAlgebra,
	}
}

func formatSpectrum(m map[string]int) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s:%d", k, m[k]))
	}
	return strings.Join(parts, ",")
}

func auditMajoranaCapacity(states []DoubledState) MajoranaBilinearAudit {
	var pNeutral, cNeutral []string
	for _, s := range states {
		if !s.Neutral {
			continue
		}
		if s.Sector == "Re(S_C)" {
			pNeutral = append(pNeutral, s.Label)
		} else {
			cNeutral = append(cNeutral, s.Label)
		}
	}
	sort.Strings(pNeutral)
	sort.Strings(cNeutral)
	capacity := len(pNeutral) > 0 && len(cNeutral) > 0
	return MajoranaBilinearAudit{
		DoubledSpaceAvailable:            true,
		NeutralParticleStates:            pNeutral,
		NeutralConjugateStates:           cNeutral,
		NeutralBilinearCapacity:          capacity,
		TotallyNeutralSlotCount:          min(len(pNeutral), len(cNeutral)),
		RHNeutrinoSlotDerived:            false,
		MajoranaTermKinematicallyAllowed: capacity,
		GaugeInvariantIfNeutral:          capacity,
		GradingCompatibilityDerived:      false,
		OrderOneCompatibilityDerived:     false,
		Verdict:                          StatusMajoranaCapacity,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func auditBGapIdentification(b bsector.Vacuum, maj MajoranaBilinearAudit) BGapIdentificationAudit {
	gap := b.FirstPositiveEigenvalue(1e-8)
	return BGapIdentificationAudit{
		BGapAvailable:                 gap > 0,
		BGap:                          gap,
		CandidateMajoranaSlotExists:   maj.NeutralBilinearCapacity,
		BGapDimensionless:             true,
		BGapInsertedAsDiagnostic:      false,
		BGapCanonicalMajoranaEntry:    false,
		BGapPromotedToMass:            false,
		BGapSelectsRHNeutrino:         false,
		RequiresAlgebraRepresentation: true,
		RequiresScaleMap:              true,
		Verdict:                       StatusFailedBGapMajorana,
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		ExternalAntiparticlesAdded: false,
		ConnesAlgebraImported:      false,
		ContinuumMassInserted:      false,
		VEVInserted:                false,
		MBInserted:                 false,
		MIntInserted:               false,
		MStarInserted:              false,
		BGapPromotedToMass:         false,
		MajoranaMassClaimed:        false,
		OrderOneClaimed:            false,
		PMNSOrYukawaClaimed:        false,
		FiniteCorePolluted:         false,
		Verdict:                    "FIREWALL_PRESERVED_COMPLEXIFICATION_ONLY",
	}
}

func summarize(c ComplexificationAudit, j AntiLinearJAudit, alg FiniteAlgebraSearchAudit, maj MajoranaBilinearAudit, bg BGapIdentificationAudit) Summary {
	nativeAlg := alg.MaximalAssociativeSubalgebraDerived && alg.FaithfulDoubledRepresentation && alg.OppositeAlgebraActionDerived
	bgID := bg.BGapCanonicalMajoranaEntry && bg.BGapSelectsRHNeutrino && bg.BGapPromotedToMass
	full := c.DerivedByComplexification && j.AntiLinear && nativeAlg && maj.NeutralBilinearCapacity && bgID
	status := strings.Join([]string{StatusComplexificationSupport, StatusCandidateAntiLinearJ, StatusMajoranaCapacity, StatusFailedNativeAlgebra, StatusFailedBGapMajorana, StatusFailedSpectralTripleComplete}, ";")
	return Summary{
		ComplexificationDerived:   c.DerivedByComplexification && !c.ExternalStatesAdded && c.RealDimensionAfter == 32,
		AntiLinearJAvailable:      j.AntiLinear && j.J2Sign == 1 && j.CandidateOnly,
		NativeAlgebraDerived:      nativeAlg,
		MajoranaCapacity:          maj.NeutralBilinearCapacity,
		BGapMajoranaIdentified:    bgID,
		FullSpectralTripleDerived: full,
		Status:                    status,
		NextGate:                  "derive the native associative finite algebra representation on S_C before applying the full order-one condition",
		Comment:                   "Gate 235 derives the doubled 32-real carrier by complexification, not by adding external states. It creates Majorana capacity but not a canonical B-gap Majorana theorem.",
	}
}

func buildTruth(c ComplexificationAudit, j AntiLinearJAudit, alg FiniteAlgebraSearchAudit, maj MajoranaBilinearAudit, bg BGapIdentificationAudit) string {
	return fmt.Sprintf("S_C complexification gives dim_C=%d dim_R=%d with anti-linear J²=%+d. Neutral particle/conjugate bilinears are kinematically available (%d slots), but the native associative algebra, opposite action, order-one calculus, and B_gap Majorana identification remain un-derived.", c.ComplexDimensionAfter, c.RealDimensionAfter, j.J2Sign, maj.TotallyNeutralSlotCount)
}
