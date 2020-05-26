#!/bin/bash
#
# This script will get dump from the database, archive it then upload to s3 bucket.
#

HOST=localhost
USER=postgres

aws s3api create-bucket --bucket goreact-bucket

pg_dumpall -h [$HOST] \
           -U [$USER] \
           --file=postgresql_backup.sql

gzip postgres_backup.sql


S3_KEY=goreact-bucket/backups/$(date "+%Y-%m-%d")-backup.gz
aws s3 cp postgresql_backup.sql s3://$S3_KEY --sse AES256

