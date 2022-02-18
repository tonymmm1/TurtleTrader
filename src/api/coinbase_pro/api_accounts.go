//Coinbase Pro API : Accounts
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type cbpAccount struct { //Get all accounts for a profile/Get a single account by id
    Id string `json:"id"`
    Currency string `json:"currency"`
    Balance string `json:"balance"`
    Hold string `json:"hold"`
    Available string `json:"available"`
    Profile_id string `json:"profile_id"`
    Trading_enabled bool `json:"trading_enabled"`
}

type cbpLedger struct { //struct to store API ledger
    Id string `json:"id"`
    Amount string `json:"amount"`
    Balance string `json:"balance"`
    Created_at string `json:"created_at"`
    Type string `json:"type"`
    Details map[string] interface {} `json:"details"`
}

type cbpHold struct { //struct to store API hold
    Id string `json:"id"`
    Created_at string `json:"created_at"`
    Amount string `json:"amount"`
    Ref string `json:"ref"`
    Type string `json:"type"`
}

type cbpPastTransfer struct { //struct to store API past transfer
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

func cbp_get_all_accounts() []cbpAccount { //Get a list of trading accounts from the profile of the API key.
    path := "/accounts"

    var accounts []cbpAccount

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &accounts); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all accounts for a profile")
    fmt.Println()
    for account := range accounts {
        fmt.Println("accounts[", account, "]")
        fmt.Println(accounts[account].Id)
        fmt.Println(accounts[account].Currency)
        fmt.Println(accounts[account].Balance)
        fmt.Println(accounts[account].Hold)
        fmt.Println(accounts[account].Available)
        fmt.Println(accounts[account].Profile_id)
        fmt.Println(accounts[account].Trading_enabled)
        fmt.Println()
    }

    return accounts
}

func cbp_get_single_account(account_id string) cbpAccount { //Information for a single account.
    path := "/accounts/" + account_id

    var account cbpAccount //store single cbpAccount

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &account); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a single account by id")
    fmt.Println()
    fmt.Println(account.Id)
    fmt.Println(account.Currency)
    fmt.Println(account.Balance)
    fmt.Println(account.Hold)
    fmt.Println(account.Available)
    fmt.Println(account.Profile_id)
    fmt.Println(account.Trading_enabled)
    fmt.Println()

    return account
}

func cbp_get_single_account_holds(account_id string) []cbpHold { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/holds" //?limit=100" //implement limit logic later

    var holds []cbpHold

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }
    if response_body == nil {
        return nil
    }

    if err := json.Unmarshal(response_body, &holds); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a single account's holds")
    fmt.Println()
    for hold := range holds {
        fmt.Println("holds[", hold,"]")
        fmt.Println(holds[hold].Id)
        fmt.Println(holds[hold].Created_at)
        fmt.Println(holds[hold].Amount)
        fmt.Println(holds[hold].Ref)
        fmt.Println(holds[hold].Type)
        fmt.Println()
    }

    return holds
}

func cbp_get_single_account_ledger(account_id string) []cbpLedger { //List the holds of an account that belong to the same profile as the API key.
    path := "/accounts/" + account_id + "/ledger" //?limit=100" //implement limit logic later

    var ledgers []cbpLedger

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &ledgers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a single account's ledger")
    fmt.Println()
    for ledger := range ledgers {
        fmt.Println(ledgers[ledger].Id)
        fmt.Println(ledgers[ledger].Amount)
        fmt.Println(ledgers[ledger].Balance)
        fmt.Println(ledgers[ledger].Created_at)
        fmt.Println(ledgers[ledger].Type)
        for k, v := range ledgers[ledger].Details {
            fmt.Println(k, ":", v)
        }
        fmt.Println()
    }

    return ledgers
}

func cbp_get_single_account_transfers(account_id string) []cbpPastTransfer { //Lists past withdrawals and deposits for an account.
    path := "/accounts/" + account_id + "/transfers" //?limit=100" //implement limit logic later

    var transfers []cbpPastTransfer

    response_status, response_body := cbp_rest_get(path)
    if response_status != CBP_STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &transfers); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get a single account's transfers")
    fmt.Println()
    for transfer := range transfers {
        fmt.Println(transfers[transfer].Id)
        fmt.Println(transfers[transfer].Type)
        fmt.Println(transfers[transfer].Created_at)
        fmt.Println(transfers[transfer].Completed_at)
        fmt.Println(transfers[transfer].Canceled_at)
        fmt.Println(transfers[transfer].Processed_at)
        fmt.Println(transfers[transfer].User_nonce)
        fmt.Println(transfers[transfer].Amount)
        for k, v := range transfers[transfer].Details {
            fmt.Println(k, ":", v)
        }
        fmt.Println()
    }

    return transfers
}
