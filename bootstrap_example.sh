docker build -t gloss:latest .
docker stack deploy -c test_stack.yml gloss_example
docker service logs --follow gloss_example_web