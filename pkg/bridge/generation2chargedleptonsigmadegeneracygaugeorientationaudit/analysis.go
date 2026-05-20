// Package generation2chargedleptonsigmadegeneracygaugeorientationaudit implements
// Gate 603: Charged-Lepton Sigma Degeneracy Gauge-or-Orientation Audit.
//
// Gate 602 unsealed the lepton wall and showed that B_flav selects the
// electron row, the third neutrino projector, and the positive CKM orientation
// sign, while leaving a sixfold charged-lepton sigma/cyclic-order degeneracy.
// Gate 603 audits whether that residual sigma label is physical signed
// discriminant data or only Fourier-coordinate redundancy for the balance.
package generation2chargedleptonsigmadegeneracygaugeorientationaudit

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2unsealedleptonwallpmnsrowbranchselectoraudit"
)

const (
	AuditID = "GATE603-CHARGED-LEPTON-SIGMA-DEGENERACY-GAUGE-ORIENTATION-AUDIT"

	StatusGate602Inherited                    = "PASS_GATE602_UNSEALED_BRANCH_ROW_SELECTOR_INHERITED"
	StatusSigmaDegeneracySourceIdentified     = "PASS_SIGMA_DEGENERACY_SOURCE_IDENTIFIED"
	StatusS3ActionAudited                     = "PASS_S3_ACTION_ON_KOIDE_FOURIER_COORDINATES_AUDITED"
	StatusPhysicalLabelsAudited               = "PASS_PHYSICAL_CHARGED_LEPTON_LABEL_ORDERING_AUDITED"
	StatusSignedDiscriminantAudited           = "PASS_SIGNED_DISCRIMINANT_AND_VANDERMONDE_ORIENTATION_AUDITED"
	StatusFourierCyclicOrientationAudited     = "PASS_FOURIER_CYCLIC_ORIENTATION_AUDITED"
	StatusPMNSCKMOrientationCouplingAudited   = "PASS_PMNS_CKM_ORIENTATION_COUPLING_AUDITED"
	StatusSigmaGaugeRedundancy                = "CONDITIONAL_SUPPORT_SIGMA_IS_FOURIER_COORDINATE_REDUNDANCY_FOR_B_FLAV"
	StatusDiscriminantOrientationSealRequired = "CONDITIONAL_SUPPORT_CHARGED_LEPTON_DISCRIMINANT_ORIENTATION_SEAL_REQUIRED_FOR_FULL_ORDER_SELECTION"
	StatusBFlavDoesNotSeeCyclicSigma          = "FAILED_ROUTE_B_FLAV_DOES_NOT_SEE_CYCLIC_SIGMA"
	StatusNoNativeSignedDiscriminantTheorem   = "FAILED_ROUTE_NO_NATIVE_SIGNED_DISCRIMINANT_ORIENTATION_THEOREM"
	StatusNoNativeMassOrderingTheorem         = "FAILED_ROUTE_NO_NATIVE_CHARGED_LEPTON_MASS_ORDERING_THEOREM"
	StatusNoNativeSigmaSelectionTheorem       = "FAILED_ROUTE_NO_NATIVE_SIGMA_SELECTION_THEOREM"
	StatusNoNativeBFlavZero                   = "FAILED_ROUTE_NO_NATIVE_B_FLAV_ZERO_THEOREM"
	StatusGate352Preserved                    = "FIREWALL_PRESERVED_GATE352_ROOT_TRACE_OBSTRUCTION_REMAINS_BINDING"
	StatusGate596Preserved                    = "FIREWALL_PRESERVED_GATE596_FOURTH_ROOT_OBSTRUCTION_REMAINS_BINDING"
	StatusGate600Preserved                    = "FIREWALL_PRESERVED_GATE600_BRANCH_CHAMBER_BOUNDARY_REMAINS_BINDING"
	StatusGate602Preserved                    = "FIREWALL_PRESERVED_GATE602_UNSEALED_LEPTON_WALL_BOUNDARY_REMAINS_BINDING"
	StatusNoKoideDerivation                   = "FIREWALL_PRESERVED_NO_KOIDE_DERIVATION"
	StatusNoMassDerivation                    = "FIREWALL_PRESERVED_NO_CHARGED_LEPTON_MASS_DERIVATION"
	StatusNoPMNSCKMDerivation                 = "FIREWALL_PRESERVED_NO_PMNS_CKM_NEUTRINO_OR_FLAVOR_DERIVATION"
	StatusNoBFlavNativePromotion              = "FIREWALL_PRESERVED_NO_B_FLAV_ZERO_NATIVE_PROMOTION"
	StatusNoNewCarrierSelector                = "FIREWALL_PRESERVED_NO_NEW_CARRIER_OR_SELECTOR_ADDED"
	StatusGate603Boundary                     = "FIREWALL_PRESERVED_GATE603_SIGMA_GAUGE_ORIENTATION_BOUNDARY"
)

