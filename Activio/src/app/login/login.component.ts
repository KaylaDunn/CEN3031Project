import { Component,OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginUserData: any = {}
  constructor(
    private _auth: AuthService,
    private _router: Router,
  ) { }
  ngOnInit() {

  }
  loginUser() {
    this._auth.loginUser(this.loginUserData)
    .subscribe(
      res => {
          console.log(this.loginUserData);
          this._router.navigate(['/logsuccess']);
        },
          
        err => console.log(err)
        )
      
  }
}

