#!/bin/bash

# Set variables
REGION="us-central1"
PROJECT_ID="go-personal-site"
REPO="my-repo"
IMAGE="my-site"
MAX_IMAGES_TO_KEEP=5

# List all image digests with their tags, sorted by creation time descending
IMAGES=$(gcloud artifacts docker images list \
  "$REGION-docker.pkg.dev/$PROJECT_ID/$REPO/$IMAGE" \
  --include-tags \
  --format="get(digest)" \
  --sort-by="~CREATE_TIME")

COUNT=0
for DIGEST in $IMAGES; do
  ((COUNT++))
  if [[ $COUNT -le $MAX_IMAGES_TO_KEEP ]]; then
    echo "Keeping: $DIGEST"
    continue
  fi

  echo "Deleting: $DIGEST"
  gcloud artifacts docker images delete \
    "$REGION-docker.pkg.dev/$PROJECT_ID/$REPO/$IMAGE@$DIGEST" \
    --quiet
done
