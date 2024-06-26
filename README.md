# go-money-api
Money tracking API written in Golang

# TODOs
1. [x] Add Update category API.
1. [x] Add Update expense API.
1. [x] Add pagination to expenses API.
1. <del>[] Add Upload category image API.<del> - For now we will used predefined set of category icons
1. [x] Create Create/Update settings API
1. [x] Create default account on user registration
1. [x] Create default categories on user registration
1. [x] Add import expenses API.
1. [x] Refactor to use DDD structure
1. [x] Add factories for DI
1. [] Add Validation.
1. [] Send Registration/confirmation email.
1. [] Add Forgotten password functionality (email + reset APIs).
1. [] Add Expiry date to auth token.
1. [] Add Refresh token API
1. [] Add Logout API
1. [x] Add settings API (choose theme, default account, pagination)
1. [] Add support for shared/multi user account (invitation)
1. [] Add tests
1. [] Use uuids/slugs instead of int ids
1. [] Add docker setup
1. [] Add Redis
1. [] Add Mailpit

Validation
    * custom validators
    * custom errors
    * email, valid address / uniqueness
    * password  - 8 symbols, 1 special symbol, 1 digit, 1 uppercase, 1 lowercase
    * permissions - user can edit/delete/view only his accounts, categories and expenses

