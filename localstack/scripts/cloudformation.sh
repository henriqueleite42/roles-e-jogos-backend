#!/bin/bash

aws configure set cli_follow_urlparam false

# ---------------------------
#
#  Cloudformation
#
# ---------------------------

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-vpc \
	--template-file "/etc/localstack/cloudformation/vpc.yaml"

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-s3 \
	--template-file "/etc/localstack/cloudformation/s3.yaml"

awslocal cloudformation deploy \
	--stack-name dev-rolesejogos-cloudfront \
	--template-file "/etc/localstack/cloudformation/cloudfront.yaml"

# ---------------------------
#
#  Parameter Store
#
# ---------------------------

awslocal ssm put-parameter \
	--name dev-port \
	--type String \
	--value 3001

awslocal ssm put-parameter \
	--name dev-website-url \
	--type String \
	--value http://localhost:3000

awslocal ssm put-parameter \
	--name dev-database-url \
	--type String \
	--value postgres

awslocal ssm put-parameter \
	--name dev-ludopedia-client-id \
	--type String \
	--value $LUDOPEDIA_CLIENT_ID

awslocal ssm put-parameter \
	--name dev-ludopedia-redirect-uri \
	--type String \
	--value $LUDOPEDIA_REDIRECT_URI

awslocal ssm put-parameter \
	--name dev-google-client-id \
	--type String \
	--value $GOOGLE_CLIENT_ID

awslocal ssm put-parameter \
	--name dev-google-redirect-uri \
	--type String \
	--value $GOOGLE_REDIRECT_URI

awslocal ssm put-parameter \
	--name dev-medias-s3-bucket-name \
	--type String \
	--value dev-rolesejogos-medias

awslocal ssm put-parameter \
	--name dev-medias-cloudfront-url \
	--type String \
	--value foo

awslocal ssm put-parameter \
	--name dev-email-address-system-messages \
	--type String \
	--value no-reply@rolesejogos.com.br

awslocal ssm put-parameter \
	--name dev-name-system-messages \
	--type String \
	--value "Rolês & Jogos"

awslocal ssm put-parameter \
	--name dev-email-template-sign-in-otp \
	--type String \
	--value foo

# ---------------------------
#
#  Secret Manager
#
# ---------------------------

awslocal secretsmanager create-secret \
	--name dev-ludopedia-client-secret \
	--secret-string $LUDOPEDIA_CLIENT_SECRET

awslocal secretsmanager create-secret \
	--name dev-google-client-secret \
	--secret-string $GOOGLE_CLIENT_SECRET

awslocal secretsmanager create-secret \
	--name dev-database-username-password \
	--secret-string '{"username":"username","password":"password"}'
