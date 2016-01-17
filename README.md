# go-otpserver

A golang version of simple OTP server.  

> Reference: [Google Authenticator](https://github.com/google/google-authenticator)

## API

### Get TOTP Authentication URL

**POST** `/totp/authurl`

**Request**

```json
{
    "issuer":"My Company",
    "accountName":"foo@bar.com",
    "secret":"WgQacDUV24W4myXP"
}
```

**Response**

```json
{
  "authURL": "otpauth://totp/My%20Company:foo@bar.com?issuer=My+Company&secret=K5TVCYLDIRKVMMRUK42G26KYKA"
}
```

Scan the QRCode below with Google Authenticator:

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
