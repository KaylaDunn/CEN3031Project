import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
import { CommentDialogComponent } from '../comment-dialog/comment-dialog.component';
import { ViewDialogComponent } from '../view-dialog/view-dialog.component';

@Component({
  selector: 'app-location-search',
  templateUrl: './location-search.component.html',
  styleUrls: ['./location-search.component.css']
})
export class LocationSearchComponent {
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