const (
	yElectron   = 2.9350283095504176e-06
	yMuon       = 0.0006068707640859305
	yTau        = 0.010205763440624986
	sin2Theta12 = 0.308
	sin2Theta23 = 0.470
	sin2Theta13 = 0.02215
	deltaCPDeg  = 212.0
	jCKM        = 3.1169935287554706e-05
)

type InheritedGate602 struct {
	SelectsElectronRow   bool
	SelectsP3            bool
	SelectsPositiveJ     bool
	SelectsFullSigma     bool
	MinimalClassSize     int
	BestResidual         float64
	NextDistinctResidual float64
	Verdict              string
}

type S3ActionRow struct {
	Sigma                  string
	Order                  []string
	DeltaDeg               float64
	R                      float64
	Q                      float64
	ElectronWallEpsilonDeg float64
	ElectronWallEpsilonRad float64
	CanonicalAscending     bool
	VandermondeX           float64
	SignVandermondeX       int
	OrientationParity      string
	BFlavInvariantValue    float64
	Verdict                string
}

type InvariantVsOrientationSensitive struct {
	Quantity             string
	InvariantUnderS3     bool
	OrientationSensitive bool
	Native               bool
	Explanation          string
	Verdict              string
}

type PhysicalLabelAudit struct {
	ObservedOrdering      string
	YukawaOrdering        string
	NativeOrdering        bool
	EnvironmentalOrdering bool
	Verdict               string
}

type SignedDiscriminantAudit struct {
	DeltaFormula                 string
	VandermondeFormula           string
	RootVandermondeFormula       string
	TraceRingSuppliesDeltaOnly   bool
	SignedVRequiresOrdering      bool
	SignDistinguishesOrientation bool
	NativeSignedVTheorem         bool
	MinimalSeal                  string
	Verdict                      string
}

type FourierCyclicOrientationAudit struct {
	RequiresCyclicConvention        bool
	CanonicalCycle                  string
	ReversedCycle                   string
	BFlavUsesUnsignedWallDistance   bool
	BFlavDependsOnCyclicOrientation bool
	Verdict                         string
}

type PMNSCKMOrientationCouplingAudit struct {
	JCKMSign                 int
	JPMNSSign                int
	CanonicalVxSign          int
	Candidate1               string
	Candidate1Value          int
	Candidate2               string
	Candidate2Value          int
	TypedASHAOperatorPresent bool
	Verdict                  string
}

type MinimalRemainingSeal struct {
	SigmaGaugeForBFlav               bool
	PhysicalFullOrderingRequiresSeal bool
	SealName                         string
	SealData                         []string
	Statement                        string
	Verdict                          string
}

type Firewalls struct {
	DerivesKoide               bool
	DerivesChargedLeptonMasses bool
	DerivesPMNS                bool
	DerivesCKM                 bool
	DerivesBFlavZero           bool
	AddsCarrier                bool
	AddsSelector               bool
	PreservesGate352           bool
	PreservesGate596           bool
	PreservesGate600           bool
	PreservesGate602           bool
	Verdict                    string
}

