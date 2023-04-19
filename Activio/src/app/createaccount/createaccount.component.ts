import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-createaccount',
  templateUrl: './createaccount.component.html',
  styleUrls: ['./createaccount.component.css']
})
export class CreateaccountComponent implements OnInit {
  registerUserData: any = {}
  constructor(private _auth: AuthService, 
    private _router: Router,) { }
  ngOnInit() {

  }
  
  registerUser() {
    this._auth.registerUser(this.registerUserData)
      .subscribe(
        res => {
          this._router.navigate(['/create-account-verif']);
        },
        err => console.log(err)
    )
  }
}
