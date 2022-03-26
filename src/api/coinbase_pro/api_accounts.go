//Coinbase Pro API : Accounts
package coinbase_pro

import (
    "fmt"
    "os"
)

type Account struct { //Get all accounts for a profile/Get a single account by id
    Id string `json:"id"`
    Currency string `json:"currency"`
    Balance string `json:"balance"`
    Hold string `json:"hold"`
    Available string `json:"available"`
    Profile_id string `json:"profile_id"`
    Trading_enabled bool `json:"trading_enabled"`
}

type Ledger struct { //struct to store API ledger
    Id string `json:"id"`
    Amount string `json:"amount"`
    Balance string `json:"balance"`
    Created_at string `json:"created_at"`
    Type string `json:"type"`
    Details map[string] interface {} `json:"details"`
}

type Hold struct { //struct to store API hold
    Id string `json:"id"`
    Created_at string `json:"created_at"`
    Amount string `json:"amount"`
    Ref string `json:"ref"`
    Type string `json:"type"`
}

type PastTransfer struct { //struct to store API past transfer
    Id string `json:"id"`
    Type string `json:"type"`
    Created_at string `json:"created_at"`
    Completed_at string `json:"completed_at"`
    Canceled_at string `json:"canceled_at"`
    Processed_at string `json:"processed_at"`
    Account_id string `json:"account_id"`
    User_id string `json:"user_id"`
    Amount string `json:"amount"`
    Idem string `json:"idem"`
    Details map[string] interface {} `json:"details"`
    User_nonce string `json:"user_nonce"`
}

/*  Accounts
*       Get all accounts for a profile      (GET)
*       Get a single account by id          (GET)
*       Get a single account's holds        (GET)
*       Get a single account's ledger       (GET)
*       Get a single account's transfers    (GET)
*/

func Get_all_accounts() []byte { //Get a list of trading accounts from the profile of the API key.
    path := "/accounts"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_single_account(account_id string) []byte { //Information for a single account.
    path := "/accounts/" + account_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_single_account_holds(account_id string) []byte { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/holds" //?limit=100" //implement limit logic later

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }
    
    return response
}

func Get_single_account_ledger(account_id string) []byte { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/ledger" //?limit=100" //implement limit logic later

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Get_single_account_transfers(account_id string) []byte { //Lists past withdrawals and deposits for an account.
    path := "/accounts/" + account_id + "/transfers" //?limit=100" //implement limit logic later

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
