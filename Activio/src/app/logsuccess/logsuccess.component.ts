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

      this.posts = response.posts;

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

  // called when user likes a post
  onLike(id: number) {
    this.http.put(`http://localhost:3000/api/auth/likepost/${id}`, null, {
      withCredentials: true
    })
    .subscribe(response => {
      console.log(response);
      window.location.reload();
    }, error => {
      console.error(error);
    }
    )
  }

  // called when user hits comment
  onComment() {
    
  }
}