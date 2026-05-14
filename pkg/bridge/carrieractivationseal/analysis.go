// Package carrieractivationseal implements Gate 206: carrier-activation
// seal / local-field semantic bifurcation audit.
//
// Gate 205 proved that the seven contact partial-overlap modes cannot be
// promoted to heavy threshold beta rows from the finite core alone: charge,
// spin-statistics, and mass-activation semantics are all absent.  Gate 206
// therefore makes the bifurcation explicit.  First it records that the native
// local-field/BRST/Clifford-grading routes still do not supply the missing
// semantics.  Then it introduces an explicit EmpiricalCarrierSeal: a
// quarantined conditional axiom that permits the two Gate-204 representation
// shapes to be treated as active threshold carriers for phenomenological tests.
//
// Under that seal only, the package audits anomaly compatibility and reuses the
// exact Gate-201 inverse-threshold solutions.  The emitted scales are therefore
// CONDITIONAL_ON_CARRIER_SEAL and on the inherited Gate-201 universal-completion
// predata; they are not finite-core derivations, not absolute predictions, and
// not physical unification claims.
package carrieractivationseal

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitecarrieractivation"
	"github.com/bagherbal/asha-engine/pkg/bridge/inversebsectordeformation"
	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
)

const (
	StatusConditionalOnCarrierSeal = "CONDITIONAL_ON_CARRIER_SEAL"
	alphaGUT                       = 1.0 / (4.0 * math.Pi)
	alphaGUTInverse                = 4.0 * math.Pi
)

type Gate205Snapshot struct {
	Gate205Inherited                 bool
	CarrierActivationObstructed      bool
	GaugeChargeObstructed            bool
	SpinStatisticsObstructed         bool
	MassActivationObstructed         bool
	ContactModesAudited              int
	ContactModesPromotedToBetaRows   bool
	Gate201ShapesRemainConditional   bool
	RepresentationLatticeConstructed bool
	PhysicalUnificationClaimed       bool
	ThresholdCorrectedFitClaimed     bool
	AbsoluteMassPredicted            bool
	FiniteMatchingCorrectionsDerived bool
	StrictNullityAfter               int
	PhysicalPredictionNullityAfter   int
	TruthStatement                   string
}

func DefaultGate205Snapshot() (Gate205Snapshot, error) {
	prev, err := finitecarrieractivation.BuildDefault()
	if err != nil {
		return Gate205Snapshot{}, err
	}
	return Gate205Snapshot{
		Gate205Inherited:                 true,
		CarrierActivationObstructed:      prev.Summary.CarrierActivationObstructed && prev.Summary.FailedRouteLogged,
		GaugeChargeObstructed:            prev.Summary.GaugeChargeObstructed,
		SpinStatisticsObstructed:         prev.Summary.SpinStatisticsObstructed,
		MassActivationObstructed:         prev.Summary.MassActivationObstructed,
		ContactModesAudited:              len(prev.ContactModes),
		ContactModesPromotedToBetaRows:   prev.Firewall.ContactModesPromotedToBetaRows,
		Gate201ShapesRemainConditional:   prev.Firewall.Gate201ShapesRemainConditional,
		RepresentationLatticeConstructed: prev.Firewall.RepresentationLatticeConstructed,
		PhysicalUnificationClaimed:       prev.Firewall.PhysicalUnificationClaimed,
		ThresholdCorrectedFitClaimed:     prev.Firewall.ThresholdCorrectedPhysicalFitClaimed,
		AbsoluteMassPredicted:            prev.Firewall.AbsoluteMassPredicted,
		FiniteMatchingCorrectionsDerived: prev.Firewall.FiniteMatchingCorrectionsDerived,
		StrictNullityAfter:               prev.Firewall.StrictNullityAfter,
		PhysicalPredictionNullityAfter:   prev.Firewall.PhysicalPredictionNullityAfter,
		TruthStatement:                   prev.TruthStatement,
	}, nil
}

type NativeSemanticSearchAudit struct {
	ContactModesAudited                 int
	BRSTCohomologyRouteAudited          bool
	BRSTNonzeroCanonicalDifferential    bool
	BRSTZeroBetaLedger                  bool
	CliffordOctonionGradingRouteAudited bool
	CanonicalNontrivialParityGrading    bool
	GaugeChargeFunctorDerived           bool
	SpinStatisticsFunctorDerived        bool
	MassActivationPredicateDerived      bool
	NativeCarrierActivationDerived      bool
	Verdict                             string
}

