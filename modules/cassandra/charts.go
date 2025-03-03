// SPDX-License-Identifier: GPL-3.0-or-later

package cassandra

import "github.com/netdata/go.d.plugin/agent/module"

const (
	prioRequestsRate = module.Priority + iota

	prioLatency

	prioCacheHitRatio
	prioCacheHitRate
	prioCacheSize

	prioStorageLiveDiskSpaceUsed
	prioStorageTotalDiskSpaceUsed

	prioCompactionCompletedTasksRate
	prioCompactionPendingTasksCount
	prioCompactionBytesCompactedRate

	prioThreadPoolsActiveTasksCount
	prioThreadPoolsPendingTasksCount
	prioThreadPoolsBlockedTasksRate
	prioThreadPoolsCurrentlyBlockedTasksCount

	prioJVMGCCount
	prioJVMGCTime

	prioDroppedMessagesOneMinuteRate
	prioRequestsTimeoutsRate
	prioRequestsUnavailablesRate
	prioRequestsFailuresRate
	prioStorageExceptionsRate
)

var baseCharts = module.Charts{
	chartClientRequestsRate.Copy(),

	chartClientRequestsLatency.Copy(),

	chartCacheHitRatio.Copy(),
	chartCacheHitRate.Copy(),
	chartCacheSize.Copy(),

	chartStorageLiveDiskSpaceUsed.Copy(),
	chartStorageTotalDiskSpaceUsed.Copy(),

	chartCompactionCompletedTasksRate.Copy(),
	chartCompactionPendingTasksCount.Copy(),
	chartCompactionBytesCompactedRate.Copy(),

	threadPoolsActiveTasksCount.Copy(),
	threadPoolsPendingTasksCount.Copy(),
	threadPoolsBlockedTasksRate.Copy(),
	threadPoolsBlockedTasksCount.Copy(),

	chartJVMGCRate.Copy(),
	chartJVMGCTime.Copy(),

	chartDroppedMessagesOneMinuteRate.Copy(),
	chartClientRequestTimeoutsRate.Copy(),
	chartClientRequestUnavailablesRate.Copy(),
	chartClientRequestFailuresRate.Copy(),
	chartStorageExceptionsRate.Copy(),
}

var (
	chartClientRequestsRate = module.Chart{
		ID:       "client_requests_rate",
		Title:    "Client requests rate",
		Units:    "requests/s",
		Fam:      "throughput",
		Ctx:      "cassandra.client_requests_rate",
		Priority: prioRequestsRate,
		Dims: module.Dims{
			{ID: "client_request_latency_reads", Name: "read", Algo: module.Incremental},
			{ID: "client_request_latency_writes", Name: "write", Algo: module.Incremental, Mul: -1},
		},
	}
)

var (
	chartClientRequestsLatency = module.Chart{
		ID:       "client_requests_latency",
		Title:    "Client requests latency",
		Units:    "seconds",
		Fam:      "latency",
		Ctx:      "cassandra.client_requests_latency",
		Priority: prioLatency,
		Dims: module.Dims{
			{ID: "client_request_total_latency_reads", Name: "read", Algo: module.Incremental, Div: 1e6},
			{ID: "client_request_total_latency_writes", Name: "write", Algo: module.Incremental, Mul: -1, Div: 1e6},
		},
	}
)

var (
	chartCacheHitRatio = module.Chart{
		ID:       "cache_hit_ratio",
		Title:    "Cache hit ratio",
		Units:    "percentage",
		Fam:      "cache",
		Ctx:      "cassandra.cache_hit_ratio",
		Priority: prioCacheHitRatio,
		Dims: module.Dims{
			{ID: "cache_hit_ratio", Name: "hit_ratio", Div: 1000},
		},
	}
	chartCacheHitRate = module.Chart{
		ID:       "cache_hit_rate",
		Title:    "Cache hit rate",
		Units:    "events/s",
		Fam:      "cache",
		Ctx:      "cassandra.cache_hit_rate",
		Priority: prioCacheHitRate,
		Type:     module.Stacked,
		Dims: module.Dims{
			{ID: "cache_hits", Name: "hits", Algo: module.Incremental},
			{ID: "cache_misses", Name: "misses", Algo: module.Incremental},
		},
	}
	chartCacheSize = module.Chart{
		ID:       "cache_size",
		Title:    "Cache size",
		Units:    "bytes",
		Fam:      "cache",
		Ctx:      "cassandra.cache_size",
		Priority: prioCacheSize,
		Dims: module.Dims{
			{ID: "cache_size", Name: "size"},
		},
	}
)

