import requests
import boto3
from botocore.exceptions import ClientError
from datetime import datetime


def my_public_ip():
    url = "https://ifconfig.me"
    response = requests.request("GET", url)
    public_ip = (response.text + "/32")
    return public_ip

def update_security_group_rule(ip,security_group_id,security_group_rule_id, from_port, to_port):
    now = datetime.now()
    date = now.strftime("%Y-%m-%d")
    rule = [
        {
            'SecurityGroupRuleId': security_group_rule_id,
            'SecurityGroupRule': {
                'IpProtocol': 'tcp',
                'FromPort': from_port,
                'ToPort': to_port,
                'CidrIpv4': f'{ip}',
                'Description': f'Giovanni: Rule updated at {date}.'
            }
        }
    ]
    try:
        response = ec2.modify_security_group_rules(GroupId=security_group_id, SecurityGroupRules=rule)
        print(f"Response code = {response['ResponseMetadata']['HTTPStatusCode']}")
        print(f"The rule was updated successfully!")
    except ClientError as e:
        print(e)


boto3.setup_default_session(profile_name='default')
ec2 = boto3.client('ec2')


if __name__ == "__main__":
    sg = "sg-example"
    sgr_id = "sgr-example"

    ip = my_public_ip()

    update_security_group_rule(ip,sg,sgr_id,22,22)
