from cgitb import text
from pygame import mixer
from gtts import gTTS

def text_to_speech(text, output):
   tts = gTTS(text)
   tts.save(output + '.mp3')
   mixer.init()
   mixer.music.load(output + '.mp3')
   mixer.music.play()

if __name__ == '__main__':
    text_to_speech('This text now is an audio', 'text_to_speech')
