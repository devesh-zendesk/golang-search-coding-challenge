# Search Engine 
A Project for doing Hnadson practices on GoLang

##Description:

* This CLI Application is developed for Searching the data fron the given JSON files.

* You have 2 JSON files as below:
    user.json
        - Id
        - Name
        - CreatedAt
        - Verified

    tickets.json
        - Id
        - CreatedAt
        - Type
        - Subject
        - AssigneeId
        - Tags

* You can search from the below feilds:
    1. UserId 
    2. User Name 
    3. User Verified Flag 
    4. User Created Date 
    5. Type of the Ticket 
    6. Tag of the Ticket 
    7. Ticket Created Date 
    8. Tickets Without Assignee 
    9. exit

* Once the application starts, it will ask you by which field do you want to search. You need to insert the digit which refers to the feild you want to search.

* After that, the application will ask you to enter the search keyword and will return all the searched data you asked for.

### Example Input and Output:

```
Search By: 
        1. UserId 
        2. User Name 
        3. User Verified Flag 
        4. User Created Date 
        5. Type of the Ticket 
        6. Tag of the Ticket 
        7. Ticket Created Date 
        8. Tickets Without Assignee 
        9. exit

2
Please enter the Name you want to search
simpson

User Details:
        User ID:  75 
        Name:  Catalina Simpson 
        Created At:  2016-06-07 09:18:00 -1000 -1000 
        Verified:  true

Ticket Details:
        Ticket ID:  25d9edca-7756-4d28-8fdd-f16f1532f6ab 
        Type:  task 
        Subject:  A Problem in Cyprus 
        Created At:  2016-03-01 05:58:09 -1100 -1100 
        Tags:  [Puerto Rico Idaho Oklahoma Louisiana]


User Details:
        User ID:  76 
        Name:  Homer Simpson 
        Created At:  2016-06-07 09:18:00 -1000 -1000 
        Verified:  true

User Details:
        User ID:  77 
        Name:  Bart Simpson 
        Created At:  2016-06-07 09:18:00 -1000 -1000 
        Verified:  true

```