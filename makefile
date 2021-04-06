SHELL := /bin/bash

export PROJECT = eatigo_assignment

# ==============================================================================
# Building containers

all: main

main:
		docker build \
			-f deployment/Dockerfile \
			-t shortly:lstest \
			.

# ==============================================================================
# Running from within docker compose

run: up build_up dev

up:
	docker-compose -f deployment/docker-compose.yml up --remove-orphans

build_up:
	docker-compose -f deployment/docker-compose.yml up --build --remove-orphans

dev:
	docker-compose -f deployment/develop.yml up --build --remove-orphans