//Coinbase Pro API : Profiles
package coinbase_pro

import (
    "fmt"
    "os"
)

type Profile struct { //Get profiles/Create a profile/Get profile by id/Get profile by id
    Id string `json:"id"`
    User_id string `json:"user_id"`
    Name string `json:"name"`
    Active bool `json:"active"`
    Is_default bool `json:"is_default"`
    Has_margin bool `json:"has_margin"`
    Created_at string `json:"created_at"`
}

/*  Profiles
*       Get profiles                    (GET)
*       Create a profile                (POST)
*       Transfer funds between profiles (POST)
*       Get profile by id               (GET)
*       Rename a profile                (PUT)
*       Delete a profile                (PUT)
*/

func Get_profiles() []byte {
    path := "/profiles"

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func create_profile(name string) []byte {
    path := "/profiles"

    status, response := rest_post_create_profile(path, name)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}

func transfer_funds_profiles(from string, to string, currency string, amount string) []byte {
    path := "/profiles/transfer"

    status, response := rest_post_transfer_funds_profiles(path, from, to, currency, amount)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }
    
    return response
}

func Get_profile(profile_id string) []byte {
    path := "/profiles/" + profile_id

    status, response := rest_get(path)
    if status != STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", status)
        os.Exit(1)
    }

    return response
}
