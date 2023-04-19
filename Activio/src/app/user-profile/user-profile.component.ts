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
  posts: any[] = [];
  constructor(private http: HttpClient) { }
  ngOnInit(): void {
    this.http.get<any>('http://localhost:3000/api/posts')

      .subscribe(response => {

        this.posts = response.posts;

      }, error => {
        console.error('Error retrieving posts.');
      })

    this.http.get<any>('http://localhost:3000/api/auth/me',{
            reportProgress: true,
            observe: 'response', // to get the post id
            withCredentials: true // authentication
    })
      .subscribe(response => {
      this.me = response.body;
       
    }, error => {
      console.error('Getting Account Info');
    })
    
  }
  onDelete(id: number) {

    this.http.delete(`http://localhost:3000/api/auth/deletepost/${id}`, {
      withCredentials: true })
      .subscribe(response => {
        console.log(id);
        window.location.reload();
      }, error => {
        console.error(error);
      }
      )
  }
  
}
