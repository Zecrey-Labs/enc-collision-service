### POST /api/v1/crypto/getEncCollision

> 解密碰撞表，因为解密时间是线性的，如果数很大，需要的时间将非常长，通过后台设计碰撞表，将解密时间降低至常量级。

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

> 解密碰撞表，由于有时候存在多个加密值，因此传入字符串数组，进行批量解密

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
