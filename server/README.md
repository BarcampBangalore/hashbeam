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
  "mysql": {
    "host": "127.0.0.1",
    "port": "3306",
    "user": "root",
    "password": "pass",
    "database": "bcb"
  },
  "firebase": {
    "topicName": "sometopicname",
    "notificationIconUrl": "https://barcampbangalore.com/bcb/apple-touch-icon.png?v=PY4NNGXQPr",
    "notificationClickedTargetUrl": "https://barcampbangalore.com/bcb/"
  }
}
```