type Analysis struct {
	Inherited           InheritedGate602
	S3Action            []S3ActionRow
	Invariants          []InvariantVsOrientationSensitive
	PhysicalLabels      PhysicalLabelAudit
	SignedDiscriminant  SignedDiscriminantAudit
	FourierCyclic       FourierCyclicOrientationAudit
	OrientationCoupling PMNSCKMOrientationCouplingAudit
	MinimalRemaining    MinimalRemainingSeal
	Firewalls           Firewalls
	Truth               string
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
	g602, err := generation2unsealedleptonwallpmnsrowbranchselectoraudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate602 predecessor: %w", err)
	}
	inherited := inherit(g602)
	s3 := auditS3Action()
	inv := auditInvariants()
	physical := auditPhysicalLabels()
	disc := auditSignedDiscriminant()
	cyclic := auditFourierCyclicOrientation()
	coupling := auditPMNSCKMOrientationCoupling()
	minimal := defineMinimalRemainingSeal()
	firewalls := auditFirewalls()
	truth := "Gate 603 identifies the sixfold sigma degeneracy left by Gate 602 as invisible to B_flav because the balance uses an unsigned electron-wall distance.  For the balance itself, sigma is Fourier-coordinate redundancy; to select a full charged-lepton cyclic/order orientation one would need an additional signed Vandermonde/discriminant orientation seal, which is not native."
	return Analysis{Inherited: inherited, S3Action: s3, Invariants: inv, PhysicalLabels: physical, SignedDiscriminant: disc, FourierCyclic: cyclic, OrientationCoupling: coupling, MinimalRemaining: minimal, Firewalls: firewalls, Truth: truth}, nil
}

func inherit(a generation2unsealedleptonwallpmnsrowbranchselectoraudit.Analysis) InheritedGate602 {
	return InheritedGate602{SelectsElectronRow: a.SelectorVerdict.SelectsElectronRow, SelectsP3: a.SelectorVerdict.SelectsThirdNeutrinoProjector, SelectsPositiveJ: a.SelectorVerdict.SelectsPositiveCKMSign, SelectsFullSigma: a.SelectorVerdict.SelectsFullChargedLeptonSigma, MinimalClassSize: a.ObservedRank.MinimalClassSize, BestResidual: a.Gap.BestAbsResidual, NextDistinctResidual: a.Gap.NextDistinctAbsResidual, Verdict: StatusGate602Inherited}
}

func auditS3Action() []S3ActionRow {
	orders := [][]string{{"e", "mu", "tau"}, {"e", "tau", "mu"}, {"mu", "e", "tau"}, {"mu", "tau", "e"}, {"tau", "e", "mu"}, {"tau", "mu", "e"}}
	out := make([]S3ActionRow, 0, len(orders))
	for _, order := range orders {
		delta, r := fourierPhase(order)
		q := (1 + r*r) / 3
		idx := indexOf(order, "e")
		_, epsDeg := nearestZeroWall(delta, idx)
		epsRad := epsDeg * math.Pi / 180
		kappa := 1 - 8*math.Pi*epsRad
		b := kappa - sin2Theta13/4 + jCKM
		vx := rootVandermonde(order)
		sgn := sign(vx)
		out = append(out, S3ActionRow{Sigma: strings.Join(order, ","), Order: append([]string(nil), order...), DeltaDeg: delta, R: r, Q: q, ElectronWallEpsilonDeg: epsDeg, ElectronWallEpsilonRad: epsRad, CanonicalAscending: strings.Join(order, ",") == "e,mu,tau", VandermondeX: vx, SignVandermondeX: sgn, OrientationParity: parityLabel(order), BFlavInvariantValue: b, Verdict: StatusS3ActionAudited})
	}
	return out
}

func auditInvariants() []InvariantVsOrientationSensitive {
	return []InvariantVsOrientationSensitive{
		{"Koide Q", true, false, false, "symmetric/projective scalar of the root vector; unchanged by relabeling", StatusS3ActionAudited},
		{"Fourier amplitude R", true, false, false, "magnitude in the democratic complement; invariant under S3 up to Fourier-coordinate rotation/reflection", StatusS3ActionAudited},
		{"Fourier phase delta", false, true, false, "coordinate angle in a chosen cyclic ordering; rotates or reflects under permutations", StatusFourierCyclicOrientationAudited},
		{"electron-wall distance epsilon_e", true, false, false, "once the physical electron wall alpha=e is selected, the unsigned distance is the same across sigma presentations", StatusSigmaGaugeRedundancy},
		{"signed Vandermonde sign(V_x)", false, true, false, "orientation of the ordered root list; changes under odd permutations", StatusDiscriminantOrientationSealRequired},
		{"B_flav at alpha=e,i=3,+J", true, false, false, "depends on the unsigned wall distance, reactor overlap, and CKM sign, not on sigma cyclic orientation", StatusBFlavDoesNotSeeCyclicSigma},
	}
}

