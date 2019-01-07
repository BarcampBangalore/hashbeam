# `config.json`:

```json
{
  "app": {
    "port": "3000",
    "jwtSecret": "bigmansathyabhat",
    "admins": {
      "admin1": "password1",
      "admin2": "password2"
    }
  },
  "mySql": {
    "host": "127.0.0.1",
    "port": "3306",
    "user": "root",
    "password": "pass$123",
    "database": "bcb"
  },
  "twitter": {
    "consumer_key": "",
    "consumer_secret": "",
    "access_token": "",
    "access_token_secret": "",
    "textToTrack": "#bcb19"
  },
  "fcm": {
    "topicName": "sometopicname",
    "notificationIconUrl": "https://barcampbangalore.com/bcb/apple-touch-icon.png?v=PY4NNGXQPr",
    "notificationClickedTargetUrl": "https://barcampbangalore.com/bcb/"
  }
}
```

## Other things you need:

- Firebase Service Key JSON file -- in the project root as `firebase-service-key.json`
