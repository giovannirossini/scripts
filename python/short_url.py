import pyshorteners

def short_url(long_url):
    #TinyURL shortener service
    type_tiny = pyshorteners.Shortener()
    short_url = type_tiny.tinyurl.short(long_url)
    return short_url

if __name__ == '__main__':
    url = short_url("https://giovannirossini.github.io")
    print(f'The shortened URL is: {url}')
