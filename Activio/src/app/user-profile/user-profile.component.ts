import { HttpClient, HttpEventType, HttpHeaders } from '@angular/common/http'; // allows for http requests
import { Component, Inject, OnInit } from '@angular/core';
import * as _ from 'cypress/types/lodash';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {
  me: any = {}
  constructor(private http: HttpClient) { }
  ngOnInit(): void {

    this.http.get<any>('http://localhost:3000/api/auth/me',{
            reportProgress: true,
            observe: 'response', // to get the post id
            withCredentials: true // authentication
    })
      .subscribe(response => {
      this.me = response.body;
        console.log(this.me.email);
        console.log(this.me.username);
        console.log(this.me.firstName)
       
    }, error => {
      console.error('Getting Account Info');
    })
    
  }
  
}