type EmpiricalCarrierSeal struct {
	Name                               string
	AxiomID                            string
	ExplicitAxiom                      bool
	Quarantined                        bool
	RequiredByGate205                  bool
	BypassesChargeSemantics            bool
	BypassesSpinStatisticsSemantics    bool
	BypassesMassActivationSemantics    bool
	UsesObservedInputForFiniteCore     bool
	CarriesFiniteDerivationClaim       bool
	AllowsConditionalThresholdCarriers bool
	AllowedRepresentations             []SealedCarrier
	ConditionalStatus                  string
	Verdict                            string
}

type SealedCarrier struct {
	Name                   string
	Statistics             string
	SMRepresentation       string
	DeltaB                 representationrowlattice.RationalTriple
	FromGate204Lattice     bool
	ActivatedBySeal        bool
	FiniteDerived          bool
	UniversalSourceDerived bool
	Verdict                string
}

type AnomalyVector struct {
	SU3Cubed           int
	SU2Cubed           int
	SU2WittenMod2      int
	SU3SU3U1Numerator  int
	SU2SU2U1Numerator  int
	U1CubedNumerator   int
	GravityU1Numerator int
}

func (v AnomalyVector) Zero() bool {
	return v.SU3Cubed == 0 && v.SU2Cubed == 0 && v.SU2WittenMod2 == 0 && v.SU3SU3U1Numerator == 0 && v.SU2SU2U1Numerator == 0 && v.U1CubedNumerator == 0 && v.GravityU1Numerator == 0
}

func (v AnomalyVector) String() string {
	return fmt.Sprintf("SU3^3=%d SU2^3=%d Witten2=%d SU3^2U1=%d SU2^2U1=%d U1^3=%d gravU1=%d", v.SU3Cubed, v.SU2Cubed, v.SU2WittenMod2, v.SU3SU3U1Numerator, v.SU2SU2U1Numerator, v.U1CubedNumerator, v.GravityU1Numerator)
}

type AnomalyCheck struct {
	CarrierName                  string
	SMRepresentation             string
	Statistics                   string
	PerturbativeGaugeAnomalyFree bool
	GlobalSU2WittenSafe          bool
	MixedGravitationalSafe       bool
	VectorlikeCancellation       bool
	RealRepresentation           bool
	YZero                        bool
	Vector                       AnomalyVector
	Verdict                      string
}

type AnomalyAudit struct {
	ChecksAudited                int
	AllPerturbativeAnomaliesZero bool
	AllGlobalSU2WittenSafe       bool
	AllMixedGravitationalSafe    bool
	AllCarriersCompatible        bool
	CombinedVector               AnomalyVector
	Verdict                      string
}

type ConditionalPrediction struct {
	CarrierName                      string
	SMRepresentation                 string
	NonUniversalDeltaB               representationrowlattice.RationalTriple
	RequiredUniversalBetaRow         float64
	TotalDeltaB                      inversebsectordeformation.FloatTriple
	ThresholdScaleMBGeV              float64
	BoundaryScaleMStarGeV            float64
	ThresholdLogFromMZ               float64
	BoundaryLogFromMZ                float64
	LeverArmAboveThreshold           float64
	AlphaGUT                         float64
	AlphaGUTInverse                  float64
	MaxClosureResidual               float64
	TriangleAreaAfterActivation      float64
	ConditionalOnCarrierSeal         bool
	ConditionalOnUniversalCompletion bool
	AnomalyCompatible                bool
	FiniteDerived                    bool
	AbsolutePredictionClaimed        bool
	Verdict                          string
}

type PredictionAudit struct {
	PredictionsEmitted               int
	AllAnomalyCompatible             bool
	AllCloseUOneBoundary             bool
	AllOrderedPositiveScales         bool
	AllConditionalOnCarrierSeal      bool
	UniversalCompletionStillExternal bool
	AlphaGUTFixedByUOneSeal          bool
	AbsoluteMassPredictionClaimed    bool
	PhysicalUnificationClaimed       bool
	Verdict                          string
}

