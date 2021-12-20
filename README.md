
#### Recomendation for local development most comfortable usage:
    make local // run all containers
    make run // it's easier way to attach debugger or rebuild/rerun project

#### 🙌👨‍💻🚀 Docker-compose files:
    docker-compose.local.yml - run postgresql, redis, aws, prometheus, grafana containrs
    docker-compose.dev.yml - run docker development environment
    docker-compose.delve.yml run development environment with delve debug

### Docker development usage:
    make docker

### Local development usage:
    make local
    make run
