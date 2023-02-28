# Activio REST API Documentation

## Introduction
These endpoints are used to interact with the Activio REST API. The API is used to interact with the Activio backend.

### API Key
- The API key is used to authenticate the user. The API key is sent in the `Authorization` header of the request. The API key is a JWT token that is generated when the user logs in. The API key is valid for 1 hour. After 1 hour, the user must log in again to get a new API key. The API key can be refreshed by calling the `/api/auth/refreshtoken` endpoint. The API key is valid for 1 hour after it is refreshed.

### Naming Conventions
The following naming conventions are used throughout the documentation:
- All API routes are prefixed with `/api/`
- All routes that require authentication are prefixed with `/api/auth/`


# API Endpoints Grouped by Tag

## User Management
- POST [api/signup](#apisignup)
- POST [api/login](#apilogin)
- GET [api/auth/refreshtoken](#apiauthrefreshtoken)

## Image Management
- POST [api/images](#apiimages)
- GET [api/image/[id]](#apiimageid)

## Miscellaneous
- GET [api/alive](#apialive)


# Endpoints

### `api/signup`
- **Description:** Creates a new user account
- **Method:** POST
- **URL:** `/api/signup`
- **Request Body:**
    - `username` - The username of the user
    - `password` - The password of the user
    - `firstName` - The first name of the user
    - `lastName` - The last name of the user
    - `email` - The email of the user (Unique)
- **Response Body:**
    - `email` - The email of the user
    - `id` - The id of the user
    - `verified` - Whether the user has verified their email

### `api/login`
- **Description:** Logs in a user
- **Method:** POST
- **URL:** `/api/login`
- **Request Body:**
    - `email` - The email of the user
    - `password` - The password of the user
- **Response Body:**
    - None
- **Response Header**
    - `Authorization` - The API key of the user (JWT token)

### `api/auth/refreshtoken`
- **Description:** Refreshes the API key of the user
- **Method:** GET
- **URL:** `/api/auth/refreshtoken`
- **Request Body:**
    - None
- **Response Header**
    - `Authorization` - The API key of the user (JWT token)

### `api/images`
- **Description:** Upload multiple images
- **Method:** POST
- **URL:** `/api/images`
- **Request Body:**
    - `images` - The images to upload multipart/form
- **Response Body:**
    - `images` - The hashed names of the images

### `api/image/[id]`
- **Description:** Get an image
- **Method:** GET
- **URL:** `/api/image/[id]`
- **Response Body:**
    - `image` - The image

## `api/alive`
- **Description:** Check if the server is alive
- **Method:** GET
- **URL:** `/api/alive`
- **Response Body:**
    - `alive` - Whether the server is alive

