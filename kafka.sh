# quick access to run kafka cli commands in docker container


# Usage:
# ./kafka.sh <command> <args>
# e.g.
# ./kafka.sh topic list
# ./kafka.sh topic describe --topic my-topic
# ./kafka.sh console-consumer --topic my-topic --from-beginning
# ./kafka.sh console-producer --topic my-topic

# Note: if you want to pass arguments to the command, you need to use the -- separator
# e.g.
# ./kafka.sh console-producer --topic my-topic -- --property parse.key=true --property key.separator=,

# get the command to run
COMMAND=$1

# get the command arguments
ARGS=${@:2}

# get the docker-compose file path
DOCKER_COMPOSE_FILE_PATH=$(dirname "$0")/docker-compose.yml

# run the command in the docker container
docker exec redpanda rpk $COMMAND $ARGS

