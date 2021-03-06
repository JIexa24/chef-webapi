import CRYPT from 'node-rsa';

export const ApiURL = '/api';
// export const ApiURL = 'http://localhost:3001/api';

export const ProfileURL = "/profile";

const CRYPTKEY="-----BEGIN PUBLIC KEY-----" +
"MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAutAUDazV/mSQ6mreFrig" +
"MILx7Hvia/NdZ+Knxm5TakiLmbpMs6hZz46hV5zOsO+xXxehk/KF/2rPFh9JB+wL" +
"f+CKIJKtm5BeaYYqyJz2JwufQbEDiXod22SJTJEGL1vcS2s3LduMP+QKDTGBLjJ6" +
"kJqHUNWYVK3d7wE7FBVXOBC7U1q7upBCcgSilNRSFIraAXPZ5lzXD5GXdOL5rQp4" +
"UAnT7IUn9OaYcEaOkkRNkEfqUj1p1wFgBon1VBCucRDAisGKXy4vKilg6Z3PYNRN" +
"BJdbky8YdOEFm7ThtxNgzUINk0eoxqeRxvPoxdUayC84/xRljawcLVfvzb8J9yRO" +
"bYrQpKi5mIYY37Ujmxy25aTKxxkZTs2c1NiqkgNNu7ecFag0i+fLWGC3LcqlQiIm" +
"7xKZs5cNsB7Gix+JVzyQdeuC9ylK7U2eaXWArMx6a3nyjzWl92oV3xFLwpElfdUx" +
"ix/OLQNYT/66aV2ztfy4KVfDlZKIJBO/ghQZj4kk8p2FPKHPza67yEDPbWSZqzMJ" +
"6g4k4I20ZhjEBZT/UDee7AT2xCPVdVEbqvdlYkO1dSll4dYM2m2yW/yUXxb1gjHu" +
"bd7N4loGmBCp3mieTsnln35OBGGtHp/N3NrcbONWczsiRqxW+YG3NzY9rt6gyNCX" +
"IFn3POPnz4iWSLi4rh2wm10CAwEAAQ==" +
"-----END PUBLIC KEY-----"

export const ENCRYPT = (data) => {
  const keys = new CRYPT()
  keys.importKey(CRYPTKEY, 'pkcs8-public-pem')
  keys.setOptions({
      environment: 'browser',
      encryptionScheme: {
          'hash': 'sha256',
      }
  })
  const encryptedData = keys.encrypt(data, 'base64')
  return encryptedData
}  