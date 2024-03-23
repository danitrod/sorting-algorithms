#!/bin/bash

podman run --memory 32m --cpus 0.5 sorting > data/results-32m-0.5cpu.csv
podman run --memory 64m --cpus 0.5 sorting > data/results-64m-0.5cpu.csv
podman run --memory 128m --cpus 0.5 sorting > data/results-128m-0.5cpu.csv
podman run --memory 2g --cpus 0.5 sorting > data/results-2g-0.5cpu.csv

podman run --memory 32m --cpus 1 sorting > data/results-32m-1cpu.csv
podman run --memory 64m --cpus 1 sorting > data/results-64m-1cpu.csv
podman run --memory 128m --cpus 1 sorting > data/results-128m-1cpu.csv
podman run --memory 2g --cpus 1 sorting > data/results-2g-1cpu.csv

podman run --memory 32m --cpus 2 sorting > data/results-32m-2cpu.csv
podman run --memory 64m --cpus 2 sorting > data/results-64m-2cpu.csv
podman run --memory 128m --cpus 2 sorting > data/results-128m-2cpu.csv
podman run --memory 2g --cpus 2 sorting > data/results-2g-2cpu.csv

podman run --memory 32m --cpus 4 sorting > data/results-32m-4cpu.csv
podman run --memory 64m --cpus 4 sorting > data/results-64m-4cpu.csv
podman run --memory 128m --cpus 4 sorting > data/results-128m-4cpu.csv
podman run --memory 2g --cpus 4 sorting > data/results-2g-4cpu.csv

podman run sorting > data/results-no-limit.csv
