# microsoft-defender-ATP-exporter
Prometheus exporter for various Microsoft Defender ATP metrics, taken from REST API

## Usage

1. Clone this repository
2. `go build ./cmd/defender-atp-exporter`
3. Set required env var
```
export AZURE_TENANT_ID="..."
export AZURE_CLIENT_ID="..."
export AZURE_CLIENT_SECRET="..."
```
Check documentation to create the application : https://docs.microsoft.com/fr-fr/windows/security/threat-protection/microsoft-defender-atp/exposed-apis-create-app-webapp
4. `./defender-atp-exporter`
5. `curl localhost:8080/metrics`


## Metrics outputs

```
# HELP defender_atp_exposure Exposure of machines
# TYPE defender_atp_exposure gauge
defender_atp_exposure{product_name="adblock_for_mac",severity="Low"} 1
defender_atp_exposure{product_name="chrome_for_mac",severity="Low"} 317
defender_atp_exposure{product_name="chrome_for_mac",severity="Medium"} 273
defender_atp_exposure{product_name="edge_chromium-based_for_mac",severity="Low"} 6
defender_atp_exposure{product_name="edge_chromium-based_for_mac",severity="Medium"} 12
defender_atp_exposure{product_name="firefox_for_mac",severity="Low"} 126
defender_atp_exposure{product_name="firefox_for_mac",severity="Medium"} 91
defender_atp_exposure{product_name="horizon_for_mac",severity="Low"} 2
defender_atp_exposure{product_name="horizon_for_mac",severity="Medium"} 1
defender_atp_exposure{product_name="intellij_idea_for_mac",severity="Low"} 1
defender_atp_exposure{product_name="safari_for_mac",severity="Low"} 3
# HELP defender_atp_vulnerabilities Number of vulnerability found on machines
# TYPE defender_atp_vulnerabilities gauge
defender_atp_vulnerabilities{machineId="098dab67c4b154f7fa67da7e05aadc574381a612",severity="Low"} 15
defender_atp_vulnerabilities{machineId="098dab67c4b154f7fa67da7e05aadc574381a612",severity="Medium"} 7
defender_atp_vulnerabilities{machineId="26b03f48d5f3f0b0e18fae49ffeba80cadf44b82",severity="Low"} 45
defender_atp_vulnerabilities{machineId="26b03f48d5f3f0b0e18fae49ffeba80cadf44b82",severity="Medium"} 32
defender_atp_vulnerabilities{machineId="84b56163b15fe90d73ba609f8d7af97f4a127de9",severity="Low"} 101
defender_atp_vulnerabilities{machineId="84b56163b15fe90d73ba609f8d7af97f4a127de9",severity="Medium"} 89
defender_atp_vulnerabilities{machineId="9e100538342763fa032f2c7b019ea21f5945e2c0",severity="Low"} 4
defender_atp_vulnerabilities{machineId="9e100538342763fa032f2c7b019ea21f5945e2c0",severity="Medium"} 8
defender_atp_vulnerabilities{machineId="af1db96242c897aeb8fb67372f997e9d6ef917a9",severity="Low"} 1
defender_atp_vulnerabilities{machineId="b279f02fa9e7d2400daa7b36299478f1ddc51c15",severity="Low"} 53
defender_atp_vulnerabilities{machineId="b279f02fa9e7d2400daa7b36299478f1ddc51c15",severity="Medium"} 46
defender_atp_vulnerabilities{machineId="d8c7d991d5c68428b3397d9b09d6574257044395",severity="Low"} 38
defender_atp_vulnerabilities{machineId="d8c7d991d5c68428b3397d9b09d6574257044395",severity="Medium"} 25
defender_atp_vulnerabilities{machineId="eb441179015d4cb63e09afa0be78280354d5915c",severity="Low"} 99
defender_atp_vulnerabilities{machineId="eb441179015d4cb63e09afa0be78280354d5915c",severity="Medium"} 85
defender_atp_vulnerabilities{machineId="ec60c92fff7f690e107261fe67a07c0ec2af7c37",severity="Low"} 1
defender_atp_vulnerabilities{machineId="ed8e575c21bd4751e1cbc75980a6f78086ef2f33",severity="Low"} 99
defender_atp_vulnerabilities{machineId="ed8e575c21bd4751e1cbc75980a6f78086ef2f33",severity="Medium"} 85
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 4.1343e-05
go_gc_duration_seconds{quantile="0.25"} 4.1343e-05
go_gc_duration_seconds{quantile="0.5"} 4.1343e-05
go_gc_duration_seconds{quantile="0.75"} 4.1343e-05
go_gc_duration_seconds{quantile="1"} 4.1343e-05
go_gc_duration_seconds_sum 4.1343e-05
go_gc_duration_seconds_count 1
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 12
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.15.5"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 3.512432e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 5.178152e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.44525e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 57917
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 6.805351150621018e-05
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 4.539616e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 3.512432e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.0882944e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 5.603328e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 20822
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 6.0850176e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6486272e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.6075982061772192e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 78739
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 27776
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 32768
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 115600
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 131072
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.470878e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 622592
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 622592
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.4728448e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 13
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```