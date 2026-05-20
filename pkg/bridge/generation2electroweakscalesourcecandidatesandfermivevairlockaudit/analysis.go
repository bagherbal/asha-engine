// Package generation2electroweakscalesourcecandidatesandfermivevairlockaudit implements
// Gate 778: Electroweak Scale Source Candidates and Fermi-VEV Airlock Audit.
//
// Gate 777 separated the sealed tree Higgs tower into a dimensionless correction
// factor C_Higgs and a dimensionful VEV scale seal v. Gate 778 audits lawful
// source candidates for v, records the Fermi-VEV convention lane, and rejects
// shortcuts that would turn dimensionless scalar bridge data into a GeV-scale
// theorem. This is a scale-source and airlock audit only. It does not derive
// the VEV, Fermi constant, W mass, absolute gauge coupling, scalar runtime
// lambda, Higgs pole mass, Yukawa operators, CKM/PMNS, flavor hierarchy, or a
// native HistoryLoopUnit theorem.
package generation2electroweakscalesourcecandidatesandfermivevairlockaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE778-ELECTROWEAK-SCALE-SOURCE-CANDIDATES-FERMI-VEV-AIRLOCK-AUDIT"

	StatusGate777VEVScaleAirlockInherited          = "PASS_GATE777_VEV_SCALE_AIRLOCK_INHERITED"
	StatusFermiVEVConventionLaneAudited            = "PASS_FERMI_VEV_CONVENTION_LANE_AUDITED"
	StatusWMassGaugeCouplingLaneAudited            = "PASS_W_MASS_GAUGE_COUPLING_LANE_AUDITED"
	StatusPotentialStationarityLaneAudited         = "PASS_POTENTIAL_STATIONARITY_LANE_AUDITED"
	StatusSpectralActionScaleCandidateAudited      = "PASS_SPECTRAL_ACTION_SCALE_CANDIDATE_AUDITED"
	StatusBoundaryRGScaleLaneAudited               = "PASS_BOUNDARY_RG_SCALE_LANE_AUDITED"
	StatusSourceRankingRecorded                    = "PASS_SOURCE_RANKING_RECORDED"
	StatusPhysicalFirewallsEnforced                = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusVCanBeTypedByFermiScaleConventionSeal    = "CONDITIONAL_SUPPORT_V_CAN_BE_TYPED_BY_FERMI_SCALE_CONVENTION_SEAL"
	StatusDimensionfulScaleRequiresSeparateAirlock = "CONDITIONAL_SUPPORT_DIMENSIONFUL_HIGGS_SCALE_REQUIRES_SEPARATE_SCALE_AIRLOCK"
	StatusNoNativeFermiConstantTheorem             = "FAILED_ROUTE_NO_NATIVE_FERMI_CONSTANT_THEOREM"
	StatusNoNativeElectroweakScaleTheorem          = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SCALE_THEOREM"
	StatusNoNativeMuSquaredSourceTheorem           = "FAILED_ROUTE_NO_NATIVE_MU_SQUARED_SOURCE_THEOREM"
	StatusCHiggsDoesNotSetMassUnits                = "FAILED_ROUTE_C_HIGGS_DOES_NOT_SET_MASS_UNITS"
	StatusTreeProxyNotPoleMass                     = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusNoYukawaOperatorOrEigenvalueTheorem      = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate778ElectroweakScaleSourceBoundary    = "FIREWALL_PRESERVED_GATE778_ELECTROWEAK_SCALE_SOURCE_BOUNDARY"
)

const (
	// Gate 777 ledger snapshots.
	cHiggsMZ         = 1.0372205204048603
	dRadialMZ        = 1.0184402389953279
	vevConventionGeV = 246.2196508
)

type Gate777Inheritance struct {
	Inherited                     bool
	TreeTowerFormula              string
	CHiggs                        float64
	DilationFactor                float64
	VEVScaleGeV                   float64
	TreeMassGeV                   float64
	DimensionlessCorrectionExists bool
	DimensionfulScaleSeal         string
	NativeVEVTheorem              bool
	PoleMassTheorem               bool
	Verdict                       string
}

type FermiVEVConventionLane struct {
	Audited               bool
	SealName              string
	Formula               string
	Input                 string
	Output                string
	VEVGeV                float64
	EquivalentGFGeVMinus2 float64
	NativeFermiTheorem    bool
	NativeVEVTheorem      bool
	LawfulExternalAirlock bool
	Verdict               string
}

