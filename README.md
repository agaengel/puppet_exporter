# Puppet Exporter for Prometheus

This exporter aims to provide metrics for prometheus based on the last_run_summary.yaml report written by the puppet agent.

## build from source
### compile binarys
make
#### run test
make test

the default port for the webserver is 9199
this can be set via -listenAddress

start with file form unitetest:
./bin/github.com/agaengel/puppet_exporter -puppetLastReportPath /Users/agaengel/go/src/github.com/agaengel/puppet_exporter/test/last_run_summary.yaml
