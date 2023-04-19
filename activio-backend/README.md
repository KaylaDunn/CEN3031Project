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
- DELETE [api/auth/deleteaccount](#apiauthdeleteaccount)
- GET [api/auth/me](#apiauthme)
- PUT [api/auth/updateuser](#apiauthupdateuser)

## Image Management
- POST [api/auth/setprofileimage](#apiauthsetprofileimage)
- GET [api/image/[id]](#apiimageid)

## Post Management
- GET [api/posts](#apiposts)
- GET [api/post/[id]](#apipostid)
- GET [api/posts/location/[location]](#apipostslocationlocation)
- GET [api/posts/user/[id]](#apipostsuserid)
- POST [api/auth/createpost](#apiauthcreatepost)
- PUT [/addImageToPost/:id](#addimagetopostid)
- DELETE [api/auth/deletepost/[id]](#apiauthdeletepostid)
- PUT [api/auth/likepost/[id]](#apiauthlikepostid)
- PUT [api/auth/comment/[id]](#apiauthcommentid)

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
    - `birthday` - The birthday of the user (Type: Date String)
    - `phoneNumber` - The phone number of the user
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

### `api/auth/deleteaccount`
- **Description:** Deletes the user's account and all of their data
- **Method:** DELETE
- **URL:** `/api/auth/deleteaccount`
- **Request Body:**
    - None
- **Response Body:**
    - On Success:
        - `message` - "User deleted successfully"
    - On Failure:
        - `error` - "Unauthorized"

### `api/auth/me`
- **Description:** Gets the user's information
- **Method:** GET
- **URL:** `/api/auth/me`
- **Request Body:**
    - None
- **Response Body:**
    - `email` - The email of the user
    - `id` - The id of the user
    - `verified` - Whether the user has verified their email
    - `firstName` - The first name of the user
    - `lastName` - The last name of the user
    - `birthday` - The birthday of the user (Type: Date String)
    - `phoneNumber` - The phone number of the user


### `api/auth/updateuser`
- **Description:** Updates the user's information
- **Method:** PUT
- **URL:** `/api/auth/updateuser`
- **Request Body:**
    - `firstName` - The first name of the user (Optional)
    - `lastName` - The last name of the user (Optional)
    - `birthday` - The birthday of the user (Type: Date String) (Optional)
    - `phoneNumber` - The phone number of the user (Optional)
    - `username` - The username of the user (Optional)
- **Response Body:**
    - On Success:
        - `message` - "User updated successfully"
    - On Failure:
        - `error` - "Unauthorized"
    
### `api/auth/setprofileimage`
- **Description:** Sets the user's profile image
- **Method:** POST
- **URL:** `/api/auth/setprofileimage`
- **Request Body:**
    - `image` - The image to set as the profile image
- **Response Body:**
    - On Success:
        - `message` - "Profile image set successfully"
        - `filename` - The name of the image on the server
    - On Failure:
        - `error` - "Unauthorized"

### `api/image/[id]`
- **Description:** Get an image
- **Method:** GET
- **URL:** `/api/image/[id]`
- **Response Body:**
    - `image` - The image

### `api/posts`
- **Description:** Get all posts
- **Method:** GET
- **URL:** `/api/posts`
- **Response Body:**
    - `posts` - Up to 10 posts in the database

### `api/post/[id]`
- **Description:** Get a post
- **Method:** GET
- **URL:** `/api/post/[id]`
- **Response Body:**
    - `post` - The post

### `api/posts/location/[location]`
- **Description:** Get all posts in a location 
- **Method:** GET
- **URL:** `/api/posts/location/[location]`
- **Response Body:**
    - `posts` - All posts in the location

### `api/posts/user/[id]`
- **Description:** Get all posts by a user
- **Method:** GET
- **URL:** `/api/posts/user/[id]`
- **Response Body:**
    - `posts` - All posts by the user

### `api/addimagetopost/[id]`
- **Description:** Adds images to a post
- **Method:** PUT
- **URL:** `/api/addimagetopost/[id]`
- **Request Body:**
    - `images` - The images to add to the post
- **Response Body:**
    - `filenames` - The names of the images that were added to the post

### `api/auth/createpost`
- **Description:** Create a post
- **Method:** POST
- **URL:** `/api/auth/createpost`
- **Request Body:**
    - `Description` - The description of the post
    - `Longitude` - The longitude of the post (Type: Double)
    - `Latitude` - The latitude of the post (Type: Double)
    - `LocationName` - The name of the location of the post
- **Response Body:**
    - `post` - The post that was created

### `api/auth/deletepost/[id]`
- **Description:** Delete a post and all related items by id
- **Method:** DELETE
- **URL:** `/api/auth/deletepost/[id]`
- **Request Body:**
    - None
- **Response Body:**
    - On Success:
        - `message` - "Post deleted successfully"
    - On Failure:
        - `error` - "Unauthorized"
        - `error` - "Post not found"

### `api/auth/likepost/[id]`
- **Description:** Like a post
- **Method:** PUT
- **URL:** `/api/auth/likepost/[id]`
- **Request Body:**
    - None
- **Response Body:**
    - On Success:
        - `message` - "Post liked successfully"
    - On Failure:
        - `error` - "Unauthorized"
        - `error` - "Post not found"
        - `error` - "Post already liked"

### `api/auth/comment/[id]`
- **Description:** Comment on a post
- **Method:** PUT
- **URL:** `/api/auth/comment/[id]`
- **Request Body:**
    - `comment` - The comment to add to the post
- **Response Body:**
    - On Success:
        - `message` - "Comment added successfully"
    - On Failure:
        - `error` - "Unauthorized"
        - `error` - "Post not found"

## `api/alive`
- **Description:** Check if the server is alive
- **Method:** GET
- **URL:** `/api/alive`
- **Response Body:**
    - `alive` - Whether the server is alive

