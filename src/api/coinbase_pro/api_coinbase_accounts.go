//Coinbase Pro API : Coinbase accounts
package coinbase_pro

import (
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

func Get_all_wallets() []byte { //Gets all the user's available Coinbase wallets
    path := "/coinbase-accounts"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func Generate_crypto_address(account_id string) []byte { //Generates a one-time crypto address for depositing crypto.
    path := "/coinbase-accounts/" + account_id + "/addresses"

    status, response := rest_post_address(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
