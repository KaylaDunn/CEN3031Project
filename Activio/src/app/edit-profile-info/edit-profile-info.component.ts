import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { HttpClient, HttpEventType, HttpHeaders } from '@angular/common/http';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
@Component({
  selector: 'app-edit-profile-info',
  templateUrl: './edit-profile-info.component.html',
  styleUrls: ['./edit-profile-info.component.css']
})
export class EditProfileInfoComponent implements OnInit {
  
  UserData: any = {}
  
  constructor(private http: HttpClient, private dialog: MatDialog, private _auth: AuthService,
    private _router: Router) { }

 
  ngOnInit() : void {
    
  }
 

  /*ProfilePic() {
    this._auth.ProfilePic(this.ImgData)
      .subscribe(
        res => {

        },

        err => console.log(err)
      )

  }*/

  User() {
    const headers = new HttpHeaders({
      'Content-Type': 'application/json'
    })
    console.log(this.UserData);
    this.http.put(`http://localhost:3000/api/auth/updateuser`, {
      withCredentials: true,
      headers: headers,
    })
      .subscribe(
        res => {
          console.log(this.UserData.email);
          console.log(this.UserData.username);
          console.log(this.UserData.firstName)
          console.log(this.UserData.lastName)
          console.log(this.UserData.birthday)
          console.log(this.UserData.phoneNumber)
          this._router.navigate(['/user-profile']);
        },
        err => console.log(err)
      )
  }

}
