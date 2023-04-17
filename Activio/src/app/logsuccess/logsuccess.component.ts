import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-logsuccess',
  templateUrl: './logsuccess.component.html',
  styleUrls: ['./logsuccess.component.css'],
})
export class LogsuccessComponent {

  posts: any[] = [];
  
  constructor(private http: HttpClient) {}
  
  ngOnInit(): void {

    // GET request for the posts
    this.http.get<any>('http://localhost:3000/api/posts')

    .subscribe(response => {

      // (debugging) log entire array
      console.log(response.posts); 

      this.posts = response.posts;

      // (debugging) iterate through array
      response.posts.forEach((post: any) => {
        console.log(post.postDescription);
      })
    }, error => {
      console.error('Error retrieving posts.');
    })

    // // temp: clear posts for testing
    // const id = 1;
    // this.http.delete(`http://localhost:3000/api/auth/deletepost/${id}`, {
    //   withCredentials: true
    // })
    //   .subscribe(response => {
    //     console.log(response);
    //   })
  }
}