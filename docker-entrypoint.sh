#!/bin/bash
ulimit -n 8192

exec "$@"