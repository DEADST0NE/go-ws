#!/bin/bash

ulimit -n 20000
ulimit -i 127150
ulimit -u 20000

# Настройки TCP для оптимизации производительности
sysctl -w net.ipv4.tcp_tw_reuse=1
sysctl -w net.ipv4.tcp_fin_timeout=30
sysctl -w net.ipv4.tcp_keepalive_time=1800
sysctl -w net.ipv4.tcp_keepalive_probes=5
sysctl -w net.ipv4.tcp_keepalive_intvl=30
sysctl -w net.core.netdev_max_backlog=5000 # Убрано, так как ключ не распознается
sysctl -w net.ipv4.ip_local_port_range='1024 65535'
sysctl -w net.ipv4.tcp_rmem='4096 87380 6291456'
sysctl -w net.ipv4.tcp_wmem='4096 65536 6291456'
sysctl -w net.core.rmem_max=65536 # Убрано, так как ключ не распознается
sysctl -w net.core.wmem_max=65536 # Убрано, так как ключ не распознается

sysctl -w fs.nr_open=1000000
sysctl -w fs.file-max=1000000

net.ipv4.tcp_fin_timeout = 5
net.ipv4.tcp_syncookies = 0
sysctl -w net.ipv4.tcp_syn_retries = 1
sysctl -w net.ipv4.tcp_synack_retries = 1
sysctl -w net.core.somaxconn=65535
sysctl -w net.ipv4.tcp_max_syn_backlog=32768
sysctl -w kernel.threads-max=3261780

# Запуск основной команды контейнера
exec "$@"