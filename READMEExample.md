<!--@startuml

skinparam ActivityStartColor White
skinparam ActivityEndColor Red
skinparam ActivityStartBorderColor Black

skinparam note {
    BorderColor black
    BackgroundColor white
}

skinparam activity {
    BorderColor black
    BackgroundColor #ADD8DB
    EndColor black
    
}

skinparam activityDiamond {
    BorderColor black
    BackgroundColor #ADD8DB
}

start
-[#black]->
:Parse and Validate Command;
note right
  Slackbot
end note
-[#black]->
if (Valid?) then (no)
    -[#black]->
    :Send error message;
    -[#black]->
    end
else (yes)
    -[#black]->
    partition Deploy {
        note left
            Helm v3 Tillerless
        end note
        :Checkout repo;
        -[#black]->
        note right
            Github
        end note
        :Decrypt secrets;
        -[#black]->
        note right
            Sops
        end note
        :Deploy Umbrella Charts;
        -[#black]->
        note right
            Process Manager
            Hub Next
        end note
        
}
:Send response;
-[#black]->
end
partition "Cron Job Delete" {
    start
    -[#black]->
    note right
        Every 30 seconds
    end note
    :Check if any release exceeds time to live;
    -[#black]->
    :Delete outdated releases;
    -[#black]->
    end
}

endif

@enduml-->
![](https://plantuml.signavio.com/png/UDgSKqrhsq0KlE-lU2GNvn0ev59iQ8cjq59e2NK_3gM7bNPiBLxj4s-VtPZY_ruanRBToe5wRtuxdfcTUJ9fvMDjnLPqNwZVUDtErOhEEB3Gzz8h1Y0FqNM0pt3d6AOi3jB1fi4MAtE4ZAoWtuQ8w1YNjxXsQelLKdYz5_hLEjaTSzXdjV_YkRxFijVPjBthNt603ojatbOStKL23M4JXxdzSTsYVzw-CUD7AmbaewDlDdXd5JJZghBHJKoRZVXbgOPetlmbPpM8hatjXCSlQDGIlBqXBH5f5FccjNECQdwFvmq1H5YeGafsYSaW4D4PX0GQRP66kMehwjLpf0nru7q1pMcz1ooqdpJdFK95cplwuaE01AJKN_VFsW_6inB5YjTAWfed_Vnl4xtQFgd3UUUrNESlI6GeP5ih9HG2JHU9pBbEBqgqoNojSa48bcQbbGkL7eKBf4GVRRHBo6cowvm-uKdFEp4xqtKlI3N7XCd91ZLr7_gycWb7-i0vPGXGN3rNsstmaEy1pmyRo9RkNZNnSdJ9_EEhwvZyWcpSaY309X2U2i0bKb-1b2dupPaT7NVkYDVQV3fkJv7Eh_JED4_r2sF-1000___-2dBQ)