type WMassGaugeCouplingLane struct {
	Audited                 bool
	Formula                 string
	RequiresAbsoluteWeakG   bool
	RequiresWMass           bool
	GaugeRatiosOrganized    bool
	AbsoluteWeakScaleNative bool
	WMassNative             bool
	LaneSealed              bool
	Verdict                 string
}

type PotentialStationarityLane struct {
	Audited                       bool
	Formula                       string
	LambdaAirlockExists           bool
	MuSquaredIndependentlySourced bool
	CircularWithoutMuSource       bool
	NativeMuSquaredSource         bool
	DeterminesVEV                 bool
	Verdict                       string
}

type SpectralActionScaleCandidate struct {
	Audited                        bool
	Candidate                      string
	DimensionfulScaleCouldSetUnits bool
	CurrentBridgeDimensionlessOnly bool
	MapsSpectralScaleToVEV         bool
	LaneCandidateOnly              bool
	Verdict                        string
}

type BoundaryRGScaleLane struct {
	Audited                       bool
	BoundaryScaleSealExists       bool
	ScalarWallDataExists          bool
	DeterminesElectroweakVEV      bool
	BoundaryScaleEqualsVEVTheorem bool
	Verdict                       string
}

type SourceRanking struct {
	BestCurrentLawfulSource string
	BestFutureNativeTargets []string
	BlockedShortcuts        []string
	Recorded                bool
	Verdict                 string
}

type DerivedLedger struct {
	CHiggs         float64
	DilationFactor float64
	VEVGeV         float64
	EquivalentGF   float64
	VHalfGeV       float64
	TreeMassGeV    float64
	Finite         bool
}

type PhysicalFirewalls struct {
	Audited                            bool
	FermiScaleNativeTheorem            bool
	VEVDerivedFromCHiggs               bool
	VEVDerivedFromLambdaRuntimeOnly    bool
	MuSquaredBridgeNativeSource        bool
	WRelationNativeWithoutInputs       bool
	TreeProxyPoleMass                  bool
	DimensionlessTowerMassScaleTheorem bool
	YukawaOperatorOrEigenvalue         bool
	Verdict                            string
}

