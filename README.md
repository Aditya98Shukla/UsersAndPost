In this repo, you have 2 json files that act as mock database such as Users and Post. The 
code in this repo basically extract information of a specific user from both Users and Post.
This is created in order to demonstrate the use of gin web framework at the basic level.

Steps to Run

1. Open VSCode with this Project as workspace. Open the terminal.
2. run these commands one by one.

`make step_1 step_2 step_3`

3. Run -> `make run` . The Output is shown below and it is running a server listening at port 8082.

![Screenshot 2023-07-19 8 34 20 PM](https://github.com/Aditya98Shukla/UsersAndPost/assets/49987108/28b8bf67-a354-4aa5-90de-763135cb0f44)

4. Open another terminal in VSCode and execute these commands one by one.

`make getUsers` -> Lists all Users as json after converting from json to struct and then struct to json.

`make getPosts` -> Lists all Posts as json after converting from json to struct and then struct to json.

5. `make getBoth ID=4`. The output is shown below
   
![Screenshot 2023-07-19 8 37 06 PM](https://github.com/Aditya98Shukla/UsersAndPost/assets/49987108/cb83eb29-9b41-4c97-975a-93bc3f3dcd51)

Title and Body is from Posts.json and remaining keys such as ID, name, lat, lng, companyName is from Users.json given the ID is 4.
