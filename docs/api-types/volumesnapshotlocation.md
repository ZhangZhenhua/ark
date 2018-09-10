# Ark Volume Snapshot Location

## Volume Snapshot Location

Ark can store volume snapshots in multiple persistent volume types in a single cluster. These are represented in the cluster via the `VolumeSnapshotLocation` CRD.

Ark must have at least one `VolumeSnapshotLocation`  per block storage provider. 

A sample YAML `VolumeSnapshotLocation` looks like the following:

```yaml
apiVersion: ark.heptio.com/v1
kind: VolumeSnapshotLocation
metadata:
  name: default
  namespace: heptio-ark
spec:
  provider: aws
  config:
    region: us-west-2
```

### Parameter Reference

The configurable parameters are as follows:

#### Main config parameters

| Key | Type | Default | Meaning |
| --- | --- | --- | --- |
| `provider` | String (Ark natively supports `aws`, `gcp`, and `azure`. Other providers may be available via external plugins.)| Required Field | The name for whichever cloud provider will be used to actually store the volume. |
| `config/region` | map[string]string<br><br>(See the corresponding [AWS][0], [GCP][1], and [Azure][2]-specific configs or your provider's documentation.) | None (Optional) | Configuration keys/values to be passed to the cloud provider for backup storage. |

#### AWS

##### config/region 

| Key | Type | Default | Meaning |
| --- | --- | --- | --- |
| `region` | string | Empty | *Example*: "us-east-1"<br><br>See [AWS documentation][3] for the full list.<br><br>Queried from the AWS S3 API if not provided. |

#### Azure

##### config/region ?? 

| Key | Type | Default | Meaning |
| --- | --- | --- | --- |

#### GCP

No parameters required.

[0]: #aws
[1]: #gcp
[2]: #azure
[3]: http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions