#!/bin/bash

# Configuration variables
PROJECT_ID="wipro-425519"
CLUSTER_NAME="automation-cluster"
ZONE="us-central1"
GSA_NAME="automation-gsa"
KSA_NAME="automation-ksa"
NAMESPACE="automation-ns"
ROLE="roles/storage.admin"  # Example role, change according to your needs

## Create GKE cluster with Workload Identity enabled
#echo "Creating GKE cluster with Workload Identity..."
#gcloud container clusters create $CLUSTER_NAME \
#  --zone $ZONE \
#  --workload-pool=wipro-425519.svc.id.goog \
#  --project $PROJECT_ID
#
## Create a Google Service Account
#echo "Creating Google Service Account..."
#gcloud iam service-accounts create $GSA_NAME \
#  --description="Service account for Kubernetes workload" \
#  --display-name=$GSA_NAME \
#  --project $PROJECT_ID
#
## Assign roles to the Google Service Account
#echo "Assigning roles to Google Service Account..."
#gcloud projects add-iam-policy-binding $PROJECT_ID \
#  --member="serviceAccount:$GSA_NAME@$PROJECT_ID.iam.gserviceaccount.com" \
#  --role=$ROLE \
#  --project $PROJECT_ID

# Create a Kubernetes Service Account in the specified namespace
echo "Creating Kubernetes Service Account..."
kubectl create serviceaccount $KSA_NAME --namespace $NAMESPACE

# Bind the Kubernetes Service Account to the Google Service Account
echo "Binding Kubernetes Service Account to Google Service Account..."
gcloud iam service-accounts add-iam-policy-binding \
  $GSA_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --role "roles/iam.workloadIdentityUser" \
  --member "serviceAccount:$PROJECT_ID.svc.id.goog[$NAMESPACE/$KSA_NAME]" \
  --project $PROJECT_ID

# Annotate the Kubernetes Service Account
echo "Annotating Kubernetes Service Account..."
kubectl annotate serviceaccount $KSA_NAME \
  --namespace $NAMESPACE \
  iam.gke.io/gcp-service-account=$GSA_NAME@$PROJECT_ID.iam.gserviceaccount.com

echo "Setup complete."
