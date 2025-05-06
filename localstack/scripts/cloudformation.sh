#!/bin/bash

aws configure set cli_follow_urlparam false

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-vpc \
	--template-file "/etc/localstack/cloudformation/vpc.yaml"

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-s3 \
	--template-file "/etc/localstack/cloudformation/s3.yaml"

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-cloudfront \
	--template-file "/etc/localstack/cloudformation/cloudfront.yaml"

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-ssm \
	--template-file "/etc/localstack/cloudformation/ssm.yaml" \
	--parameter-overrides GoogleClientIdValue=$GOOGLE_CLIENT_ID GoogleRedirectUriValue=$GOOGLE_REDIRECT_URI

awslocal secretsmanager create-secret \
	--name dev-google-client-secret \
	--secret-string $GOOGLE_CLIENT_SECRET

awslocal secretsmanager create-secret \
	--name dev-database-username-password \
	--secret-string '{"username":"username","password":"password"}'
