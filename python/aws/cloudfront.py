import json
import boto3

boto3.setup_default_session(profile_name='default')

cloudfront = boto3.client('cloudfront')

def removeDistributions(distributionId):
    try:
        distribution = cloudfront.get_distribution(
            Id=distributionId
        )
        distributionConfig = distribution['Distribution']['DistributionConfig']

        distributionConfig['Enabled'] = False

        if distributionConfig['Enabled'] == False:
            cloudfront.update_distribution(
                Id=distributionId,
                IfMatch=distribution['ETag'],
                DistributionConfig=distributionConfig
            )

        while(distribution['Distribution']['Status'] != 'Deployed'):
            print(f"Distribution {distributionId} still {distribution['Distribution']['Status']}...")

        print(f"Distribution {distributionId} is {distribution['Distribution']['Status']}")
        if distribution['Distribution']['Status'] == 'Deployed' and distributionConfig['Enabled'] == False:
            cloudfront.delete_distribution(
                Id=distributionId,
                IfMatch=distribution['ETag']
            )
            print(f"Distribution {distributionId} was be removed.")
    except Exception as error:
        print(error)

def addCNAME(distributionId,cname):
    try:
        distribution = cloudfront.get_distribution(
            Id=distributionId
        )
        distributionConfig = distribution['Distribution']['DistributionConfig']
        distributionConfig['Aliases']['Items'].append(cname)
        distributionConfig['Aliases']['Quantity'] += 1
        cloudfront.update_distribution(
            Id=distributionId,
            IfMatch=distribution['ETag'],
            DistributionConfig=distributionConfig
        )

        while(distribution['Distribution']['Status'] != 'Deployed'):
            print(f"Distribution {distributionId} still {distribution['Distribution']['Status']}...")

        print(f"Distribution {distributionId} is {distribution['Distribution']['Status']}")
        print(distributionConfig['Aliases']['Items'])
    except Exception as error:
        print(error)


distributionIds = 'EX4MPL3XXXXX'

removeDistributions(distributionIds)

addCNAME(distributionIds,'my-cname.example.com.')