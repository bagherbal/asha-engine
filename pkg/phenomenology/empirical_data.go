// Package phenomenology contains the empirical-seal layer for ASHA.
//
// This package deliberately lives outside pkg/bridge.  The bridge packages track
// ASHA-native geometric derivations.  The phenomenology package imports the
// sealed ASHA law-space and injects measured environmental data in order to
// compute conditional observables for our universe.
package phenomenology

const (
	// 2024 PDG-style empirical quarantine values supplied to this package.
	TopMassGeV     = 172.69
	HiggsMassGeV   = 125.25
	StrongCoupling = 0.1179 // alpha_s(m_Z)
	ZBosonMassGeV  = 91.1876

	// Cosmological target observables.
	TargetRelicDensity = 0.120    // Omega_c h^2
	TargetDarkEnergy   = 1.0e-120 // rho_Lambda / M_Pl^4
)

const (
	// ASHA-derived/sealed inputs inherited from the final bridge ledger.
	AshaHiggsTreeProxyGeV    = 124.925370
	AshaLambdaAtVEV          = 0.1277456365
	AshaBGapMajoranaMassGeV  = 1.46774973718e6
	AshaThresholdDeltaLambda = -0.097846792207
	AshaPfaffianVEVGeV       = 247.1513
	StandardVEVGeV           = 246.22

	// Continuum numerical constants.
	PlanckMassGeV              = 1.2209e19 // unreduced Planck mass, used as upper RG scale here
	ReducedPlanckMassGeV       = 2.435e18
	AgeOfUniverseYears         = 13.8e9
	SecondsPerYear             = 31557600.0
	HbarSecondsPerGeV          = 6.582119569e-25
	EntropyDensityTodayPerCM3  = 2891.2
	CriticalDensityH2GeVPerCM3 = 1.05375e-5
	CMBTemperatureK            = 2.7255

	// Approximate electroweak gauge couplings at m_Z used by the one-loop audit.
	// g1 is GUT-normalized: g1 = sqrt(5/3) g_Y.
	G1AtMZ = 0.4614
	G2AtMZ = 0.6517
)

const (
	StatusEmpiricalSealLoaded             = "CONDITIONAL_SUPPORT_EMPIRICAL_QUARANTINE_STATE_LOADED"
	StatusVacuumRGExecuted                = "CONDITIONAL_SUPPORT_ONE_LOOP_VACUUM_RG_EXECUTED"
	StatusAshaThresholdApplied            = "CONDITIONAL_SUPPORT_ASHA_THRESHOLD_JUMP_APPLIED"
	StatusVacuumMetastableConditional     = "CONDITIONAL_SUPPORT_CONDITIONAL_VACUUM_METASTABILITY_FOUND"
	StatusDarkYieldConstraintComputed     = "CONDITIONAL_SUPPORT_DARK_MATTER_REQUIRED_YIELD_COMPUTED"
	StatusThermalRelicOverclosureComputed = "CONDITIONAL_SUPPORT_THERMAL_STABLE_MAJORANA_OVERCLOSURE_COMPUTED"
	StatusCosmologicalFineTuningComputed  = "CONDITIONAL_SUPPORT_COSMOLOGICAL_CONSTANT_FINE_TUNING_COMPUTED"
	StatusPhenomenologyLayerComplete      = "CONDITIONAL_SUPPORT_PHENOMENOLOGY_LAYER_EXECUTED"
	StatusNotNativePrediction             = "CONDITIONAL_TENSION_OUTPUTS_DEPEND_ON_EMPIRICAL_SEALS"
	StatusDarkMatterNotPredicted          = "FAILED_ROUTE_DARK_MATTER_ABUNDANCE_NOT_NATIVELY_PREDICTED"
	StatusUniverseLifetimeNotNative       = "FAILED_ROUTE_UNIVERSE_LIFETIME_NOT_NATIVELY_PREDICTED"
	StatusCosmologicalConstantNotSolved   = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_ORGANICALLY_SOLVED"
)
