import boto3
import csv

boto3.setup_default_session(profile_name='default')
dynamodb = boto3.resource("dynamodb")
table = dynamodb.Table("Users")

users = []
response = table.scan()
while "LastEvaluatedKey" in response:
    users += response["Items"]
    response = table.scan(ExclusiveStartKey=response["LastEvaluatedKey"])
users += response["Items"]