func auditPhysicalLabels() PhysicalLabelAudit {
	return PhysicalLabelAudit{ObservedOrdering: "e < mu < tau", YukawaOrdering: fmt.Sprintf("%.15g < %.15g < %.15g", yElectron, yMuon, yTau), NativeOrdering: false, EnvironmentalOrdering: true, Verdict: StatusPhysicalLabelsAudited}
}

func auditSignedDiscriminant() SignedDiscriminantAudit {
	return SignedDiscriminantAudit{DeltaFormula: "Delta_e=prod_{i<j}(lambda_i-lambda_j)^2", VandermondeFormula: "V_e=prod_{i<j}(lambda_j-lambda_i)", RootVandermondeFormula: "V_x=prod_{i<j}(x_j-x_i)", TraceRingSuppliesDeltaOnly: true, SignedVRequiresOrdering: true, SignDistinguishesOrientation: true, NativeSignedVTheorem: false, MinimalSeal: "ChargedLeptonDiscriminantOrientationSeal = choice of sign(V_e) or sign(V_x), equivalent to cyclic/order orientation of (e,mu,tau)", Verdict: strings.Join([]string{StatusSignedDiscriminantAudited, StatusNoNativeSignedDiscriminantTheorem}, ";")}
}

func auditFourierCyclicOrientation() FourierCyclicOrientationAudit {
	return FourierCyclicOrientationAudit{RequiresCyclicConvention: true, CanonicalCycle: "(e,mu,tau)", ReversedCycle: "(e,tau,mu)", BFlavUsesUnsignedWallDistance: true, BFlavDependsOnCyclicOrientation: false, Verdict: strings.Join([]string{StatusFourierCyclicOrientationAudited, StatusBFlavDoesNotSeeCyclicSigma}, ";")}
}

func auditPMNSCKMOrientationCoupling() PMNSCKMOrientationCouplingAudit {
	jpmns := jPMNS()
	return PMNSCKMOrientationCouplingAudit{JCKMSign: +1, JPMNSSign: sign(jpmns), CanonicalVxSign: sign(rootVandermonde([]string{"e", "mu", "tau"})), Candidate1: "sgn(V_x)*sgn(J_CKM)", Candidate1Value: sign(rootVandermonde([]string{"e", "mu", "tau"})) * 1, Candidate2: "sgn(V_x)*sgn(J_CKM)*sgn(J_PMNS)", Candidate2Value: sign(rootVandermonde([]string{"e", "mu", "tau"})) * 1 * sign(jpmns), TypedASHAOperatorPresent: false, Verdict: strings.Join([]string{StatusPMNSCKMOrientationCouplingAudited, StatusNoNativeSignedDiscriminantTheorem}, ";")}
}

func defineMinimalRemainingSeal() MinimalRemainingSeal {
	return MinimalRemainingSeal{SigmaGaugeForBFlav: true, PhysicalFullOrderingRequiresSeal: true, SealName: "ChargedLeptonDiscriminantOrientationSeal", SealData: []string{"choice of signed Vandermonde sign(V_e) or sign(V_x)", "equivalent cyclic order of (e,mu,tau)", "compatible Fourier chamber orientation"}, Statement: "For B_flav, sigma is a Fourier-coordinate redundancy after alpha=e is selected.  If full cyclic/order orientation is considered physical, an extra signed-discriminant seal is required.", Verdict: strings.Join([]string{StatusSigmaGaugeRedundancy, StatusDiscriminantOrientationSealRequired, StatusNoNativeSigmaSelectionTheorem}, ";")}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, false, true, true, true, true, StatusGate603Boundary}
}

