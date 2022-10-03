from slack_sdk import WebClient
from slack_sdk.errors import SlackApiError
import json

def send_message(message, channel):
    client = WebClient(token="xoxb-examples")
    try:
        result = client.chat_postMessage(
            channel=channel,
            text = "Slack Message with Python",
            blocks = json.dumps(message)
        )
        return result
    except SlackApiError as e:
        return "Error: " + str(e)

if __name__ == '__main__':
    message = "My message to this channel is about..."
    send_message(message, "CH4NN3L_EX4MPL3_1D")