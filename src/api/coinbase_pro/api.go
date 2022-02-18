//Coinbase Pro API
package coinbase_pro

const ( //Common return codes (https://docs.cloud.coinbase.com/exchange/docs/requests)
    _STATUS_CODE_SUCCESS int = 200           //Success
    _STATUS_CODE_BAD_REQUEST = 400           //Bad Request -- Invalid request format
    _STATUS_CODE_UNAUTHORIZED = 401          //Unauthorized -- Invalid API Key
    _STATUS_CODE_FORBIDDEN = 403             //Forbidden -- You do not have access to the requested resource
    _STATUS_CODE_NOT_FOUND = 404             //Not Found
    _STATUS_CODE_INTERNAL_SERVER_ERROR = 500 //Internal Server Error -- We had a problem with our server
)

type Config struct { //Coinbase Pro configuration
    Host string
    Key string
    Password string
    Secret string
}