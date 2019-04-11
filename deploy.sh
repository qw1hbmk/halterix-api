#!/bin/bash

argument="$1"

display_usage() {
  echo
  echo "Usage: $0 [target]"
  echo
  echo " dev   deploy to dev project"
  echo " rnd   deploy to rnd project"
  echo " prod  deploy to prod project"
  echo " all  deploy to all project"
  echo
}

deploy_to_dev(){
    echo "deploying to: dev (halterix-dev)"
    gcloud app deploy --project halterix-dev
}

deploy_to_rnd(){
    echo "deploying to: rnd (spars-9-axis)"
    gcloud app deploy --project spars-9-axis
}

deploy_to_prod(){
    echo "deploying to: prod (halterix-prod-83363)"
    gcloud app deploy --project halterix-prod-83363 
}

case $argument in
    all)
      deploy_to_dev
      deploy_to_rnd
      deploy_to_prod
      ;;
    dev)
      deploy_to_dev
      ;;
    rnd)
      deploy_to_rnd
      ;;
    prod)
      deploy_to_prod
      ;;
    *)
      raise_error "Unknown argument: ${argument}"
      display_usage
      ;;
esac