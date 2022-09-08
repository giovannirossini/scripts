import json
import boto3

boto3.setup_default_session(profile_name='default')

acm = boto3.client('acm')

certificates = [
    'arn:aws:acm:us-east-1:YOUR_ACCOUNT_ID:certificate/ACM_ID'
]

def removeCertificates():
    for certificate in certificates:
        try:
            acm.delete_certificate(
                CertificateArn=certificate
            )
            print(f"Certificate {certificate} was be removed.")
        except Exception as error:
            print(error)
