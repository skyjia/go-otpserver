# go-otpserver

A golang version OTP server

## API

### Get TOTP Authentication URL

**POST** `/totp/authurl`

**Request**

```json
{
    "issuer":"MyCompany",
    "accountName":"skyjia@me.com",
    "secret":"WgQacDUV24W4myXP"
}
```

**Response**

```json
{
  "authRUL": "otpauth://totp/MyCompany:skyjia@me.com?secret=K5TVCYLDIRKVMMRUK42G26KYKA&issuer=MyCompany"
}
```


### Verify TOTP Token

**POST** `/totp/verify`

**Request**

```json
{
    "token": "123456",
    "secret": "WgQacDUV24W4myXP"
}
```

**Response**

```json
{
  "isValid": true
}
```
