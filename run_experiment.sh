#!/bin/bash
# run_experiment.sh <nodeID> [hz] [timeout]
# Example: ./run_experiment.sh node1 0.5 100

NODE=$1
HZ=${2:-0.5}
TIME=${3:-100}

python3 -u gentx.py $HZ | timeout $TIME ./mp1_node $NODE config.txt

echo "Done. Latency saved to latency_$NODE.txt"