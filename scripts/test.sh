#!/bin/bash

error=0

go test -v --race --cover ./apps/analytics/... || error=1
go test -v --race --cover ./apps/monitor/... || error=1

if [ $error -eq 0 ]; then
  printf "All tests passed\n"
else
  printf "Some tests failed\n"
  printf "exit with status %s\n" $error
fi

exit $error