func fourierPhase(order []string) (deltaDeg float64, r float64) {
	xs := make([]float64, 3)
	for i, label := range order {
		xs[i] = math.Sqrt(yukawa(label))
	}
	a := (xs[0] + xs[1] + xs[2]) / 3
	omega := 2 * math.Pi / 3
	var c complex128
	for j := 0; j < 3; j++ {
		v := xs[j]/a - 1
		c += complex(v, 0) * cmplx.Exp(complex(0, -omega*float64(j)))
	}
	delta := math.Atan2(imag(c), real(c)) * 180 / math.Pi
	return normalizeDeg(delta), cmplx.Abs(c) * 2 / (3 * math.Sqrt2)
}

func nearestZeroWall(deltaDeg float64, componentIndex int) (float64, float64) {
	walls := []float64{normalizeDeg(135 - 120*float64(componentIndex)), normalizeDeg(225 - 120*float64(componentIndex))}
	bestWall, best := walls[0], circularDistanceDeg(deltaDeg, walls[0])
	for _, w := range walls[1:] {
		if d := circularDistanceDeg(deltaDeg, w); d < best {
			best = d
			bestWall = w
		}
	}
	return bestWall, best
}

func rootVandermonde(order []string) float64 {
	xs := make([]float64, 3)
	for i, label := range order {
		xs[i] = math.Sqrt(yukawa(label))
	}
	return (xs[1] - xs[0]) * (xs[2] - xs[0]) * (xs[2] - xs[1])
}

func parityLabel(order []string) string {
	canonical := map[string]int{"e": 0, "mu": 1, "tau": 2}
	inv := 0
	for i := 0; i < len(order); i++ {
		for j := i + 1; j < len(order); j++ {
			if canonical[order[i]] > canonical[order[j]] {
				inv++
			}
		}
	}
	if inv%2 == 0 {
		return "even"
	}
	return "odd"
}

func jPMNS() float64 {
	s12, c12 := math.Sqrt(sin2Theta12), math.Sqrt(1-sin2Theta12)
	s23, c23 := math.Sqrt(sin2Theta23), math.Sqrt(1-sin2Theta23)
	s13, c13 := math.Sqrt(sin2Theta13), math.Sqrt(1-sin2Theta13)
	return c12 * c23 * c13 * c13 * s12 * s23 * s13 * math.Sin(deltaCPDeg*math.Pi/180)
}

func sign(x float64) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
func indexOf(order []string, target string) int {
	for i, s := range order {
		if s == target {
			return i
		}
	}
	return -1
}
func yukawa(label string) float64 {
	switch label {
	case "e":
		return yElectron
	case "mu":
		return yMuon
	case "tau":
		return yTau
	default:
		panic("unknown label")
	}
}
func normalizeDeg(x float64) float64 {
	y := math.Mod(x, 360)
	if y < 0 {
		y += 360
	}
	return y
}
func circularDistanceDeg(a, b float64) float64 {
	d := math.Abs(normalizeDeg(a) - normalizeDeg(b))
	if d > 180 {
		d = 360 - d
	}
	return d
}

func Statuses() []string {
	return []string{StatusGate602Inherited, StatusSigmaDegeneracySourceIdentified, StatusS3ActionAudited, StatusPhysicalLabelsAudited, StatusSignedDiscriminantAudited, StatusFourierCyclicOrientationAudited, StatusPMNSCKMOrientationCouplingAudited, StatusSigmaGaugeRedundancy, StatusDiscriminantOrientationSealRequired, StatusBFlavDoesNotSeeCyclicSigma, StatusNoNativeSignedDiscriminantTheorem, StatusNoNativeMassOrderingTheorem, StatusNoNativeSigmaSelectionTheorem, StatusNoNativeBFlavZero, StatusGate352Preserved, StatusGate596Preserved, StatusGate600Preserved, StatusGate602Preserved, StatusNoKoideDerivation, StatusNoMassDerivation, StatusNoPMNSCKMDerivation, StatusNoBFlavNativePromotion, StatusNoNewCarrierSelector, StatusGate603Boundary}
}