var (
	chartStorageLiveDiskSpaceUsed = module.Chart{
		ID:       "storage_live_disk_space_used",
		Title:    "Disk space used by live data",
		Units:    "bytes",
		Fam:      "disk usage",
		Ctx:      "cassandra.storage_live_disk_space_used",
		Priority: prioStorageLiveDiskSpaceUsed,
		Dims: module.Dims{
			{ID: "storage_load", Name: "used"},
		},
	}
	chartStorageTotalDiskSpaceUsed = module.Chart{
		ID:       "storage_total_disk_space_used",
		Title:    "Total disk space used",
		Units:    "bytes",
		Fam:      "disk usage",
		Ctx:      "cassandra.storage_total_disk_space_used",
		Priority: prioStorageTotalDiskSpaceUsed,
		Dims: module.Dims{
			{ID: "tables_total_disk_space_used", Name: "used"},
		},
	}
)

var (
	chartCompactionCompletedTasksRate = module.Chart{
		ID:       "compaction_completed_tasks_rate",
		Title:    "Completed compactions rate",
		Units:    "tasks/s",
		Fam:      "compaction",
		Ctx:      "cassandra.compaction_completed_tasks_rate",
		Priority: prioCompactionCompletedTasksRate,
		Dims: module.Dims{
			{ID: "compaction_completed_tasks", Name: "completed", Algo: module.Incremental},
		},
	}
	chartCompactionPendingTasksCount = module.Chart{
		ID:       "compaction_pending_tasks_count",
		Title:    "Pending compactions",
		Units:    "tasks",
		Fam:      "compaction",
		Ctx:      "cassandra.compaction_pending_tasks_count",
		Priority: prioCompactionPendingTasksCount,
		Dims: module.Dims{
			{ID: "compaction_pending_tasks", Name: "pending"},
		},
	}
	chartCompactionBytesCompactedRate = module.Chart{
		ID:       "compaction_compacted_rate",
		Title:    "Compaction data rate",
		Units:    "bytes/s",
		Fam:      "compaction",
		Ctx:      "cassandra.compaction_compacted_rate",
		Priority: prioCompactionBytesCompactedRate,
		Dims: module.Dims{
			{ID: "compaction_bytes_compacted", Name: "compacted", Algo: module.Incremental},
		},
	}
)

var (
	threadPoolsActiveTasksCount = module.Chart{
		ID:       "thread_pools_active_tasks_count",
		Title:    "Active tasks",
		Units:    "tasks",
		Fam:      "thread pools",
		Ctx:      "cassandra.thread_pools_active_tasks_count",
		Priority: prioThreadPoolsActiveTasksCount,
		Dims: module.Dims{
			{ID: "thread_pools_active_tasks", Name: "active"},
		},
	}
	threadPoolsPendingTasksCount = module.Chart{
		ID:       "thread_pools_pending_tasks_count",
		Title:    "Pending tasks",
		Units:    "tasks",
		Fam:      "thread pools",
		Ctx:      "cassandra.thread_pools_pending_tasks_count",
		Priority: prioThreadPoolsPendingTasksCount,
		Dims: module.Dims{
			{ID: "thread_pools_pending_tasks", Name: "pending"},
		},
	}
	threadPoolsBlockedTasksRate = module.Chart{
		ID:       "thread_pools_blocked_tasks_rate",
		Title:    "Blocked tasks rate",
		Units:    "tasks/s",
		Fam:      "thread pools",
		Ctx:      "cassandra.thread_pools_blocked_tasks_rate",
		Priority: prioThreadPoolsBlockedTasksRate,
		Dims: module.Dims{
			{ID: "thread_pools_total_blocked_tasks", Name: "blocked", Algo: module.Incremental},
		},
	}
	threadPoolsBlockedTasksCount = module.Chart{
		ID:       "thread_pools_blocked_tasks_count",
		Title:    "Blocked tasks",
		Units:    "tasks",
		Fam:      "thread pools",
		Ctx:      "cassandra.thread_pools_blocked_tasks_count",
		Priority: prioThreadPoolsCurrentlyBlockedTasksCount,
		Dims: module.Dims{
			{ID: "thread_pools_currently_blocked_tasks", Name: "blocked"},
		},
	}
)

