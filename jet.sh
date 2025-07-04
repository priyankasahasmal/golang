#!/bin/bash
jet -source=postgres \
  -host=dpg-d1ibu73ipnbc73bfqdrg-a.oregon-postgres.render.com \
  -port=5432 \
  -user=bootcamp_db_qc7b_user \
  -password=25Zq6NeZxYyLS855V7c9dqbZjKNWU2ZH \
  -dbname=bootcamp_db_qc7b \
  -schema=public \
  -path=./.gen \
  -sslmode=require
