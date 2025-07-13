package cache

const (
	// Maximum number of query length samples to keep in memory
	MaxQueryLengthSamples = 1000

	// Number of samples to remove when limit is exceeded
	QueryLengthCleanupThreshold = 500

	// Maximum number of response size samples to keep in memory
	MaxResponseSizeSamples = 1000

	// Number of response size samples to remove when limit is exceeded
	ResponseSizeCleanupThreshold = 500

	// Maximum number of strategy history entries to keep
	MaxStrategyHistoryEntries = 100

	// Maximum number of user patterns to keep per tenant
	MaxUserPatternsPerTenant = 1000

	// Maximum number of model patterns to keep
	MaxModelPatterns = 100

	// Maximum number of time pattern entries to keep (24 hours)
	MaxTimePatternEntries = 24

	// Cleanup interval for adaptive learning data (in hours)
	AdaptiveCleanupIntervalHours = 24

	// Maximum age of pattern data before cleanup (in hours)
	MaxPatternDataAgeHours = 168 // 1 week
)