type Analysis struct {
	Gate777       Gate777Inheritance
	FermiLane     FermiVEVConventionLane
	WLane         WMassGaugeCouplingLane
	PotentialLane PotentialStationarityLane
	SpectralLane  SpectralActionScaleCandidate
	BoundaryLane  BoundaryRGScaleLane
	Ranking       SourceRanking
	Ledger        DerivedLedger
	Firewalls     PhysicalFirewalls
	Truth         string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		return cloneAnalysis(cache), nil
	}
	for name, value := range map[string]float64{
		"C_Higgs":  cHiggsMZ,
		"D_radial": dRadialMZ,
		"v":        vevConventionGeV,
	} {
		if math.IsNaN(value) || math.IsInf(value, 0) || value <= 0 {
			return nil, fmt.Errorf("invalid %s ledger: %.17g", name, value)
		}
	}
	gf := 1 / (math.Sqrt2 * vevConventionGeV * vevConventionGeV)
	vHalf := vevConventionGeV / 2
	mass := vHalf * dRadialMZ
	if !finite(gf, vHalf, mass) || gf <= 0 || mass <= 0 {
		return nil, fmt.Errorf("invalid derived scale ledger")
	}

	a := &Analysis{
		Gate777: Gate777Inheritance{
			Inherited:                     true,
			TreeTowerFormula:              "m_H_tree=(v/2)sqrt(C_Higgs)",
			CHiggs:                        cHiggsMZ,
			DilationFactor:                dRadialMZ,
			VEVScaleGeV:                   vevConventionGeV,
			TreeMassGeV:                   mass,
			DimensionlessCorrectionExists: true,
			DimensionfulScaleSeal:         "v",
			NativeVEVTheorem:              false,
			PoleMassTheorem:               false,
			Verdict:                       StatusGate777VEVScaleAirlockInherited,
		},
		FermiLane: FermiVEVConventionLane{
			Audited:               true,
			SealName:              "FermiVEVScaleSeal",
			Formula:               "v=(sqrt(2)G_F)^(-1/2)",
			Input:                 "G_F",
			Output:                "v",
			VEVGeV:                vevConventionGeV,
			EquivalentGFGeVMinus2: gf,
			NativeFermiTheorem:    false,
			NativeVEVTheorem:      false,
			LawfulExternalAirlock: true,
			Verdict:               StatusFermiVEVConventionLaneAudited,
		},
		WLane: WMassGaugeCouplingLane{
			Audited:                 true,
			Formula:                 "v=2m_W/g",
			RequiresAbsoluteWeakG:   true,
			RequiresWMass:           true,
			GaugeRatiosOrganized:    true,
			AbsoluteWeakScaleNative: false,
			WMassNative:             false,
			LaneSealed:              true,
			Verdict:                 StatusWMassGaugeCouplingLaneAudited,
		},
		PotentialLane: PotentialStationarityLane{
			Audited:                       true,
			Formula:                       "v^2=-mu^2/lambda_H",
			LambdaAirlockExists:           true,
			MuSquaredIndependentlySourced: false,
			CircularWithoutMuSource:       true,
			NativeMuSquaredSource:         false,
			DeterminesVEV:                 false,
			Verdict:                       StatusPotentialStationarityLaneAudited,
		},
		SpectralLane: SpectralActionScaleCandidate{
			Audited:                        true,
			Candidate:                      "dimensionful spectral-action scale or cutoff",
			DimensionfulScaleCouldSetUnits: true,
			CurrentBridgeDimensionlessOnly: true,
			MapsSpectralScaleToVEV:         false,
			LaneCandidateOnly:              true,
			Verdict:                        StatusSpectralActionScaleCandidateAudited,
		},
		BoundaryLane: BoundaryRGScaleLane{
			Audited:                       true,
			BoundaryScaleSealExists:       true,
			ScalarWallDataExists:          true,
			DeterminesElectroweakVEV:      false,
			BoundaryScaleEqualsVEVTheorem: false,
			Verdict:                       StatusBoundaryRGScaleLaneAudited,
		},
		Ranking: SourceRanking{
			BestCurrentLawfulSource: "FermiVEVScaleSeal: v=(sqrt(2)G_F)^(-1/2)",
			BestFutureNativeTargets: []string{"mu^2 source theorem", "absolute electroweak scale theorem"},
			BlockedShortcuts: []string{
				"C_Higgs does not determine v",
				"lambda_runtime_eff does not determine v without mu^2",
				"P_rad does not determine v",
				"HistoryLoopUnit does not determine v",
				"7/72 does not determine v",
				"1/(8pi) does not determine v",
			},
			Recorded: true,
			Verdict:  StatusSourceRankingRecorded,
		},
		Ledger: DerivedLedger{
			CHiggs:         cHiggsMZ,
			DilationFactor: dRadialMZ,
			VEVGeV:         vevConventionGeV,
			EquivalentGF:   gf,
			VHalfGeV:       vHalf,
			TreeMassGeV:    mass,
			Finite:         finite(cHiggsMZ, dRadialMZ, vevConventionGeV, gf, vHalf, mass),
		},
		Firewalls: PhysicalFirewalls{
			Audited:                            true,
			FermiScaleNativeTheorem:            false,
			VEVDerivedFromCHiggs:               false,
			VEVDerivedFromLambdaRuntimeOnly:    false,
			MuSquaredBridgeNativeSource:        false,
			WRelationNativeWithoutInputs:       false,
			TreeProxyPoleMass:                  false,
			DimensionlessTowerMassScaleTheorem: false,
			YukawaOperatorOrEigenvalue:         false,
			Verdict:                            StatusGate778ElectroweakScaleSourceBoundary,
		},
		Truth: fmt.Sprintf("Gate778 keeps the Higgs tree tower split: dimensionless C_Higgs=%.16f and dimensionful VEV seal v=%.7f GeV. The lawful current scale airlock is FermiVEVScaleSeal with G_F=%.14e GeV^-2; no native Fermi, VEV, electroweak-scale, W/g, mu^2, pole-mass, or Yukawa theorem is derived.", cHiggsMZ, vevConventionGeV, gf),
	}
	cache = a
	return cloneAnalysis(a), nil
}

func cloneAnalysis(a *Analysis) *Analysis {
	clone := *a
	clone.Ranking.BestFutureNativeTargets = append([]string(nil), a.Ranking.BestFutureNativeTargets...)
	clone.Ranking.BlockedShortcuts = append([]string(nil), a.Ranking.BlockedShortcuts...)
	return &clone
}

func finite(values ...float64) bool {
	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate777VEVScaleAirlockInherited,
		StatusFermiVEVConventionLaneAudited,
		StatusWMassGaugeCouplingLaneAudited,
		StatusPotentialStationarityLaneAudited,
		StatusSpectralActionScaleCandidateAudited,
		StatusBoundaryRGScaleLaneAudited,
		StatusSourceRankingRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusVCanBeTypedByFermiScaleConventionSeal,
		StatusDimensionfulScaleRequiresSeparateAirlock,
		StatusNoNativeFermiConstantTheorem,
		StatusNoNativeElectroweakScaleTheorem,
		StatusNoNativeMuSquaredSourceTheorem,
		StatusCHiggsDoesNotSetMassUnits,
		StatusTreeProxyNotPoleMass,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate778ElectroweakScaleSourceBoundary,
	}
}

