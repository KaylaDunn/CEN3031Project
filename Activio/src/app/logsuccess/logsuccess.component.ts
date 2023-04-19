import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
import { CommentDialogComponent } from '../comment-dialog/comment-dialog.component';
import { ViewDialogComponent } from '../view-dialog/view-dialog.component';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-logsuccess',
  templateUrl: './logsuccess.component.html',
  styleUrls: ['./logsuccess.component.css'],
})

export class LogsuccessComponent {

  posts: any[] = [];
  LocationData: any = {}
  constructor(private http: HttpClient, private dialog: MatDialog, private _auth: AuthService,
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
   
    // // temp: clear posts for testing
    // const id = 1;
    // this.http.delete(`http://localhost:3000/api/auth/deletepost/${id}`, {
    //   withCredentials: true
    // })
    //   .subscribe(response => {
    //     console.log(response);
    //   })
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
  onComment(id: number) {

    const dialogRef: MatDialogRef<CommentDialogComponent> = this.dialog.open(CommentDialogComponent);

    dialogRef.afterClosed().subscribe((comment) => {
      if (comment) {
        console.log("Comment: ", comment);
        this.http.put(`http://localhost:3000/api/auth/comment/${id}`,
        {
          comment
        },
        {
          withCredentials: true
        })
        .subscribe(response => {
          console.log(response);
          window.location.reload();
        }, (error) => {
          console.log(error);
        });
      }
    })
  }

  // called when user hits view
  onView(id: number) {

    const dialogConfig = new MatDialogConfig();
    dialogConfig.data = { id: id };
    const dialogRef: MatDialogRef<ViewDialogComponent> = this.dialog.open(ViewDialogComponent, dialogConfig);
  }
}
