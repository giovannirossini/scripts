import json
import boto3

boto3.setup_default_session(profile_name='default')

client = boto3.client('route53')

def createRecord(domain, source, zoneId):
    record_name= domain
    r = client.list_resource_record_sets(
        HostedZoneId=zoneId
    )
    cfZone = '' # TODO get CF HostedZoneId
    client.change_resource_record_sets(
    ChangeBatch={
            'Changes': [
                {
                    'Action': 'CREATE',
                    'ResourceRecordSet': {
                    'Name': record_name,
                    'AliasTarget': {
                        'HostedZoneId': cfZone,
                        'DNSName': source,
                        'EvaluateTargetHealth': False
                    },
                    'Type': 'A'
                },
            },
        ],
    },
    HostedZoneId=zoneId
)

createRecord('my-record.example.com','example.cloudfront.net', 'ZONEEXAMPLE')