func FormatGate777(g Gate777Inheritance) string {
	return fmt.Sprintf("inherited=%t formula=%s C=%.17g D=%.17g v=%.10f mass=%.12f dimless=%t scale=%s nativeVEV=%t pole=%t verdict=%s", g.Inherited, g.TreeTowerFormula, g.CHiggs, g.DilationFactor, g.VEVScaleGeV, g.TreeMassGeV, g.DimensionlessCorrectionExists, g.DimensionfulScaleSeal, g.NativeVEVTheorem, g.PoleMassTheorem, g.Verdict)
}

func FormatFermiLane(f FermiVEVConventionLane) string {
	return fmt.Sprintf("audited=%t seal=%s formula=%s input=%s output=%s v=%.10f GF=%.14e nativeFermi=%t nativeVEV=%t airlock=%t verdict=%s", f.Audited, f.SealName, f.Formula, f.Input, f.Output, f.VEVGeV, f.EquivalentGFGeVMinus2, f.NativeFermiTheorem, f.NativeVEVTheorem, f.LawfulExternalAirlock, f.Verdict)
}

func FormatWLane(w WMassGaugeCouplingLane) string {
	return fmt.Sprintf("audited=%t formula=%s needsG=%t needsW=%t ratios=%t absGNative=%t WNative=%t sealed=%t verdict=%s", w.Audited, w.Formula, w.RequiresAbsoluteWeakG, w.RequiresWMass, w.GaugeRatiosOrganized, w.AbsoluteWeakScaleNative, w.WMassNative, w.LaneSealed, w.Verdict)
}

func FormatPotentialLane(p PotentialStationarityLane) string {
	return fmt.Sprintf("audited=%t formula=%s lambdaAirlock=%t muSourced=%t circular=%t nativeMu=%t determinesVEV=%t verdict=%s", p.Audited, p.Formula, p.LambdaAirlockExists, p.MuSquaredIndependentlySourced, p.CircularWithoutMuSource, p.NativeMuSquaredSource, p.DeterminesVEV, p.Verdict)
}

func FormatSpectralLane(s SpectralActionScaleCandidate) string {
	return fmt.Sprintf("audited=%t candidate=%s couldSetUnits=%t dimensionlessOnly=%t mapsToVEV=%t candidateOnly=%t verdict=%s", s.Audited, s.Candidate, s.DimensionfulScaleCouldSetUnits, s.CurrentBridgeDimensionlessOnly, s.MapsSpectralScaleToVEV, s.LaneCandidateOnly, s.Verdict)
}

func FormatBoundaryLane(b BoundaryRGScaleLane) string {
	return fmt.Sprintf("audited=%t boundarySeal=%t wallData=%t determinesVEV=%t boundaryEqualsVEV=%t verdict=%s", b.Audited, b.BoundaryScaleSealExists, b.ScalarWallDataExists, b.DeterminesElectroweakVEV, b.BoundaryScaleEqualsVEVTheorem, b.Verdict)
}

func FormatRanking(r SourceRanking) string {
	return fmt.Sprintf("bestCurrent=%s future=%s blocked=%s recorded=%t verdict=%s", r.BestCurrentLawfulSource, strings.Join(r.BestFutureNativeTargets, ","), strings.Join(r.BlockedShortcuts, ","), r.Recorded, r.Verdict)
}

func FormatLedger(l DerivedLedger) string {
	return fmt.Sprintf("C=%.17g D=%.17g v=%.10f GF=%.14e vhalf=%.10f mass=%.12f finite=%t", l.CHiggs, l.DilationFactor, l.VEVGeV, l.EquivalentGF, l.VHalfGeV, l.TreeMassGeV, l.Finite)
}

func FormatFirewalls(f PhysicalFirewalls) string {
	return fmt.Sprintf("audited=%t fermiNative=%t vFromC=%t vFromLambda=%t muNative=%t WWithoutInputs=%t pole=%t dimlessMass=%t yukawaOp=%t verdict=%s", f.Audited, f.FermiScaleNativeTheorem, f.VEVDerivedFromCHiggs, f.VEVDerivedFromLambdaRuntimeOnly, f.MuSquaredBridgeNativeSource, f.WRelationNativeWithoutInputs, f.TreeProxyPoleMass, f.DimensionlessTowerMassScaleTheorem, f.YukawaOperatorOrEigenvalue, f.Verdict)
}
