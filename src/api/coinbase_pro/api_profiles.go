//Coinbase Pro API : Profiles
package coinbase_pro

import (
    "encoding/json"
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

func Get_profiles() []Profile{
    path := "/profiles"

    var profiles []Profile

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profiles); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get profiles")
    fmt.Println()
    for profile := range profiles {
        fmt.Println("profiles[", profile, "]")
        fmt.Println(profiles[profile].Id)
        fmt.Println(profiles[profile].User_id)
        fmt.Println(profiles[profile].Name)
        fmt.Println(profiles[profile].Active)
        fmt.Println(profiles[profile].Is_default)
        fmt.Println(profiles[profile].Has_margin)
        fmt.Println(profiles[profile].Created_at)
    }
    fmt.Println()

    return profiles
}

func create_profile(name string) Profile {
    path := "/profiles"

    var profile Profile

    response_status, response_body := rest_post_create_profile(path, name)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profile); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Create a profile")
    fmt.Println()
    fmt.Println(profile.Id)
    fmt.Println(profile.User_id)
    fmt.Println(profile.Name)
    fmt.Println(profile.Active)
    fmt.Println(profile.Is_default)
    fmt.Println(profile.Has_margin)
    fmt.Println(profile.Created_at)
    fmt.Println()

    return profile
}

func transfer_funds_profiles(from string, to string, currency string, amount string) string {
    path := "/profiles/transfer"

    var profile string

    response_status, response_body := rest_post_transfer_funds_profiles(path, from, to, currency, amount)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    profile = string(response_body) //convert to string

    //debug
    fmt.Println("Transfer funds between profiles")
    fmt.Println()
    fmt.Println(profile)
    fmt.Println()

    return profile
}

func Get_profile(profile_id string) Profile {
    path := "/profiles/" + profile_id

    var profile Profile

    response_status, response_body := rest_get(path)
    if response_status != _STATUS_CODE_SUCCESS {
        fmt.Println("ERROR REST GET status code: ", response_status)
        os.Exit(1)
    }

    if err := json.Unmarshal(response_body, &profile); err != nil { //JSON unmarshal REST response body to store as struct
        fmt.Println("ERROR decoding REST response")
        os.Exit(1)
    }

    //debug
    fmt.Println("Get profile by id")
    fmt.Println()
    fmt.Println(profile.Id)
    fmt.Println(profile.User_id)
    fmt.Println(profile.Name)
    fmt.Println(profile.Active)
    fmt.Println(profile.Is_default)
    fmt.Println(profile.Has_margin)
    fmt.Println(profile.Created_at)
    fmt.Println()

    return profile
}