var (
	chartJVMGCRate = module.Chart{
		ID:       "jvm_gc_rate",
		Title:    "Garbage collections rate",
		Units:    "gc/s",
		Fam:      "garbage collection",
		Ctx:      "cassandra.jvm_gc_rate",
		Priority: prioJVMGCCount,
		Dims: module.Dims{
			{ID: "jvm_gc_parnew_count", Name: "parnew", Algo: module.Incremental},
			{ID: "jvm_gc_cms_count", Name: "cms", Algo: module.Incremental},
		},
	}
	chartJVMGCTime = module.Chart{
		ID:       "jvm_gc_time",
		Title:    "Garbage collection time",
		Units:    "seconds",
		Fam:      "garbage collection",
		Ctx:      "cassandra.jvm_gc_time",
		Priority: prioJVMGCTime,
		Dims: module.Dims{
			{ID: "jvm_gc_parnew_time", Name: "parnew", Algo: module.Incremental, Div: 1e9},
			{ID: "jvm_gc_cms_time", Name: "cms", Algo: module.Incremental, Div: 1e9},
		},
	}
)

var (
	chartDroppedMessagesOneMinuteRate = module.Chart{
		ID:       "dropped_messages_one_minute_rate",
		Title:    "Dropped messages one minute rate",
		Units:    "messages/s",
		Fam:      "errors",
		Ctx:      "cassandra.dropped_messages_one_minute_rate",
		Priority: prioDroppedMessagesOneMinuteRate,
		Dims: module.Dims{
			{ID: "dropped_messages_one_minute", Name: "dropped", Div: 1000},
		},
	}
	chartClientRequestTimeoutsRate = module.Chart{
		ID:       "client_requests_timeouts_rate",
		Title:    "Client requests timeouts rate",
		Units:    "timeouts/s",
		Fam:      "errors",
		Ctx:      "cassandra.client_requests_timeouts_rate",
		Priority: prioRequestsTimeoutsRate,
		Dims: module.Dims{
			{ID: "client_request_timeouts_reads", Name: "read", Algo: module.Incremental},
			{ID: "client_request_timeouts_writes", Name: "write", Algo: module.Incremental, Mul: -1},
		},
	}
	chartClientRequestUnavailablesRate = module.Chart{
		ID:       "client_requests_unavailables_rate",
		Title:    "Client requests unavailable exceptions rate",
		Units:    "exceptions/s",
		Fam:      "errors",
		Ctx:      "cassandra.client_requests_unavailables_rate",
		Priority: prioRequestsUnavailablesRate,
		Dims: module.Dims{
			{ID: "client_request_unavailables_reads", Name: "read", Algo: module.Incremental},
			{ID: "client_request_unavailables_writes", Name: "write", Algo: module.Incremental, Mul: -1},
		},
	}
	chartClientRequestFailuresRate = module.Chart{
		ID:       "client_requests_failures_rate",
		Title:    "Client requests failures rate",
		Units:    "failures/s",
		Fam:      "errors",
		Ctx:      "cassandra.client_requests_failures_rate",
		Priority: prioRequestsFailuresRate,
		Dims: module.Dims{
			{ID: "client_request_failures_reads", Name: "read", Algo: module.Incremental},
			{ID: "client_request_failures_writes", Name: "write", Algo: module.Incremental, Mul: -1},
		},
	}
	chartStorageExceptionsRate = module.Chart{
		ID:       "storage_exceptions_rate",
		Title:    "Storage exceptions rate",
		Units:    "exceptions/s",
		Fam:      "errors",
		Ctx:      "cassandra.storage_exceptions_rate",
		Priority: prioStorageExceptionsRate,
		Dims: module.Dims{
			{ID: "storage_exceptions", Name: "storage", Algo: module.Incremental},
		},
	}
)
