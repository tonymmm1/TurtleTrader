//Coinbase Pro API : Coinbase accounts
package coinbase_pro

import (
    "encoding/json"
    "fmt"
    "os"
)

type Wallet struct { //struct to store API wallet
    Id string `json:"id"`
    Name string `json:"name"`
    Balance string `json:"balance"`
    Currency string `json:"currency"`
    Type string `json:"type"`
    Primary bool `json:"primary"`
    Active bool `json:"active"`
    Available_on_consumer bool `json:"available_on_consumer"`
    //Ready bool `json:"ready"`
    Wire_deposit_information struct {
        Account_number string `json:"account_number"`
        Routing_number string `json:"routing_number"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country struct {
            Code string `json:"code"`
            Name string `json:"name"`
        } `json:"bank_country"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"wire_deposit_information"`
    Swift_deposit_information struct {
        Account_number string `json:"account_number"`
        Routing_number string `json:"routing_number"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country struct {
            Code string `json:"code"`
            Name string `json:"name"`
        } `json:"bank_country"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"swift_deposit_information"`
    Sepa_deposit_information struct {
        Iban string `json:"iban"`
        Swift string `json:"swift"`
        Bank_name string `json:"bank_name"`
        Bank_address string `json:"bank_address"`
        Bank_country_name string `json:"bank_country_name"`
        Account_name string `json:"account_name"`
        Account_address string `json:"account_address"`
        Reference string `json:"reference"`
    } `json:"sepa_deposit_information"`
    Uk_deposit_information struct {
        Sort_code string `json:"sort_code"`
        Account_number string `json:"account_number"`
        Bank_name string `json:"bank_name"`
        Account_name string `json:"account_name"`
        Reference string `json:"reference"`
    } `json:"uk_deposit_information"`
    //Destination_tag_name string `json:"destination_tag_name"`
    //Destination_tag_regex string `json:"destination_tag_regex"`
    Hold_balance string `json:"hold_balance"`
    Hold_currency string `json:"hold_currency"`
}

type CryptoAddress struct { //struct to store API generated crypto address
    Id string `json:"id"`
    Address string `json:"address"`
    Address_info struct {
        Address string `json:"address"`
        Destination_tag string `json:"destination_tag"`
    } `json:"address_info"`
    Name string `json:"name"`
    Created_at string `json:"created_at"`
    Updated_at string `json:"updated_at"`
    Network string `json:"network"`
    Uri_scheme string `json:"uri_scheme"`
    Resource string `json:"resource"`
    Resource_path string `json:"resource_path"`
    Warnings []struct {
        Title string `json:"title"`
        Details string `json:"details"`
        Image_url string `json:"image_url"`
    } `json:"warnings"`
    Exchange_deposit_address bool `json:"exchange_deposit_address"`
    //Legacy_address string `json:"legacy_address"`
    //Destination_tag string `json:"destination_tag"`
    //Deposit_uri string `json:"deposity_uri"`
    //Callback_url string `json:"callback_url"`
}

/*  Coinbase accounts
*       Get all Coinbase wallets    (GET)
*       Generate crypto address     (POST)
*/

func Get_all_wallets() []Wallet { //Gets all the user's available Coinbase wallets
    path := "/coinbase-accounts"

    var wallets []Wallet

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &wallets); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get all Coinbase wallets")
    fmt.Println()
    for wallet := range wallets {
        fmt.Println("wallets[", wallet, "]")
        fmt.Println()
        fmt.Println(wallets[wallet].Id)
        fmt.Println(wallets[wallet].Name)
        fmt.Println(wallets[wallet].Balance)
        fmt.Println(wallets[wallet].Currency)
        fmt.Println(wallets[wallet].Type)
        fmt.Println(wallets[wallet].Primary)
        fmt.Println(wallets[wallet].Active)
        fmt.Println(wallets[wallet].Available_on_consumer)
        /*if wallets[wallet].Ready == (true || false) {
            fmt.Println(wallets[wallet].Ready)
        }*/
        if wallets[wallet].Wire_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Wire_deposit_information.Routing_number)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_country.Code)
            fmt.Println(wallets[wallet].Wire_deposit_information.Bank_country.Name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Wire_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Wire_deposit_information.Reference)
        }
        if wallets[wallet].Swift_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Swift_deposit_information.Routing_number)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_country.Code)
            fmt.Println(wallets[wallet].Swift_deposit_information.Bank_country.Name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Swift_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Swift_deposit_information.Reference)
        }
        if wallets[wallet].Sepa_deposit_information.Account_name != ""{
            fmt.Println(wallets[wallet].Sepa_deposit_information.Iban)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Swift)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_address)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Bank_country_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Account_address)
            fmt.Println(wallets[wallet].Sepa_deposit_information.Reference)
        }
        if wallets[wallet].Uk_deposit_information.Account_name != "" {
            fmt.Println(wallets[wallet].Uk_deposit_information.Sort_code)
            fmt.Println(wallets[wallet].Uk_deposit_information.Account_number)
            fmt.Println(wallets[wallet].Uk_deposit_information.Bank_name)
            fmt.Println(wallets[wallet].Uk_deposit_information.Account_name)
            fmt.Println(wallets[wallet].Uk_deposit_information.Reference)
        }
        /*if wallets[wallet].Destination_tag_name != "" {
            fmt.Println(wallets[wallet].Destination_tag_name)
        }
        if wallets[wallet].Destination_tag_regex != "" {
            fmt.Println(wallets[wallet].Destination_tag_regex)
        }*/
        fmt.Println(wallets[wallet].Hold_balance)
        fmt.Println(wallets[wallet].Hold_currency)
        fmt.Println()
    }

    return wallets
}

func Generate_crypto_address(account_id string) CryptoAddress { //Generates a one-time crypto address for depositing crypto.
    path := "/coinbase-accounts/" + account_id + "/addresses"

    var address CryptoAddress

    response_status, response_body := rest_post_address(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &address); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Generate crypto address")
    fmt.Println()
    fmt.Println(address.Id)
    fmt.Println(address.Address)
    if address.Address_info.Address != "" {
        fmt.Println(address.Address_info.Address)
        fmt.Println(address.Address_info.Destination_tag)
    }
    fmt.Println(address.Name)
    fmt.Println(address.Created_at)
    fmt.Println(address.Updated_at)
    if address.Network != "" {
        fmt.Println(address.Network)
    }
    if address.Uri_scheme != "" {
        fmt.Println(address.Uri_scheme)
    }
    fmt.Println(address.Resource)
    fmt.Println(address.Resource_path)
    if address.Warnings != nil {
        for warning := range address.Warnings {
            fmt.Println()
            fmt.Println("address.Warnings[", warning, "]")
            fmt.Println(address.Warnings[warning].Title)
            fmt.Println(address.Warnings[warning].Details)
            fmt.Println(address.Warnings[warning].Image_url)
        }
    }
    fmt.Println(address.Exchange_deposit_address)
    /*if address.Legacy_address != "" {
        fmt.Println(address.Legacy_address)
    }
    if address.Destination_tag != "" {
        fmt.Println(address.Destination_tag)
    }
    if address.Deposit_uri != "" {
        fmt.Println(address.Deposit_uri)
    }
    if address.Callback_url != "" {
        fmt.Println(address.Callback_url)
    }*/
    fmt.Println()

    return address
}