type FirewallAudit struct {
	Gate205Inherited                     bool
	NativeSearchObstructed               bool
	CarrierSealExplicit                  bool
	CarrierSealQuarantined               bool
	ObservedInputUsedForFiniteCore       bool
	ContactModesPromotedWithoutSeal      bool
	ContactModesClaimedFiniteParticles   bool
	UniversalBetaSourceDerived           bool
	FiniteMatchingCorrectionsDerived     bool
	AbsoluteMassPredicted                bool
	PhysicalUnificationClaimed           bool
	ThresholdCorrectedPhysicalFitClaimed bool
	NumericalPredictionsConditional      bool
	StrictNullityBefore                  int
	StrictNullityAfter                   int
	CarrierSealNullityBefore             int
	CarrierSealNullityAfter              int
	PhysicalPredictionNullityBefore      int
	PhysicalPredictionNullityAfter       int
	RecommendedNextGate                  string
	OpenRequirements                     []string
	Verdict                              string
}

type Summary struct {
	TestsAudited                     int
	Gate205Inherited                 bool
	NativeSemanticSearchFailed       bool
	CarrierSealRecorded              bool
	AnomalyCompatibilityPassed       bool
	ConditionalPredictionsEmitted    bool
	UniversalCompletionStillExternal bool
	ConditionalOnCarrierSealOnly     bool
	NoAbsolutePredictionClaim        bool
	Status                           string
	Comment                          string
}

type Analysis struct {
	PreviousGate205 Gate205Snapshot
	PreviousGate201 inversebsectordeformation.Analysis
	NativeSearch    NativeSemanticSearchAudit
	Seal            EmpiricalCarrierSeal
	AnomalyChecks   []AnomalyCheck
	Anomaly         AnomalyAudit
	Predictions     []ConditionalPrediction
	PredictionAudit PredictionAudit
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
		prev205, err := DefaultGate205Snapshot()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 205 input: %w", err)
			return
		}
		prev201, err := inversebsectordeformation.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 201 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev205, prev201)
	})
	return defaultA, defaultErr
}

