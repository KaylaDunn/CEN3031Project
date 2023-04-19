import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { MatDialog} from '@angular/material/dialog';


@Component({
  selector: 'app-location-search-from-main',
  templateUrl: './location-search-from-main.component.html',
  styleUrls: ['./location-search-from-main.component.css']
})
export class LocationSearchFromMainComponent {
  posts: any[] = [];

  constructor(private http: HttpClient, private dialog: MatDialog) { }

  ngOnInit(): void {

    // GET request for the posts
    this.http.get<any>('http://localhost:3000/api/posts')

      .subscribe(response => {

        this.posts = response.posts;

      }, error => {
        console.error('Error retrieving posts.');
      })
  }
}
