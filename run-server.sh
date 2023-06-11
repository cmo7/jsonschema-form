#!/bin/bash

go run ./back/. &
cd front
pnpm dev
