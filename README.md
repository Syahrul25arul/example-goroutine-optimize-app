# Gorutine - optimize your app

## Deskription

this is a simple example of how to use goroutines to optimize your application.
in this case we create a simple app to register user, validation data from client, and send verification email to user. and you can see the process of sending an email to a user is taking a while. therefore we will use goroutines to cut the execution time of the process.

## result without use goroutine

![image](./images/time_range_for_register.png)

this is result process registration and send email without use goroutine. you can see the process took 2.8s.
