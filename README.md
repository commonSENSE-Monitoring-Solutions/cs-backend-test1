# commonSENSE Backend Development Task

## Task Description

Your task is to write an application to migrate data from the `user` table to the `user_accounts` table. The structure of the tables is shown in the `initdb.sql` file. The fields `created_at`, `updated_at`, `deleted_at` and `last_logged_in` in the `user` table should be treated as UNIX timestamp fields. You may structure the application how you wish, however the following points will be taken into consideration when reviewing your submission:

1. Application structure, architecture and design
2. Code quality and readability
3. Testing (both unit and integration)

You may use any frameworks or libraries you feel appropriate, however we recommend [Gorm](https://gorm.io) for interacting with the database and [Stretchr Testify](https://github.com/stretchr/testify) for testing.

## Submission

To submit your work, pull this repository and create a new branch with your name (e.g. `john_smith`). When you are ready to submit your work, create a pull request from your branch to the main branch. In the description of the pull request, please include the following information:

* Full Name
* Email Address that you applied for the position with (or the email address you included in your C.V.)