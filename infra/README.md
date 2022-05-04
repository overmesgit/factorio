```bash
gcloud compute instances create r0c0 --image-family cos-stable --image-project cos-cloud --metadata-from-file user-data=init --metadata=cos-metrics-enabled=true --zone=asia-northeast1-a --machine-type=e2-micro --project=factorio2022
```