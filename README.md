# go-otpserver

A golang version OTP server.  

> Reference: [Google Authenticator](https://github.com/google/google-authenticator)

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

![](url_qrcode.jpg)


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
