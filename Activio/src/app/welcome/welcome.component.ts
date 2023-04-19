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
  
  ngOnInit(): void {

    // GET request for the posts
    this.http.get<any>('http://localhost:3000/api/posts')

    .subscribe(response => {

      this.posts = response.posts;

    }, error => {
      console.error('Error retrieving posts.');
    })
  }
  Location() {
    /*this._auth.Location(this.LocationData)
      .subscribe(
        res => {
          console.log(this.LocationData.input);
          this._router.navigate(['/location-search']);
        },
        err => console.log(err)
      )
      */
  }
}

