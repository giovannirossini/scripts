from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError
import json

class Message():
  def send(message: None):
    channel_id = "EX4MPL3"
    client = WebClient(token="xoxb-examples")
    try:
        result = client.chat_postMessage(
            channel=channel_id,
            text = "Slack Message with Python",
            blocks = json.dumps(message)
        )
        return result
    except SlackApiError as e:
        return "Error: " + str(e)