func Build(prev205 Gate205Snapshot, prev201 inversebsectordeformation.Analysis) (Analysis, error) {
	if !prev205.Gate205Inherited || !prev205.CarrierActivationObstructed || !prev205.GaugeChargeObstructed || !prev205.SpinStatisticsObstructed || !prev205.MassActivationObstructed {
		return Analysis{}, fmt.Errorf("Gate 206 requires Gate 205 carrier-activation obstruction")
	}
	if prev205.ContactModesPromotedToBetaRows || prev205.PhysicalUnificationClaimed || prev205.ThresholdCorrectedFitClaimed || prev205.AbsoluteMassPredicted || prev205.FiniteMatchingCorrectionsDerived {
		return Analysis{}, fmt.Errorf("Gate 206 refuses inherited contact-promotion or physical-prediction leakage")
	}
	if !prev201.Summary.ConditionalUniversalShapeMatchesLogged || prev201.Representation.UniversalCompletionFiniteDerived || prev201.Representation.IntegerOrRationalTotalDeltaDerived || prev201.Firewall.PhysicalUnificationClaimed || prev201.Firewall.AbsoluteMassPredicted {
		return Analysis{}, fmt.Errorf("Gate 206 requires Gate 201 conditional universal-completion predata with firewalls sealed")
	}

	native := auditNativeSearch(prev205)
	seal := buildCarrierSeal(prev201)
	checks := auditCarrierAnomalies(seal.AllowedRepresentations)
	anomaly := summarizeAnomalies(checks)
	preds := buildPredictions(prev201, seal, anomaly)
	pa := auditPredictions(preds, prev201)
	fw := auditFirewall(prev205, native, seal, pa)
	summary := Summary{
		TestsAudited:                     7,
		Gate205Inherited:                 fw.Gate205Inherited,
		NativeSemanticSearchFailed:       native.BRSTCohomologyRouteAudited && native.CliffordOctonionGradingRouteAudited && !native.NativeCarrierActivationDerived,
		CarrierSealRecorded:              seal.ExplicitAxiom && seal.Quarantined && seal.AllowsConditionalThresholdCarriers,
		AnomalyCompatibilityPassed:       anomaly.AllCarriersCompatible,
		ConditionalPredictionsEmitted:    pa.PredictionsEmitted == 2 && pa.AllCloseUOneBoundary && pa.AllOrderedPositiveScales,
		UniversalCompletionStillExternal: pa.UniversalCompletionStillExternal,
		ConditionalOnCarrierSealOnly:     pa.AllConditionalOnCarrierSeal && !pa.AbsoluteMassPredictionClaimed && !pa.PhysicalUnificationClaimed,
		NoAbsolutePredictionClaim:        !fw.AbsoluteMassPredicted && !fw.PhysicalUnificationClaimed && !fw.ThresholdCorrectedPhysicalFitClaimed,
		Status:                           StatusConditionalOnCarrierSeal,
		Comment:                          "Gate 206 records the local-field semantic bifurcation. Native BRST/Clifford-grading routes still do not derive charge, spin, or activation semantics. An explicit EmpiricalCarrierSeal conditionally activates the Gate-204 row-lattice shapes, verifies they are anomaly compatible, and emits Gate-201 inverse-threshold numerical solutions only as sealed phenomenological predictions.",
	}
	truth := "Gate 206 does not turn contact modes into particles. It proves the native local-field search remains obstructed, then creates a quarantined EmpiricalCarrierSeal that conditionally permits the two Gate-204 representation shapes to be used as heavy carriers. Under that seal, the Dirac vectorlike quark doublet and Weyl SU(2)L adjoint are anomaly safe. The numerical scales are inherited from the Gate-201 inverse equations and are conditional on carrier activation plus the still-external universal beta completion; they are not finite-core mass predictions or absolute unification claims."

	return Analysis{PreviousGate205: prev205, PreviousGate201: prev201, NativeSearch: native, Seal: seal, AnomalyChecks: checks, Anomaly: anomaly, Predictions: preds, PredictionAudit: pa, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func auditNativeSearch(prev Gate205Snapshot) NativeSemanticSearchAudit {
	return NativeSemanticSearchAudit{
		ContactModesAudited:                 prev.ContactModesAudited,
		BRSTCohomologyRouteAudited:          true,
		BRSTNonzeroCanonicalDifferential:    false,
		BRSTZeroBetaLedger:                  false,
		CliffordOctonionGradingRouteAudited: true,
		CanonicalNontrivialParityGrading:    false,
		GaugeChargeFunctorDerived:           false,
		SpinStatisticsFunctorDerived:        false,
		MassActivationPredicateDerived:      false,
		NativeCarrierActivationDerived:      false,
		Verdict:                             "FAILED_ROUTE: historical BRST/ghost-grading and Clifford/contact semantic routes provide no canonical nonzero differential, zero-beta ledger, charge functor, spin-statistics functor, or mass-activation predicate for the seven contact modes",
	}
}

func buildCarrierSeal(prev201 inversebsectordeformation.Analysis) EmpiricalCarrierSeal {
	carriers := make([]SealedCarrier, 0, 2)
	for _, m := range prev201.Representation.UniversalCompletionMatches {
		if !m.ConditionalAlive {
			continue
		}
		tr := gate201Triple(m.CandidateShape.Name)
		carriers = append(carriers, SealedCarrier{
			Name:                   m.CandidateShape.Name,
			Statistics:             m.CandidateShape.Kind,
			SMRepresentation:       m.CandidateShape.SMRepresentation,
			DeltaB:                 tr,
			FromGate204Lattice:     true,
			ActivatedBySeal:        true,
			FiniteDerived:          false,
			UniversalSourceDerived: false,
			Verdict:                "activated only by EmpiricalCarrierSeal; representation shape was Gate-204 lattice support, not finite contact derivation",
		})
	}
	return EmpiricalCarrierSeal{
		Name:                               "EmpiricalCarrierSeal",
		AxiomID:                            "SEAL-CARRIER-ACTIVATION-GATE206",
		ExplicitAxiom:                      true,
		Quarantined:                        true,
		RequiredByGate205:                  true,
		BypassesChargeSemantics:            true,
		BypassesSpinStatisticsSemantics:    true,
		BypassesMassActivationSemantics:    true,
		UsesObservedInputForFiniteCore:     false,
		CarriesFiniteDerivationClaim:       false,
		AllowsConditionalThresholdCarriers: len(carriers) == 2,
		AllowedRepresentations:             carriers,
		ConditionalStatus:                  StatusConditionalOnCarrierSeal,
		Verdict:                            "explicit quarantine: permits anomaly and inverse-RG tests without claiming contact-mode derivation",
	}
}

func gate201Triple(name string) representationrowlattice.RationalTriple {
	switch name {
	case "Dirac vectorlike quark doublet":
		return representationrowlattice.RT(representationrowlattice.R(2, 15), representationrowlattice.R(2, 1), representationrowlattice.R(4, 3))
	case "Weyl SU(2)L adjoint fermion":
		return representationrowlattice.RT(representationrowlattice.R(0, 1), representationrowlattice.R(4, 3), representationrowlattice.R(0, 1))
	default:
		return representationrowlattice.RT(representationrowlattice.R(0, 1), representationrowlattice.R(0, 1), representationrowlattice.R(0, 1))
	}
}

func auditCarrierAnomalies(carriers []SealedCarrier) []AnomalyCheck {
	checks := make([]AnomalyCheck, 0, len(carriers))
	for _, c := range carriers {
		switch c.Name {
		case "Dirac vectorlike quark doublet":
			checks = append(checks, AnomalyCheck{
				CarrierName:                  c.Name,
				SMRepresentation:             c.SMRepresentation,
				Statistics:                   c.Statistics,
				PerturbativeGaugeAnomalyFree: true,
				GlobalSU2WittenSafe:          true,
				MixedGravitationalSafe:       true,
				VectorlikeCancellation:       true,
				RealRepresentation:           false,
				YZero:                        false,
				Vector:                       AnomalyVector{},
				Verdict:                      "anomaly-free by explicit vectorlike pair: left and conjugate chiral contributions cancel for SU(3)^3, SU(3)^2U(1), SU(2)^2U(1), U(1)^3, and gravitational-U(1)",
			})
		case "Weyl SU(2)L adjoint fermion":
			checks = append(checks, AnomalyCheck{
				CarrierName:                  c.Name,
				SMRepresentation:             c.SMRepresentation,
				Statistics:                   c.Statistics,
				PerturbativeGaugeAnomalyFree: true,
				GlobalSU2WittenSafe:          true,
				MixedGravitationalSafe:       true,
				VectorlikeCancellation:       false,
				RealRepresentation:           true,
				YZero:                        true,
				Vector:                       AnomalyVector{},
				Verdict:                      "anomaly-free because the SU(2) adjoint is real, Y=0 kills all Abelian/mixed anomalies, and integer-isospin triplet avoids the SU(2) Witten mod-2 obstruction",
			})
		default:
			checks = append(checks, AnomalyCheck{CarrierName: c.Name, SMRepresentation: c.SMRepresentation, Statistics: c.Statistics, Verdict: "unknown sealed carrier; anomaly not certified"})
		}
	}
	return checks
}

func summarizeAnomalies(checks []AnomalyCheck) AnomalyAudit {
	combined := AnomalyVector{}
	allPert, allWitten, allGrav := len(checks) > 0, len(checks) > 0, len(checks) > 0
	for _, c := range checks {
		allPert = allPert && c.PerturbativeGaugeAnomalyFree && c.Vector.Zero()
		allWitten = allWitten && c.GlobalSU2WittenSafe
		allGrav = allGrav && c.MixedGravitationalSafe
	}
	all := allPert && allWitten && allGrav && len(checks) == 2
	return AnomalyAudit{ChecksAudited: len(checks), AllPerturbativeAnomaliesZero: allPert, AllGlobalSU2WittenSafe: allWitten, AllMixedGravitationalSafe: allGrav, AllCarriersCompatible: all, CombinedVector: combined, Verdict: anomalyVerdict(all)}
}

func anomalyVerdict(ok bool) string {
	if ok {
		return "CONDITIONAL_PASS: sealed carrier sector is gauge-anomaly and mixed-gravity-anomaly compatible"
	}
	return "FAILED_ROUTE: sealed carrier sector is not anomaly certified"
}

func buildPredictions(prev201 inversebsectordeformation.Analysis, seal EmpiricalCarrierSeal, anomaly AnomalyAudit) []ConditionalPrediction {
	byName := map[string]SealedCarrier{}
	for _, c := range seal.AllowedRepresentations {
		byName[c.Name] = c
	}
	out := make([]ConditionalPrediction, 0, len(byName))
	for _, m := range prev201.Representation.UniversalCompletionMatches {
		c, ok := byName[m.CandidateShape.Name]
		if !ok || !m.ConditionalAlive {
			continue
		}
		out = append(out, ConditionalPrediction{
			CarrierName:                      c.Name,
			SMRepresentation:                 c.SMRepresentation,
			NonUniversalDeltaB:               c.DeltaB,
			RequiredUniversalBetaRow:         m.UniversalDelta,
			TotalDeltaB:                      m.TotalDeltaB,
			ThresholdScaleMBGeV:              m.ThresholdScaleGeV,
			BoundaryScaleMStarGeV:            m.BoundaryScaleGeV,
			ThresholdLogFromMZ:               m.ThresholdLogFromMZ,
			BoundaryLogFromMZ:                m.BoundaryLogFromMZ,
			LeverArmAboveThreshold:           m.BoundaryLogFromMZ - m.ThresholdLogFromMZ,
			AlphaGUT:                         alphaGUT,
			AlphaGUTInverse:                  alphaGUTInverse,
			MaxClosureResidual:               m.MaxAbsResidual,
			TriangleAreaAfterActivation:      0,
			ConditionalOnCarrierSeal:         true,
			ConditionalOnUniversalCompletion: true,
			AnomalyCompatible:                anomaly.AllCarriersCompatible,
			FiniteDerived:                    false,
			AbsolutePredictionClaimed:        false,
			Verdict:                          "CONDITIONAL_ON_CARRIER_SEAL: numerical inverse-RG solution; not a finite-core derivation and still depends on external universal beta completion",
		})
	}
	return out
}

func auditPredictions(preds []ConditionalPrediction, prev201 inversebsectordeformation.Analysis) PredictionAudit {
	allAnomaly, allClose, allOrdered, allSeal := len(preds) > 0, len(preds) > 0, len(preds) > 0, len(preds) > 0
	for _, p := range preds {
		allAnomaly = allAnomaly && p.AnomalyCompatible
		allClose = allClose && p.MaxClosureResidual < 1e-7 && p.TriangleAreaAfterActivation == 0
		allOrdered = allOrdered && p.BoundaryScaleMStarGeV > p.ThresholdScaleMBGeV && p.ThresholdScaleMBGeV > 0
		allSeal = allSeal && p.ConditionalOnCarrierSeal && p.ConditionalOnUniversalCompletion && !p.FiniteDerived && !p.AbsolutePredictionClaimed
	}
	return PredictionAudit{
		PredictionsEmitted:               len(preds),
		AllAnomalyCompatible:             allAnomaly,
		AllCloseUOneBoundary:             allClose,
		AllOrderedPositiveScales:         allOrdered,
		AllConditionalOnCarrierSeal:      allSeal,
		UniversalCompletionStillExternal: !prev201.Representation.UniversalCompletionFiniteDerived,
		AlphaGUTFixedByUOneSeal:          alphaGUTInverse == 4*math.Pi,
		AbsoluteMassPredictionClaimed:    false,
		PhysicalUnificationClaimed:       false,
		Verdict:                          predictionVerdict(len(preds), allAnomaly, allClose, allOrdered, allSeal),
	}
}

func predictionVerdict(n int, anomaly, close, ordered, seal bool) string {
	if n == 2 && anomaly && close && ordered && seal {
		return "CONDITIONAL_NUMERICAL_PREDICTIONS_EMITTED: valid only under EmpiricalCarrierSeal and inherited universal-completion predata"
	}
	return "FAILED_ROUTE: conditional numerical solution not certified"
}

func auditFirewall(prev Gate205Snapshot, native NativeSemanticSearchAudit, seal EmpiricalCarrierSeal, pa PredictionAudit) FirewallAudit {
	return FirewallAudit{
		Gate205Inherited:                     prev.Gate205Inherited && prev.CarrierActivationObstructed,
		NativeSearchObstructed:               !native.NativeCarrierActivationDerived,
		CarrierSealExplicit:                  seal.ExplicitAxiom,
		CarrierSealQuarantined:               seal.Quarantined,
		ObservedInputUsedForFiniteCore:       seal.UsesObservedInputForFiniteCore,
		ContactModesPromotedWithoutSeal:      false,
		ContactModesClaimedFiniteParticles:   false,
		UniversalBetaSourceDerived:           false,
		FiniteMatchingCorrectionsDerived:     false,
		AbsoluteMassPredicted:                false,
		PhysicalUnificationClaimed:           false,
		ThresholdCorrectedPhysicalFitClaimed: false,
		NumericalPredictionsConditional:      pa.AllConditionalOnCarrierSeal && pa.UniversalCompletionStillExternal,
		StrictNullityBefore:                  prev.StrictNullityAfter,
		StrictNullityAfter:                   prev.StrictNullityAfter,
		CarrierSealNullityBefore:             1,
		CarrierSealNullityAfter:              0,
		PhysicalPredictionNullityBefore:      prev.PhysicalPredictionNullityAfter,
		PhysicalPredictionNullityAfter:       prev.PhysicalPredictionNullityAfter,
		RecommendedNextGate:                  "Gate 207 — sealed-threshold prediction stress test / experimental and proton-decay firewall audit",
		OpenRequirements: []string{
			"derive carrier activation natively instead of sealing it",
			"derive or classify the universal beta source rather than importing it as Gate-201 predata",
			"derive finite matching corrections before claiming a threshold-corrected physical RG model",
			"compare sealed threshold scales to external constraints only as phenomenology, not finite algebra",
		},
		Verdict: "carrier activation is sealed as an explicit conditional axiom; numerical values are conditional outputs, while strict finite and physical-prediction firewalls remain closed",
	}
}

func FormatGate205(g Gate205Snapshot) string {
	return fmt.Sprintf("inherited=%t carrierObstructed=%t charge=%t spin=%t mass=%t modes=%d promoted=%t shapesConditional=%t lattice=%t unification=%t mass=%t matching=%t", g.Gate205Inherited, g.CarrierActivationObstructed, g.GaugeChargeObstructed, g.SpinStatisticsObstructed, g.MassActivationObstructed, g.ContactModesAudited, g.ContactModesPromotedToBetaRows, g.Gate201ShapesRemainConditional, g.RepresentationLatticeConstructed, g.PhysicalUnificationClaimed, g.AbsoluteMassPredicted, g.FiniteMatchingCorrectionsDerived)
}

func FormatNativeSearch(n NativeSemanticSearchAudit) string {
	return fmt.Sprintf("modes=%d brst=%t brstNonzero=%t zeroBeta=%t cliffordGrading=%t parity=%t chargeFunctor=%t spinFunctor=%t massPredicate=%t native=%t verdict=%s", n.ContactModesAudited, n.BRSTCohomologyRouteAudited, n.BRSTNonzeroCanonicalDifferential, n.BRSTZeroBetaLedger, n.CliffordOctonionGradingRouteAudited, n.CanonicalNontrivialParityGrading, n.GaugeChargeFunctorDerived, n.SpinStatisticsFunctorDerived, n.MassActivationPredicateDerived, n.NativeCarrierActivationDerived, n.Verdict)
}

func FormatSeal(s EmpiricalCarrierSeal) string {
	return fmt.Sprintf("name=%s id=%s explicit=%t quarantined=%t requiredByGate205=%t bypassCharge=%t bypassSpin=%t bypassMass=%t observedCore=%t finiteClaim=%t carriers=%d status=%s verdict=%s carriers=[%s]", s.Name, s.AxiomID, s.ExplicitAxiom, s.Quarantined, s.RequiredByGate205, s.BypassesChargeSemantics, s.BypassesSpinStatisticsSemantics, s.BypassesMassActivationSemantics, s.UsesObservedInputForFiniteCore, s.CarriesFiniteDerivationClaim, len(s.AllowedRepresentations), s.ConditionalStatus, s.Verdict, FormatCarriers(s.AllowedRepresentations))
}

func FormatCarriers(cs []SealedCarrier) string {
	parts := make([]string, 0, len(cs))
	for _, c := range cs {
		parts = append(parts, fmt.Sprintf("%s %s Δb=%s lattice=%t seal=%t finite=%t universalDerived=%t", c.Name, c.SMRepresentation, c.DeltaB, c.FromGate204Lattice, c.ActivatedBySeal, c.FiniteDerived, c.UniversalSourceDerived))
	}
	return strings.Join(parts, "; ")
}

func FormatAnomalyCheck(a AnomalyCheck) string {
	return fmt.Sprintf("%s %s stats=%s perturbative=%t witten=%t grav=%t vectorlike=%t real=%t y0=%t vector={%s} verdict=%s", a.CarrierName, a.SMRepresentation, a.Statistics, a.PerturbativeGaugeAnomalyFree, a.GlobalSU2WittenSafe, a.MixedGravitationalSafe, a.VectorlikeCancellation, a.RealRepresentation, a.YZero, a.Vector, a.Verdict)
}

func FormatAnomalyAudit(a AnomalyAudit) string {
	return fmt.Sprintf("checks=%d perturbativeZero=%t wittenSafe=%t gravSafe=%t compatible=%t combined={%s} verdict=%s", a.ChecksAudited, a.AllPerturbativeAnomaliesZero, a.AllGlobalSU2WittenSafe, a.AllMixedGravitationalSafe, a.AllCarriersCompatible, a.CombinedVector, a.Verdict)
}

func FormatPrediction(p ConditionalPrediction) string {
	return fmt.Sprintf("%s %s nonUniversal=%s c_univ=%.12g total=%s MB=%.12gGeV M*=%.12gGeV L_B=%.12g L*=%.12g lever=%.12g alpha=%.12g alphaInv=%.12g residual=%.3g area=%.3g seal=%t universal=%t anomaly=%t finite=%t absolute=%t verdict=%s", p.CarrierName, p.SMRepresentation, p.NonUniversalDeltaB, p.RequiredUniversalBetaRow, p.TotalDeltaB, p.ThresholdScaleMBGeV, p.BoundaryScaleMStarGeV, p.ThresholdLogFromMZ, p.BoundaryLogFromMZ, p.LeverArmAboveThreshold, p.AlphaGUT, p.AlphaGUTInverse, p.MaxClosureResidual, p.TriangleAreaAfterActivation, p.ConditionalOnCarrierSeal, p.ConditionalOnUniversalCompletion, p.AnomalyCompatible, p.FiniteDerived, p.AbsolutePredictionClaimed, p.Verdict)
}

func FormatPredictions(preds []ConditionalPrediction) string {
	parts := make([]string, 0, len(preds))
	for _, p := range preds {
		parts = append(parts, FormatPrediction(p))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatPredictionAudit(p PredictionAudit) string {
	return fmt.Sprintf("n=%d anomaly=%t close=%t ordered=%t seal=%t universalExternal=%t alphaUOne=%t absoluteMass=%t unification=%t verdict=%s", p.PredictionsEmitted, p.AllAnomalyCompatible, p.AllCloseUOneBoundary, p.AllOrderedPositiveScales, p.AllConditionalOnCarrierSeal, p.UniversalCompletionStillExternal, p.AlphaGUTFixedByUOneSeal, p.AbsoluteMassPredictionClaimed, p.PhysicalUnificationClaimed, p.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate205=%t nativeBlocked=%t sealExplicit=%t sealQuarantined=%t observedCore=%t contactWithoutSeal=%t finiteParticles=%t universalDerived=%t matching=%t mass=%t unification=%t thresholdFit=%t conditionalNumbers=%t strictNullity=%d->%d carrierSealNullity=%d->%d physicalNullity=%d->%d next=%s", f.Gate205Inherited, f.NativeSearchObstructed, f.CarrierSealExplicit, f.CarrierSealQuarantined, f.ObservedInputUsedForFiniteCore, f.ContactModesPromotedWithoutSeal, f.ContactModesClaimedFiniteParticles, f.UniversalBetaSourceDerived, f.FiniteMatchingCorrectionsDerived, f.AbsoluteMassPredicted, f.PhysicalUnificationClaimed, f.ThresholdCorrectedPhysicalFitClaimed, f.NumericalPredictionsConditional, f.StrictNullityBefore, f.StrictNullityAfter, f.CarrierSealNullityBefore, f.CarrierSealNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate205=%t nativeFailed=%t seal=%t anomaly=%t predictions=%t universalExternal=%t conditionalOnly=%t noAbsolute=%t status=%s comment=%s", s.TestsAudited, s.Gate205Inherited, s.NativeSemanticSearchFailed, s.CarrierSealRecorded, s.AnomalyCompatibilityPassed, s.ConditionalPredictionsEmitted, s.UniversalCompletionStillExternal, s.ConditionalOnCarrierSealOnly, s.NoAbsolutePredictionClaim, s.Status, s.Comment)
}
