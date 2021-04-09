#!/usr/bin/env bash
export GIT_BRANCH=${GIT_BRANCH:-$(git rev-parse --abbrev-ref HEAD 2> /dev/null)}

function build {
    echo "go: building gadget"
    (docker-compose build --no-cache publisher)
    (docker-compose build gadget)

    # in theory, we could use docker cp, but I couldn't get it to actually, you know, work.
    echo "go: copying package-lock.json"
    docker-compose run --rm --volume=[] gadget cat /usr/src/app/package-lock.json | tr -d '\r' > package-lock.json

    echo "go: touching .Dockerfile"
    touch .Dockerfile
}

function gadget {
    echo "go: go-gadget $*"
    docker-compose run --rm gadget npm run "$*"
}

function requireNpm {
    if [ ! -f ./.npmrc ]; then
        echo "Must login to npm first."
        if [[ -z ${JENKINS_HOME:-} ]]; then
            confirm "May I copy your ~/.npmrc?" || exit 1
            cp ~/.npmrc .
        else
            echo "On Jenkins, this means running prepareNpmEnv() in your Jenkinsfile"
            exit 1
        fi
    fi
}

function requireBuild {
    # "-nt" means "is newer than"
    # So if package.json, this script, or the Dockerfile are newer than .Dockerfile,
    # then we will build.
    if [[ \
          Dockerfile -nt .Dockerfile \
              || package.json -nt .Dockerfile \
              || docker-entrypoint.sh -nt .Dockerfile \
              || go -nt .Dockerfile \
        ]];
    then
        echo "go: triggering a build."
        build
    fi
}

function shell {
    echo "go: shell"
    docker-compose run gadget bash
}

function main {
    set -euo pipefail
    local option=$1; shift

    case ${option} in
        build)
            requireNpm
            build
            ;;
        go-gadget|gadget|run)
            requireNpm
            requireBuild
            gadget $*
            ;;
        just-gadget)
            gadget $*
            ;;
        shell)
            shell $*
            ;;
        *)
            ./go gadget "${option}" "$@"
            ;;
    esac
}

main $*
