

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







