#!/bin/bash

/go/bin/migrate  -url $DATABASE_URL -path ./db/migrations up
go-wrapper run
