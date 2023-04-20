import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css'],
})
export class WelcomeComponent {
  posts: any[] = [];
  LocationData: any = {}
  constructor(private http: HttpClient, private _auth: AuthService,
    private _router: Router) {}

  sortByLikes(posts: any[]) {
    return posts.sort((a, b) => b.numberOfLikes - a.numberOfLikes);
  }
  
  ngOnInit(): void {

    // GET request for the posts
    this.http.get<any>('http://localhost:3000/api/posts')

    .subscribe(response => {

      this.posts = this.sortByLikes(response.posts);

    }, error => {
      console.error('Error retrieving posts.');
    })
  }
  
  onGo(location: string) {
    this.http.get<any>(`http://localhost:3000/api/posts/location/${location}`, {
      reportProgress: true,
      withCredentials: true 
    })
      .subscribe(response => {
        this.posts = this.sortByLikes(response.posts);
      }, error => {
        console.error('Getting Account Info');
      })
  }
  
}

