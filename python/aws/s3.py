import json
import boto3

boto3.setup_default_session(profile_name='default')

s3 = boto3.client('s3')

buckets = [
    'my-bucket'
]

def removeBuckets():
    for bucket in buckets:
        try:
            s3.delete_bucket(
                Bucket=bucket
            )
            print(f'{bucket} was removed.')
        except Exception as error:
            print(error)
