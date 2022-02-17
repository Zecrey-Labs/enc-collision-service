### POST /api/v1/crypto/getEncCollision

> Decrypting the collision table, because the decryption time is linear, if the enc is large, the time will be very long. The collision table is designed in the background to reduce the decryption time to a constant level.

#### Request Params

| Name     | Type   | Size | Comment             |
| -------- | ------ | ---- | ------------------- |
| enc_data | string |      | raw encryption data |

#### Response Struct
```json
{
    "status": int, // status for the request
    "msg": string, // response message
    "err":{ //err struct 
    	"err_type": int, // err type
    	"err_msg": string, // error msg
	  },
    "result": { //result struct
  		"colision_result": int // result
    }
}
```

### POST /api/v1/crypto/getEncCollisionBatches

> Decrypt the collision table, because sometimes there are multiple encrypted values, so pass in a string array for batch decryption

#### Request Params

| Name             | Type     | Size | Comment                   |
| ---------------- | -------- | ---- | ------------------------- |
| enc_data_batches | []string |      | raw encryption data array |

#### Response Struct
```json
{
    "status": int, // status for the request
    "msg": string, // response message
    "err":{ //err struct 
    	"err_type": int, // err type
    	"err_msg": string, // error msg
	  },
    "result": [{ //result struct
                "enc_data": string
  		"colision_result": int // result
    }]
}
```
