# Configuration

```json
{
  "app": {
    "port": "3000",
    "jwtSecret": "bigmansathyabhat",
    "admins": [
      { "username": "admin1", "password": "password1" },
      { "username": "admin2", "password": "password2" }
    ]
  },
  "mongo": {
    "url": "someurl:someport/somedatabase",
    "user": "bcb19",
    "password": "bigmansathyabhat1"
  },
  "firebase": {
    "topicName": "sometopicname",
    "notificationIconUrl": "https://barcampbangalore.com/bcb/apple-touch-icon.png?v=PY4NNGXQPr",
    "notificationClickedTargetUrl": "https://barcampbangalore.com/bcb/"
  }
}
```