SUPEREGO_ROOT="$(dirname "${BASH_SOURCE}")/.."

kubectl create secret generic service-account-creds \
  --from-file=$SUPEREGO_ROOT/service-account-creds.json