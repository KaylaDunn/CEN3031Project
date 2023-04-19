import { Component } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
@Component({
  selector: 'app-edit-profile-info',
  templateUrl: './edit-profile-info.component.html',
  styleUrls: ['./edit-profile-info.component.css']
})
export class EditProfileInfoComponent {
  UserData: any = {}
  constructor(private http: HttpClient, private dialog: MatDialog, private _auth: AuthService,
    private _router: Router) { }
  ngOnInit() {

  }

  User() {
    this.http.put(`http://localhost:3000/api/auth/updateuser`, null, {
      withCredentials: true
    })
      .subscribe(
        res => {
          console.log(this.UserData.email);
          console.log(this.UserData.username);
          console.log(this.UserData.firstName)
          console.log(this.UserData.lastName)
          console.log(this.UserData.birthday)
          console.log(this.UserData.phoneNumber)
          this._router.navigate(['/create-account-verif']);
        },
        err => console.log(err)
      )
  }

}
