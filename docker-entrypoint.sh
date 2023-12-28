#!/bin/bash
ulimit -n 20000
ulimit -i 127150

exec "$@"