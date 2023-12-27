#!/bin/sh
k6 run -e ENV="$ENV" $SCRIPT --tag test_run_id=$SCRIPT
