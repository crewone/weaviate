function wait_weaviate() {
  port_number=$1
  echo "Wait for Weaviate to be ready: $port_number"
  for _ in {1..120}; do
    if curl -sf -o /dev/null localhost:$port_number/v1/.well-known/ready; then
      echo "Weaviate is ready"
      return 0
    fi

    echo "Weaviate is not ready on port $port_number, trying again in 1s"
    sleep 1
  done
  echo "ERROR: Weaviate is not ready after 120s on port $port_number"
  exit 1
}

# Create logs directory if it doesn't exist (mv it if it does)
mv logs logs.bak
mkdir -p logs

# Start local-development node and redirect output to log file
./tools/dev/restart_dev_environment.sh
./tools/dev/run_dev_server.sh local-development > logs/local-development.log 2>&1 &
sleep 3
./tools/dev/run_dev_server.sh second-node > logs/second-node.log 2>&1 &
# ./tools/dev/run_dev_server.sh third-node > logs/third-node.log 2>&1 &

wait_weaviate 8080
wait_weaviate 8081

echo "All nodes are running. You can tail the logs with:"
echo "tail -f logs/local-development.log"
echo "tail -f logs/second-node.log"
# echo "tail -f logs/third-node.log"

# du -hd 1 data-weaviate-*/*vector/