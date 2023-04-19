import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  me: any = {}
  constructor(private http: HttpClient) { }
  ngOnInit(): void {
    this.http.get<any>('http://localhost:3000/api/auth/me')
      .subscribe(response => {

        this.me = response.me;
        console.log(this.me.email)
      }, error => {
        console.error('Getting Account Info');
      })
  }
  
}
