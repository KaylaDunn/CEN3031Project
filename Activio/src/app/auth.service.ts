import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable()
export class AuthService {
  private _registerUrl = "http://localhost:3000/api/signup";
  private _loginUrl = "http://localhost:3000/api/login";
  private _userUpdate = "http://localhost:3000/api/auth/updateuser";
  private _imgUpload = "http://localhost:3000/api/auth/setprofileimage"
  constructor(private http: HttpClient) { }
  registerUser(user: {}) {
    return this.http.post<any>(this._registerUrl, user)
  }
  loginUser(user: {}) {
    return this.http.post<any>(this._loginUrl, user, { withCredentials : true })
  }
  User(user: {}) {
    return this.http.put<any>(this._userUpdate,user)
  }
  ProfilePic(img: {}) {
    return this.http.post<any>(this._imgUpload, img)
  }
  
}
